// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package xen_abi

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

// XENCryptoMintInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoMintInfo struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}

// XENCryptoStakeInfo is an auto generated low-level Go binding around an user-defined struct.
type XENCryptoStakeInfo struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"MintClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"}],\"name\":\"RankClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AUTHORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DAYS_IN_YEAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_PM_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EAA_RANK_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_RANK\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_PENALTY_PCT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_TERM_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_TERM\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_AMPLIFIER_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_IN_DAY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TERM_AMPLIFIER_THRESHOLD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_WINDOW_DAYS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_DAYS_STEP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_END\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_APY_START\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_BURN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"XEN_MIN_STAKE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeMinters\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activeStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimMintReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"other\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndShare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"pct\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimMintRewardAndStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"claimRank\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"genesisTs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAMP\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentAPY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentEAAR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentMaxTerm\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rankDelta\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaa\",\"type\":\"uint256\"}],\"name\":\"getGrossReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserMint\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.MintInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserStake\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"internalType\":\"structXENCrypto.StakeInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalRank\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalXenStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBurns\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userMints\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amplifier\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"eaaRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"term\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTs\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"apy\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Store *StoreCaller) AUTHORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "AUTHORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Store *StoreSession) AUTHORS() (string, error) {
	return _Store.Contract.AUTHORS(&_Store.CallOpts)
}

// AUTHORS is a free data retrieval call binding the contract method 0xba3ec741.
//
// Solidity: function AUTHORS() view returns(string)
func (_Store *StoreCallerSession) AUTHORS() (string, error) {
	return _Store.Contract.AUTHORS(&_Store.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Store *StoreCaller) DAYSINYEAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DAYS_IN_YEAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Store *StoreSession) DAYSINYEAR() (*big.Int, error) {
	return _Store.Contract.DAYSINYEAR(&_Store.CallOpts)
}

// DAYSINYEAR is a free data retrieval call binding the contract method 0x02378932.
//
// Solidity: function DAYS_IN_YEAR() view returns(uint256)
func (_Store *StoreCallerSession) DAYSINYEAR() (*big.Int, error) {
	return _Store.Contract.DAYSINYEAR(&_Store.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Store *StoreCaller) EAAPMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "EAA_PM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Store *StoreSession) EAAPMSTART() (*big.Int, error) {
	return _Store.Contract.EAAPMSTART(&_Store.CallOpts)
}

// EAAPMSTART is a free data retrieval call binding the contract method 0x909a2ff6.
//
// Solidity: function EAA_PM_START() view returns(uint256)
func (_Store *StoreCallerSession) EAAPMSTART() (*big.Int, error) {
	return _Store.Contract.EAAPMSTART(&_Store.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Store *StoreCaller) EAAPMSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "EAA_PM_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Store *StoreSession) EAAPMSTEP() (*big.Int, error) {
	return _Store.Contract.EAAPMSTEP(&_Store.CallOpts)
}

// EAAPMSTEP is a free data retrieval call binding the contract method 0xbcfe394f.
//
// Solidity: function EAA_PM_STEP() view returns(uint256)
func (_Store *StoreCallerSession) EAAPMSTEP() (*big.Int, error) {
	return _Store.Contract.EAAPMSTEP(&_Store.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Store *StoreCaller) EAARANKSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "EAA_RANK_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Store *StoreSession) EAARANKSTEP() (*big.Int, error) {
	return _Store.Contract.EAARANKSTEP(&_Store.CallOpts)
}

// EAARANKSTEP is a free data retrieval call binding the contract method 0xf340faed.
//
// Solidity: function EAA_RANK_STEP() view returns(uint256)
func (_Store *StoreCallerSession) EAARANKSTEP() (*big.Int, error) {
	return _Store.Contract.EAARANKSTEP(&_Store.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Store *StoreCaller) GENESISRANK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "GENESIS_RANK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Store *StoreSession) GENESISRANK() (*big.Int, error) {
	return _Store.Contract.GENESISRANK(&_Store.CallOpts)
}

