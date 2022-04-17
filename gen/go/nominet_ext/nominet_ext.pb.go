// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: nominet_ext/nominet_ext.proto

package nominet_ext

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type BillType int32

const (
	BillType_Unspecified   BillType = 0
	BillType_BillRegistrar BillType = 1
	BillType_BillCustomer  BillType = 2
)

// Enum value maps for BillType.
var (
	BillType_name = map[int32]string{
		0: "Unspecified",
		1: "BillRegistrar",
		2: "BillCustomer",
	}
	BillType_value = map[string]int32{
		"Unspecified":   0,
		"BillRegistrar": 1,
		"BillCustomer":  2,
	}
)

func (x BillType) Enum() *BillType {
	p := new(BillType)
	*p = x
	return p
}

func (x BillType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BillType) Descriptor() protoreflect.EnumDescriptor {
	return file_nominet_ext_nominet_ext_proto_enumTypes[0].Descriptor()
}

func (BillType) Type() protoreflect.EnumType {
	return &file_nominet_ext_nominet_ext_proto_enumTypes[0]
}

func (x BillType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BillType.Descriptor instead.
func (BillType) EnumDescriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{0}
}

type RegistrationStatus int32

const (
	RegistrationStatus_RegisteredUntilExpiry RegistrationStatus = 0
	RegistrationStatus_RenewalRequired       RegistrationStatus = 1
	RegistrationStatus_NoLongerRequired      RegistrationStatus = 2
)

// Enum value maps for RegistrationStatus.
var (
	RegistrationStatus_name = map[int32]string{
		0: "RegisteredUntilExpiry",
		1: "RenewalRequired",
		2: "NoLongerRequired",
	}
	RegistrationStatus_value = map[string]int32{
		"RegisteredUntilExpiry": 0,
		"RenewalRequired":       1,
		"NoLongerRequired":      2,
	}
)

func (x RegistrationStatus) Enum() *RegistrationStatus {
	p := new(RegistrationStatus)
	*p = x
	return p
}

func (x RegistrationStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RegistrationStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_nominet_ext_nominet_ext_proto_enumTypes[1].Descriptor()
}

func (RegistrationStatus) Type() protoreflect.EnumType {
	return &file_nominet_ext_nominet_ext_proto_enumTypes[1]
}

func (x RegistrationStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RegistrationStatus.Descriptor instead.
func (RegistrationStatus) EnumDescriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{1}
}

type DataQualityStatus int32

const (
	DataQualityStatus_Invalid DataQualityStatus = 0
	DataQualityStatus_Valid   DataQualityStatus = 1
)

// Enum value maps for DataQualityStatus.
var (
	DataQualityStatus_name = map[int32]string{
		0: "Invalid",
		1: "Valid",
	}
	DataQualityStatus_value = map[string]int32{
		"Invalid": 0,
		"Valid":   1,
	}
)

func (x DataQualityStatus) Enum() *DataQualityStatus {
	p := new(DataQualityStatus)
	*p = x
	return p
}

func (x DataQualityStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataQualityStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_nominet_ext_nominet_ext_proto_enumTypes[2].Descriptor()
}

func (DataQualityStatus) Type() protoreflect.EnumType {
	return &file_nominet_ext_nominet_ext_proto_enumTypes[2]
}

func (x DataQualityStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataQualityStatus.Descriptor instead.
func (DataQualityStatus) EnumDescriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{2}
}

type DomainCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstBill  BillType                `protobuf:"varint,1,opt,name=first_bill,json=firstBill,proto3,enum=epp.nominet_ext.BillType" json:"first_bill,omitempty"`
	RecurBill  BillType                `protobuf:"varint,2,opt,name=recur_bill,json=recurBill,proto3,enum=epp.nominet_ext.BillType" json:"recur_bill,omitempty"`
	AutoBill   *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=auto_bill,json=autoBill,proto3" json:"auto_bill,omitempty"`
	NextBill   *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=next_bill,json=nextBill,proto3" json:"next_bill,omitempty"`
	AutoPeriod *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=auto_period,json=autoPeriod,proto3" json:"auto_period,omitempty"`
	NextPeriod *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=next_period,json=nextPeriod,proto3" json:"next_period,omitempty"`
	Notes      []string                `protobuf:"bytes,7,rep,name=notes,proto3" json:"notes,omitempty"`
	Reseller   *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=reseller,proto3" json:"reseller,omitempty"`
}

