package ethclient

type Token int 
type Chain int

const (
	ChainEthMain Chain = iota
	ChainArbOne
	ChainArbNova
)

const (
	TokenArb Token = iota
	TokenUSDT
	TokenEth
)

func (t Token) String() string {
	switch t {
	case TokenArb:
		return "ARB"
	case TokenUSDT:
		return "USDT"
	case TokenEth:
		return "ETH"
	}
	return ""
}

var chainUrl = map[Chain]string {
	ChainEthMain: "https://eth.llamarpc.com",
	ChainArbNova:"https://nova.arbitrum.io/rpc",
	ChainArbOne: "https://arb1.arbitrum.io/rpc",
}

var contractAddress = map[Chain]map[Token]string {
	ChainEthMain: {
		TokenUSDT: "0xdAC17F958D2ee523a2206206994597C13D831ec7",
		TokenArb: "0xB50721BCf8d664c30412Cfbc6cf7a15145234ad1",
	},
	ChainArbOne: {
		TokenUSDT: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
		TokenArb: "0x912CE59144191C1204E64559FE8253a0e49E6548",
	},
}

func GetContractAddress(chain Chain, token Token) string {
	return contractAddress[chain][token]
}


func GetChainRpcEndpoint(chain Chain) string {
	return chainUrl[chain]
}