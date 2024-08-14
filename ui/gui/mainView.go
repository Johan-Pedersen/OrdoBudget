package ui

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MainView(debug bool, app fyne.App) {
	initWindow := app.NewWindow("Main")
	initWindow.Resize(fyne.NewSize(300, 150))
	month := widget.NewEntry()
	excrptPath := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Måned: ", Widget: month},
			{Text: "Bank udtræk: ", Widget: excrptPath},
		},
		OnSubmit: func() { // optional, handle form submission
			if debug {
				submitDebug()
			} else {

				intMonth, isValid := isMonthValid(month.Text)
				if isValid {
					// Submit(intMonth, excrptPath.Text)

					submit(intMonth, excrptPath.Text)
				}
			}
			handleExcrptsView(app)

			initWindow.Close()
		},
	}

	initWindow.SetContent(form)
	initWindow.ShowAndRun()
}

func isMonthValid(month string) (int64, bool) {
	intMonth, err := strconv.ParseInt(month, 10, 64)
	if err != nil {
		log.Fatal("Month not a number")
	}
	return intMonth, intMonth >= 1 && intMonth <= 12
}
