// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: chainrpc/chainkit.proto

package chainrpc

import (
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

type GetBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the requested block.
	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *GetBlockRequest) Reset() {
	*x = GetBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockRequest) ProtoMessage() {}

func (x *GetBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBlockRequest) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{0}
}

func (x *GetBlockRequest) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

// TODO(ffranr): The neutrino GetBlock response includes many
// additional helpful fields. Consider adding them here also.
type GetBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The raw bytes of the requested block.
	RawBlock []byte `protobuf:"bytes,1,opt,name=raw_block,json=rawBlock,proto3" json:"raw_block,omitempty"`
}

func (x *GetBlockResponse) Reset() {
	*x = GetBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockResponse) ProtoMessage() {}

func (x *GetBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockResponse.ProtoReflect.Descriptor instead.
func (*GetBlockResponse) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{1}
}

func (x *GetBlockResponse) GetRawBlock() []byte {
	if x != nil {
		return x.RawBlock
	}
	return nil
}

type GetBlockHeaderRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the block with the requested header.
	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *GetBlockHeaderRequest) Reset() {
	*x = GetBlockHeaderRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderRequest) ProtoMessage() {}

func (x *GetBlockHeaderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderRequest) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{2}
}

func (x *GetBlockHeaderRequest) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

type GetBlockHeaderResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The header of the block with the requested hash.
	RawBlockHeader []byte `protobuf:"bytes,1,opt,name=raw_block_header,json=rawBlockHeader,proto3" json:"raw_block_header,omitempty"`
}

func (x *GetBlockHeaderResponse) Reset() {
	*x = GetBlockHeaderResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHeaderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHeaderResponse) ProtoMessage() {}

func (x *GetBlockHeaderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHeaderResponse.ProtoReflect.Descriptor instead.
func (*GetBlockHeaderResponse) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{3}
}

func (x *GetBlockHeaderResponse) GetRawBlockHeader() []byte {
	if x != nil {
		return x.RawBlockHeader
	}
	return nil
}

type GetBestBlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetBestBlockRequest) Reset() {
	*x = GetBestBlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBestBlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBestBlockRequest) ProtoMessage() {}

func (x *GetBestBlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBestBlockRequest.ProtoReflect.Descriptor instead.
func (*GetBestBlockRequest) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{4}
}

type GetBestBlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the best block.
	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
	// The height of the best block.
	BlockHeight int32 `protobuf:"varint,2,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *GetBestBlockResponse) Reset() {
	*x = GetBestBlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBestBlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBestBlockResponse) ProtoMessage() {}

func (x *GetBestBlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBestBlockResponse.ProtoReflect.Descriptor instead.
func (*GetBestBlockResponse) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{5}
}

func (x *GetBestBlockResponse) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

func (x *GetBestBlockResponse) GetBlockHeight() int32 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

type GetBlockHashRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Block height of the target best chain block.
	BlockHeight int64 `protobuf:"varint,1,opt,name=block_height,json=blockHeight,proto3" json:"block_height,omitempty"`
}

func (x *GetBlockHashRequest) Reset() {
	*x = GetBlockHashRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHashRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHashRequest) ProtoMessage() {}

func (x *GetBlockHashRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHashRequest.ProtoReflect.Descriptor instead.
func (*GetBlockHashRequest) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{6}
}

func (x *GetBlockHashRequest) GetBlockHeight() int64 {
	if x != nil {
		return x.BlockHeight
	}
	return 0
}

type GetBlockHashResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The hash of the best block at the specified height.
	BlockHash []byte `protobuf:"bytes,1,opt,name=block_hash,json=blockHash,proto3" json:"block_hash,omitempty"`
}

func (x *GetBlockHashResponse) Reset() {
	*x = GetBlockHashResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chainrpc_chainkit_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBlockHashResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBlockHashResponse) ProtoMessage() {}

func (x *GetBlockHashResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chainrpc_chainkit_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBlockHashResponse.ProtoReflect.Descriptor instead.
func (*GetBlockHashResponse) Descriptor() ([]byte, []int) {
	return file_chainrpc_chainkit_proto_rawDescGZIP(), []int{7}
}

func (x *GetBlockHashResponse) GetBlockHash() []byte {
	if x != nil {
		return x.BlockHash
	}
	return nil
}

var File_chainrpc_chainkit_proto protoreflect.FileDescriptor

var file_chainrpc_chainkit_proto_rawDesc = []byte{
	0x0a, 0x17, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x6b, 0x69, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x72, 0x70, 0x63, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x72, 0x61, 0x77,
	0x5f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x72, 0x61,
	0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0x42,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x61, 0x77, 0x5f,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0e, 0x72, 0x61, 0x77, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x22, 0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x14, 0x47, 0x65, 0x74,
	0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x22, 0x38, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48,
	0x61, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x22, 0x35, 0x0a,
	0x14, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x61, 0x73, 0x68, 0x32, 0xc0, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x4b, 0x69,
	0x74, 0x12, 0x41, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x19, 0x2e,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x53, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70,
	0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x69,
	0x6e, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x65, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1d, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72,
	0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6c, 0x69, 0x67, 0x68, 0x74, 0x6e, 0x69, 0x6e, 0x67, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x6c, 0x6e, 0x64, 0x2f, 0x6c, 0x6e, 0x72, 0x70, 0x63,
	0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_chainrpc_chainkit_proto_rawDescOnce sync.Once
	file_chainrpc_chainkit_proto_rawDescData = file_chainrpc_chainkit_proto_rawDesc
)

func file_chainrpc_chainkit_proto_rawDescGZIP() []byte {
	file_chainrpc_chainkit_proto_rawDescOnce.Do(func() {
		file_chainrpc_chainkit_proto_rawDescData = protoimpl.X.CompressGZIP(file_chainrpc_chainkit_proto_rawDescData)
	})
	return file_chainrpc_chainkit_proto_rawDescData
}

var file_chainrpc_chainkit_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chainrpc_chainkit_proto_goTypes = []interface{}{
	(*GetBlockRequest)(nil),        // 0: chainrpc.GetBlockRequest
	(*GetBlockResponse)(nil),       // 1: chainrpc.GetBlockResponse
	(*GetBlockHeaderRequest)(nil),  // 2: chainrpc.GetBlockHeaderRequest
	(*GetBlockHeaderResponse)(nil), // 3: chainrpc.GetBlockHeaderResponse
	(*GetBestBlockRequest)(nil),    // 4: chainrpc.GetBestBlockRequest
	(*GetBestBlockResponse)(nil),   // 5: chainrpc.GetBestBlockResponse
	(*GetBlockHashRequest)(nil),    // 6: chainrpc.GetBlockHashRequest
	(*GetBlockHashResponse)(nil),   // 7: chainrpc.GetBlockHashResponse
}
var file_chainrpc_chainkit_proto_depIdxs = []int32{
	0, // 0: chainrpc.ChainKit.GetBlock:input_type -> chainrpc.GetBlockRequest
	2, // 1: chainrpc.ChainKit.GetBlockHeader:input_type -> chainrpc.GetBlockHeaderRequest
	4, // 2: chainrpc.ChainKit.GetBestBlock:input_type -> chainrpc.GetBestBlockRequest
	6, // 3: chainrpc.ChainKit.GetBlockHash:input_type -> chainrpc.GetBlockHashRequest
	1, // 4: chainrpc.ChainKit.GetBlock:output_type -> chainrpc.GetBlockResponse
	3, // 5: chainrpc.ChainKit.GetBlockHeader:output_type -> chainrpc.GetBlockHeaderResponse
	5, // 6: chainrpc.ChainKit.GetBestBlock:output_type -> chainrpc.GetBestBlockResponse
	7, // 7: chainrpc.ChainKit.GetBlockHash:output_type -> chainrpc.GetBlockHashResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chainrpc_chainkit_proto_init() }
func file_chainrpc_chainkit_proto_init() {
	if File_chainrpc_chainkit_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chainrpc_chainkit_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockRequest); i {
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
		file_chainrpc_chainkit_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockResponse); i {
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
		file_chainrpc_chainkit_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderRequest); i {
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
		file_chainrpc_chainkit_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHeaderResponse); i {
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
		file_chainrpc_chainkit_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBestBlockRequest); i {
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
		file_chainrpc_chainkit_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBestBlockResponse); i {
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
		file_chainrpc_chainkit_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHashRequest); i {
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
		file_chainrpc_chainkit_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBlockHashResponse); i {
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
			RawDescriptor: file_chainrpc_chainkit_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chainrpc_chainkit_proto_goTypes,
		DependencyIndexes: file_chainrpc_chainkit_proto_depIdxs,
		MessageInfos:      file_chainrpc_chainkit_proto_msgTypes,
	}.Build()
	File_chainrpc_chainkit_proto = out.File
	file_chainrpc_chainkit_proto_rawDesc = nil
	file_chainrpc_chainkit_proto_goTypes = nil
	file_chainrpc_chainkit_proto_depIdxs = nil
}
