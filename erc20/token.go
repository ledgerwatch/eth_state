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
const HolderBin = `0x608060405234801561001057600080fd5b50600160008181526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a03191690911790556106718061005c6000396000f3fe6080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461009a5780632a847f46146100c1578063310ec4a7146100f45780634df4d0391461012d57806383197ef01461015e5780638852671e1461017357806399d3c204146101ae578063eb5a662e146101e7578063fb1669ca1461021a575b005b3480156100a657600080fd5b506100af610244565b60408051918252519081900360200190f35b3480156100cd57600080fd5b50610098600480360360208110156100e457600080fd5b5035600160a060020a031661028c565b34801561010057600080fd5b506100986004803603604081101561011757600080fd5b50600160a060020a03813516906020013561037a565b34801561013957600080fd5b506101426103d6565b60408051600160a060020a039092168252519081900360200190f35b34801561016a57600080fd5b506100986103db565b34801561017f57600080fd5b506100986004803603604081101561019657600080fd5b50600160a060020a0381358116916020013516610417565b3480156101ba57600080fd5b50610098600480360360408110156101d157600080fd5b50600160a060020a0381351690602001356104ea565b3480156101f357600080fd5b506100af6004803603602081101561020a57600080fd5b5035600160a060020a03166105a4565b34801561022657600080fd5b506100986004803603602081101561023d57600080fd5b5035610601565b6000336001141561025457600080fd5b33600090815260208190526040902054600160a060020a0316151561027857600080fd5b503360009081526003602052604090205490565b600154600160a060020a031615156102a357600080fd5b600154600160a060020a031633146102ba57600080fd5b600160a060020a038116158015906102dc5750600160a060020a038116600114155b15156102e757600080fd5b600160a060020a03818116600090815260208190526040902054161561030c57600080fd5b600060208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a039384168084526040842080549590921673ffffffffffffffffffffffffffffffffffffffff199586161790915560019092528054909216179055565b336001141561038857600080fd5b33600090815260208190526040902054600160a060020a031615156103ac57600080fd5b336000908152600460209081526040808320600160a060020a039590951683529390529190912055565b600181565b600154600160a060020a031615156103f257600080fd5b600154600160a060020a0316331461040957600080fd5b600154600160a060020a0316ff5b600154600160a060020a0316151561042e57600080fd5b600154600160a060020a0316331461044557600080fd5b600160a060020a038216158015906104675750600160a060020a038216600114155b151561047257600080fd5b600160a060020a0381811660009081526020819052604090205481169083161461049b57600080fd5b600160a060020a039182166000818152602081905260408082208054948616835290822080549490951673ffffffffffffffffffffffffffffffffffffffff1994851617909455528154169055565b600254604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526c01000000000000000000000000600160a060020a039485168102602184015286851602603583015260498083018690528351808403909101815260699092019092528051910120163014156105a05760018054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff19918216179091556002805490911690555b5050565b600033600114156105b457600080fd5b33600090815260208190526040902054600160a060020a031615156105d857600080fd5b50336000908152600460209081526040808320600160a060020a03949094168352929052205490565b336001141561060f57600080fd5b33600090815260208190526040902054600160a060020a0316151561063357600080fd5b3360009081526003602052604090205556fea165627a7a723058205bf590cf8435c063c1ffc96525284b74541f786853d65c82109cae1dc93ff6ee0029`

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
const HolderFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"HolderCreated\",\"type\":\"event\"}]"

