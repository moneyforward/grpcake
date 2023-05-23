// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: advanced/example.proto

package advanced

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

type AdvancedType struct {
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

func (x *AdvancedType) Reset() {
	*x = AdvancedType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_advanced_example_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvancedType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedType) ProtoMessage() {}

func (x *AdvancedType) ProtoReflect() protoreflect.Message {
	mi := &file_advanced_example_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedType.ProtoReflect.Descriptor instead.
func (*AdvancedType) Descriptor() ([]byte, []int) {
	return file_advanced_example_proto_rawDescGZIP(), []int{0}
}

func (x *AdvancedType) GetInt32Field() int32 {
	if x != nil {
		return x.Int32Field
	}
	return 0
}

func (x *AdvancedType) GetInt64Field() int64 {
	if x != nil {
		return x.Int64Field
	}
	return 0
}

func (x *AdvancedType) GetUint32Field() uint32 {
	if x != nil {
		return x.Uint32Field
	}
	return 0
}

func (x *AdvancedType) GetUint64Field() uint64 {
	if x != nil {
		return x.Uint64Field
	}
	return 0
}

func (x *AdvancedType) GetSint32Field() int32 {
	if x != nil {
		return x.Sint32Field
	}
	return 0
}

func (x *AdvancedType) GetSint64Field() int64 {
	if x != nil {
		return x.Sint64Field
	}
	return 0
}

func (x *AdvancedType) GetFixed32Field() uint32 {
	if x != nil {
		return x.Fixed32Field
	}
	return 0
}

func (x *AdvancedType) GetFixed64Field() uint64 {
	if x != nil {
		return x.Fixed64Field
	}
	return 0
}

func (x *AdvancedType) GetSfixed32Field() int32 {
	if x != nil {
		return x.Sfixed32Field
	}
	return 0
}

func (x *AdvancedType) GetSfixed64Field() int64 {
	if x != nil {
		return x.Sfixed64Field
	}
	return 0
}

func (x *AdvancedType) GetFloatField() float32 {
	if x != nil {
		return x.FloatField
	}
	return 0
}

func (x *AdvancedType) GetDoubleField() float64 {
	if x != nil {
		return x.DoubleField
	}
	return 0
}

func (x *AdvancedType) GetBoolField() bool {
	if x != nil {
		return x.BoolField
	}
	return false
}

func (x *AdvancedType) GetStringField() string {
	if x != nil {
		return x.StringField
	}
	return ""
}

func (x *AdvancedType) GetBytesField() []byte {
	if x != nil {
		return x.BytesField
	}
	return nil
}

func (x *AdvancedType) GetTimestampField() *timestamppb.Timestamp {
	if x != nil {
		return x.TimestampField
	}
	return nil
}

func (x *AdvancedType) GetDurationField() *durationpb.Duration {
	if x != nil {
		return x.DurationField
	}
	return nil
}

var File_advanced_example_proto protoreflect.FileDescriptor

var file_advanced_example_proto_rawDesc = []byte{
	0x0a, 0x16, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63,
	0x65, 0x64, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x05, 0x0a, 0x0c, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x36,
	0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0b, 0x75, 0x69,
	0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x75, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x0b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c,
	0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x11, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x21, 0x0a, 0x0c, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x12, 0x52, 0x0b, 0x73, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x07, 0x52, 0x0c, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x33, 0x32, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x66, 0x69, 0x78, 0x65, 0x64,
	0x36, 0x34, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c,
	0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e,
	0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0f, 0x52, 0x0d, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x33, 0x32, 0x46, 0x69,
	0x65, 0x6c, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x66, 0x69, 0x78, 0x65, 0x64, 0x36, 0x34, 0x5f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x10, 0x52, 0x0d, 0x73, 0x66, 0x69,
	0x78, 0x65, 0x64, 0x36, 0x34, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x6c,
	0x6f, 0x61, 0x74, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x0a, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0c, 0x20, 0x01, 0x28,
	0x01, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x62, 0x6f, 0x6f, 0x6c, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x62, 0x6f, 0x6f, 0x6c, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x21, 0x0a,
	0x0c, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x46, 0x69, 0x65, 0x6c, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18,
	0x0f, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x46, 0x69, 0x65, 0x6c,
	0x64, 0x12, 0x43, 0x0a, 0x0f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32, 0x53, 0x0a, 0x0f, 0x41, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x55,
	0x6e, 0x61, 0x72, 0x79, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x16, 0x2e, 0x61, 0x64,
	0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x1a, 0x16, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x2e, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x54, 0x79, 0x70, 0x65, 0x22, 0x00, 0x42, 0x9a, 0x01,
	0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x42, 0x0c,
	0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x6f, 0x6e, 0x65, 0x79,
	0x66, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x61, 0x6b, 0x65, 0x2f,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67,
	0x2f, 0x70, 0x62, 0x2f, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0xa2, 0x02, 0x03, 0x41,
	0x58, 0x58, 0xaa, 0x02, 0x08, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0xca, 0x02, 0x08,
	0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0xe2, 0x02, 0x14, 0x41, 0x64, 0x76, 0x61, 0x6e,
	0x63, 0x65, 0x64, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea,
	0x02, 0x08, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_advanced_example_proto_rawDescOnce sync.Once
	file_advanced_example_proto_rawDescData = file_advanced_example_proto_rawDesc
)

func file_advanced_example_proto_rawDescGZIP() []byte {
	file_advanced_example_proto_rawDescOnce.Do(func() {
		file_advanced_example_proto_rawDescData = protoimpl.X.CompressGZIP(file_advanced_example_proto_rawDescData)
	})
	return file_advanced_example_proto_rawDescData
}

var file_advanced_example_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_advanced_example_proto_goTypes = []interface{}{
	(*AdvancedType)(nil),          // 0: advanced.AdvancedType
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
	(*durationpb.Duration)(nil),   // 2: google.protobuf.Duration
}
var file_advanced_example_proto_depIdxs = []int32{
	1, // 0: advanced.AdvancedType.timestamp_field:type_name -> google.protobuf.Timestamp
	2, // 1: advanced.AdvancedType.duration_field:type_name -> google.protobuf.Duration
	0, // 2: advanced.AdvancedService.UnaryExample:input_type -> advanced.AdvancedType
	0, // 3: advanced.AdvancedService.UnaryExample:output_type -> advanced.AdvancedType
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_advanced_example_proto_init() }
func file_advanced_example_proto_init() {
	if File_advanced_example_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_advanced_example_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvancedType); i {
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
			RawDescriptor: file_advanced_example_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_advanced_example_proto_goTypes,
		DependencyIndexes: file_advanced_example_proto_depIdxs,
		MessageInfos:      file_advanced_example_proto_msgTypes,
	}.Build()
	File_advanced_example_proto = out.File
	file_advanced_example_proto_rawDesc = nil
	file_advanced_example_proto_goTypes = nil
	file_advanced_example_proto_depIdxs = nil
}
