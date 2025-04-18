// Code generated by wit-bindgen-go. DO NOT EDIT.

package types

import (
	wallclock "github.com/golemcloud/golem-go/binding/wasi/clocks/wall-clock"
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// MetadataHashValueShape is used for storage in variant or result types.
type MetadataHashValueShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(MetadataHashValue{})]byte
}

// TupleListU8BoolShape is used for storage in variant or result types.
type TupleListU8BoolShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(cm.Tuple[cm.List[uint8], bool]{})]byte
}

func lower_DateTime(v wallclock.DateTime) (f0 uint64, f1 uint32) {
	f0 = (uint64)(v.Seconds)
	f1 = (uint32)(v.Nanoseconds)
	return
}

func lower_NewTimestamp(v NewTimestamp) (f0 uint32, f1 uint64, f2 uint32) {
	f0 = (uint32)(v.Tag())
	switch f0 {
	case 2: // timestamp
		v1, v2 := lower_DateTime(*cm.Case[DateTime](&v, 2))
		f1 = (uint64)(v1)
		f2 = (uint32)(v2)
	}
	return
}

// DescriptorStatShape is used for storage in variant or result types.
type DescriptorStatShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(DescriptorStat{})]byte
}

// OptionDirectoryEntryShape is used for storage in variant or result types.
type OptionDirectoryEntryShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(cm.Option[DirectoryEntry]{})]byte
}
