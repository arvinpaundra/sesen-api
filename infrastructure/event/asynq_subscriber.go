package event

import (
	"context"

	"github.com/arvinpaundra/sesen-api/core/event"
	"github.com/hibiken/asynq"
	"github.com/redis/go-redis/v9"
)

// AsynqDomainEventSubscriber subscribes to domain events using Asynq
type AsynqDomainEventSubscriber struct {
	server *asynq.Server
	mux    *asynq.ServeMux
}

func NewAsynqDomainEventSubscriber(rdc *redis.Client) *AsynqDomainEventSubscriber {
	opt := rdc.Options()

	redisOpt := asynq.RedisClientOpt{
		Addr:     opt.Addr,
		Password: opt.Password,
		DB:       opt.DB,
	}

	server := asynq.NewServer(redisOpt, asynq.Config{
		Concurrency: 10,
		Queues: map[string]int{
			"domain_events": 6,
			"default":       1,
		},
	})

	mux := asynq.NewServeMux()

	return &AsynqDomainEventSubscriber{
		server: server,
		mux:    mux,
	}
}

func (s *AsynqDomainEventSubscriber) RegisterService(eventType string, handler event.DomainEventHandler) {
	// Register the handler with Asynq mux
	s.mux.HandleFunc(eventType, func(ctx context.Context, task *asynq.Task) error {
		if err := handler.Handle(ctx, task.Payload()); err != nil {
			return err
		}

		return nil
	})
}

func (s *AsynqDomainEventSubscriber) Start(ctx context.Context) error {
	return s.server.Start(s.mux)
}

func (s *AsynqDomainEventSubscriber) Shutdown() {
	s.server.Shutdown()
}
