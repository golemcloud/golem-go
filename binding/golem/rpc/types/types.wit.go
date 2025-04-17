// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "golem:rpc/types@0.2.0".
package types

import (
	wallclock "github.com/golemcloud/golem-go/binding/wasi/clocks/wall-clock"
	"github.com/golemcloud/golem-go/binding/wasi/io/poll"
	"go.bytecodealliance.org/cm"
)

// DateTime represents the type alias "golem:rpc/types@0.2.0#datetime".
//
// See [wallclock.DateTime] for more information.
type DateTime = wallclock.DateTime

// Pollable represents the imported type alias "golem:rpc/types@0.2.0#pollable".
//
// See [poll.Pollable] for more information.
type Pollable = poll.Pollable

// UUID represents the record "golem:rpc/types@0.2.0#uuid".
//
// UUID
//
//	record uuid {
//		high-bits: u64,
//		low-bits: u64,
//	}
type UUID struct {
	_        cm.HostLayout `json:"-"`
	HighBits uint64        `json:"high-bits"`
	LowBits  uint64        `json:"low-bits"`
}

// ComponentID represents the record "golem:rpc/types@0.2.0#component-id".
//
// Represents a Golem component
//
//	record component-id {
//		uuid: uuid,
//	}
type ComponentID struct {
	_    cm.HostLayout `json:"-"`
	UUID UUID          `json:"uuid"`
}

// WorkerID represents the record "golem:rpc/types@0.2.0#worker-id".
//
// Represents a Golem worker
//
//	record worker-id {
//		component-id: component-id,
//		worker-name: string,
//	}
type WorkerID struct {
	_           cm.HostLayout `json:"-"`
	ComponentID ComponentID   `json:"component-id"`
	WorkerName  string        `json:"worker-name"`
}

// NodeIndex represents the s32 "golem:rpc/types@0.2.0#node-index".
//
//	type node-index = s32
type NodeIndex int32

// ResourceID represents the u64 "golem:rpc/types@0.2.0#resource-id".
//
//	type resource-id = u64
type ResourceID uint64

// ResourceMode represents the enum "golem:rpc/types@0.2.0#resource-mode".
//
//	enum resource-mode {
//		owned,
//		borrowed
//	}
type ResourceMode uint8

const (
	ResourceModeOwned ResourceMode = iota
	ResourceModeBorrowed
)

var _ResourceModeStrings = [2]string{
	"owned",
	"borrowed",
}

// String implements [fmt.Stringer], returning the enum case name of e.
func (e ResourceMode) String() string {
	return _ResourceModeStrings[e]
}

// MarshalText implements [encoding.TextMarshaler].
func (e ResourceMode) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalText implements [encoding.TextUnmarshaler], unmarshaling into an enum
// case. Returns an error if the supplied text is not one of the enum cases.
func (e *ResourceMode) UnmarshalText(text []byte) error {
	return _ResourceModeUnmarshalCase(e, text)
}

var _ResourceModeUnmarshalCase = cm.CaseUnmarshaler[ResourceMode](_ResourceModeStrings[:])

// WitTypeNode represents the variant "golem:rpc/types@0.2.0#wit-type-node".
//
//	variant wit-type-node {
//		record-type(list<tuple<string, node-index>>),
//		variant-type(list<tuple<string, option<node-index>>>),
//		enum-type(list<string>),
//		flags-type(list<string>),
//		tuple-type(list<node-index>),
//		list-type(node-index),
//		option-type(node-index),
//		result-type(tuple<option<node-index>, option<node-index>>),
//		prim-u8-type,
//		prim-u16-type,
//		prim-u32-type,
//		prim-u64-type,
//		prim-s8-type,
//		prim-s16-type,
//		prim-s32-type,
//		prim-s64-type,
//		prim-f32-type,
//		prim-f64-type,
//		prim-char-type,
//		prim-bool-type,
//		prim-string-type,
//		handle-type(tuple<resource-id, resource-mode>),
//	}
type WitTypeNode cm.Variant[uint8, [2]cm.Option[NodeIndex], cm.Tuple[ResourceID, ResourceMode]]

// WitTypeNodeRecordType returns a [WitTypeNode] of case "record-type".
func WitTypeNodeRecordType(data cm.List[cm.Tuple[string, NodeIndex]]) WitTypeNode {
	return cm.New[WitTypeNode](0, data)
}

// RecordType returns a non-nil *[cm.List[cm.Tuple[string, NodeIndex]]] if [WitTypeNode] represents the variant case "record-type".
func (self *WitTypeNode) RecordType() *cm.List[cm.Tuple[string, NodeIndex]] {
	return cm.Case[cm.List[cm.Tuple[string, NodeIndex]]](self, 0)
}

