package golemhost

import binding "github.com/golemcloud/golem-go/golem_go_bindings"

func GetIdempotenceMode() bool {
	return binding.GolemApi0_2_0_HostGetIdempotenceMode()
}

func SetIdempotenceMode(idempotent bool) {
	binding.GolemApi0_2_0_HostSetIdempotenceMode(idempotent)
}

func WithIdempotenceMode[T any](idempotent bool, f func() (T, error)) (T, error) {
	currentMode := GetIdempotenceMode()
	defer SetIdempotenceMode(currentMode)
	SetIdempotenceMode(idempotent)
	return f()
}
