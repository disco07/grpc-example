// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.1
// source: calulator/calculator.proto

package calculator

import (
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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calulator_calculator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calulator_calculator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_calulator_calculator_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddRequest) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calulator_calculator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calulator_calculator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_calulator_calculator_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

type ModuloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
}

func (x *ModuloRequest) Reset() {
	*x = ModuloRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calulator_calculator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuloRequest) ProtoMessage() {}

func (x *ModuloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_calulator_calculator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuloRequest.ProtoReflect.Descriptor instead.
func (*ModuloRequest) Descriptor() ([]byte, []int) {
	return file_calulator_calculator_proto_rawDescGZIP(), []int{2}
}

func (x *ModuloRequest) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

type ModuloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result int64 `protobuf:"varint,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ModuloResponse) Reset() {
	*x = ModuloResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_calulator_calculator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModuloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModuloResponse) ProtoMessage() {}

func (x *ModuloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_calulator_calculator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModuloResponse.ProtoReflect.Descriptor instead.
func (*ModuloResponse) Descriptor() ([]byte, []int) {
	return file_calulator_calculator_proto_rawDescGZIP(), []int{3}
}

func (x *ModuloResponse) GetResult() int64 {
	if x != nil {
		return x.Result
	}
	return 0
}

var File_calulator_calculator_proto protoreflect.FileDescriptor

var file_calulator_calculator_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x61, 0x6c, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x63, 0x61, 0x6c, 0x63,
	0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x63, 0x61,
	0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x28, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x62,
	0x22, 0x25, 0x0a, 0x0b, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x1d, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x75, 0x6c,
	0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x01, 0x61, 0x22, 0x28, 0x0a, 0x0e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x32, 0xcf, 0x01, 0x0a, 0x11, 0x43, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x55, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12, 0x16, 0x2e,
	0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1d,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x3a, 0x01, 0x2a, 0x22, 0x12, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x61, 0x64, 0x64, 0x12, 0x63, 0x0a,
	0x06, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x6f, 0x12, 0x19, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c,
	0x61, 0x74, 0x6f, 0x72, 0x2e, 0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2e,
	0x4d, 0x6f, 0x64, 0x75, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x20,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1a, 0x3a, 0x01, 0x2a, 0x22, 0x15, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74, 0x6f, 0x72, 0x2f, 0x6d, 0x6f, 0x64, 0x75, 0x6c, 0x6f,
	0x30, 0x01, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x30, 0x37, 0x2f, 0x67, 0x52, 0x50, 0x43, 0x2f, 0x67, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x63, 0x61, 0x6c, 0x63, 0x75, 0x6c, 0x61, 0x74,
	0x6f, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_calulator_calculator_proto_rawDescOnce sync.Once
	file_calulator_calculator_proto_rawDescData = file_calulator_calculator_proto_rawDesc
)

func file_calulator_calculator_proto_rawDescGZIP() []byte {
	file_calulator_calculator_proto_rawDescOnce.Do(func() {
		file_calulator_calculator_proto_rawDescData = protoimpl.X.CompressGZIP(file_calulator_calculator_proto_rawDescData)
	})
	return file_calulator_calculator_proto_rawDescData
}

var file_calulator_calculator_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_calulator_calculator_proto_goTypes = []any{
	(*AddRequest)(nil),     // 0: calculator.AddRequest
	(*AddResponse)(nil),    // 1: calculator.AddResponse
	(*ModuloRequest)(nil),  // 2: calculator.ModuloRequest
	(*ModuloResponse)(nil), // 3: calculator.ModuloResponse
}
var file_calulator_calculator_proto_depIdxs = []int32{
	0, // 0: calculator.CalculatorService.Add:input_type -> calculator.AddRequest
	2, // 1: calculator.CalculatorService.Modulo:input_type -> calculator.ModuloRequest
	1, // 2: calculator.CalculatorService.Add:output_type -> calculator.AddResponse
	3, // 3: calculator.CalculatorService.Modulo:output_type -> calculator.ModuloResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_calulator_calculator_proto_init() }
func file_calulator_calculator_proto_init() {
	if File_calulator_calculator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_calulator_calculator_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AddRequest); i {
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
		file_calulator_calculator_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AddResponse); i {
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
		file_calulator_calculator_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*ModuloRequest); i {
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
		file_calulator_calculator_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ModuloResponse); i {
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
			RawDescriptor: file_calulator_calculator_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_calulator_calculator_proto_goTypes,
		DependencyIndexes: file_calulator_calculator_proto_depIdxs,
		MessageInfos:      file_calulator_calculator_proto_msgTypes,
	}.Build()
	File_calulator_calculator_proto = out.File
	file_calulator_calculator_proto_rawDesc = nil
	file_calulator_calculator_proto_goTypes = nil
	file_calulator_calculator_proto_depIdxs = nil
}
