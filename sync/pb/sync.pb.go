// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sync.proto

/*
Package syncpb is a generated protocol buffer package.

It is generated from these files:
	sync.proto

It has these top-level messages:
	Chunks
	ChunkData
*/
package syncpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import corepb "github.com/nebulasio/go-nebulas/core/pb"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Chunks struct {
	ChunksRoot [][]byte `protobuf:"bytes,1,rep,name=chunks_root,json=chunksRoot" json:"chunks_root,omitempty"`
	Root       []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *Chunks) Reset()                    { *m = Chunks{} }
func (m *Chunks) String() string            { return proto.CompactTextString(m) }
func (*Chunks) ProtoMessage()               {}
func (*Chunks) Descriptor() ([]byte, []int) { return fileDescriptorSync, []int{0} }

func (m *Chunks) GetChunksRoot() [][]byte {
	if m != nil {
		return m.ChunksRoot
	}
	return nil
}

func (m *Chunks) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

type ChunkData struct {
	Blocks []*corepb.Block `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
	Root   []byte          `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
}

func (m *ChunkData) Reset()                    { *m = ChunkData{} }
func (m *ChunkData) String() string            { return proto.CompactTextString(m) }
func (*ChunkData) ProtoMessage()               {}
func (*ChunkData) Descriptor() ([]byte, []int) { return fileDescriptorSync, []int{1} }

func (m *ChunkData) GetBlocks() []*corepb.Block {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *ChunkData) GetRoot() []byte {
	if m != nil {
		return m.Root
	}
	return nil
}

func init() {
	proto.RegisterType((*Chunks)(nil), "syncpb.Chunks")
	proto.RegisterType((*ChunkData)(nil), "syncpb.ChunkData")
}

func init() { proto.RegisterFile("sync.proto", fileDescriptorSync) }

var fileDescriptorSync = []byte{
	// 176 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0xae, 0xcc, 0x4b,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0xb1, 0x0b, 0x92, 0xa4, 0x8c, 0xd3, 0x33,
	0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5, 0xf3, 0x52, 0x93, 0x4a, 0x73, 0x12, 0x8b,
	0x33, 0xf3, 0xf5, 0xd3, 0xf3, 0x75, 0xa1, 0x1c, 0xfd, 0xe4, 0xfc, 0xa2, 0x54, 0xfd, 0x82, 0x24,
	0xfd, 0xa4, 0x9c, 0xfc, 0xe4, 0x6c, 0x88, 0x66, 0x25, 0x5b, 0x2e, 0x36, 0xe7, 0x8c, 0xd2, 0xbc,
	0xec, 0x62, 0x21, 0x79, 0x2e, 0xee, 0x64, 0x30, 0x2b, 0xbe, 0x28, 0x3f, 0xbf, 0x44, 0x82, 0x51,
	0x81, 0x59, 0x83, 0x27, 0x88, 0x0b, 0x22, 0x14, 0x94, 0x9f, 0x5f, 0x22, 0x24, 0xc4, 0xc5, 0x02,
	0x96, 0x61, 0x52, 0x60, 0xd4, 0xe0, 0x09, 0x02, 0xb3, 0x95, 0xdc, 0xb8, 0x38, 0xc1, 0xda, 0x5d,
	0x12, 0x4b, 0x12, 0x85, 0x54, 0xb9, 0xd8, 0xc0, 0x46, 0x17, 0x83, 0x35, 0x73, 0x1b, 0xf1, 0xea,
	0x81, 0x6c, 0x2c, 0x48, 0xd2, 0x73, 0x02, 0x89, 0x06, 0x41, 0x25, 0xb1, 0x99, 0x93, 0xc4, 0x06,
	0x76, 0x8d, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x55, 0x61, 0xf1, 0xc2, 0xd8, 0x00, 0x00, 0x00,
}
