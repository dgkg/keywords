package main

import (
	"keywords/app/handler"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"keywords/app/config"
)

func main() {
	router := gin.Default()

	router.GET("/health-check", handler.HealthCheck)
	conf := config.New()
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

	srv.ListenAndServe()
}
