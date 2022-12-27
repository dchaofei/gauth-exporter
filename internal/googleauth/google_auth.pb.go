// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google_auth.proto

package googleauth

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

type MigrationPayload_Algorithm int32

const (
	MigrationPayload_ALGO_INVALID MigrationPayload_Algorithm = 0
	MigrationPayload_ALGO_SHA1    MigrationPayload_Algorithm = 1
)

var MigrationPayload_Algorithm_name = map[int32]string{
	0: "ALGO_INVALID",
	1: "ALGO_SHA1",
}

var MigrationPayload_Algorithm_value = map[string]int32{
	"ALGO_INVALID": 0,
	"ALGO_SHA1":    1,
}

func (x MigrationPayload_Algorithm) String() string {
	return proto.EnumName(MigrationPayload_Algorithm_name, int32(x))
}

func (MigrationPayload_Algorithm) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4375413210822c36, []int{0, 0}
}

type MigrationPayload_OtpType int32

const (
	MigrationPayload_OTP_INVALID MigrationPayload_OtpType = 0
	MigrationPayload_OTP_HOTP    MigrationPayload_OtpType = 1
	MigrationPayload_OTP_TOTP    MigrationPayload_OtpType = 2
)

var MigrationPayload_OtpType_name = map[int32]string{
	0: "OTP_INVALID",
	1: "OTP_HOTP",
	2: "OTP_TOTP",
}

var MigrationPayload_OtpType_value = map[string]int32{
	"OTP_INVALID": 0,
	"OTP_HOTP":    1,
	"OTP_TOTP":    2,
}

func (x MigrationPayload_OtpType) String() string {
	return proto.EnumName(MigrationPayload_OtpType_name, int32(x))
}

func (MigrationPayload_OtpType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4375413210822c36, []int{0, 1}
}

