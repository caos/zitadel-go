// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: zitadel/user/v2beta/user.proto

package user

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Gender int32

const (
	Gender_GENDER_UNSPECIFIED Gender = 0
	Gender_GENDER_FEMALE      Gender = 1
	Gender_GENDER_MALE        Gender = 2
	Gender_GENDER_DIVERSE     Gender = 3
)

// Enum value maps for Gender.
var (
	Gender_name = map[int32]string{
		0: "GENDER_UNSPECIFIED",
		1: "GENDER_FEMALE",
		2: "GENDER_MALE",
		3: "GENDER_DIVERSE",
	}
	Gender_value = map[string]int32{
		"GENDER_UNSPECIFIED": 0,
		"GENDER_FEMALE":      1,
		"GENDER_MALE":        2,
		"GENDER_DIVERSE":     3,
	}
)

func (x Gender) Enum() *Gender {
	p := new(Gender)
	*p = x
	return p
}

func (x Gender) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Gender) Descriptor() protoreflect.EnumDescriptor {
	return file_zitadel_user_v2beta_user_proto_enumTypes[0].Descriptor()
}

func (Gender) Type() protoreflect.EnumType {
	return &file_zitadel_user_v2beta_user_proto_enumTypes[0]
}

func (x Gender) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Gender.Descriptor instead.
func (Gender) EnumDescriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_user_proto_rawDescGZIP(), []int{0}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type SetHumanProfile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GivenName         string  `protobuf:"bytes,1,opt,name=given_name,json=givenName,proto3" json:"given_name,omitempty"`
	FamilyName        string  `protobuf:"bytes,2,opt,name=family_name,json=familyName,proto3" json:"family_name,omitempty"`
	NickName          *string `protobuf:"bytes,3,opt,name=nick_name,json=nickName,proto3,oneof" json:"nick_name,omitempty"`
	DisplayName       *string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3,oneof" json:"display_name,omitempty"`
	PreferredLanguage *string `protobuf:"bytes,5,opt,name=preferred_language,json=preferredLanguage,proto3,oneof" json:"preferred_language,omitempty"`
	Gender            *Gender `protobuf:"varint,6,opt,name=gender,proto3,enum=zitadel.user.v2beta.Gender,oneof" json:"gender,omitempty"`
}

func (x *SetHumanProfile) Reset() {
	*x = SetHumanProfile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetHumanProfile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetHumanProfile) ProtoMessage() {}

func (x *SetHumanProfile) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetHumanProfile.ProtoReflect.Descriptor instead.
func (*SetHumanProfile) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_user_proto_rawDescGZIP(), []int{1}
}

func (x *SetHumanProfile) GetGivenName() string {
	if x != nil {
		return x.GivenName
	}
	return ""
}

func (x *SetHumanProfile) GetFamilyName() string {
	if x != nil {
		return x.FamilyName
	}
	return ""
}

func (x *SetHumanProfile) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *SetHumanProfile) GetDisplayName() string {
	if x != nil && x.DisplayName != nil {
		return *x.DisplayName
	}
	return ""
}

func (x *SetHumanProfile) GetPreferredLanguage() string {
	if x != nil && x.PreferredLanguage != nil {
		return *x.PreferredLanguage
	}
	return ""
}

func (x *SetHumanProfile) GetGender() Gender {
	if x != nil && x.Gender != nil {
		return *x.Gender
	}
	return Gender_GENDER_UNSPECIFIED
}

type SetMetadataEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetMetadataEntry) Reset() {
	*x = SetMetadataEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_zitadel_user_v2beta_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetMetadataEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetMetadataEntry) ProtoMessage() {}

func (x *SetMetadataEntry) ProtoReflect() protoreflect.Message {
	mi := &file_zitadel_user_v2beta_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetMetadataEntry.ProtoReflect.Descriptor instead.
func (*SetMetadataEntry) Descriptor() ([]byte, []int) {
	return file_zitadel_user_v2beta_user_proto_rawDescGZIP(), []int{2}
}

func (x *SetMetadataEntry) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetMetadataEntry) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

var File_zitadel_user_v2beta_user_proto protoreflect.FileDescriptor

