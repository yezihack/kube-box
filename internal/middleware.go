package internal

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 中间件

func RequestCount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		httpCode := ctx.Writer.Status()
		log.Println("httpCode:", httpCode)
		HTTPRequests.WithLabelValues(strconv.Itoa(httpCode)).Inc()
	}
}
