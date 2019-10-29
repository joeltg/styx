// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types/types.proto

package types

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

type SourceList struct {
	Sources              []*SourceList_Source `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *SourceList) Reset()         { *m = SourceList{} }
func (m *SourceList) String() string { return proto.CompactTextString(m) }
func (*SourceList) ProtoMessage()    {}
func (*SourceList) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0}
}

func (m *SourceList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceList.Unmarshal(m, b)
}
func (m *SourceList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceList.Marshal(b, m, deterministic)
}
func (m *SourceList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceList.Merge(m, src)
}
func (m *SourceList) XXX_Size() int {
	return xxx_messageInfo_SourceList.Size(m)
}
func (m *SourceList) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceList.DiscardUnknown(m)
}

var xxx_messageInfo_SourceList proto.InternalMessageInfo

func (m *SourceList) GetSources() []*SourceList_Source {
	if m != nil {
		return m.Sources
	}
	return nil
}

type SourceList_Source struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Index                uint32   `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Graph                string   `protobuf:"bytes,3,opt,name=graph,proto3" json:"graph,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SourceList_Source) Reset()         { *m = SourceList_Source{} }
func (m *SourceList_Source) String() string { return proto.CompactTextString(m) }
func (*SourceList_Source) ProtoMessage()    {}
func (*SourceList_Source) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{0, 0}
}

func (m *SourceList_Source) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceList_Source.Unmarshal(m, b)
}
func (m *SourceList_Source) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceList_Source.Marshal(b, m, deterministic)
}
func (m *SourceList_Source) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceList_Source.Merge(m, src)
}
func (m *SourceList_Source) XXX_Size() int {
	return xxx_messageInfo_SourceList_Source.Size(m)
}
func (m *SourceList_Source) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceList_Source.DiscardUnknown(m)
}

var xxx_messageInfo_SourceList_Source proto.InternalMessageInfo

func (m *SourceList_Source) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *SourceList_Source) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *SourceList_Source) GetGraph() string {
	if m != nil {
		return m.Graph
	}
	return ""
}

type Blank struct {
	Origin               uint64   `protobuf:"varint,1,opt,name=origin,proto3" json:"origin,omitempty"`
	Id                   string   `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Blank) Reset()         { *m = Blank{} }
func (m *Blank) String() string { return proto.CompactTextString(m) }
func (*Blank) ProtoMessage()    {}
func (*Blank) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{1}
}

func (m *Blank) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Blank.Unmarshal(m, b)
}
func (m *Blank) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Blank.Marshal(b, m, deterministic)
}
func (m *Blank) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Blank.Merge(m, src)
}
func (m *Blank) XXX_Size() int {
	return xxx_messageInfo_Blank.Size(m)
}
func (m *Blank) XXX_DiscardUnknown() {
	xxx_messageInfo_Blank.DiscardUnknown(m)
}

var xxx_messageInfo_Blank proto.InternalMessageInfo

func (m *Blank) GetOrigin() uint64 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *Blank) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Literal struct {
	Value                string   `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Language             string   `protobuf:"bytes,2,opt,name=language,proto3" json:"language,omitempty"`
	Datatype             string   `protobuf:"bytes,3,opt,name=datatype,proto3" json:"datatype,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Literal) Reset()         { *m = Literal{} }
func (m *Literal) String() string { return proto.CompactTextString(m) }
func (*Literal) ProtoMessage()    {}
func (*Literal) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{2}
}

func (m *Literal) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Literal.Unmarshal(m, b)
}
func (m *Literal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Literal.Marshal(b, m, deterministic)
}
func (m *Literal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Literal.Merge(m, src)
}
func (m *Literal) XXX_Size() int {
	return xxx_messageInfo_Literal.Size(m)
}
func (m *Literal) XXX_DiscardUnknown() {
	xxx_messageInfo_Literal.DiscardUnknown(m)
}

var xxx_messageInfo_Literal proto.InternalMessageInfo

func (m *Literal) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Literal) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *Literal) GetDatatype() string {
	if m != nil {
		return m.Datatype
	}
	return ""
}

type Dataset struct {
	Multihash            []byte   `protobuf:"bytes,1,opt,name=multihash,proto3" json:"multihash,omitempty"`
	Graphs               []string `protobuf:"bytes,2,rep,name=graphs,proto3" json:"graphs,omitempty"`
	Count                uint32   `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	Size                 uint32   `protobuf:"varint,4,opt,name=size,proto3" json:"size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{3}
}

