// Code generated by protoc-gen-go. DO NOT EDIT.
// source: piece_store.proto

package piecestoreroutes

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type PayerBandwidthAllocation struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayerBandwidthAllocation) Reset()         { *m = PayerBandwidthAllocation{} }
func (m *PayerBandwidthAllocation) String() string { return proto.CompactTextString(m) }
func (*PayerBandwidthAllocation) ProtoMessage()    {}
func (*PayerBandwidthAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{0}
}
func (m *PayerBandwidthAllocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayerBandwidthAllocation.Unmarshal(m, b)
}
func (m *PayerBandwidthAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayerBandwidthAllocation.Marshal(b, m, deterministic)
}
func (dst *PayerBandwidthAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayerBandwidthAllocation.Merge(dst, src)
}
func (m *PayerBandwidthAllocation) XXX_Size() int {
	return xxx_messageInfo_PayerBandwidthAllocation.Size(m)
}
func (m *PayerBandwidthAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_PayerBandwidthAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_PayerBandwidthAllocation proto.InternalMessageInfo

func (m *PayerBandwidthAllocation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *PayerBandwidthAllocation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type PayerBandwidthAllocation_Data struct {
	Payer                []byte   `protobuf:"bytes,1,opt,name=payer,proto3" json:"payer,omitempty"`
	Renter               []byte   `protobuf:"bytes,2,opt,name=renter,proto3" json:"renter,omitempty"`
	MaxSize              int64    `protobuf:"varint,3,opt,name=max_size,json=maxSize" json:"max_size,omitempty"`
	Expiration           int64    `protobuf:"varint,4,opt,name=expiration" json:"expiration,omitempty"`
	SerialNumber         string   `protobuf:"bytes,5,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PayerBandwidthAllocation_Data) Reset()         { *m = PayerBandwidthAllocation_Data{} }
func (m *PayerBandwidthAllocation_Data) String() string { return proto.CompactTextString(m) }
func (*PayerBandwidthAllocation_Data) ProtoMessage()    {}
func (*PayerBandwidthAllocation_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{0, 0}
}
func (m *PayerBandwidthAllocation_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PayerBandwidthAllocation_Data.Unmarshal(m, b)
}
func (m *PayerBandwidthAllocation_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PayerBandwidthAllocation_Data.Marshal(b, m, deterministic)
}
func (dst *PayerBandwidthAllocation_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PayerBandwidthAllocation_Data.Merge(dst, src)
}
func (m *PayerBandwidthAllocation_Data) XXX_Size() int {
	return xxx_messageInfo_PayerBandwidthAllocation_Data.Size(m)
}
func (m *PayerBandwidthAllocation_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_PayerBandwidthAllocation_Data.DiscardUnknown(m)
}

var xxx_messageInfo_PayerBandwidthAllocation_Data proto.InternalMessageInfo

func (m *PayerBandwidthAllocation_Data) GetPayer() []byte {
	if m != nil {
		return m.Payer
	}
	return nil
}

func (m *PayerBandwidthAllocation_Data) GetRenter() []byte {
	if m != nil {
		return m.Renter
	}
	return nil
}

func (m *PayerBandwidthAllocation_Data) GetMaxSize() int64 {
	if m != nil {
		return m.MaxSize
	}
	return 0
}

func (m *PayerBandwidthAllocation_Data) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

func (m *PayerBandwidthAllocation_Data) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

type RenterBandwidthAllocation struct {
	Signature            []byte   `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RenterBandwidthAllocation) Reset()         { *m = RenterBandwidthAllocation{} }
func (m *RenterBandwidthAllocation) String() string { return proto.CompactTextString(m) }
func (*RenterBandwidthAllocation) ProtoMessage()    {}
func (*RenterBandwidthAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{1}
}
func (m *RenterBandwidthAllocation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenterBandwidthAllocation.Unmarshal(m, b)
}
func (m *RenterBandwidthAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenterBandwidthAllocation.Marshal(b, m, deterministic)
}
func (dst *RenterBandwidthAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenterBandwidthAllocation.Merge(dst, src)
}
func (m *RenterBandwidthAllocation) XXX_Size() int {
	return xxx_messageInfo_RenterBandwidthAllocation.Size(m)
}
func (m *RenterBandwidthAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_RenterBandwidthAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_RenterBandwidthAllocation proto.InternalMessageInfo

func (m *RenterBandwidthAllocation) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *RenterBandwidthAllocation) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type RenterBandwidthAllocation_Data struct {
	PayerAllocation      *PayerBandwidthAllocation `protobuf:"bytes,1,opt,name=payer_allocation,json=payerAllocation" json:"payer_allocation,omitempty"`
	Size                 int64                     `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Total                int64                     `protobuf:"varint,3,opt,name=total" json:"total,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *RenterBandwidthAllocation_Data) Reset()         { *m = RenterBandwidthAllocation_Data{} }
func (m *RenterBandwidthAllocation_Data) String() string { return proto.CompactTextString(m) }
func (*RenterBandwidthAllocation_Data) ProtoMessage()    {}
func (*RenterBandwidthAllocation_Data) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{1, 0}
}
func (m *RenterBandwidthAllocation_Data) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RenterBandwidthAllocation_Data.Unmarshal(m, b)
}
func (m *RenterBandwidthAllocation_Data) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RenterBandwidthAllocation_Data.Marshal(b, m, deterministic)
}
func (dst *RenterBandwidthAllocation_Data) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RenterBandwidthAllocation_Data.Merge(dst, src)
}
func (m *RenterBandwidthAllocation_Data) XXX_Size() int {
	return xxx_messageInfo_RenterBandwidthAllocation_Data.Size(m)
}
func (m *RenterBandwidthAllocation_Data) XXX_DiscardUnknown() {
	xxx_messageInfo_RenterBandwidthAllocation_Data.DiscardUnknown(m)
}

