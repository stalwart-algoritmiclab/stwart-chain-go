// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package rates

import (
	binary "encoding/binary"
	fmt "fmt"
	io "io"
	math "math"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Rates          protoreflect.MessageDescriptor
	fd_Rates_denom    protoreflect.FieldDescriptor
	fd_Rates_rate     protoreflect.FieldDescriptor
	fd_Rates_creator  protoreflect.FieldDescriptor
	fd_Rates_decimals protoreflect.FieldDescriptor
)

func init() {
	file_stwartchain_rates_rates_proto_init()
	md_Rates = File_stwartchain_rates_rates_proto.Messages().ByName("Rates")
	fd_Rates_denom = md_Rates.Fields().ByName("denom")
	fd_Rates_rate = md_Rates.Fields().ByName("rate")
	fd_Rates_creator = md_Rates.Fields().ByName("creator")
	fd_Rates_decimals = md_Rates.Fields().ByName("decimals")
}

var _ protoreflect.Message = (*fastReflection_Rates)(nil)

type fastReflection_Rates Rates

func (x *Rates) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Rates)(x)
}

func (x *Rates) slowProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_rates_rates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Rates_messageType fastReflection_Rates_messageType
var _ protoreflect.MessageType = fastReflection_Rates_messageType{}

type fastReflection_Rates_messageType struct{}

func (x fastReflection_Rates_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Rates)(nil)
}
func (x fastReflection_Rates_messageType) New() protoreflect.Message {
	return new(fastReflection_Rates)
}
func (x fastReflection_Rates_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Rates
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Rates) Descriptor() protoreflect.MessageDescriptor {
	return md_Rates
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Rates) Type() protoreflect.MessageType {
	return _fastReflection_Rates_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Rates) New() protoreflect.Message {
	return new(fastReflection_Rates)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Rates) Interface() protoreflect.ProtoMessage {
	return (*Rates)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Rates) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_Rates_denom, value) {
			return
		}
	}
	if x.Rate != float64(0) || math.Signbit(x.Rate) {
		value := protoreflect.ValueOfFloat64(x.Rate)
		if !f(fd_Rates_rate, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Rates_creator, value) {
			return
		}
	}
	if x.Decimals != int32(0) {
		value := protoreflect.ValueOfInt32(x.Decimals)
		if !f(fd_Rates_decimals, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Rates) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "stwartchain.rates.Rates.denom":
		return x.Denom != ""
	case "stwartchain.rates.Rates.rate":
		return x.Rate != float64(0) || math.Signbit(x.Rate)
	case "stwartchain.rates.Rates.creator":
		return x.Creator != ""
	case "stwartchain.rates.Rates.decimals":
		return x.Decimals != int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rates) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "stwartchain.rates.Rates.denom":
		x.Denom = ""
	case "stwartchain.rates.Rates.rate":
		x.Rate = float64(0)
	case "stwartchain.rates.Rates.creator":
		x.Creator = ""
	case "stwartchain.rates.Rates.decimals":
		x.Decimals = int32(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Rates) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "stwartchain.rates.Rates.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "stwartchain.rates.Rates.rate":
		value := x.Rate
		return protoreflect.ValueOfFloat64(value)
	case "stwartchain.rates.Rates.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	case "stwartchain.rates.Rates.decimals":
		value := x.Decimals
		return protoreflect.ValueOfInt32(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rates) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "stwartchain.rates.Rates.denom":
		x.Denom = value.Interface().(string)
	case "stwartchain.rates.Rates.rate":
		x.Rate = value.Float()
	case "stwartchain.rates.Rates.creator":
		x.Creator = value.Interface().(string)
	case "stwartchain.rates.Rates.decimals":
		x.Decimals = int32(value.Int())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rates) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "stwartchain.rates.Rates.denom":
		panic(fmt.Errorf("field denom of message stwartchain.rates.Rates is not mutable"))
	case "stwartchain.rates.Rates.rate":
		panic(fmt.Errorf("field rate of message stwartchain.rates.Rates is not mutable"))
	case "stwartchain.rates.Rates.creator":
		panic(fmt.Errorf("field creator of message stwartchain.rates.Rates is not mutable"))
	case "stwartchain.rates.Rates.decimals":
		panic(fmt.Errorf("field decimals of message stwartchain.rates.Rates is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Rates) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "stwartchain.rates.Rates.denom":
		return protoreflect.ValueOfString("")
	case "stwartchain.rates.Rates.rate":
		return protoreflect.ValueOfFloat64(float64(0))
	case "stwartchain.rates.Rates.creator":
		return protoreflect.ValueOfString("")
	case "stwartchain.rates.Rates.decimals":
		return protoreflect.ValueOfInt32(int32(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.rates.Rates"))
		}
		panic(fmt.Errorf("message stwartchain.rates.Rates does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Rates) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in stwartchain.rates.Rates", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Rates) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Rates) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Rates) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Rates) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Rates)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Rate != 0 || math.Signbit(x.Rate) {
			n += 9
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Decimals != 0 {
			n += 1 + runtime.Sov(uint64(x.Decimals))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Rates)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.Decimals != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Decimals))
			i--
			dAtA[i] = 0x20
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x1a
		}
		if x.Rate != 0 || math.Signbit(x.Rate) {
			i -= 8
			binary.LittleEndian.PutUint64(dAtA[i:], uint64(math.Float64bits(float64(x.Rate))))
			i--
			dAtA[i] = 0x11
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Rates)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rates: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Rates: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Denom = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 1 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
				}
				var v uint64
				if (iNdEx + 8) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v = uint64(binary.LittleEndian.Uint64(dAtA[iNdEx:]))
				iNdEx += 8
				x.Rate = float64(math.Float64frombits(v))
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Creator = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
				}
				x.Decimals = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Decimals |= int32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: stwartchain/rates/rates.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Rates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Denom    string  `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom,omitempty"`
	Rate     float64 `protobuf:"fixed64,2,opt,name=rate,proto3" json:"rate,omitempty"`
	Creator  string  `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	Decimals int32   `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
}

