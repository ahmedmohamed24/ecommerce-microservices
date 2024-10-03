// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: protos/v1/admin.proto

package __

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

type GetAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *GetAdminRequest) Reset() {
	*x = GetAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminRequest) ProtoMessage() {}

func (x *GetAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminRequest.ProtoReflect.Descriptor instead.
func (*GetAdminRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{0}
}

func (x *GetAdminRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

type GetAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetAdminResponse) Reset() {
	*x = GetAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAdminResponse) ProtoMessage() {}

func (x *GetAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAdminResponse.ProtoReflect.Descriptor instead.
func (*GetAdminResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{1}
}

func (x *GetAdminResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *GetAdminResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAdminResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAdminResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetAdminResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type UpdateAdminRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password  string `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UpdateAdminRequest) Reset() {
	*x = UpdateAdminRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminRequest) ProtoMessage() {}

func (x *UpdateAdminRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminRequest.ProtoReflect.Descriptor instead.
func (*UpdateAdminRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateAdminRequest) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *UpdateAdminRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateAdminRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAdminRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdateAdminResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Email     string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt int64  `protobuf:"varint,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *UpdateAdminResponse) Reset() {
	*x = UpdateAdminResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateAdminResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateAdminResponse) ProtoMessage() {}

func (x *UpdateAdminResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateAdminResponse.ProtoReflect.Descriptor instead.
func (*UpdateAdminResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateAdminResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateAdminResponse) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateAdminResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateAdminResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *UpdateAdminResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type GenerateAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email    string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *GenerateAuthTokenRequest) Reset() {
	*x = GenerateAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthTokenRequest) ProtoMessage() {}

func (x *GenerateAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{4}
}

func (x *GenerateAuthTokenRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GenerateAuthTokenRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type GenerateAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken    string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *GenerateAuthTokenResponse) Reset() {
	*x = GenerateAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateAuthTokenResponse) ProtoMessage() {}

func (x *GenerateAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{5}
}

func (x *GenerateAuthTokenResponse) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

func (x *GenerateAuthTokenResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type RefreshAuthTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	AdminId      int64  `protobuf:"varint,2,opt,name=admin_id,json=adminId,proto3" json:"admin_id,omitempty"`
}

func (x *RefreshAuthTokenRequest) Reset() {
	*x = RefreshAuthTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAuthTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAuthTokenRequest) ProtoMessage() {}

func (x *RefreshAuthTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAuthTokenRequest.ProtoReflect.Descriptor instead.
func (*RefreshAuthTokenRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{6}
}

func (x *RefreshAuthTokenRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshAuthTokenRequest) GetAdminId() int64 {
	if x != nil {
		return x.AdminId
	}
	return 0
}

type RefreshAuthTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthToken string `protobuf:"bytes,1,opt,name=auth_token,json=authToken,proto3" json:"auth_token,omitempty"`
}

func (x *RefreshAuthTokenResponse) Reset() {
	*x = RefreshAuthTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_admin_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshAuthTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshAuthTokenResponse) ProtoMessage() {}

func (x *RefreshAuthTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_admin_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshAuthTokenResponse.ProtoReflect.Descriptor instead.
func (*RefreshAuthTokenResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_admin_proto_rawDescGZIP(), []int{7}
}

func (x *RefreshAuthTokenResponse) GetAuthToken() string {
	if x != nil {
		return x.AuthToken
	}
	return ""
}

var File_protos_v1_admin_proto protoreflect.FileDescriptor

var file_protos_v1_admin_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x22, 0x30, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x8a, 0x01, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x79, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x8d, 0x01, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x4c, 0x0a, 0x18, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22,
	0x5f, 0x0a, 0x19, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x74, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x61, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x59, 0x0a, 0x17, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x49, 0x64, 0x22, 0x39, 0x0a, 0x18, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x75, 0x74,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x8f, 0x03, 0x0a, 0x0d, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64,
	0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58, 0x0a, 0x0b, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x6a, 0x0a, 0x11, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41,
	0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x28, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x29, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e,
	0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x67, 0x0a, 0x10, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x41, 0x75, 0x74, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_admin_proto_rawDescOnce sync.Once
	file_protos_v1_admin_proto_rawDescData = file_protos_v1_admin_proto_rawDesc
)

func file_protos_v1_admin_proto_rawDescGZIP() []byte {
	file_protos_v1_admin_proto_rawDescOnce.Do(func() {
		file_protos_v1_admin_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_admin_proto_rawDescData)
	})
	return file_protos_v1_admin_proto_rawDescData
}

var file_protos_v1_admin_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_v1_admin_proto_goTypes = []any{
	(*GetAdminRequest)(nil),           // 0: proto_admin_v1.GetAdminRequest
	(*GetAdminResponse)(nil),          // 1: proto_admin_v1.GetAdminResponse
	(*UpdateAdminRequest)(nil),        // 2: proto_admin_v1.UpdateAdminRequest
	(*UpdateAdminResponse)(nil),       // 3: proto_admin_v1.UpdateAdminResponse
	(*GenerateAuthTokenRequest)(nil),  // 4: proto_admin_v1.GenerateAuthTokenRequest
	(*GenerateAuthTokenResponse)(nil), // 5: proto_admin_v1.GenerateAuthTokenResponse
	(*RefreshAuthTokenRequest)(nil),   // 6: proto_admin_v1.RefreshAuthTokenRequest
	(*RefreshAuthTokenResponse)(nil),  // 7: proto_admin_v1.RefreshAuthTokenResponse
}
var file_protos_v1_admin_proto_depIdxs = []int32{
	0, // 0: proto_admin_v1.AdminsService.GetAdmin:input_type -> proto_admin_v1.GetAdminRequest
	2, // 1: proto_admin_v1.AdminsService.UpdateAdmin:input_type -> proto_admin_v1.UpdateAdminRequest
	4, // 2: proto_admin_v1.AdminsService.GenerateAuthToken:input_type -> proto_admin_v1.GenerateAuthTokenRequest
	6, // 3: proto_admin_v1.AdminsService.RefreshAuthToken:input_type -> proto_admin_v1.RefreshAuthTokenRequest
	1, // 4: proto_admin_v1.AdminsService.GetAdmin:output_type -> proto_admin_v1.GetAdminResponse
	3, // 5: proto_admin_v1.AdminsService.UpdateAdmin:output_type -> proto_admin_v1.UpdateAdminResponse
	5, // 6: proto_admin_v1.AdminsService.GenerateAuthToken:output_type -> proto_admin_v1.GenerateAuthTokenResponse
	7, // 7: proto_admin_v1.AdminsService.RefreshAuthToken:output_type -> proto_admin_v1.RefreshAuthTokenResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_v1_admin_proto_init() }
func file_protos_v1_admin_proto_init() {
	if File_protos_v1_admin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_admin_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetAdminRequest); i {
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
		file_protos_v1_admin_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAdminResponse); i {
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
		file_protos_v1_admin_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAdminRequest); i {
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
		file_protos_v1_admin_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateAdminResponse); i {
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
		file_protos_v1_admin_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateAuthTokenRequest); i {
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
		file_protos_v1_admin_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*GenerateAuthTokenResponse); i {
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
		file_protos_v1_admin_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshAuthTokenRequest); i {
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
		file_protos_v1_admin_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*RefreshAuthTokenResponse); i {
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
			RawDescriptor: file_protos_v1_admin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_admin_proto_goTypes,
		DependencyIndexes: file_protos_v1_admin_proto_depIdxs,
		MessageInfos:      file_protos_v1_admin_proto_msgTypes,
	}.Build()
	File_protos_v1_admin_proto = out.File
	file_protos_v1_admin_proto_rawDesc = nil
	file_protos_v1_admin_proto_goTypes = nil
	file_protos_v1_admin_proto_depIdxs = nil
}
