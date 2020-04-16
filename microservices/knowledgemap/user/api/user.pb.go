// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/user.proto

/*
	Package api is a generated protocol buffer package.

	It is generated from these files:
		api/user.proto

	It has these top-level messages:
		UserReq
		Empty
		UserReply
		ClassReq
		ClassReply
*/
package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UserReq struct {
	Userid string `protobuf:"bytes,1,opt,name=userid,proto3" json:"uid"`
}

func (m *UserReq) Reset()                    { *m = UserReq{} }
func (m *UserReq) String() string            { return proto.CompactTextString(m) }
func (*UserReq) ProtoMessage()               {}
func (*UserReq) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{0} }

func (m *UserReq) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

type Empty struct {
}

func (m *Empty) Reset()                    { *m = Empty{} }
func (m *Empty) String() string            { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()               {}
func (*Empty) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{1} }

type UserReply struct {
	Userid       string   `protobuf:"bytes,1,opt,name=userid,proto3" json:"userid,omitempty" bson:"_id",form:"_id"`
	Username     string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty" bson:"name"`
	Usertype     int64    `protobuf:"varint,3,opt,name=usertype,proto3" json:"usertype,omitempty"`
	Major        string   `protobuf:"bytes,4,opt,name=major,proto3" json:"major,omitempty"`
	Idcard       string   `protobuf:"bytes,5,opt,name=idcard,proto3" json:"idcard,omitempty"`
	Admissontime string   `protobuf:"bytes,6,opt,name=admissontime,proto3" json:"admissontime,omitempty"`
	Origin       string   `protobuf:"bytes,7,opt,name=origin,proto3" json:"origin,omitempty"`
	Courses      []string `protobuf:"bytes,8,rep,name=courses" json:"courses,omitempty"`
	Class        string   `protobuf:"bytes,9,opt,name=class,proto3" json:"class,omitempty"`
	Password     string   `protobuf:"bytes,10,opt,name=password,proto3" json:"password,omitempty"`
	Number       string   `protobuf:"bytes,11,opt,name=number,proto3" json:"number,omitempty"`
}

func (m *UserReply) Reset()                    { *m = UserReply{} }
func (m *UserReply) String() string            { return proto.CompactTextString(m) }
func (*UserReply) ProtoMessage()               {}
func (*UserReply) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{2} }

func (m *UserReply) GetUserid() string {
	if m != nil {
		return m.Userid
	}
	return ""
}

func (m *UserReply) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *UserReply) GetUsertype() int64 {
	if m != nil {
		return m.Usertype
	}
	return 0
}

func (m *UserReply) GetMajor() string {
	if m != nil {
		return m.Major
	}
	return ""
}

func (m *UserReply) GetIdcard() string {
	if m != nil {
		return m.Idcard
	}
	return ""
}

func (m *UserReply) GetAdmissontime() string {
	if m != nil {
		return m.Admissontime
	}
	return ""
}

func (m *UserReply) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *UserReply) GetCourses() []string {
	if m != nil {
		return m.Courses
	}
	return nil
}

func (m *UserReply) GetClass() string {
	if m != nil {
		return m.Class
	}
	return ""
}

func (m *UserReply) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserReply) GetNumber() string {
	if m != nil {
		return m.Number
	}
	return ""
}

type ClassReq struct {
	Classid string `protobuf:"bytes,1,opt,name=classid,proto3" json:"cid"`
}

func (m *ClassReq) Reset()                    { *m = ClassReq{} }
func (m *ClassReq) String() string            { return proto.CompactTextString(m) }
func (*ClassReq) ProtoMessage()               {}
func (*ClassReq) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{3} }

func (m *ClassReq) GetClassid() string {
	if m != nil {
		return m.Classid
	}
	return ""
}

type ClassReply struct {
	Students []*UserReply `protobuf:"bytes,1,rep,name=students" json:"students,omitempty"`
}

func (m *ClassReply) Reset()                    { *m = ClassReply{} }
func (m *ClassReply) String() string            { return proto.CompactTextString(m) }
func (*ClassReply) ProtoMessage()               {}
func (*ClassReply) Descriptor() ([]byte, []int) { return fileDescriptorUser, []int{4} }

func (m *ClassReply) GetStudents() []*UserReply {
	if m != nil {
		return m.Students
	}
	return nil
}

func init() {
	proto.RegisterType((*UserReq)(nil), "api.UserReq")
	proto.RegisterType((*Empty)(nil), "api.Empty")
	proto.RegisterType((*UserReply)(nil), "api.UserReply")
	proto.RegisterType((*ClassReq)(nil), "api.ClassReq")
	proto.RegisterType((*ClassReply)(nil), "api.ClassReply")
}
func (m *UserReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Userid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Userid)))
		i += copy(dAtA[i:], m.Userid)
	}
	return i, nil
}

