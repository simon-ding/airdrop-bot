package services

import (
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"time"
)

const quickSwapUrl = `https://quickswap.exchange/#/swap`

func NewQuickSwap(meta *metamask.Metamask) *QuickSwap {
	ctx, _ := chromedp.NewContext(meta.Context())
	return &QuickSwap{
		ctx:  ctx,
		meta: meta,
	}
}

type QuickSwap struct {
	ctx  context.Context
	meta *metamask.Metamask
}

func (q *QuickSwap) LinkMetamask() error {
	connectWallet := `//button[@id="connect-wallet"]`
	metaButton := `//button[@id="connect-METAMASK"]`
	err := chromedp.Run(q.ctx,
		chromedp.Navigate(quickSwapUrl),
		chromedp.Click(connectWallet),
		chromedp.Click(metaButton),
	)
	if err != nil {
		return err
	}
	err = q.meta.ConfirmLinkAccount()
	if err != nil {
		return err
	}
	log.Infof("quick swap link metamask success")
	return nil
}

func (q *QuickSwap) SwapMaticToEth(amount float64) error {
	inputArea := `//button[@class="sc-kGVuwA grroav token-amount-input"]`
	sendButton := `//button[@id="swap-button"]`
	confirmSendButton := `//button[@id="confirm-swap-or-send"]`
	err := chromedp.Run(q.ctx,
		chromedp.Sleep(2*time.Second),
		chromedp.SendKeys(inputArea, fmt.Sprintf("%v", amount)),
		chromedp.Sleep(5*time.Second),
		chromedp.Click(sendButton),
		chromedp.Click(confirmSendButton),
		chromedp.Sleep(5*time.Second),
	)
	if err != nil {
		return err
	}

	err = q.meta.ConfirmTransaction()
	if err != nil {
		return err
	}
	log.Infof("swap matic to eth success")
	return nil
}
