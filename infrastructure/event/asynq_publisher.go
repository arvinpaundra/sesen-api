package event

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/event"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

// AsynqDomainEventPublisher publishes domain events using Asynq
type AsynqDomainEventPublisher struct {
	client *asynq.Client
}

func NewAsynqDomainEventPublisher(rdc *redis.Client) *AsynqDomainEventPublisher {
	opt := rdc.Options()

	redisOpt := asynq.RedisClientOpt{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	}

	client := asynq.NewClient(redisOpt)

	return &AsynqDomainEventPublisher{
		client: client,
	}
}

func (p *AsynqDomainEventPublisher) Publish(ctx context.Context, events ...event.DomainEvent) error {
	for _, event := range events {
		payload, err := event.ToBytes()
		if err != nil {
			continue
		}

		task := asynq.NewTask(event.GetEventType(), payload)

		// Enqueue the task with retry options
		_, err = p.client.EnqueueContext(ctx, task,
			asynq.MaxRetry(3),
			asynq.Queue("domain_events"),
		)

		if err != nil {
			return err
		}
	}

	return nil
}

func (p *AsynqDomainEventPublisher) Close() error {
	return p.client.Close()
}
