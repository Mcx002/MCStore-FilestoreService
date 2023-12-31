// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.9.1
// source: auth.proto

package proto_gen

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

type SubjectType int32

const (
	SubjectType_UNKNOWN       SubjectType = 0
	SubjectType_CUSTOMER      SubjectType = 1
	SubjectType_SELLER        SubjectType = 2
	SubjectType_ADMIN         SubjectType = 3
	SubjectType_ANON_CUSTOMER SubjectType = 4
	SubjectType_ANON_SELLER   SubjectType = 5
	SubjectType_ANON_ADMIN    SubjectType = 6
)

// Enum value maps for SubjectType.
var (
	SubjectType_name = map[int32]string{
		0: "UNKNOWN",
		1: "CUSTOMER",
		2: "SELLER",
		3: "ADMIN",
		4: "ANON_CUSTOMER",
		5: "ANON_SELLER",
		6: "ANON_ADMIN",
	}
	SubjectType_value = map[string]int32{
		"UNKNOWN":       0,
		"CUSTOMER":      1,
		"SELLER":        2,
		"ADMIN":         3,
		"ANON_CUSTOMER": 4,
		"ANON_SELLER":   5,
		"ANON_ADMIN":    6,
	}
)

func (x SubjectType) Enum() *SubjectType {
	p := new(SubjectType)
	*p = x
	return p
}

func (x SubjectType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubjectType) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[0].Descriptor()
}

func (SubjectType) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[0]
}

func (x SubjectType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubjectType.Descriptor instead.
func (SubjectType) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

type AnonymousDto_AnonymousLevel int32

const (
	AnonymousDto_UNKNOWN  AnonymousDto_AnonymousLevel = 0
	AnonymousDto_CUSTOMER AnonymousDto_AnonymousLevel = 1
	AnonymousDto_SELLER   AnonymousDto_AnonymousLevel = 2
	AnonymousDto_ADMIN    AnonymousDto_AnonymousLevel = 3
)

// Enum value maps for AnonymousDto_AnonymousLevel.
var (
	AnonymousDto_AnonymousLevel_name = map[int32]string{
		0: "UNKNOWN",
		1: "CUSTOMER",
		2: "SELLER",
		3: "ADMIN",
	}
	AnonymousDto_AnonymousLevel_value = map[string]int32{
		"UNKNOWN":  0,
		"CUSTOMER": 1,
		"SELLER":   2,
		"ADMIN":    3,
	}
)

func (x AnonymousDto_AnonymousLevel) Enum() *AnonymousDto_AnonymousLevel {
	p := new(AnonymousDto_AnonymousLevel)
	*p = x
	return p
}

func (x AnonymousDto_AnonymousLevel) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AnonymousDto_AnonymousLevel) Descriptor() protoreflect.EnumDescriptor {
	return file_auth_proto_enumTypes[1].Descriptor()
}

func (AnonymousDto_AnonymousLevel) Type() protoreflect.EnumType {
	return &file_auth_proto_enumTypes[1]
}

func (x AnonymousDto_AnonymousLevel) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AnonymousDto_AnonymousLevel.Descriptor instead.
func (AnonymousDto_AnonymousLevel) EnumDescriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1, 0}
}

