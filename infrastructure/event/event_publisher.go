package event

import (
	"github.com/arvinpaundra/sesen-api/core/event"
	"github.com/redis/go-redis/v9"
)

// CreateDomainEventPublisher creates and configures an Asynq-based domain event publisher
func CreateDomainEventPublisher(rdc *redis.Client) event.DomainEventPublisher {
	return NewAsynqDomainEventPublisher(rdc)
}
