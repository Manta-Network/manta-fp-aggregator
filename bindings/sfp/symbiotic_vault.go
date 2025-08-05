// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package sfp

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

// SymbioticVaultMetaData contains all meta data concerning the SymbioticVault contract.
var SymbioticVaultMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"delegatorFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"slasherFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"vaultFactory\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DELEGATOR_FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSITOR_WHITELIST_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_LIMIT_SET_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEPOSIT_WHITELIST_SET_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"IS_DEPOSIT_LIMIT_SET_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SLASHER_FACTORY\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeBalanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeBalanceOfAt\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"hints\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeShares\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeSharesAt\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"hint\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeSharesOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeSharesOfAt\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"hint\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"activeStakeAt\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"},{\"name\":\"hint\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimBatch\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"epochs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"collateral\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentEpoch\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentEpochStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"delegator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposit\",\"inputs\":[{\"name\":\"onBehalfOf\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"depositedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mintedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositWhitelist\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochAt\",\"inputs\":[{\"name\":\"timestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochDuration\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"epochDurationInit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"initialVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isDelegatorInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDepositLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isDepositorWhitelisted\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSlasherInitialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isWithdrawalsClaimed\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"value\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"migrate\",\"inputs\":[{\"name\":\"newVersion\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"nextEpochStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onSlash\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"outputs\":[{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"previousEpochStart\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint48\",\"internalType\":\"uint48\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"redeem\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"shares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"withdrawnAssets\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mintedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDelegator\",\"inputs\":[{\"name\":\"delegator_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDepositLimit\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDepositWhitelist\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDepositorWhitelistStatus\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setIsDepositLimit\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSlasher\",\"inputs\":[{\"name\":\"slasher_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"slashableBalanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"slasher\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staticDelegateCall\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"burnedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"mintedShares\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawalShares\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalSharesOf\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawals\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawalsOf\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"epoch\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ClaimBatch\",\"inputs\":[{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"epochs\",\"type\":\"uint256[]\",\"indexed\":false,\"internalType\":\"uint256[]\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"depositor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"onBehalfOf\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"shares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OnSlash\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"captureTimestamp\",\"type\":\"uint48\",\"indexed\":false,\"internalType\":\"uint48\"},{\"name\":\"slashedAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDelegator\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositLimit\",\"inputs\":[{\"name\":\"limit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositWhitelist\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetDepositorWhitelistStatus\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetIsDepositLimit\",\"inputs\":[{\"name\":\"status\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SetSlasher\",\"inputs\":[{\"name\":\"slasher\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"withdrawer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"claimer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"burnedShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"mintedShares\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressInsufficientBalance\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadySet\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CheckpointUnorderedInsertion\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DelegatorAlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DepositLimitReached\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedInnerCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientClaim\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientRedemption\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientWithdrawal\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidAccount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCaptureEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidClaimer\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCollateral\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDelegator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEpochDuration\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidLengthEpochs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOnBehalfOf\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidTimestamp\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MathOverflowedMulDiv\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MissingRoles\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoPreviousEpoch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotDelegator\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFactory\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotSlasher\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWhitelistedDepositor\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeCastOverflowedUintDowncast\",\"inputs\":[{\"name\":\"bits\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SlasherAlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooMuchRedeem\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TooMuchWithdraw\",\"inputs\":[]}]",
	Bin: "0x60e03461013c57601f614b8b38819003918201601f1916830192916001600160401b03918285118486101761014057816060928592604097885283398101031261013c5761004c82610154565b916100648461005d60208401610154565b9201610154565b9260805260a0527ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a009081549060ff82861c1661012b5780808316036100e7575b50505060c05251614a22908161016982396080518181816111a001526116d5015260a05181818161183e0152611d04015260c05181818161077001526109a30152f35b6001600160401b0319909116811790915582519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d290602090a15f80806100a4565b845163f92ee8a960e01b8152600490fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b038216820361013c5756fe60806040526004361015610011575f80fd5b5f3560e01c806301ffc9a7146104345780631415519b1461042f5780631b66c9e11461042a5780631e9a695014610425578063248a9ca31461042057806327810b6e1461041b578063281f5752146104165780632abe3048146104115780632d73c69c1461040c5780632dd31000146104075780632f2ff15d1461040257806336568abe146103fd578063392e53cd146103f85780634105a7dd146103f357806346361671146103ee57806347e7ef24146103e957806348d3b775146103e45780634ff0876a146103df57806350861adc146103da57806350f22068146103d55780635346e34f146103d057806354fd4d50146103cb57806357ec83cc146103c657806359f769a9146103c15780635cc07076146103bc57806361a8c8c4146103b75780636da3e06d146103b25780636ec1e3f8146103ad578063715018a6146103a85780637278e31c146103a357806373790ab31461039e5780637667180814610399578063794b15b7146103945780637953b33b1461038f5780637c04c80a1461038a578063810da75d1461038557806383cd9cc31461038057806387df07881461037b5780638b0e9f3f146103765780638da5cb5b1461037157806391d148541461036c5780639d66201b146103675780639f86fd8514610362578063a1b122021461035d578063a217fddf14610358578063a21a1df914610353578063a28614661461034e578063a3b5417214610349578063a5d0322314610344578063aabc24961461033f578063aad3ec961461033a578063afba70ad14610335578063b134427114610330578063bd49c35f1461032b578063bdc8144b14610326578063bfefcd7b14610321578063c31e8dd71461031c578063ce9b793014610317578063d547741f14610312578063d8dfeb451461030d578063db38871514610308578063ecf7085814610303578063efb559d6146102fe578063f2fde38b146102f9578063f3fef3a3146102f45763f5e7ee0f146102ef575f80fd5b6123a1565b6122fb565b6122ce565b61217f565b612162565b612128565b6120fd565b6120b1565b612089565b61201f565b612005565b611f27565b611f0d565b611ee5565b611ebb565b611dfa565b611c43565b611bfd565b611bb0565b611abd565b611a83565b611a69565b611a45565b61199b565b611961565b611903565b6118cf565b61186d565b611829565b611670565b611657565b6115f1565b6115ca565b61158a565b611570565b61152f565b61125b565b6111f4565b6111cf565b61118b565b611138565b61110e565b6110e7565b610faf565b610f7d565b610ee8565b610ecf565b610e72565b610e52565b610e31565b610b73565b610b4b565b610ab0565b610a6b565b610a20565b6109d2565b61098e565b61091a565b61072b565b61065d565b610635565b6105fc565b610513565b6104c8565b61048e565b3461048a57602036600319011261048a5760043563ffffffff60e01b811680910361048a57602090637965db0b60e01b8114908115610479575b506040519015158152f35b6301ffc9a760e01b1490505f61046e565b5f80fd5b3461048a575f36600319011261048a5760206040517fc6aaadd7371d5e8f9ed6849dd66a66573a3ba37167d03f4352c9ba5693678fac8152f35b3461048a575f36600319011261048a5760206040517f9c56d972d63cbb4195b3c1484691dfc220fa96a4c47e7b6613bd82a022029e068152f35b6001600160a01b0381160361048a57565b3461048a57604036600319011261048a5760043561053081610502565b6024359061053c612bf5565b6001600160a01b038116156105eb57335f908152600e6020526040902061056790613c8a565b613c8a565b82116105d957610587610578613b9d565b610580613c16565b9084613cd7565b9182156105c4578261059892612c24565b906105af60015f805160206149ad83398151915255565b604080519182526020820192909252f35b0390f35b60405160016245ddc760e11b03198152600490fd5b604051632418411f60e11b8152600490fd5b6040516208978560e71b8152600490fd5b3461048a57602036600319011261048a576004355f525f8051602061498d8339815191526020526020600160405f200154604051908152f35b3461048a575f36600319011261048a576001546040516001600160a01b039091168152602090f35b3461048a575f36600319011261048a5761067561291a565b80156106d65760015465ffffffffffff8160a01c16915f1981019081116106d1576106a39160d01c906123fa565b81018091116106d1576106b86105c091612d9e565b60405165ffffffffffff90911681529081906020820190565b6123cb565b604051639fa56a5b60e01b8152600490fd5b600435906001600160401b038216820361048a57565b9181601f8401121561048a578235916001600160401b03831161048a576020838186019501011161048a57565b3461048a57604036600319011261048a576107446106e8565b6001600160401b0360243581811161048a576107649036906004016106fe565b505061076e612bf5565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031633036107ea575f805160206149cd833981519152549060ff8260401c169283156107da575b5050506107c8575f80fd5b60405163f92ee8a960e01b8152600490fd5b81169116101590505f80806107bd565b604051631966391b60e11b8152600490fd5b65ffffffffffff81160361048a57565b634e487b7160e01b5f52604160045260245ffd5b606081019081106001600160401b0382111761083b57604052565b61080c565b604081019081106001600160401b0382111761083b57604052565b90601f801991011681019081106001600160401b0382111761083b57604052565b6040519061016082018281106001600160401b0382111761083b57604052565b604051906108a982610840565b565b6001600160401b03811161083b57601f01601f191660200190565b9291926108d2826108ab565b916108e0604051938461085b565b82948184528183011161048a578281602093845f960137010152565b9080601f8301121561048a57816020610917933591016108c6565b90565b3461048a57606036600319011261048a5760043561093781610502565b60243590610944826107fc565b604435906001600160401b03821161048a5760209261096a6109869336906004016108fc565b6001600160a01b039092165f908152600e855260409020612f8b565b604051908152f35b3461048a575f36600319011261048a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461048a57604036600319011261048a57610a1e6024356004356109f582610502565b805f525f8051602061498d833981519152602052610a19600160405f200154613162565b613462565b005b3461048a57604036600319011261048a57602435610a3d81610502565b336001600160a01b03821603610a5957610a1e90600435613498565b60405163334bd91960e11b8152600490fd5b3461048a575f36600319011261048a5760ff60025460a01c1680610a97575b6020906040519015158152f35b5060035460a01c60ff16610a8a565b8015150361048a57565b3461048a57602036600319011261048a57600435610acd81610aa6565b610ad5612bf5565b610add613057565b5f54901515908160ff8216151514610b395760ff191660ff8216175f556040519081527f3e12b7b36c75ac9609a3f58609b331210428e1a85909132638955ba0301eec3390602090a160015f805160206149ad83398151915255005b60405163a741a04560e01b8152600490fd5b3461048a575f36600319011261048a57602065ffffffffffff60015460a01c16604051908152f35b3461048a5760408060031936011261048a576004803591610b9383610502565b610b9b612bf5565b6001600160a01b0391838316908115610e23575f5460ff811680610e0a575b610dfb5783516370a0823160e01b808252308483019081529096602093909160109190911c82169084908490819083010381845afa928315610dd3575f93610dd8575b50610c0f90602435903090339061352a565b5f5486519788523085890190815290979184918391908290819060200103918b60101c165afa908115610dd357610c4e935f92610da6575b50506123ed565b938415610d975760081c60ff1680610d82575b610d745750906105c091610d12610c76613b9d565b95610cbf610c82613c16565b610caa610c908a838b613cd7565b99610ca48a610c9e42612d9e565b9261241b565b9061367e565b5050610cb989610c9e42612d9e565b9061375e565b50506001600160a01b0381165f908152600e6020526040902090610d0c88610d07610562610cec42612d9e565b6001600160a01b039095165f908152600e6020526040902090565b61241b565b9161383a565b505081518481526020810186905233907fdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d790604090a3610d5e60015f805160206149ad83398151915255565b5191825260208201929092529081906040820190565b8251632484557960e01b8152fd5b50610d8f84610d07613b9d565b815410610c61565b50825163070f6eed60e11b8152fd5b610dc59250803d10610dcc575b610dbd818361085b565b810190612449565b5f80610c47565b503d610db3565b612458565b610c0f919350610df490853d8711610dcc57610dbd818361085b565b9290610bfd565b5082516304f63b8560e01b8152fd5b50335f9081526005602052604090205460ff1615610bba565b8251630d534ce360e11b8152fd5b3461048a575f36600319011261048a57602060ff5f54166040519015158152f35b3461048a575f36600319011261048a57602060015460d01c604051908152f35b3461048a575f36600319011261048a57602060ff60025460a01c166040519015158152f35b90604060031983011261048a57600435610eb0816107fc565b91602435906001600160401b03821161048a57610917916004016108fc565b3461048a576020610986610ee236610e97565b90612e01565b3461048a57602036600319011261048a57600435610f0581610aa6565b610f0d612bf5565b610f156130b4565b5f54901515908160ff8260081c16151514610b395761ff001916600882901b61ff0016175f556040519081527ffa7a25a0b611d4ba3c0ea990e90dc23d484a5dd7a1be4733fef2946ba74530c690602090a1610a1e60015f805160206149ad83398151915255565b3461048a575f36600319011261048a5760206001600160401b035f805160206149cd8339815191525416604051908152f35b3461048a57606036600319011261048a57610fc86106e8565b602435610fd481610502565b6001600160401b039160443583811161048a57610ff59036906004016106fe565b5f805160206149cd8339815191529491855494838616806110d65760ff8760401c169081156110c9575b506107c8577fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2966110c49661108c958716906001600160401b0319161790556110875f805160206149cd833981519152600160401b68ff000000000000000019825416179055565b612463565b5f805160206149cd833981519152805468ff0000000000000000191690556040516001600160401b0390911681529081906020820190565b0390a1005b905084861611155f61101f565b60405162dc149f60e41b8152600490fd5b3461048a57602036600319011261048a57602061098660043561110981610502565b612883565b3461048a57602036600319011261048a576004355f526006602052602060405f2054604051908152f35b3461048a575f36600319011261048a5760015465ffffffffffff61116e818360a01c169261116461291a565b9060d01c906123fa565b82018092116106d157611182602092612d9e565b60405191168152f35b3461048a575f36600319011261048a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461048a575f36600319011261048a57602060ff60035460a01c166040519015158152f35b3461048a575f36600319011261048a5761120c6139cc565b5f8051602061492d83398151915280546001600160a01b031981169091555f906001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b3461048a57604036600319011261048a5760043560243561127b816107fc565b611283612bf5565b6003546112a0906001600160a01b03165b6001600160a01b031690565b330361151d576112ae61291a565b906112b88161294e565b908215158061150c575b8015611503575b6114f1576105c0937ff9d090c096f71cd1659861a9ce5b6f384bceb4e2fa4e4a19edf6489a9b8d56c7926112fb613b9d565b946113166113088261240d565b5f52600660205260405f2090565b549082810361140a575061132a818761241b565b6113348185613a04565b9687806113c3575b50505050505b8361138f575b6040805191825265ffffffffffff929092166020820152908101839052606090a161137f60015f805160206149ad83398151915255565b6040519081529081906020820190565b5f546113be9085906113ac9060101c6001600160a01b0316611294565b6001546001600160a01b031690613ab2565b611348565b6113ff93610c4783610ca46113e96113e2611308986113f99888613a20565b80966123ed565b946113f342612d9e565b926123ed565b9261240d565b555f8080808761133c565b915061141e825f52600660205260405f2090565b549061142e81610d07848a61241b565b906114398286613a04565b978861144a575b5050505050611342565b6114bf93836114729361147f8c61147a6114686114b0998884613a20565b9788958584613a20565b9485926123ed565b6123ed565b938484106114ca575b610c4790610ca461149e9596976113f342612d9e565b6114aa6113088761240d565b556123ed565b915f52600660205260405f2090565b555f80808080611440565b916114e6610c47916114e08661149e97986123ed565b9061241b565b859450929050611488565b604051635e18d2a560e01b8152600490fd5b508282116112c9565b50611516836123df565b82106112c2565b60405163dabc4ad960e01b8152600490fd5b3461048a575f36600319011261048a5760015465ffffffffffff808260a01c169161155861291a565b600181018091116106d15761116e9160d01c906123fa565b3461048a575f36600319011261048a57602061098661291a565b3461048a57602036600319011261048a576004356115a781610502565b60018060a01b03165f526005602052602060ff60405f2054166040519015158152f35b3461048a57602036600319011261048a5760206109866004356115ec816107fc565b61294e565b3461048a57604036600319011261048a5760043561160e81610502565b6024356001600160401b0380821161048a573660238301121561048a57816004013590811161048a573660248260051b8401011161048a576105c092602461137f930190612989565b3461048a57602061098661166a36610e97565b90612ed6565b3461048a5760208060031936011261048a57600480359061169082610502565b611698612bf5565b60025460a01c60ff1661181a576040516302910f8b60e31b81526001600160a01b03838116838301908152909491908290829081906020010381887f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dd3575f916117ed575b50156117dd5760405163fbfa77cf60e01b8152838516949082818581895afa928315610dd3575f936117ae575b50503091160361179f5750600280546001600160a01b03929092166001600160a81b031990921691909117600160a01b1790557fdb2160616f776a37b24808115554e79439bf26cccbbd4438190cc6d28e80ecd15f80a2610a1e60015f805160206149ad83398151915255565b60405163b9f0f17160e01b8152fd5b6117ce929350803d106117d6575b6117c6818361085b565b810190612ae2565b905f80611732565b503d6117bc565b506040516324e5af8d60e21b8152fd5b61180d9150823d8411611813575b611805818361085b565b810190612ac2565b5f611705565b503d6117fb565b604051631380833b60e01b8152fd5b3461048a575f36600319011261048a576040517f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03168152602090f35b3461048a575f36600319011261048a5761188561291a565b61188d613b9d565b90805f52600660205260405f205482018092116106d157600181018091116106d1575f52600660205260405f205481018091116106d157602090604051908152f35b3461048a575f36600319011261048a575f8051602061492d833981519152546040516001600160a01b039091168152602090f35b3461048a57604036600319011261048a57602060ff61195560243561192781610502565b6004355f525f8051602061498d833981519152845260405f209060018060a01b03165f5260205260405f2090565b54166040519015158152f35b3461048a57602036600319011261048a5760043561197e81610502565b60018060a01b03165f52600e602052602061098660405f20613c8a565b3461048a57604036600319011261048a576004356119b881610502565b6024356001600160401b03811161048a575f916119da839236906004016106fe565b90816040519283928337810184815203915af4611a3d60806119fa612af7565b9260206040519485921515828401526040808401528051918291826060860152018484015e5f838284010152601f8019910116810103606081018452018261085b565b805190602001fd5b3461048a575f36600319011261048a57602060ff5f5460081c166040519015158152f35b3461048a575f36600319011261048a5760206040515f8152f35b3461048a575f36600319011261048a5760206040517f4a634bc14d77baf979756509ef4298c6f6318af357828612545267ee2eb792338152f35b3461048a57604036600319011261048a57600435611ada81610502565b602435611ae681610aa6565b611aee612bf5565b611af661310b565b6001600160a01b038216918215611b9e575f8381526005602052604090205460ff16151582151514610b39576001600160a01b03165f9081526005602052604090207ff991b1ecfb5115cbb36a2b2e2240c058406d2acc2fcc6e9e2dc99d845ff70a6291611b8591611b73908260ff801983541691151516179055565b60405190151581529081906020820190565b0390a2610a1e60015f805160206149ad83398151915255565b604051630da30f6560e31b8152600490fd5b3461048a57604036600319011261048a576020611bf4602435611bd281610502565b6004355f526008835260405f209060018060a01b03165f5260205260405f2090565b54604051908152f35b3461048a57604036600319011261048a57602060ff611955602435611c2181610502565b6004355f526009845260405f209060018060a01b03165f5260205260405f2090565b3461048a5760208060031936011261048a576004803590611c6382610502565b611c6b612bf5565b60035460a01c60ff16611deb576001600160a01b03828116939084611cd9575b6003805460ff60a01b1916600160a01b179055847fe7e4c932e03abddfe20f83af42c33627e816115c7ec2b168441f65dc14bfc3ba5f80a2610a1e60015f805160206149ad83398151915255565b6040516302910f8b60e31b81526001600160a01b0385168482019081528390829081906020010381857f0000000000000000000000000000000000000000000000000000000000000000165afa908115610dd3575f91611dce575b5015611dbd5760405163fbfa77cf60e01b815282818581895afa928315610dd3575f93611d9e575b505030911603611d8f5750600380546001600160a01b0319166001600160a01b039092169190911790555f808080611c8b565b604051633f3e089160e21b8152fd5b611db5929350803d106117d6576117c6818361085b565b905f80611d5c565b60405163dabc4ad960e01b81528390fd5b611de59150833d851161181357611805818361085b565b5f611d34565b60405163703fe2e560e01b8152fd5b3461048a57604036600319011261048a57600435611e1781610502565b60243590611e23612bf5565b6001600160a01b0391818316908115611ea957611e546105c094611e4683613b01565b9485915f5460101c16613ab2565b604080519182526020820184905233917f865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e946459190a360015f805160206149ad833981519152556040519081529081906020820190565b604051634e46966960e11b8152600490fd5b3461048a57602036600319011261048a576004355f526007602052602060405f2054604051908152f35b3461048a575f36600319011261048a576003546040516001600160a01b039091168152602090f35b3461048a575f36600319011261048a576020610986613b9d565b3461048a57602036600319011261048a57600435611f43612bf5565b335f9081527f2d4ccbfe9f7672a6f2537c1e5bce6adacb73706879154bbab6ce096012d6721660205260409020547f4a634bc14d77baf979756509ef4298c6f6318af357828612545267ee2eb792339060ff1615611fe757508060045414610b39576020817f854df3eb95564502c8bc871ebdd15310ee26270f955f6c6bd8cea68e75045bc092600455604051908152a160015f805160206149ad83398151915255005b6044906040519063e2517d3f60e01b82523360048301526024820152fd5b3461048a575f36600319011261048a576020610986613c16565b3461048a57602036600319011261048a5760043561203c81610502565b61204461291a565b61204d82612883565b916120588183612b97565b83018093116106d157600182018092116106d15761207591612b97565b81018091116106d157602090604051908152f35b3461048a575f36600319011261048a576002546040516001600160a01b039091168152602090f35b3461048a57604036600319011261048a57610a1e6024356004356120d482610502565b805f525f8051602061498d8339815191526020526120f8600160405f200154613162565b613498565b3461048a575f36600319011261048a575f5460405160109190911c6001600160a01b03168152602090f35b3461048a575f36600319011261048a5760206040517fbae4ee3de6c709ff9a002e774c5b78cb381560b219213c88ae0f1e207c03c0238152f35b3461048a575f36600319011261048a576020600454604051908152f35b3461048a57606036600319011261048a5760043561219c81610502565b602435906121a9826107fc565b6001600160401b039160443583811161048a576121ca9036906004016106fe565b909260409384516121da81610820565b60608152606060208201526060868201529280612239575b6105c08661222a878761222461220b8984845191612428565b928561221b602085015183612ed6565b93015190612e01565b91613cd7565b90519081529081906020820190565b810193925060208185031261048a5780359086821161048a57019260608482031261048a5784519561226a87610820565b843581811161048a578261227f9187016108fc565b8752602085013581811161048a57826122999187016108fc565b60208801528585013590811161048a576105c0966122c26122249361220b9361222a98016108fc565b878201529394506121f2565b3461048a57602036600319011261048a57610a1e6004356122ee81610502565b6122f66139cc565b612b26565b3461048a57604036600319011261048a5760043561231881610502565b602435612323612bf5565b6001600160a01b038216156105eb57801561238f57612352612343613c16565b61234b613b9d565b9083614121565b335f908152600e6020526040902090929061236c90613c8a565b831161237d57610598918391612c24565b6040516321727a3760e11b8152600490fd5b604051636c6c627d60e11b8152600490fd5b3461048a57604036600319011261048a5760206109866024356123c381610502565b600435612b97565b634e487b7160e01b5f52601160045260245ffd5b5f198101919082116106d157565b919082039182116106d157565b818102929181159184041417156106d157565b90600182018092116106d157565b919082018092116106d157565b6001600160a01b03165f908152600e60205260409020610917929190612f8b565b9081602091031261048a575190565b6040513d5f823e3d90fd5b916124b5916124a691612474614268565b61247c614268565b60015f805160206149ad833981519152556001600160a01b0394808616612869575b5036916108c6565b60208082518301019101613903565b805182166001600160a01b03161561285757604080820165ffffffffffff6124e3825165ffffffffffff1690565b16156128465760c083018051909290612504906001600160a01b0316611294565b156126f7575b5091612643610140926125e26125bd6126a196612558612530875160018060a01b031690565b5f805462010000600160b01b03191660109290921b62010000600160b01b0316919091179055565b6020860151600180546001600160a01b0319166001600160a01b039092169190911790556125b161258842612d9e565b6001805465ffffffffffff60a01b191660a09290921b65ffffffffffff60a01b16919091179055565b5165ffffffffffff1690565b600180546001600160d01b031660d09290921b6001600160d01b031916919091179055565b6126026125f26060850151151590565b60ff80195f541691151516175f55565b6126296126126080850151151590565b61ff005f5491151560081b169061ff001916175f55565b61263660a0840151600455565b516001600160a01b031690565b8481166126e7575b5060e08101516001600160a01b03168481166126d7575b506101008101516001600160a01b03168481166126c7575b506101208101516001600160a01b03168481166126b7575b5001516001600160a01b031690565b9081166126ab5750565b6126b4906133f3565b50565b6126c090613384565b505f612692565b6126d090613315565b505f61267a565b6126e09061323f565b505f612662565b6126f09061318e565b505f61264b565b60e084015161270e906001600160a01b0316611294565b156127e3575b61012084015161272c906001600160a01b0316611294565b61250a576080840151156127995760a0840151158061277a575b61276a575091612643610140926125e26125bd6126a1965b9496505050925061250a565b5163183c854560e21b8152600490fd5b50610140840151612793906001600160a01b0316611294565b15612746565b60a0840151158015906127c3575b61276a575091612643610140926125e26125bd6126a19661275e565b506101408401516127dc906001600160a01b0316611294565b15156127a7565b60608401511561281957610100840151612805906001600160a01b0316611294565b612714575163183c854560e21b8152600490fd5b610100840151612831906001600160a01b0316611294565b15612714575163183c854560e21b8152600490fd5b81516368f5f8f160e11b8152600490fd5b6040516368f7a67560e11b8152600490fd5b61287d90612875614268565b6122f6614268565b5f61249e565b6001600160a01b03165f908152600e602052604090206128a290613c8a565b6128aa613b9d565b6128b2613c16565b90600181018091116106d157600182018092116106d15761091792613a20565b65ffffffffffff91821690821603919082116106d157565b634e487b7160e01b5f52601260045260245ffd5b9065ffffffffffff80911691821561291557160490565b6128ea565b61292342612d9e565b61294a60015461294065ffffffffffff93848360a01c16906128d2565b9060d01c906128fe565b1690565b6001549065ffffffffffff91828160a01c169182848216106129775761294a92612940916128d2565b60405163b7d0949760e01b8152600490fd5b90612992612bf5565b5f926001600160a01b038316928315611ea9578115612a4a575f5b828110612a2857505f547f326b6aff1cd2fb1c8234de4f9dcfb9047c5c36eb9ef2eb34af5121e969a75d2792869290916129fd9184916129f89060101c6001600160a01b0316611294565b613ab2565b612a0e604051928392339684612a85565b0390a3906108a960015f805160206149ad83398151915255565b94612a436001916114e0612a3d898789612a70565b35613b01565b95016129ad565b60405163edf3b93360e01b8152600490fd5b634e487b7160e01b5f52603260045260245ffd5b9190811015612a805760051b0190565b612a5c565b604080825281018390529392916001600160fb1b03811161048a5760209160609160051b809183880137850101930152565b51906108a982610aa6565b9081602091031261048a575161091781610aa6565b51906108a982610502565b9081602091031261048a575161091781610502565b3d15612b21573d90612b08826108ab565b91612b16604051938461085b565b82523d5f602084013e565b606090565b6001600160a01b03908116908115612b7f575f8051602061492d83398151915280546001600160a01b031981168417909155167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e05f80a3565b604051631e4fbdf760e01b81525f6004820152602490fd5b90612bbd90825f52600860205260405f209060018060a01b03165f5260205260405f2090565b54905f52600660205260405f2054600760205260405f205490600181018091116106d157600182018092116106d15761091792613a20565b5f805160206149ad8339815191526002815414612c125760029055565b604051633ee5aeb560e01b8152600490fd5b335f908152600e6020526040902093929190612c3f42612d9e565b335f908152600e60205260409020612c5690613c8a565b918483039283116106d157612c917febff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f938593612d1c9961383a565b5050612cab612c9f42612d9e565b610cb98761147a613c16565b5050612cc5612cb942612d9e565b610ca48461147a613b9d565b5050612d6a81612d55612cde612cd961291a565b61240d565b612d34612cf3825f52600660205260405f2090565b54612d06835f52600760205260405f2090565b54612d1282828b613cd7565b9d8e998a9361241b565b612d2e855f52600660205260405f2090565b5561241b565b612d46825f52600760205260405f2090565b555f52600860205260405f2090565b9060018060a01b03165f5260205260405f2090565b612d7583825461241b565b9055604080519485526020850195909552938301526001600160a01b03909216913391606090a3565b65ffffffffffff90818111612db1571690565b604490604051906306dfcc6560e41b8252603060048301526024820152fd5b9081602091031261048a575163ffffffff8116810361048a5790565b90600163ffffffff809316019182116106d157565b90805115612ecc5780602080612e1c93518301019101612dd0565b612e2581613f48565b90612e36825165ffffffffffff1690565b9065ffffffffffff8085169216828114612ec05782119182612e6c575b5050612e63575061091790613cf8565b60209150015190565b909150612e7a600a546123df565b63ffffffff831614918215612e93575b50505f80612e53565b612eb89192506125b1612ea8612ead92612dec565b613f48565b65ffffffffffff1690565b115f80612e8a565b50505060209150015190565b5061091790613cf8565b90805115612f815780602080612ef193518301019101612dd0565b612efa81613fcd565b90612f0b825165ffffffffffff1690565b9065ffffffffffff8085169216828114612ec05782119182612f38575b5050612e63575061091790613dbf565b909150612f46600c546123df565b63ffffffff831614918215612f5f575b50505f80612f28565b612f799192506125b1612f74612ead92612dec565b613fcd565b115f80612f56565b5061091790613dbf565b9180511561304d5780602080612fa693518301019101612dd0565b91612fb18382614048565b92612fc2845165ffffffffffff1690565b9065ffffffffffff80851692168281146130415782119182612ff8575b5050612fef576109179250613e7f565b50506020015190565b90915061300583546123df565b63ffffffff83161491821561301e575b50505f80612fdf565b6130399192506125b1613033612ead92612dec565b85614048565b115f80613015565b50505050506020015190565b5061091791613e7f565b335f9081527f03b26b53c87faa1058f30bb5d2fee64f717719a0d9b7e4f4df027aa5e9b6ed69602052604090207fbae4ee3de6c709ff9a002e774c5b78cb381560b219213c88ae0f1e207c03c0239060ff905b541615611fe75750565b335f9081527fe41986b390c4ce6791b87f7423b5fab7a81a0e6d89660b6264f697d008b2243a602052604090207fc6aaadd7371d5e8f9ed6849dd66a66573a3ba37167d03f4352c9ba5693678fac9060ff906130aa565b335f9081527f0287ab4cb376ea253d7b2ce0c2cbf458f7deae9458470d74afdfbbee484e9540602052604090207f9c56d972d63cbb4195b3c1484691dfc220fa96a4c47e7b6613bd82a022029e069060ff906130aa565b5f8181525f8051602061498d83398151915260209081526040808320338452909152902060ff906130aa565b6001600160a01b0381165f9081527fb7db2dd08fcb62d0c9e08c51941cae53c267786a0b75803fb7960902fc8ef97d60205260409020545f8051602061498d8339815191529060ff16613239575f808052602091825260408082206001600160a01b038516835290925220805460ff1916600117905533906001600160a01b03165f7f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d8180a4600190565b50505f90565b6001600160a01b0381165f9081527f03b26b53c87faa1058f30bb5d2fee64f717719a0d9b7e4f4df027aa5e9b6ed69602052604090207fbae4ee3de6c709ff9a002e774c5b78cb381560b219213c88ae0f1e207c03c023905f8051602061498d8339815191529060ff905b541661330e575f828152602091825260408082206001600160a01b038616835290925220805460ff1916600117905533916001600160a01b0316907f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d5f80a4600190565b5050505f90565b6001600160a01b0381165f9081527f0287ab4cb376ea253d7b2ce0c2cbf458f7deae9458470d74afdfbbee484e9540602052604090207f9c56d972d63cbb4195b3c1484691dfc220fa96a4c47e7b6613bd82a022029e06905f8051602061498d8339815191529060ff906132aa565b6001600160a01b0381165f9081527fe41986b390c4ce6791b87f7423b5fab7a81a0e6d89660b6264f697d008b2243a602052604090207fc6aaadd7371d5e8f9ed6849dd66a66573a3ba37167d03f4352c9ba5693678fac905f8051602061498d8339815191529060ff906132aa565b6001600160a01b0381165f9081527f2d4ccbfe9f7672a6f2537c1e5bce6adacb73706879154bbab6ce096012d67216602052604090207f4a634bc14d77baf979756509ef4298c6f6318af357828612545267ee2eb79233905f8051602061498d8339815191529060ff906132aa565b5f8181525f8051602061498d833981519152602081815260408084206001600160a01b038716855290915290912060ff906132aa565b5f8181525f8051602061498d833981519152602081815260408084206001600160a01b03871685529091529091205460ff161561330e575f828152602091825260408082206001600160a01b038616835290925220805460ff1916905533916001600160a01b0316907ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b5f80a4600190565b6040516323b872dd60e01b60208201526001600160a01b03928316602482015292909116604483015260648083019390935291815260a08101918183106001600160401b0384111761083b576108a9926040526140b0565b600b54811015612a8057600b5f525f8051602061496d83398151915201905f90565b600d54811015612a8057600d5f525f8051602061494d83398151915201905f90565b8054821015612a80575f5260205f2001905f90565b600d54600160401b81101561083b576001810180600d55811015612a8057600d5f525f8051602061494d8339815191520155565b600b54600160401b81101561083b576001810180600b55811015612a8057600b5f525f8051602061496d8339815191520155565b90815491600160401b83101561083b57826136669160016108a9950181556135c6565b90919082549060031b91821b915f19901b1916179055565b919091600d9081548015613729575b50613696614185565b509290918154926136a5613b9d565b9481613714575b50156136f157505f1982018281116106d15781541115612a80575f527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb4018390559190565b613705926001600160d01b0316915061439f565b5050613710836135db565b9190565b65ffffffffffff83811691161490505f6136ac565b600160401b81101561083b5760018101808455811015612a80575f908382525f8051602061494d83398151915201555f61368d565b919091600b9081548015613805575b506137766141f0565b50929091815492613785613c16565b94816137f0575b50156137d157505f1982018281116106d15781541115612a80575f527f0175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db8018390559190565b6137e5926001600160d01b0316915061449f565b50506137108361360f565b65ffffffffffff83811691161490505f61378c565b600160401b81101561083b5760018101808455811015612a80575f908382525f8051602061496d83398151915201555f61376d565b909291926001820190815480156138c8575b50846138578461423a565b5094909284549361386783613c8a565b96816138b3575b5015613890575050505f1981019081116106d1576136668591613710936135c6565b6137109492936138ac926001600160d01b039091169190614549565b5050613643565b65ffffffffffff83811691161490505f61386e565b600160401b81101561083b578060016138e492018455836135c6565b8154905f199060031b1b191690555f61384c565b51906108a9826107fc565b908161016091031261048a5761391761087c565b9061392181612ad7565b825261392f60208201612ad7565b6020830152613940604082016138f8565b604083015261395160608201612ab7565b606083015261396260808201612ab7565b608083015260a081015160a083015261397d60c08201612ad7565b60c083015261398e60e08201612ad7565b60e08301526101006139a1818301612ad7565b908301526101206139b3818301612ad7565b908301526139c5610140809201612ad7565b9082015290565b5f8051602061492d833981519152546001600160a01b031633036139ec57565b60405163118cdaa760e01b8152336004820152602490fd5b9080821015613a11575090565b905090565b8115612915570490565b9091828202915f1984820993838086109503948086039514613aa55784831115613a9357829109815f038216809204600280826003021880830282030280830282030280830282030280830282030280830282030280920290030293600183805f03040190848311900302920304170290565b60405163227bc15360e01b8152600490fd5b5050906109179250613a16565b60405163a9059cbb60e01b60208201526001600160a01b03909216602483015260448083019390935291815260808101916001600160401b0383118284101761083b576108a9926040526140b0565b90613b0a61291a565b821015613b8b575f82815260096020908152604080832033845290915290205460ff16613b7957613b3b3383612b97565b918215613b67575f9081526009602090815260408083203384529091529020805460ff19166001179055565b6040516366c0bcbf60e11b8152600490fd5b604051630c8d9eab60e31b8152600490fd5b60405163d5b25b6360e01b8152600490fd5b600c5480613bdd57505f5b6001600160d01b03811615613bd857600d54811015612a8057600d5f525f8051602061494d833981519152015490565b505f90565b805f198101116106d157600c5f527fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c6015460301c613ba8565b600a5480613c5157505f5b6001600160d01b03811615613bd857600b54811015612a8057600b5f525f8051602061496d833981519152015490565b805f198101116106d157600a5f527fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a7015460301c613c21565b805480613cb957505f5b6001600160d01b03811615613239576001613caf92016135c6565b90549060031b1c90565b5f199080828101116106d157825f5260205f2001015460301c613c94565b9190600181018091116106d157600182018092116106d15761091792613a20565b600a54905f829160058411613d66575b613d129350614755565b80613d3557505f5b6001600160d01b03811615613bd857613caf61091791613582565b613d5a613d44613d61926123df565b600a5f525f8051602061490d8339815191520190565b5460301c90565b613d1a565b9192613d718161461c565b81039081116106d157613d1293600a5f5265ffffffffffff80835f8051602061490d833981519152015416908516105f14613dad575091613d08565b929150613db99061240d565b90613d08565b600c54905f829160058411613e26575b613dd993506147be565b80613dfc57505f5b6001600160d01b03811615613bd857613caf610917916135a4565b613d5a613e0b613e21926123df565b600c5f525f805160206148ed8339815191520190565b613de1565b9192613e318161461c565b81039081116106d157613dd993600c5f5265ffffffffffff80835f805160206148ed833981519152015416908516105f14613e6d575091613dcf565b929150613e799061240d565b90613dcf565b908154905f829160058411613ee2575b613e9a935084614827565b80613ec257505f905b6001600160d01b0382161561323957610917916001613caf92016135c6565b613d5a613ed1613edc926123df565b835f5260205f200190565b90613ea3565b9192613eed8161461c565b81039081116106d157613e9a93855f5265ffffffffffff808360205f20015416908516105f14613f1e575091613e8f565b929150613f2a9061240d565b90613e8f565b60405190613f3d82610840565b5f6020838281520152565b613f50613f30565b50613f59613f30565b50600a54811015612a8057613f8090600a5f525f8051602061490d83398151915201614161565b65ffffffffffff81511690602060018060d01b0391015116600b54811015612a8057600b5f525f8051602061496d833981519152015460405191613fc383610840565b8252602082015290565b613fd5613f30565b50613fde613f30565b50600c54811015612a805761400590600c5f525f805160206148ed83398151915201614161565b65ffffffffffff81511690602060018060d01b0391015116600d54811015612a8057600d5f525f8051602061494d833981519152015460405191613fc383610840565b9061407161406b6140929261405b613f30565b50614064613f30565b50846135c6565b50614161565b91600165ffffffffffff845116936020828060d01b039101511691016135c6565b905490604051926140a284610840565b835260031b1c602082015290565b5f806140d89260018060a01b03169360208151910182865af16140d1612af7565b9083614889565b8051908115159182614106575b50506140ee5750565b60249060405190635274afe760e01b82526004820152fd5b6141199250602080918301019101612ac2565b155f806140e5565b9190600181018091116106d157600182018092116106d157614144828285613a20565b92821561291557096141535790565b600181018091116106d15790565b9060405161416e81610840565b915465ffffffffffff8116835260301c6020830152565b600c548061419657505f905f905f90565b805f198101116106d1577fdf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c66141cf91600c5f5201614161565b9065ffffffffffff82511691602060018060d01b0391015116906001929190565b600a548061420157505f905f905f90565b805f198101116106d1577fc65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a76141cf91600a5f5201614161565b80548061424b5750505f905f905f90565b5f199080828101116106d1576141cf925f5260205f200101614161565b60ff5f805160206149cd8339815191525460401c161561428457565b604051631afcd79f60e31b8152600490fd5b600c5490600160401b82101561083b576001820180600c55821015612a8057600c5f52805160209091015160301b65ffffffffffff191665ffffffffffff91909116175f805160206148ed83398151915290910155565b600a5490600160401b82101561083b576001820180600a55821015612a8057600a5f52805160209091015160301b65ffffffffffff191665ffffffffffff91909116175f8051602061490d83398151915290910155565b8054600160401b81101561083b57614361916001820181556135c6565b61438c57815160209092015160301b65ffffffffffff191665ffffffffffff92909216919091179055565b634e487b7160e01b5f525f60045260245ffd5b600c54919291908115614476576143c06143bb613e0b846123df565b614161565b805165ffffffffffff83811694911684106144645785602093613710956143f0612ead865165ffffffffffff1690565b0361442c5750614405613e0b61441e936123df565b9065ffffffffffff82549181199060301b169116179055565b01516001600160d01b031690565b905061445f915061444c61443e61089c565b65ffffffffffff9092168252565b6001600160d01b03871681850152614296565b61441e565b604051632520601d60e01b8152600490fd5b61449a915061448661443e61089c565b6001600160d01b0384166020820152614296565b5f9190565b600a54919291908115614525576144bb6143bb613d44846123df565b805165ffffffffffff83811694911684106144645785602093613710956144eb612ead865165ffffffffffff1690565b036145005750614405613d4461441e936123df565b905061445f915061451261443e61089c565b6001600160d01b038716818501526142ed565b61449a915061453561443e61089c565b6001600160d01b03841660208201526142ed565b8054929392919082156145f8576145656143bb613ed1856123df565b90614576825165ffffffffffff1690565b65ffffffffffff84811691168110614464576137109460209488926145a4612ead875165ffffffffffff1690565b036145c5575061441e926145ba614405926123df565b905f5260205f200190565b91505061445f916145e56145d761089c565b65ffffffffffff9093168352565b6001600160d01b03881682860152614344565b61449a92506146086145d761089c565b6001600160d01b0385166020830152614344565b8015613bd857806146ee6146e76146dd6146d36146c96146bf6146b56146ab60016109179a5f908b60801c80614749575b508060401c8061473c575b508060201c8061472f575b508060101c80614722575b508060081c80614715575b508060041c80614708575b508060021c806146fb575b50821c6146f4575b811c1b6146a4818b613a16565b0160011c90565b6146a4818a613a16565b6146a48189613a16565b6146a48188613a16565b6146a48187613a16565b6146a48186613a16565b6146a48185613a16565b8092613a16565b90613a04565b8101614697565b600291509101905f61468f565b600491509101905f614684565b600891509101905f614679565b601091509101905f61466e565b602091509101905f614663565b604091509101905f614658565b9150506080905f61464d565b905b82811061476357505090565b9091808216906001818418811c83018093116106d157600a5f5265ffffffffffff80845f8051602061490d833981519152015416908616105f146147ab575050915b90614757565b9093925081018091116106d157906147a5565b905b8281106147cc57505090565b9091808216906001818418811c83018093116106d157600c5f5265ffffffffffff80845f805160206148ed833981519152015416908616105f14614814575050915b906147c0565b9093925081018091116106d1579061480e565b91905b8382106148375750505090565b909192808316906001818518811c83018093116106d157855f5265ffffffffffff808460205f20015416908516105f14614876575050925b919061482a565b9094935081018091116106d1579161486f565b906148b0575080511561489e57805190602001fd5b604051630a12f52160e11b8152600490fd5b815115806148e3575b6148c1575090565b604051639996b31560e01b81526001600160a01b039091166004820152602490fd5b50803b156148b956fedf6966c971051c3d54ec59162606531493a51404a002842f56009d7e5cf4a8c7c65a7bb8d6351c1cf70c95a316cc6a92839c986682d98bc35f958f4883f9d2a89016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300d7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb50175b7a638427703f0dbe7bb9bbf987a2551717b34e79f33b5b1008d1fa01db902dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b6268009b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a26469706673582212205eb48e48c582a0d2dd84283abb0d0c766d767ab9b676d6e9c13aeaa32e7e136664736f6c63430008190033",
}

