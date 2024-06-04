// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/stats/fee_stats.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

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

type FeeStats struct {
	Date  string         `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Index string         `protobuf:"bytes,2,opt,name=index,proto3" json:"index,omitempty"`
	Stats *FeeDailyStats `protobuf:"bytes,3,opt,name=stats,proto3" json:"stats,omitempty"`
}

func (m *FeeStats) Reset()         { *m = FeeStats{} }
func (m *FeeStats) String() string { return proto.CompactTextString(m) }
func (*FeeStats) ProtoMessage()    {}
func (*FeeStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_e52fe5b3b97d6924, []int{0}
}
func (m *FeeStats) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *FeeStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_FeeStats.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *FeeStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeeStats.Merge(m, src)
}
func (m *FeeStats) XXX_Size() int {
	return m.Size()
}
func (m *FeeStats) XXX_DiscardUnknown() {
	xxx_messageInfo_FeeStats.DiscardUnknown(m)
}

var xxx_messageInfo_FeeStats proto.InternalMessageInfo

func (m *FeeStats) GetDate() string {
	if m != nil {
		return m.Date
	}
	return ""
}

func (m *FeeStats) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *FeeStats) GetStats() *FeeDailyStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

func init() {
	proto.RegisterType((*FeeStats)(nil), "stwartchain.stats.FeeStats")
}

func init() { proto.RegisterFile("stwartchain/stats/fee_stats.proto", fileDescriptor_e52fe5b3b97d6924) }

var fileDescriptor_e52fe5b3b97d6924 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2c, 0x2e, 0x29, 0x4f,
	0x2c, 0x2a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x2e, 0x49, 0x2c, 0x29, 0xd6, 0x4f, 0x4b,
	0x4d, 0x8d, 0x07, 0xb3, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x91, 0x94, 0xe8, 0x81,
	0x25, 0xa4, 0xd4, 0xb1, 0xeb, 0x4a, 0x49, 0xcc, 0xcc, 0xa9, 0x44, 0xd6, 0xab, 0x94, 0xc3, 0xc5,
	0xe1, 0x96, 0x9a, 0x1a, 0x0c, 0x12, 0x11, 0x12, 0xe2, 0x62, 0x49, 0x49, 0x2c, 0x49, 0x95, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x85, 0x44, 0xb8, 0x58, 0x33, 0xf3, 0x52, 0x52, 0x2b,
	0x24, 0x98, 0xc0, 0x82, 0x10, 0x8e, 0x90, 0x19, 0x17, 0x2b, 0xd8, 0x10, 0x09, 0x66, 0x05, 0x46,
	0x0d, 0x6e, 0x23, 0x05, 0x3d, 0x0c, 0x17, 0xe8, 0xb9, 0xa5, 0xa6, 0xba, 0x80, 0x6c, 0x03, 0x1b,
	0x1d, 0x04, 0x51, 0xee, 0x14, 0x7d, 0xe2, 0x91, 0x1c, 0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e,
	0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51,
	0x8e, 0xe9, 0x99, 0x25, 0x39, 0x89, 0x49, 0x20, 0xcd, 0x39, 0x20, 0xa3, 0xf4, 0x4a, 0x52, 0x93,
	0x33, 0xf4, 0x33, 0xb3, 0x32, 0xf3, 0xf5, 0x73, 0x41, 0x7e, 0x48, 0x4a, 0x4c, 0xce, 0x4e, 0xcd,
	0x4b, 0xd1, 0x87, 0x58, 0xa4, 0x0b, 0xf1, 0x58, 0x05, 0xd4, 0x6b, 0x25, 0x95, 0x05, 0xa9, 0xc5,
	0x49, 0x6c, 0x60, 0x1f, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb0, 0xcb, 0xb3, 0xcb, 0x32,
	0x01, 0x00, 0x00,
}

func (m *FeeStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *FeeStats) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *FeeStats) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Stats != nil {
		{
			size, err := m.Stats.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintFeeStats(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintFeeStats(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Date) > 0 {
		i -= len(m.Date)
		copy(dAtA[i:], m.Date)
		i = encodeVarintFeeStats(dAtA, i, uint64(len(m.Date)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintFeeStats(dAtA []byte, offset int, v uint64) int {
	offset -= sovFeeStats(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *FeeStats) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Date)
	if l > 0 {
		n += 1 + l + sovFeeStats(uint64(l))
	}
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovFeeStats(uint64(l))
	}
	if m.Stats != nil {
		l = m.Stats.Size()
		n += 1 + l + sovFeeStats(uint64(l))
	}
	return n
}

func sovFeeStats(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFeeStats(x uint64) (n int) {
	return sovFeeStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *FeeStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFeeStats
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
			return fmt.Errorf("proto: FeeStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: FeeStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Date", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeStats
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
				return ErrInvalidLengthFeeStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Date = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeStats
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
				return ErrInvalidLengthFeeStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFeeStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFeeStats
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
				return ErrInvalidLengthFeeStats
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFeeStats
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Stats == nil {
				m.Stats = &FeeDailyStats{}
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFeeStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthFeeStats
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
func skipFeeStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFeeStats
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
					return 0, ErrIntOverflowFeeStats
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
					return 0, ErrIntOverflowFeeStats
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
				return 0, ErrInvalidLengthFeeStats
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupFeeStats
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthFeeStats
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthFeeStats        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFeeStats          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupFeeStats = fmt.Errorf("proto: unexpected end of group")
)