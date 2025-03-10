// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.4
// 	protoc        v5.29.1
// source: api/im/im.proto

package im

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

type Message struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	SpaceId       int64                  `protobuf:"varint,2,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	From          int64                  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To            int64                  `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Type          string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Content       string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_api_im_im_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Message) GetSpaceId() int64 {
	if x != nil {
		return x.SpaceId
	}
	return 0
}

func (x *Message) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *Message) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *Message) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Message) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpaceId       int64                  `protobuf:"varint,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	ChannelId     int64                  `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	From          int64                  `protobuf:"varint,3,opt,name=from,proto3" json:"from,omitempty"`
	To            int64                  `protobuf:"varint,4,opt,name=to,proto3" json:"to,omitempty"`
	Type          string                 `protobuf:"bytes,5,opt,name=type,proto3" json:"type,omitempty"`
	Content       string                 `protobuf:"bytes,6,opt,name=content,proto3" json:"content,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_api_im_im_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{1}
}

func (x *SendMessageRequest) GetSpaceId() int64 {
	if x != nil {
		return x.SpaceId
	}
	return 0
}

func (x *SendMessageRequest) GetChannelId() int64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *SendMessageRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *SendMessageRequest) GetTo() int64 {
	if x != nil {
		return x.To
	}
	return 0
}

func (x *SendMessageRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *SendMessageRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SendMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MessageId     int64                  `protobuf:"varint,1,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageResponse) Reset() {
	*x = SendMessageResponse{}
	mi := &file_api_im_im_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageResponse) ProtoMessage() {}

func (x *SendMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageResponse.ProtoReflect.Descriptor instead.
func (*SendMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{2}
}

func (x *SendMessageResponse) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type AckMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpaceId       int64                  `protobuf:"varint,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MessageIds    []int64                `protobuf:"varint,3,rep,packed,name=message_ids,json=messageIds,proto3" json:"message_ids,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckMessagesRequest) Reset() {
	*x = AckMessagesRequest{}
	mi := &file_api_im_im_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessagesRequest) ProtoMessage() {}

func (x *AckMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessagesRequest.ProtoReflect.Descriptor instead.
func (*AckMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{3}
}

func (x *AckMessagesRequest) GetSpaceId() int64 {
	if x != nil {
		return x.SpaceId
	}
	return 0
}

func (x *AckMessagesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AckMessagesRequest) GetMessageIds() []int64 {
	if x != nil {
		return x.MessageIds
	}
	return nil
}

type AckMessagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckMessagesResponse) Reset() {
	*x = AckMessagesResponse{}
	mi := &file_api_im_im_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckMessagesResponse) ProtoMessage() {}

