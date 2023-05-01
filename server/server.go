package server

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/binance"
	"airdrop-bot/pkg/ethclient"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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
	app := api.Group("/app")
	{
		app.POST("/muteio/swap/", HttpHandler(s.doMuteIoSwap))
		app.POST("/cbridge/send", HttpHandler(s.cbridgeSend))
	}
	api.GET("/balance/:id", HttpHandler(s.getBalance))
	api.POST("/bridge/orbiter", HttpHandler(s.orbiterBridge))
	api.GET("/address/all", HttpHandler(s.getAllAccounts))
	api.POST("/address/gen/:num", HttpHandler(s.GenAccounts))
	
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

type cbridgeSendInput struct {
	DstChain string `json:"dst_chain"`
	SrcChain string `json:"src_chain"`
	Ammount float64 `json:"ammount"`
	AccountID int `json:"account_id"`
}

func (s *Server) cbridgeSend(c *gin.Context) (interface{}, error) {
	var req cbridgeSendInput
	if err := c.BindJSON(&req); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}

	ac := db.FindAccount(req.AccountID)
	log.Infof("query account: %v	", ac.Address)
	if ac.Address == "" {
		return nil, errors.Errorf("no such account: %v", ac.ID)
	}

	src := ethclient.GetChain(req.SrcChain)
	dst := ethclient.GetChain(req.DstChain)
	
	h := ethclient.GetHandler(src)
	tx, err := h.CBridgeSend(dst, ac.PrivateKey, req.Ammount)
	if err != nil {
		return nil, errors.Wrap(err, "cbridge send")
	}
	return tx, nil
}