package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"keywords/app/model"
)

func HealthCheck(ctx *gin.Context) {
	authValue := ctx.GetHeader("Authorization")
	if len(authValue) == 0 {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": model.HealthcheckStatusRandom(), // ok, altered, down.
	})
}
