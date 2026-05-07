package v1

import (
	"net/http"

	"github.com/Nebuska/neblab/account/api/v1/dto"
	"github.com/Nebuska/neblab/account/internal/auth"
	"github.com/Nebuska/neblab/account/internal/session"
	mw "github.com/Nebuska/neblab/tasker/api/middlewares"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service auth.Service
}

func NewAuthHandler(service auth.Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (handler *AuthHandler) RegisterRoute(router *gin.RouterGroup) {
	router.POST("/register", mw.WithBody(handler.Register))
	router.POST("/login", mw.WithBody(handler.Login))
	router.POST("/logout", handler.Logout)
	router.POST("/refresh", handler.Refresh)
	router.GET("/get-jwt", handler.GetJwt)
	router.POST("/one-time-jwt", mw.WithBody(handler.OneTimeJwt))
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

func (handler *AuthHandler) Login(context *gin.Context, loginDTO dto.LoginDTO) {
	token, err := handler.service.
		Login(context.Request.Context(), loginDTO.ToServiceLoginData())
	if err != nil {
		_ = context.Error(err)
		return
	}
	cookie := dto.NewSessionCookie(token.String())
	context.SetCookieData(cookie)
	context.Redirect(http.StatusContinue, "/v1/get-jwt")
}

func (handler *AuthHandler) Logout(context *gin.Context) {
	token, err := context.Cookie(dto.SESSION_COOKIE_NAME)
	if err != nil {
		_ = context.Error(err)
		return
	}
	err = handler.service.Logout(context, session.Tokenize(token))
	if err != nil {
		_ = context.Error(err)
		return
	}
	context.SetCookieData(dto.NewSessionDeletionToken())
	context.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})

}

func (handler *AuthHandler) Refresh(context *gin.Context) {
	token, err := context.Cookie(dto.SESSION_COOKIE_NAME)
	if err != nil {
		_ = context.Error(err)
		return
	}
	refresh, err := handler.service.Refresh(context, session.Tokenize(token))
	if err != nil {
		_ = context.Error(err)
		return
	}
	cookie := dto.NewSessionCookie(refresh.String())
	context.SetCookieData(cookie)
	context.Redirect(http.StatusFound, "/v1/get-jwt")
}

func (handler *AuthHandler) GetJwt(ctx *gin.Context) {
	token, err := ctx.Cookie(dto.SESSION_COOKIE_NAME)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	jwt, err := handler.service.GetJwt(ctx, session.Tokenize(token))
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	ctx.Header("Authorization", string("Bearer "+jwt))
	ctx.JSON(http.StatusOK, gin.H{"Jwt token": token})
}

func (handler *AuthHandler) OneTimeJwt(context *gin.Context, loginDTO dto.LoginDTO) {
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
