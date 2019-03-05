// Code generated by protoc-gen-go. DO NOT EDIT.
// source: rdp_msg.proto

package rdp_msg

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

type Protocol int32

const (
	Protocol_PROTOCOL_UNKOWN Protocol = 0
	Protocol_PROTOCOL_TCP    Protocol = 1
	Protocol_PROTOCOL_UDP    Protocol = 2
	Protocol_PROTOCOL_ICMP   Protocol = 3
	Protocol_PROTOCOL_PEER   Protocol = 4
)

var Protocol_name = map[int32]string{
	0: "PROTOCOL_UNKOWN",
	1: "PROTOCOL_TCP",
	2: "PROTOCOL_UDP",
	3: "PROTOCOL_ICMP",
	4: "PROTOCOL_PEER",
}

var Protocol_value = map[string]int32{
	"PROTOCOL_UNKOWN": 0,
	"PROTOCOL_TCP":    1,
	"PROTOCOL_UDP":    2,
	"PROTOCOL_ICMP":   3,
	"PROTOCOL_PEER":   4,
}

func (x Protocol) Enum() *Protocol {
	p := new(Protocol)
	*p = x
	return p
}

func (x Protocol) String() string {
	return proto.EnumName(Protocol_name, int32(x))
}

func (x *Protocol) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Protocol_value, data, "Protocol")
	if err != nil {
		return err
	}
	*x = Protocol(value)
	return nil
}

func (Protocol) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_33d485ccf4e36eac, []int{0}
}

