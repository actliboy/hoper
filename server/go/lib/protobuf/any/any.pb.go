// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0--rc1
// source: any/any.proto

package any

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RawJson struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	B []byte `protobuf:"bytes,1,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *RawJson) Reset() {
	*x = RawJson{}
	if protoimpl.UnsafeEnabled {
		mi := &file_any_any_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawJson) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawJson) ProtoMessage() {}

func (x *RawJson) ProtoReflect() protoreflect.Message {
	mi := &file_any_any_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawJson.ProtoReflect.Descriptor instead.
func (*RawJson) Descriptor() ([]byte, []int) {
	return file_any_any_proto_rawDescGZIP(), []int{0}
}

func (x *RawJson) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

var File_any_any_proto protoreflect.FileDescriptor

var file_any_any_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x61, 0x6e, 0x79, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x03, 0x61, 0x6e, 0x79, 0x22, 0x17, 0x0a, 0x07, 0x52, 0x61, 0x77, 0x4a, 0x73, 0x6f, 0x6e, 0x12,
	0x0c, 0x0a, 0x01, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x62, 0x42, 0x50, 0x0a,
	0x1c, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x61, 0x6e, 0x79, 0x5a, 0x30, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x6f, 0x76, 0x2f, 0x68,
	0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x6c,
	0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_any_any_proto_rawDescOnce sync.Once
	file_any_any_proto_rawDescData = file_any_any_proto_rawDesc
)

func file_any_any_proto_rawDescGZIP() []byte {
	file_any_any_proto_rawDescOnce.Do(func() {
		file_any_any_proto_rawDescData = protoimpl.X.CompressGZIP(file_any_any_proto_rawDescData)
	})
	return file_any_any_proto_rawDescData
}

var file_any_any_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_any_any_proto_goTypes = []interface{}{
	(*RawJson)(nil), // 0: any.RawJson
}
var file_any_any_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_any_any_proto_init() }
func file_any_any_proto_init() {
	if File_any_any_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_any_any_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawJson); i {
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
			RawDescriptor: file_any_any_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_any_any_proto_goTypes,
		DependencyIndexes: file_any_any_proto_depIdxs,
		MessageInfos:      file_any_any_proto_msgTypes,
	}.Build()
	File_any_any_proto = out.File
	file_any_any_proto_rawDesc = nil
	file_any_any_proto_goTypes = nil
	file_any_any_proto_depIdxs = nil
}