// WitTypeNodeVariantType returns a [WitTypeNode] of case "variant-type".
func WitTypeNodeVariantType(data cm.List[cm.Tuple[string, cm.Option[NodeIndex]]]) WitTypeNode {
	return cm.New[WitTypeNode](1, data)
}

// VariantType returns a non-nil *[cm.List[cm.Tuple[string, cm.Option[NodeIndex]]]] if [WitTypeNode] represents the variant case "variant-type".
func (self *WitTypeNode) VariantType() *cm.List[cm.Tuple[string, cm.Option[NodeIndex]]] {
	return cm.Case[cm.List[cm.Tuple[string, cm.Option[NodeIndex]]]](self, 1)
}

// WitTypeNodeEnumType returns a [WitTypeNode] of case "enum-type".
func WitTypeNodeEnumType(data cm.List[string]) WitTypeNode {
	return cm.New[WitTypeNode](2, data)
}

// EnumType returns a non-nil *[cm.List[string]] if [WitTypeNode] represents the variant case "enum-type".
func (self *WitTypeNode) EnumType() *cm.List[string] {
	return cm.Case[cm.List[string]](self, 2)
}

// WitTypeNodeFlagsType returns a [WitTypeNode] of case "flags-type".
func WitTypeNodeFlagsType(data cm.List[string]) WitTypeNode {
	return cm.New[WitTypeNode](3, data)
}

// FlagsType returns a non-nil *[cm.List[string]] if [WitTypeNode] represents the variant case "flags-type".
func (self *WitTypeNode) FlagsType() *cm.List[string] {
	return cm.Case[cm.List[string]](self, 3)
}

// WitTypeNodeTupleType returns a [WitTypeNode] of case "tuple-type".
func WitTypeNodeTupleType(data cm.List[NodeIndex]) WitTypeNode {
	return cm.New[WitTypeNode](4, data)
}

// TupleType returns a non-nil *[cm.List[NodeIndex]] if [WitTypeNode] represents the variant case "tuple-type".
func (self *WitTypeNode) TupleType() *cm.List[NodeIndex] {
	return cm.Case[cm.List[NodeIndex]](self, 4)
}

// WitTypeNodeListType returns a [WitTypeNode] of case "list-type".
func WitTypeNodeListType(data NodeIndex) WitTypeNode {
	return cm.New[WitTypeNode](5, data)
}

// ListType returns a non-nil *[NodeIndex] if [WitTypeNode] represents the variant case "list-type".
func (self *WitTypeNode) ListType() *NodeIndex {
	return cm.Case[NodeIndex](self, 5)
}

// WitTypeNodeOptionType returns a [WitTypeNode] of case "option-type".
func WitTypeNodeOptionType(data NodeIndex) WitTypeNode {
	return cm.New[WitTypeNode](6, data)
}

// OptionType returns a non-nil *[NodeIndex] if [WitTypeNode] represents the variant case "option-type".
func (self *WitTypeNode) OptionType() *NodeIndex {
	return cm.Case[NodeIndex](self, 6)
}

// WitTypeNodeResultType returns a [WitTypeNode] of case "result-type".
func WitTypeNodeResultType(data [2]cm.Option[NodeIndex]) WitTypeNode {
	return cm.New[WitTypeNode](7, data)
}

// ResultType returns a non-nil *[[2]cm.Option[NodeIndex]] if [WitTypeNode] represents the variant case "result-type".
func (self *WitTypeNode) ResultType() *[2]cm.Option[NodeIndex] {
	return cm.Case[[2]cm.Option[NodeIndex]](self, 7)
}

// WitTypeNodePrimU8Type returns a [WitTypeNode] of case "prim-u8-type".
func WitTypeNodePrimU8Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](8, data)
}

// PrimU8Type returns true if [WitTypeNode] represents the variant case "prim-u8-type".
func (self *WitTypeNode) PrimU8Type() bool {
	return self.Tag() == 8
}

// WitTypeNodePrimU16Type returns a [WitTypeNode] of case "prim-u16-type".
func WitTypeNodePrimU16Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](9, data)
}

// PrimU16Type returns true if [WitTypeNode] represents the variant case "prim-u16-type".
func (self *WitTypeNode) PrimU16Type() bool {
	return self.Tag() == 9
}

// WitTypeNodePrimU32Type returns a [WitTypeNode] of case "prim-u32-type".
func WitTypeNodePrimU32Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](10, data)
}

// PrimU32Type returns true if [WitTypeNode] represents the variant case "prim-u32-type".
func (self *WitTypeNode) PrimU32Type() bool {
	return self.Tag() == 10
}

// WitTypeNodePrimU64Type returns a [WitTypeNode] of case "prim-u64-type".
func WitTypeNodePrimU64Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](11, data)
}

