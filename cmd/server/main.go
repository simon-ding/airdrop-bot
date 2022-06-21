package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	"airdrop-bot/utils"
	"path"
)

func main() {

	cfg1, err := cfg.LoadServerConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")
	awsDir := path.Join(cfg1.Dir, "aws.config")
	err = aws.WriteAwsConfig(cfg1.AWS.AccessKeyID, cfg1.AWS.SecretAccessKey, awsDir)
	if err != nil {
		log.Errorf("write aws config error: %v", err)
		return
	}

	err = utils.Unzip(asserts.Ext, path.Join(cfg1.Dir, "ext"))
	if err != nil {
		log.Errorf("unzip extension error: %v", err)
		return
	}

	db.Open(cfg1.Dir)
	log.Infof("db open success")

	server, err := NewServer(cfg1)
	if err != nil {
		log.Errorf("server error: %v", err)
		return
	}

	if err := server.Serve(); err != nil {
		log.Errorf("start server error: %v", err)
		return
	}
}
