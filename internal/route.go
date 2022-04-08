package internal

import (
	"github.com/gin-gonic/gin"
)

func AddRoute(r *gin.Engine) {
	r.GET("/", Home)
	r.GET("/ping", Ping)
	r.GET("/healthz", Healthz)
	r.GET("/check-ip", CheckIP)
	r.GET("/dry-check-ip", DryCheckIP)
	r.GET("/check-healthz", CheckHealthz)
	r.GET("/dry-check-healthz", DryCheckHealthz)
}
