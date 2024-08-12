package main

import (
	golemos "github.com/golemcloud/golem-go/os"
	"net/http"

	golemhttp "github.com/golemcloud/golem-go/net/http"
)

// Test app for testing if the API compiles

func main() {
	// net/http

	{
		http.DefaultClient.Transport = &golemhttp.WasiHttpTransport{}
	}

	// os - args

	{
		var args []string
		args = golemos.GetArgs()
		unused(args)
	}

	// os - env

	{
		var value string
		value = golemos.Getenv("ENV_VAR")
		unused(value)
	}

	{
		var value string
		var ok bool
		value, ok = golemos.LookupEnv("ENV_VAR")
		unused(value)
		unused(ok)
	}
}

func unused[T any](_ T) {}