var file_zitadel_user_v2beta_user_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76,
	0x32, 0x62, 0x65, 0x74, 0x61, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x16, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xef, 0x03, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x48,
	0x75, 0x6d, 0x61, 0x6e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x3f, 0x0a, 0x0a, 0x67,
	0x69, 0x76, 0x65, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x20, 0x92, 0x41, 0x10, 0x4a, 0x08, 0x22, 0x4d, 0x69, 0x6e, 0x6e, 0x69, 0x65, 0x22, 0x78, 0xc8,
	0x01, 0x80, 0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8,
	0x01, 0x52, 0x09, 0x67, 0x69, 0x76, 0x65, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x40, 0x0a, 0x0b,
	0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1f, 0x92, 0x41, 0x0f, 0x4a, 0x07, 0x22, 0x4d, 0x6f, 0x75, 0x73, 0x65, 0x22, 0x78,
	0xc8, 0x01, 0x80, 0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18,
	0xc8, 0x01, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x6e, 0x69, 0x63, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x16, 0x92, 0x41, 0x0b, 0x4a, 0x06, 0x22, 0x4d, 0x69, 0x6e, 0x69, 0x22, 0x78, 0xc8,
	0x01, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xc8, 0x01, 0x48, 0x00, 0x52, 0x08, 0x6e, 0x69, 0x63,
	0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x46, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1e,
	0x92, 0x41, 0x13, 0x4a, 0x0e, 0x22, 0x4d, 0x69, 0x6e, 0x6e, 0x69, 0x65, 0x20, 0x4d, 0x6f, 0x75,
	0x73, 0x65, 0x22, 0x78, 0xc8, 0x01, 0xfa, 0x42, 0x05, 0x72, 0x03, 0x18, 0xc8, 0x01, 0x48, 0x01,
	0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01,
	0x12, 0x46, 0x0a, 0x12, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x92, 0x41,
	0x08, 0x4a, 0x04, 0x22, 0x65, 0x6e, 0x22, 0x78, 0x0a, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x0a,
	0x48, 0x02, 0x52, 0x11, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x72, 0x65, 0x64, 0x4c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x88, 0x01, 0x01, 0x12, 0x4e, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x7a, 0x69, 0x74, 0x61, 0x64,
	0x65, 0x6c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x32, 0x62, 0x65, 0x74, 0x61, 0x2e, 0x47,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x42, 0x14, 0x92, 0x41, 0x11, 0x4a, 0x0f, 0x22, 0x47, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x22, 0x48, 0x03, 0x52, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x88, 0x01, 0x01, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6e, 0x69, 0x63,
	0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0f, 0x0a, 0x0d, 0x5f, 0x64, 0x69, 0x73, 0x70, 0x6c,
	0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x15, 0x0a, 0x13, 0x5f, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x72, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x42, 0x09,
	0x0a, 0x07, 0x5f, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xbb, 0x01, 0x0a, 0x10, 0x53, 0x65,
	0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x32,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x92, 0x41, 0x10,
	0x4a, 0x08, 0x22, 0x6d, 0x79, 0x2d, 0x6b, 0x65, 0x79, 0x22, 0x78, 0xc8, 0x01, 0x80, 0x01, 0x01,
	0xe0, 0x41, 0x02, 0xfa, 0x42, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x73, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x42, 0x5d, 0x92, 0x41, 0x4c, 0x32, 0x23, 0x54, 0x68, 0x65, 0x20, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x20, 0x68, 0x61, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x62, 0x65, 0x20, 0x62, 0x61, 0x73, 0x65,
	0x36, 0x34, 0x20, 0x65, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x64, 0x2e, 0x4a, 0x1e, 0x22, 0x56, 0x47,
	0x68, 0x70, 0x63, 0x79, 0x42, 0x70, 0x63, 0x79, 0x42, 0x74, 0x65, 0x53, 0x42, 0x30, 0x5a, 0x58,
	0x4e, 0x30, 0x49, 0x48, 0x5a, 0x68, 0x62, 0x48, 0x56, 0x6c, 0x22, 0x78, 0xa0, 0xc2, 0x1e, 0x80,
	0x01, 0x01, 0xe0, 0x41, 0x02, 0xfa, 0x42, 0x08, 0x7a, 0x06, 0x10, 0x01, 0x18, 0xa0, 0xc2, 0x1e,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x2a, 0x58, 0x0a, 0x06, 0x47, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x47, 0x45, 0x4e,
	0x44, 0x45, 0x52, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b,
	0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x12, 0x12, 0x0a,
	0x0e, 0x47, 0x45, 0x4e, 0x44, 0x45, 0x52, 0x5f, 0x44, 0x49, 0x56, 0x45, 0x52, 0x53, 0x45, 0x10,
	0x03, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f, 0x7a, 0x69, 0x74, 0x61, 0x64, 0x65, 0x6c, 0x2f,
	0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x32,
	0x62, 0x65, 0x74, 0x61, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_zitadel_user_v2beta_user_proto_rawDescOnce sync.Once
	file_zitadel_user_v2beta_user_proto_rawDescData = file_zitadel_user_v2beta_user_proto_rawDesc
)

func file_zitadel_user_v2beta_user_proto_rawDescGZIP() []byte {
	file_zitadel_user_v2beta_user_proto_rawDescOnce.Do(func() {
		file_zitadel_user_v2beta_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_zitadel_user_v2beta_user_proto_rawDescData)
	})
	return file_zitadel_user_v2beta_user_proto_rawDescData
}

var file_zitadel_user_v2beta_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zitadel_user_v2beta_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_zitadel_user_v2beta_user_proto_goTypes = []interface{}{
	(Gender)(0),              // 0: zitadel.user.v2beta.Gender
	(*User)(nil),             // 1: zitadel.user.v2beta.User
	(*SetHumanProfile)(nil),  // 2: zitadel.user.v2beta.SetHumanProfile
	(*SetMetadataEntry)(nil), // 3: zitadel.user.v2beta.SetMetadataEntry
}
var file_zitadel_user_v2beta_user_proto_depIdxs = []int32{
	0, // 0: zitadel.user.v2beta.SetHumanProfile.gender:type_name -> zitadel.user.v2beta.Gender
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_zitadel_user_v2beta_user_proto_init() }
func file_zitadel_user_v2beta_user_proto_init() {
	if File_zitadel_user_v2beta_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_zitadel_user_v2beta_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_zitadel_user_v2beta_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetHumanProfile); i {
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
		file_zitadel_user_v2beta_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetMetadataEntry); i {
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
	file_zitadel_user_v2beta_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zitadel_user_v2beta_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zitadel_user_v2beta_user_proto_goTypes,
		DependencyIndexes: file_zitadel_user_v2beta_user_proto_depIdxs,
		EnumInfos:         file_zitadel_user_v2beta_user_proto_enumTypes,
		MessageInfos:      file_zitadel_user_v2beta_user_proto_msgTypes,
	}.Build()
	File_zitadel_user_v2beta_user_proto = out.File
	file_zitadel_user_v2beta_user_proto_rawDesc = nil
	file_zitadel_user_v2beta_user_proto_goTypes = nil
	file_zitadel_user_v2beta_user_proto_depIdxs = nil
}
