// Code generated by wit-bindgen-go. DO NOT EDIT.

package blobstore

import (
	"github.com/golemcloud/golem-go/binding/wasi/blobstore/types"
	"go.bytecodealliance.org/cm"
)

func lower_ObjectID(v types.ObjectID) (f0 *uint8, f1 uint32, f2 *uint8, f3 uint32) {
	f0, f1 = cm.LowerString(v.Container)
	f2, f3 = cm.LowerString(v.Object)
	return
}
