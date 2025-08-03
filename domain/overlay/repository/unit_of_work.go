package repository

type UnitOfWork interface {
	Begin() (UnitOfWorkProcessor, error)
}

type UnitOfWorkProcessor interface {
	OverlayWriter() OverlayWriter
	Commit() error
	Rollback() error
}
