package services

import (
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"time"
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
	confirmMove := `//*[@id="headlessui-dialog-13"]/div/div[2]/div[2]/button[1]`
	err := chromedp.Run(a.ctx,
		chromedp.Navigate(arbitrumBridgeUrl),
		chromedp.SendKeys(input, fmt.Sprintf("%f", amount)),
		chromedp.Sleep(1*time.Second),
		chromedp.WaitEnabled(depositButton),
		chromedp.Click(depositButton),
		chromedp.Click(confirmMove),
	)
	if err != nil {
		return err
	}
	if err := a.meta.ConfirmTransaction(); err != nil {
		return errors.Wrap(err, "meta confirm transaction")
	}
	log.Infof("deposit %f eth to arb success, waiting for transaction ends...", amount)

	depositStatus := `//*[@id="root"]/div/main/div/div/div[5]/div/div/div/div/table/tbody/tr/td[2]/span`
	for {
		var status string
		chromedp.Run(a.ctx,
			chromedp.Reload(),
			chromedp.TextContent(depositStatus, &status),
		)
		if status == "success" {
			log.Infof("transaction success ends")
			break
		}
		log.Debugf("current transaction status: %v", status)
		time.Sleep(30 * time.Second)
	}
	return nil
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
