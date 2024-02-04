// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package addresbook

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

// AddressBookMetaData contains all meta data concerning the AddressBook contract.
var AddressBookMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_alias\",\"type\":\"string\"}],\"name\":\"addAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getAlias\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"removeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610fa0806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80634ba79dfe1461005157806399900d111461006d578063a39fac121461009d578063d033c456146100bb575b600080fd5b61006b60048036038101906100669190610858565b6100d7565b005b61008760048036038101906100829190610858565b610484565b6040516100949190610915565b60405180910390f35b6100a5610592565b6040516100b291906109f5565b60405180910390f35b6100d560048036038101906100d09190610b4c565b61065c565b005b60008060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050905060005b8181101561047f576000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020818154811061017857610177610ba8565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361046c576000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490506001108015610231575060018261022e9190610c10565b81105b15610356576000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001836102819190610c10565b8154811061029257610291610ba8565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020828154811061030d5761030c610ba8565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206001836103a19190610c10565b815481106103b2576103b1610ba8565b5b9060005260206000200160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006104679190610789565b61047f565b808061047790610c44565b915050610120565b505050565b6060600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805461050d90610cbb565b80601f016020809104026020016040519081016040528092919081815260200182805461053990610cbb565b80156105865780601f1061055b57610100808354040283529160200191610586565b820191906000526020600020905b81548152906001019060200180831161056957829003601f168201915b50505050509050919050565b60606000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080548060200260200160405190810160405280929190818152602001828054801561065257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610608575b5050505050905090565b6000803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090816107849190610e98565b505050565b50805461079590610cbb565b6000825580601f106107a757506107c6565b601f0160209004906000526020600020908101906107c591906107c9565b5b50565b5b808211156107e25760008160009055506001016107ca565b5090565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610825826107fa565b9050919050565b6108358161081a565b811461084057600080fd5b50565b6000813590506108528161082c565b92915050565b60006020828403121561086e5761086d6107f0565b5b600061087c84828501610843565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156108bf5780820151818401526020810190506108a4565b60008484015250505050565b6000601f19601f8301169050919050565b60006108e782610885565b6108f18185610890565b93506109018185602086016108a1565b61090a816108cb565b840191505092915050565b6000602082019050818103600083015261092f81846108dc565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61096c8161081a565b82525050565b600061097e8383610963565b60208301905092915050565b6000602082019050919050565b60006109a282610937565b6109ac8185610942565b93506109b783610953565b8060005b838110156109e85781516109cf8882610972565b97506109da8361098a565b9250506001810190506109bb565b5085935050505092915050565b60006020820190508181036000830152610a0f8184610997565b905092915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a59826108cb565b810181811067ffffffffffffffff82111715610a7857610a77610a21565b5b80604052505050565b6000610a8b6107e6565b9050610a978282610a50565b919050565b600067ffffffffffffffff821115610ab757610ab6610a21565b5b610ac0826108cb565b9050602081019050919050565b82818337600083830152505050565b6000610aef610aea84610a9c565b610a81565b905082815260208101848484011115610b0b57610b0a610a1c565b5b610b16848285610acd565b509392505050565b600082601f830112610b3357610b32610a17565b5b8135610b43848260208601610adc565b91505092915050565b60008060408385031215610b6357610b626107f0565b5b6000610b7185828601610843565b925050602083013567ffffffffffffffff811115610b9257610b916107f5565b5b610b9e85828601610b1e565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000819050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000610c1b82610bd7565b9150610c2683610bd7565b9250828203905081811115610c3e57610c3d610be1565b5b92915050565b6000610c4f82610bd7565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610c8157610c80610be1565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610cd357607f821691505b602082108103610ce657610ce5610c8c565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610d4e7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610d11565b610d588683610d11565b95508019841693508086168417925050509392505050565b6000819050919050565b6000610d95610d90610d8b84610bd7565b610d70565b610bd7565b9050919050565b6000819050919050565b610daf83610d7a565b610dc3610dbb82610d9c565b848454610d1e565b825550505050565b600090565b610dd8610dcb565b610de3818484610da6565b505050565b5b81811015610e0757610dfc600082610dd0565b600181019050610de9565b5050565b601f821115610e4c57610e1d81610cec565b610e2684610d01565b81016020851015610e35578190505b610e49610e4185610d01565b830182610de8565b50505b505050565b600082821c905092915050565b6000610e6f60001984600802610e51565b1980831691505092915050565b6000610e888383610e5e565b9150826002028217905092915050565b610ea182610885565b67ffffffffffffffff811115610eba57610eb9610a21565b5b610ec48254610cbb565b610ecf828285610e0b565b600060209050601f831160018114610f025760008415610ef0578287015190505b610efa8582610e7c565b865550610f62565b601f198416610f1086610cec565b60005b82811015610f3857848901518255600182019150602085019450602081019050610f13565b86831015610f555784890151610f51601f891682610e5e565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220817325953d09e2104181995a43a6dda3c73b230e27863e2c192a6d7f0efc8aab64736f6c63430008100033",
}

