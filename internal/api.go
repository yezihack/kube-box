package internal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/yezihack/kube-box/internal/pkg"
)

func Home(ctx *gin.Context) {
	hostname, err := GetInterfaceIpv4Addr(GetEnvValue(EnvNetworkName, DefaultNetworkName))
	if err != nil {
		hostname = "None"
	}
	ctx.JSON(200, &HomeEntity{
		Host:    hostname,
		Version: GetEnvValue(EnvVersion, DefaultVersion),
		Date:    time.Now().Format(YYYYFormat),
	})

}
func Ping(ctx *gin.Context) {
	ctx.JSON(200, &ResponseEntity{
		Code:    CodeOk,
		Message: PING,
	})
}

func Healthz(ctx *gin.Context) {
	ctx.JSON(200, &ResponseEntity{
		Code:    CodeOk,
		Message: HEALTHZOK,
	})
}

// prometheus metrics
func Metrics(ctx *gin.Context) {
	promhttp.Handler().ServeHTTP(ctx.Writer, ctx.Request)
}

// 检查 ip 是否通达
func CheckIP(ctx *gin.Context) {
	result, err := getIPNetPool()

	if err != nil {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: err.Error(),
		})
		return
	}
	// 1. 详情，2. 不通达ip列表
	ctx.IndentedJSON(Code200, result)
}
func DryCheckIP(ctx *gin.Context) {
	result, err := getIPNetPool()
	if err != nil {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: err.Error(),
		})
		return
	}
	var (
		total = 0
		succ  = 0
		fail  = 0
	)
	for _, item := range result.List {
		total++
		if item.Status {
			succ++
		} else {
			fail++
		}
	}

	rsp := DryCheckResponse{
		Category:    result.Category,
		Hostname:    result.Hostname,
		LocalIP:     result.LocalIP,
		ElapsedTime: result.ElapsedTime,
		Total:       total,
		SuccNum:     succ,
		FailNum:     fail,
		BadList:     result.BadIPs,
	}
	ctx.IndentedJSON(Code200, rsp)
}

// 检查 healthz 接口是否通达
func CheckHealthz(ctx *gin.Context) {
	result, err := getIPHealthzPool()
	if err != nil {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: err.Error(),
		})
		return
	}

	// 1. 详情，2. 不通达ip列表
	ctx.IndentedJSON(Code200, result)
}

func DryCheckHealthz(ctx *gin.Context) {
	result, err := getIPHealthzPool()
	if err != nil {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: err.Error(),
		})
		return
	}
	var (
		total = 0
		succ  = 0
		fail  = 0
	)
	for _, item := range result.List {
		total++
		if item.Status {
			succ++
		} else {
			fail++
		}
	}
	rsp := DryCheckResponse{
		Category:    result.Category,
		Hostname:    result.Hostname,
		LocalIP:     result.LocalIP,
		ElapsedTime: result.ElapsedTime,
		Total:       total,
		SuccNum:     succ,
		FailNum:     fail,
		BadList:     result.BadIPs,
	}
	ctx.IndentedJSON(Code200, rsp)
}

// 检查MYSQL连接状态
func CheckMySQLConnect(ctx *gin.Context) {
	host := ctx.Query("host")
	port := ctx.Query("port")
	user := ctx.Query("user")
	pass := ctx.Query("pass")
	if strings.EqualFold(host, "") ||
		strings.EqualFold(port, "") ||
		strings.EqualFold(user, "") ||
		strings.EqualFold(pass, "") {
		ctx.JSON(Code400, &ResponseEntity{
			Code:    Code400,
			Message: "host,port,user,pass is lock.",
		})
		return
	}
	// port change int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		ctx.JSON(Code400, &ResponseEntity{
			Code:    Code400,
			Message: "port not integer",
		})
		return
	}
	// config
	dbInfo := pkg.MySQLHandlerInfo{
		Host:     host,
		Port:     portInt,
		User:     user,
		Password: pass,
	}
	handler := pkg.NewMySQLHandler()
	defer handler.Close()
	err = handler.Connect(&dbInfo)
	if err != nil {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: err.Error(),
		})
		return
	}
	if !handler.Ping() {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: "mysql ping is error",
		})
		return
	}
	ctx.IndentedJSON(Code200, &ResponseEntity{
		Code:    Code200,
		Message: "mysql connection is ok.",
	})
}

