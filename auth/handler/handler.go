package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"keywords/auth/handler/model"
	"keywords/src/jwt"
)

var dataUser = map[string]model.User{
	"casper": {
		ID:          "1bd7f4b5-2249-4206-8d82-86ae7ccd59ee",
		Name:        "Casper",
		Password:    "Tatata",
		AccessLevel: 1,
	},
	"boss": {
		ID:          "f791ed48-64f8-4dc6-8507-845e149286f2",
		Name:        "The Boss",
		Password:    "bibibi",
		AccessLevel: 10,
	},
}

func Login(ctx *gin.Context) {
	var payload model.PayloadLogin
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, ok := dataUser[payload.Login]
	if !ok || res.Password != payload.Password {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	jwtValue, err := jwt.New(res.AccessLevel, res.ID, res.Name)
	if err != nil {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"jwt": jwtValue,
	})
}
