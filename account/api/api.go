package api

import "github.com/gin-gonic/gin"

func health(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "OK",
	})
}
