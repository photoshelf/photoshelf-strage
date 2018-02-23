// Code generated by protoc-gen-go. DO NOT EDIT.
// source: photos.proto

/*
Package protobuf is a generated protocol buffer package.

It is generated from these files:
	photos.proto

It has these top-level messages:
	Id
	Photo
	Empty
*/
package protobuf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

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

type Id struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Photo struct {
	Id    *Id    `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Image []byte `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
}

func (m *Photo) Reset()                    { *m = Photo{} }
func (m *Photo) String() string            { return proto.CompactTextString(m) }
func (*Photo) ProtoMessage()               {}
func (*Photo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Photo) GetId() *Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Photo) GetImage() []byte {
	if m != nil {
		return m.Image
	}
	return nil
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func init() {
	proto.RegisterType((*Id)(nil), "protobuf.Id")
	proto.RegisterType((*Photo)(nil), "protobuf.Photo")
	proto.RegisterType((*Empty)(nil), "protobuf.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PhotoService service

type PhotoServiceClient interface {
	Create(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Id, error)
	Update(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Empty, error)
	Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Photo, error)
	Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
}

type photoServiceClient struct {
	cc *grpc.ClientConn
}

func NewPhotoServiceClient(cc *grpc.ClientConn) PhotoServiceClient {
	return &photoServiceClient{cc}
}

func (c *photoServiceClient) Create(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/protobuf.PhotoService/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) Update(ctx context.Context, in *Photo, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protobuf.PhotoService/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) Read(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Photo, error) {
	out := new(Photo)
	err := grpc.Invoke(ctx, "/protobuf.PhotoService/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoServiceClient) Delete(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := grpc.Invoke(ctx, "/protobuf.PhotoService/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PhotoService service

type PhotoServiceServer interface {
	Create(context.Context, *Photo) (*Id, error)
	Update(context.Context, *Photo) (*Empty, error)
	Read(context.Context, *Id) (*Photo, error)
	Delete(context.Context, *Id) (*Empty, error)
}

func RegisterPhotoServiceServer(s *grpc.Server, srv PhotoServiceServer) {
	s.RegisterService(&_PhotoService_serviceDesc, srv)
}

func _PhotoService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Photo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.PhotoService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).Create(ctx, req.(*Photo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Photo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.PhotoService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).Update(ctx, req.(*Photo))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.PhotoService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).Read(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.PhotoService/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoServiceServer).Delete(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _PhotoService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.PhotoService",
	HandlerType: (*PhotoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _PhotoService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _PhotoService_Update_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _PhotoService_Read_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PhotoService_Delete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photos.proto",
}

func init() { proto.RegisterFile("photos.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xcf, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x06, 0x60, 0x1a, 0x6c, 0xd5, 0x63, 0x51, 0x88, 0x22, 0x5d, 0x19, 0x63, 0x04, 0x2f, 0xc6,
	0x2e, 0x5a, 0x9c, 0x77, 0x7a, 0xa7, 0x0e, 0xd9, 0xdd, 0xa8, 0xf8, 0x00, 0x99, 0x39, 0xd6, 0x40,
	0xd7, 0x94, 0x9a, 0x15, 0x44, 0xbc, 0xf1, 0x15, 0x7c, 0x15, 0xdf, 0xc4, 0x57, 0xf0, 0x41, 0x64,
	0x27, 0x13, 0xed, 0x70, 0x57, 0xe1, 0x24, 0x7f, 0x3e, 0xfe, 0x03, 0x61, 0xf5, 0x68, 0xac, 0x79,
	0x4a, 0xaa, 0xda, 0x58, 0xc3, 0x77, 0xe8, 0x98, 0x2d, 0x1e, 0xe2, 0x6e, 0x6e, 0x4c, 0x5e, 0x60,
	0x2a, 0x2b, 0x9d, 0xca, 0xb2, 0x34, 0x56, 0x5a, 0x6d, 0xca, 0x55, 0x4e, 0xc4, 0xc0, 0x26, 0x8a,
	0x1f, 0x81, 0xdf, 0xc8, 0x62, 0x81, 0x91, 0xd7, 0xf7, 0x06, 0xbb, 0x99, 0x1b, 0xc4, 0x05, 0xf8,
	0xd3, 0xa5, 0xc9, 0xbb, 0xc0, 0xb4, 0xa2, 0xb7, 0xbd, 0x51, 0x98, 0xfc, 0xc8, 0xc9, 0x44, 0x65,
	0x4c, 0xd3, 0x67, 0x3d, 0x97, 0x39, 0x46, 0xac, 0xef, 0x0d, 0xc2, 0xcc, 0x0d, 0x62, 0x1b, 0xfc,
	0xf1, 0xbc, 0xb2, 0xcf, 0xa3, 0x0f, 0x06, 0x21, 0x31, 0xb7, 0x58, 0x37, 0xfa, 0x1e, 0xf9, 0x25,
	0x04, 0x57, 0x35, 0x4a, 0x8b, 0xfc, 0xe0, 0xd7, 0xa2, 0x44, 0xdc, 0xc2, 0x45, 0xe7, 0xed, 0xf3,
	0xeb, 0x9d, 0x1d, 0x8a, 0x7d, 0xea, 0xde, 0x9c, 0xa6, 0x6e, 0xc3, 0x73, 0x6f, 0xc8, 0xa7, 0x10,
	0xdc, 0x55, 0xea, 0x5f, 0xe3, 0xcf, 0x05, 0x15, 0x10, 0x27, 0xc4, 0xf4, 0xe2, 0x4e, 0x9b, 0x49,
	0x5f, 0xb4, 0x4a, 0x68, 0xd1, 0xd7, 0xa5, 0x38, 0x86, 0xad, 0x0c, 0xa5, 0xe2, 0xad, 0x0a, 0xf1,
	0xba, 0x2e, 0x7a, 0x84, 0x45, 0xfc, 0x78, 0x1d, 0x73, 0x12, 0xbf, 0x81, 0xe0, 0x1a, 0x0b, 0xb4,
	0xb8, 0x19, 0x72, 0xad, 0x56, 0xd0, 0x70, 0x03, 0x34, 0x0b, 0x28, 0x7f, 0xf6, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0xd9, 0x5c, 0x73, 0x40, 0xd7, 0x01, 0x00, 0x00,
}
