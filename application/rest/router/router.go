package router

import (
	"github.com/arvinpaundra/sesen-api/application/rest/middleware"
	"github.com/arvinpaundra/sesen-api/application/rest/router/auth"
	"github.com/arvinpaundra/sesen-api/core/validator"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Register(g *gin.Engine, rdb *redis.Client, db *gorm.DB) {
	g.Use(middleware.Cors())
	g.Use(gin.Recovery())
	g.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		SkipPaths: []string{"/metrics"},
	}))

	v1 := g.Group("/v1")

	authRouter := auth.NewAuthRouter(db, rdb, validator.NewValidator())

	// public routes
	authRouter.Public(v1)

	// private routes
	authRouter.Private(v1)

}
