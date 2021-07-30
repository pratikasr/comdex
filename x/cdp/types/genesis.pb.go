// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: comdex/cdp/v1alpha1/genesis.proto

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

// GenesisState defines the cdp module's genesis state.
type GenesisState struct {
	Params        []Params  `protobuf:"bytes,1,rep,name=params,proto3" json:"params" yaml:"params"`
	Cdps          []CDP     `protobuf:"bytes,2,rep,name=cdps,proto3" json:"cdps" yaml:"cdps"`
	Deposits      []Deposit `protobuf:"bytes,3,rep,name=deposits,proto3" json:"deposits" yaml:"deposits"`
	StartingCdpId uint64    `protobuf:"varint,4,opt,name=starting_cdp_id,json=startingCdpId,proto3" json:"starting_cdp_id,omitempty"`
	DebtDenom     string    `protobuf:"bytes,5,opt,name=debt_denom,json=debtDenom,proto3" json:"debt_denom,omitempty" yaml:"debt_denom"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dab3778453fac25, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() []Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func (m *GenesisState) GetCdps() []CDP {
	if m != nil {
		return m.Cdps
	}
	return nil
}

func (m *GenesisState) GetDeposits() []Deposit {
	if m != nil {
		return m.Deposits
	}
	return nil
}

func (m *GenesisState) GetStartingCdpId() uint64 {
	if m != nil {
		return m.StartingCdpId
	}
	return 0
}

func (m *GenesisState) GetDebtDenom() string {
	if m != nil {
		return m.DebtDenom
	}
	return ""
}

type GenesisAccumulationTime struct {
	CollateralType           string                                 `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	PreviousAccumulationTime time.Time                              `protobuf:"bytes,2,opt,name=previous_accumulation_time,json=previousAccumulationTime,proto3,stdtime" json:"previous_accumulation_time" yaml:"previous_accumulation_time"`
	InterestFactor           github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=interest_factor,json=interestFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"interest_factor" yaml:"interest_factor"`
}

func (m *GenesisAccumulationTime) Reset()         { *m = GenesisAccumulationTime{} }
func (m *GenesisAccumulationTime) String() string { return proto.CompactTextString(m) }
func (*GenesisAccumulationTime) ProtoMessage()    {}
func (*GenesisAccumulationTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dab3778453fac25, []int{1}
}
func (m *GenesisAccumulationTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisAccumulationTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisAccumulationTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisAccumulationTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisAccumulationTime.Merge(m, src)
}
func (m *GenesisAccumulationTime) XXX_Size() int {
	return m.Size()
}
func (m *GenesisAccumulationTime) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisAccumulationTime.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisAccumulationTime proto.InternalMessageInfo

func (m *GenesisAccumulationTime) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisAccumulationTime) GetPreviousAccumulationTime() time.Time {
	if m != nil {
		return m.PreviousAccumulationTime
	}
	return time.Time{}
}

type GenesisTotalPrincipal struct {
	CollateralType string `protobuf:"bytes,1,opt,name=collateral_type,json=collateralType,proto3" json:"collateral_type,omitempty" yaml:"collateral_type"`
	TotalPrincipal int64  `protobuf:"varint,2,opt,name=total_principal,json=totalPrincipal,proto3" json:"total_principal,omitempty" yaml:"total_principal"`
}

func (m *GenesisTotalPrincipal) Reset()         { *m = GenesisTotalPrincipal{} }
func (m *GenesisTotalPrincipal) String() string { return proto.CompactTextString(m) }
func (*GenesisTotalPrincipal) ProtoMessage()    {}
func (*GenesisTotalPrincipal) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dab3778453fac25, []int{2}
}
func (m *GenesisTotalPrincipal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisTotalPrincipal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisTotalPrincipal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisTotalPrincipal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisTotalPrincipal.Merge(m, src)
}
func (m *GenesisTotalPrincipal) XXX_Size() int {
	return m.Size()
}
func (m *GenesisTotalPrincipal) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisTotalPrincipal.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisTotalPrincipal proto.InternalMessageInfo

func (m *GenesisTotalPrincipal) GetCollateralType() string {
	if m != nil {
		return m.CollateralType
	}
	return ""
}

func (m *GenesisTotalPrincipal) GetTotalPrincipal() int64 {
	if m != nil {
		return m.TotalPrincipal
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "comdex.cdp.v1alpha1.GenesisState")
	proto.RegisterType((*GenesisAccumulationTime)(nil), "comdex.cdp.v1alpha1.GenesisAccumulationTime")
	proto.RegisterType((*GenesisTotalPrincipal)(nil), "comdex.cdp.v1alpha1.GenesisTotalPrincipal")
}

func init() { proto.RegisterFile("comdex/cdp/v1alpha1/genesis.proto", fileDescriptor_0dab3778453fac25) }

var fileDescriptor_0dab3778453fac25 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x72, 0xd3, 0x3c,
	0x10, 0xc7, 0xe3, 0xa4, 0x5f, 0xe7, 0xab, 0x4a, 0x9b, 0xc1, 0xa5, 0xd4, 0x13, 0xc0, 0x4e, 0x7d,
	0xe8, 0xe4, 0x12, 0x7b, 0x5a, 0x38, 0x71, 0xab, 0x93, 0xa1, 0xc0, 0xa9, 0x98, 0x9c, 0xb8, 0x78,
	0x14, 0x49, 0x71, 0x35, 0xd8, 0x96, 0xb0, 0x94, 0x0e, 0x7d, 0x02, 0xae, 0x7d, 0x05, 0xde, 0xa6,
	0x27, 0xa6, 0x47, 0x06, 0x66, 0x02, 0xd3, 0xbe, 0x41, 0x9e, 0x80, 0xb1, 0x24, 0xd3, 0xa6, 0x84,
	0x1b, 0xa7, 0x58, 0xbb, 0xff, 0xfd, 0x69, 0xf7, 0xaf, 0x0d, 0xd8, 0x45, 0x2c, 0xc7, 0xe4, 0x63,
	0x88, 0x30, 0x0f, 0x4f, 0xf7, 0x61, 0xc6, 0x4f, 0xe0, 0x7e, 0x98, 0x92, 0x82, 0x08, 0x2a, 0x02,
	0x5e, 0x32, 0xc9, 0xec, 0x2d, 0x2d, 0x09, 0x10, 0xe6, 0x41, 0x2d, 0xe9, 0x3c, 0x48, 0x59, 0xca,
	0x54, 0x3e, 0xac, 0xbe, 0xb4, 0xb4, 0xe3, 0xa5, 0x8c, 0xa5, 0x19, 0x09, 0xd5, 0x69, 0x3c, 0x9d,
	0x84, 0x92, 0xe6, 0x44, 0x48, 0x98, 0x73, 0x23, 0x78, 0xb2, 0xec, 0xba, 0x0a, 0xac, 0xd3, 0xdd,
	0x65, 0x69, 0x0e, 0x4b, 0x98, 0x9b, 0x66, 0xfc, 0x2f, 0x4d, 0x70, 0xef, 0x48, 0xb7, 0xf7, 0x56,
	0x42, 0x49, 0xec, 0xd7, 0x60, 0x55, 0x0b, 0x1c, 0xab, 0xdb, 0xea, 0xad, 0x1f, 0x3c, 0x0a, 0x96,
	0xb4, 0x1b, 0x1c, 0x2b, 0x49, 0xb4, 0x7d, 0x31, 0xf3, 0x1a, 0xf3, 0x99, 0xb7, 0x71, 0x06, 0xf3,
	0xec, 0xb9, 0xaf, 0x0b, 0xfd, 0xd8, 0x10, 0xec, 0x43, 0xb0, 0x82, 0x30, 0x17, 0x4e, 0x53, 0x91,
	0x9c, 0xa5, 0xa4, 0xc1, 0xf0, 0x38, 0xda, 0x32, 0x98, 0x75, 0x8d, 0xa9, 0x6a, 0xfc, 0x58, 0x95,
	0xda, 0x6f, 0xc0, 0xff, 0x98, 0x70, 0x26, 0xa8, 0x14, 0x4e, 0x4b, 0x61, 0x1e, 0x2f, 0xc5, 0x0c,
	0xb5, 0x28, 0xda, 0x31, 0xa8, 0xb6, 0x46, 0xd5, 0xb5, 0x7e, 0xfc, 0x1b, 0x63, 0xef, 0x81, 0xb6,
	0x90, 0xb0, 0x94, 0xb4, 0x48, 0x13, 0x84, 0x79, 0x42, 0xb1, 0xb3, 0xd2, 0xb5, 0x7a, 0x2b, 0xf1,
	0x46, 0x1d, 0x1e, 0x60, 0xfe, 0x0a, 0xdb, 0xcf, 0x00, 0xc0, 0x64, 0x2c, 0x13, 0x4c, 0x0a, 0x96,
	0x3b, 0xff, 0x75, 0xad, 0xde, 0x5a, 0xb4, 0x3d, 0x9f, 0x79, 0xf7, 0x6b, 0x74, 0x9d, 0xf3, 0xe3,
	0xb5, 0xea, 0x30, 0x54, 0xdf, 0xdf, 0x9b, 0x60, 0xc7, 0x18, 0x7a, 0x88, 0xd0, 0x34, 0x9f, 0x66,
	0x50, 0x52, 0x56, 0x8c, 0x68, 0x4e, 0xec, 0x01, 0x68, 0x23, 0x96, 0x65, 0x50, 0x92, 0x12, 0x66,
	0x89, 0x3c, 0xe3, 0xc4, 0xb1, 0x14, 0xb6, 0x33, 0x9f, 0x79, 0x0f, 0xcd, 0xf0, 0x8b, 0x02, 0x3f,
	0xde, 0xbc, 0x89, 0x8c, 0xce, 0x38, 0xb1, 0x3f, 0x59, 0xa0, 0xc3, 0x4b, 0x72, 0x4a, 0xd9, 0x54,
	0x24, 0xf0, 0xd6, 0x15, 0x49, 0xb5, 0x1c, 0x4e, 0xb3, 0x6b, 0xf5, 0xd6, 0x0f, 0x3a, 0x81, 0xde,
	0x9c, 0xa0, 0xde, 0x9c, 0x60, 0x54, 0x6f, 0x4e, 0xd4, 0x37, 0x16, 0xed, 0x9a, 0x47, 0xfb, 0x2b,
	0xcb, 0x3f, 0xff, 0xe1, 0x59, 0xb1, 0x53, 0x0b, 0xfe, 0x18, 0xe7, 0x03, 0x68, 0xd3, 0x42, 0x92,
	0x92, 0x08, 0x99, 0x4c, 0x20, 0x92, 0xac, 0x74, 0x5a, 0x6a, 0x9c, 0x97, 0xd5, 0x0d, 0xdf, 0x66,
	0xde, 0x5e, 0x4a, 0xe5, 0xc9, 0x74, 0x5c, 0x3d, 0x58, 0x88, 0x98, 0xc8, 0x99, 0x30, 0x3f, 0x7d,
	0x81, 0xdf, 0x87, 0xd5, 0x78, 0x22, 0x18, 0x12, 0x74, 0x33, 0xfc, 0x1d, 0x9c, 0x1f, 0x6f, 0xd6,
	0x91, 0x17, 0x3a, 0xf0, 0xd9, 0x02, 0xdb, 0xc6, 0xdd, 0x11, 0x93, 0x30, 0x3b, 0x2e, 0x69, 0x81,
	0x28, 0x87, 0xd9, 0xbf, 0xf1, 0x76, 0x00, 0xda, 0xb2, 0xc2, 0x26, 0xbc, 0xe6, 0x2a, 0x3f, 0x5b,
	0xb7, 0x21, 0x77, 0x04, 0x7e, 0xbc, 0x29, 0x17, 0x3a, 0x89, 0x8e, 0x2e, 0xae, 0x5c, 0xeb, 0xf2,
	0xca, 0xb5, 0x7e, 0x5e, 0xb9, 0xd6, 0xf9, 0xb5, 0xdb, 0xb8, 0xbc, 0x76, 0x1b, 0x5f, 0xaf, 0xdd,
	0xc6, 0xbb, 0xfe, 0x82, 0x1f, 0xd5, 0x12, 0xf7, 0xd9, 0x64, 0x42, 0x11, 0x85, 0x99, 0x39, 0x87,
	0xfa, 0xbf, 0xaa, 0xac, 0x19, 0xaf, 0xaa, 0xc7, 0x7b, 0xfa, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x09,
	0x6d, 0x13, 0x7e, 0x54, 0x04, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.DebtDenom) > 0 {
		i -= len(m.DebtDenom)
		copy(dAtA[i:], m.DebtDenom)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.DebtDenom)))
		i--
		dAtA[i] = 0x2a
	}
	if m.StartingCdpId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.StartingCdpId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Deposits) > 0 {
		for iNdEx := len(m.Deposits) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Deposits[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Cdps) > 0 {
		for iNdEx := len(m.Cdps) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Cdps[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Params) > 0 {
		for iNdEx := len(m.Params) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Params[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GenesisAccumulationTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisAccumulationTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisAccumulationTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.InterestFactor.Size()
		i -= size
		if _, err := m.InterestFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.PreviousAccumulationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintGenesis(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x12
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GenesisTotalPrincipal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisTotalPrincipal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisTotalPrincipal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.TotalPrincipal != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TotalPrincipal))
		i--
		dAtA[i] = 0x10
	}
	if len(m.CollateralType) > 0 {
		i -= len(m.CollateralType)
		copy(dAtA[i:], m.CollateralType)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.CollateralType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Params) > 0 {
		for _, e := range m.Params {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Cdps) > 0 {
		for _, e := range m.Cdps {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Deposits) > 0 {
		for _, e := range m.Deposits {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.StartingCdpId != 0 {
		n += 1 + sovGenesis(uint64(m.StartingCdpId))
	}
	l = len(m.DebtDenom)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	return n
}

func (m *GenesisAccumulationTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.PreviousAccumulationTime)
	n += 1 + l + sovGenesis(uint64(l))
	l = m.InterestFactor.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func (m *GenesisTotalPrincipal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollateralType)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.TotalPrincipal != 0 {
		n += 1 + sovGenesis(uint64(m.TotalPrincipal))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Params = append(m.Params, Params{})
			if err := m.Params[len(m.Params)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cdps", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Cdps = append(m.Cdps, CDP{})
			if err := m.Cdps[len(m.Cdps)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deposits", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deposits = append(m.Deposits, Deposit{})
			if err := m.Deposits[len(m.Deposits)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartingCdpId", wireType)
			}
			m.StartingCdpId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartingCdpId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DebtDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DebtDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisAccumulationTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisAccumulationTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisAccumulationTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousAccumulationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.PreviousAccumulationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InterestFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InterestFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *GenesisTotalPrincipal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisTotalPrincipal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisTotalPrincipal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollateralType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollateralType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TotalPrincipal", wireType)
			}
			m.TotalPrincipal = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TotalPrincipal |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
