// Code generated by protoc-gen-go. DO NOT EDIT.
// source: atomix/lock/lock.proto

package lock

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import headers "github.com/atomix/atomix-k8s-controller/proto/atomix/headers"
import duration "github.com/golang/protobuf/ptypes/duration"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Timeout              *duration.Duration     `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{0}
}
func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (dst *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(dst, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CreateRequest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type CreateResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{1}
}
func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (dst *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(dst, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type KeepAliveRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *KeepAliveRequest) Reset()         { *m = KeepAliveRequest{} }
func (m *KeepAliveRequest) String() string { return proto.CompactTextString(m) }
func (*KeepAliveRequest) ProtoMessage()    {}
func (*KeepAliveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{2}
}
func (m *KeepAliveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAliveRequest.Unmarshal(m, b)
}
func (m *KeepAliveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAliveRequest.Marshal(b, m, deterministic)
}
func (dst *KeepAliveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAliveRequest.Merge(dst, src)
}
func (m *KeepAliveRequest) XXX_Size() int {
	return xxx_messageInfo_KeepAliveRequest.Size(m)
}
func (m *KeepAliveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAliveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAliveRequest proto.InternalMessageInfo

func (m *KeepAliveRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type KeepAliveResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *KeepAliveResponse) Reset()         { *m = KeepAliveResponse{} }
func (m *KeepAliveResponse) String() string { return proto.CompactTextString(m) }
func (*KeepAliveResponse) ProtoMessage()    {}
func (*KeepAliveResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{3}
}
func (m *KeepAliveResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeepAliveResponse.Unmarshal(m, b)
}
func (m *KeepAliveResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeepAliveResponse.Marshal(b, m, deterministic)
}
func (dst *KeepAliveResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeepAliveResponse.Merge(dst, src)
}
func (m *KeepAliveResponse) XXX_Size() int {
	return xxx_messageInfo_KeepAliveResponse.Size(m)
}
func (m *KeepAliveResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_KeepAliveResponse.DiscardUnknown(m)
}

var xxx_messageInfo_KeepAliveResponse proto.InternalMessageInfo

func (m *KeepAliveResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type CloseRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Delete               bool                   `protobuf:"varint,2,opt,name=delete,proto3" json:"delete,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CloseRequest) Reset()         { *m = CloseRequest{} }
func (m *CloseRequest) String() string { return proto.CompactTextString(m) }
func (*CloseRequest) ProtoMessage()    {}
func (*CloseRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{4}
}
func (m *CloseRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseRequest.Unmarshal(m, b)
}
func (m *CloseRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseRequest.Marshal(b, m, deterministic)
}
func (dst *CloseRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseRequest.Merge(dst, src)
}
func (m *CloseRequest) XXX_Size() int {
	return xxx_messageInfo_CloseRequest.Size(m)
}
func (m *CloseRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CloseRequest proto.InternalMessageInfo

func (m *CloseRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *CloseRequest) GetDelete() bool {
	if m != nil {
		return m.Delete
	}
	return false
}

type CloseResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CloseResponse) Reset()         { *m = CloseResponse{} }
func (m *CloseResponse) String() string { return proto.CompactTextString(m) }
func (*CloseResponse) ProtoMessage()    {}
func (*CloseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{5}
}
func (m *CloseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CloseResponse.Unmarshal(m, b)
}
func (m *CloseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CloseResponse.Marshal(b, m, deterministic)
}
func (dst *CloseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CloseResponse.Merge(dst, src)
}
func (m *CloseResponse) XXX_Size() int {
	return xxx_messageInfo_CloseResponse.Size(m)
}
func (m *CloseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CloseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CloseResponse proto.InternalMessageInfo

func (m *CloseResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type LockRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Timeout              *duration.Duration     `protobuf:"bytes,2,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *LockRequest) Reset()         { *m = LockRequest{} }
func (m *LockRequest) String() string { return proto.CompactTextString(m) }
func (*LockRequest) ProtoMessage()    {}
func (*LockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{6}
}
func (m *LockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockRequest.Unmarshal(m, b)
}
func (m *LockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockRequest.Marshal(b, m, deterministic)
}
func (dst *LockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockRequest.Merge(dst, src)
}
func (m *LockRequest) XXX_Size() int {
	return xxx_messageInfo_LockRequest.Size(m)
}
func (m *LockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockRequest proto.InternalMessageInfo

func (m *LockRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LockRequest) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

type LockResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Version              uint64                  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LockResponse) Reset()         { *m = LockResponse{} }
func (m *LockResponse) String() string { return proto.CompactTextString(m) }
func (*LockResponse) ProtoMessage()    {}
func (*LockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{7}
}
func (m *LockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockResponse.Unmarshal(m, b)
}
func (m *LockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockResponse.Marshal(b, m, deterministic)
}
func (dst *LockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockResponse.Merge(dst, src)
}
func (m *LockResponse) XXX_Size() int {
	return xxx_messageInfo_LockResponse.Size(m)
}
func (m *LockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LockResponse proto.InternalMessageInfo

func (m *LockResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *LockResponse) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type UnlockRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Version              uint64                 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *UnlockRequest) Reset()         { *m = UnlockRequest{} }
func (m *UnlockRequest) String() string { return proto.CompactTextString(m) }
func (*UnlockRequest) ProtoMessage()    {}
func (*UnlockRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{8}
}
func (m *UnlockRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockRequest.Unmarshal(m, b)
}
func (m *UnlockRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockRequest.Marshal(b, m, deterministic)
}
func (dst *UnlockRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockRequest.Merge(dst, src)
}
func (m *UnlockRequest) XXX_Size() int {
	return xxx_messageInfo_UnlockRequest.Size(m)
}
func (m *UnlockRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockRequest proto.InternalMessageInfo

func (m *UnlockRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UnlockRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type UnlockResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Unlocked             bool                    `protobuf:"varint,2,opt,name=unlocked,proto3" json:"unlocked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *UnlockResponse) Reset()         { *m = UnlockResponse{} }
func (m *UnlockResponse) String() string { return proto.CompactTextString(m) }
func (*UnlockResponse) ProtoMessage()    {}
func (*UnlockResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{9}
}
func (m *UnlockResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockResponse.Unmarshal(m, b)
}
func (m *UnlockResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockResponse.Marshal(b, m, deterministic)
}
func (dst *UnlockResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockResponse.Merge(dst, src)
}
func (m *UnlockResponse) XXX_Size() int {
	return xxx_messageInfo_UnlockResponse.Size(m)
}
func (m *UnlockResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockResponse proto.InternalMessageInfo

func (m *UnlockResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *UnlockResponse) GetUnlocked() bool {
	if m != nil {
		return m.Unlocked
	}
	return false
}

type IsLockedRequest struct {
	Header               *headers.RequestHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Version              uint64                 `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *IsLockedRequest) Reset()         { *m = IsLockedRequest{} }
func (m *IsLockedRequest) String() string { return proto.CompactTextString(m) }
func (*IsLockedRequest) ProtoMessage()    {}
func (*IsLockedRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{10}
}
func (m *IsLockedRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsLockedRequest.Unmarshal(m, b)
}
func (m *IsLockedRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsLockedRequest.Marshal(b, m, deterministic)
}
func (dst *IsLockedRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsLockedRequest.Merge(dst, src)
}
func (m *IsLockedRequest) XXX_Size() int {
	return xxx_messageInfo_IsLockedRequest.Size(m)
}
func (m *IsLockedRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IsLockedRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IsLockedRequest proto.InternalMessageInfo

func (m *IsLockedRequest) GetHeader() *headers.RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *IsLockedRequest) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

type IsLockedResponse struct {
	Header               *headers.ResponseHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	IsLocked             bool                    `protobuf:"varint,2,opt,name=isLocked,proto3" json:"isLocked,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *IsLockedResponse) Reset()         { *m = IsLockedResponse{} }
func (m *IsLockedResponse) String() string { return proto.CompactTextString(m) }
func (*IsLockedResponse) ProtoMessage()    {}
func (*IsLockedResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_lock_6337aa30f4fbad76, []int{11}
}
func (m *IsLockedResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IsLockedResponse.Unmarshal(m, b)
}
func (m *IsLockedResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IsLockedResponse.Marshal(b, m, deterministic)
}
func (dst *IsLockedResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IsLockedResponse.Merge(dst, src)
}
func (m *IsLockedResponse) XXX_Size() int {
	return xxx_messageInfo_IsLockedResponse.Size(m)
}
func (m *IsLockedResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_IsLockedResponse.DiscardUnknown(m)
}

var xxx_messageInfo_IsLockedResponse proto.InternalMessageInfo

func (m *IsLockedResponse) GetHeader() *headers.ResponseHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *IsLockedResponse) GetIsLocked() bool {
	if m != nil {
		return m.IsLocked
	}
	return false
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "atomix.lock.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "atomix.lock.CreateResponse")
	proto.RegisterType((*KeepAliveRequest)(nil), "atomix.lock.KeepAliveRequest")
	proto.RegisterType((*KeepAliveResponse)(nil), "atomix.lock.KeepAliveResponse")
	proto.RegisterType((*CloseRequest)(nil), "atomix.lock.CloseRequest")
	proto.RegisterType((*CloseResponse)(nil), "atomix.lock.CloseResponse")
	proto.RegisterType((*LockRequest)(nil), "atomix.lock.LockRequest")
	proto.RegisterType((*LockResponse)(nil), "atomix.lock.LockResponse")
	proto.RegisterType((*UnlockRequest)(nil), "atomix.lock.UnlockRequest")
	proto.RegisterType((*UnlockResponse)(nil), "atomix.lock.UnlockResponse")
	proto.RegisterType((*IsLockedRequest)(nil), "atomix.lock.IsLockedRequest")
	proto.RegisterType((*IsLockedResponse)(nil), "atomix.lock.IsLockedResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LockServiceClient is the client API for LockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LockServiceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error)
	Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error)
	Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error)
	IsLocked(ctx context.Context, in *IsLockedRequest, opts ...grpc.CallOption) (*IsLockedResponse, error)
}

type lockServiceClient struct {
	cc *grpc.ClientConn
}

func NewLockServiceClient(cc *grpc.ClientConn) LockServiceClient {
	return &lockServiceClient{cc}
}

func (c *lockServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) KeepAlive(ctx context.Context, in *KeepAliveRequest, opts ...grpc.CallOption) (*KeepAliveResponse, error) {
	out := new(KeepAliveResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/KeepAlive", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) Close(ctx context.Context, in *CloseRequest, opts ...grpc.CallOption) (*CloseResponse, error) {
	out := new(CloseResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) Lock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/Lock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) Unlock(ctx context.Context, in *UnlockRequest, opts ...grpc.CallOption) (*UnlockResponse, error) {
	out := new(UnlockResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/Unlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) IsLocked(ctx context.Context, in *IsLockedRequest, opts ...grpc.CallOption) (*IsLockedResponse, error) {
	out := new(IsLockedResponse)
	err := c.cc.Invoke(ctx, "/atomix.lock.LockService/IsLocked", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServiceServer is the server API for LockService service.
type LockServiceServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	KeepAlive(context.Context, *KeepAliveRequest) (*KeepAliveResponse, error)
	Close(context.Context, *CloseRequest) (*CloseResponse, error)
	Lock(context.Context, *LockRequest) (*LockResponse, error)
	Unlock(context.Context, *UnlockRequest) (*UnlockResponse, error)
	IsLocked(context.Context, *IsLockedRequest) (*IsLockedResponse, error)
}

func RegisterLockServiceServer(s *grpc.Server, srv LockServiceServer) {
	s.RegisterService(&_LockService_serviceDesc, srv)
}

func _LockService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_KeepAlive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).KeepAlive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/KeepAlive",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).KeepAlive(ctx, req.(*KeepAliveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).Close(ctx, req.(*CloseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_Lock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).Lock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/Lock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).Lock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_Unlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnlockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).Unlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/Unlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).Unlock(ctx, req.(*UnlockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_IsLocked_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsLockedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).IsLocked(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/atomix.lock.LockService/IsLocked",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).IsLocked(ctx, req.(*IsLockedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "atomix.lock.LockService",
	HandlerType: (*LockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LockService_Create_Handler,
		},
		{
			MethodName: "KeepAlive",
			Handler:    _LockService_KeepAlive_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _LockService_Close_Handler,
		},
		{
			MethodName: "Lock",
			Handler:    _LockService_Lock_Handler,
		},
		{
			MethodName: "Unlock",
			Handler:    _LockService_Unlock_Handler,
		},
		{
			MethodName: "IsLocked",
			Handler:    _LockService_IsLocked_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "atomix/lock/lock.proto",
}

func init() { proto.RegisterFile("atomix/lock/lock.proto", fileDescriptor_lock_6337aa30f4fbad76) }

var fileDescriptor_lock_6337aa30f4fbad76 = []byte{
	// 452 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x6d, 0xa1, 0xb8, 0x61, 0xd2, 0x94, 0xb2, 0x87, 0xca, 0x59, 0x9a, 0x08, 0xed, 0x89, 0xd3,
	0x46, 0x6a, 0x05, 0x37, 0x24, 0x90, 0x91, 0x68, 0xd4, 0x9c, 0x8c, 0x38, 0x22, 0xd5, 0x89, 0xa7,
	0xc1, 0x8a, 0xe3, 0x0d, 0xf6, 0x3a, 0x02, 0xf1, 0xab, 0xf8, 0x87, 0x28, 0xfb, 0x61, 0xbc, 0x26,
	0x39, 0xad, 0x10, 0x97, 0xc4, 0xe3, 0x79, 0xfb, 0xde, 0xbc, 0xf5, 0xcc, 0xc0, 0x65, 0x22, 0xc5,
	0x3a, 0xfb, 0x3e, 0xc9, 0xc5, 0x62, 0xa5, 0x7e, 0xf8, 0xa6, 0x14, 0x52, 0x90, 0xbe, 0x7e, 0xcf,
	0x77, 0xaf, 0xe8, 0x78, 0x29, 0xc4, 0x32, 0xc7, 0x89, 0x4a, 0xcd, 0xeb, 0x87, 0x49, 0x5a, 0x97,
	0x89, 0xcc, 0x44, 0xa1, 0xc1, 0xf4, 0xca, 0x90, 0x7c, 0xc5, 0x24, 0xc5, 0xb2, 0xb2, 0xff, 0x3a,
	0xcb, 0x7e, 0xc2, 0x20, 0x2a, 0x31, 0x91, 0x18, 0xe3, 0xb7, 0x1a, 0x2b, 0x49, 0x5e, 0x43, 0xa0,
	0x11, 0xe1, 0xf1, 0xcb, 0xe3, 0x57, 0xfd, 0xeb, 0x11, 0x37, 0x62, 0xf6, 0x9c, 0x01, 0xde, 0xaa,
	0x30, 0x36, 0x60, 0x72, 0x03, 0xa7, 0x32, 0x5b, 0xa3, 0xa8, 0x65, 0xf8, 0x48, 0x9d, 0x1b, 0x72,
	0x5d, 0x17, 0xb7, 0x75, 0xf1, 0x0f, 0xa6, 0xae, 0xd8, 0x22, 0xd9, 0x2d, 0x9c, 0x5b, 0xf1, 0x6a,
	0x23, 0x8a, 0x0a, 0xc9, 0x9b, 0x8e, 0xfa, 0xf8, 0x6f, 0x75, 0x8d, 0x74, 0xe5, 0xd9, 0x14, 0x2e,
	0xee, 0x10, 0x37, 0xef, 0xf3, 0x6c, 0xeb, 0xe9, 0x84, 0xdd, 0xc1, 0xf3, 0x16, 0x95, 0x67, 0x5d,
	0x5f, 0xe0, 0x2c, 0xca, 0x45, 0xe5, 0x7b, 0xbb, 0x97, 0x10, 0xa4, 0x98, 0xa3, 0x44, 0x75, 0xb9,
	0xbd, 0xd8, 0x44, 0xec, 0x23, 0x0c, 0x0c, 0xbd, 0x67, 0x9d, 0x3f, 0xa0, 0x3f, 0x13, 0x8b, 0xd5,
	0xff, 0x68, 0x82, 0x7b, 0x38, 0xd3, 0xd2, 0x7e, 0x16, 0x48, 0x08, 0xa7, 0x5b, 0x2c, 0xab, 0x4c,
	0x14, 0x4a, 0xfc, 0x24, 0xb6, 0x21, 0xbb, 0x87, 0xc1, 0xe7, 0x22, 0xf7, 0xb7, 0x77, 0x58, 0x21,
	0x85, 0x73, 0xab, 0xe0, 0xe9, 0x82, 0x42, 0xaf, 0x56, 0x4c, 0x98, 0x9a, 0x6f, 0xdd, 0xc4, 0x6c,
	0x0e, 0xcf, 0xa6, 0xd5, 0x4c, 0x3d, 0xff, 0x33, 0x27, 0x0f, 0x70, 0xf1, 0x47, 0xc3, 0xdf, 0x4b,
	0x66, 0xb8, 0xac, 0x17, 0x1b, 0x5f, 0xff, 0x7a, 0xac, 0x3b, 0xee, 0x13, 0x96, 0xdb, 0x6c, 0x81,
	0x24, 0x82, 0x40, 0xaf, 0x02, 0x42, 0x79, 0x6b, 0xbb, 0x71, 0x67, 0x39, 0xd1, 0x17, 0x7b, 0x73,
	0x5a, 0x9c, 0x1d, 0x91, 0x19, 0x3c, 0x6d, 0x46, 0x97, 0x8c, 0x1c, 0x6c, 0x77, 0x3b, 0xd0, 0xf1,
	0xa1, 0x74, 0xc3, 0xf6, 0x0e, 0x9e, 0xa8, 0xe1, 0x22, 0x43, 0x57, 0xb5, 0x35, 0xcf, 0x94, 0xee,
	0x4b, 0x35, 0x0c, 0x6f, 0xe1, 0x64, 0xe7, 0x91, 0x84, 0x0e, 0xaa, 0x35, 0x68, 0x74, 0xb8, 0x27,
	0xd3, 0x1c, 0x8f, 0x20, 0xd0, 0x5d, 0xd5, 0xb9, 0x13, 0xa7, 0x99, 0x3b, 0x77, 0xe2, 0xb6, 0x21,
	0x3b, 0x22, 0x53, 0xe8, 0xd9, 0x0f, 0x4a, 0xae, 0x1c, 0x68, 0xa7, 0x97, 0xe8, 0xe8, 0x40, 0xd6,
	0x52, 0xcd, 0x03, 0x35, 0xc5, 0x37, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbd, 0xfc, 0x76, 0x0c,
	0x97, 0x06, 0x00, 0x00,
}
