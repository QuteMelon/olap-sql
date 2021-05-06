// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api/proto/request.proto

package proto

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

type METRIC_TYPE int32

const (
	METRIC_TYPE_METRIC_TYPE_UNKNOWN        METRIC_TYPE = 0
	METRIC_TYPE_METRIC_TYPE_VALUE          METRIC_TYPE = 1
	METRIC_TYPE_METRIC_TYPE_COUNT          METRIC_TYPE = 2
	METRIC_TYPE_METRIC_TYPE_DISTINCT_COUNT METRIC_TYPE = 3
	METRIC_TYPE_METRIC_TYPE_SUM            METRIC_TYPE = 4
	METRIC_TYPE_METRIC_TYPE_ADD            METRIC_TYPE = 5
	METRIC_TYPE_METRIC_TYPE_SUBTRACT       METRIC_TYPE = 6
	METRIC_TYPE_METRIC_TYPE_MULTIPLY       METRIC_TYPE = 7
	METRIC_TYPE_METRIC_TYPE_DIVIDE         METRIC_TYPE = 8
	METRIC_TYPE_METRIC_TYPE_POST           METRIC_TYPE = 99
	METRIC_TYPE_METRIC_TYPE_EXTENSION      METRIC_TYPE = 100
)

var METRIC_TYPE_name = map[int32]string{
	0:   "METRIC_TYPE_UNKNOWN",
	1:   "METRIC_TYPE_VALUE",
	2:   "METRIC_TYPE_COUNT",
	3:   "METRIC_TYPE_DISTINCT_COUNT",
	4:   "METRIC_TYPE_SUM",
	5:   "METRIC_TYPE_ADD",
	6:   "METRIC_TYPE_SUBTRACT",
	7:   "METRIC_TYPE_MULTIPLY",
	8:   "METRIC_TYPE_DIVIDE",
	99:  "METRIC_TYPE_POST",
	100: "METRIC_TYPE_EXTENSION",
}

var METRIC_TYPE_value = map[string]int32{
	"METRIC_TYPE_UNKNOWN":        0,
	"METRIC_TYPE_VALUE":          1,
	"METRIC_TYPE_COUNT":          2,
	"METRIC_TYPE_DISTINCT_COUNT": 3,
	"METRIC_TYPE_SUM":            4,
	"METRIC_TYPE_ADD":            5,
	"METRIC_TYPE_SUBTRACT":       6,
	"METRIC_TYPE_MULTIPLY":       7,
	"METRIC_TYPE_DIVIDE":         8,
	"METRIC_TYPE_POST":           99,
	"METRIC_TYPE_EXTENSION":      100,
}

func (x METRIC_TYPE) String() string {
	return proto.EnumName(METRIC_TYPE_name, int32(x))
}

func (METRIC_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{0}
}

type FILTER_OPERATOR_TYPE int32

const (
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_UNKNOWN        FILTER_OPERATOR_TYPE = 0
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_IN             FILTER_OPERATOR_TYPE = 1
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_NOT_IN         FILTER_OPERATOR_TYPE = 2
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_LESS_EQUALS    FILTER_OPERATOR_TYPE = 3
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_LESS           FILTER_OPERATOR_TYPE = 4
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_GREATER_EQUALS FILTER_OPERATOR_TYPE = 5
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_GREATER        FILTER_OPERATOR_TYPE = 6
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_LIKE           FILTER_OPERATOR_TYPE = 7
	FILTER_OPERATOR_TYPE_FILTER_OPERATOR_EXTENSION      FILTER_OPERATOR_TYPE = 20
)

var FILTER_OPERATOR_TYPE_name = map[int32]string{
	0:  "FILTER_OPERATOR_UNKNOWN",
	1:  "FILTER_OPERATOR_IN",
	2:  "FILTER_OPERATOR_NOT_IN",
	3:  "FILTER_OPERATOR_LESS_EQUALS",
	4:  "FILTER_OPERATOR_LESS",
	5:  "FILTER_OPERATOR_GREATER_EQUALS",
	6:  "FILTER_OPERATOR_GREATER",
	7:  "FILTER_OPERATOR_LIKE",
	20: "FILTER_OPERATOR_EXTENSION",
}

