package os_test

import (
	"fmt"
	"testing"

	golemos "github.com/golemcloud/golem-go/os"
)

func TestGetArgsCompiles(t *testing.T) {
	_ = func() {
		var args []string
		args = golemos.GetArgs()
		fmt.Printf("%+v\n", args)
	}
}
