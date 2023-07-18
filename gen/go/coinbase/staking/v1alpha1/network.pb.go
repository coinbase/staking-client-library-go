// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/v1alpha1/network.proto

package v1alpha1

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

// A Network resource, which represents a blockchain network.
// (i.e. mainnet, testnet, etc.)
type Network struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Network.
	// Format: protocols/{protocolName}/networks/{networkName}
	// Ex: protocols/polygon/networks/testnet
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Represents if the network is the mainnet network
	// for the given protocol.
	IsMainnet bool `protobuf:"varint,2,opt,name=is_mainnet,json=isMainnet,proto3" json:"is_mainnet,omitempty"`
}

func (x *Network) Reset() {
	*x = Network{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Network) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Network) ProtoMessage() {}

func (x *Network) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Network.ProtoReflect.Descriptor instead.
func (*Network) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_network_proto_rawDescGZIP(), []int{0}
}

func (x *Network) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Network) GetIsMainnet() bool {
	if x != nil {
		return x.IsMainnet
	}
	return false
}

// The request message for ListNetworks.
type ListNetworksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the parent that owns
	// the collection of networks.
	// Format: protocols/{protocol}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// *UNUSED* The maximum number of networks to return. The service may
	// return fewer than this value.
	//
	// If unspecified, 50 networks will be returned.
	// The maximum value is 1000; values over 1000 will be floored to 1000.
	PageSize int32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// *UNUSED* A page token as part of the response of a previous call.
	// Provide this to retrieve the next page.
	//
	// When paginating, all other parameters must match the previous
	// request to list resources.
	PageToken string `protobuf:"bytes,3,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty"`
}

func (x *ListNetworksRequest) Reset() {
	*x = ListNetworksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksRequest) ProtoMessage() {}

func (x *ListNetworksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksRequest.ProtoReflect.Descriptor instead.
func (*ListNetworksRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_network_proto_rawDescGZIP(), []int{1}
}

func (x *ListNetworksRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

func (x *ListNetworksRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListNetworksRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

// The response message for ListNetworks.
type ListNetworksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of networks.
	Networks []*Network `protobuf:"bytes,1,rep,name=networks,proto3" json:"networks,omitempty"`
	// *UNUSED* A token which can be provided as `page_token` to retrieve the next page.
	// If this field is omitted, there are no additional pages.
	NextPageToken string `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListNetworksResponse) Reset() {
	*x = ListNetworksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListNetworksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListNetworksResponse) ProtoMessage() {}

func (x *ListNetworksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListNetworksResponse.ProtoReflect.Descriptor instead.
func (*ListNetworksResponse) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_network_proto_rawDescGZIP(), []int{2}
}

func (x *ListNetworksResponse) GetNetworks() []*Network {
	if x != nil {
		return x.Networks
	}
	return nil
}

func (x *ListNetworksResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

var File_coinbase_staking_v1alpha1_network_proto protoreflect.FileDescriptor

var file_coinbase_staking_v1alpha1_network_proto_rawDesc = []byte{
	0x0a, 0x27, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x9b, 0x01, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x4d, 0x61, 0x69, 0x6e, 0x6e, 0x65, 0x74, 0x3a,
	0x5d, 0xea, 0x41, 0x5a, 0x0a, 0x1c, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f,
	0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x12, 0x27, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x7d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b,
	0x73, 0x2f, 0x7b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x7d, 0x2a, 0x08, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x73, 0x32, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x22, 0x8f,
	0x01, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x24, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1e, 0x12, 0x1c,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x06, 0x70, 0x61,
	0x72, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x7e, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3e, 0x0a, 0x08, 0x6e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x52, 0x08,
	0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74,
	0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f,
	0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f,
	0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_staking_v1alpha1_network_proto_rawDescOnce sync.Once
	file_coinbase_staking_v1alpha1_network_proto_rawDescData = file_coinbase_staking_v1alpha1_network_proto_rawDesc
)

func file_coinbase_staking_v1alpha1_network_proto_rawDescGZIP() []byte {
	file_coinbase_staking_v1alpha1_network_proto_rawDescOnce.Do(func() {
		file_coinbase_staking_v1alpha1_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_staking_v1alpha1_network_proto_rawDescData)
	})
	return file_coinbase_staking_v1alpha1_network_proto_rawDescData
}

var file_coinbase_staking_v1alpha1_network_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_staking_v1alpha1_network_proto_goTypes = []interface{}{
	(*Network)(nil),              // 0: coinbase.staking.v1alpha1.Network
	(*ListNetworksRequest)(nil),  // 1: coinbase.staking.v1alpha1.ListNetworksRequest
	(*ListNetworksResponse)(nil), // 2: coinbase.staking.v1alpha1.ListNetworksResponse
}
var file_coinbase_staking_v1alpha1_network_proto_depIdxs = []int32{
	0, // 0: coinbase.staking.v1alpha1.ListNetworksResponse.networks:type_name -> coinbase.staking.v1alpha1.Network
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coinbase_staking_v1alpha1_network_proto_init() }
func file_coinbase_staking_v1alpha1_network_proto_init() {
	if File_coinbase_staking_v1alpha1_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_staking_v1alpha1_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Network); i {
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
		file_coinbase_staking_v1alpha1_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksRequest); i {
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
		file_coinbase_staking_v1alpha1_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListNetworksResponse); i {
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
			RawDescriptor: file_coinbase_staking_v1alpha1_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_staking_v1alpha1_network_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_v1alpha1_network_proto_depIdxs,
		MessageInfos:      file_coinbase_staking_v1alpha1_network_proto_msgTypes,
	}.Build()
	File_coinbase_staking_v1alpha1_network_proto = out.File
	file_coinbase_staking_v1alpha1_network_proto_rawDesc = nil
	file_coinbase_staking_v1alpha1_network_proto_goTypes = nil
	file_coinbase_staking_v1alpha1_network_proto_depIdxs = nil
}
