// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/stats/asset_stats.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type AssetDailyStats struct {
	// fee policy module
	AmountWithFee []types.Coin `protobuf:"bytes,1,rep,name=amountWithFee,proto3" json:"amountWithFee"`
	AmountNoFee   []types.Coin `protobuf:"bytes,2,rep,name=amountNoFee,proto3" json:"amountNoFee"`
	Fee           []types.Coin `protobuf:"bytes,3,rep,name=fee,proto3" json:"fee"`
	CountWithFee  int32        `protobuf:"varint,4,opt,name=countWithFee,proto3" json:"countWithFee,omitempty"`
	CountNoFee    int32        `protobuf:"varint,5,opt,name=countNoFee,proto3" json:"countNoFee,omitempty"`
	// core module
	Burned            []types.Coin `protobuf:"bytes,6,rep,name=burned,proto3" json:"burned"`
	CountBurned       uint64       `protobuf:"varint,7,opt,name=countBurned,proto3" json:"countBurned,omitempty"`
	Issued            []types.Coin `protobuf:"bytes,8,rep,name=issued,proto3" json:"issued"`
	CountIssued       uint64       `protobuf:"varint,9,opt,name=countIssued,proto3" json:"countIssued,omitempty"`
	Withdraw          []types.Coin `protobuf:"bytes,10,rep,name=withdraw,proto3" json:"withdraw"`
	CountWithdraw     uint64       `protobuf:"varint,11,opt,name=countWithdraw,proto3" json:"countWithdraw,omitempty"`
	RefReward         []types.Coin `protobuf:"bytes,12,rep,name=refReward,proto3" json:"refReward"`
	CountRefReward    uint64       `protobuf:"varint,13,opt,name=countRefReward,proto3" json:"countRefReward,omitempty"`
	SysRefReward      []types.Coin `protobuf:"bytes,16,rep,name=sysRefReward,proto3" json:"sysRefReward"`
	CountSysRefReward uint64       `protobuf:"varint,17,opt,name=countSysRefReward,proto3" json:"countSysRefReward,omitempty"`
}

func (m *AssetDailyStats) Reset()         { *m = AssetDailyStats{} }
func (m *AssetDailyStats) String() string { return proto.CompactTextString(m) }
func (*AssetDailyStats) ProtoMessage()    {}
func (*AssetDailyStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_11d69b73d68d8efd, []int{0}
}
func (m *AssetDailyStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AssetDailyStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AssetDailyStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AssetDailyStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetDailyStats.Merge(m, src)
}
func (m *AssetDailyStats) XXX_Size() int {
	return m.Size()
}
func (m *AssetDailyStats) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetDailyStats.DiscardUnknown(m)
}

var xxx_messageInfo_AssetDailyStats proto.InternalMessageInfo

func (m *AssetDailyStats) GetAmountWithFee() []types.Coin {
	if m != nil {
		return m.AmountWithFee
	}
	return nil
}

func (m *AssetDailyStats) GetAmountNoFee() []types.Coin {
	if m != nil {
		return m.AmountNoFee
	}
	return nil
}

func (m *AssetDailyStats) GetFee() []types.Coin {
	if m != nil {
		return m.Fee
	}
	return nil
}

func (m *AssetDailyStats) GetCountWithFee() int32 {
	if m != nil {
		return m.CountWithFee
	}
	return 0
}

func (m *AssetDailyStats) GetCountNoFee() int32 {
	if m != nil {
		return m.CountNoFee
	}
	return 0
}

func (m *AssetDailyStats) GetBurned() []types.Coin {
	if m != nil {
		return m.Burned
	}
	return nil
}

func (m *AssetDailyStats) GetCountBurned() uint64 {
	if m != nil {
		return m.CountBurned
	}
	return 0
}

func (m *AssetDailyStats) GetIssued() []types.Coin {
	if m != nil {
		return m.Issued
	}
	return nil
}

func (m *AssetDailyStats) GetCountIssued() uint64 {
	if m != nil {
		return m.CountIssued
	}
	return 0
}

func (m *AssetDailyStats) GetWithdraw() []types.Coin {
	if m != nil {
		return m.Withdraw
	}
	return nil
}

