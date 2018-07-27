package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xapi/app/models"
)

func GetOrders(context *gin.Context) {
	results := models.OrderList()

	context.JSON(http.StatusOK, results)
}
