package utils

import (
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"math/rand"
	"time"
)

func GbkToUtf8(s []byte) []byte {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil
	}
	return d
}

func CdpPrint(s string) chromedp.ActionFunc {
	return chromedp.ActionFunc(func(ctx context.Context) error {
		fmt.Println(s)
		return nil
	})

}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func NewMnemonic(bitSize int) (string, error) {
	entropy, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func NewMnemonic128() (string, error) {
	return NewMnemonic(128)
}

func NewEthAccount() (string, string, string) {
	mnemonic, _ := hdwallet.NewMnemonic(128)
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, _ := wallet.Derive(path, false)
	PriKey, _ := wallet.PrivateKeyHex(account)
	return mnemonic, account.Address.Hex(), PriKey
}

func OpenChanListAndAddNetwork(ctx context.Context, networkName string, meta *metamask.Metamask) error {
	chanListUrl := `https://chainlist.org/zh`
	ctx, cancel := chromedp.NewContext(ctx)
	searchPos := `//*[@id="__next"]/div/main/div/div[2]/div[1]/div[1]/div/div/div/input`
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
	err = meta.ConfirmLinkAccount()
	if err != nil {
		return err
	}

	log.Infof("link meta to chanlist success")
	err = chromedp.Run(ctx,
		chromedp.SendKeys(searchPos, networkName),
		chromedp.Sleep(3*time.Second),
		chromedp.Click(addNetworkButton),
	)
	if err != nil {
		return err
	}
	return meta.ConfirmAddNetwork()
}
