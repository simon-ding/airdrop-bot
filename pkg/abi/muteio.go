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

// MuteIoMetaData contains all meta data concerning the MuteIo contract.
var MuteIoMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_factory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountADesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenDesired\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeType\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"addLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"getAmountsOutExpanded\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"},{\"internalType\":\"uint256[]\",\"name\":\"fees\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getPairInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"pairFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveB\",\"type\":\"uint256\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountAMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountBMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountB\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountTokenMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountETHMin\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"stable\",\"type\":\"bool\"}],\"name\":\"removeLiquidityETHSupportingFeeOnTransferTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB\",\"type\":\"address\"}],\"name\":\"sortTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token1\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactETHForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForETHSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bool[]\",\"name\":\"stable\",\"type\":\"bool[]\"}],\"name\":\"swapExactTokensForTokensSupportingFeeOnTransferTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MuteIoABI is the input ABI used to generate the binding from.
// Deprecated: Use MuteIoMetaData.ABI instead.
var MuteIoABI = MuteIoMetaData.ABI

// MuteIo is an auto generated Go binding around an Ethereum contract.
type MuteIo struct {
	MuteIoCaller     // Read-only binding to the contract
	MuteIoTransactor // Write-only binding to the contract
	MuteIoFilterer   // Log filterer for contract events
}

