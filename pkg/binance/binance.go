package binance

import (
	"airdrop-bot/log"
	"context"
	"fmt"
	"math/big"

	"github.com/adshao/go-binance/v2"
	"github.com/pkg/errors"
)

func New(key, secret string) *Binance {
	client := binance.NewClient(key, secret)

	return &Binance{client: client}
}

type Binance struct {
	client *binance.Client
}

func (b *Binance) WithdrawEth(network, addr string, amount float64) error {
	req := b.client.NewCreateWithdrawService()
	req.Coin("ETH")
	req.Address(addr)
	req.Amount(fmt.Sprintf("%v", amount))
	req.Network(network)
	if _, err := req.Do(context.Background()); err != nil {
		return errors.Wrap(err, "with draw")
	}
	log.Infof("withdraw %v eth to address %v", amount, addr)
	return nil
}

func (b *Binance) EthBalance() (string, error) {
	svc := b.client.NewGetAccountService()
	resp, err := svc.Do(context.TODO())
	if err != nil {
		return "", errors.Wrap(err, "get account")
	}
	var balance string
	for _, a := range resp.Balances {
		if a.Asset == "ETH" {
			balance = a.Free
		}
	}
	return balance, nil
}


func (b *Binance) Price(symbol string) (*big.Float, error) {
	svc := b.client.NewAveragePriceService()
	svc.Symbol(symbol + "USDT")
	resp, err := svc.Do(context.TODO())
	if err != nil {
		return nil, errors.Wrap(err, "svc do")
	}
	f, _, err := big.NewFloat(0).Parse(resp.Price, 10)
	if err != nil {
		return nil, errors.Wrap(err, "parse float")
	}
	return f, nil
}

func (b *Binance) GasPrice(gwei int) (string, error) {
	return "", nil
}