package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"keywords/auth/config"
	"keywords/auth/handler"
	"keywords/db/bolt"
)

func main() {
	// get config.
	config := config.New()
	log.Println("Mode : ", config.ModeEnv)

	// create router.
	router := gin.Default()

	log.Println("creating db")
	// create db connection.
	db := bolt.New(config.DBName, 1)
	log.Println("success create db")

	// create routes.
	service := handler.New(db)
	log.Println("create handler:")
	router.POST("/login", service.Login)

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

	log.Println("listend PORT:", srv.Addr)
	// run server.
	srv.ListenAndServe()
}
