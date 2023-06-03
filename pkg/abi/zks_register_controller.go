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

// ZNSRegistrarControllerEpoch is an auto generated low-level Go binding around an user-defined struct.
type ZNSRegistrarControllerEpoch struct {
	OpenTime  *big.Int
	MinLength *big.Int
	MaxLength *big.Int
}

// ZksBaseRegisterMetaData contains all meta data concerning the ZksBaseRegister contract.
var ZksBaseRegisterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"_base\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_teamAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"label\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cost\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expires\",\"type\":\"uint256\"}],\"name\":\"NameRenewed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIN_REGISTRATION_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ONE_YEAR_DURATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WL_PRIORITY_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"YearlyBasePrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"YearlyPriceMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"openTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLength\",\"type\":\"uint256\"}],\"name\":\"addEpoch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"freeCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"allowedCount\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"}],\"name\":\"addToWL\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"canRegister\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"checkWL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentEpoch\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"openTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxLength\",\"type\":\"uint256\"}],\"internalType\":\"structZNSRegistrarController.Epoch\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"openTimeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"yearCount\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"yearCount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"registerWithConfig\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"yearCount\",\"type\":\"uint256\"}],\"name\":\"renew\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"yearCount\",\"type\":\"uint256\"}],\"name\":\"rentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"yearCount\",\"type\":\"uint256\"}],\"name\":\"rentPriceForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractBaseRegistrar\",\"name\":\"_base\",\"type\":\"address\"}],\"name\":\"setBase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nameLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_teamAddress\",\"type\":\"address\"}],\"name\":\"setTeamAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"setWLPriority\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setYearlyBasePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"valid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenContract\",\"type\":\"address\"}],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ZksBaseRegisterABI is the input ABI used to generate the binding from.
// Deprecated: Use ZksBaseRegisterMetaData.ABI instead.
var ZksBaseRegisterABI = ZksBaseRegisterMetaData.ABI

// ZksBaseRegister is an auto generated Go binding around an Ethereum contract.
type ZksBaseRegister struct {
	ZksBaseRegisterCaller     // Read-only binding to the contract
	ZksBaseRegisterTransactor // Write-only binding to the contract
	ZksBaseRegisterFilterer   // Log filterer for contract events
}

// ZksBaseRegisterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ZksBaseRegisterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ZksBaseRegisterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ZksBaseRegisterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ZksBaseRegisterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ZksBaseRegisterSession struct {
	Contract     *ZksBaseRegister  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ZksBaseRegisterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ZksBaseRegisterCallerSession struct {
	Contract *ZksBaseRegisterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ZksBaseRegisterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ZksBaseRegisterTransactorSession struct {
	Contract     *ZksBaseRegisterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ZksBaseRegisterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ZksBaseRegisterRaw struct {
	Contract *ZksBaseRegister // Generic contract binding to access the raw methods on
}

// ZksBaseRegisterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ZksBaseRegisterCallerRaw struct {
	Contract *ZksBaseRegisterCaller // Generic read-only contract binding to access the raw methods on
}

// ZksBaseRegisterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ZksBaseRegisterTransactorRaw struct {
	Contract *ZksBaseRegisterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewZksBaseRegister creates a new instance of ZksBaseRegister, bound to a specific deployed contract.
func NewZksBaseRegister(address common.Address, backend bind.ContractBackend) (*ZksBaseRegister, error) {
	contract, err := bindZksBaseRegister(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegister{ZksBaseRegisterCaller: ZksBaseRegisterCaller{contract: contract}, ZksBaseRegisterTransactor: ZksBaseRegisterTransactor{contract: contract}, ZksBaseRegisterFilterer: ZksBaseRegisterFilterer{contract: contract}}, nil
}

// NewZksBaseRegisterCaller creates a new read-only instance of ZksBaseRegister, bound to a specific deployed contract.
func NewZksBaseRegisterCaller(address common.Address, caller bind.ContractCaller) (*ZksBaseRegisterCaller, error) {
	contract, err := bindZksBaseRegister(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterCaller{contract: contract}, nil
}

// NewZksBaseRegisterTransactor creates a new write-only instance of ZksBaseRegister, bound to a specific deployed contract.
func NewZksBaseRegisterTransactor(address common.Address, transactor bind.ContractTransactor) (*ZksBaseRegisterTransactor, error) {
	contract, err := bindZksBaseRegister(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterTransactor{contract: contract}, nil
}

// NewZksBaseRegisterFilterer creates a new log filterer instance of ZksBaseRegister, bound to a specific deployed contract.
func NewZksBaseRegisterFilterer(address common.Address, filterer bind.ContractFilterer) (*ZksBaseRegisterFilterer, error) {
	contract, err := bindZksBaseRegister(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterFilterer{contract: contract}, nil
}

// bindZksBaseRegister binds a generic wrapper to an already deployed contract.
func bindZksBaseRegister(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ZksBaseRegisterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZksBaseRegister *ZksBaseRegisterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZksBaseRegister.Contract.ZksBaseRegisterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZksBaseRegister *ZksBaseRegisterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.ZksBaseRegisterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZksBaseRegister *ZksBaseRegisterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.ZksBaseRegisterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ZksBaseRegister *ZksBaseRegisterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ZksBaseRegister.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ZksBaseRegister *ZksBaseRegisterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ZksBaseRegister *ZksBaseRegisterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.contract.Transact(opts, method, params...)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) MINREGISTRATIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "MIN_REGISTRATION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZksBaseRegister.Contract.MINREGISTRATIONDURATION(&_ZksBaseRegister.CallOpts)
}

// MINREGISTRATIONDURATION is a free data retrieval call binding the contract method 0x8a95b09f.
//
// Solidity: function MIN_REGISTRATION_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) MINREGISTRATIONDURATION() (*big.Int, error) {
	return _ZksBaseRegister.Contract.MINREGISTRATIONDURATION(&_ZksBaseRegister.CallOpts)
}

// ONEYEARDURATION is a free data retrieval call binding the contract method 0x631e136c.
//
// Solidity: function ONE_YEAR_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) ONEYEARDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "ONE_YEAR_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ONEYEARDURATION is a free data retrieval call binding the contract method 0x631e136c.
//
// Solidity: function ONE_YEAR_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) ONEYEARDURATION() (*big.Int, error) {
	return _ZksBaseRegister.Contract.ONEYEARDURATION(&_ZksBaseRegister.CallOpts)
}

// ONEYEARDURATION is a free data retrieval call binding the contract method 0x631e136c.
//
// Solidity: function ONE_YEAR_DURATION() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) ONEYEARDURATION() (*big.Int, error) {
	return _ZksBaseRegister.Contract.ONEYEARDURATION(&_ZksBaseRegister.CallOpts)
}

// WLPRIORITYPERIOD is a free data retrieval call binding the contract method 0x4c5e182b.
//
// Solidity: function WL_PRIORITY_PERIOD() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) WLPRIORITYPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "WL_PRIORITY_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WLPRIORITYPERIOD is a free data retrieval call binding the contract method 0x4c5e182b.
//
// Solidity: function WL_PRIORITY_PERIOD() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) WLPRIORITYPERIOD() (*big.Int, error) {
	return _ZksBaseRegister.Contract.WLPRIORITYPERIOD(&_ZksBaseRegister.CallOpts)
}

// WLPRIORITYPERIOD is a free data retrieval call binding the contract method 0x4c5e182b.
//
// Solidity: function WL_PRIORITY_PERIOD() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) WLPRIORITYPERIOD() (*big.Int, error) {
	return _ZksBaseRegister.Contract.WLPRIORITYPERIOD(&_ZksBaseRegister.CallOpts)
}

// YearlyBasePrice is a free data retrieval call binding the contract method 0x4e36949d.
//
// Solidity: function YearlyBasePrice() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) YearlyBasePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "YearlyBasePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YearlyBasePrice is a free data retrieval call binding the contract method 0x4e36949d.
//
// Solidity: function YearlyBasePrice() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) YearlyBasePrice() (*big.Int, error) {
	return _ZksBaseRegister.Contract.YearlyBasePrice(&_ZksBaseRegister.CallOpts)
}

