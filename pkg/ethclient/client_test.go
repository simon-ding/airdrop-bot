package ethclient

import (
	"log"
	"testing"
)

func TestGetbalance(t *testing.T) {
	c := NewClient(ChainArbOne)
	err := c.Connect()
	log.Print(err)
	bal, err := c.GetBalance(TokenArb,"0x439f80d25baD0922B9636b4C534a9fDD2f6E8C11")
	log.Printf("current USDT balance is %v, err: %v", bal, err)
	bal1, err := c.GetEthBalance("0x439f80d25baD0922B9636b4C534a9fDD2f6E8C11")
	log.Printf("current Eth balance is %v, err: %v", bal1, err)
	p := c.GasPrice()
	log.Printf("gas price: %v", p)
}
