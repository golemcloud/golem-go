package std

import (
	"github.com/golemcloud/golem-go/net/http"
	"github.com/golemcloud/golem-go/os"
)

type Packages struct {
	Os      bool
	NetHttp bool
}

// Init optionally initializes standard lib's packages with the WASI environment and wrappers
func Init(modules Packages) {
	if modules.NetHttp {
		http.InitStd()
	}

	if modules.Os {
		os.InitStd()
	}
}
