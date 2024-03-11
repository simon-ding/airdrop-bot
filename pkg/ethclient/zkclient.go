package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"
	"airdrop-bot/pkg/contracts/storage"
	"airdrop-bot/utils"
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go"
	"github.com/zksync-sdk/zksync2-go/accounts"
)

func NewZkClient() *ZkClient {
	return &ZkClient{Client: NewClient(ChainZkEra)}
}

type ZkClient struct {
	*Client
	provider zksync2.Provider
}

func (c *ZkClient) Connect() error {
	// also, init ZkSync Provider, specify ZkSync2 RPC URL (e.g. testnet)
	//zp, err := zksync2.NewDefaultProvider("https://zksync-era.rpc.thirdweb.com/ed043a51ae23b0db3873f5a38b77ab28175fa496f15d3c53cf70401be89b622a")
	zp, err := zksync2.NewDefaultProvider("https://mainnet.era.zksync.io")
	if err != nil {
		return fmt.Errorf("coonect to zksync network: %v", err)
	}
	c.Client.client = zp.GetClient()
	c.provider = zp
	log.Infof("connect to zksync era rpc success")
	return nil
}

func (c *ZkClient) MuteIoSwap(privateKey string, from, to Token, amount float64) (string, error) {
	if from != TokenEth && to != TokenEth {
		return "", fmt.Errorf("only support eth swap")
	}

	var muteIOSwapContractAddress = "0x8B791913eB07C32779a16750e3868aA8495F5964"

	mute, err := abi.NewMuteIo(common.HexToAddress(muteIOSwapContractAddress), c.client)
	if err != nil {
		return "", err
	}

	publicKey := utils.GetPublicKey(privateKey)

	log.Infof("public key: %v", common.HexToAddress(publicKey))
	deadline := time.Now().Add(20 * time.Minute).Unix()

	auth, err := c.GetTransactor(privateKey)
	if err != nil {
		return "", err
	}
	//auth.GasLimit = uint64(4000000) // in units
	log.Infof("before transaction: %+v", auth)

	var txHash string
	if from == TokenEth && to == TokenUSDC {
		wei := utils.Eth2Wei(big.NewFloat(amount))
		auth.Value = wei
		path := []common.Address{
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenWETH)),
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenUSDC)),
		}
		out, err := mute.GetAmountOut(&bind.CallOpts{}, wei, path[0], path[1])
		if err != nil {
			return "", err
		}
		log.Infof("out USDC ammount: %+v", out)
		out.AmountOut.Div(out.AmountOut, big.NewInt(2))
		tx, err := mute.SwapExactETHForTokensSupportingFeeOnTransferTokens(auth, out.AmountOut, path, common.HexToAddress(publicKey), big.NewInt(deadline), []bool{false, false})
		if err != nil {
			return "", fmt.Errorf("swap to usdc: %v", err)
		}
		log.Infof("swap tx hash: %+v", tx.Hash().Hex())
		txHash = tx.Hash().Hex()
	}

	if from == TokenUSDC && to == TokenEth {

		if err := c.ApproveTokenAllowance(TokenUSDC, privateKey, muteIOSwapContractAddress); err != nil {
			return "", errors.Wrap(err, "approve allowance")
		}

		a := c.ConvertToken(from, big.NewFloat(amount))
		log.Infof("converted: %v", a)
		//auth.Value = a
		path := []common.Address{
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenUSDC)),
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenWETH)),
		}
		out, err := mute.GetAmountOut(&bind.CallOpts{}, a, path[0], path[1])
		if err != nil {
			return "", err
		}
		log.Infof("estimated ammount out: %v", out.AmountOut)
		out.AmountOut.Div(out.AmountOut, big.NewInt(2))

		tx, err := mute.SwapExactTokensForETHSupportingFeeOnTransferTokens(auth, a, out.AmountOut, path, common.HexToAddress(publicKey), big.NewInt(deadline), []bool{false, false})
		if err != nil {
			return "", fmt.Errorf("swap to eth: %v", err)
		}
		log.Infof("swap tx hash: %+v", tx.Hash().Hex())
		txHash = tx.Hash().Hex()

	}

	return txHash, nil
}

