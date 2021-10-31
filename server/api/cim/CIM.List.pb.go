// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: CIM.List.proto

package cim

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

// 最近聊天会话列表请求
type CIMRecentContactSessionReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0201
	UserId           uint64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LatestUpdateTime uint32 `protobuf:"varint,2,opt,name=latest_update_time,json=latestUpdateTime,proto3" json:"latest_update_time,omitempty"` // 最后更新时间
}

func (x *CIMRecentContactSessionReq) Reset() {
	*x = CIMRecentContactSessionReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_List_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMRecentContactSessionReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMRecentContactSessionReq) ProtoMessage() {}

func (x *CIMRecentContactSessionReq) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_List_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMRecentContactSessionReq.ProtoReflect.Descriptor instead.
func (*CIMRecentContactSessionReq) Descriptor() ([]byte, []int) {
	return file_CIM_List_proto_rawDescGZIP(), []int{0}
}

func (x *CIMRecentContactSessionReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMRecentContactSessionReq) GetLatestUpdateTime() uint32 {
	if x != nil {
		return x.LatestUpdateTime
	}
	return 0
}

type CIMRecentContactSessionRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0202
	UserId             uint64                   `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	UnreadCounts       uint32                   `protobuf:"varint,2,opt,name=unread_counts,json=unreadCounts,proto3" json:"unread_counts,omitempty"`                    // 总未读数量
	ContactSessionList []*CIMContactSessionInfo `protobuf:"bytes,3,rep,name=contact_session_list,json=contactSessionList,proto3" json:"contact_session_list,omitempty"` // 会话列表
}

func (x *CIMRecentContactSessionRsp) Reset() {
	*x = CIMRecentContactSessionRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_List_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMRecentContactSessionRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMRecentContactSessionRsp) ProtoMessage() {}

func (x *CIMRecentContactSessionRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_List_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMRecentContactSessionRsp.ProtoReflect.Descriptor instead.
func (*CIMRecentContactSessionRsp) Descriptor() ([]byte, []int) {
	return file_CIM_List_proto_rawDescGZIP(), []int{1}
}

func (x *CIMRecentContactSessionRsp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMRecentContactSessionRsp) GetUnreadCounts() uint32 {
	if x != nil {
		return x.UnreadCounts
	}
	return 0
}

func (x *CIMRecentContactSessionRsp) GetContactSessionList() []*CIMContactSessionInfo {
	if x != nil {
		return x.ContactSessionList
	}
	return nil
}

// 历史离线聊天消息请求
type CIMGetMsgListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0205
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	//   uint64 from_time_stamp = 4; // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;  // 结束时间点，单位：毫秒
	EndMsgId   uint64 `protobuf:"varint,4,opt,name=end_msg_id,json=endMsgId,proto3" json:"end_msg_id,omitempty"`     // 结束服务器消息id(不包含在查询结果中)
	LimitCount uint32 `protobuf:"varint,6,opt,name=limit_count,json=limitCount,proto3" json:"limit_count,omitempty"` // 本次查询消息的条数上线(最多100条)
}

func (x *CIMGetMsgListReq) Reset() {
	*x = CIMGetMsgListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_List_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMGetMsgListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMGetMsgListReq) ProtoMessage() {}

func (x *CIMGetMsgListReq) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_List_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMGetMsgListReq.ProtoReflect.Descriptor instead.
func (*CIMGetMsgListReq) Descriptor() ([]byte, []int) {
	return file_CIM_List_proto_rawDescGZIP(), []int{2}
}

func (x *CIMGetMsgListReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMGetMsgListReq) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetEndMsgId() uint64 {
	if x != nil {
		return x.EndMsgId
	}
	return 0
}

func (x *CIMGetMsgListReq) GetLimitCount() uint32 {
	if x != nil {
		return x.LimitCount
	}
	return 0
}

//对于群而言，如果消息数目返回的数值小于请求的cnt,则表示群的消息能拉取的到头了，更早的消息没有权限拉取。
//如果limit_count 和 msg_list.count 不一致，说明服务器消息有缺失，需要
//客户端做一个缺失标记，避免下次再次拉取。
type CIMGetMsgListRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// cmd id:		0x0206
	UserId      uint64         `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SessionType CIMSessionType `protobuf:"varint,2,opt,name=session_type,json=sessionType,proto3,enum=CIM.Def.CIMSessionType" json:"session_type,omitempty"`
	SessionId   uint64         `protobuf:"varint,3,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	EndMsgId    uint64         `protobuf:"varint,4,opt,name=end_msg_id,json=endMsgId,proto3" json:"end_msg_id,omitempty"` // 结束服务器消息id(不包含在查询结果中)
	//   uint64 from_time_stamp = 4;     // 起始时间点，单位：毫秒
	//   uint64 end_time_stamp = 5;      // 结束时间点，单位：毫秒
	MsgList []*CIMMsgInfo `protobuf:"bytes,6,rep,name=msg_list,json=msgList,proto3" json:"msg_list,omitempty"` // 消息列表
}

func (x *CIMGetMsgListRsp) Reset() {
	*x = CIMGetMsgListRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_CIM_List_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CIMGetMsgListRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CIMGetMsgListRsp) ProtoMessage() {}

func (x *CIMGetMsgListRsp) ProtoReflect() protoreflect.Message {
	mi := &file_CIM_List_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CIMGetMsgListRsp.ProtoReflect.Descriptor instead.
func (*CIMGetMsgListRsp) Descriptor() ([]byte, []int) {
	return file_CIM_List_proto_rawDescGZIP(), []int{3}
}

func (x *CIMGetMsgListRsp) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetSessionType() CIMSessionType {
	if x != nil {
		return x.SessionType
	}
	return CIMSessionType_kCIM_SESSION_TYPE_Invalid
}

func (x *CIMGetMsgListRsp) GetSessionId() uint64 {
	if x != nil {
		return x.SessionId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetEndMsgId() uint64 {
	if x != nil {
		return x.EndMsgId
	}
	return 0
}

func (x *CIMGetMsgListRsp) GetMsgList() []*CIMMsgInfo {
	if x != nil {
		return x.MsgList
	}
	return nil
}

var File_CIM_List_proto protoreflect.FileDescriptor

var file_CIM_List_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x43, 0x49, 0x4d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x43, 0x49, 0x4d, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x43, 0x49, 0x4d, 0x2e,
	0x44, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x63, 0x0a, 0x1a, 0x43, 0x49, 0x4d,
	0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x2c, 0x0a, 0x12, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x10, 0x6c, 0x61,
	0x74, 0x65, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0xac,
	0x01, 0x0a, 0x1a, 0x43, 0x49, 0x4d, 0x52, 0x65, 0x63, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64,
	0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0c, 0x75,
	0x6e, 0x72, 0x65, 0x61, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12, 0x50, 0x0a, 0x14, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x5f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x43, 0x49, 0x4d, 0x2e,
	0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x12, 0x63, 0x6f, 0x6e, 0x74, 0x61,
	0x63, 0x74, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xc5, 0x01,
	0x0a, 0x10, 0x43, 0x49, 0x4d, 0x47, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x17, 0x2e, 0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x53,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x5f, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0a, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd4, 0x01, 0x0a, 0x10, 0x43, 0x49, 0x4d, 0x47, 0x65, 0x74,
	0x4d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x73, 0x70, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x43, 0x49, 0x4d, 0x2e,
	0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x0a, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x08, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08,
	0x6d, 0x73, 0x67, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x43, 0x49, 0x4d, 0x2e, 0x44, 0x65, 0x66, 0x2e, 0x43, 0x49, 0x4d, 0x4d, 0x73, 0x67, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x22, 0x0a, 0x17,
	0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x66, 0x66, 0x65, 0x65, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x48, 0x03, 0x5a, 0x05, 0x2e, 0x3b, 0x63, 0x69, 0x6d,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_CIM_List_proto_rawDescOnce sync.Once
	file_CIM_List_proto_rawDescData = file_CIM_List_proto_rawDesc
)

func file_CIM_List_proto_rawDescGZIP() []byte {
	file_CIM_List_proto_rawDescOnce.Do(func() {
		file_CIM_List_proto_rawDescData = protoimpl.X.CompressGZIP(file_CIM_List_proto_rawDescData)
	})
	return file_CIM_List_proto_rawDescData
}

var file_CIM_List_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_CIM_List_proto_goTypes = []interface{}{
	(*CIMRecentContactSessionReq)(nil), // 0: CIM.List.CIMRecentContactSessionReq
	(*CIMRecentContactSessionRsp)(nil), // 1: CIM.List.CIMRecentContactSessionRsp
	(*CIMGetMsgListReq)(nil),           // 2: CIM.List.CIMGetMsgListReq
	(*CIMGetMsgListRsp)(nil),           // 3: CIM.List.CIMGetMsgListRsp
	(*CIMContactSessionInfo)(nil),      // 4: CIM.Def.CIMContactSessionInfo
	(CIMSessionType)(0),                // 5: CIM.Def.CIMSessionType
	(*CIMMsgInfo)(nil),                 // 6: CIM.Def.CIMMsgInfo
}
var file_CIM_List_proto_depIdxs = []int32{
	4, // 0: CIM.List.CIMRecentContactSessionRsp.contact_session_list:type_name -> CIM.Def.CIMContactSessionInfo
	5, // 1: CIM.List.CIMGetMsgListReq.session_type:type_name -> CIM.Def.CIMSessionType
	5, // 2: CIM.List.CIMGetMsgListRsp.session_type:type_name -> CIM.Def.CIMSessionType
	6, // 3: CIM.List.CIMGetMsgListRsp.msg_list:type_name -> CIM.Def.CIMMsgInfo
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_CIM_List_proto_init() }
func file_CIM_List_proto_init() {
	if File_CIM_List_proto != nil {
		return
	}
	file_CIM_Def_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_CIM_List_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMRecentContactSessionReq); i {
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
		file_CIM_List_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMRecentContactSessionRsp); i {
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
		file_CIM_List_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMGetMsgListReq); i {
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
		file_CIM_List_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CIMGetMsgListRsp); i {
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
			RawDescriptor: file_CIM_List_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_CIM_List_proto_goTypes,
		DependencyIndexes: file_CIM_List_proto_depIdxs,
		MessageInfos:      file_CIM_List_proto_msgTypes,
	}.Build()
	File_CIM_List_proto = out.File
	file_CIM_List_proto_rawDesc = nil
	file_CIM_List_proto_goTypes = nil
	file_CIM_List_proto_depIdxs = nil
}