// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/base/tendermint/v1beta1/types.proto

package tmservice

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	types "github.com/tendermint/tendermint/proto/tendermint/types"
	version "github.com/tendermint/tendermint/proto/tendermint/version"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Block is tendermint type Block, with the Header proposer address
// field converted to bech32 string.
type Block struct {
	Header     Header        `protobuf:"bytes,1,opt,name=header,proto3" json:"header"`
	Data       types.Data    `protobuf:"bytes,2,opt,name=data,proto3" json:"data"`
	LastCommit *types.Commit `protobuf:"bytes,4,opt,name=last_commit,json=lastCommit,proto3" json:"last_commit,omitempty"`
}

func (m *Block) Reset()         { *m = Block{} }
func (m *Block) String() string { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()    {}
func (*Block) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9931519c08e0d6, []int{0}
}
func (m *Block) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Block) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Block.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Block) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Block.Merge(m, src)
}
func (m *Block) XXX_Size() int {
	return m.Size()
}
func (m *Block) XXX_DiscardUnknown() {
	xxx_messageInfo_Block.DiscardUnknown(m)
}

var xxx_messageInfo_Block proto.InternalMessageInfo

func (m *Block) GetHeader() Header {
	if m != nil {
		return m.Header
	}
	return Header{}
}

func (m *Block) GetData() types.Data {
	if m != nil {
		return m.Data
	}
	return types.Data{}
}

func (m *Block) GetLastCommit() *types.Commit {
	if m != nil {
		return m.LastCommit
	}
	return nil
}

// Header defines the structure of a Tendermint block header.
type Header struct {
	// basic block info
	Version version.Consensus `protobuf:"bytes,1,opt,name=version,proto3" json:"version"`
	ChainID string            `protobuf:"bytes,2,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	Height  int64             `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Time    time.Time         `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
	// prev block info
	LastBlockId types.BlockID `protobuf:"bytes,5,opt,name=last_block_id,json=lastBlockId,proto3" json:"last_block_id"`
	// hashes of block data
	LastCommitHash []byte `protobuf:"bytes,6,opt,name=last_commit_hash,json=lastCommitHash,proto3" json:"last_commit_hash,omitempty"`
	DataHash       []byte `protobuf:"bytes,7,opt,name=data_hash,json=dataHash,proto3" json:"data_hash,omitempty"`
	// hashes from the app output from the prev block
	ValidatorsHash     []byte `protobuf:"bytes,8,opt,name=validators_hash,json=validatorsHash,proto3" json:"validators_hash,omitempty"`
	NextValidatorsHash []byte `protobuf:"bytes,9,opt,name=next_validators_hash,json=nextValidatorsHash,proto3" json:"next_validators_hash,omitempty"`
	ConsensusHash      []byte `protobuf:"bytes,10,opt,name=consensus_hash,json=consensusHash,proto3" json:"consensus_hash,omitempty"`
	AppHash            []byte `protobuf:"bytes,11,opt,name=app_hash,json=appHash,proto3" json:"app_hash,omitempty"`
	LastResultsHash    []byte `protobuf:"bytes,12,opt,name=last_results_hash,json=lastResultsHash,proto3" json:"last_results_hash,omitempty"`
	// consensus info
	EvidenceHash []byte `protobuf:"bytes,13,opt,name=evidence_hash,json=evidenceHash,proto3" json:"evidence_hash,omitempty"`
	// proposer_address is the original block proposer address, formatted as a Bech32 string.
	// In Tendermint, this type is `bytes`, but in the SDK, we convert it to a Bech32 string
	// for better UX.
	ProposerAddress string `protobuf:"bytes,14,opt,name=proposer_address,json=proposerAddress,proto3" json:"proposer_address,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_bb9931519c08e0d6, []int{1}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetVersion() version.Consensus {
	if m != nil {
		return m.Version
	}
	return version.Consensus{}
}

func (m *Header) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetTime() time.Time {
	if m != nil {
		return m.Time
	}
	return time.Time{}
}

func (m *Header) GetLastBlockId() types.BlockID {
	if m != nil {
		return m.LastBlockId
	}
	return types.BlockID{}
}

func (m *Header) GetLastCommitHash() []byte {
	if m != nil {
		return m.LastCommitHash
	}
	return nil
}

func (m *Header) GetDataHash() []byte {
	if m != nil {
		return m.DataHash
	}
	return nil
}

func (m *Header) GetValidatorsHash() []byte {
	if m != nil {
		return m.ValidatorsHash
	}
	return nil
}

