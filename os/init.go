package os

// InitStd overrides standard lib environment with the WASI values
func InitStd() {
	InitStdArgs()
	InitStdEnv()
}
