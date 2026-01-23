package middleware

import (
	"net/http"
	"strings"

	"github.com/arvinpaundra/sesen-api/config"
	"github.com/arvinpaundra/sesen-api/core/format"
	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/domain/shared/entity"
	"github.com/arvinpaundra/sesen-api/domain/user/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/user/service"
	authinfra "github.com/arvinpaundra/sesen-api/infrastructure/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Authenticate struct {
	db *gorm.DB
}

func NewAuthenticate(db *gorm.DB) *Authenticate {
	return &Authenticate{db: db}
}

func (a *Authenticate) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		if bearerToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, format.Unauthorized("bearer token is missing"))
			return
		}

		accessToken := strings.Replace(bearerToken, "Bearer ", "", 1)

		svc := service.NewCheckSession(
			authinfra.NewUserReaderRepository(a.db),
			token.NewJWT(config.GetString("JWT_SECRET")),
		)

		payload := request.CheckSessionPayload{
			AccessToken: accessToken,
		}

		res, err := svc.Execute(c.Request.Context(), payload)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, format.Unauthorized(err.Error()))
			return
		}

		session := &entity.UserAuth{
			UserId:   res.UserId,
			Email:    res.Email,
			Username: res.Username,
			Fullname: res.Fullname,
		}

		c.Set("session", session)

		c.Next()
	}
}
