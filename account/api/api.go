package api

import "github.com/gin-gonic/gin"

func health(context *gin.Context) {
	context.JSON(200, gin.H{
		"status": "OK",
	})
}

func ready(context *gin.Context) {
	// todo : migration db and redis check
	context.JSON(200, gin.H{
		"status": "OK",
	})
}

func version(context *gin.Context) {
	// todo : should be inserted from another way
	context.JSON(200, gin.H{
		"version": "v1.0 alpha",
	})
}
