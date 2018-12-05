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

// FactoryABI is the input ABI used to generate the binding from.
const FactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createHolder\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\"}],\"name\":\"createToken\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"holder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"HolderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"TokenCreated\",\"type\":\"event\"}]"

// FactoryBin is the compiled bytecode used for deploying new contracts.
const FactoryBin = `0x608060405234801561001057600080fd5b506120d8806100206000396000f3fe60806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166351e8054c8114610050578063c21ab7f91461009f575b600080fd5b34801561005c57600080fd5b506100836004803603602081101561007357600080fd5b5035600160a060020a03166100d2565b60408051600160a060020a039092168252519081900360200190f35b3480156100ab57600080fd5b50610083600480360360208110156100c257600080fd5b5035600160a060020a03166101ee565b60006060611200604051908101604052806111d98152602001610ed46111d99139805160208201818120929350600091869183f59050600160a060020a038116151561011d57600080fd5b604080517f99d3c204000000000000000000000000000000000000000000000000000000008152600160a060020a03878116600483015260248201859052915183928316916399d3c20491604480830192600092919082900301818387803b15801561018857600080fd5b505af115801561019c573d6000803e3d6000fd5b505060408051600160a060020a0380861682528a16602082015281517f4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a17160559450908190039091019150a195945050505050565b60008082306101fb61027d565b600160a060020a03928316815291166020820152604080519182900301906000f08015801561022e573d6000803e3d6000fd5b5060408051600160a060020a0380841682528616602082015281519293507fd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139929081900390910190a192915050565b604051610c468061028e8339019056fe608060405234801561001057600080fd5b50604051604080610c468339810180604052604081101561003057600080fd5b50805160209091015160018054600160a060020a03938416600160a060020a03199182161790915560028054939092169216919091179055610bcf806100776000396000f3fe60806040526004361061007f5763ffffffff60e060020a600035041663075461728114610081578063095ea7b3146100b257806318160ddd146100ff57806323b872dd1461012657806340c10f191461016957806370a08231146101a2578063a9059cbb146101d5578063c45a01551461020e578063dd62ed3e14610223575b005b34801561008d57600080fd5b5061009661025e565b60408051600160a060020a039092168252519081900360200190f35b3480156100be57600080fd5b506100eb600480360360408110156100d557600080fd5b50600160a060020a03813516906020013561026d565b604080519115158252519081900360200190f35b34801561010b57600080fd5b50610114610300565b60408051918252519081900360200190f35b34801561013257600080fd5b506100eb6004803603606081101561014957600080fd5b50600160a060020a03813581169160208101359091169060400135610306565b34801561017557600080fd5b506100eb6004803603604081101561018c57600080fd5b50600160a060020a038135169060200135610669565b3480156101ae57600080fd5b50610114600480360360208110156101c557600080fd5b5035600160a060020a0316610782565b3480156101e157600080fd5b506100eb600480360360408110156101f857600080fd5b50600160a060020a0381351690602001356107ff565b34801561021a57600080fd5b50610096610a40565b34801561022f57600080fd5b506101146004803603604081101561024657600080fd5b50600160a060020a0381358116916020013516610a4f565b600154600160a060020a031681565b60008061027933610aea565b905080600160a060020a031663310ec4a785856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b1580156102de57600080fd5b505af11580156102f2573d6000803e3d6000fd5b506001979650505050505050565b60005481565b60008061031285610aea565b9050600061031f85610aea565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561035f57600080fd5b505afa158015610373573d6000803e3d6000fd5b505050506040513d602081101561038957600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156103ea57600080fd5b505afa1580156103fe573d6000803e3d6000fd5b505050506040513d602081101561041457600080fd5b5051604080517feb5a662e0000000000000000000000000000000000000000000000000000000081523360048201529051919250600091600160a060020a0387169163eb5a662e916024808301926020929190829003018186803b15801561047b57600080fd5b505afa15801561048f573d6000803e3d6000fd5b505050506040513d60208110156104a557600080fd5b50519050868310156104b657600080fd5b8682018211156104c557600080fd5b808711156104d257600080fd5b84600160a060020a031663fb1669ca8885036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b5050505083600160a060020a031663fb1669ca8884016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561058057600080fd5b505af1158015610594573d6000803e3d6000fd5b5050604080517f310ec4a70000000000000000000000000000000000000000000000000000000081523360048201528a850360248201529051600160a060020a038916935063310ec4a79250604480830192600092919082900301818387803b15801561060057600080fd5b505af1158015610614573d6000803e3d6000fd5b5050604080518a81529051600160a060020a03808d1694508d1692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600198975050505050505050565b600154600090600160a060020a0316331461068357600080fd5b600061068e84610aea565b9050600081600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156106ce57600080fd5b505afa1580156106e2573d6000803e3d6000fd5b505050506040513d60208110156106f857600080fd5b5051905083810181111561070b57600080fd5b81600160a060020a031663fb1669ca8583016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561075657600080fd5b505af115801561076a573d6000803e3d6000fd5b50506000805487019055506001935050505092915050565b60008061078e83610aea565b905080600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107cc57600080fd5b505afa1580156107e0573d6000803e3d6000fd5b505050506040513d60208110156107f657600080fd5b50519392505050565b60008061080b33610aea565b9050600061081885610aea565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561085857600080fd5b505afa15801561086c573d6000803e3d6000fd5b505050506040513d602081101561088257600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156108e357600080fd5b505afa1580156108f7573d6000803e3d6000fd5b505050506040513d602081101561090d57600080fd5b505190508582101561091e57600080fd5b85810181111561092d57600080fd5b83600160a060020a031663fb1669ca8784036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561097857600080fd5b505af115801561098c573d6000803e3d6000fd5b5050505082600160a060020a031663fb1669ca8783016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156109db57600080fd5b505af11580156109ef573d6000803e3d6000fd5b5050604080518981529051600160a060020a038b1693503392507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b600254600160a060020a031681565b600080610a5b84610aea565b905080600160a060020a031663eb5a662e846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015610ab657600080fd5b505afa158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b5051949350505050565b600254604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526c01000000000000000000000000600160a060020a039485168102602184015293851690930260358201527f4a9e75186f08fc5e8c52bbad6aa27e7256b62a61438527f0121e7ea260534c3e6049808301829052835180840390910181526069909201909252805192019190912060009190803b838111610b9b57600080fd5b50939250505056fea165627a7a72305820398cc77383325948fd521c4c268175f6ea7f5e7ea158a15d759b852da6b8f7b20029608060405234801561001057600080fd5b506001600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061113a8061009f6000396000f3fe608060405260043610610099576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806312065fe01461009b5780632a847f46146100c6578063310ec4a7146101175780634df4d0391461017257806383197ef0146101c95780638852671e146101e057806399d3c20414610251578063eb5a662e146102ac578063fb1669ca14610311575b005b3480156100a757600080fd5b506100b061034c565b6040518082815260200191505060405180910390f35b3480156100d257600080fd5b50610115600480360360208110156100e957600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610469565b005b34801561012357600080fd5b506101706004803603604081101561013a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061078e565b005b34801561017e57600080fd5b506101876108e9565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101d557600080fd5b506101de6108ee565b005b3480156101ec57600080fd5b5061024f6004803603604081101561020357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506109e3565b005b34801561025d57600080fd5b506102aa6004803603604081101561027457600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d07565b005b3480156102b857600080fd5b506102fb600480360360208110156102cf57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e95565b6040518082815260200191505060405180910390f35b34801561031d57600080fd5b5061034a6004803603602081101561033457600080fd5b8101908080359060200190929190505050610ff1565b005b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561038a57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561042457600080fd5b600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905090565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156104c757600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561052357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415801561058d5750600173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b151561059857600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151561063157600080fd5b600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600080600173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515156107ca57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561086457600080fd5b80600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b600181565b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415151561094c57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156109a857600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16ff5b600073ffffffffffffffffffffffffffffffffffffffff16600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610a4157600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610a9d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610b075750600173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b1515610b1257600080fd5b8173ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610baa57600080fd5b6000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b3073ffffffffffffffffffffffffffffffffffffffff16600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168383604051602001808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c010000000000000000000000000281526014018281526020019350505050604051602081830303815290604052805190602001206001900473ffffffffffffffffffffffffffffffffffffffff161415610e915781600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050565b6000600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151515610ed357600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614151515610f6d57600080fd5b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b600173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415151561102d57600080fd5b600073ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515156110c757600080fd5b80600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505056fea165627a7a72305820b07d304542735e93c801d4e3e1fe1cf63b58ea0c3bd0fb00642de67898f3480b0029a165627a7a72305820a2d15a21af4627642be5b512f11ae7bb70e7133714d3a00336298f4329ea29c90029`

