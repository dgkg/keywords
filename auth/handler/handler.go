package handler

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/keywords/auth/handler/model"
	"github.com/dgkg/keywords/db"
	"github.com/dgkg/keywords/src/jwt"
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
	start := time.Now()
	res, err := s.db.GetUserByLogin(payload.Login)
	if err != nil {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	since := time.Since(start)
	log.Println("db action took:", since)

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
