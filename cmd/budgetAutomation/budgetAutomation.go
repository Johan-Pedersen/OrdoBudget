package main

import (
	ui "budgetAutomation/ui"
	"flag"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	debugMode := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()
	ui.MainView(*debugMode, app)
}
