package main

import (
	gui "budgetAutomation/ui/gui"
	"flag"

	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()

	debugMode := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()
	gui.MainView(*debugMode, app)
}