// SymbioticVaultABI is the input ABI used to generate the binding from.
// Deprecated: Use SymbioticVaultMetaData.ABI instead.
var SymbioticVaultABI = SymbioticVaultMetaData.ABI

// SymbioticVaultBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SymbioticVaultMetaData.Bin instead.
var SymbioticVaultBin = SymbioticVaultMetaData.Bin

// DeploySymbioticVault deploys a new Ethereum contract, binding an instance of SymbioticVault to it.
func DeploySymbioticVault(auth *bind.TransactOpts, backend bind.ContractBackend, delegatorFactory common.Address, slasherFactory common.Address, vaultFactory common.Address) (common.Address, *types.Transaction, *SymbioticVault, error) {
	parsed, err := SymbioticVaultMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SymbioticVaultBin), backend, delegatorFactory, slasherFactory, vaultFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SymbioticVault{SymbioticVaultCaller: SymbioticVaultCaller{contract: contract}, SymbioticVaultTransactor: SymbioticVaultTransactor{contract: contract}, SymbioticVaultFilterer: SymbioticVaultFilterer{contract: contract}}, nil
}

// SymbioticVault is an auto generated Go binding around an Ethereum contract.
type SymbioticVault struct {
	SymbioticVaultCaller     // Read-only binding to the contract
	SymbioticVaultTransactor // Write-only binding to the contract
	SymbioticVaultFilterer   // Log filterer for contract events
}

// SymbioticVaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type SymbioticVaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticVaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SymbioticVaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticVaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SymbioticVaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SymbioticVaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SymbioticVaultSession struct {
	Contract     *SymbioticVault   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SymbioticVaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SymbioticVaultCallerSession struct {
	Contract *SymbioticVaultCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// SymbioticVaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SymbioticVaultTransactorSession struct {
	Contract     *SymbioticVaultTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// SymbioticVaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type SymbioticVaultRaw struct {
	Contract *SymbioticVault // Generic contract binding to access the raw methods on
}

// SymbioticVaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SymbioticVaultCallerRaw struct {
	Contract *SymbioticVaultCaller // Generic read-only contract binding to access the raw methods on
}

// SymbioticVaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SymbioticVaultTransactorRaw struct {
	Contract *SymbioticVaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSymbioticVault creates a new instance of SymbioticVault, bound to a specific deployed contract.
func NewSymbioticVault(address common.Address, backend bind.ContractBackend) (*SymbioticVault, error) {
	contract, err := bindSymbioticVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SymbioticVault{SymbioticVaultCaller: SymbioticVaultCaller{contract: contract}, SymbioticVaultTransactor: SymbioticVaultTransactor{contract: contract}, SymbioticVaultFilterer: SymbioticVaultFilterer{contract: contract}}, nil
}

