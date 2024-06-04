// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package faucet

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_Tokens         protoreflect.MessageDescriptor
	fd_Tokens_id      protoreflect.FieldDescriptor
	fd_Tokens_denom   protoreflect.FieldDescriptor
	fd_Tokens_amount  protoreflect.FieldDescriptor
	fd_Tokens_creator protoreflect.FieldDescriptor
)

func init() {
	file_stwartchain_faucet_tokens_proto_init()
	md_Tokens = File_stwartchain_faucet_tokens_proto.Messages().ByName("Tokens")
	fd_Tokens_id = md_Tokens.Fields().ByName("id")
	fd_Tokens_denom = md_Tokens.Fields().ByName("denom")
	fd_Tokens_amount = md_Tokens.Fields().ByName("amount")
	fd_Tokens_creator = md_Tokens.Fields().ByName("creator")
}

var _ protoreflect.Message = (*fastReflection_Tokens)(nil)

type fastReflection_Tokens Tokens

func (x *Tokens) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Tokens)(x)
}

func (x *Tokens) slowProtoReflect() protoreflect.Message {
	mi := &file_stwartchain_faucet_tokens_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Tokens_messageType fastReflection_Tokens_messageType
var _ protoreflect.MessageType = fastReflection_Tokens_messageType{}

type fastReflection_Tokens_messageType struct{}

func (x fastReflection_Tokens_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Tokens)(nil)
}
func (x fastReflection_Tokens_messageType) New() protoreflect.Message {
	return new(fastReflection_Tokens)
}
func (x fastReflection_Tokens_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Tokens
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Tokens) Descriptor() protoreflect.MessageDescriptor {
	return md_Tokens
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Tokens) Type() protoreflect.MessageType {
	return _fastReflection_Tokens_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Tokens) New() protoreflect.Message {
	return new(fastReflection_Tokens)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Tokens) Interface() protoreflect.ProtoMessage {
	return (*Tokens)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Tokens) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Tokens_id, value) {
			return
		}
	}
	if x.Denom != "" {
		value := protoreflect.ValueOfString(x.Denom)
		if !f(fd_Tokens_denom, value) {
			return
		}
	}
	if x.Amount != "" {
		value := protoreflect.ValueOfString(x.Amount)
		if !f(fd_Tokens_amount, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Tokens_creator, value) {
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
func (x *fastReflection_Tokens) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "stwartchain.faucet.Tokens.id":
		return x.Id != uint64(0)
	case "stwartchain.faucet.Tokens.denom":
		return x.Denom != ""
	case "stwartchain.faucet.Tokens.amount":
		return x.Amount != ""
	case "stwartchain.faucet.Tokens.creator":
		return x.Creator != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tokens) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "stwartchain.faucet.Tokens.id":
		x.Id = uint64(0)
	case "stwartchain.faucet.Tokens.denom":
		x.Denom = ""
	case "stwartchain.faucet.Tokens.amount":
		x.Amount = ""
	case "stwartchain.faucet.Tokens.creator":
		x.Creator = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Tokens) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "stwartchain.faucet.Tokens.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "stwartchain.faucet.Tokens.denom":
		value := x.Denom
		return protoreflect.ValueOfString(value)
	case "stwartchain.faucet.Tokens.amount":
		value := x.Amount
		return protoreflect.ValueOfString(value)
	case "stwartchain.faucet.Tokens.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Tokens) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "stwartchain.faucet.Tokens.id":
		x.Id = value.Uint()
	case "stwartchain.faucet.Tokens.denom":
		x.Denom = value.Interface().(string)
	case "stwartchain.faucet.Tokens.amount":
		x.Amount = value.Interface().(string)
	case "stwartchain.faucet.Tokens.creator":
		x.Creator = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Tokens) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "stwartchain.faucet.Tokens.id":
		panic(fmt.Errorf("field id of message stwartchain.faucet.Tokens is not mutable"))
	case "stwartchain.faucet.Tokens.denom":
		panic(fmt.Errorf("field denom of message stwartchain.faucet.Tokens is not mutable"))
	case "stwartchain.faucet.Tokens.amount":
		panic(fmt.Errorf("field amount of message stwartchain.faucet.Tokens is not mutable"))
	case "stwartchain.faucet.Tokens.creator":
		panic(fmt.Errorf("field creator of message stwartchain.faucet.Tokens is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Tokens) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "stwartchain.faucet.Tokens.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "stwartchain.faucet.Tokens.denom":
		return protoreflect.ValueOfString("")
	case "stwartchain.faucet.Tokens.amount":
		return protoreflect.ValueOfString("")
	case "stwartchain.faucet.Tokens.creator":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: stwartchain.faucet.Tokens"))
		}
		panic(fmt.Errorf("message stwartchain.faucet.Tokens does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Tokens) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in stwartchain.faucet.Tokens", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Tokens) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Tokens) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Tokens) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Tokens) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Tokens)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Denom)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Amount)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
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
		x := input.Message.Interface().(*Tokens)
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
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Amount) > 0 {
			i -= len(x.Amount)
			copy(dAtA[i:], x.Amount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Amount)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Denom) > 0 {
			i -= len(x.Denom)
			copy(dAtA[i:], x.Denom)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Denom)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*Tokens)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tokens: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Tokens: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
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
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
				x.Amount = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
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
// source: stwartchain/faucet/tokens.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Tokens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Denom   string `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Amount  string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Tokens) Reset() {
	*x = Tokens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stwartchain_faucet_tokens_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tokens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tokens) ProtoMessage() {}

