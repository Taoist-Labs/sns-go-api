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

// BaseRegistrarMetaData contains all meta data concerning the BaseRegistrar contract.
var BaseRegistrarMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"labelToTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BaseRegistrarABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseRegistrarMetaData.ABI instead.
var BaseRegistrarABI = BaseRegistrarMetaData.ABI

// BaseRegistrar is an auto generated Go binding around an Ethereum contract.
type BaseRegistrar struct {
	BaseRegistrarCaller     // Read-only binding to the contract
	BaseRegistrarTransactor // Write-only binding to the contract
	BaseRegistrarFilterer   // Log filterer for contract events
}

// BaseRegistrarCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseRegistrarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRegistrarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseRegistrarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRegistrarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseRegistrarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseRegistrarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseRegistrarSession struct {
	Contract     *BaseRegistrar    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseRegistrarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseRegistrarCallerSession struct {
	Contract *BaseRegistrarCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BaseRegistrarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseRegistrarTransactorSession struct {
	Contract     *BaseRegistrarTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BaseRegistrarRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseRegistrarRaw struct {
	Contract *BaseRegistrar // Generic contract binding to access the raw methods on
}

// BaseRegistrarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseRegistrarCallerRaw struct {
	Contract *BaseRegistrarCaller // Generic read-only contract binding to access the raw methods on
}

// BaseRegistrarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseRegistrarTransactorRaw struct {
	Contract *BaseRegistrarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseRegistrar creates a new instance of BaseRegistrar, bound to a specific deployed contract.
func NewBaseRegistrar(address common.Address, backend bind.ContractBackend) (*BaseRegistrar, error) {
	contract, err := bindBaseRegistrar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseRegistrar{BaseRegistrarCaller: BaseRegistrarCaller{contract: contract}, BaseRegistrarTransactor: BaseRegistrarTransactor{contract: contract}, BaseRegistrarFilterer: BaseRegistrarFilterer{contract: contract}}, nil
}

// NewBaseRegistrarCaller creates a new read-only instance of BaseRegistrar, bound to a specific deployed contract.
func NewBaseRegistrarCaller(address common.Address, caller bind.ContractCaller) (*BaseRegistrarCaller, error) {
	contract, err := bindBaseRegistrar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRegistrarCaller{contract: contract}, nil
}

// NewBaseRegistrarTransactor creates a new write-only instance of BaseRegistrar, bound to a specific deployed contract.
func NewBaseRegistrarTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseRegistrarTransactor, error) {
	contract, err := bindBaseRegistrar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseRegistrarTransactor{contract: contract}, nil
}

// NewBaseRegistrarFilterer creates a new log filterer instance of BaseRegistrar, bound to a specific deployed contract.
func NewBaseRegistrarFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseRegistrarFilterer, error) {
	contract, err := bindBaseRegistrar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseRegistrarFilterer{contract: contract}, nil
}

// bindBaseRegistrar binds a generic wrapper to an already deployed contract.
func bindBaseRegistrar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseRegistrarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRegistrar *BaseRegistrarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRegistrar.Contract.BaseRegistrarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRegistrar *BaseRegistrarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRegistrar.Contract.BaseRegistrarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRegistrar *BaseRegistrarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRegistrar.Contract.BaseRegistrarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseRegistrar *BaseRegistrarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseRegistrar.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseRegistrar *BaseRegistrarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseRegistrar.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseRegistrar *BaseRegistrarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseRegistrar.Contract.contract.Transact(opts, method, params...)
}

// LabelToTokenId is a free data retrieval call binding the contract method 0xbd2c89a7.
//
// Solidity: function labelToTokenId(bytes32 ) view returns(uint256)
func (_BaseRegistrar *BaseRegistrarCaller) LabelToTokenId(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _BaseRegistrar.contract.Call(opts, &out, "labelToTokenId", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LabelToTokenId is a free data retrieval call binding the contract method 0xbd2c89a7.
//
// Solidity: function labelToTokenId(bytes32 ) view returns(uint256)
func (_BaseRegistrar *BaseRegistrarSession) LabelToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _BaseRegistrar.Contract.LabelToTokenId(&_BaseRegistrar.CallOpts, arg0)
}

// LabelToTokenId is a free data retrieval call binding the contract method 0xbd2c89a7.
//
// Solidity: function labelToTokenId(bytes32 ) view returns(uint256)
func (_BaseRegistrar *BaseRegistrarCallerSession) LabelToTokenId(arg0 [32]byte) (*big.Int, error) {
	return _BaseRegistrar.Contract.LabelToTokenId(&_BaseRegistrar.CallOpts, arg0)
}
