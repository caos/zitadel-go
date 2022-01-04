// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: zitadel/options.proto

package authoption

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AuthOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permission     string `protobuf:"bytes,1,opt,name=permission,proto3" json:"permission,omitempty"`
	CheckFieldName string `protobuf:"bytes,2,opt,name=check_field_name,json=checkFieldName,proto3" json:"check_field_name,omitempty"`
	Feature        string `protobuf:"bytes,3,opt,name=feature,proto3" json:"feature,omitempty"`
}

func (x *AuthOption) Reset() {
	*x = AuthOption{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_options_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthOption) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthOption) ProtoMessage() {}

func (x *AuthOption) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_options_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthOption.ProtoReflect.Descriptor instead.
func (*AuthOption) Descriptor() ([]byte, []int) {
	return file_zitadel_options_proto_rawDescGZIP(), []int{0}
}

func (x *AuthOption) GetPermission() string {
	if x != nil {
		return x.Permission
	}
	return ""
}

func (x *AuthOption) GetCheckFieldName() string {
	if x != nil {
		return x.CheckFieldName
	}
	return ""
}

func (x *AuthOption) GetFeature() string {
	if x != nil {
		return x.Feature
	}
	return ""
}

var file_zitadel_options_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*AuthOption)(nil),
		Field:         50000,
		Name:          "zitadel.v1.auth_option",
		Tag:           "bytes,50000,opt,name=auth_option",
		Filename:      "zitadel/options.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional zitadel.v1.AuthOption auth_option = 50000;
	E_AuthOption = &file_zitadel_options_proto_extTypes[0]
)

var File_zitadel_options_proto protoreflect.FileDescriptor

var file_zitadel_options_proto_rawDesc = []byte{
	0x0a, 0x15, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c,
	0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0a, 0x41, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x5f, 0x66, 0x69, 0x65,
	0x6c, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63,
	0x68, 0x65, 0x63, 0x6b, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x3a, 0x59, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd0, 0x86, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16,
	0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68,
	0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x4f, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x42, 0x4a, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x61, 0x6f, 0x73, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_zitadel_options_proto_rawDescOnce sync.Once
	file_zitadel_options_proto_rawDescData = file_zitadel_options_proto_rawDesc
)

func file_zitadel_options_proto_rawDescGZIP() []byte {
	file_zitadel_options_proto_rawDescOnce.Do(func() {
		file_zitadel_options_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_options_proto_rawDescData)
	})
	return file_zitadel_options_proto_rawDescData
}

var file_zitadel_options_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_zitadel_options_proto_goTypes = []interface{}{
	(*AuthOption)(nil),                 // 0: zitadel.v1.AuthOption
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_zitadel_options_proto_depIdxs = []int32{
	1, // 0: zitadel.v1.auth_option:extendee -> google.protobuf.MethodOptions
	0, // 1: zitadel.v1.auth_option:type_name -> zitadel.v1.AuthOption
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zitadel_options_proto_init() }
func file_zitadel_options_proto_init() {
	if File_zitadel_options_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_options_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthOption); i {
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
			RawDescriptor: file_zitadel_options_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_options_proto_goTypes,
		DependencyIndexes: file_zitadel_options_proto_depIdxs,
		MessageInfos:      file_zitadel_options_proto_msgTypes,
		ExtensionInfos:    file_zitadel_options_proto_extTypes,
	}.Build()
	File_zitadel_options_proto = out.File
	file_zitadel_options_proto_rawDesc = nil
	file_zitadel_options_proto_goTypes = nil
	file_zitadel_options_proto_depIdxs = nil
}