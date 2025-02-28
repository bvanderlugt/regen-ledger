// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: regen/ecocredit/basket/v1/state.proto

// Revision 1

package basket

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/api/cosmos/orm/v1alpha1"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Basket represents a basket in state.
type Basket struct {
	// id is the uint64 ID of the basket. It is used internally for reducing
	// storage space.
	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// basket_denom is the basket bank denom formed from name and credit type with
	// the format `eco.<prefix><credit_type_abbrev>.<name>` where prefix is the
	// prefix of the bank denom exponent, a standard SI unit derived from credit
	// type precision, and mapped as follows:
	//   0 - no prefix
	//   1 - d (deci)
	//   2 - c (centi)
	//   3 - m (milli)
	//   6 - u (micro)
	//   9 - n (nano)
	//   12 - p (pico)
	//   15 - f (femto)
	//   18 - a (atto)
	//   21 - z (zepto)
	//   24 - y (yocto)
	BasketDenom string `protobuf:"bytes,2,opt,name=basket_denom,json=basketDenom,proto3" json:"basket_denom,omitempty"`
	// name is the unique name of the basket specified in MsgCreate. Basket
	// names must be unique across all credit types and choices of exponent
	// above and beyond the uniqueness constraint on basket_denom.
	Name string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	// disable_auto_retire indicates whether or not the credits will be retired
	// upon withdraw from the basket.
	DisableAutoRetire bool `protobuf:"varint,4,opt,name=disable_auto_retire,json=disableAutoRetire,proto3" json:"disable_auto_retire,omitempty"`
	// credit_type_abbrev is the abbreviation of the credit type this basket is
	// able to hold.
	CreditTypeAbbrev string `protobuf:"bytes,5,opt,name=credit_type_abbrev,json=creditTypeAbbrev,proto3" json:"credit_type_abbrev,omitempty"`
	// date_criteria is the date criteria for batches admitted to the basket.
	DateCriteria *DateCriteria `protobuf:"bytes,6,opt,name=date_criteria,json=dateCriteria,proto3" json:"date_criteria,omitempty"`
	// Deprecated (Since Revision 1): This field is no longer used and will be
	// removed in the next version. The value of credit type precision is always
	// used as the exponent when converting credits to/from basket tokens. This
	// field will be set to the value of credit type precision until removed.
	Exponent uint32 `protobuf:"varint,7,opt,name=exponent,proto3" json:"exponent,omitempty"` // Deprecated: Do not use.
	// curator is the address of the basket curator who is able to change certain
	// basket settings.
	//
	// Since Revision 1
	Curator []byte `protobuf:"bytes,8,opt,name=curator,proto3" json:"curator,omitempty"`
}

func (m *Basket) Reset()         { *m = Basket{} }
func (m *Basket) String() string { return proto.CompactTextString(m) }
func (*Basket) ProtoMessage()    {}
func (*Basket) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{0}
}
func (m *Basket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Basket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Basket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Basket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Basket.Merge(m, src)
}
func (m *Basket) XXX_Size() int {
	return m.Size()
}
func (m *Basket) XXX_DiscardUnknown() {
	xxx_messageInfo_Basket.DiscardUnknown(m)
}

var xxx_messageInfo_Basket proto.InternalMessageInfo

func (m *Basket) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Basket) GetBasketDenom() string {
	if m != nil {
		return m.BasketDenom
	}
	return ""
}

func (m *Basket) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Basket) GetDisableAutoRetire() bool {
	if m != nil {
		return m.DisableAutoRetire
	}
	return false
}

func (m *Basket) GetCreditTypeAbbrev() string {
	if m != nil {
		return m.CreditTypeAbbrev
	}
	return ""
}

func (m *Basket) GetDateCriteria() *DateCriteria {
	if m != nil {
		return m.DateCriteria
	}
	return nil
}

// Deprecated: Do not use.
func (m *Basket) GetExponent() uint32 {
	if m != nil {
		return m.Exponent
	}
	return 0
}

