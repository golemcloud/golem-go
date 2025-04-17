package golemhost

import (
	"github.com/golemcloud/golem-go/binding/golem/api/host"
	"github.com/google/uuid"
)

type ComponentID uuid.UUID

func NewComponentID(componentID host.ComponentID) ComponentID {
	return ComponentID(NewUUID(componentID.UUID))
}

func (componentID ComponentID) ToBinding() host.ComponentID {
	return host.ComponentID{
		UUID: UUIDToBinding(uuid.UUID(componentID)),
	}
}
