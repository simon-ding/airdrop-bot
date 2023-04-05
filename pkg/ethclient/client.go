package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"
	"airdrop-bot/utils"
	"context"
	"fmt"
	"math"
	"math/big"

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
	var rpcUrl = GetChainRpcEndpoint(c.chain)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("connect to endpoint error: %v", err)
	}
	c.client = client
	return nil
}

func (c *Client) GetEthBalance(addr string) (*big.Float, error) {
	account := common.HexToAddress(addr)
	balance, err := c.client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		return nil, fmt.Errorf("get balance: %v", err)
	}

	return utils.Wei2Eth(balance), nil
}

func (c *Client) GetBalance(token Token, address string) (*big.Float, error) {
	contractAddress := GetContractAddress(c.chain, token)
	tokenAddress := common.HexToAddress(contractAddress)
	d := GetTokenDecimal(token)

	abiClient, err := abi.NewErc20(tokenAddress, c.client)
	if err != nil {
		return nil, fmt.Errorf("new usdt: %v", err)
	}
	addr := common.HexToAddress(address)
	bal, err := abiClient.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return nil, fmt.Errorf("call abi: %v", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	v := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(d)))
	return v, nil
}

func (c *Client) GasPrice() *big.Float {
	gas, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("get gas price error: %v", err)
	}
	return utils.Wei2Eth(gas)
}
