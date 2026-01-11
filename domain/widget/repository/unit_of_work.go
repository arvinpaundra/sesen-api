package repository

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	WidgetWriter() WidgetWriter

	Commit() error
	Rollback() error
}
