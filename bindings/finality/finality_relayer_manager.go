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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"VerifyFinalitySignature\",\"inputs\":[{\"name\":\"finalityBatch\",\"type\":\"tuple\",\"internalType\":\"structIFinalityRelayerManager.FinalityBatch\",\"components\":[{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l2BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"l1BlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"l1BlockNumber\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"msgHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"disputeGameType\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]},{\"name\":\"finalityNonSignerAndSignature\",\"type\":\"tuple\",\"internalType\":\"structIBLSApkRegistry.FinalityNonSignerAndSignature\",\"components\":[{\"name\":\"nonSignerPubkeys\",\"type\":\"tuple[]\",\"internalType\":\"structBN254.G1Point[]\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"apkG2\",\"type\":\"tuple\",\"internalType\":\"structBN254.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"sigma\",\"type\":\"tuple\",\"internalType\":\"structBN254.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"totalBtcStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"totalMantaStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"minGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addOrRemoveOperatorWhitelist\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAdd\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"blsApkRegistry\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIBLSApkRegistry\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"confirmBatchId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deRegisterOperator\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_initialOwner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_isDisputeGameFactory\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_blsApkRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_disputeGameFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_operatorWhitelistManager\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDisputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2OutputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelist\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"operatorWhitelistManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"registerOperator\",\"inputs\":[{\"name\":\"nodeUrl\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorDeRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OperatorRegistered\",\"inputs\":[{\"name\":\"operator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nodeUrl\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VerifyFinalitySig\",\"inputs\":[{\"name\":\"batchId\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalBtcStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"totalMantaStaking\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"signatoryRecordHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61125f806100df6000396000f3fe608060405234801561001057600080fd5b50600436106100f55760003560e01c80638c76d2bd11610097578063e03c863211610066578063e03c863214610206578063f2b4e61714610229578063f2fde38b1461023c578063fa5db51b1461024f57600080fd5b80638c76d2bd146101a45780638da5cb5b146101bb578063a1e4c636146101eb578063b9a0634d146101fe57600080fd5b80634f5b4e15116100d35780634f5b4e15146101635780635df4594614610176578063715018a61461018957806376af12a41461019157600080fd5b8063097c4af1146100fa5780632a6301641461010f5780634d9f155914610138575b600080fd5b61010d610108366004610bba565b610262565b005b60045461012390600160a01b900460ff1681565b60405190151581526020015b60405180910390f35b60035461014b906001600160a01b031681565b6040516001600160a01b03909116815260200161012f565b60005461014b906001600160a01b031681565b60025461014b906001600160a01b031681565b61010d61033e565b61010d61019f366004610c58565b610352565b6101ad60055481565b60405190815260200161012f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661014b565b61010d6101f9366004610e1f565b6104c7565b61010d6107b8565b610123610214366004610f5a565b60016020526000908152604090205460ff1681565b60045461014b906001600160a01b031681565b61010d61024a366004610f5a565b61087a565b61010d61025d366004610f7c565b6108b8565b3360009081526001602052604090205460ff1661029a5760405162461bcd60e51b815260040161029190610faf565b60405180910390fd5b6002546040516303682a4560e41b81523360048201526001600160a01b0390911690633682a45090602401600060405180830381600087803b1580156102df57600080fd5b505af11580156102f3573d6000803e3d6000fd5b50505050336001600160a01b03167f11a85ea4a40584362c3d9c17685709a2e02b466ac78d5eb00b6aff73d90f5805838360405161033292919061102b565b60405180910390a25050565b610346610a06565b6103506000610a61565b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156103985750825b905060008267ffffffffffffffff1660011480156103b55750303b155b9050811580156103c3575080155b156103e15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561040b57845460ff60401b1916600160401b1785555b6104148b610a61565b600280546001600160a01b03808c166001600160a01b031992831617909255600380548b8416908316179055600480548d1515600160a01b026001600160a81b03199091168b851617179055600080549289169290911691909117815560055583156104ba57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b60025460405163041c37a160e21b815260009182916001600160a01b0390911690631070de84906105089060808901359060208a01359089906004016110a5565b606060405180830381865afa158015610525573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105499190611166565b91509150600061056185608001518660600151610ad2565b600454909150600160a01b900460ff166106745760035460405187356024820152604481018390526000916105d3916001600160a01b0390911690879084906064015b60408051601f198184030181529190526020810180516001600160e01b031663bd05743b60e01b179052610b3e565b90508061066e5760405162461bcd60e51b815260206004820152605e60248201527f5374726174656779426173652e56657269667946696e616c6974795369676e6160448201527f747572653a206368616e67652066696e616c697a656420706572696f6473206960648201527f6e206c326f7574707574206f7261636c65207365636f6e6473206661696c0000608482015260a401610291565b50610750565b60045460405187356024820152604481018390526000916106a7916001600160a01b0390911690879084906064016105a4565b90508061074e5760405162461bcd60e51b815260206004820152606360248201527f5374726174656779426173652e56657269667946696e616c6974795369676e6160448201527f747572653a206368616e67652066696e616c697a656420706572696f6473206960648201527f6e20646973707574652067616d6520666163746f7279207365636f6e64732066608482015262185a5b60ea1b60a482015260c401610291565b505b600580547f9262d1d0a737d718e3befc723ca2b2e75b805f632947f99a2fb0af0e61fdd194916000610781836111c4565b90915550845160208087015160408051948552918401929092528201526060810184905260800160405180910390a1505050505050565b3360009081526001602052604090205460ff166107e75760405162461bcd60e51b815260040161029190610faf565b600254604051636c67cc6560e11b81523360048201526001600160a01b039091169063d8cf98ca90602401600060405180830381600087803b15801561082c57600080fd5b505af1158015610840573d6000803e3d6000fd5b50506040513381527fb2c38c6252ee2d17f80059fb47a790e20f7bd75e7ba577685375e5484f412d739250602001905060405180910390a1565b610882610a06565b6001600160a01b0381166108ac57604051631e4fbdf760e01b815260006004820152602401610291565b6108b581610a61565b50565b6000546001600160a01b031633146109515760405162461bcd60e51b815260206004820152605060248201527f53747261746567794d616e616765722e6f6e6c7946696e616c6974795768697460448201527f654c6973744d616e616765723a206e6f74207468652066696e616c697479207760648201526f3434ba32b634b9ba1036b0b730b3b2b960811b608482015260a401610291565b6001600160a01b0382166109db5760405162461bcd60e51b815260206004820152604560248201527f46696e616c69747952656c617965724d616e616765722e6164644f706572617460448201527f6f7257686974656c6973743a206f70657261746f722061646472657373206973606482015264207a65726f60d81b608482015260a401610291565b6001600160a01b03919091166000908152600160205260409020805460ff1916911515919091179055565b33610a387f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146103505760405163118cdaa760e01b8152336004820152602401610291565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b60008062015180610ae864174876e800856111dd565b610af291906111ff565b62015180610b0a69d3c21bcecceda1000000876111dd565b610b1491906111ff565b610b1e9190611216565b905061a8c0811015610b355761a8c0915050610b38565b90505b92915050565b6000806000610b4e866000610b9c565b905080610b84576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600080603f83619c4001026040850201603f5a021015949350505050565b60008060208385031215610bcd57600080fd5b823567ffffffffffffffff80821115610be557600080fd5b818501915085601f830112610bf957600080fd5b813581811115610c0857600080fd5b866020828501011115610c1a57600080fd5b60209290920196919550909350505050565b80356001600160a01b0381168114610c4357600080fd5b919050565b80358015158114610c4357600080fd5b60008060008060008060c08789031215610c7157600080fd5b610c7a87610c2c565b9550610c8860208801610c48565b9450610c9660408801610c2c565b9350610ca460608801610c2c565b9250610cb260808801610c2c565b9150610cc060a08801610c2c565b90509295509295509295565b634e487b7160e01b600052604160045260246000fd5b6040805190810167ffffffffffffffff81118282101715610d0557610d05610ccc565b60405290565b60405160a0810167ffffffffffffffff81118282101715610d0557610d05610ccc565b604051601f8201601f1916810167ffffffffffffffff81118282101715610d5757610d57610ccc565b604052919050565b600060408284031215610d7157600080fd5b610d79610ce2565b9050813581526020820135602082015292915050565b600082601f830112610da057600080fd5b610da8610ce2565b806040840185811115610dba57600080fd5b845b81811015610dd4578035845260209384019301610dbc565b509095945050505050565b600060808284031215610df157600080fd5b610df9610ce2565b9050610e058383610d8f565b8152610e148360408401610d8f565b602082015292915050565b600080600083850361010080821215610e3757600080fd5b60c0821215610e4557600080fd5b85945060c0860135915067ffffffffffffffff80831115610e6557600080fd5b918601916101208389031215610e7a57600080fd5b610e82610d0b565b833582811115610e9157600080fd5b8401601f81018a13610ea257600080fd5b8035602084821115610eb657610eb6610ccc565b610ec4818360051b01610d2e565b828152818101955060069290921b83018101918c831115610ee457600080fd5b928101925b82841015610f0d57610efb8d85610d5f565b86528186019550604084019350610ee9565b8452610f1b8c888301610ddf565b81850152505050610f2f8960a08601610d5f565b604082015260e084810135606083015292909301356080840152509396909550939092013592915050565b600060208284031215610f6c57600080fd5b610f7582610c2c565b9392505050565b60008060408385031215610f8f57600080fd5b610f9883610c2c565b9150610fa660208401610c48565b90509250929050565b60208082526056908201527f46696e616c69747952656c617965724d616e616765722e72656769737465724f60408201527f70657261746f723a207468697320616464726573732068617665206e6f742070606082015275032b936b4b9b9b4b7b7103a37903932b3b4b9ba32b9160551b608082015260a00190565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b8060005b600281101561107d57815184526020938401939091019060010161105e565b50505050565b61108e82825161105a565b60208101516110a0604084018261105a565b505050565b838152600060208460208401526040606060408501526101808401855161012060608701528181518084526101a088019150602083019350600092505b80831015611113576110ff82855180518252602090810151910152565b9285019260019290920191908401906110e2565b50602088015194506111286080880186611083565b604088015180516101008901526020015161012088015260608801516101408801526080909701516101609096019590955250939695505050505050565b600080828403606081121561117a57600080fd5b604081121561118857600080fd5b50611191610ce2565b835181526020808501519082015260409093015192949293505050565b634e487b7160e01b600052601160045260246000fd5b6000600182016111d6576111d66111ae565b5060010190565b6000826111fa57634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417610b3857610b386111ae565b80820180821115610b3857610b386111ae56fea2646970667358221220e9f394c860d6c2d1d1476b18e76cc49640bd6308a6af9a0b360266d26412568b64736f6c63430008160033",
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