// NewSymbioticVaultCaller creates a new read-only instance of SymbioticVault, bound to a specific deployed contract.
func NewSymbioticVaultCaller(address common.Address, caller bind.ContractCaller) (*SymbioticVaultCaller, error) {
	contract, err := bindSymbioticVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultCaller{contract: contract}, nil
}

// NewSymbioticVaultTransactor creates a new write-only instance of SymbioticVault, bound to a specific deployed contract.
func NewSymbioticVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*SymbioticVaultTransactor, error) {
	contract, err := bindSymbioticVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultTransactor{contract: contract}, nil
}

// NewSymbioticVaultFilterer creates a new log filterer instance of SymbioticVault, bound to a specific deployed contract.
func NewSymbioticVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*SymbioticVaultFilterer, error) {
	contract, err := bindSymbioticVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultFilterer{contract: contract}, nil
}

// bindSymbioticVault binds a generic wrapper to an already deployed contract.
func bindSymbioticVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SymbioticVaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticVault *SymbioticVaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticVault.Contract.SymbioticVaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticVault *SymbioticVaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SymbioticVaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticVault *SymbioticVaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SymbioticVaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SymbioticVault *SymbioticVaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SymbioticVault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SymbioticVault *SymbioticVaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticVault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SymbioticVault *SymbioticVaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SymbioticVault.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEFAULTADMINROLE(&_SymbioticVault.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEFAULTADMINROLE(&_SymbioticVault.CallOpts)
}

// DELEGATORFACTORY is a free data retrieval call binding the contract method 0x6da3e06d.
//
// Solidity: function DELEGATOR_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) DELEGATORFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "DELEGATOR_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DELEGATORFACTORY is a free data retrieval call binding the contract method 0x6da3e06d.
//
// Solidity: function DELEGATOR_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) DELEGATORFACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.DELEGATORFACTORY(&_SymbioticVault.CallOpts)
}

// DELEGATORFACTORY is a free data retrieval call binding the contract method 0x6da3e06d.
//
// Solidity: function DELEGATOR_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) DELEGATORFACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.DELEGATORFACTORY(&_SymbioticVault.CallOpts)
}

// DEPOSITORWHITELISTROLE is a free data retrieval call binding the contract method 0x1b66c9e1.
//
// Solidity: function DEPOSITOR_WHITELIST_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) DEPOSITORWHITELISTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "DEPOSITOR_WHITELIST_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITORWHITELISTROLE is a free data retrieval call binding the contract method 0x1b66c9e1.
//
// Solidity: function DEPOSITOR_WHITELIST_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) DEPOSITORWHITELISTROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITORWHITELISTROLE(&_SymbioticVault.CallOpts)
}

// DEPOSITORWHITELISTROLE is a free data retrieval call binding the contract method 0x1b66c9e1.
//
// Solidity: function DEPOSITOR_WHITELIST_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) DEPOSITORWHITELISTROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITORWHITELISTROLE(&_SymbioticVault.CallOpts)
}

// DEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0xa21a1df9.
//
// Solidity: function DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) DEPOSITLIMITSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "DEPOSIT_LIMIT_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0xa21a1df9.
//
// Solidity: function DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) DEPOSITLIMITSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITLIMITSETROLE(&_SymbioticVault.CallOpts)
}

// DEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0xa21a1df9.
//
// Solidity: function DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) DEPOSITLIMITSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITLIMITSETROLE(&_SymbioticVault.CallOpts)
}

// DEPOSITWHITELISTSETROLE is a free data retrieval call binding the contract method 0xdb388715.
//
// Solidity: function DEPOSIT_WHITELIST_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) DEPOSITWHITELISTSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "DEPOSIT_WHITELIST_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEPOSITWHITELISTSETROLE is a free data retrieval call binding the contract method 0xdb388715.
//
// Solidity: function DEPOSIT_WHITELIST_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) DEPOSITWHITELISTSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITWHITELISTSETROLE(&_SymbioticVault.CallOpts)
}

// DEPOSITWHITELISTSETROLE is a free data retrieval call binding the contract method 0xdb388715.
//
// Solidity: function DEPOSIT_WHITELIST_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) DEPOSITWHITELISTSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.DEPOSITWHITELISTSETROLE(&_SymbioticVault.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) FACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) FACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.FACTORY(&_SymbioticVault.CallOpts)
}

// FACTORY is a free data retrieval call binding the contract method 0x2dd31000.
//
// Solidity: function FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) FACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.FACTORY(&_SymbioticVault.CallOpts)
}

// ISDEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0x1415519b.
//
// Solidity: function IS_DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) ISDEPOSITLIMITSETROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "IS_DEPOSIT_LIMIT_SET_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ISDEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0x1415519b.
//
// Solidity: function IS_DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) ISDEPOSITLIMITSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.ISDEPOSITLIMITSETROLE(&_SymbioticVault.CallOpts)
}

// ISDEPOSITLIMITSETROLE is a free data retrieval call binding the contract method 0x1415519b.
//
// Solidity: function IS_DEPOSIT_LIMIT_SET_ROLE() view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) ISDEPOSITLIMITSETROLE() ([32]byte, error) {
	return _SymbioticVault.Contract.ISDEPOSITLIMITSETROLE(&_SymbioticVault.CallOpts)
}

// SLASHERFACTORY is a free data retrieval call binding the contract method 0x87df0788.
//
// Solidity: function SLASHER_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) SLASHERFACTORY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "SLASHER_FACTORY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SLASHERFACTORY is a free data retrieval call binding the contract method 0x87df0788.
//
// Solidity: function SLASHER_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) SLASHERFACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.SLASHERFACTORY(&_SymbioticVault.CallOpts)
}

// SLASHERFACTORY is a free data retrieval call binding the contract method 0x87df0788.
//
// Solidity: function SLASHER_FACTORY() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) SLASHERFACTORY() (common.Address, error) {
	return _SymbioticVault.Contract.SLASHERFACTORY(&_SymbioticVault.CallOpts)
}

// ActiveBalanceOf is a free data retrieval call binding the contract method 0x59f769a9.
//
// Solidity: function activeBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveBalanceOf is a free data retrieval call binding the contract method 0x59f769a9.
//
// Solidity: function activeBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveBalanceOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveBalanceOf(&_SymbioticVault.CallOpts, account)
}

// ActiveBalanceOf is a free data retrieval call binding the contract method 0x59f769a9.
//
// Solidity: function activeBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveBalanceOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveBalanceOf(&_SymbioticVault.CallOpts, account)
}

// ActiveBalanceOfAt is a free data retrieval call binding the contract method 0xefb559d6.
//
// Solidity: function activeBalanceOfAt(address account, uint48 timestamp, bytes hints) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveBalanceOfAt(opts *bind.CallOpts, account common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeBalanceOfAt", account, timestamp, hints)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveBalanceOfAt is a free data retrieval call binding the contract method 0xefb559d6.
//
// Solidity: function activeBalanceOfAt(address account, uint48 timestamp, bytes hints) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveBalanceOfAt(account common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveBalanceOfAt(&_SymbioticVault.CallOpts, account, timestamp, hints)
}

// ActiveBalanceOfAt is a free data retrieval call binding the contract method 0xefb559d6.
//
// Solidity: function activeBalanceOfAt(address account, uint48 timestamp, bytes hints) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveBalanceOfAt(account common.Address, timestamp *big.Int, hints []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveBalanceOfAt(&_SymbioticVault.CallOpts, account, timestamp, hints)
}

// ActiveShares is a free data retrieval call binding the contract method 0xbfefcd7b.
//
// Solidity: function activeShares() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveShares is a free data retrieval call binding the contract method 0xbfefcd7b.
//
// Solidity: function activeShares() view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveShares() (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveShares(&_SymbioticVault.CallOpts)
}

// ActiveShares is a free data retrieval call binding the contract method 0xbfefcd7b.
//
// Solidity: function activeShares() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveShares() (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveShares(&_SymbioticVault.CallOpts)
}

// ActiveSharesAt is a free data retrieval call binding the contract method 0x50f22068.
//
// Solidity: function activeSharesAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveSharesAt(opts *bind.CallOpts, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeSharesAt", timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveSharesAt is a free data retrieval call binding the contract method 0x50f22068.
//
// Solidity: function activeSharesAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveSharesAt(timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesAt(&_SymbioticVault.CallOpts, timestamp, hint)
}

// ActiveSharesAt is a free data retrieval call binding the contract method 0x50f22068.
//
// Solidity: function activeSharesAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveSharesAt(timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesAt(&_SymbioticVault.CallOpts, timestamp, hint)
}

// ActiveSharesOf is a free data retrieval call binding the contract method 0x9d66201b.
//
// Solidity: function activeSharesOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveSharesOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeSharesOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveSharesOf is a free data retrieval call binding the contract method 0x9d66201b.
//
// Solidity: function activeSharesOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveSharesOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesOf(&_SymbioticVault.CallOpts, account)
}

