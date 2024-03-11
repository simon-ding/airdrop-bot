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
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/pkg/errors"
)

func NewClient(chain Chain) *Client {
	return &Client{chain: chain}
}

type Client struct {
	chain  Chain
	client EthereumClient
}

func (c *Client) Connect() error {
	var rpcUrl = GetChainRpcEndpoint(c.chain)

	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return fmt.Errorf("connect to endpoint error: %v", err)
	}
	c.client = client
	log.Infof("connect to network success %v", c.chain)
	return nil
}

func (c *Client) ChainID() (*big.Int, error) {
	return c.client.ChainID(context.Background())
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
	f := big.NewFloat(0).SetInt(gas)
	return f.Quo(f, big.NewFloat(params.GWei))
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

func (c *Client) BridgeUseOrbiter(privateKey string, value *big.Float, toChain Chain) error {

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

func (c *Client) ConvertToken(t Token, ammount *big.Float) *big.Int {
	addr := GetContractAddress(c.chain, t)
	erc20, err := abi.NewErc20(common.HexToAddress(addr), c.client)
	if err != nil {
		log.Errorf("converting token: %v", err)
		return nil
	}
	d, err := erc20.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Errorf("get decimals: %v", err)
		return nil
	}
	a := ammount
	b := new(big.Float)
	b.SetInt(big.NewInt(int64(math.Pow10(int(d)))))
	a.Mul(a, b)

	r := new(big.Int)
	a.Int(r)
	return r
}

func (c *Client) GetTransactor(privateKey string) (*bind.TransactOpts, error) {
	if c.chain == ChainZkEra {
		return c.getZkTransactor(privateKey)
	}
	// publicKey := utils.GetPublicKeyFromPrivateKey(privateKey)
	// nonce, err := c.client.PendingNonceAt(context.Background(), common.HexToAddress(publicKey))
	// if err != nil {
	// 	return nil, fmt.Errorf("get nonce: %v", err)
	// }

	// gasPrice, err := c.client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	// tipCap, err := c.client.SuggestGasTipCap(context.Background())
	// if err != nil {
	// 	return nil, err
	// }
	pkey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	chainId, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	//log.Infof("chain id: %v", chainId)
	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainId)
	if err != nil {
		return nil, err
	}
	// auth.Nonce = big.NewInt(10)
	// auth.Value = big.NewInt(0)       // in wei
	// auth.GasLimit = uint64(10000000) // in units
	// auth.GasPrice = gasPrice
	// auth.GasTipCap = tipCap
	//log.Infof("before transaction: %+v", auth)

	return auth, nil
}

func (c *Client) getZkTransactor(privateKey string) (*bind.TransactOpts, error) {
	gasPrice, err := c.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	pkey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, err
	}
	chainId, err := c.client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	//log.Infof("chain id: %v", chainId)
	auth, err := bind.NewKeyedTransactorWithChainID(pkey, chainId)
	if err != nil {
		return nil, err
	}
	// auth.Nonce = big.NewInt(10)
	// auth.Value = big.NewInt(0)       // in wei
	// auth.GasLimit = uint64(10000000) // in units
	auth.GasPrice = gasPrice
	//log.Infof("before transaction: %+v", auth)

	return auth, nil
}

func (c *Client) ApproveTokenAllowance(t Token, ownerPrivateKey, spender string) error {
	contractAddr := GetContractAddress(c.chain, t)
	erc20, err := abi.NewErc20(common.HexToAddress(contractAddr), c.client)
	if err != nil {
		return nil
	}
	pubKey := common.HexToAddress(utils.GetPublicKey(ownerPrivateKey))
	spenderAddr := common.HexToAddress(spender)
	al, err := erc20.Allowance(&bind.CallOpts{}, pubKey, spenderAddr)
	if err != nil {
		return err
	}
	bal1, err := erc20.BalanceOf(&bind.CallOpts{}, pubKey)
	if err != nil {
		return errors.Wrap(err, "get balance")
	}

	if al.Int64() >= bal1.Int64() {
		log.Infof("already set allowance to max: %v", al)
		return nil
	}

	auth, err := c.GetTransactor(ownerPrivateKey)
	if err != nil {
		return err
	}

	bal, err := erc20.BalanceOf(&bind.CallOpts{}, pubKey)
	if err != nil {
		return errors.Wrap(err, "erc20 balance")
	}

	tx, err := erc20.Approve(auth, spenderAddr, bal)
	if err != nil {
		return errors.Wrap(err, "erc20 approve")
	}
	log.Infof("approve all token transaction: %v", tx.Hash().Hex())
	return nil
}

