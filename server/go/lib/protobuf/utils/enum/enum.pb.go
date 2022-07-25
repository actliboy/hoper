// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: utils/enum/enum.proto

package enum

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var file_utils_enum_enum_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         62025,
		Name:          "enum.enum_customtype",
		Tag:           "bytes,62025,opt,name=enum_customtype",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62026,
		Name:          "enum.enum_genvaluemap",
		Tag:           "varint,62026,opt,name=enum_genvaluemap",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62027,
		Name:          "enum.enum_numorder",
		Tag:           "varint,62027,opt,name=enum_numorder",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62028,
		Name:          "enum.enum_jsonmarshal",
		Tag:           "varint,62028,opt,name=enum_jsonmarshal",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62029,
		Name:          "enum.enum_errorcode",
		Tag:           "varint,62029,opt,name=enum_errorcode",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62030,
		Name:          "enum.enum_gqlgen",
		Tag:           "varint,62030,opt,name=enum_gqlgen",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62031,
		Name:          "enum.enum_prefix",
		Tag:           "varint,62031,opt,name=enum_prefix",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62033,
		Name:          "enum.enum_stringer",
		Tag:           "varint,62033,opt,name=enum_stringer",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         62032,
		Name:          "enum.enum_ext_gen",
		Tag:           "varint,62032,opt,name=enum_ext_gen",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.EnumValueOptions)(nil),
		ExtensionType: (*string)(nil),
		Field:         66002,
		Name:          "enum.enumvalue_cn",
		Tag:           "bytes,66002,opt,name=enumvalue_cn",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1001,
		Name:          "enum.enum_gqlgen_all",
		Tag:           "varint,1001,opt,name=enum_gqlgen_all",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1002,
		Name:          "enum.enum_prefix_all",
		Tag:           "varint,1002,opt,name=enum_prefix_all",
		Filename:      "utils/enum/enum.proto",
	},
	{
		ExtendedType:  (*descriptorpb.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         1003,
		Name:          "enum.enum_ext_gen_all",
		Tag:           "varint,1003,opt,name=enum_ext_gen_all",
		Filename:      "utils/enum/enum.proto",
	},
}

// Extension fields to descriptorpb.EnumOptions.
var (
	//自定义类型
	//
	// optional string enum_customtype = 62025;
	E_EnumCustomtype = &file_utils_enum_enum_proto_extTypes[0]
	// optional bool enum_genvaluemap = 62026;
	E_EnumGenvaluemap = &file_utils_enum_enum_proto_extTypes[1]
	//不用手动标序号，= iota
	//
	// optional bool enum_numorder = 62027;
	E_EnumNumorder = &file_utils_enum_enum_proto_extTypes[2]
	//生成JsonMarshal
	//
	// optional bool enum_jsonmarshal = 62028;
	E_EnumJsonmarshal = &file_utils_enum_enum_proto_extTypes[3]
	//是errorcode
	//
	// optional bool enum_errorcode = 62029;
	E_EnumErrorcode = &file_utils_enum_enum_proto_extTypes[4]
	// optional bool enum_gqlgen = 62030;
	E_EnumGqlgen = &file_utils_enum_enum_proto_extTypes[5]
	// optional bool enum_prefix = 62031;
	E_EnumPrefix = &file_utils_enum_enum_proto_extTypes[6]
	// optional bool enum_stringer = 62033;
	E_EnumStringer = &file_utils_enum_enum_proto_extTypes[7]
	// optional bool enum_ext_gen = 62032;
	E_EnumExtGen = &file_utils_enum_enum_proto_extTypes[8]
)

// Extension fields to descriptorpb.EnumValueOptions.
var (
	// optional string enumvalue_cn = 66002;
	E_EnumvalueCn = &file_utils_enum_enum_proto_extTypes[9]
)

// Extension fields to descriptorpb.FileOptions.
var (
	// optional bool enum_gqlgen_all = 1001;
	E_EnumGqlgenAll = &file_utils_enum_enum_proto_extTypes[10]
	// optional bool enum_prefix_all = 1002;
	E_EnumPrefixAll = &file_utils_enum_enum_proto_extTypes[11]
	// optional bool enum_ext_gen_all = 1003;
	E_EnumExtGenAll = &file_utils_enum_enum_proto_extTypes[12]
)

var File_utils_enum_enum_proto protoreflect.FileDescriptor

