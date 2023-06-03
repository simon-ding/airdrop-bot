// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ZksBaseRegisterImplMetaData contains all meta data concerning the ZksBaseRegisterImpl contract.
var ZksBaseRegisterImplMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"addController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getDomainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getOwnedDomains\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPrimaryDomainId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getPrimaryDomainName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getUserFromPrimaryName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"myPrimaryDomainInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"nameExpires\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"reclaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"registerOnly\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"removeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseExtension\",\"type\":\"string\"}],\"name\":\"setBaseExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newBaseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractENS\",\"name\":\"_ens\",\"type\":\"address\"}],\"name\":\"setEns\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_baseNode\",\"type\":\"bytes32\"}],\"name\":\"setNode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"setPrimaryDomain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"setResolver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ZksBaseRegisterImplABI is the input ABI used to generate the binding from.
// Deprecated: Use ZksBaseRegisterImplMetaData.ABI instead.
var ZksBaseRegisterImplABI = ZksBaseRegisterImplMetaData.ABI

// ZksBaseRegisterImpl is an auto generated Go binding around an Ethereum contract.
type ZksBaseRegisterImpl struct {
	ZksBaseRegisterImplCaller     // Read-only binding to the contract
	ZksBaseRegisterImplTransactor // Write-only binding to the contract
	ZksBaseRegisterImplFilterer   // Log filterer for contract events
}

// ZksBaseRegisterImplCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZksBaseRegisterImplCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterImplTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZksBaseRegisterImplTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterImplFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZksBaseRegisterImplFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterImplSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZksBaseRegisterImplSession struct {
	Contract     *ZksBaseRegisterImpl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ZksBaseRegisterImplCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZksBaseRegisterImplCallerSession struct {
	Contract *ZksBaseRegisterImplCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ZksBaseRegisterImplTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZksBaseRegisterImplTransactorSession struct {
	Contract     *ZksBaseRegisterImplTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ZksBaseRegisterImplRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZksBaseRegisterImplRaw struct {
	Contract *ZksBaseRegisterImpl // Generic contract binding to access the raw methods on
}

// ZksBaseRegisterImplCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZksBaseRegisterImplCallerRaw struct {
	Contract *ZksBaseRegisterImplCaller // Generic read-only contract binding to access the raw methods on
}

// ZksBaseRegisterImplTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZksBaseRegisterImplTransactorRaw struct {
	Contract *ZksBaseRegisterImplTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZksBaseRegisterImpl creates a new instance of ZksBaseRegisterImpl, bound to a specific deployed contract.
func NewZksBaseRegisterImpl(address common.Address, backend bind.ContractBackend) (*ZksBaseRegisterImpl, error) {
	contract, err := bindZksBaseRegisterImpl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterImpl{ZksBaseRegisterImplCaller: ZksBaseRegisterImplCaller{contract: contract}, ZksBaseRegisterImplTransactor: ZksBaseRegisterImplTransactor{contract: contract}, ZksBaseRegisterImplFilterer: ZksBaseRegisterImplFilterer{contract: contract}}, nil
}

// NewZksBaseRegisterImplCaller creates a new read-only instance of ZksBaseRegisterImpl, bound to a specific deployed contract.
func NewZksBaseRegisterImplCaller(address common.Address, caller bind.ContractCaller) (*ZksBaseRegisterImplCaller, error) {
	contract, err := bindZksBaseRegisterImpl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterImplCaller{contract: contract}, nil
}

// NewZksBaseRegisterImplTransactor creates a new write-only instance of ZksBaseRegisterImpl, bound to a specific deployed contract.
func NewZksBaseRegisterImplTransactor(address common.Address, transactor bind.ContractTransactor) (*ZksBaseRegisterImplTransactor, error) {
	contract, err := bindZksBaseRegisterImpl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterImplTransactor{contract: contract}, nil
}

// NewZksBaseRegisterImplFilterer creates a new log filterer instance of ZksBaseRegisterImpl, bound to a specific deployed contract.
func NewZksBaseRegisterImplFilterer(address common.Address, filterer bind.ContractFilterer) (*ZksBaseRegisterImplFilterer, error) {
	contract, err := bindZksBaseRegisterImpl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterImplFilterer{contract: contract}, nil
}

// bindZksBaseRegisterImpl binds a generic wrapper to an already deployed contract.
func bindZksBaseRegisterImpl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZksBaseRegisterImplMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZksBaseRegisterImpl.Contract.ZksBaseRegisterImplCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.ZksBaseRegisterImplTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.ZksBaseRegisterImplTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZksBaseRegisterImpl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.contract.Transact(opts, method, params...)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) Available(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "available", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Available(id *big.Int) (bool, error) {
	return _ZksBaseRegisterImpl.Contract.Available(&_ZksBaseRegisterImpl.CallOpts, id)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 id) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) Available(id *big.Int) (bool, error) {
	return _ZksBaseRegisterImpl.Contract.Available(&_ZksBaseRegisterImpl.CallOpts, id)
}

