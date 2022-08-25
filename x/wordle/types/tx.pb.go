// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wordle/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgSubmitWordle struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Word    string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (m *MsgSubmitWordle) Reset()         { *m = MsgSubmitWordle{} }
func (m *MsgSubmitWordle) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitWordle) ProtoMessage()    {}
func (*MsgSubmitWordle) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6a6b5378ea72e2, []int{0}
}
func (m *MsgSubmitWordle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitWordle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitWordle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitWordle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitWordle.Merge(m, src)
}
func (m *MsgSubmitWordle) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitWordle) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitWordle.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitWordle proto.InternalMessageInfo

func (m *MsgSubmitWordle) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitWordle) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type MsgSubmitWordleResponse struct {
}

func (m *MsgSubmitWordleResponse) Reset()         { *m = MsgSubmitWordleResponse{} }
func (m *MsgSubmitWordleResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitWordleResponse) ProtoMessage()    {}
func (*MsgSubmitWordleResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6a6b5378ea72e2, []int{1}
}
func (m *MsgSubmitWordleResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitWordleResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitWordleResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitWordleResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitWordleResponse.Merge(m, src)
}
func (m *MsgSubmitWordleResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitWordleResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitWordleResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitWordleResponse proto.InternalMessageInfo

type MsgSubmitGuess struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Word    string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
}

func (m *MsgSubmitGuess) Reset()         { *m = MsgSubmitGuess{} }
func (m *MsgSubmitGuess) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitGuess) ProtoMessage()    {}
func (*MsgSubmitGuess) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6a6b5378ea72e2, []int{2}
}
func (m *MsgSubmitGuess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitGuess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitGuess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitGuess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitGuess.Merge(m, src)
}
func (m *MsgSubmitGuess) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitGuess) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitGuess.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitGuess proto.InternalMessageInfo

func (m *MsgSubmitGuess) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgSubmitGuess) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

type MsgSubmitGuessResponse struct {
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Body  string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"`
}

func (m *MsgSubmitGuessResponse) Reset()         { *m = MsgSubmitGuessResponse{} }
func (m *MsgSubmitGuessResponse) String() string { return proto.CompactTextString(m) }
func (*MsgSubmitGuessResponse) ProtoMessage()    {}
func (*MsgSubmitGuessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1c6a6b5378ea72e2, []int{3}
}
func (m *MsgSubmitGuessResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSubmitGuessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSubmitGuessResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSubmitGuessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSubmitGuessResponse.Merge(m, src)
}
func (m *MsgSubmitGuessResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgSubmitGuessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSubmitGuessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSubmitGuessResponse proto.InternalMessageInfo

func (m *MsgSubmitGuessResponse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgSubmitGuessResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgSubmitWordle)(nil), "yazzyyaz.wordle.wordle.MsgSubmitWordle")
	proto.RegisterType((*MsgSubmitWordleResponse)(nil), "yazzyyaz.wordle.wordle.MsgSubmitWordleResponse")
	proto.RegisterType((*MsgSubmitGuess)(nil), "yazzyyaz.wordle.wordle.MsgSubmitGuess")
	proto.RegisterType((*MsgSubmitGuessResponse)(nil), "yazzyyaz.wordle.wordle.MsgSubmitGuessResponse")
}

func init() { proto.RegisterFile("wordle/tx.proto", fileDescriptor_1c6a6b5378ea72e2) }

