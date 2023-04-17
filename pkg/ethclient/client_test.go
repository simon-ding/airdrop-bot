package ethclient

import (
	"airdrop-bot/log"
	"testing"
)

func TestArb(t *testing.T) {
	a := NewArbNovaClient()
	err := a.SyncSwap()
	log.Infof("error : %v", err)
}
