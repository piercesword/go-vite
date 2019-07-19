// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vitepb/message.proto

package vitepb

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

type State_PeerStatus int32

const (
	State_Connected    State_PeerStatus = 0
	State_Disconnected State_PeerStatus = 1
)

var State_PeerStatus_name = map[int32]string{
	0: "Connected",
	1: "Disconnected",
}

var State_PeerStatus_value = map[string]int32{
	"Connected":    0,
	"Disconnected": 1,
}

func (x State_PeerStatus) String() string {
	return proto.EnumName(State_PeerStatus_name, int32(x))
}

func (State_PeerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{4, 0}
}

type Handshake struct {
	Version              int64    `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	NetId                int64    `protobuf:"varint,2,opt,name=NetId,proto3" json:"NetId,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	ID                   []byte   `protobuf:"bytes,4,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp            int64    `protobuf:"varint,5,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Genesis              []byte   `protobuf:"bytes,6,opt,name=Genesis,proto3" json:"Genesis,omitempty"`
	Height               uint64   `protobuf:"varint,7,opt,name=Height,proto3" json:"Height,omitempty"`
	Head                 []byte   `protobuf:"bytes,8,opt,name=Head,proto3" json:"Head,omitempty"`
	FileAddress          []byte   `protobuf:"bytes,9,opt,name=FileAddress,proto3" json:"FileAddress,omitempty"`
	Key                  []byte   `protobuf:"bytes,10,opt,name=Key,proto3" json:"Key,omitempty"`
	Token                []byte   `protobuf:"bytes,11,opt,name=Token,proto3" json:"Token,omitempty"`
	PublicAddress        []byte   `protobuf:"bytes,12,opt,name=PublicAddress,proto3" json:"PublicAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Handshake) Reset()         { *m = Handshake{} }
func (m *Handshake) String() string { return proto.CompactTextString(m) }
func (*Handshake) ProtoMessage()    {}
func (*Handshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{0}
}

func (m *Handshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Handshake.Unmarshal(m, b)
}
func (m *Handshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Handshake.Marshal(b, m, deterministic)
}
func (m *Handshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Handshake.Merge(m, src)
}
func (m *Handshake) XXX_Size() int {
	return xxx_messageInfo_Handshake.Size(m)
}
func (m *Handshake) XXX_DiscardUnknown() {
	xxx_messageInfo_Handshake.DiscardUnknown(m)
}

var xxx_messageInfo_Handshake proto.InternalMessageInfo

func (m *Handshake) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Handshake) GetNetId() int64 {
	if m != nil {
		return m.NetId
	}
	return 0
}

func (m *Handshake) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Handshake) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *Handshake) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Handshake) GetGenesis() []byte {
	if m != nil {
		return m.Genesis
	}
	return nil
}

func (m *Handshake) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Handshake) GetHead() []byte {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *Handshake) GetFileAddress() []byte {
	if m != nil {
		return m.FileAddress
	}
	return nil
}

func (m *Handshake) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *Handshake) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

func (m *Handshake) GetPublicAddress() []byte {
	if m != nil {
		return m.PublicAddress
	}
	return nil
}

type SyncConnHandshake struct {
	ID                   []byte   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Key                  []byte   `protobuf:"bytes,3,opt,name=Key,proto3" json:"Key,omitempty"`
	Token                []byte   `protobuf:"bytes,4,opt,name=Token,proto3" json:"Token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SyncConnHandshake) Reset()         { *m = SyncConnHandshake{} }
func (m *SyncConnHandshake) String() string { return proto.CompactTextString(m) }
func (*SyncConnHandshake) ProtoMessage()    {}
func (*SyncConnHandshake) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{1}
}

func (m *SyncConnHandshake) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SyncConnHandshake.Unmarshal(m, b)
}
func (m *SyncConnHandshake) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SyncConnHandshake.Marshal(b, m, deterministic)
}
func (m *SyncConnHandshake) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SyncConnHandshake.Merge(m, src)
}
func (m *SyncConnHandshake) XXX_Size() int {
	return xxx_messageInfo_SyncConnHandshake.Size(m)
}
func (m *SyncConnHandshake) XXX_DiscardUnknown() {
	xxx_messageInfo_SyncConnHandshake.DiscardUnknown(m)
}

var xxx_messageInfo_SyncConnHandshake proto.InternalMessageInfo

func (m *SyncConnHandshake) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *SyncConnHandshake) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *SyncConnHandshake) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *SyncConnHandshake) GetToken() []byte {
	if m != nil {
		return m.Token
	}
	return nil
}

type ChunkRequest struct {
	From                 uint64   `protobuf:"varint,1,opt,name=From,proto3" json:"From,omitempty"`
	To                   uint64   `protobuf:"varint,2,opt,name=To,proto3" json:"To,omitempty"`
	PrevHash             []byte   `protobuf:"bytes,3,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	EndHash              []byte   `protobuf:"bytes,4,opt,name=EndHash,proto3" json:"EndHash,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkRequest) Reset()         { *m = ChunkRequest{} }