func (c *ZkClient) SyncSwapEth2Usdc(privateKey string, ammount *big.Float) (string, error) {
	const syncSwapContractAddress = "0x2da10A1e27bF85cEdD8FFb1AbBe97e53391C0295"
	const classicPoolFactoryAddr = "0xf2DAd89f2788a8CD54625C60b55cD3d2D0ACa7Cb"
	//const poolMasterAddr = "0xbB05918E9B4bA9Fe2c8384d223f0844867909Ffb"

	wethAddr := GetContractAddress(ChainZkEra, TokenWETH)
	usdcAddr := GetContractAddress(ChainZkEra, TokenUSDC)

	classicPoolFact, err := abi.NewSyncSwapPoolFactory(common.HexToAddress(classicPoolFactoryAddr), c.client)
	if err != nil {
		return "", errors.Wrap(err, "new classic pool")
	}
	poolAddr, err := classicPoolFact.GetPool(&bind.CallOpts{}, common.HexToAddress(wethAddr), common.HexToAddress(usdcAddr))
	if err != nil {
		return "", errors.Wrap(err, "get pool")
	}
	log.Infof("pool addr: %v", poolAddr.Hex())
	pool, err := abi.NewSyncSwapClassicPool(poolAddr, c.client)
	if err != nil {
		return "", errors.Wrap(err, "new pool")
	}
	a, err := pool.GetReserves(&bind.CallOpts{})
	log.Infof("get reserves: %v, %v", a, err)

	tokenAddress := common.HexToAddress(syncSwapContractAddress)
	syncClient, err := abi.NewSyncSwapRouter(tokenAddress, c.client)
	if err != nil {
		return "", err
	}
	ammountWei := utils.Eth2Wei(ammount)

	// Constructs the swap paths with steps.
	// Determine withdraw mode, to withdraw native ETH or wETH on last step.
	// 0 - vault internal transfer
	// 1 - withdraw and unwrap to naitve ETH
	// 2 - withdraw and wrap to wETH
	withdrawMode := uint8(1) // 1 or 2 to withdraw to user's wallet

	var buff bytes.Buffer
	binary.Write(&buff, binary.LittleEndian, withdrawMode)
	log.Infof("buffer: %v", buff.String())
	var data []byte
	wethBytes := common.LeftPadBytes(common.HexToAddress(wethAddr).Bytes(), 32)
	publicBytes := common.LeftPadBytes(common.HexToAddress(utils.GetPublicKey(privateKey)).Bytes(), 32)
	modeBytes := common.LeftPadBytes(buff.Bytes(), 32)
	data = append(data, wethBytes...)
	data = append(data, publicBytes...)
	data = append(data, modeBytes...)
	log.Infof("sync swap data: %v", hex.EncodeToString(data))
	auth, err := c.GetTransactor(privateKey)

	auth.Value = ammountWei

	steps := []abi.IRouterSwapStep{
		{
			Pool:         poolAddr,
			Data:         data,
			Callback:     common.HexToAddress(ZERO_ADDRESS),
			CallbackData: []byte(""),
		},
	}
	path := []abi.IRouterSwapPath{
		{
			Steps:    steps,
			TokenIn:  common.HexToAddress(ZERO_ADDRESS),
			AmountIn: ammountWei,
		},
	}
	log.Infof("sync swap path: %+v", path)
	deadline := time.Now().Add(20 * time.Minute).Unix()

	tx, err := syncClient.Swap(auth, path, big.NewInt(0), big.NewInt(deadline))
	if err != nil {
		return "", errors.Wrap(err, "syncswap")
	}
	log.Infof("syncswap address: %v", tx.Hash().Hex())

	return tx.Hash().Hex(), nil
}

