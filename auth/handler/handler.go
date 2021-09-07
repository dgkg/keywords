package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"keywords/auth/handler/model"
	"keywords/src/jwt"
)

func Login(ctx *gin.Context) {
	var payload model.PayloadLogin
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	if payload.User != "root" || payload.Password != "password" {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := jwt.New("c4209493-df64-45fc-9cc1-ae845e820e29", "1", "Rob Pike")
	if err != nil {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwtValue,
	})
}
