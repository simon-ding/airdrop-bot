package ethclient

import (
	"airdrop-bot/log"
	"testing"
)

func TestZkClient(t *testing.T) {
	c := NewArbNovaClient()
	err := c.Connect()
	log.Infof("connect to zk: %v", err)

}
