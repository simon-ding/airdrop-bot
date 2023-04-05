package main

import (
	"airdrop-bot/pkg/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"path"
)

func main() {

	cfg1, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")

	db.Open(cfg1.Server.DbFile)
	log.Infof("db open success")

	awsDir := path.Join(cfg1.Server.Dir, "aws.config")
	err = aws.WriteAwsConfig(cfg1.Server.AWS.AccessKeyID, cfg1.Server.AWS.SecretAccessKey, awsDir)
	if err != nil {
		log.Errorf("write aws config error: %v", err)
		return
	}

	server, err := NewServer(&cfg1.Server)
	if err != nil {
		log.Errorf("server error: %v", err)
		return
	}
	server.FirstRunGenAccount()

	if err := server.Serve(); err != nil {
		log.Errorf("start server error: %v", err)
		return
	}
}
