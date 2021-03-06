// Code generated by protoc-gen-go. DO NOT EDIT.
// source: say.proto

package say

import (
	context "context"
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type SayRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayRequest) Reset()         { *m = SayRequest{} }
func (m *SayRequest) String() string { return proto.CompactTextString(m) }
func (*SayRequest) ProtoMessage()    {}
func (*SayRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dad255d0361606e, []int{0}
}

func (m *SayRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayRequest.Unmarshal(m, b)
}
func (m *SayRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayRequest.Marshal(b, m, deterministic)
}
func (m *SayRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayRequest.Merge(m, src)
}
func (m *SayRequest) XXX_Size() int {
	return xxx_messageInfo_SayRequest.Size(m)
}
func (m *SayRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SayRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SayRequest proto.InternalMessageInfo

func (m *SayRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type SayResponse struct {
	Audio                []byte   `protobuf:"bytes,1,opt,name=audio,proto3" json:"audio,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SayResponse) Reset()         { *m = SayResponse{} }
func (m *SayResponse) String() string { return proto.CompactTextString(m) }
func (*SayResponse) ProtoMessage()    {}
func (*SayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6dad255d0361606e, []int{1}
}

func (m *SayResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SayResponse.Unmarshal(m, b)
}
func (m *SayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SayResponse.Marshal(b, m, deterministic)
}
func (m *SayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SayResponse.Merge(m, src)
}
func (m *SayResponse) XXX_Size() int {
	return xxx_messageInfo_SayResponse.Size(m)
}
func (m *SayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SayResponse proto.InternalMessageInfo

func (m *SayResponse) GetAudio() []byte {
	if m != nil {
		return m.Audio
	}
	return nil
}

func init() {
	proto.RegisterType((*SayRequest)(nil), "say.SayRequest")
	proto.RegisterType((*SayResponse)(nil), "say.SayResponse")
}

func init() { proto.RegisterFile("say.proto", fileDescriptor_6dad255d0361606e) }

var fileDescriptor_6dad255d0361606e = []byte{
	// 141 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x4e, 0xac, 0xd4,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x4e, 0xac, 0x54, 0x52, 0xe0, 0xe2, 0x0a, 0x4e,
	0xac, 0x0c, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe2, 0x62, 0x29, 0x49, 0xad, 0x28,
	0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0xb3, 0x95, 0x94, 0xb9, 0xb8, 0xc1, 0x2a, 0x8a,
	0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x44, 0xb8, 0x58, 0x13, 0x4b, 0x53, 0x32, 0xf3, 0xc1, 0x6a,
	0x78, 0x82, 0x20, 0x1c, 0x23, 0x2b, 0x2e, 0x9e, 0x90, 0xd4, 0x8a, 0x92, 0x90, 0xfc, 0xe0, 0x82,
	0xd4, 0xd4, 0xe4, 0x0c, 0x21, 0x2d, 0x2e, 0xe6, 0xe0, 0xc4, 0x4a, 0x21, 0x7e, 0x3d, 0x90, 0x75,
	0x08, 0x0b, 0xa4, 0x04, 0x10, 0x02, 0x10, 0xf3, 0x94, 0x18, 0x92, 0xd8, 0xc0, 0xce, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x88, 0xfc, 0x18, 0xd1, 0x9b, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TextToSpeechClient is the client API for TextToSpeech service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TextToSpeechClient interface {
	Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error)
}

type textToSpeechClient struct {
	cc *grpc.ClientConn
}

func NewTextToSpeechClient(cc *grpc.ClientConn) TextToSpeechClient {
	return &textToSpeechClient{cc}
}

func (c *textToSpeechClient) Say(ctx context.Context, in *SayRequest, opts ...grpc.CallOption) (*SayResponse, error) {
	out := new(SayResponse)
	err := c.cc.Invoke(ctx, "/say.TextToSpeech/Say", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextToSpeechServer is the server API for TextToSpeech service.
type TextToSpeechServer interface {
	Say(context.Context, *SayRequest) (*SayResponse, error)
}

func RegisterTextToSpeechServer(s *grpc.Server, srv TextToSpeechServer) {
	s.RegisterService(&_TextToSpeech_serviceDesc, srv)
}

func _TextToSpeech_Say_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SayRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextToSpeechServer).Say(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/say.TextToSpeech/Say",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextToSpeechServer).Say(ctx, req.(*SayRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TextToSpeech_serviceDesc = grpc.ServiceDesc{
	ServiceName: "say.TextToSpeech",
	HandlerType: (*TextToSpeechServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Say",
			Handler:    _TextToSpeech_Say_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "say.proto",
}
