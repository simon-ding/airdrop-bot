package ethclient

import "fmt"

type Token int 
type Chain int

const (
	ChainEthMain Chain = iota
	ChainArbOne
	ChainArbNova
	ChainZkEra
)

const (
	TokenArb Token = iota
	TokenUSDT
	TokenEth
	TokenWETH
	TokenUSDC
	TokenDAI
)

func (t Token) String() string {
	switch t {
	case TokenArb:
		return "ARB"
	case TokenUSDT:
		return "USDT"
	case TokenEth:
		return "ETH"
	case TokenWETH:
		return "WETH"
	case TokenUSDC:
		return "USDC"
	case TokenDAI:
		return "DAI"
	}
	return "UNKNOWN"
}

func (c Chain) String() string {
	switch c {
	case ChainEthMain:
		return "Ethereum"
	case ChainArbOne:
		return "ArbOne"
	case ChainArbNova:
		return "ArbNova"
	case ChainZkEra:
		return "ZkEra"
	}
	return ""
}

const ZERO_ADDRESS = "0x0000000000000000000000000000000000000000"

func GetChain(name string) Chain {

	for  i := Chain(0);i < 100;i++ {
		if i.String() == name {
			return i
		}
	}
	panic(fmt.Sprintf("no chain named: %v", name))
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
		TokenWETH: "0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2",
		TokenUSDC: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
		TokenDAI: "0x6b175474e89094c44da98b954eedeac495271d0f",
	},
	ChainArbOne: {
		TokenUSDT: "0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9",
		TokenArb: "0x912CE59144191C1204E64559FE8253a0e49E6548",
	},
	ChainZkEra: {
		TokenUSDC: "0x3355df6D4c9C3035724Fd0e3914dE96A5a83aaf4",
		TokenWETH:"0x5aea5775959fbc2557cc8789bc1bf90a239d9a91",
		TokenDAI: "0x4BEf76b6b7f2823C6c1f4FcfEACD85C24548ad7e",
	},
}

func GetContractAddress(chain Chain, token Token) string {
	return contractAddress[chain][token]
}


func GetChainRpcEndpoint(chain Chain) string {
	return chainUrl[chain]
}

//https://docs.orbiter.finance/technology
var orbiterNetworkCode = map[Chain]int {
	ChainEthMain: 9001,
	ChainArbOne: 9002,
	ChainArbNova: 9016,
	ChainZkEra: 9014,
}

func GetOribiterNewworkCode(toChain Chain) int {
	return orbiterNetworkCode[toChain]
}