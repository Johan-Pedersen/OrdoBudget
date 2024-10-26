package config

import (
	"OrdoBudget/internal/request"
	"log"

	"google.golang.org/api/sheets/v4"
)

func GetConfig() *sheets.ValueRange {
	// Laes Entry Col fra sheets
	sheet := request.GetSheet()
	// Find excerpt grps to insert at

	res, err := sheet.Values.Get(request.SpreadSheetId, "Config!A2:C").Do() //
	if err != nil {
		log.Fatalf("Unable to perform get operation: %v", err)
	}

	return res
}
