package repository

import (
	"context"
)

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	Commit(ctx context.Context) error
	Rollback(ctx context.Context) error
}
