package server

import (
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/ethclient"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SwapEth2USDCInput struct {
	Ammount float64 `json:"ammount"`
	ID      int     `json:"id"`
	InToken string  `json:"token"`
}

func (s *Server) doMuteIoSwap(c *gin.Context) (interface{}, error) {
	var in SwapEth2USDCInput
	if err := c.ShouldBindJSON(&in); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}
	ac := db.FindAccount(in.ID)

	h := ethclient.NewZkClient()
	h.Connect()

	var txHash string
	if in.InToken == "" || in.InToken == ethclient.TokenEth.String() {
		hash, err := h.MuteIoSwap(ac.PrivateKey, ethclient.TokenEth, ethclient.TokenUSDC, in.Ammount)
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash
		log.Infof("swap eth to usdc success")

	} else if in.InToken == ethclient.TokenUSDC.String() {
		hash, err := h.MuteIoSwap(ac.PrivateKey, ethclient.TokenUSDC, ethclient.TokenEth, in.Ammount)
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash
		log.Infof("swap usdc to eth success")

	}

	return txHash, nil
}

func (s *Server) doSyncSwap(c *gin.Context) (interface{}, error) {
	var in SwapEth2USDCInput
	if err := c.ShouldBindJSON(&in); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}
	ac := db.FindAccount(in.ID)

	h := ethclient.NewZkClient()
	h.Connect()

	var txHash string

	if in.InToken == "" || in.InToken == ethclient.TokenEth.String() {
		hash, err := h.SyncSwapEth2Usdc(ac.PrivateKey, big.NewFloat(in.Ammount))
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash
		log.Infof("swap eth to usdc success")

	} else if in.InToken == ethclient.TokenUSDC.String() {
		hash, err := h.SyncSwapUsdc2Eth(ac.PrivateKey, big.NewFloat(in.Ammount))
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash
		log.Infof("swap usdc to eth success")

	}

	return txHash, nil

}

func (s *Server) getTransaction(c *gin.Context) (interface{}, error) {
	var apiUrl = "https://zksync2-mainnet-explorer.zksync.io/transactions?limit=50&direction=older&accountAddress="

	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	ac := db.FindAccount(idd)

	resp, err := http.Get(apiUrl + ac.Address)
	if err != nil {
		return nil, errors.Wrap(err, "http get")
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	return string(data), nil
}

type ZnsBuyDomain struct {
	Account int    `json:"account"`
	Domain  string `json:"domain"`
}

func (s *Server) ZnsBuyDomain(c *gin.Context) (interface{}, error) {
	var in ZnsBuyDomain
	if err := c.ShouldBindJSON(&in); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}
	ac := db.FindAccount(in.Account)

	h := ethclient.NewZkClient()
	h.Connect()

	tx, err := h.ZnsBuyDomain(in.Domain, ac.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "zns buy domain")
	}
	return tx, nil
}


func (s *Server) ZnsGetOwnedDomains(c *gin.Context) (interface{}, error) {
	id := c.Param("id")
	idd, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	ac := db.FindAccount(idd)

	h := ethclient.NewZkClient()
	h.Connect()

	domains, err := h.ZnsGetOwnedDomains(ac.Address)
	if err != nil {
		return nil, errors.Wrap(err, "get domains")
	}
	return domains, nil	
}