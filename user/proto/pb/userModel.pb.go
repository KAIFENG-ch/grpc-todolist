// Code generated by protoc-gen-go. DO NOT EDIT.
// source: userModel.proto

package pb

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

type UserModel struct {
	Id                   uint32   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserName             string   `protobuf:"bytes,2,opt,name=user_name,json=userName,proto3" json:"user_name,omitempty"`
	CreatedAt            int64    `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt            int64    `protobuf:"varint,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt            int64    `protobuf:"varint,5,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
	Status               string   `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Birthday             string   `protobuf:"bytes,8,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Signature            string   `protobuf:"bytes,9,opt,name=signature,proto3" json:"signature,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserModel) Reset()         { *m = UserModel{} }
func (m *UserModel) String() string { return proto.CompactTextString(m) }
func (*UserModel) ProtoMessage()    {}
func (*UserModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb74a2483174a702, []int{0}
}

func (m *UserModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserModel.Unmarshal(m, b)
}
func (m *UserModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserModel.Marshal(b, m, deterministic)
}
func (m *UserModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserModel.Merge(m, src)
}
func (m *UserModel) XXX_Size() int {
	return xxx_messageInfo_UserModel.Size(m)
}
func (m *UserModel) XXX_DiscardUnknown() {
	xxx_messageInfo_UserModel.DiscardUnknown(m)
}

var xxx_messageInfo_UserModel proto.InternalMessageInfo

func (m *UserModel) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UserModel) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *UserModel) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *UserModel) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *UserModel) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func (m *UserModel) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserModel) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserModel) GetBirthday() string {
	if m != nil {
		return m.Birthday
	}
	return ""
}

func (m *UserModel) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*UserModel)(nil), "service.UserModel")
}

func init() { proto.RegisterFile("userModel.proto", fileDescriptor_fb74a2483174a702) }

var fileDescriptor_fb74a2483174a702 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd0, 0x4d, 0x4b, 0xc4, 0x30,
	0x10, 0xc6, 0x71, 0xd2, 0xdd, 0xed, 0x36, 0x03, 0x2a, 0x0c, 0x22, 0xc1, 0x17, 0x28, 0x9e, 0x7a,
	0x5a, 0x0f, 0x7e, 0x82, 0xf5, 0xae, 0x87, 0x82, 0x17, 0x2f, 0xcb, 0x74, 0x33, 0x68, 0xa0, 0x6f,
	0x24, 0x53, 0xc1, 0x0f, 0x2f, 0x48, 0x9b, 0x14, 0x8f, 0xcf, 0xff, 0x47, 0x48, 0x08, 0x5c, 0x4d,
	0x81, 0xfd, 0xeb, 0x60, 0xb9, 0x3d, 0x8c, 0x7e, 0x90, 0x01, 0xf7, 0x81, 0xfd, 0xb7, 0x3b, 0xf3,
	0xe3, 0xaf, 0x02, 0xfd, 0xbe, 0x22, 0x5e, 0x42, 0xe6, 0xac, 0x51, 0xa5, 0xaa, 0x2e, 0xea, 0xcc,
	0x59, 0xbc, 0x03, 0x3d, 0x9f, 0x3c, 0xf5, 0xd4, 0xb1, 0xc9, 0x4a, 0x55, 0xe9, 0xba, 0x98, 0xc3,
	0x1b, 0x75, 0x8c, 0x0f, 0x00, 0x67, 0xcf, 0x24, 0x6c, 0x4f, 0x24, 0x66, 0x53, 0xaa, 0x6a, 0x53,
	0xeb, 0x54, 0x8e, 0x32, 0xf3, 0x34, 0xda, 0x95, 0xb7, 0x91, 0x53, 0x89, 0x6c, 0xb9, 0xe5, 0xc4,
	0xbb, 0xc8, 0xa9, 0x1c, 0x05, 0xaf, 0x61, 0xc7, 0x1d, 0xb9, 0xd6, 0xe4, 0xcb, 0xad, 0x71, 0xe0,
	0x0d, 0xe4, 0x41, 0x48, 0xa6, 0x60, 0xf6, 0x4b, 0x4e, 0x0b, 0x6f, 0xa1, 0x68, 0x9c, 0x97, 0x2f,
	0x4b, 0x3f, 0xa6, 0x88, 0xcf, 0x5c, 0x37, 0xde, 0x83, 0x0e, 0xee, 0xb3, 0x27, 0x99, 0x3c, 0x1b,
	0xbd, 0xe0, 0x7f, 0x78, 0xc9, 0x3f, 0xb6, 0x87, 0xa7, 0xb1, 0x69, 0xf2, 0xe5, 0x5f, 0x9e, 0xff,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xc3, 0xec, 0xae, 0x90, 0x2a, 0x01, 0x00, 0x00,
}