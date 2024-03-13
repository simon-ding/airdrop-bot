package cloudflare

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/pkg/errors"
)

func NewClient(cfg *cfg.Cloudflare) (*Client, error) {
	api, err := cloudflare.NewWithAPIToken(cfg.ApiKey)
	if err != nil {
		return nil, errors.Wrap(err, "new token")
	}
	log.Infof("%+v", cfg)

	return &Client{
		cfg: cfg,
		api: api,
	}, nil
}

type Client struct {
	cfg *cfg.Cloudflare
	api *cloudflare.API
}

func (c *Client) Get(key string) []byte {
	v, err := c.api.GetWorkersKV(context.TODO(), cloudflare.AccountIdentifier(c.cfg.AccountId), cloudflare.GetWorkersKVParams{
		NamespaceID: c.cfg.KvId,
		Key:         key,
	})
	if err != nil {
		log.Errorf("get value error, key: %v, err: %v", key, err)
		return nil
	}
	return v
}

func (c *Client) Put(key string, value []byte) error {
	_, err := c.api.WriteWorkersKVEntry(context.TODO(), cloudflare.AccountIdentifier(c.cfg.AccountId), cloudflare.WriteWorkersKVEntryParams{
		NamespaceID: c.cfg.KvId,
		Key:         key,
		Value:       value,
	})
	return err
}