func (m *ChunkRequest) String() string { return proto.CompactTextString(m) }
func (*ChunkRequest) ProtoMessage()    {}
func (*ChunkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{2}
}

func (m *ChunkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkRequest.Unmarshal(m, b)
}
func (m *ChunkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkRequest.Marshal(b, m, deterministic)
}
func (m *ChunkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkRequest.Merge(m, src)
}
func (m *ChunkRequest) XXX_Size() int {
	return xxx_messageInfo_ChunkRequest.Size(m)
}
func (m *ChunkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkRequest proto.InternalMessageInfo

func (m *ChunkRequest) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ChunkRequest) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *ChunkRequest) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *ChunkRequest) GetEndHash() []byte {
	if m != nil {
		return m.EndHash
	}
	return nil
}

type ChunkResponse struct {
	From                 uint64   `protobuf:"varint,1,opt,name=From,proto3" json:"From,omitempty"`
	To                   uint64   `protobuf:"varint,2,opt,name=To,proto3" json:"To,omitempty"`
	PrevHash             []byte   `protobuf:"bytes,3,opt,name=PrevHash,proto3" json:"PrevHash,omitempty"`
	EndHash              []byte   `protobuf:"bytes,4,opt,name=EndHash,proto3" json:"EndHash,omitempty"`
	Size                 uint64   `protobuf:"varint,5,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChunkResponse) Reset()         { *m = ChunkResponse{} }
func (m *ChunkResponse) String() string { return proto.CompactTextString(m) }
func (*ChunkResponse) ProtoMessage()    {}
func (*ChunkResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{3}
}

func (m *ChunkResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChunkResponse.Unmarshal(m, b)
}
func (m *ChunkResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChunkResponse.Marshal(b, m, deterministic)
}
func (m *ChunkResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChunkResponse.Merge(m, src)
}
func (m *ChunkResponse) XXX_Size() int {
	return xxx_messageInfo_ChunkResponse.Size(m)
}
func (m *ChunkResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ChunkResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ChunkResponse proto.InternalMessageInfo

func (m *ChunkResponse) GetFrom() uint64 {
	if m != nil {
		return m.From
	}
	return 0
}

func (m *ChunkResponse) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

func (m *ChunkResponse) GetPrevHash() []byte {
	if m != nil {
		return m.PrevHash
	}
	return nil
}

func (m *ChunkResponse) GetEndHash() []byte {
	if m != nil {
		return m.EndHash
	}
	return nil
}

func (m *ChunkResponse) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type State struct {
	Peers                []*State_Peer `protobuf:"bytes,1,rep,name=Peers,proto3" json:"Peers,omitempty"`
	Patch                bool          `protobuf:"varint,2,opt,name=Patch,proto3" json:"Patch,omitempty"`
	Head                 []byte        `protobuf:"bytes,3,opt,name=Head,proto3" json:"Head,omitempty"`
	Height               uint64        `protobuf:"varint,4,opt,name=Height,proto3" json:"Height,omitempty"`
	Timestamp            int64         `protobuf:"varint,10,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}
func (*State) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{4}
}

func (m *State) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State.Unmarshal(m, b)
}
func (m *State) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State.Marshal(b, m, deterministic)
}
func (m *State) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State.Merge(m, src)
}
func (m *State) XXX_Size() int {
	return xxx_messageInfo_State.Size(m)
}
func (m *State) XXX_DiscardUnknown() {
	xxx_messageInfo_State.DiscardUnknown(m)
}

var xxx_messageInfo_State proto.InternalMessageInfo

func (m *State) GetPeers() []*State_Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func (m *State) GetPatch() bool {
	if m != nil {
		return m.Patch
	}
	return false
}

func (m *State) GetHead() []byte {
	if m != nil {
		return m.Head
	}
	return nil
}

