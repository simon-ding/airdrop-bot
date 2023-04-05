package main

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/ethclient"
	"airdrop-bot/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func NewServer(cfg *cfg.ServerConfig) (*Server, error) {
	r := gin.New()
	return &Server{
		cfg: cfg,
		r:   r,
	}, nil
}

type Server struct {
	cfg *cfg.ServerConfig

	r    *gin.Engine
}

func (s *Server) Serve() error {
	s.r.Use(func(c *gin.Context) {
		token := c.GetHeader(cfg.AuthHeader)
		if token != s.cfg.Token {
			//token is invalid
			log.Infof("access use invalid token: %v, ip %s", token, c.ClientIP())
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	})

	s.r.GET("/balance/:id", s.getBalance)
	return s.r.Run(":8080")
}

func (s *Server) getBalance(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	ac := db.FindAccount(idd)

	arbClient := ethclient.NewArbOneClient()
	ethCLient := ethclient.NewEthClient()
	arbClient.Connect()
	ethCLient.Connect()

	etheth, err := ethCLient.GetEthBalance(ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}
	ethUSDT, err := ethCLient.GetBalance(ethclient.TokenUSDT, ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	ethArb, err := ethCLient.GetBalance(ethclient.TokenArb, ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	arbEth, err := arbClient.GetEthBalance(ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	arbUSDT, err := arbClient.GetBalance(ethclient.TokenUSDT, ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	arbArb, err := arbClient.GetBalance(ethclient.TokenArb, ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	m := map[string]map[string]string {
		"ethereum": {
			"ETH": etheth.String(),
			"ARB": ethArb.String(),
			"USDT": ethUSDT.String(),
		},
		"arbitrum": {
			"ETH": arbEth.String(),
			"ARB": arbArb.String(),
			"USDT": arbUSDT.String(),
		},
	}
	c.JSON(200, m)
}

func (s *Server) FirstRunGenAccount() {
	num := s.cfg.AccountsToGen
	accountNum := db.AccountNum()
	if accountNum == num {
		log.Infof("accounts have already been generated,skip")
		return
	}
	log.Infof("try to generate %d accounts", num-accountNum)
	for i := 0; i < num-accountNum; i++ {
		mnemonic, address, priKey := utils.NewEthAccount()
		log.Infof("generate eth account: %s", address)
		db.SaveAccount(mnemonic, address, priKey)
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
