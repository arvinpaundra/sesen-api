package trait

import "github.com/arvinpaundra/sesen-api/core/event"

// EventSource provides functionality for entities to store and manage domain events
type EventSource struct {
	domainEvents []event.DomainEvent
}

// AddDomainEvent adds a domain event to the entity
func (e *EventSource) AddDomainEvent(domainEvent event.DomainEvent) {
	e.domainEvents = append(e.domainEvents, domainEvent)
}

// GetDomainEvents returns all domain events for this entity
func (e *EventSource) GetDomainEvents() []event.DomainEvent {
	return e.domainEvents
}

// ClearDomainEvents clears all domain events from the entity
func (e *EventSource) ClearDomainEvents() {
	e.domainEvents = make([]event.DomainEvent, 0)
}

// HasDomainEvents returns true if the entity has any domain events
func (e *EventSource) HasDomainEvents() bool {
	return len(e.domainEvents) > 0
}
