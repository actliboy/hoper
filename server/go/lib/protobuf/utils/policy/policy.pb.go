// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.20.1
// source: utils/policy/policy.proto

package policy

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

var file_utils_policy_policy_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         50000,
		Name:          "policy.non_sensitive",
		Tag:           "varint,50000,opt,name=non_sensitive",
		Filename:      "utils/policy/policy.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool non_sensitive = 50000;
	E_NonSensitive = &file_utils_policy_policy_proto_extTypes[0]
)

var File_utils_policy_policy_proto protoreflect.FileDescriptor

var file_utils_policy_policy_proto_rawDesc = []byte{
	0x0a, 0x19, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x6f, 0x6c,
	0x69, 0x63, 0x79, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3a, 0x44, 0x0a, 0x0d, 0x6e, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6e,
	0x6f, 0x6e, 0x53, 0x65, 0x6e, 0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x6c, 0x69, 0x62,
	0x6f, 0x79, 0x2f, 0x68, 0x6f, 0x70, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x67, 0x6f, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x75, 0x74, 0x69, 0x6c, 0x73, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var file_utils_policy_policy_proto_goTypes = []interface{}{
	(*descriptorpb.FieldOptions)(nil), // 0: google.protobuf.FieldOptions
}
var file_utils_policy_policy_proto_depIdxs = []int32{
	0, // 0: policy.non_sensitive:extendee -> google.protobuf.FieldOptions
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_utils_policy_policy_proto_init() }
func file_utils_policy_policy_proto_init() {
	if File_utils_policy_policy_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_utils_policy_policy_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_utils_policy_policy_proto_goTypes,
		DependencyIndexes: file_utils_policy_policy_proto_depIdxs,
		ExtensionInfos:    file_utils_policy_policy_proto_extTypes,
	}.Build()
	File_utils_policy_policy_proto = out.File
	file_utils_policy_policy_proto_rawDesc = nil
	file_utils_policy_policy_proto_goTypes = nil
	file_utils_policy_policy_proto_depIdxs = nil
}
