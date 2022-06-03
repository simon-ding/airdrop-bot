package services

import (
	"airdrop-bot/metamask"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

const aaveUrl = `https://app.aave.com/?marketName=proto_arbitrum_v3`

func NewAave(ctx context.Context, meta *metamask.Metamask) *Aave {
	ctx1, _ := chromedp.NewContext(ctx)
	return &Aave{
		ctx:  ctx1,
		meta: meta,
	}
}

type Aave struct {
	ctx  context.Context
	meta *metamask.Metamask
}

func (a *Aave) OpenAndLinkMetaMask() error {
	connectWallet := `//*[@id="wallet-button"]`
	browserWallet := `/html/body/div[7]/div[3]/div[1]/button[1]`
	chromedp.Run(a.ctx,
		chromedp.Navigate(aaveUrl),
		chromedp.Click(connectWallet),
		chromedp.Click(browserWallet),
	)
	return a.meta.ConfirmLinkAccount()
}

func (a *Aave) SupplyEth(amount float64) error {
	asserts := `//*[@id="__next"]/main/div[2]/div/div[2]/div[1]/div[2]/div[3]/div`
	var nodes []*cdp.Node

	chromedp.Run(a.ctx,
		chromedp.Nodes(asserts, &nodes),
	)

	coinType := `//*[@id="__next"]/main/div[2]/div/div[2]/div[1]/div[2]/div[3]/div[%d]/div[1]/a/p`
	supplyButton := `//*[@id="__next"]/main/div[2]/div/div[2]/div[1]/div[2]/div[3]/div[%d]/div[5]/button`
	moneyInput := `/html/body/div[8]/div[3]/div[1]/div[2]/div[1]/div[1]/input`
	approveSupply := `/html/body/div[8]/div[3]/div[4]/button[1]`
	// 0 is the header
	for i := 1; i < len(nodes); i++ {
		var t string
		chromedp.Run(a.ctx,
			chromedp.TextContent(fmt.Sprintf(coinType, i+1), &t),
		)
		if t == "ETH" {
			chromedp.Run(a.ctx,
				chromedp.Click(fmt.Sprintf(supplyButton, i+1)),
				chromedp.SendKeys(moneyInput, fmt.Sprintf("%f", amount)),
				chromedp.Click(approveSupply),
			)

			a.meta.SignTransaction()

		}
	}
	return nil
}

func (a *Aave) BorrowMoney(coinType string, amount float64) error {
	availableCurrency := `//*[@id="__next"]/main/div[2]/div/div[2]/div[2]/div[2]/div[3]/div`
	types := `//*[@id="__next"]/main/div[2]/div/div[2]/div[2]/div[2]/div[3]/div[%d]/div[1]/a/p`
	borrowButton := `//*[@id="__next"]/main/div[2]/div/div[2]/div[2]/div[2]/div[3]/div[%d]/div[5]/button`
	borrowDialog := `/html/body/div[8]/div[3]/div[1]/div[2]/div[1]/div[1]/input`
	confirmBorrow := `/html/body/div[8]/div[3]/div[4]/button`
	var nodes []*cdp.Node
	chromedp.Run(a.ctx,
		chromedp.Nodes(&availableCurrency, &nodes),
	)

	//0 is the header
	for i := 0; i < len(nodes); i++ {
		var t string
		chromedp.Run(a.ctx,
			chromedp.TextContent(fmt.Sprintf(types, i+1), &t),
		)
		if t == coinType {
			chromedp.Run(a.ctx,
				chromedp.Click(fmt.Sprintf(borrowButton, i+1)),
				chromedp.SendKeys(borrowDialog, fmt.Sprintf("%f", amount)),
				chromedp.Click(confirmBorrow),
			)

			return a.meta.ConfirmTransaction()
		}
	}
	return fmt.Errorf("no coin of type %s found", coinType)
}
