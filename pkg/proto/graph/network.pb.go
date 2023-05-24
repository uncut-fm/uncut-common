// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: graph/network.proto

package graph

import (
	user "github.com/uncut-fm/uncut-common/pkg/proto/auth/user"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetNetworkByUserIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string            `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Pagination *OffsetPagination `protobuf:"bytes,2,opt,name=pagination,proto3,oneof" json:"pagination,omitempty"`
}

func (x *GetNetworkByUserIDRequest) Reset() {
	*x = GetNetworkByUserIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetNetworkByUserIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetNetworkByUserIDRequest) ProtoMessage() {}

func (x *GetNetworkByUserIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetNetworkByUserIDRequest.ProtoReflect.Descriptor instead.
func (*GetNetworkByUserIDRequest) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{0}
}

func (x *GetNetworkByUserIDRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetNetworkByUserIDRequest) GetPagination() *OffsetPagination {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type OffsetPagination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Limit  int32 `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset int32 `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *OffsetPagination) Reset() {
	*x = OffsetPagination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OffsetPagination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OffsetPagination) ProtoMessage() {}

func (x *OffsetPagination) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OffsetPagination.ProtoReflect.Descriptor instead.
func (*OffsetPagination) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{1}
}

func (x *OffsetPagination) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *OffsetPagination) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type NetworkMembersInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int32                `protobuf:"varint,1,opt,name=totalCount,proto3" json:"totalCount,omitempty"`
	Users      []*NetworkMemberUser `protobuf:"bytes,2,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *NetworkMembersInfoResponse) Reset() {
	*x = NetworkMembersInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMembersInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMembersInfoResponse) ProtoMessage() {}

func (x *NetworkMembersInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMembersInfoResponse.ProtoReflect.Descriptor instead.
func (*NetworkMembersInfoResponse) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{2}
}

func (x *NetworkMembersInfoResponse) GetTotalCount() int32 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

func (x *NetworkMembersInfoResponse) GetUsers() []*NetworkMemberUser {
	if x != nil {
		return x.Users
	}
	return nil
}

type NetworkMemberUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User         *user.User             `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	NftsInCommon []*Nft                 `protobuf:"bytes,2,rep,name=nftsInCommon,proto3" json:"nftsInCommon,omitempty"`
	JoinedOn     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=joinedOn,proto3" json:"joinedOn,omitempty"`
}

func (x *NetworkMemberUser) Reset() {
	*x = NetworkMemberUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkMemberUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkMemberUser) ProtoMessage() {}

func (x *NetworkMemberUser) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkMemberUser.ProtoReflect.Descriptor instead.
func (*NetworkMemberUser) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{3}
}

func (x *NetworkMemberUser) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *NetworkMemberUser) GetNftsInCommon() []*Nft {
	if x != nil {
		return x.NftsInCommon
	}
	return nil
}

func (x *NetworkMemberUser) GetJoinedOn() *timestamppb.Timestamp {
	if x != nil {
		return x.JoinedOn
	}
	return nil
}

type Nft struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int32       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Price   float32     `protobuf:"fixed32,2,opt,name=price,proto3" json:"price,omitempty"`
	Name    string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	TokenId string      `protobuf:"bytes,4,opt,name=tokenId,proto3" json:"tokenId,omitempty"`
	Owners  []*NftOwner `protobuf:"bytes,5,rep,name=owners,proto3" json:"owners,omitempty"`
	Creator *user.User  `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *Nft) Reset() {
	*x = Nft{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Nft) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Nft) ProtoMessage() {}

func (x *Nft) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Nft.ProtoReflect.Descriptor instead.
func (*Nft) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{4}
}

func (x *Nft) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Nft) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Nft) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Nft) GetTokenId() string {
	if x != nil {
		return x.TokenId
	}
	return ""
}

func (x *Nft) GetOwners() []*NftOwner {
	if x != nil {
		return x.Owners
	}
	return nil
}

func (x *Nft) GetCreator() *user.User {
	if x != nil {
		return x.Creator
	}
	return nil
}

type NftOwner struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User              *user.User `protobuf:"bytes,1,opt,name=user,proto3,oneof" json:"user,omitempty"`
	Nft               *Nft       `protobuf:"bytes,2,opt,name=nft,proto3" json:"nft,omitempty"`
	UserWalletAddress string     `protobuf:"bytes,3,opt,name=userWalletAddress,proto3" json:"userWalletAddress,omitempty"`
	Balance           int32      `protobuf:"varint,4,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *NftOwner) Reset() {
	*x = NftOwner{}
	if protoimpl.UnsafeEnabled {
		mi := &file_graph_network_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NftOwner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NftOwner) ProtoMessage() {}

func (x *NftOwner) ProtoReflect() protoreflect.Message {
	mi := &file_graph_network_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NftOwner.ProtoReflect.Descriptor instead.
func (*NftOwner) Descriptor() ([]byte, []int) {
	return file_graph_network_proto_rawDescGZIP(), []int{5}
}

func (x *NftOwner) GetUser() *user.User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *NftOwner) GetNft() *Nft {
	if x != nil {
		return x.Nft
	}
	return nil
}

func (x *NftOwner) GetUserWalletAddress() string {
	if x != nil {
		return x.UserWalletAddress
	}
	return ""
}

func (x *NftOwner) GetBalance() int32 {
	if x != nil {
		return x.Balance
	}
	return 0
}

