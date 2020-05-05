// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/demo/demo.proto

package demo

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type DemoRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoRequest) Reset()         { *m = DemoRequest{} }
func (m *DemoRequest) String() string { return proto.CompactTextString(m) }
func (*DemoRequest) ProtoMessage()    {}
func (*DemoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b73fb6bfa52be9, []int{0}
}

func (m *DemoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoRequest.Unmarshal(m, b)
}
func (m *DemoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoRequest.Marshal(b, m, deterministic)
}
func (m *DemoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoRequest.Merge(m, src)
}
func (m *DemoRequest) XXX_Size() int {
	return xxx_messageInfo_DemoRequest.Size(m)
}
func (m *DemoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DemoRequest proto.InternalMessageInfo

func (m *DemoRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type DemoResponse struct {
	Text                 string   `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DemoResponse) Reset()         { *m = DemoResponse{} }
func (m *DemoResponse) String() string { return proto.CompactTextString(m) }
func (*DemoResponse) ProtoMessage()    {}
func (*DemoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_08b73fb6bfa52be9, []int{1}
}

func (m *DemoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DemoResponse.Unmarshal(m, b)
}
func (m *DemoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DemoResponse.Marshal(b, m, deterministic)
}
func (m *DemoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemoResponse.Merge(m, src)
}
func (m *DemoResponse) XXX_Size() int {
	return xxx_messageInfo_DemoResponse.Size(m)
}
func (m *DemoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DemoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DemoResponse proto.InternalMessageInfo

func (m *DemoResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func init() {
	proto.RegisterType((*DemoRequest)(nil), "DemoRequest")
	proto.RegisterType((*DemoResponse)(nil), "DemoResponse")
}

func init() { proto.RegisterFile("proto/demo/demo.proto", fileDescriptor_08b73fb6bfa52be9) }

var fileDescriptor_08b73fb6bfa52be9 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0x49, 0xcd, 0x85, 0x10, 0x7a, 0x60, 0xbe, 0x92, 0x22, 0x17, 0xb7, 0x4b, 0x6a,
	0x6e, 0x7e, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90, 0x10, 0x17, 0x4b, 0x5e, 0x62, 0x6e,
	0xaa, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x98, 0xad, 0xa4, 0xc4, 0xc5, 0x03, 0x51, 0x52,
	0x5c, 0x90, 0x9f, 0x57, 0x9c, 0x0a, 0x52, 0x53, 0x92, 0x5a, 0x51, 0x22, 0xc1, 0x04, 0x51, 0x03,
	0x62, 0x1b, 0x59, 0x40, 0x8c, 0x09, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0xd2, 0xe4, 0xe2,
	0x08, 0x4e, 0xac, 0xf4, 0x48, 0xcd, 0xc9, 0xc9, 0x17, 0xe2, 0xd1, 0x43, 0xb2, 0x40, 0x8a, 0x57,
	0x0f, 0xd9, 0x2c, 0x25, 0x06, 0x27, 0x9e, 0x28, 0x2e, 0x84, 0xcb, 0x92, 0xd8, 0xc0, 0x6c, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xa2, 0xe0, 0xb3, 0xc3, 0xae, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for DemoService service

type DemoServiceClient interface {
	SayHello(ctx context.Context, in *DemoRequest, opts ...client.CallOption) (*DemoResponse, error)
}

type demoServiceClient struct {
	c           client.Client
	serviceName string
}

func NewDemoServiceClient(serviceName string, c client.Client) DemoServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "demoservice"
	}
	return &demoServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *demoServiceClient) SayHello(ctx context.Context, in *DemoRequest, opts ...client.CallOption) (*DemoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "DemoService.SayHello", in)
	out := new(DemoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DemoService service

type DemoServiceHandler interface {
	SayHello(context.Context, *DemoRequest, *DemoResponse) error
}

func RegisterDemoServiceHandler(s server.Server, hdlr DemoServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&DemoService{hdlr}, opts...))
}

type DemoService struct {
	DemoServiceHandler
}

func (h *DemoService) SayHello(ctx context.Context, in *DemoRequest, out *DemoResponse) error {
	return h.DemoServiceHandler.SayHello(ctx, in, out)
}
