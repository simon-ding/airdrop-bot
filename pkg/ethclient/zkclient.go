package ethclient

import (
	"fmt"

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
