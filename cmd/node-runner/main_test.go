package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/utils"
	"os"
	"os/exec"
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
	if cfg1.XvfbMod() {
		res := "1920x1080x24"
		log.Infof("xvfb mode, run xvfb %s...", res)
		go exec.Command("Xvfb", ":1", "-screen", "0", res).Run()
		os.Setenv("DISPLAY", ":1")
	}

	client := New(cfg1)

	client.bridgeOne("surprise fog bean render you soda ostrich add soul then asthma salmon")
}
