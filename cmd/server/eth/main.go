package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://10.0.0.64:7545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections

	// Instantiate the contract and display its name
	// NOTE update the deployment address!
	store, err := NewStorage(common.HexToAddress("0x3E43bc82b4fdc82C08dfeD0116c0565C1d0C65Eb"), client)
	if err != nil {
		log.Fatalf("Failed to instantiate Storage contract: %v", err)
	}

	privateKey, err := crypto.HexToECDSA("adef8c0ff6a27154e8ab1ad220f1a569f3e83c1df05a0e12f4018bba2849200c")
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		log.Fatal(err)
	}

	// Call the store() function
	tx, err := store.Store(auth, big.NewInt(420))
	if err != nil {
		log.Fatalf("Failed to update value: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())

	value, err := store.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to retrieve value: %v", err)
	}
	fmt.Println("Value: ", value)

}
