// 协议

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.1
// source: userinfo.proto

// 包名

package userinfo

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// 手机类型
// 枚举类型第一个字段必须为0
type PhoneType int32

const (
	PhoneType_HOME PhoneType = 0
	PhoneType_WORK PhoneType = 1
)

// Enum value maps for PhoneType.
var (
	PhoneType_name = map[int32]string{
		0: "HOME",
		1: "WORK",
	}
	PhoneType_value = map[string]int32{
		"HOME": 0,
		"WORK": 1,
	}
)

func (x PhoneType) Enum() *PhoneType {
	p := new(PhoneType)
	*p = x
	return p
}

func (x PhoneType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PhoneType) Descriptor() protoreflect.EnumDescriptor {
	return file_userinfo_proto_enumTypes[0].Descriptor()
}

func (PhoneType) Type() protoreflect.EnumType {
	return &file_userinfo_proto_enumTypes[0]
}

func (x PhoneType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PhoneType.Descriptor instead.
func (PhoneType) EnumDescriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

// 手机
type Phone struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type   PhoneType `protobuf:"varint,1,opt,name=type,proto3,enum=userinfo.PhoneType" json:"type,omitempty"`
	Number string    `protobuf:"bytes,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *Phone) Reset() {
	*x = Phone{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Phone) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Phone) ProtoMessage() {}

func (x *Phone) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Phone.ProtoReflect.Descriptor instead.
func (*Phone) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{0}
}

func (x *Phone) GetType() PhoneType {
	if x != nil {
		return x.Type
	}
	return PhoneType_HOME
}

func (x *Phone) GetNumber() string {
	if x != nil {
		return x.Number
	}
	return ""
}

// 人
type Person struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Phones []*Phone `protobuf:"bytes,3,rep,name=phones,proto3" json:"phones,omitempty"`
}

func (x *Person) Reset() {
	*x = Person{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Person) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Person) ProtoMessage() {}

func (x *Person) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Person.ProtoReflect.Descriptor instead.
func (*Person) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{1}
}

func (x *Person) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Person) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Person) GetPhones() []*Phone {
	if x != nil {
		return x.Phones
	}
	return nil
}

// 联系簿
type ContactBook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Persons []*Person `protobuf:"bytes,1,rep,name=persons,proto3" json:"persons,omitempty"`
}

func (x *ContactBook) Reset() {
	*x = ContactBook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userinfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactBook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactBook) ProtoMessage() {}

func (x *ContactBook) ProtoReflect() protoreflect.Message {
	mi := &file_userinfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactBook.ProtoReflect.Descriptor instead.
func (*ContactBook) Descriptor() ([]byte, []int) {
	return file_userinfo_proto_rawDescGZIP(), []int{2}
}

func (x *ContactBook) GetPersons() []*Person {
	if x != nil {
		return x.Persons
	}
	return nil
}

var File_userinfo_proto protoreflect.FileDescriptor

var file_userinfo_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x22, 0x48, 0x0a, 0x05, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x68, 0x6f,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x06, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x27, 0x0a, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x68,
	0x6f, 0x6e, 0x65, 0x52, 0x06, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0b, 0x43,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x2a, 0x0a, 0x07, 0x70, 0x65,
	0x72, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x52, 0x07, 0x70,
	0x65, 0x72, 0x73, 0x6f, 0x6e, 0x73, 0x2a, 0x1f, 0x0a, 0x09, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x4f, 0x4d, 0x45, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x57, 0x4f, 0x52, 0x4b, 0x10, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userinfo_proto_rawDescOnce sync.Once
	file_userinfo_proto_rawDescData = file_userinfo_proto_rawDesc
)

func file_userinfo_proto_rawDescGZIP() []byte {
	file_userinfo_proto_rawDescOnce.Do(func() {
		file_userinfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_userinfo_proto_rawDescData)
	})
	return file_userinfo_proto_rawDescData
}

var file_userinfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_userinfo_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_userinfo_proto_goTypes = []interface{}{
	(PhoneType)(0),      // 0: userinfo.PhoneType
	(*Phone)(nil),       // 1: userinfo.Phone
	(*Person)(nil),      // 2: userinfo.Person
	(*ContactBook)(nil), // 3: userinfo.ContactBook
}
var file_userinfo_proto_depIdxs = []int32{
	0, // 0: userinfo.Phone.type:type_name -> userinfo.PhoneType
	1, // 1: userinfo.Person.phones:type_name -> userinfo.Phone
	2, // 2: userinfo.ContactBook.persons:type_name -> userinfo.Person
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_userinfo_proto_init() }
func file_userinfo_proto_init() {
	if File_userinfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userinfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Phone); i {
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
		file_userinfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Person); i {
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
		file_userinfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactBook); i {
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
			RawDescriptor: file_userinfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_userinfo_proto_goTypes,
		DependencyIndexes: file_userinfo_proto_depIdxs,
		EnumInfos:         file_userinfo_proto_enumTypes,
		MessageInfos:      file_userinfo_proto_msgTypes,
	}.Build()
	File_userinfo_proto = out.File
	file_userinfo_proto_rawDesc = nil
	file_userinfo_proto_goTypes = nil
	file_userinfo_proto_depIdxs = nil
}
