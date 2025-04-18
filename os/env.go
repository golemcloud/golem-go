package os

import (
	"fmt"
	"os"
	"sync"

	"github.com/golemcloud/golem-go/binding/wasi/cli/environment"
)

var (
	env     map[string]string
	envOnce sync.Once
)

func initEnv() {
	env := map[string]string{}
	e := environment.GetEnvironment()
	for _, kv := range e.Slice() {
		env[kv[0]] = kv[1]
	}
}

// InitStdEnv overrides standard lib env vars with the ones coming from the WASI environment
func InitStdEnv() {
	// NOTE: os.ClearEnv panics with "nil pointer dereference" currently, but the env is empty anyway
	// os.Clearenv()
	e := environment.GetEnvironment()
	for _, kv := range e.Slice() {
		err := os.Setenv(kv[0], kv[1])
		if err != nil {
			panic(fmt.Sprintf("failed to initialize standard lib environment variables: %+v", err))
		}
	}
}

// Getenv is convenience wrapper to be used instead of the standard lib's os.Getenv
func Getenv(key string) string {
	value, _ := LookupEnv(key)
	return value
}

// LookupEnv is convenience wrapper to be used instead of the standard lib's os.LookupEnv
func LookupEnv(key string) (string, bool) {
	envOnce.Do(initEnv)
	value, ok := env[key]
	return value, ok
}
