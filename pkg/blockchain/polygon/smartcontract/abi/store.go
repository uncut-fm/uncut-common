// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

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

// StoreStoreNFT is an auto generated low-level Go binding around an user-defined struct.
type StoreStoreNFT struct {
	StoreId      *big.Int
	NftContract  common.Address
	TokenId      *big.Int
	Creator      common.Address
	Currency     common.Address
	ListingPrice *big.Int
	ListingFees  *big.Int
	Balance      *big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"}],\"name\":\"CollectionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"listingCurrency\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"listingFees\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"minterData\",\"type\":\"string\"}],\"name\":\"NFTMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldListingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newListingPrice\",\"type\":\"uint256\"}],\"name\":\"NFTPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"airdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"available\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"burnRemainings\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"buyAndDeliver\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newPrice\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"showId\",\"type\":\"uint256\"}],\"name\":\"createCollection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"getByStoreId\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"listingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structStore.StoreNFT\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoyaltiesWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"royaltiesWallet\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"feesWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"liquidateStore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listAll\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"nftContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"listingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"internalType\":\"structStore.StoreNFT[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collectionAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"listingCurrency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"listingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingFees\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"royaltiesPct\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"minterData\",\"type\":\"string\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"nftContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"nftTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRoyaltiesAddress\",\"type\":\"address\"}],\"name\":\"setRoyaltiesWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"storeId\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Store.Contract.DEFAULTADMINROLE(&_Store.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Store.Contract.DEFAULTADMINROLE(&_Store.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreSession) MINTERROLE() ([32]byte, error) {
	return _Store.Contract.MINTERROLE(&_Store.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) MINTERROLE() ([32]byte, error) {
	return _Store.Contract.MINTERROLE(&_Store.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Store *StoreCaller) OWNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "OWNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Store *StoreSession) OWNERROLE() ([32]byte, error) {
	return _Store.Contract.OWNERROLE(&_Store.CallOpts)
}

// OWNERROLE is a free data retrieval call binding the contract method 0xe58378bb.
//
// Solidity: function OWNER_ROLE() view returns(bytes32)
func (_Store *StoreCallerSession) OWNERROLE() ([32]byte, error) {
	return _Store.Contract.OWNERROLE(&_Store.CallOpts)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 storeId) view returns(uint256)
func (_Store *StoreCaller) Available(opts *bind.CallOpts, storeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "available", storeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 storeId) view returns(uint256)
func (_Store *StoreSession) Available(storeId *big.Int) (*big.Int, error) {
	return _Store.Contract.Available(&_Store.CallOpts, storeId)
}

// Available is a free data retrieval call binding the contract method 0x96e494e8.
//
// Solidity: function available(uint256 storeId) view returns(uint256)
func (_Store *StoreCallerSession) Available(storeId *big.Int) (*big.Int, error) {
	return _Store.Contract.Available(&_Store.CallOpts, storeId)
}

// GetByStoreId is a free data retrieval call binding the contract method 0x8c55d947.
//
// Solidity: function getByStoreId(uint256 storeId) view returns((uint256,address,uint256,address,address,uint256,uint256,uint256))
func (_Store *StoreCaller) GetByStoreId(opts *bind.CallOpts, storeId *big.Int) (StoreStoreNFT, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getByStoreId", storeId)

	if err != nil {
		return *new(StoreStoreNFT), err
	}

	out0 := *abi.ConvertType(out[0], new(StoreStoreNFT)).(*StoreStoreNFT)

	return out0, err

}

// GetByStoreId is a free data retrieval call binding the contract method 0x8c55d947.
//
// Solidity: function getByStoreId(uint256 storeId) view returns((uint256,address,uint256,address,address,uint256,uint256,uint256))
func (_Store *StoreSession) GetByStoreId(storeId *big.Int) (StoreStoreNFT, error) {
	return _Store.Contract.GetByStoreId(&_Store.CallOpts, storeId)
}

// GetByStoreId is a free data retrieval call binding the contract method 0x8c55d947.
//
// Solidity: function getByStoreId(uint256 storeId) view returns((uint256,address,uint256,address,address,uint256,uint256,uint256))
func (_Store *StoreCallerSession) GetByStoreId(storeId *big.Int) (StoreStoreNFT, error) {
	return _Store.Contract.GetByStoreId(&_Store.CallOpts, storeId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Store.Contract.GetRoleAdmin(&_Store.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Store *StoreCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Store.Contract.GetRoleAdmin(&_Store.CallOpts, role)
}

// GetRoyaltiesWallet is a free data retrieval call binding the contract method 0xaf4b0877.
//
// Solidity: function getRoyaltiesWallet() view returns(address royaltiesWallet)
func (_Store *StoreCaller) GetRoyaltiesWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getRoyaltiesWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoyaltiesWallet is a free data retrieval call binding the contract method 0xaf4b0877.
//
// Solidity: function getRoyaltiesWallet() view returns(address royaltiesWallet)
func (_Store *StoreSession) GetRoyaltiesWallet() (common.Address, error) {
	return _Store.Contract.GetRoyaltiesWallet(&_Store.CallOpts)
}

// GetRoyaltiesWallet is a free data retrieval call binding the contract method 0xaf4b0877.
//
// Solidity: function getRoyaltiesWallet() view returns(address royaltiesWallet)
func (_Store *StoreCallerSession) GetRoyaltiesWallet() (common.Address, error) {
	return _Store.Contract.GetRoyaltiesWallet(&_Store.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Store.Contract.HasRole(&_Store.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Store *StoreCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Store.Contract.HasRole(&_Store.CallOpts, role, account)
}

// ListAll is a free data retrieval call binding the contract method 0x78a9eeed.
//
// Solidity: function listAll() view returns((uint256,address,uint256,address,address,uint256,uint256,uint256)[])
func (_Store *StoreCaller) ListAll(opts *bind.CallOpts) ([]StoreStoreNFT, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "listAll")

	if err != nil {
		return *new([]StoreStoreNFT), err
	}

	out0 := *abi.ConvertType(out[0], new([]StoreStoreNFT)).(*[]StoreStoreNFT)

	return out0, err

}

// ListAll is a free data retrieval call binding the contract method 0x78a9eeed.
//
// Solidity: function listAll() view returns((uint256,address,uint256,address,address,uint256,uint256,uint256)[])
func (_Store *StoreSession) ListAll() ([]StoreStoreNFT, error) {
	return _Store.Contract.ListAll(&_Store.CallOpts)
}

// ListAll is a free data retrieval call binding the contract method 0x78a9eeed.
//
// Solidity: function listAll() view returns((uint256,address,uint256,address,address,uint256,uint256,uint256)[])
func (_Store *StoreCallerSession) ListAll() ([]StoreStoreNFT, error) {
	return _Store.Contract.ListAll(&_Store.CallOpts)
}

// NftContractAddress is a free data retrieval call binding the contract method 0x6b687579.
//
// Solidity: function nftContractAddress(uint256 storeId) view returns(address)
func (_Store *StoreCaller) NftContractAddress(opts *bind.CallOpts, storeId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "nftContractAddress", storeId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContractAddress is a free data retrieval call binding the contract method 0x6b687579.
//
// Solidity: function nftContractAddress(uint256 storeId) view returns(address)
func (_Store *StoreSession) NftContractAddress(storeId *big.Int) (common.Address, error) {
	return _Store.Contract.NftContractAddress(&_Store.CallOpts, storeId)
}

// NftContractAddress is a free data retrieval call binding the contract method 0x6b687579.
//
// Solidity: function nftContractAddress(uint256 storeId) view returns(address)
func (_Store *StoreCallerSession) NftContractAddress(storeId *big.Int) (common.Address, error) {
	return _Store.Contract.NftContractAddress(&_Store.CallOpts, storeId)
}

// NftTokenId is a free data retrieval call binding the contract method 0x2d51de72.
//
// Solidity: function nftTokenId(uint256 storeId) view returns(uint256)
func (_Store *StoreCaller) NftTokenId(opts *bind.CallOpts, storeId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "nftTokenId", storeId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftTokenId is a free data retrieval call binding the contract method 0x2d51de72.
//
// Solidity: function nftTokenId(uint256 storeId) view returns(uint256)
func (_Store *StoreSession) NftTokenId(storeId *big.Int) (*big.Int, error) {
	return _Store.Contract.NftTokenId(&_Store.CallOpts, storeId)
}

// NftTokenId is a free data retrieval call binding the contract method 0x2d51de72.
//
// Solidity: function nftTokenId(uint256 storeId) view returns(uint256)
func (_Store *StoreCallerSession) NftTokenId(storeId *big.Int) (*big.Int, error) {
	return _Store.Contract.NftTokenId(&_Store.CallOpts, storeId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreSession) Paused() (bool, error) {
	return _Store.Contract.Paused(&_Store.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Store *StoreCallerSession) Paused() (bool, error) {
	return _Store.Contract.Paused(&_Store.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 storeId) view returns(string)
func (_Store *StoreCaller) Uri(opts *bind.CallOpts, storeId *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "uri", storeId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 storeId) view returns(string)
func (_Store *StoreSession) Uri(storeId *big.Int) (string, error) {
	return _Store.Contract.Uri(&_Store.CallOpts, storeId)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 storeId) view returns(string)
func (_Store *StoreCallerSession) Uri(storeId *big.Int) (string, error) {
	return _Store.Contract.Uri(&_Store.CallOpts, storeId)
}

// Airdrop is a paid mutator transaction binding the contract method 0xd5516e7f.
//
// Solidity: function airdrop(uint256 storeId, address[] recipients, uint256[] amounts) returns()
func (_Store *StoreTransactor) Airdrop(opts *bind.TransactOpts, storeId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "airdrop", storeId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xd5516e7f.
//
// Solidity: function airdrop(uint256 storeId, address[] recipients, uint256[] amounts) returns()
func (_Store *StoreSession) Airdrop(storeId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.Airdrop(&_Store.TransactOpts, storeId, recipients, amounts)
}

// Airdrop is a paid mutator transaction binding the contract method 0xd5516e7f.
//
// Solidity: function airdrop(uint256 storeId, address[] recipients, uint256[] amounts) returns()
func (_Store *StoreTransactorSession) Airdrop(storeId *big.Int, recipients []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.Airdrop(&_Store.TransactOpts, storeId, recipients, amounts)
}

// BurnRemainings is a paid mutator transaction binding the contract method 0x89a86072.
//
// Solidity: function burnRemainings(uint256 storeId) returns()
func (_Store *StoreTransactor) BurnRemainings(opts *bind.TransactOpts, storeId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burnRemainings", storeId)
}

// BurnRemainings is a paid mutator transaction binding the contract method 0x89a86072.
//
// Solidity: function burnRemainings(uint256 storeId) returns()
func (_Store *StoreSession) BurnRemainings(storeId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnRemainings(&_Store.TransactOpts, storeId)
}

// BurnRemainings is a paid mutator transaction binding the contract method 0x89a86072.
//
// Solidity: function burnRemainings(uint256 storeId) returns()
func (_Store *StoreTransactorSession) BurnRemainings(storeId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnRemainings(&_Store.TransactOpts, storeId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 storeId) returns()
func (_Store *StoreTransactor) Buy(opts *bind.TransactOpts, storeId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "buy", storeId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 storeId) returns()
func (_Store *StoreSession) Buy(storeId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Buy(&_Store.TransactOpts, storeId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 storeId) returns()
func (_Store *StoreTransactorSession) Buy(storeId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Buy(&_Store.TransactOpts, storeId)
}

// BuyAndDeliver is a paid mutator transaction binding the contract method 0x8c8cd159.
//
// Solidity: function buyAndDeliver(uint256 storeId, address recipient, uint256 amount) returns()
func (_Store *StoreTransactor) BuyAndDeliver(opts *bind.TransactOpts, storeId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "buyAndDeliver", storeId, recipient, amount)
}

// BuyAndDeliver is a paid mutator transaction binding the contract method 0x8c8cd159.
//
// Solidity: function buyAndDeliver(uint256 storeId, address recipient, uint256 amount) returns()
func (_Store *StoreSession) BuyAndDeliver(storeId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BuyAndDeliver(&_Store.TransactOpts, storeId, recipient, amount)
}

// BuyAndDeliver is a paid mutator transaction binding the contract method 0x8c8cd159.
//
// Solidity: function buyAndDeliver(uint256 storeId, address recipient, uint256 amount) returns()
func (_Store *StoreTransactorSession) BuyAndDeliver(storeId *big.Int, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BuyAndDeliver(&_Store.TransactOpts, storeId, recipient, amount)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 storeId, uint256 newPrice) returns()
func (_Store *StoreTransactor) ChangePrice(opts *bind.TransactOpts, storeId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "changePrice", storeId, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 storeId, uint256 newPrice) returns()
func (_Store *StoreSession) ChangePrice(storeId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangePrice(&_Store.TransactOpts, storeId, newPrice)
}

// ChangePrice is a paid mutator transaction binding the contract method 0xb3de019c.
//
// Solidity: function changePrice(uint256 storeId, uint256 newPrice) returns()
func (_Store *StoreTransactorSession) ChangePrice(storeId *big.Int, newPrice *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ChangePrice(&_Store.TransactOpts, storeId, newPrice)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0a0e1d4a.
//
// Solidity: function createCollection(string name, address creator, uint256 showId) returns()
func (_Store *StoreTransactor) CreateCollection(opts *bind.TransactOpts, name string, creator common.Address, showId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "createCollection", name, creator, showId)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0a0e1d4a.
//
// Solidity: function createCollection(string name, address creator, uint256 showId) returns()
func (_Store *StoreSession) CreateCollection(name string, creator common.Address, showId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.CreateCollection(&_Store.TransactOpts, name, creator, showId)
}

// CreateCollection is a paid mutator transaction binding the contract method 0x0a0e1d4a.
//
// Solidity: function createCollection(string name, address creator, uint256 showId) returns()
func (_Store *StoreTransactorSession) CreateCollection(name string, creator common.Address, showId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.CreateCollection(&_Store.TransactOpts, name, creator, showId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GrantRole(&_Store.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.GrantRole(&_Store.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address feesWallet) returns()
func (_Store *StoreTransactor) Initialize(opts *bind.TransactOpts, owner common.Address, feesWallet common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initialize", owner, feesWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address feesWallet) returns()
func (_Store *StoreSession) Initialize(owner common.Address, feesWallet common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, owner, feesWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner, address feesWallet) returns()
func (_Store *StoreTransactorSession) Initialize(owner common.Address, feesWallet common.Address) (*types.Transaction, error) {
	return _Store.Contract.Initialize(&_Store.TransactOpts, owner, feesWallet)
}

// LiquidateStore is a paid mutator transaction binding the contract method 0xb332df4a.
//
// Solidity: function liquidateStore(address to) returns()
func (_Store *StoreTransactor) LiquidateStore(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "liquidateStore", to)
}

// LiquidateStore is a paid mutator transaction binding the contract method 0xb332df4a.
//
// Solidity: function liquidateStore(address to) returns()
func (_Store *StoreSession) LiquidateStore(to common.Address) (*types.Transaction, error) {
	return _Store.Contract.LiquidateStore(&_Store.TransactOpts, to)
}

// LiquidateStore is a paid mutator transaction binding the contract method 0xb332df4a.
//
// Solidity: function liquidateStore(address to) returns()
func (_Store *StoreTransactorSession) LiquidateStore(to common.Address) (*types.Transaction, error) {
	return _Store.Contract.LiquidateStore(&_Store.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0xc65d54af.
//
// Solidity: function mint(address collectionAddress, string tokenURI, uint256 supply, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, uint32 royaltiesPct, string minterData) returns()
func (_Store *StoreTransactor) Mint(opts *bind.TransactOpts, collectionAddress common.Address, tokenURI string, supply *big.Int, creator common.Address, listingCurrency common.Address, listingPrice *big.Int, listingFees *big.Int, royaltiesPct uint32, minterData string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mint", collectionAddress, tokenURI, supply, creator, listingCurrency, listingPrice, listingFees, royaltiesPct, minterData)
}

// Mint is a paid mutator transaction binding the contract method 0xc65d54af.
//
// Solidity: function mint(address collectionAddress, string tokenURI, uint256 supply, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, uint32 royaltiesPct, string minterData) returns()
func (_Store *StoreSession) Mint(collectionAddress common.Address, tokenURI string, supply *big.Int, creator common.Address, listingCurrency common.Address, listingPrice *big.Int, listingFees *big.Int, royaltiesPct uint32, minterData string) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, collectionAddress, tokenURI, supply, creator, listingCurrency, listingPrice, listingFees, royaltiesPct, minterData)
}

// Mint is a paid mutator transaction binding the contract method 0xc65d54af.
//
// Solidity: function mint(address collectionAddress, string tokenURI, uint256 supply, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, uint32 royaltiesPct, string minterData) returns()
func (_Store *StoreTransactorSession) Mint(collectionAddress common.Address, tokenURI string, supply *big.Int, creator common.Address, listingCurrency common.Address, listingPrice *big.Int, listingFees *big.Int, royaltiesPct uint32, minterData string) (*types.Transaction, error) {
	return _Store.Contract.Mint(&_Store.TransactOpts, collectionAddress, tokenURI, supply, creator, listingCurrency, listingPrice, listingFees, royaltiesPct, minterData)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Store *StoreTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Store *StoreSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC1155BatchReceived(&_Store.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_Store *StoreTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC1155BatchReceived(&_Store.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Store *StoreTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Store *StoreSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC1155Received(&_Store.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_Store *StoreTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _Store.Contract.OnERC1155Received(&_Store.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreSession) Pause() (*types.Transaction, error) {
	return _Store.Contract.Pause(&_Store.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Store *StoreTransactorSession) Pause() (*types.Transaction, error) {
	return _Store.Contract.Pause(&_Store.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RenounceRole(&_Store.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RenounceRole(&_Store.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RevokeRole(&_Store.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Store *StoreTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Store.Contract.RevokeRole(&_Store.TransactOpts, role, account)
}

// SetRoyaltiesWallet is a paid mutator transaction binding the contract method 0x919898aa.
//
// Solidity: function setRoyaltiesWallet(address newRoyaltiesAddress) returns()
func (_Store *StoreTransactor) SetRoyaltiesWallet(opts *bind.TransactOpts, newRoyaltiesAddress common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setRoyaltiesWallet", newRoyaltiesAddress)
}

// SetRoyaltiesWallet is a paid mutator transaction binding the contract method 0x919898aa.
//
// Solidity: function setRoyaltiesWallet(address newRoyaltiesAddress) returns()
func (_Store *StoreSession) SetRoyaltiesWallet(newRoyaltiesAddress common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetRoyaltiesWallet(&_Store.TransactOpts, newRoyaltiesAddress)
}

// SetRoyaltiesWallet is a paid mutator transaction binding the contract method 0x919898aa.
//
// Solidity: function setRoyaltiesWallet(address newRoyaltiesAddress) returns()
func (_Store *StoreTransactorSession) SetRoyaltiesWallet(newRoyaltiesAddress common.Address) (*types.Transaction, error) {
	return _Store.Contract.SetRoyaltiesWallet(&_Store.TransactOpts, newRoyaltiesAddress)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 storeId, address recipient) returns()
func (_Store *StoreTransactor) Transfer(opts *bind.TransactOpts, storeId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transfer", storeId, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 storeId, address recipient) returns()
func (_Store *StoreSession) Transfer(storeId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, storeId, recipient)
}

// Transfer is a paid mutator transaction binding the contract method 0xb7760c8f.
//
// Solidity: function transfer(uint256 storeId, address recipient) returns()
func (_Store *StoreTransactorSession) Transfer(storeId *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, storeId, recipient)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreSession) Unpause() (*types.Transaction, error) {
	return _Store.Contract.Unpause(&_Store.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Store *StoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _Store.Contract.Unpause(&_Store.TransactOpts)
}

// StoreCollectionCreatedIterator is returned from FilterCollectionCreated and is used to iterate over the raw logs and unpacked data for CollectionCreated events raised by the Store contract.
type StoreCollectionCreatedIterator struct {
	Event *StoreCollectionCreated // Event containing the contract specifics and raw log

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
func (it *StoreCollectionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCollectionCreated)
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
		it.Event = new(StoreCollectionCreated)
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
func (it *StoreCollectionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCollectionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCollectionCreated represents a CollectionCreated event raised by the Store contract.
type StoreCollectionCreated struct {
	Name              string
	CollectionAddress common.Address
	Creator           common.Address
	ShowId            *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCollectionCreated is a free log retrieval operation binding the contract event 0x996485a610f9184a9f5a38da53d48de0fe91b5625ec0e38416b6e63085b11dcf.
//
// Solidity: event CollectionCreated(string name, address collectionAddress, address creator, uint256 showId)
func (_Store *StoreFilterer) FilterCollectionCreated(opts *bind.FilterOpts) (*StoreCollectionCreatedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return &StoreCollectionCreatedIterator{contract: _Store.contract, event: "CollectionCreated", logs: logs, sub: sub}, nil
}

// WatchCollectionCreated is a free log subscription operation binding the contract event 0x996485a610f9184a9f5a38da53d48de0fe91b5625ec0e38416b6e63085b11dcf.
//
// Solidity: event CollectionCreated(string name, address collectionAddress, address creator, uint256 showId)
func (_Store *StoreFilterer) WatchCollectionCreated(opts *bind.WatchOpts, sink chan<- *StoreCollectionCreated) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "CollectionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCollectionCreated)
				if err := _Store.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
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

// ParseCollectionCreated is a log parse operation binding the contract event 0x996485a610f9184a9f5a38da53d48de0fe91b5625ec0e38416b6e63085b11dcf.
//
// Solidity: event CollectionCreated(string name, address collectionAddress, address creator, uint256 showId)
func (_Store *StoreFilterer) ParseCollectionCreated(log types.Log) (*StoreCollectionCreated, error) {
	event := new(StoreCollectionCreated)
	if err := _Store.contract.UnpackLog(event, "CollectionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNFTMintedIterator is returned from FilterNFTMinted and is used to iterate over the raw logs and unpacked data for NFTMinted events raised by the Store contract.
type StoreNFTMintedIterator struct {
	Event *StoreNFTMinted // Event containing the contract specifics and raw log

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
func (it *StoreNFTMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNFTMinted)
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
		it.Event = new(StoreNFTMinted)
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
func (it *StoreNFTMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNFTMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNFTMinted represents a NFTMinted event raised by the Store contract.
type StoreNFTMinted struct {
	StoreId         *big.Int
	NftContract     common.Address
	NftTokenId      *big.Int
	Balance         *big.Int
	Creator         common.Address
	ListingCurrency common.Address
	ListingPrice    *big.Int
	ListingFees     *big.Int
	MinterData      string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNFTMinted is a free log retrieval operation binding the contract event 0x6010acc8b4704d513e9d9cb4f509581dea40d120414683cc4d28fe32ee59b4e9.
//
// Solidity: event NFTMinted(uint256 storeId, address nftContract, uint256 nftTokenId, uint256 balance, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, string minterData)
func (_Store *StoreFilterer) FilterNFTMinted(opts *bind.FilterOpts) (*StoreNFTMintedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return &StoreNFTMintedIterator{contract: _Store.contract, event: "NFTMinted", logs: logs, sub: sub}, nil
}

// WatchNFTMinted is a free log subscription operation binding the contract event 0x6010acc8b4704d513e9d9cb4f509581dea40d120414683cc4d28fe32ee59b4e9.
//
// Solidity: event NFTMinted(uint256 storeId, address nftContract, uint256 nftTokenId, uint256 balance, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, string minterData)
func (_Store *StoreFilterer) WatchNFTMinted(opts *bind.WatchOpts, sink chan<- *StoreNFTMinted) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NFTMinted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNFTMinted)
				if err := _Store.contract.UnpackLog(event, "NFTMinted", log); err != nil {
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

// ParseNFTMinted is a log parse operation binding the contract event 0x6010acc8b4704d513e9d9cb4f509581dea40d120414683cc4d28fe32ee59b4e9.
//
// Solidity: event NFTMinted(uint256 storeId, address nftContract, uint256 nftTokenId, uint256 balance, address creator, address listingCurrency, uint256 listingPrice, uint256 listingFees, string minterData)
func (_Store *StoreFilterer) ParseNFTMinted(log types.Log) (*StoreNFTMinted, error) {
	event := new(StoreNFTMinted)
	if err := _Store.contract.UnpackLog(event, "NFTMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreNFTPriceChangedIterator is returned from FilterNFTPriceChanged and is used to iterate over the raw logs and unpacked data for NFTPriceChanged events raised by the Store contract.
type StoreNFTPriceChangedIterator struct {
	Event *StoreNFTPriceChanged // Event containing the contract specifics and raw log

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
func (it *StoreNFTPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreNFTPriceChanged)
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
		it.Event = new(StoreNFTPriceChanged)
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
func (it *StoreNFTPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreNFTPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreNFTPriceChanged represents a NFTPriceChanged event raised by the Store contract.
type StoreNFTPriceChanged struct {
	StoreId         *big.Int
	NftTokenId      *big.Int
	NftContract     common.Address
	OldListingPrice *big.Int
	NewListingPrice *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNFTPriceChanged is a free log retrieval operation binding the contract event 0x0e40d608d06669526ed8f161ee08bb3652ae241694b4c4baa06371d8752d93f9.
//
// Solidity: event NFTPriceChanged(uint256 storeId, uint256 nftTokenId, address nftContract, uint256 oldListingPrice, uint256 newListingPrice)
func (_Store *StoreFilterer) FilterNFTPriceChanged(opts *bind.FilterOpts) (*StoreNFTPriceChangedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "NFTPriceChanged")
	if err != nil {
		return nil, err
	}
	return &StoreNFTPriceChangedIterator{contract: _Store.contract, event: "NFTPriceChanged", logs: logs, sub: sub}, nil
}

// WatchNFTPriceChanged is a free log subscription operation binding the contract event 0x0e40d608d06669526ed8f161ee08bb3652ae241694b4c4baa06371d8752d93f9.
//
// Solidity: event NFTPriceChanged(uint256 storeId, uint256 nftTokenId, address nftContract, uint256 oldListingPrice, uint256 newListingPrice)
func (_Store *StoreFilterer) WatchNFTPriceChanged(opts *bind.WatchOpts, sink chan<- *StoreNFTPriceChanged) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "NFTPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreNFTPriceChanged)
				if err := _Store.contract.UnpackLog(event, "NFTPriceChanged", log); err != nil {
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

// ParseNFTPriceChanged is a log parse operation binding the contract event 0x0e40d608d06669526ed8f161ee08bb3652ae241694b4c4baa06371d8752d93f9.
//
// Solidity: event NFTPriceChanged(uint256 storeId, uint256 nftTokenId, address nftContract, uint256 oldListingPrice, uint256 newListingPrice)
func (_Store *StoreFilterer) ParseNFTPriceChanged(log types.Log) (*StoreNFTPriceChanged, error) {
	event := new(StoreNFTPriceChanged)
	if err := _Store.contract.UnpackLog(event, "NFTPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StorePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Store contract.
type StorePausedIterator struct {
	Event *StorePaused // Event containing the contract specifics and raw log

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
func (it *StorePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StorePaused)
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
		it.Event = new(StorePaused)
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
func (it *StorePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StorePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StorePaused represents a Paused event raised by the Store contract.
type StorePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Store *StoreFilterer) FilterPaused(opts *bind.FilterOpts) (*StorePausedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StorePausedIterator{contract: _Store.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Store *StoreFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StorePaused) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StorePaused)
				if err := _Store.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Store *StoreFilterer) ParsePaused(log types.Log) (*StorePaused, error) {
	event := new(StorePaused)
	if err := _Store.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Store contract.
type StoreRoleAdminChangedIterator struct {
	Event *StoreRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StoreRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleAdminChanged)
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
		it.Event = new(StoreRoleAdminChanged)
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
func (it *StoreRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleAdminChanged represents a RoleAdminChanged event raised by the Store contract.
type StoreRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StoreRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleAdminChangedIterator{contract: _Store.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StoreRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleAdminChanged)
				if err := _Store.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Store *StoreFilterer) ParseRoleAdminChanged(log types.Log) (*StoreRoleAdminChanged, error) {
	event := new(StoreRoleAdminChanged)
	if err := _Store.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Store contract.
type StoreRoleGrantedIterator struct {
	Event *StoreRoleGranted // Event containing the contract specifics and raw log

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
func (it *StoreRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleGranted)
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
		it.Event = new(StoreRoleGranted)
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
func (it *StoreRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleGranted represents a RoleGranted event raised by the Store contract.
type StoreRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoreRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleGrantedIterator{contract: _Store.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StoreRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleGranted)
				if err := _Store.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) ParseRoleGranted(log types.Log) (*StoreRoleGranted, error) {
	event := new(StoreRoleGranted)
	if err := _Store.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Store contract.
type StoreRoleRevokedIterator struct {
	Event *StoreRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StoreRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreRoleRevoked)
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
		it.Event = new(StoreRoleRevoked)
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
func (it *StoreRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreRoleRevoked represents a RoleRevoked event raised by the Store contract.
type StoreRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StoreRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StoreRoleRevokedIterator{contract: _Store.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StoreRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreRoleRevoked)
				if err := _Store.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Store *StoreFilterer) ParseRoleRevoked(log types.Log) (*StoreRoleRevoked, error) {
	event := new(StoreRoleRevoked)
	if err := _Store.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Store contract.
type StoreUnpausedIterator struct {
	Event *StoreUnpaused // Event containing the contract specifics and raw log

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
func (it *StoreUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreUnpaused)
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
		it.Event = new(StoreUnpaused)
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
func (it *StoreUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreUnpaused represents a Unpaused event raised by the Store contract.
type StoreUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Store *StoreFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StoreUnpausedIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StoreUnpausedIterator{contract: _Store.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Store *StoreFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StoreUnpaused) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreUnpaused)
				if err := _Store.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Store *StoreFilterer) ParseUnpaused(log types.Log) (*StoreUnpaused, error) {
	event := new(StoreUnpaused)
	if err := _Store.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
