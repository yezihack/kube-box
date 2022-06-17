package main

import (
	"fmt"
	"log"
	"time"

	"github.com/yezihack/kube-box/internal"

	"github.com/gin-gonic/gin"
)

var (
	version = internal.DefaultVersion
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	r := gin.Default()
	internal.AddRoute(r)
	// r.Use(gin.Logger())
	// LoggerWithFormatter middleware will write the logs to gin.DefaultWriter
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())

	port := internal.GetEnvValueToInteger(internal.EnvPort, internal.DefaultPort)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln("Server is failed, err:", err)
	}
}
