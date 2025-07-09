package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"statusFlow/gui/pages"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("StatusFlow")

	portTestPage := gui.NewPortTestPage()
	webTestPage := gui.NewWebTestPage()
	icmpTestPage := gui.NewIcmpTestPage()

	content := container.NewAppTabs(
		container.NewTabItem("TCP Port", portTestPage),
		container.NewTabItem("Website", webTestPage),
		container.NewTabItem("ICMP", icmpTestPage),
	)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(500, 400))
	myWindow.ShowAndRun()
}
