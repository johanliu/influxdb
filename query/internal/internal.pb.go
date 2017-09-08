// Code generated by protoc-gen-gogo.
// source: internal/internal.proto
// DO NOT EDIT!

/*
Package query is a generated protocol buffer package.

It is generated from these files:
	internal/internal.proto

It has these top-level messages:
	Point
	Aux
	IteratorOptions
	Measurements
	Measurement
	Interval
	IteratorStats
	VarRef
*/
package query

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type Point struct {
	Name             *string        `protobuf:"bytes,1,req,name=Name" json:"Name,omitempty"`
	Tags             *string        `protobuf:"bytes,2,req,name=Tags" json:"Tags,omitempty"`
	Time             *int64         `protobuf:"varint,3,req,name=Time" json:"Time,omitempty"`
	Nil              *bool          `protobuf:"varint,4,req,name=Nil" json:"Nil,omitempty"`
	Aux              []*Aux         `protobuf:"bytes,5,rep,name=Aux" json:"Aux,omitempty"`
	Aggregated       *uint32        `protobuf:"varint,6,opt,name=Aggregated" json:"Aggregated,omitempty"`
	FloatValue       *float64       `protobuf:"fixed64,7,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64         `protobuf:"varint,8,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string        `protobuf:"bytes,9,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool          `protobuf:"varint,10,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue    *uint64        `protobuf:"varint,12,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
	Stats            *IteratorStats `protobuf:"bytes,11,opt,name=Stats" json:"Stats,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{0} }

func (m *Point) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Point) GetTags() string {
	if m != nil && m.Tags != nil {
		return *m.Tags
	}
	return ""
}

func (m *Point) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Point) GetNil() bool {
	if m != nil && m.Nil != nil {
		return *m.Nil
	}
	return false
}

func (m *Point) GetAux() []*Aux {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *Point) GetAggregated() uint32 {
	if m != nil && m.Aggregated != nil {
		return *m.Aggregated
	}
	return 0
}

func (m *Point) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Point) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Point) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Point) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Point) GetUnsignedValue() uint64 {
	if m != nil && m.UnsignedValue != nil {
		return *m.UnsignedValue
	}
	return 0
}

func (m *Point) GetStats() *IteratorStats {
	if m != nil {
		return m.Stats
	}
	return nil
}

type Aux struct {
	DataType         *int32   `protobuf:"varint,1,req,name=DataType" json:"DataType,omitempty"`
	FloatValue       *float64 `protobuf:"fixed64,2,opt,name=FloatValue" json:"FloatValue,omitempty"`
	IntegerValue     *int64   `protobuf:"varint,3,opt,name=IntegerValue" json:"IntegerValue,omitempty"`
	StringValue      *string  `protobuf:"bytes,4,opt,name=StringValue" json:"StringValue,omitempty"`
	BooleanValue     *bool    `protobuf:"varint,5,opt,name=BooleanValue" json:"BooleanValue,omitempty"`
	UnsignedValue    *uint64  `protobuf:"varint,6,opt,name=UnsignedValue" json:"UnsignedValue,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Aux) Reset()                    { *m = Aux{} }
func (m *Aux) String() string            { return proto.CompactTextString(m) }
func (*Aux) ProtoMessage()               {}
func (*Aux) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{1} }

func (m *Aux) GetDataType() int32 {
	if m != nil && m.DataType != nil {
		return *m.DataType
	}
	return 0
}

func (m *Aux) GetFloatValue() float64 {
	if m != nil && m.FloatValue != nil {
		return *m.FloatValue
	}
	return 0
}

func (m *Aux) GetIntegerValue() int64 {
	if m != nil && m.IntegerValue != nil {
		return *m.IntegerValue
	}
	return 0
}

func (m *Aux) GetStringValue() string {
	if m != nil && m.StringValue != nil {
		return *m.StringValue
	}
	return ""
}

func (m *Aux) GetBooleanValue() bool {
	if m != nil && m.BooleanValue != nil {
		return *m.BooleanValue
	}
	return false
}

func (m *Aux) GetUnsignedValue() uint64 {
	if m != nil && m.UnsignedValue != nil {
		return *m.UnsignedValue
	}
	return 0
}