func (x *AckMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckMessagesResponse.ProtoReflect.Descriptor instead.
func (*AckMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{4}
}

func (x *AckMessagesResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type PullHistoryRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpaceId       int64                  `protobuf:"varint,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	ChannelId     int64                  `protobuf:"varint,2,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId        int64                  `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	From          int64                  `protobuf:"varint,4,opt,name=from,proto3" json:"from,omitempty"`
	Cursor        int64                  `protobuf:"varint,5,opt,name=cursor,proto3" json:"cursor,omitempty"`
	Limit         int32                  `protobuf:"varint,6,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PullHistoryRequest) Reset() {
	*x = PullHistoryRequest{}
	mi := &file_api_im_im_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullHistoryRequest) ProtoMessage() {}

func (x *PullHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullHistoryRequest.ProtoReflect.Descriptor instead.
func (*PullHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{5}
}

func (x *PullHistoryRequest) GetSpaceId() int64 {
	if x != nil {
		return x.SpaceId
	}
	return 0
}

func (x *PullHistoryRequest) GetChannelId() int64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *PullHistoryRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PullHistoryRequest) GetFrom() int64 {
	if x != nil {
		return x.From
	}
	return 0
}

func (x *PullHistoryRequest) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

func (x *PullHistoryRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type PullHistoryResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	Cursor        int64                  `protobuf:"varint,2,opt,name=cursor,proto3" json:"cursor,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PullHistoryResponse) Reset() {
	*x = PullHistoryResponse{}
	mi := &file_api_im_im_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PullHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PullHistoryResponse) ProtoMessage() {}

func (x *PullHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PullHistoryResponse.ProtoReflect.Descriptor instead.
func (*PullHistoryResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{6}
}

func (x *PullHistoryResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

func (x *PullHistoryResponse) GetCursor() int64 {
	if x != nil {
		return x.Cursor
	}
	return 0
}

type GetInboxMessagesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpaceId       int64                  `protobuf:"varint,1,opt,name=space_id,json=spaceId,proto3" json:"space_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit         int32                  `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInboxMessagesRequest) Reset() {
	*x = GetInboxMessagesRequest{}
	mi := &file_api_im_im_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInboxMessagesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxMessagesRequest) ProtoMessage() {}

func (x *GetInboxMessagesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxMessagesRequest.ProtoReflect.Descriptor instead.
func (*GetInboxMessagesRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{7}
}

func (x *GetInboxMessagesRequest) GetSpaceId() int64 {
	if x != nil {
		return x.SpaceId
	}
	return 0
}

func (x *GetInboxMessagesRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GetInboxMessagesRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetInboxMessagesResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Messages      []*Message             `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetInboxMessagesResponse) Reset() {
	*x = GetInboxMessagesResponse{}
	mi := &file_api_im_im_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetInboxMessagesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInboxMessagesResponse) ProtoMessage() {}

func (x *GetInboxMessagesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInboxMessagesResponse.ProtoReflect.Descriptor instead.
func (*GetInboxMessagesResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{8}
}

func (x *GetInboxMessagesResponse) GetMessages() []*Message {
	if x != nil {
		return x.Messages
	}
	return nil
}

type AckChannelMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChannelId     int64                  `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	MessageId     int64                  `protobuf:"varint,3,opt,name=message_id,json=messageId,proto3" json:"message_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckChannelMessageRequest) Reset() {
	*x = AckChannelMessageRequest{}
	mi := &file_api_im_im_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckChannelMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckChannelMessageRequest) ProtoMessage() {}

func (x *AckChannelMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckChannelMessageRequest.ProtoReflect.Descriptor instead.
func (*AckChannelMessageRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{9}
}

func (x *AckChannelMessageRequest) GetChannelId() int64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *AckChannelMessageRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AckChannelMessageRequest) GetMessageId() int64 {
	if x != nil {
		return x.MessageId
	}
	return 0
}

type AckChannelMessageResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AckChannelMessageResponse) Reset() {
	*x = AckChannelMessageResponse{}
	mi := &file_api_im_im_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AckChannelMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckChannelMessageResponse) ProtoMessage() {}

func (x *AckChannelMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckChannelMessageResponse.ProtoReflect.Descriptor instead.
func (*AckChannelMessageResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{10}
}

func (x *AckChannelMessageResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type GetChannelInboxRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChannelId     int64                  `protobuf:"varint,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	UserId        int64                  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChannelInboxRequest) Reset() {
	*x = GetChannelInboxRequest{}
	mi := &file_api_im_im_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChannelInboxRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelInboxRequest) ProtoMessage() {}

func (x *GetChannelInboxRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelInboxRequest.ProtoReflect.Descriptor instead.
func (*GetChannelInboxRequest) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{11}
}

func (x *GetChannelInboxRequest) GetChannelId() int64 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *GetChannelInboxRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetChannelInboxResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Current       int64                  `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	Last          int64                  `protobuf:"varint,2,opt,name=last,proto3" json:"last,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChannelInboxResponse) Reset() {
	*x = GetChannelInboxResponse{}
	mi := &file_api_im_im_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChannelInboxResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChannelInboxResponse) ProtoMessage() {}

func (x *GetChannelInboxResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_im_im_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChannelInboxResponse.ProtoReflect.Descriptor instead.
func (*GetChannelInboxResponse) Descriptor() ([]byte, []int) {
	return file_api_im_im_proto_rawDescGZIP(), []int{12}
}

func (x *GetChannelInboxResponse) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *GetChannelInboxResponse) GetLast() int64 {
	if x != nil {
		return x.Last
	}
	return 0
}

var File_api_im_im_proto protoreflect.FileDescriptor

var file_api_im_im_proto_rawDesc = string([]byte{
	0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x2f, 0x69, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x69, 0x6d, 0x22, 0xa5, 0x01, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x74, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xa0, 0x01,
	0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x74, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x34, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x69, 0x0a, 0x12, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1f, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x73, 0x18,
	0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x49, 0x64,
	0x73, 0x22, 0x2f, 0x0a, 0x13, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x22, 0xa9, 0x01, 0x0a, 0x12, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x66, 0x72, 0x6f, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d,
	0x12, 0x16, 0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x56,
	0x0a, 0x13, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x2e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x63, 0x75, 0x72, 0x73, 0x6f, 0x72, 0x22, 0x63, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x43, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x69, 0x6d, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73,
	0x22, 0x71, 0x0a, 0x18, 0x41, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x49, 0x64, 0x22, 0x35, 0x0a, 0x19, 0x41, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x68, 0x61, 0x6e, 0x6e, 0x65,
	0x6c, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x17,
	0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x61, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x04, 0x6c, 0x61, 0x73, 0x74, 0x32, 0xb8, 0x03, 0x0a, 0x09, 0x49, 0x6d, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x73, 0x12, 0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x2e,
	0x41, 0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x12, 0x1b, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x49, 0x6e, 0x62, 0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x62,
	0x6f, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x11, 0x41, 0x63, 0x6b, 0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x2e, 0x69, 0x6d, 0x2e, 0x41, 0x63, 0x6b,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x69, 0x6d, 0x2e, 0x41, 0x63, 0x6b, 0x43, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e, 0x6e,
	0x65, 0x6c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x12, 0x1a, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x69, 0x6d, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x6e,
	0x6e, 0x65, 0x6c, 0x49, 0x6e, 0x62, 0x6f, 0x78, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3e, 0x0a, 0x0b, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12,
	0x16, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x75, 0x6c, 0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x69, 0x6d, 0x2e, 0x50, 0x75, 0x6c,
	0x6c, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x08, 0x5a, 0x06, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
})

var (
	file_api_im_im_proto_rawDescOnce sync.Once
	file_api_im_im_proto_rawDescData []byte
)

func file_api_im_im_proto_rawDescGZIP() []byte {
	file_api_im_im_proto_rawDescOnce.Do(func() {
		file_api_im_im_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_im_im_proto_rawDesc), len(file_api_im_im_proto_rawDesc)))
	})
	return file_api_im_im_proto_rawDescData
}

