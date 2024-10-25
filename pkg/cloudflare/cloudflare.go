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

func (c *Client) ListAllKeys(prefix string) []string {
	resp, err := c.api.ListWorkersKVKeys(context.Background(), cloudflare.AccountIdentifier(c.cfg.AccountId), cloudflare.ListWorkersKVsParams{
		NamespaceID: c.cfg.KvId,
		Prefix:      prefix,
	})
	if err != nil {
		log.Errorf("get kv keys: %v", err)
		return nil
	}
	var keys []string
	for _, v := range resp.Result {
		keys = append(keys, v.Name)
	}
	return keys
}

func (c *Client) GetByPrefix(prefix string) [][]byte {
	keys := c.ListAllKeys(prefix)

	var res [][]byte
	for _, k := range keys {
		v := c.Get(k)
		res = append(res, v)
	}
	return res
}

func (c *Client) Delete(key string) error {
	_, err := c.api.DeleteWorkersKVEntry(context.Background(), cloudflare.AccountIdentifier(c.cfg.AccountId), cloudflare.DeleteWorkersKVEntryParams{
		NamespaceID: c.cfg.KvId,
		Key:         key,
	})
	return err
}

func (c *Client) DeleteByPrefix(prefix string) error {
	keys := c.ListAllKeys(prefix)
	_, err := c.api.DeleteWorkersKVEntries(context.Background(), cloudflare.AccountIdentifier(c.cfg.AccountId), cloudflare.DeleteWorkersKVEntriesParams{
		NamespaceID: c.cfg.KvId,
		Keys:        keys,
	})
	return err
}
