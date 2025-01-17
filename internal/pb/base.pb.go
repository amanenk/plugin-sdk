// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: internal/pb/base.proto

package pb

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

type GetExampleConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetExampleConfig) Reset() {
	*x = GetExampleConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleConfig) ProtoMessage() {}

func (x *GetExampleConfig) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleConfig.ProtoReflect.Descriptor instead.
func (*GetExampleConfig) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{0}
}

type Configure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Configure) Reset() {
	*x = Configure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure) ProtoMessage() {}

func (x *Configure) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure.ProtoReflect.Descriptor instead.
func (*Configure) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{1}
}

type GetExampleConfig_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetExampleConfig_Request) Reset() {
	*x = GetExampleConfig_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleConfig_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleConfig_Request) ProtoMessage() {}

func (x *GetExampleConfig_Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleConfig_Request.ProtoReflect.Descriptor instead.
func (*GetExampleConfig_Request) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{0, 0}
}

type GetExampleConfig_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	Config  string `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetExampleConfig_Response) Reset() {
	*x = GetExampleConfig_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetExampleConfig_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetExampleConfig_Response) ProtoMessage() {}

func (x *GetExampleConfig_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetExampleConfig_Response.ProtoReflect.Descriptor instead.
func (*GetExampleConfig_Response) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{0, 1}
}

func (x *GetExampleConfig_Response) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetExampleConfig_Response) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *GetExampleConfig_Response) GetConfig() string {
	if x != nil {
		return x.Config
	}
	return ""
}

type Configure_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Holds information such as credentials, regions, accounts, etc'
	// Marshalled spec.SourceSpec or spec.DestinationSpec
	Config []byte `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *Configure_Request) Reset() {
	*x = Configure_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure_Request) ProtoMessage() {}

func (x *Configure_Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure_Request.ProtoReflect.Descriptor instead.
func (*Configure_Request) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Configure_Request) GetConfig() []byte {
	if x != nil {
		return x.Config
	}
	return nil
}

type Configure_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error            string `protobuf:"bytes,1,opt,name=error,proto3" json:"error,omitempty"`
	JsonschemaResult []byte `protobuf:"bytes,2,opt,name=jsonschema_result,json=jsonschemaResult,proto3" json:"jsonschema_result,omitempty"`
}

func (x *Configure_Response) Reset() {
	*x = Configure_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_base_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configure_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configure_Response) ProtoMessage() {}

func (x *Configure_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_base_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configure_Response.ProtoReflect.Descriptor instead.
func (*Configure_Response) Descriptor() ([]byte, []int) {
	return file_internal_pb_base_proto_rawDescGZIP(), []int{1, 1}
}

func (x *Configure_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

func (x *Configure_Response) GetJsonschemaResult() []byte {
	if x != nil {
		return x.JsonschemaResult
	}
	return nil
}

var File_internal_pb_base_proto protoreflect.FileDescriptor

var file_internal_pb_base_proto_rawDesc = []byte{
	0x0a, 0x16, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x6f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x50,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x7d, 0x0a, 0x09, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x65, 0x1a, 0x21, 0x0a,
	0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x1a, 0x4d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x12, 0x2b, 0x0a, 0x11, 0x6a, 0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x10, 0x6a,
	0x73, 0x6f, 0x6e, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42,
	0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_base_proto_rawDescOnce sync.Once
	file_internal_pb_base_proto_rawDescData = file_internal_pb_base_proto_rawDesc
)

func file_internal_pb_base_proto_rawDescGZIP() []byte {
	file_internal_pb_base_proto_rawDescOnce.Do(func() {
		file_internal_pb_base_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_base_proto_rawDescData)
	})
	return file_internal_pb_base_proto_rawDescData
}

var file_internal_pb_base_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_pb_base_proto_goTypes = []interface{}{
	(*GetExampleConfig)(nil),          // 0: proto.GetExampleConfig
	(*Configure)(nil),                 // 1: proto.Configure
	(*GetExampleConfig_Request)(nil),  // 2: proto.GetExampleConfig.Request
	(*GetExampleConfig_Response)(nil), // 3: proto.GetExampleConfig.Response
	(*Configure_Request)(nil),         // 4: proto.Configure.Request
	(*Configure_Response)(nil),        // 5: proto.Configure.Response
}
var file_internal_pb_base_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_base_proto_init() }
func file_internal_pb_base_proto_init() {
	if File_internal_pb_base_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_base_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExampleConfig); i {
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
		file_internal_pb_base_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure); i {
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
		file_internal_pb_base_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExampleConfig_Request); i {
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
		file_internal_pb_base_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetExampleConfig_Response); i {
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
		file_internal_pb_base_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure_Request); i {
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
		file_internal_pb_base_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Configure_Response); i {
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
			RawDescriptor: file_internal_pb_base_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_base_proto_goTypes,
		DependencyIndexes: file_internal_pb_base_proto_depIdxs,
		MessageInfos:      file_internal_pb_base_proto_msgTypes,
	}.Build()
	File_internal_pb_base_proto = out.File
	file_internal_pb_base_proto_rawDesc = nil
	file_internal_pb_base_proto_goTypes = nil
	file_internal_pb_base_proto_depIdxs = nil
}
