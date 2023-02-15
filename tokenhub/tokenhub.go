// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenhub

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

// TokenhubMetaData contains all meta data concerning the Tokenhub contract.
var TokenhubMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"ParamChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceiveDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"}],\"name\":\"RefundFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"status\",\"type\":\"uint32\"}],\"name\":\"RefundSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"refundAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferInSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"senderAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"relayFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"ackRelayFee\",\"type\":\"uint256\"}],\"name\":\"TransferOutSuccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"UnexpectedPackage\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"APP_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CODE_OK\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CROSS_CHAIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ERROR_FAIL_DECODE\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GOV_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIGHT_CLIENT\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_GAS_FOR_TRANSFER_BNB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PROXY_ADMIN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REWARD_UPPER_LIMIT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOKEN_HUB\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_FAILURE_INSUFFICIENT_BALANCE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_FAILURE_NON_PAYABLE_RECIPIENT\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_FAILURE_UNKNOWN\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_IN_SUCCESS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRANSFER_OUT_CHANNELID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ackRelayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimRelayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"govHub\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleFailAckPackage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"channelId\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"msgBytes\",\"type\":\"bytes\"}],\"name\":\"handleSynPackage\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferOut\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TokenhubABI is the input ABI used to generate the binding from.
// Deprecated: Use TokenhubMetaData.ABI instead.
var TokenhubABI = TokenhubMetaData.ABI

// Tokenhub is an auto generated Go binding around an Ethereum contract.
type Tokenhub struct {
	TokenhubCaller     // Read-only binding to the contract
	TokenhubTransactor // Write-only binding to the contract
	TokenhubFilterer   // Log filterer for contract events
}

// TokenhubCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenhubCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenhubTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenhubTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenhubFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenhubFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenhubSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenhubSession struct {
	Contract     *Tokenhub         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenhubCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenhubCallerSession struct {
	Contract *TokenhubCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TokenhubTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenhubTransactorSession struct {
	Contract     *TokenhubTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenhubRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenhubRaw struct {
	Contract *Tokenhub // Generic contract binding to access the raw methods on
}

// TokenhubCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenhubCallerRaw struct {
	Contract *TokenhubCaller // Generic read-only contract binding to access the raw methods on
}

// TokenhubTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenhubTransactorRaw struct {
	Contract *TokenhubTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenhub creates a new instance of Tokenhub, bound to a specific deployed contract.
func NewTokenhub(address common.Address, backend bind.ContractBackend) (*Tokenhub, error) {
	contract, err := bindTokenhub(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tokenhub{TokenhubCaller: TokenhubCaller{contract: contract}, TokenhubTransactor: TokenhubTransactor{contract: contract}, TokenhubFilterer: TokenhubFilterer{contract: contract}}, nil
}

// NewTokenhubCaller creates a new read-only instance of Tokenhub, bound to a specific deployed contract.
func NewTokenhubCaller(address common.Address, caller bind.ContractCaller) (*TokenhubCaller, error) {
	contract, err := bindTokenhub(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenhubCaller{contract: contract}, nil
}

// NewTokenhubTransactor creates a new write-only instance of Tokenhub, bound to a specific deployed contract.
func NewTokenhubTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenhubTransactor, error) {
	contract, err := bindTokenhub(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenhubTransactor{contract: contract}, nil
}

// NewTokenhubFilterer creates a new log filterer instance of Tokenhub, bound to a specific deployed contract.
func NewTokenhubFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenhubFilterer, error) {
	contract, err := bindTokenhub(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenhubFilterer{contract: contract}, nil
}

// bindTokenhub binds a generic wrapper to an already deployed contract.
func bindTokenhub(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TokenhubMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenhub *TokenhubRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenhub.Contract.TokenhubCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenhub *TokenhubRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenhub.Contract.TokenhubTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenhub *TokenhubRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenhub.Contract.TokenhubTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tokenhub *TokenhubCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tokenhub.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tokenhub *TokenhubTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenhub.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tokenhub *TokenhubTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tokenhub.Contract.contract.Transact(opts, method, params...)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCaller) APPCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "APP_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubSession) APPCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.APPCHANNELID(&_Tokenhub.CallOpts)
}

// APPCHANNELID is a free data retrieval call binding the contract method 0x63ee4ac4.
//
// Solidity: function APP_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) APPCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.APPCHANNELID(&_Tokenhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenhub *TokenhubCaller) CODEOK(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "CODE_OK")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenhub *TokenhubSession) CODEOK() (uint32, error) {
	return _Tokenhub.Contract.CODEOK(&_Tokenhub.CallOpts)
}

// CODEOK is a free data retrieval call binding the contract method 0xab51bb96.
//
// Solidity: function CODE_OK() view returns(uint32)
func (_Tokenhub *TokenhubCallerSession) CODEOK() (uint32, error) {
	return _Tokenhub.Contract.CODEOK(&_Tokenhub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Tokenhub *TokenhubCaller) CROSSCHAIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "CROSS_CHAIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Tokenhub *TokenhubSession) CROSSCHAIN() (common.Address, error) {
	return _Tokenhub.Contract.CROSSCHAIN(&_Tokenhub.CallOpts)
}

// CROSSCHAIN is a free data retrieval call binding the contract method 0x557cf477.
//
// Solidity: function CROSS_CHAIN() view returns(address)
func (_Tokenhub *TokenhubCallerSession) CROSSCHAIN() (common.Address, error) {
	return _Tokenhub.Contract.CROSSCHAIN(&_Tokenhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenhub *TokenhubCaller) ERRORFAILDECODE(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "ERROR_FAIL_DECODE")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenhub *TokenhubSession) ERRORFAILDECODE() (uint32, error) {
	return _Tokenhub.Contract.ERRORFAILDECODE(&_Tokenhub.CallOpts)
}

// ERRORFAILDECODE is a free data retrieval call binding the contract method 0x0bee7a67.
//
// Solidity: function ERROR_FAIL_DECODE() view returns(uint32)
func (_Tokenhub *TokenhubCallerSession) ERRORFAILDECODE() (uint32, error) {
	return _Tokenhub.Contract.ERRORFAILDECODE(&_Tokenhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCaller) GOVCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "GOV_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubSession) GOVCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.GOVCHANNELID(&_Tokenhub.CallOpts)
}

// GOVCHANNELID is a free data retrieval call binding the contract method 0x96713da9.
//
// Solidity: function GOV_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) GOVCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.GOVCHANNELID(&_Tokenhub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Tokenhub *TokenhubCaller) GOVHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "GOV_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Tokenhub *TokenhubSession) GOVHUB() (common.Address, error) {
	return _Tokenhub.Contract.GOVHUB(&_Tokenhub.CallOpts)
}