type MigrationPayload struct {
	OtpParameters        []*MigrationPayload_OtpParameters `protobuf:"bytes,1,rep,name=otp_parameters,json=otpParameters,proto3" json:"otp_parameters,omitempty"`
	Version              int32                             `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	BatchSize            int32                             `protobuf:"varint,3,opt,name=batch_size,json=batchSize,proto3" json:"batch_size,omitempty"`
	BatchIndex           int32                             `protobuf:"varint,4,opt,name=batch_index,json=batchIndex,proto3" json:"batch_index,omitempty"`
	BatchId              int32                             `protobuf:"varint,5,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *MigrationPayload) Reset()         { *m = MigrationPayload{} }
func (m *MigrationPayload) String() string { return proto.CompactTextString(m) }
func (*MigrationPayload) ProtoMessage()    {}
func (*MigrationPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_4375413210822c36, []int{0}
}

func (m *MigrationPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationPayload.Unmarshal(m, b)
}
func (m *MigrationPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationPayload.Marshal(b, m, deterministic)
}
func (m *MigrationPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationPayload.Merge(m, src)
}
func (m *MigrationPayload) XXX_Size() int {
	return xxx_messageInfo_MigrationPayload.Size(m)
}
func (m *MigrationPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationPayload.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationPayload proto.InternalMessageInfo

func (m *MigrationPayload) GetOtpParameters() []*MigrationPayload_OtpParameters {
	if m != nil {
		return m.OtpParameters
	}
	return nil
}

func (m *MigrationPayload) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *MigrationPayload) GetBatchSize() int32 {
	if m != nil {
		return m.BatchSize
	}
	return 0
}

func (m *MigrationPayload) GetBatchIndex() int32 {
	if m != nil {
		return m.BatchIndex
	}
	return 0
}

func (m *MigrationPayload) GetBatchId() int32 {
	if m != nil {
		return m.BatchId
	}
	return 0
}

type MigrationPayload_OtpParameters struct {
	Secret               []byte                     `protobuf:"bytes,1,opt,name=secret,proto3" json:"secret,omitempty"`
	Name                 string                     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Issuer    string                     `protobuf:"bytes,3,opt,name=issuer,proto3" json:"issuer,omitempty"`
	Algorithm MigrationPayload_Algorithm `protobuf:"varint,4,opt,name=algorithm,proto3,enum=googleauth.MigrationPayload_Algorithm" json:"algorithm,omitempty"`
	Digits    int32                      `protobuf:"varint,5,opt,name=digits,proto3" json:"digits,omitempty"`
	Type      MigrationPayload_OtpType   `protobuf:"varint,6,opt,name=type,proto3,enum=googleauth.MigrationPayload_OtpType" json:"type,omitempty"`
	Counter   int64                      `protobuf:"varint,7,opt,name=counter,proto3" json:"counter,omitempty"`
	PlainSecret          string                     `protobuf:"bytes,8,opt,name=plainSecret,proto3" json:"plainSecret,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *MigrationPayload_OtpParameters) Reset()         { *m = MigrationPayload_OtpParameters{} }
func (m *MigrationPayload_OtpParameters) String() string { return proto.CompactTextString(m) }
func (*MigrationPayload_OtpParameters) ProtoMessage()    {}
func (*MigrationPayload_OtpParameters) Descriptor() ([]byte, []int) {
	return fileDescriptor_4375413210822c36, []int{0, 0}
}

func (m *MigrationPayload_OtpParameters) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MigrationPayload_OtpParameters.Unmarshal(m, b)
}
func (m *MigrationPayload_OtpParameters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MigrationPayload_OtpParameters.Marshal(b, m, deterministic)
}
func (m *MigrationPayload_OtpParameters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MigrationPayload_OtpParameters.Merge(m, src)
}
func (m *MigrationPayload_OtpParameters) XXX_Size() int {
	return xxx_messageInfo_MigrationPayload_OtpParameters.Size(m)
}
func (m *MigrationPayload_OtpParameters) XXX_DiscardUnknown() {
	xxx_messageInfo_MigrationPayload_OtpParameters.DiscardUnknown(m)
}

var xxx_messageInfo_MigrationPayload_OtpParameters proto.InternalMessageInfo

func (m *MigrationPayload_OtpParameters) GetSecret() []byte {
	if m != nil {
		return m.Secret
	}
	return nil
}

func (m *MigrationPayload_OtpParameters) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MigrationPayload_OtpParameters) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *MigrationPayload_OtpParameters) GetAlgorithm() MigrationPayload_Algorithm {
	if m != nil {
		return m.Algorithm
	}
	return MigrationPayload_ALGO_INVALID
}

func (m *MigrationPayload_OtpParameters) GetDigits() int32 {
	if m != nil {
		return m.Digits
	}
	return 0
}

func (m *MigrationPayload_OtpParameters) GetType() MigrationPayload_OtpType {
	if m != nil {
		return m.Type
	}
	return MigrationPayload_OTP_INVALID
}

func (m *MigrationPayload_OtpParameters) GetCounter() int64 {
	if m != nil {
		return m.Counter
	}
	return 0
}

func (m *MigrationPayload_OtpParameters) GetPlainSecret() string {
	if m != nil {
		return m.PlainSecret
	}
	return ""
}

func init() {
	proto.RegisterEnum("googleauth.MigrationPayload_Algorithm", MigrationPayload_Algorithm_name, MigrationPayload_Algorithm_value)
	proto.RegisterEnum("googleauth.MigrationPayload_OtpType", MigrationPayload_OtpType_name, MigrationPayload_OtpType_value)
	proto.RegisterType((*MigrationPayload)(nil), "googleauth.MigrationPayload")
	proto.RegisterType((*MigrationPayload_OtpParameters)(nil), "googleauth.MigrationPayload.OtpParameters")
}

func init() { proto.RegisterFile("google_auth.proto", fileDescriptor_4375413210822c36) }

var fileDescriptor_4375413210822c36 = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0xdf, 0xca, 0xd3, 0x40,
	0x10, 0xc5, 0x4d, 0x9b, 0xfe, 0xc9, 0xa4, 0xad, 0x71, 0x2f, 0x64, 0x15, 0xc4, 0x50, 0x44, 0x82,
	0x48, 0xc0, 0x0a, 0xe2, 0x6d, 0xa0, 0x60, 0x0b, 0xd5, 0xc4, 0x6d, 0xf0, 0x36, 0x6c, 0x9b, 0x25,
	0x5d, 0x48, 0xb3, 0x61, 0xb3, 0x15, 0xdb, 0x17, 0xf2, 0xdd, 0x7c, 0x0a, 0xc9, 0x26, 0x69, 0xab,
	0x17, 0xdf, 0x77, 0x97, 0xdf, 0x99, 0x39, 0x93, 0x33, 0xc3, 0xc2, 0xb3, 0x4c, 0x88, 0x2c, 0x67,
	0x09, 0x3d, 0xa9, 0x83, 0x5f, 0x4a, 0xa1, 0x04, 0x82, 0x46, 0xaa, 0x95, 0xf9, 0x1f, 0x13, 0x9c,
	0xaf, 0x3c, 0x93, 0x54, 0x71, 0x51, 0x44, 0xf4, 0x9c, 0x0b, 0x9a, 0xa2, 0xef, 0x30, 0x13, 0xaa,
	0x4c, 0x4a, 0x2a, 0xe9, 0x91, 0x29, 0x26, 0x2b, 0x6c, 0xb8, 0x7d, 0xcf, 0x5e, 0xbc, 0xf3, 0x6f,
	0x4e, 0xff, 0x7f, 0x97, 0x1f, 0xaa, 0x32, 0xba, 0x3a, 0xc8, 0x54, 0xdc, 0x23, 0xc2, 0x30, 0xfa,
	0xc9, 0x64, 0xc5, 0x45, 0x81, 0x7b, 0xae, 0xe1, 0x0d, 0x48, 0x87, 0xe8, 0x15, 0xc0, 0x8e, 0xaa,
	0xfd, 0x21, 0xa9, 0xf8, 0x85, 0xe1, 0xbe, 0x2e, 0x5a, 0x5a, 0xd9, 0xf2, 0x0b, 0x43, 0xaf, 0xc1,
	0x6e, 0xca, 0xbc, 0x48, 0xd9, 0x2f, 0x6c, 0xea, 0x7a, 0xe3, 0x58, 0xd7, 0x0a, 0x7a, 0x01, 0xe3,
	0xb6, 0x21, 0xc5, 0x83, 0x66, 0x74, 0x53, 0x4d, 0x5f, 0xfe, 0xee, 0xc1, 0xf4, 0x9f, 0x54, 0xe8,
	0x39, 0x0c, 0x2b, 0xb6, 0x97, 0x4c, 0x61, 0xc3, 0x35, 0xbc, 0x09, 0x69, 0x09, 0x21, 0x30, 0x0b,
	0x7a, 0x64, 0x3a, 0x9b, 0x45, 0xf4, 0x77, 0xdd, 0xcb, 0xab, 0xea, 0xc4, 0xa4, 0x0e, 0x65, 0x91,
	0x96, 0xd0, 0x12, 0x2c, 0x9a, 0x67, 0x42, 0x72, 0x75, 0x38, 0xea, 0x3c, 0xb3, 0xc5, 0xdb, 0x07,
	0x0f, 0x13, 0x74, 0xdd, 0xe4, 0x66, 0xac, 0xa7, 0xa7, 0x3c, 0xe3, 0xaa, 0x6a, 0x43, 0xb7, 0x84,
	0x3e, 0x83, 0xa9, 0xce, 0x25, 0xc3, 0x43, 0x3d, 0xf8, 0xcd, 0x63, 0x17, 0x8f, 0xcf, 0x25, 0x23,
	0xda, 0x51, 0x9f, 0x78, 0x2f, 0x4e, 0x85, 0x62, 0x12, 0x8f, 0x5c, 0xc3, 0xeb, 0x93, 0x0e, 0x91,
	0x0b, 0x76, 0x99, 0x53, 0x5e, 0x6c, 0x9b, 0xd5, 0xc7, 0x7a, 0x9d, 0x7b, 0x69, 0xfe, 0x1e, 0xac,
	0x6b, 0x4a, 0xe4, 0xc0, 0x24, 0xd8, 0x7c, 0x09, 0x93, 0xf5, 0xb7, 0x1f, 0xc1, 0x66, 0xbd, 0x74,
	0x9e, 0xa0, 0x29, 0x58, 0x5a, 0xd9, 0xae, 0x82, 0x0f, 0x8e, 0x31, 0xff, 0x04, 0xa3, 0xf6, 0xd7,
	0xe8, 0x29, 0xd8, 0x61, 0x1c, 0xdd, 0xb5, 0x4e, 0x60, 0x5c, 0x0b, 0xab, 0x30, 0x8e, 0x1c, 0xa3,
	0xa3, 0xb8, 0xa6, 0xde, 0x6e, 0xa8, 0xdf, 0xdf, 0xc7, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x81,
	0xbe, 0xb2, 0xcc, 0x94, 0x02, 0x00, 0x00,
}