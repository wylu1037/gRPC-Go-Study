// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: proto/InvokeJava.proto

package invoke

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

type InvokeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Target   string `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
}

func (x *InvokeRequest) Reset() {
	*x = InvokeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_InvokeJava_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeRequest) ProtoMessage() {}

func (x *InvokeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_InvokeJava_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeRequest.ProtoReflect.Descriptor instead.
func (*InvokeRequest) Descriptor() ([]byte, []int) {
	return file_proto_InvokeJava_proto_rawDescGZIP(), []int{0}
}

func (x *InvokeRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *InvokeRequest) GetTarget() string {
	if x != nil {
		return x.Target
	}
	return ""
}

type InvokeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result string `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	Code   int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *InvokeResponse) Reset() {
	*x = InvokeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_InvokeJava_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InvokeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeResponse) ProtoMessage() {}

func (x *InvokeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_InvokeJava_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeResponse.ProtoReflect.Descriptor instead.
func (*InvokeResponse) Descriptor() ([]byte, []int) {
	return file_proto_InvokeJava_proto_rawDescGZIP(), []int{1}
}

func (x *InvokeResponse) GetResult() string {
	if x != nil {
		return x.Result
	}
	return ""
}

func (x *InvokeResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

var File_proto_InvokeJava_proto protoreflect.FileDescriptor

var file_proto_InvokeJava_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4a, 0x61,
	0x76, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x22, 0x43, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a,
	0x06, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x3c, 0x0a, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x32, 0x53, 0x0a, 0x0b, 0x49, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x4a, 0x61,
	0x76, 0x61, 0x12, 0x44, 0x0a, 0x13, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x72, 0x6f, 0x73,
	0x73, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12, 0x15, 0x2e, 0x69, 0x6e, 0x76, 0x6f,
	0x6b, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x2e, 0x49, 0x6e, 0x76, 0x6f, 0x6b, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x31, 0x0a, 0x19, 0x63, 0x6e, 0x2e, 0x65,
	0x64, 0x75, 0x2e, 0x68, 0x75, 0x73, 0x74, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x5a, 0x14, 0x67, 0x52, 0x50, 0x43, 0x2d, 0x47, 0x6f, 0x2d, 0x53,
	0x74, 0x75, 0x64, 0x79, 0x2f, 0x69, 0x6e, 0x76, 0x6f, 0x6b, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_proto_InvokeJava_proto_rawDescOnce sync.Once
	file_proto_InvokeJava_proto_rawDescData = file_proto_InvokeJava_proto_rawDesc
)

func file_proto_InvokeJava_proto_rawDescGZIP() []byte {
	file_proto_InvokeJava_proto_rawDescOnce.Do(func() {
		file_proto_InvokeJava_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_InvokeJava_proto_rawDescData)
	})
	return file_proto_InvokeJava_proto_rawDescData
}

var file_proto_InvokeJava_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_InvokeJava_proto_goTypes = []interface{}{
	(*InvokeRequest)(nil),  // 0: invoke.InvokeRequest
	(*InvokeResponse)(nil), // 1: invoke.InvokeResponse
}
var file_proto_InvokeJava_proto_depIdxs = []int32{
	0, // 0: invoke.IInvokeJava.invokeCrossLanguage:input_type -> invoke.InvokeRequest
	1, // 1: invoke.IInvokeJava.invokeCrossLanguage:output_type -> invoke.InvokeResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_InvokeJava_proto_init() }
func file_proto_InvokeJava_proto_init() {
	if File_proto_InvokeJava_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_InvokeJava_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeRequest); i {
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
		file_proto_InvokeJava_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InvokeResponse); i {
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
			RawDescriptor: file_proto_InvokeJava_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_InvokeJava_proto_goTypes,
		DependencyIndexes: file_proto_InvokeJava_proto_depIdxs,
		MessageInfos:      file_proto_InvokeJava_proto_msgTypes,
	}.Build()
	File_proto_InvokeJava_proto = out.File
	file_proto_InvokeJava_proto_rawDesc = nil
	file_proto_InvokeJava_proto_goTypes = nil
	file_proto_InvokeJava_proto_depIdxs = nil
}
