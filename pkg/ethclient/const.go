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

const (
	ArbOneUrl = "https://arb1.arbitrum.io/rpc"
	ArbNovaUrl = "https://nova.arbitrum.io/rpc"
	EthMainUrl = "https://eth.llamarpc.com"
)


const (
	ContractAddressArbUsdt = "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9"
	ContractAddressArbArb = "0x912CE59144191C1204E64559FE8253a0e49E6548"
)