type Subject struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xid          string `protobuf:"bytes,1,opt,name=xid,proto3" json:"xid,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhotoProfile string `protobuf:"bytes,3,opt,name=photo_profile,json=photoProfile,proto3" json:"photo_profile,omitempty"`
	SubjectType  int32  `protobuf:"varint,4,opt,name=subject_type,json=subjectType,proto3" json:"subject_type,omitempty"`
}

func (x *Subject) Reset() {
	*x = Subject{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subject) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subject) ProtoMessage() {}

func (x *Subject) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subject.ProtoReflect.Descriptor instead.
func (*Subject) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{0}
}

func (x *Subject) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *Subject) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Subject) GetPhotoProfile() string {
	if x != nil {
		return x.PhotoProfile
	}
	return ""
}

func (x *Subject) GetSubjectType() int32 {
	if x != nil {
		return x.SubjectType
	}
	return 0
}

type AnonymousDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Xid       string                      `protobuf:"bytes,1,opt,name=xid,proto3" json:"xid,omitempty"`
	Username  string                      `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Password  string                      `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Level     AnonymousDto_AnonymousLevel `protobuf:"varint,4,opt,name=level,proto3,enum=auth.AnonymousDto_AnonymousLevel" json:"level,omitempty"`
	CreatedAt int64                       `protobuf:"varint,5,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	UpdatedAt int64                       `protobuf:"varint,6,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	Version   int32                       `protobuf:"varint,7,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *AnonymousDto) Reset() {
	*x = AnonymousDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AnonymousDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AnonymousDto) ProtoMessage() {}

func (x *AnonymousDto) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AnonymousDto.ProtoReflect.Descriptor instead.
func (*AnonymousDto) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{1}
}

func (x *AnonymousDto) GetXid() string {
	if x != nil {
		return x.Xid
	}
	return ""
}

func (x *AnonymousDto) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *AnonymousDto) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *AnonymousDto) GetLevel() AnonymousDto_AnonymousLevel {
	if x != nil {
		return x.Level
	}
	return AnonymousDto_UNKNOWN
}

func (x *AnonymousDto) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *AnonymousDto) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *AnonymousDto) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

type AuthResultDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *AuthResultDto) Reset() {
	*x = AuthResultDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthResultDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthResultDto) ProtoMessage() {}

func (x *AuthResultDto) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthResultDto.ProtoReflect.Descriptor instead.
func (*AuthResultDto) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{2}
}

func (x *AuthResultDto) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateTokenDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token    string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Audience []string `protobuf:"bytes,2,rep,name=audience,proto3" json:"audience,omitempty"`
}

func (x *ValidateTokenDto) Reset() {
	*x = ValidateTokenDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateTokenDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateTokenDto) ProtoMessage() {}

func (x *ValidateTokenDto) ProtoReflect() protoreflect.Message {
	mi := &file_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateTokenDto.ProtoReflect.Descriptor instead.
func (*ValidateTokenDto) Descriptor() ([]byte, []int) {
	return file_auth_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateTokenDto) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ValidateTokenDto) GetAudience() []string {
	if x != nil {
		return x.Audience
	}
	return nil
}

var File_auth_proto protoreflect.FileDescriptor

var file_auth_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75,
	0x74, 0x68, 0x22, 0x77, 0x0a, 0x07, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x68, 0x6f, 0x74,
	0x6f, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x75, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b,
	0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22, 0xab, 0x02, 0x0a, 0x0c,
	0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x44, 0x74, 0x6f, 0x12, 0x10, 0x0a, 0x03,
	0x78, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x78, 0x69, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61,
	0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x21, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x41, 0x6e, 0x6f,
	0x6e, 0x79, 0x6d, 0x6f, 0x75, 0x73, 0x44, 0x74, 0x6f, 0x2e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d,
	0x6f, 0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x42, 0x0a, 0x0e, 0x41, 0x6e, 0x6f, 0x6e, 0x79, 0x6d, 0x6f,
	0x75, 0x73, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f,
	0x57, 0x4e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52,
	0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09,
	0x0a, 0x05, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x22, 0x25, 0x0a, 0x0d, 0x41, 0x75, 0x74,
	0x68, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x44, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x44, 0x74, 0x6f, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08, 0x61, 0x75,
	0x64, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x2a, 0x73, 0x0a, 0x0b, 0x53, 0x75, 0x62, 0x6a, 0x65, 0x63,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e,
	0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05,
	0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x4e, 0x4f, 0x4e, 0x5f,
	0x43, 0x55, 0x53, 0x54, 0x4f, 0x4d, 0x45, 0x52, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x4e,
	0x4f, 0x4e, 0x5f, 0x53, 0x45, 0x4c, 0x4c, 0x45, 0x52, 0x10, 0x05, 0x12, 0x0e, 0x0a, 0x0a, 0x41,
	0x4e, 0x4f, 0x4e, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x06, 0x42, 0x0f, 0x5a, 0x0d, 0x73,
	0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_proto_rawDescOnce sync.Once
	file_auth_proto_rawDescData = file_auth_proto_rawDesc
)

func file_auth_proto_rawDescGZIP() []byte {
	file_auth_proto_rawDescOnce.Do(func() {
		file_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_proto_rawDescData)
	})
	return file_auth_proto_rawDescData
}

var file_auth_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_auth_proto_goTypes = []interface{}{
	(SubjectType)(0),                 // 0: auth.SubjectType
	(AnonymousDto_AnonymousLevel)(0), // 1: auth.AnonymousDto.AnonymousLevel
	(*Subject)(nil),                  // 2: auth.Subject
	(*AnonymousDto)(nil),             // 3: auth.AnonymousDto
	(*AuthResultDto)(nil),            // 4: auth.AuthResultDto
	(*ValidateTokenDto)(nil),         // 5: auth.ValidateTokenDto
}
var file_auth_proto_depIdxs = []int32{
	1, // 0: auth.AnonymousDto.level:type_name -> auth.AnonymousDto.AnonymousLevel
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auth_proto_init() }
func file_auth_proto_init() {
	if File_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Subject); i {
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
		file_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AnonymousDto); i {
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
		file_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthResultDto); i {
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
		file_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateTokenDto); i {
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
			RawDescriptor: file_auth_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_proto_goTypes,
		DependencyIndexes: file_auth_proto_depIdxs,
		EnumInfos:         file_auth_proto_enumTypes,
		MessageInfos:      file_auth_proto_msgTypes,
	}.Build()
	File_auth_proto = out.File
	file_auth_proto_rawDesc = nil
	file_auth_proto_goTypes = nil
	file_auth_proto_depIdxs = nil
}
