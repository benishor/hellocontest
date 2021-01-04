// Code generated by protoc-gen-go. DO NOT EDIT.
// source: log.proto

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

type FileInfo struct {
	FormatVersion        int32    `protobuf:"varint,1,opt,name=format_version,json=formatVersion" json:"format_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}
func (*FileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{0}
}
func (m *FileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FileInfo.Unmarshal(m, b)
}
func (m *FileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FileInfo.Marshal(b, m, deterministic)
}
func (dst *FileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FileInfo.Merge(dst, src)
}
func (m *FileInfo) XXX_Size() int {
	return xxx_messageInfo_FileInfo.Size(m)
}
func (m *FileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_FileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_FileInfo proto.InternalMessageInfo

func (m *FileInfo) GetFormatVersion() int32 {
	if m != nil {
		return m.FormatVersion
	}
	return 0
}

type Entry struct {
	// Types that are valid to be assigned to Entry:
	//	*Entry_Qso
	//	*Entry_Station
	//	*Entry_Contest
	Entry                isEntry_Entry `protobuf_oneof:"entry"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{1}
}
func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (dst *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(dst, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

type isEntry_Entry interface {
	isEntry_Entry()
}

type Entry_Qso struct {
	Qso *QSO `protobuf:"bytes,1,opt,name=qso,oneof"`
}
type Entry_Station struct {
	Station *Station `protobuf:"bytes,2,opt,name=station,oneof"`
}
type Entry_Contest struct {
	Contest *Contest `protobuf:"bytes,3,opt,name=contest,oneof"`
}

func (*Entry_Qso) isEntry_Entry()     {}
func (*Entry_Station) isEntry_Entry() {}
func (*Entry_Contest) isEntry_Entry() {}

func (m *Entry) GetEntry() isEntry_Entry {
	if m != nil {
		return m.Entry
	}
	return nil
}

func (m *Entry) GetQso() *QSO {
	if x, ok := m.GetEntry().(*Entry_Qso); ok {
		return x.Qso
	}
	return nil
}

func (m *Entry) GetStation() *Station {
	if x, ok := m.GetEntry().(*Entry_Station); ok {
		return x.Station
	}
	return nil
}

func (m *Entry) GetContest() *Contest {
	if x, ok := m.GetEntry().(*Entry_Contest); ok {
		return x.Contest
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Entry) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Entry_OneofMarshaler, _Entry_OneofUnmarshaler, _Entry_OneofSizer, []interface{}{
		(*Entry_Qso)(nil),
		(*Entry_Station)(nil),
		(*Entry_Contest)(nil),
	}
}

func _Entry_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Entry)
	// entry
	switch x := m.Entry.(type) {
	case *Entry_Qso:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Qso); err != nil {
			return err
		}
	case *Entry_Station:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Station); err != nil {
			return err
		}
	case *Entry_Contest:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Contest); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Entry.Entry has unexpected type %T", x)
	}
	return nil
}

