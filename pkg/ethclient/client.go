package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"
	"airdrop-bot/utils"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

	abiClient, err := abi.NewErc20(tokenAddress, c.client)
	if err != nil {
		return nil, fmt.Errorf("new usdt: %v", err)
	}
	addr := common.HexToAddress(address)
	bal, err := abiClient.BalanceOf(&bind.CallOpts{}, addr)
	if err != nil {
		return nil, fmt.Errorf("call abi: %v", err)
	}
	d, err := abiClient.Decimals(&bind.CallOpts{})
	if err != nil {
		return nil, fmt.Errorf("get decimals: %v", err)
	}

	fbalance := new(big.Float)
	fbalance.SetString(bal.String())
	v := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(d))))
	return v, nil
}

func (c *Client) GasPrice() *big.Float {
	gas, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Errorf("get gas price error: %v", err)
	}
	return utils.Wei2Eth(gas)
}

func (c *Client) Tranfer(from, to string, value *big.Float) error {
	log.Infof("transfer %v eth to %v", value, to)
	privateKey, err := crypto.HexToECDSA(from)
	if err != nil {
		return fmt.Errorf("encode private key: %v", err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return fmt.Errorf("get nonce: %v", err)
	}

	weiValue := utils.Eth2Wei(value)

	gasLimit := uint64(1001000) //recommended gas limit 21000 didn't work, this works...
	toAddress := common.HexToAddress(to)

	// tx := types.NewTransaction(nonce, toAddress, weiValue, gasLimit, gasPrice, nil)
	// Get suggested gas price
	tipCap, err := c.client.SuggestGasTipCap(context.Background())
	if err != nil {
		tipCap = big.NewInt(0)
		//return fmt.Errorf("get gas tip cap: %v", err)
	}
	feeCap, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("get gas price: %v", err)
	}

	log.Infof("gas tip cap: %v, gas price: %v", tipCap, feeCap)
	chainID, err := c.client.NetworkID(context.Background())
	if err != nil {
		return fmt.Errorf("get chain id: %v", err)
	}
	feeTx := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: tipCap,
		GasFeeCap: feeCap,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     weiValue,
		Data:      nil,
	}
	log.Infof("new tx: %+v", feeTx)

	tx := types.NewTx(feeTx)

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		return fmt.Errorf("sign tx: %v", err)
	}

	err = c.client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return fmt.Errorf("send transaction: %v", err)
	}
	return nil
}


func(c *Client) BridgeUseOrbiter(privateKey string, value *big.Float, toChain Chain) error {

	// min 0.005eth
	if value.Cmp(big.NewFloat(0.006)) <= 0 {
		return fmt.Errorf("less then min transfer amount")
	}

	log.Infof("orbiter bridge from %v to %v, amount: %v", c.chain, toChain, value)

	code := GetOribiterNewworkCode(toChain)
	if code == 0 {
		panic(fmt.Sprintf("no to chain code: %v", toChain))
	}
	wei := utils.Eth2Wei(value)

	wei.Add(wei, big.NewInt(int64(code)))

	orbiterAddress := "0x80c67432656d59144ceff962e8faf8926599bcf8"
	
	return c.Tranfer(privateKey, orbiterAddress, utils.Wei2Eth(wei))
}

func (c *Client) Name() string {
	return c.chain.String()
}