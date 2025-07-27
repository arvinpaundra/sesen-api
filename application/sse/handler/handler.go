package handler

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Handler struct {
	db     *gorm.DB
	logger *zap.Logger
}

func NewHandler(db *gorm.DB, logger *zap.Logger) *Handler {
	return &Handler{
		db:     db,
		logger: logger,
	}
}
