package os

import (
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