func _Entry_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Entry)
	switch tag {
	case 1: // entry.qso
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(QSO)
		err := b.DecodeMessage(msg)
		m.Entry = &Entry_Qso{msg}
		return true, err
	case 2: // entry.station
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Station)
		err := b.DecodeMessage(msg)
		m.Entry = &Entry_Station{msg}
		return true, err
	case 3: // entry.contest
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Contest)
		err := b.DecodeMessage(msg)
		m.Entry = &Entry_Contest{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Entry_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Entry)
	// entry
	switch x := m.Entry.(type) {
	case *Entry_Qso:
		s := proto.Size(x.Qso)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Station:
		s := proto.Size(x.Station)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Entry_Contest:
		s := proto.Size(x.Contest)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type QSO struct {
	Callsign             string   `protobuf:"bytes,1,opt,name=callsign" json:"callsign,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Band                 string   `protobuf:"bytes,3,opt,name=band" json:"band,omitempty"`
	Mode                 string   `protobuf:"bytes,4,opt,name=mode" json:"mode,omitempty"`
	MyReport             string   `protobuf:"bytes,5,opt,name=my_report,json=myReport" json:"my_report,omitempty"`
	MyNumber             int32    `protobuf:"varint,6,opt,name=my_number,json=myNumber" json:"my_number,omitempty"`
	TheirReport          string   `protobuf:"bytes,7,opt,name=their_report,json=theirReport" json:"their_report,omitempty"`
	TheirNumber          int32    `protobuf:"varint,8,opt,name=their_number,json=theirNumber" json:"their_number,omitempty"`
	LogTimestamp         int64    `protobuf:"varint,9,opt,name=log_timestamp,json=logTimestamp" json:"log_timestamp,omitempty"`
	MyXchange            string   `protobuf:"bytes,10,opt,name=my_xchange,json=myXchange" json:"my_xchange,omitempty"`
	TheirXchange         string   `protobuf:"bytes,11,opt,name=their_xchange,json=theirXchange" json:"their_xchange,omitempty"`
	Frequency            float64  `protobuf:"fixed64,12,opt,name=frequency" json:"frequency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QSO) Reset()         { *m = QSO{} }
func (m *QSO) String() string { return proto.CompactTextString(m) }
func (*QSO) ProtoMessage()    {}
func (*QSO) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{2}
}
func (m *QSO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QSO.Unmarshal(m, b)
}
func (m *QSO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QSO.Marshal(b, m, deterministic)
}
func (dst *QSO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QSO.Merge(dst, src)
}
func (m *QSO) XXX_Size() int {
	return xxx_messageInfo_QSO.Size(m)
}
func (m *QSO) XXX_DiscardUnknown() {
	xxx_messageInfo_QSO.DiscardUnknown(m)
}

var xxx_messageInfo_QSO proto.InternalMessageInfo

func (m *QSO) GetCallsign() string {
	if m != nil {
		return m.Callsign
	}
	return ""
}

func (m *QSO) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *QSO) GetBand() string {
	if m != nil {
		return m.Band
	}
	return ""
}

func (m *QSO) GetMode() string {
	if m != nil {
		return m.Mode
	}
	return ""
}

func (m *QSO) GetMyReport() string {
	if m != nil {
		return m.MyReport
	}
	return ""
}

func (m *QSO) GetMyNumber() int32 {
	if m != nil {
		return m.MyNumber
	}
	return 0
}

func (m *QSO) GetTheirReport() string {
	if m != nil {
		return m.TheirReport
	}
	return ""
}

func (m *QSO) GetTheirNumber() int32 {
	if m != nil {
		return m.TheirNumber
	}
	return 0
}

func (m *QSO) GetLogTimestamp() int64 {
	if m != nil {
		return m.LogTimestamp
	}
	return 0
}

func (m *QSO) GetMyXchange() string {
	if m != nil {
		return m.MyXchange
	}
	return ""
}

func (m *QSO) GetTheirXchange() string {
	if m != nil {
		return m.TheirXchange
	}
	return ""
}

func (m *QSO) GetFrequency() float64 {
	if m != nil {
		return m.Frequency
	}
	return 0
}

