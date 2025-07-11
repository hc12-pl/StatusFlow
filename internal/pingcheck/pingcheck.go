package icmpcheck

import (

	"time"

	"github.com/go-ping/ping"
)

type PingResult struct {
	Host     string
	Success  bool
	Duration time.Duration
	Error    string
}

func PingHost(host string, timeout time.Duration) PingResult {
	pinger, err := ping.NewPinger(host)
	if err != nil {
		return PingResult{
			Host:    host,
			Success: false,
			Error:   "init error: " + err.Error(),
		}
	}

	pinger.Count = 4
	pinger.Timeout = timeout

	start := time.Now()
	err = pinger.Run()
	duration := time.Since(start)

	if err != nil {
		return PingResult{
			Host:    host,
			Success: false,
			Duration: duration,
			Error:   "ping failed: " + err.Error(),
		}
	}

	stats := pinger.Statistics()

	return PingResult{
		Host:     host,
		Success:  true,
		Duration: stats.AvgRtt,
	}
}
