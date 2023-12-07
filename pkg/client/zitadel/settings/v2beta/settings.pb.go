// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/settings/v2beta/settings.proto

package settings

import (
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

type ResourceOwnerType int32

const (
	ResourceOwnerType_RESOURCE_OWNER_TYPE_UNSPECIFIED ResourceOwnerType = 0
	ResourceOwnerType_RESOURCE_OWNER_TYPE_INSTANCE    ResourceOwnerType = 1
	ResourceOwnerType_RESOURCE_OWNER_TYPE_ORG         ResourceOwnerType = 2
)

// Enum value maps for ResourceOwnerType.
var (
	ResourceOwnerType_name = map[int32]string{
		0: "RESOURCE_OWNER_TYPE_UNSPECIFIED",
		1: "RESOURCE_OWNER_TYPE_INSTANCE",
		2: "RESOURCE_OWNER_TYPE_ORG",
	}
	ResourceOwnerType_value = map[string]int32{
		"RESOURCE_OWNER_TYPE_UNSPECIFIED": 0,
		"RESOURCE_OWNER_TYPE_INSTANCE":    1,
		"RESOURCE_OWNER_TYPE_ORG":         2,
	}
)

func (x ResourceOwnerType) Enum() *ResourceOwnerType {
	p := new(ResourceOwnerType)
	*p = x
	return p
}

func (x ResourceOwnerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceOwnerType) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_settings_v2beta_settings_proto_enumTypes[0].Descriptor()
}

func (ResourceOwnerType) Type() protoreflect.EnumType {
	return &file_zitadel_settings_v2beta_settings_proto_enumTypes[0]
}

func (x ResourceOwnerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceOwnerType.Descriptor instead.
func (ResourceOwnerType) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_settings_v2beta_settings_proto_rawDescGZIP(), []int{0}
}

var File_zitadel_settings_v2beta_settings_proto protoreflect.FileDescriptor

var file_zitadel_settings_v2beta_settings_proto_rawDesc = []byte{
	0x0a, 0x26, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65,
	0x6c, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70,
	0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0x77, 0x0a, 0x11, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x1f, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52,
	0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x20, 0x0a, 0x1c, 0x52,
	0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x49, 0x4e, 0x53, 0x54, 0x41, 0x4e, 0x43, 0x45, 0x10, 0x01, 0x12, 0x1b, 0x0a,
	0x17, 0x52, 0x45, 0x53, 0x4f, 0x55, 0x52, 0x43, 0x45, 0x5f, 0x4f, 0x57, 0x4e, 0x45, 0x52, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x4f, 0x52, 0x47, 0x10, 0x02, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x32, 0x62, 0x65, 0x74,
	0x61, 0x3b, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_zitadel_settings_v2beta_settings_proto_rawDescOnce sync.Once
	file_zitadel_settings_v2beta_settings_proto_rawDescData = file_zitadel_settings_v2beta_settings_proto_rawDesc
)

func file_zitadel_settings_v2beta_settings_proto_rawDescGZIP() []byte {
	file_zitadel_settings_v2beta_settings_proto_rawDescOnce.Do(func() {
		file_zitadel_settings_v2beta_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_settings_v2beta_settings_proto_rawDescData)
	})
	return file_zitadel_settings_v2beta_settings_proto_rawDescData
}

var file_zitadel_settings_v2beta_settings_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_settings_v2beta_settings_proto_goTypes = []interface{}{
	(ResourceOwnerType)(0), // 0: zitadel.settings.v2beta.ResourceOwnerType
}
var file_zitadel_settings_v2beta_settings_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zitadel_settings_v2beta_settings_proto_init() }
func file_zitadel_settings_v2beta_settings_proto_init() {
	if File_zitadel_settings_v2beta_settings_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_settings_v2beta_settings_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_settings_v2beta_settings_proto_goTypes,
		DependencyIndexes: file_zitadel_settings_v2beta_settings_proto_depIdxs,
		EnumInfos:         file_zitadel_settings_v2beta_settings_proto_enumTypes,
	}.Build()
	File_zitadel_settings_v2beta_settings_proto = out.File
	file_zitadel_settings_v2beta_settings_proto_rawDesc = nil
	file_zitadel_settings_v2beta_settings_proto_goTypes = nil
	file_zitadel_settings_v2beta_settings_proto_depIdxs = nil
}