// GENESISRANK is a free data retrieval call binding the contract method 0x1c6f212e.
//
// Solidity: function GENESIS_RANK() view returns(uint256)
func (_Store *StoreCallerSession) GENESISRANK() (*big.Int, error) {
	return _Store.Contract.GENESISRANK(&_Store.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Store *StoreCaller) MAXPENALTYPCT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MAX_PENALTY_PCT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Store *StoreSession) MAXPENALTYPCT() (*big.Int, error) {
	return _Store.Contract.MAXPENALTYPCT(&_Store.CallOpts)
}

// MAXPENALTYPCT is a free data retrieval call binding the contract method 0x0f2e1228.
//
// Solidity: function MAX_PENALTY_PCT() view returns(uint256)
func (_Store *StoreCallerSession) MAXPENALTYPCT() (*big.Int, error) {
	return _Store.Contract.MAXPENALTYPCT(&_Store.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Store *StoreCaller) MAXTERMEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MAX_TERM_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Store *StoreSession) MAXTERMEND() (*big.Int, error) {
	return _Store.Contract.MAXTERMEND(&_Store.CallOpts)
}

// MAXTERMEND is a free data retrieval call binding the contract method 0xc0c65259.
//
// Solidity: function MAX_TERM_END() view returns(uint256)
func (_Store *StoreCallerSession) MAXTERMEND() (*big.Int, error) {
	return _Store.Contract.MAXTERMEND(&_Store.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Store *StoreCaller) MAXTERMSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MAX_TERM_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Store *StoreSession) MAXTERMSTART() (*big.Int, error) {
	return _Store.Contract.MAXTERMSTART(&_Store.CallOpts)
}

// MAXTERMSTART is a free data retrieval call binding the contract method 0x0bfae56b.
//
// Solidity: function MAX_TERM_START() view returns(uint256)
func (_Store *StoreCallerSession) MAXTERMSTART() (*big.Int, error) {
	return _Store.Contract.MAXTERMSTART(&_Store.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Store *StoreCaller) MINTERM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MIN_TERM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Store *StoreSession) MINTERM() (*big.Int, error) {
	return _Store.Contract.MINTERM(&_Store.CallOpts)
}

// MINTERM is a free data retrieval call binding the contract method 0xf0604829.
//
// Solidity: function MIN_TERM() view returns(uint256)
func (_Store *StoreCallerSession) MINTERM() (*big.Int, error) {
	return _Store.Contract.MINTERM(&_Store.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Store *StoreCaller) REWARDAMPLIFIEREND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "REWARD_AMPLIFIER_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Store *StoreSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _Store.Contract.REWARDAMPLIFIEREND(&_Store.CallOpts)
}

// REWARDAMPLIFIEREND is a free data retrieval call binding the contract method 0x543d3652.
//
// Solidity: function REWARD_AMPLIFIER_END() view returns(uint256)
func (_Store *StoreCallerSession) REWARDAMPLIFIEREND() (*big.Int, error) {
	return _Store.Contract.REWARDAMPLIFIEREND(&_Store.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Store *StoreCaller) REWARDAMPLIFIERSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "REWARD_AMPLIFIER_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Store *StoreSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _Store.Contract.REWARDAMPLIFIERSTART(&_Store.CallOpts)
}

// REWARDAMPLIFIERSTART is a free data retrieval call binding the contract method 0xe81917b4.
//
// Solidity: function REWARD_AMPLIFIER_START() view returns(uint256)
func (_Store *StoreCallerSession) REWARDAMPLIFIERSTART() (*big.Int, error) {
	return _Store.Contract.REWARDAMPLIFIERSTART(&_Store.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Store *StoreCaller) SECONDSINDAY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "SECONDS_IN_DAY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Store *StoreSession) SECONDSINDAY() (*big.Int, error) {
	return _Store.Contract.SECONDSINDAY(&_Store.CallOpts)
}

// SECONDSINDAY is a free data retrieval call binding the contract method 0x61a52a36.
//
// Solidity: function SECONDS_IN_DAY() view returns(uint256)
func (_Store *StoreCallerSession) SECONDSINDAY() (*big.Int, error) {
	return _Store.Contract.SECONDSINDAY(&_Store.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Store *StoreCaller) TERMAMPLIFIER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "TERM_AMPLIFIER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Store *StoreSession) TERMAMPLIFIER() (*big.Int, error) {
	return _Store.Contract.TERMAMPLIFIER(&_Store.CallOpts)
}

