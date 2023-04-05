package abi

//go:generate abigen --abi=erc20.abi --pkg=abi --type erc20 --out=erc20.go
//go:generate abigen --abi=arbtrum_token_distributor.abi --pkg=abi --type ArbToken --out=arbtrum_token_distributor.go 