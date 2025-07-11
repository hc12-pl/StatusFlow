package gui

import (
	"fmt"
	"net"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/dnscheck"
	"statusFlow/internal/logger"
)

func NewDnsTestPage() fyne.CanvasObject {
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter target domain (e.g., google.com)")

	ipEntry := widget.NewEntry()
	ipEntry.SetPlaceHolder("Enter target IP address (e.g., 1.1.1.1)")

	testType := widget.NewSelect([]string{"Reverse", "Normal", "CNAME", "TXT"}, nil)
	testType.SetSelected("Normal")

	resultLabel := widget.NewLabel("")
	resultLabel.Wrapping = fyne.TextWrapWord

	scroll := container.NewScroll(resultLabel)
	scroll.SetMinSize(fyne.NewSize(400, 100))

	checkBtn := widget.NewButton("Check", func() {
		host := hostEntry.Text
		ip := net.ParseIP(ipEntry.Text)
		test := testType.Selected

		go func() {
			var text string

			switch test {
			case "Reverse":
				res := dnscheck.LookupAddress(ip)
				if res.Success {
					text = fmt.Sprintf("Reverse lookup for %s found: %s\nDuration: %s", res.IP, res.Host, res.Duration)
				} else {
					text = fmt.Sprintf("No domain found for IP %s\nError: %s\nDuration: %s", res.Host, res.Error, res.Duration)
				}
			case "Normal":
				res := dnscheck.LookupHost(host)
				if res.Success {
					text = fmt.Sprintf("Host lookup for %s returned IP(s): %s\nDuration: %s", res.Host, res.IP, res.Duration)
				} else {
					text = fmt.Sprintf("No IP addresses found for %s\nError: %s\nDuration: %s", res.Host, res.Error, res.Duration)
				}
			case "CNAME":
				res := dnscheck.LookupCNAME(host)
				if res.Success {
					text = fmt.Sprintf("CNAME record for %s: %s\nDuration: %s", res.Host, res.Add, res.Duration)
				} else {
					text = fmt.Sprintf("No CNAME record for %s\nError: %s\nDuration: %s", res.Host, res.Error, res.Duration)
				}
			case "TXT":
				res := dnscheck.LookupTXT(host)
				if res.Success {
					text = fmt.Sprintf("TXT records for %s: %s\nDuration: %s", res.Host, res.Add, res.Duration)
				} else {
					text = fmt.Sprintf("No TXT records for %s\nError: %s\nDuration: %s", res.Host, res.Error, res.Duration)
				}
			}

			logger.SaveLog(text)

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "DNS check result",
				Content: text,
			})

			fyne.Do(func() {
				resultLabel.SetText(text)
			})
		}()
	})

	return container.NewVBox(
		widget.NewLabel("DNS Check"),
		hostEntry,
		ipEntry,
		testType,
		checkBtn,
		widget.NewSeparator(),
		scroll,
	)
}
