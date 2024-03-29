// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: rusprofile_parser.proto

package api

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type INNCompany struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
}

func (x *INNCompany) Reset() {
	*x = INNCompany{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_parser_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *INNCompany) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*INNCompany) ProtoMessage() {}

func (x *INNCompany) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_parser_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use INNCompany.ProtoReflect.Descriptor instead.
func (*INNCompany) Descriptor() ([]byte, []int) {
	return file_rusprofile_parser_proto_rawDescGZIP(), []int{0}
}

func (x *INNCompany) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

type CompanyInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	INN             string `protobuf:"bytes,1,opt,name=INN,proto3" json:"INN,omitempty"`
	KPP             string `protobuf:"bytes,2,opt,name=KPP,proto3" json:"KPP,omitempty"`
	CompanyName     string `protobuf:"bytes,3,opt,name=CompanyName,proto3" json:"CompanyName,omitempty"`
	FullNameManager string `protobuf:"bytes,4,opt,name=FullNameManager,proto3" json:"FullNameManager,omitempty"`
}

func (x *CompanyInfo) Reset() {
	*x = CompanyInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rusprofile_parser_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompanyInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompanyInfo) ProtoMessage() {}

func (x *CompanyInfo) ProtoReflect() protoreflect.Message {
	mi := &file_rusprofile_parser_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompanyInfo.ProtoReflect.Descriptor instead.
func (*CompanyInfo) Descriptor() ([]byte, []int) {
	return file_rusprofile_parser_proto_rawDescGZIP(), []int{1}
}

func (x *CompanyInfo) GetINN() string {
	if x != nil {
		return x.INN
	}
	return ""
}

func (x *CompanyInfo) GetKPP() string {
	if x != nil {
		return x.KPP
	}
	return ""
}

func (x *CompanyInfo) GetCompanyName() string {
	if x != nil {
		return x.CompanyName
	}
	return ""
}

func (x *CompanyInfo) GetFullNameManager() string {
	if x != nil {
		return x.FullNameManager
	}
	return ""
}

var File_rusprofile_parser_proto protoreflect.FileDescriptor

var file_rusprofile_parser_proto_rawDesc = []byte{
	0x0a, 0x17, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x72, 0x75, 0x73, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x1a, 0x1e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72,
	0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x0a, 0x49, 0x4e,
	0x4e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x22, 0x7d, 0x0a, 0x0b, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x49, 0x4e, 0x4e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x49, 0x4e, 0x4e, 0x12, 0x10, 0x0a, 0x03, 0x4b,
	0x50, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x50, 0x50, 0x12, 0x20, 0x0a,
	0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x28, 0x0a, 0x0f, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67,
	0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x46, 0x75, 0x6c, 0x6c, 0x4e, 0x61,
	0x6d, 0x65, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x32, 0x73, 0x0a, 0x09, 0x43, 0x6c, 0x75,
	0x62, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x66, 0x0a, 0x0a, 0x67, 0x65, 0x74, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x6e, 0x79, 0x12, 0x1d, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x49, 0x4e, 0x4e, 0x43, 0x6f, 0x6d, 0x70,
	0x61, 0x6e, 0x79, 0x1a, 0x1e, 0x2e, 0x72, 0x75, 0x73, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x5f, 0x70, 0x61, 0x72, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x49,
	0x6e, 0x66, 0x6f, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x22, 0x0e, 0x2f, 0x76, 0x31,
	0x2f, 0x69, 0x6e, 0x6e, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x3a, 0x01, 0x2a, 0x42, 0x07,
	0x5a, 0x05, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rusprofile_parser_proto_rawDescOnce sync.Once
	file_rusprofile_parser_proto_rawDescData = file_rusprofile_parser_proto_rawDesc
)

func file_rusprofile_parser_proto_rawDescGZIP() []byte {
	file_rusprofile_parser_proto_rawDescOnce.Do(func() {
		file_rusprofile_parser_proto_rawDescData = protoimpl.X.CompressGZIP(file_rusprofile_parser_proto_rawDescData)
	})
	return file_rusprofile_parser_proto_rawDescData
}

var file_rusprofile_parser_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rusprofile_parser_proto_goTypes = []interface{}{
	(*INNCompany)(nil),  // 0: rusprofile_parser.INNCompany
	(*CompanyInfo)(nil), // 1: rusprofile_parser.CompanyInfo
}
var file_rusprofile_parser_proto_depIdxs = []int32{
	0, // 0: rusprofile_parser.ClubsInfo.getCompany:input_type -> rusprofile_parser.INNCompany
	1, // 1: rusprofile_parser.ClubsInfo.getCompany:output_type -> rusprofile_parser.CompanyInfo
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_rusprofile_parser_proto_init() }
func file_rusprofile_parser_proto_init() {
	if File_rusprofile_parser_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_rusprofile_parser_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*INNCompany); i {
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
		file_rusprofile_parser_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompanyInfo); i {
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
			RawDescriptor: file_rusprofile_parser_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_rusprofile_parser_proto_goTypes,
		DependencyIndexes: file_rusprofile_parser_proto_depIdxs,
		MessageInfos:      file_rusprofile_parser_proto_msgTypes,
	}.Build()
	File_rusprofile_parser_proto = out.File
	file_rusprofile_parser_proto_rawDesc = nil
	file_rusprofile_parser_proto_goTypes = nil
	file_rusprofile_parser_proto_depIdxs = nil
}
