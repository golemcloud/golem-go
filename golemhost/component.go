package golemhost

import (
	"github.com/google/uuid"

	"github.com/golemcloud/golem-go/binding"
)

type ComponentID uuid.UUID

func NewComponentID(componentID binding.GolemApi0_2_0_HostComponentId) ComponentID {
	return ComponentID(NewUUID(componentID.Uuid))
}

func (componentID ComponentID) ToBinding() binding.GolemApi0_2_0_HostComponentId {
	return binding.GolemApi0_2_0_HostComponentId{
		Uuid: UUIDToBinding(uuid.UUID(componentID)),
	}
}