// AddressBookABI is the input ABI used to generate the binding from.
// Deprecated: Use AddressBookMetaData.ABI instead.
var AddressBookABI = AddressBookMetaData.ABI

// AddressBookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AddressBookMetaData.Bin instead.
var AddressBookBin = AddressBookMetaData.Bin

// DeployAddressBook deploys a new Ethereum contract, binding an instance of AddressBook to it.
func DeployAddressBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AddressBook, error) {
	parsed, err := AddressBookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AddressBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AddressBook{AddressBookCaller: AddressBookCaller{contract: contract}, AddressBookTransactor: AddressBookTransactor{contract: contract}, AddressBookFilterer: AddressBookFilterer{contract: contract}}, nil
}

// AddressBook is an auto generated Go binding around an Ethereum contract.
type AddressBook struct {
	AddressBookCaller     // Read-only binding to the contract
	AddressBookTransactor // Write-only binding to the contract
	AddressBookFilterer   // Log filterer for contract events
}

// AddressBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressBookSession struct {
	Contract     *AddressBook      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressBookCallerSession struct {
	Contract *AddressBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AddressBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressBookTransactorSession struct {
	Contract     *AddressBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AddressBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressBookRaw struct {
	Contract *AddressBook // Generic contract binding to access the raw methods on
}

// AddressBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressBookCallerRaw struct {
	Contract *AddressBookCaller // Generic read-only contract binding to access the raw methods on
}

// AddressBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressBookTransactorRaw struct {
	Contract *AddressBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddressBook creates a new instance of AddressBook, bound to a specific deployed contract.
func NewAddressBook(address common.Address, backend bind.ContractBackend) (*AddressBook, error) {
	contract, err := bindAddressBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AddressBook{AddressBookCaller: AddressBookCaller{contract: contract}, AddressBookTransactor: AddressBookTransactor{contract: contract}, AddressBookFilterer: AddressBookFilterer{contract: contract}}, nil
}

// NewAddressBookCaller creates a new read-only instance of AddressBook, bound to a specific deployed contract.
func NewAddressBookCaller(address common.Address, caller bind.ContractCaller) (*AddressBookCaller, error) {
	contract, err := bindAddressBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressBookCaller{contract: contract}, nil
}

// NewAddressBookTransactor creates a new write-only instance of AddressBook, bound to a specific deployed contract.
func NewAddressBookTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressBookTransactor, error) {
	contract, err := bindAddressBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressBookTransactor{contract: contract}, nil
}

// NewAddressBookFilterer creates a new log filterer instance of AddressBook, bound to a specific deployed contract.
func NewAddressBookFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressBookFilterer, error) {
	contract, err := bindAddressBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressBookFilterer{contract: contract}, nil
}

