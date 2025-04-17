package golemhost

import (
	"fmt"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
)

type PersistenceLevel int

const (
	PersistenceLevelPersistNothing PersistenceLevel = iota
	PersistenceLevelPersistRemoteSideEffects
	PersistenceLevelSmart
)

func NewPersistenceLevel(level host.PersistenceLevel) PersistenceLevel {
	switch level {
	case host.PersistenceLevelPersistRemoteSideEffects:
		return PersistenceLevelPersistRemoteSideEffects
	case host.PersistenceLevelPersistNothing:
		return PersistenceLevelPersistNothing
	case host.PersistenceLevelSmart:
		return PersistenceLevelSmart
	default:
		panic(fmt.Sprintf("NewPersistenceLevel: unhandled persistence level: %d", level))
	}
}

func (level PersistenceLevel) ToBinding() host.PersistenceLevel {
	switch level {
	case PersistenceLevelPersistNothing:
		return host.PersistenceLevelPersistNothing
	case PersistenceLevelPersistRemoteSideEffects:
		return host.PersistenceLevelPersistRemoteSideEffects
	case PersistenceLevelSmart:
		return host.PersistenceLevelSmart
	default:
		panic(fmt.Sprintf("ToBinding: unhandled persistence level: %d", level))
	}
}

func SetPersistenceLevel(level PersistenceLevel) {
	host.SetOplogPersistenceLevel(level.ToBinding())
}

func GetPersistenceLevel() PersistenceLevel {
	return NewPersistenceLevel(host.GetOplogPersistenceLevel())
}

func WithPersistenceLevel[T any](level PersistenceLevel, f func() (T, error)) (T, error) {
	currentLevel := GetPersistenceLevel()
	defer SetPersistenceLevel(currentLevel)
	SetPersistenceLevel(level)
	return f()
}
