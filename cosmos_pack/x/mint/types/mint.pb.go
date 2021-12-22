// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Minter represents the minting state.
type Minter struct {
	// current annual inflation rate //当年的通货膨胀率
	Inflation github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=inflation,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation"`
	// current annual expected provisions  //当前年通胀率下每年新铸造的链上资产数量
	AnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=annual_provisions,json=annualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"annual_provisions" yaml:"annual_provisions"`
	//TAT占比
	TatProbability github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=tatprobability,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"tatprobability"`
	//新的链上资产数量
	NewAnnualProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=newannual_provisions,json=newannualProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"newannual_provisions" yaml:"newannual_provisions"`
	//累计aunit的发送量
	UnitGrant github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=unitgrant,json=unitGrant,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"unitgrant" yaml:"unitgrant"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// maximum annual change in inflation rate 通货膨胀率的最大年度变化
	InflationRateChange github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=inflation_rate_change,json=inflationRateChange,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_rate_change" yaml:"inflation_rate_change"`
	// maximum inflation rate 最高通货膨胀率
	InflationMax github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=inflation_max,json=inflationMax,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_max" yaml:"inflation_max"`
	// minimum inflation rate 最低通货膨胀率
	InflationMin github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=inflation_min,json=inflationMin,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"inflation_min" yaml:"inflation_min"`
	// goal of percent bonded atoms 键合原子百分比目标(意思是达到这个百分比后开始改变通货膨胀率来产生币，每个区块产生的币的数量都不一样)
	GoalBonded github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,5,opt,name=goal_bonded,json=goalBonded,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"goal_bonded" yaml:"goal_bonded"`
	//unit上限 TAT产量的百分比
	Probability github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,6,opt,name=pro_bability,json=Probability,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"pro_bability" yaml:"pro_bability"`
	//TAT产量累加(第一年的TAT总产量)
	UnitGrant uint64 `protobuf:"varint,7,opt,name=unit_grant,json=unitgrant,proto3" json:"unit_grant,omitempty" yaml:"unit_grant"`
	// expected blocks per year  每年的预期区块
	BlocksPerYear uint64 `protobuf:"varint,11,opt,name=blocks_per_year,json=blocksPerYear,proto3" json:"blocks_per_year,omitempty" yaml:"blocks_per_year"`
	//监听开始块
	StartBlock int64 `protobuf:"varint,8,opt,name=start_block,json=startblock,proto3" json:"start_block,omitempty" yaml:"start_block"`
	//监听结束块
	EndBlock int64 `protobuf:"varint,9,opt,name=end_block,json=endblock,proto3" json:"end_block,omitempty" yaml:"end_block"`
	//监听高度
	HeightBlock int64 `protobuf:"varint,10,opt,name=height_block,json=heightblock,proto3" json:"height_block,omitempty" yaml:"height_block"`
	//第一年，第二年每个块的基本奖励
	PerReward int64 `protobuf:"varint,12,rep,name=per_reward,json=perReward,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"per_reward,omitempty" yaml:"per_reward"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2df116d183c1e223, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetBlocksPerYear() uint64 {
	if m != nil {
		return m.BlocksPerYear
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "cosmos.mint.v1beta1.Minter")
	proto.RegisterType((*Params)(nil), "cosmos.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("cosmos/mint/v1beta1/mint.proto", fileDescriptor_2df116d183c1e223) }

var fileDescriptor_2df116d183c1e223 = []byte{
	// 437 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x1c, 0xc6, 0x63, 0x28, 0x95, 0x6a, 0x98, 0x00, 0x6f, 0xa0, 0x68, 0x82, 0x64, 0xca, 0x01, 0x8d,
	0x03, 0x89, 0x26, 0x6e, 0x3b, 0x66, 0x15, 0x07, 0xc4, 0x50, 0xe5, 0x1b, 0x5c, 0xa2, 0x7f, 0x52,
	0x93, 0x59, 0x8d, 0xed, 0xca, 0xf6, 0x46, 0x7b, 0xe5, 0x09, 0x38, 0x72, 0xe4, 0x2d, 0x78, 0x85,
	0xdd, 0xd8, 0x11, 0x71, 0xa8, 0x50, 0xfb, 0x06, 0x7b, 0x02, 0x14, 0xbb, 0x6a, 0xa1, 0x20, 0xa4,
	0x4a, 0x9c, 0x92, 0xef, 0xfb, 0xff, 0xf3, 0xfd, 0xbe, 0x58, 0x32, 0x8e, 0x2a, 0x65, 0x84, 0x32,
	0x99, 0xe0, 0xd2, 0x66, 0x17, 0x47, 0x25, 0xb3, 0x70, 0xe4, 0x44, 0x3a, 0xd6, 0xca, 0x2a, 0xb2,
	0xeb, 0xe7, 0xa9, 0xb3, 0x96, 0xf3, 0xfd, 0xbd, 0x5a, 0xd5, 0xca, 0xcd, 0xb3, 0xf6, 0xcd, 0xaf,
	0x26, 0x5f, 0x11, 0xee, 0x9e, 0x72, 0x69, 0x99, 0x26, 0xaf, 0x70, 0x8f, 0xcb, 0x77, 0x0d, 0x58,
	0xae, 0x64, 0x88, 0x0e, 0xd0, 0x61, 0x2f, 0x4f, 0x2f, 0x67, 0x71, 0xf0, 0x7d, 0x16, 0x3f, 0xa9,
	0xb9, 0x3d, 0x3b, 0x2f, 0xd3, 0x4a, 0x89, 0x6c, 0xc9, 0xf6, 0x8f, 0x67, 0x66, 0x38, 0xca, 0xec,
	0x74, 0xcc, 0x4c, 0xda, 0x67, 0x15, 0x5d, 0x07, 0x90, 0xf7, 0xf8, 0x3e, 0x48, 0x79, 0x0e, 0x4d,
	0x31, 0xd6, 0xea, 0x82, 0x1b, 0xae, 0xa4, 0x09, 0x6f, 0xb8, 0xd4, 0x97, 0xdb, 0xa5, 0x5e, 0xcf,
	0xe2, 0x70, 0x0a, 0xa2, 0x39, 0x4e, 0xfe, 0x08, 0x4c, 0xe8, 0x3d, 0xef, 0x0d, 0xd6, 0xd6, 0x97,
	0x0e, 0xee, 0x0e, 0x40, 0x83, 0x30, 0xe4, 0x31, 0xc6, 0xed, 0x11, 0x14, 0x43, 0x26, 0x95, 0xf0,
	0xbf, 0x44, 0x7b, 0xad, 0xd3, 0x6f, 0x0d, 0xf2, 0x01, 0xe1, 0x07, 0xab, 0xc2, 0x85, 0x06, 0xcb,
	0x8a, 0xea, 0x0c, 0x64, 0xcd, 0x96, 0x3d, 0x5f, 0x6f, 0xdd, 0xf3, 0x91, 0xef, 0xf9, 0xd7, 0xd0,
	0x84, 0xee, 0xae, 0x7c, 0x0a, 0x96, 0x9d, 0x38, 0x97, 0x8c, 0xf0, 0xce, 0x7a, 0x5d, 0xc0, 0x24,
	0xbc, 0xe9, 0xd8, 0x2f, 0xb6, 0x66, 0xef, 0x6d, 0xb2, 0x05, 0x4c, 0x12, 0x7a, 0x67, 0xa5, 0x4f,
	0x61, 0xb2, 0x01, 0xe3, 0x32, 0xec, 0xfc, 0x37, 0x18, 0x97, 0xbf, 0xc1, 0xb8, 0x24, 0x0c, 0xdf,
	0xae, 0x15, 0x34, 0x45, 0xa9, 0xe4, 0x90, 0x0d, 0xc3, 0x5b, 0x0e, 0xd5, 0xdf, 0x1a, 0x45, 0x3c,
	0xea, 0x97, 0xa8, 0x84, 0xe2, 0x56, 0xe5, 0x4e, 0x90, 0x1c, 0xdf, 0x2d, 0x1b, 0x55, 0x8d, 0x4c,
	0x31, 0x66, 0xba, 0x98, 0x32, 0xd0, 0x61, 0xf7, 0x00, 0x1d, 0x76, 0xf2, 0xfd, 0xeb, 0x59, 0xfc,
	0xd0, 0x7f, 0xbc, 0xb1, 0x90, 0xd0, 0x1d, 0xef, 0x0c, 0x98, 0x7e, 0xc3, 0x40, 0x1f, 0x77, 0x3e,
	0x7d, 0x8e, 0x83, 0xfc, 0xe4, 0x72, 0x1e, 0xa1, 0xab, 0x79, 0x84, 0x7e, 0xcc, 0x23, 0xf4, 0x71,
	0x11, 0x05, 0x57, 0x8b, 0x28, 0xf8, 0xb6, 0x88, 0x82, 0xb7, 0x4f, 0xff, 0xd9, 0x76, 0xe2, 0x2f,
	0xa2, 0x2b, 0x5d, 0x76, 0xdd, 0xbd, 0x7a, 0xfe, 0x33, 0x00, 0x00, 0xff, 0xff, 0x34, 0x52, 0xcf,
	0x87, 0xa4, 0x03, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	fmt.Println("size:", size)
	dAtA = make([]byte, size)
	fmt.Println("dAtA:", dAtA)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	fmt.Println("n:", n)
	if err != nil {
		return nil, err
	}
	fmt.Println("dAtA[:n]:", dAtA[:n])
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AnnualProvisions.Size()
		i -= size
		if _, err := m.AnnualProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.Inflation.Size()
		i -= size
		if _, err := m.Inflation.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BlocksPerYear != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.BlocksPerYear))
		i--
		dAtA[i] = 0x30
	}
	{
		size := m.GoalBonded.Size()
		i -= size
		if _, err := m.GoalBonded.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.InflationMin.Size()
		i -= size
		if _, err := m.InflationMin.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.InflationMax.Size()
		i -= size
		if _, err := m.InflationMax.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.InflationRateChange.Size()
		i -= size
		if _, err := m.InflationRateChange.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Inflation.Size()
	// fmt.Println("L1:", l)
	n += 1 + l + sovMint(uint64(l))
	// fmt.Println("n1:", n)
	l = m.AnnualProvisions.Size()
	// fmt.Println("L2:", l)
	n += 1 + l + sovMint(uint64(l))
	// fmt.Println("n2:", n)
	// l = m.TatProbability.Size()
	// fmt.Println("L3:", l)
	// n += 1 + l + sovMint(uint64(l))
	// fmt.Println("n3:", n)
	// l = m.NewAnnualProvisions.Size()
	// fmt.Println("L4:", l)
	// n += 1 + l + sovMint(uint64(l))
	// fmt.Println("n4:", n)
	// l = m.UnitGrant.Size()
	// fmt.Println("L5:", l)
	// n += 1 + l + sovMint(uint64(l))
	// fmt.Println("n5:", n)
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.InflationRateChange.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMax.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.InflationMin.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.GoalBonded.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.BlocksPerYear != 0 {
		n += 1 + sovMint(uint64(m.BlocksPerYear))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
		fieldNum := int32(wire >> 3) // 左移 wire乘以2的3次,右移就是除以2的3次
		fmt.Println("fieldNum:", fieldNum)
		wireType := int(wire & 0x7)
		fmt.Println("wireType:", wireType)
		if wireType == 4 {
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Inflation", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Inflation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AnnualProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AnnualProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationRateChange", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationRateChange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMax", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMax.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InflationMin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InflationMin.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GoalBonded", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GoalBonded.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlocksPerYear", wireType)
			}
			m.BlocksPerYear = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlocksPerYear |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
