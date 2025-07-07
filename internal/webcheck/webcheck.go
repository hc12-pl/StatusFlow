package webcheck

import (
	"fmt"
	"net"
	"time"
)

type PortCheckResult struct {
	Host string
	Type string
	Success bool
	Duration time.Duration
	Error string
}


func TestWeb(host string, webtype string, protocol string, timeout time.Duration) PortCheckResult {
	port := 0
	switch webtype {
		case "http":
			port = 80
		case "https":
			port = 443
	}
	address := net.JoinHostPort(host, fmt.Sprintf("%d", port))
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, timeout)
	duration := time.Since((start))

	if err != nil {
		return PortCheckResult{
			Host:    host,
			Type:    webtype,
			Success:  false,
			Duration: duration,
			Error:   err.Error(),
		}
	}

	conn.Close()
	return PortCheckResult{
		Host:    host,
		Type:    webtype,
		Success:  true,
		Duration: duration,
		Error:   "",	
	}
}