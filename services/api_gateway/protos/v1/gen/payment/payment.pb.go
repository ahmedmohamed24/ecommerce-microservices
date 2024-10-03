// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: protos/v1/payment.proto

package __

import (
	httpbody "google.golang.org/genproto/googleapis/api/httpbody"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Quantity    uint32  `protobuf:"varint,1,opt,name=quantity,proto3" json:"quantity,omitempty"`
	Amount      float32 `protobuf:"fixed32,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Name        string  `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{0}
}

func (x *Product) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *Product) GetAmount() float32 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type GeneratePaymentLinkRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId  string     `protobuf:"bytes,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Products []*Product `protobuf:"bytes,2,rep,name=products,proto3" json:"products,omitempty"`
}

func (x *GeneratePaymentLinkRequest) Reset() {
	*x = GeneratePaymentLinkRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePaymentLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePaymentLinkRequest) ProtoMessage() {}

func (x *GeneratePaymentLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePaymentLinkRequest.ProtoReflect.Descriptor instead.
func (*GeneratePaymentLinkRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{1}
}

func (x *GeneratePaymentLinkRequest) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *GeneratePaymentLinkRequest) GetProducts() []*Product {
	if x != nil {
		return x.Products
	}
	return nil
}

type GeneratePaymentLinkResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PaymentLink string `protobuf:"bytes,1,opt,name=payment_link,json=paymentLink,proto3" json:"payment_link,omitempty"`
}

func (x *GeneratePaymentLinkResponse) Reset() {
	*x = GeneratePaymentLinkResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeneratePaymentLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratePaymentLinkResponse) ProtoMessage() {}

func (x *GeneratePaymentLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratePaymentLinkResponse.ProtoReflect.Descriptor instead.
func (*GeneratePaymentLinkResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{2}
}

func (x *GeneratePaymentLinkResponse) GetPaymentLink() string {
	if x != nil {
		return x.PaymentLink
	}
	return ""
}

type GetPaymentOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint64 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *GetPaymentOperationRequest) Reset() {
	*x = GetPaymentOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentOperationRequest) ProtoMessage() {}

func (x *GetPaymentOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentOperationRequest.ProtoReflect.Descriptor instead.
func (*GetPaymentOperationRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{3}
}

func (x *GetPaymentOperationRequest) GetOrderId() uint64 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

type GetPaymentOperationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId       uint32           `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	PaymentId     uint32           `protobuf:"varint,2,opt,name=payment_id,json=paymentId,proto3" json:"payment_id,omitempty"`
	PaidAmount    float32          `protobuf:"fixed32,3,opt,name=paid_amount,json=paidAmount,proto3" json:"paid_amount,omitempty"`
	CardLast_4    string           `protobuf:"bytes,4,opt,name=card_last_4,json=cardLast4,proto3" json:"card_last_4,omitempty"`
	PaymentInfo   *structpb.Struct `protobuf:"bytes,5,opt,name=payment_info,json=paymentInfo,proto3" json:"payment_info,omitempty"`
	CreatedAt     int64            `protobuf:"varint,6,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     int64            `protobuf:"varint,7,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	PaymentMethod string           `protobuf:"bytes,8,opt,name=payment_method,json=paymentMethod,proto3" json:"payment_method,omitempty"`
}

func (x *GetPaymentOperationResponse) Reset() {
	*x = GetPaymentOperationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPaymentOperationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPaymentOperationResponse) ProtoMessage() {}

func (x *GetPaymentOperationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPaymentOperationResponse.ProtoReflect.Descriptor instead.
func (*GetPaymentOperationResponse) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{4}
}

func (x *GetPaymentOperationResponse) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

func (x *GetPaymentOperationResponse) GetPaymentId() uint32 {
	if x != nil {
		return x.PaymentId
	}
	return 0
}

func (x *GetPaymentOperationResponse) GetPaidAmount() float32 {
	if x != nil {
		return x.PaidAmount
	}
	return 0
}

