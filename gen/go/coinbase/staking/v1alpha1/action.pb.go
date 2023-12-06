// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: coinbase/staking/v1alpha1/action.proto

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

// An Action resource, which represents an action you may take on a network,
// posted to a validator (e.g. stake, unstake).
type Action struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the Action.
	// Format: protocols/{protocolName}/networks/{networkName}/actions/{actionName}
	// Ex: protocols/polygon/networks/goerli/validators/stake
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Action) Reset() {
	*x = Action{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Action) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Action) ProtoMessage() {}

func (x *Action) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Action.ProtoReflect.Descriptor instead.
func (*Action) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_action_proto_rawDescGZIP(), []int{0}
}

func (x *Action) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// The request message for ListActions.
type ListActionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The resource name of the parent that owns
	// the collection of actions.
	// Format: protocols/{protocol}/networks/{network}
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
}

func (x *ListActionsRequest) Reset() {
	*x = ListActionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActionsRequest) ProtoMessage() {}

func (x *ListActionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActionsRequest.ProtoReflect.Descriptor instead.
func (*ListActionsRequest) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_action_proto_rawDescGZIP(), []int{1}
}

func (x *ListActionsRequest) GetParent() string {
	if x != nil {
		return x.Parent
	}
	return ""
}

// The response message for ListActions.
type ListActionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The list of actions.
	Actions []*Action `protobuf:"bytes,1,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *ListActionsResponse) Reset() {
	*x = ListActionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListActionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActionsResponse) ProtoMessage() {}

func (x *ListActionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_coinbase_staking_v1alpha1_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActionsResponse.ProtoReflect.Descriptor instead.
func (*ListActionsResponse) Descriptor() ([]byte, []int) {
	return file_coinbase_staking_v1alpha1_action_proto_rawDescGZIP(), []int{2}
}

func (x *ListActionsResponse) GetActions() []*Action {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_coinbase_staking_v1alpha1_action_proto protoreflect.FileDescriptor

var file_coinbase_staking_v1alpha1_action_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x66, 0x69, 0x65, 0x6c, 0x64, 0x5f, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x89, 0x01, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x3a, 0x6b,
	0xea, 0x41, 0x68, 0x0a, 0x1b, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x69,
	0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x38, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x73, 0x2f, 0x7b, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x7d, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x2f,
	0x7b, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x7d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x7b, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x7d, 0x2a, 0x07, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x32, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x51, 0x0a, 0x12, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x3b, 0x0a, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x23, 0xe0, 0x41, 0x02, 0xfa, 0x41, 0x1d, 0x12, 0x1b, 0x73, 0x74, 0x61, 0x6b, 0x69,
	0x6e, 0x67, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x22, 0x52,
	0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x50, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e,
	0x67, 0x2d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2d, 0x6c, 0x69, 0x62, 0x72, 0x61, 0x72, 0x79,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x62,
	0x61, 0x73, 0x65, 0x2f, 0x73, 0x74, 0x61, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_coinbase_staking_v1alpha1_action_proto_rawDescOnce sync.Once
	file_coinbase_staking_v1alpha1_action_proto_rawDescData = file_coinbase_staking_v1alpha1_action_proto_rawDesc
)

func file_coinbase_staking_v1alpha1_action_proto_rawDescGZIP() []byte {
	file_coinbase_staking_v1alpha1_action_proto_rawDescOnce.Do(func() {
		file_coinbase_staking_v1alpha1_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_coinbase_staking_v1alpha1_action_proto_rawDescData)
	})
	return file_coinbase_staking_v1alpha1_action_proto_rawDescData
}

var file_coinbase_staking_v1alpha1_action_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_coinbase_staking_v1alpha1_action_proto_goTypes = []interface{}{
	(*Action)(nil),              // 0: coinbase.staking.v1alpha1.Action
	(*ListActionsRequest)(nil),  // 1: coinbase.staking.v1alpha1.ListActionsRequest
	(*ListActionsResponse)(nil), // 2: coinbase.staking.v1alpha1.ListActionsResponse
}
var file_coinbase_staking_v1alpha1_action_proto_depIdxs = []int32{
	0, // 0: coinbase.staking.v1alpha1.ListActionsResponse.actions:type_name -> coinbase.staking.v1alpha1.Action
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_coinbase_staking_v1alpha1_action_proto_init() }
func file_coinbase_staking_v1alpha1_action_proto_init() {
	if File_coinbase_staking_v1alpha1_action_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_coinbase_staking_v1alpha1_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Action); i {
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
		file_coinbase_staking_v1alpha1_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActionsRequest); i {
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
		file_coinbase_staking_v1alpha1_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListActionsResponse); i {
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
			RawDescriptor: file_coinbase_staking_v1alpha1_action_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_coinbase_staking_v1alpha1_action_proto_goTypes,
		DependencyIndexes: file_coinbase_staking_v1alpha1_action_proto_depIdxs,
		MessageInfos:      file_coinbase_staking_v1alpha1_action_proto_msgTypes,
	}.Build()
	File_coinbase_staking_v1alpha1_action_proto = out.File
	file_coinbase_staking_v1alpha1_action_proto_rawDesc = nil
	file_coinbase_staking_v1alpha1_action_proto_goTypes = nil
	file_coinbase_staking_v1alpha1_action_proto_depIdxs = nil
}
