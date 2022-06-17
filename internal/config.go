package internal

// 常量
const (
	YYYYFormat = "2006-01-02 15:04:05"
)

// 默认值
const (
	DefaultVersion         = "v0.3.0"  // 版本号 默认值
	DefaultPort            = 80        // 端口 默认值
	DefaultTargetPort      = 80        // 目标端口，默认值
	DefaultDataPath        = "./data/" // 数据存储目录路径 默认值
	DefaultIpDataFileName  = "ip.data" // ip数据文件名 默认值
	DefaultNetworkName     = "eth0"    // 网卡名称 默认值
	DefaultGoNumber        = 10        // goroutine 数量
	DefaultTimeout         = 5         // 超时，单位s
	DefaultHealthzPathName = "healthz" // URL健康接口名称
	PING                   = "PONG"    // PING 返回的结果
	HEALTHZOK              = "OK"      // healthz 返回的结果
)

const (
	EnvPort            = "PORT"              // 端口 环境变量名称
	EnvTargetPort      = "TARGET_PORT"       // 目标端口 环境变量名称
	EnvVersion         = "VERSION"           // 版本 环境变量名称
	EnvDataPath        = "DATA_PATH"         // 数据存储目录路径 环境变量名称
	EnvIpDataFileName  = "IP_DATA_FILENAME"  // IP数据文件名称 环境变量名称
	EnvNetworkName     = "NETWORK_NAME"      // 网卡名称 环境变量名称
	EnvGoNumber        = "GO_NUMBER"         // goroutine 数量 环境变量名称
	EnvTimeout         = "TIMEOUT"           // 超时 环境变量名称
	EnvHealthzPathName = "HEALTHZ_PATH_NAME" // URL健康接口名称 环境变量名称
)

// 错误号
const (
	CodeOk  = 200
	Code200 = CodeOk
	Code400 = 400
	Code404 = 404
	Code500 = 500
)

const (
	CategoryCheckIP      = "CheckIP"
	CategoryCheckHealthz = "CheckHealthz"
)