func (m *Header) GetNextValidatorsHash() []byte {
	if m != nil {
		return m.NextValidatorsHash
	}
	return nil
}

func (m *Header) GetConsensusHash() []byte {
	if m != nil {
		return m.ConsensusHash
	}
	return nil
}

func (m *Header) GetAppHash() []byte {
	if m != nil {
		return m.AppHash
	}
	return nil
}

func (m *Header) GetLastResultsHash() []byte {
	if m != nil {
		return m.LastResultsHash
	}
	return nil
}

func (m *Header) GetEvidenceHash() []byte {
	if m != nil {
		return m.EvidenceHash
	}
	return nil
}

func (m *Header) GetProposerAddress() string {
	if m != nil {
		return m.ProposerAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Block)(nil), "cosmos.base.tendermint.v1beta1.Block")
	proto.RegisterType((*Header)(nil), "cosmos.base.tendermint.v1beta1.Header")
}

func init() {
	proto.RegisterFile("cosmos/base/tendermint/v1beta1/types.proto", fileDescriptor_bb9931519c08e0d6)
}

var fileDescriptor_bb9931519c08e0d6 = []byte{
	// 598 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xcd, 0x6e, 0xd3, 0x4e,
	0x14, 0xc5, 0xe3, 0x7f, 0xd3, 0x7c, 0x4c, 0xbe, 0xfa, 0x1f, 0x55, 0x95, 0x1b, 0xc0, 0x89, 0x8a,
	0x28, 0xa1, 0x12, 0x76, 0x5b, 0x36, 0xb0, 0x60, 0x41, 0x92, 0x45, 0x23, 0xb1, 0xb2, 0x10, 0x0b,
	0x36, 0xd1, 0xd8, 0x1e, 0xec, 0x51, 0x6d, 0x8f, 0xe5, 0x99, 0x44, 0xf0, 0x16, 0x7d, 0x1e, 0x9e,
	0xa0, 0xcb, 0x2e, 0x59, 0x15, 0x94, 0x48, 0x3c, 0x07, 0x9a, 0x3b, 0x36, 0x49, 0xa8, 0xc4, 0x2a,
	0x33, 0xf7, 0xfe, 0xce, 0x9d, 0x33, 0x67, 0x22, 0xa3, 0x33, 0x9f, 0x8b, 0x84, 0x0b, 0xc7, 0x23,
	0x82, 0x3a, 0x92, 0xa6, 0x01, 0xcd, 0x13, 0x96, 0x4a, 0x67, 0x79, 0xe1, 0x51, 0x49, 0x2e, 0x1c,
	0xf9, 0x35, 0xa3, 0xc2, 0xce, 0x72, 0x2e, 0x39, 0xb6, 0x34, 0x6b, 0x2b, 0xd6, 0xde, 0xb0, 0x76,
	0xc1, 0xf6, 0x0f, 0x43, 0x1e, 0x72, 0x40, 0x1d, 0xb5, 0xd2, 0xaa, 0xfe, 0xe3, 0xad, 0xa9, 0x30,
	0x6d, 0x7b, 0x66, 0xdf, 0xda, 0x3e, 0x93, 0xe6, 0x82, 0xf1, 0x74, 0xa7, 0x3f, 0x08, 0x39, 0x0f,
	0x63, 0xea, 0xc0, 0xce, 0x5b, 0x7c, 0x76, 0x24, 0x4b, 0xa8, 0x90, 0x24, 0xc9, 0x34, 0x70, 0xf2,
	0xcd, 0x40, 0xfb, 0xe3, 0x98, 0xfb, 0xd7, 0x78, 0x8a, 0x6a, 0x11, 0x25, 0x01, 0xcd, 0x4d, 0x63,
	0x68, 0x8c, 0x5a, 0x97, 0xa7, 0xf6, 0xbf, 0xfd, 0xda, 0x57, 0x40, 0x8f, 0xab, 0xb7, 0xf7, 0x83,
	0x8a, 0x5b, 0x68, 0xf1, 0x39, 0xaa, 0x06, 0x44, 0x12, 0xf3, 0x3f, 0x98, 0x71, 0xb4, 0xad, 0xd3,
	0xbe, 0xa6, 0x44, 0x92, 0x42, 0x03, 0x24, 0x7e, 0x83, 0x5a, 0x31, 0x11, 0x72, 0xee, 0xf3, 0x24,
	0x61, 0xd2, 0xac, 0x82, 0xd0, 0x7c, 0x28, 0x9c, 0x40, 0xdf, 0x45, 0x0a, 0xd6, 0xeb, 0x93, 0x5f,
	0x55, 0x54, 0xd3, 0x2e, 0xf0, 0x5b, 0x54, 0x2f, 0xee, 0x5f, 0xd8, 0x7f, 0xb2, 0x63, 0x59, 0xb7,
	0xec, 0x09, 0x4f, 0x05, 0x4d, 0xc5, 0x42, 0x14, 0x0e, 0x4a, 0x0d, 0x3e, 0x45, 0x0d, 0x3f, 0x22,
	0x2c, 0x9d, 0xb3, 0x00, 0xac, 0x37, 0xc7, 0xad, 0xd5, 0xfd, 0xa0, 0x3e, 0x51, 0xb5, 0xd9, 0xd4,
	0xad, 0x43, 0x73, 0x16, 0xe0, 0x23, 0x15, 0x12, 0x0b, 0x23, 0x69, 0xee, 0x0d, 0x8d, 0xd1, 0x9e,
	0x5b, 0xec, 0xf0, 0x6b, 0x54, 0x55, 0xc9, 0x16, 0xee, 0xfb, 0xb6, 0x8e, 0xdd, 0x2e, 0x63, 0xb7,
	0x3f, 0x94, 0xb1, 0x8f, 0x1b, 0xea, 0xe0, 0x9b, 0x1f, 0x03, 0xc3, 0x05, 0x05, 0x9e, 0xa0, 0x0e,
	0x5c, 0xdf, 0x53, 0x8f, 0xa0, 0x8e, 0xdf, 0x87, 0x11, 0xc7, 0x0f, 0x03, 0x80, 0x67, 0x9a, 0x4d,
	0x0b, 0xeb, 0x10, 0x9a, 0x2e, 0x05, 0x78, 0x84, 0x0e, 0xb6, 0x32, 0x9c, 0x47, 0x44, 0x44, 0x66,
	0x6d, 0x68, 0x8c, 0xda, 0x6e, 0x77, 0x13, 0xd7, 0x15, 0x11, 0x11, 0x7e, 0x84, 0x9a, 0x2a, 0x75,
	0x8d, 0xd4, 0x01, 0x69, 0xa8, 0x02, 0x34, 0x9f, 0xa3, 0xde, 0x92, 0xc4, 0x2c, 0x20, 0x92, 0xe7,
	0x42, 0x23, 0x0d, 0x3d, 0x65, 0x53, 0x06, 0xf0, 0x1c, 0x1d, 0xa6, 0xf4, 0x8b, 0x9c, 0xff, 0x4d,
	0x37, 0x81, 0xc6, 0xaa, 0xf7, 0x71, 0x57, 0xf1, 0x0c, 0x75, 0xfd, 0x32, 0x7c, 0xcd, 0x22, 0x60,
	0x3b, 0x7f, 0xaa, 0x80, 0x1d, 0xa3, 0x06, 0xc9, 0x32, 0x0d, 0xb4, 0x00, 0xa8, 0x93, 0x2c, 0x83,
	0xd6, 0x19, 0xfa, 0x1f, 0xee, 0x98, 0x53, 0xb1, 0x88, 0x65, 0x31, 0xa4, 0x0d, 0x4c, 0x4f, 0x35,
	0x5c, 0x5d, 0x07, 0xf6, 0x29, 0xea, 0xd0, 0x25, 0x0b, 0x68, 0xea, 0x53, 0xcd, 0x75, 0x80, 0x6b,
	0x97, 0x45, 0x80, 0x5e, 0xa0, 0x83, 0x2c, 0xe7, 0x19, 0x17, 0x34, 0x9f, 0x93, 0x20, 0xc8, 0xa9,
	0x10, 0x66, 0x57, 0xbd, 0xbd, 0xdb, 0x2b, 0xeb, 0xef, 0x74, 0x79, 0xfc, 0xfe, 0x76, 0x65, 0x19,
	0x77, 0x2b, 0xcb, 0xf8, 0xb9, 0xb2, 0x8c, 0x9b, 0xb5, 0x55, 0xb9, 0x5b, 0x5b, 0x95, 0xef, 0x6b,
	0xab, 0xf2, 0xe9, 0x32, 0x64, 0x32, 0x5a, 0x78, 0xb6, 0xcf, 0x13, 0xa7, 0xf8, 0x16, 0xe8, 0x9f,
	0x97, 0x22, 0xb8, 0x76, 0xfc, 0x98, 0xd1, 0x54, 0x3a, 0x61, 0x9e, 0xf9, 0x8e, 0x4c, 0x04, 0xcd,
	0x97, 0xcc, 0xa7, 0x5e, 0x0d, 0xfe, 0x16, 0xaf, 0x7e, 0x07, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x61,
	0xed, 0x16, 0x3d, 0x04, 0x00, 0x00,
}