// PrimU64Type returns true if [WitTypeNode] represents the variant case "prim-u64-type".
func (self *WitTypeNode) PrimU64Type() bool {
	return self.Tag() == 11
}

// WitTypeNodePrimS8Type returns a [WitTypeNode] of case "prim-s8-type".
func WitTypeNodePrimS8Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](12, data)
}

// PrimS8Type returns true if [WitTypeNode] represents the variant case "prim-s8-type".
func (self *WitTypeNode) PrimS8Type() bool {
	return self.Tag() == 12
}

// WitTypeNodePrimS16Type returns a [WitTypeNode] of case "prim-s16-type".
func WitTypeNodePrimS16Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](13, data)
}

// PrimS16Type returns true if [WitTypeNode] represents the variant case "prim-s16-type".
func (self *WitTypeNode) PrimS16Type() bool {
	return self.Tag() == 13
}

// WitTypeNodePrimS32Type returns a [WitTypeNode] of case "prim-s32-type".
func WitTypeNodePrimS32Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](14, data)
}

// PrimS32Type returns true if [WitTypeNode] represents the variant case "prim-s32-type".
func (self *WitTypeNode) PrimS32Type() bool {
	return self.Tag() == 14
}

// WitTypeNodePrimS64Type returns a [WitTypeNode] of case "prim-s64-type".
func WitTypeNodePrimS64Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](15, data)
}

// PrimS64Type returns true if [WitTypeNode] represents the variant case "prim-s64-type".
func (self *WitTypeNode) PrimS64Type() bool {
	return self.Tag() == 15
}

// WitTypeNodePrimF32Type returns a [WitTypeNode] of case "prim-f32-type".
func WitTypeNodePrimF32Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](16, data)
}

// PrimF32Type returns true if [WitTypeNode] represents the variant case "prim-f32-type".
func (self *WitTypeNode) PrimF32Type() bool {
	return self.Tag() == 16
}

// WitTypeNodePrimF64Type returns a [WitTypeNode] of case "prim-f64-type".
func WitTypeNodePrimF64Type() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](17, data)
}

// PrimF64Type returns true if [WitTypeNode] represents the variant case "prim-f64-type".
func (self *WitTypeNode) PrimF64Type() bool {
	return self.Tag() == 17
}

// WitTypeNodePrimCharType returns a [WitTypeNode] of case "prim-char-type".
func WitTypeNodePrimCharType() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](18, data)
}

// PrimCharType returns true if [WitTypeNode] represents the variant case "prim-char-type".
func (self *WitTypeNode) PrimCharType() bool {
	return self.Tag() == 18
}

// WitTypeNodePrimBoolType returns a [WitTypeNode] of case "prim-bool-type".
func WitTypeNodePrimBoolType() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](19, data)
}

// PrimBoolType returns true if [WitTypeNode] represents the variant case "prim-bool-type".
func (self *WitTypeNode) PrimBoolType() bool {
	return self.Tag() == 19
}

// WitTypeNodePrimStringType returns a [WitTypeNode] of case "prim-string-type".
func WitTypeNodePrimStringType() WitTypeNode {
	var data struct{}
	return cm.New[WitTypeNode](20, data)
}

// PrimStringType returns true if [WitTypeNode] represents the variant case "prim-string-type".
func (self *WitTypeNode) PrimStringType() bool {
	return self.Tag() == 20
}

// WitTypeNodeHandleType returns a [WitTypeNode] of case "handle-type".
func WitTypeNodeHandleType(data cm.Tuple[ResourceID, ResourceMode]) WitTypeNode {
	return cm.New[WitTypeNode](21, data)
}

// HandleType returns a non-nil *[cm.Tuple[ResourceID, ResourceMode]] if [WitTypeNode] represents the variant case "handle-type".
func (self *WitTypeNode) HandleType() *cm.Tuple[ResourceID, ResourceMode] {
	return cm.Case[cm.Tuple[ResourceID, ResourceMode]](self, 21)
}

var _WitTypeNodeStrings = [22]string{
	"record-type",
	"variant-type",
	"enum-type",
	"flags-type",
	"tuple-type",
	"list-type",
	"option-type",
	"result-type",
	"prim-u8-type",
	"prim-u16-type",
	"prim-u32-type",
	"prim-u64-type",
	"prim-s8-type",
	"prim-s16-type",
	"prim-s32-type",
	"prim-s64-type",
	"prim-f32-type",
	"prim-f64-type",
	"prim-char-type",
	"prim-bool-type",
	"prim-string-type",
	"handle-type",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v WitTypeNode) String() string {
	return _WitTypeNodeStrings[v.Tag()]
}

