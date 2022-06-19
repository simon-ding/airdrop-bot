package server

import (
	"airdrop-bot/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/cloudflare"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/services"
	"airdrop-bot/utils"
	"context"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"math/rand"
	"path"
	"strconv"
	"strings"
	"time"
)

func NewServer(cfg *cfg.Config) (*Server, error) {
	cf, err := cloudflare.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	lightsail, err := aws.CreateLightsailClient(cfg.AWS.InstanceName, cfg.AWS.Region, path.Join(cfg.Dir, "aws.config"))
	if err != nil {
		return nil, errors.Wrap(err, "create lightsail client")
	}

	return &Server{
		cfg:       cfg,
		cf:        cf,
		lightsail: lightsail,
	}, nil
}

type Action func(*metamask.Metamask, func()) error

type Server struct {
	cfg       *cfg.Config
	cf        *cloudflare.Client
	lightsail *aws.Client
}

type Para struct {
	Accounts string `json:"accounts"`
	Step     string `json:"step"`
}

func (s *Server) BridgeAccounts(accounts string) error {

	if strings.Contains(accounts, "-") {
		pp := strings.Split(accounts, "-")
		if len(pp) < 2 {
			return errors.Errorf("format error")
		}
		startS := pp[0]
		endS := pp[1]
		startInt, err := strconv.Atoi(startS)
		if err != nil {
			return errors.Errorf("format error")
		}
		endInt, err := strconv.Atoi(endS)
		if err != nil {
			return errors.Errorf("format error")
		}
		for i := startInt; i < endInt+1; i++ {
			account := db.FindAccount(i)
			if account.ID == 0 {
				log.Errorf("accounts with id %d not found in db, skip it", i+1)
				continue
			}
			if err := s.bridgeOne(account); err != nil {
				log.Errorf("bridge account %d error: %v", account.ID, err)
				return err
			}
		}

	}
	if accounts == "" {
		// all accounts
		for i := 0; i < s.cfg.AccountsToGen; i++ {
			account := db.FindAccount(i + 1)
			if account.ID == 0 {
				log.Errorf("accounts with id %d not found in db, skip it", i+1)
				continue
			}
			if err := s.bridgeOne(account); err != nil {
				log.Errorf("bridge account %d error: %v", account.ID, err)
				return err
			}
		}

	}
	return nil
}

func (s *Server) bridgeOne(a db.Account) error {

	//if db.HasArbitrumStepRun(a.ID, db.StepArbitrumDeposit) {
	//	log.Infof("account %d bridge has already been done,, skip", a.ID)
	//	return nil
	//}
	err := s.attachOrAllocateAccountIp(a)
	if err != nil {
		return errors.Wrap(err, "attachOrAllocateAccountIp")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("load-extension", path.Join(s.cfg.Dir, "ext")),
		chromedp.Flag("window-size", "1920,1080"),

		//chromedp.UserDataDir(dataDir),
	)

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
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Debugf), chromedp.WithLogf(log.Infof),
		chromedp.WithErrorf(log.Errorf))
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	meta := metamask.NewMetamask(allocCtx)
	err = meta.FirstOpenAndImportAccount(s.cfg.WalletPassword, a)
	if err != nil {
		return err
	}
	balance, err := meta.Balance()
	log.Infof("account %s current balance is %v", a.Address, balance)

	err = meta.OpenChanListAndAddNetwork("Optimism")
	if err != nil {
		return errors.Wrap(err, "add op")
	}

	hop := services.MewHop(meta.Context(), meta, "ETH", "optimism", "arbitrum")
	err = hop.LinkMetaMask()
	if err != nil {
		return err
	}
	log.Infof("hop link metamask success")

	err = hop.BridgeMoney(balance - cfg.BridgeReverse)
	if err != nil {
		return errors.Wrap(err, "arb deposit")
	}
	db.SaveArbitrumStep(a.ID, db.StepArbitrumDeposit)

	balance = meta.WaitForBalanceChange(balance)
	return nil

}