// DeployFactory deploys a new Ethereum contract, binding an instance of Factory to it.
func DeployFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Factory, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(FactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// Factory is an auto generated Go binding around an Ethereum contract.
type Factory struct {
	FactoryCaller     // Read-only binding to the contract
	FactoryTransactor // Write-only binding to the contract
	FactoryFilterer   // Log filterer for contract events
}

// FactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FactorySession struct {
	Contract     *Factory          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FactoryCallerSession struct {
	Contract *FactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// FactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FactoryTransactorSession struct {
	Contract     *FactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FactoryRaw struct {
	Contract *Factory // Generic contract binding to access the raw methods on
}

// FactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FactoryCallerRaw struct {
	Contract *FactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FactoryTransactorRaw struct {
	Contract *FactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFactory creates a new instance of Factory, bound to a specific deployed contract.
func NewFactory(address common.Address, backend bind.ContractBackend) (*Factory, error) {
	contract, err := bindFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Factory{FactoryCaller: FactoryCaller{contract: contract}, FactoryTransactor: FactoryTransactor{contract: contract}, FactoryFilterer: FactoryFilterer{contract: contract}}, nil
}

// NewFactoryCaller creates a new read-only instance of Factory, bound to a specific deployed contract.
func NewFactoryCaller(address common.Address, caller bind.ContractCaller) (*FactoryCaller, error) {
	contract, err := bindFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryCaller{contract: contract}, nil
}

// NewFactoryTransactor creates a new write-only instance of Factory, bound to a specific deployed contract.
func NewFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FactoryTransactor, error) {
	contract, err := bindFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FactoryTransactor{contract: contract}, nil
}

// NewFactoryFilterer creates a new log filterer instance of Factory, bound to a specific deployed contract.
func NewFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FactoryFilterer, error) {
	contract, err := bindFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FactoryFilterer{contract: contract}, nil
}

// bindFactory binds a generic wrapper to an already deployed contract.
func bindFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.FactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.FactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Factory *FactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Factory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Factory *FactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Factory *FactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Factory.Contract.contract.Transact(opts, method, params...)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_Factory *FactoryTransactor) CreateHolder(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createHolder", _owner)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_Factory *FactorySession) CreateHolder(_owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateHolder(&_Factory.TransactOpts, _owner)
}

