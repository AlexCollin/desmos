// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: desmos/subspaces/v1beta1/subspace.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// SubspaceType contains all the possible subspace types
type SubspaceType int32

const (
	// SubspaceTypeUnspecified identifies an unspecified type of subspace (used in
	// errors)
	SubspaceTypeUnspecified SubspaceType = 0
	// SubspaceTypeOpen identifies that users can interact inside the subspace
	// without the need to being registered in it
	SubspaceTypeOpen SubspaceType = 1
	// SubspaceTypeClosed identifies that users can't interact inside the subspace
	// without being registered in it
	SubspaceTypeClosed SubspaceType = 2
)

var SubspaceType_name = map[int32]string{
	0: "SUBSPACE_TYPE_UNSPECIFIED",
	1: "SUBSPACE_TYPE_OPEN",
	2: "SUBSPACE_TYPE_CLOSED",
}

var SubspaceType_value = map[string]int32{
	"SUBSPACE_TYPE_UNSPECIFIED": 0,
	"SUBSPACE_TYPE_OPEN":        1,
	"SUBSPACE_TYPE_CLOSED":      2,
}

func (x SubspaceType) String() string {
	return proto.EnumName(SubspaceType_name, int32(x))
}

func (SubspaceType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e657cf67bd23372d, []int{0}
}

// Subspace contains all the data of a Desmos subspace
type Subspace struct {
	// unique SHA-256 string that identifies the subspace
	ID string `protobuf:"bytes,1,opt,name=id,proto3" json:"subspace_id" yaml:"subspace_id"`
	// human readable name of the subspace
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty" yaml:"name"`
	// the address of the user that owns the subspace
	Owner string `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty" yaml:"owner"`
	// the address of the subspace creator
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty" yaml:"creator"`
	// the creation time of the subspace
	CreationTime time.Time `protobuf:"bytes,5,opt,name=creation_time,json=creationTime,proto3,stdtime" json:"creation_time" yaml:"creation_time"`
	// the type of the subspace that indicates if it need registration or not
	Type SubspaceType `protobuf:"varint,6,opt,name=type,proto3,enum=desmos.subspaces.v1beta1.SubspaceType" json:"type" yaml:"type"`
}

func (m *Subspace) Reset()         { *m = Subspace{} }
func (m *Subspace) String() string { return proto.CompactTextString(m) }
func (*Subspace) ProtoMessage()    {}
func (*Subspace) Descriptor() ([]byte, []int) {
	return fileDescriptor_e657cf67bd23372d, []int{0}
}
func (m *Subspace) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Subspace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Subspace.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Subspace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Subspace.Merge(m, src)
}
func (m *Subspace) XXX_Size() int {
	return m.Size()
}
func (m *Subspace) XXX_DiscardUnknown() {
	xxx_messageInfo_Subspace.DiscardUnknown(m)
}

var xxx_messageInfo_Subspace proto.InternalMessageInfo

func (m *Subspace) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Subspace) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Subspace) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Subspace) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Subspace) GetCreationTime() time.Time {
	if m != nil {
		return m.CreationTime
	}
	return time.Time{}
}

func (m *Subspace) GetType() SubspaceType {
	if m != nil {
		return m.Type
	}
	return SubspaceTypeUnspecified
}

// UnregisteredPair contains a subspace-user pair used to remove all the
// relationships and blocklist created by given user in the specified subspace
// during end block
type UnregisteredPair struct {
	// the id of the unreigstered subspace
	SubspaceID string `protobuf:"bytes,1,opt,name=subspace_id,json=subspaceId,proto3" json:"subspace_id,omitempty" yaml:"subspace_id"`
	// the address of the unregistered user
	User string `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty" yaml:"user"`
}

func (m *UnregisteredPair) Reset()         { *m = UnregisteredPair{} }
func (m *UnregisteredPair) String() string { return proto.CompactTextString(m) }
func (*UnregisteredPair) ProtoMessage()    {}
func (*UnregisteredPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e657cf67bd23372d, []int{1}
}
func (m *UnregisteredPair) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnregisteredPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnregisteredPair.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnregisteredPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnregisteredPair.Merge(m, src)
}
func (m *UnregisteredPair) XXX_Size() int {
	return m.Size()
}
func (m *UnregisteredPair) XXX_DiscardUnknown() {
	xxx_messageInfo_UnregisteredPair.DiscardUnknown(m)
}

var xxx_messageInfo_UnregisteredPair proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("desmos.subspaces.v1beta1.SubspaceType", SubspaceType_name, SubspaceType_value)
	proto.RegisterType((*Subspace)(nil), "desmos.subspaces.v1beta1.Subspace")
	proto.RegisterType((*UnregisteredPair)(nil), "desmos.subspaces.v1beta1.UnregisteredPair")
}

func init() {
	proto.RegisterFile("desmos/subspaces/v1beta1/subspace.proto", fileDescriptor_e657cf67bd23372d)
}

