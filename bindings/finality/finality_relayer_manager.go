// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package finality

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

// BN254G1Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G1Point struct {
	X *big.Int
	Y *big.Int
}

// BN254G2Point is an auto generated low-level Go binding around an user-defined struct.
type BN254G2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// IBLSApkRegistryFinalityNonSignerAndSignature is an auto generated low-level Go binding around an user-defined struct.
type IBLSApkRegistryFinalityNonSignerAndSignature struct {
	NonSignerPubkeys []BN254G1Point
	ApkG2            BN254G2Point
	Sigma            BN254G1Point
	TotalBtcStake    *big.Int
	TotalMantaStake  *big.Int
}

// IFinalityRelayerManagerFinalityBatch is an auto generated low-level Go binding around an user-defined struct.
type IFinalityRelayerManagerFinalityBatch struct {
	StateRoot       [32]byte
	L2BlockNumber   *big.Int
	L1BlockHash     [32]byte
	L1BlockNumber   *big.Int
	MsgHash         [32]byte
	DisputeGameType uint32
}

// FinalityRelayerManagerMetaData contains all meta data concerning the FinalityRelayerManager contract.
var FinalityRelayerManagerMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"VerifyFinalitySignature\",\"inputs\":[{\"name\":\"finalityBatch\",\"type\":\"tuple\",\"internalType\":\"structIFinalityRelayerManager.FinalityBatch\",\"components\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"disputeGameType\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"finalityNonSignerAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.FinalityNonSignerAndSignature\",\"components\":[{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalBtcStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMantaStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOrRemoveOperatorWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAdd\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deRegisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isDisputeGameFactory\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_disputeGameFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorWhitelistManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDisputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelistManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"nodeUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nodeUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyFinalitySig\",\"inputs\":[{\"name\":\"totalBtcStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalMantaStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b610eb9806100d96000396000f3fe608060405234801561001057600080fd5b50600436106100ea5760003560e01c80638da5cb5b1161008c578063e03c863211610066578063e03c8632146101e4578063f2b4e61714610207578063f2fde38b1461021a578063fa5db51b1461022d57600080fd5b80638da5cb5b14610199578063a1e4c636146101c9578063b9a0634d146101dc57600080fd5b80634f5b4e15116100c85780634f5b4e15146101585780635df459461461016b578063715018a61461017e57806376af12a41461018657600080fd5b8063097c4af1146100ef5780632a630164146101045780634d9f15591461012d575b600080fd5b6101026100fd36600461088f565b610240565b005b60045461011890600160a01b900460ff1681565b60405190151581526020015b60405180910390f35b600354610140906001600160a01b031681565b6040516001600160a01b039091168152602001610124565b600054610140906001600160a01b031681565b600254610140906001600160a01b031681565b61010261031c565b61010261019436600461092d565b610330565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610140565b6101026101d7366004610af4565b6104a2565b610102610575565b6101186101f2366004610c2f565b60016020526000908152604090205460ff1681565b600454610140906001600160a01b031681565b610102610228366004610c2f565b610637565b61010261023b366004610c51565b610675565b3360009081526001602052604090205460ff166102785760405162461bcd60e51b815260040161026f90610c84565b60405180910390fd5b6002546040516303682a4560e41b81523360048201526001600160a01b0390911690633682a45090602401600060405180830381600087803b1580156102bd57600080fd5b505af11580156102d1573d6000803e3d6000fd5b50505050336001600160a01b03167f11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f58058383604051610310929190610d00565b60405180910390a25050565b6103246107c3565b61032e600061081e565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156103765750825b905060008267ffffffffffffffff1660011480156103935750303b155b9050811580156103a1575080155b156103bf5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156103e957845460ff60401b1916600160401b1785555b6103f28b61081e565b600280546001600160a01b03808c166001600160a01b031992831617909255600380548b8416908316179055600480548d1515600160a01b026001600160a81b03199091168b8516171790556000805492891692909116919091179055831561049557845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60025460405163041c37a160e21b815260009182916001600160a01b0390911690631070de84906104e39060808901359060208a0135908990600401610d7a565b606060405180830381865afa158015610500573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105249190610e3b565b81516020808401516040805193845291830152810182905291935091507f5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e19060600160405180910390a15050505050565b3360009081526001602052604090205460ff166105a45760405162461bcd60e51b815260040161026f90610c84565b600254604051636c67cc6560e11b81523360048201526001600160a01b039091169063d8cf98ca90602401600060405180830381600087803b1580156105e957600080fd5b505af11580156105fd573d6000803e3d6000fd5b50506040513381527fb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d739250602001905060405180910390a1565b61063f6107c3565b6001600160a01b03811661066957604051631e4fbdf760e01b81526000600482015260240161026f565b6106728161081e565b50565b6000546001600160a01b0316331461070e5760405162461bcd60e51b815260206004820152605060248201527f53747261746567794d616e616765722e6f6e6c7946696e616c6974795768697460448201527f654c6973744d616e616765723a206e6f74207468652066696e616c697479207760648201526f3434ba32b634b9ba1036b0b730b3b2b960811b608482015260a40161026f565b6001600160a01b0382166107985760405162461bcd60e51b815260206004820152604560248201527f46696e616c69747952656c617965724d616e616765722e6164644f706572617460448201527f6f7257686974656c6973743a206f70657261746f722061646472657373206973606482015264207a65726f60d81b608482015260a40161026f565b6001600160a01b03919091166000908152600160205260409020805460ff1916911515919091179055565b336107f57f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461032e5760405163118cdaa760e01b815233600482015260240161026f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b600080602083850312156108a257600080fd5b823567ffffffffffffffff808211156108ba57600080fd5b818501915085601f8301126108ce57600080fd5b8135818111156108dd57600080fd5b8660208285010111156108ef57600080fd5b60209290920196919550909350505050565b80356001600160a01b038116811461091857600080fd5b919050565b8035801515811461091857600080fd5b60008060008060008060c0878903121561094657600080fd5b61094f87610901565b955061095d6020880161091d565b945061096b60408801610901565b935061097960608801610901565b925061098760808801610901565b915061099560a08801610901565b90509295509295509295565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff811182821017156109da576109da6109a1565b60405290565b60405160a0810167ffffffffffffffff811182821017156109da576109da6109a1565b604051601f8201601f1916810167ffffffffffffffff81118282101715610a2c57610a2c6109a1565b604052919050565b600060408284031215610a4657600080fd5b610a4e6109b7565b9050813581526020820135602082015292915050565b600082601f830112610a7557600080fd5b610a7d6109b7565b806040840185811115610a8f57600080fd5b845b81811015610aa9578035845260209384019301610a91565b509095945050505050565b600060808284031215610ac657600080fd5b610ace6109b7565b9050610ada8383610a64565b8152610ae98360408401610a64565b602082015292915050565b600080600083850361010080821215610b0c57600080fd5b60c0821215610b1a57600080fd5b85945060c0860135915067ffffffffffffffff80831115610b3a57600080fd5b918601916101208389031215610b4f57600080fd5b610b576109e0565b833582811115610b6657600080fd5b8401601f81018a13610b7757600080fd5b8035602084821115610b8b57610b8b6109a1565b610b99818360051b01610a03565b828152818101955060069290921b83018101918c831115610bb957600080fd5b928101925b82841015610be257610bd08d85610a34565b86528186019550604084019350610bbe565b8452610bf08c888301610ab4565b81850152505050610c048960a08601610a34565b604082015260e084810135606083015292909301356080840152509396909550939092013592915050565b600060208284031215610c4157600080fd5b610c4a82610901565b9392505050565b60008060408385031215610c6457600080fd5b610c6d83610901565b9150610c7b6020840161091d565b90509250929050565b60208082526056908201527f46696e616c69747952656c617965724d616e616765722e72656769737465724f60408201527f70657261746f723a207468697320616464726573732068617665206e6f742070606082015275032b936b4b9b9b4b7b7103a37903932b3b4b9ba32b9160551b608082015260a00190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b8060005b6002811015610d52578151845260209384019390910190600101610d33565b50505050565b610d63828251610d2f565b6020810151610d756040840182610d2f565b505050565b838152600060208460208401526040606060408501526101808401855161012060608701528181518084526101a088019150602083019350600092505b80831015610de857610dd482855180518252602090810151910152565b928501926001929092019190840190610db7565b5060208801519450610dfd6080880186610d58565b604088015180516101008901526020015161012088015260608801516101408801526080909701516101609096019590955250939695505050505050565b6000808284036060811215610e4f57600080fd5b6040811215610e5d57600080fd5b50610e666109b7565b83518152602080850151908201526040909301519294929350505056fea2646970667358221220580c42fb250dc329437dde60dc3e6f3f8fa974a6e25f789ffd148238947ee50764736f6c63430008190033",
}

// FinalityRelayerManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FinalityRelayerManagerMetaData.ABI instead.
var FinalityRelayerManagerABI = FinalityRelayerManagerMetaData.ABI

// FinalityRelayerManagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FinalityRelayerManagerMetaData.Bin instead.
var FinalityRelayerManagerBin = FinalityRelayerManagerMetaData.Bin

// DeployFinalityRelayerManager deploys a new Ethereum contract, binding an instance of FinalityRelayerManager to it.
func DeployFinalityRelayerManager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FinalityRelayerManager, error) {
	parsed, err := FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FinalityRelayerManagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FinalityRelayerManager{FinalityRelayerManagerCaller: FinalityRelayerManagerCaller{contract: contract}, FinalityRelayerManagerTransactor: FinalityRelayerManagerTransactor{contract: contract}, FinalityRelayerManagerFilterer: FinalityRelayerManagerFilterer{contract: contract}}, nil
}

// FinalityRelayerManager is an auto generated Go binding around an Ethereum contract.
type FinalityRelayerManager struct {
	FinalityRelayerManagerCaller     // Read-only binding to the contract
	FinalityRelayerManagerTransactor // Write-only binding to the contract
	FinalityRelayerManagerFilterer   // Log filterer for contract events
}

// FinalityRelayerManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FinalityRelayerManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FinalityRelayerManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FinalityRelayerManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FinalityRelayerManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FinalityRelayerManagerSession struct {
	Contract     *FinalityRelayerManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FinalityRelayerManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FinalityRelayerManagerCallerSession struct {
	Contract *FinalityRelayerManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// FinalityRelayerManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FinalityRelayerManagerTransactorSession struct {
	Contract     *FinalityRelayerManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// FinalityRelayerManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FinalityRelayerManagerRaw struct {
	Contract *FinalityRelayerManager // Generic contract binding to access the raw methods on
}

// FinalityRelayerManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FinalityRelayerManagerCallerRaw struct {
	Contract *FinalityRelayerManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FinalityRelayerManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FinalityRelayerManagerTransactorRaw struct {
	Contract *FinalityRelayerManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFinalityRelayerManager creates a new instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManager(address common.Address, backend bind.ContractBackend) (*FinalityRelayerManager, error) {
	contract, err := bindFinalityRelayerManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManager{FinalityRelayerManagerCaller: FinalityRelayerManagerCaller{contract: contract}, FinalityRelayerManagerTransactor: FinalityRelayerManagerTransactor{contract: contract}, FinalityRelayerManagerFilterer: FinalityRelayerManagerFilterer{contract: contract}}, nil
}

// NewFinalityRelayerManagerCaller creates a new read-only instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerCaller(address common.Address, caller bind.ContractCaller) (*FinalityRelayerManagerCaller, error) {
	contract, err := bindFinalityRelayerManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerCaller{contract: contract}, nil
}

// NewFinalityRelayerManagerTransactor creates a new write-only instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FinalityRelayerManagerTransactor, error) {
	contract, err := bindFinalityRelayerManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerTransactor{contract: contract}, nil
}

// NewFinalityRelayerManagerFilterer creates a new log filterer instance of FinalityRelayerManager, bound to a specific deployed contract.
func NewFinalityRelayerManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FinalityRelayerManagerFilterer, error) {
	contract, err := bindFinalityRelayerManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerFilterer{contract: contract}, nil
}

// bindFinalityRelayerManager binds a generic wrapper to an already deployed contract.
func bindFinalityRelayerManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FinalityRelayerManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityRelayerManager *FinalityRelayerManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.FinalityRelayerManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FinalityRelayerManager *FinalityRelayerManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FinalityRelayerManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.contract.Transact(opts, method, params...)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) BlsApkRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "blsApkRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) BlsApkRegistry() (common.Address, error) {
	return _FinalityRelayerManager.Contract.BlsApkRegistry(&_FinalityRelayerManager.CallOpts)
}