// ActiveSharesOf is a free data retrieval call binding the contract method 0x9d66201b.
//
// Solidity: function activeSharesOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveSharesOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesOf(&_SymbioticVault.CallOpts, account)
}

// ActiveSharesOfAt is a free data retrieval call binding the contract method 0x2d73c69c.
//
// Solidity: function activeSharesOfAt(address account, uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveSharesOfAt(opts *bind.CallOpts, account common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeSharesOfAt", account, timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveSharesOfAt is a free data retrieval call binding the contract method 0x2d73c69c.
//
// Solidity: function activeSharesOfAt(address account, uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveSharesOfAt(account common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesOfAt(&_SymbioticVault.CallOpts, account, timestamp, hint)
}

// ActiveSharesOfAt is a free data retrieval call binding the contract method 0x2d73c69c.
//
// Solidity: function activeSharesOfAt(address account, uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveSharesOfAt(account common.Address, timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveSharesOfAt(&_SymbioticVault.CallOpts, account, timestamp, hint)
}

// ActiveStake is a free data retrieval call binding the contract method 0xbd49c35f.
//
// Solidity: function activeStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStake is a free data retrieval call binding the contract method 0xbd49c35f.
//
// Solidity: function activeStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveStake() (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveStake(&_SymbioticVault.CallOpts)
}

// ActiveStake is a free data retrieval call binding the contract method 0xbd49c35f.
//
// Solidity: function activeStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveStake() (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveStake(&_SymbioticVault.CallOpts)
}

// ActiveStakeAt is a free data retrieval call binding the contract method 0x810da75d.
//
// Solidity: function activeStakeAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) ActiveStakeAt(opts *bind.CallOpts, timestamp *big.Int, hint []byte) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "activeStakeAt", timestamp, hint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ActiveStakeAt is a free data retrieval call binding the contract method 0x810da75d.
//
// Solidity: function activeStakeAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) ActiveStakeAt(timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveStakeAt(&_SymbioticVault.CallOpts, timestamp, hint)
}

// ActiveStakeAt is a free data retrieval call binding the contract method 0x810da75d.
//
// Solidity: function activeStakeAt(uint48 timestamp, bytes hint) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) ActiveStakeAt(timestamp *big.Int, hint []byte) (*big.Int, error) {
	return _SymbioticVault.Contract.ActiveStakeAt(&_SymbioticVault.CallOpts, timestamp, hint)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) Burner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "burner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) Burner() (common.Address, error) {
	return _SymbioticVault.Contract.Burner(&_SymbioticVault.CallOpts)
}

// Burner is a free data retrieval call binding the contract method 0x27810b6e.
//
// Solidity: function burner() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) Burner() (common.Address, error) {
	return _SymbioticVault.Contract.Burner(&_SymbioticVault.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) Collateral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "collateral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) Collateral() (common.Address, error) {
	return _SymbioticVault.Contract.Collateral(&_SymbioticVault.CallOpts)
}

// Collateral is a free data retrieval call binding the contract method 0xd8dfeb45.
//
// Solidity: function collateral() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) Collateral() (common.Address, error) {
	return _SymbioticVault.Contract.Collateral(&_SymbioticVault.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "currentEpoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) CurrentEpoch() (*big.Int, error) {
	return _SymbioticVault.Contract.CurrentEpoch(&_SymbioticVault.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x76671808.
//
// Solidity: function currentEpoch() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) CurrentEpoch() (*big.Int, error) {
	return _SymbioticVault.Contract.CurrentEpoch(&_SymbioticVault.CallOpts)
}

// CurrentEpochStart is a free data retrieval call binding the contract method 0x61a8c8c4.
//
// Solidity: function currentEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCaller) CurrentEpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "currentEpochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentEpochStart is a free data retrieval call binding the contract method 0x61a8c8c4.
//
// Solidity: function currentEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultSession) CurrentEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.CurrentEpochStart(&_SymbioticVault.CallOpts)
}

// CurrentEpochStart is a free data retrieval call binding the contract method 0x61a8c8c4.
//
// Solidity: function currentEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCallerSession) CurrentEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.CurrentEpochStart(&_SymbioticVault.CallOpts)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) Delegator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "delegator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) Delegator() (common.Address, error) {
	return _SymbioticVault.Contract.Delegator(&_SymbioticVault.CallOpts)
}

// Delegator is a free data retrieval call binding the contract method 0xce9b7930.
//
// Solidity: function delegator() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) Delegator() (common.Address, error) {
	return _SymbioticVault.Contract.Delegator(&_SymbioticVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) DepositLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "depositLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) DepositLimit() (*big.Int, error) {
	return _SymbioticVault.Contract.DepositLimit(&_SymbioticVault.CallOpts)
}

// DepositLimit is a free data retrieval call binding the contract method 0xecf70858.
//
// Solidity: function depositLimit() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) DepositLimit() (*big.Int, error) {
	return _SymbioticVault.Contract.DepositLimit(&_SymbioticVault.CallOpts)
}

// DepositWhitelist is a free data retrieval call binding the contract method 0x48d3b775.
//
// Solidity: function depositWhitelist() view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) DepositWhitelist(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "depositWhitelist")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DepositWhitelist is a free data retrieval call binding the contract method 0x48d3b775.
//
// Solidity: function depositWhitelist() view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) DepositWhitelist() (bool, error) {
	return _SymbioticVault.Contract.DepositWhitelist(&_SymbioticVault.CallOpts)
}

// DepositWhitelist is a free data retrieval call binding the contract method 0x48d3b775.
//
// Solidity: function depositWhitelist() view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) DepositWhitelist() (bool, error) {
	return _SymbioticVault.Contract.DepositWhitelist(&_SymbioticVault.CallOpts)
}

// EpochAt is a free data retrieval call binding the contract method 0x7953b33b.
//
// Solidity: function epochAt(uint48 timestamp) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) EpochAt(opts *bind.CallOpts, timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "epochAt", timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochAt is a free data retrieval call binding the contract method 0x7953b33b.
//
// Solidity: function epochAt(uint48 timestamp) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) EpochAt(timestamp *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.EpochAt(&_SymbioticVault.CallOpts, timestamp)
}

// EpochAt is a free data retrieval call binding the contract method 0x7953b33b.
//
// Solidity: function epochAt(uint48 timestamp) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) EpochAt(timestamp *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.EpochAt(&_SymbioticVault.CallOpts, timestamp)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCaller) EpochDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "epochDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint48)
func (_SymbioticVault *SymbioticVaultSession) EpochDuration() (*big.Int, error) {
	return _SymbioticVault.Contract.EpochDuration(&_SymbioticVault.CallOpts)
}

// EpochDuration is a free data retrieval call binding the contract method 0x4ff0876a.
//
// Solidity: function epochDuration() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCallerSession) EpochDuration() (*big.Int, error) {
	return _SymbioticVault.Contract.EpochDuration(&_SymbioticVault.CallOpts)
}

// EpochDurationInit is a free data retrieval call binding the contract method 0x46361671.
//
// Solidity: function epochDurationInit() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCaller) EpochDurationInit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "epochDurationInit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EpochDurationInit is a free data retrieval call binding the contract method 0x46361671.
//
// Solidity: function epochDurationInit() view returns(uint48)
func (_SymbioticVault *SymbioticVaultSession) EpochDurationInit() (*big.Int, error) {
	return _SymbioticVault.Contract.EpochDurationInit(&_SymbioticVault.CallOpts)
}

// EpochDurationInit is a free data retrieval call binding the contract method 0x46361671.
//
// Solidity: function epochDurationInit() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCallerSession) EpochDurationInit() (*big.Int, error) {
	return _SymbioticVault.Contract.EpochDurationInit(&_SymbioticVault.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticVault *SymbioticVaultSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SymbioticVault.Contract.GetRoleAdmin(&_SymbioticVault.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_SymbioticVault *SymbioticVaultCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _SymbioticVault.Contract.GetRoleAdmin(&_SymbioticVault.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SymbioticVault.Contract.HasRole(&_SymbioticVault.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _SymbioticVault.Contract.HasRole(&_SymbioticVault.CallOpts, role, account)
}

// IsDelegatorInitialized is a free data retrieval call binding the contract method 0x50861adc.
//
// Solidity: function isDelegatorInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) IsDelegatorInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isDelegatorInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelegatorInitialized is a free data retrieval call binding the contract method 0x50861adc.
//
// Solidity: function isDelegatorInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) IsDelegatorInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsDelegatorInitialized(&_SymbioticVault.CallOpts)
}

// IsDelegatorInitialized is a free data retrieval call binding the contract method 0x50861adc.
//
// Solidity: function isDelegatorInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) IsDelegatorInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsDelegatorInitialized(&_SymbioticVault.CallOpts)
}

// IsDepositLimit is a free data retrieval call binding the contract method 0xa1b12202.
//
// Solidity: function isDepositLimit() view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) IsDepositLimit(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isDepositLimit")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositLimit is a free data retrieval call binding the contract method 0xa1b12202.
//
// Solidity: function isDepositLimit() view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) IsDepositLimit() (bool, error) {
	return _SymbioticVault.Contract.IsDepositLimit(&_SymbioticVault.CallOpts)
}

// IsDepositLimit is a free data retrieval call binding the contract method 0xa1b12202.
//
// Solidity: function isDepositLimit() view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) IsDepositLimit() (bool, error) {
	return _SymbioticVault.Contract.IsDepositLimit(&_SymbioticVault.CallOpts)
}

// IsDepositorWhitelisted is a free data retrieval call binding the contract method 0x794b15b7.
//
// Solidity: function isDepositorWhitelisted(address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultCaller) IsDepositorWhitelisted(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isDepositorWhitelisted", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDepositorWhitelisted is a free data retrieval call binding the contract method 0x794b15b7.
//
// Solidity: function isDepositorWhitelisted(address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultSession) IsDepositorWhitelisted(account common.Address) (bool, error) {
	return _SymbioticVault.Contract.IsDepositorWhitelisted(&_SymbioticVault.CallOpts, account)
}

// IsDepositorWhitelisted is a free data retrieval call binding the contract method 0x794b15b7.
//
// Solidity: function isDepositorWhitelisted(address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultCallerSession) IsDepositorWhitelisted(account common.Address) (bool, error) {
	return _SymbioticVault.Contract.IsDepositorWhitelisted(&_SymbioticVault.CallOpts, account)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) IsInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsInitialized(&_SymbioticVault.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) IsInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsInitialized(&_SymbioticVault.CallOpts)
}

// IsSlasherInitialized is a free data retrieval call binding the contract method 0x6ec1e3f8.
//
// Solidity: function isSlasherInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) IsSlasherInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isSlasherInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSlasherInitialized is a free data retrieval call binding the contract method 0x6ec1e3f8.
//
// Solidity: function isSlasherInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) IsSlasherInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsSlasherInitialized(&_SymbioticVault.CallOpts)
}

// IsSlasherInitialized is a free data retrieval call binding the contract method 0x6ec1e3f8.
//
// Solidity: function isSlasherInitialized() view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) IsSlasherInitialized() (bool, error) {
	return _SymbioticVault.Contract.IsSlasherInitialized(&_SymbioticVault.CallOpts)
}

// IsWithdrawalsClaimed is a free data retrieval call binding the contract method 0xa5d03223.
//
// Solidity: function isWithdrawalsClaimed(uint256 epoch, address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultCaller) IsWithdrawalsClaimed(opts *bind.CallOpts, epoch *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "isWithdrawalsClaimed", epoch, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWithdrawalsClaimed is a free data retrieval call binding the contract method 0xa5d03223.
//
// Solidity: function isWithdrawalsClaimed(uint256 epoch, address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultSession) IsWithdrawalsClaimed(epoch *big.Int, account common.Address) (bool, error) {
	return _SymbioticVault.Contract.IsWithdrawalsClaimed(&_SymbioticVault.CallOpts, epoch, account)
}

// IsWithdrawalsClaimed is a free data retrieval call binding the contract method 0xa5d03223.
//
// Solidity: function isWithdrawalsClaimed(uint256 epoch, address account) view returns(bool value)
func (_SymbioticVault *SymbioticVaultCallerSession) IsWithdrawalsClaimed(epoch *big.Int, account common.Address) (bool, error) {
	return _SymbioticVault.Contract.IsWithdrawalsClaimed(&_SymbioticVault.CallOpts, epoch, account)
}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCaller) NextEpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "nextEpochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultSession) NextEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.NextEpochStart(&_SymbioticVault.CallOpts)
}

// NextEpochStart is a free data retrieval call binding the contract method 0x73790ab3.
//
// Solidity: function nextEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCallerSession) NextEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.NextEpochStart(&_SymbioticVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) Owner() (common.Address, error) {
	return _SymbioticVault.Contract.Owner(&_SymbioticVault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) Owner() (common.Address, error) {
	return _SymbioticVault.Contract.Owner(&_SymbioticVault.CallOpts)
}

// PreviousEpochStart is a free data retrieval call binding the contract method 0x281f5752.
//
// Solidity: function previousEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCaller) PreviousEpochStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "previousEpochStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousEpochStart is a free data retrieval call binding the contract method 0x281f5752.
//
// Solidity: function previousEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultSession) PreviousEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.PreviousEpochStart(&_SymbioticVault.CallOpts)
}

// PreviousEpochStart is a free data retrieval call binding the contract method 0x281f5752.
//
// Solidity: function previousEpochStart() view returns(uint48)
func (_SymbioticVault *SymbioticVaultCallerSession) PreviousEpochStart() (*big.Int, error) {
	return _SymbioticVault.Contract.PreviousEpochStart(&_SymbioticVault.CallOpts)
}

// SlashableBalanceOf is a free data retrieval call binding the contract method 0xc31e8dd7.
//
// Solidity: function slashableBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) SlashableBalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "slashableBalanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlashableBalanceOf is a free data retrieval call binding the contract method 0xc31e8dd7.
//
// Solidity: function slashableBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) SlashableBalanceOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.SlashableBalanceOf(&_SymbioticVault.CallOpts, account)
}

