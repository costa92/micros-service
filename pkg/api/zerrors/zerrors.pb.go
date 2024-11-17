// Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/costa92/micros-service.
//

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.28.3
// source: zerrors/zerrors.proto

package zerrors

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
	// 未知错误，服务器内部错误
	ErrorReason_Unknown ErrorReason = 0
	// 无效参数错误
	ErrorReason_InvalidParameter ErrorReason = 1
	// 未找到错误
	ErrorReason_NotFound ErrorReason = 2
	// 未经授权错误
	ErrorReason_Unauthorized ErrorReason = 3
	// 禁止访问错误
	ErrorReason_Forbidden ErrorReason = 4
	// 缺少幂等性令牌错误
	ErrorReason_IdempotentMissingToken ErrorReason = 5
	// 幂等性令牌已过期错误
	ErrorReason_IdempotentTokenExpired ErrorReason = 6
	// 请求路径没有找到
	ErrorReason_PageNotFound ErrorReason = 7
	// Gin 请求参数绑定失败
	ErrorReason_BindFailed ErrorReason = 8
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "Unknown",
		1: "InvalidParameter",
		2: "NotFound",
		3: "Unauthorized",
		4: "Forbidden",
		5: "IdempotentMissingToken",
		6: "IdempotentTokenExpired",
		7: "PageNotFound",
		8: "BindFailed",
	}
	ErrorReason_value = map[string]int32{
		"Unknown":                0,
		"InvalidParameter":       1,
		"NotFound":               2,
		"Unauthorized":           3,
		"Forbidden":              4,
		"IdempotentMissingToken": 5,
		"IdempotentTokenExpired": 6,
		"PageNotFound":           7,
		"BindFailed":             8,
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
	return file_zerrors_zerrors_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_zerrors_zerrors_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_zerrors_zerrors_proto_rawDescGZIP(), []int{0}
}

var File_zerrors_zerrors_proto protoreflect.FileDescriptor

var file_zerrors_zerrors_proto_rawDesc = []byte{
	0x0a, 0x15, 0x7a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x7a, 0x65, 0x72, 0x72, 0x6f, 0x72,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x7a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x1a, 0x13, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf5, 0x01, 0x0a, 0x0b, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x52,
	0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x11, 0x0a, 0x07, 0x55, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x12, 0x1a, 0x0a, 0x10, 0x49, 0x6e, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x10, 0x01, 0x1a, 0x04,
	0xa8, 0x45, 0x90, 0x03, 0x12, 0x12, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64,
	0x10, 0x02, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x55, 0x6e, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64, 0x10, 0x03, 0x1a, 0x04, 0xa8, 0x45, 0x91, 0x03,
	0x12, 0x13, 0x0a, 0x09, 0x46, 0x6f, 0x72, 0x62, 0x69, 0x64, 0x64, 0x65, 0x6e, 0x10, 0x04, 0x1a,
	0x04, 0xa8, 0x45, 0x93, 0x03, 0x12, 0x20, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6d, 0x70, 0x6f, 0x74,
	0x65, 0x6e, 0x74, 0x4d, 0x69, 0x73, 0x73, 0x69, 0x6e, 0x67, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x10,
	0x05, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x20, 0x0a, 0x16, 0x49, 0x64, 0x65, 0x6d, 0x70,
	0x6f, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65,
	0x64, 0x10, 0x06, 0x1a, 0x04, 0xa8, 0x45, 0x90, 0x03, 0x12, 0x16, 0x0a, 0x0c, 0x50, 0x61, 0x67,
	0x65, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e, 0x64, 0x10, 0x07, 0x1a, 0x04, 0xa8, 0x45, 0x94,
	0x03, 0x12, 0x14, 0x0a, 0x0a, 0x42, 0x69, 0x6e, 0x64, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x10,
	0x08, 0x1a, 0x04, 0xa8, 0x45, 0xf4, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x42, 0x3b, 0x5a,
	0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x73, 0x74,
	0x61, 0x39, 0x32, 0x2f, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x73, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x7a, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x3b, 0x7a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_zerrors_zerrors_proto_rawDescOnce sync.Once
	file_zerrors_zerrors_proto_rawDescData = file_zerrors_zerrors_proto_rawDesc
)

func file_zerrors_zerrors_proto_rawDescGZIP() []byte {
	file_zerrors_zerrors_proto_rawDescOnce.Do(func() {
		file_zerrors_zerrors_proto_rawDescData = protoimpl.X.CompressGZIP(file_zerrors_zerrors_proto_rawDescData)
	})
	return file_zerrors_zerrors_proto_rawDescData
}

var file_zerrors_zerrors_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_zerrors_zerrors_proto_goTypes = []interface{}{
	(ErrorReason)(0), // 0: zerrors.ErrorReason
}
var file_zerrors_zerrors_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_zerrors_zerrors_proto_init() }
func file_zerrors_zerrors_proto_init() {
	if File_zerrors_zerrors_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_zerrors_zerrors_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_zerrors_zerrors_proto_goTypes,
		DependencyIndexes: file_zerrors_zerrors_proto_depIdxs,
		EnumInfos:         file_zerrors_zerrors_proto_enumTypes,
	}.Build()
	File_zerrors_zerrors_proto = out.File
	file_zerrors_zerrors_proto_rawDesc = nil
	file_zerrors_zerrors_proto_goTypes = nil
	file_zerrors_zerrors_proto_depIdxs = nil
}