// WitType represents the record "golem:rpc/types@0.2.0#wit-type".
//
//	record wit-type {
//		nodes: list<wit-type-node>,
//	}
type WitType struct {
	_     cm.HostLayout        `json:"-"`
	Nodes cm.List[WitTypeNode] `json:"nodes"`
}

// URI represents the record "golem:rpc/types@0.2.0#uri".
//
//	record uri {
//		value: string,
//	}
type URI struct {
	_     cm.HostLayout `json:"-"`
	Value string        `json:"value"`
}

// WitNode represents the variant "golem:rpc/types@0.2.0#wit-node".
//
//	variant wit-node {
//		record-value(list<node-index>),
//		variant-value(tuple<u32, option<node-index>>),
//		enum-value(u32),
//		flags-value(list<bool>),
//		tuple-value(list<node-index>),
//		list-value(list<node-index>),
//		option-value(option<node-index>),
//		result-value(result<option<node-index>, option<node-index>>),
//		prim-u8(u8),
//		prim-u16(u16),
//		prim-u32(u32),
//		prim-u64(u64),
//		prim-s8(s8),
//		prim-s16(s16),
//		prim-s32(s32),
//		prim-s64(s64),
//		prim-float32(f32),
//		prim-float64(f64),
//		prim-char(char),
//		prim-bool(bool),
//		prim-string(string),
//		handle(tuple<uri, u64>),
//	}
type WitNode cm.Variant[uint8, TupleURIU64Shape, cm.Tuple[URI, uint64]]

// WitNodeRecordValue returns a [WitNode] of case "record-value".
func WitNodeRecordValue(data cm.List[NodeIndex]) WitNode {
	return cm.New[WitNode](0, data)
}

// RecordValue returns a non-nil *[cm.List[NodeIndex]] if [WitNode] represents the variant case "record-value".
func (self *WitNode) RecordValue() *cm.List[NodeIndex] {
	return cm.Case[cm.List[NodeIndex]](self, 0)
}

// WitNodeVariantValue returns a [WitNode] of case "variant-value".
func WitNodeVariantValue(data cm.Tuple[uint32, cm.Option[NodeIndex]]) WitNode {
	return cm.New[WitNode](1, data)
}

// VariantValue returns a non-nil *[cm.Tuple[uint32, cm.Option[NodeIndex]]] if [WitNode] represents the variant case "variant-value".
func (self *WitNode) VariantValue() *cm.Tuple[uint32, cm.Option[NodeIndex]] {
	return cm.Case[cm.Tuple[uint32, cm.Option[NodeIndex]]](self, 1)
}

// WitNodeEnumValue returns a [WitNode] of case "enum-value".
func WitNodeEnumValue(data uint32) WitNode {
	return cm.New[WitNode](2, data)
}

// EnumValue returns a non-nil *[uint32] if [WitNode] represents the variant case "enum-value".
func (self *WitNode) EnumValue() *uint32 {
	return cm.Case[uint32](self, 2)
}

// WitNodeFlagsValue returns a [WitNode] of case "flags-value".
func WitNodeFlagsValue(data cm.List[bool]) WitNode {
	return cm.New[WitNode](3, data)
}

// FlagsValue returns a non-nil *[cm.List[bool]] if [WitNode] represents the variant case "flags-value".
func (self *WitNode) FlagsValue() *cm.List[bool] {
	return cm.Case[cm.List[bool]](self, 3)
}

// WitNodeTupleValue returns a [WitNode] of case "tuple-value".
func WitNodeTupleValue(data cm.List[NodeIndex]) WitNode {
	return cm.New[WitNode](4, data)
}

// TupleValue returns a non-nil *[cm.List[NodeIndex]] if [WitNode] represents the variant case "tuple-value".
func (self *WitNode) TupleValue() *cm.List[NodeIndex] {
	return cm.Case[cm.List[NodeIndex]](self, 4)
}

// WitNodeListValue returns a [WitNode] of case "list-value".
func WitNodeListValue(data cm.List[NodeIndex]) WitNode {
	return cm.New[WitNode](5, data)
}

// ListValue returns a non-nil *[cm.List[NodeIndex]] if [WitNode] represents the variant case "list-value".
func (self *WitNode) ListValue() *cm.List[NodeIndex] {
	return cm.Case[cm.List[NodeIndex]](self, 5)
}

// WitNodeOptionValue returns a [WitNode] of case "option-value".
func WitNodeOptionValue(data cm.Option[NodeIndex]) WitNode {
	return cm.New[WitNode](6, data)
}

// OptionValue returns a non-nil *[cm.Option[NodeIndex]] if [WitNode] represents the variant case "option-value".
func (self *WitNode) OptionValue() *cm.Option[NodeIndex] {
	return cm.Case[cm.Option[NodeIndex]](self, 6)
}

