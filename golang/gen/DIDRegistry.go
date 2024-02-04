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

// DIDRegistryMetaData contains all meta data concerning the DIDRegistry contract.
var DIDRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"NameIsUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoIdentityWithPrefix\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotTheOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"did_document_contract_address\",\"type\":\"address\"}],\"name\":\"addNewIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"did_document_contract_address\",\"type\":\"address\"}],\"name\":\"changeIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"changeOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_name\",\"type\":\"string\"}],\"name\":\"deleteIdentity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"prefix_name\",\"type\":\"string\"}],\"name\":\"getIdentity\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"identities\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"owners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"used_names\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610f22806100206000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806342ba518b1161005b57806342ba518b1461013d578063bf951c6814610159578063c5fce76714610175578063cf1066b21461019157610088565b806301e21afb1461008d5780630a29ae6f146100bd5780632b1fd995146100f15780633ed3612d14610121575b600080fd5b6100a760048036038101906100a29190610b06565b6101c1565b6040516100b49190610b90565b60405180910390f35b6100d760048036038101906100d29190610b06565b61020a565b6040516100e8959493929190610c5e565b60405180910390f35b61010b60048036038101906101069190610b06565b61041b565b6040516101189190610b90565b60405180910390f35b61013b60048036038101906101369190610ceb565b610464565b005b61015760048036038101906101529190610ceb565b6105bb565b005b610173600480360381019061016e9190610b06565b610700565b005b61018f600480360381019061018a9190610ceb565b61085f565b005b6101ab60048036038101906101a69190610b06565b610963565b6040516101b89190610d47565b60405180910390f35b6000818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60608060008060008015156001876040516102259190610d9e565b908152602001604051809103902060009054906101000a900460ff1615150361027a576040517fe5b702ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808760405161028b9190610d9e565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690506102c581610999565b156103b5578073ffffffffffffffffffffffffffffffffffffffff166336afc6fa6040518163ffffffff1660e01b8152600401600060405180830381865afa92505050801561033757506040513d6000823e3d601f19601f820116820180604052508101906103349190610e7d565b60015b61039c5760008060016040518060400160405280601081526020017f556e6b6e6f776e20436f6e747261637400000000000000000000000000000000815250929190604051806020016040528060008152509291909550955095509550955050610412565b8983838360009850985098509850985050505050610412565b60008060016040518060400160405280601681526020017f4e6f74206120636f6e74726163742061646472657373000000000000000000008152509291906040518060200160405280600081525092919095509550955095509550505b91939590929450565b6002818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6001826040516104749190610d9e565b908152602001604051809103902060009054906101000a900460ff16156104c7576040517f89f4febf00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b336002836040516104d89190610d9e565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000836040516105359190610d9e565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600180836040516105929190610d9e565b908152602001604051809103902060006101000a81548160ff0219169083151502179055505050565b3373ffffffffffffffffffffffffffffffffffffffff166002836040516105e29190610d9e565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461065e576040517f36b6b89500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8060028360405161066f9190610d9e565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000826040516106cb9190610d9e565b908152602001604051809103902060006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690555050565b3373ffffffffffffffffffffffffffffffffffffffff166002826040516107279190610d9e565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146107a3576040517f36b6b89500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006001826040516107b59190610d9e565b908152602001604051809103902060006101000a81548160ff0219169083151502179055506002816040516107ea9190610d9e565b908152602001604051809103902060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905560008160405161082b9190610d9e565b908152602001604051809103902060006101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550565b3373ffffffffffffffffffffffffffffffffffffffff166002836040516108869190610d9e565b908152602001604051809103902060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610902576040517f36b6b89500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b806000836040516109139190610d9e565b908152602001604051809103902060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6001818051602081018201805184825260208301602085012081835280955050505050506000915054906101000a900460ff1681565b600080823b905060008111915050919050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610a13826109ca565b810181811067ffffffffffffffff82111715610a3257610a316109db565b5b80604052505050565b6000610a456109ac565b9050610a518282610a0a565b919050565b600067ffffffffffffffff821115610a7157610a706109db565b5b610a7a826109ca565b9050602081019050919050565b82818337600083830152505050565b6000610aa9610aa484610a56565b610a3b565b905082815260208101848484011115610ac557610ac46109c5565b5b610ad0848285610a87565b509392505050565b600082601f830112610aed57610aec6109c0565b5b8135610afd848260208601610a96565b91505092915050565b600060208284031215610b1c57610b1b6109b6565b5b600082013567ffffffffffffffff811115610b3a57610b396109bb565b5b610b4684828501610ad8565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610b7a82610b4f565b9050919050565b610b8a81610b6f565b82525050565b6000602082019050610ba56000830184610b81565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610be5578082015181840152602081019050610bca565b60008484015250505050565b6000610bfc82610bab565b610c068185610bb6565b9350610c16818560208601610bc7565b610c1f816109ca565b840191505092915050565b6000819050919050565b610c3d81610c2a565b82525050565b60008115159050919050565b610c5881610c43565b82525050565b600060a0820190508181036000830152610c788188610bf1565b90508181036020830152610c8c8187610bf1565b9050610c9b6040830186610c34565b610ca86060830185610c4f565b610cb56080830184610c4f565b9695505050505050565b610cc881610b6f565b8114610cd357600080fd5b50565b600081359050610ce581610cbf565b92915050565b60008060408385031215610d0257610d016109b6565b5b600083013567ffffffffffffffff811115610d2057610d1f6109bb565b5b610d2c85828601610ad8565b9250506020610d3d85828601610cd6565b9150509250929050565b6000602082019050610d5c6000830184610c4f565b92915050565b600081905092915050565b6000610d7882610bab565b610d828185610d62565b9350610d92818560208601610bc7565b80840191505092915050565b6000610daa8284610d6d565b915081905092915050565b6000610dc8610dc384610a56565b610a3b565b905082815260208101848484011115610de457610de36109c5565b5b610def848285610bc7565b509392505050565b600082601f830112610e0c57610e0b6109c0565b5b8151610e1c848260208601610db5565b91505092915050565b610e2e81610c2a565b8114610e3957600080fd5b50565b600081519050610e4b81610e25565b92915050565b610e5a81610c43565b8114610e6557600080fd5b50565b600081519050610e7781610e51565b92915050565b600080600060608486031215610e9657610e956109b6565b5b600084015167ffffffffffffffff811115610eb457610eb36109bb565b5b610ec086828701610df7565b9350506020610ed186828701610e3c565b9250506040610ee286828701610e68565b915050925092509256fea264697066735822122010a32aed8745fecfb196a5b894de5127080466f3610a4638e1f2b9f03e0e942f64736f6c63430008100033",
}

// DIDRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DIDRegistryMetaData.ABI instead.
var DIDRegistryABI = DIDRegistryMetaData.ABI

// DIDRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIDRegistryMetaData.Bin instead.
var DIDRegistryBin = DIDRegistryMetaData.Bin

// DeployDIDRegistry deploys a new Ethereum contract, binding an instance of DIDRegistry to it.
func DeployDIDRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DIDRegistry, error) {
	parsed, err := DIDRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIDRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIDRegistry{DIDRegistryCaller: DIDRegistryCaller{contract: contract}, DIDRegistryTransactor: DIDRegistryTransactor{contract: contract}, DIDRegistryFilterer: DIDRegistryFilterer{contract: contract}}, nil
}

// DIDRegistry is an auto generated Go binding around an Ethereum contract.
type DIDRegistry struct {
	DIDRegistryCaller     // Read-only binding to the contract
	DIDRegistryTransactor // Write-only binding to the contract
	DIDRegistryFilterer   // Log filterer for contract events
}

// DIDRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIDRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIDRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIDRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIDRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIDRegistrySession struct {
	Contract     *DIDRegistry      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIDRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIDRegistryCallerSession struct {
	Contract *DIDRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DIDRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIDRegistryTransactorSession struct {
	Contract     *DIDRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DIDRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIDRegistryRaw struct {
	Contract *DIDRegistry // Generic contract binding to access the raw methods on
}

// DIDRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIDRegistryCallerRaw struct {
	Contract *DIDRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// DIDRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIDRegistryTransactorRaw struct {
	Contract *DIDRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIDRegistry creates a new instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistry(address common.Address, backend bind.ContractBackend) (*DIDRegistry, error) {
	contract, err := bindDIDRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIDRegistry{DIDRegistryCaller: DIDRegistryCaller{contract: contract}, DIDRegistryTransactor: DIDRegistryTransactor{contract: contract}, DIDRegistryFilterer: DIDRegistryFilterer{contract: contract}}, nil
}

// NewDIDRegistryCaller creates a new read-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryCaller(address common.Address, caller bind.ContractCaller) (*DIDRegistryCaller, error) {
	contract, err := bindDIDRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryCaller{contract: contract}, nil
}

// NewDIDRegistryTransactor creates a new write-only instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DIDRegistryTransactor, error) {
	contract, err := bindDIDRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryTransactor{contract: contract}, nil
}

