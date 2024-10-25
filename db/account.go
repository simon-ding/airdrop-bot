package db

import (
	"airdrop-bot/log"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

type Account struct {
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Mnemonic holds the value of the "mnemonic" field.
	Mnemonic string `json:"mnemonic,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// PrivateKey holds the value of the "privateKey" field.
	PrivateKey string `json:"privateKey,omitempty"`
}

const accountsKey = "accounts"

func (c *Client) GetAllAccountsOld() []Account {
	v := c.cf.Get(accountsKey)
	if v == nil {
		log.Errorf("no account find")
		return nil
	}
	var accounts []Account
	err := json.Unmarshal(v, &accounts)
	if err != nil {
		log.Errorf("unmarshalling data error, data : %v, error: %v", string(v), err)
		return nil
	}
	return accounts
}

const accountPrefix = "account_"

func (c *Client) GetAllAccounts() []Account {
	v := c.cf.GetByPrefix(accountPrefix)
	if v == nil {
		log.Errorf("no account find")
		return nil
	}
	var accounts []Account
	for _, v1 := range v {
		var a Account
		err := json.Unmarshal(v1, &a)
		if err != nil {
			log.Errorf("unmarshalling data error, data : %v, error: %v", string(v1), err)
			return nil
		}
		accounts = append(accounts, a)
	}

	sort.Slice(accounts, func(i, j int) bool {
		return accounts[i].ID < accounts[j].ID
	})
	return accounts
}

func (c *Client) writeAccount(a Account) error {
	if a.ID == 0 || a.Address == "" || a.Mnemonic == "" || a.PrivateKey == "" {
		return fmt.Errorf("account should no empty field")
	}
	key := accountPrefix + fmt.Sprintf("%04d", a.ID)
	return c.putValue(key, a)
}

// update by id
func (c *Client) UpdateAccount(a Account) error {
	accountOld := c.GetAccount(a.ID)

	if accountOld == nil {
		return fmt.Errorf("no related account found: %v", a)
	}
	return c.writeAccount(a)
}

func (c *Client) AddAccount(a Account) error {
	accountsAll := c.GetAllAccounts()

	for _, ori := range accountsAll {
		if ori.Address == a.Address {
			return fmt.Errorf("account with same address already exist: %v", ori)
		}
	}
	if len(accountsAll) > 0 {
		maxId := accountsAll[len(accountsAll)-1].ID
		a.ID = maxId + 1
	}

	return c.writeAccount(a)
}

func (c *Client) GetAccount(id int) *Account {
	key := accountPrefix + strconv.Itoa(id)
	v := c.cf.Get(key)
	var a Account
	err := json.Unmarshal(v, &a)
	if err != nil {
		log.Errorf("unmarshal account: %v", err)
		return nil
	}
	return &a
}
