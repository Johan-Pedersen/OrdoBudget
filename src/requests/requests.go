package requests

import (
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

/*
Updates row-wise from (rowInd, colInd) to (rowInd + len(grpSums), colInd)
*/
func MultiUpdateReq(data []string, rowInd, colInd, sheetId int64) *sheets.Request {
	var rowData []*sheets.RowData

	for i := range data {
		rowData = append(rowData, &sheets.RowData{
			Values: []*sheets.CellData{
				{
					UserEnteredValue: &sheets.ExtendedValue{
						StringValue: &data[i],
					},
				},
			},
		})
	}

	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     sheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
Updates row-wise from (rowInd, colInd) to (rowInd + len(grpSums), colInd)
*/
func MultiUpdateReqNum(data []float64, rowInd, colInd, sheetId int64) *sheets.Request {
	var rowData []*sheets.RowData

	for i := range data {
		rowData = append(rowData, &sheets.RowData{
			Values: []*sheets.CellData{
				{
					UserEnteredValue: &sheets.ExtendedValue{
						NumberValue: &data[i],
					},
				},
			},
		})
	}

	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     sheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
Update single cell at (rowInd, colInd) with blank
*/

func SingleUpdateReqBlank(rowInd, colInd, sheetId int64) *sheets.Request {
	blank := ""
	rowData := &sheets.RowData{
		Values: []*sheets.CellData{
			{
				UserEnteredValue: &sheets.ExtendedValue{
					StringValue: &blank,
				},
			},
		},
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "UserEnteredValue",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     sheetId,
			},
			Rows: []*sheets.RowData{rowData},
		},
	}
	return updateReq
}

/*
Update single cell at (rowInd, colInd) with amount
*/

func SingleUpdateReq(amount float64, rowInd, colInd, sheetId int64) *sheets.Request {
	rowData := &sheets.RowData{
		Values: []*sheets.CellData{
			{
				UserEnteredValue: &sheets.ExtendedValue{
					NumberValue: &amount,
				},
			},
		},
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "UserEnteredValue",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     sheetId,
			},
			Rows: []*sheets.RowData{rowData},
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
