// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: foo/example.proto

package foo

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BasicTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IntField    int32   `protobuf:"varint,1,opt,name=int_field,json=intField,proto3" json:"int_field,omitempty"`
	LongField   int64   `protobuf:"varint,2,opt,name=long_field,json=longField,proto3" json:"long_field,omitempty"`
	FloatField  float32 `protobuf:"fixed32,3,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	DoubleField float64 `protobuf:"fixed64,4,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	BoolField   bool    `protobuf:"varint,5,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	StringField string  `protobuf:"bytes,6,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	BytesField  []byte  `protobuf:"bytes,7,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
}

func (x *BasicTypes) Reset() {
	*x = BasicTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BasicTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicTypes) ProtoMessage() {}

func (x *BasicTypes) ProtoReflect() protoreflect.Message {
	mi := &file_foo_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicTypes.ProtoReflect.Descriptor instead.
func (*BasicTypes) Descriptor() ([]byte, []int) {
	return file_foo_example_proto_rawDescGZIP(), []int{0}
}

func (x *BasicTypes) GetIntField() int32 {
	if x != nil {
		return x.IntField
	}
	return 0
}

func (x *BasicTypes) GetLongField() int64 {
	if x != nil {
		return x.LongField
	}
	return 0
}

func (x *BasicTypes) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *BasicTypes) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *BasicTypes) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *BasicTypes) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *BasicTypes) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

type AdvancedTypes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Int32Field     int32                  `protobuf:"varint,1,opt,name=int32_field,json=int32Field,proto3" json:"int32_field,omitempty"`
	Int64Field     int64                  `protobuf:"varint,2,opt,name=int64_field,json=int64Field,proto3" json:"int64_field,omitempty"`
	Uint32Field    uint32                 `protobuf:"varint,3,opt,name=uint32_field,json=uint32Field,proto3" json:"uint32_field,omitempty"`
	Uint64Field    uint64                 `protobuf:"varint,4,opt,name=uint64_field,json=uint64Field,proto3" json:"uint64_field,omitempty"`
	Sint32Field    int32                  `protobuf:"zigzag32,5,opt,name=sint32_field,json=sint32Field,proto3" json:"sint32_field,omitempty"`
	Sint64Field    int64                  `protobuf:"zigzag64,6,opt,name=sint64_field,json=sint64Field,proto3" json:"sint64_field,omitempty"`
	Fixed32Field   uint32                 `protobuf:"fixed32,7,opt,name=fixed32_field,json=fixed32Field,proto3" json:"fixed32_field,omitempty"`
	Fixed64Field   uint64                 `protobuf:"fixed64,8,opt,name=fixed64_field,json=fixed64Field,proto3" json:"fixed64_field,omitempty"`
	Sfixed32Field  int32                  `protobuf:"fixed32,9,opt,name=sfixed32_field,json=sfixed32Field,proto3" json:"sfixed32_field,omitempty"`
	Sfixed64Field  int64                  `protobuf:"fixed64,10,opt,name=sfixed64_field,json=sfixed64Field,proto3" json:"sfixed64_field,omitempty"`
	FloatField     float32                `protobuf:"fixed32,11,opt,name=float_field,json=floatField,proto3" json:"float_field,omitempty"`
	DoubleField    float64                `protobuf:"fixed64,12,opt,name=double_field,json=doubleField,proto3" json:"double_field,omitempty"`
	BoolField      bool                   `protobuf:"varint,13,opt,name=bool_field,json=boolField,proto3" json:"bool_field,omitempty"`
	StringField    string                 `protobuf:"bytes,14,opt,name=string_field,json=stringField,proto3" json:"string_field,omitempty"`
	BytesField     []byte                 `protobuf:"bytes,15,opt,name=bytes_field,json=bytesField,proto3" json:"bytes_field,omitempty"`
	TimestampField *timestamppb.Timestamp `protobuf:"bytes,16,opt,name=timestamp_field,json=timestampField,proto3" json:"timestamp_field,omitempty"`
	DurationField  *durationpb.Duration   `protobuf:"bytes,17,opt,name=duration_field,json=durationField,proto3" json:"duration_field,omitempty"`
}

func (x *AdvancedTypes) Reset() {
	*x = AdvancedTypes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_foo_example_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedTypes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedTypes) ProtoMessage() {}

func (x *AdvancedTypes) ProtoReflect() protoreflect.Message {
	mi := &file_foo_example_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedTypes.ProtoReflect.Descriptor instead.
func (*AdvancedTypes) Descriptor() ([]byte, []int) {
	return file_foo_example_proto_rawDescGZIP(), []int{1}
}

func (x *AdvancedTypes) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *AdvancedTypes) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *AdvancedTypes) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *AdvancedTypes) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *AdvancedTypes) GetSint32Field() int32 {
	if x != nil {
		return x.Sint32Field
	}
	return 0
}

func (x *AdvancedTypes) GetSint64Field() int64 {
	if x != nil {
		return x.Sint64Field
	}
	return 0
}

func (x *AdvancedTypes) GetFixed32Field() uint32 {
	if x != nil {
		return x.Fixed32Field
	}
	return 0
}

