package webcheck

import (
	"net/http"
	"strings"
	"time"
	"fmt"
)

type HTTPStatusResult struct {
	URL        string
	StatusCode int
	Status     string
	Error      string
	Duration   time.Duration
}

func CheckHTTPStatus(protocol, rawHost string, timeout time.Duration) HTTPStatusResult {
	if protocol != "http" && protocol != "https" {
		protocol = "http" // fallback
	}

	// Ensure proper URL formatting
	url := rawHost
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		url = fmt.Sprintf("%s://%s", protocol, rawHost)
	}

	client := http.Client{Timeout: timeout}
	start := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(start)

	if err != nil {
		return HTTPStatusResult{
			URL:      url,
			Error:    err.Error(),
			Duration: duration,
		}
	}
	defer resp.Body.Close()

	return HTTPStatusResult{
		URL:        url,
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Duration:   duration,
	}
}