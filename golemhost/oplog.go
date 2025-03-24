package golemhost

import "github.com/golemcloud/golem-go/binding"

type OpLogIndex binding.GolemApi1_1_6_HostOplogIndex

func OpLogCommit(replicas uint8) {
	binding.GolemApi1_1_6_HostOplogCommit(replicas)
}

func MarkBeginOperation() OpLogIndex {
	return OpLogIndex(binding.GolemApi1_1_6_HostMarkBeginOperation())
}

func MarkEndOperation(index OpLogIndex) {
	binding.GolemApi1_1_6_HostMarkEndOperation(binding.GolemApi1_1_6_HostOplogIndex(index))
}

func Atomically[T any](f func() (T, error)) (T, error) {
	index := MarkBeginOperation()
	defer MarkEndOperation(index)
	return f()
}