func (c *ZkClient) SyncSwapUsdc2Eth(privateKey string, ammount *big.Float) (string, error) {
	const syncSwapContractAddress = "0x2da10A1e27bF85cEdD8FFb1AbBe97e53391C0295"
	const classicPoolFactoryAddr = "0xf2DAd89f2788a8CD54625C60b55cD3d2D0ACa7Cb"
	//const poolMasterAddr = "0xbB05918E9B4bA9Fe2c8384d223f0844867909Ffb"

	wethAddr := GetContractAddress(ChainZkEra, TokenWETH)
	usdcAddr := GetContractAddress(ChainZkEra, TokenUSDC)

	if err := c.ApproveTokenAllowance(TokenUSDC, privateKey, syncSwapContractAddress); err != nil {
		return "", errors.Wrap(err, "approve allowance")
	}

	classicPoolFact, err := abi.NewSyncSwapPoolFactory(common.HexToAddress(classicPoolFactoryAddr), c.client)
	if err != nil {
		return "", errors.Wrap(err, "new classic pool")
	}
	poolAddr, err := classicPoolFact.GetPool(&bind.CallOpts{}, common.HexToAddress(usdcAddr), common.HexToAddress(wethAddr))
	if err != nil {
		return "", errors.Wrap(err, "get pool")
	}
	log.Infof("pool addr: %v", poolAddr.Hex())

	pool, err := abi.NewSyncSwapClassicPool(poolAddr, c.client)
	if err != nil {
		return "", errors.Wrap(err, "new pool")
	}
	a, err := pool.GetReserves(&bind.CallOpts{})
	log.Infof("get reserves: %v, %v", a, err)

	tokenAddress := common.HexToAddress(syncSwapContractAddress)
	syncClient, err := abi.NewSyncSwapRouter(tokenAddress, c.client)
	if err != nil {
		return "", err
	}

	// Constructs the swap paths with steps.
	// Determine withdraw mode, to withdraw native ETH or wETH on last step.
	// 0 - vault internal transfer
	// 1 - withdraw and unwrap to naitve ETH
	// 2 - withdraw and wrap to wETH
	withdrawMode := uint8(1) // 1 or 2 to withdraw to user's wallet

	var buff bytes.Buffer
	binary.Write(&buff, binary.LittleEndian, withdrawMode)
	log.Infof("buffer: %v", buff.String())
	var data []byte
	usdcBytes := common.LeftPadBytes(common.HexToAddress(usdcAddr).Bytes(), 32)
	publicBytes := common.LeftPadBytes(common.HexToAddress(utils.GetPublicKey(privateKey)).Bytes(), 32)
	modeBytes := common.LeftPadBytes(buff.Bytes(), 32)
	data = append(data, usdcBytes...)
	data = append(data, publicBytes...)
	data = append(data, modeBytes...)
	log.Infof("sync swap data: %v", hex.EncodeToString(data))
	auth, err := c.GetTransactor(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "get auth")
	}

	ammount1 := c.ConvertToken(TokenUSDC, ammount)

	steps := []abi.IRouterSwapStep{
		{
			Pool:         poolAddr,
			Data:         data,
			Callback:     common.HexToAddress(ZERO_ADDRESS),
			CallbackData: []byte(""),
		},
	}
	path := []abi.IRouterSwapPath{
		{
			Steps:    steps,
			TokenIn:  common.HexToAddress(usdcAddr),
			AmountIn: ammount1,
		},
	}
	log.Infof("sync swap path: %+v", path)
	deadline := time.Now().Add(20 * time.Minute).Unix()

	tx, err := syncClient.Swap(auth, path, big.NewInt(0), big.NewInt(deadline))
	if err != nil {
		return "", errors.Wrap(err, "syncswap")
	}
	log.Infof("syncswap address: %v", tx.Hash().Hex())

	return tx.Hash().Hex(), nil

}

func (c *ZkClient) ZnsBuyDomain(domain string, privateKey string) (string, error) {
	znsContractAddr := common.HexToAddress("0xCBE2093030F485adAaf5b61deb4D9cA8ADEAE509")

	pubKey := utils.GetPublicKey(privateKey)
	znszks, err := abi.NewZksBaseRegister(znsContractAddr, c.client)
	if err != nil {
		return "", errors.Wrap(err, "new zns zks")
	}
	isAvailable, canMint, err := znszks.CanRegister(&bind.CallOpts{}, domain, common.HexToAddress(pubKey))
	if err != nil {
		return "", errors.Wrap(err, "can register")
	}
	if !isAvailable || !canMint {
		return "", errors.New("You can not register this name at this moment: " + domain)
	}
	log.Infof("%s is available for register", domain)
	auth, err := c.GetTransactor(privateKey)
	if err != nil {
		return "", err
	}

	tx, err := znszks.Register(auth, domain, common.HexToAddress(pubKey), big.NewInt(1))
	if err != nil {
		return "", errors.Wrap(err, "")
	}

	return tx.Hash().Hex(), nil
}

func (c *ZkClient) ZnsGetOwnedDomains(address string) ([]string, error) {
	znsContractAddr := common.HexToAddress("0xCc788c0495894C01F01cD328CF637c7C441Ee69E")
	znszks, err := abi.NewZksBaseRegisterImpl(znsContractAddr, c.client)
	if err != nil {
		return nil, errors.Wrap(err, "new zns zks")
	}
	_, domains, err := znszks.GetOwnedDomains(&bind.CallOpts{}, common.HexToAddress(address))
	if err != nil {
		return nil, errors.Wrap(err, "get domains")
	}
	log.Infof("get domains for address: %s, domains: %v", address, domains)
	return domains, nil
}

func (c *ZkClient) DeploySimpleStorageContract(priKey string) (string, error) {
	// Create wallet
	wallet, err := accounts.NewWallet(common.Hex2Bytes(priKey), &c.provider, nil)
	if err != nil {
		log.Panic(err)
	}

	// Deploy smart contract
	hash, err := wallet.DeployWithCreate(nil, accounts.CreateTransaction{Bytecode: storage.StorageBinData})
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction: ", hash)
	return hash.Hex(), nil
}
