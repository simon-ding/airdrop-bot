package client

import (
	"airdrop-bot/cfg"
	"airdrop-bot/db"
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
	"path"
	"time"
)

func New(cfg *cfg.NodeConfig) *Client {
	c := &Client{
		cfg:  cfg,
		http: http.DefaultClient,
	}
	return c
}

type Client struct {
	cfg    *cfg.NodeConfig
	http   *http.Client
	locked bool
}

func (c *Client) Heartbeat() error {

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
	log.Infof("task accepted step %v, %+v", rr.Task, rr)
	go func() {
		t := time.After(rr.Timeout)
		var err error
		ch := make(chan error)

		go func() {
			ch <- c.BridgeOne(rr.Account, rr.Task)
		}()
		select {
		case <-t:
			err = fmt.Errorf("timeout")
		case x := <-ch:
			err = x
		}
		if err != nil {
			log.Errorf("task error: %v", err)
		}
		err2 := c.updateTaskStatus(rr.TrackID, err)
		if err != nil {
			log.Errorf("update task status error: %v", err2)
		}
		log.Infof("task done!!! task id %v, err: %v", rr.TrackID, err)
	}()
	return nil
}

func (c *Client) updateTaskStatus(trackId uint, err error) error {
	log.Infof("begin update task status: %v, err: %v", trackId, err)
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

func (c *Client) BridgeOne(mnemonic string, step string) error {
	ctx, cancel := c.startChrome()
	defer cancel()

	meta := metamask.NewMetamask(ctx)
	err := meta.FirstOpenAndImportAccount(c.cfg.WalletPassword, mnemonic)
	if err != nil {
		return err
	}

	if step == db.StepArbitrumBridge2 {
		log.Infof("bridge from arbitrum to op")
		err = meta.OpenChanListAndAddNetwork("Arbitrum One")
		if err != nil {
			return errors.Wrap(err, "add arb")
		}

		balance, err := meta.Balance()
		log.Infof("account current balance is %v", balance)
		if err == nil && balance != 0 {
			hop := services.MewHop(meta.Context(), meta, "ETH", "arbitrum", "optimism")
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
		}

	}

	log.Infof("bridge from op to arbitrum")
	err = meta.OpenChanListAndAddNetwork("Optimism")
	if err != nil {
		return errors.Wrap(err, "add op")
	}

	balance, err := meta.Balance()
	log.Infof("account current balance is %v", balance)
	if err != nil || balance == 0 {
		return fmt.Errorf("current balance %v", balance)
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
