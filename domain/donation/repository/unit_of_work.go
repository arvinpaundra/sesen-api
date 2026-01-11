package repository

import "context"

type UnitOfWork interface {
	Begin(ctx context.Context) (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	DonationWriter() DonationWriter

	Context() context.Context
	Commit() error
	Rollback() error
}