// CreateHolder is a paid mutator transaction binding the contract method 0x51e8054c.
//
// Solidity: function createHolder(_owner address) returns(address)
func (_Factory *FactoryTransactorSession) CreateHolder(_owner common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateHolder(&_Factory.TransactOpts, _owner)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_Factory *FactoryTransactor) CreateToken(opts *bind.TransactOpts, _minter common.Address) (*types.Transaction, error) {
	return _Factory.contract.Transact(opts, "createToken", _minter)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_Factory *FactorySession) CreateToken(_minter common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateToken(&_Factory.TransactOpts, _minter)
}

// CreateToken is a paid mutator transaction binding the contract method 0xc21ab7f9.
//
// Solidity: function createToken(_minter address) returns(address)
func (_Factory *FactoryTransactorSession) CreateToken(_minter common.Address) (*types.Transaction, error) {
	return _Factory.Contract.CreateToken(&_Factory.TransactOpts, _minter)
}

// FactoryHolderCreatedIterator is returned from FilterHolderCreated and is used to iterate over the raw logs and unpacked data for HolderCreated events raised by the Factory contract.
type FactoryHolderCreatedIterator struct {
	Event *FactoryHolderCreated // Event containing the contract specifics and raw log

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
func (it *FactoryHolderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryHolderCreated)
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
		it.Event = new(FactoryHolderCreated)
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
func (it *FactoryHolderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryHolderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryHolderCreated represents a HolderCreated event raised by the Factory contract.
type FactoryHolderCreated struct {
	Holder common.Address
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterHolderCreated is a free log retrieval operation binding the contract event 0x4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a1716055.
//
// Solidity: e HolderCreated(holder address, owner address)
func (_Factory *FactoryFilterer) FilterHolderCreated(opts *bind.FilterOpts) (*FactoryHolderCreatedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "HolderCreated")
	if err != nil {
		return nil, err
	}
	return &FactoryHolderCreatedIterator{contract: _Factory.contract, event: "HolderCreated", logs: logs, sub: sub}, nil
}

// WatchHolderCreated is a free log subscription operation binding the contract event 0x4c6f74edaaaad2c1bf11ab46b1a77f3644b07c39ee9575704ec557c6a1716055.
//
// Solidity: e HolderCreated(holder address, owner address)
func (_Factory *FactoryFilterer) WatchHolderCreated(opts *bind.WatchOpts, sink chan<- *FactoryHolderCreated) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "HolderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryHolderCreated)
				if err := _Factory.contract.UnpackLog(event, "HolderCreated", log); err != nil {
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

// FactoryTokenCreatedIterator is returned from FilterTokenCreated and is used to iterate over the raw logs and unpacked data for TokenCreated events raised by the Factory contract.
type FactoryTokenCreatedIterator struct {
	Event *FactoryTokenCreated // Event containing the contract specifics and raw log

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
func (it *FactoryTokenCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FactoryTokenCreated)
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
		it.Event = new(FactoryTokenCreated)
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
func (it *FactoryTokenCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FactoryTokenCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FactoryTokenCreated represents a TokenCreated event raised by the Factory contract.
type FactoryTokenCreated struct {
	Token  common.Address
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCreated is a free log retrieval operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: e TokenCreated(token address, minter address)
func (_Factory *FactoryFilterer) FilterTokenCreated(opts *bind.FilterOpts) (*FactoryTokenCreatedIterator, error) {

	logs, sub, err := _Factory.contract.FilterLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return &FactoryTokenCreatedIterator{contract: _Factory.contract, event: "TokenCreated", logs: logs, sub: sub}, nil
}

// WatchTokenCreated is a free log subscription operation binding the contract event 0xd5f9bdf12adf29dab0248c349842c3822d53ae2bb4f36352f301630d018c8139.
//
// Solidity: e TokenCreated(token address, minter address)
func (_Factory *FactoryFilterer) WatchTokenCreated(opts *bind.WatchOpts, sink chan<- *FactoryTokenCreated) (event.Subscription, error) {

	logs, sub, err := _Factory.contract.WatchLogs(opts, "TokenCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FactoryTokenCreated)
				if err := _Factory.contract.UnpackLog(event, "TokenCreated", log); err != nil {
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

// HolderABI is the input ABI used to generate the binding from.
const HolderABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\"}],\"name\":\"addTokenContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"SENTINEL_TOKEN_CONTRACTS\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"destroy\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"tokenContract\",\"type\":\"address\"},{\"name\":\"prev\",\"type\":\"address\"}],\"name\":\"removeTokenContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"init_code_hash\",\"type\":\"bytes32\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"getAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// HolderBin is the compiled bytecode used for deploying new contracts.
const HolderBin = `0x608060405234801561001057600080fd5b50600160008181526020527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a03191690911790556106718061005c6000396000f3fe6080604052600436106100985763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe0811461009a5780632a847f46146100c1578063310ec4a7146100f45780634df4d0391461012d57806383197ef01461015e5780638852671e1461017357806399d3c204146101ae578063eb5a662e146101e7578063fb1669ca1461021a575b005b3480156100a657600080fd5b506100af610244565b60408051918252519081900360200190f35b3480156100cd57600080fd5b50610098600480360360208110156100e457600080fd5b5035600160a060020a031661028c565b34801561010057600080fd5b506100986004803603604081101561011757600080fd5b50600160a060020a03813516906020013561037a565b34801561013957600080fd5b506101426103d6565b60408051600160a060020a039092168252519081900360200190f35b34801561016a57600080fd5b506100986103db565b34801561017f57600080fd5b506100986004803603604081101561019657600080fd5b50600160a060020a0381358116916020013516610417565b3480156101ba57600080fd5b50610098600480360360408110156101d157600080fd5b50600160a060020a0381351690602001356104ea565b3480156101f357600080fd5b506100af6004803603602081101561020a57600080fd5b5035600160a060020a03166105a4565b34801561022657600080fd5b506100986004803603602081101561023d57600080fd5b5035610601565b6000336001141561025457600080fd5b33600090815260208190526040902054600160a060020a0316151561027857600080fd5b503360009081526003602052604090205490565b600154600160a060020a031615156102a357600080fd5b600154600160a060020a031633146102ba57600080fd5b600160a060020a038116158015906102dc5750600160a060020a038116600114155b15156102e757600080fd5b600160a060020a03818116600090815260208190526040902054161561030c57600080fd5b600060208190527fada5013122d395ba3c54772283fb069b10426056ef8ca54750cb9bb552a59e7d8054600160a060020a039384168084526040842080549590921673ffffffffffffffffffffffffffffffffffffffff199586161790915560019092528054909216179055565b336001141561038857600080fd5b33600090815260208190526040902054600160a060020a031615156103ac57600080fd5b336000908152600460209081526040808320600160a060020a039590951683529390529190912055565b600181565b600154600160a060020a031615156103f257600080fd5b600154600160a060020a0316331461040957600080fd5b600154600160a060020a0316ff5b600154600160a060020a0316151561042e57600080fd5b600154600160a060020a0316331461044557600080fd5b600160a060020a038216158015906104675750600160a060020a038216600114155b151561047257600080fd5b600160a060020a0381811660009081526020819052604090205481169083161461049b57600080fd5b600160a060020a039182166000818152602081905260408082208054948616835290822080549490951673ffffffffffffffffffffffffffffffffffffffff1994851617909455528154169055565b600254604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526c01000000000000000000000000600160a060020a039485168102602184015286851602603583015260498083018690528351808403909101815260699092019092528051910120163014156105a05760018054600160a060020a03841673ffffffffffffffffffffffffffffffffffffffff19918216179091556002805490911690555b5050565b600033600114156105b457600080fd5b33600090815260208190526040902054600160a060020a031615156105d857600080fd5b50336000908152600460209081526040808320600160a060020a03949094168352929052205490565b336001141561060f57600080fd5b33600090815260208190526040902054600160a060020a0316151561063357600080fd5b3360009081526003602052604090205556fea165627a7a72305820568b6d030bb399dd48387d226608cee53b891856cb9ec300a25a6138acda7a070029`

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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"minter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"a\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_minter\",\"type\":\"address\"},{\"name\":\"_factory\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `0x608060405234801561001057600080fd5b50604051604080610c468339810180604052604081101561003057600080fd5b50805160209091015160018054600160a060020a03938416600160a060020a03199182161790915560028054939092169216919091179055610bcf806100776000396000f3fe60806040526004361061007f5763ffffffff60e060020a600035041663075461728114610081578063095ea7b3146100b257806318160ddd146100ff57806323b872dd1461012657806340c10f191461016957806370a08231146101a2578063a9059cbb146101d5578063c45a01551461020e578063dd62ed3e14610223575b005b34801561008d57600080fd5b5061009661025e565b60408051600160a060020a039092168252519081900360200190f35b3480156100be57600080fd5b506100eb600480360360408110156100d557600080fd5b50600160a060020a03813516906020013561026d565b604080519115158252519081900360200190f35b34801561010b57600080fd5b50610114610300565b60408051918252519081900360200190f35b34801561013257600080fd5b506100eb6004803603606081101561014957600080fd5b50600160a060020a03813581169160208101359091169060400135610306565b34801561017557600080fd5b506100eb6004803603604081101561018c57600080fd5b50600160a060020a038135169060200135610669565b3480156101ae57600080fd5b50610114600480360360208110156101c557600080fd5b5035600160a060020a0316610782565b3480156101e157600080fd5b506100eb600480360360408110156101f857600080fd5b50600160a060020a0381351690602001356107ff565b34801561021a57600080fd5b50610096610a40565b34801561022f57600080fd5b506101146004803603604081101561024657600080fd5b50600160a060020a0381358116916020013516610a4f565b600154600160a060020a031681565b60008061027933610aea565b905080600160a060020a031663310ec4a785856040518363ffffffff1660e060020a0281526004018083600160a060020a0316600160a060020a0316815260200182815260200192505050600060405180830381600087803b1580156102de57600080fd5b505af11580156102f2573d6000803e3d6000fd5b506001979650505050505050565b60005481565b60008061031285610aea565b9050600061031f85610aea565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561035f57600080fd5b505afa158015610373573d6000803e3d6000fd5b505050506040513d602081101561038957600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156103ea57600080fd5b505afa1580156103fe573d6000803e3d6000fd5b505050506040513d602081101561041457600080fd5b5051604080517feb5a662e0000000000000000000000000000000000000000000000000000000081523360048201529051919250600091600160a060020a0387169163eb5a662e916024808301926020929190829003018186803b15801561047b57600080fd5b505afa15801561048f573d6000803e3d6000fd5b505050506040513d60208110156104a557600080fd5b50519050868310156104b657600080fd5b8682018211156104c557600080fd5b808711156104d257600080fd5b84600160a060020a031663fb1669ca8885036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561051d57600080fd5b505af1158015610531573d6000803e3d6000fd5b5050505083600160a060020a031663fb1669ca8884016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561058057600080fd5b505af1158015610594573d6000803e3d6000fd5b5050604080517f310ec4a70000000000000000000000000000000000000000000000000000000081523360048201528a850360248201529051600160a060020a038916935063310ec4a79250604480830192600092919082900301818387803b15801561060057600080fd5b505af1158015610614573d6000803e3d6000fd5b5050604080518a81529051600160a060020a03808d1694508d1692507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a350600198975050505050505050565b600154600090600160a060020a0316331461068357600080fd5b600061068e84610aea565b9050600081600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156106ce57600080fd5b505afa1580156106e2573d6000803e3d6000fd5b505050506040513d60208110156106f857600080fd5b5051905083810181111561070b57600080fd5b81600160a060020a031663fb1669ca8583016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561075657600080fd5b505af115801561076a573d6000803e3d6000fd5b50506000805487019055506001935050505092915050565b60008061078e83610aea565b905080600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b1580156107cc57600080fd5b505afa1580156107e0573d6000803e3d6000fd5b505050506040513d60208110156107f657600080fd5b50519392505050565b60008061080b33610aea565b9050600061081885610aea565b9050600082600160a060020a03166312065fe06040518163ffffffff1660e060020a02815260040160206040518083038186803b15801561085857600080fd5b505afa15801561086c573d6000803e3d6000fd5b505050506040513d602081101561088257600080fd5b5051604080517f12065fe00000000000000000000000000000000000000000000000000000000081529051919250600091600160a060020a038516916312065fe0916004808301926020929190829003018186803b1580156108e357600080fd5b505afa1580156108f7573d6000803e3d6000fd5b505050506040513d602081101561090d57600080fd5b505190508582101561091e57600080fd5b85810181111561092d57600080fd5b83600160a060020a031663fb1669ca8784036040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b15801561097857600080fd5b505af115801561098c573d6000803e3d6000fd5b5050505082600160a060020a031663fb1669ca8783016040518263ffffffff1660e060020a02815260040180828152602001915050600060405180830381600087803b1580156109db57600080fd5b505af11580156109ef573d6000803e3d6000fd5b5050604080518981529051600160a060020a038b1693503392507fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9181900360200190a35060019695505050505050565b600254600160a060020a031681565b600080610a5b84610aea565b905080600160a060020a031663eb5a662e846040518263ffffffff1660e060020a0281526004018082600160a060020a0316600160a060020a0316815260200191505060206040518083038186803b158015610ab657600080fd5b505afa158015610aca573d6000803e3d6000fd5b505050506040513d6020811015610ae057600080fd5b5051949350505050565b600254604080517fff000000000000000000000000000000000000000000000000000000000000006020808301919091526c01000000000000000000000000600160a060020a039485168102602184015293851690930260358201527f4a9e75186f08fc5e8c52bbad6aa27e7256b62a61438527f0121e7ea260534c3e6049808301829052835180840390910181526069909201909252805192019190912060009190803b838111610b9b57600080fd5b50939250505056fea165627a7a72305820398cc77383325948fd521c4c268175f6ea7f5e7ea158a15d759b852da6b8f7b20029`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, _minter common.Address, _factory common.Address) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, _minter, _factory)
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

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_Token *TokenCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "factory")
	return *ret0, err
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_Token *TokenSession) Factory() (common.Address, error) {
	return _Token.Contract.Factory(&_Token.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() constant returns(address)
func (_Token *TokenCallerSession) Factory() (common.Address, error) {
	return _Token.Contract.Factory(&_Token.CallOpts)
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