// YearlyBasePrice is a free data retrieval call binding the contract method 0x4e36949d.
//
// Solidity: function YearlyBasePrice() view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) YearlyBasePrice() (*big.Int, error) {
	return _ZksBaseRegister.Contract.YearlyBasePrice(&_ZksBaseRegister.CallOpts)
}

// YearlyPriceMap is a free data retrieval call binding the contract method 0x61ef425e.
//
// Solidity: function YearlyPriceMap(uint256 ) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) YearlyPriceMap(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "YearlyPriceMap", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// YearlyPriceMap is a free data retrieval call binding the contract method 0x61ef425e.
//
// Solidity: function YearlyPriceMap(uint256 ) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) YearlyPriceMap(arg0 *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.YearlyPriceMap(&_ZksBaseRegister.CallOpts, arg0)
}

// YearlyPriceMap is a free data retrieval call binding the contract method 0x61ef425e.
//
// Solidity: function YearlyPriceMap(uint256 ) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) YearlyPriceMap(arg0 *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.YearlyPriceMap(&_ZksBaseRegister.CallOpts, arg0)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCaller) Available(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "available", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterSession) Available(name string) (bool, error) {
	return _ZksBaseRegister.Contract.Available(&_ZksBaseRegister.CallOpts, name)
}

// Available is a free data retrieval call binding the contract method 0xaeb8ce9b.
//
// Solidity: function available(string name) view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) Available(name string) (bool, error) {
	return _ZksBaseRegister.Contract.Available(&_ZksBaseRegister.CallOpts, name)
}

// CanRegister is a free data retrieval call binding the contract method 0xfaa88afd.
//
// Solidity: function canRegister(string name, address user) view returns(bool, bool)
func (_ZksBaseRegister *ZksBaseRegisterCaller) CanRegister(opts *bind.CallOpts, name string, user common.Address) (bool, bool, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "canRegister", name, user)

	if err != nil {
		return *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// CanRegister is a free data retrieval call binding the contract method 0xfaa88afd.
//
// Solidity: function canRegister(string name, address user) view returns(bool, bool)
func (_ZksBaseRegister *ZksBaseRegisterSession) CanRegister(name string, user common.Address) (bool, bool, error) {
	return _ZksBaseRegister.Contract.CanRegister(&_ZksBaseRegister.CallOpts, name, user)
}

// CanRegister is a free data retrieval call binding the contract method 0xfaa88afd.
//
// Solidity: function canRegister(string name, address user) view returns(bool, bool)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) CanRegister(name string, user common.Address) (bool, bool, error) {
	return _ZksBaseRegister.Contract.CanRegister(&_ZksBaseRegister.CallOpts, name, user)
}

// CheckWL is a free data retrieval call binding the contract method 0x18c63f24.
//
// Solidity: function checkWL(address user) view returns(uint256, uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) CheckWL(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "checkWL", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CheckWL is a free data retrieval call binding the contract method 0x18c63f24.
//
// Solidity: function checkWL(address user) view returns(uint256, uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) CheckWL(user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ZksBaseRegister.Contract.CheckWL(&_ZksBaseRegister.CallOpts, user)
}

// CheckWL is a free data retrieval call binding the contract method 0x18c63f24.
//
// Solidity: function checkWL(address user) view returns(uint256, uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) CheckWL(user common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _ZksBaseRegister.Contract.CheckWL(&_ZksBaseRegister.CallOpts, user)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns((uint256,uint256,uint256))
func (_ZksBaseRegister *ZksBaseRegisterCaller) CurrentEpoch(opts *bind.CallOpts) (ZNSRegistrarControllerEpoch, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(ZNSRegistrarControllerEpoch), err
	}

	out0 := *abi.ConvertType(out[0], new(ZNSRegistrarControllerEpoch)).(*ZNSRegistrarControllerEpoch)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns((uint256,uint256,uint256))
