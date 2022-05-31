package internal

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