// Deprecated: Use Tokens.ProtoReflect.Descriptor instead.
func (*Tokens) Descriptor() ([]byte, []int) {
	return file_stwartchain_faucet_tokens_proto_rawDescGZIP(), []int{0}
}

func (x *Tokens) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Tokens) GetDenom() string {
	if x != nil {
		return x.Denom
	}
	return ""
}

func (x *Tokens) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Tokens) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_stwartchain_faucet_tokens_proto protoreflect.FileDescriptor

var file_stwartchain_faucet_tokens_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f, 0x66, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x2f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x12, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66,
	0x61, 0x75, 0x63, 0x65, 0x74, 0x22, 0x60, 0x0a, 0x06, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0xb3, 0x01, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e,
	0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x66, 0x61, 0x75, 0x63,
	0x65, 0x74, 0x42, 0x0b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x23, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x73, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2f,
	0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0xa2, 0x02, 0x03, 0x53, 0x46, 0x58, 0xaa, 0x02, 0x12, 0x53,
	0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x2e, 0x46, 0x61, 0x75, 0x63, 0x65,
	0x74, 0xca, 0x02, 0x12, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5c,
	0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0xe2, 0x02, 0x1e, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74, 0x63,
	0x68, 0x61, 0x69, 0x6e, 0x5c, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x13, 0x53, 0x74, 0x77, 0x61, 0x72, 0x74,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x3a, 0x3a, 0x46, 0x61, 0x75, 0x63, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stwartchain_faucet_tokens_proto_rawDescOnce sync.Once
	file_stwartchain_faucet_tokens_proto_rawDescData = file_stwartchain_faucet_tokens_proto_rawDesc
)

func file_stwartchain_faucet_tokens_proto_rawDescGZIP() []byte {
	file_stwartchain_faucet_tokens_proto_rawDescOnce.Do(func() {
		file_stwartchain_faucet_tokens_proto_rawDescData = protoimpl.X.CompressGZIP(file_stwartchain_faucet_tokens_proto_rawDescData)
	})
	return file_stwartchain_faucet_tokens_proto_rawDescData
}

var file_stwartchain_faucet_tokens_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_stwartchain_faucet_tokens_proto_goTypes = []interface{}{
	(*Tokens)(nil), // 0: stwartchain.faucet.Tokens
}
var file_stwartchain_faucet_tokens_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_stwartchain_faucet_tokens_proto_init() }
func file_stwartchain_faucet_tokens_proto_init() {
	if File_stwartchain_faucet_tokens_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stwartchain_faucet_tokens_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tokens); i {
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
			RawDescriptor: file_stwartchain_faucet_tokens_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_stwartchain_faucet_tokens_proto_goTypes,
		DependencyIndexes: file_stwartchain_faucet_tokens_proto_depIdxs,
		MessageInfos:      file_stwartchain_faucet_tokens_proto_msgTypes,
	}.Build()
	File_stwartchain_faucet_tokens_proto = out.File
	file_stwartchain_faucet_tokens_proto_rawDesc = nil
	file_stwartchain_faucet_tokens_proto_goTypes = nil
	file_stwartchain_faucet_tokens_proto_depIdxs = nil
}