type Station struct {
	Callsign             string   `protobuf:"bytes,1,opt,name=callsign" json:"callsign,omitempty"`
	Operator             string   `protobuf:"bytes,2,opt,name=operator" json:"operator,omitempty"`
	Locator              string   `protobuf:"bytes,3,opt,name=locator" json:"locator,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Station) Reset()         { *m = Station{} }
func (m *Station) String() string { return proto.CompactTextString(m) }
func (*Station) ProtoMessage()    {}
func (*Station) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{3}
}
func (m *Station) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Station.Unmarshal(m, b)
}
func (m *Station) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Station.Marshal(b, m, deterministic)
}
func (dst *Station) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Station.Merge(dst, src)
}
func (m *Station) XXX_Size() int {
	return xxx_messageInfo_Station.Size(m)
}
func (m *Station) XXX_DiscardUnknown() {
	xxx_messageInfo_Station.DiscardUnknown(m)
}

var xxx_messageInfo_Station proto.InternalMessageInfo

func (m *Station) GetCallsign() string {
	if m != nil {
		return m.Callsign
	}
	return ""
}

func (m *Station) GetOperator() string {
	if m != nil {
		return m.Operator
	}
	return ""
}

func (m *Station) GetLocator() string {
	if m != nil {
		return m.Locator
	}
	return ""
}

type Contest struct {
	Name                    string   `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	EnterTheirNumber        bool     `protobuf:"varint,2,opt,name=enter_their_number,json=enterTheirNumber" json:"enter_their_number,omitempty"`
	EnterTheirXchange       bool     `protobuf:"varint,3,opt,name=enter_their_xchange,json=enterTheirXchange" json:"enter_their_xchange,omitempty"`
	RequireTheirXchange     bool     `protobuf:"varint,4,opt,name=require_their_xchange,json=requireTheirXchange" json:"require_their_xchange,omitempty"`
	AllowMultiBand          bool     `protobuf:"varint,5,opt,name=allow_multi_band,json=allowMultiBand" json:"allow_multi_band,omitempty"`
	AllowMultiMode          bool     `protobuf:"varint,6,opt,name=allow_multi_mode,json=allowMultiMode" json:"allow_multi_mode,omitempty"`
	SameCountryPoints       int32    `protobuf:"varint,7,opt,name=same_country_points,json=sameCountryPoints" json:"same_country_points,omitempty"`
	SameContinentPoints     int32    `protobuf:"varint,8,opt,name=same_continent_points,json=sameContinentPoints" json:"same_continent_points,omitempty"`
	SpecificCountryPoints   int32    `protobuf:"varint,9,opt,name=specific_country_points,json=specificCountryPoints" json:"specific_country_points,omitempty"`
	SpecificCountryPrefixes []string `protobuf:"bytes,10,rep,name=specific_country_prefixes,json=specificCountryPrefixes" json:"specific_country_prefixes,omitempty"`
	OtherPoints             int32    `protobuf:"varint,11,opt,name=other_points,json=otherPoints" json:"other_points,omitempty"`
	Multis                  *Multis  `protobuf:"bytes,12,opt,name=multis" json:"multis,omitempty"`
	XchangeMultiPattern     string   `protobuf:"bytes,13,opt,name=xchange_multi_pattern,json=xchangeMultiPattern" json:"xchange_multi_pattern,omitempty"`
	CountPerBand            bool     `protobuf:"varint,14,opt,name=count_per_band,json=countPerBand" json:"count_per_band,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *Contest) Reset()         { *m = Contest{} }
func (m *Contest) String() string { return proto.CompactTextString(m) }
func (*Contest) ProtoMessage()    {}
func (*Contest) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{4}
}
func (m *Contest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Contest.Unmarshal(m, b)
}
func (m *Contest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Contest.Marshal(b, m, deterministic)
}
func (dst *Contest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contest.Merge(dst, src)
}
func (m *Contest) XXX_Size() int {
	return xxx_messageInfo_Contest.Size(m)
}
func (m *Contest) XXX_DiscardUnknown() {
	xxx_messageInfo_Contest.DiscardUnknown(m)
}

var xxx_messageInfo_Contest proto.InternalMessageInfo

func (m *Contest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Contest) GetEnterTheirNumber() bool {
	if m != nil {
		return m.EnterTheirNumber
	}
	return false
}

func (m *Contest) GetEnterTheirXchange() bool {
	if m != nil {
		return m.EnterTheirXchange
	}
	return false
}

func (m *Contest) GetRequireTheirXchange() bool {
	if m != nil {
		return m.RequireTheirXchange
	}
	return false
}

func (m *Contest) GetAllowMultiBand() bool {
	if m != nil {
		return m.AllowMultiBand
	}
	return false
}

func (m *Contest) GetAllowMultiMode() bool {
	if m != nil {
		return m.AllowMultiMode
	}
	return false
}

func (m *Contest) GetSameCountryPoints() int32 {
	if m != nil {
		return m.SameCountryPoints
	}
	return 0
}

func (m *Contest) GetSameContinentPoints() int32 {
	if m != nil {
		return m.SameContinentPoints
	}
	return 0
}

func (m *Contest) GetSpecificCountryPoints() int32 {
	if m != nil {
		return m.SpecificCountryPoints
	}
	return 0
}

func (m *Contest) GetSpecificCountryPrefixes() []string {
	if m != nil {
		return m.SpecificCountryPrefixes
	}
	return nil
}

func (m *Contest) GetOtherPoints() int32 {
	if m != nil {
		return m.OtherPoints
	}
	return 0
}

func (m *Contest) GetMultis() *Multis {
	if m != nil {
		return m.Multis
	}
	return nil
}

func (m *Contest) GetXchangeMultiPattern() string {
	if m != nil {
		return m.XchangeMultiPattern
	}
	return ""
}

func (m *Contest) GetCountPerBand() bool {
	if m != nil {
		return m.CountPerBand
	}
	return false
}

type Multis struct {
	Dxcc                 bool     `protobuf:"varint,1,opt,name=dxcc" json:"dxcc,omitempty"`
	Wpx                  bool     `protobuf:"varint,2,opt,name=wpx" json:"wpx,omitempty"`
	Xchange              bool     `protobuf:"varint,3,opt,name=xchange" json:"xchange,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Multis) Reset()         { *m = Multis{} }
func (m *Multis) String() string { return proto.CompactTextString(m) }
func (*Multis) ProtoMessage()    {}
func (*Multis) Descriptor() ([]byte, []int) {
	return fileDescriptor_log_8b86d34e10137475, []int{5}
}
func (m *Multis) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Multis.Unmarshal(m, b)
}
func (m *Multis) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Multis.Marshal(b, m, deterministic)
}
func (dst *Multis) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Multis.Merge(dst, src)
}
func (m *Multis) XXX_Size() int {
	return xxx_messageInfo_Multis.Size(m)
}
func (m *Multis) XXX_DiscardUnknown() {
	xxx_messageInfo_Multis.DiscardUnknown(m)
}