var xxx_messageInfo_RenterBandwidthAllocation_Data proto.InternalMessageInfo

func (m *RenterBandwidthAllocation_Data) GetPayerAllocation() *PayerBandwidthAllocation {
	if m != nil {
		return m.PayerAllocation
	}
	return nil
}

func (m *RenterBandwidthAllocation_Data) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *RenterBandwidthAllocation_Data) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

type PieceStore struct {
	Bandwidthallocation  *RenterBandwidthAllocation `protobuf:"bytes,1,opt,name=bandwidthallocation" json:"bandwidthallocation,omitempty"`
	Piecedata            *PieceStore_PieceData      `protobuf:"bytes,2,opt,name=piecedata" json:"piecedata,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PieceStore) Reset()         { *m = PieceStore{} }
func (m *PieceStore) String() string { return proto.CompactTextString(m) }
func (*PieceStore) ProtoMessage()    {}
func (*PieceStore) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{2}
}
func (m *PieceStore) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceStore.Unmarshal(m, b)
}
func (m *PieceStore) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceStore.Marshal(b, m, deterministic)
}
func (dst *PieceStore) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceStore.Merge(dst, src)
}
func (m *PieceStore) XXX_Size() int {
	return xxx_messageInfo_PieceStore.Size(m)
}
func (m *PieceStore) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceStore.DiscardUnknown(m)
}

var xxx_messageInfo_PieceStore proto.InternalMessageInfo

func (m *PieceStore) GetBandwidthallocation() *RenterBandwidthAllocation {
	if m != nil {
		return m.Bandwidthallocation
	}
	return nil
}

func (m *PieceStore) GetPiecedata() *PieceStore_PieceData {
	if m != nil {
		return m.Piecedata
	}
	return nil
}

type PieceStore_PieceData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Ttl                  int64    `protobuf:"varint,2,opt,name=ttl" json:"ttl,omitempty"`
	Content              []byte   `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceStore_PieceData) Reset()         { *m = PieceStore_PieceData{} }
