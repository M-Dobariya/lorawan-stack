// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/contact_info.proto

package ttnpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/TheThingsIndustries/protoc-gen-go-json/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	golang_proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
	reflect "reflect"
	strconv "strconv"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ContactType int32

const (
	CONTACT_TYPE_OTHER     ContactType = 0
	CONTACT_TYPE_ABUSE     ContactType = 1
	CONTACT_TYPE_BILLING   ContactType = 2
	CONTACT_TYPE_TECHNICAL ContactType = 3
)

var ContactType_name = map[int32]string{
	0: "CONTACT_TYPE_OTHER",
	1: "CONTACT_TYPE_ABUSE",
	2: "CONTACT_TYPE_BILLING",
	3: "CONTACT_TYPE_TECHNICAL",
}

var ContactType_value = map[string]int32{
	"CONTACT_TYPE_OTHER":     0,
	"CONTACT_TYPE_ABUSE":     1,
	"CONTACT_TYPE_BILLING":   2,
	"CONTACT_TYPE_TECHNICAL": 3,
}

func (ContactType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3baa36b7c4d52524, []int{0}
}

type ContactMethod int32

const (
	CONTACT_METHOD_OTHER ContactMethod = 0
	CONTACT_METHOD_EMAIL ContactMethod = 1
	CONTACT_METHOD_PHONE ContactMethod = 2
)

var ContactMethod_name = map[int32]string{
	0: "CONTACT_METHOD_OTHER",
	1: "CONTACT_METHOD_EMAIL",
	2: "CONTACT_METHOD_PHONE",
}

var ContactMethod_value = map[string]int32{
	"CONTACT_METHOD_OTHER": 0,
	"CONTACT_METHOD_EMAIL": 1,
	"CONTACT_METHOD_PHONE": 2,
}

func (ContactMethod) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3baa36b7c4d52524, []int{1}
}

