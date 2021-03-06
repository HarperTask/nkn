// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/block.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type WinnerType int32

const (
	WinnerType_GENESIS_SIGNER WinnerType = 0
	WinnerType_TXN_SIGNER     WinnerType = 1
	WinnerType_BLOCK_SIGNER   WinnerType = 2
)

var WinnerType_name = map[int32]string{
	0: "GENESIS_SIGNER",
	1: "TXN_SIGNER",
	2: "BLOCK_SIGNER",
}
var WinnerType_value = map[string]int32{
	"GENESIS_SIGNER": 0,
	"TXN_SIGNER":     1,
	"BLOCK_SIGNER":   2,
}

func (x WinnerType) String() string {
	return proto.EnumName(WinnerType_name, int32(x))
}
func (WinnerType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_block_623cc3ece9da7f8b, []int{0}
}

type UnsignedHeader struct {
	Version              uint32     `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	PrevBlockHash        []byte     `protobuf:"bytes,2,opt,name=prev_block_hash,json=prevBlockHash,proto3" json:"prev_block_hash,omitempty"`
	TransactionsRoot     []byte     `protobuf:"bytes,3,opt,name=transactions_root,json=transactionsRoot,proto3" json:"transactions_root,omitempty"`
	StateRoot            []byte     `protobuf:"bytes,4,opt,name=state_root,json=stateRoot,proto3" json:"state_root,omitempty"`
	Timestamp            int64      `protobuf:"varint,5,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Height               uint32     `protobuf:"varint,6,opt,name=height,proto3" json:"height,omitempty"`
	RandomBeacon         []byte     `protobuf:"bytes,7,opt,name=random_beacon,json=randomBeacon,proto3" json:"random_beacon,omitempty"`
	WinnerHash           []byte     `protobuf:"bytes,8,opt,name=winner_hash,json=winnerHash,proto3" json:"winner_hash,omitempty"`
	WinnerType           WinnerType `protobuf:"varint,9,opt,name=winner_type,json=winnerType,proto3,enum=pb.WinnerType" json:"winner_type,omitempty"`
	SignerPk             []byte     `protobuf:"bytes,10,opt,name=signer_pk,json=signerPk,proto3" json:"signer_pk,omitempty"`
	SignerId             []byte     `protobuf:"bytes,11,opt,name=signer_id,json=signerId,proto3" json:"signer_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *UnsignedHeader) Reset()         { *m = UnsignedHeader{} }
func (m *UnsignedHeader) String() string { return proto.CompactTextString(m) }
func (*UnsignedHeader) ProtoMessage()    {}
func (*UnsignedHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_623cc3ece9da7f8b, []int{0}
}
func (m *UnsignedHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnsignedHeader.Unmarshal(m, b)
}
func (m *UnsignedHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnsignedHeader.Marshal(b, m, deterministic)
}
func (dst *UnsignedHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnsignedHeader.Merge(dst, src)
}
func (m *UnsignedHeader) XXX_Size() int {
	return xxx_messageInfo_UnsignedHeader.Size(m)
}
func (m *UnsignedHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_UnsignedHeader.DiscardUnknown(m)
}

var xxx_messageInfo_UnsignedHeader proto.InternalMessageInfo

func (m *UnsignedHeader) GetVersion() uint32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *UnsignedHeader) GetPrevBlockHash() []byte {
	if m != nil {
		return m.PrevBlockHash
	}
	return nil
}

func (m *UnsignedHeader) GetTransactionsRoot() []byte {
	if m != nil {
		return m.TransactionsRoot
	}
	return nil
}

func (m *UnsignedHeader) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *UnsignedHeader) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *UnsignedHeader) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *UnsignedHeader) GetRandomBeacon() []byte {
	if m != nil {
		return m.RandomBeacon
	}
	return nil
}

func (m *UnsignedHeader) GetWinnerHash() []byte {
	if m != nil {
		return m.WinnerHash
	}
	return nil
}

func (m *UnsignedHeader) GetWinnerType() WinnerType {
	if m != nil {
		return m.WinnerType
	}
	return WinnerType_GENESIS_SIGNER
}

func (m *UnsignedHeader) GetSignerPk() []byte {
	if m != nil {
		return m.SignerPk
	}
	return nil
}

func (m *UnsignedHeader) GetSignerId() []byte {
	if m != nil {
		return m.SignerId
	}
	return nil
}

type Header struct {
	UnsignedHeader       *UnsignedHeader `protobuf:"bytes,1,opt,name=unsigned_header,json=unsignedHeader,proto3" json:"unsigned_header,omitempty"`
	Signature            []byte          `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_623cc3ece9da7f8b, []int{1}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetUnsignedHeader() *UnsignedHeader {
	if m != nil {
		return m.UnsignedHeader
	}
	return nil
}

func (m *Header) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Block struct {
	Header               *Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	Transactions         []*Transaction `protobuf:"bytes,2,rep,name=transactions,proto3" json:"transactions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_block_623cc3ece9da7f8b, []int{2}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Block.Unmarshal(m, b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Block.Marshal(b, m, deterministic)
}
func (dst *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(dst, src)
}
func (m *Block) XXX_Size() int {
	return xxx_messageInfo_Block.Size(m)
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() *Header {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *Block) GetTransactions() []*Transaction {
	if m != nil {
		return m.Transactions
	}
	return nil
}

func init() {
	proto.RegisterType((*UnsignedHeader)(nil), "pb.UnsignedHeader")
	proto.RegisterType((*Header)(nil), "pb.Header")
	proto.RegisterType((*Block)(nil), "pb.Block")
	proto.RegisterEnum("pb.WinnerType", WinnerType_name, WinnerType_value)
}

func init() { proto.RegisterFile("pb/block.proto", fileDescriptor_block_623cc3ece9da7f8b) }

var fileDescriptor_block_623cc3ece9da7f8b = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x92, 0x51, 0x8b, 0xd3, 0x40,
	0x14, 0x85, 0x4d, 0xeb, 0x76, 0xb7, 0xb7, 0x6d, 0x5a, 0x2f, 0x22, 0x83, 0xae, 0x18, 0x2a, 0x48,
	0x50, 0x68, 0xa1, 0xfb, 0xe8, 0x5b, 0xa5, 0xec, 0x16, 0xa5, 0x4a, 0x5a, 0xd1, 0xb7, 0x38, 0x69,
	0x86, 0x4d, 0xa8, 0x9d, 0x19, 0x66, 0xa6, 0x2b, 0xfb, 0x03, 0xfc, 0xdf, 0x92, 0x9b, 0xc4, 0xa4,
	0x6f, 0x99, 0xef, 0x9c, 0x9c, 0x9b, 0xb9, 0x27, 0xe0, 0xeb, 0x64, 0x9e, 0xfc, 0x56, 0xfb, 0xc3,
	0x4c, 0x1b, 0xe5, 0x14, 0x76, 0x74, 0xf2, 0xf2, 0xb9, 0x4e, 0xe6, 0xce, 0x70, 0x69, 0xf9, 0xde,
	0xe5, 0x4a, 0x96, 0xca, 0xf4, 0x6f, 0x17, 0xfc, 0xef, 0xd2, 0xe6, 0xf7, 0x52, 0xa4, 0x77, 0x82,
	0xa7, 0xc2, 0x20, 0x83, 0xcb, 0x07, 0x61, 0x6c, 0xae, 0x24, 0xf3, 0x02, 0x2f, 0x1c, 0x45, 0xf5,
	0x11, 0xdf, 0xc1, 0x58, 0x1b, 0xf1, 0x10, 0x53, 0x74, 0x9c, 0x71, 0x9b, 0xb1, 0x4e, 0xe0, 0x85,
	0xc3, 0x68, 0x54, 0xe0, 0x65, 0x41, 0xef, 0xb8, 0xcd, 0xf0, 0x03, 0x3c, 0x6b, 0x4d, 0xb2, 0xb1,
	0x51, 0xca, 0xb1, 0x2e, 0x39, 0x27, 0x6d, 0x21, 0x52, 0xca, 0xe1, 0x6b, 0x00, 0xeb, 0xb8, 0x13,
	0xa5, 0xeb, 0x29, 0xb9, 0xfa, 0x44, 0x48, 0xbe, 0x86, 0xbe, 0xcb, 0x8f, 0xc2, 0x3a, 0x7e, 0xd4,
	0xec, 0x22, 0xf0, 0xc2, 0x6e, 0xd4, 0x00, 0x7c, 0x01, 0xbd, 0x4c, 0xe4, 0xf7, 0x99, 0x63, 0x3d,
	0xfa, 0xd4, 0xea, 0x84, 0x6f, 0x61, 0x64, 0xb8, 0x4c, 0xd5, 0x31, 0x4e, 0x04, 0xdf, 0x2b, 0xc9,
	0x2e, 0x29, 0x77, 0x58, 0xc2, 0x25, 0x31, 0x7c, 0x03, 0x83, 0x3f, 0xb9, 0x94, 0xc2, 0x94, 0x57,
	0xb9, 0x22, 0x0b, 0x94, 0x88, 0xee, 0x31, 0xff, 0x6f, 0x70, 0x8f, 0x5a, 0xb0, 0x7e, 0xe0, 0x85,
	0xfe, 0xc2, 0x9f, 0xe9, 0x64, 0xf6, 0x83, 0xf0, 0xee, 0x51, 0x8b, 0xfa, 0x85, 0xe2, 0x19, 0x5f,
	0x41, 0x9f, 0x56, 0x69, 0x62, 0x7d, 0x60, 0x40, 0x79, 0x57, 0x25, 0xf8, 0x76, 0x68, 0x89, 0x79,
	0xca, 0x06, 0x6d, 0x71, 0x9d, 0x4e, 0xf7, 0xd0, 0xab, 0xd6, 0xff, 0x11, 0xc6, 0xa7, 0xaa, 0x90,
	0x38, 0x23, 0x44, 0x35, 0x0c, 0x16, 0x58, 0x0c, 0x3e, 0xef, 0x2a, 0xf2, 0x4f, 0xe7, 0xdd, 0x5d,
	0x97, 0x33, 0xb8, 0x3b, 0x19, 0x51, 0x75, 0xd3, 0x80, 0xe9, 0x2f, 0xb8, 0xa0, 0x92, 0x70, 0x5a,
	0xac, 0xad, 0x15, 0x0d, 0x45, 0x74, 0x15, 0x59, 0x29, 0x78, 0x03, 0xc3, 0x76, 0x57, 0xac, 0x13,
	0x74, 0xc3, 0xc1, 0x62, 0x5c, 0x38, 0x77, 0x0d, 0x8f, 0xce, 0x4c, 0xef, 0x97, 0x00, 0xcd, 0x6a,
	0x10, 0xc1, 0xbf, 0x5d, 0x6d, 0x56, 0xdb, 0xf5, 0x36, 0xde, 0xae, 0x6f, 0x37, 0xab, 0x68, 0xf2,
	0x04, 0x7d, 0x80, 0xdd, 0xcf, 0x4d, 0x7d, 0xf6, 0x70, 0x02, 0xc3, 0xe5, 0x97, 0xaf, 0x9f, 0x3e,
	0xd7, 0xa4, 0x93, 0xf4, 0xe8, 0xcf, 0xbc, 0xf9, 0x17, 0x00, 0x00, 0xff, 0xff, 0x1e, 0x49, 0x85,
	0xf0, 0xc5, 0x02, 0x00, 0x00,
}
