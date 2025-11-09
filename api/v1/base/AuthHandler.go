package base

import (
	"net/http"
	"task-tracker/api/v1/base/dto"
	"task-tracker/internal/Auth"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	service Auth.Service
}

func NewAuthHandler(service Auth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Register(context *gin.Context) {
	var registerData dto.RegisterDTO
	if err := context.ShouldBind(&registerData); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	err := validator.New().Struct(&registerData)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := handler.service.Register(
		registerData.Username,
		registerData.Email,
		registerData.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
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
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request body",
		})
		return
	}
	token, err := handler.service.Login(LoginData.Username, LoginData.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	context.Header("Authorization", string("Bearer "+token))
	context.JSON(http.StatusOK, gin.H{
		"Jwt token": token,
	})

}