// 创建数据库
func CreateMySQLDB(ctx *gin.Context) {
	host := ctx.Query("host")
	port := ctx.Query("port")
	user := ctx.Query("user")
	pass := ctx.Query("pass")
	dbname := ctx.Query("dbname")
	if strings.EqualFold(host, "") ||
		strings.EqualFold(port, "") ||
		strings.EqualFold(user, "") ||
		strings.EqualFold(pass, "") ||
		strings.EqualFold(dbname, "") {
		ctx.JSON(Code400, &ResponseEntity{
			Code:    Code400,
			Message: "host,port,user,pass,dbname is lock.",
		})
		return
	}
	// port change int
	portInt, err := strconv.Atoi(port)
	if err != nil {
		ctx.JSON(Code400, &ResponseEntity{
			Code:    Code400,
			Message: "port not integer",
		})
		return
	}
	// config
	dbInfo := pkg.MySQLHandlerInfo{
		Host:     host,
		Port:     portInt,
		User:     user,
		Password: pass,
		Database: dbname,
	}
	handler := pkg.NewMySQLHandler()
	defer handler.Close()
	err = handler.Connect(&dbInfo)
	if err != nil {
		ctx.IndentedJSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: fmt.Sprintf("connect to mysql is error, describe: %v", err),
		})
		return
	}
	if !handler.Ping() {
		ctx.JSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: "mysql ping is error",
		})
		return
	}
	// create database
	err = handler.CreateDBName()
	if err != nil {
		ctx.IndentedJSON(Code500, &ResponseEntity{
			Code:    Code500,
			Message: "create database is error",
			Err:     err,
		})
		log.Println("create database is error:", err)
		return
	}
	ctx.IndentedJSON(Code200, &ResponseEntity{
		Code:    Code200,
		Message: "create database is ok",
	})
}

func getIPNetPool() (result *CheckIPResponse, err error) {
	fileDir := GetEnvValue(EnvDataPath, DefaultDataPath)
	fileName := GetEnvValue(EnvIpDataFileName, DefaultIpDataFileName)

	fullPath := JoinFullPath(fileDir, fileName)
	networkName := GetEnvValue(EnvNetworkName, DefaultNetworkName)
	timeout := GetEnvValueToInteger(EnvTimeout, DefaultTimeout)
	goNumber := GetEnvValueToInteger(EnvGoNumber, DefaultGoNumber)

	svc := NewIPService(fullPath, networkName, time.Second*time.Duration(timeout), goNumber)
	return svc.IPNetPool()
}

func getIPHealthzPool() (result *CheckIPResponse, err error) {
	fileDir := GetEnvValue(EnvDataPath, DefaultDataPath)
	fileName := GetEnvValue(EnvIpDataFileName, DefaultIpDataFileName)

	fullPath := JoinFullPath(fileDir, fileName)
	networkName := GetEnvValue(EnvNetworkName, DefaultNetworkName)
	timeout := GetEnvValueToInteger(EnvTimeout, DefaultTimeout)
	goNumber := GetEnvValueToInteger(EnvGoNumber, DefaultGoNumber)
	targetPort := GetEnvValueToInteger(EnvTargetPort, DefaultTargetPort)
	healthzPathName := GetEnvValue(EnvHealthzPathName, DefaultHealthzPathName)

	svc := NewIPService(fullPath, networkName, time.Second*time.Duration(timeout), goNumber)
	return svc.IPHealthzPool(targetPort, healthzPathName)
}