func (_ZksBaseRegister *ZksBaseRegisterSession) CurrentEpoch() (ZNSRegistrarControllerEpoch, error) {
	return _ZksBaseRegister.Contract.CurrentEpoch(&_ZksBaseRegister.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns((uint256,uint256,uint256))
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) CurrentEpoch() (ZNSRegistrarControllerEpoch, error) {
	return _ZksBaseRegister.Contract.CurrentEpoch(&_ZksBaseRegister.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterSession) IsOwner() (bool, error) {
	return _ZksBaseRegister.Contract.IsOwner(&_ZksBaseRegister.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) IsOwner() (bool, error) {
	return _ZksBaseRegister.Contract.IsOwner(&_ZksBaseRegister.CallOpts)
}

// OpenTimeInfo is a free data retrieval call binding the contract method 0xbc512637.
//
// Solidity: function openTimeInfo() view returns(uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) OpenTimeInfo(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "openTimeInfo")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// OpenTimeInfo is a free data retrieval call binding the contract method 0xbc512637.
//
// Solidity: function openTimeInfo() view returns(uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) OpenTimeInfo() (*big.Int, *big.Int, error) {
	return _ZksBaseRegister.Contract.OpenTimeInfo(&_ZksBaseRegister.CallOpts)
}

// OpenTimeInfo is a free data retrieval call binding the contract method 0xbc512637.
//
// Solidity: function openTimeInfo() view returns(uint256, uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) OpenTimeInfo() (*big.Int, *big.Int, error) {
	return _ZksBaseRegister.Contract.OpenTimeInfo(&_ZksBaseRegister.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZksBaseRegister *ZksBaseRegisterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZksBaseRegister *ZksBaseRegisterSession) Owner() (common.Address, error) {
	return _ZksBaseRegister.Contract.Owner(&_ZksBaseRegister.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) Owner() (common.Address, error) {
	return _ZksBaseRegister.Contract.Owner(&_ZksBaseRegister.CallOpts)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) RentPrice(opts *bind.CallOpts, name string, yearCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "rentPrice", name, yearCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) RentPrice(name string, yearCount *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.RentPrice(&_ZksBaseRegister.CallOpts, name, yearCount)
}

// RentPrice is a free data retrieval call binding the contract method 0x83e7f6ff.
//
// Solidity: function rentPrice(string name, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) RentPrice(name string, yearCount *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.RentPrice(&_ZksBaseRegister.CallOpts, name, yearCount)
}

// RentPriceForUser is a free data retrieval call binding the contract method 0x17ea0123.
//
// Solidity: function rentPriceForUser(string name, address user, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCaller) RentPriceForUser(opts *bind.CallOpts, name string, user common.Address, yearCount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "rentPriceForUser", name, user, yearCount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RentPriceForUser is a free data retrieval call binding the contract method 0x17ea0123.
//
// Solidity: function rentPriceForUser(string name, address user, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterSession) RentPriceForUser(name string, user common.Address, yearCount *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.RentPriceForUser(&_ZksBaseRegister.CallOpts, name, user, yearCount)
}

// RentPriceForUser is a free data retrieval call binding the contract method 0x17ea0123.
//
// Solidity: function rentPriceForUser(string name, address user, uint256 yearCount) view returns(uint256)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) RentPriceForUser(name string, user common.Address, yearCount *big.Int) (*big.Int, error) {
	return _ZksBaseRegister.Contract.RentPriceForUser(&_ZksBaseRegister.CallOpts, name, user, yearCount)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZksBaseRegister.Contract.SupportsInterface(&_ZksBaseRegister.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _ZksBaseRegister.Contract.SupportsInterface(&_ZksBaseRegister.CallOpts, interfaceID)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCaller) Valid(opts *bind.CallOpts, name string) (bool, error) {
	var out []interface{}
	err := _ZksBaseRegister.contract.Call(opts, &out, "valid", name)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterSession) Valid(name string) (bool, error) {
	return _ZksBaseRegister.Contract.Valid(&_ZksBaseRegister.CallOpts, name)
}

// Valid is a free data retrieval call binding the contract method 0x9791c097.
//
// Solidity: function valid(string name) pure returns(bool)
func (_ZksBaseRegister *ZksBaseRegisterCallerSession) Valid(name string) (bool, error) {
	return _ZksBaseRegister.Contract.Valid(&_ZksBaseRegister.CallOpts, name)
}

// AddEpoch is a paid mutator transaction binding the contract method 0xa98fb345.
//
// Solidity: function addEpoch(uint256 openTime, uint256 minLength, uint256 maxLength) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) AddEpoch(opts *bind.TransactOpts, openTime *big.Int, minLength *big.Int, maxLength *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "addEpoch", openTime, minLength, maxLength)
}

