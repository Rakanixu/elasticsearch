// Code generated by protoc-gen-go.
// source: github.com/kazoup/elastic/srv/proto/elastic/elastic.proto
// DO NOT EDIT!

/*
Package go_micro_srv_elastic is a generated protocol buffer package.

It is generated from these files:
	github.com/kazoup/elastic/srv/proto/elastic/elastic.proto

It has these top-level messages:
	DocRef
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	UpdateRequest
	UpdateResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	QueryRequest
	QueryResponse
*/
package go_micro_srv_elastic

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion1

type DocRef struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *DocRef) Reset()                    { *m = DocRef{} }
func (m *DocRef) String() string            { return proto.CompactTextString(m) }
func (*DocRef) ProtoMessage()               {}
func (*DocRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type CreateRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Data  string `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ReadRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ReadResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type UpdateRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
	Data  string `protobuf:"bytes,4,opt,name=data" json:"data,omitempty"`
}

func (m *UpdateRequest) Reset()                    { *m = UpdateRequest{} }
func (m *UpdateRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateRequest) ProtoMessage()               {}
func (*UpdateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type UpdateResponse struct {
}

func (m *UpdateResponse) Reset()                    { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()               {}
func (*UpdateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type DeleteRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Id    string `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type SearchRequest struct {
	Index  string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type   string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Query  string `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
	Limit  int64  `protobuf:"varint,4,opt,name=limit" json:"limit,omitempty"`
	Offset int64  `protobuf:"varint,5,opt,name=offset" json:"offset,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SearchResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type QueryRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index" json:"index,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=type" json:"type,omitempty"`
	Query string `protobuf:"bytes,3,opt,name=query" json:"query,omitempty"`
}

func (m *QueryRequest) Reset()                    { *m = QueryRequest{} }
func (m *QueryRequest) String() string            { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()               {}
func (*QueryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type QueryResponse struct {
	Result string `protobuf:"bytes,1,opt,name=result" json:"result,omitempty"`
}

func (m *QueryResponse) Reset()                    { *m = QueryResponse{} }
func (m *QueryResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()               {}
func (*QueryResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func init() {
	proto.RegisterType((*DocRef)(nil), "go.micro.srv.elastic.DocRef")
	proto.RegisterType((*CreateRequest)(nil), "go.micro.srv.elastic.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.srv.elastic.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "go.micro.srv.elastic.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "go.micro.srv.elastic.ReadResponse")
	proto.RegisterType((*UpdateRequest)(nil), "go.micro.srv.elastic.UpdateRequest")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.srv.elastic.UpdateResponse")
	proto.RegisterType((*DeleteRequest)(nil), "go.micro.srv.elastic.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.srv.elastic.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "go.micro.srv.elastic.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "go.micro.srv.elastic.SearchResponse")
	proto.RegisterType((*QueryRequest)(nil), "go.micro.srv.elastic.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "go.micro.srv.elastic.QueryResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Elastic service

type ElasticClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error)
}

type elasticClient struct {
	c           client.Client
	serviceName string
}

func NewElasticClient(serviceName string, c client.Client) ElasticClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.elastic"
	}
	return &elasticClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *elasticClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticClient) Update(ctx context.Context, in *UpdateRequest, opts ...client.CallOption) (*UpdateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Update", in)
	out := new(UpdateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *elasticClient) Query(ctx context.Context, in *QueryRequest, opts ...client.CallOption) (*QueryResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Elastic.Query", in)
	out := new(QueryResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Elastic service

type ElasticHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Update(context.Context, *UpdateRequest, *UpdateResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Query(context.Context, *QueryRequest, *QueryResponse) error
}

func RegisterElasticHandler(s server.Server, hdlr ElasticHandler) {
	s.Handle(s.NewHandler(&Elastic{hdlr}))
}

type Elastic struct {
	ElasticHandler
}

func (h *Elastic) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.ElasticHandler.Create(ctx, in, out)
}

func (h *Elastic) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.ElasticHandler.Read(ctx, in, out)
}

func (h *Elastic) Update(ctx context.Context, in *UpdateRequest, out *UpdateResponse) error {
	return h.ElasticHandler.Update(ctx, in, out)
}

func (h *Elastic) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.ElasticHandler.Delete(ctx, in, out)
}

func (h *Elastic) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.ElasticHandler.Search(ctx, in, out)
}

func (h *Elastic) Query(ctx context.Context, in *QueryRequest, out *QueryResponse) error {
	return h.ElasticHandler.Query(ctx, in, out)
}

var fileDescriptor0 = []byte{
	// 416 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xb4, 0x94, 0x4f, 0x4f, 0xb3, 0x40,
	0x10, 0xc6, 0xdf, 0xb6, 0x94, 0x37, 0x8e, 0xa5, 0x31, 0x9b, 0xc6, 0x90, 0x9e, 0x74, 0x6b, 0xb4,
	0x27, 0x48, 0xf4, 0xe4, 0x55, 0x6b, 0x8c, 0x17, 0x8d, 0x98, 0x1e, 0x3d, 0x50, 0xd8, 0xb6, 0x44,
	0xda, 0x45, 0x58, 0x8c, 0x35, 0xf1, 0x0b, 0xf8, 0xa9, 0x65, 0xff, 0xd0, 0x94, 0xa6, 0x60, 0xd4,
	0x7a, 0x62, 0x67, 0x78, 0xf8, 0xcd, 0xcc, 0xee, 0xb3, 0xc0, 0xf9, 0x24, 0x60, 0xd3, 0x74, 0x64,
	0x79, 0x74, 0x66, 0x3f, 0xb9, 0x6f, 0x34, 0x8d, 0x6c, 0x12, 0xba, 0x09, 0x0b, 0x3c, 0x3b, 0x89,
	0x5f, 0xec, 0x28, 0xa6, 0x8c, 0x2e, 0x33, 0xea, 0x69, 0x89, 0x2c, 0xea, 0x4c, 0xa8, 0x35, 0x0b,
	0xbc, 0x98, 0x5a, 0x99, 0xd2, 0x52, 0xef, 0xf0, 0x05, 0xe8, 0x03, 0xea, 0x39, 0x64, 0x8c, 0x3a,
	0xd0, 0x0c, 0xe6, 0x3e, 0x79, 0x35, 0x6b, 0x07, 0xb5, 0xfe, 0x8e, 0x23, 0x03, 0x84, 0x40, 0x63,
	0x8b, 0x88, 0x98, 0x75, 0x91, 0x14, 0x6b, 0xd4, 0x86, 0x7a, 0xe0, 0x9b, 0x0d, 0x91, 0xc9, 0x56,
	0xf8, 0x11, 0x8c, 0xcb, 0x98, 0xb8, 0x8c, 0x38, 0xe4, 0x39, 0x25, 0x09, 0xfb, 0x39, 0x8a, 0x6b,
	0x7c, 0x97, 0xb9, 0xa6, 0x26, 0x35, 0x7c, 0x8d, 0xf7, 0xa0, 0x9d, 0xe3, 0x93, 0x88, 0xce, 0x13,
	0x82, 0xaf, 0x61, 0xd7, 0x21, 0xae, 0xff, 0xeb, 0x72, 0xf8, 0x18, 0x5a, 0x12, 0x24, 0xc1, 0x68,
	0x1f, 0xf4, 0x98, 0x24, 0x69, 0xc8, 0x14, 0x4a, 0x45, 0x7c, 0xc2, 0x61, 0xe4, 0xff, 0xe5, 0x84,
	0x39, 0x5e, 0x4d, 0x78, 0x03, 0xc6, 0x80, 0x84, 0x64, 0x0b, 0x05, 0x39, 0x3c, 0x47, 0x29, 0xf8,
	0x3b, 0x18, 0x0f, 0xc4, 0x8d, 0xbd, 0xe9, 0xf7, 0xe1, 0x99, 0x32, 0xfb, 0x24, 0x5e, 0x28, 0xbe,
	0x0c, 0x78, 0x36, 0x0c, 0x66, 0x01, 0x13, 0x43, 0x35, 0x1c, 0x19, 0xf0, 0xcd, 0xa4, 0xe3, 0x71,
	0x42, 0x98, 0xd9, 0x14, 0x69, 0x15, 0xe1, 0x3e, 0xb4, 0xf3, 0xf2, 0x5f, 0x6c, 0xfb, 0x2d, 0xb4,
	0xee, 0x79, 0x81, 0x2d, 0xf5, 0x89, 0x4f, 0xc0, 0x50, 0xbc, 0xea, 0xc2, 0xa7, 0x1f, 0x1a, 0xfc,
	0xbf, 0x92, 0x37, 0x04, 0x0d, 0x41, 0x97, 0xf6, 0x43, 0x3d, 0x6b, 0xd3, 0x15, 0xb2, 0x0a, 0xde,
	0xef, 0x1e, 0x55, 0x8b, 0xd4, 0x11, 0xfc, 0x43, 0x77, 0xa0, 0x71, 0xeb, 0xa1, 0xc3, 0xcd, 0xfa,
	0x15, 0x7f, 0x77, 0x71, 0x95, 0x64, 0x09, 0xcc, 0xfa, 0x94, 0x26, 0x2a, 0xeb, 0xb3, 0xe0, 0xe0,
	0xb2, 0x3e, 0xd7, 0x7c, 0x28, 0xb0, 0xd2, 0x3e, 0x65, 0xd8, 0x82, 0x4f, 0xcb, 0xb0, 0x6b, 0x0e,
	0x14, 0x58, 0x69, 0x82, 0x32, 0x6c, 0xc1, 0xa1, 0x65, 0xd8, 0xa2, 0x8f, 0x32, 0xac, 0x03, 0x4d,
	0x71, 0xc2, 0xa8, 0x64, 0xcf, 0x56, 0xed, 0xd4, 0xed, 0x55, 0x6a, 0x72, 0xe6, 0x48, 0x17, 0xff,
	0xcf, 0xb3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x2e, 0x44, 0x40, 0x63, 0x7c, 0x05, 0x00, 0x00,
}