func (m *Basket) GetCurator() []byte {
	if m != nil {
		return m.Curator
	}
	return nil
}

// BasketClass describes a credit class that can be deposited in a basket.
type BasketClass struct {
	// basket_id is the ID of the basket
	BasketId uint64 `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	// class_id is the id of the credit class that is allowed to be deposited in
	// the basket
	ClassId string `protobuf:"bytes,2,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
}

func (m *BasketClass) Reset()         { *m = BasketClass{} }
func (m *BasketClass) String() string { return proto.CompactTextString(m) }
func (*BasketClass) ProtoMessage()    {}
func (*BasketClass) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{1}
}
func (m *BasketClass) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketClass) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketClass.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketClass) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketClass.Merge(m, src)
}
func (m *BasketClass) XXX_Size() int {
	return m.Size()
}
func (m *BasketClass) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketClass.DiscardUnknown(m)
}

var xxx_messageInfo_BasketClass proto.InternalMessageInfo

func (m *BasketClass) GetBasketId() uint64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

func (m *BasketClass) GetClassId() string {
	if m != nil {
		return m.ClassId
	}
	return ""
}

// BasketBalance stores the amount of credits from a batch in a basket
type BasketBalance struct {
	// basket_id is the ID of the basket
	BasketId uint64 `protobuf:"varint,1,opt,name=basket_id,json=basketId,proto3" json:"basket_id,omitempty"`
	// batch_denom is the denom of the credit batch
	BatchDenom string `protobuf:"bytes,2,opt,name=batch_denom,json=batchDenom,proto3" json:"batch_denom,omitempty"`
	// balance is the amount of ecocredits held in the basket
	Balance string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
	// batch_start_date is the start date of the batch. This field is used
	// to create an index which is used to remove the oldest credits first.
	BatchStartDate *types.Timestamp `protobuf:"bytes,4,opt,name=batch_start_date,json=batchStartDate,proto3" json:"batch_start_date,omitempty"`
}

func (m *BasketBalance) Reset()         { *m = BasketBalance{} }
func (m *BasketBalance) String() string { return proto.CompactTextString(m) }
func (*BasketBalance) ProtoMessage()    {}
func (*BasketBalance) Descriptor() ([]byte, []int) {
	return fileDescriptor_c416a19075224f85, []int{2}
}
func (m *BasketBalance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BasketBalance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BasketBalance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BasketBalance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BasketBalance.Merge(m, src)
}
func (m *BasketBalance) XXX_Size() int {
	return m.Size()
}
func (m *BasketBalance) XXX_DiscardUnknown() {
	xxx_messageInfo_BasketBalance.DiscardUnknown(m)
}

var xxx_messageInfo_BasketBalance proto.InternalMessageInfo

func (m *BasketBalance) GetBasketId() uint64 {
	if m != nil {
		return m.BasketId
	}
	return 0
}

func (m *BasketBalance) GetBatchDenom() string {
	if m != nil {
		return m.BatchDenom
	}
	return ""
}

func (m *BasketBalance) GetBalance() string {
	if m != nil {
		return m.Balance
	}
	return ""
}

func (m *BasketBalance) GetBatchStartDate() *types.Timestamp {
	if m != nil {
		return m.BatchStartDate
	}
	return nil
}

func init() {
	proto.RegisterType((*Basket)(nil), "regen.ecocredit.basket.v1.Basket")
	proto.RegisterType((*BasketClass)(nil), "regen.ecocredit.basket.v1.BasketClass")
	proto.RegisterType((*BasketBalance)(nil), "regen.ecocredit.basket.v1.BasketBalance")
}

func init() {
	proto.RegisterFile("regen/ecocredit/basket/v1/state.proto", fileDescriptor_c416a19075224f85)
}

