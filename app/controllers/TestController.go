package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const NSQ_HOST = "172.16.19.146"

func Test(context *gin.Context) {
	results := gin.H{
		"test": "test-key",
	}

	context.JSON(http.StatusOK, results)
}