// Code generated by wit-bindgen-go. DO NOT EDIT.

package oplog

import (
	"github.com/golemcloud/golem-go/binding/golem/rpc/types"
	"go.bytecodealliance.org/cm"
	"unsafe"
)

// LocalSpanDataShape is used for storage in variant or result types.
type LocalSpanDataShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(LocalSpanData{})]byte
}

// ExportedFunctionInvocationParametersShape is used for storage in variant or result types.
type ExportedFunctionInvocationParametersShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(ExportedFunctionInvocationParameters{})]byte
}

// CreateParametersShape is used for storage in variant or result types.
type CreateParametersShape struct {
	_     cm.HostLayout
	shape [unsafe.Sizeof(CreateParameters{})]byte
}

func lower_UUID(v types.UUID) (f0 uint64, f1 uint64) {
	f0 = (uint64)(v.HighBits)
	f1 = (uint64)(v.LowBits)
	return
}

func lower_ComponentID(v types.ComponentID) (f0 uint64, f1 uint64) {
	f0, f1 = lower_UUID(v.UUID)
	return
}

func lower_WorkerID(v types.WorkerID) (f0 uint64, f1 uint64, f2 *uint8, f3 uint32) {
	f0, f1 = lower_ComponentID(v.ComponentID)
	f2, f3 = cm.LowerString(v.WorkerName)
	return
}