// SlashableBalanceOf is a free data retrieval call binding the contract method 0xc31e8dd7.
//
// Solidity: function slashableBalanceOf(address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) SlashableBalanceOf(account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.SlashableBalanceOf(&_SymbioticVault.CallOpts, account)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_SymbioticVault *SymbioticVaultCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_SymbioticVault *SymbioticVaultSession) Slasher() (common.Address, error) {
	return _SymbioticVault.Contract.Slasher(&_SymbioticVault.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_SymbioticVault *SymbioticVaultCallerSession) Slasher() (common.Address, error) {
	return _SymbioticVault.Contract.Slasher(&_SymbioticVault.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticVault *SymbioticVaultCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticVault *SymbioticVaultSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SymbioticVault.Contract.SupportsInterface(&_SymbioticVault.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_SymbioticVault *SymbioticVaultCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _SymbioticVault.Contract.SupportsInterface(&_SymbioticVault.CallOpts, interfaceId)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) TotalStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "totalStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) TotalStake() (*big.Int, error) {
	return _SymbioticVault.Contract.TotalStake(&_SymbioticVault.CallOpts)
}

// TotalStake is a free data retrieval call binding the contract method 0x8b0e9f3f.
//
// Solidity: function totalStake() view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) TotalStake() (*big.Int, error) {
	return _SymbioticVault.Contract.TotalStake(&_SymbioticVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_SymbioticVault *SymbioticVaultCaller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_SymbioticVault *SymbioticVaultSession) Version() (uint64, error) {
	return _SymbioticVault.Contract.Version(&_SymbioticVault.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint64)
func (_SymbioticVault *SymbioticVaultCallerSession) Version() (uint64, error) {
	return _SymbioticVault.Contract.Version(&_SymbioticVault.CallOpts)
}

// WithdrawalShares is a free data retrieval call binding the contract method 0xafba70ad.
//
// Solidity: function withdrawalShares(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCaller) WithdrawalShares(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "withdrawalShares", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalShares is a free data retrieval call binding the contract method 0xafba70ad.
//
// Solidity: function withdrawalShares(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultSession) WithdrawalShares(epoch *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalShares(&_SymbioticVault.CallOpts, epoch)
}

// WithdrawalShares is a free data retrieval call binding the contract method 0xafba70ad.
//
// Solidity: function withdrawalShares(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCallerSession) WithdrawalShares(epoch *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalShares(&_SymbioticVault.CallOpts, epoch)
}

// WithdrawalSharesOf is a free data retrieval call binding the contract method 0xa3b54172.
//
// Solidity: function withdrawalSharesOf(uint256 epoch, address account) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCaller) WithdrawalSharesOf(opts *bind.CallOpts, epoch *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "withdrawalSharesOf", epoch, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalSharesOf is a free data retrieval call binding the contract method 0xa3b54172.
//
// Solidity: function withdrawalSharesOf(uint256 epoch, address account) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultSession) WithdrawalSharesOf(epoch *big.Int, account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalSharesOf(&_SymbioticVault.CallOpts, epoch, account)
}

// WithdrawalSharesOf is a free data retrieval call binding the contract method 0xa3b54172.
//
// Solidity: function withdrawalSharesOf(uint256 epoch, address account) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCallerSession) WithdrawalSharesOf(epoch *big.Int, account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalSharesOf(&_SymbioticVault.CallOpts, epoch, account)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCaller) Withdrawals(opts *bind.CallOpts, epoch *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "withdrawals", epoch)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultSession) Withdrawals(epoch *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.Withdrawals(&_SymbioticVault.CallOpts, epoch)
}

// Withdrawals is a free data retrieval call binding the contract method 0x5cc07076.
//
// Solidity: function withdrawals(uint256 epoch) view returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultCallerSession) Withdrawals(epoch *big.Int) (*big.Int, error) {
	return _SymbioticVault.Contract.Withdrawals(&_SymbioticVault.CallOpts, epoch)
}

// WithdrawalsOf is a free data retrieval call binding the contract method 0xf5e7ee0f.
//
// Solidity: function withdrawalsOf(uint256 epoch, address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCaller) WithdrawalsOf(opts *bind.CallOpts, epoch *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SymbioticVault.contract.Call(opts, &out, "withdrawalsOf", epoch, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalsOf is a free data retrieval call binding the contract method 0xf5e7ee0f.
//
// Solidity: function withdrawalsOf(uint256 epoch, address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultSession) WithdrawalsOf(epoch *big.Int, account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalsOf(&_SymbioticVault.CallOpts, epoch, account)
}

// WithdrawalsOf is a free data retrieval call binding the contract method 0xf5e7ee0f.
//
// Solidity: function withdrawalsOf(uint256 epoch, address account) view returns(uint256)
func (_SymbioticVault *SymbioticVaultCallerSession) WithdrawalsOf(epoch *big.Int, account common.Address) (*big.Int, error) {
	return _SymbioticVault.Contract.WithdrawalsOf(&_SymbioticVault.CallOpts, epoch, account)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 epoch) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultTransactor) Claim(opts *bind.TransactOpts, recipient common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "claim", recipient, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 epoch) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultSession) Claim(recipient common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Claim(&_SymbioticVault.TransactOpts, recipient, epoch)
}

// Claim is a paid mutator transaction binding the contract method 0xaad3ec96.
//
// Solidity: function claim(address recipient, uint256 epoch) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultTransactorSession) Claim(recipient common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Claim(&_SymbioticVault.TransactOpts, recipient, epoch)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0x7c04c80a.
//
// Solidity: function claimBatch(address recipient, uint256[] epochs) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultTransactor) ClaimBatch(opts *bind.TransactOpts, recipient common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "claimBatch", recipient, epochs)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0x7c04c80a.
//
// Solidity: function claimBatch(address recipient, uint256[] epochs) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultSession) ClaimBatch(recipient common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.ClaimBatch(&_SymbioticVault.TransactOpts, recipient, epochs)
}

// ClaimBatch is a paid mutator transaction binding the contract method 0x7c04c80a.
//
// Solidity: function claimBatch(address recipient, uint256[] epochs) returns(uint256 amount)
func (_SymbioticVault *SymbioticVaultTransactorSession) ClaimBatch(recipient common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.ClaimBatch(&_SymbioticVault.TransactOpts, recipient, epochs)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns(uint256 depositedAmount, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactor) Deposit(opts *bind.TransactOpts, onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "deposit", onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns(uint256 depositedAmount, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Deposit(&_SymbioticVault.TransactOpts, onBehalfOf, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address onBehalfOf, uint256 amount) returns(uint256 depositedAmount, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactorSession) Deposit(onBehalfOf common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Deposit(&_SymbioticVault.TransactOpts, onBehalfOf, amount)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.GrantRole(&_SymbioticVault.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.GrantRole(&_SymbioticVault.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x57ec83cc.
//
// Solidity: function initialize(uint64 initialVersion, address owner_, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactor) Initialize(opts *bind.TransactOpts, initialVersion uint64, owner_ common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "initialize", initialVersion, owner_, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x57ec83cc.
//
// Solidity: function initialize(uint64 initialVersion, address owner_, bytes data) returns()
func (_SymbioticVault *SymbioticVaultSession) Initialize(initialVersion uint64, owner_ common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Initialize(&_SymbioticVault.TransactOpts, initialVersion, owner_, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x57ec83cc.
//
// Solidity: function initialize(uint64 initialVersion, address owner_, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) Initialize(initialVersion uint64, owner_ common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Initialize(&_SymbioticVault.TransactOpts, initialVersion, owner_, data)
}

// Migrate is a paid mutator transaction binding the contract method 0x2abe3048.
//
// Solidity: function migrate(uint64 newVersion, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactor) Migrate(opts *bind.TransactOpts, newVersion uint64, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "migrate", newVersion, data)
}

// Migrate is a paid mutator transaction binding the contract method 0x2abe3048.
//
// Solidity: function migrate(uint64 newVersion, bytes data) returns()
func (_SymbioticVault *SymbioticVaultSession) Migrate(newVersion uint64, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Migrate(&_SymbioticVault.TransactOpts, newVersion, data)
}

// Migrate is a paid mutator transaction binding the contract method 0x2abe3048.
//
// Solidity: function migrate(uint64 newVersion, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) Migrate(newVersion uint64, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Migrate(&_SymbioticVault.TransactOpts, newVersion, data)
}

// OnSlash is a paid mutator transaction binding the contract method 0x7278e31c.
//
// Solidity: function onSlash(uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultTransactor) OnSlash(opts *bind.TransactOpts, amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "onSlash", amount, captureTimestamp)
}

// OnSlash is a paid mutator transaction binding the contract method 0x7278e31c.
//
// Solidity: function onSlash(uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultSession) OnSlash(amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.OnSlash(&_SymbioticVault.TransactOpts, amount, captureTimestamp)
}

// OnSlash is a paid mutator transaction binding the contract method 0x7278e31c.
//
// Solidity: function onSlash(uint256 amount, uint48 captureTimestamp) returns(uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultTransactorSession) OnSlash(amount *big.Int, captureTimestamp *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.OnSlash(&_SymbioticVault.TransactOpts, amount, captureTimestamp)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address claimer, uint256 shares) returns(uint256 withdrawnAssets, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactor) Redeem(opts *bind.TransactOpts, claimer common.Address, shares *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "redeem", claimer, shares)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address claimer, uint256 shares) returns(uint256 withdrawnAssets, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultSession) Redeem(claimer common.Address, shares *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Redeem(&_SymbioticVault.TransactOpts, claimer, shares)
}

// Redeem is a paid mutator transaction binding the contract method 0x1e9a6950.
//
// Solidity: function redeem(address claimer, uint256 shares) returns(uint256 withdrawnAssets, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactorSession) Redeem(claimer common.Address, shares *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Redeem(&_SymbioticVault.TransactOpts, claimer, shares)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SymbioticVault *SymbioticVaultTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SymbioticVault *SymbioticVaultSession) RenounceOwnership() (*types.Transaction, error) {
	return _SymbioticVault.Contract.RenounceOwnership(&_SymbioticVault.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SymbioticVault.Contract.RenounceOwnership(&_SymbioticVault.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SymbioticVault *SymbioticVaultTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SymbioticVault *SymbioticVaultSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.RenounceRole(&_SymbioticVault.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.RenounceRole(&_SymbioticVault.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.RevokeRole(&_SymbioticVault.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.RevokeRole(&_SymbioticVault.TransactOpts, role, account)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator_) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetDelegator(opts *bind.TransactOpts, delegator_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setDelegator", delegator_)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator_) returns()
func (_SymbioticVault *SymbioticVaultSession) SetDelegator(delegator_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDelegator(&_SymbioticVault.TransactOpts, delegator_)
}

// SetDelegator is a paid mutator transaction binding the contract method 0x83cd9cc3.
//
// Solidity: function setDelegator(address delegator_) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetDelegator(delegator_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDelegator(&_SymbioticVault.TransactOpts, delegator_)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetDepositLimit(opts *bind.TransactOpts, limit *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setDepositLimit", limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_SymbioticVault *SymbioticVaultSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositLimit(&_SymbioticVault.TransactOpts, limit)
}

// SetDepositLimit is a paid mutator transaction binding the contract method 0xbdc8144b.
//
// Solidity: function setDepositLimit(uint256 limit) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetDepositLimit(limit *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositLimit(&_SymbioticVault.TransactOpts, limit)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x4105a7dd.
//
// Solidity: function setDepositWhitelist(bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetDepositWhitelist(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setDepositWhitelist", status)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x4105a7dd.
//
// Solidity: function setDepositWhitelist(bool status) returns()
func (_SymbioticVault *SymbioticVaultSession) SetDepositWhitelist(status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositWhitelist(&_SymbioticVault.TransactOpts, status)
}

// SetDepositWhitelist is a paid mutator transaction binding the contract method 0x4105a7dd.
//
// Solidity: function setDepositWhitelist(bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetDepositWhitelist(status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositWhitelist(&_SymbioticVault.TransactOpts, status)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xa2861466.
//
// Solidity: function setDepositorWhitelistStatus(address account, bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetDepositorWhitelistStatus(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setDepositorWhitelistStatus", account, status)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xa2861466.
//
// Solidity: function setDepositorWhitelistStatus(address account, bool status) returns()
func (_SymbioticVault *SymbioticVaultSession) SetDepositorWhitelistStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositorWhitelistStatus(&_SymbioticVault.TransactOpts, account, status)
}

// SetDepositorWhitelistStatus is a paid mutator transaction binding the contract method 0xa2861466.
//
// Solidity: function setDepositorWhitelistStatus(address account, bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetDepositorWhitelistStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetDepositorWhitelistStatus(&_SymbioticVault.TransactOpts, account, status)
}

// SetIsDepositLimit is a paid mutator transaction binding the contract method 0x5346e34f.
//
// Solidity: function setIsDepositLimit(bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetIsDepositLimit(opts *bind.TransactOpts, status bool) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setIsDepositLimit", status)
}

// SetIsDepositLimit is a paid mutator transaction binding the contract method 0x5346e34f.
//
// Solidity: function setIsDepositLimit(bool status) returns()
func (_SymbioticVault *SymbioticVaultSession) SetIsDepositLimit(status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetIsDepositLimit(&_SymbioticVault.TransactOpts, status)
}

// SetIsDepositLimit is a paid mutator transaction binding the contract method 0x5346e34f.
//
// Solidity: function setIsDepositLimit(bool status) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetIsDepositLimit(status bool) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetIsDepositLimit(&_SymbioticVault.TransactOpts, status)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasher_) returns()
func (_SymbioticVault *SymbioticVaultTransactor) SetSlasher(opts *bind.TransactOpts, slasher_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "setSlasher", slasher_)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasher_) returns()
func (_SymbioticVault *SymbioticVaultSession) SetSlasher(slasher_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetSlasher(&_SymbioticVault.TransactOpts, slasher_)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address slasher_) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) SetSlasher(slasher_ common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.SetSlasher(&_SymbioticVault.TransactOpts, slasher_)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactor) StaticDelegateCall(opts *bind.TransactOpts, target common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "staticDelegateCall", target, data)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_SymbioticVault *SymbioticVaultSession) StaticDelegateCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.StaticDelegateCall(&_SymbioticVault.TransactOpts, target, data)
}

