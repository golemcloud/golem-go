package golemhost

import (
	"fmt"

	"github.com/google/uuid"
	"go.bytecodealliance.org/cm"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
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

func NewWorkerStatus(status host.WorkerStatus) WorkerStatus {
	switch status {
	case host.WorkerStatusRunning:
		return WorkerStatusRunning
	case host.WorkerStatusIdle:
		return WorkerStatusIdle
	case host.WorkerStatusSuspended:
		return WorkerStatusSuspended
	case host.WorkerStatusInterrupted:
		return WorkerStatusInterrupted
	case host.WorkerStatusRetrying:
		return WorkerStatusRetrying
	case host.WorkerStatusFailed:
		return WorkerStatusFailed
	case host.WorkerStatusExited:
		return WorkerStatusExited
	default:
		panic(fmt.Sprintf("NewWorkerStatus: unhandled status: %s", status.String()))
	}
}

func (ws WorkerStatus) ToBinding() host.WorkerStatus {
	switch ws {
	case WorkerStatusRunning:
		return host.WorkerStatusRunning
	case WorkerStatusIdle:
		return host.WorkerStatusIdle
	case WorkerStatusSuspended:
		return host.WorkerStatusSuspended
	case WorkerStatusInterrupted:
		return host.WorkerStatusInterrupted
	case WorkerStatusRetrying:
		return host.WorkerStatusRetrying
	case WorkerStatusFailed:
		return host.WorkerStatusFailed
	case WorkerStatusExited:
		return host.WorkerStatusExited
	default:
		panic(fmt.Sprintf("ToBinding: unhandled status: %d", ws))
	}
}

type WorkerID struct {
	ComponentID ComponentID
	WorkerName  string
}

func NewWorkerID(workerID host.WorkerID) WorkerID {
	return WorkerID{
		ComponentID: NewComponentID(workerID.ComponentID),
		WorkerName:  workerID.WorkerName,
	}
}

func (workerID WorkerID) ToBinding() host.WorkerID {
	return host.WorkerID{
		ComponentID: workerID.ComponentID.ToBinding(),
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

func NewWorkerMetadata(metadata host.WorkerMetadata) WorkerMetadata {
	envVars := make([]WorkerMetadataEnvVar, metadata.Env.Len())
	envSlice := metadata.Env.Slice()
	for i := range envSlice {
		envVars[i] = WorkerMetadataEnvVar{
			Name:  envSlice[i][0],
			Value: envSlice[i][1],
		}
	}

	return WorkerMetadata{
		WorkerId:         NewWorkerID(metadata.WorkerID),
		Args:             metadata.Args.Slice(),
		Env:              envVars,
		Status:           NewWorkerStatus(metadata.Status),
		ComponentVersion: metadata.ComponentVersion,
		RetryCount:       metadata.RetryCount,
	}
}

func GetSelfMetadata() WorkerMetadata {
	return NewWorkerMetadata(host.GetSelfMetadata())
}

func GetWorkerMetadata(workerID WorkerID) *WorkerMetadata {
	bindingMetadata := host.GetWorkerMetadata(workerID.ToBinding())
	if bindingMetadata.None() {
		return nil
	}
	metadata := NewWorkerMetadata(*bindingMetadata.Some())
	return &metadata
}

// GetWorkers enumerates all the workers optionally matching the provided filter
// NOTE: Enumerating workers of a component is a slow operation and should not be used as part of the application logic.
func GetWorkers(componentID ComponentID, filter *WorkerAnyFilter) []WorkerMetadata {
	bindingFilter := cm.None[host.WorkerAnyFilter]()
	if filter == nil {
		bindingFilter = cm.Some(filter.ToBinding())
	}

	iter := host.NewGetWorkers(componentID.ToBinding(), bindingFilter, true)

	var results []WorkerMetadata
	for {
		nextResults := iter.GetNext()
		if nextResults.None() {
			break
		}

		for _, metadata := range nextResults.Some().Slice() {
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

func (updateMode UpdateMode) ToBinding() host.UpdateMode {
	switch updateMode {
	case UpdateModeAutomatic:
		return host.UpdateModeAutomatic
	case UpdateModeSnapshotBased:
		return host.UpdateModeSnapshotBased
	default:
		panic(fmt.Sprintf("ToBinding: unhandled update mode: %d", updateMode))
	}
}

func UpdateWorker(workerID WorkerID, targetVersion uint64, updateMode UpdateMode) {
	host.UpdateWorker(workerID.ToBinding(), host.ComponentVersion(targetVersion), updateMode.ToBinding())
}

func GenerateIdempotencyKey() uuid.UUID {
	return NewUUID(host.GenerateIdempotencyKey())
}
