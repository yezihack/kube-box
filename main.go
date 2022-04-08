package main

import (
	"fmt"
	"log"

	"github.com/yezihack/kube-box/internal"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	r := gin.Default()
	internal.AddRoute(r)

	port := internal.GetEnvValueToInteger(internal.EnvPort, internal.DefaultPort)
	if err := r.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatalln("Server is failed, err:", err)
	}
}
