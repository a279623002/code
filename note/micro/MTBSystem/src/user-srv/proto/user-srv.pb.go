// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/user-srv.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address              string   `protobuf:"bytes,3,opt,name=Address,proto3" json:"Address,omitempty"`
	Phone                string   `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *User) GetPhone() string {
	if m != nil {
		return m.Phone
	}
	return ""
}

type UpdateUserReq struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserReq) Reset()         { *m = UpdateUserReq{} }
func (m *UpdateUserReq) String() string { return proto.CompactTextString(m) }
func (*UpdateUserReq) ProtoMessage()    {}
func (*UpdateUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{1}
}

func (m *UpdateUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserReq.Unmarshal(m, b)
}
func (m *UpdateUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserReq.Marshal(b, m, deterministic)
}
func (m *UpdateUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserReq.Merge(m, src)
}
func (m *UpdateUserReq) XXX_Size() int {
	return xxx_messageInfo_UpdateUserReq.Size(m)
}
func (m *UpdateUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserReq proto.InternalMessageInfo

func (m *UpdateUserReq) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type UpdateUserRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateUserRep) Reset()         { *m = UpdateUserRep{} }
func (m *UpdateUserRep) String() string { return proto.CompactTextString(m) }
func (*UpdateUserRep) ProtoMessage()    {}
func (*UpdateUserRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{2}
}

func (m *UpdateUserRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateUserRep.Unmarshal(m, b)
}
func (m *UpdateUserRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateUserRep.Marshal(b, m, deterministic)
}
func (m *UpdateUserRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateUserRep.Merge(m, src)
}
func (m *UpdateUserRep) XXX_Size() int {
	return xxx_messageInfo_UpdateUserRep.Size(m)
}
func (m *UpdateUserRep) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateUserRep.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateUserRep proto.InternalMessageInfo

type SelectUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectUserReq) Reset()         { *m = SelectUserReq{} }
func (m *SelectUserReq) String() string { return proto.CompactTextString(m) }
func (*SelectUserReq) ProtoMessage()    {}
func (*SelectUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{3}
}

func (m *SelectUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectUserReq.Unmarshal(m, b)
}
func (m *SelectUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectUserReq.Marshal(b, m, deterministic)
}
func (m *SelectUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectUserReq.Merge(m, src)
}
func (m *SelectUserReq) XXX_Size() int {
	return xxx_messageInfo_SelectUserReq.Size(m)
}
func (m *SelectUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_SelectUserReq proto.InternalMessageInfo

func (m *SelectUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type SelectUserRep struct {
	Users                *User    `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SelectUserRep) Reset()         { *m = SelectUserRep{} }
func (m *SelectUserRep) String() string { return proto.CompactTextString(m) }
func (*SelectUserRep) ProtoMessage()    {}
func (*SelectUserRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{4}
}

func (m *SelectUserRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SelectUserRep.Unmarshal(m, b)
}
func (m *SelectUserRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SelectUserRep.Marshal(b, m, deterministic)
}
func (m *SelectUserRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SelectUserRep.Merge(m, src)
}
func (m *SelectUserRep) XXX_Size() int {
	return xxx_messageInfo_SelectUserRep.Size(m)
}
func (m *SelectUserRep) XXX_DiscardUnknown() {
	xxx_messageInfo_SelectUserRep.DiscardUnknown(m)
}

var xxx_messageInfo_SelectUserRep proto.InternalMessageInfo

func (m *SelectUserRep) GetUsers() *User {
	if m != nil {
		return m.Users
	}
	return nil
}

type DeletetUserReq struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletetUserReq) Reset()         { *m = DeletetUserReq{} }
func (m *DeletetUserReq) String() string { return proto.CompactTextString(m) }
func (*DeletetUserReq) ProtoMessage()    {}
func (*DeletetUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{5}
}

