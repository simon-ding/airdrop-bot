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

// CbridgeMetaData contains all meta data concerning the Cbridge contract.
var CbridgeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"DelayPeriodUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"DelayThresholdUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"DelayedTransferAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DelayedTransferExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"length\",\"type\":\"uint256\"}],\"name\":\"EpochLengthUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cap\",\"type\":\"uint256\"}],\"name\":\"EpochVolumeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MaxSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinAddUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"MinSendUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PauserRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"srcChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"srcTransferId\",\"type\":\"bytes32\"}],\"name\":\"Relay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"resetTime\",\"type\":\"uint256\"}],\"name\":\"ResetNotification\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"transferId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"dstChainId\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxSlippage\",\"type\":\"uint32\"}],\"name\":\"Send\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"SignersUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"seqnum\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"refid\",\"type\":\"bytes32\"}],\"name\":\"WithdrawDone\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"addGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addNativeLiquidity\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"addseq\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"delayPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"delayThresholds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"delayedTransfers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumeCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"epochVolumes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"executeDelayedTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"governors\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"increaseNoticePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"isGovernor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPauser\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastOpTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minAdd\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"minSend\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimalMaxSlippage\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nativeWrap\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"noticePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"notifyResetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pausers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_relayRequest\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_account\",\"type\":\"address\"}],\"name\":\"removeGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"removePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePauser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"resetSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resetTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"send\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"_dstChainId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_nonce\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"_maxSlippage\",\"type\":\"uint32\"}],\"name\":\"sendNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_period\",\"type\":\"uint256\"}],\"name\":\"setDelayPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_thresholds\",\"type\":\"uint256[]\"}],\"name\":\"setDelayThresholds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_length\",\"type\":\"uint256\"}],\"name\":\"setEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_caps\",\"type\":\"uint256[]\"}],\"name\":\"setEpochVolumeCaps\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMaxSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinAdd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"setMinSend\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"_minimalMaxSlippage\",\"type\":\"uint32\"}],\"name\":\"setMinimalMaxSlippage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_weth\",\"type\":\"address\"}],\"name\":\"setWrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ssHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"transfers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"triggerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_triggerTime\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_newSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_newPowers\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_curSigners\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_curPowers\",\"type\":\"uint256[]\"}],\"name\":\"updateSigners\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_msg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"verifySigs\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"_wdmsg\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"_sigs\",\"type\":\"bytes[]\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_powers\",\"type\":\"uint256[]\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"withdraws\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// CbridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use CbridgeMetaData.ABI instead.
var CbridgeABI = CbridgeMetaData.ABI

// Cbridge is an auto generated Go binding around an Ethereum contract.
type Cbridge struct {
	CbridgeCaller     // Read-only binding to the contract
	CbridgeTransactor // Write-only binding to the contract
	CbridgeFilterer   // Log filterer for contract events
}

// CbridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CbridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CbridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CbridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CbridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CbridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CbridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CbridgeSession struct {
	Contract     *Cbridge          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CbridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CbridgeCallerSession struct {
	Contract *CbridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CbridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CbridgeTransactorSession struct {
	Contract     *CbridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CbridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CbridgeRaw struct {
	Contract *Cbridge // Generic contract binding to access the raw methods on
}

// CbridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CbridgeCallerRaw struct {
	Contract *CbridgeCaller // Generic read-only contract binding to access the raw methods on
}

// CbridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CbridgeTransactorRaw struct {
	Contract *CbridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCbridge creates a new instance of Cbridge, bound to a specific deployed contract.
func NewCbridge(address common.Address, backend bind.ContractBackend) (*Cbridge, error) {
	contract, err := bindCbridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cbridge{CbridgeCaller: CbridgeCaller{contract: contract}, CbridgeTransactor: CbridgeTransactor{contract: contract}, CbridgeFilterer: CbridgeFilterer{contract: contract}}, nil
}

// NewCbridgeCaller creates a new read-only instance of Cbridge, bound to a specific deployed contract.
func NewCbridgeCaller(address common.Address, caller bind.ContractCaller) (*CbridgeCaller, error) {
	contract, err := bindCbridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CbridgeCaller{contract: contract}, nil
}

// NewCbridgeTransactor creates a new write-only instance of Cbridge, bound to a specific deployed contract.
func NewCbridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*CbridgeTransactor, error) {
	contract, err := bindCbridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CbridgeTransactor{contract: contract}, nil
}

// NewCbridgeFilterer creates a new log filterer instance of Cbridge, bound to a specific deployed contract.
func NewCbridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*CbridgeFilterer, error) {
	contract, err := bindCbridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CbridgeFilterer{contract: contract}, nil
}

// bindCbridge binds a generic wrapper to an already deployed contract.
func bindCbridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CbridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cbridge *CbridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cbridge.Contract.CbridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cbridge *CbridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.Contract.CbridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cbridge *CbridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cbridge.Contract.CbridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cbridge *CbridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cbridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cbridge *CbridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cbridge *CbridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cbridge.Contract.contract.Transact(opts, method, params...)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Cbridge *CbridgeCaller) Addseq(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "addseq")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Cbridge *CbridgeSession) Addseq() (uint64, error) {
	return _Cbridge.Contract.Addseq(&_Cbridge.CallOpts)
}

// Addseq is a free data retrieval call binding the contract method 0x89e39127.
//
// Solidity: function addseq() view returns(uint64)
func (_Cbridge *CbridgeCallerSession) Addseq() (uint64, error) {
	return _Cbridge.Contract.Addseq(&_Cbridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Cbridge *CbridgeCaller) DelayPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "delayPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Cbridge *CbridgeSession) DelayPeriod() (*big.Int, error) {
	return _Cbridge.Contract.DelayPeriod(&_Cbridge.CallOpts)
}

// DelayPeriod is a free data retrieval call binding the contract method 0xb1c94d94.
//
// Solidity: function delayPeriod() view returns(uint256)
func (_Cbridge *CbridgeCallerSession) DelayPeriod() (*big.Int, error) {
	return _Cbridge.Contract.DelayPeriod(&_Cbridge.CallOpts)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) DelayThresholds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "delayThresholds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.DelayThresholds(&_Cbridge.CallOpts, arg0)
}

// DelayThresholds is a free data retrieval call binding the contract method 0x52532faa.
//
// Solidity: function delayThresholds(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) DelayThresholds(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.DelayThresholds(&_Cbridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Cbridge *CbridgeCaller) DelayedTransfers(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "delayedTransfers", arg0)

	outstruct := new(struct {
		Receiver  common.Address
		Token     common.Address
		Amount    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Token = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Cbridge *CbridgeSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Cbridge.Contract.DelayedTransfers(&_Cbridge.CallOpts, arg0)
}

// DelayedTransfers is a free data retrieval call binding the contract method 0xadc0d57f.
//
// Solidity: function delayedTransfers(bytes32 ) view returns(address receiver, address token, uint256 amount, uint256 timestamp)
func (_Cbridge *CbridgeCallerSession) DelayedTransfers(arg0 [32]byte) (struct {
	Receiver  common.Address
	Token     common.Address
	Amount    *big.Int
	Timestamp *big.Int
}, error) {
	return _Cbridge.Contract.DelayedTransfers(&_Cbridge.CallOpts, arg0)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Cbridge *CbridgeCaller) EpochLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "epochLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Cbridge *CbridgeSession) EpochLength() (*big.Int, error) {
	return _Cbridge.Contract.EpochLength(&_Cbridge.CallOpts)
}

// EpochLength is a free data retrieval call binding the contract method 0x57d775f8.
//
// Solidity: function epochLength() view returns(uint256)
func (_Cbridge *CbridgeCallerSession) EpochLength() (*big.Int, error) {
	return _Cbridge.Contract.EpochLength(&_Cbridge.CallOpts)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) EpochVolumeCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "epochVolumeCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.EpochVolumeCaps(&_Cbridge.CallOpts, arg0)
}

// EpochVolumeCaps is a free data retrieval call binding the contract method 0xb5f2bc47.
//
// Solidity: function epochVolumeCaps(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) EpochVolumeCaps(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.EpochVolumeCaps(&_Cbridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) EpochVolumes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "epochVolumes", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.EpochVolumes(&_Cbridge.CallOpts, arg0)
}