var fileDescriptor_c416a19075224f85 = []byte{
	// 567 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xed, 0xba, 0x25, 0x49, 0x27, 0x6d, 0x65, 0x16, 0x10, 0xdb, 0x20, 0x5c, 0x53, 0x09, 0x61,
	0xa1, 0x62, 0x93, 0x72, 0x41, 0xe5, 0xd4, 0xb4, 0x97, 0x4a, 0x9c, 0x4c, 0x4f, 0x5c, 0xac, 0xb5,
	0x3d, 0xa4, 0x56, 0x6d, 0xaf, 0xb5, 0xde, 0x84, 0xf6, 0x27, 0x10, 0x5f, 0xc0, 0xf7, 0x70, 0xac,
	0xc4, 0x85, 0x23, 0x6a, 0x0f, 0x5c, 0x11, 0x5f, 0x80, 0xbc, 0xeb, 0x24, 0x2d, 0x88, 0xde, 0xf2,
	0xe6, 0xbd, 0xe7, 0x99, 0x7d, 0x33, 0x81, 0xa7, 0x12, 0xc7, 0x58, 0x06, 0x98, 0x88, 0x44, 0x62,
	0x9a, 0xa9, 0x20, 0xe6, 0xf5, 0x29, 0xaa, 0x60, 0x3a, 0x0c, 0x6a, 0xc5, 0x15, 0xfa, 0x95, 0x14,
	0x4a, 0xd0, 0x4d, 0x2d, 0xf3, 0xe7, 0x32, 0xdf, 0xc8, 0xfc, 0xe9, 0x70, 0xf0, 0x38, 0x11, 0x75,
	0x21, 0xea, 0x40, 0xc8, 0x22, 0x98, 0x0e, 0x79, 0x5e, 0x9d, 0xf0, 0x61, 0x03, 0x8c, 0x73, 0xb0,
	0x35, 0x16, 0x62, 0x9c, 0x63, 0xa0, 0x51, 0x3c, 0xf9, 0x10, 0xa8, 0xac, 0xc0, 0x5a, 0xf1, 0xa2,
	0x6a, 0x05, 0xb7, 0x4c, 0xa0, 0xce, 0x2b, 0xac, 0x8d, 0x6c, 0xfb, 0xa7, 0x05, 0x9d, 0x91, 0x66,
	0xe8, 0x06, 0x58, 0x59, 0xca, 0x88, 0x4b, 0xbc, 0x95, 0xd0, 0xca, 0x52, 0xfa, 0x04, 0xd6, 0x8c,
	0x27, 0x4a, 0xb1, 0x14, 0x05, 0xb3, 0x5c, 0xe2, 0xad, 0x86, 0x7d, 0x53, 0x3b, 0x6c, 0x4a, 0x94,
	0xc2, 0x4a, 0xc9, 0x0b, 0x64, 0xcb, 0x9a, 0xd2, 0xbf, 0xa9, 0x0f, 0xf7, 0xd2, 0xac, 0xe6, 0x71,
	0x8e, 0x11, 0x9f, 0x28, 0x11, 0x49, 0x54, 0x99, 0x44, 0xb6, 0xe2, 0x12, 0xaf, 0x17, 0xde, 0x6d,
	0xa9, 0xfd, 0x89, 0x12, 0xa1, 0x26, 0xe8, 0x0e, 0x50, 0x33, 0x61, 0xd4, 0xcc, 0x15, 0xf1, 0x38,
	0x96, 0x38, 0x65, 0x77, 0xf4, 0x17, 0x6d, 0xc3, 0x1c, 0x9f, 0x57, 0xb8, 0xaf, 0xeb, 0xf4, 0x2d,
	0xac, 0xa7, 0x5c, 0x61, 0x94, 0xc8, 0x4c, 0xa1, 0xcc, 0x38, 0xeb, 0xb8, 0xc4, 0xeb, 0xef, 0x3e,
	0xf3, 0xff, 0x9b, 0xa4, 0x7f, 0xc8, 0x15, 0x1e, 0xb4, 0xf2, 0x70, 0x2d, 0xbd, 0x86, 0xa8, 0x03,
	0x3d, 0x3c, 0xab, 0x44, 0x89, 0xa5, 0x62, 0x5d, 0x97, 0x78, 0xeb, 0x23, 0x8b, 0x91, 0x70, 0x5e,
	0xa3, 0x0c, 0xba, 0xc9, 0x44, 0x72, 0x25, 0x24, 0xeb, 0xb9, 0xc4, 0x5b, 0x0b, 0x67, 0x70, 0xef,
	0xe5, 0xef, 0x2f, 0xdf, 0x3e, 0x2d, 0x3f, 0x87, 0x4e, 0x13, 0x9a, 0x4d, 0x28, 0xbd, 0x19, 0x96,
	0x4d, 0x18, 0xa1, 0x60, 0xd2, 0xb1, 0x2d, 0x46, 0x18, 0xd9, 0x46, 0xe8, 0x9b, 0xa0, 0x0f, 0x72,
	0x5e, 0xd7, 0xf4, 0x11, 0xac, 0xb6, 0x86, 0x79, 0xe8, 0x3d, 0x53, 0x38, 0x4a, 0xe9, 0x26, 0xf4,
	0x92, 0x46, 0xd5, 0x70, 0x26, 0xf6, 0xae, 0xc6, 0x47, 0xe9, 0x9e, 0xa3, 0x1b, 0x33, 0xb8, 0x0f,
	0x74, 0xee, 0xdf, 0x59, 0x88, 0xb7, 0x7f, 0x11, 0x58, 0x37, 0x7d, 0x46, 0x3c, 0xe7, 0x65, 0x82,
	0xb7, 0x77, 0xda, 0x82, 0x7e, 0xcc, 0x55, 0x72, 0x72, 0x63, 0xc7, 0xa0, 0x4b, 0x66, 0xc5, 0x0c,
	0xba, 0xb1, 0xf9, 0x50, 0xbb, 0xe5, 0x19, 0xa4, 0x87, 0x60, 0x1b, 0x6b, 0xad, 0xb8, 0x54, 0x51,
	0x13, 0xac, 0xde, 0x72, 0x7f, 0x77, 0xe0, 0x9b, 0xeb, 0xf4, 0x67, 0xd7, 0xe9, 0x1f, 0xcf, 0xae,
	0x33, 0xdc, 0xd0, 0x9e, 0x77, 0x8d, 0xa5, 0x59, 0xcc, 0xde, 0xbe, 0x7e, 0xcf, 0x1b, 0x78, 0x08,
	0x0f, 0x16, 0xef, 0xb9, 0x36, 0x12, 0x75, 0x60, 0xf0, 0x37, 0xb1, 0x68, 0x68, 0x13, 0xb6, 0x3c,
	0x0a, 0xbf, 0x5e, 0x3a, 0xe4, 0xe2, 0xd2, 0x21, 0x3f, 0x2e, 0x1d, 0xf2, 0xf9, 0xca, 0x59, 0xba,
	0xb8, 0x72, 0x96, 0xbe, 0x5f, 0x39, 0x4b, 0xef, 0x5f, 0x8f, 0x33, 0x75, 0x32, 0x89, 0xfd, 0x44,
	0x14, 0x81, 0x3e, 0x90, 0x17, 0x25, 0xaa, 0x8f, 0x42, 0x9e, 0xb6, 0x28, 0xc7, 0x74, 0x8c, 0x32,
	0x38, 0xfb, 0xe7, 0x6f, 0x12, 0x77, 0xf4, 0xe8, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x51,
	0xb2, 0x58, 0xb2, 0xc9, 0x03, 0x00, 0x00,
}

