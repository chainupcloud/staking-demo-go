// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package depositor

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

// DepositorMetaData contains all meta data concerning the Depositor contract.
var DepositorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"mainnet\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"depositContract_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nodesAmount\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"collateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"credentialsLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"pubkeys\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"withdrawal_credentials\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32[]\",\"name\":\"deposit_data_roots\",\"type\":\"bytes32[]\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositContract\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodesMaxAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nodesMinAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pubkeyLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"signatureLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// DepositorABI is the input ABI used to generate the binding from.
// Deprecated: Use DepositorMetaData.ABI instead.
var DepositorABI = DepositorMetaData.ABI

// Depositor is an auto generated Go binding around an Ethereum contract.
type Depositor struct {
	DepositorCaller     // Read-only binding to the contract
	DepositorTransactor // Write-only binding to the contract
	DepositorFilterer   // Log filterer for contract events
}

// DepositorCaller is an auto generated read-only Go binding around an Ethereum contract.
type DepositorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DepositorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DepositorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DepositorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DepositorSession struct {
	Contract     *Depositor        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DepositorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DepositorCallerSession struct {
	Contract *DepositorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DepositorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DepositorTransactorSession struct {
	Contract     *DepositorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DepositorRaw is an auto generated low-level Go binding around an Ethereum contract.
type DepositorRaw struct {
	Contract *Depositor // Generic contract binding to access the raw methods on
}

// DepositorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DepositorCallerRaw struct {
	Contract *DepositorCaller // Generic read-only contract binding to access the raw methods on
}

// DepositorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DepositorTransactorRaw struct {
	Contract *DepositorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDepositor creates a new instance of Depositor, bound to a specific deployed contract.
func NewDepositor(address common.Address, backend bind.ContractBackend) (*Depositor, error) {
	contract, err := bindDepositor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Depositor{DepositorCaller: DepositorCaller{contract: contract}, DepositorTransactor: DepositorTransactor{contract: contract}, DepositorFilterer: DepositorFilterer{contract: contract}}, nil
}

// NewDepositorCaller creates a new read-only instance of Depositor, bound to a specific deployed contract.
func NewDepositorCaller(address common.Address, caller bind.ContractCaller) (*DepositorCaller, error) {
	contract, err := bindDepositor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DepositorCaller{contract: contract}, nil
}

// NewDepositorTransactor creates a new write-only instance of Depositor, bound to a specific deployed contract.
func NewDepositorTransactor(address common.Address, transactor bind.ContractTransactor) (*DepositorTransactor, error) {
	contract, err := bindDepositor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DepositorTransactor{contract: contract}, nil
}

// NewDepositorFilterer creates a new log filterer instance of Depositor, bound to a specific deployed contract.
func NewDepositorFilterer(address common.Address, filterer bind.ContractFilterer) (*DepositorFilterer, error) {
	contract, err := bindDepositor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DepositorFilterer{contract: contract}, nil
}

// bindDepositor binds a generic wrapper to an already deployed contract.
func bindDepositor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DepositorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositor *DepositorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositor.Contract.DepositorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositor *DepositorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.Contract.DepositorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositor *DepositorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositor.Contract.DepositorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Depositor *DepositorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Depositor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Depositor *DepositorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Depositor *DepositorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Depositor.Contract.contract.Transact(opts, method, params...)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(uint256)
func (_Depositor *DepositorCaller) Collateral(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(uint256)
func (_Depositor *DepositorSession) Collateral() (*big.Int, error) {
	return _Depositor.Contract.Collateral(&_Depositor.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(uint256)
func (_Depositor *DepositorCallerSession) Collateral() (*big.Int, error) {
	return _Depositor.Contract.Collateral(&_Depositor.CallOpts)
}

// CredentialsLength is a free data retrieval call binding the contract method 0xa5ab1d31.
//
// Solidity: function credentialsLength() view returns(uint256)
func (_Depositor *DepositorCaller) CredentialsLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "credentialsLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CredentialsLength is a free data retrieval call binding the contract method 0xa5ab1d31.
//
// Solidity: function credentialsLength() view returns(uint256)
func (_Depositor *DepositorSession) CredentialsLength() (*big.Int, error) {
	return _Depositor.Contract.CredentialsLength(&_Depositor.CallOpts)
}

// CredentialsLength is a free data retrieval call binding the contract method 0xa5ab1d31.
//
// Solidity: function credentialsLength() view returns(uint256)
func (_Depositor *DepositorCallerSession) CredentialsLength() (*big.Int, error) {
	return _Depositor.Contract.CredentialsLength(&_Depositor.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Depositor *DepositorCaller) DepositContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "depositContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Depositor *DepositorSession) DepositContract() (common.Address, error) {
	return _Depositor.Contract.DepositContract(&_Depositor.CallOpts)
}

// DepositContract is a free data retrieval call binding the contract method 0xe94ad65b.
//
// Solidity: function depositContract() view returns(address)
func (_Depositor *DepositorCallerSession) DepositContract() (common.Address, error) {
	return _Depositor.Contract.DepositContract(&_Depositor.CallOpts)
}

// NodesMaxAmount is a free data retrieval call binding the contract method 0x43ba591d.
//
// Solidity: function nodesMaxAmount() view returns(uint256)
func (_Depositor *DepositorCaller) NodesMaxAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "nodesMaxAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesMaxAmount is a free data retrieval call binding the contract method 0x43ba591d.
//
// Solidity: function nodesMaxAmount() view returns(uint256)
func (_Depositor *DepositorSession) NodesMaxAmount() (*big.Int, error) {
	return _Depositor.Contract.NodesMaxAmount(&_Depositor.CallOpts)
}

// NodesMaxAmount is a free data retrieval call binding the contract method 0x43ba591d.
//
// Solidity: function nodesMaxAmount() view returns(uint256)
func (_Depositor *DepositorCallerSession) NodesMaxAmount() (*big.Int, error) {
	return _Depositor.Contract.NodesMaxAmount(&_Depositor.CallOpts)
}

// NodesMinAmount is a free data retrieval call binding the contract method 0xf4fbdfac.
//
// Solidity: function nodesMinAmount() view returns(uint256)
func (_Depositor *DepositorCaller) NodesMinAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "nodesMinAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NodesMinAmount is a free data retrieval call binding the contract method 0xf4fbdfac.
//
// Solidity: function nodesMinAmount() view returns(uint256)
func (_Depositor *DepositorSession) NodesMinAmount() (*big.Int, error) {
	return _Depositor.Contract.NodesMinAmount(&_Depositor.CallOpts)
}

// NodesMinAmount is a free data retrieval call binding the contract method 0xf4fbdfac.
//
// Solidity: function nodesMinAmount() view returns(uint256)
func (_Depositor *DepositorCallerSession) NodesMinAmount() (*big.Int, error) {
	return _Depositor.Contract.NodesMinAmount(&_Depositor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Depositor *DepositorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Depositor *DepositorSession) Owner() (common.Address, error) {
	return _Depositor.Contract.Owner(&_Depositor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Depositor *DepositorCallerSession) Owner() (common.Address, error) {
	return _Depositor.Contract.Owner(&_Depositor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Depositor *DepositorCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Depositor *DepositorSession) Paused() (bool, error) {
	return _Depositor.Contract.Paused(&_Depositor.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Depositor *DepositorCallerSession) Paused() (bool, error) {
	return _Depositor.Contract.Paused(&_Depositor.CallOpts)
}

// PubkeyLength is a free data retrieval call binding the contract method 0x4903e8be.
//
// Solidity: function pubkeyLength() view returns(uint256)
func (_Depositor *DepositorCaller) PubkeyLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "pubkeyLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PubkeyLength is a free data retrieval call binding the contract method 0x4903e8be.
//
// Solidity: function pubkeyLength() view returns(uint256)
func (_Depositor *DepositorSession) PubkeyLength() (*big.Int, error) {
	return _Depositor.Contract.PubkeyLength(&_Depositor.CallOpts)
}

// PubkeyLength is a free data retrieval call binding the contract method 0x4903e8be.
//
// Solidity: function pubkeyLength() view returns(uint256)
func (_Depositor *DepositorCallerSession) PubkeyLength() (*big.Int, error) {
	return _Depositor.Contract.PubkeyLength(&_Depositor.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_Depositor *DepositorCaller) SignatureLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Depositor.contract.Call(opts, &out, "signatureLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_Depositor *DepositorSession) SignatureLength() (*big.Int, error) {
	return _Depositor.Contract.SignatureLength(&_Depositor.CallOpts)
}

// SignatureLength is a free data retrieval call binding the contract method 0x1b9a9323.
//
// Solidity: function signatureLength() view returns(uint256)
func (_Depositor *DepositorCallerSession) SignatureLength() (*big.Int, error) {
	return _Depositor.Contract.SignatureLength(&_Depositor.CallOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x4f498c73.
//
// Solidity: function deposit(bytes[] pubkeys, bytes[] withdrawal_credentials, bytes[] signatures, bytes32[] deposit_data_roots) payable returns()
func (_Depositor *DepositorTransactor) Deposit(opts *bind.TransactOpts, pubkeys [][]byte, withdrawal_credentials [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Depositor.contract.Transact(opts, "deposit", pubkeys, withdrawal_credentials, signatures, deposit_data_roots)
}

// Deposit is a paid mutator transaction binding the contract method 0x4f498c73.
//
// Solidity: function deposit(bytes[] pubkeys, bytes[] withdrawal_credentials, bytes[] signatures, bytes32[] deposit_data_roots) payable returns()
func (_Depositor *DepositorSession) Deposit(pubkeys [][]byte, withdrawal_credentials [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Depositor.Contract.Deposit(&_Depositor.TransactOpts, pubkeys, withdrawal_credentials, signatures, deposit_data_roots)
}

// Deposit is a paid mutator transaction binding the contract method 0x4f498c73.
//
// Solidity: function deposit(bytes[] pubkeys, bytes[] withdrawal_credentials, bytes[] signatures, bytes32[] deposit_data_roots) payable returns()
func (_Depositor *DepositorTransactorSession) Deposit(pubkeys [][]byte, withdrawal_credentials [][]byte, signatures [][]byte, deposit_data_roots [][32]byte) (*types.Transaction, error) {
	return _Depositor.Contract.Deposit(&_Depositor.TransactOpts, pubkeys, withdrawal_credentials, signatures, deposit_data_roots)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Depositor *DepositorTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Depositor *DepositorSession) Pause() (*types.Transaction, error) {
	return _Depositor.Contract.Pause(&_Depositor.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Depositor *DepositorTransactorSession) Pause() (*types.Transaction, error) {
	return _Depositor.Contract.Pause(&_Depositor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositor *DepositorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositor *DepositorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Depositor.Contract.RenounceOwnership(&_Depositor.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Depositor *DepositorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Depositor.Contract.RenounceOwnership(&_Depositor.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositor *DepositorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Depositor.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositor *DepositorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Depositor.Contract.TransferOwnership(&_Depositor.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Depositor *DepositorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Depositor.Contract.TransferOwnership(&_Depositor.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Depositor *DepositorTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Depositor *DepositorSession) Unpause() (*types.Transaction, error) {
	return _Depositor.Contract.Unpause(&_Depositor.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Depositor *DepositorTransactorSession) Unpause() (*types.Transaction, error) {
	return _Depositor.Contract.Unpause(&_Depositor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositor *DepositorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Depositor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositor *DepositorSession) Receive() (*types.Transaction, error) {
	return _Depositor.Contract.Receive(&_Depositor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Depositor *DepositorTransactorSession) Receive() (*types.Transaction, error) {
	return _Depositor.Contract.Receive(&_Depositor.TransactOpts)
}

// DepositorDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the Depositor contract.
type DepositorDepositEventIterator struct {
	Event *DepositorDepositEvent // Event containing the contract specifics and raw log

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
func (it *DepositorDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositorDepositEvent)
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
		it.Event = new(DepositorDepositEvent)
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
func (it *DepositorDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositorDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositorDepositEvent represents a DepositEvent event raised by the Depositor contract.
type DepositorDepositEvent struct {
	From        common.Address
	NodesAmount *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address from, uint256 nodesAmount)
func (_Depositor *DepositorFilterer) FilterDepositEvent(opts *bind.FilterOpts) (*DepositorDepositEventIterator, error) {

	logs, sub, err := _Depositor.contract.FilterLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return &DepositorDepositEventIterator{contract: _Depositor.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address from, uint256 nodesAmount)
func (_Depositor *DepositorFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *DepositorDepositEvent) (event.Subscription, error) {

	logs, sub, err := _Depositor.contract.WatchLogs(opts, "DepositEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositorDepositEvent)
				if err := _Depositor.contract.UnpackLog(event, "DepositEvent", log); err != nil {
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

// ParseDepositEvent is a log parse operation binding the contract event 0x2d8a08b6430a894aea608bcaa6013d5d3e263bc49110605e4d4ba76930ae5c29.
//
// Solidity: event DepositEvent(address from, uint256 nodesAmount)
func (_Depositor *DepositorFilterer) ParseDepositEvent(log types.Log) (*DepositorDepositEvent, error) {
	event := new(DepositorDepositEvent)
	if err := _Depositor.contract.UnpackLog(event, "DepositEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Depositor contract.
type DepositorOwnershipTransferredIterator struct {
	Event *DepositorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *DepositorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositorOwnershipTransferred)
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
		it.Event = new(DepositorOwnershipTransferred)
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
func (it *DepositorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositorOwnershipTransferred represents a OwnershipTransferred event raised by the Depositor contract.
type DepositorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Depositor *DepositorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*DepositorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Depositor.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &DepositorOwnershipTransferredIterator{contract: _Depositor.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Depositor *DepositorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *DepositorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Depositor.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositorOwnershipTransferred)
				if err := _Depositor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Depositor *DepositorFilterer) ParseOwnershipTransferred(log types.Log) (*DepositorOwnershipTransferred, error) {
	event := new(DepositorOwnershipTransferred)
	if err := _Depositor.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositorPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Depositor contract.
type DepositorPausedIterator struct {
	Event *DepositorPaused // Event containing the contract specifics and raw log

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
func (it *DepositorPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositorPaused)
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
		it.Event = new(DepositorPaused)
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
func (it *DepositorPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositorPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositorPaused represents a Paused event raised by the Depositor contract.
type DepositorPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Depositor *DepositorFilterer) FilterPaused(opts *bind.FilterOpts) (*DepositorPausedIterator, error) {

	logs, sub, err := _Depositor.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &DepositorPausedIterator{contract: _Depositor.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Depositor *DepositorFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *DepositorPaused) (event.Subscription, error) {

	logs, sub, err := _Depositor.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositorPaused)
				if err := _Depositor.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Depositor *DepositorFilterer) ParsePaused(log types.Log) (*DepositorPaused, error) {
	event := new(DepositorPaused)
	if err := _Depositor.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DepositorUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Depositor contract.
type DepositorUnpausedIterator struct {
	Event *DepositorUnpaused // Event containing the contract specifics and raw log

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
func (it *DepositorUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DepositorUnpaused)
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
		it.Event = new(DepositorUnpaused)
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
func (it *DepositorUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DepositorUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DepositorUnpaused represents a Unpaused event raised by the Depositor contract.
type DepositorUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Depositor *DepositorFilterer) FilterUnpaused(opts *bind.FilterOpts) (*DepositorUnpausedIterator, error) {

	logs, sub, err := _Depositor.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &DepositorUnpausedIterator{contract: _Depositor.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Depositor *DepositorFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *DepositorUnpaused) (event.Subscription, error) {

	logs, sub, err := _Depositor.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DepositorUnpaused)
				if err := _Depositor.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Depositor *DepositorFilterer) ParseUnpaused(log types.Log) (*DepositorUnpaused, error) {
	event := new(DepositorUnpaused)
	if err := _Depositor.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
