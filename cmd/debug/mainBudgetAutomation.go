package main

import (
	"log"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func MainView(app fyne.App) {
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
			log.Println("Form submitted:", month.Text)
			intMonth, err := strconv.ParseInt(month.Text, 10, 64)
			if err != nil {
				log.Fatal("Month not a number")
			}

			if intMonth >= 1 && intMonth <= 12 {
				// Submit(intMonth, excrptPath.Text)

				submit(intMonth, excrptPath.Text)
				handleExcrptsView(app)

				initWindow.Close()
			}
		},
	}

	initWindow.SetContent(form)
	initWindow.ShowAndRun()
}
