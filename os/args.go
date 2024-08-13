package os

import (
	"github.com/golemcloud/golem-go/binding"
)

// GetArgs is convenience wrapper to be used instead of the standard lib's os.Args
func GetArgs() []string {
	return binding.WasiCli0_2_0_EnvironmentGetArguments()
}