func (x *DomainCreate) Reset() {
	*x = DomainCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nominet_ext_nominet_ext_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCreate) ProtoMessage() {}

func (x *DomainCreate) ProtoReflect() protoreflect.Message {
	mi := &file_nominet_ext_nominet_ext_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCreate.ProtoReflect.Descriptor instead.
func (*DomainCreate) Descriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{0}
}

func (x *DomainCreate) GetFirstBill() BillType {
	if x != nil {
		return x.FirstBill
	}
	return BillType_Unspecified
}

func (x *DomainCreate) GetRecurBill() BillType {
	if x != nil {
		return x.RecurBill
	}
	return BillType_Unspecified
}

func (x *DomainCreate) GetAutoBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoBill
	}
	return nil
}

func (x *DomainCreate) GetNextBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextBill
	}
	return nil
}

func (x *DomainCreate) GetAutoPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoPeriod
	}
	return nil
}

func (x *DomainCreate) GetNextPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextPeriod
	}
	return nil
}

func (x *DomainCreate) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *DomainCreate) GetReseller() *wrapperspb.StringValue {
	if x != nil {
		return x.Reseller
	}
	return nil
}

type DomainUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstBill          BillType                `protobuf:"varint,1,opt,name=first_bill,json=firstBill,proto3,enum=epp.nominet_ext.BillType" json:"first_bill,omitempty"`
	RecurBill          BillType                `protobuf:"varint,2,opt,name=recur_bill,json=recurBill,proto3,enum=epp.nominet_ext.BillType" json:"recur_bill,omitempty"`
	AutoBill           *wrapperspb.UInt32Value `protobuf:"bytes,3,opt,name=auto_bill,json=autoBill,proto3" json:"auto_bill,omitempty"`
	NextBill           *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=next_bill,json=nextBill,proto3" json:"next_bill,omitempty"`
	AutoPeriod         *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=auto_period,json=autoPeriod,proto3" json:"auto_period,omitempty"`
	NextPeriod         *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=next_period,json=nextPeriod,proto3" json:"next_period,omitempty"`
	RenewalNotRequired bool                    `protobuf:"varint,7,opt,name=renewal_not_required,json=renewalNotRequired,proto3" json:"renewal_not_required,omitempty"`
	Notes              []string                `protobuf:"bytes,8,rep,name=notes,proto3" json:"notes,omitempty"`
	Reseller           *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=reseller,proto3" json:"reseller,omitempty"`
}

func (x *DomainUpdate) Reset() {
	*x = DomainUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nominet_ext_nominet_ext_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainUpdate) ProtoMessage() {}

func (x *DomainUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_nominet_ext_nominet_ext_proto_msgTypes[1]
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
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{1}
}

func (x *DomainUpdate) GetFirstBill() BillType {
	if x != nil {
		return x.FirstBill
	}
	return BillType_Unspecified
}

func (x *DomainUpdate) GetRecurBill() BillType {
	if x != nil {
		return x.RecurBill
	}
	return BillType_Unspecified
}

func (x *DomainUpdate) GetAutoBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoBill
	}
	return nil
}

func (x *DomainUpdate) GetNextBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextBill
	}
	return nil
}

func (x *DomainUpdate) GetAutoPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoPeriod
	}
	return nil
}

func (x *DomainUpdate) GetNextPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextPeriod
	}
	return nil
}

func (x *DomainUpdate) GetRenewalNotRequired() bool {
	if x != nil {
		return x.RenewalNotRequired
	}
	return false
}

func (x *DomainUpdate) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *DomainUpdate) GetReseller() *wrapperspb.StringValue {
	if x != nil {
		return x.Reseller
	}
	return nil
}

type DomainCheckInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AbuseLimit uint64 `protobuf:"varint,1,opt,name=abuse_limit,json=abuseLimit,proto3" json:"abuse_limit,omitempty"`
}

func (x *DomainCheckInfo) Reset() {
	*x = DomainCheckInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nominet_ext_nominet_ext_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainCheckInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainCheckInfo) ProtoMessage() {}

func (x *DomainCheckInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nominet_ext_nominet_ext_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DomainCheckInfo.ProtoReflect.Descriptor instead.
func (*DomainCheckInfo) Descriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{2}
}