var file_api_im_im_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_im_im_proto_goTypes = []any{
	(*Message)(nil),                   // 0: im.Message
	(*SendMessageRequest)(nil),        // 1: im.SendMessageRequest
	(*SendMessageResponse)(nil),       // 2: im.SendMessageResponse
	(*AckMessagesRequest)(nil),        // 3: im.AckMessagesRequest
	(*AckMessagesResponse)(nil),       // 4: im.AckMessagesResponse
	(*PullHistoryRequest)(nil),        // 5: im.PullHistoryRequest
	(*PullHistoryResponse)(nil),       // 6: im.PullHistoryResponse
	(*GetInboxMessagesRequest)(nil),   // 7: im.GetInboxMessagesRequest
	(*GetInboxMessagesResponse)(nil),  // 8: im.GetInboxMessagesResponse
	(*AckChannelMessageRequest)(nil),  // 9: im.AckChannelMessageRequest
	(*AckChannelMessageResponse)(nil), // 10: im.AckChannelMessageResponse
	(*GetChannelInboxRequest)(nil),    // 11: im.GetChannelInboxRequest
	(*GetChannelInboxResponse)(nil),   // 12: im.GetChannelInboxResponse
}
var file_api_im_im_proto_depIdxs = []int32{
	0,  // 0: im.PullHistoryResponse.messages:type_name -> im.Message
	0,  // 1: im.GetInboxMessagesResponse.messages:type_name -> im.Message
	3,  // 2: im.ImService.AckMessages:input_type -> im.AckMessagesRequest
	7,  // 3: im.ImService.GetInboxMessages:input_type -> im.GetInboxMessagesRequest
	9,  // 4: im.ImService.AckChannelMessage:input_type -> im.AckChannelMessageRequest
	11, // 5: im.ImService.GetChannelInbox:input_type -> im.GetChannelInboxRequest
	1,  // 6: im.ImService.SendMessage:input_type -> im.SendMessageRequest
	5,  // 7: im.ImService.PullHistory:input_type -> im.PullHistoryRequest
	4,  // 8: im.ImService.AckMessages:output_type -> im.AckMessagesResponse
	8,  // 9: im.ImService.GetInboxMessages:output_type -> im.GetInboxMessagesResponse
	10, // 10: im.ImService.AckChannelMessage:output_type -> im.AckChannelMessageResponse
	12, // 11: im.ImService.GetChannelInbox:output_type -> im.GetChannelInboxResponse
	2,  // 12: im.ImService.SendMessage:output_type -> im.SendMessageResponse
	6,  // 13: im.ImService.PullHistory:output_type -> im.PullHistoryResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_im_im_proto_init() }
func file_api_im_im_proto_init() {
	if File_api_im_im_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_im_im_proto_rawDesc), len(file_api_im_im_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_im_im_proto_goTypes,
		DependencyIndexes: file_api_im_im_proto_depIdxs,
		MessageInfos:      file_api_im_im_proto_msgTypes,
	}.Build()
	File_api_im_im_proto = out.File
	file_api_im_im_proto_goTypes = nil
	file_api_im_im_proto_depIdxs = nil
}
