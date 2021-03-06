// Code generated by protoc-gen-go.
// source: github.com/google/trillian/storage/storagepb/storage.proto
// DO NOT EDIT!

/*
Package storagepb is a generated protocol buffer package.

It is generated from these files:
	github.com/google/trillian/storage/storagepb/storage.proto

It has these top-level messages:
	NodeIDProto
	SubtreeProto
*/
package storagepb

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

// NodeIDProto is the serialized form of NodeID. It's used only for persistence in storage.
// As this is long-term we prefer not to use a Go specific format.
type NodeIDProto struct {
	Path          []byte `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	PrefixLenBits int32  `protobuf:"varint,2,opt,name=prefix_len_bits,json=prefixLenBits" json:"prefix_len_bits,omitempty"`
}

func (m *NodeIDProto) Reset()                    { *m = NodeIDProto{} }
func (m *NodeIDProto) String() string            { return proto.CompactTextString(m) }
func (*NodeIDProto) ProtoMessage()               {}
func (*NodeIDProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// SubtreeProto contains nodes of a subtree.
type SubtreeProto struct {
	// subtree's prefix (must be a multiple of 8 bits)
	Prefix []byte `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// subtree's depth
	Depth    int32  `protobuf:"varint,2,opt,name=depth" json:"depth,omitempty"`
	RootHash []byte `protobuf:"bytes,3,opt,name=root_hash,json=rootHash,proto3" json:"root_hash,omitempty"`
	// map of suffix (within subtree) to subtree-leaf node hash
	Leaves map[string][]byte `protobuf:"bytes,4,rep,name=leaves" json:"leaves,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Map of suffix (within subtree) to subtree-internal node hash.
	// This structure is only used in RAM as a cache, the internal nodes of
	// the subtree are not generally stored.
	InternalNodes map[string][]byte `protobuf:"bytes,5,rep,name=internal_nodes,json=internalNodes" json:"internal_nodes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *SubtreeProto) Reset()                    { *m = SubtreeProto{} }
func (m *SubtreeProto) String() string            { return proto.CompactTextString(m) }
func (*SubtreeProto) ProtoMessage()               {}
func (*SubtreeProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SubtreeProto) GetLeaves() map[string][]byte {
	if m != nil {
		return m.Leaves
	}
	return nil
}

func (m *SubtreeProto) GetInternalNodes() map[string][]byte {
	if m != nil {
		return m.InternalNodes
	}
	return nil
}

func init() {
	proto.RegisterType((*NodeIDProto)(nil), "trillian.storage.storagepb.NodeIDProto")
	proto.RegisterType((*SubtreeProto)(nil), "trillian.storage.storagepb.SubtreeProto")
}

func init() {
	proto.RegisterFile("github.com/google/trillian/storage/storagepb/storage.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x49, 0xd3, 0x96, 0x7f, 0xa7, 0xed, 0x5f, 0x59, 0x44, 0x42, 0xbd, 0x94, 0x1e, 0xa4,
	0xa7, 0x14, 0x54, 0xa4, 0xb6, 0x17, 0x29, 0x0a, 0x16, 0x8a, 0x48, 0xbc, 0x79, 0x09, 0x1b, 0x3b,
	0x26, 0x8b, 0xeb, 0x6e, 0xd8, 0xdd, 0x16, 0xfb, 0x0d, 0xfc, 0xd8, 0x92, 0xcd, 0xa6, 0x04, 0x44,
	0xb0, 0xa7, 0xbc, 0x79, 0xbc, 0xf9, 0x25, 0x79, 0x0c, 0xcc, 0x52, 0x66, 0xb2, 0x4d, 0x12, 0xbe,
	0xca, 0x8f, 0x49, 0x2a, 0x65, 0xca, 0x71, 0x62, 0x14, 0xe3, 0x9c, 0x51, 0x31, 0xd1, 0x46, 0x2a,
	0x9a, 0x62, 0xf5, 0xcc, 0x93, 0x4a, 0x85, 0xb9, 0x92, 0x46, 0x92, 0x41, 0x95, 0x0c, 0x2b, 0x7f,
	0x9f, 0x1c, 0x2d, 0xa1, 0xfb, 0x28, 0xd7, 0xb8, 0xbc, 0x7b, 0xb2, 0x51, 0x02, 0xcd, 0x9c, 0x9a,
	0x2c, 0xf0, 0x86, 0xde, 0xb8, 0x17, 0x59, 0x4d, 0xce, 0xe1, 0x28, 0x57, 0xf8, 0xc6, 0x3e, 0x63,
	0x8e, 0x22, 0x4e, 0x98, 0xd1, 0x41, 0x63, 0xe8, 0x8d, 0x5b, 0x51, 0xbf, 0xb4, 0x57, 0x28, 0x16,
	0xcc, 0xe8, 0xd1, 0x97, 0x0f, 0xbd, 0xe7, 0x4d, 0x62, 0x14, 0x62, 0x09, 0x3b, 0x85, 0x76, 0x99,
	0x70, 0x38, 0x37, 0x91, 0x13, 0x68, 0xad, 0x31, 0x37, 0x99, 0xc3, 0x94, 0x03, 0x39, 0x83, 0x8e,
	0x92, 0xd2, 0xc4, 0x19, 0xd5, 0x59, 0xe0, 0xdb, 0x85, 0x7f, 0x85, 0xf1, 0x40, 0x75, 0x46, 0x56,
	0xd0, 0xe6, 0x48, 0xb7, 0xa8, 0x83, 0xe6, 0xd0, 0x1f, 0x77, 0x2f, 0xae, 0xc2, 0xdf, 0xff, 0x29,
	0xac, 0x7f, 0x44, 0xb8, 0xb2, 0x6b, 0xf7, 0xc2, 0xa8, 0x5d, 0xe4, 0x18, 0x24, 0x81, 0xff, 0x4c,
	0x18, 0x54, 0x82, 0xf2, 0x58, 0xc8, 0x35, 0xea, 0xa0, 0x65, 0xa9, 0xf3, 0x3f, 0x53, 0x97, 0x6e,
	0xbd, 0xe8, 0xce, 0xc1, 0xfb, 0xac, 0xee, 0x0d, 0x6e, 0xa0, 0x5b, 0x7b, 0x35, 0x39, 0x06, 0xff,
	0x1d, 0x77, 0xb6, 0x88, 0x4e, 0x54, 0xc8, 0xa2, 0x85, 0x2d, 0xe5, 0x1b, 0xb4, 0x2d, 0xf4, 0xa2,
	0x72, 0x98, 0x35, 0xa6, 0xde, 0xe0, 0x16, 0xc8, 0x4f, 0xfe, 0x21, 0x84, 0xc5, 0xf4, 0xe5, 0xfa,
	0x90, 0x7b, 0x99, 0xef, 0x55, 0xd2, 0xb6, 0x27, 0x73, 0xf9, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x01,
	0x6b, 0x32, 0xc6, 0x70, 0x02, 0x00, 0x00,
}
