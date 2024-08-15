package golemhost

import (
	"encoding/binary"
	"fmt"

	"github.com/google/uuid"

	"github.com/golemcloud/golem-go/binding"
)

func newUUID(bindingUUID binding.GolemApi0_2_0_HostUuid) uuid.UUID {
	var bs [16]byte
	binary.BigEndian.PutUint64(bs[:8], bindingUUID.HighBits)
	binary.BigEndian.PutUint64(bs[8:], bindingUUID.LowBits)

	goUUID, err := uuid.FromBytes(bs[:])
	if err != nil {
		panic(fmt.Sprintf("newUUID: uuid.FromBytes failed: error: %s,bytes: %+v", err.Error(), bs))
	}
	return goUUID
}

func uuidToBinding(goUUID uuid.UUID) binding.GolemApi0_2_0_HostUuid {
	return binding.GolemApi0_2_0_HostUuid{
		HighBits: binary.BigEndian.Uint64(goUUID[:8]),
		LowBits:  binary.BigEndian.Uint64(goUUID[8:]),
	}
}
