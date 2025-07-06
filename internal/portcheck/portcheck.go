package portcheck

import (
	"fmt"
	"net"
	"time"
)

type PortCheckResult struct {
	Host string
	Port int
	Success bool
	Duration time.Duration
	Error string
}

func TestPort(host string, port int, protocol string, timeout time.Duration) PortCheckResult {
	address := fmt.Sprintf("%s:%d", host, port)
	start := time.Now()
	conn, err := net.DialTimeout("tcp", address, timeout)
	duration := time.Since((start))

	if err != nil {
		return PortCheckResult{
			Host:    host,
			Port:    port,
			Success:  false,
			Duration: duration,
			Error:   err.Error(),
		}
	}

	conn.Close()
	return PortCheckResult{
		Host:    host,
		Port:    port,
		Success:  true,
		Duration: duration,
		Error:   "",	
	}
}