var FILTER_OPERATOR_TYPE_value = map[string]int32{
	"FILTER_OPERATOR_UNKNOWN":        0,
	"FILTER_OPERATOR_IN":             1,
	"FILTER_OPERATOR_NOT_IN":         2,
	"FILTER_OPERATOR_LESS_EQUALS":    3,
	"FILTER_OPERATOR_LESS":           4,
	"FILTER_OPERATOR_GREATER_EQUALS": 5,
	"FILTER_OPERATOR_GREATER":        6,
	"FILTER_OPERATOR_LIKE":           7,
	"FILTER_OPERATOR_EXTENSION":      20,
}

func (x FILTER_OPERATOR_TYPE) String() string {
	return proto.EnumName(FILTER_OPERATOR_TYPE_name, int32(x))
}

func (FILTER_OPERATOR_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{1}
}

type VALUE_TYPE int32

const (
	VALUE_TYPE_VALUE_UNKNOWN VALUE_TYPE = 0
	VALUE_TYPE_VALUE_STRING  VALUE_TYPE = 1
	VALUE_TYPE_VALUE_INTEGER VALUE_TYPE = 2
	VALUE_TYPE_VALUE_FLOAT   VALUE_TYPE = 3
)

var VALUE_TYPE_name = map[int32]string{
	0: "VALUE_UNKNOWN",
	1: "VALUE_STRING",
	2: "VALUE_INTEGER",
	3: "VALUE_FLOAT",
}

var VALUE_TYPE_value = map[string]int32{
	"VALUE_UNKNOWN": 0,
	"VALUE_STRING":  1,
	"VALUE_INTEGER": 2,
	"VALUE_FLOAT":   3,
}

func (x VALUE_TYPE) String() string {
	return proto.EnumName(VALUE_TYPE_name, int32(x))
}

func (VALUE_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{2}
}

type DATA_SOURCE_TYPE int32

const (
	DATA_SOURCE_TYPE_DATA_SOURCE_UNKNOWN    DATA_SOURCE_TYPE = 0
	DATA_SOURCE_TYPE_DATA_SOURCE_CLICKHOUSE DATA_SOURCE_TYPE = 1
	DATA_SOURCE_TYPE_DATA_SOURCE_DRUID      DATA_SOURCE_TYPE = 2
	DATA_SOURCE_TYPE_DATA_SOURCE_KYLIN      DATA_SOURCE_TYPE = 3
	DATA_SOURCE_TYPE_DATA_SOURCE_PRESTO     DATA_SOURCE_TYPE = 4
)

var DATA_SOURCE_TYPE_name = map[int32]string{
	0: "DATA_SOURCE_UNKNOWN",
	1: "DATA_SOURCE_CLICKHOUSE",
	2: "DATA_SOURCE_DRUID",
	3: "DATA_SOURCE_KYLIN",
	4: "DATA_SOURCE_PRESTO",
}

var DATA_SOURCE_TYPE_value = map[string]int32{
	"DATA_SOURCE_UNKNOWN":    0,
	"DATA_SOURCE_CLICKHOUSE": 1,
	"DATA_SOURCE_DRUID":      2,
	"DATA_SOURCE_KYLIN":      3,
	"DATA_SOURCE_PRESTO":     4,
}

func (x DATA_SOURCE_TYPE) String() string {
	return proto.EnumName(DATA_SOURCE_TYPE_name, int32(x))
}

func (DATA_SOURCE_TYPE) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{3}
}

type Metric struct {
	Type                 METRIC_TYPE `protobuf:"varint,1,opt,name=type,proto3,enum=proto.METRIC_TYPE" json:"type,omitempty"`
	FieldName            string      `protobuf:"bytes,2,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Name                 string      `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ExtensionValue       string      `protobuf:"bytes,4,opt,name=extension_value,json=extensionValue,proto3" json:"extension_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetType() METRIC_TYPE {
	if m != nil {
		return m.Type
	}
	return METRIC_TYPE_METRIC_TYPE_UNKNOWN
}

func (m *Metric) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Metric) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Metric) GetExtensionValue() string {
	if m != nil {
		return m.ExtensionValue
	}
	return ""
}

type Dimension struct {
	FieldName            string   `protobuf:"bytes,1,opt,name=field_name,json=fieldName,proto3" json:"field_name,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dimension) Reset()         { *m = Dimension{} }
func (m *Dimension) String() string { return proto.CompactTextString(m) }
func (*Dimension) ProtoMessage()    {}
func (*Dimension) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{1}
}

