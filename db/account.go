package db

import (
	"airdrop-bot/log"
	"encoding/json"
	"fmt"
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

func (c *Client) GetAllAccounts() []Account {
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

func (c *Client) writeAllAccounts(accounts []Account) error {
	return c.putValue(accountsKey, accounts)
}

// update by id
func (c *Client) UpdateAccount(a Account) error {
	accountsAll := c.GetAllAccounts()
	var found = false
	for i, ori := range accountsAll {
		if ori.ID == a.ID {
			found = true
			if a.Mnemonic != "" {
				accountsAll[i].Mnemonic = a.Mnemonic
			}
			if a.Address != "" {
				accountsAll[i].Address = a.Address
			}
			if a.PrivateKey != "" {
				accountsAll[i].PrivateKey = a.PrivateKey
			}
		}
	}
	if !found {
		return fmt.Errorf("no related account found: %v", a)
	}
	return c.writeAllAccounts(accountsAll)
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
	accountsAll = append(accountsAll, a)
	return c.writeAllAccounts(accountsAll)
}

func (c *Client) DeleteAccount(a Account) error {
	accountsAll := c.GetAllAccounts()

	var newAccounts []Account
	for _, ori := range accountsAll {
		if ori.ID == a.ID || ori.Address == a.Address {
			continue
		}
		newAccounts = append(newAccounts, ori)
	}
	return c.writeAllAccounts(newAccounts)
}

func (c *Client) GetAccount(id int) *Account {
	accountsAll := c.GetAllAccounts()

	for _, ori := range accountsAll {
		if ori.ID == id {
			return &ori
		}
	}
	return nil
}
