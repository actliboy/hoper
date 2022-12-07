// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: dict/dict.proto

package dict

import (
	_ "github.com/liov/hoper/server/go/lib/protobuf/utils/patch"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
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

type Dict struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primaryKey"`
	Typ    uint32 `protobuf:"varint,2,opt,name=typ,proto3" json:"typ,omitempty"`
	PId    uint64 `protobuf:"varint,3,opt,name=pId,proto3" json:"pId,omitempty" gorm:"size:20"`
	Key    string `protobuf:"bytes,4,opt,name=key,proto3" json:"key,omitempty" gorm:"index"`
	Val    string `protobuf:"bytes,5,opt,name=val,proto3" json:"val,omitempty" annotation:"值"`
	Seq    uint32 `protobuf:"varint,6,opt,name=seq,proto3" json:"seq,omitempty" annotation:"顺序"`
	CAt    string `protobuf:"bytes,16,opt,name=cAt,proto3" json:"cAt,omitempty" gorm:"type:timestamptz(6);default:now();index"`
	DAt    string `protobuf:"bytes,28,opt,name=dAt,proto3" json:"dAt,omitempty" gorm:"type:timestamptz(6);default:0001-01-01 00:00:00;index"`
	Status uint32 `protobuf:"varint,18,opt,name=status,proto3" json:"status,omitempty" gorm:"type:int2;default:0"`
}

func (x *Dict) Reset() {
	*x = Dict{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dict_dict_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Dict) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Dict) ProtoMessage() {}

func (x *Dict) ProtoReflect() protoreflect.Message {
	mi := &file_dict_dict_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Dict.ProtoReflect.Descriptor instead.
func (*Dict) Descriptor() ([]byte, []int) {
	return file_dict_dict_proto_rawDescGZIP(), []int{0}
}

func (x *Dict) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Dict) GetTyp() uint32 {
	if x != nil {
		return x.Typ
	}
	return 0
}

func (x *Dict) GetPId() uint64 {
	if x != nil {
		return x.PId
	}
	return 0
}

func (x *Dict) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Dict) GetVal() string {
	if x != nil {
		return x.Val
	}
	return ""
}

func (x *Dict) GetSeq() uint32 {
	if x != nil {
		return x.Seq
	}
	return 0
}

func (x *Dict) GetCAt() string {
	if x != nil {
		return x.CAt
	}
	return ""
}

func (x *Dict) GetDAt() string {
	if x != nil {
		return x.DAt
	}
	return ""
}

func (x *Dict) GetStatus() uint32 {
	if x != nil {
		return x.Status
	}
	return 0
}

var File_dict_dict_proto protoreflect.FileDescriptor

var file_dict_dict_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x64, 0x69, 0x63, 0x74, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x64, 0x69, 0x63, 0x74, 0x1a, 0x14, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70,
	0x61, 0x74, 0x63, 0x68, 0x2f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70,
	0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x03,
	0x0a, 0x04, 0x44, 0x69, 0x63, 0x74, 0x12, 0x28, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x42, 0x18, 0xd2, 0xb5, 0x03, 0x14, 0xa2, 0x01, 0x11, 0x67, 0x6f, 0x72, 0x6d, 0x3a,
	0x22, 0x70, 0x72, 0x69, 0x6d, 0x61, 0x72, 0x79, 0x4b, 0x65, 0x79, 0x22, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x2b, 0x0a, 0x03, 0x74, 0x79, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x19, 0xd2,
	0xb5, 0x03, 0x15, 0xa2, 0x01, 0x12, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3a, 0x22, 0xe7, 0xb1, 0xbb, 0xe5, 0x9e, 0x8b, 0x52, 0x03, 0x74, 0x79, 0x70, 0x12, 0x27, 0x0a,
	0x03, 0x70, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x42, 0x15, 0xd2, 0xb5, 0x03, 0x11,
	0xa2, 0x01, 0x0e, 0x67, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x73, 0x69, 0x7a, 0x65, 0x3a, 0x32, 0x30,
	0x22, 0x52, 0x03, 0x70, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x13, 0xd2, 0xb5, 0x03, 0x0f, 0xa2, 0x01, 0x0c, 0x67, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x29, 0x0a,
	0x03, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x17, 0xd2, 0xb5, 0x03, 0x13,
	0xa2, 0x01, 0x10, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x22, 0xe5,
	0x80, 0xbc, 0x22, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x12, 0x2c, 0x0a, 0x03, 0x73, 0x65, 0x71, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x1a, 0xd2, 0xb5, 0x03, 0x16, 0xa2, 0x01, 0x13, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x22, 0xe9, 0xa1, 0xba, 0xe5, 0xba, 0x8f,
	0x22, 0x52, 0x03, 0x73, 0x65, 0x71, 0x12, 0x47, 0x0a, 0x03, 0x63, 0x41, 0x74, 0x18, 0x10, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x35, 0xd2, 0xb5, 0x03, 0x31, 0xa2, 0x01, 0x2e, 0x67, 0x6f, 0x72, 0x6d,
	0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x74, 0x7a, 0x28, 0x36, 0x29, 0x3b, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x6e, 0x6f,
	0x77, 0x28, 0x29, 0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x52, 0x03, 0x63, 0x41, 0x74, 0x12,
	0x55, 0x0a, 0x03, 0x64, 0x41, 0x74, 0x18, 0x1c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x43, 0xd2, 0xb5,
	0x03, 0x3f, 0xa2, 0x01, 0x3c, 0x67, 0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x74, 0x7a, 0x28, 0x36, 0x29, 0x3b, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x30, 0x30, 0x30, 0x31, 0x2d, 0x30, 0x31, 0x2d, 0x30,
	0x31, 0x20, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3a, 0x30, 0x30, 0x3b, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x22, 0x52, 0x03, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x12, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x26, 0xd2, 0xb5, 0x03, 0x1d, 0xa2, 0x01, 0x1a, 0x67,
	0x6f, 0x72, 0x6d, 0x3a, 0x22, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x69, 0x6e, 0x74, 0x32, 0x3b, 0x64,
	0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x3a, 0x30, 0x22, 0x92, 0x41, 0x02, 0x40, 0x01, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x50, 0x0a, 0x17, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x64, 0x69, 0x63,
	0x74, 0x5a, 0x35, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63,
	0x74, 0x6c, 0x69, 0x62, 0x6f, 0x79, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x64, 0x69, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dict_dict_proto_rawDescOnce sync.Once
	file_dict_dict_proto_rawDescData = file_dict_dict_proto_rawDesc
)

func file_dict_dict_proto_rawDescGZIP() []byte {
	file_dict_dict_proto_rawDescOnce.Do(func() {
		file_dict_dict_proto_rawDescData = protoimpl.X.CompressGZIP(file_dict_dict_proto_rawDescData)
	})
	return file_dict_dict_proto_rawDescData
}

var file_dict_dict_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_dict_dict_proto_goTypes = []interface{}{
	(*Dict)(nil), // 0: dict.Dict
}
var file_dict_dict_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dict_dict_proto_init() }
func file_dict_dict_proto_init() {
	if File_dict_dict_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dict_dict_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Dict); i {
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
			RawDescriptor: file_dict_dict_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_dict_dict_proto_goTypes,
		DependencyIndexes: file_dict_dict_proto_depIdxs,
		MessageInfos:      file_dict_dict_proto_msgTypes,
	}.Build()
	File_dict_dict_proto = out.File
	file_dict_dict_proto_rawDesc = nil
	file_dict_dict_proto_goTypes = nil
	file_dict_dict_proto_depIdxs = nil
}
