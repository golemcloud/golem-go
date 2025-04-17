package golemhost

import "github.com/golemcloud/golem-go/binding/golem/api/host"

type OpLogIndex host.OplogIndex

func OpLogCommit(replicas uint8) {
	host.OplogCommit(replicas)
}

func MarkBeginOperation() OpLogIndex {
	return OpLogIndex(host.MarkBeginOperation())
}

func MarkEndOperation(index OpLogIndex) {
	host.MarkEndOperation(host.OplogIndex(index))
}

func Atomically[T any](f func() (T, error)) (T, error) {
	index := MarkBeginOperation()
	defer MarkEndOperation(index)
	return f()
}
