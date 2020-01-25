// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CIM.Voip.proto

package cim

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 音视频呼叫邀请
type CIMVoipInviteReq struct {
	// cmd id:		0x0401
	CreatorUserId        uint64            `protobuf:"varint,1,opt,name=creator_user_id,json=creatorUserId,proto3" json:"creator_user_id,omitempty"`
	InviteUserList       []uint64          `protobuf:"varint,2,rep,packed,name=invite_user_list,json=inviteUserList,proto3" json:"invite_user_list,omitempty"`
	InviteType           CIMVoipInviteType `protobuf:"varint,3,opt,name=invite_type,json=inviteType,proto3,enum=CIM.Def.CIMVoipInviteType" json:"invite_type,omitempty"`
	ChannelInfo          *CIMChannelInfo   `protobuf:"bytes,4,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CIMVoipInviteReq) Reset()         { *m = CIMVoipInviteReq{} }
func (m *CIMVoipInviteReq) String() string { return proto.CompactTextString(m) }
func (*CIMVoipInviteReq) ProtoMessage()    {}
func (*CIMVoipInviteReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{0}
}

func (m *CIMVoipInviteReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipInviteReq.Unmarshal(m, b)
}
func (m *CIMVoipInviteReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipInviteReq.Marshal(b, m, deterministic)
}
func (m *CIMVoipInviteReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipInviteReq.Merge(m, src)
}
func (m *CIMVoipInviteReq) XXX_Size() int {
	return xxx_messageInfo_CIMVoipInviteReq.Size(m)
}
func (m *CIMVoipInviteReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipInviteReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipInviteReq proto.InternalMessageInfo

func (m *CIMVoipInviteReq) GetCreatorUserId() uint64 {
	if m != nil {
		return m.CreatorUserId
	}
	return 0
}

func (m *CIMVoipInviteReq) GetInviteUserList() []uint64 {
	if m != nil {
		return m.InviteUserList
	}
	return nil
}

func (m *CIMVoipInviteReq) GetInviteType() CIMVoipInviteType {
	if m != nil {
		return m.InviteType
	}
	return CIMVoipInviteType_kCIM_VOIP_INVITE_TYPE_UNKNOWN
}

func (m *CIMVoipInviteReq) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 音视频呼叫应答状态
type CIMVoipInviteReply struct {
	// cmd id:		0x0402
	UserId               uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	RspCode              CIMInviteRspCode `protobuf:"varint,2,opt,name=rsp_code,json=rspCode,proto3,enum=CIM.Def.CIMInviteRspCode" json:"rsp_code,omitempty"`
	ChannelInfo          *CIMChannelInfo  `protobuf:"bytes,3,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CIMVoipInviteReply) Reset()         { *m = CIMVoipInviteReply{} }
func (m *CIMVoipInviteReply) String() string { return proto.CompactTextString(m) }
func (*CIMVoipInviteReply) ProtoMessage()    {}
func (*CIMVoipInviteReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{1}
}

func (m *CIMVoipInviteReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipInviteReply.Unmarshal(m, b)
}
func (m *CIMVoipInviteReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipInviteReply.Marshal(b, m, deterministic)
}
func (m *CIMVoipInviteReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipInviteReply.Merge(m, src)
}
func (m *CIMVoipInviteReply) XXX_Size() int {
	return xxx_messageInfo_CIMVoipInviteReply.Size(m)
}
func (m *CIMVoipInviteReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipInviteReply.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipInviteReply proto.InternalMessageInfo

func (m *CIMVoipInviteReply) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMVoipInviteReply) GetRspCode() CIMInviteRspCode {
	if m != nil {
		return m.RspCode
	}
	return CIMInviteRspCode_kCIM_VOIP_INVITE_CODE_UNKNOWN
}

func (m *CIMVoipInviteReply) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 音视频呼叫ACK
// 100 Trying->180 Ringing->200 Ok->ACK(this message)
type CIMVoipInviteReplyAck struct {
	ChannelInfo          *CIMChannelInfo `protobuf:"bytes,1,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CIMVoipInviteReplyAck) Reset()         { *m = CIMVoipInviteReplyAck{} }
func (m *CIMVoipInviteReplyAck) String() string { return proto.CompactTextString(m) }
func (*CIMVoipInviteReplyAck) ProtoMessage()    {}
func (*CIMVoipInviteReplyAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{2}
}

func (m *CIMVoipInviteReplyAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipInviteReplyAck.Unmarshal(m, b)
}
func (m *CIMVoipInviteReplyAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipInviteReplyAck.Marshal(b, m, deterministic)
}
func (m *CIMVoipInviteReplyAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipInviteReplyAck.Merge(m, src)
}
func (m *CIMVoipInviteReplyAck) XXX_Size() int {
	return xxx_messageInfo_CIMVoipInviteReplyAck.Size(m)
}
func (m *CIMVoipInviteReplyAck) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipInviteReplyAck.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipInviteReplyAck proto.InternalMessageInfo

func (m *CIMVoipInviteReplyAck) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 心跳
type CIMVoipHeartbeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIMVoipHeartbeat) Reset()         { *m = CIMVoipHeartbeat{} }
func (m *CIMVoipHeartbeat) String() string { return proto.CompactTextString(m) }
func (*CIMVoipHeartbeat) ProtoMessage()    {}
func (*CIMVoipHeartbeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{3}
}

func (m *CIMVoipHeartbeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipHeartbeat.Unmarshal(m, b)
}
func (m *CIMVoipHeartbeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipHeartbeat.Marshal(b, m, deterministic)
}
func (m *CIMVoipHeartbeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipHeartbeat.Merge(m, src)
}
func (m *CIMVoipHeartbeat) XXX_Size() int {
	return xxx_messageInfo_CIMVoipHeartbeat.Size(m)
}
func (m *CIMVoipHeartbeat) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipHeartbeat.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipHeartbeat proto.InternalMessageInfo

// 挂断请求
type CIMVoipByeReq struct {
	LocalCallTimeLen     uint64          `protobuf:"varint,1,opt,name=local_call_time_len,json=localCallTimeLen,proto3" json:"local_call_time_len,omitempty"`
	UserId               uint64          `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ChannelInfo          *CIMChannelInfo `protobuf:"bytes,3,opt,name=channel_info,json=channelInfo,proto3" json:"channel_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CIMVoipByeReq) Reset()         { *m = CIMVoipByeReq{} }
func (m *CIMVoipByeReq) String() string { return proto.CompactTextString(m) }
func (*CIMVoipByeReq) ProtoMessage()    {}
func (*CIMVoipByeReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{4}
}

func (m *CIMVoipByeReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipByeReq.Unmarshal(m, b)
}
func (m *CIMVoipByeReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipByeReq.Marshal(b, m, deterministic)
}
func (m *CIMVoipByeReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipByeReq.Merge(m, src)
}
func (m *CIMVoipByeReq) XXX_Size() int {
	return xxx_messageInfo_CIMVoipByeReq.Size(m)
}
func (m *CIMVoipByeReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipByeReq.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipByeReq proto.InternalMessageInfo

func (m *CIMVoipByeReq) GetLocalCallTimeLen() uint64 {
	if m != nil {
		return m.LocalCallTimeLen
	}
	return 0
}

func (m *CIMVoipByeReq) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMVoipByeReq) GetChannelInfo() *CIMChannelInfo {
	if m != nil {
		return m.ChannelInfo
	}
	return nil
}

// 挂断响应
type CIMVoipByeRsp struct {
	ErrorCode            CIMErrorCode `protobuf:"varint,1,opt,name=error_code,json=errorCode,proto3,enum=CIM.Def.CIMErrorCode" json:"error_code,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *CIMVoipByeRsp) Reset()         { *m = CIMVoipByeRsp{} }
func (m *CIMVoipByeRsp) String() string { return proto.CompactTextString(m) }
func (*CIMVoipByeRsp) ProtoMessage()    {}
func (*CIMVoipByeRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{5}
}

func (m *CIMVoipByeRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipByeRsp.Unmarshal(m, b)
}
func (m *CIMVoipByeRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipByeRsp.Marshal(b, m, deterministic)
}
func (m *CIMVoipByeRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipByeRsp.Merge(m, src)
}
func (m *CIMVoipByeRsp) XXX_Size() int {
	return xxx_messageInfo_CIMVoipByeRsp.Size(m)
}
func (m *CIMVoipByeRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipByeRsp.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipByeRsp proto.InternalMessageInfo

func (m *CIMVoipByeRsp) GetErrorCode() CIMErrorCode {
	if m != nil {
		return m.ErrorCode
	}
	return CIMErrorCode_kCIM_ERR_SUCCSSE
}

// 挂断通知
type CIMVoipByeNotify struct {
	UserId               uint64           `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	ByeReason            CIMVoipByeReason `protobuf:"varint,2,opt,name=byeReason,proto3,enum=CIM.Def.CIMVoipByeReason" json:"byeReason,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CIMVoipByeNotify) Reset()         { *m = CIMVoipByeNotify{} }
func (m *CIMVoipByeNotify) String() string { return proto.CompactTextString(m) }
func (*CIMVoipByeNotify) ProtoMessage()    {}
func (*CIMVoipByeNotify) Descriptor() ([]byte, []int) {
	return fileDescriptor_d1f178c14f4443d3, []int{6}
}

func (m *CIMVoipByeNotify) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CIMVoipByeNotify.Unmarshal(m, b)
}
func (m *CIMVoipByeNotify) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CIMVoipByeNotify.Marshal(b, m, deterministic)
}
func (m *CIMVoipByeNotify) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIMVoipByeNotify.Merge(m, src)
}
func (m *CIMVoipByeNotify) XXX_Size() int {
	return xxx_messageInfo_CIMVoipByeNotify.Size(m)
}
func (m *CIMVoipByeNotify) XXX_DiscardUnknown() {
	xxx_messageInfo_CIMVoipByeNotify.DiscardUnknown(m)
}

var xxx_messageInfo_CIMVoipByeNotify proto.InternalMessageInfo

func (m *CIMVoipByeNotify) GetUserId() uint64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *CIMVoipByeNotify) GetByeReason() CIMVoipByeReason {
	if m != nil {
		return m.ByeReason
	}
	return CIMVoipByeReason_kCIM_VOIP_BYE_REASON_UNKNOWN
}