func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (m *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(m, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

func (m *Dataset) GetMultihash() []byte {
	if m != nil {
		return m.Multihash
	}
	return nil
}

func (m *Dataset) GetGraphs() []string {
	if m != nil {
		return m.Graphs
	}
	return nil
}

func (m *Dataset) GetCount() uint32 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Dataset) GetSize() uint32 {
	if m != nil {
		return m.Size
	}
	return 0
}

type Value struct {
	// Types that are valid to be assigned to Node:
	//	*Value_Iri
	//	*Value_Blank
	//	*Value_Literal
	//	*Value_Dataset
	Node                 isValue_Node `protobuf_oneof:"node"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{4}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

type isValue_Node interface {
	isValue_Node()
}

type Value_Iri struct {
	Iri string `protobuf:"bytes,1,opt,name=iri,proto3,oneof"`
}

type Value_Blank struct {
	Blank *Blank `protobuf:"bytes,2,opt,name=blank,proto3,oneof"`
}

type Value_Literal struct {
	Literal *Literal `protobuf:"bytes,3,opt,name=literal,proto3,oneof"`
}

type Value_Dataset struct {
	Dataset *Dataset `protobuf:"bytes,4,opt,name=dataset,proto3,oneof"`
}

func (*Value_Iri) isValue_Node() {}

func (*Value_Blank) isValue_Node() {}

func (*Value_Literal) isValue_Node() {}

func (*Value_Dataset) isValue_Node() {}

func (m *Value) GetNode() isValue_Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Value) GetIri() string {
	if x, ok := m.GetNode().(*Value_Iri); ok {
		return x.Iri
	}
	return ""
}

func (m *Value) GetBlank() *Blank {
	if x, ok := m.GetNode().(*Value_Blank); ok {
		return x.Blank
	}
	return nil
}

func (m *Value) GetLiteral() *Literal {
	if x, ok := m.GetNode().(*Value_Literal); ok {
		return x.Literal
	}
	return nil
}

func (m *Value) GetDataset() *Dataset {
	if x, ok := m.GetNode().(*Value_Dataset); ok {
		return x.Dataset
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Value) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Value_Iri)(nil),
		(*Value_Blank)(nil),
		(*Value_Literal)(nil),
		(*Value_Dataset)(nil),
	}
}

type Index struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Subject              uint64   `protobuf:"varint,2,opt,name=subject,proto3" json:"subject,omitempty"`
	Predicate            uint64   `protobuf:"varint,3,opt,name=predicate,proto3" json:"predicate,omitempty"`
	Object               uint64   `protobuf:"varint,4,opt,name=object,proto3" json:"object,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Index) Reset()         { *m = Index{} }
func (m *Index) String() string { return proto.CompactTextString(m) }
func (*Index) ProtoMessage()    {}
func (*Index) Descriptor() ([]byte, []int) {
	return fileDescriptor_2c0f90c600ad7e2e, []int{5}
}

func (m *Index) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Index.Unmarshal(m, b)
}
func (m *Index) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Index.Marshal(b, m, deterministic)
}
func (m *Index) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Index.Merge(m, src)
}
func (m *Index) XXX_Size() int {
	return xxx_messageInfo_Index.Size(m)
}
func (m *Index) XXX_DiscardUnknown() {
	xxx_messageInfo_Index.DiscardUnknown(m)
}

var xxx_messageInfo_Index proto.InternalMessageInfo

func (m *Index) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Index) GetSubject() uint64 {
	if m != nil {
		return m.Subject
	}
	return 0
}

func (m *Index) GetPredicate() uint64 {
	if m != nil {
		return m.Predicate
	}
	return 0
}

func (m *Index) GetObject() uint64 {
	if m != nil {
		return m.Object
	}
	return 0
}

func init() {
	proto.RegisterType((*SourceList)(nil), "types.SourceList")
	proto.RegisterType((*SourceList_Source)(nil), "types.SourceList.Source")
	proto.RegisterType((*Blank)(nil), "types.Blank")
	proto.RegisterType((*Literal)(nil), "types.Literal")
	proto.RegisterType((*Dataset)(nil), "types.Dataset")
	proto.RegisterType((*Value)(nil), "types.Value")
	proto.RegisterType((*Index)(nil), "types.Index")
}

func init() { proto.RegisterFile("types/types.proto", fileDescriptor_2c0f90c600ad7e2e) }