// BlsApkRegistry is a free data retrieval call binding the contract method 0x5df45946.
//
// Solidity: function blsApkRegistry() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) BlsApkRegistry() (common.Address, error) {
	return _FinalityRelayerManager.Contract.BlsApkRegistry(&_FinalityRelayerManager.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) DisputeGameFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "disputeGameFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) DisputeGameFactory() (common.Address, error) {
	return _FinalityRelayerManager.Contract.DisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) DisputeGameFactory() (common.Address, error) {
	return _FinalityRelayerManager.Contract.DisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) IsDisputeGameFactory(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "isDisputeGameFactory")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) IsDisputeGameFactory() (bool, error) {
	return _FinalityRelayerManager.Contract.IsDisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// IsDisputeGameFactory is a free data retrieval call binding the contract method 0x2a630164.
//
// Solidity: function isDisputeGameFactory() view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) IsDisputeGameFactory() (bool, error) {
	return _FinalityRelayerManager.Contract.IsDisputeGameFactory(&_FinalityRelayerManager.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) L2OutputOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "l2OutputOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) L2OutputOracle() (common.Address, error) {
	return _FinalityRelayerManager.Contract.L2OutputOracle(&_FinalityRelayerManager.CallOpts)
}

// L2OutputOracle is a free data retrieval call binding the contract method 0x4d9f1559.
//
// Solidity: function l2OutputOracle() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) L2OutputOracle() (common.Address, error) {
	return _FinalityRelayerManager.Contract.L2OutputOracle(&_FinalityRelayerManager.CallOpts)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) OperatorWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "operatorWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelist(&_FinalityRelayerManager.CallOpts, arg0)
}

// OperatorWhitelist is a free data retrieval call binding the contract method 0xe03c8632.
//
// Solidity: function operatorWhitelist(address ) view returns(bool)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) OperatorWhitelist(arg0 common.Address) (bool, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelist(&_FinalityRelayerManager.CallOpts, arg0)
}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) OperatorWhitelistManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "operatorWhitelistManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) OperatorWhitelistManager() (common.Address, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelistManager(&_FinalityRelayerManager.CallOpts)
}

// OperatorWhitelistManager is a free data retrieval call binding the contract method 0x4f5b4e15.
//
// Solidity: function operatorWhitelistManager() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) OperatorWhitelistManager() (common.Address, error) {
	return _FinalityRelayerManager.Contract.OperatorWhitelistManager(&_FinalityRelayerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) Owner() (common.Address, error) {
	return _FinalityRelayerManager.Contract.Owner(&_FinalityRelayerManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) Owner() (common.Address, error) {
	return _FinalityRelayerManager.Contract.Owner(&_FinalityRelayerManager.CallOpts)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSignerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) VerifyFinalitySignature(opts *bind.TransactOpts, finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSignerAndSignature IBLSApkRegistryFinalityNonSignerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "VerifyFinalitySignature", finalityBatch, finalityNonSignerAndSignature, minGas)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSignerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) VerifyFinalitySignature(finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSignerAndSignature IBLSApkRegistryFinalityNonSignerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.VerifyFinalitySignature(&_FinalityRelayerManager.TransactOpts, finalityBatch, finalityNonSignerAndSignature, minGas)
}

// VerifyFinalitySignature is a paid mutator transaction binding the contract method 0xa1e4c636.
//
// Solidity: function VerifyFinalitySignature((bytes32,uint256,bytes32,uint256,bytes32,uint32) finalityBatch, ((uint256,uint256)[],(uint256[2],uint256[2]),(uint256,uint256),uint256,uint256) finalityNonSignerAndSignature, uint256 minGas) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) VerifyFinalitySignature(finalityBatch IFinalityRelayerManagerFinalityBatch, finalityNonSignerAndSignature IBLSApkRegistryFinalityNonSignerAndSignature, minGas *big.Int) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.VerifyFinalitySignature(&_FinalityRelayerManager.TransactOpts, finalityBatch, finalityNonSignerAndSignature, minGas)
}

// AddOrRemoveOperatorWhitelist is a paid mutator transaction binding the contract method 0xfa5db51b.
//
// Solidity: function addOrRemoveOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) AddOrRemoveOperatorWhitelist(opts *bind.TransactOpts, operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "addOrRemoveOperatorWhitelist", operator, isAdd)
}

// AddOrRemoveOperatorWhitelist is a paid mutator transaction binding the contract method 0xfa5db51b.
//
// Solidity: function addOrRemoveOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) AddOrRemoveOperatorWhitelist(operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.AddOrRemoveOperatorWhitelist(&_FinalityRelayerManager.TransactOpts, operator, isAdd)
}

