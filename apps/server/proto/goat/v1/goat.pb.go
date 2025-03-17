// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: goat/v1/goat.proto

package goatv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Vote int32

const (
	Vote_YES Vote = 0
	Vote_NO  Vote = 1
)

// Enum value maps for Vote.
var (
	Vote_name = map[int32]string{
		0: "YES",
		1: "NO",
	}
	Vote_value = map[string]int32{
		"YES": 0,
		"NO":  1,
	}
)

func (x Vote) Enum() *Vote {
	p := new(Vote)
	*p = x
	return p
}

func (x Vote) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Vote) Descriptor() protoreflect.EnumDescriptor {
	return file_goat_v1_goat_proto_enumTypes[0].Descriptor()
}

func (Vote) Type() protoreflect.EnumType {
	return &file_goat_v1_goat_proto_enumTypes[0]
}

func (x Vote) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Vote.Descriptor instead.
func (Vote) EnumDescriptor() ([]byte, []int) {
	return file_goat_v1_goat_proto_rawDescGZIP(), []int{0}
}

type VoteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Vote          Vote                   `protobuf:"varint,1,opt,name=Vote,proto3,enum=goat.v1.Vote" json:"Vote,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteRequest) Reset() {
	*x = VoteRequest{}
	mi := &file_goat_v1_goat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteRequest) ProtoMessage() {}

func (x *VoteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goat_v1_goat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteRequest.ProtoReflect.Descriptor instead.
func (*VoteRequest) Descriptor() ([]byte, []int) {
	return file_goat_v1_goat_proto_rawDescGZIP(), []int{0}
}

func (x *VoteRequest) GetVote() Vote {
	if x != nil {
		return x.Vote
	}
	return Vote_YES
}

type GetVotesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesRequest) Reset() {
	*x = GetVotesRequest{}
	mi := &file_goat_v1_goat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesRequest) ProtoMessage() {}

func (x *GetVotesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_goat_v1_goat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesRequest.ProtoReflect.Descriptor instead.
func (*GetVotesRequest) Descriptor() ([]byte, []int) {
	return file_goat_v1_goat_proto_rawDescGZIP(), []int{1}
}

type GetVotesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Yes           int64                  `protobuf:"varint,1,opt,name=Yes,proto3" json:"Yes,omitempty"`
	No            int64                  `protobuf:"varint,2,opt,name=No,proto3" json:"No,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetVotesResponse) Reset() {
	*x = GetVotesResponse{}
	mi := &file_goat_v1_goat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetVotesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVotesResponse) ProtoMessage() {}

func (x *GetVotesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goat_v1_goat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVotesResponse.ProtoReflect.Descriptor instead.
func (*GetVotesResponse) Descriptor() ([]byte, []int) {
	return file_goat_v1_goat_proto_rawDescGZIP(), []int{2}
}

func (x *GetVotesResponse) GetYes() int64 {
	if x != nil {
		return x.Yes
	}
	return 0
}

func (x *GetVotesResponse) GetNo() int64 {
	if x != nil {
		return x.No
	}
	return 0
}

type VoteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=Success,proto3" json:"Success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *VoteResponse) Reset() {
	*x = VoteResponse{}
	mi := &file_goat_v1_goat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *VoteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VoteResponse) ProtoMessage() {}

func (x *VoteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_goat_v1_goat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VoteResponse.ProtoReflect.Descriptor instead.
func (*VoteResponse) Descriptor() ([]byte, []int) {
	return file_goat_v1_goat_proto_rawDescGZIP(), []int{3}
}

func (x *VoteResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_goat_v1_goat_proto protoreflect.FileDescriptor

var file_goat_v1_goat_proto_rawDesc = string([]byte{
	0x0a, 0x12, 0x67, 0x6f, 0x61, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x6f, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x6f, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x22, 0x30, 0x0a,
	0x0b, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x04,
	0x56, 0x6f, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0d, 0x2e, 0x67, 0x6f, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x22,
	0x11, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x34, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x59, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x03, 0x59, 0x65, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x4e, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x4e, 0x6f, 0x22, 0x28, 0x0a, 0x0c, 0x56, 0x6f, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2a, 0x17, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x07, 0x0a, 0x03, 0x59, 0x45,
	0x53, 0x10, 0x00, 0x12, 0x06, 0x0a, 0x02, 0x4e, 0x4f, 0x10, 0x01, 0x32, 0x87, 0x01, 0x0a, 0x0b,
	0x47, 0x6f, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x56,
	0x6f, 0x74, 0x65, 0x12, 0x14, 0x2e, 0x67, 0x6f, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x6f, 0x61, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x18,
	0x2e, 0x67, 0x6f, 0x61, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x67, 0x6f, 0x61, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x70, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x61,
	0x74, 0x2e, 0x76, 0x31, 0x42, 0x09, 0x47, 0x6f, 0x61, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x19, 0x67, 0x6f, 0x61, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x61, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x6f, 0x61, 0x74, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47,
	0x58, 0x58, 0xaa, 0x02, 0x07, 0x47, 0x6f, 0x61, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x47,
	0x6f, 0x61, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x13, 0x47, 0x6f, 0x61, 0x74, 0x5c, 0x56, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x47,
	0x6f, 0x61, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_goat_v1_goat_proto_rawDescOnce sync.Once
	file_goat_v1_goat_proto_rawDescData []byte
)

func file_goat_v1_goat_proto_rawDescGZIP() []byte {
	file_goat_v1_goat_proto_rawDescOnce.Do(func() {
		file_goat_v1_goat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_goat_v1_goat_proto_rawDesc), len(file_goat_v1_goat_proto_rawDesc)))
	})
	return file_goat_v1_goat_proto_rawDescData
}

var file_goat_v1_goat_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_goat_v1_goat_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_goat_v1_goat_proto_goTypes = []any{
	(Vote)(0),                // 0: goat.v1.Vote
	(*VoteRequest)(nil),      // 1: goat.v1.VoteRequest
	(*GetVotesRequest)(nil),  // 2: goat.v1.GetVotesRequest
	(*GetVotesResponse)(nil), // 3: goat.v1.GetVotesResponse
	(*VoteResponse)(nil),     // 4: goat.v1.VoteResponse
}
var file_goat_v1_goat_proto_depIdxs = []int32{
	0, // 0: goat.v1.VoteRequest.Vote:type_name -> goat.v1.Vote
	1, // 1: goat.v1.GoatService.Vote:input_type -> goat.v1.VoteRequest
	2, // 2: goat.v1.GoatService.GetVotes:input_type -> goat.v1.GetVotesRequest
	4, // 3: goat.v1.GoatService.Vote:output_type -> goat.v1.VoteResponse
	3, // 4: goat.v1.GoatService.GetVotes:output_type -> goat.v1.GetVotesResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_goat_v1_goat_proto_init() }
func file_goat_v1_goat_proto_init() {
	if File_goat_v1_goat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_goat_v1_goat_proto_rawDesc), len(file_goat_v1_goat_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_goat_v1_goat_proto_goTypes,
		DependencyIndexes: file_goat_v1_goat_proto_depIdxs,
		EnumInfos:         file_goat_v1_goat_proto_enumTypes,
		MessageInfos:      file_goat_v1_goat_proto_msgTypes,
	}.Build()
	File_goat_v1_goat_proto = out.File
	file_goat_v1_goat_proto_goTypes = nil
	file_goat_v1_goat_proto_depIdxs = nil
}