func (x *DomainCheckInfo) GetAbuseLimit() uint64 {
	if x != nil {
		return x.AbuseLimit
	}
	return 0
}

type DomainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegistrationStatus RegistrationStatus      `protobuf:"varint,1,opt,name=registration_status,json=registrationStatus,proto3,enum=epp.nominet_ext.RegistrationStatus" json:"registration_status,omitempty"`
	FirstBill          BillType                `protobuf:"varint,2,opt,name=first_bill,json=firstBill,proto3,enum=epp.nominet_ext.BillType" json:"first_bill,omitempty"`
	RecurBill          BillType                `protobuf:"varint,3,opt,name=recur_bill,json=recurBill,proto3,enum=epp.nominet_ext.BillType" json:"recur_bill,omitempty"`
	AutoBill           *wrapperspb.UInt32Value `protobuf:"bytes,4,opt,name=auto_bill,json=autoBill,proto3" json:"auto_bill,omitempty"`
	NextBill           *wrapperspb.UInt32Value `protobuf:"bytes,5,opt,name=next_bill,json=nextBill,proto3" json:"next_bill,omitempty"`
	AutoPeriod         *wrapperspb.UInt32Value `protobuf:"bytes,6,opt,name=auto_period,json=autoPeriod,proto3" json:"auto_period,omitempty"`
	NextPeriod         *wrapperspb.UInt32Value `protobuf:"bytes,7,opt,name=next_period,json=nextPeriod,proto3" json:"next_period,omitempty"`
	RenewalNotRequired bool                    `protobuf:"varint,8,opt,name=renewal_not_required,json=renewalNotRequired,proto3" json:"renewal_not_required,omitempty"`
	Notes              []string                `protobuf:"bytes,9,rep,name=notes,proto3" json:"notes,omitempty"`
	Reseller           *wrapperspb.StringValue `protobuf:"bytes,10,opt,name=reseller,proto3" json:"reseller,omitempty"`
}

func (x *DomainInfo) Reset() {
	*x = DomainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nominet_ext_nominet_ext_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DomainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DomainInfo) ProtoMessage() {}

func (x *DomainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_nominet_ext_nominet_ext_proto_msgTypes[3]
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
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{3}
}

func (x *DomainInfo) GetRegistrationStatus() RegistrationStatus {
	if x != nil {
		return x.RegistrationStatus
	}
	return RegistrationStatus_RegisteredUntilExpiry
}

func (x *DomainInfo) GetFirstBill() BillType {
	if x != nil {
		return x.FirstBill
	}
	return BillType_Unspecified
}

func (x *DomainInfo) GetRecurBill() BillType {
	if x != nil {
		return x.RecurBill
	}
	return BillType_Unspecified
}

func (x *DomainInfo) GetAutoBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoBill
	}
	return nil
}

func (x *DomainInfo) GetNextBill() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextBill
	}
	return nil
}

func (x *DomainInfo) GetAutoPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.AutoPeriod
	}
	return nil
}

func (x *DomainInfo) GetNextPeriod() *wrapperspb.UInt32Value {
	if x != nil {
		return x.NextPeriod
	}
	return nil
}

func (x *DomainInfo) GetRenewalNotRequired() bool {
	if x != nil {
		return x.RenewalNotRequired
	}
	return false
}

func (x *DomainInfo) GetNotes() []string {
	if x != nil {
		return x.Notes
	}
	return nil
}

func (x *DomainInfo) GetReseller() *wrapperspb.StringValue {
	if x != nil {
		return x.Reseller
	}
	return nil
}

type DataQuality struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status        DataQualityStatus       `protobuf:"varint,1,opt,name=status,proto3,enum=epp.nominet_ext.DataQualityStatus" json:"status,omitempty"`
	Reason        *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	DateCommenced *timestamppb.Timestamp  `protobuf:"bytes,3,opt,name=date_commenced,json=dateCommenced,proto3" json:"date_commenced,omitempty"`
	DateToSuspend *timestamppb.Timestamp  `protobuf:"bytes,4,opt,name=date_to_suspend,json=dateToSuspend,proto3" json:"date_to_suspend,omitempty"`
	LockApplied   *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=lock_applied,json=lockApplied,proto3" json:"lock_applied,omitempty"`
	Domains       []string                `protobuf:"bytes,6,rep,name=domains,proto3" json:"domains,omitempty"`
}

