package main

import (
	stdhttp "net/http"
	"time"

	"github.com/golemcloud/golem-go/golemhost"
	"github.com/golemcloud/golem-go/net/http"
	"github.com/golemcloud/golem-go/os"
	"github.com/golemcloud/golem-go/std"
)

// Test app for testing if the API compiles with the right types
// (as tinygo build expects a main module, we cannot simply build the lib)

func main() {
	// net/http

	{
		stdhttp.DefaultClient.Transport = &http.WasiHttpTransport{}
	}

	// net/http
	{
		http.InitStdDefaultClientTransport()
	}

	// os - args

	{
		var args []string
		args = os.GetArgs()
		unused(args)
	}

	// os - env

	{
		var value string
		value = os.Getenv("ENV_VAR")
		unused(value)
	}

	{
		var value string
		var ok bool
		value, ok = os.LookupEnv("ENV_VAR")
		unused(value)
		unused(ok)
	}

	// os - init
	{
		os.InitStdEnv()
		os.InitStdArgs()
		os.InitStd()
	}

	// golemhost - idempotence

	{
		var mode bool
		mode = golemhost.GetIdempotenceMode()
		unused(mode)
	}

	{
		golemhost.SetIdempotenceMode(false)
	}

	{
		var result []string
		var err error
		result, err = golemhost.WithIdempotenceMode(true, func() ([]string, error) {
			return []string{"golem"}, nil
		})
		unused(result)
		unused(err)
	}

	// golemhost persistence

	{
		golemhost.SetPersistenceLevel(golemhost.PersistenceLevelSmart)
	}

	{
		var level golemhost.PersistenceLevel
		level = golemhost.GetPersistenceLevel()
		unused(level)
	}

	{
		var result map[string]int
		var err error
		result, err = golemhost.WithPersistenceLevel(
			golemhost.PersistenceLevelPersistRemoteSideEffects,
			func() (map[string]int, error) {
				return nil, nil
			},
		)
		unused(result)
		unused(err)
	}

	// golemhost retrypolicy

	{
		golemhost.SetRetryPolicy(golemhost.RetryPolicy{
			MaxAttempts: 10,
			MinDelay:    100 * time.Millisecond,
			MaxDelay:    5 * time.Nanosecond,
			Multiplier:  3,
		})
	}

	{
		var result golemhost.RetryPolicy
		result = golemhost.GetRetryPolicy()
		unused(result)
	}

	{
		var result string
		var err error
		result, err = golemhost.WithRetryPolicy(golemhost.RetryPolicy{
			MaxAttempts: 4,
			MinDelay:    10 * time.Microsecond,
			MaxDelay:    4 * time.Minute,
			Multiplier:  2,
		}, func() (string, error) {
			return "golem", nil
		})
		unused(result)
		unused(err)
	}

	// std init

	{
		std.Init(std.Modules{
			Os:   true,
			Http: true,
		})
	}
}

func unused[T any](_ T) {}
