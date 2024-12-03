package golemhost

import (
	"fmt"

	"github.com/google/uuid"

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

func NewWorkerStatus(status binding.GolemApi1_1_0_HostWorkerStatus) WorkerStatus {
	switch status.Kind() {
	case binding.GolemApi1_1_0_HostWorkerStatusKindRunning:
		return WorkerStatusRunning
	case binding.GolemApi1_1_0_HostWorkerStatusKindIdle:
		return WorkerStatusIdle
	case binding.GolemApi1_1_0_HostWorkerStatusKindSuspended:
		return WorkerStatusSuspended
	case binding.GolemApi1_1_0_HostWorkerStatusKindInterrupted:
		return WorkerStatusInterrupted
	case binding.GolemApi1_1_0_HostWorkerStatusKindRetrying:
		return WorkerStatusRetrying
	case binding.GolemApi1_1_0_HostWorkerStatusKindFailed:
		return WorkerStatusFailed
	case binding.GolemApi1_1_0_HostWorkerStatusKindExited:
		return WorkerStatusExited
	default:
		panic(fmt.Sprintf("NewWorkerStatus: unhandled status: %d", status.Kind()))
	}
}

func (ws WorkerStatus) ToBinding() binding.GolemApi1_1_0_HostWorkerStatus {
	switch ws {
	case WorkerStatusRunning:
		return binding.GolemApi1_1_0_HostWorkerStatusRunning()
	case WorkerStatusIdle:
		return binding.GolemApi1_1_0_HostWorkerStatusIdle()
	case WorkerStatusSuspended:
		return binding.GolemApi1_1_0_HostWorkerStatusSuspended()
	case WorkerStatusInterrupted:
		return binding.GolemApi1_1_0_HostWorkerStatusInterrupted()
	case WorkerStatusRetrying:
		return binding.GolemApi1_1_0_HostWorkerStatusRetrying()
	case WorkerStatusFailed:
		return binding.GolemApi1_1_0_HostWorkerStatusFailed()
	case WorkerStatusExited:
		return binding.GolemApi1_1_0_HostWorkerStatusExited()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled status: %d", ws))
	}
}

type WorkerID struct {
	ComponentID ComponentID
	WorkerName  string
}

func NewWorkerID(workerID binding.GolemApi1_1_0_HostWorkerId) WorkerID {
	return WorkerID{
		ComponentID: NewComponentID(workerID.ComponentId),
		WorkerName:  workerID.WorkerName,
	}
}

func (workerID WorkerID) ToBinding() binding.GolemApi1_1_0_HostWorkerId {
	return binding.GolemApi1_1_0_HostWorkerId{
		ComponentId: workerID.ComponentID.ToBinding(),
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

func NewWorkerMetadata(metadata binding.GolemApi1_1_0_HostWorkerMetadata) WorkerMetadata {
	envVars := make([]WorkerMetadataEnvVar, len(metadata.Env))
	for i := range metadata.Env {
		envVars[i] = WorkerMetadataEnvVar{
			Name:  metadata.Env[i].F0,
			Value: metadata.Env[i].F1,
		}
	}

	return WorkerMetadata{
		WorkerId:         NewWorkerID(metadata.WorkerId),
		Args:             metadata.Args,
		Env:              envVars,
		Status:           NewWorkerStatus(metadata.Status),
		ComponentVersion: metadata.ComponentVersion,
		RetryCount:       metadata.RetryCount,
	}
}

func GetSelfMetadata() WorkerMetadata {
	return NewWorkerMetadata(binding.GolemApi1_1_0_HostGetSelfMetadata())
}

func GetWorkerMetadata(workerID WorkerID) *WorkerMetadata {
	bindingMetadata := binding.GolemApi1_1_0_HostGetWorkerMetadata(workerID.ToBinding())
	if bindingMetadata.IsNone() {
		return nil
	}
	metadata := NewWorkerMetadata(bindingMetadata.Unwrap())
	return &metadata
}

// GetWorkers enumerates all the workers optionally matching the provided filter
// NOTE: Enumerating workers of a component is a slow operation and should not be used as part of the application logic.
func GetWorkers(componentID ComponentID, filter *WorkerAnyFilter) []WorkerMetadata {
	bindingFilter := binding.None[binding.GolemApi1_1_0_HostWorkerAnyFilter]()
	if filter == nil {
		bindingFilter.Set(filter.ToBinding())
	}

	iter := binding.NewGetWorkers(componentID.ToBinding(), bindingFilter, true)

	var results []WorkerMetadata
	for {
		nextResults := iter.GetNext()
		if nextResults.IsNone() {
			break
		}

		for _, metadata := range nextResults.Unwrap() {
			results = append(results, NewWorkerMetadata(metadata))
		}
	}

	return results
}

type UpdateMode int

const (
	UpdateModeAutomatic UpdateMode = iota
	UpdateModeSnapshotBased
)

func (updateMode UpdateMode) ToBinding() binding.GolemApi1_1_0_HostUpdateMode {
	switch updateMode {
	case UpdateModeAutomatic:
		return binding.GolemApi1_1_0_HostUpdateModeAutomatic()
	case UpdateModeSnapshotBased:
		return binding.GolemApi1_1_0_HostUpdateModeSnapshotBased()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled update mode: %d", updateMode))
	}
}

func UpdateWorker(workerID WorkerID, targetVersion uint64, updateMode UpdateMode) {
	binding.GolemApi1_1_0_HostUpdateWorker(workerID.ToBinding(), targetVersion, updateMode.ToBinding())
}

func GenerateIdempotencyKey() uuid.UUID {
	return NewUUID(binding.GolemApi1_1_0_HostGenerateIdempotencyKey())
}
