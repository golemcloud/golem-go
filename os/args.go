package os

import (
	golem "github.com/golemcloud/golem-go/golem_go_bindings"
)

// GetArgs is convenience wrapper to be used instead of the standard lib's os.Args
func GetArgs() []string {
	return golem.WasiCli0_2_0_EnvironmentGetArguments()
}
