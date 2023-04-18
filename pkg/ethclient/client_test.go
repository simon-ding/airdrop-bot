package ethclient

import (
	"airdrop-bot/log"
	"testing"
)

func TestArb(t *testing.T) {
	a := NewArbOneClient()
	a.Connect()
	err := a.ClaimAidoge("")
	log.Infof("error : %v", err)
}