// GOVHUB is a free data retrieval call binding the contract method 0xa9dae71c.
//
// Solidity: function GOV_HUB() view returns(address)
func (_Tokenhub *TokenhubCallerSession) GOVHUB() (common.Address, error) {
	return _Tokenhub.Contract.GOVHUB(&_Tokenhub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Tokenhub *TokenhubCaller) LIGHTCLIENT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "LIGHT_CLIENT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Tokenhub *TokenhubSession) LIGHTCLIENT() (common.Address, error) {
	return _Tokenhub.Contract.LIGHTCLIENT(&_Tokenhub.CallOpts)
}

// LIGHTCLIENT is a free data retrieval call binding the contract method 0xe613ae00.
//
// Solidity: function LIGHT_CLIENT() view returns(address)
func (_Tokenhub *TokenhubCallerSession) LIGHTCLIENT() (common.Address, error) {
	return _Tokenhub.Contract.LIGHTCLIENT(&_Tokenhub.CallOpts)
}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenhub *TokenhubCaller) MAXGASFORTRANSFERBNB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "MAX_GAS_FOR_TRANSFER_BNB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenhub *TokenhubSession) MAXGASFORTRANSFERBNB() (*big.Int, error) {
	return _Tokenhub.Contract.MAXGASFORTRANSFERBNB(&_Tokenhub.CallOpts)
}

// MAXGASFORTRANSFERBNB is a free data retrieval call binding the contract method 0xfa9e9159.
//
// Solidity: function MAX_GAS_FOR_TRANSFER_BNB() view returns(uint256)
func (_Tokenhub *TokenhubCallerSession) MAXGASFORTRANSFERBNB() (*big.Int, error) {
	return _Tokenhub.Contract.MAXGASFORTRANSFERBNB(&_Tokenhub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Tokenhub *TokenhubCaller) PROXYADMIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "PROXY_ADMIN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Tokenhub *TokenhubSession) PROXYADMIN() (common.Address, error) {
	return _Tokenhub.Contract.PROXYADMIN(&_Tokenhub.CallOpts)
}

// PROXYADMIN is a free data retrieval call binding the contract method 0xed9bc82a.
//
// Solidity: function PROXY_ADMIN() view returns(address)
func (_Tokenhub *TokenhubCallerSession) PROXYADMIN() (common.Address, error) {
	return _Tokenhub.Contract.PROXYADMIN(&_Tokenhub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Tokenhub *TokenhubCaller) RELAYERHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "RELAYER_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Tokenhub *TokenhubSession) RELAYERHUB() (common.Address, error) {
	return _Tokenhub.Contract.RELAYERHUB(&_Tokenhub.CallOpts)
}

// RELAYERHUB is a free data retrieval call binding the contract method 0xb9d86913.
//
// Solidity: function RELAYER_HUB() view returns(address)
func (_Tokenhub *TokenhubCallerSession) RELAYERHUB() (common.Address, error) {
	return _Tokenhub.Contract.RELAYERHUB(&_Tokenhub.CallOpts)
}

