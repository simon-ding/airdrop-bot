package server

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/binance"
	"airdrop-bot/pkg/ethclient"
	"airdrop-bot/utils"
	"fmt"
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

	r  *gin.Engine
	bn *binance.Binance
}

func (s *Server) Serve() error {
	s.initBinance()
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
	bot := api.Group("/bot")
	{
		bot.GET("/settings", HttpHandler(s.GetAllKeyValues))
		bot.POST("/settings", HttpHandler(s.SetKeyValue))
	}
	binance := api.Group("/binance")
	{
		binance.GET("/balance", HttpHandler(s.BinanceEthBalance))
		binance.POST("/transfer", HttpHandler(s.transferBinanceEth))
	}
	api.GET("/balance/:id", HttpHandler(s.getBalance))
	api.POST("/bridge/orbiter", HttpHandler(s.orbiterBridge))
	api.GET("/address/all", HttpHandler(s.getAllAccounts))
	api.POST("/address/gen/:num", HttpHandler(s.GenAccounts))
	api.POST("/app/muteio/swap/", HttpHandler(s.doMuteIoSwap))
	return s.r.Run(":8080")
}

func (s *Server) initBinance() {
	key, secret := db.GetBinanceKeySecret()
	client := binance.New(key, secret)
	s.bn = client
}

func (s *Server) getAllAccounts(c *gin.Context) (interface{}, error) {
	accounts := db.FetchAllAccounts()
	var m = make([]gin.H, 0, len(accounts))
	for _, a := range accounts {
		m = append(m, gin.H{
			"id":      a.ID,
			"address": a.Address,
		})
	}
	return m, nil
}

func (s *Server) getBalance(c *gin.Context) (interface{}, error) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}
	ac := db.FindAccount(idd)
	log.Infof("query account: %v	", ac.Address)

	chains := []ethclient.Chain{
		ethclient.ChainEthMain,
		ethclient.ChainArbOne,
		ethclient.ChainArbNova,
		ethclient.ChainZkEra,
	}

	var resp = gin.H{}
	resp["address"] = ac.Address

	for _, c := range chains {
		h := ethclient.GetHandler(c)

		if err := h.Connect(); err != nil {
			log.Errorf("connect to network: %v", err)
			continue
		}
		name := h.Name()
		m := h.AllTokenBalances(ac.Address)
		for k, v := range m {
			if p, err := s.bn.Price(k); err != nil {
				log.Errorf("get binance price error %v: %v", k, err)
			} else {
				f, _, err := big.NewFloat(0).Parse(v, 10)
				if err != nil {
					log.Errorf("parse float: %v", err)
				} else {
					f = f.Mul(f, p)
					m[k] = fmt.Sprintf("%v (%vUSDT)", v, f.Text('g', 4))
				}
			}
		}

		if len(m) == 0 {
			continue
		}
		resp[name] = m
	}
	return resp, nil
}

type orbiterInput struct {
	ToChain   string  `json:"to_chain"`
	FromChain string  `json:"from_chain"`
	Account   int     `json:"account"`
	Amount    float64 `json:"amount"`
}

func (s *Server) orbiterBridge(c *gin.Context) (interface{}, error) {
	var in orbiterInput
	c.ShouldBindJSON(&in)
	if err := s.doOrbiterBridge(in); err != nil {
		return nil, fmt.Errorf("orbiter bridge error: %v", err)
	}
	return nil, nil
}

func (s *Server) doOrbiterBridge(in orbiterInput) error {
	toChain := ethclient.GetChain(in.ToChain)
	fromChain := ethclient.GetChain(in.FromChain)
	ac := db.FindAccount(in.Account)
	log.Infof("query account: %v", ac.Address)

	hander := ethclient.GetHandler(fromChain)
	if err := hander.Connect(); err != nil {
		return err
	}

	err := hander.BridgeUseOrbiter(ac.PrivateKey, big.NewFloat(in.Amount), toChain)
	return err
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

func (s *Server) claimAidoge() {

}
