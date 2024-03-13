package main

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"

	"airdrop-bot/server"
)

func main() {

	cfg1, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")

	server, err := server.NewServer(&cfg1.Server)
	if err != nil {
		log.Errorf("server error: %v", err)
		return
	}

	if err := server.Serve(); err != nil {
		log.Errorf("start server error: %v", err)
		return
	}
}