func (s *Server) attachOrAllocateAccountIp(a db.Account) error {
	attachedIpName, _, err := s.lightsail.GetAttachedIp()
	if err != nil {
		return errors.Wrap(err, "get attached static ip")
	}

	ip, ok := db.AccountHasIp(db.ArbitrumService, a.ID)
	if !ok {
		//account has no associated ip, first check current ip, if reach limit then use new ip

		ip2 := db.GetOrAddIp(attachedIpName)
		count := db.IpRelatedAccountNum(db.ArbitrumService, ip2.ID)
		if count >= s.cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip

			newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
			if err := s.lightsail.CreateIp(newIpName); err != nil {
				return errors.Wrap(err, "create ip")
			}

			if err := s.lightsail.DetachIp(attachedIpName); err != nil {
				return errors.Wrap(err, "lightsail detach ip")
			}
			newIp := db.GetOrAddIp(newIpName) //save new ip to db
			rel := db.StaticIpAccountRelation{
				StaticIpID: newIp.ID,
				AccountID:  a.ID,
				Service:    db.ArbitrumService,
			}
			db.SaveAccountIpRelation(&rel)
			if err := s.lightsail.AttachIp(newIpName); err != nil {
				return errors.Wrap(err, "attch ip")
			}
		} else { //使用当前attachedIp
			rel := db.StaticIpAccountRelation{
				StaticIpID: ip2.ID,
				AccountID:  a.ID,
				Service:    db.ArbitrumService,
			}
			db.SaveAccountIpRelation(&rel)

		}

	} else { //账户存在指定ip，如果当前ip是指定ip，就是用当前ip，如果不是切换ip
		if ip.Name != attachedIpName {
			//attached ip is not the desired ip
			err := s.lightsail.DetachIp(attachedIpName)
			if err != nil {
				return errors.Wrap(err, "detach ip")
			}
			err = s.lightsail.AttachIp(ip.Name)
			if err != nil {
				return errors.Wrap(err, "attach ip")
			}
		}

	}
	if _, address, err := s.lightsail.GetAttachedIp(); err == nil {
		err := s.cf.UpdateSSRecord(address)
		if err != nil {
			log.Infof("update cloudflare dns record error: %v", err)
		} else {
			log.Infof("update cloudflare dns record success")
		}
	}
	return nil

}

func (s *Server) Do() error {
	s.FirstRunGenAccount()

	for i := 0; i < s.cfg.AccountsToGen; i++ {
		account := db.FindAccount(i + 1)
		if account.ID == 0 {
			log.Errorf("accounts with id %d not found in db, skip it", i+1)
			continue
		}

		attachedIpName, _, err := s.lightsail.GetAttachedIp()
		if err != nil {
			return errors.Wrap(err, "get attached static ip")
		}

		ip, ok := db.AccountHasIp(db.ArbitrumService, account.ID)
		if !ok {
			//account has no associated ip, first check current ip, if reach limit then use new ip

			ip2 := db.GetOrAddIp(attachedIpName)
			count := db.IpRelatedAccountNum(db.ArbitrumService, ip2.ID)
			if count >= s.cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip

				newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
				if err := s.lightsail.CreateIp(newIpName); err != nil {
					return errors.Wrap(err, "create ip")
				}

				if err := s.lightsail.DetachIp(attachedIpName); err != nil {
					return errors.Wrap(err, "lightsail detach ip")
				}
				newIp := db.GetOrAddIp(newIpName) //save new ip to db
				rel := db.StaticIpAccountRelation{
					StaticIpID: newIp.ID,
					AccountID:  account.ID,
					Service:    db.ArbitrumService,
				}
				db.SaveAccountIpRelation(&rel)
				if err := s.lightsail.AttachIp(newIpName); err != nil {
					return errors.Wrap(err, "attch ip")
				}
			} else { //使用当前attachedIp
				rel := db.StaticIpAccountRelation{
					StaticIpID: ip2.ID,
					AccountID:  account.ID,
					Service:    db.ArbitrumService,
				}
				db.SaveAccountIpRelation(&rel)

			}

		} else { //账户存在指定ip，如果当前ip是指定ip，就是用当前ip，如果不是切换ip
			if ip.Name != attachedIpName {
				//attached ip is not the desired ip
				err := s.lightsail.DetachIp(attachedIpName)
				if err != nil {
					return errors.Wrap(err, "detach ip")
				}
				err = s.lightsail.AttachIp(ip.Name)
				if err != nil {
					return errors.Wrap(err, "attach ip")
				}
			}

		}

		err = s.DoAccountAllStep(account)
		if err != nil {
			log.Errorf("do error: %v", err)
		}

	}
	return nil
}