func (c *Client) AllTokenBalances(address string) map[string]string {
	tokens := []Token{
		TokenArb,
		TokenUSDT,
		TokenEth,
		TokenWETH,
		TokenUSDC,
		TokenDAI,
	}

	var res = make(map[string]string)
	for _, t := range tokens {
		addr := GetContractAddress(c.chain, t)
		if addr == "" {
			continue
		}
		bal, err := c.GetBalance(t, address)
		if err != nil {
			log.Errorf("get token %v address %v balance: %v", t, address, err)
			continue
		}
		if bal.String() == "0" {
			continue
		}
		res[t.String()] = bal.String()
	}
	ethBal, err := c.GetEthBalance(address)
	if err == nil && ethBal.String() != "0" {
		res[TokenEth.String()] = ethBal.String()
	}
	return res
}

func (c *Client) CBridgeSend(dst Chain, privateKey string, ammount float64) (string, error) {
	if ammount <= 0.005 {
		return "", errors.Errorf("eth should be greater than 0.005")
	}
	var contractAddresses = map[Chain]string{
		ChainEthMain: "0x5427FEFA711Eff984124bFBB1AB6fbf5E3DA1820",
		ChainArbOne:  "0x1619DE6B6B20eD217a58d00f37B9d47C7663feca",
		ChainArbNova: "0xb3833Ecd19D4Ff964fA7bc3f8aC070ad5e360E56",
		ChainZkEra:   "0x54069e96C4247b37C2fbd9559CA99f08CD1CD66c",
	}
	addr := contractAddresses[c.chain]
	if addr == "" {
		return "", errors.Errorf("no contract address on chain: %v", c.chain)
	}
	cbridge, err := abi.NewCbridge(common.HexToAddress(addr), c.client)
	if err != nil {
		return "", errors.Wrap(err, "new cbridge")
	}

	auth, err := c.GetTransactor(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "get auth")
	}
	wei := utils.Eth2Wei(big.NewFloat(ammount))
	auth.Value = wei

	pubKey := utils.GetPublicKey(privateKey)

	h := GetHandler(dst)
	h.Connect()
	id, err := h.ChainID()
	if err != nil {
		return "", errors.Wrap(err, "chain id")
	}
	tx, err := cbridge.SendNative(auth, common.HexToAddress(pubKey), wei, id.Uint64(), uint64(time.Now().UnixNano()), 259946)
	if err != nil {
		return "", errors.Wrap(err, "send native")
	}
	return tx.Hash().Hex(), nil
}

func (c *Client) EstimatedGas() {
	addr := common.HexToAddress("0x0000000000000000000000000000000000000000")

	f, err := c.client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: addr,
		To:   &addr,
	})
	log.Infof("estimated: %v, %v", f, err)
}