type Rdp_Msg struct {
	// connect msg
	Version  *uint32   `protobuf:"varint,11,opt,name=version" json:"version,omitempty"`
	Protocol *Protocol `protobuf:"varint,12,opt,name=protocol,enum=Protocol" json:"protocol,omitempty"`
	DstIp    *uint32   `protobuf:"varint,13,opt,name=dst_ip,json=dstIp" json:"dst_ip,omitempty"`
	DstPort  *uint32   `protobuf:"varint,14,opt,name=dst_port,json=dstPort" json:"dst_port,omitempty"`
	SrcIp    *uint32   `protobuf:"varint,15,opt,name=src_ip,json=srcIp" json:"src_ip,omitempty"`
	SrcPort  *uint32   `protobuf:"varint,16,opt,name=src_port,json=srcPort" json:"src_port,omitempty"`
	UseRdp   *bool     `protobuf:"varint,21,opt,name=use_rdp,json=useRdp" json:"use_rdp,omitempty"`
	Encrypt  *bool     `protobuf:"varint,22,opt,name=encrypt" json:"encrypt,omitempty"`
	Uuid     *string   `protobuf:"bytes,23,opt,name=uuid" json:"uuid,omitempty"`
	// optional string mac = 24;
	ClientIp             *uint32  `protobuf:"varint,25,opt,name=client_ip,json=clientIp" json:"client_ip,omitempty"`
	Business             *string  `protobuf:"bytes,26,opt,name=business" json:"business,omitempty"`
	ClientFlag           *string  `protobuf:"bytes,27,opt,name=client_flag,json=clientFlag" json:"client_flag,omitempty"`
	UserId               *string  `protobuf:"bytes,28,opt,name=user_id,json=userId" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Rdp_Msg) Reset()         { *m = Rdp_Msg{} }
func (m *Rdp_Msg) String() string { return proto.CompactTextString(m) }
func (*Rdp_Msg) ProtoMessage()    {}
func (*Rdp_Msg) Descriptor() ([]byte, []int) {
	return fileDescriptor_33d485ccf4e36eac, []int{0}
}

func (m *Rdp_Msg) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Rdp_Msg.Unmarshal(m, b)
}
func (m *Rdp_Msg) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Rdp_Msg.Marshal(b, m, deterministic)
}
func (m *Rdp_Msg) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Rdp_Msg.Merge(m, src)
}
func (m *Rdp_Msg) XXX_Size() int {
	return xxx_messageInfo_Rdp_Msg.Size(m)
}
func (m *Rdp_Msg) XXX_DiscardUnknown() {
	xxx_messageInfo_Rdp_Msg.DiscardUnknown(m)
}

var xxx_messageInfo_Rdp_Msg proto.InternalMessageInfo

func (m *Rdp_Msg) GetVersion() uint32 {
	if m != nil && m.Version != nil {
		return *m.Version
	}
	return 0
}

func (m *Rdp_Msg) GetProtocol() Protocol {
	if m != nil && m.Protocol != nil {
		return *m.Protocol
	}
	return Protocol_PROTOCOL_UNKOWN
}

func (m *Rdp_Msg) GetDstIp() uint32 {
	if m != nil && m.DstIp != nil {
		return *m.DstIp
	}
	return 0
}

func (m *Rdp_Msg) GetDstPort() uint32 {
	if m != nil && m.DstPort != nil {
		return *m.DstPort
	}
	return 0
}

func (m *Rdp_Msg) GetSrcIp() uint32 {
	if m != nil && m.SrcIp != nil {
		return *m.SrcIp
	}
	return 0
}

func (m *Rdp_Msg) GetSrcPort() uint32 {
	if m != nil && m.SrcPort != nil {
		return *m.SrcPort
	}
	return 0
}

func (m *Rdp_Msg) GetUseRdp() bool {
	if m != nil && m.UseRdp != nil {
		return *m.UseRdp
	}
	return false
}

func (m *Rdp_Msg) GetEncrypt() bool {
	if m != nil && m.Encrypt != nil {
		return *m.Encrypt
	}
	return false
}

func (m *Rdp_Msg) GetUuid() string {
	if m != nil && m.Uuid != nil {
		return *m.Uuid
	}
	return ""
}

func (m *Rdp_Msg) GetClientIp() uint32 {
	if m != nil && m.ClientIp != nil {
		return *m.ClientIp
	}
	return 0
}

func (m *Rdp_Msg) GetBusiness() string {
	if m != nil && m.Business != nil {
		return *m.Business
	}
	return ""
}

func (m *Rdp_Msg) GetClientFlag() string {
	if m != nil && m.ClientFlag != nil {
		return *m.ClientFlag
	}
	return ""
}

func (m *Rdp_Msg) GetUserId() string {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return ""
}

func init() {
	proto.RegisterEnum("Protocol", Protocol_name, Protocol_value)
	proto.RegisterType((*Rdp_Msg)(nil), "Rdp_Msg")
}

func init() { proto.RegisterFile("rdp_msg.proto", fileDescriptor_33d485ccf4e36eac) }

var fileDescriptor_33d485ccf4e36eac = []byte{
	// 329 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x51, 0x6b, 0xea, 0x30,
	0x18, 0x86, 0x4f, 0x3d, 0x1e, 0xad, 0x9f, 0x56, 0x7b, 0x32, 0x9c, 0x51, 0x07, 0x2b, 0x83, 0x41,
	0xd9, 0x85, 0x17, 0xfb, 0x0b, 0xce, 0x41, 0xd9, 0xb4, 0x25, 0x28, 0xbb, 0x2c, 0xae, 0xe9, 0x4a,
	0xa0, 0x6b, 0x43, 0xbe, 0x74, 0xb0, 0x7f, 0xbc, 0x9f, 0x31, 0x9a, 0x6a, 0xc1, 0xbb, 0xbe, 0xcf,
	0xdb, 0x27, 0x79, 0x09, 0x38, 0x8a, 0xcb, 0xf8, 0x13, 0xb3, 0x95, 0x54, 0xa5, 0x2e, 0xef, 0x7e,
	0x3a, 0xd0, 0x67, 0x5c, 0xc6, 0x5b, 0xcc, 0x08, 0x85, 0xfe, 0x57, 0xaa, 0x50, 0x94, 0x05, 0x1d,
	0x7a, 0x96, 0xef, 0xb0, 0x73, 0x24, 0xf7, 0x60, 0x9b, 0xdf, 0x93, 0x32, 0xa7, 0x23, 0xcf, 0xf2,
	0xc7, 0x8f, 0x83, 0x55, 0x74, 0x02, 0xac, 0xad, 0xc8, 0x14, 0x7a, 0x1c, 0x75, 0x2c, 0x24, 0x75,
	0x8c, 0xff, 0x8f, 0xa3, 0x0e, 0x24, 0x99, 0x83, 0x5d, 0x63, 0x59, 0x2a, 0x4d, 0xc7, 0xcd, 0xc1,
	0x1c, 0x75, 0x54, 0x2a, 0x5d, 0x1b, 0xa8, 0x92, 0xda, 0x98, 0x34, 0x06, 0xaa, 0xa4, 0x31, 0x6a,
	0x6c, 0x0c, 0xb7, 0x31, 0x50, 0x25, 0xc6, 0x98, 0x41, 0xbf, 0xc2, 0x34, 0x56, 0x5c, 0xd2, 0xa9,
	0x67, 0xf9, 0x36, 0xeb, 0x55, 0x98, 0x32, 0x2e, 0xeb, 0xf5, 0x69, 0x91, 0xa8, 0x6f, 0xa9, 0xe9,
	0xb5, 0x29, 0xce, 0x91, 0x10, 0xe8, 0x56, 0x95, 0xe0, 0x74, 0xe6, 0x59, 0xfe, 0x80, 0x99, 0x6f,
	0xb2, 0x84, 0x41, 0x92, 0x8b, 0xb4, 0x30, 0x6b, 0xe7, 0xe6, 0x0a, 0xbb, 0x01, 0x81, 0x24, 0x0b,
	0xb0, 0xdf, 0x2b, 0x14, 0x45, 0x8a, 0x48, 0x17, 0x46, 0x6a, 0x33, 0xb9, 0x85, 0xe1, 0x49, 0xfc,
	0xc8, 0x8f, 0x19, 0x5d, 0x9a, 0x1a, 0x1a, 0xf4, 0x9c, 0x1f, 0xb3, 0xd3, 0x40, 0x15, 0x0b, 0x4e,
	0x6f, 0x4c, 0x59, 0x0f, 0x54, 0x01, 0x7f, 0x10, 0x60, 0x9f, 0xdf, 0x8c, 0x5c, 0xc1, 0x24, 0x62,
	0xe1, 0x3e, 0x5c, 0x87, 0xaf, 0xf1, 0x61, 0xf7, 0x12, 0xbe, 0xed, 0xdc, 0x3f, 0xc4, 0x85, 0x51,
	0x0b, 0xf7, 0xeb, 0xc8, 0xb5, 0x2e, 0xc8, 0xe1, 0x29, 0x72, 0x3b, 0xe4, 0x3f, 0x38, 0x2d, 0x09,
	0xd6, 0xdb, 0xc8, 0xfd, 0x7b, 0x81, 0xa2, 0xcd, 0x86, 0xb9, 0xdd, 0xdf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x51, 0x9d, 0x2a, 0xa4, 0xe5, 0x01, 0x00, 0x00,
}