func (m *Dimension) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dimension.Unmarshal(m, b)
}
func (m *Dimension) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dimension.Marshal(b, m, deterministic)
}
func (m *Dimension) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dimension.Merge(m, src)
}
func (m *Dimension) XXX_Size() int {
	return xxx_messageInfo_Dimension.Size(m)
}
func (m *Dimension) XXX_DiscardUnknown() {
	xxx_messageInfo_Dimension.DiscardUnknown(m)
}

var xxx_messageInfo_Dimension proto.InternalMessageInfo

func (m *Dimension) GetFieldName() string {
	if m != nil {
		return m.FieldName
	}
	return ""
}

func (m *Dimension) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Filter struct {
	OperatorType         FILTER_OPERATOR_TYPE `protobuf:"varint,1,opt,name=operator_type,json=operatorType,proto3,enum=proto.FILTER_OPERATOR_TYPE" json:"operator_type,omitempty"`
	ValueType            VALUE_TYPE           `protobuf:"varint,2,opt,name=value_type,json=valueType,proto3,enum=proto.VALUE_TYPE" json:"value_type,omitempty"`
	Name                 string               `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Value                []string             `protobuf:"bytes,4,rep,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Filter) Reset()         { *m = Filter{} }
func (m *Filter) String() string { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()    {}
func (*Filter) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{2}
}

func (m *Filter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Filter.Unmarshal(m, b)
}
func (m *Filter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Filter.Marshal(b, m, deterministic)
}
func (m *Filter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filter.Merge(m, src)
}
func (m *Filter) XXX_Size() int {
	return xxx_messageInfo_Filter.Size(m)
}
func (m *Filter) XXX_DiscardUnknown() {
	xxx_messageInfo_Filter.DiscardUnknown(m)
}

var xxx_messageInfo_Filter proto.InternalMessageInfo

func (m *Filter) GetOperatorType() FILTER_OPERATOR_TYPE {
	if m != nil {
		return m.OperatorType
	}
	return FILTER_OPERATOR_TYPE_FILTER_OPERATOR_UNKNOWN
}

func (m *Filter) GetValueType() VALUE_TYPE {
	if m != nil {
		return m.ValueType
	}
	return VALUE_TYPE_VALUE_UNKNOWN
}

func (m *Filter) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Filter) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type JoinOn struct {
	Key1                 string   `protobuf:"bytes,1,opt,name=key1,proto3" json:"key1,omitempty"`
	Key2                 string   `protobuf:"bytes,2,opt,name=key2,proto3" json:"key2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinOn) Reset()         { *m = JoinOn{} }
func (m *JoinOn) String() string { return proto.CompactTextString(m) }
func (*JoinOn) ProtoMessage()    {}
func (*JoinOn) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{3}
}

func (m *JoinOn) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinOn.Unmarshal(m, b)
}
func (m *JoinOn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinOn.Marshal(b, m, deterministic)
}
func (m *JoinOn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinOn.Merge(m, src)
}
func (m *JoinOn) XXX_Size() int {
	return xxx_messageInfo_JoinOn.Size(m)
}
func (m *JoinOn) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinOn.DiscardUnknown(m)
}

var xxx_messageInfo_JoinOn proto.InternalMessageInfo

func (m *JoinOn) GetKey1() string {
	if m != nil {
		return m.Key1
	}
	return ""
}

func (m *JoinOn) GetKey2() string {
	if m != nil {
		return m.Key2
	}
	return ""
}

type Join struct {
	Table1               string    `protobuf:"bytes,1,opt,name=table1,proto3" json:"table1,omitempty"`
	Table2               string    `protobuf:"bytes,2,opt,name=table2,proto3" json:"table2,omitempty"`
	On                   []*JoinOn `protobuf:"bytes,3,rep,name=on,proto3" json:"on,omitempty"`
	Filters              []*Filter `protobuf:"bytes,4,rep,name=filters,proto3" json:"filters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Join) Reset()         { *m = Join{} }
func (m *Join) String() string { return proto.CompactTextString(m) }
func (*Join) ProtoMessage()    {}
func (*Join) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{4}
}

