package main

import (
	"airdrop-bot/arbitrum"
	"airdrop-bot/metamask"
	"context"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"log"
	"os"
	"path"
	"time"
)

func main() {

	dir, _ := os.Getwd()
	dataDir := path.Join(dir, "data")
	os.Mkdir(dataDir, 0777)
	print(dataDir)
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		//chromedp.Flag("load-extension", path.Join(dir, "ext")),

		chromedp.UserDataDir(dataDir),
	)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	meta := metamask.NewMetamask(allocCtx)
	meta.OpenAndLogin(metaPasswd)

	arb := arbitrum.NewArbitrum(meta.Context(), meta)

	arb.LinkMetaMask()

	time.Sleep(time.Minute)
}

var metaPasswd = os.Args[1]

func linkMetaWithzkSync(allocCtx context.Context) error {

	etherWallet := `//*[@id="__layout"]/main/section/div/div[1]/div[1]/button[1]`
	metaMaskButton := `/html/body/aside/section/ul/li[1]/button`
	//openMetaMaskButton := `/html/body/aside/section/section/footer/a/button`

	zkSyncUrl := `https://wallet.zksync.io/`

	meta := metamask.NewMetamask(allocCtx)
	defer meta.Done()

	if err := meta.OpenAndLogin(metaPasswd); err != nil {
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
		log.Print("tab error: ", err.Error())
	}

	log.Print("will open meta wallet")

	return meta.ConfirmLinkAccount(3)

}
