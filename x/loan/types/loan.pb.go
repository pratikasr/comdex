// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/loan/v1beta1/loan.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Loan struct {
	Id           uint64                                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppId        uint64                                 `protobuf:"varint,2,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty" yaml:"app_id"`
	PairID       uint64                                 `protobuf:"varint,3,opt,name=pair_id,json=pairId,proto3" json:"pair_id,omitempty" yaml:"pair_id"`
	Owner        string                                 `protobuf:"bytes,4,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	AmountIn     github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,5,opt,name=amount_in,json=amountIn,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_in" yaml:"amount_in"`
	AmountOut    github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,6,opt,name=amount_out,json=amountOut,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount_out" yaml:"amount_out"`
	Fee          github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,7,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee" yaml:"fee"`
	CreatedAt    time.Time                              `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at" yaml:"created_at"`
	DurationDays int64                                  `protobuf:"varint,9,opt,name=duration_days,json=durationDays,proto3" json:"duration_days,omitempty" yaml:"duration_days"`
	Lender       string                                 `protobuf:"bytes,10,opt,name=lender,proto3" json:"lender,omitempty" yaml:"owner"`
	State        string                                 `protobuf:"bytes,11,opt,name=state,proto3" json:"state,omitempty" yaml:"state"`
}

func (m *Loan) Reset()         { *m = Loan{} }
func (m *Loan) String() string { return proto.CompactTextString(m) }
func (*Loan) ProtoMessage()    {}
func (*Loan) Descriptor() ([]byte, []int) {
	return fileDescriptor_e8ca162d9d63a6ca, []int{0}
}
func (m *Loan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Loan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Loan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Loan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Loan.Merge(m, src)
}
func (m *Loan) XXX_Size() int {
	return m.Size()
}
func (m *Loan) XXX_DiscardUnknown() {
	xxx_messageInfo_Loan.DiscardUnknown(m)
}

var xxx_messageInfo_Loan proto.InternalMessageInfo

func (m *Loan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Loan) GetAppId() uint64 {
	if m != nil {
		return m.AppId
	}
	return 0
}

func (m *Loan) GetPairID() uint64 {
	if m != nil {
		return m.PairID
	}
	return 0
}

func (m *Loan) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Loan) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Loan) GetDurationDays() int64 {
	if m != nil {
		return m.DurationDays
	}
	return 0
}

func (m *Loan) GetLender() string {
	if m != nil {
		return m.Lender
	}
	return ""
}

func (m *Loan) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func init() {
	proto.RegisterType((*Loan)(nil), "comdex.loan.v1beta1.Loan")
}

func init() { proto.RegisterFile("comdex/loan/v1beta1/loan.proto", fileDescriptor_e8ca162d9d63a6ca) }

