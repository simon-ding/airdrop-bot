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

// DistributionPoolInfoView is an auto generated low-level Go binding around an user-defined struct.
type DistributionPoolInfoView struct {
	MaxToken      *big.Int
	InitClaim     *big.Int
	CurrentClaim  *big.Int
	Claimed       bool
	InviteRewards *big.Int
	InviteUsers   *big.Int
	ClaimedSupply *big.Int
	ClaimedCount  *big.Int
}

// AidogeMetaData contains all meta data concerning the Aidoge contract.
var AidogeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"INIT_CLAIM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_ADDRESSES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TOKEN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_claimedUser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"_usedNonce\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"addSigner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"canClaimAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"nonce\",\"type\":\"uint128\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedPercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimedSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"name\":\"delSigner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getInfoView\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"maxToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initClaim\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentClaim\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"claimed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"inviteRewards\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"inviteUsers\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimedCount\",\"type\":\"uint256\"}],\"internalType\":\"structDistributionPool.InfoView\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSigners\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"ret\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviteRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"inviteUsers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AidogeABI is the input ABI used to generate the binding from.
// Deprecated: Use AidogeMetaData.ABI instead.
var AidogeABI = AidogeMetaData.ABI

// Aidoge is an auto generated Go binding around an Ethereum contract.
type Aidoge struct {
	AidogeCaller     // Read-only binding to the contract
	AidogeTransactor // Write-only binding to the contract
	AidogeFilterer   // Log filterer for contract events
}

