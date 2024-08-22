package golemhost

import (
	"encoding/json"

	"github.com/golemcloud/golem-go/binding"
)

type PromiseID struct {
	WorkerID WorkerID
	OplogIdx OpLogIndex
}

func NewPromise() PromiseID {
	promise := binding.GolemApi0_2_0_HostCreatePromise()
	return PromiseID{
		WorkerID: NewWorkerID(promise.WorkerId),
		OplogIdx: OpLogIndex(promise.OplogIdx),
	}
}

func (promiseID PromiseID) ToBinding() binding.GolemApi0_2_0_HostPromiseId {
	return binding.GolemApi0_2_0_HostPromiseId{
		WorkerId: promiseID.WorkerID.ToBinding(),
		OplogIdx: binding.GolemApi0_2_0_HostOplogIndex(promiseID.OplogIdx),
	}
}

func DeletePromise(promiseID PromiseID) {
	binding.GolemApi0_2_0_HostDeletePromise(promiseID.ToBinding())
}

func AwaitPromise(promiseID PromiseID) []byte {
	return binding.GolemApi0_2_0_HostAwaitPromise(promiseID.ToBinding())
}

func AwaitPromiseJSON(promiseID PromiseID, v any) error {
	return json.Unmarshal(AwaitPromise(promiseID), v)
}

func CompletePromise(promiseID PromiseID, payload []byte) bool {
	return binding.GolemApi0_2_0_HostCompletePromise(promiseID.ToBinding(), payload)
}

func CompletePromiseJSON(promiseID PromiseID, v any) (bool, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return false, err
	}
	return CompletePromise(promiseID, bs), nil
}