// bindAddressBook binds a generic wrapper to an already deployed contract.
func bindAddressBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AddressBookMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressBook *AddressBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressBook.Contract.AddressBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressBook *AddressBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressBook.Contract.AddressBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressBook *AddressBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressBook.Contract.AddressBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AddressBook *AddressBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AddressBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AddressBook *AddressBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AddressBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AddressBook *AddressBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AddressBook.Contract.contract.Transact(opts, method, params...)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressBook *AddressBookCaller) GetAddresses(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AddressBook.contract.Call(opts, &out, "getAddresses")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressBook *AddressBookSession) GetAddresses() ([]common.Address, error) {
	return _AddressBook.Contract.GetAddresses(&_AddressBook.CallOpts)
}

// GetAddresses is a free data retrieval call binding the contract method 0xa39fac12.
//
// Solidity: function getAddresses() view returns(address[])
func (_AddressBook *AddressBookCallerSession) GetAddresses() ([]common.Address, error) {
	return _AddressBook.Contract.GetAddresses(&_AddressBook.CallOpts)
}

// GetAlias is a free data retrieval call binding the contract method 0x99900d11.
//
// Solidity: function getAlias(address addr) view returns(string)
func (_AddressBook *AddressBookCaller) GetAlias(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _AddressBook.contract.Call(opts, &out, "getAlias", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetAlias is a free data retrieval call binding the contract method 0x99900d11.
//
// Solidity: function getAlias(address addr) view returns(string)
func (_AddressBook *AddressBookSession) GetAlias(addr common.Address) (string, error) {
	return _AddressBook.Contract.GetAlias(&_AddressBook.CallOpts, addr)
}

// GetAlias is a free data retrieval call binding the contract method 0x99900d11.
//
// Solidity: function getAlias(address addr) view returns(string)
func (_AddressBook *AddressBookCallerSession) GetAlias(addr common.Address) (string, error) {
	return _AddressBook.Contract.GetAlias(&_AddressBook.CallOpts, addr)
}

// AddAddress is a paid mutator transaction binding the contract method 0xd033c456.
//
// Solidity: function addAddress(address addr, string _alias) returns()
func (_AddressBook *AddressBookTransactor) AddAddress(opts *bind.TransactOpts, addr common.Address, _alias string) (*types.Transaction, error) {
	return _AddressBook.contract.Transact(opts, "addAddress", addr, _alias)
}

// AddAddress is a paid mutator transaction binding the contract method 0xd033c456.
//
// Solidity: function addAddress(address addr, string _alias) returns()
func (_AddressBook *AddressBookSession) AddAddress(addr common.Address, _alias string) (*types.Transaction, error) {
	return _AddressBook.Contract.AddAddress(&_AddressBook.TransactOpts, addr, _alias)
}

// AddAddress is a paid mutator transaction binding the contract method 0xd033c456.
//
// Solidity: function addAddress(address addr, string _alias) returns()
func (_AddressBook *AddressBookTransactorSession) AddAddress(addr common.Address, _alias string) (*types.Transaction, error) {
	return _AddressBook.Contract.AddAddress(&_AddressBook.TransactOpts, addr, _alias)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address addr) returns()
func (_AddressBook *AddressBookTransactor) RemoveAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _AddressBook.contract.Transact(opts, "removeAddress", addr)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address addr) returns()
func (_AddressBook *AddressBookSession) RemoveAddress(addr common.Address) (*types.Transaction, error) {
	return _AddressBook.Contract.RemoveAddress(&_AddressBook.TransactOpts, addr)
}

// RemoveAddress is a paid mutator transaction binding the contract method 0x4ba79dfe.
//
// Solidity: function removeAddress(address addr) returns()
func (_AddressBook *AddressBookTransactorSession) RemoveAddress(addr common.Address) (*types.Transaction, error) {
	return _AddressBook.Contract.RemoveAddress(&_AddressBook.TransactOpts, addr)
}
