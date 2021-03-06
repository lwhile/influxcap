// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pb/influxcap.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/influxcap.proto

It has these top-level messages:
	JoinRequest
*/
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

type JoinRequest struct {
	Id  string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Url string `protobuf:"bytes,2,opt,name=url" json:"url,omitempty"`
}

func (m *JoinRequest) Reset()                    { *m = JoinRequest{} }
func (m *JoinRequest) String() string            { return proto.CompactTextString(m) }
func (*JoinRequest) ProtoMessage()               {}
func (*JoinRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *JoinRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *JoinRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func init() {
	proto.RegisterType((*JoinRequest)(nil), "pb.JoinRequest")
}

func init() { proto.RegisterFile("pb/influxcap.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 97 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2a, 0x48, 0xd2, 0xcf,
	0xcc, 0x4b, 0xcb, 0x29, 0xad, 0x48, 0x4e, 0x2c, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2a, 0x48, 0x52, 0xd2, 0xe7, 0xe2, 0xf6, 0xca, 0xcf, 0xcc, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d,
	0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca,
	0x4c, 0x11, 0x12, 0xe0, 0x62, 0x2e, 0x2d, 0xca, 0x91, 0x60, 0x02, 0x0b, 0x80, 0x98, 0x49, 0x6c,
	0x60, 0xbd, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x54, 0x70, 0xdd, 0x59, 0x51, 0x00, 0x00,
	0x00,
}
