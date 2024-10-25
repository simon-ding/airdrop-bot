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

// // IPoolTokenAmount is an auto generated low-level Go binding around an user-defined struct.
// type IPoolTokenAmount struct {
// 	Token  common.Address
// 	Amount *big.Int
// }

// SyncSwapClassicPoolMetaData contains all meta data concerning the SyncSwapClassicPool contract.
var SyncSwapClassicPoolMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Expired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientLiquidityMinted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Overflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1In\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0Out\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1Out\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reserve1\",\"type\":\"uint256\"}],\"name\":\"Sync\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount[]\",\"name\":\"_amounts\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"burnSingle\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAssets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_protocolFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_reserve0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_reserve1\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"_swapFee\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"invariantLast\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"master\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"_v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_signature\",\"type\":\"bytes\"}],\"name\":\"permit2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve0\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_callback\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_callbackData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structIPool.TokenAmount\",\"name\":\"_tokenAmount\",\"type\":\"tuple\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token0\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token1\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SyncSwapClassicPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use SyncSwapClassicPoolMetaData.ABI instead.
var SyncSwapClassicPoolABI = SyncSwapClassicPoolMetaData.ABI

// SyncSwapClassicPool is an auto generated Go binding around an Ethereum contract.
type SyncSwapClassicPool struct {
	SyncSwapClassicPoolCaller     // Read-only binding to the contract
	SyncSwapClassicPoolTransactor // Write-only binding to the contract
	SyncSwapClassicPoolFilterer   // Log filterer for contract events
}

// SyncSwapClassicPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type SyncSwapClassicPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncSwapClassicPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SyncSwapClassicPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncSwapClassicPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SyncSwapClassicPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SyncSwapClassicPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SyncSwapClassicPoolSession struct {
	Contract     *SyncSwapClassicPool // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SyncSwapClassicPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SyncSwapClassicPoolCallerSession struct {
	Contract *SyncSwapClassicPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SyncSwapClassicPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SyncSwapClassicPoolTransactorSession struct {
	Contract     *SyncSwapClassicPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SyncSwapClassicPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type SyncSwapClassicPoolRaw struct {
	Contract *SyncSwapClassicPool // Generic contract binding to access the raw methods on
}

// SyncSwapClassicPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SyncSwapClassicPoolCallerRaw struct {
	Contract *SyncSwapClassicPoolCaller // Generic read-only contract binding to access the raw methods on
}

// SyncSwapClassicPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SyncSwapClassicPoolTransactorRaw struct {
	Contract *SyncSwapClassicPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSyncSwapClassicPool creates a new instance of SyncSwapClassicPool, bound to a specific deployed contract.
func NewSyncSwapClassicPool(address common.Address, backend bind.ContractBackend) (*SyncSwapClassicPool, error) {
	contract, err := bindSyncSwapClassicPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPool{SyncSwapClassicPoolCaller: SyncSwapClassicPoolCaller{contract: contract}, SyncSwapClassicPoolTransactor: SyncSwapClassicPoolTransactor{contract: contract}, SyncSwapClassicPoolFilterer: SyncSwapClassicPoolFilterer{contract: contract}}, nil
}

// NewSyncSwapClassicPoolCaller creates a new read-only instance of SyncSwapClassicPool, bound to a specific deployed contract.
func NewSyncSwapClassicPoolCaller(address common.Address, caller bind.ContractCaller) (*SyncSwapClassicPoolCaller, error) {
	contract, err := bindSyncSwapClassicPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolCaller{contract: contract}, nil
}

// NewSyncSwapClassicPoolTransactor creates a new write-only instance of SyncSwapClassicPool, bound to a specific deployed contract.
func NewSyncSwapClassicPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*SyncSwapClassicPoolTransactor, error) {
	contract, err := bindSyncSwapClassicPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolTransactor{contract: contract}, nil
}

// NewSyncSwapClassicPoolFilterer creates a new log filterer instance of SyncSwapClassicPool, bound to a specific deployed contract.
func NewSyncSwapClassicPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*SyncSwapClassicPoolFilterer, error) {
	contract, err := bindSyncSwapClassicPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolFilterer{contract: contract}, nil
}

