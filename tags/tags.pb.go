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

type BaseResponse struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseResponse) Reset()         { *m = BaseResponse{} }
func (m *BaseResponse) String() string { return proto.CompactTextString(m) }
func (*BaseResponse) ProtoMessage()    {}
func (*BaseResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f12757892b65802, []int{2}
}

func (m *BaseResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseResponse.Unmarshal(m, b)
}
func (m *BaseResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseResponse.Marshal(b, m, deterministic)
}
func (m *BaseResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseResponse.Merge(m, src)
}
func (m *BaseResponse) XXX_Size() int {
	return xxx_messageInfo_BaseResponse.Size(m)
}
func (m *BaseResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BaseResponse proto.InternalMessageInfo

func (m *BaseResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *BaseResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Tag)(nil), "ffv.Tag")
	proto.RegisterType((*Tags)(nil), "ffv.Tags")
	proto.RegisterType((*BaseResponse)(nil), "ffv.BaseResponse")
}

func init() { proto.RegisterFile("tags/tags.proto", fileDescriptor_0f12757892b65802) }

var fileDescriptor_0f12757892b65802 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x86, 0xcd, 0x87, 0xb5, 0x19, 0x45, 0x71, 0x10, 0x09, 0xe2, 0xa1, 0x04, 0x0f, 0xa9, 0x4a,
	0x02, 0xf5, 0x26, 0x9e, 0x8a, 0x47, 0x4f, 0x31, 0x27, 0x6f, 0x9b, 0x74, 0x76, 0xb3, 0xd4, 0x34,
	0x21, 0xb3, 0x0d, 0xf8, 0xdf, 0xfc, 0x71, 0x92, 0x35, 0x16, 0x0f, 0x82, 0xbd, 0x2c, 0x33, 0xc3,
	0xc3, 0xf3, 0x2e, 0xbc, 0x70, 0x66, 0x84, 0xe2, 0x74, 0x78, 0x92, 0xb6, 0x6b, 0x4c, 0x83, 0x9e,
	0x94, 0x7d, 0x34, 0x07, 0x2f, 0x17, 0x0a, 0x11, 0xfc, 0x8d, 0xa8, 0x29, 0x74, 0x66, 0x4e, 0x1c,
	0x64, 0x76, 0xc6, 0x53, 0x70, 0xf5, 0x2a, 0x74, 0xed, 0xc5, 0xd5, 0xab, 0xe8, 0x06, 0xfc, 0x5c,
	0x28, 0xc6, 0x6b, 0xf0, 0x07, 0x4b, 0xe8, 0xcc, 0xbc, 0xf8, 0x78, 0x31, 0x4d, 0xa4, 0xec, 0x93,
	0x5c, 0xa8, 0xcc, 0x5e, 0xa3, 0x27, 0x38, 0x59, 0x0a, 0xa6, 0x8c, 0xb8, 0x6d, 0x36, 0x4c, 0x78,
	0x09, 0x13, 0x36, 0xc2, 0x6c, 0x79, 0x74, 0x8f, 0x1b, 0x5e, 0xc0, 0x21, 0x75, 0x5d, 0xd3, 0x8d,
	0x01, 0xdf, 0xcb, 0xe2, 0xd3, 0x01, 0xc8, 0x85, 0x7a, 0xa5, 0xae, 0xd7, 0x25, 0x61, 0x0c, 0x47,
	0x2f, 0x7a, 0x4d, 0xc3, 0x0f, 0x77, 0x39, 0x57, 0xe7, 0x76, 0xfa, 0x1d, 0x12, 0x1d, 0xe0, 0x2d,
	0x04, 0xcf, 0x9a, 0xf7, 0x65, 0xa7, 0xa3, 0x95, 0x31, 0xf8, 0x41, 0xf9, 0x6f, 0xf6, 0x1e, 0x60,
	0xe7, 0xfd, 0x97, 0x5e, 0xde, 0xbd, 0xcd, 0x2b, 0x63, 0x5a, 0x7e, 0x4c, 0x53, 0xa5, 0x4d, 0xb5,
	0x2d, 0x92, 0xb2, 0xa9, 0xd3, 0x5a, 0x97, 0x15, 0x89, 0xf7, 0x5a, 0xaf, 0xe9, 0xa3, 0x48, 0xa5,
	0xec, 0x6d, 0x0b, 0xc5, 0xc4, 0xd6, 0xf0, 0xf0, 0x15, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x86, 0x51,
	0x28, 0x99, 0x01, 0x00, 0x00,
}
