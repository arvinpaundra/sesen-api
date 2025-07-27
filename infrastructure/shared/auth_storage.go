package shared

import (
	"github.com/arvinpaundra/sesen-api/domain/shared/entity"
	"github.com/gin-gonic/gin"
)

func NewAuthStorage(c *gin.Context) *entity.UserAuth {
	auth, ok := c.Value("session").(*entity.UserAuth)
	if !ok {
		return nil
	}

	return auth
}