var file_utils_enum_enum_proto_rawDesc = []byte{
	0x0a, 0x15, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x2f, 0x65, 0x6e, 0x75,
	0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x65, 0x6e, 0x75, 0x6d, 0x1a, 0x20, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a,
	0x47, 0x0a, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0xc9, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x6e, 0x75, 0x6d, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x74, 0x79, 0x70, 0x65, 0x3a, 0x49, 0x0a, 0x10, 0x65, 0x6e, 0x75, 0x6d,
	0x5f, 0x67, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x6d, 0x61, 0x70, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xca, 0xe4, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x47, 0x65, 0x6e, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x6d, 0x61, 0x70, 0x3a, 0x43, 0x0a, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x6e, 0x75, 0x6d, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0xcb, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x75, 0x6d,
	0x4e, 0x75, 0x6d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x3a, 0x49, 0x0a, 0x10, 0x65, 0x6e, 0x75, 0x6d,
	0x5f, 0x6a, 0x73, 0x6f, 0x6e, 0x6d, 0x61, 0x72, 0x73, 0x68, 0x61, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcc, 0xe4, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x4a, 0x73, 0x6f, 0x6e, 0x6d, 0x61, 0x72, 0x73,
	0x68, 0x61, 0x6c, 0x3a, 0x45, 0x0a, 0x0e, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xcd, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x75,
	0x6d, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x63, 0x6f, 0x64, 0x65, 0x3a, 0x3f, 0x0a, 0x0b, 0x65, 0x6e,
	0x75, 0x6d, 0x5f, 0x67, 0x71, 0x6c, 0x67, 0x65, 0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xce, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x47, 0x71, 0x6c, 0x67, 0x65, 0x6e, 0x3a, 0x3f, 0x0a, 0x0b, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75,
	0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcf, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x3a, 0x43, 0x0a, 0x0d,
	0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65, 0x72, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd1, 0xe4, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x65,
	0x72, 0x3a, 0x40, 0x0a, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65,
	0x6e, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0xd0, 0xe4, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x65, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x74,
	0x47, 0x65, 0x6e, 0x3a, 0x46, 0x0a, 0x0c, 0x65, 0x6e, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x5f, 0x63, 0x6e, 0x12, 0x21, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd2, 0x83, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x65, 0x6e, 0x75, 0x6d, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x43, 0x6e, 0x3a, 0x45, 0x0a, 0x0f, 0x65,
	0x6e, 0x75, 0x6d, 0x5f, 0x67, 0x71, 0x6c, 0x67, 0x65, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x1c,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xe9, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x47, 0x71, 0x6c, 0x67, 0x65, 0x6e, 0x41,
	0x6c, 0x6c, 0x3a, 0x45, 0x0a, 0x0f, 0x65, 0x6e, 0x75, 0x6d, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x75, 0x6d,
	0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x41, 0x6c, 0x6c, 0x3a, 0x46, 0x0a, 0x10, 0x65, 0x6e, 0x75,
	0x6d, 0x5f, 0x65, 0x78, 0x74, 0x5f, 0x67, 0x65, 0x6e, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xeb, 0x07, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x65, 0x6e, 0x75, 0x6d, 0x45, 0x78, 0x74, 0x47, 0x65, 0x6e, 0x41, 0x6c,
	0x6c, 0x42, 0x68, 0x0a, 0x1d, 0x78, 0x79, 0x7a, 0x2e, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2e, 0x65, 0x6e,
	0x75, 0x6d, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5a, 0x3b,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x6c, 0x69,
	0x62, 0x6f, 0x79, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x2f, 0x67, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x65, 0x6e, 0x75, 0x6d,
}

var file_utils_enum_enum_proto_goTypes = []interface{}{
	(*descriptorpb.EnumOptions)(nil),      // 0: google.protobuf.EnumOptions
	(*descriptorpb.EnumValueOptions)(nil), // 1: google.protobuf.EnumValueOptions
	(*descriptorpb.FileOptions)(nil),      // 2: google.protobuf.FileOptions
}
var file_utils_enum_enum_proto_depIdxs = []int32{
	0,  // 0: enum.enum_customtype:extendee -> google.protobuf.EnumOptions
	0,  // 1: enum.enum_genvaluemap:extendee -> google.protobuf.EnumOptions
	0,  // 2: enum.enum_numorder:extendee -> google.protobuf.EnumOptions
	0,  // 3: enum.enum_jsonmarshal:extendee -> google.protobuf.EnumOptions
	0,  // 4: enum.enum_errorcode:extendee -> google.protobuf.EnumOptions
	0,  // 5: enum.enum_gqlgen:extendee -> google.protobuf.EnumOptions
	0,  // 6: enum.enum_prefix:extendee -> google.protobuf.EnumOptions
	0,  // 7: enum.enum_stringer:extendee -> google.protobuf.EnumOptions
	0,  // 8: enum.enum_ext_gen:extendee -> google.protobuf.EnumOptions
	1,  // 9: enum.enumvalue_cn:extendee -> google.protobuf.EnumValueOptions
	2,  // 10: enum.enum_gqlgen_all:extendee -> google.protobuf.FileOptions
	2,  // 11: enum.enum_prefix_all:extendee -> google.protobuf.FileOptions
	2,  // 12: enum.enum_ext_gen_all:extendee -> google.protobuf.FileOptions
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	0,  // [0:13] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_utils_enum_enum_proto_init() }
func file_utils_enum_enum_proto_init() {
	if File_utils_enum_enum_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utils_enum_enum_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 13,
			NumServices:   0,
		},
		GoTypes:           file_utils_enum_enum_proto_goTypes,
		DependencyIndexes: file_utils_enum_enum_proto_depIdxs,
		ExtensionInfos:    file_utils_enum_enum_proto_extTypes,
	}.Build()
	File_utils_enum_enum_proto = out.File
	file_utils_enum_enum_proto_rawDesc = nil
	file_utils_enum_enum_proto_goTypes = nil
	file_utils_enum_enum_proto_depIdxs = nil
}
