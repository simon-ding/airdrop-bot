package ethclient

import "math/big"

type ChainHandelr interface {
	Connect() error
	GetEthBalance(addr string) (*big.Float, error)
	GetBalance(token Token, address string) (*big.Float, error)
	Tranfer(from, to string, value *big.Float) error
	BridgeUseOrbiter(privateKey string, value *big.Float, toChain Chain) error
}

var handlers = map[Chain]ChainHandelr{}

func GetHandler(c Chain) ChainHandelr {
	return handlers[c]
}

func regHandler(c Chain, handler ChainHandelr) {
	handlers[c] = handler
}

func init() {
	regHandler(ChainArbNova, NewArbNovaClient())
	regHandler(ChainArbOne, NewArbOneClient())
	regHandler(ChainEthMain, NewEthClient())
	regHandler(ChainZkEra, NewZkClient())
}