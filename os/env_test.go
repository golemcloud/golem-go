package os_test

import (
	"fmt"
	"testing"

	golemos "github.com/golemcloud/golem-go/os"
)

func TestGetenvCompiles(t *testing.T) {
	_ = func() {
		var value string
		value = golemos.Getenv("ENV_VAR")
		fmt.Println(value)
	}
}

func TestLookupEnvCompiles(t *testing.T) {
	_ = func() {
		var value string
		var ok bool
		value, ok = golemos.LookupEnv("ENV_VAR")
		fmt.Printf("%s, %b\n", value, ok)
	}
}
