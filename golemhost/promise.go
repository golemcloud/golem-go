package golemhost

import (
	"encoding/json"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
	"go.bytecodealliance.org/cm"
)

type PromiseID struct {
	WorkerID WorkerID
	OplogIdx OpLogIndex
}

func NewPromise() PromiseID {
	promise := host.CreatePromise()
	return PromiseID{
		WorkerID: NewWorkerID(promise.WorkerID),
		OplogIdx: OpLogIndex(promise.OplogIdx),
	}
}

func (promiseID PromiseID) ToBinding() host.PromiseID {
	return host.PromiseID{
		WorkerID: promiseID.WorkerID.ToBinding(),
		OplogIdx: host.OplogIndex(promiseID.OplogIdx),
	}
}

func DeletePromise(promiseID PromiseID) {
	host.DeletePromise(promiseID.ToBinding())
}

func AwaitPromise(promiseID PromiseID) []byte {
	return host.AwaitPromise(promiseID.ToBinding()).Slice()
}

func AwaitPromiseJSON(promiseID PromiseID, v any) error {
	return json.Unmarshal(AwaitPromise(promiseID), v)
}

func CompletePromise(promiseID PromiseID, payload []byte) bool {
	return host.CompletePromise(promiseID.ToBinding(), cm.ToList(payload))
}

func CompletePromiseJSON(promiseID PromiseID, v any) (bool, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return false, err
	}
	return CompletePromise(promiseID, bs), nil
}