var fileDescriptor_e8ca162d9d63a6ca = []byte{
	// 514 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x4d, 0x6f, 0xd3, 0x30,
	0x18, 0x6e, 0xd6, 0x36, 0x5b, 0xbd, 0x0f, 0x0d, 0xb3, 0x83, 0x55, 0x41, 0x5c, 0xf9, 0x30, 0xf5,
	0xb2, 0x44, 0x05, 0x71, 0x41, 0x70, 0x58, 0xd8, 0x81, 0x4a, 0x08, 0x90, 0xc5, 0x01, 0x71, 0xa9,
	0xdc, 0xda, 0x2d, 0x16, 0x4d, 0x1c, 0x25, 0x0e, 0xd0, 0x7f, 0xb1, 0x9f, 0xc1, 0x4f, 0xd9, 0x71,
	0x47, 0xc4, 0x21, 0xa0, 0xf6, 0x1f, 0xe4, 0x17, 0xa0, 0xd8, 0xae, 0xa6, 0x22, 0x71, 0xd8, 0xc9,
	0xef, 0xfb, 0x3c, 0xcf, 0xfb, 0xa1, 0xf7, 0x91, 0x41, 0x30, 0x53, 0x09, 0x17, 0xdf, 0xa3, 0xa5,
	0x62, 0x69, 0xf4, 0x75, 0x34, 0x15, 0x9a, 0x8d, 0x4c, 0x12, 0x66, 0xb9, 0xd2, 0x0a, 0x3e, 0xb4,
	0x7c, 0x68, 0x20, 0xc7, 0xf7, 0xcf, 0x16, 0x6a, 0xa1, 0x0c, 0x1f, 0x35, 0x91, 0x95, 0xf6, 0xf1,
	0x42, 0xa9, 0xc5, 0x52, 0x44, 0x26, 0x9b, 0x96, 0xf3, 0x48, 0xcb, 0x44, 0x14, 0x9a, 0x25, 0x99,
	0x15, 0x90, 0x1f, 0x5d, 0xd0, 0x79, 0xa3, 0x58, 0x0a, 0x4f, 0xc0, 0x9e, 0xe4, 0xc8, 0x1b, 0x78,
	0xc3, 0x0e, 0xdd, 0x93, 0x1c, 0x8e, 0x80, 0xcf, 0xb2, 0x6c, 0x22, 0x39, 0xda, 0x6b, 0xb0, 0xb8,
	0xbf, 0xae, 0x70, 0xf7, 0x32, 0xcb, 0xc6, 0xbc, 0xae, 0xf0, 0xf1, 0x8a, 0x25, 0xcb, 0xe7, 0xc4,
	0x0a, 0x08, 0xed, 0xb2, 0x06, 0x87, 0xcf, 0xc0, 0x7e, 0xc6, 0x64, 0xde, 0xd4, 0xb4, 0x4d, 0xcd,
	0xa3, 0x75, 0x85, 0xfd, 0xf7, 0x4c, 0xe6, 0xe3, 0xab, 0xba, 0xc2, 0x27, 0xb6, 0xc8, 0x49, 0x08,
	0xf5, 0x9b, 0x68, 0xcc, 0xe1, 0x39, 0xe8, 0xaa, 0x6f, 0xa9, 0xc8, 0x51, 0x67, 0xe0, 0x0d, 0x7b,
	0xf1, 0x69, 0x5d, 0xe1, 0x23, 0x2b, 0x35, 0x30, 0xa1, 0x96, 0x86, 0x13, 0xd0, 0x63, 0x89, 0x2a,
	0x53, 0x3d, 0x91, 0x29, 0xea, 0x1a, 0x6d, 0x7c, 0x53, 0xe1, 0xd6, 0xaf, 0x0a, 0x9f, 0x2f, 0xa4,
	0xfe, 0x5c, 0x4e, 0xc3, 0x99, 0x4a, 0xa2, 0x99, 0x2a, 0x12, 0x55, 0xb8, 0xe7, 0xa2, 0xe0, 0x5f,
	0x22, 0xbd, 0xca, 0x44, 0x11, 0x8e, 0x53, 0x5d, 0x57, 0xf8, 0xd4, 0x6d, 0xbe, 0x6d, 0x44, 0xe8,
	0x81, 0x8d, 0xc7, 0x29, 0x9c, 0x02, 0xe0, 0x70, 0x55, 0x6a, 0xe4, 0x9b, 0x09, 0xaf, 0xee, 0x3d,
	0xe1, 0xc1, 0xce, 0x04, 0x55, 0x6a, 0x42, 0xdd, 0xde, 0xef, 0x4a, 0x0d, 0xdf, 0x82, 0xf6, 0x5c,
	0x08, 0xb4, 0x6f, 0x9a, 0xbf, 0xb8, 0x77, 0x73, 0x60, 0x9b, 0xcf, 0x85, 0x20, 0xb4, 0x69, 0x04,
	0x3f, 0x02, 0x30, 0xcb, 0x05, 0xd3, 0x82, 0x4f, 0x98, 0x46, 0x07, 0x03, 0x6f, 0x78, 0xf8, 0xa4,
	0x1f, 0x5a, 0xd7, 0xc3, 0xad, 0xeb, 0xe1, 0x87, 0xad, 0xeb, 0xf1, 0xe3, 0x66, 0xe4, 0xdd, 0x96,
	0x77, 0xb5, 0xe4, 0xfa, 0x37, 0xf6, 0x68, 0xcf, 0x01, 0x97, 0x1a, 0xbe, 0x04, 0xc7, 0xbc, 0xcc,
	0x99, 0x96, 0x2a, 0x9d, 0x70, 0xb6, 0x2a, 0x50, 0x6f, 0xe0, 0x0d, 0xdb, 0x31, 0xaa, 0x2b, 0x7c,
	0x66, 0x8b, 0x77, 0x68, 0x42, 0x8f, 0xb6, 0xf9, 0x15, 0x5b, 0x15, 0x70, 0x08, 0xfc, 0xa5, 0x48,
	0xb9, 0xc8, 0x11, 0xf8, 0x8f, 0xad, 0x8e, 0x6f, 0xfc, 0x2f, 0x34, 0xd3, 0x02, 0x1d, 0xfe, 0x2b,
	0x34, 0x30, 0xa1, 0x96, 0x8e, 0x5f, 0xdf, 0xac, 0x03, 0xef, 0x76, 0x1d, 0x78, 0x7f, 0xd6, 0x81,
	0x77, 0xbd, 0x09, 0x5a, 0xb7, 0x9b, 0xa0, 0xf5, 0x73, 0x13, 0xb4, 0x3e, 0x85, 0x3b, 0xf7, 0x6b,
	0xfe, 0xc6, 0x85, 0x9a, 0xcf, 0xe5, 0x4c, 0xb2, 0xa5, 0xcb, 0x23, 0xf7, 0x9b, 0xcc, 0x2d, 0xa7,
	0xbe, 0x39, 0xcc, 0xd3, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x88, 0xcc, 0xc9, 0x69, 0x03,
	0x00, 0x00,
}