func (x *DataQuality) Reset() {
	*x = DataQuality{}
	if protoimpl.UnsafeEnabled {
		mi := &file_nominet_ext_nominet_ext_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataQuality) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataQuality) ProtoMessage() {}

func (x *DataQuality) ProtoReflect() protoreflect.Message {
	mi := &file_nominet_ext_nominet_ext_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataQuality.ProtoReflect.Descriptor instead.
func (*DataQuality) Descriptor() ([]byte, []int) {
	return file_nominet_ext_nominet_ext_proto_rawDescGZIP(), []int{4}
}

func (x *DataQuality) GetStatus() DataQualityStatus {
	if x != nil {
		return x.Status
	}
	return DataQualityStatus_Invalid
}

func (x *DataQuality) GetReason() *wrapperspb.StringValue {
	if x != nil {
		return x.Reason
	}
	return nil
}

func (x *DataQuality) GetDateCommenced() *timestamppb.Timestamp {
	if x != nil {
		return x.DateCommenced
	}
	return nil
}

func (x *DataQuality) GetDateToSuspend() *timestamppb.Timestamp {
	if x != nil {
		return x.DateToSuspend
	}
	return nil
}

func (x *DataQuality) GetLockApplied() *wrapperspb.BoolValue {
	if x != nil {
		return x.LockApplied
	}
	return nil
}

func (x *DataQuality) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

var File_nominet_ext_nominet_ext_proto protoreflect.FileDescriptor

var file_nominet_ext_nominet_ext_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2f, 0x6e, 0x6f,
	0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc6, 0x03, 0x0a, 0x0c, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d,
	0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x0a,
	0x72, 0x65, 0x63, 0x75, 0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x19, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65,
	0x78, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63,
	0x75, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x42, 0x69, 0x6c,
	0x6c, 0x12, 0x39, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0b,
	0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52,
	0x0a, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f,
	0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73,
	0x12, 0x38, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0xf8, 0x03, 0x0a, 0x0c, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x38, 0x0a, 0x0a, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x19, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78,
	0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73,
	0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x75, 0x72, 0x5f, 0x62,
	0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x70, 0x70, 0x2e,
	0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x42, 0x69, 0x6c, 0x6c, 0x12,
	0x39, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x39, 0x0a, 0x09, 0x6e, 0x65,
	0x78, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x6e, 0x65, 0x78,
	0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x70, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e,
	0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x6f, 0x50, 0x65,
	0x72, 0x69, 0x6f, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74,
	0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x65, 0x72,
	0x69, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x5f, 0x6e,
	0x6f, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x08,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x08, 0x72,
	0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73,
	0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0x32, 0x0a, 0x0f, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x62, 0x75, 0x73,
	0x65, 0x5f, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x61,
	0x62, 0x75, 0x73, 0x65, 0x4c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0xcc, 0x04, 0x0a, 0x0a, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x54, 0x0a, 0x13, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69,
	0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x12, 0x72, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x38,
	0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74,
	0x5f, 0x65, 0x78, 0x74, 0x2e, 0x42, 0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x66,
	0x69, 0x72, 0x73, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x38, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x75,
	0x72, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x65,
	0x70, 0x70, 0x2e, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x42,
	0x69, 0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x72, 0x65, 0x63, 0x75, 0x72, 0x42, 0x69,
	0x6c, 0x6c, 0x12, 0x39, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x08, 0x61, 0x75, 0x74, 0x6f, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x39, 0x0a,
	0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x62, 0x69, 0x6c, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x42, 0x69, 0x6c, 0x6c, 0x12, 0x3d, 0x0a, 0x0b, 0x61, 0x75, 0x74, 0x6f,
	0x5f, 0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x55, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x61, 0x75, 0x74,
	0x6f, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x3d, 0x0a, 0x0b, 0x6e, 0x65, 0x78, 0x74, 0x5f,
	0x70, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0a, 0x6e, 0x65, 0x78, 0x74,
	0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61,
	0x6c, 0x5f, 0x6e, 0x6f, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x12, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x4e, 0x6f, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65,
	0x73, 0x18, 0x09, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x38,
	0x0a, 0x08, 0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08,
	0x72, 0x65, 0x73, 0x65, 0x6c, 0x6c, 0x65, 0x72, 0x22, 0xdf, 0x02, 0x0a, 0x0b, 0x44, 0x61, 0x74,
	0x61, 0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x65, 0x70, 0x70, 0x2e, 0x6e,
	0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x51,
	0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x72, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x41, 0x0a, 0x0e, 0x64, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0d,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x63, 0x65, 0x64, 0x12, 0x42, 0x0a,
	0x0f, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x6f, 0x5f, 0x73, 0x75, 0x73, 0x70, 0x65, 0x6e, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x53, 0x75, 0x73, 0x70, 0x65, 0x6e,
	0x64, 0x12, 0x3d, 0x0a, 0x0c, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x65,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f, 0x6f, 0x6c, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x52, 0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x65, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2a, 0x40, 0x0a, 0x08, 0x42, 0x69,
	0x6c, 0x6c, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x6e, 0x73, 0x70, 0x65, 0x63,
	0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x42, 0x69, 0x6c, 0x6c, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x72, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x42, 0x69,
	0x6c, 0x6c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x10, 0x02, 0x2a, 0x5a, 0x0a, 0x12,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x65, 0x64,
	0x55, 0x6e, 0x74, 0x69, 0x6c, 0x45, 0x78, 0x70, 0x69, 0x72, 0x79, 0x10, 0x00, 0x12, 0x13, 0x0a,
	0x0f, 0x52, 0x65, 0x6e, 0x65, 0x77, 0x61, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64,
	0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4e, 0x6f, 0x4c, 0x6f, 0x6e, 0x67, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x10, 0x02, 0x2a, 0x2b, 0x0a, 0x11, 0x44, 0x61, 0x74, 0x61,
	0x51, 0x75, 0x61, 0x6c, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x56, 0x61,
	0x6c, 0x69, 0x64, 0x10, 0x01, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x73, 0x32, 0x30, 0x37, 0x39, 0x36, 0x30, 0x2f, 0x65, 0x70, 0x70,
	0x2d, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x65, 0x70,
	0x70, 0x2f, 0x6e, 0x6f, 0x6d, 0x69, 0x6e, 0x65, 0x74, 0x5f, 0x65, 0x78, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_nominet_ext_nominet_ext_proto_rawDescOnce sync.Once
	file_nominet_ext_nominet_ext_proto_rawDescData = file_nominet_ext_nominet_ext_proto_rawDesc
)

