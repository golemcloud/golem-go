package os

import (
	"os"

	"github.com/golemcloud/golem-go/binding"
)

// InitStdArgs overrides standard lib args with the ones coming from the WASI environment
func InitStdArgs() {
	os.Args = GetArgs()
}

// GetArgs is convenience wrapper to be used instead of the standard lib's os.Args
func GetArgs() []string {
	return binding.WasiCli0_2_0_EnvironmentGetArguments()
}
