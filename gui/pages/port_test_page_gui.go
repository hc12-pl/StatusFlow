package gui

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/logger"
	"statusFlow/internal/portcheck"
)

func NewPortTestPage() fyne.CanvasObject {

	protocolSelect := widget.NewSelect([]string{"tcp", "udp"}, nil)
	protocolSelect.SetSelected("tcp")
	hostEntry := widget.NewEntry()
	hostEntry.SetPlaceHolder("Enter host (e.g., google.com)")

	portEntry := widget.NewEntry()
	portEntry.SetPlaceHolder("Enter port (e.g., 80)")

	timeoutEntry := widget.NewEntry()
	timeoutEntry.SetPlaceHolder("Enter timeout in seconds (e.g., 2)")

	resultLabel := widget.NewLabel("")

	checkBtn := widget.NewButton("Check Port", func() {
		
		host := hostEntry.Text
		portStr := portEntry.Text
		timeoutStr := timeoutEntry.Text

		port, err := strconv.Atoi(portStr)
		if err != nil {
			resultLabel.SetText("Incorrect port")
			logger.SaveLog(err.Error())
			return
		}

		timeoutSec, err := strconv.Atoi(timeoutStr)
		if err != nil {
			resultLabel.SetText("Incorrect timeout")
			logger.SaveLog(err.Error())
			return
		}

		resultLabel.SetText("⏳ Checking...")


		go func() {
			res := portcheck.TestPort(host, port, protocolSelect.Selected, time.Duration(timeoutSec)*time.Second)
			text := ""

			if res.Success {
				text = fmt.Sprintf("%s Port %d on %s is open (time: %s)", protocolSelect.Selected, res.Port, res.Host, res.Duration)
			} else {
				text = fmt.Sprintf("Error: %s", res.Error)
			}
			logger.SaveLog(text)

			fyne.CurrentApp().SendNotification(&fyne.Notification{
				Title:   "Port check result",
				Content: text,
			})

			fyne.Do(func() {
   				resultLabel.SetText(text)
			})

		}()
		
	})
	

	return container.NewVBox(
		widget.NewLabel("TCP Port Check"),
		hostEntry,
		portEntry,
		protocolSelect,
		timeoutEntry,
		checkBtn,
		widget.NewSeparator(),
		resultLabel,
	)
}