// bindSyncSwapClassicPool binds a generic wrapper to an already deployed contract.
func bindSyncSwapClassicPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SyncSwapClassicPoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncSwapClassicPool *SyncSwapClassicPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncSwapClassicPool.Contract.SyncSwapClassicPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncSwapClassicPool *SyncSwapClassicPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.SyncSwapClassicPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncSwapClassicPool *SyncSwapClassicPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.SyncSwapClassicPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SyncSwapClassicPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _SyncSwapClassicPool.Contract.DOMAINSEPARATOR(&_SyncSwapClassicPool.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _SyncSwapClassicPool.Contract.DOMAINSEPARATOR(&_SyncSwapClassicPool.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "allowance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Allowance(&_SyncSwapClassicPool.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address , address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Allowance(&_SyncSwapClassicPool.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.BalanceOf(&_SyncSwapClassicPool.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.BalanceOf(&_SyncSwapClassicPool.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Decimals() (uint8, error) {
	return _SyncSwapClassicPool.Contract.Decimals(&_SyncSwapClassicPool.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Decimals() (uint8, error) {
	return _SyncSwapClassicPool.Contract.Decimals(&_SyncSwapClassicPool.CallOpts)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetAmountIn(opts *bind.CallOpts, _tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getAmountIn", _tokenOut, _amountOut, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetAmountIn(&_SyncSwapClassicPool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountIn is a free data retrieval call binding the contract method 0xa287c795.
//
// Solidity: function getAmountIn(address _tokenOut, uint256 _amountOut, address _sender) view returns(uint256 _amountIn)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetAmountIn(_tokenOut common.Address, _amountOut *big.Int, _sender common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetAmountIn(&_SyncSwapClassicPool.CallOpts, _tokenOut, _amountOut, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetAmountOut(opts *bind.CallOpts, _tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getAmountOut", _tokenIn, _amountIn, _sender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetAmountOut(&_SyncSwapClassicPool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAmountOut is a free data retrieval call binding the contract method 0xff9c8ac6.
//
// Solidity: function getAmountOut(address _tokenIn, uint256 _amountIn, address _sender) view returns(uint256 _amountOut)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetAmountOut(_tokenIn common.Address, _amountIn *big.Int, _sender common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetAmountOut(&_SyncSwapClassicPool.CallOpts, _tokenIn, _amountIn, _sender)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetAssets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getAssets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetAssets() ([]common.Address, error) {
	return _SyncSwapClassicPool.Contract.GetAssets(&_SyncSwapClassicPool.CallOpts)
}

// GetAssets is a free data retrieval call binding the contract method 0x67e4ac2c.
//
// Solidity: function getAssets() view returns(address[] assets)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetAssets() ([]common.Address, error) {
	return _SyncSwapClassicPool.Contract.GetAssets(&_SyncSwapClassicPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getProtocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetProtocolFee() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetProtocolFee(&_SyncSwapClassicPool.CallOpts)
}

// GetProtocolFee is a free data retrieval call binding the contract method 0xa5a41031.
//
// Solidity: function getProtocolFee() view returns(uint24 _protocolFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetProtocolFee() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetProtocolFee(&_SyncSwapClassicPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetReserves(opts *bind.CallOpts) (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		Reserve0 *big.Int
		Reserve1 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Reserve0 = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Reserve1 = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _SyncSwapClassicPool.Contract.GetReserves(&_SyncSwapClassicPool.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint256 _reserve0, uint256 _reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetReserves() (struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
}, error) {
	return _SyncSwapClassicPool.Contract.GetReserves(&_SyncSwapClassicPool.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) GetSwapFee(opts *bind.CallOpts, _sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "getSwapFee", _sender, _tokenIn, _tokenOut, data)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetSwapFee(&_SyncSwapClassicPool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// GetSwapFee is a free data retrieval call binding the contract method 0x8b4c5470.
//
// Solidity: function getSwapFee(address _sender, address _tokenIn, address _tokenOut, bytes data) view returns(uint24 _swapFee)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) GetSwapFee(_sender common.Address, _tokenIn common.Address, _tokenOut common.Address, data []byte) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.GetSwapFee(&_SyncSwapClassicPool.CallOpts, _sender, _tokenIn, _tokenOut, data)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) InvariantLast(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "invariantLast")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) InvariantLast() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.InvariantLast(&_SyncSwapClassicPool.CallOpts)
}

// InvariantLast is a free data retrieval call binding the contract method 0x07f293f7.
//
// Solidity: function invariantLast() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) InvariantLast() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.InvariantLast(&_SyncSwapClassicPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Master(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "master")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Master() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Master(&_SyncSwapClassicPool.CallOpts)
}

// Master is a free data retrieval call binding the contract method 0xee97f7f3.
//
// Solidity: function master() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Master() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Master(&_SyncSwapClassicPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Name() (string, error) {
	return _SyncSwapClassicPool.Contract.Name(&_SyncSwapClassicPool.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Name() (string, error) {
	return _SyncSwapClassicPool.Contract.Name(&_SyncSwapClassicPool.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Nonces(&_SyncSwapClassicPool.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Nonces(&_SyncSwapClassicPool.CallOpts, arg0)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) PoolType(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) PoolType() (uint16, error) {
	return _SyncSwapClassicPool.Contract.PoolType(&_SyncSwapClassicPool.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(uint16)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) PoolType() (uint16, error) {
	return _SyncSwapClassicPool.Contract.PoolType(&_SyncSwapClassicPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Reserve0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "reserve0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Reserve0() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Reserve0(&_SyncSwapClassicPool.CallOpts)
}

// Reserve0 is a free data retrieval call binding the contract method 0x443cb4bc.
//
// Solidity: function reserve0() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Reserve0() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Reserve0(&_SyncSwapClassicPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Reserve1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "reserve1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Reserve1() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Reserve1(&_SyncSwapClassicPool.CallOpts)
}

// Reserve1 is a free data retrieval call binding the contract method 0x5a76f25e.
//
// Solidity: function reserve1() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Reserve1() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.Reserve1(&_SyncSwapClassicPool.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) SupportsInterface(opts *bind.CallOpts, interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "supportsInterface", interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _SyncSwapClassicPool.Contract.SupportsInterface(&_SyncSwapClassicPool.CallOpts, interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceID) pure returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) SupportsInterface(interfaceID [4]byte) (bool, error) {
	return _SyncSwapClassicPool.Contract.SupportsInterface(&_SyncSwapClassicPool.CallOpts, interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Symbol() (string, error) {
	return _SyncSwapClassicPool.Contract.Symbol(&_SyncSwapClassicPool.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Symbol() (string, error) {
	return _SyncSwapClassicPool.Contract.Symbol(&_SyncSwapClassicPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Token0(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "token0")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Token0() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Token0(&_SyncSwapClassicPool.CallOpts)
}

// Token0 is a free data retrieval call binding the contract method 0x0dfe1681.
//
// Solidity: function token0() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Token0() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Token0(&_SyncSwapClassicPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Token1(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "token1")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Token1() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Token1(&_SyncSwapClassicPool.CallOpts)
}

// Token1 is a free data retrieval call binding the contract method 0xd21220a7.
//
// Solidity: function token1() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Token1() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Token1(&_SyncSwapClassicPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) TotalSupply() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.TotalSupply(&_SyncSwapClassicPool.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) TotalSupply() (*big.Int, error) {
	return _SyncSwapClassicPool.Contract.TotalSupply(&_SyncSwapClassicPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SyncSwapClassicPool.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Vault() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Vault(&_SyncSwapClassicPool.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_SyncSwapClassicPool *SyncSwapClassicPoolCallerSession) Vault() (common.Address, error) {
	return _SyncSwapClassicPool.Contract.Vault(&_SyncSwapClassicPool.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "approve", _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Approve(&_SyncSwapClassicPool.TransactOpts, _spender, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Approve(_spender common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Approve(&_SyncSwapClassicPool.TransactOpts, _spender, _amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Burn(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "burn", _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Burn(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Burn is a paid mutator transaction binding the contract method 0xf66eab5b.
//
// Solidity: function burn(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256)[] _amounts)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Burn(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Burn(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) BurnSingle(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "burnSingle", _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.BurnSingle(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// BurnSingle is a paid mutator transaction binding the contract method 0x27b0bcea.
//
// Solidity: function burnSingle(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) BurnSingle(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.BurnSingle(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Mint(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "mint", _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Mint(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Mint is a paid mutator transaction binding the contract method 0x03e7286a.
//
// Solidity: function mint(bytes _data, address _sender, address _callback, bytes _callbackData) returns(uint256)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Mint(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Mint(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Permit(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "permit", _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Permit(&_SyncSwapClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address _owner, address _spender, uint256 _amount, uint256 _deadline, uint8 _v, bytes32 _r, bytes32 _s) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Permit(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Permit(&_SyncSwapClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _v, _r, _s)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Permit2(opts *bind.TransactOpts, _owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "permit2", _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Permit2(&_SyncSwapClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Permit2 is a paid mutator transaction binding the contract method 0x2c0198cc.
//
// Solidity: function permit2(address _owner, address _spender, uint256 _amount, uint256 _deadline, bytes _signature) returns()
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Permit2(_owner common.Address, _spender common.Address, _amount *big.Int, _deadline *big.Int, _signature []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Permit2(&_SyncSwapClassicPool.TransactOpts, _owner, _spender, _amount, _deadline, _signature)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Swap(opts *bind.TransactOpts, _data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "swap", _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Swap(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Swap is a paid mutator transaction binding the contract method 0x7132bb7f.
//
// Solidity: function swap(bytes _data, address _sender, address _callback, bytes _callbackData) returns((address,uint256) _tokenAmount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Swap(_data []byte, _sender common.Address, _callback common.Address, _callbackData []byte) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Swap(&_SyncSwapClassicPool.TransactOpts, _data, _sender, _callback, _callbackData)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "transfer", _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Transfer(&_SyncSwapClassicPool.TransactOpts, _to, _amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) Transfer(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.Transfer(&_SyncSwapClassicPool.TransactOpts, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.contract.Transact(opts, "transferFrom", _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.TransferFrom(&_SyncSwapClassicPool.TransactOpts, _from, _to, _amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _amount) returns(bool)
func (_SyncSwapClassicPool *SyncSwapClassicPoolTransactorSession) TransferFrom(_from common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _SyncSwapClassicPool.Contract.TransferFrom(&_SyncSwapClassicPool.TransactOpts, _from, _to, _amount)
}

// SyncSwapClassicPoolApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolApprovalIterator struct {
	Event *SyncSwapClassicPoolApproval // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolApproval)
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
		it.Event = new(SyncSwapClassicPoolApproval)
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
func (it *SyncSwapClassicPoolApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolApproval represents a Approval event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SyncSwapClassicPoolApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolApprovalIterator{contract: _SyncSwapClassicPool.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolApproval)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseApproval(log types.Log) (*SyncSwapClassicPoolApproval, error) {
	event := new(SyncSwapClassicPoolApproval)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncSwapClassicPoolBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolBurnIterator struct {
	Event *SyncSwapClassicPoolBurn // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolBurn)
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
		it.Event = new(SyncSwapClassicPoolBurn)
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
func (it *SyncSwapClassicPoolBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolBurn represents a Burn event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolBurn struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterBurn(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SyncSwapClassicPoolBurnIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolBurnIterator{contract: _SyncSwapClassicPool.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolBurn, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Burn", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolBurn)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xd175a80c109434bb89948928ab2475a6647c94244cb70002197896423c883363.
//
// Solidity: event Burn(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseBurn(log types.Log) (*SyncSwapClassicPoolBurn, error) {
	event := new(SyncSwapClassicPoolBurn)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncSwapClassicPoolMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolMintIterator struct {
	Event *SyncSwapClassicPoolMint // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolMint)
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
		it.Event = new(SyncSwapClassicPoolMint)
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
func (it *SyncSwapClassicPoolMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolMint represents a Mint event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolMint struct {
	Sender    common.Address
	Amount0   *big.Int
	Amount1   *big.Int
	Liquidity *big.Int
	To        common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterMint(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SyncSwapClassicPoolMintIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolMintIterator{contract: _SyncSwapClassicPool.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolMint, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Mint", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolMint)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xa8137fff86647d8a402117b9c5dbda627f721d3773338fb9678c83e54ed39080.
//
// Solidity: event Mint(address indexed sender, uint256 amount0, uint256 amount1, uint256 liquidity, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseMint(log types.Log) (*SyncSwapClassicPoolMint, error) {
	event := new(SyncSwapClassicPoolMint)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncSwapClassicPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolSwapIterator struct {
	Event *SyncSwapClassicPoolSwap // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolSwap)
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
		it.Event = new(SyncSwapClassicPoolSwap)
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
func (it *SyncSwapClassicPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolSwap represents a Swap event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolSwap struct {
	Sender     common.Address
	Amount0In  *big.Int
	Amount1In  *big.Int
	Amount0Out *big.Int
	Amount1Out *big.Int
	To         common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*SyncSwapClassicPoolSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolSwapIterator{contract: _SyncSwapClassicPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolSwap)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xd78ad95fa46c994b6551d0da85fc275fe613ce37657fb8d5e3d130840159d822.
//
// Solidity: event Swap(address indexed sender, uint256 amount0In, uint256 amount1In, uint256 amount0Out, uint256 amount1Out, address indexed to)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseSwap(log types.Log) (*SyncSwapClassicPoolSwap, error) {
	event := new(SyncSwapClassicPoolSwap)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncSwapClassicPoolSyncIterator is returned from FilterSync and is used to iterate over the raw logs and unpacked data for Sync events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolSyncIterator struct {
	Event *SyncSwapClassicPoolSync // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolSyncIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolSync)
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
		it.Event = new(SyncSwapClassicPoolSync)
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
func (it *SyncSwapClassicPoolSyncIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolSyncIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolSync represents a Sync event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolSync struct {
	Reserve0 *big.Int
	Reserve1 *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSync is a free log retrieval operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterSync(opts *bind.FilterOpts) (*SyncSwapClassicPoolSyncIterator, error) {

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolSyncIterator{contract: _SyncSwapClassicPool.contract, event: "Sync", logs: logs, sub: sub}, nil
}

// WatchSync is a free log subscription operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchSync(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolSync) (event.Subscription, error) {

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Sync")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolSync)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Sync", log); err != nil {
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

// ParseSync is a log parse operation binding the contract event 0xcf2aa50876cdfbb541206f89af0ee78d44a2abf8d328e37fa4917f982149848a.
//
// Solidity: event Sync(uint256 reserve0, uint256 reserve1)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseSync(log types.Log) (*SyncSwapClassicPoolSync, error) {
	event := new(SyncSwapClassicPoolSync)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Sync", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SyncSwapClassicPoolTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolTransferIterator struct {
	Event *SyncSwapClassicPoolTransfer // Event containing the contract specifics and raw log

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
func (it *SyncSwapClassicPoolTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SyncSwapClassicPoolTransfer)
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
		it.Event = new(SyncSwapClassicPoolTransfer)
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
func (it *SyncSwapClassicPoolTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SyncSwapClassicPoolTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SyncSwapClassicPoolTransfer represents a Transfer event raised by the SyncSwapClassicPool contract.
type SyncSwapClassicPoolTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SyncSwapClassicPoolTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SyncSwapClassicPoolTransferIterator{contract: _SyncSwapClassicPool.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SyncSwapClassicPoolTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SyncSwapClassicPool.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SyncSwapClassicPoolTransfer)
				if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_SyncSwapClassicPool *SyncSwapClassicPoolFilterer) ParseTransfer(log types.Log) (*SyncSwapClassicPoolTransfer, error) {
	event := new(SyncSwapClassicPoolTransfer)
	if err := _SyncSwapClassicPool.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