func (m *Loan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Loan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Loan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.Lender) > 0 {
		i -= len(m.Lender)
		copy(dAtA[i:], m.Lender)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Lender)))
		i--
		dAtA[i] = 0x52
	}
	if m.DurationDays != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.DurationDays))
		i--
		dAtA[i] = 0x48
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLoan(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLoan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.AmountOut.Size()
		i -= size
		if _, err := m.AmountOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLoan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	{
		size := m.AmountIn.Size()
		i -= size
		if _, err := m.AmountIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLoan(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintLoan(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x22
	}
	if m.PairID != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.PairID))
		i--
		dAtA[i] = 0x18
	}
	if m.AppId != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.AppId))
		i--
		dAtA[i] = 0x10
	}
	if m.Id != 0 {
		i = encodeVarintLoan(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLoan(dAtA []byte, offset int, v uint64) int {
	offset -= sovLoan(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Loan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLoan(uint64(m.Id))
	}
	if m.AppId != 0 {
		n += 1 + sovLoan(uint64(m.AppId))
	}
	if m.PairID != 0 {
		n += 1 + sovLoan(uint64(m.PairID))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = m.AmountIn.Size()
	n += 1 + l + sovLoan(uint64(l))
	l = m.AmountOut.Size()
	n += 1 + l + sovLoan(uint64(l))
	l = m.Fee.Size()
	n += 1 + l + sovLoan(uint64(l))
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovLoan(uint64(l))
	if m.DurationDays != 0 {
		n += 1 + sovLoan(uint64(m.DurationDays))
	}
	l = len(m.Lender)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovLoan(uint64(l))
	}
	return n
}

func sovLoan(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLoan(x uint64) (n int) {
	return sovLoan(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Loan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLoan
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
			return fmt.Errorf("proto: Loan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Loan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppId", wireType)
			}
			m.AppId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AppId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairID", wireType)
			}
			m.PairID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountIn", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountOut", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DurationDays", wireType)
			}
			m.DurationDays = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DurationDays |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Lender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLoan
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
				return ErrInvalidLengthLoan
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLoan
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLoan(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLoan
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
func skipLoan(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLoan
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowLoan
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthLoan
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLoan
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLoan
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLoan        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLoan          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLoan = fmt.Errorf("proto: unexpected end of group")
)
