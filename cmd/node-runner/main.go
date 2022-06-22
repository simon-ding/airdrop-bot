package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/utils"
	"github.com/robfig/cron/v3"
	"path"
	"time"
)

func main() {
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

	c := cron.New()
	c.AddFunc("@every 10s", func() {
		if err := client.heartbeat(); err != nil {
			log.Errorf("heartbeat error: %v", err)
		}
	})
	c.Start()
	for {
		time.Sleep(time.Hour)
	}
}