func (x *GetPaymentOperationResponse) GetCardLast_4() string {
	if x != nil {
		return x.CardLast_4
	}
	return ""
}

func (x *GetPaymentOperationResponse) GetPaymentInfo() *structpb.Struct {
	if x != nil {
		return x.PaymentInfo
	}
	return nil
}

func (x *GetPaymentOperationResponse) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *GetPaymentOperationResponse) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *GetPaymentOperationResponse) GetPaymentMethod() string {
	if x != nil {
		return x.PaymentMethod
	}
	return ""
}

type SuccessPaymentOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *SuccessPaymentOperationRequest) Reset() {
	*x = SuccessPaymentOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SuccessPaymentOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SuccessPaymentOperationRequest) ProtoMessage() {}

func (x *SuccessPaymentOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SuccessPaymentOperationRequest.ProtoReflect.Descriptor instead.
func (*SuccessPaymentOperationRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{5}
}

func (x *SuccessPaymentOperationRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type FailPaymentOperationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Session string `protobuf:"bytes,1,opt,name=session,proto3" json:"session,omitempty"`
}

func (x *FailPaymentOperationRequest) Reset() {
	*x = FailPaymentOperationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FailPaymentOperationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FailPaymentOperationRequest) ProtoMessage() {}

func (x *FailPaymentOperationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FailPaymentOperationRequest.ProtoReflect.Descriptor instead.
func (*FailPaymentOperationRequest) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{6}
}

func (x *FailPaymentOperationRequest) GetSession() string {
	if x != nil {
		return x.Session
	}
	return ""
}

type PaymentDoneSuccessfully struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrderId uint32 `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (x *PaymentDoneSuccessfully) Reset() {
	*x = PaymentDoneSuccessfully{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_v1_payment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentDoneSuccessfully) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentDoneSuccessfully) ProtoMessage() {}

func (x *PaymentDoneSuccessfully) ProtoReflect() protoreflect.Message {
	mi := &file_protos_v1_payment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentDoneSuccessfully.ProtoReflect.Descriptor instead.
func (*PaymentDoneSuccessfully) Descriptor() ([]byte, []int) {
	return file_protos_v1_payment_proto_rawDescGZIP(), []int{7}
}

func (x *PaymentDoneSuccessfully) GetOrderId() uint32 {
	if x != nil {
		return x.OrderId
	}
	return 0
}

var File_protos_v1_payment_proto protoreflect.FileDescriptor

var file_protos_v1_payment_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x62, 0x6f, 0x64, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x73, 0x0a, 0x07, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x6e, 0x0a, 0x1a, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x22, 0x40, 0x0a, 0x1b, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x37, 0x0a, 0x1a, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x22, 0xb9, 0x02, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x09, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1f,
	0x0a, 0x0b, 0x70, 0x61, 0x69, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x0a, 0x70, 0x61, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0b, 0x63, 0x61, 0x72, 0x64, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x61, 0x72, 0x64, 0x4c, 0x61, 0x73, 0x74, 0x34, 0x12,
	0x3a, 0x0a, 0x0c, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x0b,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x22, 0x3a, 0x0a, 0x1e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x1b,
	0x46, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x34, 0x0a, 0x17, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74,
	0x44, 0x6f, 0x6e, 0x65, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x64, 0x32, 0xc0, 0x03, 0x0a, 0x0e,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x74,
	0x0a, 0x13, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x2c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x74, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x5f, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x63, 0x0a, 0x17, 0x53, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x30, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x70, 0x61,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x12,
	0x5d, 0x0a, 0x14, 0x46, 0x61, 0x69, 0x6c, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x76, 0x31, 0x2e, 0x46, 0x61, 0x69, 0x6c, 0x50,
	0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x00, 0x42, 0x03,
	0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_v1_payment_proto_rawDescOnce sync.Once
	file_protos_v1_payment_proto_rawDescData = file_protos_v1_payment_proto_rawDesc
)

func file_protos_v1_payment_proto_rawDescGZIP() []byte {
	file_protos_v1_payment_proto_rawDescOnce.Do(func() {
		file_protos_v1_payment_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_v1_payment_proto_rawDescData)
	})
	return file_protos_v1_payment_proto_rawDescData
}

var file_protos_v1_payment_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_protos_v1_payment_proto_goTypes = []any{
	(*Product)(nil),                        // 0: proto_payment_v1.Product
	(*GeneratePaymentLinkRequest)(nil),     // 1: proto_payment_v1.GeneratePaymentLinkRequest
	(*GeneratePaymentLinkResponse)(nil),    // 2: proto_payment_v1.GeneratePaymentLinkResponse
	(*GetPaymentOperationRequest)(nil),     // 3: proto_payment_v1.GetPaymentOperationRequest
	(*GetPaymentOperationResponse)(nil),    // 4: proto_payment_v1.GetPaymentOperationResponse
	(*SuccessPaymentOperationRequest)(nil), // 5: proto_payment_v1.SuccessPaymentOperationRequest
	(*FailPaymentOperationRequest)(nil),    // 6: proto_payment_v1.FailPaymentOperationRequest
	(*PaymentDoneSuccessfully)(nil),        // 7: proto_payment_v1.PaymentDoneSuccessfully
	(*structpb.Struct)(nil),                // 8: google.protobuf.Struct
	(*httpbody.HttpBody)(nil),              // 9: google.api.HttpBody
}
var file_protos_v1_payment_proto_depIdxs = []int32{
	0, // 0: proto_payment_v1.GeneratePaymentLinkRequest.products:type_name -> proto_payment_v1.Product
	8, // 1: proto_payment_v1.GetPaymentOperationResponse.payment_info:type_name -> google.protobuf.Struct
	1, // 2: proto_payment_v1.PaymentService.GeneratePaymentLink:input_type -> proto_payment_v1.GeneratePaymentLinkRequest
	3, // 3: proto_payment_v1.PaymentService.GetPaymentOperation:input_type -> proto_payment_v1.GetPaymentOperationRequest
	5, // 4: proto_payment_v1.PaymentService.SuccessPaymentOperation:input_type -> proto_payment_v1.SuccessPaymentOperationRequest
	6, // 5: proto_payment_v1.PaymentService.FailPaymentOperation:input_type -> proto_payment_v1.FailPaymentOperationRequest
	2, // 6: proto_payment_v1.PaymentService.GeneratePaymentLink:output_type -> proto_payment_v1.GeneratePaymentLinkResponse
	4, // 7: proto_payment_v1.PaymentService.GetPaymentOperation:output_type -> proto_payment_v1.GetPaymentOperationResponse
	9, // 8: proto_payment_v1.PaymentService.SuccessPaymentOperation:output_type -> google.api.HttpBody
	9, // 9: proto_payment_v1.PaymentService.FailPaymentOperation:output_type -> google.api.HttpBody
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_v1_payment_proto_init() }
func file_protos_v1_payment_proto_init() {
	if File_protos_v1_payment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_v1_payment_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Product); i {
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
		file_protos_v1_payment_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GeneratePaymentLinkRequest); i {
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
		file_protos_v1_payment_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GeneratePaymentLinkResponse); i {
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
		file_protos_v1_payment_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentOperationRequest); i {
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
		file_protos_v1_payment_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetPaymentOperationResponse); i {
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
		file_protos_v1_payment_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*SuccessPaymentOperationRequest); i {
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
		file_protos_v1_payment_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*FailPaymentOperationRequest); i {
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
		file_protos_v1_payment_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*PaymentDoneSuccessfully); i {
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
			RawDescriptor: file_protos_v1_payment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_v1_payment_proto_goTypes,
		DependencyIndexes: file_protos_v1_payment_proto_depIdxs,
		MessageInfos:      file_protos_v1_payment_proto_msgTypes,
	}.Build()
	File_protos_v1_payment_proto = out.File
	file_protos_v1_payment_proto_rawDesc = nil
	file_protos_v1_payment_proto_goTypes = nil
	file_protos_v1_payment_proto_depIdxs = nil
}
