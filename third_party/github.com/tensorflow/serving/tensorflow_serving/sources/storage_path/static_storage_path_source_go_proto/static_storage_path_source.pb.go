// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tensorflow_serving/sources/storage_path/static_storage_path_source.proto

package static_storage_path_source_go_proto

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

// Config proto for StaticStoragePathSource.
type StaticStoragePathSourceConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The single servable name, version number and path to supply statically.
	ServableName string `protobuf:"bytes,1,opt,name=servable_name,json=servableName,proto3" json:"servable_name,omitempty"`
	VersionNum   int64  `protobuf:"varint,2,opt,name=version_num,json=versionNum,proto3" json:"version_num,omitempty"`
	VersionPath  string `protobuf:"bytes,3,opt,name=version_path,json=versionPath,proto3" json:"version_path,omitempty"`
}

func (x *StaticStoragePathSourceConfig) Reset() {
	*x = StaticStoragePathSourceConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StaticStoragePathSourceConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StaticStoragePathSourceConfig) ProtoMessage() {}

func (x *StaticStoragePathSourceConfig) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StaticStoragePathSourceConfig.ProtoReflect.Descriptor instead.
func (*StaticStoragePathSourceConfig) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescGZIP(), []int{0}
}

func (x *StaticStoragePathSourceConfig) GetServableName() string {
	if x != nil {
		return x.ServableName
	}
	return ""
}

func (x *StaticStoragePathSourceConfig) GetVersionNum() int64 {
	if x != nil {
		return x.VersionNum
	}
	return 0
}

func (x *StaticStoragePathSourceConfig) GetVersionPath() string {
	if x != nil {
		return x.VersionPath
	}
	return ""
}

var File_tensorflow_serving_sources_storage_path_static_storage_path_source_proto protoreflect.FileDescriptor

var file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDesc = []byte{
	0x0a, 0x48, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63,
	0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x22, 0x88,
	0x01, 0x0a, 0x1d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x61, 0x62, 0x6c,
	0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x4e, 0x75, 0x6d, 0x12, 0x21, 0x0a, 0x0c, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x74, 0x68, 0x42, 0x6b, 0x5a, 0x69, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f,
	0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x5f, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x67, 0x6f,
	0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescOnce sync.Once
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescData = file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDesc
)

func file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescData)
	})
	return file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDescData
}

var file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_goTypes = []interface{}{
	(*StaticStoragePathSourceConfig)(nil), // 0: tensorflow.serving.StaticStoragePathSourceConfig
}
var file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_init() }
func file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_init() {
	if File_tensorflow_serving_sources_storage_path_static_storage_path_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StaticStoragePathSourceConfig); i {
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
			RawDescriptor: file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_sources_storage_path_static_storage_path_source_proto = out.File
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_rawDesc = nil
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_goTypes = nil
	file_tensorflow_serving_sources_storage_path_static_storage_path_source_proto_depIdxs = nil
}
