package store

import (
	"encoding/binary"
)

var (
	BabylonBlockHeaderKeyPrefix         = []byte{0x01}
	EthBlockHeaderKeyPrefix             = []byte{0x02}
	TxMessageKeyPrefix                  = []byte{0x03}
	EthScannedHeightKeyPrefix           = []byte{0x04}
	BabylonScannedHeightKeyPrefix       = []byte{0x05}
	NewFinalityProviderKeyPrefix        = []byte{0x06}
	CreateBTCDelegationKeyPrefix        = []byte{0x07}
	CommitPubRandListKeyPrefix          = []byte{0x08}
	SignatureKeyPrefix                  = []byte{0x09}
	ContractEventKeyPrefix              = []byte{0x10}
	ActiveMemberKeyPrefix               = []byte{0x11}
	BtcUndelegateKeyPrefix              = []byte{0x12}
	BTCDelegateAmountKeyPrefix          = []byte{0x13}
	SelectiveSlashingEvidenceKeyPrefix  = []byte{0x14}
	BabylonDelegationKeyPrefix          = []byte{0x15}
	SubmitFinalitySignatureKeyMsgPrefix = []byte{0x16}
	StakeDetailsKeyPrefix               = []byte{0x17}
	BatchStakeDetailsKeyPrefix          = []byte{0x18}
	SymbioticFpBlobKeyPrefix            = []byte{0x19}
	OutputProposedKeyPrefix             = []byte{0x20}
	SubmitFinalitySignatureKeyPrefix    = []byte{0x21}
	CelestiaScannedHeightKeyPrefix      = []byte{0x22}
	CelestiaBlockHeaderKeyPrefix        = []byte{0x23}
	StakeDetailsByTimestampKeyPrefix    = []byte{0x24}
	FinalityOutputProposedKeyPrefix     = []byte{0x25}
	BatchStakeDetailsByApiKeyPrefix     = []byte{0x26}
)

func getBabylonBlockHeaderKey(number int64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, uint64(number))
	return append(BabylonBlockHeaderKeyPrefix, numberBz...)
}

func getEthBlockHeaderKey(number int64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, uint64(number))
	return append(EthBlockHeaderKeyPrefix, numberBz...)
}

func getSignatureKey(number int64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, uint64(number))
	return append(SignatureKeyPrefix, numberBz...)
}

func getTxMessageKey(txHash []byte) []byte {
	return append(TxMessageKeyPrefix, txHash[:]...)
}

func getContractEventKey(txHash []byte) []byte {
	return append(ContractEventKeyPrefix, txHash[:]...)
}

func getNewFinalityProviderKey(txHash []byte) []byte {
	return append(NewFinalityProviderKeyPrefix, txHash[:]...)
}

func getCreateBTCDelegationKey(txHash []byte) []byte {
	return append(CreateBTCDelegationKeyPrefix, txHash[:]...)
}

func getCommitPubRandListKey(txHash []byte) []byte {
	return append(CommitPubRandListKeyPrefix, txHash[:]...)
}

func getBtcUndelegateKey(txHash []byte) []byte {
	return append(BtcUndelegateKeyPrefix, txHash[:]...)
}

func getSelectiveSlashingEvidenceKey(txHash []byte) []byte {
	return append(SelectiveSlashingEvidenceKeyPrefix, txHash[:]...)
}

func getBTCDelegateAmountKey() []byte {
	return append(BTCDelegateAmountKeyPrefix)
}

func getBabylonDelegationKey(btcTx []byte) []byte {
	return append(BabylonDelegationKeyPrefix, btcTx[:]...)
}

func getEthScannedHeightKey() []byte {
	return EthScannedHeightKeyPrefix
}

func getBabylonScannedHeightKey() []byte {
	return BabylonScannedHeightKeyPrefix
}

func getActiveMemberKey() []byte {
	return ActiveMemberKeyPrefix
}

func getSubmitFinalitySignatureMsgKey(txHash []byte) []byte {
	return append(SubmitFinalitySignatureKeyMsgPrefix, txHash[:]...)
}

func getSubmitFinalitySignatureKey(timestamp uint64) []byte {
	timestampBz := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, timestamp)
	return append(SubmitFinalitySignatureKeyPrefix, timestampBz...)
}

func getStakeDetailsKey() []byte {
	return StakeDetailsKeyPrefix
}

func getBatchStakeDetailsKey(batchId uint64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, batchId)
	return append(BatchStakeDetailsKeyPrefix, numberBz...)
}

func getStakeDetailsApiKey(batchId uint64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, batchId)
	return append(BatchStakeDetailsByApiKeyPrefix, numberBz...)
}

func getSymbioticFpBlobKey(timestamp uint64) []byte {
	timestampBz := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, timestamp)
	return append(SymbioticFpBlobKeyPrefix, timestampBz...)
}

func getOutputProposedKey(timestamp uint64) []byte {
	timestampBz := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, timestamp)
	return append(OutputProposedKeyPrefix, timestampBz...)
}

func getCelestiaScannedHeightKey() []byte {
	return CelestiaScannedHeightKeyPrefix
}

func getCelestiaBlockHeaderKey(number uint64) []byte {
	numberBz := make([]byte, 8)
	binary.BigEndian.PutUint64(numberBz, number)
	return append(CelestiaBlockHeaderKeyPrefix, numberBz...)
}

func getStakeDetailsByTimestampKey(timestamp uint64) []byte {
	timestampBz := make([]byte, 8)
	binary.BigEndian.PutUint64(timestampBz, timestamp)
	return append(StakeDetailsByTimestampKeyPrefix, timestampBz...)
}

func getFinalityOutputProposedKey() []byte {
	return FinalityOutputProposedKeyPrefix
}