// WitNodeResultValue returns a [WitNode] of case "result-value".
func WitNodeResultValue(data cm.Result[cm.Option[NodeIndex], cm.Option[NodeIndex], cm.Option[NodeIndex]]) WitNode {
	return cm.New[WitNode](7, data)
}

// ResultValue returns a non-nil *[cm.Result[cm.Option[NodeIndex], cm.Option[NodeIndex], cm.Option[NodeIndex]]] if [WitNode] represents the variant case "result-value".
func (self *WitNode) ResultValue() *cm.Result[cm.Option[NodeIndex], cm.Option[NodeIndex], cm.Option[NodeIndex]] {
	return cm.Case[cm.Result[cm.Option[NodeIndex], cm.Option[NodeIndex], cm.Option[NodeIndex]]](self, 7)
}

// WitNodePrimU8 returns a [WitNode] of case "prim-u8".
func WitNodePrimU8(data uint8) WitNode {
	return cm.New[WitNode](8, data)
}

// PrimU8 returns a non-nil *[uint8] if [WitNode] represents the variant case "prim-u8".
func (self *WitNode) PrimU8() *uint8 {
	return cm.Case[uint8](self, 8)
}

// WitNodePrimU16 returns a [WitNode] of case "prim-u16".
func WitNodePrimU16(data uint16) WitNode {
	return cm.New[WitNode](9, data)
}

// PrimU16 returns a non-nil *[uint16] if [WitNode] represents the variant case "prim-u16".
func (self *WitNode) PrimU16() *uint16 {
	return cm.Case[uint16](self, 9)
}

// WitNodePrimU32 returns a [WitNode] of case "prim-u32".
func WitNodePrimU32(data uint32) WitNode {
	return cm.New[WitNode](10, data)
}

// PrimU32 returns a non-nil *[uint32] if [WitNode] represents the variant case "prim-u32".
func (self *WitNode) PrimU32() *uint32 {
	return cm.Case[uint32](self, 10)
}

// WitNodePrimU64 returns a [WitNode] of case "prim-u64".
func WitNodePrimU64(data uint64) WitNode {
	return cm.New[WitNode](11, data)
}

// PrimU64 returns a non-nil *[uint64] if [WitNode] represents the variant case "prim-u64".
func (self *WitNode) PrimU64() *uint64 {
	return cm.Case[uint64](self, 11)
}

// WitNodePrimS8 returns a [WitNode] of case "prim-s8".
func WitNodePrimS8(data int8) WitNode {
	return cm.New[WitNode](12, data)
}

// PrimS8 returns a non-nil *[int8] if [WitNode] represents the variant case "prim-s8".
func (self *WitNode) PrimS8() *int8 {
	return cm.Case[int8](self, 12)
}

// WitNodePrimS16 returns a [WitNode] of case "prim-s16".
func WitNodePrimS16(data int16) WitNode {
	return cm.New[WitNode](13, data)
}

// PrimS16 returns a non-nil *[int16] if [WitNode] represents the variant case "prim-s16".
func (self *WitNode) PrimS16() *int16 {
	return cm.Case[int16](self, 13)
}

// WitNodePrimS32 returns a [WitNode] of case "prim-s32".
func WitNodePrimS32(data int32) WitNode {
	return cm.New[WitNode](14, data)
}

// PrimS32 returns a non-nil *[int32] if [WitNode] represents the variant case "prim-s32".
func (self *WitNode) PrimS32() *int32 {
	return cm.Case[int32](self, 14)
}

// WitNodePrimS64 returns a [WitNode] of case "prim-s64".
func WitNodePrimS64(data int64) WitNode {
	return cm.New[WitNode](15, data)
}

// PrimS64 returns a non-nil *[int64] if [WitNode] represents the variant case "prim-s64".
func (self *WitNode) PrimS64() *int64 {
	return cm.Case[int64](self, 15)
}

// WitNodePrimFloat32 returns a [WitNode] of case "prim-float32".
func WitNodePrimFloat32(data float32) WitNode {
	return cm.New[WitNode](16, data)
}

// PrimFloat32 returns a non-nil *[float32] if [WitNode] represents the variant case "prim-float32".
func (self *WitNode) PrimFloat32() *float32 {
	return cm.Case[float32](self, 16)
}

// WitNodePrimFloat64 returns a [WitNode] of case "prim-float64".
func WitNodePrimFloat64(data float64) WitNode {
	return cm.New[WitNode](17, data)
}

// PrimFloat64 returns a non-nil *[float64] if [WitNode] represents the variant case "prim-float64".
func (self *WitNode) PrimFloat64() *float64 {
	return cm.Case[float64](self, 17)
}

// WitNodePrimChar returns a [WitNode] of case "prim-char".
func WitNodePrimChar(data rune) WitNode {
	return cm.New[WitNode](18, data)
}

