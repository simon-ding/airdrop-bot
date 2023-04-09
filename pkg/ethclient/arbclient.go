package ethclient

func NewArbOneClient() *ArbOneClient {
	return &ArbOneClient{Client: NewClient(ChainArbOne)}
}

type ArbOneClient struct {
	*Client
}

func (a *ArbOneClient) TokenSum() {
	// addr := common.HexToAddress("0x67a24CE4321aB3aF51c2D0a4801c3E111D88C9d9")
	// instance, err := abi.NewArbToken(addr, a.client)
	// if err != nil {
	// 	log.Errorf("err: ", err)
	// }
	// instance.ParseCanClaim(nil)
}


func NewArbNovaClient() *ArbNovaClient {
	return &ArbNovaClient{Client: NewClient(ChainArbNova)}
}

type ArbNovaClient struct {
	*Client
}
