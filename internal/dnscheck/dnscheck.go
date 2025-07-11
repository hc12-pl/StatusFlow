package dnscheck

import (
	"net"
	"strings"
	"time"
)

type DNSCheckResult struct {
	Host     string
	IP       string
	Add      string
	Success  bool
	Duration time.Duration
	Error    string
}

func LookupAddress(IP net.IP) DNSCheckResult {
	start := time.Now()
	host, err := net.LookupAddr(IP.String())
	duration := time.Since(start)

	if err != nil || len(host) == 0 {
		return DNSCheckResult{
			Host:     IP.String(),
			IP:       IP.String(),
			Success:  false,
			Duration: duration,
			Error:    getErrorMsg(err, "No host found for IP"),
		}
	}

	return DNSCheckResult{
		Host:     host[0],
		IP:       IP.String(),
		Success:  true,
		Duration: duration,
	}
}

func LookupHost(host string) DNSCheckResult {
	start := time.Now()
	ips, err := net.LookupIP(host)
	duration := time.Since(start)

	if err != nil || len(ips) == 0 {
		return DNSCheckResult{
			Host:     host,
			Success:  false,
			Duration: duration,
			Error:    getErrorMsg(err, "No IP found for host"),
		}
	}

	ipStrs := make([]string, len(ips))
	for i, ip := range ips {
		ipStrs[i] = ip.String()
	}

	return DNSCheckResult{
		Host:     host,
		IP:       "[" + strings.Join(ipStrs, ", ") + "]",
		Success:  true,
		Duration: duration,
	}
}

func LookupCNAME(host string) DNSCheckResult {
	start := time.Now()
	cname, err := net.LookupCNAME(host)
	duration := time.Since(start)

	if err != nil {
		return DNSCheckResult{
			Host:     host,
			Success:  false,
			Duration: duration,
			Error:    err.Error(),
		}
	}

	return DNSCheckResult{
		Host:     host,
		Add:      cname,
		Success:  true,
		Duration: duration,
	}
}

func LookupTXT(host string) DNSCheckResult {
	start := time.Now()
	records, err := net.LookupTXT(host)
	duration := time.Since(start)

	if err != nil || len(records) == 0 {
		return DNSCheckResult{
			Host:     host,
			Success:  false,
			Duration: duration,
			Error:    getErrorMsg(err, "No TXT records found"),
		}
	}

	return DNSCheckResult{
		Host:     host,
		Add:      strings.Join(records, ", "),
		Success:  true,
		Duration: duration,
	}
}

func getErrorMsg(err error, fallback string) string {
	if err != nil {
		return err.Error()
	}
	return fallback
}