func file_nominet_ext_nominet_ext_proto_rawDescGZIP() []byte {
	file_nominet_ext_nominet_ext_proto_rawDescOnce.Do(func() {
		file_nominet_ext_nominet_ext_proto_rawDescData = protoimpl.X.CompressGZIP(file_nominet_ext_nominet_ext_proto_rawDescData)
	})
	return file_nominet_ext_nominet_ext_proto_rawDescData
}

var file_nominet_ext_nominet_ext_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_nominet_ext_nominet_ext_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_nominet_ext_nominet_ext_proto_goTypes = []interface{}{
	(BillType)(0),                  // 0: epp.nominet_ext.BillType
	(RegistrationStatus)(0),        // 1: epp.nominet_ext.RegistrationStatus
	(DataQualityStatus)(0),         // 2: epp.nominet_ext.DataQualityStatus
	(*DomainCreate)(nil),           // 3: epp.nominet_ext.DomainCreate
	(*DomainUpdate)(nil),           // 4: epp.nominet_ext.DomainUpdate
	(*DomainCheckInfo)(nil),        // 5: epp.nominet_ext.DomainCheckInfo
	(*DomainInfo)(nil),             // 6: epp.nominet_ext.DomainInfo
	(*DataQuality)(nil),            // 7: epp.nominet_ext.DataQuality
	(*wrapperspb.UInt32Value)(nil), // 8: google.protobuf.UInt32Value
	(*wrapperspb.StringValue)(nil), // 9: google.protobuf.StringValue
	(*timestamppb.Timestamp)(nil),  // 10: google.protobuf.Timestamp
	(*wrapperspb.BoolValue)(nil),   // 11: google.protobuf.BoolValue
}
var file_nominet_ext_nominet_ext_proto_depIdxs = []int32{
	0,  // 0: epp.nominet_ext.DomainCreate.first_bill:type_name -> epp.nominet_ext.BillType
	0,  // 1: epp.nominet_ext.DomainCreate.recur_bill:type_name -> epp.nominet_ext.BillType
	8,  // 2: epp.nominet_ext.DomainCreate.auto_bill:type_name -> google.protobuf.UInt32Value
	8,  // 3: epp.nominet_ext.DomainCreate.next_bill:type_name -> google.protobuf.UInt32Value
	8,  // 4: epp.nominet_ext.DomainCreate.auto_period:type_name -> google.protobuf.UInt32Value
	8,  // 5: epp.nominet_ext.DomainCreate.next_period:type_name -> google.protobuf.UInt32Value
	9,  // 6: epp.nominet_ext.DomainCreate.reseller:type_name -> google.protobuf.StringValue
	0,  // 7: epp.nominet_ext.DomainUpdate.first_bill:type_name -> epp.nominet_ext.BillType
	0,  // 8: epp.nominet_ext.DomainUpdate.recur_bill:type_name -> epp.nominet_ext.BillType
	8,  // 9: epp.nominet_ext.DomainUpdate.auto_bill:type_name -> google.protobuf.UInt32Value
	8,  // 10: epp.nominet_ext.DomainUpdate.next_bill:type_name -> google.protobuf.UInt32Value
	8,  // 11: epp.nominet_ext.DomainUpdate.auto_period:type_name -> google.protobuf.UInt32Value
	8,  // 12: epp.nominet_ext.DomainUpdate.next_period:type_name -> google.protobuf.UInt32Value
	9,  // 13: epp.nominet_ext.DomainUpdate.reseller:type_name -> google.protobuf.StringValue
	1,  // 14: epp.nominet_ext.DomainInfo.registration_status:type_name -> epp.nominet_ext.RegistrationStatus
	0,  // 15: epp.nominet_ext.DomainInfo.first_bill:type_name -> epp.nominet_ext.BillType
	0,  // 16: epp.nominet_ext.DomainInfo.recur_bill:type_name -> epp.nominet_ext.BillType
	8,  // 17: epp.nominet_ext.DomainInfo.auto_bill:type_name -> google.protobuf.UInt32Value
	8,  // 18: epp.nominet_ext.DomainInfo.next_bill:type_name -> google.protobuf.UInt32Value
	8,  // 19: epp.nominet_ext.DomainInfo.auto_period:type_name -> google.protobuf.UInt32Value
	8,  // 20: epp.nominet_ext.DomainInfo.next_period:type_name -> google.protobuf.UInt32Value
	9,  // 21: epp.nominet_ext.DomainInfo.reseller:type_name -> google.protobuf.StringValue
	2,  // 22: epp.nominet_ext.DataQuality.status:type_name -> epp.nominet_ext.DataQualityStatus
	9,  // 23: epp.nominet_ext.DataQuality.reason:type_name -> google.protobuf.StringValue
	10, // 24: epp.nominet_ext.DataQuality.date_commenced:type_name -> google.protobuf.Timestamp
	10, // 25: epp.nominet_ext.DataQuality.date_to_suspend:type_name -> google.protobuf.Timestamp
	11, // 26: epp.nominet_ext.DataQuality.lock_applied:type_name -> google.protobuf.BoolValue
	27, // [27:27] is the sub-list for method output_type
	27, // [27:27] is the sub-list for method input_type
	27, // [27:27] is the sub-list for extension type_name
	27, // [27:27] is the sub-list for extension extendee
	0,  // [0:27] is the sub-list for field type_name
}

func init() { file_nominet_ext_nominet_ext_proto_init() }
func file_nominet_ext_nominet_ext_proto_init() {
	if File_nominet_ext_nominet_ext_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_nominet_ext_nominet_ext_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCreate); i {
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
		file_nominet_ext_nominet_ext_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_nominet_ext_nominet_ext_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DomainCheckInfo); i {
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
		file_nominet_ext_nominet_ext_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_nominet_ext_nominet_ext_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataQuality); i {
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
			RawDescriptor: file_nominet_ext_nominet_ext_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_nominet_ext_nominet_ext_proto_goTypes,
		DependencyIndexes: file_nominet_ext_nominet_ext_proto_depIdxs,
		EnumInfos:         file_nominet_ext_nominet_ext_proto_enumTypes,
		MessageInfos:      file_nominet_ext_nominet_ext_proto_msgTypes,
	}.Build()
	File_nominet_ext_nominet_ext_proto = out.File
	file_nominet_ext_nominet_ext_proto_rawDesc = nil
	file_nominet_ext_nominet_ext_proto_goTypes = nil
	file_nominet_ext_nominet_ext_proto_depIdxs = nil
}
