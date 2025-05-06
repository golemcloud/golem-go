package golemhost

import "github.com/golemcloud/golem-go/binding/golem/api/host"

func Fork(newName string) host.ForkResult {
	return host.Fork(newName)
}
