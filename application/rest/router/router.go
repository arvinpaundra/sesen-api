package router

import (
	"github.com/arvinpaundra/sesen-api/application/rest/middleware"
	"github.com/arvinpaundra/sesen-api/application/rest/router/auth"
	"github.com/arvinpaundra/sesen-api/core/validator"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(g *gin.Engine, db *gorm.DB) {
	g.Use(middleware.Cors())
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	v1 := g.Group("/v1")

	authRouter := auth.NewAuthRouter(db, validator.NewValidator())

	// public routes
	authRouter.Public(v1)

	// private routes
	authRouter.Private(v1)

}
