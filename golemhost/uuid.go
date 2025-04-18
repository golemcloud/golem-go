package golemhost

import (
	"encoding/binary"
	"fmt"

	"github.com/golemcloud/golem-go/binding/golem/api/host"
	"github.com/google/uuid"
)

func NewUUID(bindingUUID host.UUID) uuid.UUID {
	var bs [16]byte
	binary.BigEndian.PutUint64(bs[:8], bindingUUID.HighBits)
	binary.BigEndian.PutUint64(bs[8:], bindingUUID.LowBits)

	goUUID, err := uuid.FromBytes(bs[:])
	if err != nil {
		panic(fmt.Sprintf("NewUUID: uuid.FromBytes failed: error: %s,bytes: %+v", err.Error(), bs))
	}
	return goUUID
}

func UUIDToBinding(goUUID uuid.UUID) host.UUID {
	return host.UUID{
		HighBits: binary.BigEndian.Uint64(goUUID[:8]),
		LowBits:  binary.BigEndian.Uint64(goUUID[8:]),
	}
}