func (m *State) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *State) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type State_Peer struct {
	ID                   []byte           `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Status               State_PeerStatus `protobuf:"varint,3,opt,name=Status,proto3,enum=vitepb.State_PeerStatus" json:"Status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *State_Peer) Reset()         { *m = State_Peer{} }
func (m *State_Peer) String() string { return proto.CompactTextString(m) }
func (*State_Peer) ProtoMessage()    {}
func (*State_Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{4, 0}
}

func (m *State_Peer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_State_Peer.Unmarshal(m, b)
}
func (m *State_Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_State_Peer.Marshal(b, m, deterministic)
}
func (m *State_Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_State_Peer.Merge(m, src)
}
func (m *State_Peer) XXX_Size() int {
	return xxx_messageInfo_State_Peer.Size(m)
}
func (m *State_Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_State_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_State_Peer proto.InternalMessageInfo

func (m *State_Peer) GetID() []byte {
	if m != nil {
		return m.ID
	}
	return nil
}

func (m *State_Peer) GetStatus() State_PeerStatus {
	if m != nil {
		return m.Status
	}
	return State_Connected
}

type HashHeight struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Height               uint64   `protobuf:"varint,2,opt,name=Height,proto3" json:"Height,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HashHeight) Reset()         { *m = HashHeight{} }
func (m *HashHeight) String() string { return proto.CompactTextString(m) }
func (*HashHeight) ProtoMessage()    {}
func (*HashHeight) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{5}
}

func (m *HashHeight) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashHeight.Unmarshal(m, b)
}
func (m *HashHeight) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashHeight.Marshal(b, m, deterministic)
}
func (m *HashHeight) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashHeight.Merge(m, src)
}
func (m *HashHeight) XXX_Size() int {
	return xxx_messageInfo_HashHeight.Size(m)
}
func (m *HashHeight) XXX_DiscardUnknown() {
	xxx_messageInfo_HashHeight.DiscardUnknown(m)
}

var xxx_messageInfo_HashHeight proto.InternalMessageInfo

func (m *HashHeight) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *HashHeight) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type HashHeightPoint struct {
	Point                *HashHeight `protobuf:"bytes,1,opt,name=Point,proto3" json:"Point,omitempty"`
	Size                 uint64      `protobuf:"varint,2,opt,name=Size,proto3" json:"Size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *HashHeightPoint) Reset()         { *m = HashHeightPoint{} }
func (m *HashHeightPoint) String() string { return proto.CompactTextString(m) }
func (*HashHeightPoint) ProtoMessage()    {}
func (*HashHeightPoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{6}
}

func (m *HashHeightPoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashHeightPoint.Unmarshal(m, b)
}
func (m *HashHeightPoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashHeightPoint.Marshal(b, m, deterministic)
}
func (m *HashHeightPoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashHeightPoint.Merge(m, src)
}
func (m *HashHeightPoint) XXX_Size() int {
	return xxx_messageInfo_HashHeightPoint.Size(m)
}
func (m *HashHeightPoint) XXX_DiscardUnknown() {
	xxx_messageInfo_HashHeightPoint.DiscardUnknown(m)
}

var xxx_messageInfo_HashHeightPoint proto.InternalMessageInfo

func (m *HashHeightPoint) GetPoint() *HashHeight {
	if m != nil {
		return m.Point
	}
	return nil
}

func (m *HashHeightPoint) GetSize() uint64 {
	if m != nil {
		return m.Size
	}
	return 0
}

type HashHeightList struct {
	Points               []*HashHeightPoint `protobuf:"bytes,1,rep,name=Points,proto3" json:"Points,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HashHeightList) Reset()         { *m = HashHeightList{} }
func (m *HashHeightList) String() string { return proto.CompactTextString(m) }
func (*HashHeightList) ProtoMessage()    {}
func (*HashHeightList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{7}
}

func (m *HashHeightList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HashHeightList.Unmarshal(m, b)
}
func (m *HashHeightList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HashHeightList.Marshal(b, m, deterministic)
}
func (m *HashHeightList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HashHeightList.Merge(m, src)
}
func (m *HashHeightList) XXX_Size() int {
	return xxx_messageInfo_HashHeightList.Size(m)
}
func (m *HashHeightList) XXX_DiscardUnknown() {
	xxx_messageInfo_HashHeightList.DiscardUnknown(m)
}

var xxx_messageInfo_HashHeightList proto.InternalMessageInfo

func (m *HashHeightList) GetPoints() []*HashHeightPoint {
	if m != nil {
		return m.Points
	}
	return nil
}

type GetHashHeightList struct {
	From                 []*HashHeight `protobuf:"bytes,1,rep,name=From,proto3" json:"From,omitempty"`
	Step                 uint64        `protobuf:"varint,2,opt,name=Step,proto3" json:"Step,omitempty"`
	To                   uint64        `protobuf:"varint,3,opt,name=To,proto3" json:"To,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *GetHashHeightList) Reset()         { *m = GetHashHeightList{} }
func (m *GetHashHeightList) String() string { return proto.CompactTextString(m) }
func (*GetHashHeightList) ProtoMessage()    {}
func (*GetHashHeightList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{8}
}

func (m *GetHashHeightList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetHashHeightList.Unmarshal(m, b)
}
func (m *GetHashHeightList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetHashHeightList.Marshal(b, m, deterministic)
}
func (m *GetHashHeightList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetHashHeightList.Merge(m, src)
}
func (m *GetHashHeightList) XXX_Size() int {
	return xxx_messageInfo_GetHashHeightList.Size(m)
}
func (m *GetHashHeightList) XXX_DiscardUnknown() {
	xxx_messageInfo_GetHashHeightList.DiscardUnknown(m)
}

var xxx_messageInfo_GetHashHeightList proto.InternalMessageInfo

func (m *GetHashHeightList) GetFrom() []*HashHeight {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetHashHeightList) GetStep() uint64 {
	if m != nil {
		return m.Step
	}
	return 0
}

func (m *GetHashHeightList) GetTo() uint64 {
	if m != nil {
		return m.To
	}
	return 0
}

type GetSnapshotBlocks struct {
	From                 *HashHeight `protobuf:"bytes,1,opt,name=From,proto3" json:"From,omitempty"`
	Count                uint64      `protobuf:"varint,2,opt,name=Count,proto3" json:"Count,omitempty"`
	Forward              bool        `protobuf:"varint,3,opt,name=Forward,proto3" json:"Forward,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetSnapshotBlocks) Reset()         { *m = GetSnapshotBlocks{} }
func (m *GetSnapshotBlocks) String() string { return proto.CompactTextString(m) }
func (*GetSnapshotBlocks) ProtoMessage()    {}
func (*GetSnapshotBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{9}
}

func (m *GetSnapshotBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSnapshotBlocks.Unmarshal(m, b)
}
func (m *GetSnapshotBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSnapshotBlocks.Marshal(b, m, deterministic)
}
func (m *GetSnapshotBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSnapshotBlocks.Merge(m, src)
}
func (m *GetSnapshotBlocks) XXX_Size() int {
	return xxx_messageInfo_GetSnapshotBlocks.Size(m)
}
func (m *GetSnapshotBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSnapshotBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_GetSnapshotBlocks proto.InternalMessageInfo

func (m *GetSnapshotBlocks) GetFrom() *HashHeight {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetSnapshotBlocks) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetSnapshotBlocks) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

type SnapshotBlocks struct {
	Blocks               []*SnapshotBlock `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SnapshotBlocks) Reset()         { *m = SnapshotBlocks{} }
func (m *SnapshotBlocks) String() string { return proto.CompactTextString(m) }
func (*SnapshotBlocks) ProtoMessage()    {}
func (*SnapshotBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{10}
}

func (m *SnapshotBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SnapshotBlocks.Unmarshal(m, b)
}
func (m *SnapshotBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SnapshotBlocks.Marshal(b, m, deterministic)
}
func (m *SnapshotBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SnapshotBlocks.Merge(m, src)
}
func (m *SnapshotBlocks) XXX_Size() int {
	return xxx_messageInfo_SnapshotBlocks.Size(m)
}
func (m *SnapshotBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_SnapshotBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_SnapshotBlocks proto.InternalMessageInfo

func (m *SnapshotBlocks) GetBlocks() []*SnapshotBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type GetAccountBlocks struct {
	Address              []byte      `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	From                 *HashHeight `protobuf:"bytes,2,opt,name=From,proto3" json:"From,omitempty"`
	Count                uint64      `protobuf:"varint,3,opt,name=Count,proto3" json:"Count,omitempty"`
	Forward              bool        `protobuf:"varint,4,opt,name=Forward,proto3" json:"Forward,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *GetAccountBlocks) Reset()         { *m = GetAccountBlocks{} }
func (m *GetAccountBlocks) String() string { return proto.CompactTextString(m) }
func (*GetAccountBlocks) ProtoMessage()    {}
func (*GetAccountBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{11}
}

func (m *GetAccountBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetAccountBlocks.Unmarshal(m, b)
}
func (m *GetAccountBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetAccountBlocks.Marshal(b, m, deterministic)
}
func (m *GetAccountBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAccountBlocks.Merge(m, src)
}
func (m *GetAccountBlocks) XXX_Size() int {
	return xxx_messageInfo_GetAccountBlocks.Size(m)
}
func (m *GetAccountBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAccountBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_GetAccountBlocks proto.InternalMessageInfo

func (m *GetAccountBlocks) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *GetAccountBlocks) GetFrom() *HashHeight {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *GetAccountBlocks) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *GetAccountBlocks) GetForward() bool {
	if m != nil {
		return m.Forward
	}
	return false
}

type AccountBlocks struct {
	Blocks               []*AccountBlock `protobuf:"bytes,1,rep,name=Blocks,proto3" json:"Blocks,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *AccountBlocks) Reset()         { *m = AccountBlocks{} }
func (m *AccountBlocks) String() string { return proto.CompactTextString(m) }
func (*AccountBlocks) ProtoMessage()    {}
func (*AccountBlocks) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{12}
}

func (m *AccountBlocks) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AccountBlocks.Unmarshal(m, b)
}
func (m *AccountBlocks) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AccountBlocks.Marshal(b, m, deterministic)
}
func (m *AccountBlocks) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AccountBlocks.Merge(m, src)
}
func (m *AccountBlocks) XXX_Size() int {
	return xxx_messageInfo_AccountBlocks.Size(m)
}
func (m *AccountBlocks) XXX_DiscardUnknown() {
	xxx_messageInfo_AccountBlocks.DiscardUnknown(m)
}

var xxx_messageInfo_AccountBlocks proto.InternalMessageInfo

func (m *AccountBlocks) GetBlocks() []*AccountBlock {
	if m != nil {
		return m.Blocks
	}
	return nil
}

type NewSnapshotBlock struct {
	Block                *SnapshotBlock `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	TTL                  int32          `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *NewSnapshotBlock) Reset()         { *m = NewSnapshotBlock{} }
func (m *NewSnapshotBlock) String() string { return proto.CompactTextString(m) }
func (*NewSnapshotBlock) ProtoMessage()    {}
func (*NewSnapshotBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{13}
}

func (m *NewSnapshotBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewSnapshotBlock.Unmarshal(m, b)
}
func (m *NewSnapshotBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewSnapshotBlock.Marshal(b, m, deterministic)
}
func (m *NewSnapshotBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewSnapshotBlock.Merge(m, src)
}
func (m *NewSnapshotBlock) XXX_Size() int {
	return xxx_messageInfo_NewSnapshotBlock.Size(m)
}
func (m *NewSnapshotBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_NewSnapshotBlock.DiscardUnknown(m)
}

var xxx_messageInfo_NewSnapshotBlock proto.InternalMessageInfo

func (m *NewSnapshotBlock) GetBlock() *SnapshotBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *NewSnapshotBlock) GetTTL() int32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type NewAccountBlock struct {
	Block                *AccountBlock `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	TTL                  int32         `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *NewAccountBlock) Reset()         { *m = NewAccountBlock{} }
func (m *NewAccountBlock) String() string { return proto.CompactTextString(m) }
func (*NewAccountBlock) ProtoMessage()    {}
func (*NewAccountBlock) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{14}
}

func (m *NewAccountBlock) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountBlock.Unmarshal(m, b)
}
func (m *NewAccountBlock) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountBlock.Marshal(b, m, deterministic)
}
func (m *NewAccountBlock) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountBlock.Merge(m, src)
}
func (m *NewAccountBlock) XXX_Size() int {
	return xxx_messageInfo_NewAccountBlock.Size(m)
}
func (m *NewAccountBlock) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountBlock.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountBlock proto.InternalMessageInfo

func (m *NewAccountBlock) GetBlock() *AccountBlock {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *NewAccountBlock) GetTTL() int32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type NewAccountBlockBytes struct {
	Block                []byte   `protobuf:"bytes,1,opt,name=Block,proto3" json:"Block,omitempty"`
	TTL                  int32    `protobuf:"varint,2,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewAccountBlockBytes) Reset()         { *m = NewAccountBlockBytes{} }
func (m *NewAccountBlockBytes) String() string { return proto.CompactTextString(m) }
func (*NewAccountBlockBytes) ProtoMessage()    {}
func (*NewAccountBlockBytes) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{15}
}

func (m *NewAccountBlockBytes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewAccountBlockBytes.Unmarshal(m, b)
}
func (m *NewAccountBlockBytes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewAccountBlockBytes.Marshal(b, m, deterministic)
}
func (m *NewAccountBlockBytes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewAccountBlockBytes.Merge(m, src)
}
func (m *NewAccountBlockBytes) XXX_Size() int {
	return xxx_messageInfo_NewAccountBlockBytes.Size(m)
}
func (m *NewAccountBlockBytes) XXX_DiscardUnknown() {
	xxx_messageInfo_NewAccountBlockBytes.DiscardUnknown(m)
}

var xxx_messageInfo_NewAccountBlockBytes proto.InternalMessageInfo

func (m *NewAccountBlockBytes) GetBlock() []byte {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *NewAccountBlockBytes) GetTTL() int32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

type Trace struct {
	Hash                 []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Path                 [][]byte `protobuf:"bytes,2,rep,name=Path,proto3" json:"Path,omitempty"`
	TTL                  uint32   `protobuf:"varint,3,opt,name=TTL,proto3" json:"TTL,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Trace) Reset()         { *m = Trace{} }
func (m *Trace) String() string { return proto.CompactTextString(m) }
func (*Trace) ProtoMessage()    {}
func (*Trace) Descriptor() ([]byte, []int) {
	return fileDescriptor_2a6a8486deb9ab39, []int{16}
}

func (m *Trace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Trace.Unmarshal(m, b)
}
func (m *Trace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Trace.Marshal(b, m, deterministic)
}
func (m *Trace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Trace.Merge(m, src)
}
func (m *Trace) XXX_Size() int {
	return xxx_messageInfo_Trace.Size(m)
}
func (m *Trace) XXX_DiscardUnknown() {
	xxx_messageInfo_Trace.DiscardUnknown(m)
}

var xxx_messageInfo_Trace proto.InternalMessageInfo

func (m *Trace) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

func (m *Trace) GetPath() [][]byte {
	if m != nil {
		return m.Path
	}
	return nil
}

func (m *Trace) GetTTL() uint32 {
	if m != nil {
		return m.TTL
	}
	return 0
}

func init() {
	proto.RegisterEnum("vitepb.State_PeerStatus", State_PeerStatus_name, State_PeerStatus_value)
	proto.RegisterType((*Handshake)(nil), "vitepb.Handshake")
	proto.RegisterType((*SyncConnHandshake)(nil), "vitepb.SyncConnHandshake")
	proto.RegisterType((*ChunkRequest)(nil), "vitepb.ChunkRequest")
	proto.RegisterType((*ChunkResponse)(nil), "vitepb.ChunkResponse")
	proto.RegisterType((*State)(nil), "vitepb.State")
	proto.RegisterType((*State_Peer)(nil), "vitepb.State.Peer")
	proto.RegisterType((*HashHeight)(nil), "vitepb.HashHeight")
	proto.RegisterType((*HashHeightPoint)(nil), "vitepb.HashHeightPoint")
	proto.RegisterType((*HashHeightList)(nil), "vitepb.HashHeightList")
	proto.RegisterType((*GetHashHeightList)(nil), "vitepb.GetHashHeightList")
	proto.RegisterType((*GetSnapshotBlocks)(nil), "vitepb.GetSnapshotBlocks")
	proto.RegisterType((*SnapshotBlocks)(nil), "vitepb.SnapshotBlocks")
	proto.RegisterType((*GetAccountBlocks)(nil), "vitepb.GetAccountBlocks")
	proto.RegisterType((*AccountBlocks)(nil), "vitepb.AccountBlocks")
	proto.RegisterType((*NewSnapshotBlock)(nil), "vitepb.NewSnapshotBlock")
	proto.RegisterType((*NewAccountBlock)(nil), "vitepb.NewAccountBlock")
	proto.RegisterType((*NewAccountBlockBytes)(nil), "vitepb.NewAccountBlockBytes")
	proto.RegisterType((*Trace)(nil), "vitepb.Trace")
}

func init() { proto.RegisterFile("vitepb/message.proto", fileDescriptor_2a6a8486deb9ab39) }

var fileDescriptor_2a6a8486deb9ab39 = []byte{
	// 788 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x55, 0x4b, 0x6f, 0x1a, 0x57,
	0x14, 0xee, 0xbc, 0x30, 0x1c, 0x03, 0xc5, 0x57, 0xb4, 0x1d, 0xd1, 0x2e, 0xd0, 0xa8, 0xaa, 0x50,
	0x5b, 0xe3, 0xca, 0xdd, 0x74, 0xd3, 0x56, 0xd8, 0x8e, 0x8d, 0x15, 0x8b, 0x90, 0x0b, 0xca, 0xd6,
	0x1a, 0x86, 0x23, 0x33, 0xc2, 0xcc, 0x90, 0xb9, 0x17, 0x5b, 0x8e, 0x94, 0x5d, 0x56, 0xf9, 0x3b,
	0xf9, 0x83, 0xd1, 0x7d, 0xcc, 0x0b, 0x43, 0x94, 0x4d, 0x76, 0xe7, 0x35, 0xe7, 0x3b, 0xdf, 0xb9,
	0x1f, 0x07, 0x68, 0x3f, 0x84, 0x1c, 0xd7, 0xb3, 0x93, 0x15, 0x32, 0xe6, 0xdf, 0x61, 0x7f, 0x9d,
	0xc4, 0x3c, 0x26, 0x15, 0x15, 0xed, 0x74, 0x74, 0xd6, 0x0f, 0x82, 0x78, 0x13, 0xf1, 0xdb, 0xd9,
	0x7d, 0x1c, 0x2c, 0x55, 0x4d, 0xe7, 0x67, 0x9d, 0x63, 0x91, 0xbf, 0x66, 0x8b, 0xb8, 0x94, 0xf4,
	0x3e, 0x99, 0x50, 0x1b, 0xfa, 0xd1, 0x9c, 0x2d, 0xfc, 0x25, 0x12, 0x17, 0x0e, 0xde, 0x60, 0xc2,
	0xc2, 0x38, 0x72, 0x8d, 0xae, 0xd1, 0xb3, 0x68, 0xea, 0x92, 0x36, 0x38, 0x23, 0xe4, 0xd7, 0x73,
	0xd7, 0x94, 0x71, 0xe5, 0x10, 0x02, 0xf6, 0xc8, 0x5f, 0xa1, 0x6b, 0x75, 0x8d, 0x5e, 0x8d, 0x4a,
	0x9b, 0x34, 0xc1, 0xbc, 0xbe, 0x70, 0xed, 0xae, 0xd1, 0xab, 0x53, 0xf3, 0xfa, 0x82, 0xfc, 0x02,
	0xb5, 0x69, 0xb8, 0x42, 0xc6, 0xfd, 0xd5, 0xda, 0x75, 0xe4, 0xd7, 0x79, 0x40, 0x20, 0x5e, 0x61,
	0x84, 0x2c, 0x64, 0x6e, 0x45, 0x7e, 0x92, 0xba, 0xe4, 0x47, 0xa8, 0x0c, 0x31, 0xbc, 0x5b, 0x70,
	0xf7, 0xa0, 0x6b, 0xf4, 0x6c, 0xaa, 0x3d, 0x81, 0x39, 0x44, 0x7f, 0xee, 0x56, 0x65, 0xb9, 0xb4,
	0x49, 0x17, 0x0e, 0x2f, 0xc3, 0x7b, 0x1c, 0xcc, 0xe7, 0x09, 0x32, 0xe6, 0xd6, 0x64, 0xaa, 0x18,
	0x22, 0x2d, 0xb0, 0x5e, 0xe2, 0x93, 0x0b, 0x32, 0x23, 0x4c, 0xc1, 0x68, 0x1a, 0x2f, 0x31, 0x72,
	0x0f, 0x65, 0x4c, 0x39, 0xe4, 0x57, 0x68, 0x8c, 0x37, 0xb3, 0xfb, 0x30, 0x48, 0x7b, 0xd5, 0x65,
	0xb6, 0x1c, 0xf4, 0x42, 0x38, 0x9a, 0x3c, 0x45, 0xc1, 0x79, 0x1c, 0x45, 0xf9, 0xf2, 0x14, 0x71,
	0x63, 0x37, 0x71, 0x73, 0x9b, 0xb8, 0x1e, 0xc8, 0xda, 0x31, 0x90, 0x5d, 0x18, 0xc8, 0x5b, 0x40,
	0xfd, 0x7c, 0xb1, 0x89, 0x96, 0x14, 0xdf, 0x6e, 0x90, 0x49, 0xfa, 0x97, 0x49, 0xbc, 0x92, 0x38,
	0x36, 0x95, 0xb6, 0x40, 0x9e, 0xc6, 0x12, 0xc2, 0xa6, 0xe6, 0x34, 0x26, 0x1d, 0xa8, 0x8e, 0x13,
	0x7c, 0x18, 0xfa, 0x6c, 0xa1, 0x01, 0x32, 0x5f, 0x2c, 0xfc, 0x45, 0x34, 0x97, 0x29, 0x85, 0x93,
	0xba, 0xde, 0x7b, 0x68, 0x68, 0x24, 0xb6, 0x8e, 0x23, 0x86, 0xdf, 0x0e, 0x4a, 0x74, 0x9e, 0x84,
	0xef, 0x50, 0xca, 0xc1, 0xa6, 0xd2, 0xf6, 0x3e, 0x9a, 0xe0, 0x4c, 0xb8, 0xcf, 0x91, 0xf4, 0xc0,
	0x19, 0x23, 0x26, 0xcc, 0x35, 0xba, 0x56, 0xef, 0xf0, 0x94, 0xf4, 0x95, 0x80, 0xfb, 0x32, 0xdb,
	0x17, 0x29, 0xaa, 0x0a, 0xc4, 0xca, 0xc6, 0x3e, 0x0f, 0x16, 0x72, 0xa0, 0x2a, 0x55, 0x4e, 0xa6,
	0x10, 0xab, 0xa0, 0x90, 0x5c, 0x4d, 0x76, 0x49, 0x4d, 0xa5, 0x47, 0x82, 0xad, 0x47, 0xea, 0x0c,
	0xc1, 0x16, 0x40, 0xcf, 0x9e, 0xf6, 0x2f, 0xa8, 0x88, 0x61, 0x36, 0x4c, 0x62, 0x34, 0x4f, 0xdd,
	0xe7, 0x23, 0xaa, 0x3c, 0xd5, 0x75, 0xde, 0x31, 0x40, 0x1e, 0x25, 0x0d, 0xa8, 0x09, 0xed, 0x60,
	0xc0, 0x71, 0xde, 0xfa, 0x8e, 0xb4, 0xa0, 0x7e, 0x11, 0xb2, 0x20, 0x8b, 0x18, 0xde, 0x3f, 0x00,
	0x62, 0x51, 0x05, 0xc9, 0x8b, 0x2d, 0x1a, 0x9a, 0x90, 0x58, 0x61, 0x4e, 0xc8, 0x2c, 0x12, 0xf2,
	0x5e, 0xc1, 0xf7, 0xf9, 0x97, 0xe3, 0x38, 0x8c, 0xb8, 0xdc, 0xa7, 0x30, 0xe4, 0xf7, 0x85, 0x7d,
	0xe6, 0x75, 0x54, 0x15, 0x64, 0xef, 0x62, 0x16, 0xde, 0x65, 0x00, 0xcd, 0xbc, 0xf0, 0x26, 0x64,
	0x9c, 0x9c, 0x40, 0x45, 0x96, 0xa7, 0x0f, 0xf4, 0xd3, 0xf3, 0x86, 0x32, 0x4f, 0x75, 0x99, 0x77,
	0x0b, 0x47, 0x57, 0xc8, 0xb7, 0xba, 0xfc, 0x96, 0xa9, 0xcb, 0xda, 0x33, 0x94, 0x52, 0x9c, 0x98,
	0x89, 0xe3, 0x3a, 0x9b, 0x89, 0xe3, 0x5a, 0xab, 0xd0, 0x4a, 0x55, 0xe8, 0x2d, 0x25, 0xc0, 0x44,
	0x1f, 0xb8, 0x33, 0x71, 0xdf, 0x58, 0x01, 0xc0, 0xf8, 0x22, 0x40, 0x1b, 0x9c, 0x73, 0x71, 0x34,
	0x35, 0x82, 0x72, 0x84, 0x78, 0x2f, 0xe3, 0xe4, 0xd1, 0x4f, 0x94, 0x8e, 0xaa, 0x34, 0x75, 0xbd,
	0xff, 0xa1, 0xb9, 0x85, 0x74, 0x0c, 0x15, 0x65, 0x69, 0x32, 0x3f, 0x64, 0x72, 0x28, 0xd6, 0x51,
	0x5d, 0xe4, 0x7d, 0x30, 0xa0, 0x75, 0x85, 0x7c, 0xa0, 0x6e, 0xb5, 0xee, 0xe1, 0xc2, 0x41, 0x7a,
	0x72, 0xd4, 0x33, 0xa7, 0x6e, 0xc6, 0xc3, 0xfc, 0x5a, 0x1e, 0xd6, 0x1e, 0x1e, 0x76, 0x99, 0xc7,
	0xbf, 0xd0, 0x28, 0x8f, 0xf0, 0xe7, 0x16, 0x8d, 0x76, 0x0a, 0x55, 0x2c, 0xcb, 0x58, 0xbc, 0x86,
	0xd6, 0x08, 0x1f, 0x4b, 0x0c, 0xc9, 0x1f, 0xe0, 0x48, 0x43, 0xef, 0x7c, 0xcf, 0x1e, 0x54, 0x8d,
	0xb8, 0x80, 0xd3, 0xe9, 0x8d, 0xa4, 0xe5, 0x50, 0x61, 0x0a, 0xed, 0x8e, 0xf0, 0xb1, 0x88, 0x46,
	0x7e, 0x2f, 0x77, 0xdc, 0x3d, 0xd2, 0xde, 0x86, 0xff, 0x41, 0x7b, 0xab, 0xe1, 0xd9, 0x13, 0x47,
	0x79, 0x37, 0xf2, 0xae, 0xf5, 0xfd, 0xdf, 0x0f, 0xc0, 0x99, 0x26, 0x7e, 0x80, 0x3b, 0x7f, 0x81,
	0x04, 0xec, 0xb1, 0xcf, 0xc5, 0xed, 0xb1, 0x44, 0x4c, 0xd8, 0x69, 0x0b, 0xf1, 0x02, 0x0d, 0xd9,
	0x62, 0x56, 0x91, 0xff, 0xb3, 0x7f, 0x7f, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x61, 0x22, 0x50, 0xa0,
	0xc0, 0x07, 0x00, 0x00,
}
