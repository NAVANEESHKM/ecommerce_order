// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: order_proto/create_order.proto

package ecommerce_order

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

type CustomerOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string      `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CustomerId    string      `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	PaymentId     string      `protobuf:"bytes,3,opt,name=PaymentId,proto3" json:"PaymentId,omitempty"`
	PaymentStatus string      `protobuf:"bytes,4,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
	Status        string      `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Currency      string      `protobuf:"bytes,6,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Items         []string    `protobuf:"bytes,7,rep,name=Items,proto3" json:"Items,omitempty"`
	Shipping      []*Shipping `protobuf:"bytes,8,rep,name=Shipping,proto3" json:"Shipping,omitempty"`
	Carrier       string      `protobuf:"bytes,9,opt,name=Carrier,proto3" json:"Carrier,omitempty"`
	Tracking      string      `protobuf:"bytes,10,opt,name=Tracking,proto3" json:"Tracking,omitempty"`
}

func (x *CustomerOrder) Reset() {
	*x = CustomerOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_create_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerOrder) ProtoMessage() {}

func (x *CustomerOrder) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_create_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerOrder.ProtoReflect.Descriptor instead.
func (*CustomerOrder) Descriptor() ([]byte, []int) {
	return file_order_proto_create_order_proto_rawDescGZIP(), []int{0}
}

func (x *CustomerOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomerOrder) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerOrder) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *CustomerOrder) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *CustomerOrder) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CustomerOrder) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *CustomerOrder) GetItems() []string {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *CustomerOrder) GetShipping() []*Shipping {
	if x != nil {
		return x.Shipping
	}
	return nil
}

func (x *CustomerOrder) GetCarrier() string {
	if x != nil {
		return x.Carrier
	}
	return ""
}

func (x *CustomerOrder) GetTracking() string {
	if x != nil {
		return x.Tracking
	}
	return ""
}

type Shipping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address []*Address `protobuf:"bytes,1,rep,name=Address,proto3" json:"Address,omitempty"`
	Origin  []*Origin  `protobuf:"bytes,2,rep,name=Origin,proto3" json:"Origin,omitempty"`
}

func (x *Shipping) Reset() {
	*x = Shipping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_create_order_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shipping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shipping) ProtoMessage() {}

func (x *Shipping) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_create_order_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shipping.ProtoReflect.Descriptor instead.
func (*Shipping) Descriptor() ([]byte, []int) {
	return file_order_proto_create_order_proto_rawDescGZIP(), []int{1}
}

func (x *Shipping) GetAddress() []*Address {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *Shipping) GetOrigin() []*Origin {
	if x != nil {
		return x.Origin
	}
	return nil
}

type Address struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street1 string `protobuf:"bytes,1,opt,name=Street1,proto3" json:"Street1,omitempty"`
	Street2 string `protobuf:"bytes,2,opt,name=Street2,proto3" json:"Street2,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=Country,proto3" json:"Country,omitempty"`
	Zip     string `protobuf:"bytes,6,opt,name=Zip,proto3" json:"Zip,omitempty"`
}

func (x *Address) Reset() {
	*x = Address{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_create_order_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Address) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Address) ProtoMessage() {}

func (x *Address) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_create_order_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Address.ProtoReflect.Descriptor instead.
func (*Address) Descriptor() ([]byte, []int) {
	return file_order_proto_create_order_proto_rawDescGZIP(), []int{2}
}

func (x *Address) GetStreet1() string {
	if x != nil {
		return x.Street1
	}
	return ""
}

func (x *Address) GetStreet2() string {
	if x != nil {
		return x.Street2
	}
	return ""
}

func (x *Address) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Address) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Address) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Address) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

type Origin struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Street1 string `protobuf:"bytes,1,opt,name=Street1,proto3" json:"Street1,omitempty"`
	Street2 string `protobuf:"bytes,2,opt,name=Street2,proto3" json:"Street2,omitempty"`
	City    string `protobuf:"bytes,3,opt,name=City,proto3" json:"City,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Country string `protobuf:"bytes,5,opt,name=Country,proto3" json:"Country,omitempty"`
	Zip     string `protobuf:"bytes,6,opt,name=Zip,proto3" json:"Zip,omitempty"`
}

func (x *Origin) Reset() {
	*x = Origin{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_create_order_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Origin) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Origin) ProtoMessage() {}

func (x *Origin) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_create_order_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Origin.ProtoReflect.Descriptor instead.
func (*Origin) Descriptor() ([]byte, []int) {
	return file_order_proto_create_order_proto_rawDescGZIP(), []int{3}
}

func (x *Origin) GetStreet1() string {
	if x != nil {
		return x.Street1
	}
	return ""
}

func (x *Origin) GetStreet2() string {
	if x != nil {
		return x.Street2
	}
	return ""
}

func (x *Origin) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

func (x *Origin) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Origin) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Origin) GetZip() string {
	if x != nil {
		return x.Zip
	}
	return ""
}

type CustomerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	CustomerId    string `protobuf:"bytes,2,opt,name=CustomerId,proto3" json:"CustomerId,omitempty"`
	PaymentId     string `protobuf:"bytes,3,opt,name=PaymentId,proto3" json:"PaymentId,omitempty"`
	PaymentStatus string `protobuf:"bytes,4,opt,name=PaymentStatus,proto3" json:"PaymentStatus,omitempty"`
	Status        string `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Currency      string `protobuf:"bytes,6,opt,name=Currency,proto3" json:"Currency,omitempty"`
}

func (x *CustomerResponse) Reset() {
	*x = CustomerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_proto_create_order_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CustomerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CustomerResponse) ProtoMessage() {}

func (x *CustomerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_order_proto_create_order_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CustomerResponse.ProtoReflect.Descriptor instead.
func (*CustomerResponse) Descriptor() ([]byte, []int) {
	return file_order_proto_create_order_proto_rawDescGZIP(), []int{4}
}

func (x *CustomerResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CustomerResponse) GetCustomerId() string {
	if x != nil {
		return x.CustomerId
	}
	return ""
}

func (x *CustomerResponse) GetPaymentId() string {
	if x != nil {
		return x.PaymentId
	}
	return ""
}

func (x *CustomerResponse) GetPaymentStatus() string {
	if x != nil {
		return x.PaymentStatus
	}
	return ""
}

func (x *CustomerResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *CustomerResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

var File_order_proto_create_order_proto protoreflect.FileDescriptor

var file_order_proto_create_order_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x02,
	0x0a, 0x0d, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x49, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x31, 0x0a,
	0x08, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x68,
	0x69, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x52, 0x08, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69, 0x6e, 0x67,
	0x12, 0x18, 0x0a, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x43, 0x61, 0x72, 0x72, 0x69, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x54, 0x72,
	0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x67, 0x0a, 0x08, 0x53, 0x68, 0x69, 0x70, 0x70, 0x69,
	0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x52, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x22,
	0x93, 0x01, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x53,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x32, 0x12,
	0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43,
	0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x5a, 0x69, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x5a, 0x69, 0x70, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x31, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x53, 0x74, 0x72, 0x65, 0x65, 0x74, 0x31, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x74,
	0x72, 0x65, 0x65, 0x74, 0x32, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x53, 0x74, 0x72,
	0x65, 0x65, 0x74, 0x32, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x5a, 0x69, 0x70, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x5a, 0x69, 0x70, 0x22, 0xba, 0x01, 0x0a, 0x10, 0x43,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x49, 0x64, 0x12,
	0x1e, 0x0a, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x24, 0x0a,
	0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x42, 0x30, 0x5a, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x73,
	0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4e, 0x41,
	0x56, 0x41, 0x4e, 0x45, 0x45, 0x53, 0x48, 0x4b, 0x4d, 0x2f, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x72, 0x63, 0x65, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_order_proto_create_order_proto_rawDescOnce sync.Once
	file_order_proto_create_order_proto_rawDescData = file_order_proto_create_order_proto_rawDesc
)

func file_order_proto_create_order_proto_rawDescGZIP() []byte {
	file_order_proto_create_order_proto_rawDescOnce.Do(func() {
		file_order_proto_create_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_proto_create_order_proto_rawDescData)
	})
	return file_order_proto_create_order_proto_rawDescData
}

var file_order_proto_create_order_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_order_proto_create_order_proto_goTypes = []interface{}{
	(*CustomerOrder)(nil),    // 0: order_proto.CustomerOrder
	(*Shipping)(nil),         // 1: order_proto.Shipping
	(*Address)(nil),          // 2: order_proto.Address
	(*Origin)(nil),           // 3: order_proto.Origin
	(*CustomerResponse)(nil), // 4: order_proto.CustomerResponse
}
var file_order_proto_create_order_proto_depIdxs = []int32{
	1, // 0: order_proto.CustomerOrder.Shipping:type_name -> order_proto.Shipping
	2, // 1: order_proto.Shipping.Address:type_name -> order_proto.Address
	3, // 2: order_proto.Shipping.Origin:type_name -> order_proto.Origin
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_order_proto_create_order_proto_init() }
func file_order_proto_create_order_proto_init() {
	if File_order_proto_create_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_proto_create_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerOrder); i {
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
		file_order_proto_create_order_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shipping); i {
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
		file_order_proto_create_order_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Address); i {
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
		file_order_proto_create_order_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Origin); i {
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
		file_order_proto_create_order_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CustomerResponse); i {
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
			RawDescriptor: file_order_proto_create_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_order_proto_create_order_proto_goTypes,
		DependencyIndexes: file_order_proto_create_order_proto_depIdxs,
		MessageInfos:      file_order_proto_create_order_proto_msgTypes,
	}.Build()
	File_order_proto_create_order_proto = out.File
	file_order_proto_create_order_proto_rawDesc = nil
	file_order_proto_create_order_proto_goTypes = nil
	file_order_proto_create_order_proto_depIdxs = nil
}