// EpochVolumes is a free data retrieval call binding the contract method 0x60216b00.
//
// Solidity: function epochVolumes(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) EpochVolumes(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.EpochVolumes(&_Cbridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Cbridge *CbridgeCaller) Governors(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "governors", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Cbridge *CbridgeSession) Governors(arg0 common.Address) (bool, error) {
	return _Cbridge.Contract.Governors(&_Cbridge.CallOpts, arg0)
}

// Governors is a free data retrieval call binding the contract method 0xe3eece26.
//
// Solidity: function governors(address ) view returns(bool)
func (_Cbridge *CbridgeCallerSession) Governors(arg0 common.Address) (bool, error) {
	return _Cbridge.Contract.Governors(&_Cbridge.CallOpts, arg0)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Cbridge *CbridgeCaller) IsGovernor(opts *bind.CallOpts, _account common.Address) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "isGovernor", _account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Cbridge *CbridgeSession) IsGovernor(_account common.Address) (bool, error) {
	return _Cbridge.Contract.IsGovernor(&_Cbridge.CallOpts, _account)
}

// IsGovernor is a free data retrieval call binding the contract method 0xe43581b8.
//
// Solidity: function isGovernor(address _account) view returns(bool)
func (_Cbridge *CbridgeCallerSession) IsGovernor(_account common.Address) (bool, error) {
	return _Cbridge.Contract.IsGovernor(&_Cbridge.CallOpts, _account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Cbridge *CbridgeCaller) IsPauser(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "isPauser", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Cbridge *CbridgeSession) IsPauser(account common.Address) (bool, error) {
	return _Cbridge.Contract.IsPauser(&_Cbridge.CallOpts, account)
}

// IsPauser is a free data retrieval call binding the contract method 0x46fbf68e.
//
// Solidity: function isPauser(address account) view returns(bool)
func (_Cbridge *CbridgeCallerSession) IsPauser(account common.Address) (bool, error) {
	return _Cbridge.Contract.IsPauser(&_Cbridge.CallOpts, account)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) LastOpTimestamps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "lastOpTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.LastOpTimestamps(&_Cbridge.CallOpts, arg0)
}

// LastOpTimestamps is a free data retrieval call binding the contract method 0xf8321383.
//
// Solidity: function lastOpTimestamps(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) LastOpTimestamps(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.LastOpTimestamps(&_Cbridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) MaxSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "maxSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MaxSend(&_Cbridge.CallOpts, arg0)
}

// MaxSend is a free data retrieval call binding the contract method 0x618ee055.
//
// Solidity: function maxSend(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) MaxSend(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MaxSend(&_Cbridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) MinAdd(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "minAdd", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MinAdd(&_Cbridge.CallOpts, arg0)
}

// MinAdd is a free data retrieval call binding the contract method 0xccde517a.
//
// Solidity: function minAdd(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) MinAdd(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MinAdd(&_Cbridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Cbridge *CbridgeCaller) MinSend(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "minSend", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Cbridge *CbridgeSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MinSend(&_Cbridge.CallOpts, arg0)
}

// MinSend is a free data retrieval call binding the contract method 0xf8b30d7d.
//
// Solidity: function minSend(address ) view returns(uint256)
func (_Cbridge *CbridgeCallerSession) MinSend(arg0 common.Address) (*big.Int, error) {
	return _Cbridge.Contract.MinSend(&_Cbridge.CallOpts, arg0)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Cbridge *CbridgeCaller) MinimalMaxSlippage(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "minimalMaxSlippage")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Cbridge *CbridgeSession) MinimalMaxSlippage() (uint32, error) {
	return _Cbridge.Contract.MinimalMaxSlippage(&_Cbridge.CallOpts)
}

// MinimalMaxSlippage is a free data retrieval call binding the contract method 0x2fd1b0a4.
//
// Solidity: function minimalMaxSlippage() view returns(uint32)
func (_Cbridge *CbridgeCallerSession) MinimalMaxSlippage() (uint32, error) {
	return _Cbridge.Contract.MinimalMaxSlippage(&_Cbridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Cbridge *CbridgeCaller) NativeWrap(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "nativeWrap")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Cbridge *CbridgeSession) NativeWrap() (common.Address, error) {
	return _Cbridge.Contract.NativeWrap(&_Cbridge.CallOpts)
}

// NativeWrap is a free data retrieval call binding the contract method 0x457bfa2f.
//
// Solidity: function nativeWrap() view returns(address)
func (_Cbridge *CbridgeCallerSession) NativeWrap() (common.Address, error) {
	return _Cbridge.Contract.NativeWrap(&_Cbridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Cbridge *CbridgeCaller) NoticePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "noticePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Cbridge *CbridgeSession) NoticePeriod() (*big.Int, error) {
	return _Cbridge.Contract.NoticePeriod(&_Cbridge.CallOpts)
}

// NoticePeriod is a free data retrieval call binding the contract method 0x9b14d4c6.
//
// Solidity: function noticePeriod() view returns(uint256)
func (_Cbridge *CbridgeCallerSession) NoticePeriod() (*big.Int, error) {
	return _Cbridge.Contract.NoticePeriod(&_Cbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cbridge *CbridgeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cbridge *CbridgeSession) Owner() (common.Address, error) {
	return _Cbridge.Contract.Owner(&_Cbridge.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cbridge *CbridgeCallerSession) Owner() (common.Address, error) {
	return _Cbridge.Contract.Owner(&_Cbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cbridge *CbridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cbridge *CbridgeSession) Paused() (bool, error) {
	return _Cbridge.Contract.Paused(&_Cbridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Cbridge *CbridgeCallerSession) Paused() (bool, error) {
	return _Cbridge.Contract.Paused(&_Cbridge.CallOpts)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Cbridge *CbridgeCaller) Pausers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "pausers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Cbridge *CbridgeSession) Pausers(arg0 common.Address) (bool, error) {
	return _Cbridge.Contract.Pausers(&_Cbridge.CallOpts, arg0)
}

// Pausers is a free data retrieval call binding the contract method 0x80f51c12.
//
// Solidity: function pausers(address ) view returns(bool)
func (_Cbridge *CbridgeCallerSession) Pausers(arg0 common.Address) (bool, error) {
	return _Cbridge.Contract.Pausers(&_Cbridge.CallOpts, arg0)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Cbridge *CbridgeCaller) ResetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "resetTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Cbridge *CbridgeSession) ResetTime() (*big.Int, error) {
	return _Cbridge.Contract.ResetTime(&_Cbridge.CallOpts)
}

// ResetTime is a free data retrieval call binding the contract method 0x65a114f1.
//
// Solidity: function resetTime() view returns(uint256)
func (_Cbridge *CbridgeCallerSession) ResetTime() (*big.Int, error) {
	return _Cbridge.Contract.ResetTime(&_Cbridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Cbridge *CbridgeCaller) SsHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "ssHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Cbridge *CbridgeSession) SsHash() ([32]byte, error) {
	return _Cbridge.Contract.SsHash(&_Cbridge.CallOpts)
}

// SsHash is a free data retrieval call binding the contract method 0xd0790da9.
//
// Solidity: function ssHash() view returns(bytes32)
func (_Cbridge *CbridgeCallerSession) SsHash() ([32]byte, error) {
	return _Cbridge.Contract.SsHash(&_Cbridge.CallOpts)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeCaller) Transfers(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "transfers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Cbridge.Contract.Transfers(&_Cbridge.CallOpts, arg0)
}

// Transfers is a free data retrieval call binding the contract method 0x3c64f04b.
//
// Solidity: function transfers(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeCallerSession) Transfers(arg0 [32]byte) (bool, error) {
	return _Cbridge.Contract.Transfers(&_Cbridge.CallOpts, arg0)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Cbridge *CbridgeCaller) TriggerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "triggerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Cbridge *CbridgeSession) TriggerTime() (*big.Int, error) {
	return _Cbridge.Contract.TriggerTime(&_Cbridge.CallOpts)
}

