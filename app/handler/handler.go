package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"keywords/app/model"
)

func HealthCheck(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"status": model.HealthcheckStatusRandom(), // ok, altered, down.
	})
}

func DisplayStats(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"stats": model.HealthcheckStatusRandom(), // ok, altered, down.
	})
}

type Stats struct {
	BytesReceived uint64
}
