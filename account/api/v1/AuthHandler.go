package v1

import (
	"net/http"

	"github.com/Nebuska/neblab/account/api/v1/dto"
	"github.com/Nebuska/neblab/account/internal/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) Register(context *gin.Context, registerDTO dto.RegisterDTO) {
	err := handler.service.
		Register(context.Request.Context(), registerDTO.ToServiceRegisterData())
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Successfully registered"})
}

func (handler *AuthHandler) GetJwt(context *gin.Context, loginDTO dto.LoginDTO) {
	token, err := handler.service.
		OneTime(context.Request.Context(), loginDTO.ToServiceLoginData())
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.Header("Authorization", string("Bearer "+token))
	context.JSON(http.StatusOK, gin.H{
		"Jwt token": token,
	})
}