// AddEpoch is a paid mutator transaction binding the contract method 0xa98fb345.
//
// Solidity: function addEpoch(uint256 openTime, uint256 minLength, uint256 maxLength) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) AddEpoch(openTime *big.Int, minLength *big.Int, maxLength *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.AddEpoch(&_ZksBaseRegister.TransactOpts, openTime, minLength, maxLength)
}

// AddEpoch is a paid mutator transaction binding the contract method 0xa98fb345.
//
// Solidity: function addEpoch(uint256 openTime, uint256 minLength, uint256 maxLength) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) AddEpoch(openTime *big.Int, minLength *big.Int, maxLength *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.AddEpoch(&_ZksBaseRegister.TransactOpts, openTime, minLength, maxLength)
}

// AddToWL is a paid mutator transaction binding the contract method 0x51fc2402.
//
// Solidity: function addToWL(uint256 minLength, uint256 freeCount, uint256 allowedCount, address[] users) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) AddToWL(opts *bind.TransactOpts, minLength *big.Int, freeCount *big.Int, allowedCount *big.Int, users []common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "addToWL", minLength, freeCount, allowedCount, users)
}

// AddToWL is a paid mutator transaction binding the contract method 0x51fc2402.
//
// Solidity: function addToWL(uint256 minLength, uint256 freeCount, uint256 allowedCount, address[] users) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) AddToWL(minLength *big.Int, freeCount *big.Int, allowedCount *big.Int, users []common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.AddToWL(&_ZksBaseRegister.TransactOpts, minLength, freeCount, allowedCount, users)
}

// AddToWL is a paid mutator transaction binding the contract method 0x51fc2402.
//
// Solidity: function addToWL(uint256 minLength, uint256 freeCount, uint256 allowedCount, address[] users) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) AddToWL(minLength *big.Int, freeCount *big.Int, allowedCount *big.Int, users []common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.AddToWL(&_ZksBaseRegister.TransactOpts, minLength, freeCount, allowedCount, users)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address owner, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) Register(opts *bind.TransactOpts, name string, owner common.Address, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "register", name, owner, yearCount)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address owner, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) Register(name string, owner common.Address, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Register(&_ZksBaseRegister.TransactOpts, name, owner, yearCount)
}

// Register is a paid mutator transaction binding the contract method 0xd393c871.
//
// Solidity: function register(string name, address owner, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) Register(name string, owner common.Address, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Register(&_ZksBaseRegister.TransactOpts, name, owner, yearCount)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xb2d3ed39.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 yearCount, address resolver, address addr) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) RegisterWithConfig(opts *bind.TransactOpts, name string, owner common.Address, yearCount *big.Int, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "registerWithConfig", name, owner, yearCount, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xb2d3ed39.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 yearCount, address resolver, address addr) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) RegisterWithConfig(name string, owner common.Address, yearCount *big.Int, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.RegisterWithConfig(&_ZksBaseRegister.TransactOpts, name, owner, yearCount, resolver, addr)
}

// RegisterWithConfig is a paid mutator transaction binding the contract method 0xb2d3ed39.
//
// Solidity: function registerWithConfig(string name, address owner, uint256 yearCount, address resolver, address addr) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) RegisterWithConfig(name string, owner common.Address, yearCount *big.Int, resolver common.Address, addr common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.RegisterWithConfig(&_ZksBaseRegister.TransactOpts, name, owner, yearCount, resolver, addr)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) Renew(opts *bind.TransactOpts, name string, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "renew", name, yearCount)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) Renew(name string, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Renew(&_ZksBaseRegister.TransactOpts, name, yearCount)
}

