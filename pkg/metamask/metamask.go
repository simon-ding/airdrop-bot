package metamask

import (
	"airdrop-bot/log"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"strconv"
	"strings"
	"time"
)

const metaExtUrl = `chrome-extension://nkbihfbeogaeaoehlefnkodbefgpgknn/home.html#`

const metaChromePluginStoreUrl = "https://chrome.google.com/webstore/detail/metamask/nkbihfbeogaeaoehlefnkodbefgpgknn"

const (
	accountPos               = `//*[@id="app-content"]/div/div[3]/div[4]/div[3]/div[%d]`
	avatarPos                = `//*[@id="app-content"]/div/div[1]/div/div[2]/div[2]/div`
	accountExpandPos         = `//*[@id="app-content"]/div/div[3]/div/div/div/div[1]/button`
	showLinkedSitesButtonPos = `//*[@id="popover-content"]/div[2]/button[3]`
	accountDetail            = `//*[@id="popover-content"]/div[2]/button[2]`
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

func (m *Metamask) FirstOpenAndImportAccount(walletPass string, mnemonic string) error {
	log.Infof("begin open metamask and import account")
	words := strings.Split(mnemonic, " ")
	if len(words) != 12 {
		return fmt.Errorf("mnemonic word len is not 12")
	}

	beginButton := `//*[@id="app-content"]/div/div[2]/div/div/div/button`
	importButton := `//*[@id="app-content"]/div/div[2]/div/div/div[2]/div/div[2]/div[1]/button`
	agreeButton := `//*[@id="app-content"]/div/div[2]/div/div/div/div[5]/div[1]/footer/button[2]`
	mnemonicInput := `//*[@id="import-srp__srp-word-%d"]`

	chromedp.Run(m.ctx,
		chromedp.Navigate(metaExtUrl),
		chromedp.Click(beginButton),
		chromedp.Click(importButton),
		chromedp.Click(agreeButton),
	)

	firstPasswd := `//*[@id="password"]`
	secondPasswd := `//*[@id="confirm-password"]`
	tosCheckbox := `//*[@id="create-new-vault__terms-checkbox"]`
	importButton2 := `//*[@id="app-content"]/div/div[2]/div/div/div[2]/form/button`
	allDoneButton := `//*[@id="app-content"]/div/div[2]/div/div/button`
	closeNewFeatureButton := `//*[@id="popover-content"]/div/div/section/div[1]/div/button`
	addressPos := `//*[@id="app-content"]/div/span/div[1]/div/div/div/div[3]/div[2]/div/div/div[1]`
	for i := 0; i < len(words); i++ {
		mnemonicInput := fmt.Sprintf(mnemonicInput, i)
		chromedp.Run(m.ctx,
			chromedp.SendKeys(mnemonicInput, words[i]),
		)
	}

	var address string
	chromedp.Run(m.ctx,
		chromedp.SendKeys(firstPasswd, walletPass),
		chromedp.SendKeys(secondPasswd, walletPass),
		chromedp.Click(tosCheckbox),
		chromedp.Click(importButton2),
		chromedp.Click(allDoneButton),
		chromedp.Click(closeNewFeatureButton),
		chromedp.Click(accountExpandPos),
		chromedp.Click(accountDetail),
		chromedp.TextContent(addressPos, &address),
	)
	log.Infof("import account done")
	return nil
}

func (m *Metamask) OpenAndLogin(passWd string) error {
	unlock := `/html/body/div[1]/div/div[3]/div/div/button`
	return chromedp.Run(m.ctx,
		chromedp.Navigate(metaExtUrl),
		chromedp.SendKeys(`//*[@id="password"]`, passWd),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitEnabled(unlock),
		chromedp.Click(unlock), //login
		chromedp.Sleep(2*time.Second),
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

func (m *Metamask) ConfirmLinkAccount(accounts ...int) error {
	confirmButton := `//*[@id="app-content"]/div/div[2]/div/div[3]/div[2]/button[2]`
	linkButton := `//*[@id="app-content"]/div/div[2]/div/div[2]/div[2]/div[2]/footer/button[2]`
	accountsPos := `//*[@id="app-content"]/div/div[2]/div/div[2]/div[2]/div[2]/div/div`
	checkBoxPos := `//*[@id="app-content"]/div/div[2]/div/div[2]/div[2]/div[2]/div/div[%d]/div/input`

	log.Infof("confirm link metamask wallet")
	var nodes []*cdp.Node
	chromedp.Run(m.ctx,
		chromedp.Reload(),
		chromedp.Nodes(accountsPos, &nodes),
	)
	if len(accounts) > 0 {
		for i := 0; i < len(nodes); i++ {
			html := ""
			currentCheckbox := fmt.Sprintf(checkBoxPos, i+1)
			chromedp.Run(m.ctx,
				chromedp.OuterHTML(currentCheckbox, &html),
			)
			if strings.Contains(strings.ToLower(html), "checked") {
				//unchecked account
				log.Infof("account %d checked, uncheck it", i+1)
				chromedp.Run(m.ctx, chromedp.Click(currentCheckbox))
			}
		}

		for i := 0; i < len(accounts); i++ {
			chromedp.Run(m.ctx, chromedp.Click(fmt.Sprintf(checkBoxPos, accounts[i])))
		}
	}

	return chromedp.Run(m.ctx,
		chromedp.WaitReady(confirmButton),
		chromedp.Click(confirmButton),
		chromedp.WaitReady(linkButton),
		chromedp.Click(linkButton),
	)
}

func (m *Metamask) ConfirmAddNetwork() error {
	approveButton := `//button[text()='Approve']`
	switchNetwork := `//button[text()='Switch network']`
	log.Infof("confirm add l2 network")
	return chromedp.Run(m.ctx,
		chromedp.Sleep(time.Second),
		chromedp.Reload(),
		chromedp.Click(approveButton),
		chromedp.Sleep(time.Second),
		chromedp.Click(switchNetwork),
	)
}

func (m *Metamask) SwitchAccount(n int) error {
	acc := fmt.Sprintf(accountPos, n)
	return chromedp.Run(m.ctx,
		chromedp.Click(avatarPos),
		chromedp.Click(acc),
	)
}

//UnlinkAllSites 取消账户所有网站的链接
func (m *Metamask) UnlinkAllSites() error {
	linkedSites := `//*[@id="popover-content"]/div/div/section/div[2]/main/div`
	var nodes []*cdp.Node
	err := chromedp.Run(m.ctx,
		chromedp.Click(accountExpandPos),
		chromedp.Click(showLinkedSitesButtonPos),
		chromedp.Nodes(linkedSites, &nodes),
		cdpPrint(fmt.Sprint(nodes)),
	)
	if err != nil {
		return err
	}

	site := `//*[@id="popover-content"]/div/div/section/div[2]/main/div[%d]/div/span`
	unlinkButton := `//*[@id="popover-content"]/div/div/section/div[2]/main/div[%d]/a`
	confirmButton := `//*[@id="popover-content"]/div/div/section/div[2]/div/button[2]`
	var title string
	for i := 0; i < len(nodes); i++ {
		chromedp.Run(m.ctx,
			chromedp.TextContent(fmt.Sprintf(site, i+1), &title),
		)
		log.Infof("try to unlink %s", title)
		chromedp.Run(m.ctx,
			chromedp.Click(fmt.Sprintf(unlinkButton, i+1)),
			chromedp.Click(confirmButton),
		)
	}
	return nil
}

// SwitchNetwork 切换钱包网络
func (m *Metamask) SwitchNetwork(net string) error {
	networkListExpendPos := `//*[@id="app-content"]/div/div[1]/div/div[2]/div[1]/div`
	networkList := `//*[@id="app-content"]/div/div[2]/div/div[2]/li`

	var nodes []*cdp.Node
	chromedp.Run(m.ctx,
		chromedp.Click(networkListExpendPos),
		chromedp.Nodes(networkList, &nodes),
	)

	networkName := `//*[@id="app-content"]/div/div[2]/div/div[2]/li[%d]/span`
	for i := 0; i < len(nodes); i++ {
		var name string
		chromedp.Run(m.ctx,
			chromedp.TextContent(fmt.Sprintf(networkName, i+1), &name),
		)

		if strings.Contains(name, net) {
			chromedp.Run(m.ctx,
				chromedp.Click(fmt.Sprintf(`//*[@id="app-content"]/div/div[2]/div/div[2]/li[%d]`, i+1)),
			)
			return nil
		}
	}
	return errors.New("network not found: " + net)
}

func (m *Metamask) ConfirmTransaction() error {
	confirmButton := `//button[@data-testid='page-container-footer-next']`
	log.Infof("confirm metamask transaction")
	return chromedp.Run(m.ctx,
		chromedp.Sleep(5*time.Second),
		chromedp.Reload(),
		chromedp.WaitEnabled(confirmButton),
		chromedp.Click(confirmButton),
	)
}

func (m *Metamask) SignTransaction() error {
	log.Infof("begin sign transaction")
	signButton := `//*[@id="app-content"]/div/div[3]/div/div[4]/button[2]`
	return chromedp.Run(m.ctx,
		chromedp.Click(signButton),
	)
}

func (m *Metamask) Balance() (float64, error) {
	var s string
	balancePos := `//span[@class='currency-display-component__text']`
	chromedp.Run(m.ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.Reload(),
		chromedp.TextContent(balancePos, &s),
	)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, errors.Wrap(err, "parse float")
	}
	return f, nil
}

func (m *Metamask) OpenChanListAndAddNetwork(networkName string) error {
	chanListUrl := `https://chainlist.org/zh`
	ctx, cancel := chromedp.NewContext(m.ctx)
	searchPos := `//input[@placeholder="ETH, Fantom, ..."]`
	addNetworkButton := `//*[@id="__next"]/div/main/div/div[2]/div[2]/div[1]/div[3]/button`
	defer cancel()

	linkWallet := `//*[@id="__next"]/div[1]/main/div/div[2]/div[1]/button`
	log.Infof("try to open chanlist and add %s network", networkName)

	err := chromedp.Run(ctx,
		chromedp.Navigate(chanListUrl),
		chromedp.Click(linkWallet),
		chromedp.Sleep(time.Second),
	)
	if err != nil {
		return err
	}
	err = m.ConfirmLinkAccount()
	if err != nil {
		return err
	}

	log.Infof("link meta to chanlist success")
	err = chromedp.Run(ctx,
		chromedp.SendKeys(searchPos, networkName),
		chromedp.Sleep(3*time.Second),
		chromedp.Click(addNetworkButton),
		chromedp.Sleep(3*time.Second),
	)
	if err != nil {
		return err
	}
	return m.ConfirmAddNetwork()
}

func (m *Metamask) WaitForBalanceChange(ori float64) float64 {
	for {
		log.Infof("checking balance")
		balance1, err := m.Balance()
		if err != nil {
			log.Errorf("check balance error: %v", err)
			continue
		}
		if ori != balance1 {
			log.Infof("balance has now changed, now balance is %v", balance1)
			return balance1
		}
		time.Sleep(2 * time.Second)
	}
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

func cdpPrint(s string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		fmt.Println(s)
		return nil
	})

}
