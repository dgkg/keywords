package main

import (
	"keywords/app/handler"
	"keywords/src/jwt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"keywords/app/config"
	logTime "keywords/src/log"
)

func main() {
	router := gin.Default()
	conf := config.New()

	router.GET("/health-check", logTime.MiddlewareLogTime(), jwt.MiddlewareJWT(conf.JWTSignKey), handler.HealthCheck)
	log.Println("Mode : ", conf.ModeEnv)

	srv := &http.Server{
		Addr:              ":" + conf.Port,
		Handler:           router,
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Second,
		ReadHeaderTimeout: time.Second,
		IdleTimeout:       time.Second,
		MaxHeaderBytes:    8 << 10,
	}

	log.Println(srv.ListenAndServe())
}
