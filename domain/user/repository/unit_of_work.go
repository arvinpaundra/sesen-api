package repository

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	UserWriter() UserWriter

	Context() context.Context
	Commit() error
	Rollback() error
}