type EthereumClient interface {
	// Close closes the underlying RPC connection.
	Close()

	// ChainID retrieves the current chain ID for transaction replay protection.
	ChainID(ctx context.Context) (*big.Int, error)
	// BlockNumber returns the most recent block number
	BlockNumber(ctx context.Context) (uint64, error)
	// PeerCount returns the number of p2p peers as reported by the net_peerCount method
	PeerCount(ctx context.Context) (uint64, error)
	// HeaderByHash returns the block header with the given hash.
	HeaderByHash(ctx context.Context, hash common.Hash) (*types.Header, error)
	// HeaderByNumber returns a block header from the current canonical chain. If number is
	// nil, the latest known header is returned.
	HeaderByNumber(ctx context.Context, number *big.Int) (*types.Header, error)
	// TransactionCount returns the total number of transactions in the given block.
	TransactionCount(ctx context.Context, blockHash common.Hash) (uint, error)
	// SyncProgress retrieves the current progress of the sync algorithm. If there's
	// no sync currently running, it returns nil.
	SyncProgress(ctx context.Context) (*ethereum.SyncProgress, error)
	// SubscribeNewHead subscribes to notifications about the current blockchain head
	// on the given channel.
	SubscribeNewHead(ctx context.Context, ch chan<- *types.Header) (ethereum.Subscription, error)

	// NetworkID returns the network ID for this client.
	NetworkID(ctx context.Context) (*big.Int, error)
	// BalanceAt returns the wei balance of the given account.
	// The block number can be nil, in which case the balance is taken from the latest known block.
	BalanceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error)
	// StorageAt returns the value of key in the contract storage of the given account.
	// The block number can be nil, in which case the value is taken from the latest known block.
	StorageAt(ctx context.Context, account common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error)
	// CodeAt returns the contract code of the given account.
	// The block number can be nil, in which case the code is taken from the latest known block.
	CodeAt(ctx context.Context, account common.Address, blockNumber *big.Int) ([]byte, error)
	// NonceAt returns the account nonce of the given account.
	// The block number can be nil, in which case the nonce is taken from the latest known block.
	NonceAt(ctx context.Context, account common.Address, blockNumber *big.Int) (uint64, error)

	// FilterLogs performs the same function as FilterLogsL2, and that method should be used instead.
	// This method is designed to be compatible with bind.ContractBackend.
	FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error)
	// SubscribeFilterLogs performs the same function as SubscribeFilterLogsL2, and that method should be used instead.
	// This method is designed to be compatible with bind.ContractBackend.
	SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error)

	// PendingBalanceAt returns the wei balance of the given account in the pending state.
	PendingBalanceAt(ctx context.Context, account common.Address) (*big.Int, error)
	// PendingStorageAt returns the value of key in the contract storage of the given account in the pending state.
	PendingStorageAt(ctx context.Context, account common.Address, key common.Hash) ([]byte, error)
	// PendingCodeAt returns the contract code of the given account in the pending state.
	PendingCodeAt(ctx context.Context, account common.Address) ([]byte, error)
	// PendingNonceAt returns the account nonce of the given account in the pending state.
	// This is the nonce that should be used for the next transaction.
	PendingNonceAt(ctx context.Context, account common.Address) (uint64, error)
	// PendingTransactionCount returns the total number of transactions in the pending state.
	PendingTransactionCount(ctx context.Context) (uint, error)

	// CallContract executes a message call transaction, which is directly executed in the VM
	// of the node, but never mined into the blockchain.
	//
	// blockNumber selects the block height at which the call runs. It can be nil, in which
	// case the code is taken from the latest known block. Note that state from very old
	// blocks might not be available.
	CallContract(ctx context.Context, msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	// SuggestGasPrice retrieves the currently suggested gas price to allow a timely
	// execution of a transaction.
	SuggestGasPrice(ctx context.Context) (*big.Int, error)
	// SuggestGasTipCap retrieves the currently suggested gas tip cap after 1559 to
	// allow a timely execution of a transaction.
	SuggestGasTipCap(ctx context.Context) (*big.Int, error)
	// EstimateGas tries to estimate the gas needed to execute a transaction based on
	// the current pending state of the backend blockchain. There is no guarantee that this is
	// the true gas limit requirement as other transactions may be added or removed by miners,
	// but it should provide a basis for setting a reasonable default.
	EstimateGas(ctx context.Context, msg ethereum.CallMsg) (uint64, error)
	// SendTransaction injects a signed transaction into the pending pool for execution.
	//
	// If the transaction was a contract creation use the TransactionReceipt method to get the
	// contract address after the transaction has been mined.
	SendTransaction(ctx context.Context, tx *types.Transaction) error

}