// GetDomainName is a free data retrieval call binding the contract method 0x74b96931.
//
// Solidity: function getDomainName(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) GetDomainName(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "getDomainName", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetDomainName is a free data retrieval call binding the contract method 0x74b96931.
//
// Solidity: function getDomainName(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) GetDomainName(tokenId *big.Int) (string, error) {
	return _ZksBaseRegisterImpl.Contract.GetDomainName(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// GetDomainName is a free data retrieval call binding the contract method 0x74b96931.
//
// Solidity: function getDomainName(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) GetDomainName(tokenId *big.Int) (string, error) {
	return _ZksBaseRegisterImpl.Contract.GetDomainName(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// GetOwnedDomains is a free data retrieval call binding the contract method 0x1a6c821b.
//
// Solidity: function getOwnedDomains(address user) view returns(uint256[], string[])
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) GetOwnedDomains(opts *bind.CallOpts, user common.Address) ([]*big.Int, []string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "getOwnedDomains", user)

	if err != nil {
		return *new([]*big.Int), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)

	return out0, out1, err

}

// GetOwnedDomains is a free data retrieval call binding the contract method 0x1a6c821b.
//
// Solidity: function getOwnedDomains(address user) view returns(uint256[], string[])
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) GetOwnedDomains(user common.Address) ([]*big.Int, []string, error) {
	return _ZksBaseRegisterImpl.Contract.GetOwnedDomains(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetOwnedDomains is a free data retrieval call binding the contract method 0x1a6c821b.
//
// Solidity: function getOwnedDomains(address user) view returns(uint256[], string[])
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) GetOwnedDomains(user common.Address) ([]*big.Int, []string, error) {
	return _ZksBaseRegisterImpl.Contract.GetOwnedDomains(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetPrimaryDomainId is a free data retrieval call binding the contract method 0x7bd4a448.
//
// Solidity: function getPrimaryDomainId(address user) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) GetPrimaryDomainId(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "getPrimaryDomainId", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrimaryDomainId is a free data retrieval call binding the contract method 0x7bd4a448.
//
// Solidity: function getPrimaryDomainId(address user) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) GetPrimaryDomainId(user common.Address) (*big.Int, error) {
	return _ZksBaseRegisterImpl.Contract.GetPrimaryDomainId(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetPrimaryDomainId is a free data retrieval call binding the contract method 0x7bd4a448.
//
// Solidity: function getPrimaryDomainId(address user) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) GetPrimaryDomainId(user common.Address) (*big.Int, error) {
	return _ZksBaseRegisterImpl.Contract.GetPrimaryDomainId(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetPrimaryDomainName is a free data retrieval call binding the contract method 0xd72ede13.
//
// Solidity: function getPrimaryDomainName(address user) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) GetPrimaryDomainName(opts *bind.CallOpts, user common.Address) (string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "getPrimaryDomainName", user)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetPrimaryDomainName is a free data retrieval call binding the contract method 0xd72ede13.
//
// Solidity: function getPrimaryDomainName(address user) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) GetPrimaryDomainName(user common.Address) (string, error) {
	return _ZksBaseRegisterImpl.Contract.GetPrimaryDomainName(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetPrimaryDomainName is a free data retrieval call binding the contract method 0xd72ede13.
//
// Solidity: function getPrimaryDomainName(address user) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) GetPrimaryDomainName(user common.Address) (string, error) {
	return _ZksBaseRegisterImpl.Contract.GetPrimaryDomainName(&_ZksBaseRegisterImpl.CallOpts, user)
}

// GetUserFromPrimaryName is a free data retrieval call binding the contract method 0xb18b9262.
//
// Solidity: function getUserFromPrimaryName(string name) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) GetUserFromPrimaryName(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "getUserFromPrimaryName", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserFromPrimaryName is a free data retrieval call binding the contract method 0xb18b9262.
//
// Solidity: function getUserFromPrimaryName(string name) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) GetUserFromPrimaryName(name string) (common.Address, error) {
	return _ZksBaseRegisterImpl.Contract.GetUserFromPrimaryName(&_ZksBaseRegisterImpl.CallOpts, name)
}

// GetUserFromPrimaryName is a free data retrieval call binding the contract method 0xb18b9262.
//
// Solidity: function getUserFromPrimaryName(string name) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) GetUserFromPrimaryName(name string) (common.Address, error) {
	return _ZksBaseRegisterImpl.Contract.GetUserFromPrimaryName(&_ZksBaseRegisterImpl.CallOpts, name)
}

// MyPrimaryDomainInfo is a free data retrieval call binding the contract method 0xbc6e71b9.
//
// Solidity: function myPrimaryDomainInfo() view returns(uint256, string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) MyPrimaryDomainInfo(opts *bind.CallOpts) (*big.Int, string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "myPrimaryDomainInfo")

	if err != nil {
		return *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// MyPrimaryDomainInfo is a free data retrieval call binding the contract method 0xbc6e71b9.
//
// Solidity: function myPrimaryDomainInfo() view returns(uint256, string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) MyPrimaryDomainInfo() (*big.Int, string, error) {
	return _ZksBaseRegisterImpl.Contract.MyPrimaryDomainInfo(&_ZksBaseRegisterImpl.CallOpts)
}

// MyPrimaryDomainInfo is a free data retrieval call binding the contract method 0xbc6e71b9.
//
// Solidity: function myPrimaryDomainInfo() view returns(uint256, string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) MyPrimaryDomainInfo() (*big.Int, string, error) {
	return _ZksBaseRegisterImpl.Contract.MyPrimaryDomainInfo(&_ZksBaseRegisterImpl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Name() (string, error) {
	return _ZksBaseRegisterImpl.Contract.Name(&_ZksBaseRegisterImpl.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) Name() (string, error) {
	return _ZksBaseRegisterImpl.Contract.Name(&_ZksBaseRegisterImpl.CallOpts)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) NameExpires(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "nameExpires", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _ZksBaseRegisterImpl.Contract.NameExpires(&_ZksBaseRegisterImpl.CallOpts, id)
}

// NameExpires is a free data retrieval call binding the contract method 0xd6e4fa86.
//
// Solidity: function nameExpires(uint256 id) view returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) NameExpires(id *big.Int) (*big.Int, error) {
	return _ZksBaseRegisterImpl.Contract.NameExpires(&_ZksBaseRegisterImpl.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ZksBaseRegisterImpl.Contract.OwnerOf(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ZksBaseRegisterImpl.Contract.OwnerOf(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZksBaseRegisterImpl.Contract.SupportsInterface(&_ZksBaseRegisterImpl.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) view returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZksBaseRegisterImpl.Contract.SupportsInterface(&_ZksBaseRegisterImpl.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Symbol() (string, error) {
	return _ZksBaseRegisterImpl.Contract.Symbol(&_ZksBaseRegisterImpl.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) Symbol() (string, error) {
	return _ZksBaseRegisterImpl.Contract.Symbol(&_ZksBaseRegisterImpl.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ZksBaseRegisterImpl.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ZksBaseRegisterImpl.Contract.TokenURI(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ZksBaseRegisterImpl.Contract.TokenURI(&_ZksBaseRegisterImpl.CallOpts, tokenId)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) AddController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "addController", controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.AddController(&_ZksBaseRegisterImpl.TransactOpts, controller)
}

// AddController is a paid mutator transaction binding the contract method 0xa7fc7a07.
//
// Solidity: function addController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) AddController(controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.AddController(&_ZksBaseRegisterImpl.TransactOpts, controller)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) Reclaim(opts *bind.TransactOpts, id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "reclaim", id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Reclaim(&_ZksBaseRegisterImpl.TransactOpts, id, owner)
}

// Reclaim is a paid mutator transaction binding the contract method 0x28ed4f6c.
//
// Solidity: function reclaim(uint256 id, address owner) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) Reclaim(id *big.Int, owner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Reclaim(&_ZksBaseRegisterImpl.TransactOpts, id, owner)
}

// Register is a paid mutator transaction binding the contract method 0x8d7714b4.
//
// Solidity: function register(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) Register(opts *bind.TransactOpts, name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "register", name, id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0x8d7714b4.
//
// Solidity: function register(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Register(name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Register(&_ZksBaseRegisterImpl.TransactOpts, name, id, owner, duration)
}

// Register is a paid mutator transaction binding the contract method 0x8d7714b4.
//
// Solidity: function register(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) Register(name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Register(&_ZksBaseRegisterImpl.TransactOpts, name, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x11020055.
//
// Solidity: function registerOnly(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) RegisterOnly(opts *bind.TransactOpts, name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "registerOnly", name, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x11020055.
//
// Solidity: function registerOnly(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) RegisterOnly(name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.RegisterOnly(&_ZksBaseRegisterImpl.TransactOpts, name, id, owner, duration)
}

// RegisterOnly is a paid mutator transaction binding the contract method 0x11020055.
//
// Solidity: function registerOnly(string name, uint256 id, address owner, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) RegisterOnly(name string, id *big.Int, owner common.Address, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.RegisterOnly(&_ZksBaseRegisterImpl.TransactOpts, name, id, owner, duration)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) RemoveController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "removeController", controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.RemoveController(&_ZksBaseRegisterImpl.TransactOpts, controller)
}

// RemoveController is a paid mutator transaction binding the contract method 0xf6a74ed7.
//
// Solidity: function removeController(address controller) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) RemoveController(controller common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.RemoveController(&_ZksBaseRegisterImpl.TransactOpts, controller)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) Renew(opts *bind.TransactOpts, id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "renew", id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Renew(&_ZksBaseRegisterImpl.TransactOpts, id, duration)
}

// Renew is a paid mutator transaction binding the contract method 0xc475abff.
//
// Solidity: function renew(uint256 id, uint256 duration) returns(uint256)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) Renew(id *big.Int, duration *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.Renew(&_ZksBaseRegisterImpl.TransactOpts, id, duration)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetBaseExtension(opts *bind.TransactOpts, _newBaseExtension string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setBaseExtension", _newBaseExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetBaseExtension(_newBaseExtension string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetBaseExtension(&_ZksBaseRegisterImpl.TransactOpts, _newBaseExtension)
}

// SetBaseExtension is a paid mutator transaction binding the contract method 0xda3ef23f.
//
// Solidity: function setBaseExtension(string _newBaseExtension) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetBaseExtension(_newBaseExtension string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetBaseExtension(&_ZksBaseRegisterImpl.TransactOpts, _newBaseExtension)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetBaseURI(opts *bind.TransactOpts, _newBaseURI string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setBaseURI", _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetBaseURI(&_ZksBaseRegisterImpl.TransactOpts, _newBaseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _newBaseURI) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetBaseURI(_newBaseURI string) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetBaseURI(&_ZksBaseRegisterImpl.TransactOpts, _newBaseURI)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(address _ens) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetEns(opts *bind.TransactOpts, _ens common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setEns", _ens)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(address _ens) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetEns(_ens common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetEns(&_ZksBaseRegisterImpl.TransactOpts, _ens)
}

// SetEns is a paid mutator transaction binding the contract method 0x6e8f2be0.
//
// Solidity: function setEns(address _ens) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetEns(_ens common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetEns(&_ZksBaseRegisterImpl.TransactOpts, _ens)
}

// SetNode is a paid mutator transaction binding the contract method 0x4cc23928.
//
// Solidity: function setNode(bytes32 _baseNode) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetNode(opts *bind.TransactOpts, _baseNode [32]byte) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setNode", _baseNode)
}

// SetNode is a paid mutator transaction binding the contract method 0x4cc23928.
//
// Solidity: function setNode(bytes32 _baseNode) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetNode(_baseNode [32]byte) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetNode(&_ZksBaseRegisterImpl.TransactOpts, _baseNode)
}