func (x *AdvancedTypes) GetFixed64Field() uint64 {
	if x != nil {
		return x.Fixed64Field
	}
	return 0
}

func (x *AdvancedTypes) GetSfixed32Field() int32 {
	if x != nil {
		return x.Sfixed32Field
	}
	return 0
}

func (x *AdvancedTypes) GetSfixed64Field() int64 {
	if x != nil {
		return x.Sfixed64Field
	}
	return 0
}

func (x *AdvancedTypes) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *AdvancedTypes) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *AdvancedTypes) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *AdvancedTypes) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *AdvancedTypes) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *AdvancedTypes) GetTimestampField() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampField
	}
	return nil
}

func (x *AdvancedTypes) GetDurationField() *durationpb.Duration {
	if x != nil {
		return x.DurationField
	}
	return nil
}

var File_foo_example_proto protoreflect.FileDescriptor

var file_foo_example_proto_rawDesc = []byte{
	0x0a, 0x11, 0x66, 0x6f, 0x6f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x03, 0x66, 0x6f, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xef, 0x01, 0x0a, 0x0a, 0x42, 0x61,
	0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x69, 0x6e, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f, 0x6e, 0x67, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f,
	0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79,
	0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x22, 0xa3, 0x05, 0x0a, 0x0d,
	0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x11, 0x52, 0x0b, 0x73, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b,
	0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66,
	0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33,
	0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0f, 0x52, 0x0d, 0x73,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46,
	0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f,
	0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x40, 0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x32, 0xc1, 0x02, 0x0a, 0x0e, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x61,
	0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63,
	0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x13, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73,
	0x1a, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x00, 0x28, 0x01, 0x12, 0x3b, 0x0a, 0x13, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x66,
	0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a, 0x0f, 0x2e,
	0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x73, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x44, 0x0a, 0x1a, 0x42, 0x69, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x12, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65,
	0x73, 0x1a, 0x0f, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x3b, 0x0a, 0x0f, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x12, 0x2e, 0x66, 0x6f,
	0x6f, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x73, 0x1a,
	0x12, 0x2e, 0x66, 0x6f, 0x6f, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x22, 0x00, 0x42, 0x7c, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6f, 0x6f,
	0x42, 0x0c, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e,
	0x65, 0x79, 0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x6b,
	0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x62, 0x2f, 0x66, 0x6f, 0x6f, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa,
	0x02, 0x03, 0x46, 0x6f, 0x6f, 0xca, 0x02, 0x03, 0x46, 0x6f, 0x6f, 0xe2, 0x02, 0x0f, 0x46, 0x6f,
	0x6f, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x03,
	0x46, 0x6f, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_foo_example_proto_rawDescOnce sync.Once
	file_foo_example_proto_rawDescData = file_foo_example_proto_rawDesc
)

func file_foo_example_proto_rawDescGZIP() []byte {
	file_foo_example_proto_rawDescOnce.Do(func() {
		file_foo_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_foo_example_proto_rawDescData)
	})
	return file_foo_example_proto_rawDescData
}

var file_foo_example_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_foo_example_proto_goTypes = []interface{}{
	(*BasicTypes)(nil),            // 0: foo.BasicTypes
	(*AdvancedTypes)(nil),         // 1: foo.AdvancedTypes
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 3: google.protobuf.Duration
}
var file_foo_example_proto_depIdxs = []int32{
	2, // 0: foo.AdvancedTypes.timestamp_field:type_name -> google.protobuf.Timestamp
	3, // 1: foo.AdvancedTypes.duration_field:type_name -> google.protobuf.Duration
	0, // 2: foo.ExampleService.UnaryExample:input_type -> foo.BasicTypes
	0, // 3: foo.ExampleService.ClientStreamExample:input_type -> foo.BasicTypes
	0, // 4: foo.ExampleService.ServerStreamExample:input_type -> foo.BasicTypes
	0, // 5: foo.ExampleService.BidirectionalStreamExample:input_type -> foo.BasicTypes
	1, // 6: foo.ExampleService.AdvancedExample:input_type -> foo.AdvancedTypes
	0, // 7: foo.ExampleService.UnaryExample:output_type -> foo.BasicTypes
	0, // 8: foo.ExampleService.ClientStreamExample:output_type -> foo.BasicTypes
	0, // 9: foo.ExampleService.ServerStreamExample:output_type -> foo.BasicTypes
	0, // 10: foo.ExampleService.BidirectionalStreamExample:output_type -> foo.BasicTypes
	1, // 11: foo.ExampleService.AdvancedExample:output_type -> foo.AdvancedTypes
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_foo_example_proto_init() }
func file_foo_example_proto_init() {
	if File_foo_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_foo_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BasicTypes); i {
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
		file_foo_example_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedTypes); i {
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
			RawDescriptor: file_foo_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_foo_example_proto_goTypes,
		DependencyIndexes: file_foo_example_proto_depIdxs,
		MessageInfos:      file_foo_example_proto_msgTypes,
	}.Build()
	File_foo_example_proto = out.File
	file_foo_example_proto_rawDesc = nil
	file_foo_example_proto_goTypes = nil
	file_foo_example_proto_depIdxs = nil
}
