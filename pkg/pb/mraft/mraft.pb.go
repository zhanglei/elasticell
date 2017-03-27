// Code generated by protoc-gen-gogo.
// source: mraft.proto
// DO NOT EDIT!

/*
	Package mraft is a generated protocol buffer package.

	It is generated from these files:
		mraft.proto

	It has these top-level messages:
		RaftMessage
*/
package mraft

import (
	"fmt"

	proto "github.com/golang/protobuf/proto"

	math "math"

	metapb "github.com/deepfabric/elasticell/pkg/pb/metapb"

	raftpb "github.com/coreos/etcd/raft/raftpb"

	io "io"
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

type RaftMessage struct {
	CellID    *uint64          `protobuf:"varint,1,opt,name=cellID" json:"cellID,omitempty"`
	FromPeer  metapb.Peer      `protobuf:"bytes,2,opt,name=fromPeer" json:"fromPeer"`
	ToPeer    metapb.Peer      `protobuf:"bytes,3,opt,name=toPeer" json:"toPeer"`
	Message   raftpb.Message   `protobuf:"bytes,4,opt,name=message" json:"message"`
	CellEpoch metapb.CellEpoch `protobuf:"bytes,5,opt,name=cellEpoch" json:"cellEpoch"`
	// true means to_peer is a tombstone peer and it should remove itself.
	IsTombstone *bool `protobuf:"varint,6,opt,name=isTombstone" json:"isTombstone,omitempty"`
	// Region key range [start_key, end_key).
	Start            []byte `protobuf:"bytes,7,opt,name=start" json:"start,omitempty"`
	End              []byte `protobuf:"bytes,8,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RaftMessage) Reset()                    { *m = RaftMessage{} }
func (m *RaftMessage) String() string            { return proto.CompactTextString(m) }
func (*RaftMessage) ProtoMessage()               {}
func (*RaftMessage) Descriptor() ([]byte, []int) { return fileDescriptorMraft, []int{0} }

func (m *RaftMessage) GetCellID() uint64 {
	if m != nil && m.CellID != nil {
		return *m.CellID
	}
	return 0
}

func (m *RaftMessage) GetFromPeer() metapb.Peer {
	if m != nil {
		return m.FromPeer
	}
	return metapb.Peer{}
}

func (m *RaftMessage) GetToPeer() metapb.Peer {
	if m != nil {
		return m.ToPeer
	}
	return metapb.Peer{}
}

func (m *RaftMessage) GetMessage() raftpb.Message {
	if m != nil {
		return m.Message
	}
	return raftpb.Message{}
}

func (m *RaftMessage) GetCellEpoch() metapb.CellEpoch {
	if m != nil {
		return m.CellEpoch
	}
	return metapb.CellEpoch{}
}

func (m *RaftMessage) GetIsTombstone() bool {
	if m != nil && m.IsTombstone != nil {
		return *m.IsTombstone
	}
	return false
}

func (m *RaftMessage) GetStart() []byte {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *RaftMessage) GetEnd() []byte {
	if m != nil {
		return m.End
	}
	return nil
}

func init() {
	proto.RegisterType((*RaftMessage)(nil), "mraft.RaftMessage")
}
func (m *RaftMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RaftMessage) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.CellID != nil {
		dAtA[i] = 0x8
		i++
		i = encodeVarintMraft(dAtA, i, uint64(*m.CellID))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintMraft(dAtA, i, uint64(m.FromPeer.Size()))
	n1, err := m.FromPeer.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintMraft(dAtA, i, uint64(m.ToPeer.Size()))
	n2, err := m.ToPeer.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x22
	i++
	i = encodeVarintMraft(dAtA, i, uint64(m.Message.Size()))
	n3, err := m.Message.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x2a
	i++
	i = encodeVarintMraft(dAtA, i, uint64(m.CellEpoch.Size()))
	n4, err := m.CellEpoch.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if m.IsTombstone != nil {
		dAtA[i] = 0x30
		i++
		if *m.IsTombstone {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i++
	}
	if m.Start != nil {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintMraft(dAtA, i, uint64(len(m.Start)))
		i += copy(dAtA[i:], m.Start)
	}
	if m.End != nil {
		dAtA[i] = 0x42
		i++
		i = encodeVarintMraft(dAtA, i, uint64(len(m.End)))
		i += copy(dAtA[i:], m.End)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeFixed64Mraft(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Mraft(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintMraft(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *RaftMessage) Size() (n int) {
	var l int
	_ = l
	if m.CellID != nil {
		n += 1 + sovMraft(uint64(*m.CellID))
	}
	l = m.FromPeer.Size()
	n += 1 + l + sovMraft(uint64(l))
	l = m.ToPeer.Size()
	n += 1 + l + sovMraft(uint64(l))
	l = m.Message.Size()
	n += 1 + l + sovMraft(uint64(l))
	l = m.CellEpoch.Size()
	n += 1 + l + sovMraft(uint64(l))
	if m.IsTombstone != nil {
		n += 2
	}
	if m.Start != nil {
		l = len(m.Start)
		n += 1 + l + sovMraft(uint64(l))
	}
	if m.End != nil {
		l = len(m.End)
		n += 1 + l + sovMraft(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovMraft(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMraft(x uint64) (n int) {
	return sovMraft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RaftMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMraft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RaftMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RaftMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellID", wireType)
			}
			var v uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.CellID = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromPeer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.FromPeer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToPeer", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ToPeer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Message", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Message.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CellEpoch", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CellEpoch.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsTombstone", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			b := bool(v != 0)
			m.IsTombstone = &b
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Start = append(m.Start[:0], dAtA[iNdEx:postIndex]...)
			if m.Start == nil {
				m.Start = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMraft
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.End = append(m.End[:0], dAtA[iNdEx:postIndex]...)
			if m.End == nil {
				m.End = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMraft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMraft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipMraft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMraft
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
					return 0, ErrIntOverflowMraft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowMraft
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthMraft
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMraft
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipMraft(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthMraft = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMraft   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("mraft.proto", fileDescriptorMraft) }

var fileDescriptorMraft = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8e, 0x4d, 0x4e, 0xc3, 0x30,
	0x10, 0x85, 0xeb, 0xfe, 0xa4, 0xc1, 0xa9, 0x04, 0x98, 0x08, 0x59, 0x5d, 0x84, 0x88, 0x55, 0x84,
	0x84, 0x2d, 0x21, 0x71, 0x81, 0x02, 0x0b, 0x16, 0x48, 0x28, 0xe2, 0x02, 0x49, 0xea, 0xa4, 0x95,
	0xea, 0x4e, 0x64, 0x9b, 0xbb, 0x70, 0x0f, 0x2e, 0xd1, 0x65, 0x4f, 0x80, 0x20, 0x5c, 0x04, 0xc5,
	0x71, 0x28, 0x1b, 0x36, 0xf6, 0xbc, 0x37, 0xdf, 0xcc, 0x3c, 0x1c, 0x48, 0x95, 0x95, 0x86, 0xd5,
	0x0a, 0x0c, 0x90, 0x89, 0x15, 0xf3, 0x33, 0x29, 0x4c, 0x56, 0xe7, 0xbc, 0xfb, 0xba, 0xde, 0x3c,
	0xac, 0xa0, 0x02, 0x5b, 0xf2, 0xb6, 0x72, 0xee, 0x75, 0xb5, 0x36, 0xab, 0xd7, 0x9c, 0x15, 0x20,
	0x79, 0x01, 0x4a, 0x80, 0xe6, 0xc2, 0x14, 0x4b, 0xde, 0xee, 0xb1, 0x4f, 0x9d, 0xf3, 0xc3, 0x81,
	0xcb, 0xf7, 0x21, 0x0e, 0xd2, 0xac, 0x34, 0x4f, 0x42, 0xeb, 0xac, 0x12, 0xe4, 0x1c, 0x7b, 0x85,
	0xd8, 0x6c, 0x1e, 0xef, 0x29, 0x8a, 0x51, 0x32, 0x4e, 0x9d, 0x22, 0x0c, 0xfb, 0xa5, 0x02, 0xf9,
	0x2c, 0x84, 0xa2, 0xc3, 0x18, 0x25, 0xc1, 0xcd, 0x8c, 0xb9, 0x34, 0xad, 0xb7, 0x18, 0xef, 0x3e,
	0x2e, 0x06, 0xe9, 0x2f, 0x43, 0xae, 0xb0, 0x67, 0xc0, 0xd2, 0xa3, 0x7f, 0x69, 0x47, 0x10, 0x8e,
	0xa7, 0xb2, 0x3b, 0x4f, 0xc7, 0x16, 0x3e, 0x66, 0x5d, 0x50, 0xe6, 0x52, 0x39, 0xbe, 0xa7, 0xc8,
	0x2d, 0x3e, 0x6a, 0x63, 0x3d, 0xd4, 0x50, 0xac, 0xe8, 0xc4, 0x8e, 0x9c, 0xf6, 0xfb, 0xef, 0xfa,
	0x86, 0x1b, 0x3a, 0x90, 0x24, 0xc6, 0xc1, 0x5a, 0xbf, 0x80, 0xcc, 0xb5, 0x81, 0xad, 0xa0, 0x5e,
	0x8c, 0x12, 0x3f, 0xfd, 0x6b, 0x91, 0x10, 0x4f, 0xb4, 0xc9, 0x94, 0xa1, 0xd3, 0x18, 0x25, 0xb3,
	0xb4, 0x13, 0xe4, 0x04, 0x8f, 0xc4, 0x76, 0x49, 0x7d, 0xeb, 0xb5, 0xe5, 0x22, 0xdc, 0x7f, 0x45,
	0x83, 0x5d, 0x13, 0xa1, 0x7d, 0x13, 0xa1, 0xcf, 0x26, 0x42, 0x6f, 0xdf, 0xd1, 0xe0, 0x27, 0x00,
	0x00, 0xff, 0xff, 0xf1, 0x0f, 0xbe, 0x9e, 0xba, 0x01, 0x00, 0x00,
}