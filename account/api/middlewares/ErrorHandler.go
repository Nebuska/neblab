package middlewares

import (
	"errors"
	"net/http"

	apiDto "github.com/Nebuska/neblab/account/api/v1/dto"
	authDto "github.com/Nebuska/neblab/account/internal/auth/dto"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			return
		}
		requestLogger := c.MustGet("logger").(zerolog.Logger)

		err := c.Errors.Last()
		switch {
		case errors.Is(err, apiDto.ErrBodyBindingFailOnRegister):
			requestLogger.Warn().Err(err).Send()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Request's body is faulty",
			})
		case errors.Is(err, apiDto.ErrBodyBindingFailOnLogin):
			requestLogger.Warn().Err(err).Send()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Request's body is faulty",
			})
		case errors.Is(err, authDto.ErrEmailAlreadyExists):
			requestLogger.Warn().Err(err).Send()
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Email already exists",
			})
		default:
			requestLogger.Error().Err(err).Msg("!!! NOT EXPECTED ERROR !!!")
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal server error",
			})
		}
	}
}