// NewDIDRegistryFilterer creates a new log filterer instance of DIDRegistry, bound to a specific deployed contract.
func NewDIDRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DIDRegistryFilterer, error) {
	contract, err := bindDIDRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIDRegistryFilterer{contract: contract}, nil
}

// bindDIDRegistry binds a generic wrapper to an already deployed contract.
func bindDIDRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DIDRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.DIDRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DIDRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIDRegistry *DIDRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIDRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIDRegistry *DIDRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIDRegistry.Contract.contract.Transact(opts, method, params...)
}

// GetIdentity is a free data retrieval call binding the contract method 0x0a29ae6f.
//
// Solidity: function getIdentity(string prefix_name) view returns(string, string, uint256, bool, bool)
func (_DIDRegistry *DIDRegistryCaller) GetIdentity(opts *bind.CallOpts, prefix_name string) (string, string, *big.Int, bool, bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "getIdentity", prefix_name)

	if err != nil {
		return *new(string), *new(string), *new(*big.Int), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetIdentity is a free data retrieval call binding the contract method 0x0a29ae6f.
//
// Solidity: function getIdentity(string prefix_name) view returns(string, string, uint256, bool, bool)
func (_DIDRegistry *DIDRegistrySession) GetIdentity(prefix_name string) (string, string, *big.Int, bool, bool, error) {
	return _DIDRegistry.Contract.GetIdentity(&_DIDRegistry.CallOpts, prefix_name)
}

// GetIdentity is a free data retrieval call binding the contract method 0x0a29ae6f.
//
// Solidity: function getIdentity(string prefix_name) view returns(string, string, uint256, bool, bool)
func (_DIDRegistry *DIDRegistryCallerSession) GetIdentity(prefix_name string) (string, string, *big.Int, bool, bool, error) {
	return _DIDRegistry.Contract.GetIdentity(&_DIDRegistry.CallOpts, prefix_name)
}

// Identities is a free data retrieval call binding the contract method 0x01e21afb.
//
// Solidity: function identities(string ) view returns(address)
func (_DIDRegistry *DIDRegistryCaller) Identities(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "identities", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Identities is a free data retrieval call binding the contract method 0x01e21afb.
//
// Solidity: function identities(string ) view returns(address)
func (_DIDRegistry *DIDRegistrySession) Identities(arg0 string) (common.Address, error) {
	return _DIDRegistry.Contract.Identities(&_DIDRegistry.CallOpts, arg0)
}

// Identities is a free data retrieval call binding the contract method 0x01e21afb.
//
// Solidity: function identities(string ) view returns(address)
func (_DIDRegistry *DIDRegistryCallerSession) Identities(arg0 string) (common.Address, error) {
	return _DIDRegistry.Contract.Identities(&_DIDRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x2b1fd995.
//
// Solidity: function owners(string ) view returns(address)
func (_DIDRegistry *DIDRegistryCaller) Owners(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "owners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owners is a free data retrieval call binding the contract method 0x2b1fd995.
//
// Solidity: function owners(string ) view returns(address)
func (_DIDRegistry *DIDRegistrySession) Owners(arg0 string) (common.Address, error) {
	return _DIDRegistry.Contract.Owners(&_DIDRegistry.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x2b1fd995.
//
// Solidity: function owners(string ) view returns(address)
func (_DIDRegistry *DIDRegistryCallerSession) Owners(arg0 string) (common.Address, error) {
	return _DIDRegistry.Contract.Owners(&_DIDRegistry.CallOpts, arg0)
}

// UsedNames is a free data retrieval call binding the contract method 0xcf1066b2.
//
// Solidity: function used_names(string ) view returns(bool)
func (_DIDRegistry *DIDRegistryCaller) UsedNames(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _DIDRegistry.contract.Call(opts, &out, "used_names", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedNames is a free data retrieval call binding the contract method 0xcf1066b2.
//
// Solidity: function used_names(string ) view returns(bool)
func (_DIDRegistry *DIDRegistrySession) UsedNames(arg0 string) (bool, error) {
	return _DIDRegistry.Contract.UsedNames(&_DIDRegistry.CallOpts, arg0)
}

// UsedNames is a free data retrieval call binding the contract method 0xcf1066b2.
//
// Solidity: function used_names(string ) view returns(bool)
func (_DIDRegistry *DIDRegistryCallerSession) UsedNames(arg0 string) (bool, error) {
	return _DIDRegistry.Contract.UsedNames(&_DIDRegistry.CallOpts, arg0)
}

// AddNewIdentity is a paid mutator transaction binding the contract method 0x3ed3612d.
//
// Solidity: function addNewIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistryTransactor) AddNewIdentity(opts *bind.TransactOpts, prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "addNewIdentity", prefix_name, did_document_contract_address)
}

// AddNewIdentity is a paid mutator transaction binding the contract method 0x3ed3612d.
//
// Solidity: function addNewIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistrySession) AddNewIdentity(prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.AddNewIdentity(&_DIDRegistry.TransactOpts, prefix_name, did_document_contract_address)
}

// AddNewIdentity is a paid mutator transaction binding the contract method 0x3ed3612d.
//
// Solidity: function addNewIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) AddNewIdentity(prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.AddNewIdentity(&_DIDRegistry.TransactOpts, prefix_name, did_document_contract_address)
}

// ChangeIdentity is a paid mutator transaction binding the contract method 0xc5fce767.
//
// Solidity: function changeIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistryTransactor) ChangeIdentity(opts *bind.TransactOpts, prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "changeIdentity", prefix_name, did_document_contract_address)
}

// ChangeIdentity is a paid mutator transaction binding the contract method 0xc5fce767.
//
// Solidity: function changeIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistrySession) ChangeIdentity(prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ChangeIdentity(&_DIDRegistry.TransactOpts, prefix_name, did_document_contract_address)
}

// ChangeIdentity is a paid mutator transaction binding the contract method 0xc5fce767.
//
// Solidity: function changeIdentity(string prefix_name, address did_document_contract_address) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) ChangeIdentity(prefix_name string, did_document_contract_address common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ChangeIdentity(&_DIDRegistry.TransactOpts, prefix_name, did_document_contract_address)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x42ba518b.
//
// Solidity: function changeOwnership(string prefix_name, address newOwner) returns()
func (_DIDRegistry *DIDRegistryTransactor) ChangeOwnership(opts *bind.TransactOpts, prefix_name string, newOwner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "changeOwnership", prefix_name, newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x42ba518b.
//
// Solidity: function changeOwnership(string prefix_name, address newOwner) returns()
func (_DIDRegistry *DIDRegistrySession) ChangeOwnership(prefix_name string, newOwner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ChangeOwnership(&_DIDRegistry.TransactOpts, prefix_name, newOwner)
}

// ChangeOwnership is a paid mutator transaction binding the contract method 0x42ba518b.
//
// Solidity: function changeOwnership(string prefix_name, address newOwner) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) ChangeOwnership(prefix_name string, newOwner common.Address) (*types.Transaction, error) {
	return _DIDRegistry.Contract.ChangeOwnership(&_DIDRegistry.TransactOpts, prefix_name, newOwner)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xbf951c68.
//
// Solidity: function deleteIdentity(string prefix_name) returns()
func (_DIDRegistry *DIDRegistryTransactor) DeleteIdentity(opts *bind.TransactOpts, prefix_name string) (*types.Transaction, error) {
	return _DIDRegistry.contract.Transact(opts, "deleteIdentity", prefix_name)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xbf951c68.
//
// Solidity: function deleteIdentity(string prefix_name) returns()
func (_DIDRegistry *DIDRegistrySession) DeleteIdentity(prefix_name string) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DeleteIdentity(&_DIDRegistry.TransactOpts, prefix_name)
}

// DeleteIdentity is a paid mutator transaction binding the contract method 0xbf951c68.
//
// Solidity: function deleteIdentity(string prefix_name) returns()
func (_DIDRegistry *DIDRegistryTransactorSession) DeleteIdentity(prefix_name string) (*types.Transaction, error) {
	return _DIDRegistry.Contract.DeleteIdentity(&_DIDRegistry.TransactOpts, prefix_name)
}