var File_graph_network_proto protoreflect.FileDescriptor

var file_graph_network_proto_rawDesc = []byte{
	0x0a, 0x13, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x61, 0x70, 0x68, 0x1a, 0x14, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x81, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f,
	0x72, 0x6b, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3c, 0x0a, 0x0a, 0x70, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17,
	0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x50, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x61, 0x67,
	0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x40, 0x0a, 0x10, 0x4f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x22, 0x6c, 0x0a, 0x1a, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x9b, 0x01, 0x0a, 0x11, 0x4e, 0x65, 0x74, 0x77,
	0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2e, 0x0a,
	0x0c, 0x6e, 0x66, 0x74, 0x73, 0x49, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x66, 0x74, 0x52,
	0x0c, 0x6e, 0x66, 0x74, 0x73, 0x49, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x12, 0x36, 0x0a,
	0x08, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x4f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x6a, 0x6f, 0x69,
	0x6e, 0x65, 0x64, 0x4f, 0x6e, 0x22, 0xa8, 0x01, 0x0a, 0x03, 0x4e, 0x66, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49,
	0x64, 0x12, 0x27, 0x0a, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x66, 0x74, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x52, 0x06, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x73, 0x12, 0x24, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72,
	0x22, 0x9e, 0x01, 0x0a, 0x08, 0x4e, 0x66, 0x74, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x23, 0x0a,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x48, 0x00, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x88,
	0x01, 0x01, 0x12, 0x1c, 0x0a, 0x03, 0x6e, 0x66, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0a, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x66, 0x74, 0x52, 0x03, 0x6e, 0x66, 0x74,
	0x12, 0x2c, 0x0a, 0x11, 0x75, 0x73, 0x65, 0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x75, 0x73, 0x65,
	0x72, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18,
	0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x75, 0x73, 0x65,
	0x72, 0x32, 0x66, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x5b, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x44, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x47, 0x65, 0x74, 0x4e, 0x65,
	0x74, 0x77, 0x6f, 0x72, 0x6b, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x61, 0x70, 0x68, 0x2e, 0x4e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x11, 0x5a, 0x0f, 0x70, 0x6b, 0x67,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x61, 0x70, 0x68, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_graph_network_proto_rawDescOnce sync.Once
	file_graph_network_proto_rawDescData = file_graph_network_proto_rawDesc
)

func file_graph_network_proto_rawDescGZIP() []byte {
	file_graph_network_proto_rawDescOnce.Do(func() {
		file_graph_network_proto_rawDescData = protoimpl.X.CompressGZIP(file_graph_network_proto_rawDescData)
	})
	return file_graph_network_proto_rawDescData
}

var file_graph_network_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_graph_network_proto_goTypes = []interface{}{
	(*GetNetworkByUserIDRequest)(nil),  // 0: graph.GetNetworkByUserIDRequest
	(*OffsetPagination)(nil),           // 1: graph.OffsetPagination
	(*NetworkMembersInfoResponse)(nil), // 2: graph.NetworkMembersInfoResponse
	(*NetworkMemberUser)(nil),          // 3: graph.NetworkMemberUser
	(*Nft)(nil),                        // 4: graph.Nft
	(*NftOwner)(nil),                   // 5: graph.NftOwner
	(*user.User)(nil),                  // 6: user.User
	(*timestamppb.Timestamp)(nil),      // 7: google.protobuf.Timestamp
}
var file_graph_network_proto_depIdxs = []int32{
	1,  // 0: graph.GetNetworkByUserIDRequest.pagination:type_name -> graph.OffsetPagination
	3,  // 1: graph.NetworkMembersInfoResponse.users:type_name -> graph.NetworkMemberUser
	6,  // 2: graph.NetworkMemberUser.user:type_name -> user.User
	4,  // 3: graph.NetworkMemberUser.nftsInCommon:type_name -> graph.Nft
	7,  // 4: graph.NetworkMemberUser.joinedOn:type_name -> google.protobuf.Timestamp
	5,  // 5: graph.Nft.owners:type_name -> graph.NftOwner
	6,  // 6: graph.Nft.creator:type_name -> user.User
	6,  // 7: graph.NftOwner.user:type_name -> user.User
	4,  // 8: graph.NftOwner.nft:type_name -> graph.Nft
	0,  // 9: graph.Network.GetNetworkByUserID:input_type -> graph.GetNetworkByUserIDRequest
	2,  // 10: graph.Network.GetNetworkByUserID:output_type -> graph.NetworkMembersInfoResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_graph_network_proto_init() }
func file_graph_network_proto_init() {
	if File_graph_network_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_graph_network_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetNetworkByUserIDRequest); i {
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
		file_graph_network_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OffsetPagination); i {
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
		file_graph_network_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMembersInfoResponse); i {
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
		file_graph_network_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkMemberUser); i {
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
		file_graph_network_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Nft); i {
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
		file_graph_network_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NftOwner); i {
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
	file_graph_network_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_graph_network_proto_msgTypes[5].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_graph_network_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_graph_network_proto_goTypes,
		DependencyIndexes: file_graph_network_proto_depIdxs,
		MessageInfos:      file_graph_network_proto_msgTypes,
	}.Build()
	File_graph_network_proto = out.File
	file_graph_network_proto_rawDesc = nil
	file_graph_network_proto_goTypes = nil
	file_graph_network_proto_depIdxs = nil
}
