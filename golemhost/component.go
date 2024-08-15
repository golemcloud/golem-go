package golemhost

import (
	"github.com/google/uuid"

	"github.com/golemcloud/golem-go/binding"
)

type ComponentID uuid.UUID

func newComponentID(componentID binding.GolemApi0_2_0_HostComponentId) ComponentID {
	return ComponentID(newUUID(componentID.Uuid))
}

func (componentID ComponentID) toBinding() binding.GolemApi0_2_0_HostComponentId {
	return binding.GolemApi0_2_0_HostComponentId{
		Uuid: uuidToBinding(uuid.UUID(componentID)),
	}
}
