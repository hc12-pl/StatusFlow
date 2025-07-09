package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/pingcheck"
)


func NewIcmpTestPage() fyne.CanvasObject {

	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter target (e.g., google.com)")

	timeoutEntry := widget.NewEntry()
	timeoutEntry.SetPlaceHolder("Enter timeout in seconds (e.g., 2)")

	resultLabel := widget.NewLabel("")

	checkBtn := widget.NewButton("Ping target", func() {
		target := hostEntry.Text
		timeout := timeoutEntry.Text


		timeoutSec, err := strconv.Atoi(timeout)

		if err != nil {
			resultLabel.SetText("Incorrect timeout")
			return
		}

		resultLabel.SetText("Checking...")


		go func() {
			res := icmpcheck.PingHost(target, time.Duration(timeoutSec)*time.Second)
			text := ""

			if res.Success {
				text = fmt.Sprintf("%s is responding (time: %s)", res.Host, res.Duration)
			} else {
				text = fmt.Sprintf("Error: %s", res.Error)
			}

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title: "ICMP check result",
				Content: text,
			})

			fyne.Do(func() {
				resultLabel.SetText(text)
			})
		}()
	})

	return container.NewVBox(
		widget.NewLabel("ICMP Check"),
		hostEntry,
		timeoutEntry,
		checkBtn,
		widget.NewSeparator(),
		resultLabel,
	)
}