// HolderFactoryBin is the compiled bytecode used for deploying new contracts.
const HolderFactoryBin = `0x608060405234801561001057600080fd5b506113f8806100206000396000f3fe6080604052600436106100405763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166351e8054c8114610045575b600080fd5b34801561005157600080fd5b506100856004803603602081101561006857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100ae565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b60006060611200604051908101604052806111d981526020016101f46111d9913980516020820181902091925090600082848783f5905073ffffffffffffffffffffffffffffffffffffffff8116151561010757600080fd5b604080517f99d3c20400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff888116600483015260248201859052915183928316916399d3c20491604480830192600092919082900301818387803b15801561017f57600080fd5b505af1158015610193573d6000803e3d6000fd5b50506040805173ffffffffffffffffffffffffffffffffffffffff80861682528b16602082015281517f4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a17160559450908190039091019150a1969550505050505056fe608060405234801561001057600080fd5b506001600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061113a8061009f6000396000f3fe608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312065fe01461009b5780632a847f46146100c6578063310ec4a7146101175780634df4d0391461017257806383197ef0146101c95780638852671e146101e057806399d3c20414610251578063eb5a662e146102ac578063fb1669ca14610311575b005b3480156100a757600080fd5b506100b061034c565b6040518082815260200191505060405180910390f35b3480156100d257600080fd5b50610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610469565b005b34801561012357600080fd5b506101706004803603604081101561013a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061078e565b005b34801561017e57600080fd5b506101876108e9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101d557600080fd5b506101de6108ee565b005b3480156101ec57600080fd5b5061024f6004803603604081101561020357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109e3565b005b34801561025d57600080fd5b506102aa6004803603604081101561027457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d07565b005b3480156102b857600080fd5b506102fb600480360360208110156102cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e95565b6040518082815260200191505060405180910390f35b34801561031d57600080fd5b5061034a6004803603602081101561033457600080fd5b8101908080359060200190929190505050610ff1565b005b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561038a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561042457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156104c757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561058d5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b151561059857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561063157600080fd5b600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515156107ca57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561086457600080fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b600181565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561094c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109a857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610a4157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610b075750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1515610b1257600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610baa57600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b3073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019350505050604051602081830303815290604052805190602001206001900473ffffffffffffffffffffffffffffffffffffffff161415610e915781600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050565b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610ed357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610f6d57600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561102d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156110c757600080fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505056fea165627a7a72305820b07d304542735e93c801d4e3e1fe1cf63b58ea0c3bd0fb00642de67898f3480b0029a165627a7a723058208cd64d7e9b1528402d2e054abe9f10e9ae9864281515aa41d3709e275b52da8f0029`

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