func (s *Server) DoAccountAllStep(account db.Account) error {
	log.Infof("---- [START] begin doing transactions of account %d ----", account.ID)
	if db.HasArbitrumStepAllRun(account.ID) {
		log.Infof("arbitrum steps have all been run, skip account...")
		return nil
	}
	if s.cfg.RunSingleStep {
		log.Infof("single step mode enabled, will only run one step")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("load-extension", path.Join(s.cfg.Dir, "ext")),
		chromedp.Flag("window-size", "1920,1080"),

		//chromedp.UserDataDir(dataDir),
	)

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

	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Debugf), chromedp.WithLogf(log.Infof),
		chromedp.WithErrorf(log.Errorf))
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	meta := metamask.NewMetamask(allocCtx)
	err := meta.FirstOpenAndImportAccount(s.cfg.WalletPassword, account)
	if err != nil {
		return err
	}
	balance, err := meta.Balance()
	log.Infof("account %s current balance is %v", account.Address, balance)

	arb := services.NewArbitrum(meta.Context(), meta)
	err = arb.LinkMetaMask()
	if err != nil {
		return err
	}
	const depositReverse = 0.005

	if !db.HasArbitrumStepRun(account.ID, db.StepArbitrumDeposit) {
		err = arb.Deposit(balance - depositReverse)
		if err != nil {
			return errors.Wrap(err, "arb deposit")
		}
		db.SaveArbitrumStep(account.ID, db.StepArbitrumDeposit)

		balance = meta.WaitForBalanceChange(balance)
		if s.cfg.RunSingleStep {
			log.Infof("run single step mode, exiting")
			return nil
		}
	}

	if err := arb.AddL2NetworkAndWaitTransaction(); err != nil {
		return errors.Wrap(err, "add l2 network")
	}

	actions := map[int]Action{
		0: s.aaveSupplyAndBorrowMoney,
		1: s.gmxSwapCoin,
	}

	allNil := func(m map[int]Action) bool {
		for _, v := range m {
			if v != nil {
				return false
			}
		}
		return true
	}

	for !allNil(actions) {
		n := rand.Intn(len(actions))
		act := actions[n]
		if act == nil {
			continue
		}

		actions[n] = nil
		var callback func()
		if n == 0 {
			if db.HasArbitrumStepRun(account.ID, db.StepAaveSupplyAndBorrow) {
				log.Infof("arbitrum supply and borrow has run, skip it")
				continue
			}
			callback = func() {
				db.SaveArbitrumStep(account.ID, db.StepAaveSupplyAndBorrow)
			}
		} else if n == 1 {
			if db.HasArbitrumStepRun(account.ID, db.StepGmxSwap) {
				log.Infof("arbitrum gmx swap has run, skip it")
				continue
			}
			callback = func() {
				db.SaveArbitrumStep(account.ID, db.StepGmxSwap)
			}
		}
		if err := act(meta, callback); err != nil {
			return err
		}
		balance = meta.WaitForBalanceChange(balance)
		if s.cfg.RunSingleStep {
			log.Infof("run single step mode, exiting")
			return nil
		}
	}
	log.Infof("all steps have done")

	return nil
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