// PrimChar returns a non-nil *[rune] if [WitNode] represents the variant case "prim-char".
func (self *WitNode) PrimChar() *rune {
	return cm.Case[rune](self, 18)
}

// WitNodePrimBool returns a [WitNode] of case "prim-bool".
func WitNodePrimBool(data bool) WitNode {
	return cm.New[WitNode](19, data)
}

// PrimBool returns a non-nil *[bool] if [WitNode] represents the variant case "prim-bool".
func (self *WitNode) PrimBool() *bool {
	return cm.Case[bool](self, 19)
}

// WitNodePrimString returns a [WitNode] of case "prim-string".
func WitNodePrimString(data string) WitNode {
	return cm.New[WitNode](20, data)
}

// PrimString returns a non-nil *[string] if [WitNode] represents the variant case "prim-string".
func (self *WitNode) PrimString() *string {
	return cm.Case[string](self, 20)
}

// WitNodeHandle returns a [WitNode] of case "handle".
func WitNodeHandle(data cm.Tuple[URI, uint64]) WitNode {
	return cm.New[WitNode](21, data)
}

// Handle returns a non-nil *[cm.Tuple[URI, uint64]] if [WitNode] represents the variant case "handle".
func (self *WitNode) Handle() *cm.Tuple[URI, uint64] {
	return cm.Case[cm.Tuple[URI, uint64]](self, 21)
}

var _WitNodeStrings = [22]string{
	"record-value",
	"variant-value",
	"enum-value",
	"flags-value",
	"tuple-value",
	"list-value",
	"option-value",
	"result-value",
	"prim-u8",
	"prim-u16",
	"prim-u32",
	"prim-u64",
	"prim-s8",
	"prim-s16",
	"prim-s32",
	"prim-s64",
	"prim-float32",
	"prim-float64",
	"prim-char",
	"prim-bool",
	"prim-string",
	"handle",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v WitNode) String() string {
	return _WitNodeStrings[v.Tag()]
}

// WitValue represents the record "golem:rpc/types@0.2.0#wit-value".
//
//	record wit-value {
//		nodes: list<wit-node>,
//	}
type WitValue struct {
	_     cm.HostLayout    `json:"-"`
	Nodes cm.List[WitNode] `json:"nodes"`
}

// ValueAndType represents the record "golem:rpc/types@0.2.0#value-and-type".
//
//	record value-and-type {
//		value: wit-value,
//		typ: wit-type,
//	}
type ValueAndType struct {
	_     cm.HostLayout `json:"-"`
	Value WitValue      `json:"value"`
	Typ   WitType       `json:"typ"`
}

// RPCError represents the variant "golem:rpc/types@0.2.0#rpc-error".
//
//	variant rpc-error {
//		protocol-error(string),
//		denied(string),
//		not-found(string),
//		remote-internal-error(string),
//	}
type RPCError cm.Variant[uint8, string, string]

// RPCErrorProtocolError returns a [RPCError] of case "protocol-error".
func RPCErrorProtocolError(data string) RPCError {
	return cm.New[RPCError](0, data)
}

// ProtocolError returns a non-nil *[string] if [RPCError] represents the variant case "protocol-error".
func (self *RPCError) ProtocolError() *string {
	return cm.Case[string](self, 0)
}

// RPCErrorDenied returns a [RPCError] of case "denied".
func RPCErrorDenied(data string) RPCError {
	return cm.New[RPCError](1, data)
}

// Denied returns a non-nil *[string] if [RPCError] represents the variant case "denied".
func (self *RPCError) Denied() *string {
	return cm.Case[string](self, 1)
}

// RPCErrorNotFound returns a [RPCError] of case "not-found".
func RPCErrorNotFound(data string) RPCError {
	return cm.New[RPCError](2, data)
}

// NotFound returns a non-nil *[string] if [RPCError] represents the variant case "not-found".
func (self *RPCError) NotFound() *string {
	return cm.Case[string](self, 2)
}

// RPCErrorRemoteInternalError returns a [RPCError] of case "remote-internal-error".
func RPCErrorRemoteInternalError(data string) RPCError {
	return cm.New[RPCError](3, data)
}

// RemoteInternalError returns a non-nil *[string] if [RPCError] represents the variant case "remote-internal-error".
func (self *RPCError) RemoteInternalError() *string {
	return cm.Case[string](self, 3)
}

var _RPCErrorStrings = [4]string{
	"protocol-error",
	"denied",
	"not-found",
	"remote-internal-error",
}

// String implements [fmt.Stringer], returning the variant case name of v.
func (v RPCError) String() string {
	return _RPCErrorStrings[v.Tag()]
}