var fileDescriptor_e657cf67bd23372d = []byte{
	// 584 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x93, 0x31, 0x6f, 0xd3, 0x4e,
	0x18, 0xc6, 0x7d, 0x69, 0xda, 0x7f, 0xff, 0xd7, 0x52, 0xac, 0x53, 0x44, 0x8d, 0x91, 0x7c, 0x96,
	0x41, 0xa5, 0x82, 0x62, 0xd3, 0x32, 0x20, 0x85, 0x89, 0x24, 0x46, 0x8a, 0x84, 0xda, 0x60, 0x37,
	0x03, 0x2c, 0x91, 0x1d, 0x5f, 0x8d, 0xa5, 0xd8, 0x67, 0xf9, 0x1c, 0xa0, 0x3b, 0x43, 0xd5, 0xa9,
	0x0b, 0x12, 0x4b, 0xa5, 0x4a, 0x7c, 0x01, 0x3e, 0x46, 0xc7, 0x8e, 0x4c, 0x06, 0xa5, 0x0b, 0xea,
	0x98, 0x4f, 0x80, 0x7c, 0x8e, 0x8b, 0x23, 0xd4, 0xed, 0xde, 0xe7, 0xfd, 0x3d, 0xf7, 0xea, 0xbd,
	0xc7, 0x86, 0x0f, 0x3d, 0xc2, 0x42, 0xca, 0x0c, 0x36, 0x76, 0x59, 0xec, 0x0c, 0x09, 0x33, 0x3e,
	0x6c, 0xbb, 0x24, 0x75, 0xb6, 0xaf, 0x15, 0x3d, 0x4e, 0x68, 0x4a, 0x91, 0x54, 0x80, 0xfa, 0x35,
	0xa8, 0xcf, 0x40, 0xb9, 0xe1, 0x53, 0x9f, 0x72, 0xc8, 0xc8, 0x4f, 0x05, 0x2f, 0x63, 0x9f, 0x52,
	0x7f, 0x44, 0x0c, 0x5e, 0xb9, 0xe3, 0x03, 0x23, 0x0d, 0x42, 0xc2, 0x52, 0x27, 0x8c, 0x0b, 0x40,
	0xfb, 0xb2, 0x00, 0x97, 0xed, 0xd9, 0x65, 0xe8, 0x05, 0xac, 0x05, 0x9e, 0x04, 0x54, 0xb0, 0xf9,
	0x7f, 0xeb, 0xf1, 0x24, 0xc3, 0xb5, 0x6e, 0xe7, 0x2a, 0xc3, 0x2b, 0xe5, 0xb0, 0x41, 0xe0, 0x4d,
	0x33, 0x8c, 0x0e, 0x9d, 0x70, 0xd4, 0xd4, 0x2a, 0xa2, 0x66, 0xd5, 0x02, 0x0f, 0xdd, 0x87, 0xf5,
	0xc8, 0x09, 0x89, 0x54, 0xe3, 0xf6, 0xdb, 0xd3, 0x0c, 0xaf, 0x14, 0x64, 0xae, 0x6a, 0x16, 0x6f,
	0xa2, 0x0d, 0xb8, 0x48, 0x3f, 0x46, 0x24, 0x91, 0x16, 0x38, 0x25, 0x4e, 0x33, 0xbc, 0x5a, 0x50,
	0x5c, 0xd6, 0xac, 0xa2, 0x8d, 0xb6, 0xe0, 0x7f, 0xc3, 0x84, 0x38, 0x29, 0x4d, 0xa4, 0x3a, 0x27,
	0xd1, 0x34, 0xc3, 0x6b, 0x05, 0x39, 0x6b, 0x68, 0x56, 0x89, 0xa0, 0x04, 0xde, 0xe2, 0xc7, 0x80,
	0x46, 0x83, 0x7c, 0x41, 0x69, 0x51, 0x05, 0x9b, 0x2b, 0x3b, 0xb2, 0x5e, 0x6c, 0xaf, 0x97, 0xdb,
	0xeb, 0xfb, 0xe5, 0xf6, 0xad, 0xed, 0xf3, 0x0c, 0x0b, 0x57, 0x19, 0x9e, 0x37, 0x4e, 0x33, 0xdc,
	0xa8, 0x0c, 0x29, 0x65, 0xed, 0xe4, 0x27, 0x06, 0xd6, 0x6a, 0xa9, 0xe5, 0xb7, 0x20, 0x1b, 0xd6,
	0xd3, 0xc3, 0x98, 0x48, 0x4b, 0x2a, 0xd8, 0x5c, 0xdb, 0xd9, 0xd0, 0x6f, 0x0a, 0x46, 0x2f, 0x5f,
	0x77, 0xff, 0x30, 0x26, 0xad, 0xf5, 0xab, 0x0c, 0x73, 0xdf, 0xdf, 0xe7, 0xc9, 0x2b, 0xcd, 0xe2,
	0x62, 0x73, 0xf9, 0xeb, 0x19, 0x06, 0xbf, 0xcf, 0x30, 0xd0, 0x3e, 0x03, 0x28, 0xf6, 0xa3, 0x84,
	0xf8, 0x01, 0x4b, 0x49, 0x42, 0xbc, 0x9e, 0x13, 0x24, 0xc8, 0x84, 0xd5, 0x2c, 0x66, 0x41, 0x3d,
	0x98, 0x64, 0x18, 0x96, 0x43, 0xba, 0x9d, 0x1b, 0x12, 0x82, 0x65, 0xd5, 0xe5, 0x49, 0x8d, 0x19,
	0x49, 0xfe, 0x4d, 0x2a, 0x57, 0x35, 0x8b, 0x37, 0x9b, 0xf5, 0xa3, 0x33, 0x2c, 0x3c, 0xfa, 0x0e,
	0xe0, 0x6a, 0x75, 0x01, 0xd4, 0x84, 0x77, 0xed, 0x7e, 0xcb, 0xee, 0xbd, 0x6c, 0x9b, 0x83, 0xfd,
	0xb7, 0x3d, 0x73, 0xd0, 0xdf, 0xb5, 0x7b, 0x66, 0xbb, 0xfb, 0xaa, 0x6b, 0x76, 0x44, 0x41, 0xbe,
	0x77, 0x7c, 0xaa, 0xae, 0x57, 0x0d, 0xfd, 0x88, 0xc5, 0x64, 0x18, 0x1c, 0x04, 0xc4, 0x43, 0x5b,
	0x10, 0xcd, 0x7b, 0xf7, 0x7a, 0xe6, 0xae, 0x08, 0xe4, 0xc6, 0xf1, 0xa9, 0x2a, 0x56, 0x4d, 0x7b,
	0x31, 0x89, 0xd0, 0x53, 0xd8, 0x98, 0xa7, 0xdb, 0xaf, 0xf7, 0x6c, 0xb3, 0x23, 0xd6, 0xe4, 0x3b,
	0xc7, 0xa7, 0x2a, 0xaa, 0xf2, 0xed, 0x11, 0x65, 0xc4, 0x93, 0xeb, 0x47, 0xdf, 0x14, 0xa1, 0xf5,
	0xe6, 0x7c, 0xa2, 0x80, 0x8b, 0x89, 0x02, 0x7e, 0x4d, 0x14, 0x70, 0x72, 0xa9, 0x08, 0x17, 0x97,
	0x8a, 0xf0, 0xe3, 0x52, 0x11, 0xde, 0x3d, 0xf7, 0x83, 0xf4, 0xfd, 0xd8, 0xd5, 0x87, 0x34, 0x34,
	0x8a, 0xb8, 0x9e, 0x8c, 0x1c, 0x97, 0xcd, 0xce, 0xc6, 0x27, 0x83, 0xa5, 0x8e, 0x1f, 0x44, 0x7e,
	0xe5, 0x37, 0xcc, 0x53, 0x61, 0xee, 0x12, 0xff, 0x80, 0x9e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff,
	0xc0, 0x55, 0xb0, 0xa8, 0xa7, 0x03, 0x00, 0x00,
}

