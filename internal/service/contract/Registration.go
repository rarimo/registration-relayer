// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// SparseMerkleTreeNode is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeNode struct {
	NodeType   uint8
	ChildLeft  uint64
	ChildRight uint64
	NodeHash   [32]byte
	Key        [32]byte
	Value      [32]byte
}

// SparseMerkleTreeProof is an auto generated low-level Go binding around an user-defined struct.
type SparseMerkleTreeProof struct {
	Root         [32]byte
	Siblings     [][32]byte
	Existence    bool
	Key          [32]byte
	Value        [32]byte
	AuxExistence bool
	AuxKey       [32]byte
	AuxValue     [32]byte
}

// VerifierHelperProofPoints is an auto generated low-level Go binding around an user-defined struct.
type VerifierHelperProofPoints struct {
	A [2]*big.Int
	B [2][2]*big.Int
	C [2]*big.Int
}

// RegistrationMetaData contains all meta data concerning the Registration contract.
var RegistrationMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hashedRSAKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hashedInternalKey\",\"type\":\"bytes32\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"E\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"treeHeight_\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifier_\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"icaoMasterTreeMerkleRoot_\",\"type\":\"bytes32\"}],\"name\":\"__Registration_init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getNodeByKey\",\"outputs\":[{\"components\":[{\"internalType\":\"enumSparseMerkleTree.NodeType\",\"name\":\"nodeType\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"childLeft\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"childRight\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"nodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Node\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key_\",\"type\":\"bytes32\"}],\"name\":\"getProof\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32[]\",\"name\":\"siblings\",\"type\":\"bytes32[]\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"auxKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"auxValue\",\"type\":\"bytes32\"}],\"internalType\":\"structSparseMerkleTree.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"hashedRSAKeyToInternalKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"icaoMasterTreeMerkleRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"internalKeyToHashedRSAKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hashedInternalKey_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"s_\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"n_\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"a\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2][2]\",\"name\":\"b\",\"type\":\"uint256[2][2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"c\",\"type\":\"uint256[2]\"}],\"internalType\":\"structVerifierHelper.ProofPoints\",\"name\":\"zkPoints_\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"group1Hash_\",\"type\":\"uint256\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RegistrationABI is the input ABI used to generate the binding from.
// Deprecated: Use RegistrationMetaData.ABI instead.
var RegistrationABI = RegistrationMetaData.ABI

// Registration is an auto generated Go binding around an Ethereum contract.
type Registration struct {
	RegistrationCaller     // Read-only binding to the contract
	RegistrationTransactor // Write-only binding to the contract
	RegistrationFilterer   // Log filterer for contract events
}

// RegistrationCaller is an auto generated read-only Go binding around an Ethereum contract.
type RegistrationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RegistrationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RegistrationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RegistrationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RegistrationSession struct {
	Contract     *Registration     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RegistrationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RegistrationCallerSession struct {
	Contract *RegistrationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// RegistrationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RegistrationTransactorSession struct {
	Contract     *RegistrationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// RegistrationRaw is an auto generated low-level Go binding around an Ethereum contract.
type RegistrationRaw struct {
	Contract *Registration // Generic contract binding to access the raw methods on
}

// RegistrationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RegistrationCallerRaw struct {
	Contract *RegistrationCaller // Generic read-only contract binding to access the raw methods on
}

// RegistrationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RegistrationTransactorRaw struct {
	Contract *RegistrationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRegistration creates a new instance of Registration, bound to a specific deployed contract.
func NewRegistration(address common.Address, backend bind.ContractBackend) (*Registration, error) {
	contract, err := bindRegistration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Registration{RegistrationCaller: RegistrationCaller{contract: contract}, RegistrationTransactor: RegistrationTransactor{contract: contract}, RegistrationFilterer: RegistrationFilterer{contract: contract}}, nil
}

// NewRegistrationCaller creates a new read-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationCaller(address common.Address, caller bind.ContractCaller) (*RegistrationCaller, error) {
	contract, err := bindRegistration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationCaller{contract: contract}, nil
}

