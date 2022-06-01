package main

import (
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/pkg/errors"
)

func createEthAccount() (Account, error) {
	mnemonic, err := hdwallet.NewMnemonic(256)
	if err != nil {
		return Account{}, errors.Wrap(err, "new mnemonic")
	}

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return Account{}, errors.Wrap(err, "new hd wallet")
	}

	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		return Account{}, errors.Wrap(err, "new account")
	}
	pri, err := wallet.PrivateKeyHex(account)
	if err != nil {
		return Account{}, errors.Wrap(err, "private key")
	}
	return Account{
		Address:    account.Address.Hex(),
		PrivateKey: pri,
		Mnemonic:   mnemonic,
	}, nil
}

type Account struct {
	Address    string
	PrivateKey string
	Mnemonic   string
}