var fileDescriptor_1c6a6b5378ea72e2 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xcf, 0x2f, 0x4a,
	0xc9, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xab, 0x4c, 0xac,
	0xaa, 0xaa, 0xac, 0x4c, 0xac, 0xd2, 0x83, 0xc8, 0x40, 0x29, 0x25, 0x7b, 0x2e, 0x7e, 0xdf, 0xe2,
	0xf4, 0xe0, 0xd2, 0xa4, 0xdc, 0xcc, 0x92, 0x70, 0xb0, 0x90, 0x90, 0x04, 0x17, 0x7b, 0x72, 0x51,
	0x6a, 0x62, 0x49, 0x7e, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x8c, 0x2b, 0x24, 0xc4,
	0xc5, 0x02, 0xd2, 0x26, 0xc1, 0x04, 0x16, 0x06, 0xb3, 0x95, 0x24, 0xb9, 0xc4, 0xd1, 0x0c, 0x08,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55, 0xb2, 0xe3, 0xe2, 0x83, 0x4b, 0xb9, 0x97, 0xa6,
	0x16, 0x17, 0x93, 0x68, 0xb4, 0x13, 0x97, 0x18, 0xaa, 0x7e, 0x98, 0xc9, 0x42, 0x22, 0x5c, 0xac,
	0x25, 0x99, 0x25, 0x39, 0xa9, 0x50, 0x53, 0x20, 0x1c, 0x90, 0x19, 0x49, 0xf9, 0x29, 0x95, 0x30,
	0x33, 0x40, 0x6c, 0xa3, 0x6b, 0x8c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x19, 0x5c, 0x3c, 0x28,
	0x9e, 0x54, 0xd7, 0xc3, 0x1e, 0x20, 0x7a, 0x68, 0x9e, 0x91, 0xd2, 0x27, 0x52, 0x21, 0xdc, 0x6d,
	0xa9, 0x5c, 0xdc, 0xc8, 0x5e, 0x56, 0x23, 0xa8, 0x1f, 0xac, 0x4e, 0x4a, 0x8f, 0x38, 0x75, 0x30,
	0x6b, 0x9c, 0x9c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6,
	0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x33, 0x3d,
	0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x3f, 0x12, 0x64, 0x66, 0x64, 0x62, 0x95,
	0x3e, 0x34, 0x3d, 0x54, 0xc0, 0x18, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0xc4, 0x61,
	0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x17, 0xfa, 0xcc, 0x62, 0x2f, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	SubmitWordle(ctx context.Context, in *MsgSubmitWordle, opts ...grpc.CallOption) (*MsgSubmitWordleResponse, error)
	SubmitGuess(ctx context.Context, in *MsgSubmitGuess, opts ...grpc.CallOption) (*MsgSubmitGuessResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) SubmitWordle(ctx context.Context, in *MsgSubmitWordle, opts ...grpc.CallOption) (*MsgSubmitWordleResponse, error) {
	out := new(MsgSubmitWordleResponse)
	err := c.cc.Invoke(ctx, "/yazzyyaz.wordle.wordle.Msg/SubmitWordle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) SubmitGuess(ctx context.Context, in *MsgSubmitGuess, opts ...grpc.CallOption) (*MsgSubmitGuessResponse, error) {
	out := new(MsgSubmitGuessResponse)
	err := c.cc.Invoke(ctx, "/yazzyyaz.wordle.wordle.Msg/SubmitGuess", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	SubmitWordle(context.Context, *MsgSubmitWordle) (*MsgSubmitWordleResponse, error)
	SubmitGuess(context.Context, *MsgSubmitGuess) (*MsgSubmitGuessResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) SubmitWordle(ctx context.Context, req *MsgSubmitWordle) (*MsgSubmitWordleResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitWordle not implemented")
}
func (*UnimplementedMsgServer) SubmitGuess(ctx context.Context, req *MsgSubmitGuess) (*MsgSubmitGuessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitGuess not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_SubmitWordle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitWordle)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitWordle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yazzyyaz.wordle.wordle.Msg/SubmitWordle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitWordle(ctx, req.(*MsgSubmitWordle))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_SubmitGuess_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgSubmitGuess)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).SubmitGuess(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/yazzyyaz.wordle.wordle.Msg/SubmitGuess",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).SubmitGuess(ctx, req.(*MsgSubmitGuess))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "yazzyyaz.wordle.wordle.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitWordle",
			Handler:    _Msg_SubmitWordle_Handler,
		},
		{
			MethodName: "SubmitGuess",
			Handler:    _Msg_SubmitGuess_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "wordle/tx.proto",
}

func (m *MsgSubmitWordle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitWordle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitWordle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Word) > 0 {
		i -= len(m.Word)
		copy(dAtA[i:], m.Word)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Word)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitWordleResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitWordleResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitWordleResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgSubmitGuess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitGuess) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitGuess) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Word) > 0 {
		i -= len(m.Word)
		copy(dAtA[i:], m.Word)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Word)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSubmitGuessResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSubmitGuessResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSubmitGuessResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Body) > 0 {
		i -= len(m.Body)
		copy(dAtA[i:], m.Body)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Body)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgSubmitWordle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Word)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitWordleResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgSubmitGuess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Word)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgSubmitGuessResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Body)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgSubmitWordle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitWordle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitWordle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Word", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Word = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitWordleResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitWordleResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitWordleResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitGuess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitGuess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitGuess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Word", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Word = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgSubmitGuessResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgSubmitGuessResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSubmitGuessResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Body = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