// TERMAMPLIFIER is a free data retrieval call binding the contract method 0x72475f94.
//
// Solidity: function TERM_AMPLIFIER() view returns(uint256)
func (_Store *StoreCallerSession) TERMAMPLIFIER() (*big.Int, error) {
	return _Store.Contract.TERMAMPLIFIER(&_Store.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Store *StoreCaller) TERMAMPLIFIERTHRESHOLD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "TERM_AMPLIFIER_THRESHOLD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Store *StoreSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _Store.Contract.TERMAMPLIFIERTHRESHOLD(&_Store.CallOpts)
}

// TERMAMPLIFIERTHRESHOLD is a free data retrieval call binding the contract method 0xb21d35f2.
//
// Solidity: function TERM_AMPLIFIER_THRESHOLD() view returns(uint256)
func (_Store *StoreCallerSession) TERMAMPLIFIERTHRESHOLD() (*big.Int, error) {
	return _Store.Contract.TERMAMPLIFIERTHRESHOLD(&_Store.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Store *StoreCaller) WITHDRAWALWINDOWDAYS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "WITHDRAWAL_WINDOW_DAYS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Store *StoreSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _Store.Contract.WITHDRAWALWINDOWDAYS(&_Store.CallOpts)
}

// WITHDRAWALWINDOWDAYS is a free data retrieval call binding the contract method 0xc56f0bab.
//
// Solidity: function WITHDRAWAL_WINDOW_DAYS() view returns(uint256)
func (_Store *StoreCallerSession) WITHDRAWALWINDOWDAYS() (*big.Int, error) {
	return _Store.Contract.WITHDRAWALWINDOWDAYS(&_Store.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Store *StoreCaller) XENAPYDAYSSTEP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "XEN_APY_DAYS_STEP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Store *StoreSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _Store.Contract.XENAPYDAYSSTEP(&_Store.CallOpts)
}

// XENAPYDAYSSTEP is a free data retrieval call binding the contract method 0x32870fda.
//
// Solidity: function XEN_APY_DAYS_STEP() view returns(uint256)
func (_Store *StoreCallerSession) XENAPYDAYSSTEP() (*big.Int, error) {
	return _Store.Contract.XENAPYDAYSSTEP(&_Store.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Store *StoreCaller) XENAPYEND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "XEN_APY_END")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Store *StoreSession) XENAPYEND() (*big.Int, error) {
	return _Store.Contract.XENAPYEND(&_Store.CallOpts)
}

// XENAPYEND is a free data retrieval call binding the contract method 0x7e7aa62e.
//
// Solidity: function XEN_APY_END() view returns(uint256)
func (_Store *StoreCallerSession) XENAPYEND() (*big.Int, error) {
	return _Store.Contract.XENAPYEND(&_Store.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Store *StoreCaller) XENAPYSTART(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "XEN_APY_START")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Store *StoreSession) XENAPYSTART() (*big.Int, error) {
	return _Store.Contract.XENAPYSTART(&_Store.CallOpts)
}

// XENAPYSTART is a free data retrieval call binding the contract method 0xfed74269.
//
// Solidity: function XEN_APY_START() view returns(uint256)
func (_Store *StoreCallerSession) XENAPYSTART() (*big.Int, error) {
	return _Store.Contract.XENAPYSTART(&_Store.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Store *StoreCaller) XENMINBURN(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "XEN_MIN_BURN")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Store *StoreSession) XENMINBURN() (*big.Int, error) {
	return _Store.Contract.XENMINBURN(&_Store.CallOpts)
}

// XENMINBURN is a free data retrieval call binding the contract method 0x110d7fc2.
//
// Solidity: function XEN_MIN_BURN() view returns(uint256)
func (_Store *StoreCallerSession) XENMINBURN() (*big.Int, error) {
	return _Store.Contract.XENMINBURN(&_Store.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Store *StoreCaller) XENMINSTAKE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "XEN_MIN_STAKE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Store *StoreSession) XENMINSTAKE() (*big.Int, error) {
	return _Store.Contract.XENMINSTAKE(&_Store.CallOpts)
}