func (m *Empty) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Empty) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *UserReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UserReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Userid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Userid)))
		i += copy(dAtA[i:], m.Userid)
	}
	if len(m.Username) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Username)))
		i += copy(dAtA[i:], m.Username)
	}
	if m.Usertype != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintUser(dAtA, i, uint64(m.Usertype))
	}
	if len(m.Major) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Major)))
		i += copy(dAtA[i:], m.Major)
	}
	if len(m.Idcard) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Idcard)))
		i += copy(dAtA[i:], m.Idcard)
	}
	if len(m.Admissontime) > 0 {
		dAtA[i] = 0x32
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Admissontime)))
		i += copy(dAtA[i:], m.Admissontime)
	}
	if len(m.Origin) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Origin)))
		i += copy(dAtA[i:], m.Origin)
	}
	if len(m.Courses) > 0 {
		for _, s := range m.Courses {
			dAtA[i] = 0x42
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.Class) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Class)))
		i += copy(dAtA[i:], m.Class)
	}
	if len(m.Password) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Password)))
		i += copy(dAtA[i:], m.Password)
	}
	if len(m.Number) > 0 {
		dAtA[i] = 0x5a
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Number)))
		i += copy(dAtA[i:], m.Number)
	}
	return i, nil
}

func (m *ClassReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Classid) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintUser(dAtA, i, uint64(len(m.Classid)))
		i += copy(dAtA[i:], m.Classid)
	}
	return i, nil
}

func (m *ClassReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClassReply) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Students) > 0 {
		for _, msg := range m.Students {
			dAtA[i] = 0xa
			i++
			i = encodeVarintUser(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintUser(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *UserReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *Empty) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *UserReply) Size() (n int) {
	var l int
	_ = l
	l = len(m.Userid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Username)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if m.Usertype != 0 {
		n += 1 + sovUser(uint64(m.Usertype))
	}
	l = len(m.Major)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Idcard)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Admissontime)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Origin)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	if len(m.Courses) > 0 {
		for _, s := range m.Courses {
			l = len(s)
			n += 1 + l + sovUser(uint64(l))
		}
	}
	l = len(m.Class)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Password)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	l = len(m.Number)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *ClassReq) Size() (n int) {
	var l int
	_ = l
	l = len(m.Classid)
	if l > 0 {
		n += 1 + l + sovUser(uint64(l))
	}
	return n
}

func (m *ClassReply) Size() (n int) {
	var l int
	_ = l
	if len(m.Students) > 0 {
		for _, e := range m.Students {
			l = e.Size()
			n += 1 + l + sovUser(uint64(l))
		}
	}
	return n
}