// NewRegistrationTransactor creates a new write-only instance of Registration, bound to a specific deployed contract.
func NewRegistrationTransactor(address common.Address, transactor bind.ContractTransactor) (*RegistrationTransactor, error) {
	contract, err := bindRegistration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RegistrationTransactor{contract: contract}, nil
}

// NewRegistrationFilterer creates a new log filterer instance of Registration, bound to a specific deployed contract.
func NewRegistrationFilterer(address common.Address, filterer bind.ContractFilterer) (*RegistrationFilterer, error) {
	contract, err := bindRegistration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RegistrationFilterer{contract: contract}, nil
}

// bindRegistration binds a generic wrapper to an already deployed contract.
func bindRegistration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RegistrationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.RegistrationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Registration *RegistrationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Registration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Registration *RegistrationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Registration *RegistrationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Registration.Contract.contract.Transact(opts, method, params...)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(uint256)
func (_Registration *RegistrationCaller) E(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "E")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(uint256)
func (_Registration *RegistrationSession) E() (*big.Int, error) {
	return _Registration.Contract.E(&_Registration.CallOpts)
}

// E is a free data retrieval call binding the contract method 0x92bbf6e8.
//
// Solidity: function E() view returns(uint256)
func (_Registration *RegistrationCallerSession) E() (*big.Int, error) {
	return _Registration.Contract.E(&_Registration.CallOpts)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationCaller) GetNodeByKey(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeNode, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getNodeByKey", key_)

	if err != nil {
		return *new(SparseMerkleTreeNode), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeNode)).(*SparseMerkleTreeNode)

	return out0, err

}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Registration.Contract.GetNodeByKey(&_Registration.CallOpts, key_)
}

// GetNodeByKey is a free data retrieval call binding the contract method 0x083a8580.
//
// Solidity: function getNodeByKey(bytes32 key_) view returns((uint8,uint64,uint64,bytes32,bytes32,bytes32))
func (_Registration *RegistrationCallerSession) GetNodeByKey(key_ [32]byte) (SparseMerkleTreeNode, error) {
	return _Registration.Contract.GetNodeByKey(&_Registration.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationCaller) GetProof(opts *bind.CallOpts, key_ [32]byte) (SparseMerkleTreeProof, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getProof", key_)

	if err != nil {
		return *new(SparseMerkleTreeProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SparseMerkleTreeProof)).(*SparseMerkleTreeProof)

	return out0, err

}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Registration.Contract.GetProof(&_Registration.CallOpts, key_)
}

// GetProof is a free data retrieval call binding the contract method 0x1b80bb3a.
//
// Solidity: function getProof(bytes32 key_) view returns((bytes32,bytes32[],bool,bytes32,bytes32,bool,bytes32,bytes32))
func (_Registration *RegistrationCallerSession) GetProof(key_ [32]byte) (SparseMerkleTreeProof, error) {
	return _Registration.Contract.GetProof(&_Registration.CallOpts, key_)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationCaller) GetRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationSession) GetRoot() ([32]byte, error) {
	return _Registration.Contract.GetRoot(&_Registration.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(bytes32)
func (_Registration *RegistrationCallerSession) GetRoot() ([32]byte, error) {
	return _Registration.Contract.GetRoot(&_Registration.CallOpts)
}

// HashedRSAKeyToInternalKey is a free data retrieval call binding the contract method 0x1c36d26a.
//
// Solidity: function hashedRSAKeyToInternalKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationCaller) HashedRSAKeyToInternalKey(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "hashedRSAKeyToInternalKey", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashedRSAKeyToInternalKey is a free data retrieval call binding the contract method 0x1c36d26a.
//
// Solidity: function hashedRSAKeyToInternalKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationSession) HashedRSAKeyToInternalKey(arg0 [32]byte) ([32]byte, error) {
	return _Registration.Contract.HashedRSAKeyToInternalKey(&_Registration.CallOpts, arg0)
}

// HashedRSAKeyToInternalKey is a free data retrieval call binding the contract method 0x1c36d26a.
//
// Solidity: function hashedRSAKeyToInternalKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationCallerSession) HashedRSAKeyToInternalKey(arg0 [32]byte) ([32]byte, error) {
	return _Registration.Contract.HashedRSAKeyToInternalKey(&_Registration.CallOpts, arg0)
}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationCaller) IcaoMasterTreeMerkleRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "icaoMasterTreeMerkleRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationSession) IcaoMasterTreeMerkleRoot() ([32]byte, error) {
	return _Registration.Contract.IcaoMasterTreeMerkleRoot(&_Registration.CallOpts)
}