// SetNode is a paid mutator transaction binding the contract method 0x4cc23928.
//
// Solidity: function setNode(bytes32 _baseNode) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetNode(_baseNode [32]byte) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetNode(&_ZksBaseRegisterImpl.TransactOpts, _baseNode)
}

// SetPrimaryDomain is a paid mutator transaction binding the contract method 0x0d537e8d.
//
// Solidity: function setPrimaryDomain(uint256 tokenId) returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetPrimaryDomain(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setPrimaryDomain", tokenId)
}

// SetPrimaryDomain is a paid mutator transaction binding the contract method 0x0d537e8d.
//
// Solidity: function setPrimaryDomain(uint256 tokenId) returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetPrimaryDomain(tokenId *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetPrimaryDomain(&_ZksBaseRegisterImpl.TransactOpts, tokenId)
}

// SetPrimaryDomain is a paid mutator transaction binding the contract method 0x0d537e8d.
//
// Solidity: function setPrimaryDomain(uint256 tokenId) returns(bool)
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetPrimaryDomain(tokenId *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetPrimaryDomain(&_ZksBaseRegisterImpl.TransactOpts, tokenId)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactor) SetResolver(opts *bind.TransactOpts, resolver common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.contract.Transact(opts, "setResolver", resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetResolver(&_ZksBaseRegisterImpl.TransactOpts, resolver)
}

// SetResolver is a paid mutator transaction binding the contract method 0x4e543b26.
//
// Solidity: function setResolver(address resolver) returns()
func (_ZksBaseRegisterImpl *ZksBaseRegisterImplTransactorSession) SetResolver(resolver common.Address) (*types.Transaction, error) {
	return _ZksBaseRegisterImpl.Contract.SetResolver(&_ZksBaseRegisterImpl.TransactOpts, resolver)
}
