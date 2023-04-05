package ethclient

func NewEthClient() *ArbOneClient {
	return &ArbOneClient{Client: NewClient(ChainEthMain)}
}

type EthClient struct {
	*Client
}
