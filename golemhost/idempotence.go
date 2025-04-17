package golemhost

import "github.com/golemcloud/golem-go/binding/golem/api/host"

func GetIdempotenceMode() bool {
	return host.GetIdempotenceMode()
}

func SetIdempotenceMode(idempotent bool) {
	host.SetIdempotenceMode(idempotent)
}

func WithIdempotenceMode[T any](idempotent bool, f func() (T, error)) (T, error) {
	currentMode := GetIdempotenceMode()
	defer SetIdempotenceMode(currentMode)
	SetIdempotenceMode(idempotent)
	return f()
}
