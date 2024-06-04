// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/rates/rates.proto

package types

import (
	encoding_binary "encoding/binary"
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

type Rates struct {
	Denom    string  `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Rate     float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Creator  string  `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Decimals int32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (m *Rates) Reset()         { *m = Rates{} }
func (m *Rates) String() string { return proto.CompactTextString(m) }
func (*Rates) ProtoMessage()    {}
func (*Rates) Descriptor() ([]byte, []int) {
	return fileDescriptor_934b59a15f825390, []int{0}
}
func (m *Rates) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Rates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Rates.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Rates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rates.Merge(m, src)
}
func (m *Rates) XXX_Size() int {
	return m.Size()
}
func (m *Rates) XXX_DiscardUnknown() {
	xxx_messageInfo_Rates.DiscardUnknown(m)
}

var xxx_messageInfo_Rates proto.InternalMessageInfo

func (m *Rates) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *Rates) GetRate() float64 {
	if m != nil {
		return m.Rate
	}
	return 0
}

func (m *Rates) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Rates) GetDecimals() int32 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func init() {
	proto.RegisterType((*Rates)(nil), "stwartchain.rates.Rates")
}

func init() { proto.RegisterFile("stwartchain/rates/rates.proto", fileDescriptor_934b59a15f825390) }

var fileDescriptor_934b59a15f825390 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0x29, 0x4f,
	0x2c, 0x2a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x4a, 0x2c, 0x49, 0x2d, 0x86, 0x90, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x82, 0x48, 0xd2, 0x7a, 0x60, 0x09, 0xa5, 0x74, 0x2e, 0xd6,
	0x20, 0x10, 0x43, 0x48, 0x84, 0x8b, 0x35, 0x25, 0x35, 0x2f, 0x3f, 0x57, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc2, 0x11, 0x12, 0xe2, 0x62, 0x01, 0xa9, 0x93, 0x60, 0x52, 0x60, 0xd4, 0x60,
	0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x93, 0x8b, 0x52, 0x13, 0x4b, 0xf2, 0x8b, 0x24, 0x98,
	0xc1, 0x6a, 0x61, 0x5c, 0x21, 0x29, 0x2e, 0x8e, 0x94, 0xd4, 0xe4, 0xcc, 0xdc, 0xc4, 0x9c, 0x62,
	0x09, 0x16, 0x05, 0x46, 0x0d, 0xd6, 0x20, 0x38, 0xdf, 0x29, 0xfa, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86,
	0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x1c, 0xd3, 0x33, 0x4b, 0x72, 0x12, 0x93, 0xf4, 0x8a, 0x4b, 0x12,
	0x73, 0x40, 0xce, 0xd3, 0x2b, 0x49, 0x4d, 0xce, 0xd0, 0xcf, 0xcc, 0xca, 0xcc, 0xd7, 0xcf, 0x05,
	0xf9, 0x23, 0x29, 0x31, 0x39, 0x3b, 0x35, 0x2f, 0x45, 0x1f, 0xe2, 0x78, 0x5d, 0x88, 0xe7, 0x2a,
	0xa0, 0xde, 0x2b, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xfb, 0xcf, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0x36, 0xda, 0xd4, 0x73, 0x00, 0x01, 0x00, 0x00,
}

func (m *Rates) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Rates) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Rates) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Decimals != 0 {
		i = encodeVarintRates(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRates(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Rate != 0 {
		i -= 8
		encoding_binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(m.Rate))))
		i--
		dAtA[i] = 0x11
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintRates(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRates(dAtA []byte, offset int, v uint64) int {
	offset -= sovRates(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Rates) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovRates(uint64(l))
	}
	if m.Rate != 0 {
		n += 9
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRates(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovRates(uint64(m.Decimals))
	}
	return n
}

func sovRates(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRates(x uint64) (n int) {
	return sovRates(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Rates) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRates
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
			return fmt.Errorf("proto: Rates: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Rates: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRates
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
				return ErrInvalidLengthRates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 1 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
			}
			var v uint64
			if (iNdEx + 8) > l {
				return io.ErrUnexpectedEOF
			}
			v = uint64(encoding_binary.LittleEndian.Uint64(dAtA[iNdEx:]))
			iNdEx += 8
			m.Rate = float64(math.Float64frombits(v))
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRates
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
				return ErrInvalidLengthRates
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRates
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRates
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRates(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRates
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
func skipRates(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRates
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
					return 0, ErrIntOverflowRates
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
					return 0, ErrIntOverflowRates
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
				return 0, ErrInvalidLengthRates
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRates
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRates
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRates        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRates          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRates = fmt.Errorf("proto: unexpected end of group")
)