func (m *Join) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Join.Unmarshal(m, b)
}
func (m *Join) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Join.Marshal(b, m, deterministic)
}
func (m *Join) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Join.Merge(m, src)
}
func (m *Join) XXX_Size() int {
	return xxx_messageInfo_Join.Size(m)
}
func (m *Join) XXX_DiscardUnknown() {
	xxx_messageInfo_Join.DiscardUnknown(m)
}

var xxx_messageInfo_Join proto.InternalMessageInfo

func (m *Join) GetTable1() string {
	if m != nil {
		return m.Table1
	}
	return ""
}

func (m *Join) GetTable2() string {
	if m != nil {
		return m.Table2
	}
	return ""
}

func (m *Join) GetOn() []*JoinOn {
	if m != nil {
		return m.On
	}
	return nil
}

func (m *Join) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

type DataSource struct {
	Type                 DATA_SOURCE_TYPE `protobuf:"varint,1,opt,name=type,proto3,enum=proto.DATA_SOURCE_TYPE" json:"type,omitempty"`
	Name                 string           `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	SubRequest           *Request         `protobuf:"bytes,3,opt,name=sub_request,json=subRequest,proto3" json:"sub_request,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DataSource) Reset()         { *m = DataSource{} }
func (m *DataSource) String() string { return proto.CompactTextString(m) }
func (*DataSource) ProtoMessage()    {}
func (*DataSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{5}
}

func (m *DataSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DataSource.Unmarshal(m, b)
}
func (m *DataSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DataSource.Marshal(b, m, deterministic)
}
func (m *DataSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DataSource.Merge(m, src)
}
func (m *DataSource) XXX_Size() int {
	return xxx_messageInfo_DataSource.Size(m)
}
func (m *DataSource) XXX_DiscardUnknown() {
	xxx_messageInfo_DataSource.DiscardUnknown(m)
}

var xxx_messageInfo_DataSource proto.InternalMessageInfo

func (m *DataSource) GetType() DATA_SOURCE_TYPE {
	if m != nil {
		return m.Type
	}
	return DATA_SOURCE_TYPE_DATA_SOURCE_UNKNOWN
}

func (m *DataSource) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DataSource) GetSubRequest() *Request {
	if m != nil {
		return m.SubRequest
	}
	return nil
}

