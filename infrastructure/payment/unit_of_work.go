package payment

import (
	"context"

	"github.com/arvinpaundra/sesen-api/domain/payment/repository"
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
	if u.hasTx(ctx) {
		tx := u.getTx(ctx)
		return &UnitOfWorkProcessor{tx: tx, isNested: true}, nil
	}

	tx := u.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &UnitOfWorkProcessor{tx: tx, isNested: false}, nil
}

type UnitOfWorkProcessor struct {
	tx       *gorm.DB
	isNested bool
}

func (up *UnitOfWorkProcessor) Commit(ctx context.Context) error {
	if up.isNested {
		return nil
	}

	return up.tx.Commit().Error
}

func (up *UnitOfWorkProcessor) Rollback(ctx context.Context) error {
	if up.isNested {
		return nil
	}

	return up.tx.Rollback().Error
}

func (u *UnitOfWork) getTx(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(shared.TxKey{}).(*gorm.DB); ok {
		return tx
	}
	return u.db
}

func (u *UnitOfWork) hasTx(ctx context.Context) bool {
	tx, ok := ctx.Value(shared.TxKey{}).(*gorm.DB)
	return ok && tx != nil
}
