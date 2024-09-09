// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/feature.proto

package feature

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

type InstanceFeature int32

const (
	InstanceFeature_INSTANCE_FEATURE_UNSPECIFIED       InstanceFeature = 0
	InstanceFeature_INSTANCE_FEATURE_LOGIN_DEFAULT_ORG InstanceFeature = 1
)

// Enum value maps for InstanceFeature.
var (
	InstanceFeature_name = map[int32]string{
		0: "INSTANCE_FEATURE_UNSPECIFIED",
		1: "INSTANCE_FEATURE_LOGIN_DEFAULT_ORG",
	}
	InstanceFeature_value = map[string]int32{
		"INSTANCE_FEATURE_UNSPECIFIED":       0,
		"INSTANCE_FEATURE_LOGIN_DEFAULT_ORG": 1,
	}
)

func (x InstanceFeature) Enum() *InstanceFeature {
	p := new(InstanceFeature)
	*p = x
	return p
}

func (x InstanceFeature) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (InstanceFeature) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_feature_proto_enumTypes[0].Descriptor()
}

func (InstanceFeature) Type() protoreflect.EnumType {
	return &file_zitadel_feature_proto_enumTypes[0]
}

func (x InstanceFeature) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use InstanceFeature.Descriptor instead.
func (InstanceFeature) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_feature_proto_rawDescGZIP(), []int{0}
}

var File_zitadel_feature_proto protoreflect.FileDescriptor

var file_zitadel_feature_proto_rawDesc = []byte{
	0x0a, 0x15, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x2a, 0x5b, 0x0a, 0x0f, 0x49,
	0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x20,
	0x0a, 0x1c, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x46, 0x45, 0x41, 0x54, 0x55,
	0x52, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00,
	0x12, 0x26, 0x0a, 0x22, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x5f, 0x46, 0x45, 0x41,
	0x54, 0x55, 0x52, 0x45, 0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x44, 0x45, 0x46, 0x41, 0x55,
	0x4c, 0x54, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x01, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a,
	0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_feature_proto_rawDescOnce sync.Once
	file_zitadel_feature_proto_rawDescData = file_zitadel_feature_proto_rawDesc
)

func file_zitadel_feature_proto_rawDescGZIP() []byte {
	file_zitadel_feature_proto_rawDescOnce.Do(func() {
		file_zitadel_feature_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_feature_proto_rawDescData)
	})
	return file_zitadel_feature_proto_rawDescData
}

var file_zitadel_feature_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_feature_proto_goTypes = []interface{}{
	(InstanceFeature)(0), // 0: zitadel.feature.v1.InstanceFeature
}
var file_zitadel_feature_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zitadel_feature_proto_init() }
func file_zitadel_feature_proto_init() {
	if File_zitadel_feature_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_feature_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_feature_proto_goTypes,
		DependencyIndexes: file_zitadel_feature_proto_depIdxs,
		EnumInfos:         file_zitadel_feature_proto_enumTypes,
	}.Build()
	File_zitadel_feature_proto = out.File
	file_zitadel_feature_proto_rawDesc = nil
	file_zitadel_feature_proto_goTypes = nil
	file_zitadel_feature_proto_depIdxs = nil
}