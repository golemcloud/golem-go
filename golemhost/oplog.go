package golemhost

import "github.com/golemcloud/golem-go/binding"

type OpLogIndex binding.GolemApi0_2_0_HostOplogIndex

func OpLogCommit(replicas uint8) {
	binding.GolemApi0_2_0_HostOplogCommit(replicas)
}

func MarkBeginOperation() OpLogIndex {
	return OpLogIndex(binding.GolemApi0_2_0_HostMarkBeginOperation())
}

func MarkEndOperation(index OpLogIndex) {
	binding.GolemApi0_2_0_HostMarkEndOperation(binding.GolemApi0_2_0_HostOplogIndex(index))
}

func Atomically[T any](f func() (T, error)) (T, error) {
	index := MarkBeginOperation()
	defer MarkEndOperation(index)
	return f()
}
