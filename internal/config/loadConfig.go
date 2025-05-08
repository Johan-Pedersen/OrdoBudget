package config

import (
	"OrdoBudget/internal/logtrace"
	"OrdoBudget/internal/request"

	"google.golang.org/api/sheets/v4"
)

func GetConfig() []*sheets.GridData {
	// Laes Entry Col fra sheets
	sheet := request.GetSheet()
	// Find excerpt grps to insert at

	// This is a monster and I'm sorry
	res, err := sheet.Get(request.SpreadSheetId).Ranges("Config!A2:C").
    Fields("sheets(data(rowData(values(userEnteredValue(stringValue,boolValue),effectiveFormat(backgroundColor)))))").Do()
	if err != nil {
		logtrace.Error(err.Error())
	}

	return res.Sheets[0].Data
}
