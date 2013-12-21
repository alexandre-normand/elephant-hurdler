// Code generated by protoc-gen-go.
// source: ClientDatanodeProtocol.proto
// DO NOT EDIT!

package ClientDatanodeProtocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"
import hdfs "hdfs.pb"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type GetReplicaVisibleLengthRequestProto struct {
	Block            *hdfs.ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *GetReplicaVisibleLengthRequestProto) Reset()         { *m = GetReplicaVisibleLengthRequestProto{} }
func (m *GetReplicaVisibleLengthRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthRequestProto) ProtoMessage()    {}

func (m *GetReplicaVisibleLengthRequestProto) GetBlock() *hdfs.ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

type GetReplicaVisibleLengthResponseProto struct {
	Length           *uint64 `protobuf:"varint,1,req,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetReplicaVisibleLengthResponseProto) Reset()         { *m = GetReplicaVisibleLengthResponseProto{} }
func (m *GetReplicaVisibleLengthResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthResponseProto) ProtoMessage()    {}

func (m *GetReplicaVisibleLengthResponseProto) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

type RefreshNamenodesRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesRequestProto) Reset()         { *m = RefreshNamenodesRequestProto{} }
func (m *RefreshNamenodesRequestProto) String() string { return proto.CompactTextString(m) }
func (*RefreshNamenodesRequestProto) ProtoMessage()    {}

type RefreshNamenodesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesResponseProto) Reset()         { *m = RefreshNamenodesResponseProto{} }
func (m *RefreshNamenodesResponseProto) String() string { return proto.CompactTextString(m) }
func (*RefreshNamenodesResponseProto) ProtoMessage()    {}

type DeleteBlockPoolRequestProto struct {
	BlockPool        *string `protobuf:"bytes,1,req,name=blockPool" json:"blockPool,omitempty"`
	Force            *bool   `protobuf:"varint,2,req,name=force" json:"force,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteBlockPoolRequestProto) Reset()         { *m = DeleteBlockPoolRequestProto{} }
func (m *DeleteBlockPoolRequestProto) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockPoolRequestProto) ProtoMessage()    {}

func (m *DeleteBlockPoolRequestProto) GetBlockPool() string {
	if m != nil && m.BlockPool != nil {
		return *m.BlockPool
	}
	return ""
}

func (m *DeleteBlockPoolRequestProto) GetForce() bool {
	if m != nil && m.Force != nil {
		return *m.Force
	}
	return false
}

type DeleteBlockPoolResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DeleteBlockPoolResponseProto) Reset()         { *m = DeleteBlockPoolResponseProto{} }
func (m *DeleteBlockPoolResponseProto) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockPoolResponseProto) ProtoMessage()    {}

type GetBlockLocalPathInfoRequestProto struct {
	Block            *hdfs.ExtendedBlockProto        `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	Token            *hdfs.BlockTokenIdentifierProto `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *GetBlockLocalPathInfoRequestProto) Reset()         { *m = GetBlockLocalPathInfoRequestProto{} }
func (m *GetBlockLocalPathInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoRequestProto) ProtoMessage()    {}

func (m *GetBlockLocalPathInfoRequestProto) GetBlock() *hdfs.ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoRequestProto) GetToken() *hdfs.BlockTokenIdentifierProto {
	if m != nil {
		return m.Token
	}
	return nil
}

type GetBlockLocalPathInfoResponseProto struct {
	Block            *hdfs.ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	LocalPath        *string                  `protobuf:"bytes,2,req,name=localPath" json:"localPath,omitempty"`
	LocalMetaPath    *string                  `protobuf:"bytes,3,req,name=localMetaPath" json:"localMetaPath,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *GetBlockLocalPathInfoResponseProto) Reset()         { *m = GetBlockLocalPathInfoResponseProto{} }
func (m *GetBlockLocalPathInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoResponseProto) ProtoMessage()    {}

func (m *GetBlockLocalPathInfoResponseProto) GetBlock() *hdfs.ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalPath() string {
	if m != nil && m.LocalPath != nil {
		return *m.LocalPath
	}
	return ""
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalMetaPath() string {
	if m != nil && m.LocalMetaPath != nil {
		return *m.LocalMetaPath
	}
	return ""
}

type GetHdfsBlockLocationsRequestProto struct {
	Blocks           []*hdfs.ExtendedBlockProto        `protobuf:"bytes,1,rep,name=blocks" json:"blocks,omitempty"`
	Tokens           []*hdfs.BlockTokenIdentifierProto `protobuf:"bytes,2,rep,name=tokens" json:"tokens,omitempty"`
	XXX_unrecognized []byte                            `json:"-"`
}

func (m *GetHdfsBlockLocationsRequestProto) Reset()         { *m = GetHdfsBlockLocationsRequestProto{} }
func (m *GetHdfsBlockLocationsRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetHdfsBlockLocationsRequestProto) ProtoMessage()    {}

func (m *GetHdfsBlockLocationsRequestProto) GetBlocks() []*hdfs.ExtendedBlockProto {
	if m != nil {
		return m.Blocks
	}
	return nil
}

func (m *GetHdfsBlockLocationsRequestProto) GetTokens() []*hdfs.BlockTokenIdentifierProto {
	if m != nil {
		return m.Tokens
	}
	return nil
}

type GetHdfsBlockLocationsResponseProto struct {
	VolumeIds        [][]byte `protobuf:"bytes,1,rep,name=volumeIds" json:"volumeIds,omitempty"`
	VolumeIndexes    []uint32 `protobuf:"varint,2,rep,name=volumeIndexes" json:"volumeIndexes,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *GetHdfsBlockLocationsResponseProto) Reset()         { *m = GetHdfsBlockLocationsResponseProto{} }
func (m *GetHdfsBlockLocationsResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetHdfsBlockLocationsResponseProto) ProtoMessage()    {}

func (m *GetHdfsBlockLocationsResponseProto) GetVolumeIds() [][]byte {
	if m != nil {
		return m.VolumeIds
	}
	return nil
}

func (m *GetHdfsBlockLocationsResponseProto) GetVolumeIndexes() []uint32 {
	if m != nil {
		return m.VolumeIndexes
	}
	return nil
}

func init() {
}