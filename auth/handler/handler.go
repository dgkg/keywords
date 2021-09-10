package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"keywords/auth/handler/model"
	"keywords/db"
	"keywords/src/jwt"
)

type Service struct {
	db db.Storer
}

func New(db db.Storer) *Service {
	return &Service{
		db: db,
	}
}

func (s *Service) Login(ctx *gin.Context) {
	var payload model.PayloadLogin
	err := ctx.BindJSON(&payload)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	res, err := s.db.GetUserByLogin(payload.Login)
	if err != nil {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if res.Password != payload.Password {
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
