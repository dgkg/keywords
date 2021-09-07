package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"keywords/auth/config"
	"keywords/auth/handler"
)

func main() {
	// get config.
	config := config.New()
	log.Println("Mode : ", config.ModeEnv)

	// create router.
	router := gin.Default()

	// create routes.
	router.POST("/login", handler.Login)

	// init server.
	srv := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           router,
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Second,
		ReadHeaderTimeout: time.Second,
		IdleTimeout:       time.Second,
		MaxHeaderBytes:    8 << 10,
	}
	// run server.
	srv.ListenAndServe()
}