// REWARDUPPERLIMIT is a free data retrieval call binding the contract method 0x43a368b9.
//
// Solidity: function REWARD_UPPER_LIMIT() view returns(uint256)
func (_Tokenhub *TokenhubCaller) REWARDUPPERLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "REWARD_UPPER_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// REWARDUPPERLIMIT is a free data retrieval call binding the contract method 0x43a368b9.
//
// Solidity: function REWARD_UPPER_LIMIT() view returns(uint256)
func (_Tokenhub *TokenhubSession) REWARDUPPERLIMIT() (*big.Int, error) {
	return _Tokenhub.Contract.REWARDUPPERLIMIT(&_Tokenhub.CallOpts)
}

// REWARDUPPERLIMIT is a free data retrieval call binding the contract method 0x43a368b9.
//
// Solidity: function REWARD_UPPER_LIMIT() view returns(uint256)
func (_Tokenhub *TokenhubCallerSession) REWARDUPPERLIMIT() (*big.Int, error) {
	return _Tokenhub.Contract.REWARDUPPERLIMIT(&_Tokenhub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Tokenhub *TokenhubCaller) TOKENHUB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TOKEN_HUB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Tokenhub *TokenhubSession) TOKENHUB() (common.Address, error) {
	return _Tokenhub.Contract.TOKENHUB(&_Tokenhub.CallOpts)
}

// TOKENHUB is a free data retrieval call binding the contract method 0x6d3358a1.
//
// Solidity: function TOKEN_HUB() view returns(address)
func (_Tokenhub *TokenhubCallerSession) TOKENHUB() (common.Address, error) {
	return _Tokenhub.Contract.TOKENHUB(&_Tokenhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFERINCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_IN_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINCHANNELID(&_Tokenhub.CallOpts)
}

// TRANSFERINCHANNELID is a free data retrieval call binding the contract method 0x70fd5bad.
//
// Solidity: function TRANSFER_IN_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFERINCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINCHANNELID(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILUREINSUFFICIENTBALANCE is a free data retrieval call binding the contract method 0xa7c9f02d.
//
// Solidity: function TRANSFER_IN_FAILURE_INSUFFICIENT_BALANCE() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFERINFAILUREINSUFFICIENTBALANCE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_IN_FAILURE_INSUFFICIENT_BALANCE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINFAILUREINSUFFICIENTBALANCE is a free data retrieval call binding the contract method 0xa7c9f02d.
//
// Solidity: function TRANSFER_IN_FAILURE_INSUFFICIENT_BALANCE() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFERINFAILUREINSUFFICIENTBALANCE() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILUREINSUFFICIENTBALANCE(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILUREINSUFFICIENTBALANCE is a free data retrieval call binding the contract method 0xa7c9f02d.
//
// Solidity: function TRANSFER_IN_FAILURE_INSUFFICIENT_BALANCE() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFERINFAILUREINSUFFICIENTBALANCE() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILUREINSUFFICIENTBALANCE(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILURENONPAYABLERECIPIENT is a free data retrieval call binding the contract method 0xebf71d53.
//
// Solidity: function TRANSFER_IN_FAILURE_NON_PAYABLE_RECIPIENT() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFERINFAILURENONPAYABLERECIPIENT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_IN_FAILURE_NON_PAYABLE_RECIPIENT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINFAILURENONPAYABLERECIPIENT is a free data retrieval call binding the contract method 0xebf71d53.
//
// Solidity: function TRANSFER_IN_FAILURE_NON_PAYABLE_RECIPIENT() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFERINFAILURENONPAYABLERECIPIENT() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILURENONPAYABLERECIPIENT(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILURENONPAYABLERECIPIENT is a free data retrieval call binding the contract method 0xebf71d53.
//
// Solidity: function TRANSFER_IN_FAILURE_NON_PAYABLE_RECIPIENT() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFERINFAILURENONPAYABLERECIPIENT() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILURENONPAYABLERECIPIENT(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILUREUNKNOWN is a free data retrieval call binding the contract method 0xf0148472.
//
// Solidity: function TRANSFER_IN_FAILURE_UNKNOWN() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFERINFAILUREUNKNOWN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_IN_FAILURE_UNKNOWN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINFAILUREUNKNOWN is a free data retrieval call binding the contract method 0xf0148472.
//
// Solidity: function TRANSFER_IN_FAILURE_UNKNOWN() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFERINFAILUREUNKNOWN() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILUREUNKNOWN(&_Tokenhub.CallOpts)
}

// TRANSFERINFAILUREUNKNOWN is a free data retrieval call binding the contract method 0xf0148472.
//
// Solidity: function TRANSFER_IN_FAILURE_UNKNOWN() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFERINFAILUREUNKNOWN() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINFAILUREUNKNOWN(&_Tokenhub.CallOpts)
}

// TRANSFERINSUCCESS is a free data retrieval call binding the contract method 0xa496fba2.
//
// Solidity: function TRANSFER_IN_SUCCESS() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFERINSUCCESS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_IN_SUCCESS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFERINSUCCESS is a free data retrieval call binding the contract method 0xa496fba2.
//
// Solidity: function TRANSFER_IN_SUCCESS() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFERINSUCCESS() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINSUCCESS(&_Tokenhub.CallOpts)
}

