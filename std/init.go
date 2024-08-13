package std

import (
	"github.com/golemcloud/golem-go/net/http"
	"github.com/golemcloud/golem-go/os"
)

type Modules struct {
	Os   bool
	Http bool
}

// Init optionally initializes standard lib's modules with the WASI environment and wrappers
func Init(modules Modules) {
	if modules.Http {
		http.InitStd()
	}

	if modules.Os {
		os.InitStd()
	}
}
