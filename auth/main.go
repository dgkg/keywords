package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/dgkg/keywords/auth/config"
	"github.com/dgkg/keywords/auth/handler"
	"github.com/dgkg/keywords/db"
	"github.com/dgkg/keywords/db/bolt"
	"github.com/dgkg/keywords/db/moke"
)

const (
	EnvProd = "production"
	EnvTest = "testing"
)

func main() {

	// get config.
	config := config.New()
	log.Println("Mode : ", config.ModeEnv)

	// create router.
	router := gin.Default()

	log.Println("creating db")
	var db db.Storer
	// create db connection.
	if config.ModeEnv == EnvTest {
		log.Println("create moke DB.")
		db = moke.New()
	} else {
		log.Println("create Bolt DB.")
		db = bolt.New(config.DBName)
	}
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
