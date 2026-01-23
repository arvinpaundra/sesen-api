package donation

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/donation/repository"
	"github.com/arvinpaundra/sesen-api/infrastructure/shared"
	"gorm.io/gorm"
)

var _ repository.UnitOfWork = (*UnitOfWork)(nil)

type UnitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin(ctx context.Context) (repository.UnitOfWorkProcessor, error) {
	tx := u.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	txCtx := context.WithValue(ctx, shared.TxKey{}, tx)

	return &UnitOfWorkProcessor{tx: tx, ctx: txCtx}, nil
}

type UnitOfWorkProcessor struct {
	tx  *gorm.DB
	ctx context.Context
}

func (up *UnitOfWorkProcessor) DonationWriter() repository.DonationWriter {
	return NewDonationWriterRepository(up.tx)
}

func (up *UnitOfWorkProcessor) Context() context.Context {
	return up.ctx
}

func (up *UnitOfWorkProcessor) Commit() error {
	return up.tx.Commit().Error
}

func (up *UnitOfWorkProcessor) Rollback() error {
	return up.tx.Rollback().Error
}
