package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func NewArbOneClient() *ArbOneClient {
	return &ArbOneClient{Client: NewClient(ChainArbOne)}
}

type ArbOneClient struct {
	*Client
}

func (a *ArbOneClient) TokenSum() {
	// addr := common.HexToAddress("0x67a24CE4321aB3aF51c2D0a4801c3E111D88C9d9")
	// instance, err := abi.NewArbToken(addr, a.client)
	// if err != nil {
	// 	log.Errorf("err: ", err)
	// }
	// instance.ParseCanClaim(nil)
}

func NewArbNovaClient() *ArbNovaClient {
	return &ArbNovaClient{Client: NewClient(ChainArbNova)}
}

type ArbNovaClient struct {
	*Client
}

const syncSwapContractAddress = "0x2da10A1e27bF85cEdD8FFb1AbBe97e53391C0295"

func (a *ArbNovaClient) SyncSwap() error {
	tokenAddress := common.HexToAddress(syncSwapContractAddress)
	syncClient, err := abi.NewSyncSwapRouter(tokenAddress, a.client)
	if err != nil {
		return err
	}
	addr, err := syncClient.WETH(&bind.CallOpts{})
	if err != nil {
		return err
	}
	log.Infof("vault address: %v", addr.Hex())
	return nil
}
