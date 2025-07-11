package gui

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/logger"
	"statusFlow/internal/webcheck"
)

func NewWebTestPage() fyne.CanvasObject {
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter host (e.g., example.com)")

	protocolSelect := widget.NewSelect([]string{"http", "https"}, nil)
	protocolSelect.SetSelected("http")

	timeoutEntry := widget.NewEntry()
	timeoutEntry.SetPlaceHolder("Timeout in seconds (e.g., 2)")

	resultLabel := widget.NewLabel("")

	checkBtn := widget.NewButton("Check Website", func() {
		rawHost := strings.TrimSpace(hostEntry.Text)
		protocol := protocolSelect.Selected
		timeoutStr := timeoutEntry.Text

		timeoutSec, err := strconv.Atoi(timeoutStr)
		if err != nil {
			resultLabel.SetText("❌ Invalid timeout")
			return
		}

		if rawHost == "" {
			resultLabel.SetText("❌ Host cannot be empty")
			return
		}

		resultLabel.SetText("⏳ Checking...")
		go func() {
			res := webcheck.CheckHTTPStatus(protocol, rawHost, time.Duration(timeoutSec)*time.Second)

			var text string
			if res.Error != "" {
				text = fmt.Sprintf("❌ Error: %s", res.Error)
			} else {
				text = fmt.Sprintf("✅ %s - Status: %d (%s) in %s",
					res.URL, res.StatusCode, res.Status, res.Duration)
			}
			logger.SaveLog(text)

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Web Check Result",
				Content: text,
			})

			fyne.Do(func() {
				resultLabel.SetText(text)
			})
		}()
	})

	return container.NewVBox(
		widget.NewLabel("HTTP/HTTPS Web Status Checker"),
		hostEntry,
		protocolSelect,
		timeoutEntry,
		checkBtn,
		widget.NewSeparator(),
		resultLabel,
	)
}
