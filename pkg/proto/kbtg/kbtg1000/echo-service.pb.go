// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.20.1
// source: api/server/echo-service.proto

package kbtg1000

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Echo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Description string                 `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status      string                 `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	CreateDate  *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"`
	UpdateDate  *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_date,json=updateDate,proto3" json:"update_date,omitempty"`
}

func (x *Echo) Reset() {
	*x = Echo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_echo_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Echo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Echo) ProtoMessage() {}

func (x *Echo) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_echo_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Echo.ProtoReflect.Descriptor instead.
func (*Echo) Descriptor() ([]byte, []int) {
	return file_api_server_echo_service_proto_rawDescGZIP(), []int{0}
}

func (x *Echo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Echo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Echo) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Echo) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Echo) GetCreateDate() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateDate
	}
	return nil
}

func (x *Echo) GetUpdateDate() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateDate
	}
	return nil
}

type EchoList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Api      string  `protobuf:"bytes,1,opt,name=api,proto3" json:"api,omitempty"`
	EchoList []*Echo `protobuf:"bytes,2,rep,name=echoList,proto3" json:"echoList,omitempty"`
}

func (x *EchoList) Reset() {
	*x = EchoList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_server_echo_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoList) ProtoMessage() {}

func (x *EchoList) ProtoReflect() protoreflect.Message {
	mi := &file_api_server_echo_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoList.ProtoReflect.Descriptor instead.
func (*EchoList) Descriptor() ([]byte, []int) {
	return file_api_server_echo_service_proto_rawDescGZIP(), []int{1}
}

func (x *EchoList) GetApi() string {
	if x != nil {
		return x.Api
	}
	return ""
}

func (x *EchoList) GetEchoList() []*Echo {
	if x != nil {
		return x.EchoList
	}
	return nil
}

var File_api_server_echo_service_proto protoreflect.FileDescriptor

var file_api_server_echo_service_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x65, 0x63, 0x68,
	0x6f, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x6b, 0x62, 0x74, 0x67, 0x31, 0x30, 0x30, 0x30, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xde, 0x01, 0x0a, 0x04, 0x45,
	0x63, 0x68, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3b,
	0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65, 0x22, 0x48, 0x0a, 0x08, 0x45,
	0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x70, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x70, 0x69, 0x12, 0x2a, 0x0a, 0x08, 0x65, 0x63, 0x68,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6b, 0x62,
	0x74, 0x67, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x08, 0x65, 0x63, 0x68,
	0x6f, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x3c, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x0e, 0x2e,
	0x6b, 0x62, 0x74, 0x67, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x1a, 0x12, 0x2e,
	0x6b, 0x62, 0x74, 0x67, 0x31, 0x30, 0x30, 0x30, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x6b, 0x62, 0x74, 0x67, 0x2f, 0x6b, 0x62, 0x74, 0x67,
	0x31, 0x30, 0x30, 0x30, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_server_echo_service_proto_rawDescOnce sync.Once
	file_api_server_echo_service_proto_rawDescData = file_api_server_echo_service_proto_rawDesc
)

func file_api_server_echo_service_proto_rawDescGZIP() []byte {
	file_api_server_echo_service_proto_rawDescOnce.Do(func() {
		file_api_server_echo_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_server_echo_service_proto_rawDescData)
	})
	return file_api_server_echo_service_proto_rawDescData
}

var file_api_server_echo_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_api_server_echo_service_proto_goTypes = []interface{}{
	(*Echo)(nil),                  // 0: kbtg1000.Echo
	(*EchoList)(nil),              // 1: kbtg1000.EchoList
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_api_server_echo_service_proto_depIdxs = []int32{
	2, // 0: kbtg1000.Echo.create_date:type_name -> google.protobuf.Timestamp
	2, // 1: kbtg1000.Echo.update_date:type_name -> google.protobuf.Timestamp
	0, // 2: kbtg1000.EchoList.echoList:type_name -> kbtg1000.Echo
	0, // 3: kbtg1000.EchoService.Query:input_type -> kbtg1000.Echo
	1, // 4: kbtg1000.EchoService.Query:output_type -> kbtg1000.EchoList
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_server_echo_service_proto_init() }
func file_api_server_echo_service_proto_init() {
	if File_api_server_echo_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_server_echo_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Echo); i {
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
		file_api_server_echo_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoList); i {
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
			RawDescriptor: file_api_server_echo_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_server_echo_service_proto_goTypes,
		DependencyIndexes: file_api_server_echo_service_proto_depIdxs,
		MessageInfos:      file_api_server_echo_service_proto_msgTypes,
	}.Build()
	File_api_server_echo_service_proto = out.File
	file_api_server_echo_service_proto_rawDesc = nil
	file_api_server_echo_service_proto_goTypes = nil
	file_api_server_echo_service_proto_depIdxs = nil
}