func (x *Rates) Reset() {
	*x = Rates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_rates_rates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rates) ProtoMessage() {}

// Deprecated: Use Rates.ProtoReflect.Descriptor instead.
func (*Rates) Descriptor() ([]byte, []int) {
	return file_stwartchain_rates_rates_proto_rawDescGZIP(), []int{0}
}

func (x *Rates) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Rates) GetRate() float64 {
	if x != nil {
		return x.Rate
	}
	return 0
}

func (x *Rates) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Rates) GetDecimals() int32 {
	if x != nil {
		return x.Decimals
	}
	return 0
}

var File_stwartchain_rates_rates_proto protoreflect.FileDescriptor

var file_stwartchain_rates_rates_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x72, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x11, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x72, 0x61, 0x74,
	0x65, 0x73, 0x22, 0x67, 0x0a, 0x05, 0x52, 0x61, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x64, 0x65, 0x6e, 0x6f,
	0x6d, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52,
	0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x1a, 0x0a, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x65, 0x63, 0x69, 0x6d, 0x61, 0x6c, 0x73, 0x42, 0xac, 0x01, 0x0a, 0x15,
	0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e,
	0x72, 0x61, 0x74, 0x65, 0x73, 0x42, 0x0a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x22, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x2f, 0x72, 0x61, 0x74, 0x65, 0x73, 0xa2, 0x02, 0x03, 0x53, 0x52, 0x58, 0xaa, 0x02, 0x11,
	0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x52, 0x61, 0x74, 0x65,
	0x73, 0xca, 0x02, 0x11, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c,
	0x52, 0x61, 0x74, 0x65, 0x73, 0xe2, 0x02, 0x1d, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5c, 0x52, 0x61, 0x74, 0x65, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x52, 0x61, 0x74, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_stwartchain_rates_rates_proto_rawDescOnce sync.Once
	file_stwartchain_rates_rates_proto_rawDescData = file_stwartchain_rates_rates_proto_rawDesc
)

func file_stwartchain_rates_rates_proto_rawDescGZIP() []byte {
	file_stwartchain_rates_rates_proto_rawDescOnce.Do(func() {
		file_stwartchain_rates_rates_proto_rawDescData = protoimpl.X.CompressGZIP(file_stwartchain_rates_rates_proto_rawDescData)
	})
	return file_stwartchain_rates_rates_proto_rawDescData
}

var file_stwartchain_rates_rates_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stwartchain_rates_rates_proto_goTypes = []interface{}{
	(*Rates)(nil), // 0: stwartchain.rates.Rates
}
var file_stwartchain_rates_rates_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stwartchain_rates_rates_proto_init() }
func file_stwartchain_rates_rates_proto_init() {
	if File_stwartchain_rates_rates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stwartchain_rates_rates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_stwartchain_rates_rates_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stwartchain_rates_rates_proto_goTypes,
		DependencyIndexes: file_stwartchain_rates_rates_proto_depIdxs,
		MessageInfos:      file_stwartchain_rates_rates_proto_msgTypes,
	}.Build()
	File_stwartchain_rates_rates_proto = out.File
	file_stwartchain_rates_rates_proto_rawDesc = nil
	file_stwartchain_rates_rates_proto_goTypes = nil
	file_stwartchain_rates_rates_proto_depIdxs = nil
}