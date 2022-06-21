package main

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"github.com/robfig/cron/v3"
	"time"
)

func main() {
	cfg1, err := cfg.LoadNodeConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
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
