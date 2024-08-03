package main

import (
	ui "budgetAutomation/ui"
	"flag"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	var debug bool
	flag.BoolVar(&debug, "debug", "Run in debug mode")

	flag.Parse()

	ui.MainView(app)
}
