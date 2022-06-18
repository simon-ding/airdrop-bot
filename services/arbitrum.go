package services

import (
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/utils"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"time"
)

const arbitrumBridgeUrl = `https://bridge.arbitrum.io/`
const input = `//input[@placeholder='Enter amount here']`
const depositButton = `//button[text()='Deposit']`
const confirmMove = `//*[@id="headlessui-dialog-13"]/div/div[2]/div[2]/button[1]`

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
	log.Infof("run arbitrum deposit: %v", amount)
	err := chromedp.Run(a.ctx,
		chromedp.Navigate(arbitrumBridgeUrl),
		chromedp.Sleep(5*time.Second),
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
	return nil
}

func (a *Arbitrum) DepositUsingHop(amount float64) error {
	url := `https://app.hop.exchange/#/send?token=ETH&destNetwork=arbitrum&sourceNetwork=ethereum`
	ctx, cancel := chromedp.NewContext(a.ctx)
	defer cancel()
	input := `//input[@class='MuiInputBase-root MuiInput-root jss91 jss94 MuiInputBase-formControl MuiInput-formControl MuiInputBase-adornedEnd']`
	sendButton := `//button[@class="MuiButtonBase-root MuiButton-root MuiButton-text jss23 jss112 jss58"]`
	sendButton2 := `//button[@class="MuiButtonBase-root MuiButton-root MuiButton-text jss23 jss166 jss163"]`
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Sleep(5*time.Second),
		chromedp.SendKeys(input, fmt.Sprintf("%f", amount)),
		chromedp.Sleep(time.Second),
		chromedp.Click(sendButton),
		chromedp.Click(sendButton2),
	)
	if err != nil {
		return err
	}
	if err := a.meta.ConfirmTransaction(); err != nil {
		return errors.Wrap(err, "meta confirm transaction")
	}
	t := 2 * time.Minute
	log.Infof("deposit %f eth to arb (using hop) success, wait %v", amount, t)
	time.Sleep(t)
	return nil
}

func (a *Arbitrum) AddL2NetworkAndWaitTransaction() error {
	//l2Network := `//button[@class="mr-4 text-white hover:text-navy hover:bg-gray-200 cursor-pointer z-50 rounded-md text-sm font-medium"]`
	log.Infof("adding arbitrum l2 network")

	//chromedp.Run(a.ctx,
	//	utils.CdpPrint("111"),
	//	chromedp.Navigate(arbitrumBridgeUrl),
	//	utils.CdpPrint("222"),
	//	chromedp.Sleep(5*time.Second),
	//	utils.CdpPrint("333"),
	//
	//	chromedp.WaitReady(l2Network),
	//	utils.CdpPrint("444"),
	//
	//	chromedp.Click(l2Network),
	//	utils.CdpPrint("555"),
	//
	//	chromedp.Sleep(5*time.Second),
	//)
	if err := utils.OpenChanListAndAddNetwork(a.ctx, "Arbitrum One", a.meta); err != nil {
		return err
	}
	log.Infof("checking l2 balance")

	for {
		time.Sleep(3 * time.Second)
		balance, err := a.meta.Balance()
		if err != nil {
			return err
		}
		log.Infof("now l2 balance is %v", balance)
		if balance > 0 {
			log.Infof("transaction ends, now l2 balance is: %v", balance)
			break
		}
		time.Sleep(10 * time.Second) //wait transaction end
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