// XENMINSTAKE is a free data retrieval call binding the contract method 0x2a62d966.
//
// Solidity: function XEN_MIN_STAKE() view returns(uint256)
func (_Store *StoreCallerSession) XENMINSTAKE() (*big.Int, error) {
	return _Store.Contract.XENMINSTAKE(&_Store.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Store *StoreCaller) ActiveMinters(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "activeMinters")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Store *StoreSession) ActiveMinters() (*big.Int, error) {
	return _Store.Contract.ActiveMinters(&_Store.CallOpts)
}

// ActiveMinters is a free data retrieval call binding the contract method 0xb4800cdc.
//
// Solidity: function activeMinters() view returns(uint256)
func (_Store *StoreCallerSession) ActiveMinters() (*big.Int, error) {
	return _Store.Contract.ActiveMinters(&_Store.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Store *StoreCaller) ActiveStakes(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "activeStakes")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Store *StoreSession) ActiveStakes() (*big.Int, error) {
	return _Store.Contract.ActiveStakes(&_Store.CallOpts)
}

// ActiveStakes is a free data retrieval call binding the contract method 0xed2f2369.
//
// Solidity: function activeStakes() view returns(uint256)
func (_Store *StoreCallerSession) ActiveStakes() (*big.Int, error) {
	return _Store.Contract.ActiveStakes(&_Store.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCallerSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Store *StoreCaller) GenesisTs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "genesisTs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Store *StoreSession) GenesisTs() (*big.Int, error) {
	return _Store.Contract.GenesisTs(&_Store.CallOpts)
}