// Renew is a paid mutator transaction binding the contract method 0xacf1a841.
//
// Solidity: function renew(string name, uint256 yearCount) payable returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) Renew(name string, yearCount *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Renew(&_ZksBaseRegister.TransactOpts, name, yearCount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.RenounceOwnership(&_ZksBaseRegister.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.RenounceOwnership(&_ZksBaseRegister.TransactOpts)
}

// SetBase is a paid mutator transaction binding the contract method 0xbe17fe62.
//
// Solidity: function setBase(address _base) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) SetBase(opts *bind.TransactOpts, _base common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "setBase", _base)
}

// SetBase is a paid mutator transaction binding the contract method 0xbe17fe62.
//
// Solidity: function setBase(address _base) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) SetBase(_base common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetBase(&_ZksBaseRegister.TransactOpts, _base)
}

// SetBase is a paid mutator transaction binding the contract method 0xbe17fe62.
//
// Solidity: function setBase(address _base) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) SetBase(_base common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetBase(&_ZksBaseRegister.TransactOpts, _base)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 nameLength, uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) SetPrice(opts *bind.TransactOpts, nameLength *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "setPrice", nameLength, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 nameLength, uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) SetPrice(nameLength *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetPrice(&_ZksBaseRegister.TransactOpts, nameLength, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0xf7d97577.
//
// Solidity: function setPrice(uint256 nameLength, uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) SetPrice(nameLength *big.Int, price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetPrice(&_ZksBaseRegister.TransactOpts, nameLength, price)
}

// SetTeamAddress is a paid mutator transaction binding the contract method 0x6690864e.
//
// Solidity: function setTeamAddress(address _teamAddress) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) SetTeamAddress(opts *bind.TransactOpts, _teamAddress common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "setTeamAddress", _teamAddress)
}

// SetTeamAddress is a paid mutator transaction binding the contract method 0x6690864e.
//
// Solidity: function setTeamAddress(address _teamAddress) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) SetTeamAddress(_teamAddress common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetTeamAddress(&_ZksBaseRegister.TransactOpts, _teamAddress)
}

// SetTeamAddress is a paid mutator transaction binding the contract method 0x6690864e.
//
// Solidity: function setTeamAddress(address _teamAddress) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) SetTeamAddress(_teamAddress common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetTeamAddress(&_ZksBaseRegister.TransactOpts, _teamAddress)
}

// SetWLPriority is a paid mutator transaction binding the contract method 0x64a0588b.
//
// Solidity: function setWLPriority(uint256 _time) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) SetWLPriority(opts *bind.TransactOpts, _time *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "setWLPriority", _time)
}

// SetWLPriority is a paid mutator transaction binding the contract method 0x64a0588b.
//
// Solidity: function setWLPriority(uint256 _time) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) SetWLPriority(_time *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetWLPriority(&_ZksBaseRegister.TransactOpts, _time)
}

// SetWLPriority is a paid mutator transaction binding the contract method 0x64a0588b.
//
// Solidity: function setWLPriority(uint256 _time) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) SetWLPriority(_time *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetWLPriority(&_ZksBaseRegister.TransactOpts, _time)
}

// SetYearlyBasePrice is a paid mutator transaction binding the contract method 0x8580f341.
//
// Solidity: function setYearlyBasePrice(uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) SetYearlyBasePrice(opts *bind.TransactOpts, price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "setYearlyBasePrice", price)
}

// SetYearlyBasePrice is a paid mutator transaction binding the contract method 0x8580f341.
//
// Solidity: function setYearlyBasePrice(uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) SetYearlyBasePrice(price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetYearlyBasePrice(&_ZksBaseRegister.TransactOpts, price)
}

