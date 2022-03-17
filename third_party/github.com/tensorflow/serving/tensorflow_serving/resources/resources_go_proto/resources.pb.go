// Representations for resources used by servables, and available in a system.
//
// Each of the string-typed values are free-form, so that they can be extended
// by third parties. However we strongly recommend using the values defined in
// resource_values.h when possible, for standardization.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: tensorflow_serving/resources/resources.proto

package resources_go_proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// One kind of resource on one device (or type of device).
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of device on which the resource resides, e.g. CPU or GPU.
	Device string `protobuf:"bytes,1,opt,name=device,proto3" json:"device,omitempty"`
	// A specific instance of the device of type 'device' to which the resources
	// are bound (instances are assumed to be numbered 0, 1, ...).
	//
	// When representing the resources required by a servable that has yet to be
	// loaded, this field is optional. If not set, it denotes that the servable's
	// resources are not (yet) bound to a specific instance.
	DeviceInstance *wrapperspb.UInt32Value `protobuf:"bytes,2,opt,name=device_instance,json=deviceInstance,proto3" json:"device_instance,omitempty"`
	// The kind of resource on the device (instance), e.g. RAM or compute share.
	//
	// A given type of resource should have a standard unit that represents the
	// smallest useful quantization. We strongly recommend including the unit
	// (e.g. bytes or millicores) in this string, as in "ram_bytes".
	Kind string `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_resources_resources_proto_rawDescGZIP(), []int{0}
}

func (x *Resource) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *Resource) GetDeviceInstance() *wrapperspb.UInt32Value {
	if x != nil {
		return x.DeviceInstance
	}
	return nil
}

func (x *Resource) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

// An allocation of one or more kinds of resources, along with the quantity of
// each. Used to denote the resources that a servable (or collection of
// servables) will use or is currently using. Also used to denote resources
// available to the serving system for loading more servables.
type ResourceAllocation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceQuantities []*ResourceAllocation_Entry `protobuf:"bytes,1,rep,name=resource_quantities,json=resourceQuantities,proto3" json:"resource_quantities,omitempty"`
}

func (x *ResourceAllocation) Reset() {
	*x = ResourceAllocation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceAllocation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAllocation) ProtoMessage() {}

func (x *ResourceAllocation) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAllocation.ProtoReflect.Descriptor instead.
func (*ResourceAllocation) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_resources_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceAllocation) GetResourceQuantities() []*ResourceAllocation_Entry {
	if x != nil {
		return x.ResourceQuantities
	}
	return nil
}

// A collection of resources, each with a quantity. Treated as a resource->
// quantity map, i.e. no resource can repeat and the order is immaterial.
type ResourceAllocation_Entry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Resource *Resource `protobuf:"bytes,1,opt,name=resource,proto3" json:"resource,omitempty"`
	Quantity uint64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *ResourceAllocation_Entry) Reset() {
	*x = ResourceAllocation_Entry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceAllocation_Entry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceAllocation_Entry) ProtoMessage() {}

func (x *ResourceAllocation_Entry) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_serving_resources_resources_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceAllocation_Entry.ProtoReflect.Descriptor instead.
func (*ResourceAllocation_Entry) Descriptor() ([]byte, []int) {
	return file_tensorflow_serving_resources_resources_proto_rawDescGZIP(), []int{1, 0}
}

func (x *ResourceAllocation_Entry) GetResource() *Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceAllocation_Entry) GetQuantity() uint64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_tensorflow_serving_resources_resources_proto protoreflect.FileDescriptor

var file_tensorflow_serving_resources_resources_proto_rawDesc = []byte{
	0x0a, 0x2c, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12,
	0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x7d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x45, 0x0a, 0x0f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0e, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e,
	0x64, 0x22, 0xd2, 0x01, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x41, 0x6c,
	0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x5d, 0x0a, 0x13, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x41, 0x6c, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x12, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x51, 0x75, 0x61,
	0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x5d, 0x0a, 0x05, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x4f, 0x5a, 0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x73, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x5f, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_serving_resources_resources_proto_rawDescOnce sync.Once
	file_tensorflow_serving_resources_resources_proto_rawDescData = file_tensorflow_serving_resources_resources_proto_rawDesc
)

func file_tensorflow_serving_resources_resources_proto_rawDescGZIP() []byte {
	file_tensorflow_serving_resources_resources_proto_rawDescOnce.Do(func() {
		file_tensorflow_serving_resources_resources_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_serving_resources_resources_proto_rawDescData)
	})
	return file_tensorflow_serving_resources_resources_proto_rawDescData
}

var file_tensorflow_serving_resources_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_tensorflow_serving_resources_resources_proto_goTypes = []interface{}{
	(*Resource)(nil),                 // 0: tensorflow.serving.Resource
	(*ResourceAllocation)(nil),       // 1: tensorflow.serving.ResourceAllocation
	(*ResourceAllocation_Entry)(nil), // 2: tensorflow.serving.ResourceAllocation.Entry
	(*wrapperspb.UInt32Value)(nil),   // 3: google.protobuf.UInt32Value
}
var file_tensorflow_serving_resources_resources_proto_depIdxs = []int32{
	3, // 0: tensorflow.serving.Resource.device_instance:type_name -> google.protobuf.UInt32Value
	2, // 1: tensorflow.serving.ResourceAllocation.resource_quantities:type_name -> tensorflow.serving.ResourceAllocation.Entry
	0, // 2: tensorflow.serving.ResourceAllocation.Entry.resource:type_name -> tensorflow.serving.Resource
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_tensorflow_serving_resources_resources_proto_init() }
func file_tensorflow_serving_resources_resources_proto_init() {
	if File_tensorflow_serving_resources_resources_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_serving_resources_resources_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
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
		file_tensorflow_serving_resources_resources_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceAllocation); i {
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
		file_tensorflow_serving_resources_resources_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceAllocation_Entry); i {
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
			RawDescriptor: file_tensorflow_serving_resources_resources_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_serving_resources_resources_proto_goTypes,
		DependencyIndexes: file_tensorflow_serving_resources_resources_proto_depIdxs,
		MessageInfos:      file_tensorflow_serving_resources_resources_proto_msgTypes,
	}.Build()
	File_tensorflow_serving_resources_resources_proto = out.File
	file_tensorflow_serving_resources_resources_proto_rawDesc = nil
	file_tensorflow_serving_resources_resources_proto_goTypes = nil
	file_tensorflow_serving_resources_resources_proto_depIdxs = nil
}
