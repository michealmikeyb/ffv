// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tags/tags.proto

package tags

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

type Tag struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tag) Reset()         { *m = Tag{} }
func (m *Tag) String() string { return proto.CompactTextString(m) }
func (*Tag) ProtoMessage()    {}
func (*Tag) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f12757892b65802, []int{0}
}

func (m *Tag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tag.Unmarshal(m, b)
}
func (m *Tag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tag.Marshal(b, m, deterministic)
}
func (m *Tag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tag.Merge(m, src)
}
func (m *Tag) XXX_Size() int {
	return xxx_messageInfo_Tag.Size(m)
}
func (m *Tag) XXX_DiscardUnknown() {
	xxx_messageInfo_Tag.DiscardUnknown(m)
}

var xxx_messageInfo_Tag proto.InternalMessageInfo

func (m *Tag) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Tag) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Tags struct {
	Tags                 []*Tag   `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Tags) Reset()         { *m = Tags{} }
func (m *Tags) String() string { return proto.CompactTextString(m) }
func (*Tags) ProtoMessage()    {}
func (*Tags) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f12757892b65802, []int{1}
}

func (m *Tags) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Tags.Unmarshal(m, b)
}
func (m *Tags) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Tags.Marshal(b, m, deterministic)
}
func (m *Tags) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Tags.Merge(m, src)
}
func (m *Tags) XXX_Size() int {
	return xxx_messageInfo_Tags.Size(m)
}
func (m *Tags) XXX_DiscardUnknown() {
	xxx_messageInfo_Tags.DiscardUnknown(m)
}

var xxx_messageInfo_Tags proto.InternalMessageInfo

func (m *Tags) GetTags() []*Tag {
	if m != nil {
		return m.Tags
	}
	return nil
}

type TagBaseResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagBaseResponse) Reset()         { *m = TagBaseResponse{} }
func (m *TagBaseResponse) String() string { return proto.CompactTextString(m) }
func (*TagBaseResponse) ProtoMessage()    {}
func (*TagBaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f12757892b65802, []int{2}
}

func (m *TagBaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TagBaseResponse.Unmarshal(m, b)
}
func (m *TagBaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TagBaseResponse.Marshal(b, m, deterministic)
}
func (m *TagBaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagBaseResponse.Merge(m, src)
}
func (m *TagBaseResponse) XXX_Size() int {
	return xxx_messageInfo_TagBaseResponse.Size(m)
}
func (m *TagBaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TagBaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TagBaseResponse proto.InternalMessageInfo

func (m *TagBaseResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *TagBaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Tag)(nil), "ffv.Tag")
	proto.RegisterType((*Tags)(nil), "ffv.Tags")
	proto.RegisterType((*TagBaseResponse)(nil), "ffv.TagBaseResponse")
}

func init() { proto.RegisterFile("tags/tags.proto", fileDescriptor_0f12757892b65802) }

var fileDescriptor_0f12757892b65802 = []byte{
	// 264 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0x4d, 0x13, 0x6b, 0x33, 0x82, 0x85, 0xa5, 0x48, 0x10, 0x0f, 0x25, 0x78, 0x68, 0x29,
	0x66, 0xa1, 0xde, 0xbc, 0x08, 0xc5, 0xa3, 0xa7, 0x98, 0x93, 0xb7, 0x49, 0x3a, 0xbb, 0x59, 0x6a,
	0x9a, 0x90, 0xd9, 0x06, 0xfc, 0x87, 0xfe, 0x2c, 0xc9, 0x9a, 0x16, 0x2f, 0x82, 0x5e, 0x96, 0x99,
	0xc7, 0xdb, 0xef, 0x0d, 0x3c, 0x98, 0x5a, 0xd4, 0x2c, 0xfb, 0x27, 0x69, 0xda, 0xda, 0xd6, 0xc2,
	0x57, 0xaa, 0x8b, 0x97, 0xe0, 0x67, 0xa8, 0x85, 0x80, 0x60, 0x8f, 0x15, 0x45, 0xde, 0xdc, 0x5b,
	0x84, 0xa9, 0x9b, 0xc5, 0x15, 0x8c, 0xcc, 0x36, 0x1a, 0x39, 0x65, 0x64, 0xb6, 0xf1, 0x1d, 0x04,
	0x19, 0x6a, 0x16, 0xb7, 0x10, 0xf4, 0x94, 0xc8, 0x9b, 0xfb, 0x8b, 0xcb, 0xf5, 0x24, 0x51, 0xaa,
	0x4b, 0x32, 0xd4, 0xa9, 0x53, 0xe3, 0x27, 0x98, 0x66, 0xa8, 0x37, 0xc8, 0x94, 0x12, 0x37, 0xf5,
	0x9e, 0x49, 0x5c, 0xc3, 0x98, 0x2d, 0xda, 0x03, 0x0f, 0xf8, 0x61, 0x13, 0x33, 0x38, 0xa7, 0xb6,
	0xad, 0xdb, 0x21, 0xe3, 0x7b, 0x59, 0x7f, 0x7a, 0x00, 0x19, 0xea, 0x57, 0x6a, 0x3b, 0x53, 0x90,
	0x58, 0xc1, 0xc5, 0x8b, 0xd9, 0x51, 0x7f, 0xe4, 0x29, 0xea, 0x66, 0x76, 0x9c, 0x7e, 0xe6, 0xc4,
	0x67, 0xe2, 0x1e, 0xc2, 0x67, 0xc3, 0xff, 0xb0, 0x4f, 0x06, 0x36, 0x8b, 0xf0, 0xe8, 0xe1, 0x5f,
	0xed, 0x12, 0xe0, 0x44, 0xff, 0xcb, 0x87, 0xcd, 0xea, 0x6d, 0x59, 0x5a, 0xdb, 0xf0, 0xa3, 0x94,
	0xda, 0xd8, 0xf2, 0x90, 0x27, 0x45, 0x5d, 0xc9, 0xca, 0x14, 0x25, 0xe1, 0x7b, 0x65, 0x76, 0xf4,
	0x91, 0x4b, 0xa5, 0x3a, 0x57, 0x4a, 0x3e, 0x76, 0xad, 0x3c, 0x7c, 0x05, 0x00, 0x00, 0xff, 0xff,
	0x76, 0x4c, 0x02, 0x70, 0xa8, 0x01, 0x00, 0x00,
}
