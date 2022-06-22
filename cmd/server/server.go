package main

import (
	"airdrop-bot/aws"
	"airdrop-bot/binance"
	"airdrop-bot/cfg"
	"airdrop-bot/cloudflare"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/services"
	"airdrop-bot/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"
)

func NewServer(cfg *cfg.ServerConfig) (*Server, error) {
	cf, err := cloudflare.NewClient(cfg.Cloudflare)
	if err != nil {
		return nil, err
	}

	bn := binance.New(&cfg.Binance)

	return &Server{
		cfg: cfg,
		cf:  cf,
		bn:  bn,
	}, nil
}

type Server struct {
	cfg *cfg.ServerConfig
	cf  *cloudflare.Client
	bn  *binance.Binance
	r   *gin.Engine
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
		var a db.Account
		db.DB.First(&a, step.AccountID)
		rsp = cfg.HeartbeatResp{
			Account: a.Mnemonic,
			Task:    db.StepArbitrumBridge,
			TrackID: step.ID,
		}
		log.Infof("bridge task of account %v has activated", a.ID)
		step.Status = db.StatusRunning
		db.DB.Save(&step)
	} else {
		log.Infof("no pending tasks")
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
	var step db.StepRun
	db.DB.First(&step, ts.TrackID)
	if ts.Result == cfg.ResultSuccess {
		step.Status = db.StatusSuccess
	} else if ts.Result == cfg.ResultFail {
		step.Status = db.StatusFailed
		step.Reason = ts.Reason
	}
	db.DB.Save(&step)
	log.Infof("success update task status: %+v", step)

	c.String(http.StatusOK, "success")
}

func (s *Server) bridgeAccounts(c *gin.Context) {
	if err := s.doBridges(); err != nil {
		log.Errorf("do bridges: %v", err)
		c.String(http.StatusBadRequest, "")
	}
	log.Infof("---- ALL BRIDGES SUCCESS ----")
}

func (s *Server) doBridges() error {
	for i := 0; i < s.cfg.AccountsToGen; i++ {
		account := db.FindAccount(i + 1)
		if account.ID == 0 {
			log.Errorf("accounts with id %d not found in db, skip it", i+1)
			continue
		}
		if err := s.bridgeOne(account, 0); err != nil {
			log.Errorf("bridge account %d error: %v", account.ID, err)
			return err
		}
	}
	return nil
}

