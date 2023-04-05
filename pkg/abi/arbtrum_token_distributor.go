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

// ArbTokenMetaData contains all meta data concerning the ArbToken contract.
var ArbTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodEnd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegateTo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CanClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HasClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSweepReceiver\",\"type\":\"address\"}],\"name\":\"SweepReceiverSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claimAndDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_claimableAmount\",\"type\":\"uint256[]\"}],\"name\":\"setRecipients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"}],\"name\":\"setSweepReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArbTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbTokenMetaData.ABI instead.
var ArbTokenABI = ArbTokenMetaData.ABI

// ArbToken is an auto generated Go binding around an Ethereum contract.
type ArbToken struct {
	ArbTokenCaller     // Read-only binding to the contract
	ArbTokenTransactor // Write-only binding to the contract
	ArbTokenFilterer   // Log filterer for contract events
}

// ArbTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbTokenSession struct {
	Contract     *ArbToken         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbTokenCallerSession struct {
	Contract *ArbTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ArbTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTokenTransactorSession struct {
	Contract     *ArbTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ArbTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbTokenRaw struct {
	Contract *ArbToken // Generic contract binding to access the raw methods on
}

// ArbTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbTokenCallerRaw struct {
	Contract *ArbTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTokenTransactorRaw struct {
	Contract *ArbTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArbToken creates a new instance of ArbToken, bound to a specific deployed contract.
func NewArbToken(address common.Address, backend bind.ContractBackend) (*ArbToken, error) {
	contract, err := bindArbToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArbToken{ArbTokenCaller: ArbTokenCaller{contract: contract}, ArbTokenTransactor: ArbTokenTransactor{contract: contract}, ArbTokenFilterer: ArbTokenFilterer{contract: contract}}, nil
}

// NewArbTokenCaller creates a new read-only instance of ArbToken, bound to a specific deployed contract.
func NewArbTokenCaller(address common.Address, caller bind.ContractCaller) (*ArbTokenCaller, error) {
	contract, err := bindArbToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTokenCaller{contract: contract}, nil
}

// NewArbTokenTransactor creates a new write-only instance of ArbToken, bound to a specific deployed contract.
func NewArbTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTokenTransactor, error) {
	contract, err := bindArbToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTokenTransactor{contract: contract}, nil
}

// NewArbTokenFilterer creates a new log filterer instance of ArbToken, bound to a specific deployed contract.
func NewArbTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbTokenFilterer, error) {
	contract, err := bindArbToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbTokenFilterer{contract: contract}, nil
}

// bindArbToken binds a generic wrapper to an already deployed contract.
func bindArbToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbToken *ArbTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbToken.Contract.ArbTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbToken *ArbTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbToken.Contract.ArbTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbToken *ArbTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbToken.Contract.ArbTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArbToken *ArbTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArbToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArbToken *ArbTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArbToken *ArbTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArbToken.Contract.contract.Transact(opts, method, params...)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_ArbToken *ArbTokenCaller) ClaimPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "claimPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_ArbToken *ArbTokenSession) ClaimPeriodEnd() (*big.Int, error) {
	return _ArbToken.Contract.ClaimPeriodEnd(&_ArbToken.CallOpts)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_ArbToken *ArbTokenCallerSession) ClaimPeriodEnd() (*big.Int, error) {
	return _ArbToken.Contract.ClaimPeriodEnd(&_ArbToken.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_ArbToken *ArbTokenCaller) ClaimPeriodStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "claimPeriodStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_ArbToken *ArbTokenSession) ClaimPeriodStart() (*big.Int, error) {
	return _ArbToken.Contract.ClaimPeriodStart(&_ArbToken.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_ArbToken *ArbTokenCallerSession) ClaimPeriodStart() (*big.Int, error) {
	return _ArbToken.Contract.ClaimPeriodStart(&_ArbToken.CallOpts)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_ArbToken *ArbTokenCaller) ClaimableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "claimableTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_ArbToken *ArbTokenSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _ArbToken.Contract.ClaimableTokens(&_ArbToken.CallOpts, arg0)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_ArbToken *ArbTokenCallerSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _ArbToken.Contract.ClaimableTokens(&_ArbToken.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbToken *ArbTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbToken *ArbTokenSession) Owner() (common.Address, error) {
	return _ArbToken.Contract.Owner(&_ArbToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ArbToken *ArbTokenCallerSession) Owner() (common.Address, error) {
	return _ArbToken.Contract.Owner(&_ArbToken.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_ArbToken *ArbTokenCaller) SweepReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "sweepReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_ArbToken *ArbTokenSession) SweepReceiver() (common.Address, error) {
	return _ArbToken.Contract.SweepReceiver(&_ArbToken.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_ArbToken *ArbTokenCallerSession) SweepReceiver() (common.Address, error) {
	return _ArbToken.Contract.SweepReceiver(&_ArbToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ArbToken *ArbTokenCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ArbToken *ArbTokenSession) Token() (common.Address, error) {
	return _ArbToken.Contract.Token(&_ArbToken.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ArbToken *ArbTokenCallerSession) Token() (common.Address, error) {
	return _ArbToken.Contract.Token(&_ArbToken.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_ArbToken *ArbTokenCaller) TotalClaimable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ArbToken.contract.Call(opts, &out, "totalClaimable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_ArbToken *ArbTokenSession) TotalClaimable() (*big.Int, error) {
	return _ArbToken.Contract.TotalClaimable(&_ArbToken.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_ArbToken *ArbTokenCallerSession) TotalClaimable() (*big.Int, error) {
	return _ArbToken.Contract.TotalClaimable(&_ArbToken.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ArbToken *ArbTokenTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ArbToken *ArbTokenSession) Claim() (*types.Transaction, error) {
	return _ArbToken.Contract.Claim(&_ArbToken.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_ArbToken *ArbTokenTransactorSession) Claim() (*types.Transaction, error) {
	return _ArbToken.Contract.Claim(&_ArbToken.TransactOpts)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ArbToken *ArbTokenTransactor) ClaimAndDelegate(opts *bind.TransactOpts, delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "claimAndDelegate", delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ArbToken *ArbTokenSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ArbToken.Contract.ClaimAndDelegate(&_ArbToken.TransactOpts, delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_ArbToken *ArbTokenTransactorSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _ArbToken.Contract.ClaimAndDelegate(&_ArbToken.TransactOpts, delegatee, expiry, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbToken *ArbTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbToken *ArbTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbToken.Contract.RenounceOwnership(&_ArbToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ArbToken *ArbTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ArbToken.Contract.RenounceOwnership(&_ArbToken.TransactOpts)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_ArbToken *ArbTokenTransactor) SetRecipients(opts *bind.TransactOpts, _recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "setRecipients", _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_ArbToken *ArbTokenSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _ArbToken.Contract.SetRecipients(&_ArbToken.TransactOpts, _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_ArbToken *ArbTokenTransactorSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _ArbToken.Contract.SetRecipients(&_ArbToken.TransactOpts, _recipients, _claimableAmount)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_ArbToken *ArbTokenTransactor) SetSweepReciever(opts *bind.TransactOpts, _sweepReceiver common.Address) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "setSweepReciever", _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_ArbToken *ArbTokenSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _ArbToken.Contract.SetSweepReciever(&_ArbToken.TransactOpts, _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_ArbToken *ArbTokenTransactorSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _ArbToken.Contract.SetSweepReciever(&_ArbToken.TransactOpts, _sweepReceiver)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_ArbToken *ArbTokenTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_ArbToken *ArbTokenSession) Sweep() (*types.Transaction, error) {
	return _ArbToken.Contract.Sweep(&_ArbToken.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_ArbToken *ArbTokenTransactorSession) Sweep() (*types.Transaction, error) {
	return _ArbToken.Contract.Sweep(&_ArbToken.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbToken *ArbTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbToken *ArbTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbToken.Contract.TransferOwnership(&_ArbToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ArbToken *ArbTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ArbToken.Contract.TransferOwnership(&_ArbToken.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ArbToken *ArbTokenTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ArbToken.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ArbToken *ArbTokenSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ArbToken.Contract.Withdraw(&_ArbToken.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_ArbToken *ArbTokenTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _ArbToken.Contract.Withdraw(&_ArbToken.TransactOpts, amount)
}

// ArbTokenCanClaimIterator is returned from FilterCanClaim and is used to iterate over the raw logs and unpacked data for CanClaim events raised by the ArbToken contract.
type ArbTokenCanClaimIterator struct {
	Event *ArbTokenCanClaim // Event containing the contract specifics and raw log

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
func (it *ArbTokenCanClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenCanClaim)
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
		it.Event = new(ArbTokenCanClaim)
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
func (it *ArbTokenCanClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenCanClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenCanClaim represents a CanClaim event raised by the ArbToken contract.
type ArbTokenCanClaim struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanClaim is a free log retrieval operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) FilterCanClaim(opts *bind.FilterOpts, recipient []common.Address) (*ArbTokenCanClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbTokenCanClaimIterator{contract: _ArbToken.contract, event: "CanClaim", logs: logs, sub: sub}, nil
}

// WatchCanClaim is a free log subscription operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) WatchCanClaim(opts *bind.WatchOpts, sink chan<- *ArbTokenCanClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenCanClaim)
				if err := _ArbToken.contract.UnpackLog(event, "CanClaim", log); err != nil {
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

// ParseCanClaim is a log parse operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) ParseCanClaim(log types.Log) (*ArbTokenCanClaim, error) {
	event := new(ArbTokenCanClaim)
	if err := _ArbToken.contract.UnpackLog(event, "CanClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTokenHasClaimedIterator is returned from FilterHasClaimed and is used to iterate over the raw logs and unpacked data for HasClaimed events raised by the ArbToken contract.
type ArbTokenHasClaimedIterator struct {
	Event *ArbTokenHasClaimed // Event containing the contract specifics and raw log

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
func (it *ArbTokenHasClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenHasClaimed)
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
		it.Event = new(ArbTokenHasClaimed)
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
func (it *ArbTokenHasClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenHasClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenHasClaimed represents a HasClaimed event raised by the ArbToken contract.
type ArbTokenHasClaimed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHasClaimed is a free log retrieval operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) FilterHasClaimed(opts *bind.FilterOpts, recipient []common.Address) (*ArbTokenHasClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbTokenHasClaimedIterator{contract: _ArbToken.contract, event: "HasClaimed", logs: logs, sub: sub}, nil
}

// WatchHasClaimed is a free log subscription operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) WatchHasClaimed(opts *bind.WatchOpts, sink chan<- *ArbTokenHasClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenHasClaimed)
				if err := _ArbToken.contract.UnpackLog(event, "HasClaimed", log); err != nil {
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

// ParseHasClaimed is a log parse operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) ParseHasClaimed(log types.Log) (*ArbTokenHasClaimed, error) {
	event := new(ArbTokenHasClaimed)
	if err := _ArbToken.contract.UnpackLog(event, "HasClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ArbToken contract.
type ArbTokenOwnershipTransferredIterator struct {
	Event *ArbTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenOwnershipTransferred)
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
		it.Event = new(ArbTokenOwnershipTransferred)
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
func (it *ArbTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenOwnershipTransferred represents a OwnershipTransferred event raised by the ArbToken contract.
type ArbTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbToken *ArbTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbTokenOwnershipTransferredIterator{contract: _ArbToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ArbToken *ArbTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenOwnershipTransferred)
				if err := _ArbToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ArbToken *ArbTokenFilterer) ParseOwnershipTransferred(log types.Log) (*ArbTokenOwnershipTransferred, error) {
	event := new(ArbTokenOwnershipTransferred)
	if err := _ArbToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTokenSweepReceiverSetIterator is returned from FilterSweepReceiverSet and is used to iterate over the raw logs and unpacked data for SweepReceiverSet events raised by the ArbToken contract.
type ArbTokenSweepReceiverSetIterator struct {
	Event *ArbTokenSweepReceiverSet // Event containing the contract specifics and raw log

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
func (it *ArbTokenSweepReceiverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenSweepReceiverSet)
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
		it.Event = new(ArbTokenSweepReceiverSet)
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
func (it *ArbTokenSweepReceiverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenSweepReceiverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenSweepReceiverSet represents a SweepReceiverSet event raised by the ArbToken contract.
type ArbTokenSweepReceiverSet struct {
	NewSweepReceiver common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSweepReceiverSet is a free log retrieval operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_ArbToken *ArbTokenFilterer) FilterSweepReceiverSet(opts *bind.FilterOpts, newSweepReceiver []common.Address) (*ArbTokenSweepReceiverSetIterator, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return &ArbTokenSweepReceiverSetIterator{contract: _ArbToken.contract, event: "SweepReceiverSet", logs: logs, sub: sub}, nil
}

// WatchSweepReceiverSet is a free log subscription operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_ArbToken *ArbTokenFilterer) WatchSweepReceiverSet(opts *bind.WatchOpts, sink chan<- *ArbTokenSweepReceiverSet, newSweepReceiver []common.Address) (event.Subscription, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenSweepReceiverSet)
				if err := _ArbToken.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
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

// ParseSweepReceiverSet is a log parse operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_ArbToken *ArbTokenFilterer) ParseSweepReceiverSet(log types.Log) (*ArbTokenSweepReceiverSet, error) {
	event := new(ArbTokenSweepReceiverSet)
	if err := _ArbToken.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTokenSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the ArbToken contract.
type ArbTokenSweptIterator struct {
	Event *ArbTokenSwept // Event containing the contract specifics and raw log

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
func (it *ArbTokenSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenSwept)
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
		it.Event = new(ArbTokenSwept)
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
func (it *ArbTokenSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenSwept represents a Swept event raised by the ArbToken contract.
type ArbTokenSwept struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_ArbToken *ArbTokenFilterer) FilterSwept(opts *bind.FilterOpts) (*ArbTokenSweptIterator, error) {

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return &ArbTokenSweptIterator{contract: _ArbToken.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_ArbToken *ArbTokenFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *ArbTokenSwept) (event.Subscription, error) {

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenSwept)
				if err := _ArbToken.contract.UnpackLog(event, "Swept", log); err != nil {
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

// ParseSwept is a log parse operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_ArbToken *ArbTokenFilterer) ParseSwept(log types.Log) (*ArbTokenSwept, error) {
	event := new(ArbTokenSwept)
	if err := _ArbToken.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbTokenWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the ArbToken contract.
type ArbTokenWithdrawalIterator struct {
	Event *ArbTokenWithdrawal // Event containing the contract specifics and raw log

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
func (it *ArbTokenWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbTokenWithdrawal)
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
		it.Event = new(ArbTokenWithdrawal)
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
func (it *ArbTokenWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbTokenWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbTokenWithdrawal represents a Withdrawal event raised by the ArbToken contract.
type ArbTokenWithdrawal struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) FilterWithdrawal(opts *bind.FilterOpts, recipient []common.Address) (*ArbTokenWithdrawalIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.FilterLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbTokenWithdrawalIterator{contract: _ArbToken.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ArbTokenWithdrawal, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ArbToken.contract.WatchLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbTokenWithdrawal)
				if err := _ArbToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_ArbToken *ArbTokenFilterer) ParseWithdrawal(log types.Log) (*ArbTokenWithdrawal, error) {
	event := new(ArbTokenWithdrawal)
	if err := _ArbToken.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
