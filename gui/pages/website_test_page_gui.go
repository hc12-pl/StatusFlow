package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/webcheck"
)

func NewWebTestPage() fyne.CanvasObject {
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter host (e.g., example.com)")

	protocolSelect := widget.NewSelect([]string{"http", "https"}, func(value string) {})
	protocolSelect.SetSelected("https")

	timeoutEntry := widget.NewEntry()
	timeoutEntry.SetPlaceHolder("Timeout in seconds (e.g., 2)")

	resultLabel := widget.NewLabel("")

	checkBtn := widget.NewButton("Check Website", func() {
		host := hostEntry.Text
		webtype := protocolSelect.Selected
		timeoutStr := timeoutEntry.Text

		timeoutSec, err := strconv.Atoi(timeoutStr)
		if err != nil {
			resultLabel.SetText("Invalid timeout")
			return
		}

		resultLabel.SetText("⏳ Checking...")
		go func() {
			res := webcheck.TestWeb(host, webtype, "tcp", time.Duration(timeoutSec)*time.Second)
			text := ""

			if res.Success {
				text = fmt.Sprintf("[%s] %s is reachable (time: %s)", res.Type, res.Host, res.Duration)
			} else {
				text = fmt.Sprintf("Error: %s", res.Error)
			}

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Web check result",
				Content: text,
			})

			fyne.Do(func() {
				resultLabel.SetText(text)
			})
		}()
	})

	return container.NewVBox(
		widget.NewLabel("HTTP/HTTPS Web Check"),
		hostEntry,
		protocolSelect,
		timeoutEntry,
		checkBtn,
		widget.NewSeparator(),
		resultLabel,
	)
}
