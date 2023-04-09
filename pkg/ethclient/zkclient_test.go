package ethclient

import (
	"airdrop-bot/log"
	"testing"
)

func TestZkClient(t *testing.T) {
 c := NewZkClient()
 err := c.Connect()
 log.Infof("connect to zk: %v", err)

 b, err := c.GetEthBalance("0xA8B6315741d11C70fbd1097263350A2a9836b026")
 log.Infof("get eth balance: %v, error: %v", b, err)
}
