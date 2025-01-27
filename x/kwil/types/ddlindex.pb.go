// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kwil/ddlindex.proto

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

type Ddlindex struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Position int64  `protobuf:"varint,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (m *Ddlindex) Reset()         { *m = Ddlindex{} }
func (m *Ddlindex) String() string { return proto.CompactTextString(m) }
func (*Ddlindex) ProtoMessage()    {}
func (*Ddlindex) Descriptor() ([]byte, []int) {
	return fileDescriptor_53b3ff87c0dd0fde, []int{0}
}
func (m *Ddlindex) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Ddlindex) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Ddlindex.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Ddlindex) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ddlindex.Merge(m, src)
}
func (m *Ddlindex) XXX_Size() int {
	return m.Size()
}
func (m *Ddlindex) XXX_DiscardUnknown() {
	xxx_messageInfo_Ddlindex.DiscardUnknown(m)
}

var xxx_messageInfo_Ddlindex proto.InternalMessageInfo

func (m *Ddlindex) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Ddlindex) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Ddlindex) GetPosition() int64 {
	if m != nil {
		return m.Position
	}
	return 0
}

func init() {
	proto.RegisterType((*Ddlindex)(nil), "kwil.kwil.Ddlindex")
}

func init() { proto.RegisterFile("kwil/ddlindex.proto", fileDescriptor_53b3ff87c0dd0fde) }

var fileDescriptor_53b3ff87c0dd0fde = []byte{
	// 153 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2e, 0xcf, 0xcc,
	0xd1, 0x4f, 0x49, 0xc9, 0xc9, 0xcc, 0x4b, 0x49, 0xad, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0xe2, 0x04, 0x09, 0xea, 0x81, 0x08, 0xa5, 0x00, 0x2e, 0x0e, 0x17, 0xa8, 0xa4, 0x90, 0x08, 0x17,
	0x2b, 0x98, 0x21, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x08, 0x09, 0x71, 0xb1, 0xe4,
	0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x81, 0x05, 0xc1, 0x6c, 0x21, 0x29, 0x2e, 0x8e, 0x82, 0xfc, 0xe2,
	0xcc, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x38, 0xdf, 0x49, 0xfb,
	0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e,
	0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0x04, 0xc1, 0x6e, 0xa9, 0xd0, 0x07,
	0x53, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x07, 0x19, 0x03, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x9e, 0x4b, 0xae, 0x95, 0xa7, 0x00, 0x00, 0x00,
}

func (m *Ddlindex) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Ddlindex) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Ddlindex) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Position != 0 {
		i = encodeVarintDdlindex(dAtA, i, uint64(m.Position))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintDdlindex(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintDdlindex(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDdlindex(dAtA []byte, offset int, v uint64) int {
	offset -= sovDdlindex(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Ddlindex) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovDdlindex(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovDdlindex(uint64(l))
	}
	if m.Position != 0 {
		n += 1 + sovDdlindex(uint64(m.Position))
	}
	return n
}

func sovDdlindex(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDdlindex(x uint64) (n int) {
	return sovDdlindex(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Ddlindex) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDdlindex
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
			return fmt.Errorf("proto: Ddlindex: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Ddlindex: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdlindex
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
				return ErrInvalidLengthDdlindex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDdlindex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdlindex
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
				return ErrInvalidLengthDdlindex
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDdlindex
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Position", wireType)
			}
			m.Position = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDdlindex
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Position |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipDdlindex(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDdlindex
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
func skipDdlindex(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDdlindex
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
					return 0, ErrIntOverflowDdlindex
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
					return 0, ErrIntOverflowDdlindex
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
				return 0, ErrInvalidLengthDdlindex
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDdlindex
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDdlindex
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDdlindex        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDdlindex          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDdlindex = fmt.Errorf("proto: unexpected end of group")
)
