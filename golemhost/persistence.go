package golemhost

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding"
)

type PersistenceLevel int

const (
	PersistenceLevelPersistNothing PersistenceLevel = iota
	PersistenceLevelPersistRemoteSideEffects
	PersistenceLevelSmart
)

func NewPersistenceLevel(level binding.GolemApi1_1_0_HostPersistenceLevel) PersistenceLevel {
	switch level.Kind() {
	case binding.GolemApi1_1_0_HostPersistenceLevelKindPersistRemoteSideEffects:
		return PersistenceLevelPersistRemoteSideEffects
	case binding.GolemApi1_1_0_HostPersistenceLevelKindPersistNothing:
		return PersistenceLevelPersistNothing
	case binding.GolemApi1_1_0_HostPersistenceLevelKindSmart:
		return PersistenceLevelSmart
	default:
		panic(fmt.Sprintf("NewPersistenceLevel: unhandled persistence level: %d", level))
	}
}

func (level PersistenceLevel) ToBinding() binding.GolemApi1_1_0_HostPersistenceLevel {
	switch level {
	case PersistenceLevelPersistNothing:
		return binding.GolemApi1_1_0_HostPersistenceLevelPersistNothing()
	case PersistenceLevelPersistRemoteSideEffects:
		return binding.GolemApi1_1_0_HostPersistenceLevelPersistRemoteSideEffects()
	case PersistenceLevelSmart:
		return binding.GolemApi1_1_0_HostPersistenceLevelSmart()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled persistence level: %d", level))
	}
}

func SetPersistenceLevel(level PersistenceLevel) {
	binding.GolemApi1_1_0_HostSetOplogPersistenceLevel(level.ToBinding())
}

func GetPersistenceLevel() PersistenceLevel {
	return NewPersistenceLevel(binding.GolemApi1_1_0_HostGetOplogPersistenceLevel())
}

func WithPersistenceLevel[T any](level PersistenceLevel, f func() (T, error)) (T, error) {
	currentLevel := GetPersistenceLevel()
	defer SetPersistenceLevel(currentLevel)
	SetPersistenceLevel(level)
	return f()
}
