// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package token

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// HolderABI is the input ABI used to generate the binding from.
const HolderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"addTokenContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SENTINEL_TOKEN_CONTRACTS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\"},{\"name\":\"prev\",\"type\":\"address\"}],\"name\":\"removeTokenContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"init_code_hash\",\"type\":\"bytes32\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"getAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// HolderBin is the compiled bytecode used for deploying new contracts.
const HolderBin = `0x608060405234801561001057600080fd5b50600160008181526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a031916909117905561064d8061005c6000396000f3fe6080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461009a5780632a847f46146100c1578063310ec4a7146100f45780634df4d0391461012d57806383197ef01461015e5780638852671e1461017357806399d3c204146101ae578063eb5a662e146101e7578063fb1669ca1461021a575b005b3480156100a657600080fd5b506100af610244565b60408051918252519081900360200190f35b3480156100cd57600080fd5b50610098600480360360208110156100e457600080fd5b5035600160a060020a031661028c565b34801561010057600080fd5b506100986004803603604081101561011757600080fd5b50600160a060020a03813516906020013561037a565b34801561013957600080fd5b506101426103d6565b60408051600160a060020a039092168252519081900360200190f35b34801561016a57600080fd5b506100986103db565b34801561017f57600080fd5b506100986004803603604081101561019657600080fd5b50600160a060020a0381358116916020013516610417565b3480156101ba57600080fd5b50610098600480360360408110156101d157600080fd5b50600160a060020a0381351690602001356104ea565b3480156101f357600080fd5b506100af6004803603602081101561020a57600080fd5b5035600160a060020a0316610580565b34801561022657600080fd5b506100986004803603602081101561023d57600080fd5b50356105dd565b6000336001141561025457600080fd5b33600090815260208190526040902054600160a060020a0316151561027857600080fd5b503360009081526003602052604090205490565b600154600160a060020a031615156102a357600080fd5b600154600160a060020a031633146102ba57600080fd5b600160a060020a038116158015906102dc5750600160a060020a038116600114155b15156102e757600080fd5b600160a060020a03818116600090815260208190526040902054161561030c57600080fd5b600060208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a039384168084526040842080549590921673ffffffffffffffffffffffffffffffffffffffff199586161790915560019092528054909216179055565b336001141561038857600080fd5b33600090815260208190526040902054600160a060020a031615156103ac57600080fd5b336000908152600460209081526040808320600160a060020a039590951683529390529190912055565b600181565b600154600160a060020a031615156103f257600080fd5b600154600160a060020a0316331461040957600080fd5b600154600160a060020a0316ff5b600154600160a060020a0316151561042e57600080fd5b600154600160a060020a0316331461044557600080fd5b600160a060020a038216158015906104675750600160a060020a038216600114155b151561047257600080fd5b600160a060020a0381811660009081526020819052604090205481169083161461049b57600080fd5b600160a060020a039182166000818152602081905260408082208054948616835290822080549490951673ffffffffffffffffffffffffffffffffffffffff1994851617909455528154169055565b600254604080516c01000000000000000000000000600160a060020a0393841681026020808401919091528685169091026034830152604880830186905283518084039091018152606890920190925280519101201630141561057c5760018054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff19918216179091556002805490911690555b5050565b6000336001141561059057600080fd5b33600090815260208190526040902054600160a060020a031615156105b457600080fd5b50336000908152600460209081526040808320600160a060020a03949094168352929052205490565b33600114156105eb57600080fd5b33600090815260208190526040902054600160a060020a0316151561060f57600080fd5b3360009081526003602052604090205556fea165627a7a72305820e80912be97a8dd8ed62f3e119156048504592f1df14008172078fbd49ecb9ea00029`

// DeployHolder deploys a new Ethereum contract, binding an instance of Holder to it.
func DeployHolder(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Holder, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HolderBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Holder{HolderCaller: HolderCaller{contract: contract}, HolderTransactor: HolderTransactor{contract: contract}, HolderFilterer: HolderFilterer{contract: contract}}, nil
}

