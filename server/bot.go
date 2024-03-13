package server

import (
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/binance"
	"airdrop-bot/pkg/ethclient"
	"airdrop-bot/utils"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *Server) BinanceEthBalance(c *gin.Context) (interface{}, error) {
	bn := s.db.GetBinanceSetting()

	client := binance.New(bn.Key, bn.Secret)
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
	ac := s.db.GetAccount(req.AccountID)
	log.Infof("query account: %v	", ac.Address)
	if ac.Address == "" {
		return nil, errors.Errorf("no such account: %v", ac.ID)
	}

	err := s.bn.WithdrawEth(req.Network, ac.Address, req.Ammount)

	return nil, errors.Wrap(err, "binance withdraw")
}

func (s *Server) GenAccounts(c *gin.Context) (interface{}, error) {
	n := c.Param("num")
	num, err := strconv.Atoi(n)
	if err != nil {
		return nil, err
	}
	s.genAccounts(num)
	return nil, nil
}

func (s *Server) genAccounts(num int) {
	accountNum := len(s.db.GetAllAccounts())
	if accountNum == num {
		log.Infof("accounts have already been generated,skip")
		return
	}
	log.Infof("try to generate %d accounts", num-accountNum)
	for i := 0; i < num-accountNum; i++ {
		mnemonic, address, priKey := utils.NewEthAccount()
		log.Infof("generate eth account: %s", address)
		s.db.AddAccount(db.Account{
			Mnemonic:   mnemonic,
			Address:    address,
			PrivateKey: priKey,
		})
	}
}

func (s *Server) isGasFeeAcceptable() bool {
	fee, err := utils.GetGasFee(nil)
	if err != nil {
		log.Errorf("get gas fee error: %v", err)
		return false
	}
	if len(fee.Speeds) == 0 {
		log.Infof("api return no fee: %+v", fee)
		return false
	}
	log.Infof("current gas fee is %+v", fee.Speeds)

	if fee.Speeds[1].EstimatedFee > 7 {
		log.Errorf("gas fee too high: %v", fee.Speeds[1].EstimatedFee)
		return false
	}
	return true
}

func (s *Server) getGasPrices(c *gin.Context) (interface{}, error) {
	c1 := ethclient.NewEthClient()
	c1.Connect()
	p1 := c1.GasPrice()

	c2 := ethclient.NewArbOneClient()
	c2.Connect()
	p2 := c2.GasPrice()

	c3 := ethclient.NewZkClient()
	c3.Connect()
	p3 := c3.GasPrice()

	// c4 := ethclient.NewArbNovaClient()
	// c4.Connect()
	// p4 := c4.GasPrice()

	return map[string]string{
		ethclient.ChainEthMain.String(): p1.Text('g', 3),
		ethclient.ChainArbOne.String():  p2.Text('g', 3),
		ethclient.ChainZkEra.String():   p3.Text('g', 3),
		//ethclient.ChainArbNova.String(): p4.Text('g', 3),
	}, nil
}
