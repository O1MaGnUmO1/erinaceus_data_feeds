// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregator

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

// AggregatorMetaData contains all meta data concerning the Aggregator contract.
var AggregatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_link\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"_minSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"int256\",\"name\":\"_maxSubmissionValue\",\"type\":\"int256\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"current\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"}],\"name\":\"AnswerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AvailableFundsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"roundId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"startedBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"}],\"name\":\"NewRound\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdateRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"OracleAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"whitelisted\",\"type\":\"bool\"}],\"name\":\"OraclePermissionsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"requester\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"authorized\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"delay\",\"type\":\"uint32\"}],\"name\":\"RequesterPermissionsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"paymentAmount\",\"type\":\"uint128\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"minSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"maxSubmissionCount\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"restartDelay\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timeout\",\"type\":\"uint32\"}],\"name\":\"RoundDetailsUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"int256\",\"name\":\"submission\",\"type\":\"int256\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"round\",\"type\":\"uint32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"SubmissionReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"current\",\"type\":\"address\"}],\"name\":\"ValidatorUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"acceptAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allocatedFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"availableFunds\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_removed\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_added\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_addedAdmins\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"}],\"name\":\"changeOracles\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracles\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"}],\"name\":\"getTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestAnswer\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRound\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"linkToken\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minSubmissionValue\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleCount\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_queriedRoundId\",\"type\":\"uint32\"}],\"name\":\"oracleRoundState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_eligibleToSubmit\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_roundId\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"_latestSubmission\",\"type\":\"int256\"},{\"internalType\":\"uint64\",\"name\":\"_startedAt\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"_timeout\",\"type\":\"uint64\"},{\"internalType\":\"uint128\",\"name\":\"_availableFunds\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"_oracleCount\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paymentAmount\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestNewRound\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"\",\"type\":\"uint80\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"restartDelay\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_requester\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_authorized\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"_delay\",\"type\":\"uint32\"}],\"name\":\"setRequesterPermissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newValidator\",\"type\":\"address\"}],\"name\":\"setValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_roundId\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"_submission\",\"type\":\"int256\"}],\"name\":\"submit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timeout\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_newAdmin\",\"type\":\"address\"}],\"name\":\"transferAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateAvailableFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_paymentAmount\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"_minSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_maxSubmissions\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_restartDelay\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_timeout\",\"type\":\"uint32\"}],\"name\":\"updateFutureRounds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"validator\",\"outputs\":[{\"internalType\":\"contractAggregatorValidatorInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPayment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"withdrawablePayment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AggregatorABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregatorMetaData.ABI instead.
var AggregatorABI = AggregatorMetaData.ABI

// Aggregator is an auto generated Go binding around an Ethereum contract.
type Aggregator struct {
	AggregatorCaller     // Read-only binding to the contract
	AggregatorTransactor // Write-only binding to the contract
	AggregatorFilterer   // Log filterer for contract events
}

// AggregatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregatorSession struct {
	Contract     *Aggregator       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AggregatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregatorCallerSession struct {
	Contract *AggregatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AggregatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregatorTransactorSession struct {
	Contract     *AggregatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AggregatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregatorRaw struct {
	Contract *Aggregator // Generic contract binding to access the raw methods on
}

// AggregatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregatorCallerRaw struct {
	Contract *AggregatorCaller // Generic read-only contract binding to access the raw methods on
}

// AggregatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregatorTransactorRaw struct {
	Contract *AggregatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregator creates a new instance of Aggregator, bound to a specific deployed contract.
func NewAggregator(address common.Address, backend bind.ContractBackend) (*Aggregator, error) {
	contract, err := bindAggregator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Aggregator{AggregatorCaller: AggregatorCaller{contract: contract}, AggregatorTransactor: AggregatorTransactor{contract: contract}, AggregatorFilterer: AggregatorFilterer{contract: contract}}, nil
}

// NewAggregatorCaller creates a new read-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorCaller(address common.Address, caller bind.ContractCaller) (*AggregatorCaller, error) {
	contract, err := bindAggregator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorCaller{contract: contract}, nil
}

// NewAggregatorTransactor creates a new write-only instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregatorTransactor, error) {
	contract, err := bindAggregator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregatorTransactor{contract: contract}, nil
}

// NewAggregatorFilterer creates a new log filterer instance of Aggregator, bound to a specific deployed contract.
func NewAggregatorFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregatorFilterer, error) {
	contract, err := bindAggregator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregatorFilterer{contract: contract}, nil
}

// bindAggregator binds a generic wrapper to an already deployed contract.
func bindAggregator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AggregatorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.AggregatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.AggregatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Aggregator *AggregatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Aggregator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Aggregator *AggregatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Aggregator *AggregatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Aggregator.Contract.contract.Transact(opts, method, params...)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Aggregator *AggregatorCaller) AllocatedFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "allocatedFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Aggregator *AggregatorSession) AllocatedFunds() (*big.Int, error) {
	return _Aggregator.Contract.AllocatedFunds(&_Aggregator.CallOpts)
}

// AllocatedFunds is a free data retrieval call binding the contract method 0xd4cc54e4.
//
// Solidity: function allocatedFunds() view returns(uint128)
func (_Aggregator *AggregatorCallerSession) AllocatedFunds() (*big.Int, error) {
	return _Aggregator.Contract.AllocatedFunds(&_Aggregator.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Aggregator *AggregatorCaller) AvailableFunds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "availableFunds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Aggregator *AggregatorSession) AvailableFunds() (*big.Int, error) {
	return _Aggregator.Contract.AvailableFunds(&_Aggregator.CallOpts)
}

