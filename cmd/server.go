package main

import (
	"airdrop-bot/pkg/aws"
	"airdrop-bot/pkg/binance"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/ent"
	"airdrop-bot/ent/steprun"
	"airdrop-bot/log"
	"airdrop-bot/pkg/metamask"
	"airdrop-bot/utils"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"path"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func NewServer(cfg *cfg.ServerConfig) (*Server, error) {
	cf, err := cloudflare.NewClient(cfg.Cloudflare)
	if err != nil {
		return nil, err
	}

	bn := binance.New(&cfg.Binance)
	r := gin.New()
	return &Server{
		cfg: cfg,
		cf:  cf,
		bn:  bn,
		r:   r,
	}, nil
}

type Server struct {
	cfg  *cfg.ServerConfig
	cf   *cloudflare.Client
	bn   *binance.Binance
	r    *gin.Engine
	lock bool
}

func (s *Server) Serve() error {
	s.r.Use(func(c *gin.Context) {
		token := c.GetHeader(cfg.AuthHeader)
		if token != s.cfg.Token {
			//token is invalid
			log.Infof("access use invalid token: %v, ip %s", token, c.ClientIP())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})
	s.r.POST(cfg.ApiHeartbeatUrl, s.heartBeat)
	s.r.POST(cfg.ApiTaskStatusUpdateUrl, s.taskStatusUpdate)
	s.r.POST(cfg.ApiDoBridges, s.bridgeAccounts)
	return s.r.Run(":8080")
}

func (s *Server) heartBeat(c *gin.Context) {
	var hb cfg.Heartbeat

	if err := c.ShouldBindJSON(&hb); err != nil {
		log.Errorf("bind body error: %v", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	node := db.AddOrUpdateNode(&hb)

	var rsp cfg.HeartbeatResp
	step := db.FindFirstPendingTask(node.ID)

	if step != nil {
		a, err := db.Client.Account.Get(context.TODO(), step.Edges.Account.ID)
		if err != nil {
			log.Errorf("query account error: %v", err)
		}

		rsp = cfg.HeartbeatResp{
			Account: a.Mnemonic,
			Task:    step.Step,
			TrackID: step.ID,
			Timeout: time.Minute * 15,
		}
		log.Infof("bridge task of account %v has activated", a.ID)
		db.Client.StepRun.Update().Where(steprun.ID(step.ID)).SetStatus(steprun.StatusRunning).SaveX(context.TODO())
	}
	c.JSON(http.StatusOK, rsp)
}

func (s *Server) taskStatusUpdate(c *gin.Context) {
	var ts cfg.TaskStatus
	if err := c.ShouldBindJSON(&ts); err != nil {
		log.Errorf("bind body error: %v", err)
		c.String(http.StatusBadRequest, "")
		return
	}
	var ss steprun.Status
	switch ts.Result {
	case cfg.ResultFail:
		ss = steprun.StatusFailed
	case cfg.ResultSuccess:
		ss = steprun.StatusSuccess
	}
	db.Client.StepRun.Update().Where(steprun.ID(ts.TrackID)).SetStatus(ss).SetReason(ts.Reason).ExecX(context.Background())

	log.Infof("success update task status: %+v", ts)

	c.String(http.StatusOK, "success")
}

type DoBridgeParam struct {
	From     int  `json:"from"`
	To       int  `json:"to"`
	Transfer bool `json:"transfer"`
}

func (s *Server) bridgeAccounts(c *gin.Context) {
	if s.lock {
		c.String(http.StatusBadRequest, "running")
		return
	}

	var b DoBridgeParam
	if err := c.ShouldBindJSON(&b); err != nil {
		log.Errorf("json error: %v", err)
		c.String(http.StatusBadRequest, "json error")
		return
	}

	s.lock = true
	go func() {
		defer func() {
			s.lock = false
		}()

		log.Infof("---------- BRIDGE START ------------")
		if err := s.doBridges(b); err != nil {
			log.Errorf("do bridges: %v", err)
			c.String(http.StatusBadRequest, "")
		}
		log.Infof("---- ALL BRIDGES SUCCESS ----")
	}()
	c.String(http.StatusOK, "")
}

func (s *Server) doBridges(p DoBridgeParam) error {
	for i := p.From; i <= p.To; i++ {
		account := db.FindAccount(i)
		if account.ID == 0 {
			log.Errorf("accounts with id %d not found in db, skip it", i)
			continue
		}
		if p.Transfer {
			if err := s.transferEth(account); err != nil {
				log.Errorf("withdraw from binance error: %v", err)
			}
		}
		log.Infof("begin bridge account: %v", account.Address)

		if err := s.bridgeOne(account); err != nil {
			log.Errorf("bridge account %d error: %v", account.ID, err)
			return err
		}

	}
	return nil
}

func (s *Server) transferEth(a *ent.Account) error {
	if db.StepBeenDone(a.ID, db.StepArbitrumTransfer) {
		log.Infof("account %d step %s has already been done, skip...", a.ID, db.StepArbitrumTransfer)
		return nil
	}

	amount := 0.015 + rand.Float64()/100
	log.Infof("transfer enabled, transfer %v eth to address %v", amount, a.Address)

	err := s.bn.WithdrawEth("ARBITRUM", a.Address, amount)
	if err != nil {
		return err
	}
	return db.Client.StepRun.Create().
		SetAccount(a).SetStep(db.StepArbitrumTransfer).SetStatus(steprun.StatusSuccess).Exec(context.Background())
}

func (s *Server) pickNodeToRun() (*ent.Node, error) {
	node := db.PickANodeRandom()
	log.Infof("node %v selected for the task: %+v", node.Name, node)
	lightsail, err := aws.CreateLightsailClient(node.Name, node.Region, path.Join(s.cfg.Dir, "aws.config"))
	if err != nil {
		return nil, errors.Wrap(err, "create lightsail client")
	}
	if err := lightsail.AttachNewIP(); err != nil {
		return nil, err
	}

	//更新cloudflare dns
	if _, address, err := lightsail.GetAttachedIp(); err == nil {
		err := s.cf.UpdateDnsRecord(node.DnsName, address)
		if err != nil {
			log.Infof("update cloudflare dns record error: %v", err)
		} else {
			log.Infof("update cloudflare dns record success")
		}
	}

	return node, nil
}

func (s *Server) bridgeOne(a *ent.Account) error {
	node, err := s.pickNodeToRun()
	if err != nil {
		return err
	}

	if err := s.doOneStep(a, node, db.StepArbitrumBridge, 0); err != nil {
		return err
	}

	if err := s.doOneStep(a, node, db.StepArbitrumBridge2, 0); err != nil {
		return err
	}

	return nil
}

func (s *Server) doOneStep(a *ent.Account, node *ent.Node, sp string, retry int) error {

	if db.StepBeenDone(a.ID, sp) {
		log.Infof("account %d step %s has already been done, skip...", a.ID, sp)
		return nil
	}

	if !s.cfg.Owlracle.Disable {
		for {
			if s.isGasFeeAcceptable() {
				log.Infof("gas fee is acceptable, continue")
				break
			}
			log.Infof("will try again in 5min")
			time.Sleep(5 * time.Minute)
		}
	}

	step := db.Client.StepRun.Create().
		SetAccount(a).SetStep(sp).SetStatus(steprun.StatusPending).SetNode(node).SaveX(context.Background())
	log.Infof("add new task to db, account %d, step %+v", a.ID, step)

	i := 0
loop1:
	for {
		time.Sleep(10 * time.Second)
		step = db.Client.StepRun.GetX(context.Background(), step.ID)
		switch step.Status {
		case steprun.StatusRunning:
			if i == 0 {
				log.Infof("bridge task for account %v begin to run", a.ID)
			}
			i++
		case steprun.StatusSuccess:
			log.Infof("bridge task for account %v succeed", a.ID)
			break loop1
		case steprun.StatusFailed: //retry
			log.Errorf("bridge task for account %v failed, reason: %v", a.ID, step.Reason)
			if retry == 3 {
				return fmt.Errorf("maxium retry exceed for account: %v", a.ID)
			}
			retry++
			return s.doOneStep(a, node, sp, retry)
		}
	}

	t := time.Minute * time.Duration(rand.Intn(10))
	log.Infof("--- sleep for %v ---", t)
	time.Sleep(t)

	return nil
}

// func (s *Server) attachOrAllocateAccountIp(a db.Account) (*db.StaticIp, error) {

// 	var allocatedIP *db.StaticIp
// 	ip, ok := db.AccountHasIp(a.StaticIpID)
// 	if !ok {
// 		log.Infof("account has no associated ip: %v", a.ID)
// 		/*
// 			账户没有关联的ip，从节点中随机选取一个节点。如果节点ip关联的账号没有达到最大限制，则使用此ip，反之创建一个新的ip
// 		*/

// 		//等待分配ip
// 		node := db.PickANodeRandom()
// 		log.Infof("node %v selected for the task: %+v", node.NodeName, node)
// 		lightsail, err := aws.CreateLightsailClient(node.NodeName, node.Region, path.Join(s.cfg.Dir, "aws.config"))
// 		if err != nil {
// 			return nil, errors.Wrap(err, "create lightsail client")
// 		}
// 		attachedIpName, _, err := lightsail.GetAttachedIp()
// 		if err != nil {
// 			return nil, errors.Wrap(err, "get attached static ip")
// 		}

// 		ip2 := db.GetOrAddIp(attachedIpName, node.ID)
// 		count := db.IpRelatedAccountNum(ip2.ID)
// 		if count >= s.cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip
// 			log.Infof("attached ip has reached upper accounts limit: %v", ip2.Name)
// 			newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
// 			if err := lightsail.CreateIp(newIpName); err != nil {
// 				return nil, errors.Wrap(err, "create ip")
// 			}

// 			if err := lightsail.DetachIp(attachedIpName); err != nil {
// 				return nil, errors.Wrap(err, "lightsail detach ip")
// 			}
// 			newIp := db.GetOrAddIp(newIpName, node.ID) //save new ip to db
// 			newIp.NodeID = node.ID
// 			db.DB.Save(&newIp) //更新node id

// 			a.StaticIpID = newIp.ID
// 			db.DB.Save(&a)

// 			if err := lightsail.AttachIp(newIpName); err != nil {
// 				return nil, errors.Wrap(err, "attch ip")
// 			}
// 			allocatedIP = &newIp
// 		} else { //使用当前attachedIp
// 			log.Infof("num of accounts associated with current ip %v, use attached ip", count)

// 			a.StaticIpID = ip2.ID
// 			db.DB.Save(&a)

// 			allocatedIP = &ip2
// 		}
// 		if _, address, err := lightsail.GetAttachedIp(); err == nil {
// 			err := s.cf.UpdateDnsRecord(node.DnsName, address)
// 			if err != nil {
// 				log.Infof("update cloudflare dns record error: %v", err)
// 			} else {
// 				log.Infof("update cloudflare dns record success")
// 			}
// 		}

// 	} else { //账户存在指定ip，如果当前ip是指定ip，就是用当前ip，如果不是切换ip
// 		log.Infof("accounts has associated ip, use this ip: %+v", ip)

// 		allocatedIP = ip
// 		var node db.Node
// 		db.DB.First(&node, ip.NodeID)
// 		lightsail, err := aws.CreateLightsailClient(node.NodeName, node.Region, path.Join(s.cfg.Dir, "aws.config"))
// 		if err != nil {
// 			return nil, errors.Wrap(err, "create lightsail client")
// 		}
// 		attachedIpName, _, err := lightsail.GetAttachedIp()
// 		if err != nil {
// 			return nil, errors.Wrap(err, "get attached static ip")
// 		}

// 		if ip.Name != attachedIpName {
// 			//attached ip is not the desired ip
// 			log.Infof("attached ip is not the desired ip: %v", attachedIpName)
// 			err := lightsail.DetachIp(attachedIpName)
// 			if err != nil {
// 				return nil, errors.Wrap(err, "detach ip")
// 			}
// 			err = lightsail.AttachIp(ip.Name)
// 			if err != nil {
// 				return nil, errors.Wrap(err, "attach ip")
// 			}
// 		}
// 		//更新cloudflare dns
// 		if _, address, err := lightsail.GetAttachedIp(); err == nil {
// 			err := s.cf.UpdateDnsRecord(node.DnsName, address)
// 			if err != nil {
// 				log.Infof("update cloudflare dns record error: %v", err)
// 			} else {
// 				log.Infof("update cloudflare dns record success")
// 			}
// 		}

// 	}
// 	return allocatedIP, nil
// }

func (s *Server) FirstRunGenAccount() {
	num := s.cfg.AccountsToGen
	accountNum := db.AccountNum()
	if accountNum == num {
		log.Infof("accounts have already been generated,skip")
		return
	}
	log.Infof("try to generate %d accounts", num-accountNum)
	for i := 0; i < num-accountNum; i++ {
		mnemonic, address, priKey := utils.NewEthAccount()
		log.Infof("generate eth account: %s", address)
		db.SaveAccount(mnemonic, address, priKey)
	}
}

func (s *Server) isGasFeeAcceptable() bool {
	fee, err := utils.GetGasFee(&s.cfg.Owlracle)
	if err != nil {
		log.Errorf("get gas fee error: %v", err)
		return false
	}
	if len(fee.Speeds) == 0 {
		log.Infof("api return no fee: %+v", fee)
		return false
	}
	log.Infof("current gas fee is %+v", fee.Speeds)

	if fee.Speeds[1].EstimatedFee > 7 {
		log.Errorf("gas fee too high: %v", fee.Speeds[1].EstimatedFee)
		return false
	}
	return true
}

func (s *Server) aaveSupplyAndBorrowMoney(meta *metamask.Metamask, callback func()) error {
	log.Infof("run aave supply and borrow money")
	aave := services.NewAave(meta.Context(), meta)
	if err := aave.OpenAndLinkMetaMask(); err != nil {
		return errors.Wrap(err, "aave link metamask")
	}
	const aaveSupply = 0.01
	if err := aave.SupplyEth(aaveSupply); err != nil {
		return errors.Wrap(err, "supply aave")
	}
	n := rand.Intn(3)
	var aaveBorrowUsdt = 5 + n

	if err := aave.BorrowMoney("USDT", float64(aaveBorrowUsdt)); err != nil {
		return errors.Wrap(err, "borrow money")
	}
	callback()
	return nil
}

func (s *Server) gmxSwapCoin(meta *metamask.Metamask, callback func()) error {
	log.Infof("run gmx swap coin")
	gmx := services.NewGmx(meta.Context(), meta)
	if err := gmx.OpenAndLinkMeta(); err != nil {
		return errors.Wrap(err, "gmx link metamask")
	}
	n2 := rand.Intn(5)
	var gmxSwap = float64(10+n2) / 1000.
	if err := gmx.SwapCoin("ETH", "USDT", gmxSwap); err != nil {
		return errors.Wrap(err, "gmx swap coin")
	}
	callback()
	return nil
}
