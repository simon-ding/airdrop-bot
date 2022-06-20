package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/aws"
	"airdrop-bot/cfg"
	"airdrop-bot/db"
	"airdrop-bot/log"
	server2 "airdrop-bot/server"
	"airdrop-bot/utils"
	"flag"
	"os"
	"os/exec"
	"path"
)

var accounts = flag.String("accounts", "", "accounts: 1-30  1")
var genAccounts = flag.Bool("genAccounts", false, "-genAccounts")
var transferEth = flag.Bool("transfer", false, "transfer eth from binance to accounts")

func main() {

	flag.Parse()
	cfg, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
	}
	log.Info("read config success")
	awsDir := path.Join(cfg.Dir, "aws.config")
	err = aws.WriteAwsConfig(cfg.AWS.AccessKeyID, cfg.AWS.SecretAccessKey, awsDir)
	if err != nil {
		log.Errorf("write aws config error: %v", err)
		return
	}

	err = utils.Unzip(asserts.Ext, path.Join(cfg.Dir, "ext"))
	if err != nil {
		log.Errorf("unzip extension error: %v", err)
		return
	}

	db.Open(cfg.Dir)
	log.Infof("db open success")
	if cfg.XvfbMod() {
		res := "1920x1080x24"
		log.Infof("xvfb mode, run xvfb %s...", res)
		go exec.Command("Xvfb", ":1", "-screen", "0", res).Run()
		os.Setenv("DISPLAY", ":1")
	}

	server, err := server2.NewServer(cfg)
	if err != nil {
		log.Errorf("new server: %v", err)
		return
	}

	if *transferEth {
		err := server.TransferEthToAccounts()
		if err != nil {
			log.Errorf("transfer eth error: %v", err)
		}
	}

	if *genAccounts {
		server.FirstRunGenAccount()
	}
	err = server.BridgeAccounts(*accounts)
	if err != nil {
		log.Errorf("bridge accounts: %v", err)
	}
}
