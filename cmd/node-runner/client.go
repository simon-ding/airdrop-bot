package main

import (
	"airdrop-bot/cfg"
	"airdrop-bot/log"
	"airdrop-bot/metamask"
	"airdrop-bot/services"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/chromedp"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path"
)

func New(cfg *cfg.NodeConfig) *Client {
	c := &Client{
		cfg:  cfg,
		http: http.DefaultClient,
	}
	return c
}

type Client struct {
	cfg  *cfg.NodeConfig
	http *http.Client
}

func (c *Client) heartbeat() error {
	n := cfg.Heartbeat{
		NodeName:  c.cfg.NodeName,
		DnsName:   c.cfg.DnsName,
		AWSRegion: c.cfg.Region,
	}

	url := c.cfg.ServerUrl + cfg.ApiHeartbeatUrl
	data, err := json.Marshal(&n)
	if err != nil {
		return errors.Wrap(err, "json marshal")
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(cfg.AuthHeader, c.cfg.Token)

	rsp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request")
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", rsp.Status)
	}
	data, err = ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}
	var rr cfg.HeartbeatResp

	err = json.Unmarshal(data, &rr)
	if err != nil {
		return err
	}
	if rr.Account == "" {
		return nil
	}
	log.Infof("task accepted step %v, track id %v", rr.Task, rr.TrackID)
	go func() {
		err := c.bridgeOne(rr.Account)
		if err != nil {
			log.Errorf("bridge error: %v", err)
		}
		err = c.updateTaskStatus(rr.TrackID, err)
		if err != nil {
			log.Errorf("update task status error: %v", err)
		}
	}()
	return nil
}

func (c *Client) updateTaskStatus(trackId uint, err error) error {
	ts := cfg.TaskStatus{
		TrackID: trackId,
	}
	if err != nil {
		ts.Result = cfg.ResultFail
		ts.Reason = err.Error()
	} else {
		ts.Result = cfg.ResultSuccess
	}
	url := c.cfg.ServerUrl + cfg.ApiTaskStatusUpdateUrl
	data, err := json.Marshal(&ts)
	if err != nil {
		return errors.Wrap(err, "json marshal")
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(cfg.AuthHeader, c.cfg.Token)

	rsp, err := c.http.Do(req)
	if err != nil {
		return errors.Wrap(err, "http request")
	}
	defer rsp.Body.Close()
	if rsp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status %v", rsp.Status)
	}
	return nil

}

func (c *Client) startChrome() (context.Context, func()) {
	if c.cfg.XvfbMod() {
		res := "1920x1080x24"
		log.Infof("xvfb mode, run xvfb %s...", res)
		go exec.Command("Xvfb", ":1", "-screen", "0", res).Run()
		os.Setenv("DISPLAY", ":1")
	}

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", false),
		chromedp.Flag("disable-extensions", false),
		chromedp.Flag("restore-on-startup", false),
		chromedp.Flag("disable-web-security", true),
		chromedp.Flag("load-extension", path.Join(c.cfg.Dir, "ext")),
		chromedp.Flag("window-size", "1920,1080"),

		//chromedp.UserDataDir(dataDir),
	)

	ctx, cancel1 := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Debugf), chromedp.WithLogf(log.Infof),
		chromedp.WithErrorf(log.Errorf))

	allocCtx, cancel2 := chromedp.NewExecAllocator(ctx, opts...)
	return allocCtx, func() {
		cancel1()
		cancel2()
	}
}

func (c *Client) bridgeOne(mnemonic string) error {
	ctx, cancel := c.startChrome()
	defer cancel()

	meta := metamask.NewMetamask(ctx)
	err := meta.FirstOpenAndImportAccount(c.cfg.WalletPassword, mnemonic)
	if err != nil {
		return err
	}
	balance, err := meta.Balance()
	log.Infof("account current balance is %v", balance)

	err = meta.OpenChanListAndAddNetwork("Optimism")
	if err != nil {
		return errors.Wrap(err, "add op")
	}

	hop := services.MewHop(meta.Context(), meta, "ETH", "optimism", "arbitrum")
	err = hop.LinkMetaMask()
	if err != nil {
		return err
	}
	log.Infof("hop link metamask success")

	err = hop.BridgeMoney(balance - cfg.BridgeReverse)
	if err != nil {
		return errors.Wrap(err, "arb deposit")
	}

	balance = meta.WaitForBalanceChange(balance)
	return nil

}