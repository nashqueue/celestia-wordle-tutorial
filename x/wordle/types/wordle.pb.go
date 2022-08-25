// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wordle/wordle.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Wordle struct {
	Index     string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Word      string `protobuf:"bytes,2,opt,name=word,proto3" json:"word,omitempty"`
	Submitter string `protobuf:"bytes,3,opt,name=submitter,proto3" json:"submitter,omitempty"`
}

func (m *Wordle) Reset()         { *m = Wordle{} }
func (m *Wordle) String() string { return proto.CompactTextString(m) }
func (*Wordle) ProtoMessage()    {}
func (*Wordle) Descriptor() ([]byte, []int) {
	return fileDescriptor_c04b02df72f111cd, []int{0}
}
func (m *Wordle) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wordle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wordle.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wordle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wordle.Merge(m, src)
}
func (m *Wordle) XXX_Size() int {
	return m.Size()
}
func (m *Wordle) XXX_DiscardUnknown() {
	xxx_messageInfo_Wordle.DiscardUnknown(m)
}

var xxx_messageInfo_Wordle proto.InternalMessageInfo

func (m *Wordle) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Wordle) GetWord() string {
	if m != nil {
		return m.Word
	}
	return ""
}

func (m *Wordle) GetSubmitter() string {
	if m != nil {
		return m.Submitter
	}
	return ""
}

func init() {
	proto.RegisterType((*Wordle)(nil), "yazzyyaz.wordle.wordle.Wordle")
}

func init() { proto.RegisterFile("wordle/wordle.proto", fileDescriptor_c04b02df72f111cd) }

var fileDescriptor_c04b02df72f111cd = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0xcf, 0x2f, 0x4a,
	0xc9, 0x49, 0xd5, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x62, 0x95, 0x89, 0x55,
	0x55, 0x95, 0x95, 0x89, 0x55, 0x7a, 0x50, 0x61, 0x08, 0xa5, 0x14, 0xc0, 0xc5, 0x16, 0x0e, 0x66,
	0x09, 0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06,
	0x41, 0x38, 0x42, 0x42, 0x5c, 0x2c, 0x20, 0x95, 0x12, 0x4c, 0x60, 0x41, 0x30, 0x5b, 0x48, 0x86,
	0x8b, 0xb3, 0xb8, 0x34, 0x29, 0x37, 0xb3, 0xa4, 0x24, 0xb5, 0x48, 0x82, 0x19, 0x2c, 0x81, 0x10,
	0x70, 0x72, 0x3e, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xcd, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x48, 0x90, 0x73, 0x22, 0x13, 0xab, 0xa0,
	0xae, 0xd4, 0xaf, 0x80, 0x31, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xae, 0x36, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x9e, 0xcc, 0x66, 0x50, 0xcc, 0x00, 0x00, 0x00,
}

func (m *Wordle) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wordle) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wordle) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Submitter) > 0 {
		i -= len(m.Submitter)
		copy(dAtA[i:], m.Submitter)
		i = encodeVarintWordle(dAtA, i, uint64(len(m.Submitter)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Word) > 0 {
		i -= len(m.Word)
		copy(dAtA[i:], m.Word)
		i = encodeVarintWordle(dAtA, i, uint64(len(m.Word)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintWordle(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWordle(dAtA []byte, offset int, v uint64) int {
	offset -= sovWordle(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Wordle) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovWordle(uint64(l))
	}
	l = len(m.Word)
	if l > 0 {
		n += 1 + l + sovWordle(uint64(l))
	}
	l = len(m.Submitter)
	if l > 0 {
		n += 1 + l + sovWordle(uint64(l))
	}
	return n
}

func sovWordle(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWordle(x uint64) (n int) {
	return sovWordle(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wordle) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWordle
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
			return fmt.Errorf("proto: Wordle: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wordle: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordle
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
				return ErrInvalidLengthWordle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWordle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Word", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordle
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
				return ErrInvalidLengthWordle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWordle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Word = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submitter", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWordle
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
				return ErrInvalidLengthWordle
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWordle
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Submitter = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWordle(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWordle
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
func skipWordle(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWordle
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
					return 0, ErrIntOverflowWordle
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
					return 0, ErrIntOverflowWordle
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
				return 0, ErrInvalidLengthWordle
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWordle
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWordle
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWordle        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWordle          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWordle = fmt.Errorf("proto: unexpected end of group")
)