var xxx_messageInfo_Multis proto.InternalMessageInfo

func (m *Multis) GetDxcc() bool {
	if m != nil {
		return m.Dxcc
	}
	return false
}

func (m *Multis) GetWpx() bool {
	if m != nil {
		return m.Wpx
	}
	return false
}

func (m *Multis) GetXchange() bool {
	if m != nil {
		return m.Xchange
	}
	return false
}

func init() {
	proto.RegisterType((*FileInfo)(nil), "pb.FileInfo")
	proto.RegisterType((*Entry)(nil), "pb.Entry")
	proto.RegisterType((*QSO)(nil), "pb.QSO")
	proto.RegisterType((*Station)(nil), "pb.Station")
	proto.RegisterType((*Contest)(nil), "pb.Contest")
	proto.RegisterType((*Multis)(nil), "pb.Multis")
}

func init() { proto.RegisterFile("log.proto", fileDescriptor_log_8b86d34e10137475) }

var fileDescriptor_log_8b86d34e10137475 = []byte{
	// 682 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xcf, 0x6b, 0xdb, 0x3e,
	0x14, 0x6f, 0xea, 0x24, 0xb6, 0x5f, 0x7e, 0x90, 0x2a, 0x94, 0xfa, 0xdb, 0xf6, 0x0b, 0x9d, 0xb7,
	0xb1, 0x1c, 0x46, 0x60, 0x1d, 0xec, 0xb0, 0x63, 0xcb, 0x46, 0x37, 0xe8, 0xda, 0xaa, 0x65, 0x0c,
	0x76, 0x30, 0x8e, 0xab, 0x24, 0x06, 0x5b, 0x72, 0x65, 0x65, 0x8d, 0x4f, 0xfb, 0xa7, 0x77, 0xdc,
	0x61, 0xe8, 0x49, 0x4e, 0x9a, 0x0c, 0x76, 0x93, 0x3e, 0x3f, 0xde, 0xf3, 0x7b, 0xfa, 0x60, 0xf0,
	0x33, 0x31, 0x1b, 0x17, 0x52, 0x28, 0x41, 0x76, 0x8b, 0x49, 0xf8, 0x06, 0xbc, 0x8f, 0x69, 0xc6,
	0x3e, 0xf1, 0xa9, 0x20, 0x2f, 0xa1, 0x3f, 0x15, 0x32, 0x8f, 0x55, 0xf4, 0x83, 0xc9, 0x32, 0x15,
	0x3c, 0x68, 0x9c, 0x34, 0x46, 0x2d, 0xda, 0x33, 0xe8, 0x57, 0x03, 0x86, 0x3f, 0xa1, 0xf5, 0x81,
	0x2b, 0x59, 0x91, 0x23, 0x70, 0x1e, 0x4a, 0x81, 0xa2, 0xce, 0xa9, 0x3b, 0x2e, 0x26, 0xe3, 0x9b,
	0xdb, 0xab, 0x8b, 0x1d, 0xaa, 0x51, 0xf2, 0x0a, 0xdc, 0x52, 0xc5, 0x4a, 0x57, 0xd9, 0x45, 0x41,
	0x47, 0x0b, 0x6e, 0x0d, 0x74, 0xb1, 0x43, 0x6b, 0x56, 0x0b, 0x13, 0xc1, 0x15, 0x2b, 0x55, 0xe0,
	0xac, 0x85, 0xe7, 0x06, 0xd2, 0x42, 0xcb, 0x9e, 0xb9, 0xd0, 0x62, 0xba, 0x6f, 0xf8, 0x6b, 0x17,
	0x9c, 0x9b, 0xdb, 0x2b, 0x72, 0x08, 0x5e, 0x12, 0x67, 0x59, 0x99, 0xce, 0xcc, 0x97, 0xfa, 0x74,
	0x75, 0x27, 0xc7, 0xe0, 0xab, 0x34, 0x67, 0xa5, 0x8a, 0xf3, 0x02, 0x3f, 0xc0, 0xa1, 0x6b, 0x80,
	0x10, 0x68, 0x4e, 0x62, 0x7e, 0x8f, 0x0d, 0x7d, 0x8a, 0x67, 0x8d, 0xe5, 0xe2, 0x9e, 0x05, 0x4d,
	0x83, 0xe9, 0x33, 0x39, 0x02, 0x3f, 0xaf, 0x22, 0xc9, 0x0a, 0x21, 0x55, 0xd0, 0x32, 0x2d, 0xf2,
	0x8a, 0xe2, 0xdd, 0x92, 0x7c, 0x91, 0x4f, 0x98, 0x0c, 0xda, 0xb8, 0x29, 0x2f, 0xaf, 0xbe, 0xe0,
	0x9d, 0x3c, 0x83, 0xae, 0x9a, 0xb3, 0x54, 0xd6, 0x66, 0x17, 0xcd, 0x1d, 0xc4, 0xac, 0x7f, 0x25,
	0xb1, 0x25, 0x3c, 0x2c, 0x61, 0x24, 0xb6, 0xca, 0x73, 0xe8, 0x65, 0x62, 0x16, 0xad, 0x27, 0xf1,
	0x71, 0x92, 0x6e, 0x26, 0x66, 0x77, 0xab, 0x61, 0xfe, 0x07, 0xc8, 0xab, 0x68, 0x99, 0xcc, 0x63,
	0x3e, 0x63, 0x01, 0x60, 0x23, 0x3f, 0xaf, 0xbe, 0x19, 0x40, 0xd7, 0x30, 0x6d, 0x6a, 0x45, 0x07,
	0x15, 0xa6, 0x77, 0x2d, 0x3a, 0x06, 0x7f, 0x2a, 0xd9, 0xc3, 0x82, 0xf1, 0xa4, 0x0a, 0xba, 0x27,
	0x8d, 0x51, 0x83, 0xae, 0x81, 0xcf, 0x4d, 0xaf, 0x37, 0xe8, 0x87, 0xdf, 0xc1, 0xb5, 0xcf, 0xf7,
	0xcf, 0xcd, 0x1f, 0x82, 0x27, 0x0a, 0x26, 0x63, 0x25, 0x24, 0x2e, 0xde, 0xa7, 0xab, 0x3b, 0x09,
	0xc0, 0xcd, 0x44, 0x82, 0x94, 0x59, 0x7d, 0x7d, 0x0d, 0x7f, 0x37, 0xc1, 0xb5, 0x6f, 0xae, 0x5f,
	0x82, 0xc7, 0x39, 0xb3, 0x95, 0xf1, 0x4c, 0x5e, 0x03, 0x61, 0x5c, 0x31, 0x19, 0x6d, 0xac, 0x4c,
	0xd7, 0xf7, 0xe8, 0x00, 0x99, 0xbb, 0x27, 0x7b, 0x1b, 0xc3, 0xf0, 0xa9, 0xba, 0x9e, 0xdc, 0x41,
	0xf9, 0xde, 0x5a, 0x5e, 0x8f, 0x7f, 0x0a, 0xfb, 0x7a, 0xd8, 0x54, 0xb2, 0x2d, 0x47, 0x13, 0x1d,
	0x43, 0x4b, 0x6e, 0x78, 0x46, 0x30, 0x88, 0xb3, 0x4c, 0x3c, 0x46, 0xf9, 0x22, 0x53, 0x69, 0x84,
	0x79, 0x6a, 0xa1, 0xbc, 0x8f, 0xf8, 0xa5, 0x86, 0xcf, 0x74, 0xb2, 0xb6, 0x94, 0x98, 0xb2, 0xf6,
	0xb6, 0xf2, 0x52, 0xe7, 0x6d, 0x0c, 0xc3, 0x32, 0xce, 0x59, 0x94, 0x88, 0x85, 0x4e, 0x7a, 0x54,
	0x88, 0x94, 0xab, 0x12, 0xc3, 0xd3, 0xa2, 0x7b, 0x9a, 0x3a, 0x37, 0xcc, 0x35, 0x12, 0xfa, 0xbb,
	0xad, 0x9e, 0xab, 0x94, 0x33, 0xae, 0x6a, 0x87, 0xc9, 0xd2, 0xd0, 0x38, 0x2c, 0x67, 0x3d, 0xef,
	0xe0, 0xa0, 0x2c, 0x58, 0x92, 0x4e, 0xd3, 0x64, 0xbb, 0x8f, 0x8f, 0xae, 0xfd, 0x9a, 0xde, 0xec,
	0xf5, 0x1e, 0xfe, 0xfb, 0xdb, 0x27, 0xd9, 0x34, 0x5d, 0xb2, 0x32, 0x80, 0x13, 0x67, 0xe4, 0xd3,
	0x83, 0x6d, 0xa7, 0xa5, 0x75, 0xd4, 0x85, 0x9a, 0x33, 0x59, 0x37, 0xea, 0x98, 0xa8, 0x23, 0x66,
	0xcb, 0x87, 0xd0, 0xc6, 0xf5, 0x94, 0x18, 0xbf, 0xce, 0x29, 0xe8, 0xbf, 0x00, 0x6e, 0xa6, 0xa4,
	0x96, 0xd1, 0xe3, 0xda, 0x87, 0xb1, 0xab, 0x2c, 0x62, 0xa5, 0x98, 0xe4, 0x41, 0x0f, 0x93, 0x32,
	0xb4, 0x24, 0xba, 0xae, 0x0d, 0x45, 0x5e, 0x40, 0x1f, 0xbf, 0x36, 0x2a, 0x98, 0x34, 0x8f, 0xd4,
	0xc7, 0xd5, 0x77, 0x11, 0xbd, 0x66, 0x52, 0x3f, 0x51, 0x78, 0x01, 0x6d, 0xd3, 0x4b, 0x87, 0xef,
	0x7e, 0x99, 0x24, 0x18, 0x3e, 0x8f, 0xe2, 0x99, 0x0c, 0xc0, 0x79, 0x2c, 0x96, 0x36, 0x6d, 0xfa,
	0xa8, 0x83, 0xbc, 0x19, 0xaa, 0xfa, 0x3a, 0x69, 0xe3, 0xbf, 0xf5, 0xed, 0x9f, 0x00, 0x00, 0x00,
	0xff, 0xff, 0xc7, 0x98, 0x87, 0xc6, 0x68, 0x05, 0x00, 0x00,
}