func (m *Basket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Basket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Basket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Curator) > 0 {
		i -= len(m.Curator)
		copy(dAtA[i:], m.Curator)
		i = encodeVarintState(dAtA, i, uint64(len(m.Curator)))
		i--
		dAtA[i] = 0x42
	}
	if m.Exponent != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Exponent))
		i--
		dAtA[i] = 0x38
	}
	if m.DateCriteria != nil {
		{
			size, err := m.DateCriteria.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.CreditTypeAbbrev) > 0 {
		i -= len(m.CreditTypeAbbrev)
		copy(dAtA[i:], m.CreditTypeAbbrev)
		i = encodeVarintState(dAtA, i, uint64(len(m.CreditTypeAbbrev)))
		i--
		dAtA[i] = 0x2a
	}
	if m.DisableAutoRetire {
		i--
		if m.DisableAutoRetire {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintState(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BasketDenom) > 0 {
		i -= len(m.BasketDenom)
		copy(dAtA[i:], m.BasketDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BasketDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BasketClass) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketClass) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketClass) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClassId) > 0 {
		i -= len(m.ClassId)
		copy(dAtA[i:], m.ClassId)
		i = encodeVarintState(dAtA, i, uint64(len(m.ClassId)))
		i--
		dAtA[i] = 0x12
	}
	if m.BasketId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BasketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BasketBalance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BasketBalance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BasketBalance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BatchStartDate != nil {
		{
			size, err := m.BatchStartDate.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintState(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Balance) > 0 {
		i -= len(m.Balance)
		copy(dAtA[i:], m.Balance)
		i = encodeVarintState(dAtA, i, uint64(len(m.Balance)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.BatchDenom) > 0 {
		i -= len(m.BatchDenom)
		copy(dAtA[i:], m.BatchDenom)
		i = encodeVarintState(dAtA, i, uint64(len(m.BatchDenom)))
		i--
		dAtA[i] = 0x12
	}
	if m.BasketId != 0 {
		i = encodeVarintState(dAtA, i, uint64(m.BasketId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	offset -= sovState(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Basket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovState(uint64(m.Id))
	}
	l = len(m.BasketDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.DisableAutoRetire {
		n += 2
	}
	l = len(m.CreditTypeAbbrev)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.DateCriteria != nil {
		l = m.DateCriteria.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Exponent != 0 {
		n += 1 + sovState(uint64(m.Exponent))
	}
	l = len(m.Curator)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *BasketClass) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasketId != 0 {
		n += 1 + sovState(uint64(m.BasketId))
	}
	l = len(m.ClassId)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func (m *BasketBalance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasketId != 0 {
		n += 1 + sovState(uint64(m.BasketId))
	}
	l = len(m.BatchDenom)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	l = len(m.Balance)
	if l > 0 {
		n += 1 + l + sovState(uint64(l))
	}
	if m.BatchStartDate != nil {
		l = m.BatchStartDate.Size()
		n += 1 + l + sovState(uint64(l))
	}
	return n
}

func sovState(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Basket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: Basket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Basket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BasketDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisableAutoRetire", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.DisableAutoRetire = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreditTypeAbbrev", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CreditTypeAbbrev = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DateCriteria", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DateCriteria == nil {
				m.DateCriteria = &DateCriteria{}
			}
			if err := m.DateCriteria.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Exponent", wireType)
			}
			m.Exponent = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Exponent |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Curator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Curator = append(m.Curator[:0], dAtA[iNdEx:postIndex]...)
			if m.Curator == nil {
				m.Curator = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *BasketClass) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: BasketClass: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketClass: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketId", wireType)
			}
			m.BasketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClassId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func (m *BasketBalance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: BasketBalance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BasketBalance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasketId", wireType)
			}
			m.BasketId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BasketId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BatchDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Balance", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Balance = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BatchStartDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthState
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BatchStartDate == nil {
				m.BatchStartDate = &types.Timestamp{}
			}
			if err := m.BatchStartDate.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupState
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthState
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthState        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupState = fmt.Errorf("proto: unexpected end of group")
)