// IcaoMasterTreeMerkleRoot is a free data retrieval call binding the contract method 0x28093985.
//
// Solidity: function icaoMasterTreeMerkleRoot() view returns(bytes32)
func (_Registration *RegistrationCallerSession) IcaoMasterTreeMerkleRoot() ([32]byte, error) {
	return _Registration.Contract.IcaoMasterTreeMerkleRoot(&_Registration.CallOpts)
}

// InternalKeyToHashedRSAKey is a free data retrieval call binding the contract method 0xcde00543.
//
// Solidity: function internalKeyToHashedRSAKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationCaller) InternalKeyToHashedRSAKey(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "internalKeyToHashedRSAKey", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// InternalKeyToHashedRSAKey is a free data retrieval call binding the contract method 0xcde00543.
//
// Solidity: function internalKeyToHashedRSAKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationSession) InternalKeyToHashedRSAKey(arg0 [32]byte) ([32]byte, error) {
	return _Registration.Contract.InternalKeyToHashedRSAKey(&_Registration.CallOpts, arg0)
}

// InternalKeyToHashedRSAKey is a free data retrieval call binding the contract method 0xcde00543.
//
// Solidity: function internalKeyToHashedRSAKey(bytes32 ) view returns(bytes32)
func (_Registration *RegistrationCallerSession) InternalKeyToHashedRSAKey(arg0 [32]byte) ([32]byte, error) {
	return _Registration.Contract.InternalKeyToHashedRSAKey(&_Registration.CallOpts, arg0)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Registration *RegistrationCaller) Verifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Registration.contract.Call(opts, &out, "verifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Registration *RegistrationSession) Verifier() (common.Address, error) {
	return _Registration.Contract.Verifier(&_Registration.CallOpts)
}

// Verifier is a free data retrieval call binding the contract method 0x2b7ac3f3.
//
// Solidity: function verifier() view returns(address)
func (_Registration *RegistrationCallerSession) Verifier() (common.Address, error) {
	return _Registration.Contract.Verifier(&_Registration.CallOpts)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address verifier_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationTransactor) RegistrationInit(opts *bind.TransactOpts, treeHeight_ *big.Int, verifier_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "__Registration_init", treeHeight_, verifier_, icaoMasterTreeMerkleRoot_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address verifier_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationSession) RegistrationInit(treeHeight_ *big.Int, verifier_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationInit(&_Registration.TransactOpts, treeHeight_, verifier_, icaoMasterTreeMerkleRoot_)
}

// RegistrationInit is a paid mutator transaction binding the contract method 0x759bf7e1.
//
// Solidity: function __Registration_init(uint256 treeHeight_, address verifier_, bytes32 icaoMasterTreeMerkleRoot_) returns()
func (_Registration *RegistrationTransactorSession) RegistrationInit(treeHeight_ *big.Int, verifier_ common.Address, icaoMasterTreeMerkleRoot_ [32]byte) (*types.Transaction, error) {
	return _Registration.Contract.RegistrationInit(&_Registration.TransactOpts, treeHeight_, verifier_, icaoMasterTreeMerkleRoot_)
}

