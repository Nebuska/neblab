package middlewares

import (
	"errors"
	"net/http"
	"task-tracker/pkg/appError"
	"task-tracker/pkg/appError/errorCodes"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) == 0 {
			return
		}
		//todo handle the error
		requestLogger := c.MustGet("logger").(zerolog.Logger)
		var err appError.AppError
		if errors.As(c.Errors.Last().Err, &err) {
			switch err.ErrorCode {
			case errorCodes.Undefined:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusInternalServerError, gin.H{"message": "unexpected error"})
			case errorCodes.Forbidden:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			case errorCodes.NotFound:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
			case errorCodes.DataValidationError:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusBadRequest, gin.H{"message": "data validation error"})
			case errorCodes.ConflictingData:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusConflict, gin.H{"message": "data conflicting error"})
			case errorCodes.InternalError:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusInternalServerError, gin.H{"message": "internal server error"})
			case errorCodes.BadRequest:
				requestLogger.Warn().Err(err).Msg("")
				c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
			}
		}
	}
}