// WasmRPC represents the imported resource "golem:rpc/types@0.2.0#wasm-rpc".
//
//	resource wasm-rpc
type WasmRPC cm.Resource

// ResourceDrop represents the imported resource-drop for resource "wasm-rpc".
//
// Drops a resource handle.
//
//go:nosplit
func (self WasmRPC) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_WasmRPCResourceDrop((uint32)(self0))
	return
}

// NewWasmRPC represents the imported constructor for resource "wasm-rpc".
//
//	constructor(worker-id: worker-id)
//
//go:nosplit
func NewWasmRPC(workerID WorkerID) (result WasmRPC) {
	workerId0, workerId1, workerId2, workerId3 := lower_WorkerID(workerID)
	result0 := wasmimport_NewWasmRPC((uint64)(workerId0), (uint64)(workerId1), (*uint8)(workerId2), (uint32)(workerId3))
	result = cm.Reinterpret[WasmRPC]((uint32)(result0))
	return
}

// WasmRPCEphemeral represents the imported static function "ephemeral".
//
//	ephemeral: static func(component-id: component-id) -> wasm-rpc
//
//go:nosplit
func WasmRPCEphemeral(componentID ComponentID) (result WasmRPC) {
	componentId0, componentId1 := lower_ComponentID(componentID)
	result0 := wasmimport_WasmRPCEphemeral((uint64)(componentId0), (uint64)(componentId1))
	result = cm.Reinterpret[WasmRPC]((uint32)(result0))
	return
}

// AsyncInvokeAndAwait represents the imported method "async-invoke-and-await".
//
//	async-invoke-and-await: func(function-name: string, function-params: list<wit-value>)
//	-> future-invoke-result
//
//go:nosplit
func (self WasmRPC) AsyncInvokeAndAwait(functionName string, functionParams cm.List[WitValue]) (result FutureInvokeResult) {
	self0 := cm.Reinterpret[uint32](self)
	functionName0, functionName1 := cm.LowerString(functionName)
	functionParams0, functionParams1 := cm.LowerList(functionParams)
	result0 := wasmimport_WasmRPCAsyncInvokeAndAwait((uint32)(self0), (*uint8)(functionName0), (uint32)(functionName1), (*WitValue)(functionParams0), (uint32)(functionParams1))
	result = cm.Reinterpret[FutureInvokeResult]((uint32)(result0))
	return
}

// Invoke represents the imported method "invoke".
//
//	invoke: func(function-name: string, function-params: list<wit-value>) -> result<_,
//	rpc-error>
//
//go:nosplit
func (self WasmRPC) Invoke(functionName string, functionParams cm.List[WitValue]) (result cm.Result[RPCError, struct{}, RPCError]) {
	self0 := cm.Reinterpret[uint32](self)
	functionName0, functionName1 := cm.LowerString(functionName)
	functionParams0, functionParams1 := cm.LowerList(functionParams)
	wasmimport_WasmRPCInvoke((uint32)(self0), (*uint8)(functionName0), (uint32)(functionName1), (*WitValue)(functionParams0), (uint32)(functionParams1), &result)
	return
}

// InvokeAndAwait represents the imported method "invoke-and-await".
//
//	invoke-and-await: func(function-name: string, function-params: list<wit-value>)
//	-> result<wit-value, rpc-error>
//
//go:nosplit
func (self WasmRPC) InvokeAndAwait(functionName string, functionParams cm.List[WitValue]) (result cm.Result[RPCErrorShape, WitValue, RPCError]) {
	self0 := cm.Reinterpret[uint32](self)
	functionName0, functionName1 := cm.LowerString(functionName)
	functionParams0, functionParams1 := cm.LowerList(functionParams)
	wasmimport_WasmRPCInvokeAndAwait((uint32)(self0), (*uint8)(functionName0), (uint32)(functionName1), (*WitValue)(functionParams0), (uint32)(functionParams1), &result)
	return
}

// ScheduleCancelableInvocation represents the imported method "schedule-cancelable-invocation".
//
// Schedule invocation for later. Call cancel on the returned resource to cancel the
// invocation before the scheduled time.
//
//	schedule-cancelable-invocation: func(scheduled-time: datetime, function-name: string,
//	function-params: list<wit-value>) -> cancellation-token
//
//go:nosplit
func (self WasmRPC) ScheduleCancelableInvocation(scheduledTime DateTime, functionName string, functionParams cm.List[WitValue]) (result CancellationToken) {
	self0 := cm.Reinterpret[uint32](self)
	scheduledTime0, scheduledTime1 := lower_DateTime(scheduledTime)
	functionName0, functionName1 := cm.LowerString(functionName)
	functionParams0, functionParams1 := cm.LowerList(functionParams)
	result0 := wasmimport_WasmRPCScheduleCancelableInvocation((uint32)(self0), (uint64)(scheduledTime0), (uint32)(scheduledTime1), (*uint8)(functionName0), (uint32)(functionName1), (*WitValue)(functionParams0), (uint32)(functionParams1))
	result = cm.Reinterpret[CancellationToken]((uint32)(result0))
	return
}

