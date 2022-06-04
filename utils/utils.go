package utils

import (
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
