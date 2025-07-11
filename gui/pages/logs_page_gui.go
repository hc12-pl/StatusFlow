package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"statusFlow/internal/logger"
)

func NewLogsPage() fyne.CanvasObject {
	logs := widget.NewMultiLineEntry()
	logs.SetPlaceHolder("Logs will appear here...")
	logs.Wrapping = fyne.TextWrapWord
	scroll := container.NewScroll(logs)
	scroll.SetMinSize(fyne.NewSize(500, 300))

	// Load logs from the database
	logEntries, err := logger.GetAllLogs()
	if err != nil {
		logs.SetText("Error loading logs: " + err.Error())
		return container.NewVBox(logs)
	}

	for _, entry := range logEntries {
		logs.SetText(logs.Text + entry + "\n")
	}
	return container.NewVBox(
		widget.NewLabel("Application Logs"),
		scroll,
	)
}