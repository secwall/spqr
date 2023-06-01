// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protos/key_range.proto

package proto

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

type KeyRangeStatus int32

const (
	KeyRangeStatus_LOCKED    KeyRangeStatus = 0
	KeyRangeStatus_AVAILABLE KeyRangeStatus = 1
)

var KeyRangeStatus_name = map[int32]string{
	0: "LOCKED",
	1: "AVAILABLE",
}

var KeyRangeStatus_value = map[string]int32{
	"LOCKED":    0,
	"AVAILABLE": 1,
}

func (x KeyRangeStatus) String() string {
	return proto.EnumName(KeyRangeStatus_name, int32(x))
}

func (KeyRangeStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{0}
}

type KeyRange struct {
	LowerBound           string   `protobuf:"bytes,1,opt,name=lower_bound,json=lowerBound,proto3" json:"lower_bound,omitempty"`
	UpperBound           string   `protobuf:"bytes,2,opt,name=upper_bound,json=upperBound,proto3" json:"upper_bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyRange) Reset()         { *m = KeyRange{} }
func (m *KeyRange) String() string { return proto.CompactTextString(m) }
func (*KeyRange) ProtoMessage()    {}
func (*KeyRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{0}
}

func (m *KeyRange) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRange.Unmarshal(m, b)
}
func (m *KeyRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRange.Marshal(b, m, deterministic)
}
func (m *KeyRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRange.Merge(m, src)
}
func (m *KeyRange) XXX_Size() int {
	return xxx_messageInfo_KeyRange.Size(m)
}
func (m *KeyRange) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRange.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRange proto.InternalMessageInfo

func (m *KeyRange) GetLowerBound() string {
	if m != nil {
		return m.LowerBound
	}
	return ""
}

func (m *KeyRange) GetUpperBound() string {
	if m != nil {
		return m.UpperBound
	}
	return ""
}

// key range info is mapped to shard
type KeyRangeInfo struct {
	KeyRange             *KeyRange `protobuf:"bytes,1,opt,name=key_range,json=keyRange,proto3" json:"key_range,omitempty"`
	Krid                 string    `protobuf:"bytes,2,opt,name=krid,proto3" json:"krid,omitempty"`
	ShardId              string    `protobuf:"bytes,3,opt,name=shardId,proto3" json:"shardId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *KeyRangeInfo) Reset()         { *m = KeyRangeInfo{} }
func (m *KeyRangeInfo) String() string { return proto.CompactTextString(m) }
func (*KeyRangeInfo) ProtoMessage()    {}
func (*KeyRangeInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{1}
}

func (m *KeyRangeInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRangeInfo.Unmarshal(m, b)
}
func (m *KeyRangeInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRangeInfo.Marshal(b, m, deterministic)
}
func (m *KeyRangeInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRangeInfo.Merge(m, src)
}
func (m *KeyRangeInfo) XXX_Size() int {
	return xxx_messageInfo_KeyRangeInfo.Size(m)
}
func (m *KeyRangeInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRangeInfo.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRangeInfo proto.InternalMessageInfo

func (m *KeyRangeInfo) GetKeyRange() *KeyRange {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *KeyRangeInfo) GetKrid() string {
	if m != nil {
		return m.Krid
	}
	return ""
}

func (m *KeyRangeInfo) GetShardId() string {
	if m != nil {
		return m.ShardId
	}
	return ""
}

type ListKeyRangeRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListKeyRangeRequest) Reset()         { *m = ListKeyRangeRequest{} }
func (m *ListKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*ListKeyRangeRequest) ProtoMessage()    {}
func (*ListKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{2}
}

func (m *ListKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListKeyRangeRequest.Unmarshal(m, b)
}
func (m *ListKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *ListKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListKeyRangeRequest.Merge(m, src)
}
func (m *ListKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_ListKeyRangeRequest.Size(m)
}
func (m *ListKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListKeyRangeRequest proto.InternalMessageInfo

type AddKeyRangeRequest struct {
	KeyRangeInfo         *KeyRangeInfo `protobuf:"bytes,1,opt,name=key_range_info,json=keyRangeInfo,proto3" json:"key_range_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *AddKeyRangeRequest) Reset()         { *m = AddKeyRangeRequest{} }
func (m *AddKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*AddKeyRangeRequest) ProtoMessage()    {}
func (*AddKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{3}
}

func (m *AddKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddKeyRangeRequest.Unmarshal(m, b)
}
func (m *AddKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *AddKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddKeyRangeRequest.Merge(m, src)
}
func (m *AddKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_AddKeyRangeRequest.Size(m)
}
func (m *AddKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AddKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AddKeyRangeRequest proto.InternalMessageInfo

func (m *AddKeyRangeRequest) GetKeyRangeInfo() *KeyRangeInfo {
	if m != nil {
		return m.KeyRangeInfo
	}
	return nil
}

type SplitKeyRangeRequest struct {
	KeyRangeInfo         *KeyRangeInfo `protobuf:"bytes,1,opt,name=key_range_info,json=keyRangeInfo,proto3" json:"key_range_info,omitempty"`
	Bound                []byte        `protobuf:"bytes,2,opt,name=bound,proto3" json:"bound,omitempty"`
	SourceId             string        `protobuf:"bytes,3,opt,name=source_id,json=sourceId,proto3" json:"source_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *SplitKeyRangeRequest) Reset()         { *m = SplitKeyRangeRequest{} }
func (m *SplitKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*SplitKeyRangeRequest) ProtoMessage()    {}
func (*SplitKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{4}
}

func (m *SplitKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SplitKeyRangeRequest.Unmarshal(m, b)
}
func (m *SplitKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SplitKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *SplitKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SplitKeyRangeRequest.Merge(m, src)
}
func (m *SplitKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_SplitKeyRangeRequest.Size(m)
}
func (m *SplitKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SplitKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SplitKeyRangeRequest proto.InternalMessageInfo

func (m *SplitKeyRangeRequest) GetKeyRangeInfo() *KeyRangeInfo {
	if m != nil {
		return m.KeyRangeInfo
	}
	return nil
}

func (m *SplitKeyRangeRequest) GetBound() []byte {
	if m != nil {
		return m.Bound
	}
	return nil
}

func (m *SplitKeyRangeRequest) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

type MergeKeyRangeRequest struct {
	Bound                []byte   `protobuf:"bytes,1,opt,name=bound,proto3" json:"bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MergeKeyRangeRequest) Reset()         { *m = MergeKeyRangeRequest{} }
func (m *MergeKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*MergeKeyRangeRequest) ProtoMessage()    {}
func (*MergeKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{5}
}

func (m *MergeKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MergeKeyRangeRequest.Unmarshal(m, b)
}
func (m *MergeKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MergeKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *MergeKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MergeKeyRangeRequest.Merge(m, src)
}
func (m *MergeKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_MergeKeyRangeRequest.Size(m)
}
func (m *MergeKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MergeKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MergeKeyRangeRequest proto.InternalMessageInfo

func (m *MergeKeyRangeRequest) GetBound() []byte {
	if m != nil {
		return m.Bound
	}
	return nil
}

type MoveKeyRangeRequest struct {
	KeyRange             *KeyRangeInfo `protobuf:"bytes,1,opt,name=key_range,json=keyRange,proto3" json:"key_range,omitempty"`
	ToShardId            string        `protobuf:"bytes,2,opt,name=toShardId,proto3" json:"toShardId,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MoveKeyRangeRequest) Reset()         { *m = MoveKeyRangeRequest{} }
func (m *MoveKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*MoveKeyRangeRequest) ProtoMessage()    {}
func (*MoveKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{6}
}

func (m *MoveKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoveKeyRangeRequest.Unmarshal(m, b)
}
func (m *MoveKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoveKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *MoveKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoveKeyRangeRequest.Merge(m, src)
}
func (m *MoveKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_MoveKeyRangeRequest.Size(m)
}
func (m *MoveKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MoveKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MoveKeyRangeRequest proto.InternalMessageInfo

func (m *MoveKeyRangeRequest) GetKeyRange() *KeyRangeInfo {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

func (m *MoveKeyRangeRequest) GetToShardId() string {
	if m != nil {
		return m.ToShardId
	}
	return ""
}

type DropKeyRangeRequest struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropKeyRangeRequest) Reset()         { *m = DropKeyRangeRequest{} }
func (m *DropKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*DropKeyRangeRequest) ProtoMessage()    {}
func (*DropKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{7}
}

func (m *DropKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropKeyRangeRequest.Unmarshal(m, b)
}
func (m *DropKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *DropKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropKeyRangeRequest.Merge(m, src)
}
func (m *DropKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_DropKeyRangeRequest.Size(m)
}
func (m *DropKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropKeyRangeRequest proto.InternalMessageInfo

func (m *DropKeyRangeRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type DropAllKeyRangesRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DropAllKeyRangesRequest) Reset()         { *m = DropAllKeyRangesRequest{} }
func (m *DropAllKeyRangesRequest) String() string { return proto.CompactTextString(m) }
func (*DropAllKeyRangesRequest) ProtoMessage()    {}
func (*DropAllKeyRangesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{8}
}

func (m *DropAllKeyRangesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropAllKeyRangesRequest.Unmarshal(m, b)
}
func (m *DropAllKeyRangesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropAllKeyRangesRequest.Marshal(b, m, deterministic)
}
func (m *DropAllKeyRangesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropAllKeyRangesRequest.Merge(m, src)
}
func (m *DropAllKeyRangesRequest) XXX_Size() int {
	return xxx_messageInfo_DropAllKeyRangesRequest.Size(m)
}
func (m *DropAllKeyRangesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DropAllKeyRangesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DropAllKeyRangesRequest proto.InternalMessageInfo

type DropAllKeyRangesResponse struct {
	KeyRange             []*KeyRangeInfo `protobuf:"bytes,1,rep,name=key_range,json=keyRange,proto3" json:"key_range,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *DropAllKeyRangesResponse) Reset()         { *m = DropAllKeyRangesResponse{} }
func (m *DropAllKeyRangesResponse) String() string { return proto.CompactTextString(m) }
func (*DropAllKeyRangesResponse) ProtoMessage()    {}
func (*DropAllKeyRangesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{9}
}

func (m *DropAllKeyRangesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DropAllKeyRangesResponse.Unmarshal(m, b)
}
func (m *DropAllKeyRangesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DropAllKeyRangesResponse.Marshal(b, m, deterministic)
}
func (m *DropAllKeyRangesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DropAllKeyRangesResponse.Merge(m, src)
}
func (m *DropAllKeyRangesResponse) XXX_Size() int {
	return xxx_messageInfo_DropAllKeyRangesResponse.Size(m)
}
func (m *DropAllKeyRangesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DropAllKeyRangesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DropAllKeyRangesResponse proto.InternalMessageInfo

func (m *DropAllKeyRangesResponse) GetKeyRange() []*KeyRangeInfo {
	if m != nil {
		return m.KeyRange
	}
	return nil
}

type LockKeyRangeRequest struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LockKeyRangeRequest) Reset()         { *m = LockKeyRangeRequest{} }
func (m *LockKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*LockKeyRangeRequest) ProtoMessage()    {}
func (*LockKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{10}
}

func (m *LockKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LockKeyRangeRequest.Unmarshal(m, b)
}
func (m *LockKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LockKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *LockKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LockKeyRangeRequest.Merge(m, src)
}
func (m *LockKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_LockKeyRangeRequest.Size(m)
}
func (m *LockKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LockKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LockKeyRangeRequest proto.InternalMessageInfo

func (m *LockKeyRangeRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type UnlockKeyRangeRequest struct {
	Id                   []string `protobuf:"bytes,1,rep,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UnlockKeyRangeRequest) Reset()         { *m = UnlockKeyRangeRequest{} }
func (m *UnlockKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*UnlockKeyRangeRequest) ProtoMessage()    {}
func (*UnlockKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{11}
}

func (m *UnlockKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UnlockKeyRangeRequest.Unmarshal(m, b)
}
func (m *UnlockKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UnlockKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *UnlockKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnlockKeyRangeRequest.Merge(m, src)
}
func (m *UnlockKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_UnlockKeyRangeRequest.Size(m)
}
func (m *UnlockKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UnlockKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UnlockKeyRangeRequest proto.InternalMessageInfo

func (m *UnlockKeyRangeRequest) GetId() []string {
	if m != nil {
		return m.Id
	}
	return nil
}

type KeyRangeReply struct {
	KeyRangesInfo        []*KeyRangeInfo `protobuf:"bytes,1,rep,name=key_ranges_info,json=keyRangesInfo,proto3" json:"key_ranges_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *KeyRangeReply) Reset()         { *m = KeyRangeReply{} }
func (m *KeyRangeReply) String() string { return proto.CompactTextString(m) }
func (*KeyRangeReply) ProtoMessage()    {}
func (*KeyRangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{12}
}

func (m *KeyRangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyRangeReply.Unmarshal(m, b)
}
func (m *KeyRangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyRangeReply.Marshal(b, m, deterministic)
}
func (m *KeyRangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyRangeReply.Merge(m, src)
}
func (m *KeyRangeReply) XXX_Size() int {
	return xxx_messageInfo_KeyRangeReply.Size(m)
}
func (m *KeyRangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyRangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_KeyRangeReply proto.InternalMessageInfo

func (m *KeyRangeReply) GetKeyRangesInfo() []*KeyRangeInfo {
	if m != nil {
		return m.KeyRangesInfo
	}
	return nil
}

type ModifyReply struct {
	OperationId          string   `protobuf:"bytes,1,opt,name=operation_id,json=operationId,proto3" json:"operation_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ModifyReply) Reset()         { *m = ModifyReply{} }
func (m *ModifyReply) String() string { return proto.CompactTextString(m) }
func (*ModifyReply) ProtoMessage()    {}
func (*ModifyReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{13}
}

func (m *ModifyReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ModifyReply.Unmarshal(m, b)
}
func (m *ModifyReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ModifyReply.Marshal(b, m, deterministic)
}
func (m *ModifyReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ModifyReply.Merge(m, src)
}
func (m *ModifyReply) XXX_Size() int {
	return xxx_messageInfo_ModifyReply.Size(m)
}
func (m *ModifyReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ModifyReply.DiscardUnknown(m)
}

var xxx_messageInfo_ModifyReply proto.InternalMessageInfo

func (m *ModifyReply) GetOperationId() string {
	if m != nil {
		return m.OperationId
	}
	return ""
}

type ResolveKeyRangeRequest struct {
	Bound                string   `protobuf:"bytes,1,opt,name=bound,proto3" json:"bound,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveKeyRangeRequest) Reset()         { *m = ResolveKeyRangeRequest{} }
func (m *ResolveKeyRangeRequest) String() string { return proto.CompactTextString(m) }
func (*ResolveKeyRangeRequest) ProtoMessage()    {}
func (*ResolveKeyRangeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{14}
}

func (m *ResolveKeyRangeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveKeyRangeRequest.Unmarshal(m, b)
}
func (m *ResolveKeyRangeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveKeyRangeRequest.Marshal(b, m, deterministic)
}
func (m *ResolveKeyRangeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveKeyRangeRequest.Merge(m, src)
}
func (m *ResolveKeyRangeRequest) XXX_Size() int {
	return xxx_messageInfo_ResolveKeyRangeRequest.Size(m)
}
func (m *ResolveKeyRangeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveKeyRangeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveKeyRangeRequest proto.InternalMessageInfo

func (m *ResolveKeyRangeRequest) GetBound() string {
	if m != nil {
		return m.Bound
	}
	return ""
}

type ResolveKeyRangeReply struct {
	KeyRangeD            []string `protobuf:"bytes,1,rep,name=key_range_d,json=keyRangeD,proto3" json:"key_range_d,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResolveKeyRangeReply) Reset()         { *m = ResolveKeyRangeReply{} }
func (m *ResolveKeyRangeReply) String() string { return proto.CompactTextString(m) }
func (*ResolveKeyRangeReply) ProtoMessage()    {}
func (*ResolveKeyRangeReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_4cf7c6b005d83a8c, []int{15}
}

func (m *ResolveKeyRangeReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResolveKeyRangeReply.Unmarshal(m, b)
}
func (m *ResolveKeyRangeReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResolveKeyRangeReply.Marshal(b, m, deterministic)
}
func (m *ResolveKeyRangeReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResolveKeyRangeReply.Merge(m, src)
}
func (m *ResolveKeyRangeReply) XXX_Size() int {
	return xxx_messageInfo_ResolveKeyRangeReply.Size(m)
}
func (m *ResolveKeyRangeReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ResolveKeyRangeReply.DiscardUnknown(m)
}

var xxx_messageInfo_ResolveKeyRangeReply proto.InternalMessageInfo

func (m *ResolveKeyRangeReply) GetKeyRangeD() []string {
	if m != nil {
		return m.KeyRangeD
	}
	return nil
}

func init() {
	proto.RegisterEnum("spqr.KeyRangeStatus", KeyRangeStatus_name, KeyRangeStatus_value)
	proto.RegisterType((*KeyRange)(nil), "spqr.KeyRange")
	proto.RegisterType((*KeyRangeInfo)(nil), "spqr.KeyRangeInfo")
	proto.RegisterType((*ListKeyRangeRequest)(nil), "spqr.ListKeyRangeRequest")
	proto.RegisterType((*AddKeyRangeRequest)(nil), "spqr.AddKeyRangeRequest")
	proto.RegisterType((*SplitKeyRangeRequest)(nil), "spqr.SplitKeyRangeRequest")
	proto.RegisterType((*MergeKeyRangeRequest)(nil), "spqr.MergeKeyRangeRequest")
	proto.RegisterType((*MoveKeyRangeRequest)(nil), "spqr.MoveKeyRangeRequest")
	proto.RegisterType((*DropKeyRangeRequest)(nil), "spqr.DropKeyRangeRequest")
	proto.RegisterType((*DropAllKeyRangesRequest)(nil), "spqr.DropAllKeyRangesRequest")
	proto.RegisterType((*DropAllKeyRangesResponse)(nil), "spqr.DropAllKeyRangesResponse")
	proto.RegisterType((*LockKeyRangeRequest)(nil), "spqr.LockKeyRangeRequest")
	proto.RegisterType((*UnlockKeyRangeRequest)(nil), "spqr.UnlockKeyRangeRequest")
	proto.RegisterType((*KeyRangeReply)(nil), "spqr.KeyRangeReply")
	proto.RegisterType((*ModifyReply)(nil), "spqr.ModifyReply")
	proto.RegisterType((*ResolveKeyRangeRequest)(nil), "spqr.ResolveKeyRangeRequest")
	proto.RegisterType((*ResolveKeyRangeReply)(nil), "spqr.ResolveKeyRangeReply")
}

func init() { proto.RegisterFile("protos/key_range.proto", fileDescriptor_4cf7c6b005d83a8c) }

var fileDescriptor_4cf7c6b005d83a8c = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x55, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x4d, 0xfa, 0xef, 0x57, 0x8f, 0x9d, 0xb4, 0xbf, 0x4d, 0x5a, 0x52, 0xb7, 0x14, 0x58, 0x09,
	0x51, 0x51, 0x94, 0xa0, 0x22, 0x21, 0x84, 0x50, 0xd5, 0x84, 0xf4, 0x10, 0xc5, 0x01, 0xc9, 0x11,
	0x1c, 0xb8, 0x44, 0x69, 0xbc, 0x29, 0xab, 0x58, 0x5e, 0xd7, 0xeb, 0x04, 0xe5, 0xcc, 0xe7, 0xe2,
	0xbb, 0x21, 0x7b, 0xbb, 0xb6, 0x63, 0x6f, 0x28, 0x07, 0x4e, 0xed, 0xce, 0xbe, 0x7d, 0xf3, 0x3c,
	0xf3, 0x66, 0x02, 0x87, 0x7e, 0xc0, 0x42, 0xc6, 0x5b, 0x33, 0xb2, 0x1c, 0x05, 0x63, 0xef, 0x96,
	0x34, 0xe3, 0x00, 0xda, 0xe2, 0xfe, 0x5d, 0x80, 0x2d, 0xd8, 0xed, 0x93, 0xa5, 0x1d, 0xc5, 0xd1,
	0x13, 0xd0, 0x5d, 0xf6, 0x83, 0x04, 0xa3, 0x1b, 0x36, 0xf7, 0x9c, 0x46, 0xf9, 0x69, 0xf9, 0x4c,
	0xb3, 0x21, 0x0e, 0x75, 0xa2, 0x48, 0x04, 0x98, 0xfb, 0x7e, 0x02, 0xd8, 0x10, 0x80, 0x38, 0x14,
	0x03, 0x30, 0x05, 0x43, 0xb2, 0xf5, 0xbc, 0x29, 0x43, 0xe7, 0xa0, 0x25, 0x69, 0x63, 0x3e, 0xfd,
	0xa2, 0xda, 0x8c, 0xf2, 0x36, 0x25, 0xcc, 0xde, 0x9d, 0xc9, 0xf4, 0x08, 0xb6, 0x66, 0x01, 0x95,
	0xb4, 0xf1, 0xff, 0xa8, 0x01, 0xff, 0xf1, 0xef, 0xe3, 0xc0, 0xe9, 0x39, 0x8d, 0xcd, 0x38, 0x2c,
	0x8f, 0xf8, 0x00, 0x6a, 0x16, 0xe5, 0x61, 0xc2, 0x43, 0xee, 0xe6, 0x84, 0x87, 0xf8, 0x13, 0xa0,
	0xb6, 0xe3, 0xe4, 0xa2, 0xe8, 0x1d, 0x54, 0x13, 0x1d, 0x23, 0xea, 0x4d, 0xd9, 0xbd, 0x18, 0xb4,
	0x2a, 0x26, 0xd2, 0x6c, 0x1b, 0xb3, 0xcc, 0x09, 0xff, 0x2c, 0x43, 0x7d, 0xe8, 0xbb, 0x34, 0xfc,
	0x67, 0x94, 0xa8, 0x0e, 0xdb, 0x69, 0xfd, 0x0c, 0x5b, 0x1c, 0xd0, 0x31, 0x68, 0x9c, 0xcd, 0x83,
	0x09, 0x19, 0x51, 0xf9, 0xad, 0xbb, 0x22, 0xd0, 0x73, 0xf0, 0x2b, 0xa8, 0x0f, 0x48, 0x70, 0x4b,
	0xf2, 0x22, 0x12, 0xaa, 0x72, 0x86, 0x0a, 0x3b, 0x50, 0x1b, 0xb0, 0x45, 0x01, 0xdc, 0x2a, 0x36,
	0x43, 0x25, 0x36, 0x6d, 0xc8, 0x09, 0x68, 0x21, 0x1b, 0xde, 0x97, 0x5f, 0x74, 0x25, 0x0d, 0xe0,
	0xe7, 0x50, 0xeb, 0x06, 0xcc, 0xcf, 0x67, 0xa9, 0xc2, 0x06, 0x8d, 0xf4, 0x6c, 0x9e, 0x69, 0xf6,
	0x06, 0x75, 0xf0, 0x11, 0x3c, 0x8a, 0x60, 0x6d, 0xd7, 0x95, 0x48, 0x2e, 0x7b, 0xd5, 0x87, 0x46,
	0xf1, 0x8a, 0xfb, 0xcc, 0xe3, 0x24, 0x2f, 0x76, 0xf3, 0x21, 0xb1, 0x91, 0x1c, 0x8b, 0x4d, 0x66,
	0x0f, 0xc9, 0x79, 0x01, 0x07, 0x5f, 0x3c, 0xf7, 0x2f, 0x80, 0x7d, 0xa8, 0xa4, 0x10, 0xdf, 0x5d,
	0xa2, 0xf7, 0xb0, 0x97, 0x28, 0xe2, 0xb2, 0xe3, 0xeb, 0x74, 0x55, 0xa4, 0x2e, 0x1e, 0xbb, 0xe8,
	0x35, 0xe8, 0x03, 0xe6, 0xd0, 0xe9, 0x52, 0x50, 0x3d, 0x03, 0x83, 0xf9, 0x24, 0x18, 0x87, 0x94,
	0x79, 0x23, 0x2a, 0x27, 0x4d, 0x4f, 0x62, 0x3d, 0x07, 0x37, 0xe1, 0xd0, 0x26, 0x9c, 0xb9, 0x8b,
	0x3f, 0xf7, 0x5c, 0x93, 0x3d, 0x7f, 0x0b, 0xf5, 0x02, 0x3e, 0x4a, 0x75, 0x0a, 0x7a, 0x6a, 0x53,
	0xf9, 0x7d, 0x9a, 0x54, 0xd7, 0x7d, 0x79, 0x0e, 0x55, 0xf9, 0x60, 0x18, 0x8e, 0xc3, 0x39, 0x47,
	0x00, 0x3b, 0xd6, 0xe7, 0x8f, 0xfd, 0xeb, 0xee, 0x7e, 0x09, 0x55, 0x40, 0x6b, 0x7f, 0x6d, 0xf7,
	0xac, 0x76, 0xc7, 0xba, 0xde, 0x2f, 0x5f, 0xfc, 0xda, 0x86, 0xbd, 0x04, 0x4d, 0x82, 0x05, 0x9d,
	0x10, 0x74, 0x05, 0x46, 0x76, 0x0e, 0xd1, 0x91, 0xa8, 0x86, 0x62, 0x36, 0xcd, 0x5a, 0x6e, 0xf4,
	0x23, 0x81, 0xb8, 0x84, 0x2e, 0xc1, 0xc8, 0x76, 0x2e, 0x61, 0x28, 0x36, 0xc9, 0xfc, 0x5f, 0x5c,
	0x65, 0x6a, 0x89, 0x4b, 0xe8, 0x03, 0xe8, 0x99, 0x91, 0x47, 0x0d, 0x81, 0x29, 0x6e, 0x01, 0xf5,
	0xeb, 0x4b, 0x30, 0xb2, 0x36, 0x96, 0xd9, 0x15, 0xd6, 0x56, 0xbf, 0x1f, 0xc2, 0x7e, 0xde, 0xc4,
	0xe8, 0x71, 0xca, 0xa1, 0xf0, 0xbd, 0x79, 0xba, 0xee, 0x5a, 0x78, 0x1f, 0x97, 0x50, 0x07, 0xaa,
	0xab, 0x2e, 0x45, 0xc7, 0xe2, 0x8d, 0xd2, 0xbb, 0x6a, 0x61, 0x57, 0x50, 0x59, 0x59, 0x5c, 0xc8,
	0x14, 0x28, 0xd5, 0x36, 0x5b, 0xcb, 0xb0, 0xb2, 0x75, 0x24, 0x83, 0x6a, 0x15, 0xad, 0x2d, 0x6e,
	0x76, 0x13, 0xc9, 0xe2, 0x2a, 0xb6, 0x93, 0xfa, 0xfd, 0x00, 0xf6, 0x72, 0xae, 0x46, 0x27, 0x02,
	0xa7, 0x1e, 0x0e, 0xd3, 0x5c, 0x73, 0x1b, 0xd3, 0x75, 0x8c, 0x6f, 0x10, 0x5d, 0xb7, 0xe2, 0x1f,
	0xc0, 0x9b, 0x9d, 0xf8, 0xcf, 0x9b, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x38, 0xb7, 0x22,
	0x21, 0x07, 0x00, 0x00,
}