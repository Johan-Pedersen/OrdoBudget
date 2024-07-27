package main

import (
	ui "budgetAutomation/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	ui.MainView(app)
}
