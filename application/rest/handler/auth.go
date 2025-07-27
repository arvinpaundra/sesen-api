package handler

import (
	"net/http"

	"github.com/arvinpaundra/sesen-api/core/format"
	"github.com/arvinpaundra/sesen-api/core/token"
	"github.com/arvinpaundra/sesen-api/core/validator"
	"github.com/arvinpaundra/sesen-api/domain/auth/constant"
	"github.com/arvinpaundra/sesen-api/domain/auth/dto/request"
	"github.com/arvinpaundra/sesen-api/domain/auth/service"
	"github.com/arvinpaundra/sesen-api/infrastructure/auth"
	"github.com/arvinpaundra/sesen-api/infrastructure/shared"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type AuthHandler struct {
	db        *gorm.DB
	validator *validator.Validator
}

func NewAuthHandler(db *gorm.DB, validator *validator.Validator) *AuthHandler {
	return &AuthHandler{
		db:        db,
		validator: validator,
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var payload request.UserRegister

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
		return
	}

	verrs := h.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	svc := service.NewUserRegister(
		auth.NewUserReaderRepository(h.db),
		auth.NewUserWriterRepository(h.db),
	)

	err = svc.Execute(c.Request.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrEmailAlreadyExists, constant.ErrUsernameAlreadyExists:
			c.JSON(http.StatusConflict, format.Conflict(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
		}
		return
	}

	c.JSON(http.StatusCreated, format.SuccessCreated("user registered successfully", nil))
}

func (h *AuthHandler) Login(c *gin.Context) {
	var payload request.UserLogin

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
		return
	}

	verrs := h.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	svc := service.NewUserLogin(
		auth.NewUserReaderRepository(h.db),
		auth.NewUserWriterRepository(h.db),
		token.NewJWT(viper.GetString("JWT_SECRET")),
		auth.NewUnitOfWork(h.db),
	)

	result, err := svc.Execute(c.Request.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrUserNotFound:
			c.JSON(http.StatusNotFound, format.NotFound(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
		}
		return
	}

	c.JSON(http.StatusOK, format.SuccessOK("user logged in successfully", result))
}

func (h *AuthHandler) Logout(c *gin.Context) {
	var payload request.UserLogout

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
		return
	}

	verrs := h.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	svc := service.NewUserLogout(
		auth.NewUserReaderRepository(h.db),
		auth.NewUserWriterRepository(h.db),
		token.NewJWT(viper.GetString("JWT_SECRET")),
		shared.NewAuthStorage(c),
		auth.NewUnitOfWork(h.db),
	)

	err = svc.Execute(c.Request.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrUserNotFound:
			c.JSON(http.StatusNotFound, format.NotFound(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
		}
		return
	}

	c.JSON(http.StatusOK, format.SuccessOK("user logged out successfully", nil))
}

func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var payload request.RefreshToken

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, format.UnprocessableEntity(err.Error()))
		return
	}

	verrs := h.validator.Validate(payload)
	if verrs != nil {
		c.JSON(http.StatusBadRequest, format.BadRequest("invalid request body", verrs))
		return
	}

	svc := service.NewRefreshToken(
		auth.NewUserReaderRepository(h.db),
		auth.NewUserWriterRepository(h.db),
		token.NewJWT(viper.GetString("JWT_SECRET")),
		auth.NewUnitOfWork(h.db),
	)

	result, err := svc.Execute(c.Request.Context(), payload)
	if err != nil {
		switch err {
		case constant.ErrInvalidRefreshToken:
			c.JSON(http.StatusUnauthorized, format.Unauthorized(err.Error()))
		default:
			c.JSON(http.StatusInternalServerError, format.InternalServerError())
		}
		return
	}

	c.JSON(http.StatusOK, format.SuccessOK("token refreshed successfully", result))
}