// GenesisTs is a free data retrieval call binding the contract method 0xe3af6d0a.
//
// Solidity: function genesisTs() view returns(uint256)
func (_Store *StoreCallerSession) GenesisTs() (*big.Int, error) {
	return _Store.Contract.GenesisTs(&_Store.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Store *StoreCaller) GetCurrentAMP(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCurrentAMP")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Store *StoreSession) GetCurrentAMP() (*big.Int, error) {
	return _Store.Contract.GetCurrentAMP(&_Store.CallOpts)
}

// GetCurrentAMP is a free data retrieval call binding the contract method 0x99202454.
//
// Solidity: function getCurrentAMP() view returns(uint256)
func (_Store *StoreCallerSession) GetCurrentAMP() (*big.Int, error) {
	return _Store.Contract.GetCurrentAMP(&_Store.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Store *StoreCaller) GetCurrentAPY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCurrentAPY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Store *StoreSession) GetCurrentAPY() (*big.Int, error) {
	return _Store.Contract.GetCurrentAPY(&_Store.CallOpts)
}

// GetCurrentAPY is a free data retrieval call binding the contract method 0x962ca496.
//
// Solidity: function getCurrentAPY() view returns(uint256)
func (_Store *StoreCallerSession) GetCurrentAPY() (*big.Int, error) {
	return _Store.Contract.GetCurrentAPY(&_Store.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Store *StoreCaller) GetCurrentEAAR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCurrentEAAR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Store *StoreSession) GetCurrentEAAR() (*big.Int, error) {
	return _Store.Contract.GetCurrentEAAR(&_Store.CallOpts)
}

// GetCurrentEAAR is a free data retrieval call binding the contract method 0x8979c87c.
//
// Solidity: function getCurrentEAAR() view returns(uint256)
func (_Store *StoreCallerSession) GetCurrentEAAR() (*big.Int, error) {
	return _Store.Contract.GetCurrentEAAR(&_Store.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Store *StoreCaller) GetCurrentMaxTerm(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getCurrentMaxTerm")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Store *StoreSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _Store.Contract.GetCurrentMaxTerm(&_Store.CallOpts)
}

// GetCurrentMaxTerm is a free data retrieval call binding the contract method 0x45125715.
//
// Solidity: function getCurrentMaxTerm() view returns(uint256)
func (_Store *StoreCallerSession) GetCurrentMaxTerm() (*big.Int, error) {
	return _Store.Contract.GetCurrentMaxTerm(&_Store.CallOpts)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Store *StoreCaller) GetGrossReward(opts *bind.CallOpts, rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getGrossReward", rankDelta, amplifier, term, eaa)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Store *StoreSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _Store.Contract.GetGrossReward(&_Store.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetGrossReward is a free data retrieval call binding the contract method 0xb0fd1fc2.
//
// Solidity: function getGrossReward(uint256 rankDelta, uint256 amplifier, uint256 term, uint256 eaa) pure returns(uint256)
func (_Store *StoreCallerSession) GetGrossReward(rankDelta *big.Int, amplifier *big.Int, term *big.Int, eaa *big.Int) (*big.Int, error) {
	return _Store.Contract.GetGrossReward(&_Store.CallOpts, rankDelta, amplifier, term, eaa)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Store *StoreCaller) GetUserMint(opts *bind.CallOpts) (XENCryptoMintInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getUserMint")

	if err != nil {
		return *new(XENCryptoMintInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoMintInfo)).(*XENCryptoMintInfo)

	return out0, err

}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Store *StoreSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _Store.Contract.GetUserMint(&_Store.CallOpts)
}

// GetUserMint is a free data retrieval call binding the contract method 0x7010d7a1.
//
// Solidity: function getUserMint() view returns((address,uint256,uint256,uint256,uint256,uint256))
func (_Store *StoreCallerSession) GetUserMint() (XENCryptoMintInfo, error) {
	return _Store.Contract.GetUserMint(&_Store.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Store *StoreCaller) GetUserStake(opts *bind.CallOpts) (XENCryptoStakeInfo, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getUserStake")

	if err != nil {
		return *new(XENCryptoStakeInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(XENCryptoStakeInfo)).(*XENCryptoStakeInfo)

	return out0, err

}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Store *StoreSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _Store.Contract.GetUserStake(&_Store.CallOpts)
}

// GetUserStake is a free data retrieval call binding the contract method 0x16f9c8fd.
//
// Solidity: function getUserStake() view returns((uint256,uint256,uint256,uint256))
func (_Store *StoreCallerSession) GetUserStake() (XENCryptoStakeInfo, error) {
	return _Store.Contract.GetUserStake(&_Store.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Store *StoreCaller) GlobalRank(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "globalRank")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Store *StoreSession) GlobalRank() (*big.Int, error) {
	return _Store.Contract.GlobalRank(&_Store.CallOpts)
}

// GlobalRank is a free data retrieval call binding the contract method 0x1c244082.
//
// Solidity: function globalRank() view returns(uint256)
func (_Store *StoreCallerSession) GlobalRank() (*big.Int, error) {
	return _Store.Contract.GlobalRank(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Store *StoreCaller) TotalXenStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalXenStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Store *StoreSession) TotalXenStaked() (*big.Int, error) {
	return _Store.Contract.TotalXenStaked(&_Store.CallOpts)
}

// TotalXenStaked is a free data retrieval call binding the contract method 0x069612a5.
//
// Solidity: function totalXenStaked() view returns(uint256)
func (_Store *StoreCallerSession) TotalXenStaked() (*big.Int, error) {
	return _Store.Contract.TotalXenStaked(&_Store.CallOpts)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Store *StoreCaller) UserBurns(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "userBurns", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Store *StoreSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.UserBurns(&_Store.CallOpts, arg0)
}

// UserBurns is a free data retrieval call binding the contract method 0xce653d5f.
//
// Solidity: function userBurns(address ) view returns(uint256)
func (_Store *StoreCallerSession) UserBurns(arg0 common.Address) (*big.Int, error) {
	return _Store.Contract.UserBurns(&_Store.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Store *StoreCaller) UserMints(opts *bind.CallOpts, arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "userMints", arg0)

	outstruct := new(struct {
		User       common.Address
		Term       *big.Int
		MaturityTs *big.Int
		Rank       *big.Int
		Amplifier  *big.Int
		EaaRate    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.User = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Term = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Amplifier = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EaaRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Store *StoreSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _Store.Contract.UserMints(&_Store.CallOpts, arg0)
}

// UserMints is a free data retrieval call binding the contract method 0xdf282331.
//
// Solidity: function userMints(address ) view returns(address user, uint256 term, uint256 maturityTs, uint256 rank, uint256 amplifier, uint256 eaaRate)
func (_Store *StoreCallerSession) UserMints(arg0 common.Address) (struct {
	User       common.Address
	Term       *big.Int
	MaturityTs *big.Int
	Rank       *big.Int
	Amplifier  *big.Int
	EaaRate    *big.Int
}, error) {
	return _Store.Contract.UserMints(&_Store.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Store *StoreCaller) UserStakes(opts *bind.CallOpts, arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "userStakes", arg0)

	outstruct := new(struct {
		Term       *big.Int
		MaturityTs *big.Int
		Amount     *big.Int
		Apy        *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Term = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MaturityTs = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Apy = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Store *StoreSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _Store.Contract.UserStakes(&_Store.CallOpts, arg0)
}

// UserStakes is a free data retrieval call binding the contract method 0x8da7ad23.
//
// Solidity: function userStakes(address ) view returns(uint256 term, uint256 maturityTs, uint256 amount, uint256 apy)
func (_Store *StoreCallerSession) UserStakes(arg0 common.Address) (struct {
	Term       *big.Int
	MaturityTs *big.Int
	Amount     *big.Int
	Apy        *big.Int
}, error) {
	return _Store.Contract.UserStakes(&_Store.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Store *StoreTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Store *StoreSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_Store *StoreTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, user, amount)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Store *StoreTransactor) ClaimMintReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "claimMintReward")
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Store *StoreSession) ClaimMintReward() (*types.Transaction, error) {
	return _Store.Contract.ClaimMintReward(&_Store.TransactOpts)
}

// ClaimMintReward is a paid mutator transaction binding the contract method 0x52c7f8dc.
//
// Solidity: function claimMintReward() returns()
func (_Store *StoreTransactorSession) ClaimMintReward() (*types.Transaction, error) {
	return _Store.Contract.ClaimMintReward(&_Store.TransactOpts)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Store *StoreTransactor) ClaimMintRewardAndShare(opts *bind.TransactOpts, other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "claimMintRewardAndShare", other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Store *StoreSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimMintRewardAndShare(&_Store.TransactOpts, other, pct)
}

// ClaimMintRewardAndShare is a paid mutator transaction binding the contract method 0x1c560305.
//
// Solidity: function claimMintRewardAndShare(address other, uint256 pct) returns()
func (_Store *StoreTransactorSession) ClaimMintRewardAndShare(other common.Address, pct *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimMintRewardAndShare(&_Store.TransactOpts, other, pct)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Store *StoreTransactor) ClaimMintRewardAndStake(opts *bind.TransactOpts, pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "claimMintRewardAndStake", pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Store *StoreSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimMintRewardAndStake(&_Store.TransactOpts, pct, term)
}

// ClaimMintRewardAndStake is a paid mutator transaction binding the contract method 0x5bccb4c4.
//
// Solidity: function claimMintRewardAndStake(uint256 pct, uint256 term) returns()
func (_Store *StoreTransactorSession) ClaimMintRewardAndStake(pct *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimMintRewardAndStake(&_Store.TransactOpts, pct, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Store *StoreTransactor) ClaimRank(opts *bind.TransactOpts, term *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "claimRank", term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Store *StoreSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimRank(&_Store.TransactOpts, term)
}

// ClaimRank is a paid mutator transaction binding the contract method 0x9ff054df.
//
// Solidity: function claimRank(uint256 term) returns()
func (_Store *StoreTransactorSession) ClaimRank(term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ClaimRank(&_Store.TransactOpts, term)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Store *StoreTransactor) Stake(opts *bind.TransactOpts, amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "stake", amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Store *StoreSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Stake(&_Store.TransactOpts, amount, term)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 amount, uint256 term) returns()
func (_Store *StoreTransactorSession) Stake(amount *big.Int, term *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Stake(&_Store.TransactOpts, amount, term)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Store *StoreTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Store.Contract.Withdraw(&_Store.TransactOpts)
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

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
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
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
		it.Event = new(StoreApproval)
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
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreMintClaimedIterator is returned from FilterMintClaimed and is used to iterate over the raw logs and unpacked data for MintClaimed events raised by the Store contract.
type StoreMintClaimedIterator struct {
	Event *StoreMintClaimed // Event containing the contract specifics and raw log

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
func (it *StoreMintClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMintClaimed)
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
		it.Event = new(StoreMintClaimed)
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
func (it *StoreMintClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMintClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMintClaimed represents a MintClaimed event raised by the Store contract.
type StoreMintClaimed struct {
	User         common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMintClaimed is a free log retrieval operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Store *StoreFilterer) FilterMintClaimed(opts *bind.FilterOpts, user []common.Address) (*StoreMintClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StoreMintClaimedIterator{contract: _Store.contract, event: "MintClaimed", logs: logs, sub: sub}, nil
}

// WatchMintClaimed is a free log subscription operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Store *StoreFilterer) WatchMintClaimed(opts *bind.WatchOpts, sink chan<- *StoreMintClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "MintClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMintClaimed)
				if err := _Store.contract.UnpackLog(event, "MintClaimed", log); err != nil {
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

// ParseMintClaimed is a log parse operation binding the contract event 0xd74752b13281df13701575f3a507e9b1242e0b5fb040143211c481c1fce573a6.
//
// Solidity: event MintClaimed(address indexed user, uint256 rewardAmount)
func (_Store *StoreFilterer) ParseMintClaimed(log types.Log) (*StoreMintClaimed, error) {
	event := new(StoreMintClaimed)
	if err := _Store.contract.UnpackLog(event, "MintClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRankClaimedIterator is returned from FilterRankClaimed and is used to iterate over the raw logs and unpacked data for RankClaimed events raised by the Store contract.
type StoreRankClaimedIterator struct {
	Event *StoreRankClaimed // Event containing the contract specifics and raw log

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
func (it *StoreRankClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRankClaimed)
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
		it.Event = new(StoreRankClaimed)
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
func (it *StoreRankClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRankClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRankClaimed represents a RankClaimed event raised by the Store contract.
type StoreRankClaimed struct {
	User common.Address
	Term *big.Int
	Rank *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRankClaimed is a free log retrieval operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Store *StoreFilterer) FilterRankClaimed(opts *bind.FilterOpts, user []common.Address) (*StoreRankClaimedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return &StoreRankClaimedIterator{contract: _Store.contract, event: "RankClaimed", logs: logs, sub: sub}, nil
}

// WatchRankClaimed is a free log subscription operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Store *StoreFilterer) WatchRankClaimed(opts *bind.WatchOpts, sink chan<- *StoreRankClaimed, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RankClaimed", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRankClaimed)
				if err := _Store.contract.UnpackLog(event, "RankClaimed", log); err != nil {
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

// ParseRankClaimed is a log parse operation binding the contract event 0xe9149e1b5059238baed02fa659dbf4bd932fbcf760a431330df4d934bc942f37.
//
// Solidity: event RankClaimed(address indexed user, uint256 term, uint256 rank)
func (_Store *StoreFilterer) ParseRankClaimed(log types.Log) (*StoreRankClaimed, error) {
	event := new(StoreRankClaimed)
	if err := _Store.contract.UnpackLog(event, "RankClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Store contract.
type StoreStakedIterator struct {
	Event *StoreStaked // Event containing the contract specifics and raw log

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
func (it *StoreStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreStaked)
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
		it.Event = new(StoreStaked)
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
func (it *StoreStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreStaked represents a Staked event raised by the Store contract.
type StoreStaked struct {
	User   common.Address
	Amount *big.Int
	Term   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Store *StoreFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address) (*StoreStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return &StoreStakedIterator{contract: _Store.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Store *StoreFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StoreStaked, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Staked", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreStaked)
				if err := _Store.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 amount, uint256 term)
func (_Store *StoreFilterer) ParseStaked(log types.Log) (*StoreStaked, error) {
	event := new(StoreStaked)
	if err := _Store.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Store contract.
type StoreWithdrawnIterator struct {
	Event *StoreWithdrawn // Event containing the contract specifics and raw log

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
func (it *StoreWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreWithdrawn)
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
		it.Event = new(StoreWithdrawn)
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
func (it *StoreWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreWithdrawn represents a Withdrawn event raised by the Store contract.
type StoreWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Reward *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Store *StoreFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*StoreWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &StoreWithdrawnIterator{contract: _Store.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Store *StoreFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StoreWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreWithdrawn)
				if err := _Store.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x92ccf450a286a957af52509bc1c9939d1a6a481783e142e41e2499f0bb66ebc6.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount, uint256 reward)
func (_Store *StoreFilterer) ParseWithdrawn(log types.Log) (*StoreWithdrawn, error) {
	event := new(StoreWithdrawn)
	if err := _Store.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