type ContactInfo struct {
	ContactType          ContactType      `protobuf:"varint,1,opt,name=contact_type,json=contactType,proto3,enum=ttn.lorawan.v3.ContactType" json:"contact_type,omitempty"`
	ContactMethod        ContactMethod    `protobuf:"varint,2,opt,name=contact_method,json=contactMethod,proto3,enum=ttn.lorawan.v3.ContactMethod" json:"contact_method,omitempty"`
	Value                string           `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Public               bool             `protobuf:"varint,4,opt,name=public,proto3" json:"public,omitempty"`
	ValidatedAt          *types.Timestamp `protobuf:"bytes,5,opt,name=validated_at,json=validatedAt,proto3" json:"validated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ContactInfo) Reset()      { *m = ContactInfo{} }
func (*ContactInfo) ProtoMessage() {}
func (*ContactInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_3baa36b7c4d52524, []int{0}
}
func (m *ContactInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfo.Unmarshal(m, b)
}
func (m *ContactInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfo.Marshal(b, m, deterministic)
}
func (m *ContactInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfo.Merge(m, src)
}
func (m *ContactInfo) XXX_Size() int {
	return xxx_messageInfo_ContactInfo.Size(m)
}
func (m *ContactInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfo proto.InternalMessageInfo

func (m *ContactInfo) GetContactType() ContactType {
	if m != nil {
		return m.ContactType
	}
	return CONTACT_TYPE_OTHER
}

func (m *ContactInfo) GetContactMethod() ContactMethod {
	if m != nil {
		return m.ContactMethod
	}
	return CONTACT_METHOD_OTHER
}

func (m *ContactInfo) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *ContactInfo) GetPublic() bool {
	if m != nil {
		return m.Public
	}
	return false
}

func (m *ContactInfo) GetValidatedAt() *types.Timestamp {
	if m != nil {
		return m.ValidatedAt
	}
	return nil
}

type ContactInfoValidation struct {
	Id                   string             `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token                string             `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Entity               *EntityIdentifiers `protobuf:"bytes,3,opt,name=entity,proto3" json:"entity,omitempty"`
	ContactInfo          []*ContactInfo     `protobuf:"bytes,4,rep,name=contact_info,json=contactInfo,proto3" json:"contact_info,omitempty"`
	CreatedAt            *types.Timestamp   `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ExpiresAt            *types.Timestamp   `protobuf:"bytes,6,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ContactInfoValidation) Reset()      { *m = ContactInfoValidation{} }
func (*ContactInfoValidation) ProtoMessage() {}
func (*ContactInfoValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_3baa36b7c4d52524, []int{1}
}
func (m *ContactInfoValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContactInfoValidation.Unmarshal(m, b)
}
func (m *ContactInfoValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContactInfoValidation.Marshal(b, m, deterministic)
}
func (m *ContactInfoValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContactInfoValidation.Merge(m, src)
}
func (m *ContactInfoValidation) XXX_Size() int {
	return xxx_messageInfo_ContactInfoValidation.Size(m)
}
func (m *ContactInfoValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ContactInfoValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ContactInfoValidation proto.InternalMessageInfo

func (m *ContactInfoValidation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ContactInfoValidation) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *ContactInfoValidation) GetEntity() *EntityIdentifiers {
	if m != nil {
		return m.Entity
	}
	return nil
}

func (m *ContactInfoValidation) GetContactInfo() []*ContactInfo {
	if m != nil {
		return m.ContactInfo
	}
	return nil
}

func (m *ContactInfoValidation) GetCreatedAt() *types.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ContactInfoValidation) GetExpiresAt() *types.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func init() {
	proto.RegisterEnum("ttn.lorawan.v3.ContactType", ContactType_name, ContactType_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ContactType", ContactType_name, ContactType_value)
	proto.RegisterEnum("ttn.lorawan.v3.ContactMethod", ContactMethod_name, ContactMethod_value)
	golang_proto.RegisterEnum("ttn.lorawan.v3.ContactMethod", ContactMethod_name, ContactMethod_value)
	proto.RegisterType((*ContactInfo)(nil), "ttn.lorawan.v3.ContactInfo")
	golang_proto.RegisterType((*ContactInfo)(nil), "ttn.lorawan.v3.ContactInfo")
	proto.RegisterType((*ContactInfoValidation)(nil), "ttn.lorawan.v3.ContactInfoValidation")
	golang_proto.RegisterType((*ContactInfoValidation)(nil), "ttn.lorawan.v3.ContactInfoValidation")
}

func init() {
	proto.RegisterFile("lorawan-stack/api/contact_info.proto", fileDescriptor_3baa36b7c4d52524)
}
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/contact_info.proto", fileDescriptor_3baa36b7c4d52524)
}

var fileDescriptor_3baa36b7c4d52524 = []byte{
	// 779 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0xcf, 0x8f, 0xdb, 0x44,
	0x14, 0xce, 0x78, 0x77, 0xa3, 0xdd, 0xc9, 0x36, 0x32, 0x03, 0x44, 0x5e, 0x97, 0xba, 0x21, 0x05,
	0x29, 0x8a, 0x14, 0x5b, 0xca, 0x9e, 0x8a, 0x04, 0xc2, 0x09, 0x16, 0xb1, 0x94, 0x4d, 0x2a, 0x63,
	0x90, 0xe0, 0x12, 0x39, 0xce, 0xc4, 0x19, 0x92, 0xcc, 0x18, 0x7b, 0x92, 0xd6, 0xb7, 0x6a, 0x4f,
	0x08, 0x09, 0x09, 0x89, 0x3b, 0x17, 0x6e, 0xfd, 0x2b, 0x38, 0xf0, 0x1f, 0x70, 0xe1, 0xca, 0x96,
	0x03, 0xea, 0x89, 0x73, 0x4f, 0x55, 0x26, 0x4e, 0x62, 0x37, 0xea, 0x8f, 0x9b, 0xe7, 0x7d, 0xdf,
	0x7b, 0x6f, 0xe6, 0xfb, 0x9e, 0x1f, 0xfc, 0x68, 0xce, 0x22, 0xef, 0xa1, 0x47, 0x9b, 0x31, 0xf7,
	0xfc, 0x99, 0xe1, 0x85, 0xc4, 0xf0, 0x19, 0xe5, 0x9e, 0xcf, 0x87, 0x84, 0x4e, 0x98, 0x1e, 0x46,
	0x8c, 0x33, 0x54, 0xe6, 0x9c, 0xea, 0x29, 0x53, 0x5f, 0x5d, 0xaa, 0x66, 0x40, 0xf8, 0x74, 0x39,
	0xd2, 0x7d, 0xb6, 0x30, 0x30, 0x5d, 0xb1, 0x24, 0x8c, 0xd8, 0xa3, 0xc4, 0x10, 0x64, 0xbf, 0x19,
	0x60, 0xda, 0x5c, 0x79, 0x73, 0x32, 0xf6, 0x38, 0x36, 0x0e, 0x3e, 0x36, 0x25, 0xd5, 0x66, 0xa6,
	0x44, 0xc0, 0x02, 0xb6, 0x49, 0x1e, 0x2d, 0x27, 0xe2, 0x24, 0x0e, 0xe2, 0x2b, 0xa5, 0x77, 0x32,
	0x74, 0x77, 0x8a, 0xdd, 0x29, 0xa1, 0x41, 0x6c, 0xd3, 0xf1, 0x32, 0xe6, 0x11, 0xc1, 0x71, 0xb6,
	0x75, 0xc0, 0x9a, 0xdf, 0xc7, 0x8c, 0x1a, 0x1e, 0xa5, 0x8c, 0x7b, 0x9c, 0x30, 0x1a, 0xa7, 0x45,
	0x3e, 0x08, 0x18, 0x0b, 0xe6, 0x58, 0xbc, 0xf2, 0x10, 0xbd, 0x9d, 0xa2, 0xbb, 0x8b, 0xe0, 0x45,
	0xc8, 0x93, 0x14, 0xbc, 0xfb, 0x32, 0xc8, 0xc9, 0x02, 0xc7, 0xdc, 0x5b, 0x84, 0x29, 0xe1, 0xde,
	0xa1, 0x90, 0x64, 0x8c, 0x29, 0x27, 0x13, 0x82, 0xa3, 0xb4, 0x45, 0xed, 0x37, 0x09, 0x96, 0x3a,
	0x1b, 0x79, 0x6d, 0x3a, 0x61, 0xa8, 0x0b, 0xcf, 0xb7, 0x6a, 0xf3, 0x24, 0xc4, 0x0a, 0xa8, 0x82,
	0x7a, 0xb9, 0x75, 0x5b, 0xcf, 0xcb, 0xad, 0xa7, 0x29, 0x6e, 0x12, 0xe2, 0xf6, 0xe9, 0xf3, 0xf6,
	0xc9, 0x35, 0x90, 0x64, 0xe0, 0x94, 0xfc, 0x7d, 0x18, 0xf5, 0x61, 0x79, 0x5b, 0x69, 0x81, 0xf9,
	0x94, 0x8d, 0x15, 0x49, 0xd4, 0xba, 0xf3, 0x8a, 0x5a, 0x57, 0x82, 0x94, 0xa9, 0x76, 0xcb, 0xcf,
	0x02, 0x48, 0x83, 0x27, 0x2b, 0x6f, 0xbe, 0xc4, 0xca, 0x51, 0x15, 0xd4, 0xcf, 0x04, 0x2f, 0x3a,
	0x52, 0x1e, 0x4b, 0xce, 0x26, 0x8c, 0x2a, 0xb0, 0x18, 0x2e, 0x47, 0x73, 0xe2, 0x2b, 0xc7, 0x55,
	0x50, 0x3f, 0x75, 0xd2, 0x13, 0xfa, 0x14, 0x9e, 0x6f, 0x8d, 0x1e, 0x0f, 0x3d, 0xae, 0x9c, 0x54,
	0x41, 0xbd, 0xd4, 0x52, 0xf5, 0x8d, 0x7c, 0xfa, 0x56, 0x3e, 0xdd, 0xdd, 0xca, 0xe7, 0x94, 0x76,
	0x7c, 0x93, 0xd7, 0xfe, 0x94, 0xe0, 0xfb, 0x19, 0x81, 0xbe, 0xd9, 0x40, 0x84, 0x51, 0x74, 0x01,
	0x25, 0x32, 0x16, 0x02, 0x9d, 0xb5, 0xcf, 0x9e, 0xb7, 0x8b, 0xd1, 0xb1, 0x0c, 0x94, 0xcf, 0x1d,
	0x89, 0x8c, 0xd1, 0x5d, 0x78, 0xc2, 0xd9, 0x0c, 0x53, 0xf1, 0xe4, 0x1c, 0xba, 0x89, 0xa3, 0xfb,
	0xb0, 0xb8, 0x76, 0x82, 0x27, 0xe2, 0x35, 0xa5, 0xd6, 0x87, 0x2f, 0x8b, 0x62, 0x09, 0xd4, 0xde,
	0xfb, 0xe5, 0xa4, 0x09, 0xe8, 0xb3, 0xbd, 0x43, 0xeb, 0xff, 0x41, 0x39, 0xae, 0x1e, 0xd5, 0x4b,
	0xaf, 0x74, 0x68, 0x7d, 0xe7, 0x9d, 0x2f, 0xc2, 0xe1, 0xfb, 0x10, 0xfa, 0x11, 0x7e, 0x7b, 0x35,
	0xce, 0x52, 0xb6, 0xc9, 0xd7, 0xa9, 0xf8, 0x51, 0x48, 0x22, 0x1c, 0xaf, 0x53, 0x8b, 0x6f, 0x4e,
	0x4d, 0xd9, 0x26, 0x6f, 0xfc, 0x0c, 0x76, 0x73, 0x26, 0xa6, 0xa3, 0x02, 0x51, 0x67, 0xd0, 0x77,
	0xcd, 0x8e, 0x3b, 0x74, 0xbf, 0x7d, 0x60, 0x0d, 0x07, 0x6e, 0xd7, 0x72, 0xe4, 0xc2, 0x41, 0xdc,
	0x6c, 0x7f, 0xfd, 0x95, 0x25, 0x03, 0xa4, 0xc0, 0xf7, 0x72, 0xf1, 0xb6, 0xdd, 0xeb, 0xd9, 0xfd,
	0x2f, 0x65, 0x09, 0xa9, 0xb0, 0x92, 0x43, 0x5c, 0xab, 0xd3, 0xed, 0xdb, 0x1d, 0xb3, 0x27, 0x1f,
	0xa9, 0xca, 0xb3, 0x27, 0x17, 0xb2, 0x02, 0x1a, 0xe7, 0x59, 0xc6, 0x8f, 0xbf, 0x6b, 0x85, 0x46,
	0x02, 0x6f, 0xe5, 0xe6, 0x2e, 0xdb, 0xe0, 0xca, 0x72, 0xbb, 0x83, 0x2f, 0x76, 0x57, 0x3a, 0x44,
	0xac, 0x2b, 0xd3, 0xee, 0xe5, 0x2f, 0x95, 0x22, 0x0f, 0xba, 0x83, 0xbe, 0x25, 0x4b, 0xaa, 0xfa,
	0xec, 0xc9, 0x05, 0x52, 0x40, 0xa3, 0x9c, 0xc7, 0xd7, 0xad, 0x5b, 0x3f, 0x49, 0xf0, 0xdd, 0xac,
	0x3b, 0x38, 0x20, 0x31, 0x8f, 0x12, 0x74, 0x0d, 0xe0, 0x3b, 0x0e, 0xfe, 0x61, 0x89, 0x63, 0x9e,
	0x99, 0xb2, 0x37, 0x4f, 0x86, 0xfa, 0xf1, 0x6b, 0xbc, 0xdf, 0x57, 0xaa, 0xdd, 0xbb, 0xfe, 0xeb,
	0xdf, 0x5f, 0xa5, 0x3b, 0x35, 0x25, 0xb7, 0x4f, 0xb7, 0x5b, 0x90, 0x30, 0xfa, 0x09, 0x68, 0x20,
	0x0a, 0x4f, 0xd3, 0x14, 0x8c, 0xde, 0xae, 0xae, 0x5a, 0x39, 0x98, 0x00, 0x6b, 0xbd, 0xa6, 0xb6,
	0xfd, 0x5a, 0xaf, 0xeb, 0xd7, 0xb6, 0xff, 0xfe, 0x47, 0x2b, 0x3c, 0xbe, 0xd1, 0xc0, 0x7f, 0x37,
	0x5a, 0xe1, 0xff, 0x1b, 0x0d, 0xfc, 0xf2, 0x54, 0x2b, 0xfc, 0xf1, 0x54, 0x03, 0xdf, 0x19, 0x01,
	0xd3, 0xf9, 0x14, 0x73, 0xb1, 0x57, 0x75, 0x8a, 0xf9, 0x43, 0x16, 0xcd, 0x8c, 0xfc, 0x4e, 0x5b,
	0x5d, 0x1a, 0xe1, 0x2c, 0x30, 0x38, 0xa7, 0xe1, 0x68, 0x54, 0x14, 0xfd, 0x2f, 0x5f, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe2, 0xbe, 0x2c, 0xc1, 0x41, 0x06, 0x00, 0x00,
}

func (x ContactType) String() string {
	s, ok := ContactType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ContactMethod) String() string {
	s, ok := ContactMethod_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ContactInfoRegistryClient is the client API for ContactInfoRegistry service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ContactInfoRegistryClient interface {
	// Request validation for the non-validated contact info for the given entity.
	RequestValidation(ctx context.Context, in *EntityIdentifiers, opts ...grpc.CallOption) (*ContactInfoValidation, error)
	// Validate confirms a contact info validation.
	Validate(ctx context.Context, in *ContactInfoValidation, opts ...grpc.CallOption) (*types.Empty, error)
}

type contactInfoRegistryClient struct {
	cc *grpc.ClientConn
}

func NewContactInfoRegistryClient(cc *grpc.ClientConn) ContactInfoRegistryClient {
	return &contactInfoRegistryClient{cc}
}

func (c *contactInfoRegistryClient) RequestValidation(ctx context.Context, in *EntityIdentifiers, opts ...grpc.CallOption) (*ContactInfoValidation, error) {
	out := new(ContactInfoValidation)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ContactInfoRegistry/RequestValidation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *contactInfoRegistryClient) Validate(ctx context.Context, in *ContactInfoValidation, opts ...grpc.CallOption) (*types.Empty, error) {
	out := new(types.Empty)
	err := c.cc.Invoke(ctx, "/ttn.lorawan.v3.ContactInfoRegistry/Validate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContactInfoRegistryServer is the server API for ContactInfoRegistry service.
type ContactInfoRegistryServer interface {
	// Request validation for the non-validated contact info for the given entity.
	RequestValidation(context.Context, *EntityIdentifiers) (*ContactInfoValidation, error)
	// Validate confirms a contact info validation.
	Validate(context.Context, *ContactInfoValidation) (*types.Empty, error)
}

// UnimplementedContactInfoRegistryServer can be embedded to have forward compatible implementations.
type UnimplementedContactInfoRegistryServer struct {
}

func (*UnimplementedContactInfoRegistryServer) RequestValidation(ctx context.Context, req *EntityIdentifiers) (*ContactInfoValidation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestValidation not implemented")
}
func (*UnimplementedContactInfoRegistryServer) Validate(ctx context.Context, req *ContactInfoValidation) (*types.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Validate not implemented")
}

func RegisterContactInfoRegistryServer(s *grpc.Server, srv ContactInfoRegistryServer) {
	s.RegisterService(&_ContactInfoRegistry_serviceDesc, srv)
}

func _ContactInfoRegistry_RequestValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EntityIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactInfoRegistryServer).RequestValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ContactInfoRegistry/RequestValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactInfoRegistryServer).RequestValidation(ctx, req.(*EntityIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContactInfoRegistry_Validate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContactInfoValidation)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContactInfoRegistryServer).Validate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ttn.lorawan.v3.ContactInfoRegistry/Validate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContactInfoRegistryServer).Validate(ctx, req.(*ContactInfoValidation))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContactInfoRegistry_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.ContactInfoRegistry",
	HandlerType: (*ContactInfoRegistryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestValidation",
			Handler:    _ContactInfoRegistry_RequestValidation_Handler,
		},
		{
			MethodName: "Validate",
			Handler:    _ContactInfoRegistry_Validate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "lorawan-stack/api/contact_info.proto",
}

func (this *ContactInfo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ContactInfo{`,
		`ContactType:` + fmt.Sprintf("%v", this.ContactType) + `,`,
		`ContactMethod:` + fmt.Sprintf("%v", this.ContactMethod) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`Public:` + fmt.Sprintf("%v", this.Public) + `,`,
		`ValidatedAt:` + strings.Replace(fmt.Sprintf("%v", this.ValidatedAt), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ContactInfoValidation) String() string {
	if this == nil {
		return "nil"
	}
	repeatedStringForContactInfo := "[]*ContactInfo{"
	for _, f := range this.ContactInfo {
		repeatedStringForContactInfo += strings.Replace(f.String(), "ContactInfo", "ContactInfo", 1) + ","
	}
	repeatedStringForContactInfo += "}"
	s := strings.Join([]string{`&ContactInfoValidation{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Token:` + fmt.Sprintf("%v", this.Token) + `,`,
		`Entity:` + strings.Replace(fmt.Sprintf("%v", this.Entity), "EntityIdentifiers", "EntityIdentifiers", 1) + `,`,
		`ContactInfo:` + repeatedStringForContactInfo + `,`,
		`CreatedAt:` + strings.Replace(fmt.Sprintf("%v", this.CreatedAt), "Timestamp", "types.Timestamp", 1) + `,`,
		`ExpiresAt:` + strings.Replace(fmt.Sprintf("%v", this.ExpiresAt), "Timestamp", "types.Timestamp", 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringContactInfo(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
