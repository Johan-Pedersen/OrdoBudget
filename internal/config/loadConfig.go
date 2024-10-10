package config

import (
	"budgetAutomation/internal/requests"
	"log"

	"google.golang.org/api/sheets/v4"
)

func GetConfig() *sheets.ValueRange {
	// Laes Entry Col fra sheets
	sheet := requests.GetSheet()
	// Find excerpt grps to insert at

	res, err := sheet.Values.Get(requests.GetSpreadsheetId(), "Config!A2:C").Do() //
	if err != nil {
		log.Fatalf("Unable to perform get operation: %v", err)
	}

	return res
}