// TriggerTime is a free data retrieval call binding the contract method 0x370fb47b.
//
// Solidity: function triggerTime() view returns(uint256)
func (_Cbridge *CbridgeCallerSession) TriggerTime() (*big.Int, error) {
	return _Cbridge.Contract.TriggerTime(&_Cbridge.CallOpts)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Cbridge *CbridgeCaller) VerifySigs(opts *bind.CallOpts, _msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "verifySigs", _msg, _sigs, _signers, _powers)

	if err != nil {
		return err
	}

	return err

}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Cbridge *CbridgeSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Cbridge.Contract.VerifySigs(&_Cbridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// VerifySigs is a free data retrieval call binding the contract method 0x682dbc22.
//
// Solidity: function verifySigs(bytes _msg, bytes[] _sigs, address[] _signers, uint256[] _powers) view returns()
func (_Cbridge *CbridgeCallerSession) VerifySigs(_msg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) error {
	return _Cbridge.Contract.VerifySigs(&_Cbridge.CallOpts, _msg, _sigs, _signers, _powers)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeCaller) Withdraws(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Cbridge.contract.Call(opts, &out, "withdraws", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Cbridge.Contract.Withdraws(&_Cbridge.CallOpts, arg0)
}

// Withdraws is a free data retrieval call binding the contract method 0xe09ab428.
//
// Solidity: function withdraws(bytes32 ) view returns(bool)
func (_Cbridge *CbridgeCallerSession) Withdraws(arg0 [32]byte) (bool, error) {
	return _Cbridge.Contract.Withdraws(&_Cbridge.CallOpts, arg0)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Cbridge *CbridgeTransactor) AddGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "addGovernor", _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Cbridge *CbridgeSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.AddGovernor(&_Cbridge.TransactOpts, _account)
}

// AddGovernor is a paid mutator transaction binding the contract method 0x3c4a25d0.
//
// Solidity: function addGovernor(address _account) returns()
func (_Cbridge *CbridgeTransactorSession) AddGovernor(_account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.AddGovernor(&_Cbridge.TransactOpts, _account)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Cbridge *CbridgeTransactor) AddLiquidity(opts *bind.TransactOpts, _token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "addLiquidity", _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Cbridge *CbridgeSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.AddLiquidity(&_Cbridge.TransactOpts, _token, _amount)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x56688700.
//
// Solidity: function addLiquidity(address _token, uint256 _amount) returns()
func (_Cbridge *CbridgeTransactorSession) AddLiquidity(_token common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.AddLiquidity(&_Cbridge.TransactOpts, _token, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Cbridge *CbridgeTransactor) AddNativeLiquidity(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "addNativeLiquidity", _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Cbridge *CbridgeSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.AddNativeLiquidity(&_Cbridge.TransactOpts, _amount)
}

// AddNativeLiquidity is a paid mutator transaction binding the contract method 0x7044c89e.
//
// Solidity: function addNativeLiquidity(uint256 _amount) payable returns()
func (_Cbridge *CbridgeTransactorSession) AddNativeLiquidity(_amount *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.AddNativeLiquidity(&_Cbridge.TransactOpts, _amount)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Cbridge *CbridgeTransactor) AddPauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "addPauser", account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Cbridge *CbridgeSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.AddPauser(&_Cbridge.TransactOpts, account)
}

// AddPauser is a paid mutator transaction binding the contract method 0x82dc1ec4.
//
// Solidity: function addPauser(address account) returns()
func (_Cbridge *CbridgeTransactorSession) AddPauser(account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.AddPauser(&_Cbridge.TransactOpts, account)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Cbridge *CbridgeTransactor) ExecuteDelayedTransfer(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "executeDelayedTransfer", id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Cbridge *CbridgeSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Cbridge.Contract.ExecuteDelayedTransfer(&_Cbridge.TransactOpts, id)
}

// ExecuteDelayedTransfer is a paid mutator transaction binding the contract method 0x9e25fc5c.
//
// Solidity: function executeDelayedTransfer(bytes32 id) returns()
func (_Cbridge *CbridgeTransactorSession) ExecuteDelayedTransfer(id [32]byte) (*types.Transaction, error) {
	return _Cbridge.Contract.ExecuteDelayedTransfer(&_Cbridge.TransactOpts, id)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Cbridge *CbridgeTransactor) IncreaseNoticePeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "increaseNoticePeriod", period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Cbridge *CbridgeSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.IncreaseNoticePeriod(&_Cbridge.TransactOpts, period)
}

// IncreaseNoticePeriod is a paid mutator transaction binding the contract method 0xf20c922a.
//
// Solidity: function increaseNoticePeriod(uint256 period) returns()
func (_Cbridge *CbridgeTransactorSession) IncreaseNoticePeriod(period *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.IncreaseNoticePeriod(&_Cbridge.TransactOpts, period)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Cbridge *CbridgeTransactor) NotifyResetSigners(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "notifyResetSigners")
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Cbridge *CbridgeSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Cbridge.Contract.NotifyResetSigners(&_Cbridge.TransactOpts)
}

