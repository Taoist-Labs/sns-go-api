// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sns

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
)

// SnsMetaData contains all meta data concerning the Sns contract.
var SnsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"addr\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"node\",\"type\":\"bytes32\"}],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SnsABI is the input ABI used to generate the binding from.
// Deprecated: Use SnsMetaData.ABI instead.
var SnsABI = SnsMetaData.ABI

// Sns is an auto generated Go binding around an Ethereum contract.
type Sns struct {
	SnsCaller     // Read-only binding to the contract
	SnsTransactor // Write-only binding to the contract
	SnsFilterer   // Log filterer for contract events
}

// SnsCaller is an auto generated read-only Go binding around an Ethereum contract.
type SnsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SnsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SnsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SnsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SnsSession struct {
	Contract     *Sns              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SnsCallerSession struct {
	Contract *SnsCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SnsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SnsTransactorSession struct {
	Contract     *SnsTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SnsRaw is an auto generated low-level Go binding around an Ethereum contract.
type SnsRaw struct {
	Contract *Sns // Generic contract binding to access the raw methods on
}

// SnsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SnsCallerRaw struct {
	Contract *SnsCaller // Generic read-only contract binding to access the raw methods on
}

// SnsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SnsTransactorRaw struct {
	Contract *SnsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSns creates a new instance of Sns, bound to a specific deployed contract.
func NewSns(address common.Address, backend bind.ContractBackend) (*Sns, error) {
	contract, err := bindSns(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Sns{SnsCaller: SnsCaller{contract: contract}, SnsTransactor: SnsTransactor{contract: contract}, SnsFilterer: SnsFilterer{contract: contract}}, nil
}

// NewSnsCaller creates a new read-only instance of Sns, bound to a specific deployed contract.
func NewSnsCaller(address common.Address, caller bind.ContractCaller) (*SnsCaller, error) {
	contract, err := bindSns(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SnsCaller{contract: contract}, nil
}

// NewSnsTransactor creates a new write-only instance of Sns, bound to a specific deployed contract.
func NewSnsTransactor(address common.Address, transactor bind.ContractTransactor) (*SnsTransactor, error) {
	contract, err := bindSns(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SnsTransactor{contract: contract}, nil
}

// NewSnsFilterer creates a new log filterer instance of Sns, bound to a specific deployed contract.
func NewSnsFilterer(address common.Address, filterer bind.ContractFilterer) (*SnsFilterer, error) {
	contract, err := bindSns(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SnsFilterer{contract: contract}, nil
}

// bindSns binds a generic wrapper to an already deployed contract.
func bindSns(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SnsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sns *SnsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sns.Contract.SnsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sns *SnsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sns.Contract.SnsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sns *SnsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sns.Contract.SnsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Sns *SnsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Sns.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Sns *SnsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Sns.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Sns *SnsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Sns.Contract.contract.Transact(opts, method, params...)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Sns *SnsCaller) Addr(opts *bind.CallOpts, node [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Sns.contract.Call(opts, &out, "addr", node)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Sns *SnsSession) Addr(node [32]byte) (common.Address, error) {
	return _Sns.Contract.Addr(&_Sns.CallOpts, node)
}

// Addr is a free data retrieval call binding the contract method 0x3b3b57de.
//
// Solidity: function addr(bytes32 node) view returns(address)
func (_Sns *SnsCallerSession) Addr(node [32]byte) (common.Address, error) {
	return _Sns.Contract.Addr(&_Sns.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Sns *SnsCaller) Name(opts *bind.CallOpts, node [32]byte) (string, error) {
	var out []interface{}
	err := _Sns.contract.Call(opts, &out, "name", node)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Sns *SnsSession) Name(node [32]byte) (string, error) {
	return _Sns.Contract.Name(&_Sns.CallOpts, node)
}

// Name is a free data retrieval call binding the contract method 0x691f3431.
//
// Solidity: function name(bytes32 node) view returns(string)
func (_Sns *SnsCallerSession) Name(node [32]byte) (string, error) {
	return _Sns.Contract.Name(&_Sns.CallOpts, node)
}