// ConfirmBatchId is a free data retrieval call binding the contract method 0x8c76d2bd.
//
// Solidity: function confirmBatchId() view returns(uint256)
func (_FinalityRelayerManager *FinalityRelayerManagerCaller) ConfirmBatchId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FinalityRelayerManager.contract.Call(opts, &out, "confirmBatchId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ConfirmBatchId is a free data retrieval call binding the contract method 0x8c76d2bd.
//
// Solidity: function confirmBatchId() view returns(uint256)
func (_FinalityRelayerManager *FinalityRelayerManagerSession) ConfirmBatchId() (*big.Int, error) {
	return _FinalityRelayerManager.Contract.ConfirmBatchId(&_FinalityRelayerManager.CallOpts)
}

// ConfirmBatchId is a free data retrieval call binding the contract method 0x8c76d2bd.
//
// Solidity: function confirmBatchId() view returns(uint256)
func (_FinalityRelayerManager *FinalityRelayerManagerCallerSession) ConfirmBatchId() (*big.Int, error) {
	return _FinalityRelayerManager.Contract.ConfirmBatchId(&_FinalityRelayerManager.CallOpts)
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
	BatchId             *big.Int
	TotalBtcStaking     *big.Int
	TotalMantaStaking   *big.Int
	SignatoryRecordHash [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterVerifyFinalitySig is a free log retrieval operation binding the contract event 0x9262d1d0a737d718e3befc723ca2b2e75b805f632947f99a2fb0af0e61fdd194.
//
// Solidity: event VerifyFinalitySig(uint256 batchId, uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) FilterVerifyFinalitySig(opts *bind.FilterOpts) (*FinalityRelayerManagerVerifyFinalitySigIterator, error) {

	logs, sub, err := _FinalityRelayerManager.contract.FilterLogs(opts, "VerifyFinalitySig")
	if err != nil {
		return nil, err
	}
	return &FinalityRelayerManagerVerifyFinalitySigIterator{contract: _FinalityRelayerManager.contract, event: "VerifyFinalitySig", logs: logs, sub: sub}, nil
}

// WatchVerifyFinalitySig is a free log subscription operation binding the contract event 0x9262d1d0a737d718e3befc723ca2b2e75b805f632947f99a2fb0af0e61fdd194.
//
// Solidity: event VerifyFinalitySig(uint256 batchId, uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
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

// ParseVerifyFinalitySig is a log parse operation binding the contract event 0x9262d1d0a737d718e3befc723ca2b2e75b805f632947f99a2fb0af0e61fdd194.
//
// Solidity: event VerifyFinalitySig(uint256 batchId, uint256 totalBtcStaking, uint256 totalMantaStaking, bytes32 signatoryRecordHash)
func (_FinalityRelayerManager *FinalityRelayerManagerFilterer) ParseVerifyFinalitySig(log types.Log) (*FinalityRelayerManagerVerifyFinalitySig, error) {
	event := new(FinalityRelayerManagerVerifyFinalitySig)
	if err := _FinalityRelayerManager.contract.UnpackLog(event, "VerifyFinalitySig", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
