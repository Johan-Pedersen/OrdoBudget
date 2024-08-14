package ui

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func handleExcrptsView(app fyne.App) {
	window := app.NewWindow("Handle excerpts")

	// Er dette en case for databinding ?
	excrpt := widget.NewEntry()

	// excrpt.Disable()

	// Denne skal dynamisk opdateres
	excrpt.SetText("hew")

	excrptGrps := widget.NewTree(
		// ChildUIs
		func(tni widget.TreeNodeID) []widget.TreeNodeID {
			// excrptgrpsAsstring := GetExcrptGrpsAsString(tni)

			fmt.Printf("tni child UIs: %v\n", tni)
			if tni == "" {
				return getParentsAsString()
			} else {
				return getExcrptGrpsAsString(tni)
			}
		},
		// IsBranch
		func(tni widget.TreeNodeID) bool {
			// return true

			if tni == "" {
				return true
			} else {
				getChild := getExcrptGrps(tni)
				fmt.Printf("tni: %v\n", tni)
				return getChild != nil
			}
		},
		// CreateNode
		func(branch bool) fyne.CanvasObject {
			if branch {
				return widget.NewLabel("Branch template")
			}
			return widget.NewLabel("Leaf template")
		},
		// UpdateNode
		func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
			text := id
			if branch {
				text += " (branch)"
			}
			o.(*widget.Label).SetText(text)
		},
	)
	grid := container.NewGridWithColumns(2, excrptGrps, widget.NewLabel(""), widget.NewLabel(""), excrpt)
	window.SetContent(grid)
	window.Resize(fyne.NewSize(500, 1500))
	window.SetFixedSize(true)
	window.Show()
}
