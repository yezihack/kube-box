package internal

type ResponseEntity struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Err     error  `json:"err,omitempty"`
}
type HomeEntity struct {
	Date    string `json:"date,omitempty"`
	Host    string `json:"host,omitempty"`
	Version string `json:"version,omitempty"`
}

// check 网络返回结构
type RespCheckNet struct {
	Category string              `json:"category"`
	LocalIP  string              `json:"local_ip,omitempty"`
	Describe []*DescribeCheckNet `json:"describe,omitempty"`
	BadIP    []string            `json:"bad_ip"`
}
type DescribeCheckNet struct {
	TargetIP   string `json:"target_ip,omitempty"`
	RequestURL string `json:"request_url,omitempty"`
	RRTMs      int    `json:"rrt_ms"`
	Status     bool   `json:"status"`
	Err        error  `json:"err,omitempty"`
}

// check ip 返回结构体
type CheckIPResponse struct {
	Hostname    string         `json:"hostname,omitempty"`
	Category    string         `json:"category,omitempty"` // 分类
	LocalIP     string         `json:"local_ip,omitempty"` // 本地IP
	ElapsedTime string         `json:"elapsed_time,omitempty"`
	List        []*CheckIPItem `json:"list,omitempty"`
	BadIPs      []string       `json:"bad_ips,omitempty"`
}

// 子项 check ip 返回结构体
type CheckIPItem struct {
	HostAddr    string `json:"host_addr,omitempty"`
	RequestAdrr string `json:"request_adrr,omitempty"` // 请求URL
	RequestMS   int    `json:"request_ms"`             // 请求耗时
	Status      bool   `json:"status"`                 // true 通达，false 不通达
	Err         string `json:"err,omitempty"`          // 请求发生的错误
}

type DryCheckResponse struct {
	Hostname    string   `json:"hostname,omitempty"`     // 主机名称
	Category    string   `json:"category,omitempty"`     // 分类
	LocalIP     string   `json:"local_ip,omitempty"`     // 本地IP
	ElapsedTime string   `json:"elapsed_time,omitempty"` // 总耗时
	Total       int      `json:"total,omitempty"`        // 总条数
	SuccNum     int      `json:"succ_num,omitempty"`     // 成功条数
	FailNum     int      `json:"fail_num,omitempty"`     // 失败条数
	BadList     []string `json:"bad_list,omitempty"`     // 失败IP列表
}
