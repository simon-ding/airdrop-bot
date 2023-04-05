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

// ArbMetaData contains all meta data concerning the Arb contract.
var ArbMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_claimPeriodEnd\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"delegateTo\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"CanClaim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"HasClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newSweepReceiver\",\"type\":\"address\"}],\"name\":\"SweepReceiverSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Swept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"claimAndDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodEnd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimPeriodStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"claimableTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_claimableAmount\",\"type\":\"uint256[]\"}],\"name\":\"setRecipients\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_sweepReceiver\",\"type\":\"address\"}],\"name\":\"setSweepReciever\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweep\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sweepReceiver\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20VotesUpgradeable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ArbABI is the input ABI used to generate the binding from.
// Deprecated: Use ArbMetaData.ABI instead.
var ArbABI = ArbMetaData.ABI

// Arb is an auto generated Go binding around an Ethereum contract.
type Arb struct {
	ArbCaller     // Read-only binding to the contract
	ArbTransactor // Write-only binding to the contract
	ArbFilterer   // Log filterer for contract events
}

// ArbCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArbCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArbTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArbFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArbSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArbSession struct {
	Contract     *Arb              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArbCallerSession struct {
	Contract *ArbCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ArbTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArbTransactorSession struct {
	Contract     *ArbTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArbRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArbRaw struct {
	Contract *Arb // Generic contract binding to access the raw methods on
}

// ArbCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArbCallerRaw struct {
	Contract *ArbCaller // Generic read-only contract binding to access the raw methods on
}

// ArbTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArbTransactorRaw struct {
	Contract *ArbTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArb creates a new instance of Arb, bound to a specific deployed contract.
func NewArb(address common.Address, backend bind.ContractBackend) (*Arb, error) {
	contract, err := bindArb(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Arb{ArbCaller: ArbCaller{contract: contract}, ArbTransactor: ArbTransactor{contract: contract}, ArbFilterer: ArbFilterer{contract: contract}}, nil
}

// NewArbCaller creates a new read-only instance of Arb, bound to a specific deployed contract.
func NewArbCaller(address common.Address, caller bind.ContractCaller) (*ArbCaller, error) {
	contract, err := bindArb(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArbCaller{contract: contract}, nil
}

// NewArbTransactor creates a new write-only instance of Arb, bound to a specific deployed contract.
func NewArbTransactor(address common.Address, transactor bind.ContractTransactor) (*ArbTransactor, error) {
	contract, err := bindArb(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArbTransactor{contract: contract}, nil
}

// NewArbFilterer creates a new log filterer instance of Arb, bound to a specific deployed contract.
func NewArbFilterer(address common.Address, filterer bind.ContractFilterer) (*ArbFilterer, error) {
	contract, err := bindArb(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArbFilterer{contract: contract}, nil
}

// bindArb binds a generic wrapper to an already deployed contract.
func bindArb(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArbMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.ArbCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.ArbTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Arb *ArbCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Arb.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Arb *ArbTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Arb *ArbTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Arb.Contract.contract.Transact(opts, method, params...)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Arb *ArbCaller) ClaimPeriodEnd(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "claimPeriodEnd")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Arb *ArbSession) ClaimPeriodEnd() (*big.Int, error) {
	return _Arb.Contract.ClaimPeriodEnd(&_Arb.CallOpts)
}

// ClaimPeriodEnd is a free data retrieval call binding the contract method 0x3da082a0.
//
// Solidity: function claimPeriodEnd() view returns(uint256)
func (_Arb *ArbCallerSession) ClaimPeriodEnd() (*big.Int, error) {
	return _Arb.Contract.ClaimPeriodEnd(&_Arb.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Arb *ArbCaller) ClaimPeriodStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "claimPeriodStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Arb *ArbSession) ClaimPeriodStart() (*big.Int, error) {
	return _Arb.Contract.ClaimPeriodStart(&_Arb.CallOpts)
}

// ClaimPeriodStart is a free data retrieval call binding the contract method 0x58c13b7e.
//
// Solidity: function claimPeriodStart() view returns(uint256)
func (_Arb *ArbCallerSession) ClaimPeriodStart() (*big.Int, error) {
	return _Arb.Contract.ClaimPeriodStart(&_Arb.CallOpts)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Arb *ArbCaller) ClaimableTokens(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "claimableTokens", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Arb *ArbSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _Arb.Contract.ClaimableTokens(&_Arb.CallOpts, arg0)
}

// ClaimableTokens is a free data retrieval call binding the contract method 0x84d24226.
//
// Solidity: function claimableTokens(address ) view returns(uint256)
func (_Arb *ArbCallerSession) ClaimableTokens(arg0 common.Address) (*big.Int, error) {
	return _Arb.Contract.ClaimableTokens(&_Arb.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Arb *ArbCallerSession) Owner() (common.Address, error) {
	return _Arb.Contract.Owner(&_Arb.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Arb *ArbCaller) SweepReceiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "sweepReceiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Arb *ArbSession) SweepReceiver() (common.Address, error) {
	return _Arb.Contract.SweepReceiver(&_Arb.CallOpts)
}

// SweepReceiver is a free data retrieval call binding the contract method 0xf6e0df9f.
//
// Solidity: function sweepReceiver() view returns(address)
func (_Arb *ArbCallerSession) SweepReceiver() (common.Address, error) {
	return _Arb.Contract.SweepReceiver(&_Arb.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Arb *ArbCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Arb *ArbSession) Token() (common.Address, error) {
	return _Arb.Contract.Token(&_Arb.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Arb *ArbCallerSession) Token() (common.Address, error) {
	return _Arb.Contract.Token(&_Arb.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Arb *ArbCaller) TotalClaimable(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Arb.contract.Call(opts, &out, "totalClaimable")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Arb *ArbSession) TotalClaimable() (*big.Int, error) {
	return _Arb.Contract.TotalClaimable(&_Arb.CallOpts)
}

// TotalClaimable is a free data retrieval call binding the contract method 0x4838ed19.
//
// Solidity: function totalClaimable() view returns(uint256)
func (_Arb *ArbCallerSession) TotalClaimable() (*big.Int, error) {
	return _Arb.Contract.TotalClaimable(&_Arb.CallOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Arb *ArbTransactor) Claim(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "claim")
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Arb *ArbSession) Claim() (*types.Transaction, error) {
	return _Arb.Contract.Claim(&_Arb.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x4e71d92d.
//
// Solidity: function claim() returns()
func (_Arb *ArbTransactorSession) Claim() (*types.Transaction, error) {
	return _Arb.Contract.Claim(&_Arb.TransactOpts)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Arb *ArbTransactor) ClaimAndDelegate(opts *bind.TransactOpts, delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "claimAndDelegate", delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Arb *ArbSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Arb.Contract.ClaimAndDelegate(&_Arb.TransactOpts, delegatee, expiry, v, r, s)
}

// ClaimAndDelegate is a paid mutator transaction binding the contract method 0x78e2b594.
//
// Solidity: function claimAndDelegate(address delegatee, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Arb *ArbTransactorSession) ClaimAndDelegate(delegatee common.Address, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Arb.Contract.ClaimAndDelegate(&_Arb.TransactOpts, delegatee, expiry, v, r, s)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Arb *ArbTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Arb.Contract.RenounceOwnership(&_Arb.TransactOpts)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Arb *ArbTransactor) SetRecipients(opts *bind.TransactOpts, _recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "setRecipients", _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Arb *ArbSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Arb.Contract.SetRecipients(&_Arb.TransactOpts, _recipients, _claimableAmount)
}

// SetRecipients is a paid mutator transaction binding the contract method 0xae373c1b.
//
// Solidity: function setRecipients(address[] _recipients, uint256[] _claimableAmount) returns()
func (_Arb *ArbTransactorSession) SetRecipients(_recipients []common.Address, _claimableAmount []*big.Int) (*types.Transaction, error) {
	return _Arb.Contract.SetRecipients(&_Arb.TransactOpts, _recipients, _claimableAmount)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Arb *ArbTransactor) SetSweepReciever(opts *bind.TransactOpts, _sweepReceiver common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "setSweepReciever", _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Arb *ArbSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _Arb.Contract.SetSweepReciever(&_Arb.TransactOpts, _sweepReceiver)
}

// SetSweepReciever is a paid mutator transaction binding the contract method 0xb438abde.
//
// Solidity: function setSweepReciever(address _sweepReceiver) returns()
func (_Arb *ArbTransactorSession) SetSweepReciever(_sweepReceiver common.Address) (*types.Transaction, error) {
	return _Arb.Contract.SetSweepReciever(&_Arb.TransactOpts, _sweepReceiver)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Arb *ArbTransactor) Sweep(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "sweep")
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Arb *ArbSession) Sweep() (*types.Transaction, error) {
	return _Arb.Contract.Sweep(&_Arb.TransactOpts)
}

// Sweep is a paid mutator transaction binding the contract method 0x35faa416.
//
// Solidity: function sweep() returns()
func (_Arb *ArbTransactorSession) Sweep() (*types.Transaction, error) {
	return _Arb.Contract.Sweep(&_Arb.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Arb *ArbTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Arb.Contract.TransferOwnership(&_Arb.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Arb *ArbTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Arb.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Arb *ArbSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Withdraw(&_Arb.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_Arb *ArbTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _Arb.Contract.Withdraw(&_Arb.TransactOpts, amount)
}

// ArbCanClaimIterator is returned from FilterCanClaim and is used to iterate over the raw logs and unpacked data for CanClaim events raised by the Arb contract.
type ArbCanClaimIterator struct {
	Event *ArbCanClaim // Event containing the contract specifics and raw log

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
func (it *ArbCanClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbCanClaim)
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
		it.Event = new(ArbCanClaim)
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
func (it *ArbCanClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbCanClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbCanClaim represents a CanClaim event raised by the Arb contract.
type ArbCanClaim struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanClaim is a free log retrieval operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) FilterCanClaim(opts *bind.FilterOpts, recipient []common.Address) (*ArbCanClaimIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbCanClaimIterator{contract: _Arb.contract, event: "CanClaim", logs: logs, sub: sub}, nil
}

// WatchCanClaim is a free log subscription operation binding the contract event 0x87aeeb9eda09a064caef63d00f62c15063631980bfc422ad7dd30c8a79f0cbb7.
//
// Solidity: event CanClaim(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) WatchCanClaim(opts *bind.WatchOpts, sink chan<- *ArbCanClaim, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "CanClaim", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbCanClaim)
				if err := _Arb.contract.UnpackLog(event, "CanClaim", log); err != nil {
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
func (_Arb *ArbFilterer) ParseCanClaim(log types.Log) (*ArbCanClaim, error) {
	event := new(ArbCanClaim)
	if err := _Arb.contract.UnpackLog(event, "CanClaim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbHasClaimedIterator is returned from FilterHasClaimed and is used to iterate over the raw logs and unpacked data for HasClaimed events raised by the Arb contract.
type ArbHasClaimedIterator struct {
	Event *ArbHasClaimed // Event containing the contract specifics and raw log

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
func (it *ArbHasClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbHasClaimed)
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
		it.Event = new(ArbHasClaimed)
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
func (it *ArbHasClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbHasClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbHasClaimed represents a HasClaimed event raised by the Arb contract.
type ArbHasClaimed struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterHasClaimed is a free log retrieval operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) FilterHasClaimed(opts *bind.FilterOpts, recipient []common.Address) (*ArbHasClaimedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbHasClaimedIterator{contract: _Arb.contract, event: "HasClaimed", logs: logs, sub: sub}, nil
}

// WatchHasClaimed is a free log subscription operation binding the contract event 0x8629b200ebe43db58ad688b85131d53251f3f3be4c14933b4641aeebacf1c08c.
//
// Solidity: event HasClaimed(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) WatchHasClaimed(opts *bind.WatchOpts, sink chan<- *ArbHasClaimed, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "HasClaimed", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbHasClaimed)
				if err := _Arb.contract.UnpackLog(event, "HasClaimed", log); err != nil {
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
func (_Arb *ArbFilterer) ParseHasClaimed(log types.Log) (*ArbHasClaimed, error) {
	event := new(ArbHasClaimed)
	if err := _Arb.contract.UnpackLog(event, "HasClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Arb contract.
type ArbOwnershipTransferredIterator struct {
	Event *ArbOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ArbOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbOwnershipTransferred)
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
		it.Event = new(ArbOwnershipTransferred)
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
func (it *ArbOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbOwnershipTransferred represents a OwnershipTransferred event raised by the Arb contract.
type ArbOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ArbOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ArbOwnershipTransferredIterator{contract: _Arb.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Arb *ArbFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ArbOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbOwnershipTransferred)
				if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Arb *ArbFilterer) ParseOwnershipTransferred(log types.Log) (*ArbOwnershipTransferred, error) {
	event := new(ArbOwnershipTransferred)
	if err := _Arb.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbSweepReceiverSetIterator is returned from FilterSweepReceiverSet and is used to iterate over the raw logs and unpacked data for SweepReceiverSet events raised by the Arb contract.
type ArbSweepReceiverSetIterator struct {
	Event *ArbSweepReceiverSet // Event containing the contract specifics and raw log

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
func (it *ArbSweepReceiverSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSweepReceiverSet)
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
		it.Event = new(ArbSweepReceiverSet)
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
func (it *ArbSweepReceiverSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbSweepReceiverSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbSweepReceiverSet represents a SweepReceiverSet event raised by the Arb contract.
type ArbSweepReceiverSet struct {
	NewSweepReceiver common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSweepReceiverSet is a free log retrieval operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_Arb *ArbFilterer) FilterSweepReceiverSet(opts *bind.FilterOpts, newSweepReceiver []common.Address) (*ArbSweepReceiverSetIterator, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return &ArbSweepReceiverSetIterator{contract: _Arb.contract, event: "SweepReceiverSet", logs: logs, sub: sub}, nil
}

// WatchSweepReceiverSet is a free log subscription operation binding the contract event 0xbea8251f76064f657f2a744bf08a1b5700486d06eb94922162892eb022d95ef6.
//
// Solidity: event SweepReceiverSet(address indexed newSweepReceiver)
func (_Arb *ArbFilterer) WatchSweepReceiverSet(opts *bind.WatchOpts, sink chan<- *ArbSweepReceiverSet, newSweepReceiver []common.Address) (event.Subscription, error) {

	var newSweepReceiverRule []interface{}
	for _, newSweepReceiverItem := range newSweepReceiver {
		newSweepReceiverRule = append(newSweepReceiverRule, newSweepReceiverItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "SweepReceiverSet", newSweepReceiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbSweepReceiverSet)
				if err := _Arb.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
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
func (_Arb *ArbFilterer) ParseSweepReceiverSet(log types.Log) (*ArbSweepReceiverSet, error) {
	event := new(ArbSweepReceiverSet)
	if err := _Arb.contract.UnpackLog(event, "SweepReceiverSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbSweptIterator is returned from FilterSwept and is used to iterate over the raw logs and unpacked data for Swept events raised by the Arb contract.
type ArbSweptIterator struct {
	Event *ArbSwept // Event containing the contract specifics and raw log

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
func (it *ArbSweptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbSwept)
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
		it.Event = new(ArbSwept)
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
func (it *ArbSweptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbSweptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbSwept represents a Swept event raised by the Arb contract.
type ArbSwept struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSwept is a free log retrieval operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_Arb *ArbFilterer) FilterSwept(opts *bind.FilterOpts) (*ArbSweptIterator, error) {

	logs, sub, err := _Arb.contract.FilterLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return &ArbSweptIterator{contract: _Arb.contract, event: "Swept", logs: logs, sub: sub}, nil
}

// WatchSwept is a free log subscription operation binding the contract event 0x7f221332ee403570bf4d61630b58189ea566ff1635269001e9df6a890f413dd8.
//
// Solidity: event Swept(uint256 amount)
func (_Arb *ArbFilterer) WatchSwept(opts *bind.WatchOpts, sink chan<- *ArbSwept) (event.Subscription, error) {

	logs, sub, err := _Arb.contract.WatchLogs(opts, "Swept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbSwept)
				if err := _Arb.contract.UnpackLog(event, "Swept", log); err != nil {
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
func (_Arb *ArbFilterer) ParseSwept(log types.Log) (*ArbSwept, error) {
	event := new(ArbSwept)
	if err := _Arb.contract.UnpackLog(event, "Swept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ArbWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the Arb contract.
type ArbWithdrawalIterator struct {
	Event *ArbWithdrawal // Event containing the contract specifics and raw log

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
func (it *ArbWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ArbWithdrawal)
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
		it.Event = new(ArbWithdrawal)
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
func (it *ArbWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ArbWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ArbWithdrawal represents a Withdrawal event raised by the Arb contract.
type ArbWithdrawal struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) FilterWithdrawal(opts *bind.FilterOpts, recipient []common.Address) (*ArbWithdrawalIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.FilterLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ArbWithdrawalIterator{contract: _Arb.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed recipient, uint256 amount)
func (_Arb *ArbFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *ArbWithdrawal, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Arb.contract.WatchLogs(opts, "Withdrawal", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ArbWithdrawal)
				if err := _Arb.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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
func (_Arb *ArbFilterer) ParseWithdrawal(log types.Log) (*ArbWithdrawal, error) {
	event := new(ArbWithdrawal)
	if err := _Arb.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