func (m *DeletetUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletetUserReq.Unmarshal(m, b)
}
func (m *DeletetUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletetUserReq.Marshal(b, m, deterministic)
}
func (m *DeletetUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletetUserReq.Merge(m, src)
}
func (m *DeletetUserReq) XXX_Size() int {
	return xxx_messageInfo_DeletetUserReq.Size(m)
}
func (m *DeletetUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletetUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_DeletetUserReq proto.InternalMessageInfo

func (m *DeletetUserReq) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type DeletetUserRep struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeletetUserRep) Reset()         { *m = DeletetUserRep{} }
func (m *DeletetUserRep) String() string { return proto.CompactTextString(m) }
func (*DeletetUserRep) ProtoMessage()    {}
func (*DeletetUserRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{6}
}

func (m *DeletetUserRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeletetUserRep.Unmarshal(m, b)
}
func (m *DeletetUserRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeletetUserRep.Marshal(b, m, deterministic)
}
func (m *DeletetUserRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeletetUserRep.Merge(m, src)
}
func (m *DeletetUserRep) XXX_Size() int {
	return xxx_messageInfo_DeletetUserRep.Size(m)
}
func (m *DeletetUserRep) XXX_DiscardUnknown() {
	xxx_messageInfo_DeletetUserRep.DiscardUnknown(m)
}

var xxx_messageInfo_DeletetUserRep proto.InternalMessageInfo

type InsertUserReq struct {
	Users                *User    `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertUserReq) Reset()         { *m = InsertUserReq{} }
func (m *InsertUserReq) String() string { return proto.CompactTextString(m) }
func (*InsertUserReq) ProtoMessage()    {}
func (*InsertUserReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{7}
}

func (m *InsertUserReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserReq.Unmarshal(m, b)
}
func (m *InsertUserReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserReq.Marshal(b, m, deterministic)
}
func (m *InsertUserReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserReq.Merge(m, src)
}
func (m *InsertUserReq) XXX_Size() int {
	return xxx_messageInfo_InsertUserReq.Size(m)
}
func (m *InsertUserReq) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserReq.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserReq proto.InternalMessageInfo

func (m *InsertUserReq) GetUsers() *User {
	if m != nil {
		return m.Users
	}
	return nil
}

type InsertUserRep struct {
	Users                *User    `protobuf:"bytes,1,opt,name=users,proto3" json:"users,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InsertUserRep) Reset()         { *m = InsertUserRep{} }
func (m *InsertUserRep) String() string { return proto.CompactTextString(m) }
func (*InsertUserRep) ProtoMessage()    {}
func (*InsertUserRep) Descriptor() ([]byte, []int) {
	return fileDescriptor_739aee4c9d5a2387, []int{8}
}

func (m *InsertUserRep) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InsertUserRep.Unmarshal(m, b)
}
func (m *InsertUserRep) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InsertUserRep.Marshal(b, m, deterministic)
}
func (m *InsertUserRep) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InsertUserRep.Merge(m, src)
}
func (m *InsertUserRep) XXX_Size() int {
	return xxx_messageInfo_InsertUserRep.Size(m)
}
func (m *InsertUserRep) XXX_DiscardUnknown() {
	xxx_messageInfo_InsertUserRep.DiscardUnknown(m)
}

var xxx_messageInfo_InsertUserRep proto.InternalMessageInfo

func (m *InsertUserRep) GetUsers() *User {
	if m != nil {
		return m.Users
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "pb.User")
	proto.RegisterType((*UpdateUserReq)(nil), "pb.UpdateUserReq")
	proto.RegisterType((*UpdateUserRep)(nil), "pb.UpdateUserRep")
	proto.RegisterType((*SelectUserReq)(nil), "pb.SelectUserReq")
	proto.RegisterType((*SelectUserRep)(nil), "pb.SelectUserRep")
	proto.RegisterType((*DeletetUserReq)(nil), "pb.DeletetUserReq")
	proto.RegisterType((*DeletetUserRep)(nil), "pb.DeletetUserRep")
	proto.RegisterType((*InsertUserReq)(nil), "pb.InsertUserReq")
	proto.RegisterType((*InsertUserRep)(nil), "pb.InsertUserRep")
}

