package server

import (
	"airdrop-bot/db"
	"airdrop-bot/pkg/ethclient"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type SwapEth2USDCInput struct {
	Ammount float64 `json:"ammount"`
	ID      int    `json:"id"`
}

func (s *Server) SwapEth2USDC(c *gin.Context) (interface{}, error) {
	var in SwapEth2USDCInput
	if err := c.ShouldBindJSON(&in); err != nil {
		return nil, errors.Wrap(err, "bind json")
	}
	ac := db.FindAccount(in.ID)

	h := ethclient.NewZkClient()
	h.Connect()

	err := h.MuteIoSwap(ac.PrivateKey, ethclient.TokenEth, ethclient.TokenUSDC, in.Ammount)
	if err != nil {
		return nil, errors.Wrap(err, "mute.io swap")
	}
	return nil, nil
}
