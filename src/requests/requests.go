package requests

import (
	"fmt"

	"google.golang.org/api/sheets/v4"
)

/*
Parameter navne her giver ikke rigtig mening med to / from
*/
func CutPasteSingleReq(fromRow, fromCol, toRow, toCol int64) *sheets.Request {
	cutPasteReq := &sheets.Request{
		CutPaste: &sheets.CutPasteRequest{
			Source: &sheets.GridRange{
				EndColumnIndex:   fromCol + 1,
				EndRowIndex:      fromRow + 1,
				SheetId:          1472288449,
				StartColumnIndex: fromCol,
				StartRowIndex:    fromRow,
			},
			Destination: &sheets.GridCoordinate{
				ColumnIndex: toCol,
				RowIndex:    toRow,
				SheetId:     1472288449,
			},
			PasteType: "PASTE_NORMAL", // Adjust paste type as needed
		},
	}
	return cutPasteReq
}

func UpdateReq(grpSums []float64, toRow, toCol, sheetId int64) *sheets.Request {
	var rowData []*sheets.RowData

	for i := range grpSums {
		fmt.Println(grpSums[i])
		rowData = append(rowData, &sheets.RowData{
			Values: []*sheets.CellData{
				{
					UserEnteredValue: &sheets.ExtendedValue{
						// Den fucker op fordi den tager referencen til v vaiablen, så hvad end den ender med at vœre til sidst er det de alle sammen er
						NumberValue: &grpSums[i],
					},
				},
			},
		})
	}

	for _, v := range rowData {
		fmt.Println(*v.Values[0].UserEnteredValue.NumberValue)
	}

	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: toCol,
				RowIndex:    toRow,
				SheetId:     sheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
For each excerpt, get date, amount, describtion
*/
func GetExcerpts() {
	// readRange := "Udtrœk!B2:C12"
	// // Get from "udtrœks sheet"
	// valRange, err := srv.Spreadsheets.Values.Get(1472288449, readRange).Do()
	// if err != nil {
	// 	log.Fatalf("Unable to perform get: %v", err)
	// }
	//
	// return valRange
}