// AvailableFunds is a free data retrieval call binding the contract method 0x46fcff4c.
//
// Solidity: function availableFunds() view returns(uint128)
func (_Aggregator *AggregatorCallerSession) AvailableFunds() (*big.Int, error) {
	return _Aggregator.Contract.AvailableFunds(&_Aggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregator *AggregatorCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregator *AggregatorSession) Decimals() (uint8, error) {
	return _Aggregator.Contract.Decimals(&_Aggregator.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Aggregator *AggregatorCallerSession) Decimals() (uint8, error) {
	return _Aggregator.Contract.Decimals(&_Aggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregator *AggregatorCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregator *AggregatorSession) Description() (string, error) {
	return _Aggregator.Contract.Description(&_Aggregator.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_Aggregator *AggregatorCallerSession) Description() (string, error) {
	return _Aggregator.Contract.Description(&_Aggregator.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Aggregator *AggregatorCaller) GetAdmin(opts *bind.CallOpts, _oracle common.Address) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getAdmin", _oracle)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Aggregator *AggregatorSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _Aggregator.Contract.GetAdmin(&_Aggregator.CallOpts, _oracle)
}

// GetAdmin is a free data retrieval call binding the contract method 0x64efb22b.
//
// Solidity: function getAdmin(address _oracle) view returns(address)
func (_Aggregator *AggregatorCallerSession) GetAdmin(_oracle common.Address) (common.Address, error) {
	return _Aggregator.Contract.GetAdmin(&_Aggregator.CallOpts, _oracle)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Aggregator *AggregatorCaller) GetAnswer(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getAnswer", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Aggregator *AggregatorSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetAnswer(&_Aggregator.CallOpts, _roundId)
}

// GetAnswer is a free data retrieval call binding the contract method 0xb5ab58dc.
//
// Solidity: function getAnswer(uint256 _roundId) view returns(int256)
func (_Aggregator *AggregatorCallerSession) GetAnswer(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetAnswer(&_Aggregator.CallOpts, _roundId)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Aggregator *AggregatorCaller) GetOracles(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getOracles")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Aggregator *AggregatorSession) GetOracles() ([]common.Address, error) {
	return _Aggregator.Contract.GetOracles(&_Aggregator.CallOpts)
}

// GetOracles is a free data retrieval call binding the contract method 0x40884c52.
//
// Solidity: function getOracles() view returns(address[])
func (_Aggregator *AggregatorCallerSession) GetOracles() ([]common.Address, error) {
	return _Aggregator.Contract.GetOracles(&_Aggregator.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorCaller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregator.Contract.GetRoundData(&_Aggregator.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorCallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregator.Contract.GetRoundData(&_Aggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Aggregator *AggregatorCaller) GetTimestamp(opts *bind.CallOpts, _roundId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "getTimestamp", _roundId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Aggregator *AggregatorSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetTimestamp(&_Aggregator.CallOpts, _roundId)
}

// GetTimestamp is a free data retrieval call binding the contract method 0xb633620c.
//
// Solidity: function getTimestamp(uint256 _roundId) view returns(uint256)
func (_Aggregator *AggregatorCallerSession) GetTimestamp(_roundId *big.Int) (*big.Int, error) {
	return _Aggregator.Contract.GetTimestamp(&_Aggregator.CallOpts, _roundId)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Aggregator *AggregatorCaller) LatestAnswer(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "latestAnswer")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Aggregator *AggregatorSession) LatestAnswer() (*big.Int, error) {
	return _Aggregator.Contract.LatestAnswer(&_Aggregator.CallOpts)
}

// LatestAnswer is a free data retrieval call binding the contract method 0x50d25bcd.
//
// Solidity: function latestAnswer() view returns(int256)
func (_Aggregator *AggregatorCallerSession) LatestAnswer() (*big.Int, error) {
	return _Aggregator.Contract.LatestAnswer(&_Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Aggregator *AggregatorCaller) LatestRound(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "latestRound")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Aggregator *AggregatorSession) LatestRound() (*big.Int, error) {
	return _Aggregator.Contract.LatestRound(&_Aggregator.CallOpts)
}

// LatestRound is a free data retrieval call binding the contract method 0x668a0f02.
//
// Solidity: function latestRound() view returns(uint256)
func (_Aggregator *AggregatorCallerSession) LatestRound() (*big.Int, error) {
	return _Aggregator.Contract.LatestRound(&_Aggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregator.Contract.LatestRoundData(&_Aggregator.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_Aggregator *AggregatorCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _Aggregator.Contract.LatestRoundData(&_Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Aggregator *AggregatorCaller) LatestTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "latestTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Aggregator *AggregatorSession) LatestTimestamp() (*big.Int, error) {
	return _Aggregator.Contract.LatestTimestamp(&_Aggregator.CallOpts)
}

// LatestTimestamp is a free data retrieval call binding the contract method 0x8205bf6a.
//
// Solidity: function latestTimestamp() view returns(uint256)
func (_Aggregator *AggregatorCallerSession) LatestTimestamp() (*big.Int, error) {
	return _Aggregator.Contract.LatestTimestamp(&_Aggregator.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Aggregator *AggregatorCaller) LinkToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "linkToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Aggregator *AggregatorSession) LinkToken() (common.Address, error) {
	return _Aggregator.Contract.LinkToken(&_Aggregator.CallOpts)
}

// LinkToken is a free data retrieval call binding the contract method 0x57970e93.
//
// Solidity: function linkToken() view returns(address)
func (_Aggregator *AggregatorCallerSession) LinkToken() (common.Address, error) {
	return _Aggregator.Contract.LinkToken(&_Aggregator.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorCaller) MaxSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "maxSubmissionCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorSession) MaxSubmissionCount() (uint32, error) {
	return _Aggregator.Contract.MaxSubmissionCount(&_Aggregator.CallOpts)
}

// MaxSubmissionCount is a free data retrieval call binding the contract method 0x58609e44.
//
// Solidity: function maxSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorCallerSession) MaxSubmissionCount() (uint32, error) {
	return _Aggregator.Contract.MaxSubmissionCount(&_Aggregator.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorCaller) MaxSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "maxSubmissionValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorSession) MaxSubmissionValue() (*big.Int, error) {
	return _Aggregator.Contract.MaxSubmissionValue(&_Aggregator.CallOpts)
}

// MaxSubmissionValue is a free data retrieval call binding the contract method 0x23ca2903.
//
// Solidity: function maxSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorCallerSession) MaxSubmissionValue() (*big.Int, error) {
	return _Aggregator.Contract.MaxSubmissionValue(&_Aggregator.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorCaller) MinSubmissionCount(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "minSubmissionCount")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorSession) MinSubmissionCount() (uint32, error) {
	return _Aggregator.Contract.MinSubmissionCount(&_Aggregator.CallOpts)
}

// MinSubmissionCount is a free data retrieval call binding the contract method 0xc9374500.
//
// Solidity: function minSubmissionCount() view returns(uint32)
func (_Aggregator *AggregatorCallerSession) MinSubmissionCount() (uint32, error) {
	return _Aggregator.Contract.MinSubmissionCount(&_Aggregator.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorCaller) MinSubmissionValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "minSubmissionValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorSession) MinSubmissionValue() (*big.Int, error) {
	return _Aggregator.Contract.MinSubmissionValue(&_Aggregator.CallOpts)
}

// MinSubmissionValue is a free data retrieval call binding the contract method 0x7c2b0b21.
//
// Solidity: function minSubmissionValue() view returns(int256)
func (_Aggregator *AggregatorCallerSession) MinSubmissionValue() (*big.Int, error) {
	return _Aggregator.Contract.MinSubmissionValue(&_Aggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Aggregator *AggregatorCaller) OracleCount(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "oracleCount")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Aggregator *AggregatorSession) OracleCount() (uint8, error) {
	return _Aggregator.Contract.OracleCount(&_Aggregator.CallOpts)
}

// OracleCount is a free data retrieval call binding the contract method 0x613d8fcc.
//
// Solidity: function oracleCount() view returns(uint8)
func (_Aggregator *AggregatorCallerSession) OracleCount() (uint8, error) {
	return _Aggregator.Contract.OracleCount(&_Aggregator.CallOpts)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Aggregator *AggregatorCaller) OracleRoundState(opts *bind.CallOpts, _oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "oracleRoundState", _oracle, _queriedRoundId)

	outstruct := new(struct {
		EligibleToSubmit bool
		RoundId          uint32
		LatestSubmission *big.Int
		StartedAt        uint64
		Timeout          uint64
		AvailableFunds   *big.Int
		OracleCount      uint8
		PaymentAmount    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EligibleToSubmit = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.RoundId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.LatestSubmission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.Timeout = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.AvailableFunds = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.OracleCount = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.PaymentAmount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Aggregator *AggregatorSession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _Aggregator.Contract.OracleRoundState(&_Aggregator.CallOpts, _oracle, _queriedRoundId)
}

// OracleRoundState is a free data retrieval call binding the contract method 0x88aa80e7.
//
// Solidity: function oracleRoundState(address _oracle, uint32 _queriedRoundId) view returns(bool _eligibleToSubmit, uint32 _roundId, int256 _latestSubmission, uint64 _startedAt, uint64 _timeout, uint128 _availableFunds, uint8 _oracleCount, uint128 _paymentAmount)
func (_Aggregator *AggregatorCallerSession) OracleRoundState(_oracle common.Address, _queriedRoundId uint32) (struct {
	EligibleToSubmit bool
	RoundId          uint32
	LatestSubmission *big.Int
	StartedAt        uint64
	Timeout          uint64
	AvailableFunds   *big.Int
	OracleCount      uint8
	PaymentAmount    *big.Int
}, error) {
	return _Aggregator.Contract.OracleRoundState(&_Aggregator.CallOpts, _oracle, _queriedRoundId)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Aggregator *AggregatorCallerSession) Owner() (common.Address, error) {
	return _Aggregator.Contract.Owner(&_Aggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Aggregator *AggregatorCaller) PaymentAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "paymentAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Aggregator *AggregatorSession) PaymentAmount() (*big.Int, error) {
	return _Aggregator.Contract.PaymentAmount(&_Aggregator.CallOpts)
}

// PaymentAmount is a free data retrieval call binding the contract method 0xc35905c6.
//
// Solidity: function paymentAmount() view returns(uint128)
func (_Aggregator *AggregatorCallerSession) PaymentAmount() (*big.Int, error) {
	return _Aggregator.Contract.PaymentAmount(&_Aggregator.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Aggregator *AggregatorCaller) RestartDelay(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "restartDelay")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Aggregator *AggregatorSession) RestartDelay() (uint32, error) {
	return _Aggregator.Contract.RestartDelay(&_Aggregator.CallOpts)
}

// RestartDelay is a free data retrieval call binding the contract method 0x357ebb02.
//
// Solidity: function restartDelay() view returns(uint32)
func (_Aggregator *AggregatorCallerSession) RestartDelay() (uint32, error) {
	return _Aggregator.Contract.RestartDelay(&_Aggregator.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Aggregator *AggregatorCaller) Timeout(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "timeout")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Aggregator *AggregatorSession) Timeout() (uint32, error) {
	return _Aggregator.Contract.Timeout(&_Aggregator.CallOpts)
}

// Timeout is a free data retrieval call binding the contract method 0x70dea79a.
//
// Solidity: function timeout() view returns(uint32)
func (_Aggregator *AggregatorCallerSession) Timeout() (uint32, error) {
	return _Aggregator.Contract.Timeout(&_Aggregator.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Aggregator *AggregatorCaller) Validator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "validator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Aggregator *AggregatorSession) Validator() (common.Address, error) {
	return _Aggregator.Contract.Validator(&_Aggregator.CallOpts)
}

// Validator is a free data retrieval call binding the contract method 0x3a5381b5.
//
// Solidity: function validator() view returns(address)
func (_Aggregator *AggregatorCallerSession) Validator() (common.Address, error) {
	return _Aggregator.Contract.Validator(&_Aggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregator *AggregatorCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregator *AggregatorSession) Version() (*big.Int, error) {
	return _Aggregator.Contract.Version(&_Aggregator.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Aggregator *AggregatorCallerSession) Version() (*big.Int, error) {
	return _Aggregator.Contract.Version(&_Aggregator.CallOpts)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Aggregator *AggregatorCaller) WithdrawablePayment(opts *bind.CallOpts, _oracle common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Aggregator.contract.Call(opts, &out, "withdrawablePayment", _oracle)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Aggregator *AggregatorSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _Aggregator.Contract.WithdrawablePayment(&_Aggregator.CallOpts, _oracle)
}

// WithdrawablePayment is a free data retrieval call binding the contract method 0xe2e40317.
//
// Solidity: function withdrawablePayment(address _oracle) view returns(uint256)
func (_Aggregator *AggregatorCallerSession) WithdrawablePayment(_oracle common.Address) (*big.Int, error) {
	return _Aggregator.Contract.WithdrawablePayment(&_Aggregator.CallOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Aggregator *AggregatorTransactor) AcceptAdmin(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "acceptAdmin", _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Aggregator *AggregatorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.AcceptAdmin(&_Aggregator.TransactOpts, _oracle)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0x628806ef.
//
// Solidity: function acceptAdmin(address _oracle) returns()
func (_Aggregator *AggregatorTransactorSession) AcceptAdmin(_oracle common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.AcceptAdmin(&_Aggregator.TransactOpts, _oracle)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Aggregator *AggregatorTransactor) ChangeOracles(opts *bind.TransactOpts, _removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "changeOracles", _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Aggregator *AggregatorSession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.ChangeOracles(&_Aggregator.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// ChangeOracles is a paid mutator transaction binding the contract method 0x3969c20f.
//
// Solidity: function changeOracles(address[] _removed, address[] _added, address[] _addedAdmins, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay) returns()
func (_Aggregator *AggregatorTransactorSession) ChangeOracles(_removed []common.Address, _added []common.Address, _addedAdmins []common.Address, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.ChangeOracles(&_Aggregator.TransactOpts, _removed, _added, _addedAdmins, _minSubmissions, _maxSubmissions, _restartDelay)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Aggregator *AggregatorTransactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "onTokenTransfer", arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Aggregator *AggregatorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnTokenTransfer(&_Aggregator.TransactOpts, arg0, arg1, _data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 , bytes _data) returns()
func (_Aggregator *AggregatorTransactorSession) OnTokenTransfer(arg0 common.Address, arg1 *big.Int, _data []byte) (*types.Transaction, error) {
	return _Aggregator.Contract.OnTokenTransfer(&_Aggregator.TransactOpts, arg0, arg1, _data)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Aggregator *AggregatorTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Aggregator.Contract.RenounceOwnership(&_Aggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Aggregator *AggregatorTransactor) RequestNewRound(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "requestNewRound")
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Aggregator *AggregatorSession) RequestNewRound() (*types.Transaction, error) {
	return _Aggregator.Contract.RequestNewRound(&_Aggregator.TransactOpts)
}

// RequestNewRound is a paid mutator transaction binding the contract method 0x98e5b12a.
//
// Solidity: function requestNewRound() returns(uint80)
func (_Aggregator *AggregatorTransactorSession) RequestNewRound() (*types.Transaction, error) {
	return _Aggregator.Contract.RequestNewRound(&_Aggregator.TransactOpts)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Aggregator *AggregatorTransactor) SetRequesterPermissions(opts *bind.TransactOpts, _requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setRequesterPermissions", _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Aggregator *AggregatorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.SetRequesterPermissions(&_Aggregator.TransactOpts, _requester, _authorized, _delay)
}

// SetRequesterPermissions is a paid mutator transaction binding the contract method 0x20ed0275.
//
// Solidity: function setRequesterPermissions(address _requester, bool _authorized, uint32 _delay) returns()
func (_Aggregator *AggregatorTransactorSession) SetRequesterPermissions(_requester common.Address, _authorized bool, _delay uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.SetRequesterPermissions(&_Aggregator.TransactOpts, _requester, _authorized, _delay)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Aggregator *AggregatorTransactor) SetValidator(opts *bind.TransactOpts, _newValidator common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "setValidator", _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Aggregator *AggregatorSession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetValidator(&_Aggregator.TransactOpts, _newValidator)
}

// SetValidator is a paid mutator transaction binding the contract method 0x1327d3d8.
//
// Solidity: function setValidator(address _newValidator) returns()
func (_Aggregator *AggregatorTransactorSession) SetValidator(_newValidator common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.SetValidator(&_Aggregator.TransactOpts, _newValidator)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Aggregator *AggregatorTransactor) Submit(opts *bind.TransactOpts, _roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "submit", _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Aggregator *AggregatorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Submit(&_Aggregator.TransactOpts, _roundId, _submission)
}

// Submit is a paid mutator transaction binding the contract method 0x202ee0ed.
//
// Solidity: function submit(uint256 _roundId, int256 _submission) returns()
func (_Aggregator *AggregatorTransactorSession) Submit(_roundId *big.Int, _submission *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.Submit(&_Aggregator.TransactOpts, _roundId, _submission)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Aggregator *AggregatorTransactor) TransferAdmin(opts *bind.TransactOpts, _oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "transferAdmin", _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Aggregator *AggregatorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferAdmin(&_Aggregator.TransactOpts, _oracle, _newAdmin)
}

// TransferAdmin is a paid mutator transaction binding the contract method 0xe9ee6eeb.
//
// Solidity: function transferAdmin(address _oracle, address _newAdmin) returns()
func (_Aggregator *AggregatorTransactorSession) TransferAdmin(_oracle common.Address, _newAdmin common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferAdmin(&_Aggregator.TransactOpts, _oracle, _newAdmin)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Aggregator *AggregatorTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Aggregator.Contract.TransferOwnership(&_Aggregator.TransactOpts, newOwner)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Aggregator *AggregatorTransactor) UpdateAvailableFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateAvailableFunds")
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Aggregator *AggregatorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateAvailableFunds(&_Aggregator.TransactOpts)
}

// UpdateAvailableFunds is a paid mutator transaction binding the contract method 0x4f8fc3b5.
//
// Solidity: function updateAvailableFunds() returns()
func (_Aggregator *AggregatorTransactorSession) UpdateAvailableFunds() (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateAvailableFunds(&_Aggregator.TransactOpts)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Aggregator *AggregatorTransactor) UpdateFutureRounds(opts *bind.TransactOpts, _paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "updateFutureRounds", _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Aggregator *AggregatorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateFutureRounds(&_Aggregator.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// UpdateFutureRounds is a paid mutator transaction binding the contract method 0x38aa4c72.
//
// Solidity: function updateFutureRounds(uint128 _paymentAmount, uint32 _minSubmissions, uint32 _maxSubmissions, uint32 _restartDelay, uint32 _timeout) returns()
func (_Aggregator *AggregatorTransactorSession) UpdateFutureRounds(_paymentAmount *big.Int, _minSubmissions uint32, _maxSubmissions uint32, _restartDelay uint32, _timeout uint32) (*types.Transaction, error) {
	return _Aggregator.Contract.UpdateFutureRounds(&_Aggregator.TransactOpts, _paymentAmount, _minSubmissions, _maxSubmissions, _restartDelay, _timeout)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactor) WithdrawFunds(opts *bind.TransactOpts, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "withdrawFunds", _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.WithdrawFunds(&_Aggregator.TransactOpts, _recipient, _amount)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0xc1075329.
//
// Solidity: function withdrawFunds(address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactorSession) WithdrawFunds(_recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.WithdrawFunds(&_Aggregator.TransactOpts, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactor) WithdrawPayment(opts *bind.TransactOpts, _oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.contract.Transact(opts, "withdrawPayment", _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.WithdrawPayment(&_Aggregator.TransactOpts, _oracle, _recipient, _amount)
}

// WithdrawPayment is a paid mutator transaction binding the contract method 0x3d3d7714.
//
// Solidity: function withdrawPayment(address _oracle, address _recipient, uint256 _amount) returns()
func (_Aggregator *AggregatorTransactorSession) WithdrawPayment(_oracle common.Address, _recipient common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Aggregator.Contract.WithdrawPayment(&_Aggregator.TransactOpts, _oracle, _recipient, _amount)
}

// AggregatorAnswerUpdatedIterator is returned from FilterAnswerUpdated and is used to iterate over the raw logs and unpacked data for AnswerUpdated events raised by the Aggregator contract.
type AggregatorAnswerUpdatedIterator struct {
	Event *AggregatorAnswerUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorAnswerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorAnswerUpdated)
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
		it.Event = new(AggregatorAnswerUpdated)
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
func (it *AggregatorAnswerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorAnswerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorAnswerUpdated represents a AnswerUpdated event raised by the Aggregator contract.
type AggregatorAnswerUpdated struct {
	Current   *big.Int
	RoundId   *big.Int
	UpdatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAnswerUpdated is a free log retrieval operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Aggregator *AggregatorFilterer) FilterAnswerUpdated(opts *bind.FilterOpts, current []*big.Int, roundId []*big.Int) (*AggregatorAnswerUpdatedIterator, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorAnswerUpdatedIterator{contract: _Aggregator.contract, event: "AnswerUpdated", logs: logs, sub: sub}, nil
}

// WatchAnswerUpdated is a free log subscription operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Aggregator *AggregatorFilterer) WatchAnswerUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorAnswerUpdated, current []*big.Int, roundId []*big.Int) (event.Subscription, error) {

	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}
	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "AnswerUpdated", currentRule, roundIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorAnswerUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
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

// ParseAnswerUpdated is a log parse operation binding the contract event 0x0559884fd3a460db3073b7fc896cc77986f16e378210ded43186175bf646fc5f.
//
// Solidity: event AnswerUpdated(int256 indexed current, uint256 indexed roundId, uint256 updatedAt)
func (_Aggregator *AggregatorFilterer) ParseAnswerUpdated(log types.Log) (*AggregatorAnswerUpdated, error) {
	event := new(AggregatorAnswerUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "AnswerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorAvailableFundsUpdatedIterator is returned from FilterAvailableFundsUpdated and is used to iterate over the raw logs and unpacked data for AvailableFundsUpdated events raised by the Aggregator contract.
type AggregatorAvailableFundsUpdatedIterator struct {
	Event *AggregatorAvailableFundsUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorAvailableFundsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorAvailableFundsUpdated)
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
		it.Event = new(AggregatorAvailableFundsUpdated)
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
func (it *AggregatorAvailableFundsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorAvailableFundsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorAvailableFundsUpdated represents a AvailableFundsUpdated event raised by the Aggregator contract.
type AggregatorAvailableFundsUpdated struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAvailableFundsUpdated is a free log retrieval operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Aggregator *AggregatorFilterer) FilterAvailableFundsUpdated(opts *bind.FilterOpts, amount []*big.Int) (*AggregatorAvailableFundsUpdatedIterator, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorAvailableFundsUpdatedIterator{contract: _Aggregator.contract, event: "AvailableFundsUpdated", logs: logs, sub: sub}, nil
}

// WatchAvailableFundsUpdated is a free log subscription operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Aggregator *AggregatorFilterer) WatchAvailableFundsUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorAvailableFundsUpdated, amount []*big.Int) (event.Subscription, error) {

	var amountRule []interface{}
	for _, amountItem := range amount {
		amountRule = append(amountRule, amountItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "AvailableFundsUpdated", amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorAvailableFundsUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
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

// ParseAvailableFundsUpdated is a log parse operation binding the contract event 0xfe25c73e3b9089fac37d55c4c7efcba6f04af04cebd2fc4d6d7dbb07e1e5234f.
//
// Solidity: event AvailableFundsUpdated(uint256 indexed amount)
func (_Aggregator *AggregatorFilterer) ParseAvailableFundsUpdated(log types.Log) (*AggregatorAvailableFundsUpdated, error) {
	event := new(AggregatorAvailableFundsUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "AvailableFundsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorNewRoundIterator is returned from FilterNewRound and is used to iterate over the raw logs and unpacked data for NewRound events raised by the Aggregator contract.
type AggregatorNewRoundIterator struct {
	Event *AggregatorNewRound // Event containing the contract specifics and raw log

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
func (it *AggregatorNewRoundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorNewRound)
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
		it.Event = new(AggregatorNewRound)
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
func (it *AggregatorNewRoundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorNewRoundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorNewRound represents a NewRound event raised by the Aggregator contract.
type AggregatorNewRound struct {
	RoundId   *big.Int
	StartedBy common.Address
	StartedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewRound is a free log retrieval operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Aggregator *AggregatorFilterer) FilterNewRound(opts *bind.FilterOpts, roundId []*big.Int, startedBy []common.Address) (*AggregatorNewRoundIterator, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorNewRoundIterator{contract: _Aggregator.contract, event: "NewRound", logs: logs, sub: sub}, nil
}

// WatchNewRound is a free log subscription operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Aggregator *AggregatorFilterer) WatchNewRound(opts *bind.WatchOpts, sink chan<- *AggregatorNewRound, roundId []*big.Int, startedBy []common.Address) (event.Subscription, error) {

	var roundIdRule []interface{}
	for _, roundIdItem := range roundId {
		roundIdRule = append(roundIdRule, roundIdItem)
	}
	var startedByRule []interface{}
	for _, startedByItem := range startedBy {
		startedByRule = append(startedByRule, startedByItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "NewRound", roundIdRule, startedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorNewRound)
				if err := _Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
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

// ParseNewRound is a log parse operation binding the contract event 0x0109fc6f55cf40689f02fbaad7af7fe7bbac8a3d2186600afc7d3e10cac60271.
//
// Solidity: event NewRound(uint256 indexed roundId, address indexed startedBy, uint256 startedAt)
func (_Aggregator *AggregatorFilterer) ParseNewRound(log types.Log) (*AggregatorNewRound, error) {
	event := new(AggregatorNewRound)
	if err := _Aggregator.contract.UnpackLog(event, "NewRound", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorOracleAdminUpdateRequestedIterator is returned from FilterOracleAdminUpdateRequested and is used to iterate over the raw logs and unpacked data for OracleAdminUpdateRequested events raised by the Aggregator contract.
type AggregatorOracleAdminUpdateRequestedIterator struct {
	Event *AggregatorOracleAdminUpdateRequested // Event containing the contract specifics and raw log

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
func (it *AggregatorOracleAdminUpdateRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOracleAdminUpdateRequested)
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
		it.Event = new(AggregatorOracleAdminUpdateRequested)
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
func (it *AggregatorOracleAdminUpdateRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOracleAdminUpdateRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOracleAdminUpdateRequested represents a OracleAdminUpdateRequested event raised by the Aggregator contract.
type AggregatorOracleAdminUpdateRequested struct {
	Oracle   common.Address
	Admin    common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdateRequested is a free log retrieval operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Aggregator *AggregatorFilterer) FilterOracleAdminUpdateRequested(opts *bind.FilterOpts, oracle []common.Address) (*AggregatorOracleAdminUpdateRequestedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOracleAdminUpdateRequestedIterator{contract: _Aggregator.contract, event: "OracleAdminUpdateRequested", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdateRequested is a free log subscription operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Aggregator *AggregatorFilterer) WatchOracleAdminUpdateRequested(opts *bind.WatchOpts, sink chan<- *AggregatorOracleAdminUpdateRequested, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OracleAdminUpdateRequested", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOracleAdminUpdateRequested)
				if err := _Aggregator.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
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

// ParseOracleAdminUpdateRequested is a log parse operation binding the contract event 0xb79bf2e89c2d70dde91d2991fb1ea69b7e478061ad7c04ed5b02b96bc52b8104.
//
// Solidity: event OracleAdminUpdateRequested(address indexed oracle, address admin, address newAdmin)
func (_Aggregator *AggregatorFilterer) ParseOracleAdminUpdateRequested(log types.Log) (*AggregatorOracleAdminUpdateRequested, error) {
	event := new(AggregatorOracleAdminUpdateRequested)
	if err := _Aggregator.contract.UnpackLog(event, "OracleAdminUpdateRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorOracleAdminUpdatedIterator is returned from FilterOracleAdminUpdated and is used to iterate over the raw logs and unpacked data for OracleAdminUpdated events raised by the Aggregator contract.
type AggregatorOracleAdminUpdatedIterator struct {
	Event *AggregatorOracleAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorOracleAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOracleAdminUpdated)
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
		it.Event = new(AggregatorOracleAdminUpdated)
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
func (it *AggregatorOracleAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOracleAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOracleAdminUpdated represents a OracleAdminUpdated event raised by the Aggregator contract.
type AggregatorOracleAdminUpdated struct {
	Oracle   common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOracleAdminUpdated is a free log retrieval operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Aggregator *AggregatorFilterer) FilterOracleAdminUpdated(opts *bind.FilterOpts, oracle []common.Address, newAdmin []common.Address) (*AggregatorOracleAdminUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOracleAdminUpdatedIterator{contract: _Aggregator.contract, event: "OracleAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchOracleAdminUpdated is a free log subscription operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Aggregator *AggregatorFilterer) WatchOracleAdminUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorOracleAdminUpdated, oracle []common.Address, newAdmin []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OracleAdminUpdated", oracleRule, newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOracleAdminUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
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

// ParseOracleAdminUpdated is a log parse operation binding the contract event 0x0c5055390645c15a4be9a21b3f8d019153dcb4a0c125685da6eb84048e2fe904.
//
// Solidity: event OracleAdminUpdated(address indexed oracle, address indexed newAdmin)
func (_Aggregator *AggregatorFilterer) ParseOracleAdminUpdated(log types.Log) (*AggregatorOracleAdminUpdated, error) {
	event := new(AggregatorOracleAdminUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "OracleAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorOraclePermissionsUpdatedIterator is returned from FilterOraclePermissionsUpdated and is used to iterate over the raw logs and unpacked data for OraclePermissionsUpdated events raised by the Aggregator contract.
type AggregatorOraclePermissionsUpdatedIterator struct {
	Event *AggregatorOraclePermissionsUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorOraclePermissionsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOraclePermissionsUpdated)
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
		it.Event = new(AggregatorOraclePermissionsUpdated)
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
func (it *AggregatorOraclePermissionsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOraclePermissionsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOraclePermissionsUpdated represents a OraclePermissionsUpdated event raised by the Aggregator contract.
type AggregatorOraclePermissionsUpdated struct {
	Oracle      common.Address
	Whitelisted bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOraclePermissionsUpdated is a free log retrieval operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Aggregator *AggregatorFilterer) FilterOraclePermissionsUpdated(opts *bind.FilterOpts, oracle []common.Address, whitelisted []bool) (*AggregatorOraclePermissionsUpdatedIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOraclePermissionsUpdatedIterator{contract: _Aggregator.contract, event: "OraclePermissionsUpdated", logs: logs, sub: sub}, nil
}

// WatchOraclePermissionsUpdated is a free log subscription operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Aggregator *AggregatorFilterer) WatchOraclePermissionsUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorOraclePermissionsUpdated, oracle []common.Address, whitelisted []bool) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}
	var whitelistedRule []interface{}
	for _, whitelistedItem := range whitelisted {
		whitelistedRule = append(whitelistedRule, whitelistedItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OraclePermissionsUpdated", oracleRule, whitelistedRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOraclePermissionsUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
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

// ParseOraclePermissionsUpdated is a log parse operation binding the contract event 0x18dd09695e4fbdae8d1a5edb11221eb04564269c29a089b9753a6535c54ba92e.
//
// Solidity: event OraclePermissionsUpdated(address indexed oracle, bool indexed whitelisted)
func (_Aggregator *AggregatorFilterer) ParseOraclePermissionsUpdated(log types.Log) (*AggregatorOraclePermissionsUpdated, error) {
	event := new(AggregatorOraclePermissionsUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "OraclePermissionsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Aggregator contract.
type AggregatorOwnershipTransferredIterator struct {
	Event *AggregatorOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregatorOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorOwnershipTransferred)
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
		it.Event = new(AggregatorOwnershipTransferred)
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
func (it *AggregatorOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorOwnershipTransferred represents a OwnershipTransferred event raised by the Aggregator contract.
type AggregatorOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregatorOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorOwnershipTransferredIterator{contract: _Aggregator.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Aggregator *AggregatorFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregatorOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorOwnershipTransferred)
				if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Aggregator *AggregatorFilterer) ParseOwnershipTransferred(log types.Log) (*AggregatorOwnershipTransferred, error) {
	event := new(AggregatorOwnershipTransferred)
	if err := _Aggregator.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorRequesterPermissionsSetIterator is returned from FilterRequesterPermissionsSet and is used to iterate over the raw logs and unpacked data for RequesterPermissionsSet events raised by the Aggregator contract.
type AggregatorRequesterPermissionsSetIterator struct {
	Event *AggregatorRequesterPermissionsSet // Event containing the contract specifics and raw log

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
func (it *AggregatorRequesterPermissionsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorRequesterPermissionsSet)
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
		it.Event = new(AggregatorRequesterPermissionsSet)
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
func (it *AggregatorRequesterPermissionsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorRequesterPermissionsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorRequesterPermissionsSet represents a RequesterPermissionsSet event raised by the Aggregator contract.
type AggregatorRequesterPermissionsSet struct {
	Requester  common.Address
	Authorized bool
	Delay      uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRequesterPermissionsSet is a free log retrieval operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Aggregator *AggregatorFilterer) FilterRequesterPermissionsSet(opts *bind.FilterOpts, requester []common.Address) (*AggregatorRequesterPermissionsSetIterator, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorRequesterPermissionsSetIterator{contract: _Aggregator.contract, event: "RequesterPermissionsSet", logs: logs, sub: sub}, nil
}

// WatchRequesterPermissionsSet is a free log subscription operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Aggregator *AggregatorFilterer) WatchRequesterPermissionsSet(opts *bind.WatchOpts, sink chan<- *AggregatorRequesterPermissionsSet, requester []common.Address) (event.Subscription, error) {

	var requesterRule []interface{}
	for _, requesterItem := range requester {
		requesterRule = append(requesterRule, requesterItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "RequesterPermissionsSet", requesterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorRequesterPermissionsSet)
				if err := _Aggregator.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
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

// ParseRequesterPermissionsSet is a log parse operation binding the contract event 0xc3df5a754e002718f2e10804b99e6605e7c701d95cec9552c7680ca2b6f2820a.
//
// Solidity: event RequesterPermissionsSet(address indexed requester, bool authorized, uint32 delay)
func (_Aggregator *AggregatorFilterer) ParseRequesterPermissionsSet(log types.Log) (*AggregatorRequesterPermissionsSet, error) {
	event := new(AggregatorRequesterPermissionsSet)
	if err := _Aggregator.contract.UnpackLog(event, "RequesterPermissionsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorRoundDetailsUpdatedIterator is returned from FilterRoundDetailsUpdated and is used to iterate over the raw logs and unpacked data for RoundDetailsUpdated events raised by the Aggregator contract.
type AggregatorRoundDetailsUpdatedIterator struct {
	Event *AggregatorRoundDetailsUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorRoundDetailsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorRoundDetailsUpdated)
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
		it.Event = new(AggregatorRoundDetailsUpdated)
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
func (it *AggregatorRoundDetailsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorRoundDetailsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorRoundDetailsUpdated represents a RoundDetailsUpdated event raised by the Aggregator contract.
type AggregatorRoundDetailsUpdated struct {
	PaymentAmount      *big.Int
	MinSubmissionCount uint32
	MaxSubmissionCount uint32
	RestartDelay       uint32
	Timeout            uint32
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRoundDetailsUpdated is a free log retrieval operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Aggregator *AggregatorFilterer) FilterRoundDetailsUpdated(opts *bind.FilterOpts, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (*AggregatorRoundDetailsUpdatedIterator, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorRoundDetailsUpdatedIterator{contract: _Aggregator.contract, event: "RoundDetailsUpdated", logs: logs, sub: sub}, nil
}

// WatchRoundDetailsUpdated is a free log subscription operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Aggregator *AggregatorFilterer) WatchRoundDetailsUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorRoundDetailsUpdated, paymentAmount []*big.Int, minSubmissionCount []uint32, maxSubmissionCount []uint32) (event.Subscription, error) {

	var paymentAmountRule []interface{}
	for _, paymentAmountItem := range paymentAmount {
		paymentAmountRule = append(paymentAmountRule, paymentAmountItem)
	}
	var minSubmissionCountRule []interface{}
	for _, minSubmissionCountItem := range minSubmissionCount {
		minSubmissionCountRule = append(minSubmissionCountRule, minSubmissionCountItem)
	}
	var maxSubmissionCountRule []interface{}
	for _, maxSubmissionCountItem := range maxSubmissionCount {
		maxSubmissionCountRule = append(maxSubmissionCountRule, maxSubmissionCountItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "RoundDetailsUpdated", paymentAmountRule, minSubmissionCountRule, maxSubmissionCountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorRoundDetailsUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
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

// ParseRoundDetailsUpdated is a log parse operation binding the contract event 0x56800c9d1ed723511246614d15e58cfcde15b6a33c245b5c961b689c1890fd8f.
//
// Solidity: event RoundDetailsUpdated(uint128 indexed paymentAmount, uint32 indexed minSubmissionCount, uint32 indexed maxSubmissionCount, uint32 restartDelay, uint32 timeout)
func (_Aggregator *AggregatorFilterer) ParseRoundDetailsUpdated(log types.Log) (*AggregatorRoundDetailsUpdated, error) {
	event := new(AggregatorRoundDetailsUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "RoundDetailsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorSubmissionReceivedIterator is returned from FilterSubmissionReceived and is used to iterate over the raw logs and unpacked data for SubmissionReceived events raised by the Aggregator contract.
type AggregatorSubmissionReceivedIterator struct {
	Event *AggregatorSubmissionReceived // Event containing the contract specifics and raw log

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
func (it *AggregatorSubmissionReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorSubmissionReceived)
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
		it.Event = new(AggregatorSubmissionReceived)
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
func (it *AggregatorSubmissionReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorSubmissionReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorSubmissionReceived represents a SubmissionReceived event raised by the Aggregator contract.
type AggregatorSubmissionReceived struct {
	Submission *big.Int
	Round      uint32
	Oracle     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubmissionReceived is a free log retrieval operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Aggregator *AggregatorFilterer) FilterSubmissionReceived(opts *bind.FilterOpts, submission []*big.Int, round []uint32, oracle []common.Address) (*AggregatorSubmissionReceivedIterator, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorSubmissionReceivedIterator{contract: _Aggregator.contract, event: "SubmissionReceived", logs: logs, sub: sub}, nil
}

// WatchSubmissionReceived is a free log subscription operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Aggregator *AggregatorFilterer) WatchSubmissionReceived(opts *bind.WatchOpts, sink chan<- *AggregatorSubmissionReceived, submission []*big.Int, round []uint32, oracle []common.Address) (event.Subscription, error) {

	var submissionRule []interface{}
	for _, submissionItem := range submission {
		submissionRule = append(submissionRule, submissionItem)
	}
	var roundRule []interface{}
	for _, roundItem := range round {
		roundRule = append(roundRule, roundItem)
	}
	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "SubmissionReceived", submissionRule, roundRule, oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorSubmissionReceived)
				if err := _Aggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
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

// ParseSubmissionReceived is a log parse operation binding the contract event 0x92e98423f8adac6e64d0608e519fd1cefb861498385c6dee70d58fc926ddc68c.
//
// Solidity: event SubmissionReceived(int256 indexed submission, uint32 indexed round, address indexed oracle)
func (_Aggregator *AggregatorFilterer) ParseSubmissionReceived(log types.Log) (*AggregatorSubmissionReceived, error) {
	event := new(AggregatorSubmissionReceived)
	if err := _Aggregator.contract.UnpackLog(event, "SubmissionReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregatorValidatorUpdatedIterator is returned from FilterValidatorUpdated and is used to iterate over the raw logs and unpacked data for ValidatorUpdated events raised by the Aggregator contract.
type AggregatorValidatorUpdatedIterator struct {
	Event *AggregatorValidatorUpdated // Event containing the contract specifics and raw log

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
func (it *AggregatorValidatorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregatorValidatorUpdated)
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
		it.Event = new(AggregatorValidatorUpdated)
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
func (it *AggregatorValidatorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregatorValidatorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregatorValidatorUpdated represents a ValidatorUpdated event raised by the Aggregator contract.
type AggregatorValidatorUpdated struct {
	Previous common.Address
	Current  common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterValidatorUpdated is a free log retrieval operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Aggregator *AggregatorFilterer) FilterValidatorUpdated(opts *bind.FilterOpts, previous []common.Address, current []common.Address) (*AggregatorValidatorUpdatedIterator, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Aggregator.contract.FilterLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return &AggregatorValidatorUpdatedIterator{contract: _Aggregator.contract, event: "ValidatorUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorUpdated is a free log subscription operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Aggregator *AggregatorFilterer) WatchValidatorUpdated(opts *bind.WatchOpts, sink chan<- *AggregatorValidatorUpdated, previous []common.Address, current []common.Address) (event.Subscription, error) {

	var previousRule []interface{}
	for _, previousItem := range previous {
		previousRule = append(previousRule, previousItem)
	}
	var currentRule []interface{}
	for _, currentItem := range current {
		currentRule = append(currentRule, currentItem)
	}

	logs, sub, err := _Aggregator.contract.WatchLogs(opts, "ValidatorUpdated", previousRule, currentRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregatorValidatorUpdated)
				if err := _Aggregator.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
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

// ParseValidatorUpdated is a log parse operation binding the contract event 0xcfac5dc75b8d9a7e074162f59d9adcd33da59f0fe8dfb21580db298fc0fdad0d.
//
// Solidity: event ValidatorUpdated(address indexed previous, address indexed current)
func (_Aggregator *AggregatorFilterer) ParseValidatorUpdated(log types.Log) (*AggregatorValidatorUpdated, error) {
	event := new(AggregatorValidatorUpdated)
	if err := _Aggregator.contract.UnpackLog(event, "ValidatorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
