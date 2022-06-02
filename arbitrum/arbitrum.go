package arbitrum

import (
	"airdrop-bot/metamask"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
)

const arbitrumBridgeUrl = `https://bridge.arbitrum.io/`
const input = `//*[@id="root"]/div/main/div/div/div[4]/div[1]/div[1]/div/div[2]/div/input`
const depositButton = `//*[@id="root"]/div/main/div/div/div[4]/button`

func NewArbitrum(ctx context.Context, meta *metamask.Metamask) *Arbitrum {
	ctx1, _ := chromedp.NewContext(ctx)
	return &Arbitrum{
		ctx:  ctx1,
		meta: meta,
	}
}

type Arbitrum struct {
	ctx  context.Context
	meta *metamask.Metamask
}

func (a *Arbitrum) Deposit(amount float64) error {
	return chromedp.Run(a.ctx,
		chromedp.Navigate(arbitrumBridgeUrl),
		chromedp.SendKeys(input, fmt.Sprintf("%f", amount)),
		chromedp.Click(depositButton),
	)
}

func (a *Arbitrum) LinkMetaMask() error {

	metaMaskIcon := `//*[@id="WEB3_CONNECT_MODAL_ID"]/div/div/div[2]/div[1]/div`
	tosButton := `//*[@id="headlessui-dialog-8"]/div/div[2]/div[2]/button`
	chromedp.Run(a.ctx,
		chromedp.Navigate(arbitrumBridgeUrl),
		chromedp.Click(metaMaskIcon),
	)
	a.meta.ConfirmLinkAccount()

	return chromedp.Run(a.ctx, chromedp.Click(tosButton))
}
