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
const input = `//*[@id="root"]/div/main/div/div/div[4]/div[1]/div[1]/div/div[2]/div/input`
const depositButton = `//*[@id="root"]/div/main/div/div/div[4]/button`
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
	log.Infof("run arbitrum deposit")
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
	if err := utils.OpenChanListAndAddNetwork(a.ctx, "arbitrum", a.meta); err != nil {
		return err
	}

	for {
		time.Sleep(3 * time.Second)
		balance, err := a.meta.Balance()
		if err != nil {
			return err
		}
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