func (m *PieceStore_PieceData) String() string { return proto.CompactTextString(m) }
func (*PieceStore_PieceData) ProtoMessage()    {}
func (*PieceStore_PieceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{2, 0}
}
func (m *PieceStore_PieceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceStore_PieceData.Unmarshal(m, b)
}
func (m *PieceStore_PieceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceStore_PieceData.Marshal(b, m, deterministic)
}
func (dst *PieceStore_PieceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceStore_PieceData.Merge(dst, src)
}
func (m *PieceStore_PieceData) XXX_Size() int {
	return xxx_messageInfo_PieceStore_PieceData.Size(m)
}
func (m *PieceStore_PieceData) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceStore_PieceData.DiscardUnknown(m)
}

var xxx_messageInfo_PieceStore_PieceData proto.InternalMessageInfo

func (m *PieceStore_PieceData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceStore_PieceData) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *PieceStore_PieceData) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceId struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceId) Reset()         { *m = PieceId{} }
func (m *PieceId) String() string { return proto.CompactTextString(m) }
func (*PieceId) ProtoMessage()    {}
func (*PieceId) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{3}
}
func (m *PieceId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceId.Unmarshal(m, b)
}
func (m *PieceId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceId.Marshal(b, m, deterministic)
}
func (dst *PieceId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceId.Merge(dst, src)
}
func (m *PieceId) XXX_Size() int {
	return xxx_messageInfo_PieceId.Size(m)
}
func (m *PieceId) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceId.DiscardUnknown(m)
}

var xxx_messageInfo_PieceId proto.InternalMessageInfo

func (m *PieceId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PieceSummary struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Expiration           int64    `protobuf:"varint,3,opt,name=expiration" json:"expiration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceSummary) Reset()         { *m = PieceSummary{} }
func (m *PieceSummary) String() string { return proto.CompactTextString(m) }
func (*PieceSummary) ProtoMessage()    {}
func (*PieceSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{4}
}
func (m *PieceSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceSummary.Unmarshal(m, b)
}
func (m *PieceSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceSummary.Marshal(b, m, deterministic)
}
func (dst *PieceSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceSummary.Merge(dst, src)
}
func (m *PieceSummary) XXX_Size() int {
	return xxx_messageInfo_PieceSummary.Size(m)
}
func (m *PieceSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceSummary proto.InternalMessageInfo

func (m *PieceSummary) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceSummary) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceSummary) GetExpiration() int64 {
	if m != nil {
		return m.Expiration
	}
	return 0
}

type PieceRetrieval struct {
	Bandwidthallocation  *RenterBandwidthAllocation `protobuf:"bytes,1,opt,name=bandwidthallocation" json:"bandwidthallocation,omitempty"`
	PieceData            *PieceRetrieval_PieceData  `protobuf:"bytes,2,opt,name=pieceData" json:"pieceData,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *PieceRetrieval) Reset()         { *m = PieceRetrieval{} }
func (m *PieceRetrieval) String() string { return proto.CompactTextString(m) }
func (*PieceRetrieval) ProtoMessage()    {}
func (*PieceRetrieval) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{5}
}
func (m *PieceRetrieval) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceRetrieval.Unmarshal(m, b)
}
func (m *PieceRetrieval) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceRetrieval.Marshal(b, m, deterministic)
}
func (dst *PieceRetrieval) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceRetrieval.Merge(dst, src)
}
func (m *PieceRetrieval) XXX_Size() int {
	return xxx_messageInfo_PieceRetrieval.Size(m)
}
func (m *PieceRetrieval) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceRetrieval.DiscardUnknown(m)
}

var xxx_messageInfo_PieceRetrieval proto.InternalMessageInfo

func (m *PieceRetrieval) GetBandwidthallocation() *RenterBandwidthAllocation {
	if m != nil {
		return m.Bandwidthallocation
	}
	return nil
}

func (m *PieceRetrieval) GetPieceData() *PieceRetrieval_PieceData {
	if m != nil {
		return m.PieceData
	}
	return nil
}

