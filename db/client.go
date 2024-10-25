package db

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/pkg/cloudflare"
	"encoding/json"
	"fmt"
)

type Client struct {
	cf *cloudflare.Client
}

func NewClient(cfg *cfg.Cloudflare) (*Client, error) {
	cf, err := cloudflare.NewClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("new cf: %v", err)
	}
	return &Client{cf: cf}, nil
}

type BinanceSetting struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

func (c *Client) putValue(key string, v interface{}) error {
	d, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return c.cf.Put(key, d)
}

const binanceSettingKey = "binance_setting"

func (c *Client) GetBinanceSetting() BinanceSetting {
	var bs BinanceSetting
	v := c.cf.Get(binanceSettingKey)
	err := json.Unmarshal(v, &bs)
	if err != nil {
		log.Errorf("marshalling: %v", err)
	}
	return bs
}

func (c *Client) SaveBinanceSetting(bs BinanceSetting) error {
	return c.putValue(binanceSettingKey, bs)
}