// AidogeCaller is an auto generated read-only Go binding around an Ethereum contract.
type AidogeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AidogeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AidogeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AidogeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AidogeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AidogeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AidogeSession struct {
	Contract     *Aidoge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AidogeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AidogeCallerSession struct {
	Contract *AidogeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AidogeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AidogeTransactorSession struct {
	Contract     *AidogeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AidogeRaw is an auto generated low-level Go binding around an Ethereum contract.
type AidogeRaw struct {
	Contract *Aidoge // Generic contract binding to access the raw methods on
}

// AidogeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AidogeCallerRaw struct {
	Contract *AidogeCaller // Generic read-only contract binding to access the raw methods on
}

// AidogeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AidogeTransactorRaw struct {
	Contract *AidogeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAidoge creates a new instance of Aidoge, bound to a specific deployed contract.
func NewAidoge(address common.Address, backend bind.ContractBackend) (*Aidoge, error) {
	contract, err := bindAidoge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aidoge{AidogeCaller: AidogeCaller{contract: contract}, AidogeTransactor: AidogeTransactor{contract: contract}, AidogeFilterer: AidogeFilterer{contract: contract}}, nil
}

// NewAidogeCaller creates a new read-only instance of Aidoge, bound to a specific deployed contract.
func NewAidogeCaller(address common.Address, caller bind.ContractCaller) (*AidogeCaller, error) {
	contract, err := bindAidoge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AidogeCaller{contract: contract}, nil
}

// NewAidogeTransactor creates a new write-only instance of Aidoge, bound to a specific deployed contract.
func NewAidogeTransactor(address common.Address, transactor bind.ContractTransactor) (*AidogeTransactor, error) {
	contract, err := bindAidoge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AidogeTransactor{contract: contract}, nil
}

// NewAidogeFilterer creates a new log filterer instance of Aidoge, bound to a specific deployed contract.
func NewAidogeFilterer(address common.Address, filterer bind.ContractFilterer) (*AidogeFilterer, error) {
	contract, err := bindAidoge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AidogeFilterer{contract: contract}, nil
}

// bindAidoge binds a generic wrapper to an already deployed contract.
func bindAidoge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AidogeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aidoge *AidogeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aidoge.Contract.AidogeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aidoge *AidogeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aidoge.Contract.AidogeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aidoge *AidogeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aidoge.Contract.AidogeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aidoge *AidogeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aidoge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aidoge *AidogeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aidoge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aidoge *AidogeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aidoge.Contract.contract.Transact(opts, method, params...)
}

// INITCLAIM is a free data retrieval call binding the contract method 0xe14204b5.
//
// Solidity: function INIT_CLAIM() view returns(uint256)
func (_Aidoge *AidogeCaller) INITCLAIM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "INIT_CLAIM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITCLAIM is a free data retrieval call binding the contract method 0xe14204b5.
//
// Solidity: function INIT_CLAIM() view returns(uint256)
func (_Aidoge *AidogeSession) INITCLAIM() (*big.Int, error) {
	return _Aidoge.Contract.INITCLAIM(&_Aidoge.CallOpts)
}

// INITCLAIM is a free data retrieval call binding the contract method 0xe14204b5.
//
// Solidity: function INIT_CLAIM() view returns(uint256)
func (_Aidoge *AidogeCallerSession) INITCLAIM() (*big.Int, error) {
	return _Aidoge.Contract.INITCLAIM(&_Aidoge.CallOpts)
}

// MAXADDRESSES is a free data retrieval call binding the contract method 0x634da787.
//
// Solidity: function MAX_ADDRESSES() view returns(uint256)
func (_Aidoge *AidogeCaller) MAXADDRESSES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "MAX_ADDRESSES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXADDRESSES is a free data retrieval call binding the contract method 0x634da787.
//
// Solidity: function MAX_ADDRESSES() view returns(uint256)
func (_Aidoge *AidogeSession) MAXADDRESSES() (*big.Int, error) {
	return _Aidoge.Contract.MAXADDRESSES(&_Aidoge.CallOpts)
}

// MAXADDRESSES is a free data retrieval call binding the contract method 0x634da787.
//
// Solidity: function MAX_ADDRESSES() view returns(uint256)
func (_Aidoge *AidogeCallerSession) MAXADDRESSES() (*big.Int, error) {
	return _Aidoge.Contract.MAXADDRESSES(&_Aidoge.CallOpts)
}

// MAXTOKEN is a free data retrieval call binding the contract method 0x6e1bd323.
//
// Solidity: function MAX_TOKEN() view returns(uint256)
func (_Aidoge *AidogeCaller) MAXTOKEN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "MAX_TOKEN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOKEN is a free data retrieval call binding the contract method 0x6e1bd323.
//
// Solidity: function MAX_TOKEN() view returns(uint256)
func (_Aidoge *AidogeSession) MAXTOKEN() (*big.Int, error) {
	return _Aidoge.Contract.MAXTOKEN(&_Aidoge.CallOpts)
}

// MAXTOKEN is a free data retrieval call binding the contract method 0x6e1bd323.
//
// Solidity: function MAX_TOKEN() view returns(uint256)
func (_Aidoge *AidogeCallerSession) MAXTOKEN() (*big.Int, error) {
	return _Aidoge.Contract.MAXTOKEN(&_Aidoge.CallOpts)
}

// ClaimedUser is a free data retrieval call binding the contract method 0x97b63b04.
//
// Solidity: function _claimedUser(address ) view returns(bool)
func (_Aidoge *AidogeCaller) ClaimedUser(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "_claimedUser", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ClaimedUser is a free data retrieval call binding the contract method 0x97b63b04.
//
// Solidity: function _claimedUser(address ) view returns(bool)
func (_Aidoge *AidogeSession) ClaimedUser(arg0 common.Address) (bool, error) {
	return _Aidoge.Contract.ClaimedUser(&_Aidoge.CallOpts, arg0)
}

// ClaimedUser is a free data retrieval call binding the contract method 0x97b63b04.
//
// Solidity: function _claimedUser(address ) view returns(bool)
func (_Aidoge *AidogeCallerSession) ClaimedUser(arg0 common.Address) (bool, error) {
	return _Aidoge.Contract.ClaimedUser(&_Aidoge.CallOpts, arg0)
}

// UsedNonce is a free data retrieval call binding the contract method 0x25d3ab5c.
//
// Solidity: function _usedNonce(uint256 ) view returns(bool)
func (_Aidoge *AidogeCaller) UsedNonce(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "_usedNonce", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNonce is a free data retrieval call binding the contract method 0x25d3ab5c.
//
// Solidity: function _usedNonce(uint256 ) view returns(bool)
func (_Aidoge *AidogeSession) UsedNonce(arg0 *big.Int) (bool, error) {
	return _Aidoge.Contract.UsedNonce(&_Aidoge.CallOpts, arg0)
}

// UsedNonce is a free data retrieval call binding the contract method 0x25d3ab5c.
//
// Solidity: function _usedNonce(uint256 ) view returns(bool)
func (_Aidoge *AidogeCallerSession) UsedNonce(arg0 *big.Int) (bool, error) {
	return _Aidoge.Contract.UsedNonce(&_Aidoge.CallOpts, arg0)
}

// CanClaimAmount is a free data retrieval call binding the contract method 0x8431165e.
//
// Solidity: function canClaimAmount() view returns(uint256)
func (_Aidoge *AidogeCaller) CanClaimAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "canClaimAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CanClaimAmount is a free data retrieval call binding the contract method 0x8431165e.
//
// Solidity: function canClaimAmount() view returns(uint256)
func (_Aidoge *AidogeSession) CanClaimAmount() (*big.Int, error) {
	return _Aidoge.Contract.CanClaimAmount(&_Aidoge.CallOpts)
}

// CanClaimAmount is a free data retrieval call binding the contract method 0x8431165e.
//
// Solidity: function canClaimAmount() view returns(uint256)
func (_Aidoge *AidogeCallerSession) CanClaimAmount() (*big.Int, error) {
	return _Aidoge.Contract.CanClaimAmount(&_Aidoge.CallOpts)
}

// ClaimedCount is a free data retrieval call binding the contract method 0xc08fa1a4.
//
// Solidity: function claimedCount() view returns(uint256)
func (_Aidoge *AidogeCaller) ClaimedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "claimedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedCount is a free data retrieval call binding the contract method 0xc08fa1a4.
//
// Solidity: function claimedCount() view returns(uint256)
func (_Aidoge *AidogeSession) ClaimedCount() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedCount(&_Aidoge.CallOpts)
}

// ClaimedCount is a free data retrieval call binding the contract method 0xc08fa1a4.
//
// Solidity: function claimedCount() view returns(uint256)
func (_Aidoge *AidogeCallerSession) ClaimedCount() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedCount(&_Aidoge.CallOpts)
}

// ClaimedPercentage is a free data retrieval call binding the contract method 0xef45e240.
//
// Solidity: function claimedPercentage() view returns(uint256)
func (_Aidoge *AidogeCaller) ClaimedPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "claimedPercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedPercentage is a free data retrieval call binding the contract method 0xef45e240.
//
// Solidity: function claimedPercentage() view returns(uint256)
func (_Aidoge *AidogeSession) ClaimedPercentage() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedPercentage(&_Aidoge.CallOpts)
}