// Holder is an auto generated Go binding around an Ethereum contract.
type Holder struct {
	HolderCaller     // Read-only binding to the contract
	HolderTransactor // Write-only binding to the contract
	HolderFilterer   // Log filterer for contract events
}

// HolderCaller is an auto generated read-only Go binding around an Ethereum contract.
type HolderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HolderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HolderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HolderSession struct {
	Contract     *Holder           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HolderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HolderCallerSession struct {
	Contract *HolderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HolderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HolderTransactorSession struct {
	Contract     *HolderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HolderRaw is an auto generated low-level Go binding around an Ethereum contract.
type HolderRaw struct {
	Contract *Holder // Generic contract binding to access the raw methods on
}

// HolderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HolderCallerRaw struct {
	Contract *HolderCaller // Generic read-only contract binding to access the raw methods on
}

// HolderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HolderTransactorRaw struct {
	Contract *HolderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHolder creates a new instance of Holder, bound to a specific deployed contract.
func NewHolder(address common.Address, backend bind.ContractBackend) (*Holder, error) {
	contract, err := bindHolder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Holder{HolderCaller: HolderCaller{contract: contract}, HolderTransactor: HolderTransactor{contract: contract}, HolderFilterer: HolderFilterer{contract: contract}}, nil
}

// NewHolderCaller creates a new read-only instance of Holder, bound to a specific deployed contract.
func NewHolderCaller(address common.Address, caller bind.ContractCaller) (*HolderCaller, error) {
	contract, err := bindHolder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HolderCaller{contract: contract}, nil
}

// NewHolderTransactor creates a new write-only instance of Holder, bound to a specific deployed contract.
func NewHolderTransactor(address common.Address, transactor bind.ContractTransactor) (*HolderTransactor, error) {
	contract, err := bindHolder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HolderTransactor{contract: contract}, nil
}

// NewHolderFilterer creates a new log filterer instance of Holder, bound to a specific deployed contract.
func NewHolderFilterer(address common.Address, filterer bind.ContractFilterer) (*HolderFilterer, error) {
	contract, err := bindHolder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HolderFilterer{contract: contract}, nil
}

// bindHolder binds a generic wrapper to an already deployed contract.
func bindHolder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Holder *HolderRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Holder.Contract.HolderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Holder *HolderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.Contract.HolderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Holder *HolderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Holder.Contract.HolderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Holder *HolderCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Holder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Holder *HolderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Holder *HolderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Holder.Contract.contract.Transact(opts, method, params...)
}

// SENTINELTOKENCONTRACTS is a free data retrieval call binding the contract method 0x4df4d039.
//
// Solidity: function SENTINEL_TOKEN_CONTRACTS() constant returns(address)
func (_Holder *HolderCaller) SENTINELTOKENCONTRACTS(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "SENTINEL_TOKEN_CONTRACTS")
	return *ret0, err
}

// SENTINELTOKENCONTRACTS is a free data retrieval call binding the contract method 0x4df4d039.
//
// Solidity: function SENTINEL_TOKEN_CONTRACTS() constant returns(address)
func (_Holder *HolderSession) SENTINELTOKENCONTRACTS() (common.Address, error) {
	return _Holder.Contract.SENTINELTOKENCONTRACTS(&_Holder.CallOpts)
}

// SENTINELTOKENCONTRACTS is a free data retrieval call binding the contract method 0x4df4d039.
//
// Solidity: function SENTINEL_TOKEN_CONTRACTS() constant returns(address)
func (_Holder *HolderCallerSession) SENTINELTOKENCONTRACTS() (common.Address, error) {
	return _Holder.Contract.SENTINELTOKENCONTRACTS(&_Holder.CallOpts)
}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(_spender address) constant returns(uint256)
func (_Holder *HolderCaller) GetAllowance(opts *bind.CallOpts, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "getAllowance", _spender)
	return *ret0, err
}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(_spender address) constant returns(uint256)
func (_Holder *HolderSession) GetAllowance(_spender common.Address) (*big.Int, error) {
	return _Holder.Contract.GetAllowance(&_Holder.CallOpts, _spender)
}