func sovUser(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozUser(x uint64) (n int) {
	return sovUser(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UserReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *Empty) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: Empty: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Empty: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *UserReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: UserReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UserReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Userid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Userid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Username", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Username = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Usertype", wireType)
			}
			m.Usertype = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Usertype |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Major", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Major = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Idcard", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Idcard = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admissontime", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admissontime = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Origin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Origin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Courses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Courses = append(m.Courses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Class", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Class = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Password", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Password = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Number", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Number = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *ClassReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: ClassReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Classid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Classid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func (m *ClassReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUser
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
			return fmt.Errorf("proto: ClassReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClassReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Students", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUser
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
				return ErrInvalidLengthUser
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Students = append(m.Students, &UserReply{})
			if err := m.Students[len(m.Students)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipUser(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthUser
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
func skipUser(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
					return 0, ErrIntOverflowUser
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
				return 0, ErrInvalidLengthUser
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowUser
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
				next, err := skipUser(dAtA[start:])
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
	ErrInvalidLengthUser = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUser   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/user.proto", fileDescriptorUser) }

var fileDescriptorUser = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xbd, 0x6e, 0xdb, 0x30,
	0x10, 0x80, 0xa3, 0x28, 0xb6, 0xec, 0x73, 0x7e, 0x0a, 0xa2, 0x2d, 0x58, 0x0f, 0xb6, 0xcb, 0xc9,
	0x48, 0x1b, 0x19, 0x49, 0x97, 0xa2, 0x63, 0x8a, 0x0e, 0x5d, 0x09, 0x74, 0x2e, 0x28, 0x91, 0x56,
	0x59, 0x84, 0x22, 0x43, 0x4a, 0x28, 0xfc, 0x26, 0x7d, 0xa4, 0x8e, 0x5d, 0xba, 0x1a, 0x85, 0xbb,
	0x75, 0xcc, 0x13, 0x04, 0x24, 0x2d, 0x21, 0xc9, 0x24, 0x7e, 0x77, 0xf7, 0xdd, 0x1d, 0x70, 0x82,
	0x53, 0x66, 0xe4, 0xaa, 0x75, 0xc2, 0xe6, 0xc6, 0xea, 0x46, 0xa3, 0x94, 0x19, 0x39, 0xbd, 0xa8,
	0x64, 0xf3, 0xad, 0x2d, 0xf2, 0x52, 0xab, 0x55, 0xa5, 0x2b, 0xbd, 0x0a, 0xb9, 0xa2, 0x5d, 0x07,
	0x0a, 0x10, 0x5e, 0xd1, 0x21, 0xe7, 0x90, 0x7d, 0x71, 0xc2, 0x52, 0x71, 0x8b, 0xe6, 0x30, 0xf4,
	0xcd, 0x24, 0xc7, 0xc9, 0x22, 0x59, 0x8e, 0xaf, 0xb3, 0xff, 0xdb, 0x79, 0xda, 0x4a, 0x4e, 0xf7,
	0x61, 0x92, 0xc1, 0xe0, 0x93, 0x32, 0xcd, 0x86, 0xfc, 0x39, 0x84, 0x71, 0xb4, 0xcc, 0xcd, 0x06,
	0x5d, 0x3e, 0xf1, 0x5e, 0xdd, 0x6d, 0xe7, 0x2f, 0x0a, 0xa7, 0xeb, 0x0f, 0xe4, 0xab, 0xe4, 0xe4,
	0xed, 0x5a, 0x5b, 0x15, 0x9f, 0x5d, 0x27, 0xf4, 0x06, 0x46, 0xfe, 0x55, 0x33, 0x25, 0xf0, 0x61,
	0x90, 0xce, 0xee, 0xb6, 0xf3, 0x49, 0x94, 0x7c, 0x94, 0xd0, 0xbe, 0x00, 0x4d, 0x63, 0x71, 0xb3,
	0x31, 0x02, 0xa7, 0x8b, 0x64, 0x99, 0xd2, 0x9e, 0xd1, 0x73, 0x18, 0x28, 0xf6, 0x5d, 0x5b, 0x7c,
	0xe4, 0xbb, 0xd0, 0x08, 0xe8, 0x25, 0x0c, 0x25, 0x2f, 0x99, 0xe5, 0x78, 0x10, 0xc2, 0x7b, 0x42,
	0x04, 0x8e, 0x19, 0x57, 0xd2, 0x39, 0x5d, 0x37, 0x52, 0x09, 0x3c, 0x0c, 0xd9, 0x47, 0x31, 0xef,
	0x6a, 0x2b, 0x2b, 0x59, 0xe3, 0x2c, 0xba, 0x91, 0x10, 0x86, 0xac, 0xd4, 0xad, 0x75, 0xc2, 0xe1,
	0xd1, 0x22, 0x5d, 0x8e, 0x69, 0x87, 0x7e, 0x87, 0xf2, 0x86, 0x39, 0x87, 0xc7, 0x71, 0x87, 0x00,
	0x7e, 0x6b, 0xc3, 0x9c, 0xfb, 0xa1, 0x2d, 0xc7, 0x10, 0x12, 0x3d, 0xfb, 0x19, 0x75, 0xab, 0x0a,
	0x61, 0xf1, 0x24, 0xce, 0x88, 0x44, 0x2e, 0x60, 0xf4, 0xd1, 0xcb, 0xfe, 0x1a, 0xaf, 0x21, 0x0b,
	0x8d, 0x1e, 0x9f, 0xa3, 0x94, 0x9c, 0x76, 0x71, 0xf2, 0x1e, 0x60, 0x5f, 0xee, 0xcf, 0x70, 0x0e,
	0x23, 0xd7, 0xb4, 0x5c, 0xd4, 0x8d, 0xc3, 0xc9, 0x22, 0x5d, 0x4e, 0xae, 0x4e, 0x73, 0x66, 0x64,
	0xde, 0x1f, 0x8a, 0xf6, 0xf9, 0x2b, 0x01, 0x47, 0x3e, 0xec, 0x1d, 0xff, 0xfd, 0x5c, 0xaf, 0x35,
	0x3a, 0x7e, 0x50, 0x7d, 0x3b, 0x7d, 0xe2, 0x92, 0x03, 0x74, 0x09, 0x27, 0x1e, 0xc3, 0xc4, 0x20,
	0x9c, 0x84, 0x92, 0x6e, 0xe1, 0xe9, 0xd9, 0x43, 0x0c, 0xca, 0xf5, 0xb3, 0x5f, 0xbb, 0x59, 0xf2,
	0x7b, 0x37, 0x4b, 0xfe, 0xee, 0x66, 0xc9, 0xcf, 0x7f, 0xb3, 0x83, 0x62, 0x18, 0xfe, 0xba, 0x77,
	0xf7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd5, 0x32, 0x65, 0x75, 0xbb, 0x02, 0x00, 0x00,
}