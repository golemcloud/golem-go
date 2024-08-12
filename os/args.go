package os

import (
	binding "github.com/golemcloud/golem-go/golem_go_bindings"
)

// GetArgs is convenience wrapper to be used instead of the standard lib's os.Args
func GetArgs() []string {
	return binding.WasiCli0_2_0_EnvironmentGetArguments()
}
