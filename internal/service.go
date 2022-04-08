package internal

import "time"

type IPService interface {
	IPNetPool() (result *CheckIPResponse, err error)
	IPHealthzPool(port int, URLPathName string) (result *CheckIPResponse, err error)
}

func NewIPService(filePath, networkName string, timeout time.Duration, goroutineNumber int) IPService {
	return &Check{
		FilePath:        filePath,
		NetworkCardName: networkName,
		Timeout:         timeout,
		GoroutineNumber: goroutineNumber,
	}
}
