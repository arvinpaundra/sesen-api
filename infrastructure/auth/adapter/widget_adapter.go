package adapter

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/auth/repository"
	"github.com/arvinpaundra/sesen-api/domain/widget/service"
	infra "github.com/arvinpaundra/sesen-api/infrastructure/widget"
	"gorm.io/gorm"
)

var _ repository.WidgetMapper = (*WidgetAdapter)(nil)

type WidgetAdapter struct {
	db *gorm.DB
}

func NewWidgetAdapter(db *gorm.DB) *WidgetAdapter {
	return &WidgetAdapter{db: db}
}

func (w *WidgetAdapter) CreateDefaultWidgets(ctx context.Context, userID, username string) error {
	svc := service.NewCreateDefaultWidgets(
		infra.NewWidgetReaderRepository(w.db),
		infra.NewWidgetWriterRepository(w.db),
		infra.NewUnitOfWork(w.db),
	)

	command := service.CreateDefaultWidgetsCommand{
		UserID:   userID,
		Username: username,
	}

	return svc.Execute(ctx, command)
}
