package event

import (
	"context"
	"encoding/json"
	"time"

	"github.com/arvinpaundra/sesen-api/core/util"
)

// DomainEvent represents a domain event that occurred in the system
type DomainEvent interface {
	// GetEventID returns the unique ID of the event
	GetEventID() string
	// GetEventType returns the type/name of the event
	GetEventType() string
	// GetOccurredAt returns when the event occurred
	GetOccurredAt() time.Time
	// ToBytes serializes the event to bytes for queue storage
	ToBytes() ([]byte, error)
}

// BaseDomainEvent provides a base implementation for domain events
type BaseDomainEvent struct {
	EventID    string    `json:"event_id"`
	EventType  string    `json:"event_type"`
	OccurredAt time.Time `json:"occurred_at"`
}

func NewBaseDomainEvent(eventType string) *BaseDomainEvent {
	return &BaseDomainEvent{
		EventID:    util.GenerateUUID(),
		EventType:  eventType,
		OccurredAt: time.Now().UTC(),
	}
}

func (e *BaseDomainEvent) GetEventID() string {
	return e.EventID
}

func (e *BaseDomainEvent) GetEventType() string {
	return e.EventType
}

func (e *BaseDomainEvent) GetOccurredAt() time.Time {
	return e.OccurredAt
}

func (e *BaseDomainEvent) ToBytes() ([]byte, error) {
	return json.Marshal(e)
}

// DomainEventPublisher publishes domain events to the queue
type DomainEventPublisher interface {
	Publish(ctx context.Context, events ...DomainEvent) error
}

// DomainEventHandler handles domain events
type DomainEventHandler interface {
	Handle(ctx context.Context, eventData []byte) error
}
