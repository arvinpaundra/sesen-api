package router

import (
	"github.com/arvinpaundra/sesen-api/application/sse/handler"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Register(g *gin.Engine, db *gorm.DB, logger *zap.Logger) {
	handler := handler.NewHandler(db, logger)

	g.GET("/message", handler.ShowDonationMessageHandler)
}