func (m *Block) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Block) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Block) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LastCommit != nil {
		{
			size, err := m.LastCommit.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size, err := m.Header.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ProposerAddress) > 0 {
		i -= len(m.ProposerAddress)
		copy(dAtA[i:], m.ProposerAddress)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ProposerAddress)))
		i--
		dAtA[i] = 0x72
	}
	if len(m.EvidenceHash) > 0 {
		i -= len(m.EvidenceHash)
		copy(dAtA[i:], m.EvidenceHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.EvidenceHash)))
		i--
		dAtA[i] = 0x6a
	}
	if len(m.LastResultsHash) > 0 {
		i -= len(m.LastResultsHash)
		copy(dAtA[i:], m.LastResultsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastResultsHash)))
		i--
		dAtA[i] = 0x62
	}
	if len(m.AppHash) > 0 {
		i -= len(m.AppHash)
		copy(dAtA[i:], m.AppHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AppHash)))
		i--
		dAtA[i] = 0x5a
	}
	if len(m.ConsensusHash) > 0 {
		i -= len(m.ConsensusHash)
		copy(dAtA[i:], m.ConsensusHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ConsensusHash)))
		i--
		dAtA[i] = 0x52
	}
	if len(m.NextValidatorsHash) > 0 {
		i -= len(m.NextValidatorsHash)
		copy(dAtA[i:], m.NextValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.NextValidatorsHash)))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.ValidatorsHash) > 0 {
		i -= len(m.ValidatorsHash)
		copy(dAtA[i:], m.ValidatorsHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorsHash)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.DataHash) > 0 {
		i -= len(m.DataHash)
		copy(dAtA[i:], m.DataHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.DataHash)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.LastCommitHash) > 0 {
		i -= len(m.LastCommitHash)
		copy(dAtA[i:], m.LastCommitHash)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.LastCommitHash)))
		i--
		dAtA[i] = 0x32
	}
	{
		size, err := m.LastBlockId.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	n5, err5 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err5 != nil {
		return 0, err5
	}
	i -= n5
	i = encodeVarintTypes(dAtA, i, uint64(n5))
	i--
	dAtA[i] = 0x22
	if m.Height != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.Version.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTypes(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Block) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Header.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = m.Data.Size()
	n += 1 + l + sovTypes(uint64(l))
	if m.LastCommit != nil {
		l = m.LastCommit.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Version.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovTypes(uint64(m.Height))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	l = m.LastBlockId.Size()
	n += 1 + l + sovTypes(uint64(l))
	l = len(m.LastCommitHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.DataHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.NextValidatorsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ConsensusHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AppHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.LastResultsHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.EvidenceHash)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.ProposerAddress)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Block) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Block: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Block: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.LastCommit == nil {
				m.LastCommit = &types.Commit{}
			}
			if err := m.LastCommit.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Version.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastBlockId", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LastBlockId.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastCommitHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastCommitHash = append(m.LastCommitHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastCommitHash == nil {
				m.LastCommitHash = []byte{}
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DataHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DataHash = append(m.DataHash[:0], dAtA[iNdEx:postIndex]...)
			if m.DataHash == nil {
				m.DataHash = []byte{}
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorsHash = append(m.ValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorsHash == nil {
				m.ValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextValidatorsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NextValidatorsHash = append(m.NextValidatorsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.NextValidatorsHash == nil {
				m.NextValidatorsHash = []byte{}
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusHash = append(m.ConsensusHash[:0], dAtA[iNdEx:postIndex]...)
			if m.ConsensusHash == nil {
				m.ConsensusHash = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AppHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AppHash = append(m.AppHash[:0], dAtA[iNdEx:postIndex]...)
			if m.AppHash == nil {
				m.AppHash = []byte{}
			}
			iNdEx = postIndex
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastResultsHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LastResultsHash = append(m.LastResultsHash[:0], dAtA[iNdEx:postIndex]...)
			if m.LastResultsHash == nil {
				m.LastResultsHash = []byte{}
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvidenceHash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EvidenceHash = append(m.EvidenceHash[:0], dAtA[iNdEx:postIndex]...)
			if m.EvidenceHash == nil {
				m.EvidenceHash = []byte{}
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposerAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProposerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