func init() {
	proto.RegisterType((*CIMVoipInviteReq)(nil), "CIM.Voip.CIMVoipInviteReq")
	proto.RegisterType((*CIMVoipInviteReply)(nil), "CIM.Voip.CIMVoipInviteReply")
	proto.RegisterType((*CIMVoipInviteReplyAck)(nil), "CIM.Voip.CIMVoipInviteReplyAck")
	proto.RegisterType((*CIMVoipHeartbeat)(nil), "CIM.Voip.CIMVoipHeartbeat")
	proto.RegisterType((*CIMVoipByeReq)(nil), "CIM.Voip.CIMVoipByeReq")
	proto.RegisterType((*CIMVoipByeRsp)(nil), "CIM.Voip.CIMVoipByeRsp")
	proto.RegisterType((*CIMVoipByeNotify)(nil), "CIM.Voip.CIMVoipByeNotify")
}

func init() { proto.RegisterFile("CIM.Voip.proto", fileDescriptor_d1f178c14f4443d3) }

var fileDescriptor_d1f178c14f4443d3 = []byte{
	// 427 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xd6, 0xc6, 0x51, 0x7f, 0x26, 0x24, 0x44, 0x8b, 0xaa, 0x98, 0x9e, 0xac, 0x1c, 0x90, 0x2f,
	0xe4, 0x50, 0x2a, 0x21, 0xc1, 0x89, 0x98, 0x4a, 0xb5, 0xd4, 0x72, 0x30, 0x85, 0x03, 0x17, 0x6b,
	0xb3, 0x1e, 0xab, 0x2b, 0x36, 0xde, 0x65, 0x77, 0x83, 0xe4, 0xe7, 0xe0, 0x01, 0x78, 0x32, 0xde,
	0x05, 0xd9, 0xeb, 0xa4, 0x31, 0x41, 0x48, 0x15, 0xb7, 0xd9, 0x99, 0x6f, 0xc6, 0xdf, 0x8f, 0x0c,
	0x93, 0x24, 0xbd, 0x5d, 0x7c, 0x56, 0x42, 0x2f, 0xb4, 0x51, 0x4e, 0xd1, 0x93, 0xed, 0xfb, 0x7c,
	0xdc, 0x54, 0xef, 0xb1, 0xf4, 0x83, 0xf9, 0x2f, 0x02, 0xd3, 0x24, 0xbd, 0x6d, 0x46, 0x69, 0xf5,
	0x5d, 0x38, 0xcc, 0xf0, 0x1b, 0x7d, 0x01, 0x4f, 0xb9, 0x41, 0xe6, 0x94, 0xc9, 0x37, 0x16, 0x4d,
	0x2e, 0x8a, 0x90, 0x44, 0x24, 0x1e, 0x66, 0xe3, 0xae, 0xfd, 0xc9, 0xa2, 0x49, 0x0b, 0x1a, 0xc3,
	0x54, 0xb4, 0x4b, 0x1e, 0x26, 0x85, 0x75, 0xe1, 0x20, 0x0a, 0xe2, 0x61, 0x36, 0xf1, 0xfd, 0x06,
	0x77, 0x23, 0xac, 0xa3, 0x6f, 0x61, 0xd4, 0x21, 0x5d, 0xad, 0x31, 0x0c, 0x22, 0x12, 0x4f, 0x2e,
	0xce, 0x17, 0x5b, 0x2e, 0x3d, 0x06, 0x77, 0xb5, 0xc6, 0x0c, 0xc4, 0xae, 0xa6, 0x6f, 0xe0, 0x09,
	0xbf, 0x67, 0x55, 0x85, 0x32, 0x17, 0x55, 0xa9, 0xc2, 0x61, 0x44, 0xe2, 0xd1, 0xc5, 0x6c, 0x7f,
	0x3b, 0xf1, 0xf3, 0xb4, 0x2a, 0x55, 0x36, 0xe2, 0x0f, 0x8f, 0xf9, 0x4f, 0x02, 0xf4, 0x0f, 0x7d,
	0x5a, 0xd6, 0x74, 0x06, 0xc7, 0x7d, 0x65, 0x47, 0x1b, 0x2f, 0xe9, 0x12, 0x4e, 0x8c, 0xd5, 0x39,
	0x57, 0x05, 0x86, 0x83, 0x96, 0xe5, 0xf3, 0xfd, 0xef, 0x74, 0x37, 0xac, 0x4e, 0x54, 0x81, 0xd9,
	0xb1, 0xf1, 0xc5, 0x01, 0xc3, 0xe0, 0x11, 0x0c, 0x3f, 0xc2, 0xd9, 0x21, 0xc1, 0x77, 0xfc, 0xeb,
	0xc1, 0x51, 0xf2, 0x88, 0xa3, 0x74, 0x97, 0xea, 0x35, 0x32, 0xe3, 0x56, 0xc8, 0xdc, 0xfc, 0x07,
	0x81, 0x71, 0xd7, 0x5c, 0xd6, 0x6d, 0xce, 0x2f, 0xe1, 0x99, 0x54, 0x9c, 0xc9, 0x9c, 0x33, 0x29,
	0x73, 0x27, 0xd6, 0x98, 0x4b, 0xac, 0x3a, 0x47, 0xa6, 0xed, 0x28, 0x61, 0x52, 0xde, 0x89, 0x35,
	0xde, 0x60, 0xb5, 0x6f, 0xda, 0xa0, 0x67, 0xda, 0xff, 0xc8, 0xbf, 0xea, 0x91, 0xb2, 0x9a, 0x5e,
	0x02, 0xa0, 0x31, 0xca, 0xf8, 0x0c, 0x48, 0x9b, 0xc1, 0xd9, 0xfe, 0xa9, 0xab, 0x66, 0xda, 0xfa,
	0x7f, 0x8a, 0xdb, 0x72, 0x5e, 0xec, 0x04, 0x2f, 0x6b, 0xfc, 0xa0, 0x9c, 0x28, 0xff, 0x11, 0xf2,
	0x6b, 0x38, 0x5d, 0x35, 0x0e, 0x30, 0xab, 0xaa, 0xbf, 0xa5, 0xbc, 0xb3, 0xa8, 0x01, 0x64, 0x0f,
	0xd8, 0x65, 0x04, 0x33, 0xae, 0xd6, 0x0b, 0xae, 0xca, 0x12, 0x91, 0xdf, 0x33, 0xe7, 0xff, 0xa2,
	0xd5, 0xa6, 0xbc, 0x0e, 0xbe, 0x04, 0x5c, 0xac, 0x57, 0x47, 0x6d, 0xe3, 0xd5, 0xef, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xe9, 0xde, 0x88, 0x94, 0x81, 0x03, 0x00, 0x00,
}