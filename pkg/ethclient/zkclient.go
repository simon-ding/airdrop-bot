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
	zp, err := zksync2.NewDefaultProvider("https://mainnet.era.zksync.io")
	if err != nil {
		return fmt.Errorf("coonect to zksync network: %v", err)
	}
	c.Client.client = zp.GetClient()
	c.provider = zp
	return nil
}

func (c *ZkClient) MuteIoSwap(privateKey string, from, to Token, amount float64) error {
	if from != TokenEth && to != TokenEth {
		return fmt.Errorf("only support eth swap")
	}

	var muteIOSwapContractAddress = "0x8B791913eB07C32779a16750e3868aA8495F5964"

	mute, err := abi.NewMuteIo(common.HexToAddress(muteIOSwapContractAddress), c.client)
	if err != nil {
		return err
	}

	publicKey := utils.GetPublicKey(privateKey)

	log.Infof("public key: %v", common.HexToAddress(publicKey))
	deadline := time.Now().Add(20 * time.Minute).Unix()

	auth, err := c.GetTransactor(privateKey)
	if err != nil {
		return err
	}
	auth.GasLimit = uint64(4000000) // in units
	log.Infof("before transaction: %+v", auth)

	if from == TokenEth && to == TokenUSDC {
		wei := utils.Eth2Wei(big.NewFloat(amount))
		auth.Value = wei
		path := []common.Address{
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenWETH)),
			common.HexToAddress(GetContractAddress(ChainZkEra, TokenUSDC)),
		}
		out, err := mute.GetAmountOut(&bind.CallOpts{}, wei, path[0], path[1])
		if err != nil {
			return err
		}
		log.Infof("out USDC ammount: %+v", out)

		tx, err := mute.SwapExactETHForTokensSupportingFeeOnTransferTokens(auth, big.NewInt(0), path, common.HexToAddress(publicKey), big.NewInt(deadline), []bool{false, false})
		if err != nil {
			return fmt.Errorf("swap usdc: %v", err)
		}
		log.Infof("swap tx hash: %+v", tx.Hash().Hex())
	}

	if from == TokenUSDC && to == TokenEth {

		if err := c.ApproveTokenAllowance(TokenUSDC, privateKey, muteIOSwapContractAddress); err != nil {
			return errors.Wrap(err, "approve allowance")
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
			return err
		}
		log.Infof("estimated ammount out: %v", out.AmountOut)

		tx, err := mute.SwapExactTokensForETHSupportingFeeOnTransferTokens(auth, a, big.NewInt(0), path, common.HexToAddress(publicKey), big.NewInt(deadline), []bool{false, false})
		if err != nil {
			return fmt.Errorf("swap usdc: %v", err)
		}
		log.Infof("swap tx hash: %+v", tx.Hash().Hex())

	}

	return nil
}
