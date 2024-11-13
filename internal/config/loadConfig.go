package config

import (
	"OrdoBudget/internal/request"
	"log"

	"google.golang.org/api/sheets/v4"
)

func GetConfig() []*sheets.GridData {
	// Laes Entry Col fra sheets
	sheet := request.GetSheet()
	// Find excerpt grps to insert at

	// This is a monster and I'm sorry
	res, err := sheet.Get(request.SpreadSheetId).Ranges("Config!A2:C").Fields("sheets(data(rowData(values(userEnteredValue(stringValue,boolValue),effectiveFormat(backgroundColor)))))").Do()
	if err != nil {
		log.Fatalf("Unable to perform get operation: %v", err)
	}

	return res.Sheets[0].Data
}
