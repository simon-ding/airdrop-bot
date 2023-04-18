package utils

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/metamask"
	"archive/zip"
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
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
	searchPos := `//input[@placeholder="ETH, Fantom, ..."]`
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

func Unzip(src []byte, dest string) error {
	log.Infof("try to unzip file to dir: %s", dest)
	r, err := zip.NewReader(bytes.NewReader(src), int64(len(src)))
	if err != nil {
		return err
	}
	os.RemoveAll(dest)

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(dest, f.Name)

		// Check for ZipSlip (Directory traversal)
		if !strings.HasPrefix(path, filepath.Clean(dest)+string(os.PathSeparator)) {
			return fmt.Errorf("illegal file path: %s", path)
		}

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}
	log.Infof("unzip file success")
	return nil
}

func Wei2Eth(wei *big.Int) *big.Float {
	fbalance := new(big.Float)
	fbalance.SetString(wei.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(params.Ether))
	return ethValue
}

func Eth2Wei(eth *big.Float) *big.Int {
	coin := new(big.Float)
	coin.SetInt(big.NewInt(params.Ether))

	eth.Mul(eth, coin)

	wei := new(big.Int)
	eth.Int(wei) // store converted number in result

	return wei
}

func GetPublicKey(privateKey string) string {
	pKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Errorf("encode private key: %v", err)
		return ""
	}
	publicKey := pKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return ""
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return address
}

func SignMsg(msg string, priKey string) string {
	privateKey, err := crypto.HexToECDSA(priKey)
	if err != nil {
		panic(err)
	}
	data := []byte(msg)
	hash := crypto.Keccak256Hash(data)
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		panic(err)
	}
	return hexutil.Encode(signature)
}