func init() { proto.RegisterFile("proto/user-srv.proto", fileDescriptor_739aee4c9d5a2387) }

var fileDescriptor_739aee4c9d5a2387 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x4f, 0x4b, 0xc3, 0x40,
	0x10, 0xc5, 0x9b, 0x98, 0xfa, 0x67, 0x42, 0xab, 0x0e, 0x3d, 0x2c, 0x41, 0x34, 0xec, 0xa9, 0x97,
	0xa6, 0x50, 0x0b, 0x9e, 0x05, 0x2f, 0xde, 0x24, 0xa5, 0x17, 0x6f, 0x4d, 0x77, 0xc0, 0x40, 0x4d,
	0xd6, 0xdd, 0xd8, 0xcf, 0xec, 0xc7, 0x90, 0xdd, 0x58, 0xe3, 0x66, 0x91, 0xde, 0x32, 0xbf, 0xbc,
	0x37, 0xef, 0x65, 0x37, 0x30, 0x91, 0xaa, 0x6e, 0xea, 0xf9, 0xa7, 0x26, 0x35, 0xd3, 0x6a, 0x9f,
	0xd9, 0x11, 0x43, 0x59, 0xf0, 0x57, 0x88, 0xd6, 0x9a, 0x14, 0x8e, 0x21, 0x2c, 0x05, 0x0b, 0xd2,
	0x60, 0x3a, 0xcc, 0xc3, 0x52, 0x20, 0x42, 0x54, 0x6d, 0xde, 0x89, 0x85, 0x69, 0x30, 0xbd, 0xc8,
	0xed, 0x33, 0x32, 0x38, 0x7b, 0x14, 0x42, 0x91, 0xd6, 0xec, 0xc4, 0xe2, 0xc3, 0x88, 0x13, 0x18,
	0xbe, 0xbc, 0xd5, 0x15, 0xb1, 0xc8, 0xf2, 0x76, 0xe0, 0x33, 0x18, 0xad, 0xa5, 0xd8, 0x34, 0x64,
	0x12, 0x72, 0xfa, 0xc0, 0x1b, 0x88, 0x4c, 0x05, 0x1b, 0x13, 0x2f, 0xce, 0x33, 0x59, 0x64, 0xf6,
	0x95, 0xa5, 0xfc, 0xd2, 0x95, 0x4b, 0x7e, 0x07, 0xa3, 0x15, 0xed, 0x68, 0xdb, 0x1c, 0xfc, 0xbd,
	0x92, 0x7c, 0xee, 0x0a, 0x24, 0xde, 0xc2, 0xd0, 0xac, 0xd2, 0x5e, 0x42, 0x8b, 0x79, 0x0a, 0xe3,
	0x27, 0xda, 0x51, 0x43, 0xff, 0xae, 0xbc, 0xea, 0x29, 0xa4, 0x09, 0x79, 0xae, 0x34, 0xa9, 0x5f,
	0xcb, 0xb1, 0x90, 0x9e, 0xe1, 0x68, 0xab, 0xc5, 0x57, 0x00, 0xb1, 0x99, 0x57, 0xa4, 0xf6, 0xe5,
	0x96, 0x70, 0x09, 0xd0, 0x2d, 0xc0, 0x6b, 0x23, 0x77, 0x1a, 0x24, 0x1e, 0x92, 0x7c, 0x80, 0x0f,
	0x10, 0xff, 0x69, 0x8e, 0x68, 0x34, 0xee, 0xc7, 0x26, 0x3e, 0x33, 0xc6, 0x25, 0x40, 0x77, 0x8a,
	0x6d, 0x9c, 0x73, 0xec, 0x89, 0x87, 0x7e, 0x5c, 0xdd, 0x6d, 0xb5, 0x2e, 0xe7, 0xb2, 0x13, 0x0f,
	0x49, 0x3e, 0x28, 0x4e, 0xed, 0x9f, 0x77, 0xff, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x1f, 0x19,
	0x4b, 0x91, 0x02, 0x00, 0x00,
}