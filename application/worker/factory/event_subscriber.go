package factory

import (
	"github.com/arvinpaundra/sesen-api/application/worker/handler"
	authevent "github.com/arvinpaundra/sesen-api/domain/auth/event"
	infraevent "github.com/arvinpaundra/sesen-api/infrastructure/event"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// CreateDomainEventSubscriber creates and configures an Asynq-based domain event subscriber
// with all necessary domain event handlers registered
func CreateDomainEventSubscriber(db *gorm.DB, rdc *redis.Client) *infraevent.AsynqDomainEventSubscriber {
	subscriber := infraevent.NewAsynqDomainEventSubscriber(rdc)

	// Register all domain event handlers here
	// This keeps the wiring centralized and makes it easier to manage
	setupOverlay := handler.NewSetupOverlayHandler(db)

	userRegisteredComposite := handler.NewCompositeHandler(authevent.UserRegisteredEventType).
		AddHandler(setupOverlay)

	subscriber.RegisterService(authevent.UserRegisteredEventType, userRegisteredComposite)

	return subscriber
}