func (m *AssetDailyStats) GetCountWithdraw() uint64 {
	if m != nil {
		return m.CountWithdraw
	}
	return 0
}

func (m *AssetDailyStats) GetRefReward() []types.Coin {
	if m != nil {
		return m.RefReward
	}
	return nil
}

func (m *AssetDailyStats) GetCountRefReward() uint64 {
	if m != nil {
		return m.CountRefReward
	}
	return 0
}

func (m *AssetDailyStats) GetSysRefReward() []types.Coin {
	if m != nil {
		return m.SysRefReward
	}
	return nil
}

func (m *AssetDailyStats) GetCountSysRefReward() uint64 {
	if m != nil {
		return m.CountSysRefReward
	}
	return 0
}

func init() {
	proto.RegisterType((*AssetDailyStats)(nil), "stwartchain.stats.AssetDailyStats")
}

func init() {
	proto.RegisterFile("stwartchain/stats/asset_stats.proto", fileDescriptor_11d69b73d68d8efd)
}

var fileDescriptor_11d69b73d68d8efd = []byte{
	// 441 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x31, 0x8f, 0xd3, 0x30,
	0x1c, 0xc5, 0x13, 0xda, 0x0b, 0x77, 0xff, 0xb4, 0x40, 0x2d, 0x86, 0x70, 0x83, 0x89, 0x0e, 0x84,
	0x3a, 0x40, 0xac, 0xc2, 0xc0, 0x80, 0x18, 0xda, 0x03, 0x24, 0x16, 0x86, 0xdc, 0x80, 0x04, 0x03,
	0x72, 0x12, 0x5f, 0x63, 0x68, 0xe3, 0x53, 0xec, 0x12, 0xfa, 0x01, 0xd8, 0xf9, 0x58, 0x37, 0x76,
	0x64, 0x42, 0xa8, 0xfd, 0x22, 0xc8, 0x76, 0x95, 0x26, 0xdc, 0x92, 0x2d, 0x7a, 0xff, 0xf7, 0x7e,
	0xcf, 0x8a, 0xf4, 0xe0, 0x91, 0x54, 0x15, 0x2d, 0x55, 0x9a, 0x53, 0x5e, 0x10, 0xa9, 0xa8, 0x92,
	0x84, 0x4a, 0xc9, 0xd4, 0x17, 0xf3, 0x1d, 0x5d, 0x95, 0x42, 0x09, 0x34, 0x6a, 0x98, 0x22, 0x73,
	0x38, 0xbd, 0x3f, 0x17, 0x73, 0x61, 0xae, 0x44, 0x7f, 0x59, 0xe3, 0x29, 0x4e, 0x85, 0x5c, 0x0a,
	0x49, 0x12, 0x2a, 0x19, 0xf9, 0x3e, 0x49, 0x98, 0xa2, 0x13, 0x92, 0x0a, 0x5e, 0xd8, 0xfb, 0xd9,
	0x4f, 0x0f, 0xee, 0x4e, 0x35, 0xfe, 0x0d, 0xe5, 0x8b, 0xf5, 0x85, 0x26, 0xa1, 0xb7, 0x30, 0xa4,
	0x4b, 0xb1, 0x2a, 0xd4, 0x47, 0xae, 0xf2, 0x77, 0x8c, 0x05, 0x6e, 0xd8, 0x1b, 0xfb, 0xcf, 0x1f,
	0x44, 0x96, 0x15, 0x69, 0x56, 0xb4, 0x67, 0x45, 0xe7, 0x82, 0x17, 0xb3, 0xfe, 0xf5, 0x9f, 0x87,
	0x4e, 0xdc, 0x4e, 0xa1, 0x29, 0xf8, 0x56, 0xf8, 0x20, 0x34, 0xe4, 0x56, 0x37, 0x48, 0x33, 0x83,
	0x26, 0xd0, 0xbb, 0x64, 0x2c, 0xe8, 0x75, 0x8b, 0x6a, 0x2f, 0x3a, 0x83, 0x41, 0xda, 0x7c, 0x7b,
	0x3f, 0x74, 0xc7, 0x47, 0x71, 0x4b, 0x43, 0x18, 0x20, 0x3d, 0x3c, 0xec, 0xc8, 0x38, 0x1a, 0x0a,
	0x7a, 0x09, 0x5e, 0xb2, 0x2a, 0x0b, 0x96, 0x05, 0x5e, 0xb7, 0xe6, 0xbd, 0x1d, 0x85, 0xe0, 0x1b,
	0xcc, 0xcc, 0xa6, 0x6f, 0x87, 0xee, 0xb8, 0x1f, 0x37, 0x25, 0x8d, 0xe6, 0x52, 0xae, 0x58, 0x16,
	0x1c, 0x77, 0x44, 0x5b, 0x7b, 0x8d, 0x7e, 0x6f, 0xd3, 0x27, 0x0d, 0xb4, 0x95, 0xd0, 0x2b, 0x38,
	0xae, 0xb8, 0xca, 0xb3, 0x92, 0x56, 0x01, 0x74, 0x83, 0xd7, 0x01, 0xf4, 0x18, 0x86, 0xf5, 0x2f,
	0x32, 0x04, 0xdf, 0x14, 0xb4, 0x45, 0xf4, 0x1a, 0x4e, 0x4a, 0x76, 0x19, 0xb3, 0x8a, 0x96, 0x59,
	0x30, 0xe8, 0xd6, 0x71, 0x48, 0xa0, 0x27, 0x70, 0xc7, 0xf0, 0xe2, 0x9a, 0x31, 0x34, 0x2d, 0xff,
	0xa9, 0xe8, 0x1c, 0x06, 0x72, 0x2d, 0x0f, 0xae, 0x7b, 0xdd, 0x9a, 0x5a, 0x21, 0xf4, 0x14, 0x46,
	0x06, 0x7b, 0xd1, 0x24, 0x8d, 0x4c, 0xdf, 0xcd, 0xc3, 0xec, 0xf3, 0xf5, 0x16, 0xbb, 0x9b, 0x2d,
	0x76, 0xff, 0x6e, 0xb1, 0xfb, 0x6b, 0x87, 0x9d, 0xcd, 0x0e, 0x3b, 0xbf, 0x77, 0xd8, 0xf9, 0x34,
	0x9d, 0x73, 0xb5, 0xa0, 0x89, 0x5e, 0xd9, 0x42, 0x6f, 0x2e, 0x52, 0x2c, 0xcd, 0x09, 0xff, 0xca,
	0x05, 0x59, 0xea, 0x89, 0x26, 0x34, 0xfd, 0xc6, 0x8a, 0x8c, 0xd8, 0x45, 0x3e, 0xb3, 0xbb, 0xfd,
	0xb1, 0x5f, 0xae, 0x5a, 0x5f, 0x31, 0x99, 0x78, 0x66, 0x6b, 0x2f, 0xfe, 0x05, 0x00, 0x00, 0xff,
	0xff, 0x15, 0xec, 0x7b, 0x39, 0xdb, 0x03, 0x00, 0x00,
}

