package golemhost

import (
	"github.com/google/uuid"

	"github.com/golemcloud/golem-go/binding"
)

type ComponentID uuid.UUID

func NewComponentID(componentID binding.GolemApi1_1_6_HostComponentId) ComponentID {
	return ComponentID(NewUUID(componentID.Uuid))
}

func (componentID ComponentID) ToBinding() binding.GolemApi1_1_6_HostComponentId {
	return binding.GolemApi1_1_6_HostComponentId{
		Uuid: UUIDToBinding(uuid.UUID(componentID)),
	}
}
