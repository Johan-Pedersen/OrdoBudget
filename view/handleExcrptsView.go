package view

import (
	"budgetAutomation/controller"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func handleExcrptsView(app fyne.App) {
	window := app.NewWindow("Handle excerpts")

	// Er dette en case for databinding ?
	excrpt := widget.NewEntry()

	excrpt.Disable()

	// Denne skal dynamisk opdateres
	excrpt.SetText("hew")
	// tree := widget.NewTree(
	// 	func(id widget.TreeNodeID) []widget.TreeNodeID {
	// 		switch id {
	// 		case "":
	// 			return []widget.TreeNodeID{"a", "b", "c"}
	// 		case "a":
	// 			return []widget.TreeNodeID{"a1", "a2"}
	// 		}
	// 		return []string{}
	// 	},
	// 	func(id widget.TreeNodeID) bool {
	// 		fmt.Printf("id: %v\n", id)
	// 		return id == ""
	// 		return id == "" || id == "a"
	// 	},
	// 	func(branch bool) fyne.CanvasObject {
	// 		if branch {
	// 			return widget.NewLabel("Branch template")
	// 		}
	// 		return widget.NewLabel("Leaf template")
	// 	},
	// 	func(id widget.TreeNodeID, branch bool, o fyne.CanvasObject) {
	// 		text := id
	// 		if branch {
	// 			text += " (branch)"
	// 		}
	// 		o.(*widget.Label).SetText(text)
	// 	})
	excrptGrps := widget.NewTree(
		// ChildUIs
		func(tni widget.TreeNodeID) []widget.TreeNodeID {
			// excrptgrpsAsstring := controller.GetExcrptGrpsAsString(tni)

			fmt.Printf("tni child UIs: %v\n", tni)
			if tni == "" {
				return controller.GetParentsAsString()
			} else {
				return controller.GetExcrptGrpsAsString(tni)
			}
		},
		// IsBranch
		func(tni widget.TreeNodeID) bool {
			// return true

			if tni == "" {
				return true
			} else {
				getChild := controller.GetExcrptGrps(tni)
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
	window.SetContent(excrptGrps)
	window.Resize(fyne.NewSize(500, 1500))
	window.SetFixedSize(true)
	window.Show()
}
