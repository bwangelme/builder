// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: builder/pbs/recordpb/record.proto

package recordpb

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

type StatusType int32

const (
	StatusType_Init     StatusType = 0
	StatusType_Building StatusType = 1
	StatusType_Success  StatusType = 2
	StatusType_Failed   StatusType = 3
)

// Enum value maps for StatusType.
var (
	StatusType_name = map[int32]string{
		0: "Init",
		1: "Building",
		2: "Success",
		3: "Failed",
	}
	StatusType_value = map[string]int32{
		"Init":     0,
		"Building": 1,
		"Success":  2,
		"Failed":   3,
	}
)

func (x StatusType) Enum() *StatusType {
	p := new(StatusType)
	*p = x
	return p
}

func (x StatusType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (StatusType) Descriptor() protoreflect.EnumDescriptor {
	return file_builder_pbs_recordpb_record_proto_enumTypes[0].Descriptor()
}

func (StatusType) Type() protoreflect.EnumType {
	return &file_builder_pbs_recordpb_record_proto_enumTypes[0]
}

func (x StatusType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use StatusType.Descriptor instead.
func (StatusType) EnumDescriptor() ([]byte, []int) {
	return file_builder_pbs_recordpb_record_proto_rawDescGZIP(), []int{0}
}

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Giturl   string `protobuf:"bytes,2,opt,name=giturl,proto3" json:"giturl,omitempty"`
	Revision string `protobuf:"bytes,3,opt,name=revision,proto3" json:"revision,omitempty"`
	Image    string `protobuf:"bytes,4,opt,name=image,proto3" json:"image,omitempty"`
	Tag      string `protobuf:"bytes,5,opt,name=tag,proto3" json:"tag,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_builder_pbs_recordpb_record_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_builder_pbs_recordpb_record_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_builder_pbs_recordpb_record_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Record) GetGiturl() string {
	if x != nil {
		return x.Giturl
	}
	return ""
}

func (x *Record) GetRevision() string {
	if x != nil {
		return x.Revision
	}
	return ""
}

func (x *Record) GetImage() string {
	if x != nil {
		return x.Image
	}
	return ""
}

func (x *Record) GetTag() string {
	if x != nil {
		return x.Tag
	}
	return ""
}

var File_builder_pbs_recordpb_record_proto protoreflect.FileDescriptor

var file_builder_pbs_recordpb_record_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x73, 0x2f, 0x72, 0x65,
	0x63, 0x6f, 0x72, 0x64, 0x70, 0x62, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x70, 0x62, 0x22, 0x74, 0x0a,
	0x06, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x69, 0x74, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x69, 0x74, 0x75, 0x72, 0x6c, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x69,
	0x6d, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6d, 0x61, 0x67,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x74, 0x61, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x74, 0x61, 0x67, 0x2a, 0x3d, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x42,
	0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x10, 0x02, 0x12, 0x0a, 0x0a, 0x06, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64,
	0x10, 0x03, 0x42, 0x16, 0x5a, 0x14, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x65, 0x72, 0x2f, 0x70, 0x62,
	0x73, 0x2f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_builder_pbs_recordpb_record_proto_rawDescOnce sync.Once
	file_builder_pbs_recordpb_record_proto_rawDescData = file_builder_pbs_recordpb_record_proto_rawDesc
)

func file_builder_pbs_recordpb_record_proto_rawDescGZIP() []byte {
	file_builder_pbs_recordpb_record_proto_rawDescOnce.Do(func() {
		file_builder_pbs_recordpb_record_proto_rawDescData = protoimpl.X.CompressGZIP(file_builder_pbs_recordpb_record_proto_rawDescData)
	})
	return file_builder_pbs_recordpb_record_proto_rawDescData
}

var file_builder_pbs_recordpb_record_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_builder_pbs_recordpb_record_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_builder_pbs_recordpb_record_proto_goTypes = []interface{}{
	(StatusType)(0), // 0: recordpb.StatusType
	(*Record)(nil),  // 1: recordpb.Record
}
var file_builder_pbs_recordpb_record_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_builder_pbs_recordpb_record_proto_init() }
func file_builder_pbs_recordpb_record_proto_init() {
	if File_builder_pbs_recordpb_record_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_builder_pbs_recordpb_record_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
			RawDescriptor: file_builder_pbs_recordpb_record_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_builder_pbs_recordpb_record_proto_goTypes,
		DependencyIndexes: file_builder_pbs_recordpb_record_proto_depIdxs,
		EnumInfos:         file_builder_pbs_recordpb_record_proto_enumTypes,
		MessageInfos:      file_builder_pbs_recordpb_record_proto_msgTypes,
	}.Build()
	File_builder_pbs_recordpb_record_proto = out.File
	file_builder_pbs_recordpb_record_proto_rawDesc = nil
	file_builder_pbs_recordpb_record_proto_goTypes = nil
	file_builder_pbs_recordpb_record_proto_depIdxs = nil
}