func (s *Server) bridgeOne(a db.Account, retry int) error {
	if db.StepBeenDone(a.ID, db.StepArbitrumBridge) {
		log.Infof("account %d bridge has already been done, skip...", a.ID)
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

	ip, err := s.attachOrAllocateAccountIp(a)
	if err != nil {
		return errors.Wrap(err, "attachOrAllocateAccountIp")
	}
	step := db.StepRun{
		Service:    db.ArbitrumService,
		Step:       db.StepArbitrumBridge,
		Status:     db.StatusPending,
		Reason:     "",
		AccountID:  a.ID,
		NodeID:     ip.NodeID,
		StaticIpID: ip.ID,
	}
	db.DB.Save(&step)

loop:
	for {
		time.Sleep(10 * time.Second)
		db.DB.First(&step, step.ID)
		switch step.Status {
		case db.StatusRunning:
			log.Infof("bridge task for account %s begin to run", a.ID)
		case db.StatusSuccess:
			log.Infof("bridge task for account %s succeed", a.ID)
			break loop
		case db.StatusFailed: //retry
			log.Errorf("bridge task for account %s failed, reason: %v", a.ID, step.Reason)
			if retry == 3 {
				return fmt.Errorf("maxium retry exceed for account: %v", a.ID)
			}
			retry++
			return s.bridgeOne(a, retry)
		}
	}

	return nil
}

func (s *Server) attachOrAllocateAccountIp(a db.Account) (*db.StaticIp, error) {

	var allocatedIP *db.StaticIp
	ip, ok := db.AccountHasIp(db.ArbitrumService, a.ID)
	if !ok {
		/*
			账户没有关联的ip，从节点中随机选取一个节点。如果节点ip关联的账号没有达到最大限制，则使用此ip，反之创建一个新的ip
		*/
		//account has no associated ip, first check current ip, if reach limit then use new ip
		node := db.PickANode()
		lightsail, err := aws.CreateLightsailClient(node.NodeName, node.Region, path.Join(s.cfg.Dir, "aws.config"))
		if err != nil {
			return nil, errors.Wrap(err, "create lightsail client")
		}
		attachedIpName, _, err := lightsail.GetAttachedIp()
		if err != nil {
			return nil, errors.Wrap(err, "get attached static ip")
		}

		ip2 := db.GetOrAddIp(attachedIpName)
		count := db.IpRelatedAccountNum(db.ArbitrumService, ip2.ID)
		if count >= s.cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip

			newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
			if err := lightsail.CreateIp(newIpName); err != nil {
				return nil, errors.Wrap(err, "create ip")
			}

			if err := lightsail.DetachIp(attachedIpName); err != nil {
				return nil, errors.Wrap(err, "lightsail detach ip")
			}
			newIp := db.GetOrAddIp(newIpName) //save new ip to db
			newIp.NodeID = node.ID
			db.DB.Save(&newIp) //更新node id

			rel := db.StaticIpAccountRelation{
				StaticIpID: newIp.ID,
				AccountID:  a.ID,
				Service:    db.ArbitrumService,
			}
			db.SaveAccountIpRelation(&rel)
			if err := lightsail.AttachIp(newIpName); err != nil {
				return nil, errors.Wrap(err, "attch ip")
			}
			allocatedIP = &newIp
		} else { //使用当前attachedIp
			rel := db.StaticIpAccountRelation{
				StaticIpID: ip2.ID,
				AccountID:  a.ID,
				Service:    db.ArbitrumService,
			}
			db.SaveAccountIpRelation(&rel)

			allocatedIP = &ip2
		}
		if _, address, err := lightsail.GetAttachedIp(); err == nil {
			err := s.cf.UpdateDnsRecord(node.DnsName, address)
			if err != nil {
				log.Infof("update cloudflare dns record error: %v", err)
			} else {
				log.Infof("update cloudflare dns record success")
			}
		}

	} else { //账户存在指定ip，如果当前ip是指定ip，就是用当前ip，如果不是切换ip

		allocatedIP = ip
		var node db.Node
		db.DB.First(&node, ip.NodeID)
		lightsail, err := aws.CreateLightsailClient(node.NodeName, node.Region, path.Join(s.cfg.Dir, "aws.config"))
		if err != nil {
			return nil, errors.Wrap(err, "create lightsail client")
		}
		attachedIpName, _, err := lightsail.GetAttachedIp()
		if err != nil {
			return nil, errors.Wrap(err, "get attached static ip")
		}

		if ip.Name != attachedIpName {
			//attached ip is not the desired ip
			err := lightsail.DetachIp(attachedIpName)
			if err != nil {
				return nil, errors.Wrap(err, "detach ip")
			}
			err = lightsail.AttachIp(ip.Name)
			if err != nil {
				return nil, errors.Wrap(err, "attach ip")
			}
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

	}
	return allocatedIP, nil
}

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
		a := db.Account{
			Mnemonic:   mnemonic,
			Address:    address,
			PrivateKey: priKey,
			Services:   nil,
		}
		log.Infof("generate eth account: %s", address)
		db.SaveAccount(&a)
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

func (s *Server) TransferEthToAccounts() error {
	balance, err := s.bn.EthBalance()
	if err != nil {
		return errors.Wrap(err, "get balance")
	}
	bb, err := strconv.ParseFloat(balance, 64)
	if err != nil {
		return errors.Wrap(err, "parse float")
	}

	var accounts []db.Account
	db.DB.Find(&accounts)
	perAccount := bb / float64(len(accounts))

	p := perAccount
	for p < 1 {
		p *= 10
	}
	p = p / perAccount * 10

	for _, a := range accounts {
		r := float64(rand.Intn(10)-5)/p + perAccount - 1/(p*10)

		err := s.bn.WithdrawEth(a.Address, r)
		if err != nil {
			log.Errorf("transfer to account %v error: %v", a.Address, err)
			continue
		}
	}
	return nil
}