// ClaimedPercentage is a free data retrieval call binding the contract method 0xef45e240.
//
// Solidity: function claimedPercentage() view returns(uint256)
func (_Aidoge *AidogeCallerSession) ClaimedPercentage() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedPercentage(&_Aidoge.CallOpts)
}

// ClaimedSupply is a free data retrieval call binding the contract method 0xbfc2aa2a.
//
// Solidity: function claimedSupply() view returns(uint256)
func (_Aidoge *AidogeCaller) ClaimedSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "claimedSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimedSupply is a free data retrieval call binding the contract method 0xbfc2aa2a.
//
// Solidity: function claimedSupply() view returns(uint256)
func (_Aidoge *AidogeSession) ClaimedSupply() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedSupply(&_Aidoge.CallOpts)
}

// ClaimedSupply is a free data retrieval call binding the contract method 0xbfc2aa2a.
//
// Solidity: function claimedSupply() view returns(uint256)
func (_Aidoge *AidogeCallerSession) ClaimedSupply() (*big.Int, error) {
	return _Aidoge.Contract.ClaimedSupply(&_Aidoge.CallOpts)
}

// GetInfoView is a free data retrieval call binding the contract method 0xfd40dbe3.
//
// Solidity: function getInfoView(address user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256))
func (_Aidoge *AidogeCaller) GetInfoView(opts *bind.CallOpts, user common.Address) (DistributionPoolInfoView, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "getInfoView", user)

	if err != nil {
		return *new(DistributionPoolInfoView), err
	}

	out0 := *abi.ConvertType(out[0], new(DistributionPoolInfoView)).(*DistributionPoolInfoView)

	return out0, err

}

// GetInfoView is a free data retrieval call binding the contract method 0xfd40dbe3.
//
// Solidity: function getInfoView(address user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256))
func (_Aidoge *AidogeSession) GetInfoView(user common.Address) (DistributionPoolInfoView, error) {
	return _Aidoge.Contract.GetInfoView(&_Aidoge.CallOpts, user)
}

// GetInfoView is a free data retrieval call binding the contract method 0xfd40dbe3.
//
// Solidity: function getInfoView(address user) view returns((uint256,uint256,uint256,bool,uint256,uint256,uint256,uint256))
func (_Aidoge *AidogeCallerSession) GetInfoView(user common.Address) (DistributionPoolInfoView, error) {
	return _Aidoge.Contract.GetInfoView(&_Aidoge.CallOpts, user)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[] ret)
func (_Aidoge *AidogeCaller) GetSigners(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "getSigners")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[] ret)
func (_Aidoge *AidogeSession) GetSigners() ([]common.Address, error) {
	return _Aidoge.Contract.GetSigners(&_Aidoge.CallOpts)
}

