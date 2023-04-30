package server

import (
	"airdrop-bot/db"
	"airdrop-bot/ent/settings"
	"airdrop-bot/log"
	"airdrop-bot/pkg/binance"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Server) GetAllKeyValues(c *gin.Context) (interface{}, error) {
	settings := db.GetClient().Settings.Query().AllX(context.Background())

	return settings, nil
}

type setKeyValueRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func (s *Server) SetKeyValue(c *gin.Context) (interface{}, error) {
	var req setKeyValueRequest

	if err := c.BindJSON(&req); err != nil {
		return nil, err
	}
	if req.Key == "" || req.Value == "" {
		return nil, errors.Errorf("key or value cannot be empty")
	}

	kv, err := db.GetClient().Settings.Query().Where(settings.Key(req.Key)).First(context.TODO())
	if err != nil {
		db.GetClient().Settings.Create().SetKey(req.Key).SetValue(req.Value).SaveX(context.Background())
		return nil, errors.Wrap(err, "get settings")
	}
	log.Infof("exist key value: %v", kv)
	db.GetClient().Settings.UpdateOneID(kv.ID).SetValue(req.Value).SaveX(context.Background())
	return nil, nil
}

func (s *Server) BinanceEthBalance(c *gin.Context) (interface{}, error) {
	key, secret := db.GetBinanceKeySecret()

	client := binance.New(key, secret)
	b, err := client.EthBalance()

	return b, err
}

type transferBinanceEthInput struct {
	Network   string  `json:"network"`
	AccountID int     `json:"account_id"`
	Ammount   float64 `json:"ammount"`
}

func (s *Server) transferBinanceEth(c *gin.Context) (interface{}, error) {
	var req transferBinanceEthInput
	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}
	if req.Network != "arbitrum" {
		return nil, errors.Errorf("not supported network: %v", req.Network)
	}
	ac := db.FindAccount(req.AccountID)
	log.Infof("query account: %v	", ac.Address)
	if ac.Address == "" {
		return nil, errors.Errorf("no such account: %v", ac.ID)
	}
	
	err := s.bn.WithdrawEth(req.Network, ac.Address, req.Ammount)

	return nil, errors.Wrap(err, "binance withdraw")
}
