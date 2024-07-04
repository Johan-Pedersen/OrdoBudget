package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Container")
	myWindow.Resize(fyne.NewSize(300, 150))
	moEntry := widget.NewEntry()
	excrptEntry := widget.NewEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{ // we can specify items in the constructor
			{Text: "Måned: ", Widget: moEntry},
			{Text: "Bank udtræk: ", Widget: excrptEntry},
		},
		OnSubmit: func() { // optional, handle form submission
			log.Println("Form submitted:", moEntry.Text)
			myWindow.Close()
		},
	}

	myWindow.SetContent(form)
	myWindow.ShowAndRun()
}
