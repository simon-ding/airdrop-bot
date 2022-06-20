package binance

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"context"
	"fmt"
	"github.com/adshao/go-binance/v2"
	"github.com/pkg/errors"
)

func New(cfg *cfg.Binance) *Binance {
	client := binance.NewClient(cfg.ApiKey, cfg.SecretKey)

	return &Binance{client: client}
}

type Binance struct {
	client *binance.Client
}

func (b *Binance) WithdrawEth(addr string, amount float64) error {
	req := b.client.NewCreateWithdrawService()
	req.Coin("ETH")
	req.Address(addr)
	req.Amount(fmt.Sprintf("%v", amount))
	req.Network("Optimism")
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
