package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	HTTPRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "kubebox_requests_total",
			Help: "Number of the http requests received since the server started",
		},
		[]string{"status"},
	)
)

func init() {
	prometheus.MustRegister(HTTPRequests)
}

func AddRoute(r *gin.Engine) {
	r.Use(RequestCount())
	r.GET("/", Home)
	r.GET("/ping", Ping)
	r.GET("/healthz", Healthz)
	r.GET("/metrics", Metrics)
	r.GET("/check-ip", CheckIP)
	r.GET("/dry-check-ip", DryCheckIP)
	r.GET("/check-healthz", CheckHealthz)
	r.GET("/dry-check-healthz", DryCheckHealthz)
	r.GET("/check-mysql", CheckMySQLConnect)
	r.GET("/create-mysql-db", CreateMySQLDB)
}
