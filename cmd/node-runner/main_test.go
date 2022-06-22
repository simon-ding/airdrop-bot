package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/utils"
	"path"
	"testing"
)

func TestNew(t *testing.T) {
	cfg1, err := cfg.LoadNodeConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
		return
	}

	err = utils.Unzip(asserts.Ext, path.Join(cfg1.Dir, "ext"))
	if err != nil {
		log.Errorf("unzip extension error: %v", err)
		return
	}

	client := New(cfg1)

	client.bridgeOne("surprise fog bean render you soda ostrich add soul then asthma salmon")
}