type PieceRetrieval_PieceData struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Size                 int64    `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	Offset               int64    `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceRetrieval_PieceData) Reset()         { *m = PieceRetrieval_PieceData{} }
func (m *PieceRetrieval_PieceData) String() string { return proto.CompactTextString(m) }
func (*PieceRetrieval_PieceData) ProtoMessage()    {}
func (*PieceRetrieval_PieceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{5, 0}
}
func (m *PieceRetrieval_PieceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceRetrieval_PieceData.Unmarshal(m, b)
}
func (m *PieceRetrieval_PieceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceRetrieval_PieceData.Marshal(b, m, deterministic)
}
func (dst *PieceRetrieval_PieceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceRetrieval_PieceData.Merge(dst, src)
}
func (m *PieceRetrieval_PieceData) XXX_Size() int {
	return xxx_messageInfo_PieceRetrieval_PieceData.Size(m)
}
func (m *PieceRetrieval_PieceData) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceRetrieval_PieceData.DiscardUnknown(m)
}

var xxx_messageInfo_PieceRetrieval_PieceData proto.InternalMessageInfo

func (m *PieceRetrieval_PieceData) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *PieceRetrieval_PieceData) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceRetrieval_PieceData) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type PieceRetrievalStream struct {
	Size                 int64    `protobuf:"varint,1,opt,name=size" json:"size,omitempty"`
	Content              []byte   `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceRetrievalStream) Reset()         { *m = PieceRetrievalStream{} }
func (m *PieceRetrievalStream) String() string { return proto.CompactTextString(m) }
func (*PieceRetrievalStream) ProtoMessage()    {}
func (*PieceRetrievalStream) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{6}
}
func (m *PieceRetrievalStream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceRetrievalStream.Unmarshal(m, b)
}
func (m *PieceRetrievalStream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceRetrievalStream.Marshal(b, m, deterministic)
}
func (dst *PieceRetrievalStream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceRetrievalStream.Merge(dst, src)
}
func (m *PieceRetrievalStream) XXX_Size() int {
	return xxx_messageInfo_PieceRetrievalStream.Size(m)
}
func (m *PieceRetrievalStream) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceRetrievalStream.DiscardUnknown(m)
}

var xxx_messageInfo_PieceRetrievalStream proto.InternalMessageInfo

func (m *PieceRetrievalStream) GetSize() int64 {
	if m != nil {
		return m.Size
	}
	return 0
}

func (m *PieceRetrievalStream) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type PieceDelete struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceDelete) Reset()         { *m = PieceDelete{} }
func (m *PieceDelete) String() string { return proto.CompactTextString(m) }
func (*PieceDelete) ProtoMessage()    {}
func (*PieceDelete) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{7}
}
func (m *PieceDelete) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceDelete.Unmarshal(m, b)
}
func (m *PieceDelete) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceDelete.Marshal(b, m, deterministic)
}
func (dst *PieceDelete) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceDelete.Merge(dst, src)
}
func (m *PieceDelete) XXX_Size() int {
	return xxx_messageInfo_PieceDelete.Size(m)
}
func (m *PieceDelete) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceDelete.DiscardUnknown(m)
}

var xxx_messageInfo_PieceDelete proto.InternalMessageInfo

func (m *PieceDelete) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type PieceDeleteSummary struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceDeleteSummary) Reset()         { *m = PieceDeleteSummary{} }
func (m *PieceDeleteSummary) String() string { return proto.CompactTextString(m) }
func (*PieceDeleteSummary) ProtoMessage()    {}
func (*PieceDeleteSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{8}
}
func (m *PieceDeleteSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceDeleteSummary.Unmarshal(m, b)
}
func (m *PieceDeleteSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceDeleteSummary.Marshal(b, m, deterministic)
}
func (dst *PieceDeleteSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceDeleteSummary.Merge(dst, src)
}
func (m *PieceDeleteSummary) XXX_Size() int {
	return xxx_messageInfo_PieceDeleteSummary.Size(m)
}
func (m *PieceDeleteSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceDeleteSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceDeleteSummary proto.InternalMessageInfo

func (m *PieceDeleteSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PieceStoreSummary struct {
	Message              string   `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	TotalReceived        int64    `protobuf:"varint,2,opt,name=totalReceived" json:"totalReceived,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PieceStoreSummary) Reset()         { *m = PieceStoreSummary{} }
func (m *PieceStoreSummary) String() string { return proto.CompactTextString(m) }
func (*PieceStoreSummary) ProtoMessage()    {}
func (*PieceStoreSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_piece_store_fd7b526146f7ffd0, []int{9}
}
func (m *PieceStoreSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PieceStoreSummary.Unmarshal(m, b)
}
func (m *PieceStoreSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PieceStoreSummary.Marshal(b, m, deterministic)
}
func (dst *PieceStoreSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PieceStoreSummary.Merge(dst, src)
}
func (m *PieceStoreSummary) XXX_Size() int {
	return xxx_messageInfo_PieceStoreSummary.Size(m)
}
func (m *PieceStoreSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_PieceStoreSummary.DiscardUnknown(m)
}

var xxx_messageInfo_PieceStoreSummary proto.InternalMessageInfo

func (m *PieceStoreSummary) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PieceStoreSummary) GetTotalReceived() int64 {
	if m != nil {
		return m.TotalReceived
	}
	return 0
}

func init() {
	proto.RegisterType((*PayerBandwidthAllocation)(nil), "piecestoreroutes.PayerBandwidthAllocation")
	proto.RegisterType((*PayerBandwidthAllocation_Data)(nil), "piecestoreroutes.PayerBandwidthAllocation.Data")
	proto.RegisterType((*RenterBandwidthAllocation)(nil), "piecestoreroutes.RenterBandwidthAllocation")
	proto.RegisterType((*RenterBandwidthAllocation_Data)(nil), "piecestoreroutes.RenterBandwidthAllocation.Data")
	proto.RegisterType((*PieceStore)(nil), "piecestoreroutes.PieceStore")
	proto.RegisterType((*PieceStore_PieceData)(nil), "piecestoreroutes.PieceStore.PieceData")
	proto.RegisterType((*PieceId)(nil), "piecestoreroutes.PieceId")
	proto.RegisterType((*PieceSummary)(nil), "piecestoreroutes.PieceSummary")
	proto.RegisterType((*PieceRetrieval)(nil), "piecestoreroutes.PieceRetrieval")
	proto.RegisterType((*PieceRetrieval_PieceData)(nil), "piecestoreroutes.PieceRetrieval.PieceData")
	proto.RegisterType((*PieceRetrievalStream)(nil), "piecestoreroutes.PieceRetrievalStream")
	proto.RegisterType((*PieceDelete)(nil), "piecestoreroutes.PieceDelete")
	proto.RegisterType((*PieceDeleteSummary)(nil), "piecestoreroutes.PieceDeleteSummary")
	proto.RegisterType((*PieceStoreSummary)(nil), "piecestoreroutes.PieceStoreSummary")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for PieceStoreRoutes service

type PieceStoreRoutesClient interface {
	Piece(ctx context.Context, in *PieceId, opts ...grpc.CallOption) (*PieceSummary, error)
	Retrieve(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error)
	Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error)
	Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error)
}

type pieceStoreRoutesClient struct {
	cc *grpc.ClientConn
}

func NewPieceStoreRoutesClient(cc *grpc.ClientConn) PieceStoreRoutesClient {
	return &pieceStoreRoutesClient{cc}
}

func (c *pieceStoreRoutesClient) Piece(ctx context.Context, in *PieceId, opts ...grpc.CallOption) (*PieceSummary, error) {
	out := new(PieceSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Piece", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *pieceStoreRoutesClient) Retrieve(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_RetrieveClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[0], c.cc, "/piecestoreroutes.PieceStoreRoutes/Retrieve", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesRetrieveClient{stream}
	return x, nil
}

type PieceStoreRoutes_RetrieveClient interface {
	Send(*PieceRetrieval) error
	Recv() (*PieceRetrievalStream, error)
	grpc.ClientStream
}

type pieceStoreRoutesRetrieveClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesRetrieveClient) Send(m *PieceRetrieval) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pieceStoreRoutesRetrieveClient) Recv() (*PieceRetrievalStream, error) {
	m := new(PieceRetrievalStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Store(ctx context.Context, opts ...grpc.CallOption) (PieceStoreRoutes_StoreClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_PieceStoreRoutes_serviceDesc.Streams[1], c.cc, "/piecestoreroutes.PieceStoreRoutes/Store", opts...)
	if err != nil {
		return nil, err
	}
	x := &pieceStoreRoutesStoreClient{stream}
	return x, nil
}

type PieceStoreRoutes_StoreClient interface {
	Send(*PieceStore) error
	CloseAndRecv() (*PieceStoreSummary, error)
	grpc.ClientStream
}

type pieceStoreRoutesStoreClient struct {
	grpc.ClientStream
}

func (x *pieceStoreRoutesStoreClient) Send(m *PieceStore) error {
	return x.ClientStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreClient) CloseAndRecv() (*PieceStoreSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(PieceStoreSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *pieceStoreRoutesClient) Delete(ctx context.Context, in *PieceDelete, opts ...grpc.CallOption) (*PieceDeleteSummary, error) {
	out := new(PieceDeleteSummary)
	err := grpc.Invoke(ctx, "/piecestoreroutes.PieceStoreRoutes/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PieceStoreRoutes service

type PieceStoreRoutesServer interface {
	Piece(context.Context, *PieceId) (*PieceSummary, error)
	Retrieve(PieceStoreRoutes_RetrieveServer) error
	Store(PieceStoreRoutes_StoreServer) error
	Delete(context.Context, *PieceDelete) (*PieceDeleteSummary, error)
}

func RegisterPieceStoreRoutesServer(s *grpc.Server, srv PieceStoreRoutesServer) {
	s.RegisterService(&_PieceStoreRoutes_serviceDesc, srv)
}

func _PieceStoreRoutes_Piece_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Piece(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Piece",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Piece(ctx, req.(*PieceId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PieceStoreRoutes_Retrieve_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PieceStoreRoutesServer).Retrieve(&pieceStoreRoutesRetrieveServer{stream})
}

type PieceStoreRoutes_RetrieveServer interface {
	Send(*PieceRetrievalStream) error
	Recv() (*PieceRetrieval, error)
	grpc.ServerStream
}

type pieceStoreRoutesRetrieveServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesRetrieveServer) Send(m *PieceRetrievalStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pieceStoreRoutesRetrieveServer) Recv() (*PieceRetrieval, error) {
	m := new(PieceRetrieval)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PieceStoreRoutes_Store_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PieceStoreRoutesServer).Store(&pieceStoreRoutesStoreServer{stream})
}

type PieceStoreRoutes_StoreServer interface {
	SendAndClose(*PieceStoreSummary) error
	Recv() (*PieceStore, error)
	grpc.ServerStream
}

type pieceStoreRoutesStoreServer struct {
	grpc.ServerStream
}

func (x *pieceStoreRoutesStoreServer) SendAndClose(m *PieceStoreSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *pieceStoreRoutesStoreServer) Recv() (*PieceStore, error) {
	m := new(PieceStore)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PieceStoreRoutes_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PieceDelete)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PieceStoreRoutesServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/piecestoreroutes.PieceStoreRoutes/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PieceStoreRoutesServer).Delete(ctx, req.(*PieceDelete))
	}
	return interceptor(ctx, in, info, handler)
}

var _PieceStoreRoutes_serviceDesc = grpc.ServiceDesc{
	ServiceName: "piecestoreroutes.PieceStoreRoutes",
	HandlerType: (*PieceStoreRoutesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Piece",
			Handler:    _PieceStoreRoutes_Piece_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _PieceStoreRoutes_Delete_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Retrieve",
			Handler:       _PieceStoreRoutes_Retrieve_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Store",
			Handler:       _PieceStoreRoutes_Store_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "piece_store.proto",
}

func init() { proto.RegisterFile("piece_store.proto", fileDescriptor_piece_store_fd7b526146f7ffd0) }

var fileDescriptor_piece_store_fd7b526146f7ffd0 = []byte{
	// 609 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0xcd, 0xc8, 0x7f, 0xf1, 0x8d, 0x93, 0xcf, 0x99, 0x2f, 0x04, 0x59, 0x24, 0xc1, 0x28, 0x21,
	0x98, 0x14, 0x4c, 0x49, 0x9f, 0xa0, 0x45, 0xd0, 0x66, 0x93, 0x86, 0x31, 0xdd, 0x14, 0x8a, 0x19,
	0x5b, 0x37, 0xe9, 0x80, 0x7e, 0xcc, 0x68, 0x9c, 0x3a, 0x59, 0x76, 0xd3, 0x17, 0xe8, 0xab, 0x95,
	0x3e, 0x40, 0xdf, 0xa2, 0xab, 0xa2, 0xd1, 0x58, 0xfe, 0x95, 0xbd, 0x69, 0x77, 0x73, 0xcf, 0x9d,
	0x39, 0xf7, 0xdc, 0x73, 0x67, 0x24, 0x38, 0x1c, 0x09, 0x1c, 0x62, 0x3f, 0x51, 0xb1, 0xc4, 0xee,
	0x48, 0xc6, 0x2a, 0xa6, 0x4d, 0x0d, 0x69, 0x44, 0xc6, 0x63, 0x85, 0x89, 0xfb, 0x8b, 0x80, 0x7d,
	0xc7, 0x9f, 0x50, 0xbe, 0xe1, 0x91, 0xff, 0x45, 0xf8, 0xea, 0xf3, 0xeb, 0x20, 0x88, 0x87, 0x5c,
	0x89, 0x38, 0xa2, 0x27, 0x50, 0x4f, 0xc4, 0x43, 0xc4, 0xd5, 0x58, 0xa2, 0x4d, 0xda, 0xa4, 0xd3,
	0x60, 0x33, 0x80, 0x52, 0x28, 0xfb, 0x5c, 0x71, 0xdb, 0xd2, 0x09, 0xbd, 0x76, 0xbe, 0x13, 0x28,
	0x7b, 0x5c, 0x71, 0x7a, 0x04, 0x95, 0x51, 0x4a, 0x6b, 0x8e, 0x65, 0x01, 0x3d, 0x86, 0xaa, 0xc4,
	0x48, 0xa1, 0x34, 0x87, 0x4c, 0x44, 0x5b, 0xb0, 0x1b, 0xf2, 0x49, 0x3f, 0x11, 0xcf, 0x68, 0x97,
	0xda, 0xa4, 0x53, 0x62, 0xb5, 0x90, 0x4f, 0x7a, 0xe2, 0x19, 0xe9, 0x19, 0x00, 0x4e, 0x46, 0x42,
	0x6a, 0x45, 0x76, 0x59, 0x27, 0xe7, 0x10, 0x7a, 0x0e, 0xfb, 0x09, 0x4a, 0xc1, 0x83, 0x7e, 0x34,
	0x0e, 0x07, 0x28, 0xed, 0x4a, 0x9b, 0x74, 0xea, 0xac, 0x91, 0x81, 0xb7, 0x1a, 0x73, 0x7f, 0x12,
	0x68, 0x31, 0x5d, 0xea, 0xef, 0xb4, 0xf9, 0x6d, 0xda, 0xe6, 0x07, 0x68, 0xea, 0xce, 0xfa, 0x3c,
	0xa7, 0xd3, 0x0c, 0x7b, 0xd7, 0x57, 0xdd, 0x65, 0xaf, 0xbb, 0x45, 0x3e, 0xb3, 0xff, 0x34, 0xc7,
	0x9c, 0x22, 0x0a, 0x65, 0xed, 0x85, 0xa5, 0xdb, 0xd5, 0xeb, 0xd4, 0x51, 0x15, 0x2b, 0x1e, 0x18,
	0x83, 0xb2, 0xc0, 0xfd, 0x4d, 0x00, 0xee, 0xd2, 0x42, 0xbd, 0xb4, 0x10, 0xfd, 0x04, 0xff, 0x0f,
	0xa6, 0x05, 0x56, 0x24, 0xbd, 0x58, 0x95, 0x54, 0x68, 0x0a, 0x5b, 0xc7, 0x43, 0x3d, 0xa8, 0x6b,
	0x8a, 0xdc, 0x90, 0xbd, 0xeb, 0xcb, 0x35, 0x7d, 0xe6, 0x7a, 0xb2, 0x65, 0xea, 0x14, 0x9b, 0x1d,
	0x74, 0xde, 0x42, 0x3d, 0xc7, 0xe9, 0x01, 0x58, 0xc2, 0xd7, 0x02, 0xeb, 0xcc, 0x12, 0x3e, 0x6d,
	0x42, 0x49, 0xa9, 0xc0, 0x74, 0x9e, 0x2e, 0xa9, 0x0d, 0xb5, 0x61, 0x1c, 0x29, 0x8c, 0x94, 0x6e,
	0xbd, 0xc1, 0xa6, 0xa1, 0xdb, 0x82, 0x9a, 0x26, 0xba, 0xf1, 0x97, 0x69, 0x5c, 0x06, 0x8d, 0x4c,
	0xc6, 0x38, 0x0c, 0xb9, 0x7c, 0x5a, 0x29, 0xb3, 0xce, 0xe1, 0xc5, 0xab, 0x56, 0x5a, 0xbe, 0x6a,
	0xee, 0x57, 0x0b, 0x0e, 0x34, 0x29, 0x43, 0x25, 0x05, 0x3e, 0xf2, 0xe0, 0x5f, 0xfb, 0xfd, 0xce,
	0xf8, 0xed, 0xcd, 0xfc, 0xbe, 0x2a, 0xf0, 0x3b, 0xd7, 0xb4, 0xe2, 0xb9, 0xb7, 0xd5, 0xf3, 0x75,
	0x66, 0x1c, 0x43, 0x35, 0xbe, 0xbf, 0x4f, 0x50, 0x19, 0x23, 0x4c, 0xe4, 0x7a, 0x70, 0xb4, 0x58,
	0xaf, 0xa7, 0x24, 0xf2, 0x30, 0xe7, 0x20, 0x73, 0x1c, 0x73, 0x93, 0xb3, 0x16, 0x27, 0x77, 0x0a,
	0x7b, 0x99, 0x1c, 0x0c, 0x50, 0xe1, 0xca, 0xf4, 0xba, 0x40, 0xe7, 0xd2, 0xd3, 0x19, 0xda, 0x50,
	0x0b, 0x31, 0x49, 0xf8, 0x03, 0x9a, 0xad, 0xd3, 0xd0, 0xed, 0xc1, 0xe1, 0xec, 0xd2, 0x6d, 0xdd,
	0x4e, 0x2f, 0x60, 0x5f, 0xbf, 0x1e, 0x86, 0x43, 0x14, 0x8f, 0xe8, 0x9b, 0xc6, 0x17, 0xc1, 0xeb,
	0x1f, 0x16, 0x34, 0x67, 0xac, 0x4c, 0x7b, 0x4d, 0x3d, 0xa8, 0x68, 0x8c, 0xb6, 0x0a, 0xe6, 0x70,
	0xe3, 0x3b, 0x67, 0x45, 0x4f, 0x22, 0x13, 0xe6, 0xee, 0xd0, 0x8f, 0xb0, 0x6b, 0xfc, 0x43, 0xda,
	0xde, 0x36, 0x50, 0xe7, 0x72, 0xdb, 0x8e, 0x6c, 0x04, 0xee, 0x4e, 0x87, 0xbc, 0x24, 0xf4, 0x16,
	0x2a, 0xd9, 0xb7, 0xe0, 0x64, 0xd3, 0xcb, 0x74, 0xce, 0x37, 0x65, 0x73, 0xa5, 0x1d, 0x42, 0xdf,
	0x43, 0xd5, 0x4c, 0xe9, 0xb4, 0xe0, 0x48, 0x96, 0x76, 0x2e, 0x36, 0xa6, 0x73, 0xca, 0x41, 0x55,
	0xff, 0x8b, 0x5e, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd1, 0xd4, 0x7d, 0xa0, 0x06, 0x00,
	0x00,
}
