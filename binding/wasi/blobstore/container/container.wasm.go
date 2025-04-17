// Code generated by wit-bindgen-go. DO NOT EDIT.

package container

import (
	"go.bytecodealliance.org/cm"
)

// This file contains wasmimport and wasmexport declarations for "wasi:blobstore".

//go:wasmimport wasi:blobstore/container [resource-drop]container
//go:noescape
func wasmimport_ContainerResourceDrop(self0 uint32)

//go:wasmimport wasi:blobstore/container [method]container.clear
//go:noescape
func wasmimport_ContainerClear(self0 uint32, result *cm.Result[Error, struct{}, Error])

//go:wasmimport wasi:blobstore/container [method]container.delete-object
//go:noescape
func wasmimport_ContainerDeleteObject(self0 uint32, name0 *uint8, name1 uint32, result *cm.Result[Error, struct{}, Error])

//go:wasmimport wasi:blobstore/container [method]container.delete-objects
//go:noescape
func wasmimport_ContainerDeleteObjects(self0 uint32, names0 *ObjectName, names1 uint32, result *cm.Result[Error, struct{}, Error])

//go:wasmimport wasi:blobstore/container [method]container.get-data
//go:noescape
func wasmimport_ContainerGetData(self0 uint32, name0 *uint8, name1 uint32, start0 uint64, end0 uint64, result *cm.Result[string, IncomingValue, Error])

//go:wasmimport wasi:blobstore/container [method]container.has-object
//go:noescape
func wasmimport_ContainerHasObject(self0 uint32, name0 *uint8, name1 uint32, result *cm.Result[string, bool, Error])

//go:wasmimport wasi:blobstore/container [method]container.info
//go:noescape
func wasmimport_ContainerInfo(self0 uint32, result *cm.Result[ContainerMetadataShape, ContainerMetadata, Error])

//go:wasmimport wasi:blobstore/container [method]container.list-objects
//go:noescape
func wasmimport_ContainerListObjects(self0 uint32, result *cm.Result[string, StreamObjectNames, Error])

//go:wasmimport wasi:blobstore/container [method]container.name
//go:noescape
func wasmimport_ContainerName(self0 uint32, result *cm.Result[string, string, Error])

//go:wasmimport wasi:blobstore/container [method]container.object-info
//go:noescape
func wasmimport_ContainerObjectInfo(self0 uint32, name0 *uint8, name1 uint32, result *cm.Result[ObjectMetadataShape, ObjectMetadata, Error])

//go:wasmimport wasi:blobstore/container [method]container.write-data
//go:noescape
func wasmimport_ContainerWriteData(self0 uint32, name0 *uint8, name1 uint32, data0 uint32, result *cm.Result[Error, struct{}, Error])

//go:wasmimport wasi:blobstore/container [resource-drop]stream-object-names
//go:noescape
func wasmimport_StreamObjectNamesResourceDrop(self0 uint32)

//go:wasmimport wasi:blobstore/container [method]stream-object-names.read-stream-object-names
//go:noescape
func wasmimport_StreamObjectNamesReadStreamObjectNames(self0 uint32, len0 uint64, result *cm.Result[TupleListObjectNameBoolShape, cm.Tuple[cm.List[ObjectName], bool], Error])

//go:wasmimport wasi:blobstore/container [method]stream-object-names.skip-stream-object-names
//go:noescape
func wasmimport_StreamObjectNamesSkipStreamObjectNames(self0 uint32, num0 uint64, result *cm.Result[TupleU64BoolShape, cm.Tuple[uint64, bool], Error])
