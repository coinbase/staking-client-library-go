// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/orchestration/v1/api.proto

package v1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_coinbase_staking_orchestration_v1_api_proto protoreflect.FileDescriptor

var file_coinbase_staking_orchestration_v1_api_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2f, 0x76, 0x31, 0x2f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31,
	0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x36, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x30, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x37, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x5f,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x86, 0x14,
	0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0xf9, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x38, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f,
	0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x92, 0x41, 0x5a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x1a, 0x18,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41, 0x00, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x12, 0x0d, 0x2f,
	0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x12, 0x8d, 0x02, 0x0a,
	0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x36, 0x2e,
	0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x8b,
	0x01, 0x92, 0x41, 0x56, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x6e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x1a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2a,
	0x0c, 0x6c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x4a, 0x0b, 0x0a,
	0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0xca, 0x02, 0x0a,
	0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x73, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xb6, 0x01, 0x92, 0x41, 0x70, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x54,
	0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x1a, 0x1e, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70,
	0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x74, 0x61,
	0x72, 0x67, 0x65, 0x74, 0x73, 0x2a, 0x12, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30,
	0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x34, 0x12, 0x32, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65,
	0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x2a, 0x2f, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x12, 0x90, 0x02, 0x0a, 0x0b, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x36, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x91, 0x01, 0x92, 0x41, 0x52, 0x0a, 0x06,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x16,
	0x4c, 0x69, 0x73, 0x74, 0x20, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2a, 0x0b, 0x6c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b,
	0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x12,
	0x2b, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x2a, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0xfa, 0x01, 0x0a,
	0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12,
	0x38, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x80, 0x01, 0x92, 0x41, 0x38, 0x0a, 0x08, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x20, 0x77,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2a, 0x0e, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41, 0x0f, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x2c, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2d, 0x3a, 0x08, 0x77, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f,
	0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0xd8, 0x01, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x35, 0x2e, 0x63, 0x6f, 0x69, 0x6e,
	0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63,
	0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x22, 0x65, 0x92,
	0x41, 0x32, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x0c, 0x47, 0x65,
	0x74, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2a, 0x0b, 0x67, 0x65, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x23, 0x12, 0x21, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x70, 0x72, 0x6f,
	0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x2f, 0x2a, 0x7d, 0x12, 0xf9, 0x01, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72,
	0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x12, 0x37, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73,
	0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57,
	0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x38, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x75, 0x92, 0x41, 0x40, 0x0a, 0x08,
	0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f,
	0x77, 0x73, 0x2a, 0x0d, 0x6c, 0x69, 0x73, 0x74, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x73, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0xda, 0x41,
	0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x23, 0x12, 0x21, 0x2f,
	0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x73,
	0x12, 0xa9, 0x02, 0x0a, 0x13, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x65, 0x70, 0x12, 0x3d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68,
	0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x65, 0x72,
	0x66, 0x6f, 0x72, 0x6d, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x53, 0x74, 0x65, 0x70,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x22, 0xa5, 0x01, 0x92, 0x41, 0x71, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6b,
	0x66, 0x6c, 0x6f, 0x77, 0x12, 0x23, 0x50, 0x65, 0x72, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x74, 0x68,
	0x65, 0x20, 0x6e, 0x65, 0x78, 0x74, 0x20, 0x73, 0x74, 0x65, 0x70, 0x20, 0x69, 0x6e, 0x20, 0x61,
	0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x1a, 0x23, 0x50, 0x65, 0x72, 0x66, 0x6f,
	0x72, 0x6d, 0x20, 0x74, 0x68, 0x65, 0x20, 0x6e, 0x65, 0x78, 0x74, 0x20, 0x73, 0x74, 0x65, 0x70,
	0x20, 0x69, 0x6e, 0x20, 0x61, 0x20, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x2a, 0x0e,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x4a, 0x0b,
	0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04, 0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x2b, 0x3a, 0x01, 0x2a, 0x22, 0x26, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x3d,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x2a, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x74, 0x65, 0x70, 0x12, 0xe8, 0x02, 0x0a,
	0x12, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74,
	0x65, 0x78, 0x74, 0x12, 0x3c, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3d, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0xd4, 0x01, 0x92, 0x41, 0xad, 0x01, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x12, 0x3c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73,
	0x20, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x69, 0x6e, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x20, 0x64, 0x61, 0x74, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x1a, 0x3c, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x73, 0x20, 0x70,
	0x6f, 0x69, 0x6e, 0x74, 0x2d, 0x69, 0x6e, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x20, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20,
	0x64, 0x61, 0x74, 0x61, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x2a, 0x12, 0x56, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x4a, 0x0b, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x04,
	0x0a, 0x02, 0x4f, 0x4b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d, 0x12, 0x1b, 0x2f, 0x76, 0x31, 0x2f,
	0x76, 0x69, 0x65, 0x77, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e, 0x74, 0x65,
	0x78, 0x74, 0x3a, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x1d, 0xca, 0x41, 0x1a, 0x61, 0x70, 0x69, 0x2e,
	0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x42, 0xe5, 0x07, 0x92, 0x41, 0x99, 0x07, 0x12, 0x65, 0x0a,
	0x15, 0x4f, 0x72, 0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x20,
	0x74, 0x68, 0x61, 0x74, 0x20, 0x63, 0x61, 0x6e, 0x20, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x20, 0x6e,
	0x6f, 0x6e, 0x2d, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6c, 0x20, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x20, 0x65, 0x78, 0x70, 0x65, 0x72, 0x69, 0x65, 0x6e, 0x63, 0x65, 0x73,
	0x20, 0x66, 0x6f, 0x72, 0x20, 0x79, 0x6f, 0x75, 0x72, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2e,
	0x32, 0x02, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x22, 0x08, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2a, 0x01, 0x02, 0x32, 0x10, 0x61,
	0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x3a,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x52, 0x52, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x4b, 0x0a, 0x2c, 0x54, 0x68, 0x65, 0x20,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x65,
	0x64, 0x20, 0x68, 0x61, 0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x70, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f,
	0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x57, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x50, 0x0a, 0x31,
	0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x69, 0x66, 0x20, 0x61, 0x75, 0x74, 0x68,
	0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x64,
	0x0a, 0x03, 0x34, 0x30, 0x33, 0x12, 0x5d, 0x0a, 0x3e, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65,
	0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x64, 0x6f,
	0x65, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x68, 0x61, 0x76, 0x65, 0x20, 0x70, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x20, 0x74, 0x6f, 0x20, 0x74, 0x68, 0x65, 0x20, 0x72, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x52, 0x4c, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x45, 0x0a, 0x26, 0x52,
	0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x69, 0x73, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66,
	0x6f, 0x75, 0x6e, 0x64, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x56, 0x0a, 0x03, 0x34, 0x32, 0x39, 0x12, 0x4f, 0x0a, 0x30, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x20, 0x72, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x20, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x20, 0x68, 0x61, 0x73, 0x20,
	0x62, 0x65, 0x65, 0x6e, 0x20, 0x72, 0x65, 0x61, 0x63, 0x68, 0x65, 0x64, 0x2e, 0x12, 0x1b, 0x0a,
	0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x55, 0x0a, 0x03, 0x35, 0x30,
	0x30, 0x12, 0x4e, 0x0a, 0x2f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68,
	0x65, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x20, 0x68, 0x61, 0x70, 0x70,
	0x65, 0x6e, 0x73, 0x2e, 0x12, 0x1b, 0x0a, 0x19, 0x1a, 0x17, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x72, 0x70, 0x63, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x6a, 0x1d, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x11, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x6a, 0x1b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x10, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x6a, 0x19, 0x0a,
	0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x6a, 0x28, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x6b,
	0x69, 0x6e, 0x67, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x12, 0x17, 0x53, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x20, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x73, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x6a, 0x29, 0x0a, 0x0e, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x17, 0x53, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x20, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x6a, 0x27, 0x0a,
	0x08, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x12, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x66,
	0x6c, 0x6f, 0x77, 0x20, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x20, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x73, 0x74,
	0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x6f, 0x72,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_coinbase_staking_orchestration_v1_api_proto_goTypes = []interface{}{
	(*ListProtocolsRequest)(nil),       // 0: coinbase.staking.orchestration.v1.ListProtocolsRequest
	(*ListNetworksRequest)(nil),        // 1: coinbase.staking.orchestration.v1.ListNetworksRequest
	(*ListStakingTargetsRequest)(nil),  // 2: coinbase.staking.orchestration.v1.ListStakingTargetsRequest
	(*ListActionsRequest)(nil),         // 3: coinbase.staking.orchestration.v1.ListActionsRequest
	(*CreateWorkflowRequest)(nil),      // 4: coinbase.staking.orchestration.v1.CreateWorkflowRequest
	(*GetWorkflowRequest)(nil),         // 5: coinbase.staking.orchestration.v1.GetWorkflowRequest
	(*ListWorkflowsRequest)(nil),       // 6: coinbase.staking.orchestration.v1.ListWorkflowsRequest
	(*PerformWorkflowStepRequest)(nil), // 7: coinbase.staking.orchestration.v1.PerformWorkflowStepRequest
	(*ViewStakingContextRequest)(nil),  // 8: coinbase.staking.orchestration.v1.ViewStakingContextRequest
	(*ListProtocolsResponse)(nil),      // 9: coinbase.staking.orchestration.v1.ListProtocolsResponse
	(*ListNetworksResponse)(nil),       // 10: coinbase.staking.orchestration.v1.ListNetworksResponse
	(*ListStakingTargetsResponse)(nil), // 11: coinbase.staking.orchestration.v1.ListStakingTargetsResponse
	(*ListActionsResponse)(nil),        // 12: coinbase.staking.orchestration.v1.ListActionsResponse
	(*Workflow)(nil),                   // 13: coinbase.staking.orchestration.v1.Workflow
	(*ListWorkflowsResponse)(nil),      // 14: coinbase.staking.orchestration.v1.ListWorkflowsResponse
	(*ViewStakingContextResponse)(nil), // 15: coinbase.staking.orchestration.v1.ViewStakingContextResponse
}
var file_coinbase_staking_orchestration_v1_api_proto_depIdxs = []int32{
	0,  // 0: coinbase.staking.orchestration.v1.StakingService.ListProtocols:input_type -> coinbase.staking.orchestration.v1.ListProtocolsRequest
	1,  // 1: coinbase.staking.orchestration.v1.StakingService.ListNetworks:input_type -> coinbase.staking.orchestration.v1.ListNetworksRequest
	2,  // 2: coinbase.staking.orchestration.v1.StakingService.ListStakingTargets:input_type -> coinbase.staking.orchestration.v1.ListStakingTargetsRequest
	3,  // 3: coinbase.staking.orchestration.v1.StakingService.ListActions:input_type -> coinbase.staking.orchestration.v1.ListActionsRequest
	4,  // 4: coinbase.staking.orchestration.v1.StakingService.CreateWorkflow:input_type -> coinbase.staking.orchestration.v1.CreateWorkflowRequest
	5,  // 5: coinbase.staking.orchestration.v1.StakingService.GetWorkflow:input_type -> coinbase.staking.orchestration.v1.GetWorkflowRequest
	6,  // 6: coinbase.staking.orchestration.v1.StakingService.ListWorkflows:input_type -> coinbase.staking.orchestration.v1.ListWorkflowsRequest
	7,  // 7: coinbase.staking.orchestration.v1.StakingService.PerformWorkflowStep:input_type -> coinbase.staking.orchestration.v1.PerformWorkflowStepRequest
	8,  // 8: coinbase.staking.orchestration.v1.StakingService.ViewStakingContext:input_type -> coinbase.staking.orchestration.v1.ViewStakingContextRequest
	9,  // 9: coinbase.staking.orchestration.v1.StakingService.ListProtocols:output_type -> coinbase.staking.orchestration.v1.ListProtocolsResponse
	10, // 10: coinbase.staking.orchestration.v1.StakingService.ListNetworks:output_type -> coinbase.staking.orchestration.v1.ListNetworksResponse
	11, // 11: coinbase.staking.orchestration.v1.StakingService.ListStakingTargets:output_type -> coinbase.staking.orchestration.v1.ListStakingTargetsResponse
	12, // 12: coinbase.staking.orchestration.v1.StakingService.ListActions:output_type -> coinbase.staking.orchestration.v1.ListActionsResponse
	13, // 13: coinbase.staking.orchestration.v1.StakingService.CreateWorkflow:output_type -> coinbase.staking.orchestration.v1.Workflow
	13, // 14: coinbase.staking.orchestration.v1.StakingService.GetWorkflow:output_type -> coinbase.staking.orchestration.v1.Workflow
	14, // 15: coinbase.staking.orchestration.v1.StakingService.ListWorkflows:output_type -> coinbase.staking.orchestration.v1.ListWorkflowsResponse
	13, // 16: coinbase.staking.orchestration.v1.StakingService.PerformWorkflowStep:output_type -> coinbase.staking.orchestration.v1.Workflow
	15, // 17: coinbase.staking.orchestration.v1.StakingService.ViewStakingContext:output_type -> coinbase.staking.orchestration.v1.ViewStakingContextResponse
	9,  // [9:18] is the sub-list for method output_type
	0,  // [0:9] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_coinbase_staking_orchestration_v1_api_proto_init() }
func file_coinbase_staking_orchestration_v1_api_proto_init() {
	if File_coinbase_staking_orchestration_v1_api_proto != nil {
		return
	}
	file_coinbase_staking_orchestration_v1_protocol_proto_init()
	file_coinbase_staking_orchestration_v1_network_proto_init()
	file_coinbase_staking_orchestration_v1_action_proto_init()
	file_coinbase_staking_orchestration_v1_staking_target_proto_init()
	file_coinbase_staking_orchestration_v1_workflow_proto_init()
	file_coinbase_staking_orchestration_v1_staking_context_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_staking_orchestration_v1_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coinbase_staking_orchestration_v1_api_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_orchestration_v1_api_proto_depIdxs,
	}.Build()
	File_coinbase_staking_orchestration_v1_api_proto = out.File
	file_coinbase_staking_orchestration_v1_api_proto_rawDesc = nil
	file_coinbase_staking_orchestration_v1_api_proto_goTypes = nil
	file_coinbase_staking_orchestration_v1_api_proto_depIdxs = nil
}
