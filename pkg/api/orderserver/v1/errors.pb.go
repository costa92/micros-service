// Copyright 2024 costalong <costa9293@gmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0--rc1
// source: orderserver/v1/errors.proto

package v1

import (
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 订单找不到 ，可能是订单不存在或输入的订单标识有误
	ErrorReason_OrderNotFound ErrorReason = 0
	// 订单已存在，无法创建用户
	ErrorReason_OrderAlreadyExists ErrorReason = 1
	// 创建订单失败，可能是由于服务器或其他问题导致的创建过程中的错误
	ErrorReason_OrderCreateFailed ErrorReason = 2
	// 订单状态已经是完成状态，无法再次修改
	ErrorReason_OrderStatusCompleted ErrorReason = 3
	// 订单状态已经支付，无法再次支付
	ErrorReason_OrderStatusPaid ErrorReason = 4
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "OrderNotFound",
		1: "OrderAlreadyExists",
		2: "OrderCreateFailed",
		3: "OrderStatusCompleted",
		4: "OrderStatusPaid",
	}
	ErrorReason_value = map[string]int32{
		"OrderNotFound":        0,
		"OrderAlreadyExists":   1,
		"OrderCreateFailed":    2,
		"OrderStatusCompleted": 3,
		"OrderStatusPaid":      4,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_orderserver_v1_errors_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_orderserver_v1_errors_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_orderserver_v1_errors_proto_rawDescGZIP(), []int{0}
}

var File_orderserver_v1_errors_proto protoreflect.FileDescriptor

var file_orderserver_v1_errors_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x76, 0x31,
	0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x66,
	0x61, 0x6b, 0x65, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2a, 0xa2, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x17, 0x0a, 0x0d, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0x94, 0x03, 0x12, 0x1c, 0x0a, 0x12, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73,
	0x10, 0x01, 0x1a, 0x04, 0xa8, 0x45, 0x99, 0x03, 0x12, 0x1b, 0x0a, 0x11, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x1a,
	0x04, 0xa8, 0x45, 0x9d, 0x04, 0x12, 0x1e, 0x0a, 0x14, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x10, 0x03, 0x1a,
	0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x19, 0x0a, 0x0f, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x50, 0x61, 0x69, 0x64, 0x10, 0x04, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03,
	0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x74, 0x61, 0x39, 0x32, 0x2f, 0x6d, 0x69, 0x63,
	0x72, 0x6f, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orderserver_v1_errors_proto_rawDescOnce sync.Once
	file_orderserver_v1_errors_proto_rawDescData = file_orderserver_v1_errors_proto_rawDesc
)

func file_orderserver_v1_errors_proto_rawDescGZIP() []byte {
	file_orderserver_v1_errors_proto_rawDescOnce.Do(func() {
		file_orderserver_v1_errors_proto_rawDescData = protoimpl.X.CompressGZIP(file_orderserver_v1_errors_proto_rawDescData)
	})
	return file_orderserver_v1_errors_proto_rawDescData
}

var file_orderserver_v1_errors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orderserver_v1_errors_proto_goTypes = []any{
	(ErrorReason)(0), // 0: fakeserver.v1.ErrorReason
}
var file_orderserver_v1_errors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_orderserver_v1_errors_proto_init() }
func file_orderserver_v1_errors_proto_init() {
	if File_orderserver_v1_errors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orderserver_v1_errors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orderserver_v1_errors_proto_goTypes,
		DependencyIndexes: file_orderserver_v1_errors_proto_depIdxs,
		EnumInfos:         file_orderserver_v1_errors_proto_enumTypes,
	}.Build()
	File_orderserver_v1_errors_proto = out.File
	file_orderserver_v1_errors_proto_rawDesc = nil
	file_orderserver_v1_errors_proto_goTypes = nil
	file_orderserver_v1_errors_proto_depIdxs = nil
}