var fileDescriptor_2c0f90c600ad7e2e = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0x4b, 0xea, 0xdb, 0x30,
	0x10, 0xc6, 0xe3, 0x77, 0x3d, 0x79, 0x40, 0x45, 0x29, 0x22, 0x74, 0x61, 0x4c, 0x17, 0xa6, 0x8b,
	0x04, 0xdc, 0x1b, 0x84, 0x2c, 0x52, 0xc8, 0x4a, 0x85, 0x76, 0xad, 0x58, 0xc2, 0x51, 0xeb, 0xda,
	0xc6, 0x92, 0x4b, 0xdb, 0x03, 0xf4, 0x1e, 0xbd, 0x69, 0xd1, 0x48, 0x89, 0xe1, 0xbf, 0x31, 0xfa,
	0x8d, 0xbe, 0x19, 0x7f, 0x33, 0x23, 0x78, 0x6d, 0x7e, 0x8f, 0x52, 0x1f, 0xf1, 0x7b, 0x18, 0xa7,
	0xc1, 0x0c, 0x24, 0x41, 0x28, 0xff, 0x06, 0x00, 0x9f, 0x87, 0x79, 0x6a, 0xe4, 0x55, 0x69, 0x43,
	0x6a, 0xc8, 0x34, 0x92, 0xa6, 0x41, 0x11, 0x55, 0xeb, 0x9a, 0x1e, 0x5c, 0xd2, 0xa2, 0xf1, 0x47,
	0xf6, 0x10, 0xee, 0xcf, 0x90, 0xba, 0x10, 0xd9, 0x41, 0xa8, 0x04, 0x0d, 0x8a, 0xa0, 0x8a, 0x59,
	0xa8, 0x04, 0x79, 0x03, 0x89, 0xea, 0x85, 0xfc, 0x45, 0xc3, 0x22, 0xa8, 0xb6, 0xcc, 0x81, 0x8d,
	0xb6, 0x13, 0x1f, 0xef, 0x34, 0x2a, 0x82, 0x2a, 0x67, 0x0e, 0xca, 0x23, 0x24, 0xa7, 0x8e, 0xf7,
	0xdf, 0xc9, 0x5b, 0x48, 0x87, 0x49, 0xb5, 0xaa, 0xf7, 0x85, 0x3c, 0xf9, 0xe2, 0x21, 0xe6, 0x84,
	0x4a, 0x94, 0x5f, 0x21, 0xbb, 0x2a, 0x23, 0x27, 0xde, 0xd9, 0x8a, 0x3f, 0x79, 0x37, 0x4b, 0xcc,
	0xc8, 0x99, 0x03, 0xb2, 0x87, 0x57, 0x1d, 0xef, 0xdb, 0x99, 0xb7, 0xd2, 0xa7, 0x3d, 0xd9, 0xde,
	0x09, 0x6e, 0xb8, 0xed, 0xcd, 0xdb, 0x78, 0x72, 0xa9, 0x20, 0x3b, 0x73, 0xc3, 0xb5, 0x34, 0xe4,
	0x1d, 0xe4, 0x3f, 0xe6, 0xce, 0xa8, 0x3b, 0xd7, 0x77, 0x2c, 0xbe, 0x61, 0x4b, 0xc0, 0x3a, 0x45,
	0xef, 0x9a, 0x86, 0x45, 0x54, 0xe5, 0xcc, 0x93, 0xb5, 0xd3, 0x0c, 0x73, 0x6f, 0xb0, 0xf2, 0x96,
	0x39, 0x20, 0x04, 0x62, 0xad, 0xfe, 0x48, 0x1a, 0x63, 0x10, 0xcf, 0xe5, 0xbf, 0x00, 0x92, 0x2f,
	0x68, 0x96, 0x40, 0xa4, 0x26, 0xe5, 0x1a, 0xb8, 0xac, 0x98, 0x05, 0xf2, 0x1e, 0x92, 0x9b, 0x1d,
	0x09, 0xba, 0x5f, 0xd7, 0x1b, 0xbf, 0x0a, 0x1c, 0xd3, 0x65, 0xc5, 0xdc, 0x25, 0xf9, 0x00, 0x59,
	0xe7, 0xe6, 0x80, 0xff, 0x5b, 0xd7, 0x3b, 0xaf, 0xf3, 0xd3, 0xb9, 0xac, 0xd8, 0x43, 0x60, 0xb5,
	0xc2, 0xb5, 0x86, 0x36, 0x16, 0xad, 0x6f, 0xd8, 0x6a, 0xbd, 0xe0, 0x94, 0x42, 0xdc, 0x0f, 0x42,
	0x96, 0x2d, 0x24, 0x9f, 0x70, 0x6f, 0x2f, 0xb7, 0x4b, 0x21, 0xd3, 0xf3, 0xed, 0x9b, 0x6c, 0x0c,
	0x1a, 0x8c, 0xd9, 0x03, 0xed, 0xd8, 0xc6, 0x49, 0x0a, 0xd5, 0x70, 0xe3, 0xc6, 0x1b, 0xb3, 0x25,
	0x80, 0x0b, 0x76, 0x69, 0xb1, 0x5f, 0x30, 0xd2, 0x2d, 0xc5, 0x87, 0xf9, 0xf1, 0x7f, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x15, 0x72, 0xcc, 0x95, 0xad, 0x02, 0x00, 0x00,
}