// NotifyResetSigners is a paid mutator transaction binding the contract method 0x25c38b9f.
//
// Solidity: function notifyResetSigners() returns()
func (_Cbridge *CbridgeTransactorSession) NotifyResetSigners() (*types.Transaction, error) {
	return _Cbridge.Contract.NotifyResetSigners(&_Cbridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cbridge *CbridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cbridge *CbridgeSession) Pause() (*types.Transaction, error) {
	return _Cbridge.Contract.Pause(&_Cbridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Cbridge *CbridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Cbridge.Contract.Pause(&_Cbridge.TransactOpts)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactor) Relay(opts *bind.TransactOpts, _relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "relay", _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.Relay(&_Cbridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// Relay is a paid mutator transaction binding the contract method 0xcdd1b25d.
//
// Solidity: function relay(bytes _relayRequest, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactorSession) Relay(_relayRequest []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.Relay(&_Cbridge.TransactOpts, _relayRequest, _sigs, _signers, _powers)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Cbridge *CbridgeTransactor) RemoveGovernor(opts *bind.TransactOpts, _account common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "removeGovernor", _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Cbridge *CbridgeSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.RemoveGovernor(&_Cbridge.TransactOpts, _account)
}

// RemoveGovernor is a paid mutator transaction binding the contract method 0xeecdac88.
//
// Solidity: function removeGovernor(address _account) returns()
func (_Cbridge *CbridgeTransactorSession) RemoveGovernor(_account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.RemoveGovernor(&_Cbridge.TransactOpts, _account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Cbridge *CbridgeTransactor) RemovePauser(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "removePauser", account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Cbridge *CbridgeSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.RemovePauser(&_Cbridge.TransactOpts, account)
}

// RemovePauser is a paid mutator transaction binding the contract method 0x6b2c0f55.
//
// Solidity: function removePauser(address account) returns()
func (_Cbridge *CbridgeTransactorSession) RemovePauser(account common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.RemovePauser(&_Cbridge.TransactOpts, account)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Cbridge *CbridgeTransactor) RenounceGovernor(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "renounceGovernor")
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Cbridge *CbridgeSession) RenounceGovernor() (*types.Transaction, error) {
	return _Cbridge.Contract.RenounceGovernor(&_Cbridge.TransactOpts)
}

// RenounceGovernor is a paid mutator transaction binding the contract method 0xe026049c.
//
// Solidity: function renounceGovernor() returns()
func (_Cbridge *CbridgeTransactorSession) RenounceGovernor() (*types.Transaction, error) {
	return _Cbridge.Contract.RenounceGovernor(&_Cbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cbridge *CbridgeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cbridge *CbridgeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cbridge.Contract.RenounceOwnership(&_Cbridge.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cbridge *CbridgeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cbridge.Contract.RenounceOwnership(&_Cbridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Cbridge *CbridgeTransactor) RenouncePauser(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "renouncePauser")
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Cbridge *CbridgeSession) RenouncePauser() (*types.Transaction, error) {
	return _Cbridge.Contract.RenouncePauser(&_Cbridge.TransactOpts)
}

// RenouncePauser is a paid mutator transaction binding the contract method 0x6ef8d66d.
//
// Solidity: function renouncePauser() returns()
func (_Cbridge *CbridgeTransactorSession) RenouncePauser() (*types.Transaction, error) {
	return _Cbridge.Contract.RenouncePauser(&_Cbridge.TransactOpts)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactor) ResetSigners(opts *bind.TransactOpts, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "resetSigners", _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.ResetSigners(&_Cbridge.TransactOpts, _signers, _powers)
}

// ResetSigners is a paid mutator transaction binding the contract method 0xa7bdf45a.
//
// Solidity: function resetSigners(address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactorSession) ResetSigners(_signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.ResetSigners(&_Cbridge.TransactOpts, _signers, _powers)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Cbridge *CbridgeTransactor) Send(opts *bind.TransactOpts, _receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "send", _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Cbridge *CbridgeSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.Send(&_Cbridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// Send is a paid mutator transaction binding the contract method 0xa5977fbb.
//
// Solidity: function send(address _receiver, address _token, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) returns()
func (_Cbridge *CbridgeTransactorSession) Send(_receiver common.Address, _token common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.Send(&_Cbridge.TransactOpts, _receiver, _token, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Cbridge *CbridgeTransactor) SendNative(opts *bind.TransactOpts, _receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "sendNative", _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Cbridge *CbridgeSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.SendNative(&_Cbridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SendNative is a paid mutator transaction binding the contract method 0x3f2e5fc3.
//
// Solidity: function sendNative(address _receiver, uint256 _amount, uint64 _dstChainId, uint64 _nonce, uint32 _maxSlippage) payable returns()
func (_Cbridge *CbridgeTransactorSession) SendNative(_receiver common.Address, _amount *big.Int, _dstChainId uint64, _nonce uint64, _maxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.SendNative(&_Cbridge.TransactOpts, _receiver, _amount, _dstChainId, _nonce, _maxSlippage)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Cbridge *CbridgeTransactor) SetDelayPeriod(opts *bind.TransactOpts, _period *big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setDelayPeriod", _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Cbridge *CbridgeSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetDelayPeriod(&_Cbridge.TransactOpts, _period)
}

// SetDelayPeriod is a paid mutator transaction binding the contract method 0x3d572107.
//
// Solidity: function setDelayPeriod(uint256 _period) returns()
func (_Cbridge *CbridgeTransactorSession) SetDelayPeriod(_period *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetDelayPeriod(&_Cbridge.TransactOpts, _period)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Cbridge *CbridgeTransactor) SetDelayThresholds(opts *bind.TransactOpts, _tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setDelayThresholds", _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Cbridge *CbridgeSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetDelayThresholds(&_Cbridge.TransactOpts, _tokens, _thresholds)
}

// SetDelayThresholds is a paid mutator transaction binding the contract method 0x17bdbae5.
//
// Solidity: function setDelayThresholds(address[] _tokens, uint256[] _thresholds) returns()
func (_Cbridge *CbridgeTransactorSession) SetDelayThresholds(_tokens []common.Address, _thresholds []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetDelayThresholds(&_Cbridge.TransactOpts, _tokens, _thresholds)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Cbridge *CbridgeTransactor) SetEpochLength(opts *bind.TransactOpts, _length *big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setEpochLength", _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Cbridge *CbridgeSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetEpochLength(&_Cbridge.TransactOpts, _length)
}

// SetEpochLength is a paid mutator transaction binding the contract method 0x54eea796.
//
// Solidity: function setEpochLength(uint256 _length) returns()
func (_Cbridge *CbridgeTransactorSession) SetEpochLength(_length *big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetEpochLength(&_Cbridge.TransactOpts, _length)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Cbridge *CbridgeTransactor) SetEpochVolumeCaps(opts *bind.TransactOpts, _tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setEpochVolumeCaps", _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Cbridge *CbridgeSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetEpochVolumeCaps(&_Cbridge.TransactOpts, _tokens, _caps)
}

// SetEpochVolumeCaps is a paid mutator transaction binding the contract method 0x47b16c6c.
//
// Solidity: function setEpochVolumeCaps(address[] _tokens, uint256[] _caps) returns()
func (_Cbridge *CbridgeTransactorSession) SetEpochVolumeCaps(_tokens []common.Address, _caps []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetEpochVolumeCaps(&_Cbridge.TransactOpts, _tokens, _caps)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactor) SetMaxSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setMaxSend", _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMaxSend(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMaxSend is a paid mutator transaction binding the contract method 0x878fe1ce.
//
// Solidity: function setMaxSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactorSession) SetMaxSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMaxSend(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactor) SetMinAdd(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setMinAdd", _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinAdd(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMinAdd is a paid mutator transaction binding the contract method 0xe999e5f4.
//
// Solidity: function setMinAdd(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactorSession) SetMinAdd(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinAdd(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactor) SetMinSend(opts *bind.TransactOpts, _tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setMinSend", _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinSend(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMinSend is a paid mutator transaction binding the contract method 0x08992741.
//
// Solidity: function setMinSend(address[] _tokens, uint256[] _amounts) returns()
func (_Cbridge *CbridgeTransactorSession) SetMinSend(_tokens []common.Address, _amounts []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinSend(&_Cbridge.TransactOpts, _tokens, _amounts)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Cbridge *CbridgeTransactor) SetMinimalMaxSlippage(opts *bind.TransactOpts, _minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setMinimalMaxSlippage", _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Cbridge *CbridgeSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinimalMaxSlippage(&_Cbridge.TransactOpts, _minimalMaxSlippage)
}

// SetMinimalMaxSlippage is a paid mutator transaction binding the contract method 0x48234126.
//
// Solidity: function setMinimalMaxSlippage(uint32 _minimalMaxSlippage) returns()
func (_Cbridge *CbridgeTransactorSession) SetMinimalMaxSlippage(_minimalMaxSlippage uint32) (*types.Transaction, error) {
	return _Cbridge.Contract.SetMinimalMaxSlippage(&_Cbridge.TransactOpts, _minimalMaxSlippage)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Cbridge *CbridgeTransactor) SetWrap(opts *bind.TransactOpts, _weth common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "setWrap", _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Cbridge *CbridgeSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.SetWrap(&_Cbridge.TransactOpts, _weth)
}

// SetWrap is a paid mutator transaction binding the contract method 0x9ff9001a.
//
// Solidity: function setWrap(address _weth) returns()
func (_Cbridge *CbridgeTransactorSession) SetWrap(_weth common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.SetWrap(&_Cbridge.TransactOpts, _weth)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cbridge *CbridgeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cbridge *CbridgeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.TransferOwnership(&_Cbridge.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cbridge *CbridgeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cbridge.Contract.TransferOwnership(&_Cbridge.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cbridge *CbridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cbridge *CbridgeSession) Unpause() (*types.Transaction, error) {
	return _Cbridge.Contract.Unpause(&_Cbridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Cbridge *CbridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Cbridge.Contract.Unpause(&_Cbridge.TransactOpts)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Cbridge *CbridgeTransactor) UpdateSigners(opts *bind.TransactOpts, _triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "updateSigners", _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Cbridge *CbridgeSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.UpdateSigners(&_Cbridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// UpdateSigners is a paid mutator transaction binding the contract method 0xba2cb25c.
//
// Solidity: function updateSigners(uint256 _triggerTime, address[] _newSigners, uint256[] _newPowers, bytes[] _sigs, address[] _curSigners, uint256[] _curPowers) returns()
func (_Cbridge *CbridgeTransactorSession) UpdateSigners(_triggerTime *big.Int, _newSigners []common.Address, _newPowers []*big.Int, _sigs [][]byte, _curSigners []common.Address, _curPowers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.UpdateSigners(&_Cbridge.TransactOpts, _triggerTime, _newSigners, _newPowers, _sigs, _curSigners, _curPowers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactor) Withdraw(opts *bind.TransactOpts, _wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.contract.Transact(opts, "withdraw", _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.Withdraw(&_Cbridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Withdraw is a paid mutator transaction binding the contract method 0xa21a9280.
//
// Solidity: function withdraw(bytes _wdmsg, bytes[] _sigs, address[] _signers, uint256[] _powers) returns()
func (_Cbridge *CbridgeTransactorSession) Withdraw(_wdmsg []byte, _sigs [][]byte, _signers []common.Address, _powers []*big.Int) (*types.Transaction, error) {
	return _Cbridge.Contract.Withdraw(&_Cbridge.TransactOpts, _wdmsg, _sigs, _signers, _powers)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cbridge *CbridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cbridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cbridge *CbridgeSession) Receive() (*types.Transaction, error) {
	return _Cbridge.Contract.Receive(&_Cbridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Cbridge *CbridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Cbridge.Contract.Receive(&_Cbridge.TransactOpts)
}

// CbridgeDelayPeriodUpdatedIterator is returned from FilterDelayPeriodUpdated and is used to iterate over the raw logs and unpacked data for DelayPeriodUpdated events raised by the Cbridge contract.
type CbridgeDelayPeriodUpdatedIterator struct {
	Event *CbridgeDelayPeriodUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeDelayPeriodUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeDelayPeriodUpdated)
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
		it.Event = new(CbridgeDelayPeriodUpdated)
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
func (it *CbridgeDelayPeriodUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeDelayPeriodUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeDelayPeriodUpdated represents a DelayPeriodUpdated event raised by the Cbridge contract.
type CbridgeDelayPeriodUpdated struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelayPeriodUpdated is a free log retrieval operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Cbridge *CbridgeFilterer) FilterDelayPeriodUpdated(opts *bind.FilterOpts) (*CbridgeDelayPeriodUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeDelayPeriodUpdatedIterator{contract: _Cbridge.contract, event: "DelayPeriodUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayPeriodUpdated is a free log subscription operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Cbridge *CbridgeFilterer) WatchDelayPeriodUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeDelayPeriodUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "DelayPeriodUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeDelayPeriodUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
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

// ParseDelayPeriodUpdated is a log parse operation binding the contract event 0xc0a39f234199b125fb93713c4d067bdcebbf691087f87b79c0feb92b156ba8b6.
//
// Solidity: event DelayPeriodUpdated(uint256 period)
func (_Cbridge *CbridgeFilterer) ParseDelayPeriodUpdated(log types.Log) (*CbridgeDelayPeriodUpdated, error) {
	event := new(CbridgeDelayPeriodUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "DelayPeriodUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeDelayThresholdUpdatedIterator is returned from FilterDelayThresholdUpdated and is used to iterate over the raw logs and unpacked data for DelayThresholdUpdated events raised by the Cbridge contract.
type CbridgeDelayThresholdUpdatedIterator struct {
	Event *CbridgeDelayThresholdUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeDelayThresholdUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeDelayThresholdUpdated)
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
		it.Event = new(CbridgeDelayThresholdUpdated)
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
func (it *CbridgeDelayThresholdUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeDelayThresholdUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeDelayThresholdUpdated represents a DelayThresholdUpdated event raised by the Cbridge contract.
type CbridgeDelayThresholdUpdated struct {
	Token     common.Address
	Threshold *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelayThresholdUpdated is a free log retrieval operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Cbridge *CbridgeFilterer) FilterDelayThresholdUpdated(opts *bind.FilterOpts) (*CbridgeDelayThresholdUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeDelayThresholdUpdatedIterator{contract: _Cbridge.contract, event: "DelayThresholdUpdated", logs: logs, sub: sub}, nil
}

// WatchDelayThresholdUpdated is a free log subscription operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Cbridge *CbridgeFilterer) WatchDelayThresholdUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeDelayThresholdUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "DelayThresholdUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeDelayThresholdUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
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

// ParseDelayThresholdUpdated is a log parse operation binding the contract event 0xceaad6533bfb481492fb3e08ef19297f46611b8fa9de5ef4cf8dc23a56ad09ce.
//
// Solidity: event DelayThresholdUpdated(address token, uint256 threshold)
func (_Cbridge *CbridgeFilterer) ParseDelayThresholdUpdated(log types.Log) (*CbridgeDelayThresholdUpdated, error) {
	event := new(CbridgeDelayThresholdUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "DelayThresholdUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeDelayedTransferAddedIterator is returned from FilterDelayedTransferAdded and is used to iterate over the raw logs and unpacked data for DelayedTransferAdded events raised by the Cbridge contract.
type CbridgeDelayedTransferAddedIterator struct {
	Event *CbridgeDelayedTransferAdded // Event containing the contract specifics and raw log

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
func (it *CbridgeDelayedTransferAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeDelayedTransferAdded)
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
		it.Event = new(CbridgeDelayedTransferAdded)
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
func (it *CbridgeDelayedTransferAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeDelayedTransferAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeDelayedTransferAdded represents a DelayedTransferAdded event raised by the Cbridge contract.
type CbridgeDelayedTransferAdded struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferAdded is a free log retrieval operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Cbridge *CbridgeFilterer) FilterDelayedTransferAdded(opts *bind.FilterOpts) (*CbridgeDelayedTransferAddedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return &CbridgeDelayedTransferAddedIterator{contract: _Cbridge.contract, event: "DelayedTransferAdded", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferAdded is a free log subscription operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Cbridge *CbridgeFilterer) WatchDelayedTransferAdded(opts *bind.WatchOpts, sink chan<- *CbridgeDelayedTransferAdded) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "DelayedTransferAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeDelayedTransferAdded)
				if err := _Cbridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
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

// ParseDelayedTransferAdded is a log parse operation binding the contract event 0xcbcfffe5102114216a85d3aceb14ad4b81a3935b1b5c468fadf3889eb9c5dce6.
//
// Solidity: event DelayedTransferAdded(bytes32 id)
func (_Cbridge *CbridgeFilterer) ParseDelayedTransferAdded(log types.Log) (*CbridgeDelayedTransferAdded, error) {
	event := new(CbridgeDelayedTransferAdded)
	if err := _Cbridge.contract.UnpackLog(event, "DelayedTransferAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeDelayedTransferExecutedIterator is returned from FilterDelayedTransferExecuted and is used to iterate over the raw logs and unpacked data for DelayedTransferExecuted events raised by the Cbridge contract.
type CbridgeDelayedTransferExecutedIterator struct {
	Event *CbridgeDelayedTransferExecuted // Event containing the contract specifics and raw log

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
func (it *CbridgeDelayedTransferExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeDelayedTransferExecuted)
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
		it.Event = new(CbridgeDelayedTransferExecuted)
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
func (it *CbridgeDelayedTransferExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeDelayedTransferExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeDelayedTransferExecuted represents a DelayedTransferExecuted event raised by the Cbridge contract.
type CbridgeDelayedTransferExecuted struct {
	Id       [32]byte
	Receiver common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDelayedTransferExecuted is a free log retrieval operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) FilterDelayedTransferExecuted(opts *bind.FilterOpts) (*CbridgeDelayedTransferExecutedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return &CbridgeDelayedTransferExecutedIterator{contract: _Cbridge.contract, event: "DelayedTransferExecuted", logs: logs, sub: sub}, nil
}

// WatchDelayedTransferExecuted is a free log subscription operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) WatchDelayedTransferExecuted(opts *bind.WatchOpts, sink chan<- *CbridgeDelayedTransferExecuted) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "DelayedTransferExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeDelayedTransferExecuted)
				if err := _Cbridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
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

// ParseDelayedTransferExecuted is a log parse operation binding the contract event 0x3b40e5089937425d14cdd96947e5661868357e224af59bd8b24a4b8a330d4426.
//
// Solidity: event DelayedTransferExecuted(bytes32 id, address receiver, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) ParseDelayedTransferExecuted(log types.Log) (*CbridgeDelayedTransferExecuted, error) {
	event := new(CbridgeDelayedTransferExecuted)
	if err := _Cbridge.contract.UnpackLog(event, "DelayedTransferExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeEpochLengthUpdatedIterator is returned from FilterEpochLengthUpdated and is used to iterate over the raw logs and unpacked data for EpochLengthUpdated events raised by the Cbridge contract.
type CbridgeEpochLengthUpdatedIterator struct {
	Event *CbridgeEpochLengthUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeEpochLengthUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeEpochLengthUpdated)
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
		it.Event = new(CbridgeEpochLengthUpdated)
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
func (it *CbridgeEpochLengthUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeEpochLengthUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeEpochLengthUpdated represents a EpochLengthUpdated event raised by the Cbridge contract.
type CbridgeEpochLengthUpdated struct {
	Length *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEpochLengthUpdated is a free log retrieval operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Cbridge *CbridgeFilterer) FilterEpochLengthUpdated(opts *bind.FilterOpts) (*CbridgeEpochLengthUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeEpochLengthUpdatedIterator{contract: _Cbridge.contract, event: "EpochLengthUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochLengthUpdated is a free log subscription operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Cbridge *CbridgeFilterer) WatchEpochLengthUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeEpochLengthUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "EpochLengthUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeEpochLengthUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
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

// ParseEpochLengthUpdated is a log parse operation binding the contract event 0x2664fec2ff76486ac58ed087310855b648b15b9d19f3de8529e95f7c46b7d6b3.
//
// Solidity: event EpochLengthUpdated(uint256 length)
func (_Cbridge *CbridgeFilterer) ParseEpochLengthUpdated(log types.Log) (*CbridgeEpochLengthUpdated, error) {
	event := new(CbridgeEpochLengthUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "EpochLengthUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeEpochVolumeUpdatedIterator is returned from FilterEpochVolumeUpdated and is used to iterate over the raw logs and unpacked data for EpochVolumeUpdated events raised by the Cbridge contract.
type CbridgeEpochVolumeUpdatedIterator struct {
	Event *CbridgeEpochVolumeUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeEpochVolumeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeEpochVolumeUpdated)
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
		it.Event = new(CbridgeEpochVolumeUpdated)
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
func (it *CbridgeEpochVolumeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeEpochVolumeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeEpochVolumeUpdated represents a EpochVolumeUpdated event raised by the Cbridge contract.
type CbridgeEpochVolumeUpdated struct {
	Token common.Address
	Cap   *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterEpochVolumeUpdated is a free log retrieval operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Cbridge *CbridgeFilterer) FilterEpochVolumeUpdated(opts *bind.FilterOpts) (*CbridgeEpochVolumeUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeEpochVolumeUpdatedIterator{contract: _Cbridge.contract, event: "EpochVolumeUpdated", logs: logs, sub: sub}, nil
}

// WatchEpochVolumeUpdated is a free log subscription operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Cbridge *CbridgeFilterer) WatchEpochVolumeUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeEpochVolumeUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "EpochVolumeUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeEpochVolumeUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
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

// ParseEpochVolumeUpdated is a log parse operation binding the contract event 0x608e49c22994f20b5d3496dca088b88dfd81b4a3e8cc3809ea1e10a320107e89.
//
// Solidity: event EpochVolumeUpdated(address token, uint256 cap)
func (_Cbridge *CbridgeFilterer) ParseEpochVolumeUpdated(log types.Log) (*CbridgeEpochVolumeUpdated, error) {
	event := new(CbridgeEpochVolumeUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "EpochVolumeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeGovernorAddedIterator is returned from FilterGovernorAdded and is used to iterate over the raw logs and unpacked data for GovernorAdded events raised by the Cbridge contract.
type CbridgeGovernorAddedIterator struct {
	Event *CbridgeGovernorAdded // Event containing the contract specifics and raw log

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
func (it *CbridgeGovernorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeGovernorAdded)
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
		it.Event = new(CbridgeGovernorAdded)
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
func (it *CbridgeGovernorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeGovernorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeGovernorAdded represents a GovernorAdded event raised by the Cbridge contract.
type CbridgeGovernorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorAdded is a free log retrieval operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Cbridge *CbridgeFilterer) FilterGovernorAdded(opts *bind.FilterOpts) (*CbridgeGovernorAddedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return &CbridgeGovernorAddedIterator{contract: _Cbridge.contract, event: "GovernorAdded", logs: logs, sub: sub}, nil
}

// WatchGovernorAdded is a free log subscription operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Cbridge *CbridgeFilterer) WatchGovernorAdded(opts *bind.WatchOpts, sink chan<- *CbridgeGovernorAdded) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "GovernorAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeGovernorAdded)
				if err := _Cbridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
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

// ParseGovernorAdded is a log parse operation binding the contract event 0xdc5a48d79e2e147530ff63ecdbed5a5a66adb9d5cf339384d5d076da197c40b5.
//
// Solidity: event GovernorAdded(address account)
func (_Cbridge *CbridgeFilterer) ParseGovernorAdded(log types.Log) (*CbridgeGovernorAdded, error) {
	event := new(CbridgeGovernorAdded)
	if err := _Cbridge.contract.UnpackLog(event, "GovernorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeGovernorRemovedIterator is returned from FilterGovernorRemoved and is used to iterate over the raw logs and unpacked data for GovernorRemoved events raised by the Cbridge contract.
type CbridgeGovernorRemovedIterator struct {
	Event *CbridgeGovernorRemoved // Event containing the contract specifics and raw log

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
func (it *CbridgeGovernorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeGovernorRemoved)
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
		it.Event = new(CbridgeGovernorRemoved)
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
func (it *CbridgeGovernorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeGovernorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeGovernorRemoved represents a GovernorRemoved event raised by the Cbridge contract.
type CbridgeGovernorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGovernorRemoved is a free log retrieval operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Cbridge *CbridgeFilterer) FilterGovernorRemoved(opts *bind.FilterOpts) (*CbridgeGovernorRemovedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return &CbridgeGovernorRemovedIterator{contract: _Cbridge.contract, event: "GovernorRemoved", logs: logs, sub: sub}, nil
}

// WatchGovernorRemoved is a free log subscription operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Cbridge *CbridgeFilterer) WatchGovernorRemoved(opts *bind.WatchOpts, sink chan<- *CbridgeGovernorRemoved) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "GovernorRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeGovernorRemoved)
				if err := _Cbridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
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

// ParseGovernorRemoved is a log parse operation binding the contract event 0x1ebe834e73d60a5fec822c1e1727d34bc79f2ad977ed504581cc1822fe20fb5b.
//
// Solidity: event GovernorRemoved(address account)
func (_Cbridge *CbridgeFilterer) ParseGovernorRemoved(log types.Log) (*CbridgeGovernorRemoved, error) {
	event := new(CbridgeGovernorRemoved)
	if err := _Cbridge.contract.UnpackLog(event, "GovernorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Cbridge contract.
type CbridgeLiquidityAddedIterator struct {
	Event *CbridgeLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *CbridgeLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeLiquidityAdded)
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
		it.Event = new(CbridgeLiquidityAdded)
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
func (it *CbridgeLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeLiquidityAdded represents a LiquidityAdded event raised by the Cbridge contract.
type CbridgeLiquidityAdded struct {
	Seqnum   uint64
	Provider common.Address
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) FilterLiquidityAdded(opts *bind.FilterOpts) (*CbridgeLiquidityAddedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return &CbridgeLiquidityAddedIterator{contract: _Cbridge.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *CbridgeLiquidityAdded) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "LiquidityAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeLiquidityAdded)
				if err := _Cbridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xd5d28426c3248963b1719df49aa4c665120372e02c8249bbea03d019c39ce764.
//
// Solidity: event LiquidityAdded(uint64 seqnum, address provider, address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) ParseLiquidityAdded(log types.Log) (*CbridgeLiquidityAdded, error) {
	event := new(CbridgeLiquidityAdded)
	if err := _Cbridge.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeMaxSendUpdatedIterator is returned from FilterMaxSendUpdated and is used to iterate over the raw logs and unpacked data for MaxSendUpdated events raised by the Cbridge contract.
type CbridgeMaxSendUpdatedIterator struct {
	Event *CbridgeMaxSendUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeMaxSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeMaxSendUpdated)
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
		it.Event = new(CbridgeMaxSendUpdated)
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
func (it *CbridgeMaxSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeMaxSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeMaxSendUpdated represents a MaxSendUpdated event raised by the Cbridge contract.
type CbridgeMaxSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMaxSendUpdated is a free log retrieval operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) FilterMaxSendUpdated(opts *bind.FilterOpts) (*CbridgeMaxSendUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeMaxSendUpdatedIterator{contract: _Cbridge.contract, event: "MaxSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMaxSendUpdated is a free log subscription operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) WatchMaxSendUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeMaxSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "MaxSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeMaxSendUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
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

// ParseMaxSendUpdated is a log parse operation binding the contract event 0x4f12d1a5bfb3ccd3719255d4d299d808d50cdca9a0a5c2b3a5aaa7edde73052c.
//
// Solidity: event MaxSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) ParseMaxSendUpdated(log types.Log) (*CbridgeMaxSendUpdated, error) {
	event := new(CbridgeMaxSendUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "MaxSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeMinAddUpdatedIterator is returned from FilterMinAddUpdated and is used to iterate over the raw logs and unpacked data for MinAddUpdated events raised by the Cbridge contract.
type CbridgeMinAddUpdatedIterator struct {
	Event *CbridgeMinAddUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeMinAddUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeMinAddUpdated)
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
		it.Event = new(CbridgeMinAddUpdated)
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
func (it *CbridgeMinAddUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeMinAddUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeMinAddUpdated represents a MinAddUpdated event raised by the Cbridge contract.
type CbridgeMinAddUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinAddUpdated is a free log retrieval operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) FilterMinAddUpdated(opts *bind.FilterOpts) (*CbridgeMinAddUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeMinAddUpdatedIterator{contract: _Cbridge.contract, event: "MinAddUpdated", logs: logs, sub: sub}, nil
}

// WatchMinAddUpdated is a free log subscription operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) WatchMinAddUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeMinAddUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "MinAddUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeMinAddUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
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

// ParseMinAddUpdated is a log parse operation binding the contract event 0xc56b0d14c4940515800d94ebbd0f3f5d8cc58ba1109c12536bd993b72e466e4f.
//
// Solidity: event MinAddUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) ParseMinAddUpdated(log types.Log) (*CbridgeMinAddUpdated, error) {
	event := new(CbridgeMinAddUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "MinAddUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeMinSendUpdatedIterator is returned from FilterMinSendUpdated and is used to iterate over the raw logs and unpacked data for MinSendUpdated events raised by the Cbridge contract.
type CbridgeMinSendUpdatedIterator struct {
	Event *CbridgeMinSendUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeMinSendUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeMinSendUpdated)
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
		it.Event = new(CbridgeMinSendUpdated)
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
func (it *CbridgeMinSendUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeMinSendUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeMinSendUpdated represents a MinSendUpdated event raised by the Cbridge contract.
type CbridgeMinSendUpdated struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinSendUpdated is a free log retrieval operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) FilterMinSendUpdated(opts *bind.FilterOpts) (*CbridgeMinSendUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeMinSendUpdatedIterator{contract: _Cbridge.contract, event: "MinSendUpdated", logs: logs, sub: sub}, nil
}

// WatchMinSendUpdated is a free log subscription operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) WatchMinSendUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeMinSendUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "MinSendUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeMinSendUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
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

// ParseMinSendUpdated is a log parse operation binding the contract event 0x8b59d386e660418a48d742213ad5ce7c4dd51ae81f30e4e2c387f17d907010c9.
//
// Solidity: event MinSendUpdated(address token, uint256 amount)
func (_Cbridge *CbridgeFilterer) ParseMinSendUpdated(log types.Log) (*CbridgeMinSendUpdated, error) {
	event := new(CbridgeMinSendUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "MinSendUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cbridge contract.
type CbridgeOwnershipTransferredIterator struct {
	Event *CbridgeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CbridgeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeOwnershipTransferred)
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
		it.Event = new(CbridgeOwnershipTransferred)
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
func (it *CbridgeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeOwnershipTransferred represents a OwnershipTransferred event raised by the Cbridge contract.
type CbridgeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cbridge *CbridgeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CbridgeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CbridgeOwnershipTransferredIterator{contract: _Cbridge.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cbridge *CbridgeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CbridgeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeOwnershipTransferred)
				if err := _Cbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Cbridge *CbridgeFilterer) ParseOwnershipTransferred(log types.Log) (*CbridgeOwnershipTransferred, error) {
	event := new(CbridgeOwnershipTransferred)
	if err := _Cbridge.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Cbridge contract.
type CbridgePausedIterator struct {
	Event *CbridgePaused // Event containing the contract specifics and raw log

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
func (it *CbridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgePaused)
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
		it.Event = new(CbridgePaused)
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
func (it *CbridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgePaused represents a Paused event raised by the Cbridge contract.
type CbridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cbridge *CbridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*CbridgePausedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CbridgePausedIterator{contract: _Cbridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cbridge *CbridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CbridgePaused) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgePaused)
				if err := _Cbridge.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Cbridge *CbridgeFilterer) ParsePaused(log types.Log) (*CbridgePaused, error) {
	event := new(CbridgePaused)
	if err := _Cbridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgePauserAddedIterator is returned from FilterPauserAdded and is used to iterate over the raw logs and unpacked data for PauserAdded events raised by the Cbridge contract.
type CbridgePauserAddedIterator struct {
	Event *CbridgePauserAdded // Event containing the contract specifics and raw log

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
func (it *CbridgePauserAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgePauserAdded)
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
		it.Event = new(CbridgePauserAdded)
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
func (it *CbridgePauserAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgePauserAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgePauserAdded represents a PauserAdded event raised by the Cbridge contract.
type CbridgePauserAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserAdded is a free log retrieval operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Cbridge *CbridgeFilterer) FilterPauserAdded(opts *bind.FilterOpts) (*CbridgePauserAddedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return &CbridgePauserAddedIterator{contract: _Cbridge.contract, event: "PauserAdded", logs: logs, sub: sub}, nil
}

// WatchPauserAdded is a free log subscription operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Cbridge *CbridgeFilterer) WatchPauserAdded(opts *bind.WatchOpts, sink chan<- *CbridgePauserAdded) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "PauserAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgePauserAdded)
				if err := _Cbridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
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

// ParsePauserAdded is a log parse operation binding the contract event 0x6719d08c1888103bea251a4ed56406bd0c3e69723c8a1686e017e7bbe159b6f8.
//
// Solidity: event PauserAdded(address account)
func (_Cbridge *CbridgeFilterer) ParsePauserAdded(log types.Log) (*CbridgePauserAdded, error) {
	event := new(CbridgePauserAdded)
	if err := _Cbridge.contract.UnpackLog(event, "PauserAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgePauserRemovedIterator is returned from FilterPauserRemoved and is used to iterate over the raw logs and unpacked data for PauserRemoved events raised by the Cbridge contract.
type CbridgePauserRemovedIterator struct {
	Event *CbridgePauserRemoved // Event containing the contract specifics and raw log

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
func (it *CbridgePauserRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgePauserRemoved)
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
		it.Event = new(CbridgePauserRemoved)
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
func (it *CbridgePauserRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgePauserRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgePauserRemoved represents a PauserRemoved event raised by the Cbridge contract.
type CbridgePauserRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPauserRemoved is a free log retrieval operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Cbridge *CbridgeFilterer) FilterPauserRemoved(opts *bind.FilterOpts) (*CbridgePauserRemovedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return &CbridgePauserRemovedIterator{contract: _Cbridge.contract, event: "PauserRemoved", logs: logs, sub: sub}, nil
}

// WatchPauserRemoved is a free log subscription operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Cbridge *CbridgeFilterer) WatchPauserRemoved(opts *bind.WatchOpts, sink chan<- *CbridgePauserRemoved) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "PauserRemoved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgePauserRemoved)
				if err := _Cbridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
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

// ParsePauserRemoved is a log parse operation binding the contract event 0xcd265ebaf09df2871cc7bd4133404a235ba12eff2041bb89d9c714a2621c7c7e.
//
// Solidity: event PauserRemoved(address account)
func (_Cbridge *CbridgeFilterer) ParsePauserRemoved(log types.Log) (*CbridgePauserRemoved, error) {
	event := new(CbridgePauserRemoved)
	if err := _Cbridge.contract.UnpackLog(event, "PauserRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeRelayIterator is returned from FilterRelay and is used to iterate over the raw logs and unpacked data for Relay events raised by the Cbridge contract.
type CbridgeRelayIterator struct {
	Event *CbridgeRelay // Event containing the contract specifics and raw log

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
func (it *CbridgeRelayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeRelay)
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
		it.Event = new(CbridgeRelay)
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
func (it *CbridgeRelayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeRelayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeRelay represents a Relay event raised by the Cbridge contract.
type CbridgeRelay struct {
	TransferId    [32]byte
	Sender        common.Address
	Receiver      common.Address
	Token         common.Address
	Amount        *big.Int
	SrcChainId    uint64
	SrcTransferId [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRelay is a free log retrieval operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Cbridge *CbridgeFilterer) FilterRelay(opts *bind.FilterOpts) (*CbridgeRelayIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return &CbridgeRelayIterator{contract: _Cbridge.contract, event: "Relay", logs: logs, sub: sub}, nil
}

// WatchRelay is a free log subscription operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Cbridge *CbridgeFilterer) WatchRelay(opts *bind.WatchOpts, sink chan<- *CbridgeRelay) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "Relay")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeRelay)
				if err := _Cbridge.contract.UnpackLog(event, "Relay", log); err != nil {
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

// ParseRelay is a log parse operation binding the contract event 0x79fa08de5149d912dce8e5e8da7a7c17ccdf23dd5d3bfe196802e6eb86347c7c.
//
// Solidity: event Relay(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 srcChainId, bytes32 srcTransferId)
func (_Cbridge *CbridgeFilterer) ParseRelay(log types.Log) (*CbridgeRelay, error) {
	event := new(CbridgeRelay)
	if err := _Cbridge.contract.UnpackLog(event, "Relay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeResetNotificationIterator is returned from FilterResetNotification and is used to iterate over the raw logs and unpacked data for ResetNotification events raised by the Cbridge contract.
type CbridgeResetNotificationIterator struct {
	Event *CbridgeResetNotification // Event containing the contract specifics and raw log

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
func (it *CbridgeResetNotificationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeResetNotification)
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
		it.Event = new(CbridgeResetNotification)
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
func (it *CbridgeResetNotificationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeResetNotificationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeResetNotification represents a ResetNotification event raised by the Cbridge contract.
type CbridgeResetNotification struct {
	ResetTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterResetNotification is a free log retrieval operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Cbridge *CbridgeFilterer) FilterResetNotification(opts *bind.FilterOpts) (*CbridgeResetNotificationIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return &CbridgeResetNotificationIterator{contract: _Cbridge.contract, event: "ResetNotification", logs: logs, sub: sub}, nil
}

// WatchResetNotification is a free log subscription operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Cbridge *CbridgeFilterer) WatchResetNotification(opts *bind.WatchOpts, sink chan<- *CbridgeResetNotification) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "ResetNotification")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeResetNotification)
				if err := _Cbridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
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

// ParseResetNotification is a log parse operation binding the contract event 0x68e825132f7d4bc837dea2d64ac9fc19912bf0224b67f9317d8f1a917f5304a1.
//
// Solidity: event ResetNotification(uint256 resetTime)
func (_Cbridge *CbridgeFilterer) ParseResetNotification(log types.Log) (*CbridgeResetNotification, error) {
	event := new(CbridgeResetNotification)
	if err := _Cbridge.contract.UnpackLog(event, "ResetNotification", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeSendIterator is returned from FilterSend and is used to iterate over the raw logs and unpacked data for Send events raised by the Cbridge contract.
type CbridgeSendIterator struct {
	Event *CbridgeSend // Event containing the contract specifics and raw log

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
func (it *CbridgeSendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeSend)
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
		it.Event = new(CbridgeSend)
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
func (it *CbridgeSendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeSendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeSend represents a Send event raised by the Cbridge contract.
type CbridgeSend struct {
	TransferId  [32]byte
	Sender      common.Address
	Receiver    common.Address
	Token       common.Address
	Amount      *big.Int
	DstChainId  uint64
	Nonce       uint64
	MaxSlippage uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSend is a free log retrieval operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Cbridge *CbridgeFilterer) FilterSend(opts *bind.FilterOpts) (*CbridgeSendIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return &CbridgeSendIterator{contract: _Cbridge.contract, event: "Send", logs: logs, sub: sub}, nil
}

// WatchSend is a free log subscription operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Cbridge *CbridgeFilterer) WatchSend(opts *bind.WatchOpts, sink chan<- *CbridgeSend) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "Send")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeSend)
				if err := _Cbridge.contract.UnpackLog(event, "Send", log); err != nil {
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

// ParseSend is a log parse operation binding the contract event 0x89d8051e597ab4178a863a5190407b98abfeff406aa8db90c59af76612e58f01.
//
// Solidity: event Send(bytes32 transferId, address sender, address receiver, address token, uint256 amount, uint64 dstChainId, uint64 nonce, uint32 maxSlippage)
func (_Cbridge *CbridgeFilterer) ParseSend(log types.Log) (*CbridgeSend, error) {
	event := new(CbridgeSend)
	if err := _Cbridge.contract.UnpackLog(event, "Send", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeSignersUpdatedIterator is returned from FilterSignersUpdated and is used to iterate over the raw logs and unpacked data for SignersUpdated events raised by the Cbridge contract.
type CbridgeSignersUpdatedIterator struct {
	Event *CbridgeSignersUpdated // Event containing the contract specifics and raw log

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
func (it *CbridgeSignersUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeSignersUpdated)
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
		it.Event = new(CbridgeSignersUpdated)
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
func (it *CbridgeSignersUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeSignersUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeSignersUpdated represents a SignersUpdated event raised by the Cbridge contract.
type CbridgeSignersUpdated struct {
	Signers []common.Address
	Powers  []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSignersUpdated is a free log retrieval operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Cbridge *CbridgeFilterer) FilterSignersUpdated(opts *bind.FilterOpts) (*CbridgeSignersUpdatedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return &CbridgeSignersUpdatedIterator{contract: _Cbridge.contract, event: "SignersUpdated", logs: logs, sub: sub}, nil
}

// WatchSignersUpdated is a free log subscription operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Cbridge *CbridgeFilterer) WatchSignersUpdated(opts *bind.WatchOpts, sink chan<- *CbridgeSignersUpdated) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "SignersUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeSignersUpdated)
				if err := _Cbridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
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

// ParseSignersUpdated is a log parse operation binding the contract event 0xf126123539a68393c55697f617e7d1148e371988daed246c2f41da99965a23f8.
//
// Solidity: event SignersUpdated(address[] _signers, uint256[] _powers)
func (_Cbridge *CbridgeFilterer) ParseSignersUpdated(log types.Log) (*CbridgeSignersUpdated, error) {
	event := new(CbridgeSignersUpdated)
	if err := _Cbridge.contract.UnpackLog(event, "SignersUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Cbridge contract.
type CbridgeUnpausedIterator struct {
	Event *CbridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *CbridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeUnpaused)
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
		it.Event = new(CbridgeUnpaused)
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
func (it *CbridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeUnpaused represents a Unpaused event raised by the Cbridge contract.
type CbridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cbridge *CbridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CbridgeUnpausedIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CbridgeUnpausedIterator{contract: _Cbridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cbridge *CbridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CbridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeUnpaused)
				if err := _Cbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Cbridge *CbridgeFilterer) ParseUnpaused(log types.Log) (*CbridgeUnpaused, error) {
	event := new(CbridgeUnpaused)
	if err := _Cbridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CbridgeWithdrawDoneIterator is returned from FilterWithdrawDone and is used to iterate over the raw logs and unpacked data for WithdrawDone events raised by the Cbridge contract.
type CbridgeWithdrawDoneIterator struct {
	Event *CbridgeWithdrawDone // Event containing the contract specifics and raw log

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
func (it *CbridgeWithdrawDoneIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CbridgeWithdrawDone)
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
		it.Event = new(CbridgeWithdrawDone)
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
func (it *CbridgeWithdrawDoneIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CbridgeWithdrawDoneIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CbridgeWithdrawDone represents a WithdrawDone event raised by the Cbridge contract.
type CbridgeWithdrawDone struct {
	WithdrawId [32]byte
	Seqnum     uint64
	Receiver   common.Address
	Token      common.Address
	Amount     *big.Int
	Refid      [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawDone is a free log retrieval operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Cbridge *CbridgeFilterer) FilterWithdrawDone(opts *bind.FilterOpts) (*CbridgeWithdrawDoneIterator, error) {

	logs, sub, err := _Cbridge.contract.FilterLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return &CbridgeWithdrawDoneIterator{contract: _Cbridge.contract, event: "WithdrawDone", logs: logs, sub: sub}, nil
}

// WatchWithdrawDone is a free log subscription operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Cbridge *CbridgeFilterer) WatchWithdrawDone(opts *bind.WatchOpts, sink chan<- *CbridgeWithdrawDone) (event.Subscription, error) {

	logs, sub, err := _Cbridge.contract.WatchLogs(opts, "WithdrawDone")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CbridgeWithdrawDone)
				if err := _Cbridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
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

// ParseWithdrawDone is a log parse operation binding the contract event 0x48a1ab26f3aa7b62bb6b6e8eed182f292b84eb7b006c0254386b268af20774be.
//
// Solidity: event WithdrawDone(bytes32 withdrawId, uint64 seqnum, address receiver, address token, uint256 amount, bytes32 refid)
func (_Cbridge *CbridgeFilterer) ParseWithdrawDone(log types.Log) (*CbridgeWithdrawDone, error) {
	event := new(CbridgeWithdrawDone)
	if err := _Cbridge.contract.UnpackLog(event, "WithdrawDone", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
