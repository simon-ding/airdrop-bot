package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/pkg/errors"
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

	acc, _ := createEthAccount()
	log.Println("created account: ", acc)

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

	err := linkMetaWithzkSync(allocCtx)

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

func linkMetaWithzkSync(allocCtx context.Context) error {

	etherWallet := `//*[@id="__layout"]/main/section/div/div[1]/div[1]/button[1]`
	metaMaskButton := `/html/body/aside/section/ul/li[1]/button`
	//openMetaMaskButton := `/html/body/aside/section/section/footer/a/button`

	zkSyncUrl := `https://wallet.zksync.io/`

	meta := NewMetamask(allocCtx)
	defer meta.Done()

	if err := meta.OpenAndLogin(metaPasswd); err != nil {
		return errors.Wrap(err, "open meta")
	}
	meta.UnlinkAllSites()

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
		log.Print("tab error: ", string(GbkToUtf8([]byte(err.Error()))))
	}

	log.Print("will open meta wallet")

	return meta.ConfirmLinkAccount()

}

func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}