// GetSigners is a free data retrieval call binding the contract method 0x94cf795e.
//
// Solidity: function getSigners() view returns(address[] ret)
func (_Aidoge *AidogeCallerSession) GetSigners() ([]common.Address, error) {
	return _Aidoge.Contract.GetSigners(&_Aidoge.CallOpts)
}

// InviteRewards is a free data retrieval call binding the contract method 0x8d2c5d44.
//
// Solidity: function inviteRewards(address ) view returns(uint256)
func (_Aidoge *AidogeCaller) InviteRewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "inviteRewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteRewards is a free data retrieval call binding the contract method 0x8d2c5d44.
//
// Solidity: function inviteRewards(address ) view returns(uint256)
func (_Aidoge *AidogeSession) InviteRewards(arg0 common.Address) (*big.Int, error) {
	return _Aidoge.Contract.InviteRewards(&_Aidoge.CallOpts, arg0)
}

// InviteRewards is a free data retrieval call binding the contract method 0x8d2c5d44.
//
// Solidity: function inviteRewards(address ) view returns(uint256)
func (_Aidoge *AidogeCallerSession) InviteRewards(arg0 common.Address) (*big.Int, error) {
	return _Aidoge.Contract.InviteRewards(&_Aidoge.CallOpts, arg0)
}

// InviteUsers is a free data retrieval call binding the contract method 0x5b8aca6c.
//
// Solidity: function inviteUsers(address ) view returns(uint256)
func (_Aidoge *AidogeCaller) InviteUsers(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "inviteUsers", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InviteUsers is a free data retrieval call binding the contract method 0x5b8aca6c.
//
// Solidity: function inviteUsers(address ) view returns(uint256)
func (_Aidoge *AidogeSession) InviteUsers(arg0 common.Address) (*big.Int, error) {
	return _Aidoge.Contract.InviteUsers(&_Aidoge.CallOpts, arg0)
}

// InviteUsers is a free data retrieval call binding the contract method 0x5b8aca6c.
//
// Solidity: function inviteUsers(address ) view returns(uint256)
func (_Aidoge *AidogeCallerSession) InviteUsers(arg0 common.Address) (*big.Int, error) {
	return _Aidoge.Contract.InviteUsers(&_Aidoge.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aidoge *AidogeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aidoge *AidogeSession) Owner() (common.Address, error) {
	return _Aidoge.Contract.Owner(&_Aidoge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aidoge *AidogeCallerSession) Owner() (common.Address, error) {
	return _Aidoge.Contract.Owner(&_Aidoge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Aidoge *AidogeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aidoge.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Aidoge *AidogeSession) Token() (common.Address, error) {
	return _Aidoge.Contract.Token(&_Aidoge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Aidoge *AidogeCallerSession) Token() (common.Address, error) {
	return _Aidoge.Contract.Token(&_Aidoge.CallOpts)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address val) returns()
func (_Aidoge *AidogeTransactor) AddSigner(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "addSigner", val)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address val) returns()
func (_Aidoge *AidogeSession) AddSigner(val common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.AddSigner(&_Aidoge.TransactOpts, val)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address val) returns()
func (_Aidoge *AidogeTransactorSession) AddSigner(val common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.AddSigner(&_Aidoge.TransactOpts, val)
}

// Claim is a paid mutator transaction binding the contract method 0xf6303ef5.
//
// Solidity: function claim(uint128 nonce, bytes signature, address referrer) returns()
func (_Aidoge *AidogeTransactor) Claim(opts *bind.TransactOpts, nonce *big.Int, signature []byte, referrer common.Address) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "claim", nonce, signature, referrer)
}

// Claim is a paid mutator transaction binding the contract method 0xf6303ef5.
//
// Solidity: function claim(uint128 nonce, bytes signature, address referrer) returns()
func (_Aidoge *AidogeSession) Claim(nonce *big.Int, signature []byte, referrer common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.Claim(&_Aidoge.TransactOpts, nonce, signature, referrer)
}

// Claim is a paid mutator transaction binding the contract method 0xf6303ef5.
//
// Solidity: function claim(uint128 nonce, bytes signature, address referrer) returns()
func (_Aidoge *AidogeTransactorSession) Claim(nonce *big.Int, signature []byte, referrer common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.Claim(&_Aidoge.TransactOpts, nonce, signature, referrer)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address signer) returns(bool)
func (_Aidoge *AidogeTransactor) DelSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "delSigner", signer)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address signer) returns(bool)
func (_Aidoge *AidogeSession) DelSigner(signer common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.DelSigner(&_Aidoge.TransactOpts, signer)
}

