package base

import (
	"github.com/Nebuska/task-tracker/api/v1/base/dto"
	"github.com/Nebuska/task-tracker/internal/aAuth"
	"github.com/Nebuska/task-tracker/pkg/appError"
	"github.com/Nebuska/task-tracker/pkg/appError/errorCodes"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service aAuth.Service
}

func NewAuthHandler(service aAuth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerData dto.RegisterDTO
	if err := context.ShouldBind(&registerData); err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "AuthHandler", err.Error()))
		return
	}

	token, err := handler.service.Register(
		registerData.Username,
		registerData.Email,
		registerData.Password)
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.Header("Authorization", string("Bearer "+token))
	context.JSON(http.StatusOK, gin.H{
		"Jwt token": token,
	})
}

func (handler *AuthHandler) Login(context *gin.Context) {
	var LoginData dto.LoginDTO
	if err := context.ShouldBind(&LoginData); err != nil {
		_ = context.Error(appError.New(errorCodes.BadRequest, "AuthHandler", err.Error()))
		return
	}
	token, err := handler.service.Login(LoginData.Username, LoginData.Password)
	if err != nil {
		_ = context.Error(err)
	}
	context.Header("Authorization", string("Bearer "+token))
	context.JSON(http.StatusOK, gin.H{
		"Jwt token": token,
	})

}
