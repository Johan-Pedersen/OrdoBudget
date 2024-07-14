package main

import (
	"budgetAutomation/view"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	view.MainView(app)
}
