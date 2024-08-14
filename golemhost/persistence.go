package golemhost

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding"
)

type PersistenceLevel binding.GolemApi0_2_0_HostPersistenceLevelKind

const (
	PersistenceLevelPersistNothing PersistenceLevel = iota
	PersistenceLevelPersistRemoteSideEffects
	PersistenceLevelSmart
)

func newPersistenceLevel(level binding.GolemApi0_2_0_HostPersistenceLevel) PersistenceLevel {
	switch level.Kind() {
	case binding.GolemApi0_2_0_HostPersistenceLevelKindPersistRemoteSideEffects:
		return PersistenceLevelPersistRemoteSideEffects
	case binding.GolemApi0_2_0_HostPersistenceLevelKindPersistNothing:
		return PersistenceLevelPersistNothing
	case binding.GolemApi0_2_0_HostPersistenceLevelKindSmart:
		return PersistenceLevelSmart
	default:
		panic(fmt.Sprintf("newPersistenceLevel: illegal persistence level: %d", level))
	}
}

func (level PersistenceLevel) toBindingLevel() binding.GolemApi0_2_0_HostPersistenceLevel {
	switch level {
	case PersistenceLevelPersistNothing:
		return binding.GolemApi0_2_0_HostPersistenceLevelPersistNothing()
	case PersistenceLevelPersistRemoteSideEffects:
		return binding.GolemApi0_2_0_HostPersistenceLevelPersistRemoteSideEffects()
	case PersistenceLevelSmart:
		return binding.GolemApi0_2_0_HostPersistenceLevelSmart()
	default:
		panic(fmt.Sprintf("toBindingLevel: illegal persistence level: %d", level))
	}
}

func SetPersistenceLevel(level PersistenceLevel) {
	binding.GolemApi0_2_0_HostSetOplogPersistenceLevel(level.toBindingLevel())
}

func GetPersistenceLevel() PersistenceLevel {
	return newPersistenceLevel(binding.GolemApi0_2_0_HostGetOplogPersistenceLevel())
}

func WithPersistenceLevel[T any](level PersistenceLevel, f func() (T, error)) (T, error) {
	currentLevel := GetPersistenceLevel()
	defer SetPersistenceLevel(currentLevel)
	SetPersistenceLevel(level)
	return f()
}
