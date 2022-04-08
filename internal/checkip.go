package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/ivpusic/grpool"
	"github.com/unknwon/com"
)

type Check struct {
	FilePath        string        // 文件路径地址，全路径. /work/data/ip.data
	Timeout         time.Duration // 超时设置
	NetworkCardName string        // 网卡名称
	GoroutineNumber int           // goroutine 数量
}

/*
第一种检测：
1. 获取 ip.data 文件里的IP表表，ip.data一行一个IP地址。
2. 检查网络是否通达。
3. 通达的网络存储在 bad_map 里
4. 不通达的网络存储在 good_map 里

map 结构设计
key: ip addr
value: ipInfo_struct

ipInfo_struct:
state: 1 为通达，0为不通达
timeout: 单位ms, 请求耗时
*/
func (c *Check) IPNetPool() (result *CheckIPResponse, err error) {
	result = new(CheckIPResponse)
	start := time.Now()
	defer func() {
		result.ElapsedTime = time.Since(start).String()
	}()
	// 1. get text line
	lines, err := c.getIPData()
	if err != nil {
		return
	}
	size := len(lines)
	// check URL
	if size == 0 {
		err = errors.New("ip is null")
		return
	}

	// 获取本地网卡IP
	localIP, err := GetInterfaceIpv4Addr(c.NetworkCardName)
	if err != nil {
		return
	}
	// chan 接受值，带缓冲
	poolCh := make(chan *CheckIPItem, size)
	// pool
	pool := grpool.NewPool(size, c.GoroutineNumber)
	defer pool.Release()
	pool.WaitCount(size)

	// check 子函数
	jobCheck := func(hostAddr string) func() {
		return func() {
			defer pool.JobDone()
			item := &CheckIPItem{
				HostAddr: hostAddr,
			}
			// ping ip
			ok, rttMs, err := c.pingIPNet(hostAddr)
			fmt.Printf("add:%s, err:%v\n", hostAddr, err)
			if err != nil {
				item.Err = err.Error()
			} else {
				item.Status = ok
				item.RequestMS = int(rttMs)
			}
			poolCh <- item
		}
	}
	// check ip
	for _, hostAddr := range lines {
		pool.JobQueue <- jobCheck(hostAddr)
	}
	pool.WaitAll()
	close(poolCh) // 关闭通道，接受数据

	result.Hostname = GetHostname()
	result.Category = CategoryCheckIP
	result.LocalIP = localIP
	lst := make([]*CheckIPItem, 0, size/2)
	badIPs := make([]string, 0)

	for item := range poolCh {
		lst = append(lst, item)
		if !item.Status {
			badIPs = append(badIPs, item.HostAddr)
		}
	}

	result.List = lst
	result.BadIPs = badIPs
	return
}

/*
第二种检测：
1. 获取 ip.data 文件里的IP表表，ip.data一行一个IP地址。
2. 拼接健康URL: http://ip:9110/healthz
3. 通达的网络存储在 bad_map 里
4. 不通达的网络存储在 good_map 里

map 结构设计
key: ip addr
value: ipInfo_struct

ipInfo_struct:
state: 1 为通达，0为不通达
timeout: 单位ms, 请求耗时
*/
// port 端口号，URLPathName 路径名称，如 http://127.0.0.1/healthz, 名称即: healthz
func (c *Check) IPHealthzPool(port int, URLPathName string) (result *CheckIPResponse, err error) {
	result = new(CheckIPResponse)
	start := time.Now()
	defer func() {
		result.ElapsedTime = time.Since(start).String()
	}()
	// 1. get text line
	lines, err := c.getIPData()
	if err != nil {
		return
	}
	// size
	size := len(lines)
	// check URL
	if size == 0 {
		err = errors.New("ip is null")
		return
	}
	// 获取本地网卡IP
	localIP, err := GetInterfaceIpv4Addr(c.NetworkCardName)
	if err != nil {
		return
	}
	// chan cache ch
	poolCh := make(chan *CheckIPItem, size)
	// init pool
	pool := grpool.NewPool(size, c.GoroutineNumber)
	defer pool.Release()
	pool.WaitCount(size)

	// check 子函数
	jobCheck := func(hostAddr string) func() {
		return func() {
			defer pool.JobDone()
			requestURL := fmt.Sprintf("http://%s:%d/%s", hostAddr, port, URLPathName)
			item := &CheckIPItem{
				RequestAdrr: requestURL,
				HostAddr:    hostAddr,
			}
			// check healthz
			ok, requestTime, err := c.checkHealthz(requestURL)
			if err != nil {
				item.Err = err.Error()
			} else {
				item.Status = ok
				item.RequestMS = int(requestTime)
			}
			poolCh <- item
		}
	}

	// check ip
	for _, hostAddr := range lines {
		pool.JobQueue <- jobCheck(hostAddr)
	}
	pool.WaitAll()
	close(poolCh)

	result.Hostname = GetHostname()
	result.Category = CategoryCheckHealthz
	result.LocalIP = localIP
	lst := make([]*CheckIPItem, 0, size/2)
	badIPs := make([]string, 0)
	for item := range poolCh {
		lst = append(lst, item)
		if !item.Status {
			badIPs = append(badIPs, item.HostAddr)
		}
	}
	result.List = lst
	result.BadIPs = badIPs
	return
}

// 获取文本里的内容。
func (c *Check) getIPData() (ls []string, err error) {
	exist := com.IsExist(c.FilePath)
	if !exist {
		err = errors.New(fmt.Sprintf("%s is not exist", c.FilePath))
		return
	}
	isFile := com.IsFile(c.FilePath)
	if !isFile {
		err = errors.New(fmt.Sprintf("%s is not file", c.FilePath))
		return
	}
	ls, err = ReadTextLine(c.FilePath)
	if err != nil {
		return
	}
	return
}

// 检查IP是否通达
func (c *Check) pingIPNet(ipstr string) (result bool, rttMs int64, err error) {
	ok, stats, err := PingNetwork(ipstr)
	if err != nil {
		return
	}
	rttMs = stats.MaxRtt.Milliseconds()
	if !ok {
		err = errors.New(fmt.Sprintf("Packet Loss:%f", stats.PacketLoss))
		return
	}
	result = true
	return
}

// 检查 healthz 接口是否通达
func (c *Check) checkHealthz(requestURL string) (ok bool, rttMs int64, err error) {
	// send url
	start := time.Now()
	defer func() {
		rttMs = time.Since(start).Milliseconds()
	}()
	result, statuCode, err := CurlGet(requestURL, c.Timeout)
	if err != nil {
		return
	}

	if statuCode != Code200 {
		err = errors.New(fmt.Sprintf("URL:%s, statusCode:%d", requestURL, statuCode))
		return
	}
	// deal json
	rsp := ResponseEntity{}
	if err = json.Unmarshal(result, &rsp); err != nil {
		return
	}
	// check result
	if rsp.Code == CodeOk {
		ok = true
		return
	}
	return
}
