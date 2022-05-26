package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"time"

	"github.com/chromedp/chromedp"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func main() {

	dir, _ := os.Getwd()
	dataDir := path.Join(dir, "data")
	os.Mkdir(dataDir, 0777)
	print(dataDir)

	err := linkMetaWithzkSync(dataDir)

	if err != nil {
		log.Fatal(err)
	}
	print("done")
}

func cdpPrint(s string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		fmt.Println(s)
		return nil
	})

}

const metaExtUrl = `chrome-extension://nkbihfbeogaeaoehlefnkodbefgpgknn/home.html#`

func getTasks() []chromedp.Action {
	acts := openMetaWalletActions()

	acts = append(acts, createAccounts(10)...)

	acts = append(acts, chromedp.ActionFunc(func(ctx context.Context) error {
		time.Sleep(1 * time.Minute)
		return nil
	}))
	return acts
}

var metaPasswd = os.Args[1]

func openMetaWalletActions() []chromedp.Action {
	return []chromedp.Action{
		chromedp.Navigate(metaExtUrl),
		chromedp.SendKeys(`//*[@id="password"]`, metaPasswd),
		chromedp.Click(`//*[@id="app-content"]/div/div[3]/div/div/button`, chromedp.NodeVisible), //login
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`),            //avatar
	}
}

func createAccounts(n int) []chromedp.Action {

	actions := []chromedp.Action{
		chromedp.Click(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`), // click avatar
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div[6]`),            //创建账户
		chromedp.Click(`//*[@id="app-content"]/div/div[3]/div[6]`),
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`), //账户名
		chromedp.SendKeys(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`, fmt.Sprintf("account_%d", n)),
	}

	for i := 0; i < n; i++ {
		act := []chromedp.Action{
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`), //avatar
			chromedp.Click(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`),     // click avatar
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div[6]`),                //创建账户
			chromedp.Click(`//*[@id="app-content"]/div/div[3]/div[6]`),
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`), //账户名
			chromedp.SendKeys(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`, fmt.Sprintf("account_%d", i)),
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/div/button[2]`),
			chromedp.Click(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/div/button[2]`), //创建button
		}
		actions = append(actions, act...)
	}
	return actions
}

func linkMetaWithzkSync(dataDir string) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.UserDataDir(dataDir),
	)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	allocCtx, cancel := chromedp.NewExecAllocator(ctx, opts...)
	defer cancel()

	etherWallet := `//*[@id="__layout"]/main/section/div/div[1]/div[1]/button[1]`
	metaMaskButton := `/html/body/aside/section/ul/li[1]/button`
	//openMetaMaskButton := `/html/body/aside/section/section/footer/a/button`

	zkSyncUrl := `https://wallet.zksync.io/`

	ctx2, cancel1 := chromedp.NewContext(allocCtx)
	defer cancel1()

	if err := chromedp.Run(ctx2,
		openMetaWalletActions()...,
	); err != nil {
		log.Println(err)
	}

	ctx1, c1 := chromedp.NewContext(ctx2)
	defer c1()

	err := chromedp.Run(ctx1,
		chromedp.Navigate(zkSyncUrl),
		chromedp.WaitReady(etherWallet),
		chromedp.Click(etherWallet),
		chromedp.WaitReady(metaMaskButton),
		chromedp.Click(metaMaskButton),
		//chromedp.WaitReady(openMetaMaskButton),
		//chromedp.Click(openMetaMaskButton),

	)
	if err != nil {
		log.Print("tab error: ", string(GbkToUtf8([]byte(err.Error()))))
	}

	log.Print("will open meta wallet")
	var metaCtx = metaPluginCtx(ctx1)

	confirmButton := `//*[@id="app-content"]/div/div[2]/div/div[3]/div[2]/button[2]`
	title := ""
	chromedp.Run(metaCtx,
		cdpPrint("in extension"),
		chromedp.Sleep(5*time.Second),
		chromedp.Reload(),
		chromedp.Title(&title),
		cdpPrint("TITLE: "+title),
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[2]/div/div[2]/div[1]/div[2]`),
		chromedp.Click(confirmButton),
	)

	return err

}

func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}

func metaPluginCtx(ctx context.Context) context.Context {
	targets, err := chromedp.Targets(ctx)
	if err != nil {
		log.Println("target error: ", err)
	}
	for _, t := range targets {
		if t.Type == "iframe" && t.Title == "MetaMask" {
			metaCtx, _ := chromedp.NewContext(ctx, chromedp.WithTargetID(t.TargetID))
			return metaCtx
		}
	}
	log.Println("no meta plugin installed")
	return nil
}