// MuteIoCaller is an auto generated read-only Go binding around an Ethereum contract.
type MuteIoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MuteIoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MuteIoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MuteIoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MuteIoSession struct {
	Contract     *MuteIo           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuteIoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MuteIoCallerSession struct {
	Contract *MuteIoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MuteIoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MuteIoTransactorSession struct {
	Contract     *MuteIoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MuteIoRaw is an auto generated low-level Go binding around an Ethereum contract.
type MuteIoRaw struct {
	Contract *MuteIo // Generic contract binding to access the raw methods on
}

// MuteIoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MuteIoCallerRaw struct {
	Contract *MuteIoCaller // Generic read-only contract binding to access the raw methods on
}

// MuteIoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MuteIoTransactorRaw struct {
	Contract *MuteIoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMuteIo creates a new instance of MuteIo, bound to a specific deployed contract.
func NewMuteIo(address common.Address, backend bind.ContractBackend) (*MuteIo, error) {
	contract, err := bindMuteIo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MuteIo{MuteIoCaller: MuteIoCaller{contract: contract}, MuteIoTransactor: MuteIoTransactor{contract: contract}, MuteIoFilterer: MuteIoFilterer{contract: contract}}, nil
}

// NewMuteIoCaller creates a new read-only instance of MuteIo, bound to a specific deployed contract.
func NewMuteIoCaller(address common.Address, caller bind.ContractCaller) (*MuteIoCaller, error) {
	contract, err := bindMuteIo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MuteIoCaller{contract: contract}, nil
}

// NewMuteIoTransactor creates a new write-only instance of MuteIo, bound to a specific deployed contract.
func NewMuteIoTransactor(address common.Address, transactor bind.ContractTransactor) (*MuteIoTransactor, error) {
	contract, err := bindMuteIo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MuteIoTransactor{contract: contract}, nil
}

// NewMuteIoFilterer creates a new log filterer instance of MuteIo, bound to a specific deployed contract.
func NewMuteIoFilterer(address common.Address, filterer bind.ContractFilterer) (*MuteIoFilterer, error) {
	contract, err := bindMuteIo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MuteIoFilterer{contract: contract}, nil
}

// bindMuteIo binds a generic wrapper to an already deployed contract.
func bindMuteIo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MuteIoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuteIo *MuteIoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuteIo.Contract.MuteIoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuteIo *MuteIoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuteIo.Contract.MuteIoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuteIo *MuteIoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuteIo.Contract.MuteIoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MuteIo *MuteIoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MuteIo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MuteIo *MuteIoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MuteIo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MuteIo *MuteIoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MuteIo.Contract.contract.Transact(opts, method, params...)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MuteIo *MuteIoCaller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MuteIo *MuteIoSession) WETH() (common.Address, error) {
	return _MuteIo.Contract.WETH(&_MuteIo.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_MuteIo *MuteIoCallerSession) WETH() (common.Address, error) {
	return _MuteIo.Contract.WETH(&_MuteIo.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MuteIo *MuteIoCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MuteIo *MuteIoSession) Factory() (common.Address, error) {
	return _MuteIo.Contract.Factory(&_MuteIo.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_MuteIo *MuteIoCallerSession) Factory() (common.Address, error) {
	return _MuteIo.Contract.Factory(&_MuteIo.CallOpts)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_MuteIo *MuteIoCaller) GetAmountOut(opts *bind.CallOpts, amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "getAmountOut", amountIn, tokenIn, tokenOut)

	outstruct := new(struct {
		AmountOut *big.Int
		Stable    bool
		Fee       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountOut = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_MuteIo *MuteIoSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _MuteIo.Contract.GetAmountOut(&_MuteIo.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x5e1e6325.
//
// Solidity: function getAmountOut(uint256 amountIn, address tokenIn, address tokenOut) view returns(uint256 amountOut, bool stable, uint256 fee)
func (_MuteIo *MuteIoCallerSession) GetAmountOut(amountIn *big.Int, tokenIn common.Address, tokenOut common.Address) (struct {
	AmountOut *big.Int
	Stable    bool
	Fee       *big.Int
}, error) {
	return _MuteIo.Contract.GetAmountOut(&_MuteIo.CallOpts, amountIn, tokenIn, tokenOut)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_MuteIo *MuteIoCaller) GetAmountsOut(opts *bind.CallOpts, amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "getAmountsOut", amountIn, path, stable)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_MuteIo *MuteIoSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	return _MuteIo.Contract.GetAmountsOut(&_MuteIo.CallOpts, amountIn, path, stable)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x84e21aff.
//
// Solidity: function getAmountsOut(uint256 amountIn, address[] path, bool[] stable) view returns(uint256[] amounts)
func (_MuteIo *MuteIoCallerSession) GetAmountsOut(amountIn *big.Int, path []common.Address, stable []bool) ([]*big.Int, error) {
	return _MuteIo.Contract.GetAmountsOut(&_MuteIo.CallOpts, amountIn, path, stable)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_MuteIo *MuteIoCaller) GetAmountsOutExpanded(opts *bind.CallOpts, amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "getAmountsOutExpanded", amountIn, path)

	outstruct := new(struct {
		Amounts []*big.Int
		Stable  []bool
		Fees    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amounts = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Stable = *abi.ConvertType(out[1], new([]bool)).(*[]bool)
	outstruct.Fees = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_MuteIo *MuteIoSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _MuteIo.Contract.GetAmountsOutExpanded(&_MuteIo.CallOpts, amountIn, path)
}

// GetAmountsOutExpanded is a free data retrieval call binding the contract method 0xcad9446c.
//
// Solidity: function getAmountsOutExpanded(uint256 amountIn, address[] path) view returns(uint256[] amounts, bool[] stable, uint256[] fees)
func (_MuteIo *MuteIoCallerSession) GetAmountsOutExpanded(amountIn *big.Int, path []common.Address) (struct {
	Amounts []*big.Int
	Stable  []bool
	Fees    []*big.Int
}, error) {
	return _MuteIo.Contract.GetAmountsOutExpanded(&_MuteIo.CallOpts, amountIn, path)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_MuteIo *MuteIoCaller) GetPairInfo(opts *bind.CallOpts, path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "getPairInfo", path, stable)

	outstruct := new(struct {
		TokenA   common.Address
		TokenB   common.Address
		Pair     common.Address
		ReserveA *big.Int
		ReserveB *big.Int
		Fee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenA = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenB = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Pair = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.ReserveA = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_MuteIo *MuteIoSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _MuteIo.Contract.GetPairInfo(&_MuteIo.CallOpts, path, stable)
}

// GetPairInfo is a free data retrieval call binding the contract method 0xad97ec1b.
//
// Solidity: function getPairInfo(address[] path, bool stable) view returns(address tokenA, address tokenB, address pair, uint256 reserveA, uint256 reserveB, uint256 fee)
func (_MuteIo *MuteIoCallerSession) GetPairInfo(path []common.Address, stable bool) (struct {
	TokenA   common.Address
	TokenB   common.Address
	Pair     common.Address
	ReserveA *big.Int
	ReserveB *big.Int
	Fee      *big.Int
}, error) {
	return _MuteIo.Contract.GetPairInfo(&_MuteIo.CallOpts, path, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_MuteIo *MuteIoCaller) GetReserves(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "getReserves", tokenA, tokenB, stable)

	outstruct := new(struct {
		ReserveA *big.Int
		ReserveB *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_MuteIo *MuteIoSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _MuteIo.Contract.GetReserves(&_MuteIo.CallOpts, tokenA, tokenB, stable)
}

// GetReserves is a free data retrieval call binding the contract method 0x5e60dab5.
//
// Solidity: function getReserves(address tokenA, address tokenB, bool stable) view returns(uint256 reserveA, uint256 reserveB)
func (_MuteIo *MuteIoCallerSession) GetReserves(tokenA common.Address, tokenB common.Address, stable bool) (struct {
	ReserveA *big.Int
	ReserveB *big.Int
}, error) {
	return _MuteIo.Contract.GetReserves(&_MuteIo.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_MuteIo *MuteIoCaller) PairFor(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "pairFor", tokenA, tokenB, stable)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_MuteIo *MuteIoSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _MuteIo.Contract.PairFor(&_MuteIo.CallOpts, tokenA, tokenB, stable)
}

// PairFor is a free data retrieval call binding the contract method 0x4c1ee03e.
//
// Solidity: function pairFor(address tokenA, address tokenB, bool stable) view returns(address pair)
func (_MuteIo *MuteIoCallerSession) PairFor(tokenA common.Address, tokenB common.Address, stable bool) (common.Address, error) {
	return _MuteIo.Contract.PairFor(&_MuteIo.CallOpts, tokenA, tokenB, stable)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_MuteIo *MuteIoCaller) Quote(opts *bind.CallOpts, amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "quote", amountA, reserveA, reserveB)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_MuteIo *MuteIoSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _MuteIo.Contract.Quote(&_MuteIo.CallOpts, amountA, reserveA, reserveB)
}

// Quote is a free data retrieval call binding the contract method 0xad615dec.
//
// Solidity: function quote(uint256 amountA, uint256 reserveA, uint256 reserveB) pure returns(uint256 amountB)
func (_MuteIo *MuteIoCallerSession) Quote(amountA *big.Int, reserveA *big.Int, reserveB *big.Int) (*big.Int, error) {
	return _MuteIo.Contract.Quote(&_MuteIo.CallOpts, amountA, reserveA, reserveB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_MuteIo *MuteIoCaller) SortTokens(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	var out []interface{}
	err := _MuteIo.contract.Call(opts, &out, "sortTokens", tokenA, tokenB)

	outstruct := new(struct {
		Token0 common.Address
		Token1 common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token0 = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_MuteIo *MuteIoSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _MuteIo.Contract.SortTokens(&_MuteIo.CallOpts, tokenA, tokenB)
}

// SortTokens is a free data retrieval call binding the contract method 0x544caa56.
//
// Solidity: function sortTokens(address tokenA, address tokenB) pure returns(address token0, address token1)
func (_MuteIo *MuteIoCallerSession) SortTokens(tokenA common.Address, tokenB common.Address) (struct {
	Token0 common.Address
	Token1 common.Address
}, error) {
	return _MuteIo.Contract.SortTokens(&_MuteIo.CallOpts, tokenA, tokenB)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_MuteIo *MuteIoTransactor) AddLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "addLiquidity", tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_MuteIo *MuteIoSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.AddLiquidity(&_MuteIo.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x1ead1418.
//
// Solidity: function addLiquidity(address tokenA, address tokenB, uint256 amountADesired, uint256 amountBDesired, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, uint256 feeType, bool stable) returns(uint256 amountA, uint256 amountB, uint256 liquidity)
func (_MuteIo *MuteIoTransactorSession) AddLiquidity(tokenA common.Address, tokenB common.Address, amountADesired *big.Int, amountBDesired *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.AddLiquidity(&_MuteIo.TransactOpts, tokenA, tokenB, amountADesired, amountBDesired, amountAMin, amountBMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_MuteIo *MuteIoTransactor) AddLiquidityETH(opts *bind.TransactOpts, token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "addLiquidityETH", token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_MuteIo *MuteIoSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.AddLiquidityETH(&_MuteIo.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// AddLiquidityETH is a paid mutator transaction binding the contract method 0x3a8e53ff.
//
// Solidity: function addLiquidityETH(address token, uint256 amountTokenDesired, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, uint256 feeType, bool stable) payable returns(uint256 amountToken, uint256 amountETH, uint256 liquidity)
func (_MuteIo *MuteIoTransactorSession) AddLiquidityETH(token common.Address, amountTokenDesired *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, feeType *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.AddLiquidityETH(&_MuteIo.TransactOpts, token, amountTokenDesired, amountTokenMin, amountETHMin, to, deadline, feeType, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_MuteIo *MuteIoTransactor) RemoveLiquidity(opts *bind.TransactOpts, tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "removeLiquidity", tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_MuteIo *MuteIoSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidity(&_MuteIo.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x0fc87d25.
//
// Solidity: function removeLiquidity(address tokenA, address tokenB, uint256 liquidity, uint256 amountAMin, uint256 amountBMin, address to, uint256 deadline, bool stable) returns(uint256 amountA, uint256 amountB)
func (_MuteIo *MuteIoTransactorSession) RemoveLiquidity(tokenA common.Address, tokenB common.Address, liquidity *big.Int, amountAMin *big.Int, amountBMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidity(&_MuteIo.TransactOpts, tokenA, tokenB, liquidity, amountAMin, amountBMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_MuteIo *MuteIoTransactor) RemoveLiquidityETH(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "removeLiquidityETH", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_MuteIo *MuteIoSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidityETH(&_MuteIo.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETH is a paid mutator transaction binding the contract method 0xe43b4ee2.
//
// Solidity: function removeLiquidityETH(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountToken, uint256 amountETH)
func (_MuteIo *MuteIoTransactorSession) RemoveLiquidityETH(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidityETH(&_MuteIo.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_MuteIo *MuteIoTransactor) RemoveLiquidityETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "removeLiquidityETHSupportingFeeOnTransferTokens", token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_MuteIo *MuteIoSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// RemoveLiquidityETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0xaac57b19.
//
// Solidity: function removeLiquidityETHSupportingFeeOnTransferTokens(address token, uint256 liquidity, uint256 amountTokenMin, uint256 amountETHMin, address to, uint256 deadline, bool stable) returns(uint256 amountETH)
func (_MuteIo *MuteIoTransactorSession) RemoveLiquidityETHSupportingFeeOnTransferTokens(token common.Address, liquidity *big.Int, amountTokenMin *big.Int, amountETHMin *big.Int, to common.Address, deadline *big.Int, stable bool) (*types.Transaction, error) {
	return _MuteIo.Contract.RemoveLiquidityETHSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, token, liquidity, amountTokenMin, amountETHMin, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactETHForTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_MuteIo *MuteIoSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactETHForTokens(&_MuteIo.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x0bbaa8d2.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactorSession) SwapExactETHForTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactETHForTokens(&_MuteIo.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_MuteIo *MuteIoTransactor) SwapExactETHForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactETHForTokensSupportingFeeOnTransferTokens", amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_MuteIo *MuteIoSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactETHForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x51cbf10f.
//
// Solidity: function swapExactETHForTokensSupportingFeeOnTransferTokens(uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) payable returns()
func (_MuteIo *MuteIoTransactorSession) SwapExactETHForTokensSupportingFeeOnTransferTokens(amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactETHForTokensSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactTokensForETH", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForETH(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0xacc8723d.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactorSession) SwapExactTokensForETH(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForETH(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoTransactor) SwapExactTokensForETHSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactTokensForETHSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForETHSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x3f464b16.
//
// Solidity: function swapExactTokensForETHSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoTransactorSession) SwapExactTokensForETHSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForETHSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactTokensForTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0xc694d55d.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns(uint256[] amounts)
func (_MuteIo *MuteIoTransactorSession) SwapExactTokensForTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoTransactor) SwapExactTokensForTokensSupportingFeeOnTransferTokens(opts *bind.TransactOpts, amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.contract.Transact(opts, "swapExactTokensForTokensSupportingFeeOnTransferTokens", amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// SwapExactTokensForTokensSupportingFeeOnTransferTokens is a paid mutator transaction binding the contract method 0x19a13c5c.
//
// Solidity: function swapExactTokensForTokensSupportingFeeOnTransferTokens(uint256 amountIn, uint256 amountOutMin, address[] path, address to, uint256 deadline, bool[] stable) returns()
func (_MuteIo *MuteIoTransactorSession) SwapExactTokensForTokensSupportingFeeOnTransferTokens(amountIn *big.Int, amountOutMin *big.Int, path []common.Address, to common.Address, deadline *big.Int, stable []bool) (*types.Transaction, error) {
	return _MuteIo.Contract.SwapExactTokensForTokensSupportingFeeOnTransferTokens(&_MuteIo.TransactOpts, amountIn, amountOutMin, path, to, deadline, stable)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MuteIo *MuteIoTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _MuteIo.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MuteIo *MuteIoSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MuteIo.Contract.Fallback(&_MuteIo.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_MuteIo *MuteIoTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _MuteIo.Contract.Fallback(&_MuteIo.TransactOpts, calldata)
}
