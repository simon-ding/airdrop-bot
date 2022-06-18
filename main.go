package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	server2 "airdrop-bot/server"
	"airdrop-bot/services"
	"airdrop-bot/utils"
	"context"
	"flag"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"time"
)

var accounts = flag.String("accounts", "", "accounts: 1-30  1")

func main() {

	flag.Parse()
	cfg, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")
	awsDir := path.Join(cfg.Dir, "aws.config")
	err = aws.WriteAwsConfig(cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, awsDir)
	if err != nil {
		log.Errorf("write aws config error: %v", err)
		return
	}

	err = utils.Unzip(asserts.Ext, path.Join(cfg.Dir, "ext"))
	if err != nil {
		log.Errorf("unzip extension error: %v", err)
		return
	}

	db.Open(cfg.Dir)
	log.Infof("db open success")
	if cfg.XvfbMod() {
		res := "1920x1080x24"
		log.Infof("xvfb mode, run xvfb %s...", res)
		go exec.Command("Xvfb", ":1", "-screen", "0", res).Run()
		os.Setenv("DISPLAY", ":1")
	}

	server, err := server2.NewServer(cfg)
	if err != nil {
		log.Errorf("new server: %v", err)
		return
	}
	err = server.BridgeAccounts(*accounts)
	if err != nil {
		log.Errorf("bridge accounts: %v", err)
		return
	}
}

func do(cfg *cfg.Config) error {

	FirstRunGenAccount(cfg.AccountsToGen)

	lightsail, err := aws.CreateLightsailClient("airdrop-ubuntu-1", "", path.Join(cfg.Dir, "aws.config"))
	if err != nil {
		return errors.Wrap(err, "create lightsail client")
	}

	for i := 0; i < cfg.AccountsToGen; i++ {
		account := db.FindAccount(i + 1)
		if account.ID == 0 {
			log.Errorf("accounts with id %d not found in db, skip it", i+1)
			continue
		}

		attachedIpName, _, err := lightsail.GetAttachedIp()
		if err != nil {
			return errors.Wrap(err, "get attached static ip")
		}

		ip, ok := db.AccountHasIp(db.ArbitrumService, account.ID)
		if !ok {
			//account has no associated ip, first check current ip, if reach limit then use new ip

			ip2 := db.GetOrAddIp(attachedIpName)
			count := db.IpRelatedAccountNum(db.ArbitrumService, ip2.ID)
			if count >= cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip

				newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
				if err := lightsail.CreateIp(newIpName); err != nil {
					return errors.Wrap(err, "create ip")
				}

				if err := lightsail.DetachIp(attachedIpName); err != nil {
					return errors.Wrap(err, "lightsail detach ip")
				}
				newIp := db.GetOrAddIp(newIpName) //save new ip to db
				rel := db.StaticIpAccountRelation{
					StaticIpID: newIp.ID,
					AccountID:  account.ID,
					Service:    db.ArbitrumService,
				}
				db.SaveAccountIpRelation(&rel)
				if err := lightsail.AttachIp(newIpName); err != nil {
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
				err := lightsail.DetachIp(attachedIpName)
				if err != nil {
					return errors.Wrap(err, "detach ip")
				}
				err = lightsail.AttachIp(ip.Name)
				if err != nil {
					return errors.Wrap(err, "attach ip")
				}
			}

		}

		err = DoAllStep(cfg, account)
		if err != nil {
			log.Errorf("do error: %v", err)
		}

	}
	return nil
}

func FirstRunGenAccount(num int) {
	accountNum := db.AccountNum()
	if accountNum == num {
		log.Infof("accounts have already been generated,skip")
		return
	}
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

func isGasFeeAccepted(cfg *cfg.Config) bool {
	fee, err := utils.GetGasFee(&cfg.Owlracle)
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

type Action func(*metamask.Metamask, func()) error

func DoAllStep(cfg *cfg.Config, account db.Account) error {
	log.Infof("---- [START] begin doing transactions of account %d ----", account.ID)
	if db.HasArbitrumStepAllRun(account.ID) {
		log.Infof("arbitrum steps have all been run, skip account...")
		return nil
	}
	if cfg.RunSingleStep {
		log.Infof("single step mode enabled, will only run one step")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("load-extension", path.Join(cfg.Dir, "ext")),
		chromedp.Flag("window-size", "1920,1080"),

		//chromedp.UserDataDir(dataDir),
	)

	if !cfg.Owlracle.Disable {
		for {
			if isGasFeeAccepted(cfg) {
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
	err := meta.FirstOpenAndImportAccount(cfg.WalletPassword, account)
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
		if cfg.RunSingleStep {
			log.Infof("run single step mode, exiting")
			return nil
		}
	}

	if err := arb.AddL2NetworkAndWaitTransaction(); err != nil {
		return errors.Wrap(err, "add l2 network")
	}

	actions := map[int]Action{
		0: aaveSupplyAndBorrowMoney,
		1: gmxSwapCoin,
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
		if cfg.RunSingleStep {
			log.Infof("run single step mode, exiting")
			return nil
		}
	}
	log.Infof("all steps have done")

	return nil
}

func aaveSupplyAndBorrowMoney(meta *metamask.Metamask, callback func()) error {
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

func gmxSwapCoin(meta *metamask.Metamask, callback func()) error {
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

func generateAccounts(num int, dir string) error {
	for i := 0; i < num; i++ {
		mnemonic, err := utils.NewMnemonic128()
		if err != nil {
			return err
		}
		a := db.Account{Mnemonic: mnemonic}
		db.SaveAccount(&a)

		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("disable-extensions", false),
			chromedp.Flag("restore-on-startup", false),
			chromedp.Flag("disable-web-security", true),
			chromedp.Flag("load-extension", path.Join(dir, "ext")),

			//chromedp.UserDataDir(dataDir),
		)

		ctx, cancel1 := chromedp.NewContext(context.Background())

		allocCtx, cancel2 := chromedp.NewExecAllocator(ctx, opts...)

		meta := metamask.NewMetamask(allocCtx)

		err = meta.FirstOpenAndImportAccount("qwer1234", a)
		cancel1()
		cancel2()
	}
	return nil
}

func linkMetaWithzkSync(allocCtx context.Context) error {

	etherWallet := `//*[@id="__layout"]/main/section/div/div[1]/div[1]/button[1]`
	metaMaskButton := `/html/body/aside/section/ul/li[1]/button`
	//openMetaMaskButton := `/html/body/aside/section/section/footer/a/button`

	zkSyncUrl := `https://wallet.zksync.io/`

	meta := metamask.NewMetamask(allocCtx)
	defer meta.Done()

	if err := meta.OpenAndLogin("metaPasswd"); err != nil {
		return errors.Wrap(err, "open meta")
	}

	zkSyncCtx, c1 := chromedp.NewContext(meta.Context())
	defer c1()

	err := chromedp.Run(zkSyncCtx,
		chromedp.Navigate(zkSyncUrl),
		chromedp.WaitReady(etherWallet),
		chromedp.Click(etherWallet),
		chromedp.WaitReady(metaMaskButton),
		chromedp.Click(metaMaskButton),
		//chromedp.WaitReady(openMetaMaskButton),
		//chromedp.Click(openMetaMaskButton),

	)
	if err != nil {
		log.Errorf("tab error: %v", err)
	}

	log.Info("will open meta wallet")

	return meta.ConfirmLinkAccount(3)

}