// GetAllowance is a free data retrieval call binding the contract method 0xeb5a662e.
//
// Solidity: function getAllowance(_spender address) constant returns(uint256)
func (_Holder *HolderCallerSession) GetAllowance(_spender common.Address) (*big.Int, error) {
	return _Holder.Contract.GetAllowance(&_Holder.CallOpts, _spender)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Holder *HolderCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Holder.contract.Call(opts, out, "getBalance")
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Holder *HolderSession) GetBalance() (*big.Int, error) {
	return _Holder.Contract.GetBalance(&_Holder.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() constant returns(uint256)
func (_Holder *HolderCallerSession) GetBalance() (*big.Int, error) {
	return _Holder.Contract.GetBalance(&_Holder.CallOpts)
}

// AddTokenContract is a paid mutator transaction binding the contract method 0x2a847f46.
//
// Solidity: function addTokenContract(tokenContract address) returns()
func (_Holder *HolderTransactor) AddTokenContract(opts *bind.TransactOpts, tokenContract common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "addTokenContract", tokenContract)
}

// AddTokenContract is a paid mutator transaction binding the contract method 0x2a847f46.
//
// Solidity: function addTokenContract(tokenContract address) returns()
func (_Holder *HolderSession) AddTokenContract(tokenContract common.Address) (*types.Transaction, error) {
	return _Holder.Contract.AddTokenContract(&_Holder.TransactOpts, tokenContract)
}

// AddTokenContract is a paid mutator transaction binding the contract method 0x2a847f46.
//
// Solidity: function addTokenContract(tokenContract address) returns()
func (_Holder *HolderTransactorSession) AddTokenContract(tokenContract common.Address) (*types.Transaction, error) {
	return _Holder.Contract.AddTokenContract(&_Holder.TransactOpts, tokenContract)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Holder *HolderTransactor) Destroy(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "destroy")
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Holder *HolderSession) Destroy() (*types.Transaction, error) {
	return _Holder.Contract.Destroy(&_Holder.TransactOpts)
}

// Destroy is a paid mutator transaction binding the contract method 0x83197ef0.
//
// Solidity: function destroy() returns()
func (_Holder *HolderTransactorSession) Destroy() (*types.Transaction, error) {
	return _Holder.Contract.Destroy(&_Holder.TransactOpts)
}

// RemoveTokenContract is a paid mutator transaction binding the contract method 0x8852671e.
//
// Solidity: function removeTokenContract(tokenContract address, prev address) returns()
func (_Holder *HolderTransactor) RemoveTokenContract(opts *bind.TransactOpts, tokenContract common.Address, prev common.Address) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "removeTokenContract", tokenContract, prev)
}

// RemoveTokenContract is a paid mutator transaction binding the contract method 0x8852671e.
//
// Solidity: function removeTokenContract(tokenContract address, prev address) returns()
func (_Holder *HolderSession) RemoveTokenContract(tokenContract common.Address, prev common.Address) (*types.Transaction, error) {
	return _Holder.Contract.RemoveTokenContract(&_Holder.TransactOpts, tokenContract, prev)
}

