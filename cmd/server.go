package main

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/ethclient"
	"airdrop-bot/utils"
	"math/big"
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

	r *gin.Engine
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
	api := s.r.Group("/api/v1")

	api.GET("/balance/:id", s.getBalance)
	api.POST("/bridge/orbiter", s.orbiterBridge)
	return s.r.Run(":8080")
}

func (s *Server) getBalance(c *gin.Context) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return
	}
	ac := db.FindAccount(idd)
	log.Infof("query account: %v", ac.Address)

	arbClient := ethclient.NewArbOneClient()
	ethCLient := ethclient.NewEthClient()
	nova := ethclient.NewArbNovaClient()
	zkera := ethclient.NewZkClient()
	arbClient.Connect()
	ethCLient.Connect()
	nova.Connect()
	zkera.Connect()

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

	novaEth, err := nova.GetEthBalance(ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}
	zkEth, err := zkera.GetEthBalance(ac.Address)
	if err != nil {
		log.Errorf("get balance: %v", err)
	}

	c.JSON(200, gin.H{
		"ethereum": map[string]string{
			"ETH":  etheth.String(),
			"ARB":  ethArb.String(),
			"USDT": ethUSDT.String(),
		},
		"arbitrumOne": map[string]string{
			"ETH":  arbEth.String(),
			"ARB":  arbArb.String(),
			"USDT": arbUSDT.String(),
		},
		"abitrumNova": map[string]string{
			"ETH": novaEth.String(),
		},
		"zksync": map[string]string{
			"ETH": zkEth.String(),
		},
		"address": ac.Address,
	})
}

type orbiterInput struct {
	ToChain   string
	FromChain string
	Account   int
	Amount    float64
}

func (s *Server) orbiterBridge(c *gin.Context) {
	var in orbiterInput
	c.ShouldBindJSON(&in)
	if err := s.doOrbiterBridge(in); err != nil {
		log.Errorf("orbiter bridge error: %v", err)
		c.JSON(501, "error")
	} else {
		c.JSON(200, "success")
	}

}

func (s *Server) doOrbiterBridge(in orbiterInput) error {
	toChain := ethclient.GetChain(in.ToChain)
	fromChain := ethclient.GetChain(in.FromChain)
	ac := db.FindAccount(in.Account)
	log.Infof("query account: %v", ac.Address)

	hander := ethclient.GetHandler(fromChain)

	err := hander.BridgeUseOrbiter(ac.PrivateKey, big.NewFloat(in.Amount), toChain)
	return err
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
