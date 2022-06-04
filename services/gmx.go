package services

import (
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/utils"
	"context"
	"fmt"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const gmxUrl = `https://gmx.io/trade`

func NewGmx(ctx context.Context, meta *metamask.Metamask) *Gmx {
	ctx1, _ := chromedp.NewContext(ctx)
	return &Gmx{
		meta: meta,
		ctx:  ctx1,
	}
}

type Gmx struct {
	meta *metamask.Metamask
	ctx  context.Context
}

func (g *Gmx) OpenAndLinkMeta() error {
	connectWallet := `//*[@id="root"]/div[1]/div/header/div[1]/div[2]/div/button`
	mm := `//*[@id="root"]/div[4]/div[2]/div[3]/button[1]`

	chromedp.Run(g.ctx,
		chromedp.Navigate(gmxUrl),
		chromedp.Click(connectWallet),
		chromedp.Click(mm),
	)
	return g.meta.ConfirmLinkAccount()
}

func (g *Gmx) SwapCoin(from, to string, amount float64) error {
	inputText := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[2]/div[2]/div[1]/input`
	approve := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[6]/button`
	fromCoin := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[2]/div[2]/div[2]/div/div`
	fromTypes := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[2]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div`
	swapButton := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[1]/div[1]/div[3]`

	var nodes []*cdp.Node
	chromedp.Run(g.ctx,
		chromedp.Navigate(gmxUrl),

		chromedp.WaitReady(`//*[@id="headlessui-menu-button-3"]/button`),
		chromedp.Click(swapButton),

		chromedp.Click(fromCoin),
		chromedp.Nodes(fromTypes, &nodes),
	)
	name := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[2]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div[%d]/div[1]/div/div`
	selected := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[2]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div[%d]`
	found := false
	for i := 1; i < len(nodes); i++ {
		//0 is the header
		var s string
		chromedp.Run(g.ctx,
			chromedp.TextContent(fmt.Sprintf(name, i+1), &s),
		)
		if s == from {
			chromedp.Run(g.ctx, chromedp.Click(fmt.Sprintf(selected, i+1)))
			found = true
			log.Infof("select from coin %d", i+1)
			break
		}

	}
	if !found {
		return fmt.Errorf("from coin type not found: %s", from)
	}

	listToCoins := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[4]/div[2]/div[2]/div/div`
	toCoinTypes := `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[4]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div`
	var nodes1 []*cdp.Node
	chromedp.Run(g.ctx,
		chromedp.Sleep(2*time.Second),
		utils.CdpPrint("111"),
		chromedp.Click(listToCoins),
		utils.CdpPrint("222"),

		chromedp.Nodes(toCoinTypes, &nodes1),
	)

	name = `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[4]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div[%d]/div[1]/div/div`
	selected = `//*[@id="root"]/div[1]/div/div/div[1]/div[2]/div[1]/div[1]/div[4]/div[2]/div[2]/div/div[1]/div[2]/div[3]/div/div[%d]`

	found = false
	for i := 1; i < len(nodes1); i++ {
		//0 is the header
		var s string
		chromedp.Run(g.ctx,
			chromedp.TextContent(fmt.Sprintf(name, i+1), &s),
		)
		if s == to {
			chromedp.Run(g.ctx, chromedp.Click(fmt.Sprintf(selected, i+1)))
			found = true
			log.Infof("select to coin %d", i+1)
			break
		}

	}
	if !found {
		return fmt.Errorf("to coin type not found: %s", to)
	}

	chromedp.Run(g.ctx,
		chromedp.SendKeys(inputText, fmt.Sprintf("%f", amount)),
		chromedp.Click(approve),
	)
	return g.meta.ConfirmTransaction()
}
