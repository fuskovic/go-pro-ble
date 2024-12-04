// turbo_transfer.proto/Open GoPro, Version 2.0 (C) Copyright 2021 GoPro, Inc. (http://gopro.com/OpenGoPro).

// This copyright was auto-generated on Fri Oct  4 17:02:52 UTC 2024

//**********************************************************************************************************************
//
// This file is automatically generated!!! Do not modify manually.
//
//********************************************************************************************************************

//*
// Defines the structure of protobuf messages for enabling and disabling Turbo Transfer feature

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v3.6.1
// source: turbo_transfer.proto

package __

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// *
// Enable/disable display of "Transferring Media" UI
//
// Response: @ref ResponseGeneric
type RequestSetTurboActive struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Active *bool `protobuf:"varint,1,req,name=active" json:"active,omitempty"` // Enable or disable Turbo Transfer feature
}

func (x *RequestSetTurboActive) Reset() {
	*x = RequestSetTurboActive{}
	if protoimpl.UnsafeEnabled {
		mi := &file_turbo_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestSetTurboActive) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestSetTurboActive) ProtoMessage() {}

func (x *RequestSetTurboActive) ProtoReflect() protoreflect.Message {
	mi := &file_turbo_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestSetTurboActive.ProtoReflect.Descriptor instead.
func (*RequestSetTurboActive) Descriptor() ([]byte, []int) {
	return file_turbo_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *RequestSetTurboActive) GetActive() bool {
	if x != nil && x.Active != nil {
		return *x.Active
	}
	return false
}

var File_turbo_transfer_proto protoreflect.FileDescriptor

var file_turbo_transfer_proto_rawDesc = []byte{
	0x0a, 0x14, 0x74, 0x75, 0x72, 0x62, 0x6f, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6f, 0x70, 0x65, 0x6e, 0x5f, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x22, 0x2f, 0x0a, 0x15, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x53, 0x65, 0x74,
	0x54, 0x75, 0x72, 0x62, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x08, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x76, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e,
}

var (
	file_turbo_transfer_proto_rawDescOnce sync.Once
	file_turbo_transfer_proto_rawDescData = file_turbo_transfer_proto_rawDesc
)

func file_turbo_transfer_proto_rawDescGZIP() []byte {
	file_turbo_transfer_proto_rawDescOnce.Do(func() {
		file_turbo_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_turbo_transfer_proto_rawDescData)
	})
	return file_turbo_transfer_proto_rawDescData
}

var file_turbo_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_turbo_transfer_proto_goTypes = []interface{}{
	(*RequestSetTurboActive)(nil), // 0: open_gopro.RequestSetTurboActive
}
var file_turbo_transfer_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_turbo_transfer_proto_init() }
func file_turbo_transfer_proto_init() {
	if File_turbo_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_turbo_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestSetTurboActive); i {
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
			RawDescriptor: file_turbo_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_turbo_transfer_proto_goTypes,
		DependencyIndexes: file_turbo_transfer_proto_depIdxs,
		MessageInfos:      file_turbo_transfer_proto_msgTypes,
	}.Build()
	File_turbo_transfer_proto = out.File
	file_turbo_transfer_proto_rawDesc = nil
	file_turbo_transfer_proto_goTypes = nil
	file_turbo_transfer_proto_depIdxs = nil
}
