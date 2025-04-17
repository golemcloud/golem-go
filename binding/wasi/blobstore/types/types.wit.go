// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package types represents the imported interface "wasi:blobstore/types".
//
// Types used by blobstore
package types

import (
	"github.com/golemcloud/golem-go/binding/wasi/io/streams"
	"go.bytecodealliance.org/cm"
)

// InputStream represents the imported type alias "wasi:blobstore/types#input-stream".
//
// See [streams.InputStream] for more information.
type InputStream = streams.InputStream

// OutputStream represents the imported type alias "wasi:blobstore/types#output-stream".
//
// See [streams.OutputStream] for more information.
type OutputStream = streams.OutputStream

// ContainerName represents the string "wasi:blobstore/types#container-name".
//
// name of a container, a collection of objects.
// The container name may be any valid UTF-8 string.
//
//	type container-name = string
type ContainerName string

// ObjectName represents the string "wasi:blobstore/types#object-name".
//
// name of an object within a container
// The object name may be any valid UTF-8 string.
//
//	type object-name = string
type ObjectName string

// Timestamp represents the u64 "wasi:blobstore/types#timestamp".
//
// TODO: define timestamp to include seconds since
// Unix epoch and nanoseconds
// https://github.com/WebAssembly/wasi-blob-store/issues/7
//
//	type timestamp = u64
type Timestamp uint64

// ObjectSize represents the u64 "wasi:blobstore/types#object-size".
//
// size of an object, in bytes
//
//	type object-size = u64
type ObjectSize uint64

// Error represents the string "wasi:blobstore/types#error".
//
//	type error = string
type Error string

// ContainerMetadata represents the record "wasi:blobstore/types#container-metadata".
//
// information about a container
//
//	record container-metadata {
//		name: container-name,
//		created-at: timestamp,
//	}
type ContainerMetadata struct {
	_ cm.HostLayout `json:"-"`
	// the container's name
	Name ContainerName `json:"name"`

	// date and time container was created
	CreatedAt Timestamp `json:"created-at"`
}

// ObjectMetadata represents the record "wasi:blobstore/types#object-metadata".
//
// information about an object
//
//	record object-metadata {
//		name: object-name,
//		container: container-name,
//		created-at: timestamp,
//		size: object-size,
//	}
type ObjectMetadata struct {
	_ cm.HostLayout `json:"-"`
	// the object's name
	Name ObjectName `json:"name"`

	// the object's parent container
	Container ContainerName `json:"container"`

	// date and time the object was created
	CreatedAt Timestamp `json:"created-at"`

	// size of the object, in bytes
	Size ObjectSize `json:"size"`
}

// ObjectID represents the record "wasi:blobstore/types#object-id".
//
// identifier for an object that includes its container name
//
//	record object-id {
//		container: container-name,
//		object: object-name,
//	}
type ObjectID struct {
	_         cm.HostLayout `json:"-"`
	Container ContainerName `json:"container"`
	Object    ObjectName    `json:"object"`
}

// OutgoingValue represents the imported resource "wasi:blobstore/types#outgoing-value".
//
// A data is the data stored in a data blob. The value can be of any type
// that can be represented in a byte array. It provides a way to write the value
// to the output-stream defined in the `wasi-io` interface.
// Soon: switch to `resource value { ... }`
//
//	resource outgoing-value
type OutgoingValue cm.Resource

// ResourceDrop represents the imported resource-drop for resource "outgoing-value".
//
// Drops a resource handle.
//
//go:nosplit
func (self OutgoingValue) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingValueResourceDrop((uint32)(self0))
	return
}

// OutgoingValueNewOutgoingValue represents the imported static function "new-outgoing-value".
//
//	new-outgoing-value: static func() -> outgoing-value
//
//go:nosplit
func OutgoingValueNewOutgoingValue() (result OutgoingValue) {
	result0 := wasmimport_OutgoingValueNewOutgoingValue()
	result = cm.Reinterpret[OutgoingValue]((uint32)(result0))
	return
}

// OutgoingValueWriteBody represents the imported method "outgoing-value-write-body".
//
//	outgoing-value-write-body: func() -> result<output-stream>
//
//go:nosplit
func (self OutgoingValue) OutgoingValueWriteBody() (result cm.Result[OutputStream, OutputStream, struct{}]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_OutgoingValueOutgoingValueWriteBody((uint32)(self0), &result)
	return
}

// IncomingValue represents the imported resource "wasi:blobstore/types#incoming-value".
//
// A incoming-value is a wrapper around a value. It provides a way to read the value
// from the input-stream defined in the `wasi-io` interface.
//
// The incoming-value provides two ways to consume the value:
// 1. `incoming-value-consume-sync` consumes the value synchronously and returns the
// value as a list of bytes.
// 2. `incoming-value-consume-async` consumes the value asynchronously and returns
// the
// value as an input-stream.
// Soon: switch to `resource incoming-value { ... }`
//
//	resource incoming-value
type IncomingValue cm.Resource

// ResourceDrop represents the imported resource-drop for resource "incoming-value".
//
// Drops a resource handle.
//
//go:nosplit
func (self IncomingValue) ResourceDrop() {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_IncomingValueResourceDrop((uint32)(self0))
	return
}

// IncomingValueConsumeAsync represents the imported method "incoming-value-consume-async".
//
//	incoming-value-consume-async: func() -> result<incoming-value-async-body, error>
//
//go:nosplit
func (self IncomingValue) IncomingValueConsumeAsync() (result cm.Result[string, IncomingValueAsyncBody, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_IncomingValueIncomingValueConsumeAsync((uint32)(self0), &result)
	return
}

// IncomingValueConsumeSync represents the imported method "incoming-value-consume-sync".
//
//	incoming-value-consume-sync: func() -> result<incoming-value-sync-body, error>
//
//go:nosplit
func (self IncomingValue) IncomingValueConsumeSync() (result cm.Result[IncomingValueSyncBody, IncomingValueSyncBody, Error]) {
	self0 := cm.Reinterpret[uint32](self)
	wasmimport_IncomingValueIncomingValueConsumeSync((uint32)(self0), &result)
	return
}

// Size represents the imported method "size".
//
//	size: func() -> u64
//
//go:nosplit
func (self IncomingValue) Size() (result uint64) {
	self0 := cm.Reinterpret[uint32](self)
	result0 := wasmimport_IncomingValueSize((uint32)(self0))
	result = (uint64)((uint64)(result0))
	return
}

// IncomingValueAsyncBody represents the imported type alias "wasi:blobstore/types#incoming-value-async-body".
//
// See [InputStream] for more information.
type IncomingValueAsyncBody = InputStream

// IncomingValueSyncBody represents the list "wasi:blobstore/types#incoming-value-sync-body".
//
//	type incoming-value-sync-body = list<u8>
type IncomingValueSyncBody cm.List[uint8]