// AddOrRemoveOperatorWhitelist is a paid mutator transaction binding the contract method 0xfa5db51b.
//
// Solidity: function addOrRemoveOperatorWhitelist(address operator, bool isAdd) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) AddOrRemoveOperatorWhitelist(operator common.Address, isAdd bool) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.AddOrRemoveOperatorWhitelist(&_FinalityRelayerManager.TransactOpts, operator, isAdd)
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) DeRegisterOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "deRegisterOperator")
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) DeRegisterOperator() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.DeRegisterOperator(&_FinalityRelayerManager.TransactOpts)
}

// DeRegisterOperator is a paid mutator transaction binding the contract method 0xb9a0634d.
//
// Solidity: function deRegisterOperator() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) DeRegisterOperator() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.DeRegisterOperator(&_FinalityRelayerManager.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x76af12a4.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory, address _operatorWhitelistManager) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) Initialize(opts *bind.TransactOpts, _initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address, _operatorWhitelistManager common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "initialize", _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory, _operatorWhitelistManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x76af12a4.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory, address _operatorWhitelistManager) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) Initialize(_initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address, _operatorWhitelistManager common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.Initialize(&_FinalityRelayerManager.TransactOpts, _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory, _operatorWhitelistManager)
}

// Initialize is a paid mutator transaction binding the contract method 0x76af12a4.
//
// Solidity: function initialize(address _initialOwner, bool _isDisputeGameFactory, address _blsApkRegistry, address _l2OutputOracle, address _disputeGameFactory, address _operatorWhitelistManager) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) Initialize(_initialOwner common.Address, _isDisputeGameFactory bool, _blsApkRegistry common.Address, _l2OutputOracle common.Address, _disputeGameFactory common.Address, _operatorWhitelistManager common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.Initialize(&_FinalityRelayerManager.TransactOpts, _initialOwner, _isDisputeGameFactory, _blsApkRegistry, _l2OutputOracle, _disputeGameFactory, _operatorWhitelistManager)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) RegisterOperator(opts *bind.TransactOpts, nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "registerOperator", nodeUrl)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) RegisterOperator(nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RegisterOperator(&_FinalityRelayerManager.TransactOpts, nodeUrl)
}

// RegisterOperator is a paid mutator transaction binding the contract method 0x097c4af1.
//
// Solidity: function registerOperator(string nodeUrl) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) RegisterOperator(nodeUrl string) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RegisterOperator(&_FinalityRelayerManager.TransactOpts, nodeUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RenounceOwnership(&_FinalityRelayerManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.RenounceOwnership(&_FinalityRelayerManager.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.TransferOwnership(&_FinalityRelayerManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FinalityRelayerManager *FinalityRelayerManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FinalityRelayerManager.Contract.TransferOwnership(&_FinalityRelayerManager.TransactOpts, newOwner)
}

// FinalityRelayerManagerInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerInitializedIterator struct {
	Event *FinalityRelayerManagerInitialized // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerInitialized)
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
		it.Event = new(FinalityRelayerManagerInitialized)
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
func (it *FinalityRelayerManagerInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerInitialized represents a Initialized event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterInitialized(opts *bind.FilterOpts) (*FinalityRelayerManagerInitializedIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerInitializedIterator{contract: _FinalityRelayerManager.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerInitialized) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerInitialized)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseInitialized(log types.Log) (*FinalityRelayerManagerInitialized, error) {
	event := new(FinalityRelayerManagerInitialized)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOperatorDeRegisteredIterator is returned from FilterOperatorDeRegistered and is used to iterate over the raw logs and unpacked data for OperatorDeRegistered events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorDeRegisteredIterator struct {
	Event *FinalityRelayerManagerOperatorDeRegistered // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOperatorDeRegistered)
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
		it.Event = new(FinalityRelayerManagerOperatorDeRegistered)
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
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOperatorDeRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOperatorDeRegistered represents a OperatorDeRegistered event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorDeRegistered struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorDeRegistered is a free log retrieval operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOperatorDeRegistered(opts *bind.FilterOpts) (*FinalityRelayerManagerOperatorDeRegisteredIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOperatorDeRegisteredIterator{contract: _FinalityRelayerManager.contract, event: "OperatorDeRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorDeRegistered is a free log subscription operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOperatorDeRegistered(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOperatorDeRegistered) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OperatorDeRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOperatorDeRegistered)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
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

// ParseOperatorDeRegistered is a log parse operation binding the contract event 0xb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d73.
//
// Solidity: event OperatorDeRegistered(address operator)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOperatorDeRegistered(log types.Log) (*FinalityRelayerManagerOperatorDeRegistered, error) {
	event := new(FinalityRelayerManagerOperatorDeRegistered)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorDeRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOperatorRegisteredIterator is returned from FilterOperatorRegistered and is used to iterate over the raw logs and unpacked data for OperatorRegistered events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorRegisteredIterator struct {
	Event *FinalityRelayerManagerOperatorRegistered // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOperatorRegistered)
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
		it.Event = new(FinalityRelayerManagerOperatorRegistered)
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
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOperatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOperatorRegistered represents a OperatorRegistered event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOperatorRegistered struct {
	Operator common.Address
	NodeUrl  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorRegistered is a free log retrieval operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOperatorRegistered(opts *bind.FilterOpts, operator []common.Address) (*FinalityRelayerManagerOperatorRegisteredIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOperatorRegisteredIterator{contract: _FinalityRelayerManager.contract, event: "OperatorRegistered", logs: logs, sub: sub}, nil
}

// WatchOperatorRegistered is a free log subscription operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOperatorRegistered(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOperatorRegistered, operator []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OperatorRegistered", operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOperatorRegistered)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
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

// ParseOperatorRegistered is a log parse operation binding the contract event 0x11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805.
//
// Solidity: event OperatorRegistered(address indexed operator, string nodeUrl)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOperatorRegistered(log types.Log) (*FinalityRelayerManagerOperatorRegistered, error) {
	event := new(FinalityRelayerManagerOperatorRegistered)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OperatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOwnershipTransferredIterator struct {
	Event *FinalityRelayerManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerOwnershipTransferred)
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
		it.Event = new(FinalityRelayerManagerOwnershipTransferred)
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
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerOwnershipTransferred represents a OwnershipTransferred event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FinalityRelayerManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerOwnershipTransferredIterator{contract: _FinalityRelayerManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerOwnershipTransferred)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseOwnershipTransferred(log types.Log) (*FinalityRelayerManagerOwnershipTransferred, error) {
	event := new(FinalityRelayerManagerOwnershipTransferred)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FinalityRelayerManagerVerifyFinalitySigIterator is returned from FilterVerifyFinalitySig and is used to iterate over the raw logs and unpacked data for VerifyFinalitySig events raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerVerifyFinalitySigIterator struct {
	Event *FinalityRelayerManagerVerifyFinalitySig // Event containing the contract specifics and raw log

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
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FinalityRelayerManagerVerifyFinalitySig)
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
		it.Event = new(FinalityRelayerManagerVerifyFinalitySig)
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
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FinalityRelayerManagerVerifyFinalitySigIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FinalityRelayerManagerVerifyFinalitySig represents a VerifyFinalitySig event raised by the FinalityRelayerManager contract.
type FinalityRelayerManagerVerifyFinalitySig struct {
	TotalBtcStaking     *big.Int
	TotalMantaStaking   *big.Int
	SignatoryRecordHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVerifyFinalitySig is a free log retrieval operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterVerifyFinalitySig(opts *bind.FilterOpts) (*FinalityRelayerManagerVerifyFinalitySigIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "VerifyFinalitySig")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerVerifyFinalitySigIterator{contract: _FinalityRelayerManager.contract, event: "VerifyFinalitySig", logs: logs, sub: sub}, nil
}

// WatchVerifyFinalitySig is a free log subscription operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) WatchVerifyFinalitySig(opts *bind.WatchOpts, sink chan<- *FinalityRelayerManagerVerifyFinalitySig) (event.Subscription, error) {

	logs, sub, err := _FinalityRelayerManager.contract.WatchLogs(opts, "VerifyFinalitySig")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FinalityRelayerManagerVerifyFinalitySig)
				if err := _FinalityRelayerManager.contract.UnpackLog(event, "VerifyFinalitySig", log); err != nil {
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

// ParseVerifyFinalitySig is a log parse operation binding the contract event 0x5867a1f09ebc8c9fa2b0ab07694a570b9bb77b2603f5939e40b08b76e49b94e1.
//
// Solidity: event VerifyFinalitySig(uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseVerifyFinalitySig(log types.Log) (*FinalityRelayerManagerVerifyFinalitySig, error) {
	event := new(FinalityRelayerManagerVerifyFinalitySig)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "VerifyFinalitySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