type Request struct {
	Metrics              []*Metric    `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty"`
	Dimensions           []*Dimension `protobuf:"bytes,2,rep,name=dimensions,proto3" json:"dimensions,omitempty"`
	Filters              []*Filter    `protobuf:"bytes,3,rep,name=filters,proto3" json:"filters,omitempty"`
	Joins                []*Join      `protobuf:"bytes,4,rep,name=joins,proto3" json:"joins,omitempty"`
	DataSource           *DataSource  `protobuf:"bytes,5,opt,name=data_source,json=dataSource,proto3" json:"data_source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_336e959edb61afc4, []int{6}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetMetrics() []*Metric {
	if m != nil {
		return m.Metrics
	}
	return nil
}

func (m *Request) GetDimensions() []*Dimension {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *Request) GetFilters() []*Filter {
	if m != nil {
		return m.Filters
	}
	return nil
}

func (m *Request) GetJoins() []*Join {
	if m != nil {
		return m.Joins
	}
	return nil
}

func (m *Request) GetDataSource() *DataSource {
	if m != nil {
		return m.DataSource
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.METRIC_TYPE", METRIC_TYPE_name, METRIC_TYPE_value)
	proto.RegisterEnum("proto.FILTER_OPERATOR_TYPE", FILTER_OPERATOR_TYPE_name, FILTER_OPERATOR_TYPE_value)
	proto.RegisterEnum("proto.VALUE_TYPE", VALUE_TYPE_name, VALUE_TYPE_value)
	proto.RegisterEnum("proto.DATA_SOURCE_TYPE", DATA_SOURCE_TYPE_name, DATA_SOURCE_TYPE_value)
	proto.RegisterType((*Metric)(nil), "proto.Metric")
	proto.RegisterType((*Dimension)(nil), "proto.Dimension")
	proto.RegisterType((*Filter)(nil), "proto.Filter")
	proto.RegisterType((*JoinOn)(nil), "proto.JoinOn")
	proto.RegisterType((*Join)(nil), "proto.Join")
	proto.RegisterType((*DataSource)(nil), "proto.DataSource")
	proto.RegisterType((*Request)(nil), "proto.Request")
}

func init() { proto.RegisterFile("api/proto/request.proto", fileDescriptor_336e959edb61afc4) }

var fileDescriptor_336e959edb61afc4 = []byte{
	// 833 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x55, 0x51, 0x77, 0xa3, 0x44,
	0x14, 0x16, 0x48, 0x52, 0x73, 0xb3, 0x6d, 0xa7, 0xb3, 0xd9, 0x34, 0xdb, 0x9e, 0xae, 0x95, 0x07,
	0xb7, 0xa7, 0x9e, 0x6d, 0xd6, 0xf8, 0xee, 0x91, 0x0d, 0xb4, 0x62, 0x29, 0xc4, 0x61, 0xa8, 0xd6,
	0x17, 0x0e, 0x49, 0x66, 0x15, 0x4d, 0x20, 0x0b, 0x44, 0xdd, 0x97, 0xfd, 0x05, 0xfa, 0x0b, 0xf4,
	0xd9, 0x5f, 0xe6, 0x0f, 0xf1, 0x30, 0x40, 0x99, 0x62, 0xdc, 0x27, 0xe6, 0x7e, 0xdf, 0xfd, 0xb8,
	0xf7, 0x7e, 0xdc, 0x49, 0xe0, 0x30, 0x58, 0x87, 0xa3, 0x75, 0x12, 0x67, 0xf1, 0x28, 0x61, 0x6f,
	0x36, 0x2c, 0xcd, 0x2e, 0x78, 0x84, 0xdb, 0xfc, 0xa1, 0xfe, 0x2e, 0x41, 0xe7, 0x86, 0x65, 0x49,
	0x38, 0xc7, 0x9f, 0x40, 0x2b, 0x7b, 0xbb, 0x66, 0x43, 0xe9, 0x54, 0x3a, 0xdb, 0x1b, 0xe3, 0x22,
	0xef, 0xe2, 0xc6, 0xa0, 0xc4, 0x9c, 0xf8, 0xf4, 0x6e, 0x6a, 0x10, 0xce, 0xe3, 0x13, 0x80, 0xd7,
	0x21, 0x5b, 0x2e, 0xfc, 0x28, 0x58, 0xb1, 0xa1, 0x7c, 0x2a, 0x9d, 0x75, 0x49, 0x97, 0x23, 0x76,
	0xb0, 0x62, 0x18, 0x43, 0x8b, 0x13, 0x0a, 0x27, 0xf8, 0x19, 0x3f, 0x87, 0x7d, 0xf6, 0x5b, 0xc6,
	0xa2, 0x34, 0x8c, 0x23, 0xff, 0x97, 0x60, 0xb9, 0x61, 0xc3, 0x16, 0xa7, 0xf7, 0xee, 0xe1, 0xdb,
	0x1c, 0x55, 0xbf, 0x80, 0xae, 0x1e, 0xae, 0x0a, 0xa4, 0x51, 0x48, 0xfa, 0xbf, 0x42, 0x72, 0x5d,
	0x48, 0xfd, 0x5b, 0x82, 0xce, 0x65, 0xb8, 0xcc, 0x58, 0x82, 0xbf, 0x84, 0xdd, 0x78, 0xcd, 0x92,
	0x20, 0x8b, 0x13, 0x5f, 0x98, 0xeb, 0xb8, 0x9c, 0xeb, 0xd2, 0xb4, 0xa8, 0x41, 0x7c, 0x67, 0x6a,
	0x10, 0x8d, 0x3a, 0xa4, 0x18, 0xf0, 0x51, 0xa5, 0xa0, 0xf9, 0xa0, 0x2f, 0x01, 0x78, 0xaf, 0x85,
	0x5c, 0xe6, 0xf2, 0x83, 0x52, 0x7e, 0xab, 0x59, 0x9e, 0x51, 0x88, 0xba, 0x3c, 0x89, 0x2b, 0xb6,
	0xcd, 0xde, 0x87, 0x76, 0x35, 0xb1, 0x72, 0xd6, 0x25, 0x45, 0xa0, 0xbe, 0x84, 0xce, 0xd7, 0x71,
	0x18, 0x39, 0x51, 0xae, 0xf9, 0x99, 0xbd, 0xfd, 0xac, 0x9c, 0x8f, 0x9f, 0x4b, 0x6c, 0x5c, 0x8d,
	0x96, 0x9f, 0xd5, 0x77, 0xd0, 0xca, 0x15, 0x78, 0x00, 0x9d, 0x2c, 0x98, 0x2d, 0x59, 0xa5, 0x28,
	0xa3, 0x7b, 0xbc, 0x52, 0x95, 0x11, 0x3e, 0x01, 0x39, 0x8e, 0x86, 0xca, 0xa9, 0x72, 0xd6, 0x1b,
	0xef, 0x96, 0xdd, 0x17, 0xa5, 0x89, 0x1c, 0x47, 0xf8, 0x39, 0xec, 0xbc, 0xe6, 0x86, 0xa5, 0xbc,
	0xc1, 0x3a, 0xa7, 0xb0, 0x91, 0x54, 0xac, 0xfa, 0x0e, 0x40, 0x0f, 0xb2, 0xc0, 0x8d, 0x37, 0xc9,
	0x9c, 0xe1, 0x4f, 0x1f, 0x2c, 0xcb, 0x61, 0xa9, 0xd1, 0x35, 0xaa, 0xf9, 0xae, 0xe3, 0x91, 0x89,
	0x21, 0x6e, 0xcc, 0x96, 0x2f, 0x85, 0x47, 0xd0, 0x4b, 0x37, 0x33, 0xbf, 0x5c, 0x4a, 0xee, 0x58,
	0x6f, 0xbc, 0x57, 0xbe, 0x87, 0x14, 0x28, 0x81, 0x74, 0x33, 0x2b, 0xcf, 0xea, 0x3f, 0x12, 0xec,
	0x94, 0xe7, 0xbc, 0xe9, 0x15, 0x5f, 0xda, 0x74, 0x28, 0x3d, 0x68, 0xba, 0x58, 0x65, 0x52, 0xb1,
	0xf9, 0x27, 0x5c, 0x54, 0xfb, 0x94, 0x0e, 0x65, 0x9e, 0x8b, 0xaa, 0x66, 0x2b, 0x82, 0x08, 0x39,
	0xa2, 0x1f, 0xca, 0xfb, 0xfc, 0xc0, 0x1f, 0x43, 0xfb, 0xa7, 0x38, 0x8c, 0x2a, 0xdb, 0x7a, 0x82,
	0xb5, 0xa4, 0x60, 0xf0, 0x18, 0x7a, 0x8b, 0x20, 0x0b, 0xfc, 0x94, 0x7b, 0x36, 0x6c, 0xf3, 0x19,
	0xab, 0x0d, 0xaa, 0xcd, 0x24, 0xb0, 0xb8, 0x3f, 0x9f, 0xff, 0x29, 0x43, 0x4f, 0xb8, 0x73, 0xf8,
	0x10, 0x1e, 0x0b, 0xa1, 0xef, 0xd9, 0xd7, 0xb6, 0xf3, 0xad, 0x8d, 0x3e, 0xc0, 0x4f, 0xe0, 0x40,
	0x24, 0xf8, 0x42, 0x22, 0xa9, 0x09, 0x4f, 0x1c, 0xcf, 0xa6, 0x48, 0xc6, 0xcf, 0xe0, 0x48, 0x84,
	0x75, 0xd3, 0xa5, 0xa6, 0x3d, 0xa1, 0x25, 0xaf, 0xe0, 0xc7, 0xb0, 0x2f, 0xf2, 0xae, 0x77, 0x83,
	0x5a, 0x4d, 0x50, 0xd3, 0x75, 0xd4, 0xc6, 0x43, 0xe8, 0x3f, 0xcc, 0x7c, 0x45, 0x89, 0x36, 0xa1,
	0xa8, 0xd3, 0x64, 0x6e, 0x3c, 0x8b, 0x9a, 0x53, 0xeb, 0x0e, 0xed, 0xe0, 0x01, 0xe0, 0x87, 0xd5,
	0x6f, 0x4d, 0xdd, 0x40, 0x1f, 0xe2, 0x3e, 0x20, 0x11, 0x9f, 0x3a, 0x2e, 0x45, 0x73, 0xfc, 0x14,
	0x9e, 0x88, 0xa8, 0xf1, 0x1d, 0x35, 0x6c, 0xd7, 0x74, 0x6c, 0xb4, 0x38, 0xff, 0x4b, 0x86, 0xfe,
	0xb6, 0x9b, 0x8b, 0x8f, 0xe1, 0xb0, 0x89, 0xd7, 0x56, 0x0d, 0x00, 0x37, 0x49, 0xd3, 0x46, 0x12,
	0x3e, 0x82, 0x41, 0x13, 0xb7, 0x1d, 0x9a, 0x73, 0x32, 0xfe, 0x08, 0x8e, 0x9b, 0x9c, 0x65, 0xb8,
	0xae, 0x6f, 0x7c, 0xe3, 0x69, 0x96, 0x8b, 0x94, 0x7c, 0xda, 0x6d, 0x09, 0xa8, 0x85, 0x55, 0x78,
	0xd6, 0x64, 0xae, 0x88, 0xa1, 0xe5, 0x40, 0xa9, 0x6e, 0x6f, 0xeb, 0xb7, 0xcc, 0x29, 0x8c, 0xfc,
	0xcf, 0xab, 0xcd, 0x6b, 0x03, 0xed, 0xe0, 0x13, 0x78, 0xda, 0x64, 0x6a, 0x7b, 0xfa, 0xe7, 0x1e,
	0x40, 0xfd, 0xc3, 0x84, 0x0f, 0x60, 0xb7, 0x88, 0x6a, 0x27, 0x10, 0x3c, 0x2a, 0x20, 0x97, 0x12,
	0xd3, 0xbe, 0x42, 0x52, 0x9d, 0x64, 0xda, 0xd4, 0xb8, 0x32, 0x08, 0x92, 0xf1, 0x3e, 0xf4, 0x0a,
	0xe8, 0xd2, 0x72, 0x34, 0x8a, 0x94, 0xf3, 0x3f, 0x24, 0x40, 0xcd, 0xab, 0x9d, 0x2f, 0xa6, 0x88,
	0xd5, 0x35, 0x8e, 0x60, 0x20, 0x12, 0x13, 0xcb, 0x9c, 0x5c, 0x7f, 0xe5, 0x78, 0x6e, 0xb9, 0x9d,
	0x22, 0xa7, 0x13, 0xcf, 0xd4, 0x91, 0xdc, 0x84, 0xaf, 0xef, 0x2c, 0xd3, 0x46, 0x4a, 0xfe, 0xdd,
	0x44, 0x78, 0x4a, 0x0c, 0x97, 0x3a, 0xa8, 0xf5, 0x6a, 0xf4, 0xfd, 0x8b, 0xd3, 0x1f, 0xc2, 0xec,
	0xc7, 0xcd, 0xec, 0x62, 0x1e, 0xaf, 0x46, 0xc1, 0xaf, 0x41, 0xc6, 0x92, 0x79, 0xbc, 0x8c, 0x93,
	0x35, 0x8b, 0x46, 0xf1, 0x32, 0x58, 0xbf, 0x48, 0xdf, 0x2c, 0x47, 0xf7, 0x7f, 0x7d, 0xb3, 0x0e,
	0x7f, 0x7c, 0xfe, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0x65, 0x47, 0x39, 0x0e, 0x07, 0x00,
	0x00,
}
