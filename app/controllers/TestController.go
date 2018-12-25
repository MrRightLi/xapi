package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(context *gin.Context) {
	results := gin.H{
		"test": "test-key",
	}

	context.JSON(http.StatusOK, results)
}
