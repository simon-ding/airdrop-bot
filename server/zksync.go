package server

import (
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/pkg/ethclient"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SwapEth2USDCInput struct {
	Ammount float64 `json:"ammount"`
	ID      int    `json:"id"`
	Token string `json:"token"`
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
	if in.Token == "" || in.Token == ethclient.TokenEth.String() {
		hash, err := h.MuteIoSwap(ac.PrivateKey, ethclient.TokenEth, ethclient.TokenUSDC, in.Ammount)
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash
	} else if in.Token == ethclient.TokenUSDC.String() {
		hash, err := h.MuteIoSwap(ac.PrivateKey, ethclient.TokenUSDC, ethclient.TokenEth, in.Ammount)
		if err != nil {
			return nil, errors.Wrap(err, "mute.io swap")
		}
		txHash = hash

	}

	log.Infof("swap eth to usdc success")
	return txHash, nil
}
