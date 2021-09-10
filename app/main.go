package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/keywords/app/config"
	"github.com/dgkg/keywords/app/handler"
	"github.com/dgkg/keywords/src/jwt"
	logTime "github.com/dgkg/keywords/src/log"
)

func main() {
	router := gin.Default()
	conf := config.New()

	router.GET("/health-check", logTime.MiddlewareLogTime(), jwt.MiddlewareJWT(conf.JWTSignKey, 1), handler.HealthCheck)
	router.GET("/stats", jwt.MiddlewareJWT(conf.JWTSignKey, 3), handler.DisplayStats)
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

	log.Println("listend PORT:", srv.Addr)

	log.Println(srv.ListenAndServe())
}