// SetYearlyBasePrice is a paid mutator transaction binding the contract method 0x8580f341.
//
// Solidity: function setYearlyBasePrice(uint256 price) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) SetYearlyBasePrice(price *big.Int) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.SetYearlyBasePrice(&_ZksBaseRegister.TransactOpts, price)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.TransferOwnership(&_ZksBaseRegister.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.TransferOwnership(&_ZksBaseRegister.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) Withdraw() (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Withdraw(&_ZksBaseRegister.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) Withdraw() (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.Withdraw(&_ZksBaseRegister.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _tokenContract) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactor) WithdrawToken(opts *bind.TransactOpts, _tokenContract common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.contract.Transact(opts, "withdrawToken", _tokenContract)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _tokenContract) returns()
func (_ZksBaseRegister *ZksBaseRegisterSession) WithdrawToken(_tokenContract common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.WithdrawToken(&_ZksBaseRegister.TransactOpts, _tokenContract)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0x89476069.
//
// Solidity: function withdrawToken(address _tokenContract) returns()
func (_ZksBaseRegister *ZksBaseRegisterTransactorSession) WithdrawToken(_tokenContract common.Address) (*types.Transaction, error) {
	return _ZksBaseRegister.Contract.WithdrawToken(&_ZksBaseRegister.TransactOpts, _tokenContract)
}

// ZksBaseRegisterNameRegisteredIterator is returned from FilterNameRegistered and is used to iterate over the raw logs and unpacked data for NameRegistered events raised by the ZksBaseRegister contract.
type ZksBaseRegisterNameRegisteredIterator struct {
	Event *ZksBaseRegisterNameRegistered // Event containing the contract specifics and raw log

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
func (it *ZksBaseRegisterNameRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksBaseRegisterNameRegistered)
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
		it.Event = new(ZksBaseRegisterNameRegistered)
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
func (it *ZksBaseRegisterNameRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksBaseRegisterNameRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksBaseRegisterNameRegistered represents a NameRegistered event raised by the ZksBaseRegister contract.
type ZksBaseRegisterNameRegistered struct {
	Name    string
	Label   [32]byte
	Owner   common.Address
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRegistered is a free log retrieval operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) FilterNameRegistered(opts *bind.FilterOpts, label [][32]byte, owner []common.Address) (*ZksBaseRegisterNameRegisteredIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.FilterLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterNameRegisteredIterator{contract: _ZksBaseRegister.contract, event: "NameRegistered", logs: logs, sub: sub}, nil
}

// WatchNameRegistered is a free log subscription operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) WatchNameRegistered(opts *bind.WatchOpts, sink chan<- *ZksBaseRegisterNameRegistered, label [][32]byte, owner []common.Address) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.WatchLogs(opts, "NameRegistered", labelRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksBaseRegisterNameRegistered)
				if err := _ZksBaseRegister.contract.UnpackLog(event, "NameRegistered", log); err != nil {
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

// ParseNameRegistered is a log parse operation binding the contract event 0xca6abbe9d7f11422cb6ca7629fbf6fe9efb1c621f71ce8f02b9f2a230097404f.
//
// Solidity: event NameRegistered(string name, bytes32 indexed label, address indexed owner, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) ParseNameRegistered(log types.Log) (*ZksBaseRegisterNameRegistered, error) {
	event := new(ZksBaseRegisterNameRegistered)
	if err := _ZksBaseRegister.contract.UnpackLog(event, "NameRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZksBaseRegisterNameRenewedIterator is returned from FilterNameRenewed and is used to iterate over the raw logs and unpacked data for NameRenewed events raised by the ZksBaseRegister contract.
type ZksBaseRegisterNameRenewedIterator struct {
	Event *ZksBaseRegisterNameRenewed // Event containing the contract specifics and raw log

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
func (it *ZksBaseRegisterNameRenewedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksBaseRegisterNameRenewed)
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
		it.Event = new(ZksBaseRegisterNameRenewed)
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
func (it *ZksBaseRegisterNameRenewedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksBaseRegisterNameRenewedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksBaseRegisterNameRenewed represents a NameRenewed event raised by the ZksBaseRegister contract.
type ZksBaseRegisterNameRenewed struct {
	Name    string
	Label   [32]byte
	Cost    *big.Int
	Expires *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNameRenewed is a free log retrieval operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) FilterNameRenewed(opts *bind.FilterOpts, label [][32]byte) (*ZksBaseRegisterNameRenewedIterator, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.FilterLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterNameRenewedIterator{contract: _ZksBaseRegister.contract, event: "NameRenewed", logs: logs, sub: sub}, nil
}

// WatchNameRenewed is a free log subscription operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) WatchNameRenewed(opts *bind.WatchOpts, sink chan<- *ZksBaseRegisterNameRenewed, label [][32]byte) (event.Subscription, error) {

	var labelRule []interface{}
	for _, labelItem := range label {
		labelRule = append(labelRule, labelItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.WatchLogs(opts, "NameRenewed", labelRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksBaseRegisterNameRenewed)
				if err := _ZksBaseRegister.contract.UnpackLog(event, "NameRenewed", log); err != nil {
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

// ParseNameRenewed is a log parse operation binding the contract event 0x3da24c024582931cfaf8267d8ed24d13a82a8068d5bd337d30ec45cea4e506ae.
//
// Solidity: event NameRenewed(string name, bytes32 indexed label, uint256 cost, uint256 expires)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) ParseNameRenewed(log types.Log) (*ZksBaseRegisterNameRenewed, error) {
	event := new(ZksBaseRegisterNameRenewed)
	if err := _ZksBaseRegister.contract.UnpackLog(event, "NameRenewed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZksBaseRegisterNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the ZksBaseRegister contract.
type ZksBaseRegisterNewPriceOracleIterator struct {
	Event *ZksBaseRegisterNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ZksBaseRegisterNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksBaseRegisterNewPriceOracle)
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
		it.Event = new(ZksBaseRegisterNewPriceOracle)
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
func (it *ZksBaseRegisterNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksBaseRegisterNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksBaseRegisterNewPriceOracle represents a NewPriceOracle event raised by the ZksBaseRegister contract.
type ZksBaseRegisterNewPriceOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) FilterNewPriceOracle(opts *bind.FilterOpts, oracle []common.Address) (*ZksBaseRegisterNewPriceOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.FilterLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterNewPriceOracleIterator{contract: _ZksBaseRegister.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ZksBaseRegisterNewPriceOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.WatchLogs(opts, "NewPriceOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksBaseRegisterNewPriceOracle)
				if err := _ZksBaseRegister.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xf261845a790fe29bbd6631e2ca4a5bdc83e6eed7c3271d9590d97287e00e9123.
//
// Solidity: event NewPriceOracle(address indexed oracle)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) ParseNewPriceOracle(log types.Log) (*ZksBaseRegisterNewPriceOracle, error) {
	event := new(ZksBaseRegisterNewPriceOracle)
	if err := _ZksBaseRegister.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ZksBaseRegisterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ZksBaseRegister contract.
type ZksBaseRegisterOwnershipTransferredIterator struct {
	Event *ZksBaseRegisterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ZksBaseRegisterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ZksBaseRegisterOwnershipTransferred)
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
		it.Event = new(ZksBaseRegisterOwnershipTransferred)
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
func (it *ZksBaseRegisterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ZksBaseRegisterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ZksBaseRegisterOwnershipTransferred represents a OwnershipTransferred event raised by the ZksBaseRegister contract.
type ZksBaseRegisterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ZksBaseRegisterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ZksBaseRegisterOwnershipTransferredIterator{contract: _ZksBaseRegister.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ZksBaseRegister *ZksBaseRegisterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ZksBaseRegisterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ZksBaseRegister.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ZksBaseRegisterOwnershipTransferred)
				if err := _ZksBaseRegister.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ZksBaseRegister *ZksBaseRegisterFilterer) ParseOwnershipTransferred(log types.Log) (*ZksBaseRegisterOwnershipTransferred, error) {
	event := new(ZksBaseRegisterOwnershipTransferred)
	if err := _ZksBaseRegister.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