type IteratorOptions struct {
	Expr             *string        `protobuf:"bytes,1,opt,name=Expr" json:"Expr,omitempty"`
	Aux              []string       `protobuf:"bytes,2,rep,name=Aux" json:"Aux,omitempty"`
	Fields           []*VarRef      `protobuf:"bytes,17,rep,name=Fields" json:"Fields,omitempty"`
	Sources          []*Measurement `protobuf:"bytes,3,rep,name=Sources" json:"Sources,omitempty"`
	Interval         *Interval      `protobuf:"bytes,4,opt,name=Interval" json:"Interval,omitempty"`
	Dimensions       []string       `protobuf:"bytes,5,rep,name=Dimensions" json:"Dimensions,omitempty"`
	GroupBy          []string       `protobuf:"bytes,19,rep,name=GroupBy" json:"GroupBy,omitempty"`
	Fill             *int32         `protobuf:"varint,6,opt,name=Fill" json:"Fill,omitempty"`
	FillValue        *float64       `protobuf:"fixed64,7,opt,name=FillValue" json:"FillValue,omitempty"`
	Condition        *string        `protobuf:"bytes,8,opt,name=Condition" json:"Condition,omitempty"`
	StartTime        *int64         `protobuf:"varint,9,opt,name=StartTime" json:"StartTime,omitempty"`
	EndTime          *int64         `protobuf:"varint,10,opt,name=EndTime" json:"EndTime,omitempty"`
	Location         *string        `protobuf:"bytes,21,opt,name=Location" json:"Location,omitempty"`
	Ascending        *bool          `protobuf:"varint,11,opt,name=Ascending" json:"Ascending,omitempty"`
	Limit            *int64         `protobuf:"varint,12,opt,name=Limit" json:"Limit,omitempty"`
	Offset           *int64         `protobuf:"varint,13,opt,name=Offset" json:"Offset,omitempty"`
	SLimit           *int64         `protobuf:"varint,14,opt,name=SLimit" json:"SLimit,omitempty"`
	SOffset          *int64         `protobuf:"varint,15,opt,name=SOffset" json:"SOffset,omitempty"`
	StripName        *bool          `protobuf:"varint,22,opt,name=StripName" json:"StripName,omitempty"`
	Dedupe           *bool          `protobuf:"varint,16,opt,name=Dedupe" json:"Dedupe,omitempty"`
	MaxSeriesN       *int64         `protobuf:"varint,18,opt,name=MaxSeriesN" json:"MaxSeriesN,omitempty"`
	Ordered          *bool          `protobuf:"varint,20,opt,name=Ordered" json:"Ordered,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *IteratorOptions) Reset()                    { *m = IteratorOptions{} }
func (m *IteratorOptions) String() string            { return proto.CompactTextString(m) }
func (*IteratorOptions) ProtoMessage()               {}
func (*IteratorOptions) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{2} }

func (m *IteratorOptions) GetExpr() string {
	if m != nil && m.Expr != nil {
		return *m.Expr
	}
	return ""
}

func (m *IteratorOptions) GetAux() []string {
	if m != nil {
		return m.Aux
	}
	return nil
}

func (m *IteratorOptions) GetFields() []*VarRef {
	if m != nil {
		return m.Fields
	}
	return nil
}

func (m *IteratorOptions) GetSources() []*Measurement {
	if m != nil {
		return m.Sources
	}
	return nil
}

func (m *IteratorOptions) GetInterval() *Interval {
	if m != nil {
		return m.Interval
	}
	return nil
}

func (m *IteratorOptions) GetDimensions() []string {
	if m != nil {
		return m.Dimensions
	}
	return nil
}

func (m *IteratorOptions) GetGroupBy() []string {
	if m != nil {
		return m.GroupBy
	}
	return nil
}

func (m *IteratorOptions) GetFill() int32 {
	if m != nil && m.Fill != nil {
		return *m.Fill
	}
	return 0
}

func (m *IteratorOptions) GetFillValue() float64 {
	if m != nil && m.FillValue != nil {
		return *m.FillValue
	}
	return 0
}

func (m *IteratorOptions) GetCondition() string {
	if m != nil && m.Condition != nil {
		return *m.Condition
	}
	return ""
}

func (m *IteratorOptions) GetStartTime() int64 {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return 0
}

func (m *IteratorOptions) GetEndTime() int64 {
	if m != nil && m.EndTime != nil {
		return *m.EndTime
	}
	return 0
}

func (m *IteratorOptions) GetLocation() string {
	if m != nil && m.Location != nil {
		return *m.Location
	}
	return ""
}

func (m *IteratorOptions) GetAscending() bool {
	if m != nil && m.Ascending != nil {
		return *m.Ascending
	}
	return false
}

func (m *IteratorOptions) GetLimit() int64 {
	if m != nil && m.Limit != nil {
		return *m.Limit
	}
	return 0
}

func (m *IteratorOptions) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *IteratorOptions) GetSLimit() int64 {
	if m != nil && m.SLimit != nil {
		return *m.SLimit
	}
	return 0
}

func (m *IteratorOptions) GetSOffset() int64 {
	if m != nil && m.SOffset != nil {
		return *m.SOffset
	}
	return 0
}

func (m *IteratorOptions) GetStripName() bool {
	if m != nil && m.StripName != nil {
		return *m.StripName
	}
	return false
}

func (m *IteratorOptions) GetDedupe() bool {
	if m != nil && m.Dedupe != nil {
		return *m.Dedupe
	}
	return false
}

func (m *IteratorOptions) GetMaxSeriesN() int64 {
	if m != nil && m.MaxSeriesN != nil {
		return *m.MaxSeriesN
	}
	return 0
}

func (m *IteratorOptions) GetOrdered() bool {
	if m != nil && m.Ordered != nil {
		return *m.Ordered
	}
	return false
}

type Measurements struct {
	Items            []*Measurement `protobuf:"bytes,1,rep,name=Items" json:"Items,omitempty"`
	XXX_unrecognized []byte         `json:"-"`
}

func (m *Measurements) Reset()                    { *m = Measurements{} }
func (m *Measurements) String() string            { return proto.CompactTextString(m) }
func (*Measurements) ProtoMessage()               {}
func (*Measurements) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{3} }

func (m *Measurements) GetItems() []*Measurement {
	if m != nil {
		return m.Items
	}
	return nil
}

type Measurement struct {
	Database         *string `protobuf:"bytes,1,opt,name=Database" json:"Database,omitempty"`
	RetentionPolicy  *string `protobuf:"bytes,2,opt,name=RetentionPolicy" json:"RetentionPolicy,omitempty"`
	Name             *string `protobuf:"bytes,3,opt,name=Name" json:"Name,omitempty"`
	Regex            *string `protobuf:"bytes,4,opt,name=Regex" json:"Regex,omitempty"`
	IsTarget         *bool   `protobuf:"varint,5,opt,name=IsTarget" json:"IsTarget,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Measurement) Reset()                    { *m = Measurement{} }
func (m *Measurement) String() string            { return proto.CompactTextString(m) }
func (*Measurement) ProtoMessage()               {}
func (*Measurement) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{4} }

