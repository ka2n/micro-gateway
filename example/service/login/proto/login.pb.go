// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/login.proto

/*
Package go_micro_example_login is a generated protocol buffer package.

It is generated from these files:
	proto/login.proto

It has these top-level messages:
	Request
	Response
	GetResponse
*/
package go_micro_example_login

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/ka2n/micro-gateway/protobuf/go/micro/gw"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	Hello string `protobuf:"bytes,1,opt,name=hello" json:"hello,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

type Response struct {
	JwtToken string `protobuf:"bytes,1,opt,name=jwt_token,json=jwtToken" json:"jwt_token,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetJwtToken() string {
	if m != nil {
		return m.JwtToken
	}
	return ""
}

type GetResponse struct {
	Hello string `protobuf:"bytes,1,opt,name=hello" json:"hello,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetResponse) GetHello() string {
	if m != nil {
		return m.Hello
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.example.login.Request")
	proto.RegisterType((*Response)(nil), "go.micro.example.login.Response")
	proto.RegisterType((*GetResponse)(nil), "go.micro.example.login.GetResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Login service

type LoginClient interface {
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*GetResponse, error)
}

type loginClient struct {
	c           client.Client
	serviceName string
}

func NewLoginClient(serviceName string, c client.Client) LoginClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.example.login"
	}
	return &loginClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *loginClient) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "Login.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Login.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginHandler interface {
	Create(context.Context, *Request, *Response) error
	Get(context.Context, *Request, *GetResponse) error
}

func RegisterLoginHandler(s server.Server, hdlr LoginHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Login{hdlr}, opts...))
}

type Login struct {
	LoginHandler
}

func (h *Login) Create(ctx context.Context, in *Request, out *Response) error {
	return h.LoginHandler.Create(ctx, in, out)
}

func (h *Login) Get(ctx context.Context, in *Request, out *GetResponse) error {
	return h.LoginHandler.Get(ctx, in, out)
}

func init() { proto.RegisterFile("proto/login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x12, 0x2c, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0xcf, 0xc9, 0x4f, 0xcf, 0xcc, 0xd3, 0x03, 0xb3, 0x85, 0xc4, 0xd2, 0xf3, 0xf5, 0x72,
	0x33, 0x93, 0x8b, 0xf2, 0xf5, 0x52, 0x2b, 0x12, 0x73, 0x0b, 0x72, 0x52, 0xf5, 0xc0, 0xb2, 0x52,
	0xb2, 0xe9, 0xf9, 0xfa, 0x60, 0x71, 0xfd, 0xf4, 0x72, 0xfd, 0xc4, 0xbc, 0xbc, 0xfc, 0x92, 0xc4,
	0x92, 0xcc, 0xfc, 0xbc, 0x62, 0x88, 0x36, 0x25, 0x79, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2, 0xd4,
	0xe2, 0x12, 0x21, 0x11, 0x2e, 0xd6, 0x8c, 0xd4, 0x9c, 0x9c, 0x7c, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x08, 0x47, 0x49, 0x9d, 0x8b, 0x23, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x55,
	0x48, 0x9a, 0x8b, 0x33, 0xab, 0xbc, 0x24, 0xbe, 0x24, 0x3f, 0x3b, 0x35, 0x0f, 0xaa, 0x8a, 0x23,
	0xab, 0xbc, 0x24, 0x04, 0xc4, 0x57, 0x52, 0xe6, 0xe2, 0x76, 0x4f, 0x2d, 0x81, 0xab, 0xc5, 0x6a,
	0x9a, 0xd1, 0x5c, 0x26, 0x2e, 0x56, 0x1f, 0x90, 0xbb, 0x84, 0x72, 0xb8, 0xd8, 0x9c, 0x8b, 0x52,
	0x13, 0x4b, 0x52, 0x85, 0xe4, 0xf5, 0xb0, 0x3b, 0x5d, 0x0f, 0xea, 0x30, 0x29, 0x05, 0xdc, 0x0a,
	0x20, 0x96, 0x29, 0xc9, 0x37, 0xed, 0xfe, 0xbe, 0x93, 0x49, 0x52, 0x89, 0x53, 0xbf, 0xcc, 0x10,
	0x12, 0x2a, 0x56, 0x8c, 0x5a, 0x5e, 0x9c, 0x49, 0x89, 0xc5, 0x99, 0xc9, 0x8e, 0xa5, 0x25, 0x19,
	0x42, 0x79, 0x5c, 0xcc, 0xee, 0xa9, 0x25, 0x84, 0xad, 0x52, 0xc6, 0xa5, 0x00, 0xc9, 0x6b, 0x4a,
	0x8a, 0x60, 0xdb, 0xa4, 0x85, 0xd8, 0x40, 0xb6, 0xe5, 0xa6, 0x7a, 0x09, 0x82, 0xed, 0x09, 0xcf,
	0x2c, 0xc9, 0xc8, 0x2f, 0x2d, 0x71, 0x2d, 0x2a, 0xca, 0x2f, 0x92, 0x32, 0xbb, 0x04, 0x52, 0x62,
	0xc0, 0x25, 0xc4, 0x85, 0xe4, 0x08, 0x56, 0x30, 0x93, 0x4b, 0x8a, 0x0b, 0x53, 0x03, 0x54, 0x2e,
	0x89, 0x0d, 0x1c, 0x2b, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x11, 0x13, 0xe4, 0xe1,
	0x01, 0x00, 0x00,
}
