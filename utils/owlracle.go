package utils

import (
	"airdrop-bot/cfg"
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"time"
)

type GasFee struct {
	Timestamp time.Time  `json:"timestamp"`
	LastBlock int        `json:"lastBlock"`
	AvgTime   float64    `json:"avgTime"`
	AvgTx     float64    `json:"avgTx"`
	AvgGas    float64    `json:"avgGas"`
	BaseFee   float64    `json:"baseFee"`
	Speeds    []GasSpeed `json:"speeds"`
}

type GasSpeed struct {
	Acceptance   float64 `json:"acceptance"`
	GasPrice     float64 `json:"gasPrice"`
	EstimatedFee float64 `json:"estimatedFee"`
}

const owlracleUrl = "https://owlracle.info/eth/gas"

func GetGasFee(cfg *cfg.Owlracle) (*GasFee, error) {
	req, err := http.NewRequest(http.MethodGet, owlracleUrl, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}
	req.Header.Add("apikey", cfg.ApiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response")
	}
	var gas GasFee
	err = json.Unmarshal(data, &gas)
	return &gas, err
}