func (m *Measurement) GetDatabase() string {
	if m != nil && m.Database != nil {
		return *m.Database
	}
	return ""
}

func (m *Measurement) GetRetentionPolicy() string {
	if m != nil && m.RetentionPolicy != nil {
		return *m.RetentionPolicy
	}
	return ""
}

func (m *Measurement) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Measurement) GetRegex() string {
	if m != nil && m.Regex != nil {
		return *m.Regex
	}
	return ""
}

func (m *Measurement) GetIsTarget() bool {
	if m != nil && m.IsTarget != nil {
		return *m.IsTarget
	}
	return false
}

type Interval struct {
	Duration         *int64 `protobuf:"varint,1,opt,name=Duration" json:"Duration,omitempty"`
	Offset           *int64 `protobuf:"varint,2,opt,name=Offset" json:"Offset,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Interval) Reset()                    { *m = Interval{} }
func (m *Interval) String() string            { return proto.CompactTextString(m) }
func (*Interval) ProtoMessage()               {}
func (*Interval) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{5} }

func (m *Interval) GetDuration() int64 {
	if m != nil && m.Duration != nil {
		return *m.Duration
	}
	return 0
}

func (m *Interval) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

type IteratorStats struct {
	SeriesN          *int64 `protobuf:"varint,1,opt,name=SeriesN" json:"SeriesN,omitempty"`
	PointN           *int64 `protobuf:"varint,2,opt,name=PointN" json:"PointN,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IteratorStats) Reset()                    { *m = IteratorStats{} }
func (m *IteratorStats) String() string            { return proto.CompactTextString(m) }
func (*IteratorStats) ProtoMessage()               {}
func (*IteratorStats) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{6} }

