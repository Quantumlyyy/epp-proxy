// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: isnic/isnic.proto

package isnic

import (
	common "github.com/as207960/epp-proxy/gen/go/epp/common"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ContactStatus int32

const (
	ContactStatus_Ok              ContactStatus = 0
	ContactStatus_OkUnconfirmed   ContactStatus = 1
	ContactStatus_PendingCreate   ContactStatus = 2
	ContactStatus_ServerExpired   ContactStatus = 3
	ContactStatus_ServerSuspended ContactStatus = 4
)

// Enum value maps for ContactStatus.
var (
	ContactStatus_name = map[int32]string{
		0: "Ok",
		1: "OkUnconfirmed",
		2: "PendingCreate",
		3: "ServerExpired",
		4: "ServerSuspended",
	}
	ContactStatus_value = map[string]int32{
		"Ok":              0,
		"OkUnconfirmed":   1,
		"PendingCreate":   2,
		"ServerExpired":   3,
		"ServerSuspended": 4,
	}
)

func (x ContactStatus) Enum() *ContactStatus {
	p := new(ContactStatus)
	*p = x
	return p
}

func (x ContactStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContactStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_isnic_isnic_proto_enumTypes[0].Descriptor()
}

func (ContactStatus) Type() protoreflect.EnumType {
	return &file_isnic_isnic_proto_enumTypes[0]
}

func (x ContactStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContactStatus.Descriptor instead.
func (ContactStatus) EnumDescriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{0}
}

type PaymentInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to PaymentMethod:
	//	*PaymentInfo_Prepaid
	//	*PaymentInfo_Card_
	PaymentMethod isPaymentInfo_PaymentMethod `protobuf_oneof:"payment_method"`
}

func (x *PaymentInfo) Reset() {
	*x = PaymentInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo) ProtoMessage() {}

func (x *PaymentInfo) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo.ProtoReflect.Descriptor instead.
func (*PaymentInfo) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{0}
}

func (m *PaymentInfo) GetPaymentMethod() isPaymentInfo_PaymentMethod {
	if m != nil {
		return m.PaymentMethod
	}
	return nil
}

func (x *PaymentInfo) GetPrepaid() uint32 {
	if x, ok := x.GetPaymentMethod().(*PaymentInfo_Prepaid); ok {
		return x.Prepaid
	}
	return 0
}

func (x *PaymentInfo) GetCard() *PaymentInfo_Card {
	if x, ok := x.GetPaymentMethod().(*PaymentInfo_Card_); ok {
		return x.Card
	}
	return nil
}

type isPaymentInfo_PaymentMethod interface {
	isPaymentInfo_PaymentMethod()
}

type PaymentInfo_Prepaid struct {
	Prepaid uint32 `protobuf:"varint,1,opt,name=prepaid,proto3,oneof"`
}

type PaymentInfo_Card_ struct {
	Card *PaymentInfo_Card `protobuf:"bytes,2,opt,name=card,proto3,oneof"`
}

func (*PaymentInfo_Prepaid) isPaymentInfo_PaymentMethod() {}

func (*PaymentInfo_Card_) isPaymentInfo_PaymentMethod() {}

type DomainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneContact *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=zone_contact,json=zoneContact,proto3" json:"zone_contact,omitempty"`
}

func (x *DomainInfo) Reset() {
	*x = DomainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainInfo) ProtoMessage() {}

func (x *DomainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainInfo.ProtoReflect.Descriptor instead.
func (*DomainInfo) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{1}
}

func (x *DomainInfo) GetZoneContact() *wrapperspb.StringValue {
	if x != nil {
		return x.ZoneContact
	}
	return nil
}

type DomainUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RemoveAllNs bool     `protobuf:"varint,1,opt,name=remove_all_ns,json=removeAllNs,proto3" json:"remove_all_ns,omitempty"`
	NewMasterNs []string `protobuf:"bytes,2,rep,name=new_master_ns,json=newMasterNs,proto3" json:"new_master_ns,omitempty"`
}

func (x *DomainUpdate) Reset() {
	*x = DomainUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainUpdate) ProtoMessage() {}

func (x *DomainUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainUpdate.ProtoReflect.Descriptor instead.
func (*DomainUpdate) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{2}
}

func (x *DomainUpdate) GetRemoveAllNs() bool {
	if x != nil {
		return x.RemoveAllNs
	}
	return false
}

func (x *DomainUpdate) GetNewMasterNs() []string {
	if x != nil {
		return x.NewMasterNs
	}
	return nil
}

type ContactInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Statuses                       []ContactStatus         `protobuf:"varint,1,rep,packed,name=statuses,proto3,enum=epp.isnic.ContactStatus" json:"statuses,omitempty"`
	Mobile                         *common.Phone           `protobuf:"bytes,2,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Sid                            *wrapperspb.StringValue `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"`
	AutoUpdateFromNationalRegistry bool                    `protobuf:"varint,4,opt,name=auto_update_from_national_registry,json=autoUpdateFromNationalRegistry,proto3" json:"auto_update_from_national_registry,omitempty"`
	PaperInvoices                  bool                    `protobuf:"varint,5,opt,name=paper_invoices,json=paperInvoices,proto3" json:"paper_invoices,omitempty"`
}

func (x *ContactInfo) Reset() {
	*x = ContactInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactInfo) ProtoMessage() {}

func (x *ContactInfo) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactInfo.ProtoReflect.Descriptor instead.
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{3}
}

func (x *ContactInfo) GetStatuses() []ContactStatus {
	if x != nil {
		return x.Statuses
	}
	return nil
}

func (x *ContactInfo) GetMobile() *common.Phone {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *ContactInfo) GetSid() *wrapperspb.StringValue {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *ContactInfo) GetAutoUpdateFromNationalRegistry() bool {
	if x != nil {
		return x.AutoUpdateFromNationalRegistry
	}
	return false
}

func (x *ContactInfo) GetPaperInvoices() bool {
	if x != nil {
		return x.PaperInvoices
	}
	return false
}

type ContactCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile                         *common.Phone           `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Sid                            *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=sid,proto3" json:"sid,omitempty"`
	AutoUpdateFromNationalRegistry bool                    `protobuf:"varint,3,opt,name=auto_update_from_national_registry,json=autoUpdateFromNationalRegistry,proto3" json:"auto_update_from_national_registry,omitempty"`
	PaperInvoices                  bool                    `protobuf:"varint,4,opt,name=paper_invoices,json=paperInvoices,proto3" json:"paper_invoices,omitempty"`
	Lang                           *wrapperspb.StringValue `protobuf:"bytes,5,opt,name=lang,proto3" json:"lang,omitempty"`
}

func (x *ContactCreate) Reset() {
	*x = ContactCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactCreate) ProtoMessage() {}

func (x *ContactCreate) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactCreate.ProtoReflect.Descriptor instead.
func (*ContactCreate) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{4}
}

func (x *ContactCreate) GetMobile() *common.Phone {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *ContactCreate) GetSid() *wrapperspb.StringValue {
	if x != nil {
		return x.Sid
	}
	return nil
}

func (x *ContactCreate) GetAutoUpdateFromNationalRegistry() bool {
	if x != nil {
		return x.AutoUpdateFromNationalRegistry
	}
	return false
}

func (x *ContactCreate) GetPaperInvoices() bool {
	if x != nil {
		return x.PaperInvoices
	}
	return false
}

func (x *ContactCreate) GetLang() *wrapperspb.StringValue {
	if x != nil {
		return x.Lang
	}
	return nil
}

type ContactUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mobile                         *common.Phone           `protobuf:"bytes,1,opt,name=mobile,proto3" json:"mobile,omitempty"`
	AutoUpdateFromNationalRegistry *wrapperspb.BoolValue   `protobuf:"bytes,2,opt,name=auto_update_from_national_registry,json=autoUpdateFromNationalRegistry,proto3" json:"auto_update_from_national_registry,omitempty"`
	PaperInvoices                  *wrapperspb.BoolValue   `protobuf:"bytes,3,opt,name=paper_invoices,json=paperInvoices,proto3" json:"paper_invoices,omitempty"`
	Lang                           *wrapperspb.StringValue `protobuf:"bytes,4,opt,name=lang,proto3" json:"lang,omitempty"`
	Password                       string                  `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *ContactUpdate) Reset() {
	*x = ContactUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ContactUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContactUpdate) ProtoMessage() {}

func (x *ContactUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContactUpdate.ProtoReflect.Descriptor instead.
func (*ContactUpdate) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{5}
}

func (x *ContactUpdate) GetMobile() *common.Phone {
	if x != nil {
		return x.Mobile
	}
	return nil
}

func (x *ContactUpdate) GetAutoUpdateFromNationalRegistry() *wrapperspb.BoolValue {
	if x != nil {
		return x.AutoUpdateFromNationalRegistry
	}
	return nil
}

func (x *ContactUpdate) GetPaperInvoices() *wrapperspb.BoolValue {
	if x != nil {
		return x.PaperInvoices
	}
	return nil
}

func (x *ContactUpdate) GetLang() *wrapperspb.StringValue {
	if x != nil {
		return x.Lang
	}
	return nil
}

func (x *ContactUpdate) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type HostInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ZoneContact *wrapperspb.StringValue `protobuf:"bytes,1,opt,name=zone_contact,json=zoneContact,proto3" json:"zone_contact,omitempty"`
}

func (x *HostInfo) Reset() {
	*x = HostInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HostInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HostInfo) ProtoMessage() {}

func (x *HostInfo) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HostInfo.ProtoReflect.Descriptor instead.
func (*HostInfo) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{6}
}

func (x *HostInfo) GetZoneContact() *wrapperspb.StringValue {
	if x != nil {
		return x.ZoneContact
	}
	return nil
}

type PaymentInfo_Card struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Cvc string `protobuf:"bytes,2,opt,name=cvc,proto3" json:"cvc,omitempty"`
}

func (x *PaymentInfo_Card) Reset() {
	*x = PaymentInfo_Card{}
	if protoimpl.UnsafeEnabled {
		mi := &file_isnic_isnic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaymentInfo_Card) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaymentInfo_Card) ProtoMessage() {}

func (x *PaymentInfo_Card) ProtoReflect() protoreflect.Message {
	mi := &file_isnic_isnic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaymentInfo_Card.ProtoReflect.Descriptor instead.
func (*PaymentInfo_Card) Descriptor() ([]byte, []int) {
	return file_isnic_isnic_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PaymentInfo_Card) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *PaymentInfo_Card) GetCvc() string {
	if x != nil {
		return x.Cvc
	}
	return ""
}

var File_isnic_isnic_proto protoreflect.FileDescriptor

var file_isnic_isnic_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x73, 0x6e, 0x69, 0x63, 0x2f, 0x69, 0x73, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x65, 0x70, 0x70, 0x2e, 0x69, 0x73, 0x6e, 0x69, 0x63, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x98, 0x01, 0x0a, 0x0b, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x07, 0x70, 0x72, 0x65, 0x70, 0x61, 0x69, 0x64, 0x12,
	0x31, 0x0a, 0x04, 0x63, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x65, 0x70, 0x70, 0x2e, 0x69, 0x73, 0x6e, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x64, 0x48, 0x00, 0x52, 0x04, 0x63, 0x61,
	0x72, 0x64, 0x1a, 0x28, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x76,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x63, 0x76, 0x63, 0x42, 0x10, 0x0a, 0x0e,
	0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x22, 0x4d,
	0x0a, 0x0a, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a, 0x0c,
	0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x22, 0x56, 0x0a,
	0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a,
	0x0d, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x6c, 0x6c, 0x5f, 0x6e, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x4e,
	0x73, 0x12, 0x22, 0x0a, 0x0d, 0x6e, 0x65, 0x77, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x6e, 0x65, 0x77, 0x4d, 0x61, 0x73,
	0x74, 0x65, 0x72, 0x4e, 0x73, 0x22, 0x91, 0x02, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x34, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x18, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x69, 0x73,
	0x6e, 0x69, 0x63, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x65, 0x73, 0x12, 0x29, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x70,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x22, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x22, 0x8f, 0x02, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x63, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a, 0x06, 0x6d,
	0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x70,
	0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x06,
	0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x2e, 0x0a, 0x03, 0x73, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x4a, 0x0a, 0x22, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x6f, 0x6d, 0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x6f,
	0x69, 0x63, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x70, 0x61, 0x70, 0x65,
	0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x61, 0x6e,
	0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x22, 0xb3, 0x02, 0x0a, 0x0d,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x29, 0x0a,
	0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x65, 0x70, 0x70, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x50, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12, 0x66, 0x0a, 0x22, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x66, 0x72, 0x6f, 0x6d, 0x5f, 0x6e, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x1e, 0x61, 0x75, 0x74, 0x6f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x6f, 0x6d,
	0x4e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x79,
	0x12, 0x41, 0x0a, 0x0e, 0x70, 0x61, 0x70, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x76, 0x6f, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x70, 0x61, 0x70, 0x65, 0x72, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x04, 0x6c, 0x61, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x04, 0x6c, 0x61, 0x6e, 0x67, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x4b, 0x0a, 0x08, 0x48, 0x6f, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3f, 0x0a,
	0x0c, 0x7a, 0x6f, 0x6e, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x0b, 0x7a, 0x6f, 0x6e, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2a, 0x65,
	0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x06, 0x0a, 0x02, 0x4f, 0x6b, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4f, 0x6b, 0x55, 0x6e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x65, 0x64, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x65,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x10, 0x02, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x64, 0x10, 0x03,
	0x12, 0x13, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x64, 0x10, 0x04, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x32, 0x30, 0x37, 0x39, 0x36, 0x30, 0x2f, 0x65, 0x70, 0x70,
	0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x70,
	0x70, 0x2f, 0x69, 0x73, 0x6e, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_isnic_isnic_proto_rawDescOnce sync.Once
	file_isnic_isnic_proto_rawDescData = file_isnic_isnic_proto_rawDesc
)

func file_isnic_isnic_proto_rawDescGZIP() []byte {
	file_isnic_isnic_proto_rawDescOnce.Do(func() {
		file_isnic_isnic_proto_rawDescData = protoimpl.X.CompressGZIP(file_isnic_isnic_proto_rawDescData)
	})
	return file_isnic_isnic_proto_rawDescData
}

var file_isnic_isnic_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_isnic_isnic_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_isnic_isnic_proto_goTypes = []interface{}{
	(ContactStatus)(0),             // 0: epp.isnic.ContactStatus
	(*PaymentInfo)(nil),            // 1: epp.isnic.PaymentInfo
	(*DomainInfo)(nil),             // 2: epp.isnic.DomainInfo
	(*DomainUpdate)(nil),           // 3: epp.isnic.DomainUpdate
	(*ContactInfo)(nil),            // 4: epp.isnic.ContactInfo
	(*ContactCreate)(nil),          // 5: epp.isnic.ContactCreate
	(*ContactUpdate)(nil),          // 6: epp.isnic.ContactUpdate
	(*HostInfo)(nil),               // 7: epp.isnic.HostInfo
	(*PaymentInfo_Card)(nil),       // 8: epp.isnic.PaymentInfo.Card
	(*wrapperspb.StringValue)(nil), // 9: google.protobuf.StringValue
	(*common.Phone)(nil),           // 10: epp.common.Phone
	(*wrapperspb.BoolValue)(nil),   // 11: google.protobuf.BoolValue
}
var file_isnic_isnic_proto_depIdxs = []int32{
	8,  // 0: epp.isnic.PaymentInfo.card:type_name -> epp.isnic.PaymentInfo.Card
	9,  // 1: epp.isnic.DomainInfo.zone_contact:type_name -> google.protobuf.StringValue
	0,  // 2: epp.isnic.ContactInfo.statuses:type_name -> epp.isnic.ContactStatus
	10, // 3: epp.isnic.ContactInfo.mobile:type_name -> epp.common.Phone
	9,  // 4: epp.isnic.ContactInfo.sid:type_name -> google.protobuf.StringValue
	10, // 5: epp.isnic.ContactCreate.mobile:type_name -> epp.common.Phone
	9,  // 6: epp.isnic.ContactCreate.sid:type_name -> google.protobuf.StringValue
	9,  // 7: epp.isnic.ContactCreate.lang:type_name -> google.protobuf.StringValue
	10, // 8: epp.isnic.ContactUpdate.mobile:type_name -> epp.common.Phone
	11, // 9: epp.isnic.ContactUpdate.auto_update_from_national_registry:type_name -> google.protobuf.BoolValue
	11, // 10: epp.isnic.ContactUpdate.paper_invoices:type_name -> google.protobuf.BoolValue
	9,  // 11: epp.isnic.ContactUpdate.lang:type_name -> google.protobuf.StringValue
	9,  // 12: epp.isnic.HostInfo.zone_contact:type_name -> google.protobuf.StringValue
	13, // [13:13] is the sub-list for method output_type
	13, // [13:13] is the sub-list for method input_type
	13, // [13:13] is the sub-list for extension type_name
	13, // [13:13] is the sub-list for extension extendee
	0,  // [0:13] is the sub-list for field type_name
}

func init() { file_isnic_isnic_proto_init() }
func file_isnic_isnic_proto_init() {
	if File_isnic_isnic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_isnic_isnic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentInfo); i {
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
		file_isnic_isnic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainInfo); i {
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
		file_isnic_isnic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainUpdate); i {
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
		file_isnic_isnic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactInfo); i {
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
		file_isnic_isnic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactCreate); i {
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
		file_isnic_isnic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ContactUpdate); i {
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
		file_isnic_isnic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HostInfo); i {
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
		file_isnic_isnic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaymentInfo_Card); i {
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
	file_isnic_isnic_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*PaymentInfo_Prepaid)(nil),
		(*PaymentInfo_Card_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_isnic_isnic_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_isnic_isnic_proto_goTypes,
		DependencyIndexes: file_isnic_isnic_proto_depIdxs,
		EnumInfos:         file_isnic_isnic_proto_enumTypes,
		MessageInfos:      file_isnic_isnic_proto_msgTypes,
	}.Build()
	File_isnic_isnic_proto = out.File
	file_isnic_isnic_proto_rawDesc = nil
	file_isnic_isnic_proto_goTypes = nil
	file_isnic_isnic_proto_depIdxs = nil
}