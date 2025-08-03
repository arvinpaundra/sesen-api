package overlay

import (
	"github.com/arvinpaundra/sesen-api/domain/overlay/repository"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db *gorm.DB
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{db: db}
}

func (u *UnitOfWork) Begin() (repository.UnitOfWorkProcessor, error) {
	tx := u.db.Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &UnitOfWorkProcessor{tx: tx}, nil
}

type UnitOfWorkProcessor struct {
	tx *gorm.DB
}

func (u *UnitOfWorkProcessor) OverlayWriter() repository.OverlayWriter {
	return NewOverlayWriterRepository(u.tx)
}

func (u *UnitOfWorkProcessor) Rollback() error {
	return u.tx.Rollback().Error
}

func (u *UnitOfWorkProcessor) Commit() error {
	return u.tx.Commit().Error
}
