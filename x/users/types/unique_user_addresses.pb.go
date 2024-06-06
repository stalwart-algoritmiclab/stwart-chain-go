// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stwartchain/users/unique_user_addresses.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type UniqueUserAddresses struct {
	Addresses []string `protobuf:"bytes,1,rep,name=addresses,proto3" json:"addresses,omitempty"`
}

func (m *UniqueUserAddresses) Reset()         { *m = UniqueUserAddresses{} }
func (m *UniqueUserAddresses) String() string { return proto.CompactTextString(m) }
func (*UniqueUserAddresses) ProtoMessage()    {}
func (*UniqueUserAddresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_ea4eaa7f387b5972, []int{0}
}
func (m *UniqueUserAddresses) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UniqueUserAddresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UniqueUserAddresses.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UniqueUserAddresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UniqueUserAddresses.Merge(m, src)
}
func (m *UniqueUserAddresses) XXX_Size() int {
	return m.Size()
}
func (m *UniqueUserAddresses) XXX_DiscardUnknown() {
	xxx_messageInfo_UniqueUserAddresses.DiscardUnknown(m)
}

var xxx_messageInfo_UniqueUserAddresses proto.InternalMessageInfo

func (m *UniqueUserAddresses) GetAddresses() []string {
	if m != nil {
		return m.Addresses
	}
	return nil
}

func init() {
	proto.RegisterType((*UniqueUserAddresses)(nil), "stwartchain.users.UniqueUserAddresses")
}

func init() {
	proto.RegisterFile("stwartchain/users/unique_user_addresses.proto", fileDescriptor_ea4eaa7f387b5972)
}

var fileDescriptor_ea4eaa7f387b5972 = []byte{
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x2d, 0x2e, 0x29, 0x4f,
	0x2c, 0x2a, 0x49, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x2f, 0x2d, 0x4e, 0x2d, 0x2a, 0xd6, 0x2f, 0xcd,
	0xcb, 0x2c, 0x2c, 0x4d, 0x8d, 0x07, 0x71, 0xe2, 0x13, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x53,
	0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x91, 0x94, 0xeb, 0x81, 0x95, 0x2b, 0x19,
	0x73, 0x09, 0x87, 0x82, 0x75, 0x84, 0x16, 0xa7, 0x16, 0x39, 0xc2, 0xd4, 0x0b, 0xc9, 0x70, 0x71,
	0xc2, 0x35, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x06, 0x21, 0x04, 0x9c, 0xa2, 0x4f, 0x3c, 0x92,
	0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c,
	0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0xca, 0x31, 0x3d, 0xb3, 0x24, 0x27, 0x31, 0x49, 0xaf,
	0xb8, 0x24, 0x31, 0x07, 0x64, 0x95, 0x5e, 0x49, 0x6a, 0x72, 0x86, 0x7e, 0x66, 0x56, 0x66, 0xbe,
	0x7e, 0x2e, 0xc8, 0x8d, 0x49, 0x89, 0xc9, 0xd9, 0xa9, 0x79, 0x29, 0xfa, 0x10, 0x87, 0xe8, 0x42,
	0x1c, 0x5e, 0x01, 0x75, 0x7a, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0xad, 0xc6, 0x80,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x87, 0x36, 0x23, 0x3d, 0xdc, 0x00, 0x00, 0x00,
}

func (m *UniqueUserAddresses) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UniqueUserAddresses) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UniqueUserAddresses) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for iNdEx := len(m.Addresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Addresses[iNdEx])
			copy(dAtA[i:], m.Addresses[iNdEx])
			i = encodeVarintUniqueUserAddresses(dAtA, i, uint64(len(m.Addresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintUniqueUserAddresses(dAtA []byte, offset int, v uint64) int {
	offset -= sovUniqueUserAddresses(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UniqueUserAddresses) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Addresses) > 0 {
		for _, s := range m.Addresses {
			l = len(s)
			n += 1 + l + sovUniqueUserAddresses(uint64(l))
		}
	}
	return n
}

func sovUniqueUserAddresses(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUniqueUserAddresses(x uint64) (n int) {
	return sovUniqueUserAddresses(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UniqueUserAddresses) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUniqueUserAddresses
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
			return fmt.Errorf("proto: UniqueUserAddresses: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UniqueUserAddresses: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Addresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUniqueUserAddresses
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
				return ErrInvalidLengthUniqueUserAddresses
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUniqueUserAddresses
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Addresses = append(m.Addresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUniqueUserAddresses(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUniqueUserAddresses
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
func skipUniqueUserAddresses(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUniqueUserAddresses
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
					return 0, ErrIntOverflowUniqueUserAddresses
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
					return 0, ErrIntOverflowUniqueUserAddresses
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
				return 0, ErrInvalidLengthUniqueUserAddresses
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUniqueUserAddresses
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUniqueUserAddresses
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUniqueUserAddresses        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUniqueUserAddresses          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUniqueUserAddresses = fmt.Errorf("proto: unexpected end of group")
)
