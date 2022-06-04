package main

import (
	"airdrop-bot/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/services"
	"airdrop-bot/utils"
	"context"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"math/rand"
	"os"
	"path"
	"time"
)

func main() {
	dir, _ := os.Getwd()

	cfg, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")

	db.Open(dir)
	log.Infof("db open success")

	//fee, err := utils.GetGasFee(&cfg.Owlracle)
	//if err != nil {
	//	log.Errorf("get gas fee error: %v", err)
	//}
	//log.Infof("current gas fee is %+v", *fee)

	dataDir := path.Join(dir, "data")
	os.Mkdir(dataDir, 0777)
	log.Infof("use chrome user dir: %v", dataDir)
	if err := do(dir, cfg); err != nil {
		log.Errorf("do error: %v", err)
		return
	}
	time.Sleep(time.Minute)
}

func do(dir string, cfg *cfg.Config) error {
	lightsail, err := aws.CreateLightsailClient("", path.Join(dir, "asserts", "aws.config"))
	if err != nil {
		return errors.Wrap(err, "create lightsail client")
	}

	for i := 0; i < 50; i++ {
		account := db.FindAccount(i + 1)

		attachedIpName, err := lightsail.GetAttachedIp()
		if err != nil {
			return errors.Wrap(err, "get attached static ip")
		}

		ip, ok := db.AccountHasIp(db.ArbitrumService, account.ID)
		if !ok {
			//account has no associated ip, first check current ip, if reach limit then use new ip

			ip2 := db.GetOrAddIp(attachedIpName)
			count := db.IpRelatedAccountNum(db.ArbitrumService, ip2.ID)
			if count >= cfg.AccountsPerIp { //当前ip超过允许的最大数量，换新的ip
				if err := lightsail.DetachIp(ip.Name); err != nil {
					return errors.Wrap(err, "lightsail detach ip")
				}
				newIpName := "airdrop_bot_" + utils.RandStringRunes(4)
				if err := lightsail.CreateIp(newIpName); err != nil {
					return errors.Wrap(err, "create ip")
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

		err = DoAllStep(dir, cfg, account)
		if err != nil {
			log.Errorf("do error: %v", err)
		}

	}
	return nil
}

type Action func(*metamask.Metamask) error

func DoAllStep(dir string, cfg *cfg.Config, account db.Account) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("load-extension", path.Join(dir, "ext")),

		//chromedp.UserDataDir(dataDir),
	)

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	meta := metamask.NewMetamask(allocCtx)
	err := meta.FirstOpenAndImportAccount(cfg.WalletPassword, account)
	if err != nil {
		return err
	}
	balance, err := meta.Balance()
	log.Info(balance)

	arb := services.NewArbitrum(meta.Context(), meta)
	err = arb.LinkMetaMask()
	if err != nil {
		return err
	}
	const depositReverse = 0.005

	err = arb.Deposit(balance - depositReverse)
	if err != nil {
		return errors.Wrap(err, "arb deposit")
	}
	actions := map[int]Action{
		1: aaveSupplyAndBorrowMoney,
		2: gmxSwapCoin,
	}

	for len(actions) != 0 {
		n := rand.Intn(len(actions))
		act := actions[n]
		if err := act(meta); err != nil {
			return err
		}
		delete(actions, n)
	}

	return nil
}

func aaveSupplyAndBorrowMoney(meta *metamask.Metamask) error {
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
	return nil
}

func gmxSwapCoin(meta *metamask.Metamask) error {
	gmx := services.NewGmx(meta.Context(), meta)
	n2 := rand.Intn(5)
	var gmxSwap = float64(10+n2) / 1000.
	if err := gmx.SwapCoin("ETH", "USDT", gmxSwap); err != nil {
		return errors.Wrap(err, "gmx swap coin")
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