// RemoveTokenContract is a paid mutator transaction binding the contract method 0x8852671e.
//
// Solidity: function removeTokenContract(tokenContract address, prev address) returns()
func (_Holder *HolderTransactorSession) RemoveTokenContract(tokenContract common.Address, prev common.Address) (*types.Transaction, error) {
	return _Holder.Contract.RemoveTokenContract(&_Holder.TransactOpts, tokenContract, prev)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(_spender address, value uint256) returns()
func (_Holder *HolderTransactor) SetAllowance(opts *bind.TransactOpts, _spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "setAllowance", _spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(_spender address, value uint256) returns()
func (_Holder *HolderSession) SetAllowance(_spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.SetAllowance(&_Holder.TransactOpts, _spender, value)
}

// SetAllowance is a paid mutator transaction binding the contract method 0x310ec4a7.
//
// Solidity: function setAllowance(_spender address, value uint256) returns()
func (_Holder *HolderTransactorSession) SetAllowance(_spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.SetAllowance(&_Holder.TransactOpts, _spender, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xfb1669ca.
//
// Solidity: function setBalance(value uint256) returns()
func (_Holder *HolderTransactor) SetBalance(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "setBalance", value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xfb1669ca.
//
// Solidity: function setBalance(value uint256) returns()
func (_Holder *HolderSession) SetBalance(value *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.SetBalance(&_Holder.TransactOpts, value)
}

// SetBalance is a paid mutator transaction binding the contract method 0xfb1669ca.
//
// Solidity: function setBalance(value uint256) returns()
func (_Holder *HolderTransactorSession) SetBalance(value *big.Int) (*types.Transaction, error) {
	return _Holder.Contract.SetBalance(&_Holder.TransactOpts, value)
}

// SetOwner is a paid mutator transaction binding the contract method 0x99d3c204.
//
// Solidity: function setOwner(_owner address, init_code_hash bytes32) returns()
func (_Holder *HolderTransactor) SetOwner(opts *bind.TransactOpts, _owner common.Address, init_code_hash [32]byte) (*types.Transaction, error) {
	return _Holder.contract.Transact(opts, "setOwner", _owner, init_code_hash)
}

// SetOwner is a paid mutator transaction binding the contract method 0x99d3c204.
//
// Solidity: function setOwner(_owner address, init_code_hash bytes32) returns()
func (_Holder *HolderSession) SetOwner(_owner common.Address, init_code_hash [32]byte) (*types.Transaction, error) {
	return _Holder.Contract.SetOwner(&_Holder.TransactOpts, _owner, init_code_hash)
}

// SetOwner is a paid mutator transaction binding the contract method 0x99d3c204.
//
// Solidity: function setOwner(_owner address, init_code_hash bytes32) returns()
func (_Holder *HolderTransactorSession) SetOwner(_owner common.Address, init_code_hash [32]byte) (*types.Transaction, error) {
	return _Holder.Contract.SetOwner(&_Holder.TransactOpts, _owner, init_code_hash)
}

// HolderFactoryABI is the input ABI used to generate the binding from.
const HolderFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// HolderFactoryBin is the compiled bytecode used for deploying new contracts.
const HolderFactoryBin = `0x608060405234801561001057600080fd5b506101ba806100206000396000f3fe6080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166351e8054c8114610045575b600080fd5b34801561005157600080fd5b506100856004803603602081101561006857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100ae565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6040805180820190915260018082526000602083018190529190828082848783f5905073ffffffffffffffffffffffffffffffffffffffff811615156100f357600080fd5b604080517f99d3c20400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff888116600483015260248201859052915183928316916399d3c20491604480830192600092919082900301818387803b15801561016b57600080fd5b505af115801561017f573d6000803e3d6000fd5b5092999850505050505050505056fea165627a7a72305820d75436cf6d00f79e58bd9a4c829ab939f8a253452b44816f96a6d0fb8675d91e0029`

// DeployHolderFactory deploys a new Ethereum contract, binding an instance of HolderFactory to it.
func DeployHolderFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *HolderFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HolderFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HolderFactory{HolderFactoryCaller: HolderFactoryCaller{contract: contract}, HolderFactoryTransactor: HolderFactoryTransactor{contract: contract}, HolderFactoryFilterer: HolderFactoryFilterer{contract: contract}}, nil
}

// HolderFactory is an auto generated Go binding around an Ethereum contract.
type HolderFactory struct {
	HolderFactoryCaller     // Read-only binding to the contract
	HolderFactoryTransactor // Write-only binding to the contract
	HolderFactoryFilterer   // Log filterer for contract events
}

// HolderFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type HolderFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HolderFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HolderFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HolderFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HolderFactorySession struct {
	Contract     *HolderFactory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HolderFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HolderFactoryCallerSession struct {
	Contract *HolderFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// HolderFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HolderFactoryTransactorSession struct {
	Contract     *HolderFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// HolderFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type HolderFactoryRaw struct {
	Contract *HolderFactory // Generic contract binding to access the raw methods on
}

// HolderFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HolderFactoryCallerRaw struct {
	Contract *HolderFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// HolderFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HolderFactoryTransactorRaw struct {
	Contract *HolderFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHolderFactory creates a new instance of HolderFactory, bound to a specific deployed contract.
func NewHolderFactory(address common.Address, backend bind.ContractBackend) (*HolderFactory, error) {
	contract, err := bindHolderFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HolderFactory{HolderFactoryCaller: HolderFactoryCaller{contract: contract}, HolderFactoryTransactor: HolderFactoryTransactor{contract: contract}, HolderFactoryFilterer: HolderFactoryFilterer{contract: contract}}, nil
}

// NewHolderFactoryCaller creates a new read-only instance of HolderFactory, bound to a specific deployed contract.
func NewHolderFactoryCaller(address common.Address, caller bind.ContractCaller) (*HolderFactoryCaller, error) {
	contract, err := bindHolderFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HolderFactoryCaller{contract: contract}, nil
}

// NewHolderFactoryTransactor creates a new write-only instance of HolderFactory, bound to a specific deployed contract.
func NewHolderFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*HolderFactoryTransactor, error) {
	contract, err := bindHolderFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HolderFactoryTransactor{contract: contract}, nil
}

// NewHolderFactoryFilterer creates a new log filterer instance of HolderFactory, bound to a specific deployed contract.
func NewHolderFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*HolderFactoryFilterer, error) {
	contract, err := bindHolderFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HolderFactoryFilterer{contract: contract}, nil
}

// bindHolderFactory binds a generic wrapper to an already deployed contract.
func bindHolderFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HolderFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HolderFactory *HolderFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HolderFactory.Contract.HolderFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HolderFactory *HolderFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HolderFactory.Contract.HolderFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HolderFactory *HolderFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HolderFactory.Contract.HolderFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HolderFactory *HolderFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HolderFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HolderFactory *HolderFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HolderFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HolderFactory *HolderFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HolderFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_HolderFactory *HolderFactoryTransactor) CreateHolder(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _HolderFactory.contract.Transact(opts, "createHolder", _owner)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_HolderFactory *HolderFactorySession) CreateHolder(_owner common.Address) (*types.Transaction, error) {
	return _HolderFactory.Contract.CreateHolder(&_HolderFactory.TransactOpts, _owner)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_HolderFactory *HolderFactoryTransactorSession) CreateHolder(_owner common.Address) (*types.Transaction, error) {
	return _HolderFactory.Contract.CreateHolder(&_HolderFactory.TransactOpts, _owner)
}

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x608060405234801561001057600080fd5b5060038054600160a060020a03191633179055610a7a806100326000396000f3fe6080604052600436106100745763ffffffff60e060020a600035041663075461728114610076578063095ea7b3146100a757806318160ddd146100f457806323b872dd1461011b57806340c10f191461015e57806370a0823114610197578063a9059cbb146101ca578063dd62ed3e14610203575b005b34801561008257600080fd5b5061008b61023e565b60408051600160a060020a039092168252519081900360200190f35b3480156100b357600080fd5b506100e0600480360360408110156100ca57600080fd5b50600160a060020a03813516906020013561024d565b604080519115158252519081900360200190f35b34801561010057600080fd5b506101096102e0565b60408051918252519081900360200190f35b34801561012757600080fd5b506100e06004803603606081101561013e57600080fd5b50600160a060020a038135811691602081013590911690604001356102e6565b34801561016a57600080fd5b506100e06004803603604081101561018157600080fd5b50600160a060020a038135169060200135610649565b3480156101a357600080fd5b50610109600480360360208110156101ba57600080fd5b5035600160a060020a0316610762565b3480156101d657600080fd5b506100e0600480360360408110156101ed57600080fd5b50600160a060020a038135169060200135610774565b34801561020f57600080fd5b506101096004803603604081101561022657600080fd5b50600160a060020a03813581169160200135166109b5565b600354600160a060020a031681565b600080610259336109d2565b905080600160a060020a031663310ec4a785856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b1580156102be57600080fd5b505af11580156102d2573d6000803e3d6000fd5b506001979650505050505050565b60005481565b6000806102f2856109d2565b905060006102ff856109d2565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561033f57600080fd5b505afa158015610353573d6000803e3d6000fd5b505050506040513d602081101561036957600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156103ca57600080fd5b505afa1580156103de573d6000803e3d6000fd5b505050506040513d60208110156103f457600080fd5b5051604080517feb5a662e0000000000000000000000000000000000000000000000000000000081523360048201529051919250600091600160a060020a0387169163eb5a662e916024808301926020929190829003018186803b15801561045b57600080fd5b505afa15801561046f573d6000803e3d6000fd5b505050506040513d602081101561048557600080fd5b505190508683101561049657600080fd5b8682018211156104a557600080fd5b808711156104b257600080fd5b84600160a060020a031663fb1669ca8885036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156104fd57600080fd5b505af1158015610511573d6000803e3d6000fd5b5050505083600160a060020a031663fb1669ca8884016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561056057600080fd5b505af1158015610574573d6000803e3d6000fd5b5050604080517f310ec4a70000000000000000000000000000000000000000000000000000000081523360048201528a850360248201529051600160a060020a038916935063310ec4a79250604480830192600092919082900301818387803b1580156105e057600080fd5b505af11580156105f4573d6000803e3d6000fd5b5050604080518a81529051600160a060020a03808d1694508d1692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600198975050505050505050565b600354600090600160a060020a0316331461066357600080fd5b600061066e846109d2565b9050600081600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156106ae57600080fd5b505afa1580156106c2573d6000803e3d6000fd5b505050506040513d60208110156106d857600080fd5b505190508381018111156106eb57600080fd5b81600160a060020a031663fb1669ca8583016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561073657600080fd5b505af115801561074a573d6000803e3d6000fd5b50506000805487019055506001935050505092915050565b60016020526000908152604090205481565b600080610780336109d2565b9050600061078d856109d2565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107cd57600080fd5b505afa1580156107e1573d6000803e3d6000fd5b505050506040513d60208110156107f757600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b15801561085857600080fd5b505afa15801561086c573d6000803e3d6000fd5b505050506040513d602081101561088257600080fd5b505190508582101561089357600080fd5b8581018111156108a257600080fd5b83600160a060020a031663fb1669ca8784036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156108ed57600080fd5b505af1158015610901573d6000803e3d6000fd5b5050505082600160a060020a031663fb1669ca8783016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561095057600080fd5b505af1158015610964573d6000803e3d6000fd5b5050604080518981529051600160a060020a038b1693503392507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b600260209081526000928352604080842090915290825290205481565b604080516f09494949000000000000000000000000602080830191909152600160a060020a0384166c010000000000000000000000000260348301526000604880840182905284518085039091018152606890930190935281519101208190630949494990803b838111610a4557600080fd5b5094935050505056fea165627a7a72305820e7e56b0e51a8c9ef913495f70b1ce27f3ca35762b6c61f9df2ec9b0e5b147ad60029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
	TokenFilterer   // Log filterer for contract events
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}, TokenFilterer: TokenFilterer{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// NewTokenFilterer creates a new log filterer instance of Token, bound to a specific deployed contract.
func NewTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFilterer, error) {
	contract, err := bindToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFilterer{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", arg0, arg1)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance( address,  address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, arg0, arg1)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", arg0)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf( address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, arg0)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_Token *TokenCaller) Minter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "minter")
	return *ret0, err
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_Token *TokenSession) Minter() (common.Address, error) {
	return _Token.Contract.Minter(&_Token.CallOpts)
}

// Minter is a free data retrieval call binding the contract method 0x07546172.
//
// Solidity: function minter() constant returns(address)
func (_Token *TokenCallerSession) Minter() (common.Address, error) {
	return _Token.Contract.Minter(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _value uint256) returns(bool)
func (_Token *TokenSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) Mint(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_Token *TokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(success bool)
func (_Token *TokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Token contract.
type TokenTransferIterator struct {
	Event *TokenTransfer // Event containing the contract specifics and raw log

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
func (it *TokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenTransfer)
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
		it.Event = new(TokenTransfer)
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
func (it *TokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenTransfer represents a Transfer event raised by the Token contract.
type TokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TokenTransferIterator{contract: _Token.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_Token *TokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Token.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenTransfer)
				if err := _Token.contract.UnpackLog(event, "Transfer", log); err != nil {
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
