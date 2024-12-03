package golemhost

import "github.com/golemcloud/golem-go/binding"

type OpLogIndex binding.GolemApi1_1_0_HostOplogIndex

func OpLogCommit(replicas uint8) {
	binding.GolemApi1_1_0_HostOplogCommit(replicas)
}

func MarkBeginOperation() OpLogIndex {
	return OpLogIndex(binding.GolemApi1_1_0_HostMarkBeginOperation())
}

func MarkEndOperation(index OpLogIndex) {
	binding.GolemApi1_1_0_HostMarkEndOperation(binding.GolemApi1_1_0_HostOplogIndex(index))
}

func Atomically[T any](f func() (T, error)) (T, error) {
	index := MarkBeginOperation()
	defer MarkEndOperation(index)
	return f()
}
