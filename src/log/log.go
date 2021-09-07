package log

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddlewareLogTime() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		ctx.Next()
		log.Println("time for response:", time.Since(start))
	}
}