// HolderFactoryHolderCreatedIterator is returned from FilterHolderCreated and is used to iterate over the raw logs and unpacked data for HolderCreated events raised by the HolderFactory contract.
type HolderFactoryHolderCreatedIterator struct {
	Event *HolderFactoryHolderCreated // Event containing the contract specifics and raw log

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
func (it *HolderFactoryHolderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HolderFactoryHolderCreated)
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
		it.Event = new(HolderFactoryHolderCreated)
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
func (it *HolderFactoryHolderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HolderFactoryHolderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HolderFactoryHolderCreated represents a HolderCreated event raised by the HolderFactory contract.
type HolderFactoryHolderCreated struct {
	Holder common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHolderCreated is a free log retrieval operation binding the contract event 0x4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a1716055.
//
// Solidity: e HolderCreated(holder address, owner address)
func (_HolderFactory *HolderFactoryFilterer) FilterHolderCreated(opts *bind.FilterOpts) (*HolderFactoryHolderCreatedIterator, error) {

	logs, sub, err := _HolderFactory.contract.FilterLogs(opts, "HolderCreated")
	if err != nil {
		return nil, err
	}
	return &HolderFactoryHolderCreatedIterator{contract: _HolderFactory.contract, event: "HolderCreated", logs: logs, sub: sub}, nil
}

// WatchHolderCreated is a free log subscription operation binding the contract event 0x4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a1716055.
//
// Solidity: e HolderCreated(holder address, owner address)
func (_HolderFactory *HolderFactoryFilterer) WatchHolderCreated(opts *bind.WatchOpts, sink chan<- *HolderFactoryHolderCreated) (event.Subscription, error) {

	logs, sub, err := _HolderFactory.contract.WatchLogs(opts, "HolderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HolderFactoryHolderCreated)
				if err := _HolderFactory.contract.UnpackLog(event, "HolderCreated", log); err != nil {
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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x608060405234801561001057600080fd5b50604051602080610c0a8339810180604052602081101561003057600080fd5b505160018054600160a060020a031916600160a060020a03909216919091179055610baa806100606000396000f3fe6080604052600436106100745763ffffffff60e060020a600035041663075461728114610076578063095ea7b3146100a757806318160ddd146100f457806323b872dd1461011b57806340c10f191461015e57806370a0823114610197578063a9059cbb146101ca578063dd62ed3e14610203575b005b34801561008257600080fd5b5061008b61023e565b60408051600160a060020a039092168252519081900360200190f35b3480156100b357600080fd5b506100e0600480360360408110156100ca57600080fd5b50600160a060020a03813516906020013561024d565b604080519115158252519081900360200190f35b34801561010057600080fd5b506101096102e0565b60408051918252519081900360200190f35b34801561012757600080fd5b506100e06004803603606081101561013e57600080fd5b50600160a060020a038135811691602081013590911690604001356102e6565b34801561016a57600080fd5b506100e06004803603604081101561018157600080fd5b50600160a060020a038135169060200135610649565b3480156101a357600080fd5b50610109600480360360208110156101ba57600080fd5b5035600160a060020a0316610762565b3480156101d657600080fd5b506100e0600480360360408110156101ed57600080fd5b50600160a060020a0381351690602001356107df565b34801561020f57600080fd5b506101096004803603604081101561022657600080fd5b50600160a060020a0381358116916020013516610a20565b600154600160a060020a031681565b60008061025933610abb565b905080600160a060020a031663310ec4a785856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b1580156102be57600080fd5b505af11580156102d2573d6000803e3d6000fd5b506001979650505050505050565b60005481565b6000806102f285610abb565b905060006102ff85610abb565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561033f57600080fd5b505afa158015610353573d6000803e3d6000fd5b505050506040513d602081101561036957600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156103ca57600080fd5b505afa1580156103de573d6000803e3d6000fd5b505050506040513d60208110156103f457600080fd5b5051604080517feb5a662e0000000000000000000000000000000000000000000000000000000081523360048201529051919250600091600160a060020a0387169163eb5a662e916024808301926020929190829003018186803b15801561045b57600080fd5b505afa15801561046f573d6000803e3d6000fd5b505050506040513d602081101561048557600080fd5b505190508683101561049657600080fd5b8682018211156104a557600080fd5b808711156104b257600080fd5b84600160a060020a031663fb1669ca8885036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156104fd57600080fd5b505af1158015610511573d6000803e3d6000fd5b5050505083600160a060020a031663fb1669ca8884016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561056057600080fd5b505af1158015610574573d6000803e3d6000fd5b5050604080517f310ec4a70000000000000000000000000000000000000000000000000000000081523360048201528a850360248201529051600160a060020a038916935063310ec4a79250604480830192600092919082900301818387803b1580156105e057600080fd5b505af11580156105f4573d6000803e3d6000fd5b5050604080518a81529051600160a060020a03808d1694508d1692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600198975050505050505050565b600154600090600160a060020a0316331461066357600080fd5b600061066e84610abb565b9050600081600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156106ae57600080fd5b505afa1580156106c2573d6000803e3d6000fd5b505050506040513d60208110156106d857600080fd5b505190508381018111156106eb57600080fd5b81600160a060020a031663fb1669ca8583016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561073657600080fd5b505af115801561074a573d6000803e3d6000fd5b50506000805487019055506001935050505092915050565b60008061076e83610abb565b905080600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107ac57600080fd5b505afa1580156107c0573d6000803e3d6000fd5b505050506040513d60208110156107d657600080fd5b50519392505050565b6000806107eb33610abb565b905060006107f885610abb565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561083857600080fd5b505afa15801561084c573d6000803e3d6000fd5b505050506040513d602081101561086257600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156108c357600080fd5b505afa1580156108d7573d6000803e3d6000fd5b505050506040513d60208110156108ed57600080fd5b50519050858210156108fe57600080fd5b85810181111561090d57600080fd5b83600160a060020a031663fb1669ca8784036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561095857600080fd5b505af115801561096c573d6000803e3d6000fd5b5050505082600160a060020a031663fb1669ca8783016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156109bb57600080fd5b505af11580156109cf573d6000803e3d6000fd5b5050604080518981529051600160a060020a038b1693503392507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b600080610a2c84610abb565b905080600160a060020a031663eb5a662e846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015610a8757600080fd5b505afa158015610a9b573d6000803e3d6000fd5b505050506040513d6020811015610ab157600080fd5b5051949350505050565b604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526f094949490000000000000000000000006021830152600160a060020a0384166c010000000000000000000000000260358301527f4a9e75186f08fc5e8c52bbad6aa27e7256b62a61438527f0121e7ea260534c3e6049808401829052845180850390910181526069909301909352815191012060009190630949494990803b848111610b7557600080fd5b5094935050505056fea165627a7a72305820ecc072490329bc94e2ea1763aab10d47f63f7a414ccccf54e8f510263570357d0029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, _minter common.Address) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, _minter)
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
// Solidity: function allowance(a address, _spender address) constant returns(uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, a common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", a, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(a address, _spender address) constant returns(uint256)
func (_Token *TokenSession) Allowance(a common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, a, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(a address, _spender address) constant returns(uint256)
func (_Token *TokenCallerSession) Allowance(a common.Address, _spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, a, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(a address) constant returns(uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, a common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", a)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(a address) constant returns(uint256)
func (_Token *TokenSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, a)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(a address) constant returns(uint256)
func (_Token *TokenCallerSession) BalanceOf(a common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, a)
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
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_Token *TokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
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
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_Token *TokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
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

// TokenFactoryABI is the input ABI used to generate the binding from.
const TokenFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"createToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"TokenCreated\",\"type\":\"event\"}]"

// TokenFactoryBin is the compiled bytecode used for deploying new contracts.
const TokenFactoryBin = `0x608060405234801561001057600080fd5b50610d96806100206000396000f3fe6080604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663c21ab7f98114610045575b600080fd5b34801561005157600080fd5b506100856004803603602081101561006857600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166100ae565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b600080826100ba610150565b73ffffffffffffffffffffffffffffffffffffffff909116815260405190819003602001906000f0801580156100f4573d6000803e3d6000fd5b506040805173ffffffffffffffffffffffffffffffffffffffff80841682528616602082015281519293507fd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139929081900390910190a192915050565b604051610c0a806101618339019056fe608060405234801561001057600080fd5b50604051602080610c0a8339810180604052602081101561003057600080fd5b505160018054600160a060020a031916600160a060020a03909216919091179055610baa806100606000396000f3fe6080604052600436106100745763ffffffff60e060020a600035041663075461728114610076578063095ea7b3146100a757806318160ddd146100f457806323b872dd1461011b57806340c10f191461015e57806370a0823114610197578063a9059cbb146101ca578063dd62ed3e14610203575b005b34801561008257600080fd5b5061008b61023e565b60408051600160a060020a039092168252519081900360200190f35b3480156100b357600080fd5b506100e0600480360360408110156100ca57600080fd5b50600160a060020a03813516906020013561024d565b604080519115158252519081900360200190f35b34801561010057600080fd5b506101096102e0565b60408051918252519081900360200190f35b34801561012757600080fd5b506100e06004803603606081101561013e57600080fd5b50600160a060020a038135811691602081013590911690604001356102e6565b34801561016a57600080fd5b506100e06004803603604081101561018157600080fd5b50600160a060020a038135169060200135610649565b3480156101a357600080fd5b50610109600480360360208110156101ba57600080fd5b5035600160a060020a0316610762565b3480156101d657600080fd5b506100e0600480360360408110156101ed57600080fd5b50600160a060020a0381351690602001356107df565b34801561020f57600080fd5b506101096004803603604081101561022657600080fd5b50600160a060020a0381358116916020013516610a20565b600154600160a060020a031681565b60008061025933610abb565b905080600160a060020a031663310ec4a785856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b1580156102be57600080fd5b505af11580156102d2573d6000803e3d6000fd5b506001979650505050505050565b60005481565b6000806102f285610abb565b905060006102ff85610abb565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561033f57600080fd5b505afa158015610353573d6000803e3d6000fd5b505050506040513d602081101561036957600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156103ca57600080fd5b505afa1580156103de573d6000803e3d6000fd5b505050506040513d60208110156103f457600080fd5b5051604080517feb5a662e0000000000000000000000000000000000000000000000000000000081523360048201529051919250600091600160a060020a0387169163eb5a662e916024808301926020929190829003018186803b15801561045b57600080fd5b505afa15801561046f573d6000803e3d6000fd5b505050506040513d602081101561048557600080fd5b505190508683101561049657600080fd5b8682018211156104a557600080fd5b808711156104b257600080fd5b84600160a060020a031663fb1669ca8885036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156104fd57600080fd5b505af1158015610511573d6000803e3d6000fd5b5050505083600160a060020a031663fb1669ca8884016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561056057600080fd5b505af1158015610574573d6000803e3d6000fd5b5050604080517f310ec4a70000000000000000000000000000000000000000000000000000000081523360048201528a850360248201529051600160a060020a038916935063310ec4a79250604480830192600092919082900301818387803b1580156105e057600080fd5b505af11580156105f4573d6000803e3d6000fd5b5050604080518a81529051600160a060020a03808d1694508d1692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600198975050505050505050565b600154600090600160a060020a0316331461066357600080fd5b600061066e84610abb565b9050600081600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156106ae57600080fd5b505afa1580156106c2573d6000803e3d6000fd5b505050506040513d60208110156106d857600080fd5b505190508381018111156106eb57600080fd5b81600160a060020a031663fb1669ca8583016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561073657600080fd5b505af115801561074a573d6000803e3d6000fd5b50506000805487019055506001935050505092915050565b60008061076e83610abb565b905080600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107ac57600080fd5b505afa1580156107c0573d6000803e3d6000fd5b505050506040513d60208110156107d657600080fd5b50519392505050565b6000806107eb33610abb565b905060006107f885610abb565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561083857600080fd5b505afa15801561084c573d6000803e3d6000fd5b505050506040513d602081101561086257600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156108c357600080fd5b505afa1580156108d7573d6000803e3d6000fd5b505050506040513d60208110156108ed57600080fd5b50519050858210156108fe57600080fd5b85810181111561090d57600080fd5b83600160a060020a031663fb1669ca8784036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561095857600080fd5b505af115801561096c573d6000803e3d6000fd5b5050505082600160a060020a031663fb1669ca8783016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156109bb57600080fd5b505af11580156109cf573d6000803e3d6000fd5b5050604080518981529051600160a060020a038b1693503392507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b600080610a2c84610abb565b905080600160a060020a031663eb5a662e846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015610a8757600080fd5b505afa158015610a9b573d6000803e3d6000fd5b505050506040513d6020811015610ab157600080fd5b5051949350505050565b604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526f094949490000000000000000000000006021830152600160a060020a0384166c010000000000000000000000000260358301527f4a9e75186f08fc5e8c52bbad6aa27e7256b62a61438527f0121e7ea260534c3e6049808401829052845180850390910181526069909301909352815191012060009190630949494990803b848111610b7557600080fd5b5094935050505056fea165627a7a72305820ecc072490329bc94e2ea1763aab10d47f63f7a414ccccf54e8f510263570357d0029a165627a7a723058209fe64c62e7dd87acf2959900ef15ea8542c54688ffbae2b6ec116ca0935c6dbe0029`

// DeployTokenFactory deploys a new Ethereum contract, binding an instance of TokenFactory to it.
func DeployTokenFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// TokenFactory is an auto generated Go binding around an Ethereum contract.
type TokenFactory struct {
	TokenFactoryCaller     // Read-only binding to the contract
	TokenFactoryTransactor // Write-only binding to the contract
	TokenFactoryFilterer   // Log filterer for contract events
}

// TokenFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenFactorySession struct {
	Contract     *TokenFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenFactoryCallerSession struct {
	Contract *TokenFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenFactoryTransactorSession struct {
	Contract     *TokenFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenFactoryRaw struct {
	Contract *TokenFactory // Generic contract binding to access the raw methods on
}

// TokenFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenFactoryCallerRaw struct {
	Contract *TokenFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// TokenFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenFactoryTransactorRaw struct {
	Contract *TokenFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenFactory creates a new instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactory(address common.Address, backend bind.ContractBackend) (*TokenFactory, error) {
	contract, err := bindTokenFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenFactory{TokenFactoryCaller: TokenFactoryCaller{contract: contract}, TokenFactoryTransactor: TokenFactoryTransactor{contract: contract}, TokenFactoryFilterer: TokenFactoryFilterer{contract: contract}}, nil
}

// NewTokenFactoryCaller creates a new read-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryCaller(address common.Address, caller bind.ContractCaller) (*TokenFactoryCaller, error) {
	contract, err := bindTokenFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryCaller{contract: contract}, nil
}

// NewTokenFactoryTransactor creates a new write-only instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenFactoryTransactor, error) {
	contract, err := bindTokenFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTransactor{contract: contract}, nil
}

// NewTokenFactoryFilterer creates a new log filterer instance of TokenFactory, bound to a specific deployed contract.
func NewTokenFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenFactoryFilterer, error) {
	contract, err := bindTokenFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenFactoryFilterer{contract: contract}, nil
}

// bindTokenFactory binds a generic wrapper to an already deployed contract.
func bindTokenFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.TokenFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.TokenFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenFactory *TokenFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenFactory *TokenFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenFactory *TokenFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenFactory.Contract.contract.Transact(opts, method, params...)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_TokenFactory *TokenFactoryTransactor) CreateToken(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _TokenFactory.contract.Transact(opts, "createToken", _minter)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_TokenFactory *TokenFactorySession) CreateToken(_minter common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _minter)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_TokenFactory *TokenFactoryTransactorSession) CreateToken(_minter common.Address) (*types.Transaction, error) {
	return _TokenFactory.Contract.CreateToken(&_TokenFactory.TransactOpts, _minter)
}

// TokenFactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the TokenFactory contract.
type TokenFactoryTokenCreatedIterator struct {
	Event *TokenFactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *TokenFactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenFactoryTokenCreated)
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
		it.Event = new(TokenFactoryTokenCreated)
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
func (it *TokenFactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenFactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenFactoryTokenCreated represents a TokenCreated event raised by the TokenFactory contract.
type TokenFactoryTokenCreated struct {
	Token  common.Address
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: e TokenCreated(token address, minter address)
func (_TokenFactory *TokenFactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts) (*TokenFactoryTokenCreatedIterator, error) {

	logs, sub, err := _TokenFactory.contract.FilterLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return &TokenFactoryTokenCreatedIterator{contract: _TokenFactory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: e TokenCreated(token address, minter address)
func (_TokenFactory *TokenFactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *TokenFactoryTokenCreated) (event.Subscription, error) {

	logs, sub, err := _TokenFactory.contract.WatchLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenFactoryTokenCreated)
				if err := _TokenFactory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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
