package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

type SubmitFinalitySignature struct {
	SFS     SubmitFinalitySignatureMsgValue
	SFSByte []byte
	TxHash  []byte `json:"tx_hash"`
}

func (s *Storage) SetSubmitFinalitySignatureMsg(msg SubmitFinalitySignature) error {
	bz, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.db.Put(getSubmitFinalitySignatureMsgKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetSubmitFinalitySignatureMsg(txHash []byte) (bool, SubmitFinalitySignature) {
	sFSB, err := s.db.Get(getSubmitFinalitySignatureMsgKey(txHash), nil)
	if err != nil {
		return handleError2(SubmitFinalitySignature{}, err)
	}
	var sFs SubmitFinalitySignature
	if err = json.Unmarshal(sFSB, &sFs); err != nil {
		return false, SubmitFinalitySignature{}
	}
	return true, sFs
}

type WrapperSFs struct {
	SubmitFinalitySignature SubmitFinalitySignatureMsgParams `json:"submit_finality_signature"`
	BlockNumber             uint64                           `json:"block_number"`
	Timestamp               uint64                           `json:"timestamp"`
	TransactionHash         []byte                           `json:"transaction_hash"`
}

type BatchSubmitFinalitySignatures struct {
	Signatures []WrapperSFs `json:"signatures"`
}

type SubmitFinalitySignatureMsgParams struct {
	FpPubkeyHex    string `json:"fp_pubkey_hex"`
	L1BlockNumber  uint64 `json:"l1_block_number"`
	L1BlockHashHex string `json:"l1_block_hash_hex"`
	Height         uint64 `json:"height"`
	PubRand        []byte `json:"pub_rand"`
	Proof          Proof  `json:"proof"`
	BlockHash      []byte `json:"block_hash"`
	StateRoot      string `json:"state_root"`
	Signature      []byte `json:"signature"`
}

type Proof struct {
	Total    uint64   `json:"total"`
	Index    uint64   `json:"index"`
	LeafHash []byte   `json:"leaf_hash"`
	Aunts    [][]byte `json:"aunts"`
}

func (s *Storage) SetBabylonSubmitFinalitySignature(sfs WrapperSFs) error {
	var batchSignatures BatchSubmitFinalitySignatures
	bsb, err := s.db.Get(getSubmitFinalitySignatureKey(sfs.Timestamp), nil)
	if err != nil {
		if errors.Is(err, leveldb.ErrNotFound) {
			batchSignatures.Signatures = append(batchSignatures.Signatures, sfs)
			bsz, err := json.Marshal(batchSignatures)
			if err != nil {
				return err
			}
			return s.db.Put(getSubmitFinalitySignatureKey(sfs.Timestamp), bsz, nil)
		}
		return err
	}

	err = json.Unmarshal(bsb, &batchSignatures)
	if err != nil {
		return err
	}
	batchSignatures.Signatures = append(batchSignatures.Signatures, sfs)
	bsz, err := json.Marshal(batchSignatures)
	if err != nil {
		return err
	}
	return s.db.Put(getSubmitFinalitySignatureKey(sfs.Timestamp), bsz, nil)
}

func (s *Storage) GetBabylonFinalitySignatureByTimestamp(start uint64, end uint64) ([]WrapperSFs, error) {
	var batchSubmitFinalitySignatures []BatchSubmitFinalitySignatures
	var results []WrapperSFs
	iter := s.db.NewIterator(&util.Range{Start: getSubmitFinalitySignatureKey(start), Limit: getSubmitFinalitySignatureKey(end)}, nil)
	defer iter.Release()

	for iter.Next() {
		var bsfs BatchSubmitFinalitySignatures
		if err := json.Unmarshal(iter.Value(), &bsfs); err != nil {
			return nil, err
		}
		batchSubmitFinalitySignatures = append(batchSubmitFinalitySignatures, bsfs)
	}

	if err := iter.Error(); err != nil {
		return nil, err
	}

	for _, batchSignatures := range batchSubmitFinalitySignatures {
		for _, signature := range batchSignatures.Signatures {
			results = append(results, signature)
		}
	}

	return results, nil
}

type SubmitFinalitySignatureMsgValue struct {
	Sender                  string `json:"sender"`
	Contract                string `json:"contract"`
	SubmitFinalitySignature []byte `json:"submit_finality_signature"`
}

var (
	ErrInvalidLengthTx = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx   = fmt.Errorf("proto: integer overflow")
)

func (m *SubmitFinalitySignatureMsgValue) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MessageValue: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MessageValue: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Msg", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubmitFinalitySignature = append(m.SubmitFinalitySignature[:0], dAtA[iNdEx:postIndex]...)
			if m.SubmitFinalitySignature == nil {
				m.SubmitFinalitySignature = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	for i := 0; i < l; i++ {
		b := dAtA[i]
		if b < 0x80 {
			return i + 1, nil
		}
	}
	return l, io.ErrUnexpectedEOF
}
