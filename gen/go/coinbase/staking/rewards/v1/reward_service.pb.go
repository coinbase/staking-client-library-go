// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/rewards/v1/reward_service.proto

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

var File_coinbase_staking_rewards_v1_reward_service_proto protoreflect.FileDescriptor

var file_coinbase_staking_rewards_v1_reward_service_proto_rawDesc = []byte{
	0x0a, 0x30, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x1b, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x1a,
	0x28, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61,
	0x72, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f,
	0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x84, 0x0e, 0x0a, 0x0d, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xdd, 0x09, 0x0a, 0x0b,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0x2f, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xea,
	0x08, 0x92, 0x41, 0xb5, 0x08, 0x0a, 0x06, 0x52, 0x65, 0x77, 0x61, 0x72, 0x64, 0x12, 0x17, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x72,
	0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x1a, 0x80, 0x01, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x20, 0x6f,
	0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x20, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x20, 0x6f,
	0x66, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x20, 0x66, 0x6f, 0x72,
	0x20, 0x61, 0x20, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x20, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2c, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20,
	0x74, 0x69, 0x6d, 0x65, 0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x2c, 0x20, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x2c, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x6d, 0x6f, 0x72, 0x65, 0x2e, 0x4a, 0xf5, 0x05, 0x0a, 0x03, 0x32, 0x30,
	0x30, 0x12, 0xed, 0x05, 0x0a, 0x02, 0x4f, 0x4b, 0x12, 0xe6, 0x05, 0x32, 0xe3, 0x05, 0x7b, 0x22,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x22, 0x3a, 0x5b, 0x7b, 0x22, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x22, 0x3a, 0x22, 0x62, 0x65, 0x65, 0x66, 0x4b, 0x47, 0x42, 0x57, 0x65, 0x53,
	0x70, 0x48, 0x7a, 0x59, 0x42, 0x48, 0x5a, 0x58, 0x77, 0x70, 0x35, 0x53, 0x6f, 0x37, 0x77, 0x64,
	0x51, 0x47, 0x58, 0x36, 0x6d, 0x75, 0x34, 0x5a, 0x48, 0x43, 0x73, 0x48, 0x33, 0x75, 0x54, 0x61,
	0x72, 0x22, 0x2c, 0x22, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x3a, 0x22, 0x35, 0x33, 0x33, 0x22,
	0x2c, 0x22, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x6e, 0x69,
	0x74, 0x22, 0x3a, 0x22, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x2c, 0x22, 0x70, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x6e, 0x75, 0x6c,
	0x6c, 0x2c, 0x22, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x31, 0x31, 0x2d, 0x31, 0x36, 0x54, 0x30, 0x30,
	0x3a, 0x31, 0x33, 0x3a, 0x34, 0x34, 0x5a, 0x22, 0x2c, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45,
	0x61, 0x72, 0x6e, 0x65, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x22,
	0x3a, 0x7b, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a, 0x22, 0x32, 0x32, 0x34, 0x2e,
	0x37, 0x30, 0x39, 0x38, 0x31, 0x34, 0x35, 0x22, 0x2c, 0x22, 0x65, 0x78, 0x70, 0x22, 0x3a, 0x22,
	0x39, 0x22, 0x2c, 0x22, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22, 0x3a, 0x22, 0x53, 0x4f, 0x4c,
	0x22, 0x2c, 0x22, 0x72, 0x61, 0x77, 0x4e, 0x75, 0x6d, 0x65, 0x72, 0x69, 0x63, 0x22, 0x3a, 0x22,
	0x32, 0x32, 0x34, 0x37, 0x30, 0x39, 0x38, 0x31, 0x34, 0x35, 0x30, 0x39, 0x22, 0x7d, 0x2c, 0x22,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x55, 0x73, 0x64, 0x22, 0x3a,
	0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0x22, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x6f, 0x6c, 0x22, 0x3a, 0x22, 0x73, 0x6f, 0x6c, 0x61, 0x6e, 0x61, 0x22, 0x7d, 0x2c, 0x7b,
	0x22, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x3a, 0x22, 0x62, 0x65, 0x65, 0x66, 0x4b,
	0x47, 0x42, 0x57, 0x65, 0x53, 0x70, 0x48, 0x7a, 0x59, 0x42, 0x48, 0x5a, 0x58, 0x77, 0x70, 0x35,
	0x53, 0x6f, 0x37, 0x77, 0x64, 0x51, 0x47, 0x58, 0x36, 0x6d, 0x75, 0x34, 0x5a, 0x48, 0x43, 0x73,
	0x48, 0x33, 0x75, 0x54, 0x61, 0x72, 0x22, 0x2c, 0x22, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x3a,
	0x22, 0x35, 0x33, 0x32, 0x22, 0x2c, 0x22, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x6e, 0x69, 0x74, 0x22, 0x3a, 0x22, 0x65, 0x70, 0x6f, 0x63, 0x68, 0x22, 0x2c,
	0x22, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65,
	0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0x22, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e,
	0x64, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x3a, 0x22, 0x32, 0x30, 0x32, 0x33, 0x2d, 0x31, 0x31, 0x2d,
	0x31, 0x33, 0x54, 0x31, 0x39, 0x3a, 0x33, 0x38, 0x3a, 0x33, 0x36, 0x5a, 0x22, 0x2c, 0x22, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64, 0x4e, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x55, 0x6e, 0x69, 0x74, 0x22, 0x3a, 0x7b, 0x22, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x3a,
	0x22, 0x32, 0x32, 0x35, 0x2e, 0x30, 0x37, 0x39, 0x34, 0x32, 0x34, 0x31, 0x22, 0x2c, 0x22, 0x65,
	0x78, 0x70, 0x22, 0x3a, 0x22, 0x39, 0x22, 0x2c, 0x22, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x72, 0x22,
	0x3a, 0x22, 0x53, 0x4f, 0x4c, 0x22, 0x2c, 0x22, 0x72, 0x61, 0x77, 0x4e, 0x75, 0x6d, 0x65, 0x72,
	0x69, 0x63, 0x22, 0x3a, 0x22, 0x32, 0x32, 0x35, 0x30, 0x37, 0x39, 0x34, 0x32, 0x34, 0x30, 0x39,
	0x34, 0x22, 0x7d, 0x2c, 0x22, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x45, 0x61, 0x72, 0x6e, 0x65, 0x64,
	0x55, 0x73, 0x64, 0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0x22, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x22, 0x3a, 0x6e, 0x75, 0x6c, 0x6c, 0x2c, 0x22,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x22, 0x3a, 0x22, 0x73, 0x6f, 0x6c, 0x61, 0x6e,
	0x61, 0x22, 0x7d, 0x5d, 0x2c, 0x22, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x22, 0x3a, 0x22, 0x56, 0x41, 0x71, 0x6c, 0x2d, 0x77, 0x74, 0x64, 0x69, 0x4a,
	0x57, 0x6b, 0x57, 0x49, 0x49, 0x39, 0x62, 0x4a, 0x42, 0x44, 0x6e, 0x45, 0x39, 0x6f, 0x45, 0x63,
	0x2d, 0x38, 0x49, 0x6c, 0x67, 0x55, 0x30, 0x44, 0x74, 0x4b, 0x62, 0x78, 0x53, 0x44, 0x74, 0x42,
	0x67, 0x3d, 0x3a, 0x31, 0x3a, 0x31, 0x37, 0x30, 0x30, 0x32, 0x34, 0x31, 0x32, 0x37, 0x37, 0x22,
	0x7d, 0x4a, 0x96, 0x01, 0x0a, 0x03, 0x34, 0x30, 0x30, 0x12, 0x8e, 0x01, 0x0a, 0x2c, 0x54, 0x68,
	0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x20, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x65, 0x64, 0x20, 0x68, 0x61, 0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20,
	0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x5e, 0x32, 0x5c, 0x7b, 0x22,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x33, 0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x22, 0x3a, 0x22, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x2e, 0x20, 0x3c, 0x52, 0x65,
	0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74,
	0x61, 0x6e, 0x63, 0x65, 0x20, 0x68, 0x65, 0x72, 0x65, 0x3e, 0x2e, 0x22, 0x2c, 0x22, 0x64, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3a, 0x5b, 0x5d, 0x7d, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72,
	0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x22, 0x12, 0x20, 0x2f, 0x76, 0x31, 0x2f, 0x7b,
	0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73,
	0x2f, 0x2a, 0x7d, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x12, 0xf3, 0x03, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x72, 0x65,
	0x77, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61,
	0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x03, 0x92, 0x41,
	0xcf, 0x02, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x12, 0x20, 0x4c, 0x69, 0x73, 0x74, 0x20,
	0x61, 0x6e, 0x64, 0x20, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x1a, 0x56, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x20, 0x62, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2c, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x20,
	0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x20, 0x66, 0x6f, 0x72, 0x20, 0x74, 0x69, 0x6d, 0x65,
	0x20, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x20, 0x61, 0x6e, 0x64, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x2e, 0x4a, 0x33, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x2c, 0x0a, 0x02, 0x4f, 0x4b,
	0x12, 0x26, 0x0a, 0x24, 0x1a, 0x22, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x74, 0x61, 0x6b, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x96, 0x01, 0x0a, 0x03, 0x34, 0x30, 0x30,
	0x12, 0x8e, 0x01, 0x0a, 0x2c, 0x54, 0x68, 0x65, 0x20, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x20, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x65, 0x64, 0x20, 0x68, 0x61, 0x73, 0x20, 0x69,
	0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x12, 0x5e, 0x32, 0x5c, 0x7b, 0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x33, 0x2c, 0x22,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x3a, 0x22, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72,
	0x20, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x66, 0x61, 0x69, 0x6c,
	0x65, 0x64, 0x2e, 0x20, 0x3c, 0x52, 0x65, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x20, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x20, 0x68, 0x65, 0x72, 0x65,
	0x3e, 0x2e, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x3a, 0x5b, 0x5d,
	0x7d, 0xda, 0x41, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x21,
	0x12, 0x1f, 0x2f, 0x76, 0x31, 0x2f, 0x7b, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x3d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x65,
	0x73, 0x1a, 0x1d, 0xca, 0x41, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d,
	0x42, 0xa1, 0x05, 0x92, 0x41, 0xd1, 0x04, 0x12, 0x5d, 0x0a, 0x0f, 0x52, 0x65, 0x77, 0x61, 0x72,
	0x64, 0x73, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x46, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x20, 0x74, 0x68, 0x61, 0x74, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x73, 0x20, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x6f, 0x6e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x2c, 0x20, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x72, 0x65, 0x6c,
	0x61, 0x74, 0x65, 0x64, 0x20, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x20, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x32, 0x02, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x65, 0x76, 0x65,
	0x6c, 0x6f, 0x70, 0x65, 0x72, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x22, 0x08, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2a, 0x01, 0x02, 0x32,
	0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f,
	0x6e, 0x3a, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a,
	0x73, 0x6f, 0x6e, 0x52, 0x4c, 0x0a, 0x03, 0x34, 0x30, 0x31, 0x12, 0x45, 0x0a, 0x31, 0x52, 0x65,
	0x74, 0x75, 0x72, 0x6e, 0x65, 0x64, 0x20, 0x69, 0x66, 0x20, 0x61, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x20, 0x69, 0x73, 0x20, 0x69, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x12,
	0x10, 0x32, 0x0e, 0x22, 0x55, 0x6e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x22, 0x52, 0x76, 0x0a, 0x03, 0x35, 0x30, 0x30, 0x12, 0x6f, 0x0a, 0x2f, 0x52, 0x65, 0x74, 0x75,
	0x72, 0x6e, 0x65, 0x64, 0x20, 0x77, 0x68, 0x65, 0x6e, 0x20, 0x61, 0x6e, 0x20, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x20, 0x68, 0x61, 0x70, 0x70, 0x65, 0x6e, 0x73, 0x2e, 0x12, 0x3c, 0x32, 0x3a, 0x7b,
	0x22, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x3a, 0x33, 0x2c, 0x22, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0x3a, 0x22, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x2e, 0x22, 0x2c, 0x22, 0x64, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x22, 0x3a, 0x5b, 0x5d, 0x7d, 0x6a, 0x6c, 0x0a, 0x06, 0x52, 0x65, 0x77,
	0x61, 0x72, 0x64, 0x12, 0x62, 0x41, 0x20, 0x68, 0x69, 0x67, 0x68, 0x2d, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x20, 0x76, 0x69, 0x65, 0x77, 0x20, 0x6f, 0x66, 0x20, 0x61, 0x6e, 0x20, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x27, 0x73, 0x20, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x20, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x64, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x73,
	0x6f, 0x6d, 0x65, 0x20, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x69,
	0x6d, 0x65, 0x20, 0x28, 0x65, 0x78, 0x3a, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x20, 0x61, 0x6e, 0x20,
	0x45, 0x70, 0x6f, 0x63, 0x68, 0x29, 0x2e, 0x6a, 0x6f, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x12, 0x66, 0x41, 0x20, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x20, 0x6f, 0x66, 0x20,
	0x61, 0x6e, 0x20, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x27, 0x73, 0x20, 0x73, 0x74, 0x61,
	0x6b, 0x69, 0x6e, 0x67, 0x2d, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x20, 0x62, 0x61, 0x6c,
	0x61, 0x6e, 0x63, 0x65, 0x20, 0x61, 0x74, 0x20, 0x61, 0x20, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63,
	0x75, 0x6c, 0x61, 0x72, 0x20, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x69,
	0x6d, 0x65, 0x2e, 0x20, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x20, 0x63, 0x6f, 0x6d, 0x69,
	0x6e, 0x67, 0x20, 0x73, 0x6f, 0x6f, 0x6e, 0x2e, 0x5a, 0x4a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x62, 0x68, 0x71, 0x2e, 0x6e, 0x65, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f,
	0x72, 0x65, 0x77, 0x61, 0x72, 0x64, 0x73, 0x2d, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x69, 0x6e,
	0x67, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x72, 0x65, 0x77, 0x61, 0x72, 0x64,
	0x73, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_coinbase_staking_rewards_v1_reward_service_proto_goTypes = []interface{}{
	(*ListRewardsRequest)(nil),  // 0: coinbase.staking.rewards.v1.ListRewardsRequest
	(*ListStakesRequest)(nil),   // 1: coinbase.staking.rewards.v1.ListStakesRequest
	(*ListRewardsResponse)(nil), // 2: coinbase.staking.rewards.v1.ListRewardsResponse
	(*ListStakesResponse)(nil),  // 3: coinbase.staking.rewards.v1.ListStakesResponse
}
var file_coinbase_staking_rewards_v1_reward_service_proto_depIdxs = []int32{
	0, // 0: coinbase.staking.rewards.v1.RewardService.ListRewards:input_type -> coinbase.staking.rewards.v1.ListRewardsRequest
	1, // 1: coinbase.staking.rewards.v1.RewardService.ListStakes:input_type -> coinbase.staking.rewards.v1.ListStakesRequest
	2, // 2: coinbase.staking.rewards.v1.RewardService.ListRewards:output_type -> coinbase.staking.rewards.v1.ListRewardsResponse
	3, // 3: coinbase.staking.rewards.v1.RewardService.ListStakes:output_type -> coinbase.staking.rewards.v1.ListStakesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_coinbase_staking_rewards_v1_reward_service_proto_init() }
func file_coinbase_staking_rewards_v1_reward_service_proto_init() {
	if File_coinbase_staking_rewards_v1_reward_service_proto != nil {
		return
	}
	file_coinbase_staking_rewards_v1_reward_proto_init()
	file_coinbase_staking_rewards_v1_stake_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_coinbase_staking_rewards_v1_reward_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_coinbase_staking_rewards_v1_reward_service_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_rewards_v1_reward_service_proto_depIdxs,
	}.Build()
	File_coinbase_staking_rewards_v1_reward_service_proto = out.File
	file_coinbase_staking_rewards_v1_reward_service_proto_rawDesc = nil
	file_coinbase_staking_rewards_v1_reward_service_proto_goTypes = nil
	file_coinbase_staking_rewards_v1_reward_service_proto_depIdxs = nil
}
