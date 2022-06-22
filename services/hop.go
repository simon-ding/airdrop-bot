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

const url = "https://app.hop.exchange/#/send?token=%s&sourceNetwork=%s&destNetwork=%s"

func MewHop(ctx0 context.Context, meta *metamask.Metamask, token, sourceNet, destNet string) *Hop {
	ctx, _ := chromedp.NewContext(ctx0)
	url1 := fmt.Sprintf(url, token, sourceNet, destNet)
	return &Hop{
		meta: meta,
		ctx:  ctx,
		url:  url1,
	}
}

type Hop struct {
	meta *metamask.Metamask
	ctx  context.Context
	url  string
}

func (h *Hop) LinkMetaMask() error {
	connectWallet := `//button[@class="MuiButtonBase-root MuiButton-root MuiButton-text jss27 jss30 sc-iqseJM bKgVTO"]`
	metamaskButton := `//button[@class="bn-onboard-custom bn-onboard-icon-button svelte-1799bj2"]`
	log.Infof("try to link metamask")
	err := chromedp.Run(h.ctx,
		chromedp.Navigate(h.url),
		chromedp.Click(connectWallet),
		chromedp.Click(metamaskButton),
		chromedp.Sleep(2*time.Second),
	)
	if err != nil {
		return err
	}

	return h.meta.ConfirmLinkAccount()
}

func (h *Hop) BridgeMoney(amount float64) error {
	input := `//input[@class="MuiInputBase-input MuiInput-input jss92 jss95 MuiInputBase-inputAdornedEnd"]`
	sendButton := `//button[@class="MuiButtonBase-root MuiButton-root MuiButton-text jss27 jss112 jss58"]`
	sendButton2 := `//button[@class="MuiButtonBase-root MuiButton-root MuiButton-text jss27 jss134 jss131"]`
	log.Infof("bridge money %f", amount)
	err := chromedp.Run(h.ctx,
		chromedp.Sleep(5*time.Second),
		chromedp.SendKeys(input, fmt.Sprintf("%f", amount)),
		chromedp.Sleep(time.Second),
		chromedp.Click(sendButton),
		chromedp.Click(sendButton2),
	)
	if err != nil {
		return err
	}
	if err := h.meta.ConfirmTransaction(); err != nil {
		return errors.Wrap(err, "meta confirm transaction")
	}
	t := 2 * time.Minute
	log.Infof("deposit %f eth to arb (using hop) success, wait %v", amount, t)
	time.Sleep(t)
	return nil

}