func (m *AssetDailyStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AssetDailyStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AssetDailyStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CountSysRefReward != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountSysRefReward))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x88
	}
	if len(m.SysRefReward) > 0 {
		for iNdEx := len(m.SysRefReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SysRefReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1
			i--
			dAtA[i] = 0x82
		}
	}
	if m.CountRefReward != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountRefReward))
		i--
		dAtA[i] = 0x68
	}
	if len(m.RefReward) > 0 {
		for iNdEx := len(m.RefReward) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RefReward[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x62
		}
	}
	if m.CountWithdraw != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountWithdraw))
		i--
		dAtA[i] = 0x58
	}
	if len(m.Withdraw) > 0 {
		for iNdEx := len(m.Withdraw) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Withdraw[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.CountIssued != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountIssued))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Issued) > 0 {
		for iNdEx := len(m.Issued) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Issued[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x42
		}
	}
	if m.CountBurned != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountBurned))
		i--
		dAtA[i] = 0x38
	}
	if len(m.Burned) > 0 {
		for iNdEx := len(m.Burned) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Burned[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if m.CountNoFee != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountNoFee))
		i--
		dAtA[i] = 0x28
	}
	if m.CountWithFee != 0 {
		i = encodeVarintAssetStats(dAtA, i, uint64(m.CountWithFee))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Fee) > 0 {
		for iNdEx := len(m.Fee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Fee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.AmountNoFee) > 0 {
		for iNdEx := len(m.AmountNoFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AmountNoFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.AmountWithFee) > 0 {
		for iNdEx := len(m.AmountWithFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AmountWithFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintAssetStats(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAssetStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovAssetStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AssetDailyStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AmountWithFee) > 0 {
		for _, e := range m.AmountWithFee {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if len(m.AmountNoFee) > 0 {
		for _, e := range m.AmountNoFee {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if len(m.Fee) > 0 {
		for _, e := range m.Fee {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountWithFee != 0 {
		n += 1 + sovAssetStats(uint64(m.CountWithFee))
	}
	if m.CountNoFee != 0 {
		n += 1 + sovAssetStats(uint64(m.CountNoFee))
	}
	if len(m.Burned) > 0 {
		for _, e := range m.Burned {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountBurned != 0 {
		n += 1 + sovAssetStats(uint64(m.CountBurned))
	}
	if len(m.Issued) > 0 {
		for _, e := range m.Issued {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountIssued != 0 {
		n += 1 + sovAssetStats(uint64(m.CountIssued))
	}
	if len(m.Withdraw) > 0 {
		for _, e := range m.Withdraw {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountWithdraw != 0 {
		n += 1 + sovAssetStats(uint64(m.CountWithdraw))
	}
	if len(m.RefReward) > 0 {
		for _, e := range m.RefReward {
			l = e.Size()
			n += 1 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountRefReward != 0 {
		n += 1 + sovAssetStats(uint64(m.CountRefReward))
	}
	if len(m.SysRefReward) > 0 {
		for _, e := range m.SysRefReward {
			l = e.Size()
			n += 2 + l + sovAssetStats(uint64(l))
		}
	}
	if m.CountSysRefReward != 0 {
		n += 2 + sovAssetStats(uint64(m.CountSysRefReward))
	}
	return n
}

func sovAssetStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssetStats(x uint64) (n int) {
	return sovAssetStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AssetDailyStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssetStats
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
			return fmt.Errorf("proto: AssetDailyStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AssetDailyStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountWithFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountWithFee = append(m.AmountWithFee, types.Coin{})
			if err := m.AmountWithFee[len(m.AmountWithFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountNoFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AmountNoFee = append(m.AmountNoFee, types.Coin{})
			if err := m.AmountNoFee[len(m.AmountNoFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = append(m.Fee, types.Coin{})
			if err := m.Fee[len(m.Fee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountWithFee", wireType)
			}
			m.CountWithFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountWithFee |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountNoFee", wireType)
			}
			m.CountNoFee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountNoFee |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Burned", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Burned = append(m.Burned, types.Coin{})
			if err := m.Burned[len(m.Burned)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountBurned", wireType)
			}
			m.CountBurned = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountBurned |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = append(m.Issued, types.Coin{})
			if err := m.Issued[len(m.Issued)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountIssued", wireType)
			}
			m.CountIssued = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountIssued |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdraw", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Withdraw = append(m.Withdraw, types.Coin{})
			if err := m.Withdraw[len(m.Withdraw)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountWithdraw", wireType)
			}
			m.CountWithdraw = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountWithdraw |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RefReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RefReward = append(m.RefReward, types.Coin{})
			if err := m.RefReward[len(m.RefReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountRefReward", wireType)
			}
			m.CountRefReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountRefReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SysRefReward", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
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
				return ErrInvalidLengthAssetStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAssetStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SysRefReward = append(m.SysRefReward, types.Coin{})
			if err := m.SysRefReward[len(m.SysRefReward)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 17:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CountSysRefReward", wireType)
			}
			m.CountSysRefReward = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssetStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CountSysRefReward |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAssetStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssetStats
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
func skipAssetStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAssetStats
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
					return 0, ErrIntOverflowAssetStats
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
					return 0, ErrIntOverflowAssetStats
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
				return 0, ErrInvalidLengthAssetStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAssetStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAssetStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAssetStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAssetStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAssetStats = fmt.Errorf("proto: unexpected end of group")
)
