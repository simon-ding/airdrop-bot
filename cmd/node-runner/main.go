package main

import (
	"airdrop-bot/asserts"
	"airdrop-bot/cfg"
	client2 "airdrop-bot/cmd/node-runner/client"
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
		go func() {
			b, err := exec.Command("Xvfb", ":1", "-screen", "0", res).Output()
			if err != nil {
				log.Panicf("run xvfb error: %v:　%v", string(b), err)
			}
		}()
		os.Setenv("DISPLAY", ":1")
	}

	client := client2.New(&cfg1.Node)

	c := cron.New()
	c.AddFunc("@every 10s", func() {
		if err := client.Heartbeat(); err != nil {
			log.Errorf("heartbeat error: %v", err)
		}
	})
	c.Start()
	for {
		time.Sleep(time.Hour)
	}
}
