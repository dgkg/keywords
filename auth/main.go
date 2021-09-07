package main

import (
	"bytes"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type Config struct {
	Port    string
	ModeEnv string
}

var config Config

func init() {
	viper.SetConfigType("yaml")
	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	config.Port = viper.GetString("port")
	config.ModeEnv = viper.GetString("mode")
}

func main() {
	router := gin.Default()
	router.POST("/login", Login)
	log.Println("Mode : ", config.ModeEnv)

	srv := &http.Server{
		Addr:              ":" + config.Port,
		Handler:           router,
		ReadTimeout:       time.Second,
		WriteTimeout:      time.Second,
		ReadHeaderTimeout: time.Second,
		IdleTimeout:       time.Second,
		MaxHeaderBytes:    8 << 10,
	}

	srv.ListenAndServe()
}

func Login(ctx *gin.Context) {
	var payload PayloadLogin
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
	ctx.JSON(http.StatusOK, gin.H{
		"jwt": "jwt_value",
	})
}

type PayloadLogin struct {
	User     string `json:"user"`
	Password string `json:"pass"`
}

// any approach to require this configuration into your program.
var yamlExample = []byte(`
mode: production
port: 8080
`)
