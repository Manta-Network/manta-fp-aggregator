package synchronizer

import (
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/Manta-Network/manta-fp-aggregator/common"
	"github.com/Manta-Network/manta-fp-aggregator/store"

	types3 "github.com/babylonlabs-io/babylon/x/btccheckpoint/types"
	types2 "github.com/babylonlabs-io/babylon/x/btcstaking/types"
	"github.com/babylonlabs-io/babylon/x/finality/types"
)

func (syncer *BabylonSynchronizer) ProcessNewFinalityProvider(txMessage store.TxMessage) error {
	var err error
	var mCFP types2.MsgCreateFinalityProvider

	if txMessage.Type == common.MsgCreateFinalityProvider {
		err = mCFP.Unmarshal(txMessage.Data)
		if err != nil {
			syncer.log.Error("failed to unmarshal NewFinalityProvider message value", "err", err)
			return err
		}
		err = syncer.db.SetCreateFinalityProviderMsg(store.CreateFinalityProvider{
			FP:     mCFP,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store NewFinalityProvider message", "err", err)
			return err
		}
		syncer.log.Info("success to store NewFinalityProvider message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}

func (syncer *BabylonSynchronizer) ProcessCreateBTCDelegation(txMessage store.TxMessage) error {
	var mCBD types2.MsgCreateBTCDelegation
	var txInfo types3.TransactionInfo

	if txMessage.Type == common.MsgCreateBTCDelegation {
		mCBD.Unmarshal(txMessage.Data)
		txInfo.Unmarshal(mCBD.StakingTx)
		btcTx, err := types2.NewBtcTransaction(txInfo.Transaction)
		if err != nil {
			syncer.log.Error("failed to new btc transaction", "err", err)
			return err
		}
		err = syncer.db.SetCreateBTCDelegationMsg(store.CreateBTCDelegation{
			CBD:    mCBD,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store CreateBTCDelegation message", "err", err)
			return err
		}
		err = syncer.db.SetBabylonDelegationKey(txMessage.TransactionHash, []byte(btcTx.Transaction.TxHash().String()))
		if err != nil {
			syncer.log.Error("failed to store babylon delegation key", "err", err)
			return err
		}
		syncer.log.Info("success to store CreateBTCDelegation message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}
	return nil
}

func (syncer *BabylonSynchronizer) ProcessCommitPubRandList(txMessage store.TxMessage) error {
	var err error
	var mCPR types.MsgCommitPubRandList

	if txMessage.Type == common.MsgCommitPubRandList {
		mCPR.Unmarshal(txMessage.Data)
		err = syncer.db.SetCommitPubRandListMsg(store.CommitPubRandList{
			CPR:    mCPR,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store CommitPubRandList message", "err", err)
			return err
		}
		syncer.log.Info("success to store CommitPubRandList message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}

func (syncer *BabylonSynchronizer) ProcessBTCUndelegate(txMessage store.TxMessage) error {
	var err error
	var mBUD types2.MsgBTCUndelegate

	if txMessage.Type == common.MsgBTCUndelegate {
		mBUD.Unmarshal(txMessage.Data)
		err = syncer.db.SetBtcUndelegateMsg(store.BtcUndelegate{
			BU:     mBUD,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store BTCUndelegate message", "err", err)
			return err
		}
		syncer.log.Info("success to store BTCUndelegate message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}

func (syncer *BabylonSynchronizer) ProcessSelectiveSlashingEvidence(txMessage store.TxMessage) error {
	var err error
	var mSSE types2.MsgSelectiveSlashingEvidence

	if txMessage.Type == common.MsgSelectiveSlashingEvidence {
		mSSE.Unmarshal(txMessage.Data)
		err = syncer.db.SetSelectiveSlashingEvidenceMsg(store.SelectiveSlashingEvidence{
			SSE:    mSSE,
			TxHash: txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store SelectiveSlashingEvidenc message", "err", err)
			return err
		}
		syncer.log.Info("success to store SelectiveSlashingEvidenc message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}

func (syncer *BabylonSynchronizer) ProcessSubmitFinalitySignature(txMessage store.TxMessage) error {
	var err error
	var mSFS store.SubmitFinalitySignatureMsgValue

	if txMessage.Type == common.MsgSubmitFinalitySignatureType {
		mSFS.Unmarshal(txMessage.Data)
		err = syncer.db.SetSubmitFinalitySignatureMsg(store.SubmitFinalitySignature{
			SFS:     mSFS,
			SFSByte: txMessage.Data,
			TxHash:  txMessage.TransactionHash,
		})
		if err != nil {
			syncer.log.Error("failed to store SubmitFinalitySignature message", "err", err)
			return err
		}
		syncer.log.Info("success to store SubmitFinalitySignature message", "tx_hash", hexutil.Encode(txMessage.TransactionHash))
	}

	return nil
}