// TRANSFERINSUCCESS is a free data retrieval call binding the contract method 0xa496fba2.
//
// Solidity: function TRANSFER_IN_SUCCESS() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFERINSUCCESS() (uint8, error) {
	return _Tokenhub.Contract.TRANSFERINSUCCESS(&_Tokenhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCaller) TRANSFEROUTCHANNELID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "TRANSFER_OUT_CHANNELID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.TRANSFEROUTCHANNELID(&_Tokenhub.CallOpts)
}

// TRANSFEROUTCHANNELID is a free data retrieval call binding the contract method 0xfc3e5908.
//
// Solidity: function TRANSFER_OUT_CHANNELID() view returns(uint8)
func (_Tokenhub *TokenhubCallerSession) TRANSFEROUTCHANNELID() (uint8, error) {
	return _Tokenhub.Contract.TRANSFEROUTCHANNELID(&_Tokenhub.CallOpts)
}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Tokenhub *TokenhubCaller) AckRelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "ackRelayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Tokenhub *TokenhubSession) AckRelayFee() (*big.Int, error) {
	return _Tokenhub.Contract.AckRelayFee(&_Tokenhub.CallOpts)
}

// AckRelayFee is a free data retrieval call binding the contract method 0x6ab31754.
//
// Solidity: function ackRelayFee() view returns(uint256)
func (_Tokenhub *TokenhubCallerSession) AckRelayFee() (*big.Int, error) {
	return _Tokenhub.Contract.AckRelayFee(&_Tokenhub.CallOpts)
}

// GovHub is a free data retrieval call binding the contract method 0x3f9c44ec.
//
// Solidity: function govHub() view returns(address)
func (_Tokenhub *TokenhubCaller) GovHub(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "govHub")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GovHub is a free data retrieval call binding the contract method 0x3f9c44ec.
//
// Solidity: function govHub() view returns(address)
func (_Tokenhub *TokenhubSession) GovHub() (common.Address, error) {
	return _Tokenhub.Contract.GovHub(&_Tokenhub.CallOpts)
}

// GovHub is a free data retrieval call binding the contract method 0x3f9c44ec.
//
// Solidity: function govHub() view returns(address)
func (_Tokenhub *TokenhubCallerSession) GovHub() (common.Address, error) {
	return _Tokenhub.Contract.GovHub(&_Tokenhub.CallOpts)
}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Tokenhub *TokenhubCaller) RelayFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tokenhub.contract.Call(opts, &out, "relayFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Tokenhub *TokenhubSession) RelayFee() (*big.Int, error) {
	return _Tokenhub.Contract.RelayFee(&_Tokenhub.CallOpts)
}

// RelayFee is a free data retrieval call binding the contract method 0x71d30863.
//
// Solidity: function relayFee() view returns(uint256)
func (_Tokenhub *TokenhubCallerSession) RelayFee() (*big.Int, error) {
	return _Tokenhub.Contract.RelayFee(&_Tokenhub.CallOpts)
}

// ClaimRelayFee is a paid mutator transaction binding the contract method 0x9bfa7645.
//
// Solidity: function claimRelayFee(uint256 amount) returns(uint256)
func (_Tokenhub *TokenhubTransactor) ClaimRelayFee(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "claimRelayFee", amount)
}

// ClaimRelayFee is a paid mutator transaction binding the contract method 0x9bfa7645.
//
// Solidity: function claimRelayFee(uint256 amount) returns(uint256)
func (_Tokenhub *TokenhubSession) ClaimRelayFee(amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.Contract.ClaimRelayFee(&_Tokenhub.TransactOpts, amount)
}

