// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: user-srv.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type RegistAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Password string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *RegistAccountReq) Reset() {
	*x = RegistAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistAccountReq) ProtoMessage() {}

func (x *RegistAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistAccountReq.ProtoReflect.Descriptor instead.
func (*RegistAccountReq) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{0}
}

func (x *RegistAccountReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *RegistAccountReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *RegistAccountReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegistAccountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegistAccountRsp) Reset() {
	*x = RegistAccountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistAccountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistAccountRsp) ProtoMessage() {}

func (x *RegistAccountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistAccountRsp.ProtoReflect.Descriptor instead.
func (*RegistAccountRsp) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{1}
}

type LoginAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginAccountReq) Reset() {
	*x = LoginAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountReq) ProtoMessage() {}

func (x *LoginAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountReq.ProtoReflect.Descriptor instead.
func (*LoginAccountReq) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{2}
}

func (x *LoginAccountReq) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginAccountReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type LoginAccountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID   int64  `protobuf:"varint,1,opt,name=userID,proto3" json:"userID,omitempty"`
	UserName string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	Email    string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Phone    string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *LoginAccountRsp) Reset() {
	*x = LoginAccountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginAccountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginAccountRsp) ProtoMessage() {}

func (x *LoginAccountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginAccountRsp.ProtoReflect.Descriptor instead.
func (*LoginAccountRsp) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{3}
}

func (x *LoginAccountRsp) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

func (x *LoginAccountRsp) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *LoginAccountRsp) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *LoginAccountRsp) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type ResetAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetAccountReq) Reset() {
	*x = ResetAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetAccountReq) ProtoMessage() {}

func (x *ResetAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetAccountReq.ProtoReflect.Descriptor instead.
func (*ResetAccountReq) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{4}
}

type ResetAccountRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ResetAccountRsp) Reset() {
	*x = ResetAccountRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetAccountRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetAccountRsp) ProtoMessage() {}

func (x *ResetAccountRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetAccountRsp.ProtoReflect.Descriptor instead.
func (*ResetAccountRsp) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{5}
}

type WantScoreReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	MovieId int64 `protobuf:"varint,2,opt,name=movieId,proto3" json:"movieId,omitempty"` // 订单编号
	Score   int64 `protobuf:"varint,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *WantScoreReq) Reset() {
	*x = WantScoreReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WantScoreReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WantScoreReq) ProtoMessage() {}

func (x *WantScoreReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WantScoreReq.ProtoReflect.Descriptor instead.
func (*WantScoreReq) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{6}
}

func (x *WantScoreReq) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *WantScoreReq) GetMovieId() int64 {
	if x != nil {
		return x.MovieId
	}
	return 0
}

func (x *WantScoreReq) GetScore() int64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type WantScoreRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WantScoreRsp) Reset() {
	*x = WantScoreRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WantScoreRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WantScoreRsp) ProtoMessage() {}

func (x *WantScoreRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WantScoreRsp.ProtoReflect.Descriptor instead.
func (*WantScoreRsp) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{7}
}

type UpdateUserProfileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserImage string `protobuf:"bytes,1,opt,name=userImage,proto3" json:"userImage,omitempty"`
	UserName  string `protobuf:"bytes,2,opt,name=userName,proto3" json:"userName,omitempty"`
	UserEmail string `protobuf:"bytes,3,opt,name=userEmail,proto3" json:"userEmail,omitempty"`
	UserPhone string `protobuf:"bytes,4,opt,name=userPhone,proto3" json:"userPhone,omitempty"`
	UserID    int64  `protobuf:"varint,5,opt,name=userID,proto3" json:"userID,omitempty"`
}

func (x *UpdateUserProfileReq) Reset() {
	*x = UpdateUserProfileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserProfileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserProfileReq) ProtoMessage() {}

func (x *UpdateUserProfileReq) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserProfileReq.ProtoReflect.Descriptor instead.
func (*UpdateUserProfileReq) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateUserProfileReq) GetUserImage() string {
	if x != nil {
		return x.UserImage
	}
	return ""
}

func (x *UpdateUserProfileReq) GetUserName() string {
	if x != nil {
		return x.UserName
	}
	return ""
}

func (x *UpdateUserProfileReq) GetUserEmail() string {
	if x != nil {
		return x.UserEmail
	}
	return ""
}

func (x *UpdateUserProfileReq) GetUserPhone() string {
	if x != nil {
		return x.UserPhone
	}
	return ""
}

func (x *UpdateUserProfileReq) GetUserID() int64 {
	if x != nil {
		return x.UserID
	}
	return 0
}

type UpdateUserProfileRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateUserProfileRsp) Reset() {
	*x = UpdateUserProfileRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_srv_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUserProfileRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserProfileRsp) ProtoMessage() {}

func (x *UpdateUserProfileRsp) ProtoReflect() protoreflect.Message {
	mi := &file_user_srv_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserProfileRsp.ProtoReflect.Descriptor instead.
func (*UpdateUserProfileRsp) Descriptor() ([]byte, []int) {
	return file_user_srv_proto_rawDescGZIP(), []int{9}
}

var File_user_srv_proto protoreflect.FileDescriptor

var file_user_srv_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x72, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x60, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x12, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x43, 0x0a, 0x0f,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x71, 0x0a, 0x0f, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x73, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14,
	0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x22, 0x11, 0x0a, 0x0f, 0x52, 0x65, 0x73, 0x65, 0x74,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x56, 0x0a, 0x0c, 0x57, 0x61,
	0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x6d, 0x6f, 0x76, 0x69, 0x65, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x63, 0x6f,
	0x72, 0x65, 0x22, 0x0e, 0x0a, 0x0c, 0x57, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x73, 0x70, 0x22, 0xa4, 0x01, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x09, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x73, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0x16, 0x0a, 0x14, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x73,
	0x70, 0x32, 0xcf, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x0a, 0x0d, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0c, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0c, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x15, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x35, 0x0a,
	0x09, 0x57, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x57, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x57, 0x61, 0x6e, 0x74, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x52,
	0x73, 0x70, 0x22, 0x00, 0x12, 0x4d, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x52, 0x73,
	0x70, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_srv_proto_rawDescOnce sync.Once
	file_user_srv_proto_rawDescData = file_user_srv_proto_rawDesc
)

func file_user_srv_proto_rawDescGZIP() []byte {
	file_user_srv_proto_rawDescOnce.Do(func() {
		file_user_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_srv_proto_rawDescData)
	})
	return file_user_srv_proto_rawDescData
}

var file_user_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_user_srv_proto_goTypes = []interface{}{
	(*RegistAccountReq)(nil),     // 0: user.RegistAccountReq
	(*RegistAccountRsp)(nil),     // 1: user.RegistAccountRsp
	(*LoginAccountReq)(nil),      // 2: user.LoginAccountReq
	(*LoginAccountRsp)(nil),      // 3: user.LoginAccountRsp
	(*ResetAccountReq)(nil),      // 4: user.ResetAccountReq
	(*ResetAccountRsp)(nil),      // 5: user.ResetAccountRsp
	(*WantScoreReq)(nil),         // 6: user.WantScoreReq
	(*WantScoreRsp)(nil),         // 7: user.WantScoreRsp
	(*UpdateUserProfileReq)(nil), // 8: user.UpdateUserProfileReq
	(*UpdateUserProfileRsp)(nil), // 9: user.UpdateUserProfileRsp
}
var file_user_srv_proto_depIdxs = []int32{
	0, // 0: user.User.RegistAccount:input_type -> user.RegistAccountReq
	2, // 1: user.User.LoginAccount:input_type -> user.LoginAccountReq
	4, // 2: user.User.ResetAccount:input_type -> user.ResetAccountReq
	6, // 3: user.User.WantScore:input_type -> user.WantScoreReq
	8, // 4: user.User.UpdateUserProfile:input_type -> user.UpdateUserProfileReq
	1, // 5: user.User.RegistAccount:output_type -> user.RegistAccountRsp
	3, // 6: user.User.LoginAccount:output_type -> user.LoginAccountRsp
	5, // 7: user.User.ResetAccount:output_type -> user.ResetAccountRsp
	7, // 8: user.User.WantScore:output_type -> user.WantScoreRsp
	9, // 9: user.User.UpdateUserProfile:output_type -> user.UpdateUserProfileRsp
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_srv_proto_init() }
func file_user_srv_proto_init() {
	if File_user_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistAccountReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistAccountRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginAccountRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetAccountReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetAccountRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WantScoreReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WantScoreRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserProfileReq); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_user_srv_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUserProfileRsp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_user_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_srv_proto_goTypes,
		DependencyIndexes: file_user_srv_proto_depIdxs,
		MessageInfos:      file_user_srv_proto_msgTypes,
	}.Build()
	File_user_srv_proto = out.File
	file_user_srv_proto_rawDesc = nil
	file_user_srv_proto_goTypes = nil
	file_user_srv_proto_depIdxs = nil
}