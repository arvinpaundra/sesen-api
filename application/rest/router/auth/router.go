package auth

import (
	"github.com/arvinpaundra/sesen-api/application/rest/handler"
	"github.com/arvinpaundra/sesen-api/application/rest/middleware"
	"github.com/arvinpaundra/sesen-api/core/validator"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type AuthRouter struct {
	db  *gorm.DB
	rdb *redis.Client
	vld *validator.Validator
}

func NewAuthRouter(
	db *gorm.DB,
	rdb *redis.Client,
	vld *validator.Validator,
) *AuthRouter {
	return &AuthRouter{
		db:  db,
		rdb: rdb,
		vld: vld,
	}
}

func (r *AuthRouter) Public(g *gin.RouterGroup) {
	h := handler.NewAuthHandler(r.db, r.rdb, r.vld)

	auth := g.Group("/auth")
	{
		auth.POST("/login", h.Login)
		auth.POST("/register", h.Register)
		auth.POST("/refresh-token", h.RefreshToken)
	}
}

func (r *AuthRouter) Private(g *gin.RouterGroup) {
	h := handler.NewAuthHandler(r.db, r.rdb, r.vld)
	m := middleware.NewAuthenticate(r.db)

	auth := g.Group("/me", m.Authenticate())
	{
		auth.POST("/logout", h.Logout)
	}
}
