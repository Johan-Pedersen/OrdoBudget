package view

import (
	excrptgrps "budgetAutomation/internal/excrptGrps"
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

func getExcrptGrps(parentName string) []excrptgrps.ExcrptGrp {
	return excrptgrps.GetChildren(parentName)
}

func getExcrptGrpsAsString(parentName string) []string {
	excrptgrps := excrptgrps.GetChildren(parentName)

	var names []string

	for _, eg := range excrptgrps {
		names = append(names, eg.Name)
	}

	return names
}

func getParentsAsString() []string {
	parents := excrptgrps.GetParents()

	var names []string
	for _, egp := range parents {
		names = append(names, egp.Name)
	}
	fmt.Printf("len(parents): %v\n", len(parents))
	fmt.Printf("len(names): %v\n", len(names))
	return names
}

func getParents() []excrptgrps.ExcrptGrpParent {
	return excrptgrps.GetParents()
}
