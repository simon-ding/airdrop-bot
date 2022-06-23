package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/utils"
	"github.com/robfig/cron/v3"
	"os"
	"os/exec"
	"path"
	"time"
)

func main() {
	cfg1, err := cfg.LoadConfig()
	if err != nil {
		log.Panicf("read config: %v", err)
		return
	}

	err = utils.Unzip(asserts.Ext, path.Join(cfg1.Node.Dir, "ext"))
	if err != nil {
		log.Errorf("unzip extension error: %v", err)
		return
	}

	if cfg1.Node.XvfbMod() {
		res := "1920x1080x24"
		log.Infof("xvfb mode, run xvfb %s...", res)
		go exec.Command("Xvfb", ":1", "-screen", "0", res).Run()
		os.Setenv("DISPLAY", ":1")
	}

	client := New(&cfg1.Node)

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