// StaticDelegateCall is a paid mutator transaction binding the contract method 0x9f86fd85.
//
// Solidity: function staticDelegateCall(address target, bytes data) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) StaticDelegateCall(target common.Address, data []byte) (*types.Transaction, error) {
	return _SymbioticVault.Contract.StaticDelegateCall(&_SymbioticVault.TransactOpts, target, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SymbioticVault *SymbioticVaultTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SymbioticVault *SymbioticVaultSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.TransferOwnership(&_SymbioticVault.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SymbioticVault *SymbioticVaultTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SymbioticVault.Contract.TransferOwnership(&_SymbioticVault.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address claimer, uint256 amount) returns(uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactor) Withdraw(opts *bind.TransactOpts, claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.contract.Transact(opts, "withdraw", claimer, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address claimer, uint256 amount) returns(uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultSession) Withdraw(claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Withdraw(&_SymbioticVault.TransactOpts, claimer, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address claimer, uint256 amount) returns(uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultTransactorSession) Withdraw(claimer common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SymbioticVault.Contract.Withdraw(&_SymbioticVault.TransactOpts, claimer, amount)
}

// SymbioticVaultClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the SymbioticVault contract.
type SymbioticVaultClaimIterator struct {
	Event *SymbioticVaultClaim // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultClaim)
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
		it.Event = new(SymbioticVaultClaim)
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
func (it *SymbioticVaultClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultClaim represents a Claim event raised by the SymbioticVault contract.
type SymbioticVaultClaim struct {
	Claimer   common.Address
	Recipient common.Address
	Epoch     *big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed claimer, address indexed recipient, uint256 epoch, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) FilterClaim(opts *bind.FilterOpts, claimer []common.Address, recipient []common.Address) (*SymbioticVaultClaimIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "Claim", claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultClaimIterator{contract: _SymbioticVault.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed claimer, address indexed recipient, uint256 epoch, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *SymbioticVaultClaim, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "Claim", claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultClaim)
				if err := _SymbioticVault.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x865ca08d59f5cb456e85cd2f7ef63664ea4f73327414e9d8152c4158b0e94645.
//
// Solidity: event Claim(address indexed claimer, address indexed recipient, uint256 epoch, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) ParseClaim(log types.Log) (*SymbioticVaultClaim, error) {
	event := new(SymbioticVaultClaim)
	if err := _SymbioticVault.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultClaimBatchIterator is returned from FilterClaimBatch and is used to iterate over the raw logs and unpacked data for ClaimBatch events raised by the SymbioticVault contract.
type SymbioticVaultClaimBatchIterator struct {
	Event *SymbioticVaultClaimBatch // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultClaimBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultClaimBatch)
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
		it.Event = new(SymbioticVaultClaimBatch)
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
func (it *SymbioticVaultClaimBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultClaimBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultClaimBatch represents a ClaimBatch event raised by the SymbioticVault contract.
type SymbioticVaultClaimBatch struct {
	Claimer   common.Address
	Recipient common.Address
	Epochs    []*big.Int
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterClaimBatch is a free log retrieval operation binding the contract event 0x326b6aff1cd2fb1c8234de4f9dcfb9047c5c36eb9ef2eb34af5121e969a75d27.
//
// Solidity: event ClaimBatch(address indexed claimer, address indexed recipient, uint256[] epochs, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) FilterClaimBatch(opts *bind.FilterOpts, claimer []common.Address, recipient []common.Address) (*SymbioticVaultClaimBatchIterator, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "ClaimBatch", claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultClaimBatchIterator{contract: _SymbioticVault.contract, event: "ClaimBatch", logs: logs, sub: sub}, nil
}

// WatchClaimBatch is a free log subscription operation binding the contract event 0x326b6aff1cd2fb1c8234de4f9dcfb9047c5c36eb9ef2eb34af5121e969a75d27.
//
// Solidity: event ClaimBatch(address indexed claimer, address indexed recipient, uint256[] epochs, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) WatchClaimBatch(opts *bind.WatchOpts, sink chan<- *SymbioticVaultClaimBatch, claimer []common.Address, recipient []common.Address) (event.Subscription, error) {

	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "ClaimBatch", claimerRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultClaimBatch)
				if err := _SymbioticVault.contract.UnpackLog(event, "ClaimBatch", log); err != nil {
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

// ParseClaimBatch is a log parse operation binding the contract event 0x326b6aff1cd2fb1c8234de4f9dcfb9047c5c36eb9ef2eb34af5121e969a75d27.
//
// Solidity: event ClaimBatch(address indexed claimer, address indexed recipient, uint256[] epochs, uint256 amount)
func (_SymbioticVault *SymbioticVaultFilterer) ParseClaimBatch(log types.Log) (*SymbioticVaultClaimBatch, error) {
	event := new(SymbioticVaultClaimBatch)
	if err := _SymbioticVault.contract.UnpackLog(event, "ClaimBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the SymbioticVault contract.
type SymbioticVaultDepositIterator struct {
	Event *SymbioticVaultDeposit // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultDeposit)
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
		it.Event = new(SymbioticVaultDeposit)
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
func (it *SymbioticVaultDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultDeposit represents a Deposit event raised by the SymbioticVault contract.
type SymbioticVaultDeposit struct {
	Depositor  common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Shares     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed depositor, address indexed onBehalfOf, uint256 amount, uint256 shares)
func (_SymbioticVault *SymbioticVaultFilterer) FilterDeposit(opts *bind.FilterOpts, depositor []common.Address, onBehalfOf []common.Address) (*SymbioticVaultDepositIterator, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "Deposit", depositorRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultDepositIterator{contract: _SymbioticVault.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed depositor, address indexed onBehalfOf, uint256 amount, uint256 shares)
func (_SymbioticVault *SymbioticVaultFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *SymbioticVaultDeposit, depositor []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var depositorRule []interface{}
	for _, depositorItem := range depositor {
		depositorRule = append(depositorRule, depositorItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "Deposit", depositorRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultDeposit)
				if err := _SymbioticVault.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xdcbc1c05240f31ff3ad067ef1ee35ce4997762752e3a095284754544f4c709d7.
//
// Solidity: event Deposit(address indexed depositor, address indexed onBehalfOf, uint256 amount, uint256 shares)
func (_SymbioticVault *SymbioticVaultFilterer) ParseDeposit(log types.Log) (*SymbioticVaultDeposit, error) {
	event := new(SymbioticVaultDeposit)
	if err := _SymbioticVault.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SymbioticVault contract.
type SymbioticVaultInitializedIterator struct {
	Event *SymbioticVaultInitialized // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultInitialized)
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
		it.Event = new(SymbioticVaultInitialized)
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
func (it *SymbioticVaultInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultInitialized represents a Initialized event raised by the SymbioticVault contract.
type SymbioticVaultInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SymbioticVault *SymbioticVaultFilterer) FilterInitialized(opts *bind.FilterOpts) (*SymbioticVaultInitializedIterator, error) {

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultInitializedIterator{contract: _SymbioticVault.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SymbioticVault *SymbioticVaultFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SymbioticVaultInitialized) (event.Subscription, error) {

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultInitialized)
				if err := _SymbioticVault.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SymbioticVault *SymbioticVaultFilterer) ParseInitialized(log types.Log) (*SymbioticVaultInitialized, error) {
	event := new(SymbioticVaultInitialized)
	if err := _SymbioticVault.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultOnSlashIterator is returned from FilterOnSlash and is used to iterate over the raw logs and unpacked data for OnSlash events raised by the SymbioticVault contract.
type SymbioticVaultOnSlashIterator struct {
	Event *SymbioticVaultOnSlash // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultOnSlashIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultOnSlash)
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
		it.Event = new(SymbioticVaultOnSlash)
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
func (it *SymbioticVaultOnSlashIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultOnSlashIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultOnSlash represents a OnSlash event raised by the SymbioticVault contract.
type SymbioticVaultOnSlash struct {
	Amount           *big.Int
	CaptureTimestamp *big.Int
	SlashedAmount    *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOnSlash is a free log retrieval operation binding the contract event 0xf9d090c096f71cd1659861a9ce5b6f384bceb4e2fa4e4a19edf6489a9b8d56c7.
//
// Solidity: event OnSlash(uint256 amount, uint48 captureTimestamp, uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultFilterer) FilterOnSlash(opts *bind.FilterOpts) (*SymbioticVaultOnSlashIterator, error) {

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "OnSlash")
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultOnSlashIterator{contract: _SymbioticVault.contract, event: "OnSlash", logs: logs, sub: sub}, nil
}

// WatchOnSlash is a free log subscription operation binding the contract event 0xf9d090c096f71cd1659861a9ce5b6f384bceb4e2fa4e4a19edf6489a9b8d56c7.
//
// Solidity: event OnSlash(uint256 amount, uint48 captureTimestamp, uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultFilterer) WatchOnSlash(opts *bind.WatchOpts, sink chan<- *SymbioticVaultOnSlash) (event.Subscription, error) {

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "OnSlash")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultOnSlash)
				if err := _SymbioticVault.contract.UnpackLog(event, "OnSlash", log); err != nil {
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

// ParseOnSlash is a log parse operation binding the contract event 0xf9d090c096f71cd1659861a9ce5b6f384bceb4e2fa4e4a19edf6489a9b8d56c7.
//
// Solidity: event OnSlash(uint256 amount, uint48 captureTimestamp, uint256 slashedAmount)
func (_SymbioticVault *SymbioticVaultFilterer) ParseOnSlash(log types.Log) (*SymbioticVaultOnSlash, error) {
	event := new(SymbioticVaultOnSlash)
	if err := _SymbioticVault.contract.UnpackLog(event, "OnSlash", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SymbioticVault contract.
type SymbioticVaultOwnershipTransferredIterator struct {
	Event *SymbioticVaultOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultOwnershipTransferred)
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
		it.Event = new(SymbioticVaultOwnershipTransferred)
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
func (it *SymbioticVaultOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultOwnershipTransferred represents a OwnershipTransferred event raised by the SymbioticVault contract.
type SymbioticVaultOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SymbioticVault *SymbioticVaultFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SymbioticVaultOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultOwnershipTransferredIterator{contract: _SymbioticVault.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SymbioticVault *SymbioticVaultFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SymbioticVaultOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultOwnershipTransferred)
				if err := _SymbioticVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SymbioticVault *SymbioticVaultFilterer) ParseOwnershipTransferred(log types.Log) (*SymbioticVaultOwnershipTransferred, error) {
	event := new(SymbioticVaultOwnershipTransferred)
	if err := _SymbioticVault.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the SymbioticVault contract.
type SymbioticVaultRoleAdminChangedIterator struct {
	Event *SymbioticVaultRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultRoleAdminChanged)
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
		it.Event = new(SymbioticVaultRoleAdminChanged)
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
func (it *SymbioticVaultRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultRoleAdminChanged represents a RoleAdminChanged event raised by the SymbioticVault contract.
type SymbioticVaultRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SymbioticVault *SymbioticVaultFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*SymbioticVaultRoleAdminChangedIterator, error) {

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

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultRoleAdminChangedIterator{contract: _SymbioticVault.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_SymbioticVault *SymbioticVaultFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *SymbioticVaultRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultRoleAdminChanged)
				if err := _SymbioticVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_SymbioticVault *SymbioticVaultFilterer) ParseRoleAdminChanged(log types.Log) (*SymbioticVaultRoleAdminChanged, error) {
	event := new(SymbioticVaultRoleAdminChanged)
	if err := _SymbioticVault.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the SymbioticVault contract.
type SymbioticVaultRoleGrantedIterator struct {
	Event *SymbioticVaultRoleGranted // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultRoleGranted)
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
		it.Event = new(SymbioticVaultRoleGranted)
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
func (it *SymbioticVaultRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultRoleGranted represents a RoleGranted event raised by the SymbioticVault contract.
type SymbioticVaultRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticVault *SymbioticVaultFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SymbioticVaultRoleGrantedIterator, error) {

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

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultRoleGrantedIterator{contract: _SymbioticVault.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticVault *SymbioticVaultFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *SymbioticVaultRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultRoleGranted)
				if err := _SymbioticVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_SymbioticVault *SymbioticVaultFilterer) ParseRoleGranted(log types.Log) (*SymbioticVaultRoleGranted, error) {
	event := new(SymbioticVaultRoleGranted)
	if err := _SymbioticVault.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the SymbioticVault contract.
type SymbioticVaultRoleRevokedIterator struct {
	Event *SymbioticVaultRoleRevoked // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultRoleRevoked)
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
		it.Event = new(SymbioticVaultRoleRevoked)
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
func (it *SymbioticVaultRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultRoleRevoked represents a RoleRevoked event raised by the SymbioticVault contract.
type SymbioticVaultRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticVault *SymbioticVaultFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*SymbioticVaultRoleRevokedIterator, error) {

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

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultRoleRevokedIterator{contract: _SymbioticVault.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_SymbioticVault *SymbioticVaultFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *SymbioticVaultRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultRoleRevoked)
				if err := _SymbioticVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_SymbioticVault *SymbioticVaultFilterer) ParseRoleRevoked(log types.Log) (*SymbioticVaultRoleRevoked, error) {
	event := new(SymbioticVaultRoleRevoked)
	if err := _SymbioticVault.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetDelegatorIterator is returned from FilterSetDelegator and is used to iterate over the raw logs and unpacked data for SetDelegator events raised by the SymbioticVault contract.
type SymbioticVaultSetDelegatorIterator struct {
	Event *SymbioticVaultSetDelegator // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetDelegatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetDelegator)
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
		it.Event = new(SymbioticVaultSetDelegator)
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
func (it *SymbioticVaultSetDelegatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetDelegatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetDelegator represents a SetDelegator event raised by the SymbioticVault contract.
type SymbioticVaultSetDelegator struct {
	Delegator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSetDelegator is a free log retrieval operation binding the contract event 0xdb2160616f776a37b24808115554e79439bf26cccbbd4438190cc6d28e80ecd1.
//
// Solidity: event SetDelegator(address indexed delegator)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetDelegator(opts *bind.FilterOpts, delegator []common.Address) (*SymbioticVaultSetDelegatorIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetDelegator", delegatorRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetDelegatorIterator{contract: _SymbioticVault.contract, event: "SetDelegator", logs: logs, sub: sub}, nil
}

// WatchSetDelegator is a free log subscription operation binding the contract event 0xdb2160616f776a37b24808115554e79439bf26cccbbd4438190cc6d28e80ecd1.
//
// Solidity: event SetDelegator(address indexed delegator)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetDelegator(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetDelegator, delegator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetDelegator", delegatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetDelegator)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetDelegator", log); err != nil {
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

// ParseSetDelegator is a log parse operation binding the contract event 0xdb2160616f776a37b24808115554e79439bf26cccbbd4438190cc6d28e80ecd1.
//
// Solidity: event SetDelegator(address indexed delegator)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetDelegator(log types.Log) (*SymbioticVaultSetDelegator, error) {
	event := new(SymbioticVaultSetDelegator)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetDelegator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetDepositLimitIterator is returned from FilterSetDepositLimit and is used to iterate over the raw logs and unpacked data for SetDepositLimit events raised by the SymbioticVault contract.
type SymbioticVaultSetDepositLimitIterator struct {
	Event *SymbioticVaultSetDepositLimit // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetDepositLimit)
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
		it.Event = new(SymbioticVaultSetDepositLimit)
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
func (it *SymbioticVaultSetDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetDepositLimit represents a SetDepositLimit event raised by the SymbioticVault contract.
type SymbioticVaultSetDepositLimit struct {
	Limit *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSetDepositLimit is a free log retrieval operation binding the contract event 0x854df3eb95564502c8bc871ebdd15310ee26270f955f6c6bd8cea68e75045bc0.
//
// Solidity: event SetDepositLimit(uint256 limit)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetDepositLimit(opts *bind.FilterOpts) (*SymbioticVaultSetDepositLimitIterator, error) {

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetDepositLimit")
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetDepositLimitIterator{contract: _SymbioticVault.contract, event: "SetDepositLimit", logs: logs, sub: sub}, nil
}

// WatchSetDepositLimit is a free log subscription operation binding the contract event 0x854df3eb95564502c8bc871ebdd15310ee26270f955f6c6bd8cea68e75045bc0.
//
// Solidity: event SetDepositLimit(uint256 limit)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetDepositLimit(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetDepositLimit) (event.Subscription, error) {

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetDepositLimit)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositLimit", log); err != nil {
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

// ParseSetDepositLimit is a log parse operation binding the contract event 0x854df3eb95564502c8bc871ebdd15310ee26270f955f6c6bd8cea68e75045bc0.
//
// Solidity: event SetDepositLimit(uint256 limit)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetDepositLimit(log types.Log) (*SymbioticVaultSetDepositLimit, error) {
	event := new(SymbioticVaultSetDepositLimit)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetDepositWhitelistIterator is returned from FilterSetDepositWhitelist and is used to iterate over the raw logs and unpacked data for SetDepositWhitelist events raised by the SymbioticVault contract.
type SymbioticVaultSetDepositWhitelistIterator struct {
	Event *SymbioticVaultSetDepositWhitelist // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetDepositWhitelistIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetDepositWhitelist)
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
		it.Event = new(SymbioticVaultSetDepositWhitelist)
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
func (it *SymbioticVaultSetDepositWhitelistIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetDepositWhitelistIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetDepositWhitelist represents a SetDepositWhitelist event raised by the SymbioticVault contract.
type SymbioticVaultSetDepositWhitelist struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetDepositWhitelist is a free log retrieval operation binding the contract event 0x3e12b7b36c75ac9609a3f58609b331210428e1a85909132638955ba0301eec33.
//
// Solidity: event SetDepositWhitelist(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetDepositWhitelist(opts *bind.FilterOpts) (*SymbioticVaultSetDepositWhitelistIterator, error) {

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetDepositWhitelistIterator{contract: _SymbioticVault.contract, event: "SetDepositWhitelist", logs: logs, sub: sub}, nil
}

// WatchSetDepositWhitelist is a free log subscription operation binding the contract event 0x3e12b7b36c75ac9609a3f58609b331210428e1a85909132638955ba0301eec33.
//
// Solidity: event SetDepositWhitelist(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetDepositWhitelist(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetDepositWhitelist) (event.Subscription, error) {

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetDepositWhitelist")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetDepositWhitelist)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositWhitelist", log); err != nil {
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

// ParseSetDepositWhitelist is a log parse operation binding the contract event 0x3e12b7b36c75ac9609a3f58609b331210428e1a85909132638955ba0301eec33.
//
// Solidity: event SetDepositWhitelist(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetDepositWhitelist(log types.Log) (*SymbioticVaultSetDepositWhitelist, error) {
	event := new(SymbioticVaultSetDepositWhitelist)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositWhitelist", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetDepositorWhitelistStatusIterator is returned from FilterSetDepositorWhitelistStatus and is used to iterate over the raw logs and unpacked data for SetDepositorWhitelistStatus events raised by the SymbioticVault contract.
type SymbioticVaultSetDepositorWhitelistStatusIterator struct {
	Event *SymbioticVaultSetDepositorWhitelistStatus // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetDepositorWhitelistStatusIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetDepositorWhitelistStatus)
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
		it.Event = new(SymbioticVaultSetDepositorWhitelistStatus)
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
func (it *SymbioticVaultSetDepositorWhitelistStatusIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetDepositorWhitelistStatusIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetDepositorWhitelistStatus represents a SetDepositorWhitelistStatus event raised by the SymbioticVault contract.
type SymbioticVaultSetDepositorWhitelistStatus struct {
	Account common.Address
	Status  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetDepositorWhitelistStatus is a free log retrieval operation binding the contract event 0xf991b1ecfb5115cbb36a2b2e2240c058406d2acc2fcc6e9e2dc99d845ff70a62.
//
// Solidity: event SetDepositorWhitelistStatus(address indexed account, bool status)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetDepositorWhitelistStatus(opts *bind.FilterOpts, account []common.Address) (*SymbioticVaultSetDepositorWhitelistStatusIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetDepositorWhitelistStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetDepositorWhitelistStatusIterator{contract: _SymbioticVault.contract, event: "SetDepositorWhitelistStatus", logs: logs, sub: sub}, nil
}

// WatchSetDepositorWhitelistStatus is a free log subscription operation binding the contract event 0xf991b1ecfb5115cbb36a2b2e2240c058406d2acc2fcc6e9e2dc99d845ff70a62.
//
// Solidity: event SetDepositorWhitelistStatus(address indexed account, bool status)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetDepositorWhitelistStatus(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetDepositorWhitelistStatus, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetDepositorWhitelistStatus", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetDepositorWhitelistStatus)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositorWhitelistStatus", log); err != nil {
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

// ParseSetDepositorWhitelistStatus is a log parse operation binding the contract event 0xf991b1ecfb5115cbb36a2b2e2240c058406d2acc2fcc6e9e2dc99d845ff70a62.
//
// Solidity: event SetDepositorWhitelistStatus(address indexed account, bool status)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetDepositorWhitelistStatus(log types.Log) (*SymbioticVaultSetDepositorWhitelistStatus, error) {
	event := new(SymbioticVaultSetDepositorWhitelistStatus)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetDepositorWhitelistStatus", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetIsDepositLimitIterator is returned from FilterSetIsDepositLimit and is used to iterate over the raw logs and unpacked data for SetIsDepositLimit events raised by the SymbioticVault contract.
type SymbioticVaultSetIsDepositLimitIterator struct {
	Event *SymbioticVaultSetIsDepositLimit // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetIsDepositLimitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetIsDepositLimit)
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
		it.Event = new(SymbioticVaultSetIsDepositLimit)
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
func (it *SymbioticVaultSetIsDepositLimitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetIsDepositLimitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetIsDepositLimit represents a SetIsDepositLimit event raised by the SymbioticVault contract.
type SymbioticVaultSetIsDepositLimit struct {
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetIsDepositLimit is a free log retrieval operation binding the contract event 0xfa7a25a0b611d4ba3c0ea990e90dc23d484a5dd7a1be4733fef2946ba74530c6.
//
// Solidity: event SetIsDepositLimit(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetIsDepositLimit(opts *bind.FilterOpts) (*SymbioticVaultSetIsDepositLimitIterator, error) {

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetIsDepositLimit")
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetIsDepositLimitIterator{contract: _SymbioticVault.contract, event: "SetIsDepositLimit", logs: logs, sub: sub}, nil
}

// WatchSetIsDepositLimit is a free log subscription operation binding the contract event 0xfa7a25a0b611d4ba3c0ea990e90dc23d484a5dd7a1be4733fef2946ba74530c6.
//
// Solidity: event SetIsDepositLimit(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetIsDepositLimit(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetIsDepositLimit) (event.Subscription, error) {

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetIsDepositLimit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetIsDepositLimit)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetIsDepositLimit", log); err != nil {
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

// ParseSetIsDepositLimit is a log parse operation binding the contract event 0xfa7a25a0b611d4ba3c0ea990e90dc23d484a5dd7a1be4733fef2946ba74530c6.
//
// Solidity: event SetIsDepositLimit(bool status)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetIsDepositLimit(log types.Log) (*SymbioticVaultSetIsDepositLimit, error) {
	event := new(SymbioticVaultSetIsDepositLimit)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetIsDepositLimit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultSetSlasherIterator is returned from FilterSetSlasher and is used to iterate over the raw logs and unpacked data for SetSlasher events raised by the SymbioticVault contract.
type SymbioticVaultSetSlasherIterator struct {
	Event *SymbioticVaultSetSlasher // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultSetSlasherIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultSetSlasher)
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
		it.Event = new(SymbioticVaultSetSlasher)
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
func (it *SymbioticVaultSetSlasherIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultSetSlasherIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultSetSlasher represents a SetSlasher event raised by the SymbioticVault contract.
type SymbioticVaultSetSlasher struct {
	Slasher common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSetSlasher is a free log retrieval operation binding the contract event 0xe7e4c932e03abddfe20f83af42c33627e816115c7ec2b168441f65dc14bfc3ba.
//
// Solidity: event SetSlasher(address indexed slasher)
func (_SymbioticVault *SymbioticVaultFilterer) FilterSetSlasher(opts *bind.FilterOpts, slasher []common.Address) (*SymbioticVaultSetSlasherIterator, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "SetSlasher", slasherRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultSetSlasherIterator{contract: _SymbioticVault.contract, event: "SetSlasher", logs: logs, sub: sub}, nil
}

// WatchSetSlasher is a free log subscription operation binding the contract event 0xe7e4c932e03abddfe20f83af42c33627e816115c7ec2b168441f65dc14bfc3ba.
//
// Solidity: event SetSlasher(address indexed slasher)
func (_SymbioticVault *SymbioticVaultFilterer) WatchSetSlasher(opts *bind.WatchOpts, sink chan<- *SymbioticVaultSetSlasher, slasher []common.Address) (event.Subscription, error) {

	var slasherRule []interface{}
	for _, slasherItem := range slasher {
		slasherRule = append(slasherRule, slasherItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "SetSlasher", slasherRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultSetSlasher)
				if err := _SymbioticVault.contract.UnpackLog(event, "SetSlasher", log); err != nil {
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

// ParseSetSlasher is a log parse operation binding the contract event 0xe7e4c932e03abddfe20f83af42c33627e816115c7ec2b168441f65dc14bfc3ba.
//
// Solidity: event SetSlasher(address indexed slasher)
func (_SymbioticVault *SymbioticVaultFilterer) ParseSetSlasher(log types.Log) (*SymbioticVaultSetSlasher, error) {
	event := new(SymbioticVaultSetSlasher)
	if err := _SymbioticVault.contract.UnpackLog(event, "SetSlasher", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SymbioticVaultWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SymbioticVault contract.
type SymbioticVaultWithdrawIterator struct {
	Event *SymbioticVaultWithdraw // Event containing the contract specifics and raw log

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
func (it *SymbioticVaultWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SymbioticVaultWithdraw)
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
		it.Event = new(SymbioticVaultWithdraw)
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
func (it *SymbioticVaultWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SymbioticVaultWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SymbioticVaultWithdraw represents a Withdraw event raised by the SymbioticVault contract.
type SymbioticVaultWithdraw struct {
	Withdrawer   common.Address
	Claimer      common.Address
	Amount       *big.Int
	BurnedShares *big.Int
	MintedShares *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed withdrawer, address indexed claimer, uint256 amount, uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultFilterer) FilterWithdraw(opts *bind.FilterOpts, withdrawer []common.Address, claimer []common.Address) (*SymbioticVaultWithdrawIterator, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SymbioticVault.contract.FilterLogs(opts, "Withdraw", withdrawerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &SymbioticVaultWithdrawIterator{contract: _SymbioticVault.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed withdrawer, address indexed claimer, uint256 amount, uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SymbioticVaultWithdraw, withdrawer []common.Address, claimer []common.Address) (event.Subscription, error) {

	var withdrawerRule []interface{}
	for _, withdrawerItem := range withdrawer {
		withdrawerRule = append(withdrawerRule, withdrawerItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _SymbioticVault.contract.WatchLogs(opts, "Withdraw", withdrawerRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SymbioticVaultWithdraw)
				if err := _SymbioticVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0xebff2602b3f468259e1e99f613fed6691f3a6526effe6ef3e768ba7ae7a36c4f.
//
// Solidity: event Withdraw(address indexed withdrawer, address indexed claimer, uint256 amount, uint256 burnedShares, uint256 mintedShares)
func (_SymbioticVault *SymbioticVaultFilterer) ParseWithdraw(log types.Log) (*SymbioticVaultWithdraw, error) {
	event := new(SymbioticVaultWithdraw)
	if err := _SymbioticVault.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
