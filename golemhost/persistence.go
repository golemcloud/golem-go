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

func NewPersistenceLevel(level binding.GolemApi1_1_6_HostPersistenceLevel) PersistenceLevel {
	switch level.Kind() {
	case binding.GolemApi1_1_6_HostPersistenceLevelKindPersistRemoteSideEffects:
		return PersistenceLevelPersistRemoteSideEffects
	case binding.GolemApi1_1_6_HostPersistenceLevelKindPersistNothing:
		return PersistenceLevelPersistNothing
	case binding.GolemApi1_1_6_HostPersistenceLevelKindSmart:
		return PersistenceLevelSmart
	default:
		panic(fmt.Sprintf("NewPersistenceLevel: unhandled persistence level: %d", level))
	}
}

func (level PersistenceLevel) ToBinding() binding.GolemApi1_1_6_HostPersistenceLevel {
	switch level {
	case PersistenceLevelPersistNothing:
		return binding.GolemApi1_1_6_HostPersistenceLevelPersistNothing()
	case PersistenceLevelPersistRemoteSideEffects:
		return binding.GolemApi1_1_6_HostPersistenceLevelPersistRemoteSideEffects()
	case PersistenceLevelSmart:
		return binding.GolemApi1_1_6_HostPersistenceLevelSmart()
	default:
		panic(fmt.Sprintf("ToBinding: unhandled persistence level: %d", level))
	}
}

func SetPersistenceLevel(level PersistenceLevel) {
	binding.GolemApi1_1_6_HostSetOplogPersistenceLevel(level.ToBinding())
}

func GetPersistenceLevel() PersistenceLevel {
	return NewPersistenceLevel(binding.GolemApi1_1_6_HostGetOplogPersistenceLevel())
}

func WithPersistenceLevel[T any](level PersistenceLevel, f func() (T, error)) (T, error) {
	currentLevel := GetPersistenceLevel()
	defer SetPersistenceLevel(currentLevel)
	SetPersistenceLevel(level)
	return f()
}