func (m *IteratorStats) GetSeriesN() int64 {
	if m != nil && m.SeriesN != nil {
		return *m.SeriesN
	}
	return 0
}

func (m *IteratorStats) GetPointN() int64 {
	if m != nil && m.PointN != nil {
		return *m.PointN
	}
	return 0
}

type VarRef struct {
	Val              *string `protobuf:"bytes,1,req,name=Val" json:"Val,omitempty"`
	Type             *int32  `protobuf:"varint,2,opt,name=Type" json:"Type,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *VarRef) Reset()                    { *m = VarRef{} }
func (m *VarRef) String() string            { return proto.CompactTextString(m) }
func (*VarRef) ProtoMessage()               {}
func (*VarRef) Descriptor() ([]byte, []int) { return fileDescriptorInternal, []int{7} }

func (m *VarRef) GetVal() string {
	if m != nil && m.Val != nil {
		return *m.Val
	}
	return ""
}

func (m *VarRef) GetType() int32 {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return 0
}

func init() {
	proto.RegisterType((*Point)(nil), "query.Point")
	proto.RegisterType((*Aux)(nil), "query.Aux")
	proto.RegisterType((*IteratorOptions)(nil), "query.IteratorOptions")
	proto.RegisterType((*Measurements)(nil), "query.Measurements")
	proto.RegisterType((*Measurement)(nil), "query.Measurement")
	proto.RegisterType((*Interval)(nil), "query.Interval")
	proto.RegisterType((*IteratorStats)(nil), "query.IteratorStats")
	proto.RegisterType((*VarRef)(nil), "query.VarRef")
}

func init() { proto.RegisterFile("internal/internal.proto", fileDescriptorInternal) }

var fileDescriptorInternal = []byte{
	// 769 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x55, 0xe1, 0x8e, 0xe3, 0x34,
	0x10, 0x56, 0x92, 0x4d, 0xb7, 0x71, 0xb7, 0xf4, 0x30, 0x65, 0xb1, 0xd0, 0x09, 0x45, 0x11, 0x48,
	0x11, 0xa0, 0x22, 0xed, 0x2f, 0x7e, 0x21, 0xf5, 0xd8, 0x5b, 0x54, 0xe9, 0xae, 0x7b, 0x72, 0x97,
	0xfd, 0x6f, 0x9a, 0xd9, 0xc8, 0x52, 0xea, 0x14, 0xc7, 0x41, 0xed, 0x03, 0xf0, 0x10, 0x3c, 0x16,
	0x4f, 0xc2, 0x2b, 0x20, 0x8f, 0x9d, 0x34, 0x5d, 0x81, 0xf6, 0x7e, 0x75, 0xbe, 0x6f, 0xa6, 0x63,
	0xcf, 0xcc, 0x37, 0x0e, 0xf9, 0x42, 0x2a, 0x03, 0x5a, 0x89, 0xea, 0x87, 0xce, 0x58, 0xec, 0x75,
	0x6d, 0x6a, 0x1a, 0xff, 0xde, 0x82, 0x3e, 0x66, 0xff, 0x84, 0x24, 0xfe, 0x50, 0x4b, 0x65, 0x28,
	0x25, 0x17, 0x6b, 0xb1, 0x03, 0x16, 0xa4, 0x61, 0x9e, 0x70, 0xb4, 0x2d, 0xf7, 0x20, 0xca, 0x86,
	0x85, 0x8e, 0xb3, 0x36, 0x72, 0x72, 0x07, 0x2c, 0x4a, 0xc3, 0x3c, 0xe2, 0x68, 0xd3, 0x57, 0x24,
	0x5a, 0xcb, 0x8a, 0x5d, 0xa4, 0x61, 0x3e, 0xe6, 0xd6, 0xa4, 0xaf, 0x49, 0xb4, 0x6c, 0x0f, 0x2c,
	0x4e, 0xa3, 0x7c, 0x72, 0x43, 0x16, 0x78, 0xd8, 0x62, 0xd9, 0x1e, 0xb8, 0xa5, 0xe9, 0x57, 0x84,
	0x2c, 0xcb, 0x52, 0x43, 0x29, 0x0c, 0x14, 0x6c, 0x94, 0x06, 0xf9, 0x94, 0x0f, 0x18, 0xeb, 0xbf,
	0xab, 0x6a, 0x61, 0x1e, 0x45, 0xd5, 0x02, 0xbb, 0x4c, 0x83, 0x3c, 0xe0, 0x03, 0x86, 0x66, 0xe4,
	0x6a, 0xa5, 0x0c, 0x94, 0xa0, 0x5d, 0xc4, 0x38, 0x0d, 0xf2, 0x88, 0x9f, 0x71, 0x34, 0x25, 0x93,
	0x8d, 0xd1, 0x52, 0x95, 0x2e, 0x24, 0x49, 0x83, 0x3c, 0xe1, 0x43, 0xca, 0x66, 0x79, 0x53, 0xd7,
	0x15, 0x08, 0xe5, 0x42, 0x48, 0x1a, 0xe4, 0x63, 0x7e, 0xc6, 0xd1, 0xaf, 0xc9, 0xf4, 0x57, 0xd5,
	0xc8, 0x52, 0x41, 0xe1, 0x82, 0xae, 0xd2, 0x20, 0xbf, 0xe0, 0xe7, 0x24, 0xfd, 0x96, 0xc4, 0x1b,
	0x23, 0x4c, 0xc3, 0x26, 0x69, 0x90, 0x4f, 0x6e, 0xe6, 0xbe, 0xde, 0x95, 0x01, 0x2d, 0x4c, 0xad,
	0xd1, 0xc7, 0x5d, 0x48, 0xf6, 0x77, 0x80, 0xad, 0xa1, 0x5f, 0x92, 0xf1, 0xad, 0x30, 0xe2, 0xe1,
	0xb8, 0x77, 0x3d, 0x8f, 0x79, 0x8f, 0x9f, 0xd5, 0x1f, 0xbe, 0x58, 0x7f, 0xf4, 0x72, 0xfd, 0x17,
	0x2f, 0xd7, 0x1f, 0x7f, 0x4c, 0xfd, 0xa3, 0xff, 0xa8, 0x3f, 0xfb, 0x33, 0x26, 0xb3, 0xae, 0xd8,
	0xfb, 0xbd, 0x91, 0xb5, 0x42, 0x9d, 0xbc, 0x3d, 0xec, 0x35, 0x0b, 0xf0, 0x60, 0xb4, 0xad, 0x4e,
	0xac, 0x2a, 0xc2, 0x34, 0xca, 0x13, 0xa7, 0x84, 0x6f, 0xc8, 0xe8, 0x4e, 0x42, 0x55, 0x34, 0xec,
	0x53, 0x94, 0xca, 0xd4, 0xb7, 0xee, 0x51, 0x68, 0x0e, 0x4f, 0xdc, 0x3b, 0xe9, 0xf7, 0xe4, 0x72,
	0x53, 0xb7, 0x7a, 0x0b, 0x0d, 0x8b, 0x30, 0x8e, 0xfa, 0xb8, 0xf7, 0x20, 0x9a, 0x56, 0xc3, 0x0e,
	0x94, 0xe1, 0x5d, 0x08, 0xfd, 0x8e, 0x8c, 0x6d, 0x2b, 0xf4, 0x1f, 0xa2, 0xc2, 0xba, 0x27, 0x37,
	0xb3, 0x6e, 0x22, 0x9e, 0xe6, 0x7d, 0x80, 0xed, 0xf5, 0xad, 0xdc, 0x81, 0x6a, 0xec, 0xad, 0x51,
	0xb0, 0x09, 0x1f, 0x30, 0x94, 0x91, 0xcb, 0x5f, 0x74, 0xdd, 0xee, 0xdf, 0x1c, 0xd9, 0x67, 0xe8,
	0xec, 0xa0, 0xad, 0xf0, 0x4e, 0x56, 0x15, 0xb6, 0x24, 0xe6, 0x68, 0xd3, 0xd7, 0x24, 0xb1, 0xbf,
	0x43, 0xe1, 0x9e, 0x08, 0xeb, 0xfd, 0xb9, 0x56, 0x85, 0xb4, 0x1d, 0x42, 0xd1, 0x26, 0xfc, 0x44,
	0x58, 0xef, 0xc6, 0x08, 0x6d, 0x70, 0xbd, 0x12, 0x1c, 0xe9, 0x89, 0xb0, 0xf7, 0x78, 0xab, 0x0a,
	0xf4, 0x11, 0xf4, 0x75, 0xd0, 0x2a, 0xe9, 0x5d, 0xbd, 0x15, 0x98, 0xf4, 0x73, 0x4c, 0xda, 0x63,
	0x9b, 0x73, 0xd9, 0x6c, 0x41, 0x15, 0x52, 0x95, 0xa8, 0xce, 0x31, 0x3f, 0x11, 0x74, 0x4e, 0xe2,
	0x77, 0x72, 0x27, 0x0d, 0xaa, 0x3a, 0xe2, 0x0e, 0xd0, 0x6b, 0x32, 0xba, 0x7f, 0x7a, 0x6a, 0xc0,
	0xb0, 0x29, 0xd2, 0x1e, 0x59, 0x7e, 0xe3, 0xc2, 0x3f, 0x71, 0xbc, 0x43, 0xf6, 0x66, 0x1b, 0xff,
	0x87, 0x99, 0xbb, 0x99, 0x87, 0xae, 0x22, 0x2d, 0xf7, 0xf8, 0xb0, 0x5c, 0xbb, 0xd3, 0x7b, 0xc2,
	0xe6, 0xbb, 0x85, 0xa2, 0xdd, 0x03, 0x7b, 0x85, 0x2e, 0x8f, 0xec, 0x44, 0xde, 0x8b, 0xc3, 0x06,
	0xb4, 0x84, 0x66, 0xcd, 0x28, 0xa6, 0x1c, 0x30, 0xf6, 0xbc, 0x7b, 0x5d, 0x80, 0x86, 0x82, 0xcd,
	0xf1, 0x8f, 0x1d, 0xcc, 0x7e, 0x24, 0x57, 0x03, 0x41, 0x34, 0x34, 0x27, 0xf1, 0xca, 0xc0, 0xae,
	0x61, 0xc1, 0xff, 0x8a, 0xc6, 0x05, 0x64, 0x7f, 0x05, 0x64, 0x32, 0xa0, 0xbb, 0xed, 0xfc, 0x4d,
	0x34, 0xe0, 0x15, 0xdc, 0x63, 0x9a, 0x93, 0x19, 0x07, 0x03, 0xca, 0x36, 0xf8, 0x43, 0x5d, 0xc9,
	0xed, 0x11, 0x57, 0x34, 0xe1, 0xcf, 0xe9, 0xfe, 0x4d, 0x8d, 0xdc, 0x0e, 0x60, 0xd5, 0x73, 0x12,
	0x73, 0x28, 0xe1, 0xe0, 0x37, 0xd2, 0x01, 0x7b, 0xde, 0xaa, 0x79, 0x10, 0xba, 0x04, 0xe3, 0xf7,
	0xb0, 0xc7, 0xd9, 0x4f, 0x27, 0x39, 0xe3, 0xbd, 0x5a, 0xed, 0x66, 0x1d, 0x60, 0x67, 0x7a, 0x3c,
	0x98, 0x5b, 0x38, 0x9c, 0x5b, 0xb6, 0x24, 0xd3, 0xb3, 0x97, 0x08, 0x07, 0xe6, 0xbb, 0x1b, 0xf8,
	0x81, 0xf9, 0xd6, 0x5e, 0x93, 0x11, 0x7e, 0x0d, 0xd6, 0x5d, 0x0a, 0x87, 0xb2, 0x05, 0x19, 0xb9,
	0x8d, 0xb4, 0x2b, 0xfc, 0x28, 0x2a, 0xff, 0x95, 0xb0, 0x26, 0x7e, 0x10, 0xec, 0x23, 0x16, 0xba,
	0x35, 0xb0, 0xf6, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x55, 0x55, 0xec, 0xe3, 0x77, 0x06, 0x00,
	0x00,
}