// ClaimRelayFee is a paid mutator transaction binding the contract method 0x9bfa7645.
//
// Solidity: function claimRelayFee(uint256 amount) returns(uint256)
func (_Tokenhub *TokenhubTransactorSession) ClaimRelayFee(amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.Contract.ClaimRelayFee(&_Tokenhub.TransactOpts, amount)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubTransactor) HandleAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "handleAckPackage", channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleAckPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// HandleAckPackage is a paid mutator transaction binding the contract method 0x831d65d1.
//
// Solidity: function handleAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubTransactorSession) HandleAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleAckPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubTransactor) HandleFailAckPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "handleFailAckPackage", channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleFailAckPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// HandleFailAckPackage is a paid mutator transaction binding the contract method 0xc8509d81.
//
// Solidity: function handleFailAckPackage(uint8 channelId, bytes msgBytes) returns()
func (_Tokenhub *TokenhubTransactorSession) HandleFailAckPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleFailAckPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenhub *TokenhubTransactor) HandleSynPackage(opts *bind.TransactOpts, channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "handleSynPackage", channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenhub *TokenhubSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleSynPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// HandleSynPackage is a paid mutator transaction binding the contract method 0x1182b875.
//
// Solidity: function handleSynPackage(uint8 channelId, bytes msgBytes) returns(bytes)
func (_Tokenhub *TokenhubTransactorSession) HandleSynPackage(channelId uint8, msgBytes []byte) (*types.Transaction, error) {
	return _Tokenhub.Contract.HandleSynPackage(&_Tokenhub.TransactOpts, channelId, msgBytes)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenhub *TokenhubTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenhub *TokenhubSession) Initialize() (*types.Transaction, error) {
	return _Tokenhub.Contract.Initialize(&_Tokenhub.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Tokenhub *TokenhubTransactorSession) Initialize() (*types.Transaction, error) {
	return _Tokenhub.Contract.Initialize(&_Tokenhub.TransactOpts)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address recipient, uint256 amount) payable returns(bool)
func (_Tokenhub *TokenhubTransactor) TransferOut(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.contract.Transact(opts, "transferOut", recipient, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address recipient, uint256 amount) payable returns(bool)
func (_Tokenhub *TokenhubSession) TransferOut(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.Contract.TransferOut(&_Tokenhub.TransactOpts, recipient, amount)
}

// TransferOut is a paid mutator transaction binding the contract method 0x76890c58.
//
// Solidity: function transferOut(address recipient, uint256 amount) payable returns(bool)
func (_Tokenhub *TokenhubTransactorSession) TransferOut(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tokenhub.Contract.TransferOut(&_Tokenhub.TransactOpts, recipient, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenhub *TokenhubTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tokenhub.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenhub *TokenhubSession) Receive() (*types.Transaction, error) {
	return _Tokenhub.Contract.Receive(&_Tokenhub.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tokenhub *TokenhubTransactorSession) Receive() (*types.Transaction, error) {
	return _Tokenhub.Contract.Receive(&_Tokenhub.TransactOpts)
}

// TokenhubInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Tokenhub contract.
type TokenhubInitializedIterator struct {
	Event *TokenhubInitialized // Event containing the contract specifics and raw log

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
func (it *TokenhubInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubInitialized)
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
		it.Event = new(TokenhubInitialized)
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
func (it *TokenhubInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubInitialized represents a Initialized event raised by the Tokenhub contract.
type TokenhubInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenhub *TokenhubFilterer) FilterInitialized(opts *bind.FilterOpts) (*TokenhubInitializedIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TokenhubInitializedIterator{contract: _Tokenhub.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenhub *TokenhubFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TokenhubInitialized) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubInitialized)
				if err := _Tokenhub.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Tokenhub *TokenhubFilterer) ParseInitialized(log types.Log) (*TokenhubInitialized, error) {
	event := new(TokenhubInitialized)
	if err := _Tokenhub.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubParamChangeIterator is returned from FilterParamChange and is used to iterate over the raw logs and unpacked data for ParamChange events raised by the Tokenhub contract.
type TokenhubParamChangeIterator struct {
	Event *TokenhubParamChange // Event containing the contract specifics and raw log

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
func (it *TokenhubParamChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubParamChange)
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
		it.Event = new(TokenhubParamChange)
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
func (it *TokenhubParamChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubParamChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubParamChange represents a ParamChange event raised by the Tokenhub contract.
type TokenhubParamChange struct {
	Key   string
	Value []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterParamChange is a free log retrieval operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenhub *TokenhubFilterer) FilterParamChange(opts *bind.FilterOpts) (*TokenhubParamChangeIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return &TokenhubParamChangeIterator{contract: _Tokenhub.contract, event: "ParamChange", logs: logs, sub: sub}, nil
}

// WatchParamChange is a free log subscription operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenhub *TokenhubFilterer) WatchParamChange(opts *bind.WatchOpts, sink chan<- *TokenhubParamChange) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "ParamChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubParamChange)
				if err := _Tokenhub.contract.UnpackLog(event, "ParamChange", log); err != nil {
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

// ParseParamChange is a log parse operation binding the contract event 0xf1ce9b2cbf50eeb05769a29e2543fd350cab46894a7dd9978a12d534bb20e633.
//
// Solidity: event ParamChange(string key, bytes value)
func (_Tokenhub *TokenhubFilterer) ParseParamChange(log types.Log) (*TokenhubParamChange, error) {
	event := new(TokenhubParamChange)
	if err := _Tokenhub.contract.UnpackLog(event, "ParamChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubReceiveDepositIterator is returned from FilterReceiveDeposit and is used to iterate over the raw logs and unpacked data for ReceiveDeposit events raised by the Tokenhub contract.
type TokenhubReceiveDepositIterator struct {
	Event *TokenhubReceiveDeposit // Event containing the contract specifics and raw log

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
func (it *TokenhubReceiveDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubReceiveDeposit)
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
		it.Event = new(TokenhubReceiveDeposit)
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
func (it *TokenhubReceiveDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubReceiveDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubReceiveDeposit represents a ReceiveDeposit event raised by the Tokenhub contract.
type TokenhubReceiveDeposit struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceiveDeposit is a free log retrieval operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address from, uint256 amount)
func (_Tokenhub *TokenhubFilterer) FilterReceiveDeposit(opts *bind.FilterOpts) (*TokenhubReceiveDepositIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "ReceiveDeposit")
	if err != nil {
		return nil, err
	}
	return &TokenhubReceiveDepositIterator{contract: _Tokenhub.contract, event: "ReceiveDeposit", logs: logs, sub: sub}, nil
}

// WatchReceiveDeposit is a free log subscription operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address from, uint256 amount)
func (_Tokenhub *TokenhubFilterer) WatchReceiveDeposit(opts *bind.WatchOpts, sink chan<- *TokenhubReceiveDeposit) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "ReceiveDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubReceiveDeposit)
				if err := _Tokenhub.contract.UnpackLog(event, "ReceiveDeposit", log); err != nil {
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

// ParseReceiveDeposit is a log parse operation binding the contract event 0xd4f7d34af79a91579ffbb26e18ffb9866c734383ca40131b18e2ca4db8f6649c.
//
// Solidity: event ReceiveDeposit(address from, uint256 amount)
func (_Tokenhub *TokenhubFilterer) ParseReceiveDeposit(log types.Log) (*TokenhubReceiveDeposit, error) {
	event := new(TokenhubReceiveDeposit)
	if err := _Tokenhub.contract.UnpackLog(event, "ReceiveDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubRefundFailureIterator is returned from FilterRefundFailure and is used to iterate over the raw logs and unpacked data for RefundFailure events raised by the Tokenhub contract.
type TokenhubRefundFailureIterator struct {
	Event *TokenhubRefundFailure // Event containing the contract specifics and raw log

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
func (it *TokenhubRefundFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubRefundFailure)
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
		it.Event = new(TokenhubRefundFailure)
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
func (it *TokenhubRefundFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubRefundFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubRefundFailure represents a RefundFailure event raised by the Tokenhub contract.
type TokenhubRefundFailure struct {
	RefundAddr common.Address
	Amount     *big.Int
	Status     uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRefundFailure is a free log retrieval operation binding the contract event 0x71db90364b0bc6bf11dca109991a1372160441d45337fe10262f37a33111b1de.
//
// Solidity: event RefundFailure(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) FilterRefundFailure(opts *bind.FilterOpts) (*TokenhubRefundFailureIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "RefundFailure")
	if err != nil {
		return nil, err
	}
	return &TokenhubRefundFailureIterator{contract: _Tokenhub.contract, event: "RefundFailure", logs: logs, sub: sub}, nil
}

// WatchRefundFailure is a free log subscription operation binding the contract event 0x71db90364b0bc6bf11dca109991a1372160441d45337fe10262f37a33111b1de.
//
// Solidity: event RefundFailure(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) WatchRefundFailure(opts *bind.WatchOpts, sink chan<- *TokenhubRefundFailure) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "RefundFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubRefundFailure)
				if err := _Tokenhub.contract.UnpackLog(event, "RefundFailure", log); err != nil {
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

// ParseRefundFailure is a log parse operation binding the contract event 0x71db90364b0bc6bf11dca109991a1372160441d45337fe10262f37a33111b1de.
//
// Solidity: event RefundFailure(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) ParseRefundFailure(log types.Log) (*TokenhubRefundFailure, error) {
	event := new(TokenhubRefundFailure)
	if err := _Tokenhub.contract.UnpackLog(event, "RefundFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubRefundSuccessIterator is returned from FilterRefundSuccess and is used to iterate over the raw logs and unpacked data for RefundSuccess events raised by the Tokenhub contract.
type TokenhubRefundSuccessIterator struct {
	Event *TokenhubRefundSuccess // Event containing the contract specifics and raw log

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
func (it *TokenhubRefundSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubRefundSuccess)
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
		it.Event = new(TokenhubRefundSuccess)
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
func (it *TokenhubRefundSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubRefundSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubRefundSuccess represents a RefundSuccess event raised by the Tokenhub contract.
type TokenhubRefundSuccess struct {
	RefundAddr common.Address
	Amount     *big.Int
	Status     uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRefundSuccess is a free log retrieval operation binding the contract event 0x83c8bea6a6f19d0f5bfda8365c0a0df62e317763865535f2ab651c00e715e41c.
//
// Solidity: event RefundSuccess(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) FilterRefundSuccess(opts *bind.FilterOpts) (*TokenhubRefundSuccessIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "RefundSuccess")
	if err != nil {
		return nil, err
	}
	return &TokenhubRefundSuccessIterator{contract: _Tokenhub.contract, event: "RefundSuccess", logs: logs, sub: sub}, nil
}

// WatchRefundSuccess is a free log subscription operation binding the contract event 0x83c8bea6a6f19d0f5bfda8365c0a0df62e317763865535f2ab651c00e715e41c.
//
// Solidity: event RefundSuccess(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) WatchRefundSuccess(opts *bind.WatchOpts, sink chan<- *TokenhubRefundSuccess) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "RefundSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubRefundSuccess)
				if err := _Tokenhub.contract.UnpackLog(event, "RefundSuccess", log); err != nil {
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

// ParseRefundSuccess is a log parse operation binding the contract event 0x83c8bea6a6f19d0f5bfda8365c0a0df62e317763865535f2ab651c00e715e41c.
//
// Solidity: event RefundSuccess(address refundAddr, uint256 amount, uint32 status)
func (_Tokenhub *TokenhubFilterer) ParseRefundSuccess(log types.Log) (*TokenhubRefundSuccess, error) {
	event := new(TokenhubRefundSuccess)
	if err := _Tokenhub.contract.UnpackLog(event, "RefundSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubRewardToIterator is returned from FilterRewardTo and is used to iterate over the raw logs and unpacked data for RewardTo events raised by the Tokenhub contract.
type TokenhubRewardToIterator struct {
	Event *TokenhubRewardTo // Event containing the contract specifics and raw log

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
func (it *TokenhubRewardToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubRewardTo)
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
		it.Event = new(TokenhubRewardTo)
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
func (it *TokenhubRewardToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubRewardToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubRewardTo represents a RewardTo event raised by the Tokenhub contract.
type TokenhubRewardTo struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardTo is a free log retrieval operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address to, uint256 amount)
func (_Tokenhub *TokenhubFilterer) FilterRewardTo(opts *bind.FilterOpts) (*TokenhubRewardToIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "RewardTo")
	if err != nil {
		return nil, err
	}
	return &TokenhubRewardToIterator{contract: _Tokenhub.contract, event: "RewardTo", logs: logs, sub: sub}, nil
}

// WatchRewardTo is a free log subscription operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address to, uint256 amount)
func (_Tokenhub *TokenhubFilterer) WatchRewardTo(opts *bind.WatchOpts, sink chan<- *TokenhubRewardTo) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "RewardTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubRewardTo)
				if err := _Tokenhub.contract.UnpackLog(event, "RewardTo", log); err != nil {
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

// ParseRewardTo is a log parse operation binding the contract event 0xa641bcd8a48e29cb86bb641e1ad9cb6642ccd0227d91ec198044193b7f8416b7.
//
// Solidity: event RewardTo(address to, uint256 amount)
func (_Tokenhub *TokenhubFilterer) ParseRewardTo(log types.Log) (*TokenhubRewardTo, error) {
	event := new(TokenhubRewardTo)
	if err := _Tokenhub.contract.UnpackLog(event, "RewardTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubTransferInSuccessIterator is returned from FilterTransferInSuccess and is used to iterate over the raw logs and unpacked data for TransferInSuccess events raised by the Tokenhub contract.
type TokenhubTransferInSuccessIterator struct {
	Event *TokenhubTransferInSuccess // Event containing the contract specifics and raw log

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
func (it *TokenhubTransferInSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubTransferInSuccess)
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
		it.Event = new(TokenhubTransferInSuccess)
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
func (it *TokenhubTransferInSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubTransferInSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubTransferInSuccess represents a TransferInSuccess event raised by the Tokenhub contract.
type TokenhubTransferInSuccess struct {
	RefundAddr common.Address
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransferInSuccess is a free log retrieval operation binding the contract event 0x03a8880d7007ddfd8cf67f90f0a204392ac88639f784464160381464c7072c68.
//
// Solidity: event TransferInSuccess(address refundAddr, uint256 amount)
func (_Tokenhub *TokenhubFilterer) FilterTransferInSuccess(opts *bind.FilterOpts) (*TokenhubTransferInSuccessIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "TransferInSuccess")
	if err != nil {
		return nil, err
	}
	return &TokenhubTransferInSuccessIterator{contract: _Tokenhub.contract, event: "TransferInSuccess", logs: logs, sub: sub}, nil
}

// WatchTransferInSuccess is a free log subscription operation binding the contract event 0x03a8880d7007ddfd8cf67f90f0a204392ac88639f784464160381464c7072c68.
//
// Solidity: event TransferInSuccess(address refundAddr, uint256 amount)
func (_Tokenhub *TokenhubFilterer) WatchTransferInSuccess(opts *bind.WatchOpts, sink chan<- *TokenhubTransferInSuccess) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "TransferInSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubTransferInSuccess)
				if err := _Tokenhub.contract.UnpackLog(event, "TransferInSuccess", log); err != nil {
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

// ParseTransferInSuccess is a log parse operation binding the contract event 0x03a8880d7007ddfd8cf67f90f0a204392ac88639f784464160381464c7072c68.
//
// Solidity: event TransferInSuccess(address refundAddr, uint256 amount)
func (_Tokenhub *TokenhubFilterer) ParseTransferInSuccess(log types.Log) (*TokenhubTransferInSuccess, error) {
	event := new(TokenhubTransferInSuccess)
	if err := _Tokenhub.contract.UnpackLog(event, "TransferInSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubTransferOutSuccessIterator is returned from FilterTransferOutSuccess and is used to iterate over the raw logs and unpacked data for TransferOutSuccess events raised by the Tokenhub contract.
type TokenhubTransferOutSuccessIterator struct {
	Event *TokenhubTransferOutSuccess // Event containing the contract specifics and raw log

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
func (it *TokenhubTransferOutSuccessIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubTransferOutSuccess)
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
		it.Event = new(TokenhubTransferOutSuccess)
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
func (it *TokenhubTransferOutSuccessIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubTransferOutSuccessIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubTransferOutSuccess represents a TransferOutSuccess event raised by the Tokenhub contract.
type TokenhubTransferOutSuccess struct {
	SenderAddr  common.Address
	Amount      *big.Int
	RelayFee    *big.Int
	AckRelayFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferOutSuccess is a free log retrieval operation binding the contract event 0xc293bf5e3f32f34a15a1deaf52ef8b4ebfafa68a817c72a5131a2084b78c3a29.
//
// Solidity: event TransferOutSuccess(address senderAddr, uint256 amount, uint256 relayFee, uint256 ackRelayFee)
func (_Tokenhub *TokenhubFilterer) FilterTransferOutSuccess(opts *bind.FilterOpts) (*TokenhubTransferOutSuccessIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "TransferOutSuccess")
	if err != nil {
		return nil, err
	}
	return &TokenhubTransferOutSuccessIterator{contract: _Tokenhub.contract, event: "TransferOutSuccess", logs: logs, sub: sub}, nil
}

// WatchTransferOutSuccess is a free log subscription operation binding the contract event 0xc293bf5e3f32f34a15a1deaf52ef8b4ebfafa68a817c72a5131a2084b78c3a29.
//
// Solidity: event TransferOutSuccess(address senderAddr, uint256 amount, uint256 relayFee, uint256 ackRelayFee)
func (_Tokenhub *TokenhubFilterer) WatchTransferOutSuccess(opts *bind.WatchOpts, sink chan<- *TokenhubTransferOutSuccess) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "TransferOutSuccess")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubTransferOutSuccess)
				if err := _Tokenhub.contract.UnpackLog(event, "TransferOutSuccess", log); err != nil {
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

// ParseTransferOutSuccess is a log parse operation binding the contract event 0xc293bf5e3f32f34a15a1deaf52ef8b4ebfafa68a817c72a5131a2084b78c3a29.
//
// Solidity: event TransferOutSuccess(address senderAddr, uint256 amount, uint256 relayFee, uint256 ackRelayFee)
func (_Tokenhub *TokenhubFilterer) ParseTransferOutSuccess(log types.Log) (*TokenhubTransferOutSuccess, error) {
	event := new(TokenhubTransferOutSuccess)
	if err := _Tokenhub.contract.UnpackLog(event, "TransferOutSuccess", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TokenhubUnexpectedPackageIterator is returned from FilterUnexpectedPackage and is used to iterate over the raw logs and unpacked data for UnexpectedPackage events raised by the Tokenhub contract.
type TokenhubUnexpectedPackageIterator struct {
	Event *TokenhubUnexpectedPackage // Event containing the contract specifics and raw log

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
func (it *TokenhubUnexpectedPackageIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenhubUnexpectedPackage)
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
		it.Event = new(TokenhubUnexpectedPackage)
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
func (it *TokenhubUnexpectedPackageIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenhubUnexpectedPackageIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenhubUnexpectedPackage represents a UnexpectedPackage event raised by the Tokenhub contract.
type TokenhubUnexpectedPackage struct {
	ChannelId uint8
	MsgBytes  []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnexpectedPackage is a free log retrieval operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenhub *TokenhubFilterer) FilterUnexpectedPackage(opts *bind.FilterOpts) (*TokenhubUnexpectedPackageIterator, error) {

	logs, sub, err := _Tokenhub.contract.FilterLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return &TokenhubUnexpectedPackageIterator{contract: _Tokenhub.contract, event: "UnexpectedPackage", logs: logs, sub: sub}, nil
}

// WatchUnexpectedPackage is a free log subscription operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenhub *TokenhubFilterer) WatchUnexpectedPackage(opts *bind.WatchOpts, sink chan<- *TokenhubUnexpectedPackage) (event.Subscription, error) {

	logs, sub, err := _Tokenhub.contract.WatchLogs(opts, "UnexpectedPackage")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenhubUnexpectedPackage)
				if err := _Tokenhub.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
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

// ParseUnexpectedPackage is a log parse operation binding the contract event 0xaa5ba621c8b3d7d05bb9e51a7506108251d4d5dbe542ca66fc7bb52aacb02b65.
//
// Solidity: event UnexpectedPackage(uint8 channelId, bytes msgBytes)
func (_Tokenhub *TokenhubFilterer) ParseUnexpectedPackage(log types.Log) (*TokenhubUnexpectedPackage, error) {
	event := new(TokenhubUnexpectedPackage)
	if err := _Tokenhub.contract.UnpackLog(event, "UnexpectedPackage", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