// Register is a paid mutator transaction binding the contract method 0x05433c72.
//
// Solidity: function register(uint256 hashedInternalKey_, bytes s_, bytes n_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 group1Hash_) returns()
func (_Registration *RegistrationTransactor) Register(opts *bind.TransactOpts, hashedInternalKey_ *big.Int, s_ []byte, n_ []byte, zkPoints_ VerifierHelperProofPoints, group1Hash_ *big.Int) (*types.Transaction, error) {
	return _Registration.contract.Transact(opts, "register", hashedInternalKey_, s_, n_, zkPoints_, group1Hash_)
}

// Register is a paid mutator transaction binding the contract method 0x05433c72.
//
// Solidity: function register(uint256 hashedInternalKey_, bytes s_, bytes n_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 group1Hash_) returns()
func (_Registration *RegistrationSession) Register(hashedInternalKey_ *big.Int, s_ []byte, n_ []byte, zkPoints_ VerifierHelperProofPoints, group1Hash_ *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, hashedInternalKey_, s_, n_, zkPoints_, group1Hash_)
}

// Register is a paid mutator transaction binding the contract method 0x05433c72.
//
// Solidity: function register(uint256 hashedInternalKey_, bytes s_, bytes n_, (uint256[2],uint256[2][2],uint256[2]) zkPoints_, uint256 group1Hash_) returns()
func (_Registration *RegistrationTransactorSession) Register(hashedInternalKey_ *big.Int, s_ []byte, n_ []byte, zkPoints_ VerifierHelperProofPoints, group1Hash_ *big.Int) (*types.Transaction, error) {
	return _Registration.Contract.Register(&_Registration.TransactOpts, hashedInternalKey_, s_, n_, zkPoints_, group1Hash_)
}

// RegistrationInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Registration contract.
type RegistrationInitializedIterator struct {
	Event *RegistrationInitialized // Event containing the contract specifics and raw log

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
func (it *RegistrationInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationInitialized)
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
		it.Event = new(RegistrationInitialized)
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
func (it *RegistrationInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationInitialized represents a Initialized event raised by the Registration contract.
type RegistrationInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registration *RegistrationFilterer) FilterInitialized(opts *bind.FilterOpts) (*RegistrationInitializedIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &RegistrationInitializedIterator{contract: _Registration.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Registration *RegistrationFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *RegistrationInitialized) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationInitialized)
				if err := _Registration.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Registration *RegistrationFilterer) ParseInitialized(log types.Log) (*RegistrationInitialized, error) {
	event := new(RegistrationInitialized)
	if err := _Registration.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RegistrationRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Registration contract.
type RegistrationRegisteredIterator struct {
	Event *RegistrationRegistered // Event containing the contract specifics and raw log

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
func (it *RegistrationRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RegistrationRegistered)
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
		it.Event = new(RegistrationRegistered)
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
func (it *RegistrationRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RegistrationRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RegistrationRegistered represents a Registered event raised by the Registration contract.
type RegistrationRegistered struct {
	HashedRSAKey      [32]byte
	HashedInternalKey [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 hashedRSAKey, bytes32 hashedInternalKey)
func (_Registration *RegistrationFilterer) FilterRegistered(opts *bind.FilterOpts) (*RegistrationRegisteredIterator, error) {

	logs, sub, err := _Registration.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RegistrationRegisteredIterator{contract: _Registration.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 hashedRSAKey, bytes32 hashedInternalKey)
func (_Registration *RegistrationFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RegistrationRegistered) (event.Subscription, error) {

	logs, sub, err := _Registration.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RegistrationRegistered)
				if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x8828766b1a3f99eda4bad24a0979129565597036bac9a94cdccd1d1e6ef17eb2.
//
// Solidity: event Registered(bytes32 hashedRSAKey, bytes32 hashedInternalKey)
func (_Registration *RegistrationFilterer) ParseRegistered(log types.Log) (*RegistrationRegistered, error) {
	event := new(RegistrationRegistered)
	if err := _Registration.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
