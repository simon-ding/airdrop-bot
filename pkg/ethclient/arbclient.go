package ethclient

import (
	"airdrop-bot/log"
	"airdrop-bot/pkg/abi"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

func (a *ArbOneClient) ClaimAidoge(privateKey string) error {

	contractAddr := "0x7c20acfd25467dE0B92d03E4C4d304f18B8408E1"
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return fmt.Errorf("encode private key: %v", err)
	}

	publicKey := priKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	doge, err := abi.NewAidoge(common.HexToAddress(contractAddr), a.client)
	if err != nil {
		return err
	}
	am, err := doge.CanClaimAmount(&bind.CallOpts{})
	log.Infof("results %v: %v", am, err)

	// nonce1, err := a.client.PendingNonceAt(context.Background(), common.BytesToAddress(publicKeyBytes))
	// if err != nil {
	// 	return fmt.Errorf("get nonce: %v", err)
	// }

	// gasPrice, err := a.client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return err
	// }
	auth := bind.NewKeyedTransactor(priKey)
	//auth.Nonce = big.NewInt(int64(nonce1) + 4)
	//auth.Value = big.NewInt(0)      // in wei
	//auth.GasLimit = uint64(4000000) // in units
	//auth.GasPrice = gasPrice
	//auth.GasFeeCap = big.NewInt(104950000)
	//ethmodAddr := "0xf6303ef5"
	log.Infof("auth: %+v", auth)

	nonce := big.NewInt(time.Now().UTC().UnixNano())
	addr := "0x0000000000000000000000000000000000000000"

	log.Infof("nonce: %v", nonce)
	hash := crypto.Keccak256(
		common.LeftPadBytes(common.HexToAddress(contractAddr).Bytes(), 32),
		common.LeftPadBytes(publicKeyBytes, 32),
		common.LeftPadBytes(nonce.Bytes(), 32),
	)
	s := common.Bytes2Hex(hash[:])
	log.Infof("sig: %v, len: %d", s, len(hash))

	sig, err := crypto.Sign(hash, priKey)
	if err != nil {
		return err
	}
	log.Infof("sig: %v", common.Bytes2Hex(sig))

	tx, err := doge.Claim(auth, nonce, sig, common.HexToAddress(addr))
	log.Infof("tx: %v, error: %v", tx, err)
	return nil
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
