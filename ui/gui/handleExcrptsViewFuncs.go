package ui

import (
	excrptgrps "budgetAutomation/internal/excrptGrps"
	"fmt"
)

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
