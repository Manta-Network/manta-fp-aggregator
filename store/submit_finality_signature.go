package store

import (
	"encoding/json"
	"fmt"
	"io"
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
	return s.db.Put(getSubmitFinalitySignatureKey(msg.TxHash), bz, nil)
}

func (s *Storage) GetSubmitFinalitySignatureMsg(txHash []byte) (bool, SubmitFinalitySignature) {
	sFSB, err := s.db.Get(getSubmitFinalitySignatureKey(txHash), nil)
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
}

type SubmitFinalitySignatureMsgParams struct {
	FpPubkeyHex   string `json:"fp_pubkey_hex"`
	L1BlockNumber uint64 `json:"l1_block_number"`
	L1BlockHash   string `json:"l1_block_hash"`
	L2BlockNumber uint64 `json:"l2_block_number"`
	PubRand       []byte `json:"pub_rand"`
	Proof         Proof  `json:"proof"`
	StateRoot     string `json:"state_root"`
	Signature     []byte `json:"signature"`
}

type Proof struct {
	Total    uint64   `json:"total"`
	Index    uint64   `json:"index"`
	LeafHash []byte   `json:"leaf_hash"`
	Aunts    [][]byte `json:"aunts"`
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
