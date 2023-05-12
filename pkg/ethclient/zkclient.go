package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"
	"airdrop-bot/utils"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/zksync-sdk/zksync2-go"
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
	zp, err := zksync2.NewDefaultProvider("https://zksync-era.rpc.thirdweb.com/ed043a51ae23b0db3873f5a38b77ab28175fa496f15d3c53cf70401be89b622a")
	//zp, err := zksync2.NewDefaultProvider("https://mainnet.era.zksync.io")
	if err != nil {
		return fmt.Errorf("coonect to zksync network: %v", err)
	}
	c.Client.client = zp.GetClient()
	c.provider = zp
	log.Infof("connect to zksync era rpc success")
	return nil
}

func (c *ZkClient) MuteIoSwap(privateKey string, from, to Token, amount float64)(string, error) {
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

		a := c.ConvertToken(from, amount)
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
			return "", fmt.Errorf("swap tp eth: %v", err)
		}
		log.Infof("swap tx hash: %+v", tx.Hash().Hex())
		txHash = tx.Hash().Hex()

	}

	return txHash, nil
}


func (c *ZkClient) SyncSwap(privateKey string, ammount *big.Float) error {
	const syncSwapContractAddress = "0x2da10A1e27bF85cEdD8FFb1AbBe97e53391C0295"
	const classicPoolFactoryAddr = "0xf2DAd89f2788a8CD54625C60b55cD3d2D0ACa7Cb"
	//const poolMasterAddr = "0xbB05918E9B4bA9Fe2c8384d223f0844867909Ffb"

	wethAddr := GetContractAddress(ChainZkEra, TokenWETH)
	usdcAddr := GetContractAddress(ChainZkEra, TokenUSDC)
	classicPoolFact, err := abi.NewSyncSwapPoolFactory(common.HexToAddress(classicPoolFactoryAddr), c.client)
	if err != nil {
		return errors.Wrap(err, "new classic pool")
	}
	poolAddr, err := classicPoolFact.GetPool(&bind.CallOpts{}, common.HexToAddress(wethAddr), common.HexToAddress(usdcAddr))
	if err != nil {
		return errors.Wrap(err, "get pool")
	}
	pool, err := abi.NewSyncSwapClassicPool(poolAddr, c.client)
	if err != nil {
		return errors.Wrap(err, "new pool")
	}
	a, err := pool.GetReserves(&bind.CallOpts{})
	log.Infof("get reserves: %v, %v", a, err)

	tokenAddress := common.HexToAddress(syncSwapContractAddress)
	syncClient, err := abi.NewSyncSwapRouter(tokenAddress, c.client)
	if err != nil {
		return err
	}
	// auth, err := c.GetTransactor(privateKey)
	// path = []abi.IRouterSwapPath{
		
	// }
	// syncClient.Swap(auth, )
	addr, err := syncClient.WETH(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Infof("vault address: %v", addr.Hex())
	return nil
}
