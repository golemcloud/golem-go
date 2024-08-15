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
func Init(packages Packages) {
	if packages.NetHttp {
		http.InitStd()
	}

	if packages.Os {
		os.InitStd()
	}
}