func (this *Subspace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Subspace)
	if !ok {
		that2, ok := that.(Subspace)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	if !this.CreationTime.Equal(that1.CreationTime) {
		return false
	}
	if this.Type != that1.Type {
		return false
	}
	return true
}
func (m *Subspace) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Subspace) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Subspace) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		i = encodeVarintSubspace(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x30
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.CreationTime, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintSubspace(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x2a
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ID) > 0 {
		i -= len(m.ID)
		copy(dAtA[i:], m.ID)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.ID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *UnregisteredPair) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnregisteredPair) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnregisteredPair) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.SubspaceID) > 0 {
		i -= len(m.SubspaceID)
		copy(dAtA[i:], m.SubspaceID)
		i = encodeVarintSubspace(dAtA, i, uint64(len(m.SubspaceID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSubspace(dAtA []byte, offset int, v uint64) int {
	offset -= sovSubspace(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Subspace) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.CreationTime)
	n += 1 + l + sovSubspace(uint64(l))
	if m.Type != 0 {
		n += 1 + sovSubspace(uint64(m.Type))
	}
	return n
}

func (m *UnregisteredPair) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SubspaceID)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovSubspace(uint64(l))
	}
	return n
}

func sovSubspace(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSubspace(x uint64) (n int) {
	return sovSubspace(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Subspace) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubspace
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
			return fmt.Errorf("proto: Subspace: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Subspace: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreationTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.CreationTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= SubspaceType(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSubspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubspace
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
func (m *UnregisteredPair) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSubspace
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
			return fmt.Errorf("proto: UnregisteredPair: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnregisteredPair: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubspaceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubspaceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSubspace
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
				return ErrInvalidLengthSubspace
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSubspace
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSubspace(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSubspace
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
func skipSubspace(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSubspace
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
					return 0, ErrIntOverflowSubspace
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
					return 0, ErrIntOverflowSubspace
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
				return 0, ErrInvalidLengthSubspace
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSubspace
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSubspace
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSubspace        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSubspace          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSubspace = fmt.Errorf("proto: unexpected end of group")
)
