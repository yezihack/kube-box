package internal

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/go-ping/ping"
	"github.com/pkg/errors"
)

// 优先获取环境变量，否则获取默认值
func GetEnvValue(envKey, defaultVal string) string {
	val, ok := os.LookupEnv(envKey)
	if ok {
		return val
	}
	return defaultVal
}

// 获取主机名
func GetHostname() string {
	name, err := os.Hostname()
	if err != nil {
		return "None"
	}
	return name
}

// 优先获取环境变量，否则获取默认值
func GetEnvValueToInteger(envKey string, defaultVal int) int {
	val, ok := os.LookupEnv(envKey)
	if ok {
		i, err := strconv.Atoi(val)
		if err != nil {
			return defaultVal
		}
		return i
	}
	return defaultVal
}

// 检测网络是否通达
func PingNetwork(target string) (ok bool, stats *ping.Statistics, err error) {
	pinger, err := ping.NewPinger(target)
	if err != nil {
		return
	}
	ICMPCOUNT := 3
	pinger.Count = ICMPCOUNT
	pinger.Timeout = time.Duration(time.Second * 3)
	pinger.SetPrivileged(true) // true -> icmp, false -> udp
	pinger.Run()               // blocks until finished
	stats = pinger.Statistics()
	// 有回包，就是说明IP是可用的
	if stats.PacketsRecv >= 1 {
		ok = true
		return
	}
	return
}

// 按行读取文本数据。
func ReadTextLine(textfile string) (ls []string, err error) {
	file, err := os.Open(textfile)
	if err != nil {
		return
	}
	defer file.Close()
	ls = make([]string, 0, 10)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text() // or
		//line := scanner.Bytes()
		ls = append(ls, line)
	}
	if err = scanner.Err(); err != nil {
		return
	}
	return
}

// curl 发起 get请求
func CurlGet(uri string, timeout time.Duration) (result []byte, statuCode int, err error) {
	// 创建一个 http 客户端
	cli := &http.Client{}
	// 写入 uri 请求信息
	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 设置超时
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req = req.WithContext(ctx)
	// 发起请求
	resp, err := cli.Do(req)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	// 关闭连接
	defer resp.Body.Close()
	// 读取 body
	result, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	statuCode = resp.StatusCode
	return
}

// 判断 path 路径结尾处是否有 / 符号
func IsExistPathLastPole(pathStr string) bool {
	if len(pathStr) == 0 {
		return false
	}
	s := pathStr[len(pathStr)-1:]
	if s == "/" {
		return true
	}
	return false
}

// 拼接完整的文件路径
func JoinFullPath(dir, file string) string {
	if !IsExistPathLastPole(dir) {
		return dir + "/" + file
	}
	return dir + file
}

// LocalIPs return all non-loopback IPv4 addresses
func LocalIPv4s() ([]string, error) {
	var ips []string
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
			ips = append(ips, ipnet.IP.String())
		}
	}

	return ips, nil
}

// 获取指定网卡的IP
func GetLocalIpV4() string {
	inters, err := net.Interfaces()
	if err != nil {
		panic(err)
	}
	for _, inter := range inters {
		// 判断网卡是否开启，过滤本地环回接口
		if inter.Flags&net.FlagUp != 0 && !strings.HasPrefix(inter.Name, "lo") {
			// 获取网卡下所有的地址
			addrs, err := inter.Addrs()
			if err != nil {
				continue
			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
					//判断是否存在IPV4 IP 如果没有过滤
					if ipnet.IP.To4() != nil {
						return ipnet.IP.String()
					}
				}
			}
		}
	}
	return ""
}

func GetInterfaceIpv4Addr(interfaceName string) (addr string, err error) {
	var (
		ief      *net.Interface
		addrs    []net.Addr
		ipv4Addr net.IP
	)
	if ief, err = net.InterfaceByName(interfaceName); err != nil { // get interface
		return
	}
	if addrs, err = ief.Addrs(); err != nil { // get addresses
		return
	}
	for _, addr := range addrs { // get ipv4 address
		if ipv4Addr = addr.(*net.IPNet).IP.To4(); ipv4Addr != nil {
			break
		}
	}
	if ipv4Addr == nil {
		return "", errors.New(fmt.Sprintf("interface %s don't have an ipv4 address\n", interfaceName))
	}
	return ipv4Addr.String(), nil
}
