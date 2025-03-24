package golemhost

import (
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"

	"github.com/golemcloud/golem-go/binding"
)

func NewUUID(bindingUUID binding.GolemApi1_1_6_HostUuid) uuid.UUID {
	var bs [16]byte
	binary.BigEndian.PutUint64(bs[:8], bindingUUID.HighBits)
	binary.BigEndian.PutUint64(bs[8:], bindingUUID.LowBits)

	goUUID, err := uuid.FromBytes(bs[:])
	if err != nil {
		panic(fmt.Sprintf("NewUUID: uuid.FromBytes failed: error: %s,bytes: %+v", err.Error(), bs))
	}
	return goUUID
}

func UUIDToBinding(goUUID uuid.UUID) binding.GolemApi1_1_6_HostUuid {
	return binding.GolemApi1_1_6_HostUuid{
		HighBits: binary.BigEndian.Uint64(goUUID[:8]),
		LowBits:  binary.BigEndian.Uint64(goUUID[8:]),
	}
}
