package os

import (
	"fmt"
	"os"
	"sync"

	"github.com/golemcloud/golem-go/binding"
)

var (
	env     map[string]string
	envOnce sync.Once
)

func initEnv() {
	env := map[string]string{}
	e := binding.WasiCli0_2_0_EnvironmentGetEnvironment()
	for i := range e {
		env[e[i].F0] = e[i].F1
	}
}

// InitStdEnv overrides standard lib env vars with the ones coming from the WASI environment
func InitStdEnv() {
	// NOTE: os.ClearEnv panics with "nil pointer dereference" currently, but the env is empty anyway
	// os.Clearenv()
	e := binding.WasiCli0_2_0_EnvironmentGetEnvironment()
	for i := range e {
		err := os.Setenv(e[i].F0, e[i].F1)
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