// ScheduleInvocation represents the imported method "schedule-invocation".
//
// Schedule invocation for later
//
//	schedule-invocation: func(scheduled-time: datetime, function-name: string, function-params:
//	list<wit-value>)
//
//go:nosplit
func (self WasmRPC) ScheduleInvocation(scheduledTime DateTime, functionName string, functionParams cm.List[WitValue]) {
	self0 := cm.Reinterpret[uint32](self)
	scheduledTime0, scheduledTime1 := lower_DateTime(scheduledTime)
	functionName0, functionName1 := cm.LowerString(functionName)
	functionParams0, functionParams1 := cm.LowerList(functionParams)
	wasmimport_WasmRPCScheduleInvocation((uint32)(self0), (uint64)(scheduledTime0), (uint32)(scheduledTime1), (*uint8)(functionName0), (uint32)(functionName1), (*WitValue)(functionParams0), (uint32)(functionParams1))
	return
}

// FutureInvokeResult represents the imported resource "golem:rpc/types@0.2.0#future-invoke-result".
//
//	resource future-invoke-result
type FutureInvokeResult cm.Resource

// ResourceDrop represents the imported resource-drop for resource "future-invoke-result".
//
// Drops a resource handle.
//
//go:nosplit
func (self FutureInvokeResult) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FutureInvokeResultResourceDrop((uint32)(self0))
	return
}

// Get represents the imported method "get".
//
//	get: func() -> option<result<wit-value, rpc-error>>
//
//go:nosplit
func (self FutureInvokeResult) Get() (result cm.Option[cm.Result[RPCErrorShape, WitValue, RPCError]]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_FutureInvokeResultGet((uint32)(self0), &result)
	return
}

// Subscribe represents the imported method "subscribe".
//
//	subscribe: func() -> pollable
//
//go:nosplit
func (self FutureInvokeResult) Subscribe() (result Pollable) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_FutureInvokeResultSubscribe((uint32)(self0))
	result = cm.Reinterpret[Pollable]((uint32)(result0))
	return
}

// CancellationToken represents the imported resource "golem:rpc/types@0.2.0#cancellation-token".
//
//	resource cancellation-token
type CancellationToken cm.Resource

// ResourceDrop represents the imported resource-drop for resource "cancellation-token".
//
// Drops a resource handle.
//
//go:nosplit
func (self CancellationToken) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CancellationTokenResourceDrop((uint32)(self0))
	return
}

// Cancel represents the imported method "cancel".
//
//	cancel: func()
//
//go:nosplit
func (self CancellationToken) Cancel() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_CancellationTokenCancel((uint32)(self0))
	return
}

// ParseUUID represents the imported function "parse-uuid".
//
// Parses a UUID from a string
//
//	parse-uuid: func(uuid: string) -> result<uuid, string>
//
//go:nosplit
func ParseUUID(uuid string) (result cm.Result[UUIDShape, UUID, string]) {
	uuid0, uuid1 := cm.LowerString(uuid)
	wasmimport_ParseUUID((*uint8)(uuid0), (uint32)(uuid1), &result)
	return
}

// UUIDToString represents the imported function "uuid-to-string".
//
// Converts a UUID to a string
//
//	uuid-to-string: func(uuid: uuid) -> string
//
//go:nosplit
func UUIDToString(uuid UUID) (result string) {
	uuid0, uuid1 := lower_UUID(uuid)
	wasmimport_UUIDToString((uint64)(uuid0), (uint64)(uuid1), &result)
	return
}

// ExtractValue represents the imported function "extract-value".
//
//	extract-value: func(vnt: value-and-type) -> wit-value
//
//go:nosplit
func ExtractValue(vnt ValueAndType) (result WitValue) {
	vnt0, vnt1, vnt2, vnt3 := lower_ValueAndType(vnt)
	wasmimport_ExtractValue((*WitNode)(vnt0), (uint32)(vnt1), (*WitTypeNode)(vnt2), (uint32)(vnt3), &result)
	return
}

// ExtractType represents the imported function "extract-type".
//
//	extract-type: func(vnt: value-and-type) -> wit-type
//
//go:nosplit
func ExtractType(vnt ValueAndType) (result WitType) {
	vnt0, vnt1, vnt2, vnt3 := lower_ValueAndType(vnt)
	wasmimport_ExtractType((*WitNode)(vnt0), (uint32)(vnt1), (*WitTypeNode)(vnt2), (uint32)(vnt3), &result)
	return
}
