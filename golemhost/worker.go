package golemhost

import (
	"fmt"
	"github.com/golemcloud/golem-go/binding"
)

type WorkerStatus int

const (
	WorkerStatusRunning = iota
	WorkerStatusIdle
	WorkerStatusSuspended
	WorkerStatusInterrupted
	WorkerStatusRetrying
	WorkerStatusFailed
	WorkerStatusExited
)

func newWorkerStatus(status binding.GolemApi0_2_0_HostWorkerStatus) WorkerStatus {
	switch status.Kind() {
	case binding.GolemApi0_2_0_HostWorkerStatusKindRunning:
		return WorkerStatusRunning
	case binding.GolemApi0_2_0_HostWorkerStatusKindIdle:
		return WorkerStatusIdle
	case binding.GolemApi0_2_0_HostWorkerStatusKindSuspended:
		return WorkerStatusSuspended
	case binding.GolemApi0_2_0_HostWorkerStatusKindInterrupted:
		return WorkerStatusInterrupted
	case binding.GolemApi0_2_0_HostWorkerStatusKindRetrying:
		return WorkerStatusRetrying
	case binding.GolemApi0_2_0_HostWorkerStatusKindFailed:
		return WorkerStatusFailed
	case binding.GolemApi0_2_0_HostWorkerStatusKindExited:
		return WorkerStatusExited
	default:
		panic(fmt.Sprintf("newWorkerStatus: unhandled status: %d", status.Kind()))
	}
}

func (ws WorkerStatus) toBinding() binding.GolemApi0_2_0_HostWorkerStatus {
	switch ws {
	case WorkerStatusRunning:
		return binding.GolemApi0_2_0_HostWorkerStatusRunning()
	case WorkerStatusIdle:
		return binding.GolemApi0_2_0_HostWorkerStatusIdle()
	case WorkerStatusSuspended:
		return binding.GolemApi0_2_0_HostWorkerStatusSuspended()
	case WorkerStatusInterrupted:
		return binding.GolemApi0_2_0_HostWorkerStatusInterrupted()
	case WorkerStatusRetrying:
		return binding.GolemApi0_2_0_HostWorkerStatusRetrying()
	case WorkerStatusFailed:
		return binding.GolemApi0_2_0_HostWorkerStatusFailed()
	case WorkerStatusExited:
		return binding.GolemApi0_2_0_HostWorkerStatusExited()
	default:
		panic(fmt.Sprintf("toBinding: unhandled status: %d", ws))
	}
}

type WorkerID struct {
	ComponentID ComponentID
	WorkerName  string
}

func newWorkerID(workerID binding.GolemApi0_2_0_HostWorkerId) WorkerID {
	return WorkerID{
		ComponentID: newComponentID(workerID.ComponentId),
		WorkerName:  workerID.WorkerName,
	}
}

func (workerID WorkerID) toBinding() binding.GolemApi0_2_0_HostWorkerId {
	return binding.GolemApi0_2_0_HostWorkerId{
		ComponentId: workerID.ComponentID.toBinding(),
		WorkerName:  workerID.WorkerName,
	}
}

type WorkerMetadataEnvVar struct {
	Name  string
	Value string
}

type WorkerMetadata struct {
	WorkerId         WorkerID
	Args             []string
	Env              []WorkerMetadataEnvVar
	Status           WorkerStatus
	ComponentVersion uint64
	RetryCount       uint64
}

func newWorkerMetadata(metadata binding.GolemApi0_2_0_HostWorkerMetadata) WorkerMetadata {
	envVars := make([]WorkerMetadataEnvVar, len(metadata.Env))
	for i := range metadata.Env {
		envVars[i] = WorkerMetadataEnvVar{
			Name:  metadata.Env[i].F0,
			Value: metadata.Env[i].F1,
		}
	}

	return WorkerMetadata{
		WorkerId:         newWorkerID(metadata.WorkerId),
		Args:             metadata.Args,
		Env:              envVars,
		Status:           newWorkerStatus(metadata.Status),
		ComponentVersion: metadata.ComponentVersion,
		RetryCount:       metadata.RetryCount,
	}
}

func GetSelfMetadata() WorkerMetadata {
	return newWorkerMetadata(binding.GolemApi0_2_0_HostGetSelfMetadata())
}

func GetWorkerMetadata(workerID WorkerID) *WorkerMetadata {
	bindingMetadata := binding.GolemApi0_2_0_HostGetWorkerMetadata(workerID.toBinding())
	if bindingMetadata.IsNone() {
		return nil
	}
	metadata := newWorkerMetadata(bindingMetadata.Unwrap())
	return &metadata
}