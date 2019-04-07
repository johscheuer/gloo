// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	core "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

//
//@solo-kit:resource.short_name=sec
//@solo-kit:resource.plural_name=secrets
//
//Certain plugins such as the AWS Lambda Plugin require the use of secrets for authentication, configuration of SSL Certificates, and other data that should not be stored in plaintext configuration.
//
//Gloo runs an independent (goroutine) controller to monitor secrets. Secrets are stored in their own secret storage layer. Gloo can monitor secrets stored in the following secret storage services:
//
// Kubernetes Secrets
// Hashicorp Vault
// Plaintext files (recommended only for testing)
// Secrets must adhere to a structure, specified by the plugin that requires them.
//
//Gloo's secret backend can be configured in Gloo's bootstrap options
type Secret struct {
	// Types that are valid to be assigned to Kind:
	//	*Secret_Aws
	//	*Secret_Azure
	//	*Secret_Tls
	//	*Secret_Extension
	Kind isSecret_Kind `protobuf_oneof:"kind"`
	// Metadata contains the object metadata for this resource
	Metadata             core.Metadata `protobuf:"bytes,7,opt,name=metadata,proto3" json:"metadata"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{0}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Secret.Unmarshal(m, b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return xxx_messageInfo_Secret.Size(m)
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

type isSecret_Kind interface {
	isSecret_Kind()
	Equal(interface{}) bool
}

type Secret_Aws struct {
	Aws *AwsSecret `protobuf:"bytes,1,opt,name=aws,proto3,oneof"`
}
type Secret_Azure struct {
	Azure *AzureSecret `protobuf:"bytes,2,opt,name=azure,proto3,oneof"`
}
type Secret_Tls struct {
	Tls *TlsSecret `protobuf:"bytes,3,opt,name=tls,proto3,oneof"`
}
type Secret_Extension struct {
	Extension *Extension `protobuf:"bytes,4,opt,name=extension,proto3,oneof"`
}

func (*Secret_Aws) isSecret_Kind()       {}
func (*Secret_Azure) isSecret_Kind()     {}
func (*Secret_Tls) isSecret_Kind()       {}
func (*Secret_Extension) isSecret_Kind() {}

func (m *Secret) GetKind() isSecret_Kind {
	if m != nil {
		return m.Kind
	}
	return nil
}

func (m *Secret) GetAws() *AwsSecret {
	if x, ok := m.GetKind().(*Secret_Aws); ok {
		return x.Aws
	}
	return nil
}

func (m *Secret) GetAzure() *AzureSecret {
	if x, ok := m.GetKind().(*Secret_Azure); ok {
		return x.Azure
	}
	return nil
}

func (m *Secret) GetTls() *TlsSecret {
	if x, ok := m.GetKind().(*Secret_Tls); ok {
		return x.Tls
	}
	return nil
}

func (m *Secret) GetExtension() *Extension {
	if x, ok := m.GetKind().(*Secret_Extension); ok {
		return x.Extension
	}
	return nil
}

func (m *Secret) GetMetadata() core.Metadata {
	if m != nil {
		return m.Metadata
	}
	return core.Metadata{}
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Secret) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Secret_OneofMarshaler, _Secret_OneofUnmarshaler, _Secret_OneofSizer, []interface{}{
		(*Secret_Aws)(nil),
		(*Secret_Azure)(nil),
		(*Secret_Tls)(nil),
		(*Secret_Extension)(nil),
	}
}

func _Secret_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Aws); err != nil {
			return err
		}
	case *Secret_Azure:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Azure); err != nil {
			return err
		}
	case *Secret_Tls:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Tls); err != nil {
			return err
		}
	case *Secret_Extension:
		_ = b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Extension); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Secret.Kind has unexpected type %T", x)
	}
	return nil
}

func _Secret_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Secret)
	switch tag {
	case 1: // kind.aws
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AwsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Aws{msg}
		return true, err
	case 2: // kind.azure
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(AzureSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Azure{msg}
		return true, err
	case 3: // kind.tls
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(TlsSecret)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Tls{msg}
		return true, err
	case 4: // kind.extension
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Extension)
		err := b.DecodeMessage(msg)
		m.Kind = &Secret_Extension{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Secret_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Secret)
	// kind
	switch x := m.Kind.(type) {
	case *Secret_Aws:
		s := proto.Size(x.Aws)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Azure:
		s := proto.Size(x.Azure)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Tls:
		s := proto.Size(x.Tls)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Secret_Extension:
		s := proto.Size(x.Extension)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type AwsSecret struct {
	AccessKey            string   `protobuf:"bytes,1,opt,name=access_key,json=accessKey,proto3" json:"access_key,omitempty"`
	SecretKey            string   `protobuf:"bytes,2,opt,name=secret_key,json=secretKey,proto3" json:"secret_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AwsSecret) Reset()         { *m = AwsSecret{} }
func (m *AwsSecret) String() string { return proto.CompactTextString(m) }
func (*AwsSecret) ProtoMessage()    {}
func (*AwsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{1}
}
func (m *AwsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AwsSecret.Unmarshal(m, b)
}
func (m *AwsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AwsSecret.Marshal(b, m, deterministic)
}
func (m *AwsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwsSecret.Merge(m, src)
}
func (m *AwsSecret) XXX_Size() int {
	return xxx_messageInfo_AwsSecret.Size(m)
}
func (m *AwsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AwsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AwsSecret proto.InternalMessageInfo

func (m *AwsSecret) GetAccessKey() string {
	if m != nil {
		return m.AccessKey
	}
	return ""
}

func (m *AwsSecret) GetSecretKey() string {
	if m != nil {
		return m.SecretKey
	}
	return ""
}

type AzureSecret struct {
	ApiKeys              map[string]string `protobuf:"bytes,1,rep,name=api_keys,json=apiKeys,proto3" json:"api_keys,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *AzureSecret) Reset()         { *m = AzureSecret{} }
func (m *AzureSecret) String() string { return proto.CompactTextString(m) }
func (*AzureSecret) ProtoMessage()    {}
func (*AzureSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{2}
}
func (m *AzureSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AzureSecret.Unmarshal(m, b)
}
func (m *AzureSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AzureSecret.Marshal(b, m, deterministic)
}
func (m *AzureSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AzureSecret.Merge(m, src)
}
func (m *AzureSecret) XXX_Size() int {
	return xxx_messageInfo_AzureSecret.Size(m)
}
func (m *AzureSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_AzureSecret.DiscardUnknown(m)
}

var xxx_messageInfo_AzureSecret proto.InternalMessageInfo

func (m *AzureSecret) GetApiKeys() map[string]string {
	if m != nil {
		return m.ApiKeys
	}
	return nil
}

type TlsSecret struct {
	CertChain            string   `protobuf:"bytes,1,opt,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	PrivateKey           string   `protobuf:"bytes,2,opt,name=private_key,json=privateKey,proto3" json:"private_key,omitempty"`
	RootCa               string   `protobuf:"bytes,3,opt,name=root_ca,json=rootCa,proto3" json:"root_ca,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TlsSecret) Reset()         { *m = TlsSecret{} }
func (m *TlsSecret) String() string { return proto.CompactTextString(m) }
func (*TlsSecret) ProtoMessage()    {}
func (*TlsSecret) Descriptor() ([]byte, []int) {
	return fileDescriptor_c2f79c35f1213791, []int{3}
}
func (m *TlsSecret) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TlsSecret.Unmarshal(m, b)
}
func (m *TlsSecret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TlsSecret.Marshal(b, m, deterministic)
}
func (m *TlsSecret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TlsSecret.Merge(m, src)
}
func (m *TlsSecret) XXX_Size() int {
	return xxx_messageInfo_TlsSecret.Size(m)
}
func (m *TlsSecret) XXX_DiscardUnknown() {
	xxx_messageInfo_TlsSecret.DiscardUnknown(m)
}

var xxx_messageInfo_TlsSecret proto.InternalMessageInfo

func (m *TlsSecret) GetCertChain() string {
	if m != nil {
		return m.CertChain
	}
	return ""
}

func (m *TlsSecret) GetPrivateKey() string {
	if m != nil {
		return m.PrivateKey
	}
	return ""
}

func (m *TlsSecret) GetRootCa() string {
	if m != nil {
		return m.RootCa
	}
	return ""
}

func init() {
	proto.RegisterType((*Secret)(nil), "gloo.solo.io.Secret")
	proto.RegisterType((*AwsSecret)(nil), "gloo.solo.io.AwsSecret")
	proto.RegisterType((*AzureSecret)(nil), "gloo.solo.io.AzureSecret")
	proto.RegisterMapType((map[string]string)(nil), "gloo.solo.io.AzureSecret.ApiKeysEntry")
	proto.RegisterType((*TlsSecret)(nil), "gloo.solo.io.TlsSecret")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo/projects/gloo/api/v1/secret.proto", fileDescriptor_c2f79c35f1213791)
}

var fileDescriptor_c2f79c35f1213791 = []byte{
	// 454 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x8d, 0x93, 0x34, 0xa9, 0x27, 0x3d, 0xa0, 0x55, 0x45, 0x4d, 0x24, 0x28, 0xca, 0x01, 0x55,
	0x42, 0xac, 0x49, 0x91, 0xa0, 0x54, 0xe2, 0x90, 0x54, 0x95, 0x8a, 0x22, 0x2e, 0x86, 0x13, 0x97,
	0x68, 0xbb, 0x59, 0xb9, 0x4b, 0x5c, 0x8f, 0xb5, 0xbb, 0x49, 0x31, 0xdf, 0xc0, 0x81, 0xcf, 0xe0,
	0x53, 0xf8, 0x0a, 0x0e, 0x7c, 0x09, 0xda, 0x5d, 0xc7, 0xb1, 0x50, 0x90, 0xe0, 0x64, 0xcf, 0xbc,
	0xf7, 0x46, 0xf3, 0xde, 0x0e, 0xbc, 0x4e, 0xa5, 0xb9, 0x59, 0x5d, 0x53, 0x8e, 0xb7, 0xb1, 0xc6,
	0x0c, 0x9f, 0x49, 0x8c, 0xd3, 0x0c, 0x31, 0x2e, 0x14, 0x7e, 0x12, 0xdc, 0x68, 0x5f, 0xb1, 0x42,
	0xc6, 0xeb, 0x71, 0xac, 0x05, 0x57, 0xc2, 0xd0, 0x42, 0xa1, 0x41, 0x72, 0x60, 0x11, 0x6a, 0x45,
	0x54, 0xe2, 0xf0, 0x30, 0xc5, 0x14, 0x1d, 0x10, 0xdb, 0x3f, 0xcf, 0x19, 0xbe, 0xf9, 0xaf, 0xf1,
	0xe2, 0xb3, 0x11, 0xb9, 0x96, 0x98, 0xeb, 0x4a, 0x3e, 0xde, 0x21, 0x77, 0xdf, 0xa5, 0x34, 0x1b,
	0xd1, 0xad, 0x30, 0x6c, 0xc1, 0x0c, 0xf3, 0x92, 0xd1, 0xb7, 0x36, 0xf4, 0xde, 0xbb, 0x35, 0xc9,
	0x53, 0xe8, 0xb0, 0x3b, 0x1d, 0x05, 0x8f, 0x83, 0x93, 0xc1, 0xe9, 0x11, 0x6d, 0xae, 0x4b, 0x27,
	0x77, 0xda, 0xb3, 0xae, 0x5a, 0x89, 0x65, 0x91, 0x31, 0xec, 0xb1, 0x2f, 0x2b, 0x25, 0xa2, 0xb6,
	0xa3, 0x3f, 0xf8, 0x83, 0x6e, 0xa1, 0x5a, 0xe0, 0x99, 0x76, 0xbe, 0xc9, 0x74, 0xd4, 0xd9, 0x35,
	0xff, 0x43, 0xd6, 0x98, 0x6f, 0x32, 0x4d, 0x5e, 0x41, 0x58, 0xdb, 0x8b, 0xba, 0xbb, 0x24, 0x97,
	0x1b, 0xf8, 0xaa, 0x95, 0x6c, 0xb9, 0xe4, 0x0c, 0xf6, 0x37, 0x16, 0xa3, 0xbe, 0xd3, 0xdd, 0xa7,
	0x1c, 0x95, 0xa8, 0x75, 0xef, 0x2a, 0x74, 0xda, 0xfd, 0xf1, 0xf3, 0xb8, 0x95, 0xd4, 0xec, 0x69,
	0x0f, 0xba, 0x4b, 0x99, 0x2f, 0x46, 0x6f, 0x21, 0xac, 0xed, 0x92, 0x87, 0x00, 0x8c, 0x73, 0xa1,
	0xf5, 0x7c, 0x29, 0x4a, 0x97, 0x4d, 0x98, 0x84, 0xbe, 0x33, 0x13, 0xa5, 0x85, 0xfd, 0x23, 0x3b,
	0xb8, 0xed, 0x61, 0xdf, 0x99, 0x89, 0x72, 0xf4, 0x35, 0x80, 0x41, 0x23, 0x0b, 0x32, 0x81, 0x7d,
	0x56, 0x48, 0xcb, 0xb5, 0x39, 0x77, 0x4e, 0x06, 0xa7, 0x4f, 0xfe, 0x1a, 0x1c, 0x9d, 0x14, 0x72,
	0x26, 0x4a, 0x7d, 0x99, 0x1b, 0x55, 0x26, 0x7d, 0xe6, 0xab, 0xe1, 0x39, 0x1c, 0x34, 0x01, 0x72,
	0x0f, 0x3a, 0xdb, 0xcd, 0xec, 0x2f, 0x39, 0x84, 0xbd, 0x35, 0xcb, 0x56, 0xa2, 0x5a, 0xc7, 0x17,
	0xe7, 0xed, 0xb3, 0x60, 0xb4, 0x80, 0xb0, 0x0e, 0xda, 0xae, 0xce, 0x85, 0x32, 0x73, 0x7e, 0xc3,
	0x64, 0xbe, 0x71, 0x66, 0x3b, 0x17, 0xb6, 0x41, 0x8e, 0x61, 0x50, 0x28, 0xb9, 0x66, 0x46, 0x34,
	0xac, 0x41, 0xd5, 0xb2, 0xd6, 0x8f, 0xa0, 0xaf, 0x10, 0xcd, 0x9c, 0x33, 0xf7, 0xa4, 0x61, 0xd2,
	0xb3, 0xe5, 0x05, 0x9b, 0xbe, 0xfc, 0xfe, 0xeb, 0x51, 0xf0, 0xf1, 0xf9, 0xbf, 0x9d, 0x72, 0xb1,
	0x4c, 0xab, 0xcb, 0xbc, 0xee, 0xb9, 0x8b, 0x7c, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfd, 0x1c,
	0x02, 0x9e, 0x64, 0x03, 0x00, 0x00,
}

func (this *Secret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret)
	if !ok {
		that2, ok := that.(Secret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Kind == nil {
		if this.Kind != nil {
			return false
		}
	} else if this.Kind == nil {
		return false
	} else if !this.Kind.Equal(that1.Kind) {
		return false
	}
	if !this.Metadata.Equal(&that1.Metadata) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Secret_Aws) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Aws)
	if !ok {
		that2, ok := that.(Secret_Aws)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Aws.Equal(that1.Aws) {
		return false
	}
	return true
}
func (this *Secret_Azure) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Azure)
	if !ok {
		that2, ok := that.(Secret_Azure)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Azure.Equal(that1.Azure) {
		return false
	}
	return true
}
func (this *Secret_Tls) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Tls)
	if !ok {
		that2, ok := that.(Secret_Tls)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Tls.Equal(that1.Tls) {
		return false
	}
	return true
}
func (this *Secret_Extension) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Secret_Extension)
	if !ok {
		that2, ok := that.(Secret_Extension)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Extension.Equal(that1.Extension) {
		return false
	}
	return true
}
func (this *AwsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AwsSecret)
	if !ok {
		that2, ok := that.(AwsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.AccessKey != that1.AccessKey {
		return false
	}
	if this.SecretKey != that1.SecretKey {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *AzureSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*AzureSecret)
	if !ok {
		that2, ok := that.(AzureSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.ApiKeys) != len(that1.ApiKeys) {
		return false
	}
	for i := range this.ApiKeys {
		if this.ApiKeys[i] != that1.ApiKeys[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *TlsSecret) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TlsSecret)
	if !ok {
		that2, ok := that.(TlsSecret)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.CertChain != that1.CertChain {
		return false
	}
	if this.PrivateKey != that1.PrivateKey {
		return false
	}
	if this.RootCa != that1.RootCa {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