// DelSigner is a paid mutator transaction binding the contract method 0x883b524f.
//
// Solidity: function delSigner(address signer) returns(bool)
func (_Aidoge *AidogeTransactorSession) DelSigner(signer common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.DelSigner(&_Aidoge.TransactOpts, signer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_Aidoge *AidogeTransactor) Initialize(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "initialize", token_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_Aidoge *AidogeSession) Initialize(token_ common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.Initialize(&_Aidoge.TransactOpts, token_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address token_) returns()
func (_Aidoge *AidogeTransactorSession) Initialize(token_ common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.Initialize(&_Aidoge.TransactOpts, token_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aidoge *AidogeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aidoge *AidogeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aidoge.Contract.RenounceOwnership(&_Aidoge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aidoge *AidogeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aidoge.Contract.RenounceOwnership(&_Aidoge.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aidoge *AidogeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aidoge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aidoge *AidogeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.TransferOwnership(&_Aidoge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aidoge *AidogeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aidoge.Contract.TransferOwnership(&_Aidoge.TransactOpts, newOwner)
}

// AidogeClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the Aidoge contract.
type AidogeClaimIterator struct {
	Event *AidogeClaim // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AidogeClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AidogeClaim)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AidogeClaim)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AidogeClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AidogeClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AidogeClaim represents a Claim event raised by the Aidoge contract.
type AidogeClaim struct {
	User      common.Address
	Nonce     *big.Int
	Amount    *big.Int
	Referrer  common.Address
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x43df200563d66c23d96ffc5da991b910cb5fc0e45018c64c7fe558830feec392.
//
// Solidity: event Claim(address indexed user, uint128 nonce, uint256 amount, address referrer, uint256 timestamp)
func (_Aidoge *AidogeFilterer) FilterClaim(opts *bind.FilterOpts, user []common.Address) (*AidogeClaimIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Aidoge.contract.FilterLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return &AidogeClaimIterator{contract: _Aidoge.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x43df200563d66c23d96ffc5da991b910cb5fc0e45018c64c7fe558830feec392.
//
// Solidity: event Claim(address indexed user, uint128 nonce, uint256 amount, address referrer, uint256 timestamp)
func (_Aidoge *AidogeFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AidogeClaim, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Aidoge.contract.WatchLogs(opts, "Claim", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AidogeClaim)
				if err := _Aidoge.contract.UnpackLog(event, "Claim", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseClaim is a log parse operation binding the contract event 0x43df200563d66c23d96ffc5da991b910cb5fc0e45018c64c7fe558830feec392.
//
// Solidity: event Claim(address indexed user, uint128 nonce, uint256 amount, address referrer, uint256 timestamp)
func (_Aidoge *AidogeFilterer) ParseClaim(log types.Log) (*AidogeClaim, error) {
	event := new(AidogeClaim)
	if err := _Aidoge.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AidogeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Aidoge contract.
type AidogeInitializedIterator struct {
	Event *AidogeInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AidogeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AidogeInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AidogeInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AidogeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AidogeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AidogeInitialized represents a Initialized event raised by the Aidoge contract.
type AidogeInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aidoge *AidogeFilterer) FilterInitialized(opts *bind.FilterOpts) (*AidogeInitializedIterator, error) {

	logs, sub, err := _Aidoge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AidogeInitializedIterator{contract: _Aidoge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aidoge *AidogeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AidogeInitialized) (event.Subscription, error) {

	logs, sub, err := _Aidoge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AidogeInitialized)
				if err := _Aidoge.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Aidoge *AidogeFilterer) ParseInitialized(log types.Log) (*AidogeInitialized, error) {
	event := new(AidogeInitialized)
	if err := _Aidoge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AidogeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aidoge contract.
type AidogeOwnershipTransferredIterator struct {
	Event *AidogeOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AidogeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AidogeOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AidogeOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AidogeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AidogeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AidogeOwnershipTransferred represents a OwnershipTransferred event raised by the Aidoge contract.
type AidogeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aidoge *AidogeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AidogeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aidoge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AidogeOwnershipTransferredIterator{contract: _Aidoge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aidoge *AidogeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AidogeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aidoge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AidogeOwnershipTransferred)
				if err := _Aidoge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aidoge *AidogeFilterer) ParseOwnershipTransferred(log types.Log) (*AidogeOwnershipTransferred, error) {
	event := new(AidogeOwnershipTransferred)
	if err := _Aidoge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
