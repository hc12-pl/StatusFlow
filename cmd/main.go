package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"statusFlow/gui/pages"
	"statusFlow/internal/database"
)

func main() {
	// Initialize the database connection
	database.InitDB()
	defer database.DB.Close()

	myApp := app.New()
	myWindow := myApp.NewWindow("StatusFlow")

	portTestPage := gui.NewPortTestPage()
	webTestPage := gui.NewWebTestPage()
	icmpTestPage := gui.NewIcmpTestPage()
	logsPage := gui.NewLogsPage()

	content := container.NewAppTabs(
		container.NewTabItem("TCP Port", portTestPage),
		container.NewTabItem("Website", webTestPage),
		container.NewTabItem("ICMP", icmpTestPage),
		container.NewTabItem("Logs", logsPage),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.ShowAndRun()
}
