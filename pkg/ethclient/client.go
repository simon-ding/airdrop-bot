package ethclient

import (
	"airdrop-bot/pkg/abi"
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func NewClient(chain Chain) *Client {
	return &Client{chain: chain}
}

type Client struct {
	chain  Chain
	client *ethclient.Client
}

func (c *Client) Connect() error {
	var rpcUrl = ""
	switch c.chain {
	case ChainEthMain:
		rpcUrl = EthMainUrl
	case ChainArbOne:
		rpcUrl = ArbOneUrl
	case ChainArbNova:
		rpcUrl = ArbNovaUrl
	default:
		rpcUrl = EthMainUrl
	}

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("connect to endpoint error: %v", err)
	}
	c.client = client
	return nil
}

func (c *Client) GetBalance(token Token) int {
	return 0
}

func (c *Client) getArbUsdtBalance(address string) (int64, error) {
	contractAddress := ContractAddressArbUsdt
    tokenAddress := common.HexToAddress(contractAddress)

	abiClient, err := abi.NewErc20(tokenAddress, c.client)
	if err != nil {
		return 0, fmt.Errorf("new usdt: %v", err)
	}
	addr := common.HexToAddress(address)
	bal, err := abiClient.BalanceOf(&bind.CallOpts{},addr)
	if err != nil {
		return 0, fmt.Errorf("call abi: %v", err)
	}
	return bal.Int64(), nil
}

func (c *Client) getContractAddress(token Token) string {

	if c.chain == ChainArbOne {
		switch token {
		case TokenUSDT:
			return ContractAddressArbUsdt
		}
	}
	return ""
}
