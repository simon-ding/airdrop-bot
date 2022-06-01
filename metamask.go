package main

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"time"
)

const metaExtUrl = `chrome-extension://nkbihfbeogaeaoehlefnkodbefgpgknn/home.html#`

const metaChromePluginStoreUrl = "https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn"

const (
	accountPos               = `//*[@id="app-content"]/div/div[3]/div[4]/div[3]/div[%d]`
	avatarPos                = `//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`
	accountExpandPos         = `//*[@id="app-content"]/div/div[3]/div/div/div/div[1]/button`
	showLinkedSitesButtonPos = `//*[@id="popover-content"]/div[2]/button[3]`
)

func NewMetamask(ctx context.Context) *Metamask {
	ctx1, can := chromedp.NewContext(ctx)
	return &Metamask{
		ctx:    ctx1,
		cancel: can,
	}
}

type Metamask struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func (m *Metamask) Open() error {
	return chromedp.Run(m.ctx,
		chromedp.Navigate(metaExtUrl),
	)
}

func (m *Metamask) Context() context.Context {
	return m.ctx
}
func (m *Metamask) Done() {
	m.cancel()
}

func (m *Metamask) OpenAndLogin(passWd string) error {
	return chromedp.Run(m.ctx,
		chromedp.Navigate(metaExtUrl),
		chromedp.SendKeys(`//*[@id="password"]`, passWd),
		chromedp.Click(`//*[@id="app-content"]/div/div[3]/div/div/button`, chromedp.NodeVisible), //login
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`),            //avatar
	)
}

func (m *Metamask) CreateAccounts(n int) error {
	actions := []chromedp.Action{
		chromedp.Click(avatarPos),                                      // click avatar
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div[6]`), //创建账户
		chromedp.Click(`//*[@id="app-content"]/div/div[3]/div[6]`),
		chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`), //账户名
		chromedp.SendKeys(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`, fmt.Sprintf("account_%d", n)),
	}

	for i := 0; i < n; i++ {
		act := []chromedp.Action{
			chromedp.WaitReady(avatarPos), //avatar
			chromedp.Click(`//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`), // click avatar
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div[6]`),            //创建账户
			chromedp.Click(`//*[@id="app-content"]/div/div[3]/div[6]`),
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`), //账户名
			chromedp.SendKeys(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/input`, fmt.Sprintf("account_%d", i)),
			chromedp.WaitReady(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/div/button[2]`),
			chromedp.Click(`//*[@id="app-content"]/div/div[3]/div/div/div/div[2]/div/button[2]`), //创建button
		}
		actions = append(actions, act...)
	}

	return chromedp.Run(m.ctx,
		actions...,
	)
}

func (m *Metamask) ConfirmLinkAccount() error {
	confirmButton := `//*[@id="app-content"]/div/div[2]/div/div[3]/div[2]/button[2]`
	linkButton := `//*[@id="app-content"]/div/div[2]/div/div[2]/div[2]/div[2]/footer/button[2]`

	return chromedp.Run(m.ctx,
		chromedp.Reload(),
		chromedp.WaitReady(confirmButton),
		chromedp.Click(confirmButton),
		chromedp.WaitReady(linkButton),
		chromedp.Click(linkButton),
		chromedp.Sleep(time.Minute),
	)
}

func (m *Metamask) SwitchAccount(n int) error {
	acc := fmt.Sprintf(accountPos, n)
	return chromedp.Run(m.ctx,
		chromedp.Click(avatarPos),
		chromedp.Click(acc),
	)
}

func (m *Metamask) UnlinkAllSites() error {
	linkedSites := `//*[@id="popover-content"]/div/div/section/div[2]/main`
	var nodes []*cdp.Node
	return chromedp.Run(m.ctx,
		chromedp.Click(accountExpandPos),
		chromedp.Click(showLinkedSitesButtonPos),
		chromedp.Nodes(linkedSites, &nodes),
		cdpPrint(fmt.Sprint(nodes)),
	)
}

func (m *Metamask) installMetaChromePlugin() error {
	ctx1, cancel1 := chromedp.NewContext(m.ctx)
	defer cancel1()
	installButtom := `/html/body/div[3]/div[2]/div/div/div[2]/div[2]/div`

	chromedp.Run(ctx1,
		chromedp.Navigate(metaChromePluginStoreUrl),
		cdpPrint("store opened"),
		//chromedp.Click(`/html/body/div[3]/div[2]/div/div/div[2]/div[1]/div/div[1]/span[1]/a`),
		chromedp.Click(installButtom),
		cdpPrint("clicked"),
	)

	targets, err := chromedp.Targets(ctx1)
	if err != nil {
		return err
	}
	fmt.Println(targets)
	return nil

}
