package main

import (
	"bytes"
	"errors"
	"log"
	"math/rand"
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
	router.GET("/health-check", HealthCheck)
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

func HealthCheck(ctx *gin.Context) {
	authValue := ctx.GetHeader("Authorization")
	if len(authValue) == 0 {
		log.Println(errors.New("user not authorized"))
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status": HealthcheckStatusRandom(), // ok, altered, down.
	})
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

type StatusHealthcheck uint8

func (h StatusHealthcheck) String() string {
	return healthcheckTbl[h]
}

const (
	StatusOK StatusHealthcheck = iota + 1
	StatusAltered
	StatusDown
)

var healthcheckTbl = [...]string{
	0:             "unknown",
	StatusOK:      "ok",
	StatusAltered: "altered",
	StatusDown:    "down",
}

func HealthcheckStatusRandom() string {
	res := rand.Intn(len(healthcheckTbl))
	if res == 0 {
		res++
	}
	return healthcheckTbl[res]
}

// any approach to require this configuration into your program.
var yamlExample = []byte(`
mode: production
port: 8081
`)
