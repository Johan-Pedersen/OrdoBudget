package request

import (
	"budgetAutomation/internal/parser"
	"budgetAutomation/internal/util"
	"log"
	"os"

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
					UserEnteredFormat: &sheets.CellFormat{
						NumberFormat: &sheets.NumberFormat{
							Type: "NUMBER",
						},
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
func MultiUpdateReqDate(data []float64, rowInd, colInd, sheetId int64) *sheets.Request {
	var rowData []*sheets.RowData

	for i := range data {
		rowData = append(rowData, &sheets.RowData{
			Values: []*sheets.CellData{
				{
					UserEnteredValue: &sheets.ExtendedValue{
						NumberValue: &data[i],
					},

					UserEnteredFormat: &sheets.CellFormat{
						NumberFormat: &sheets.NumberFormat{
							Type:    "DATE",
							Pattern: "yyyy/mm/dd",
						},
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

func UpdateExcrptSheet(path string, month int64) []*sheets.Request {
	// open the csv file
	file, err := os.Open(path)
	if err != nil {
		print("could not open excerpt file")
		log.Fatalln("coud not open excerpt file.", err)
	}
	defer file.Close()

	excrpts := parser.ReadExcrptCsv(file, month)

	var dates []float64

	for _, exc := range excrpts {
		dates = append(dates, util.ConvertDateToFloat(exc.Date))
	}

	var amounts []float64
	for _, exc := range excrpts {
		amounts = append(amounts, exc.Amount)
	}

	var descriptions []string
	for _, exc := range excrpts {
		descriptions = append(descriptions, exc.Description)
	}

	var balances []float64
	for _, exc := range excrpts {
		balances = append(balances, exc.Balance)
	}

	return []*sheets.Request{
		MultiUpdateReqDate(dates, 1, 0, 1472288449),
		MultiUpdateReqNum(amounts, 1, 1, 1472288449),
		MultiUpdateReq(descriptions, 1, 2, 1472288449),
		MultiUpdateReqNum(balances, 1, 3, 1472288449),
	}
}
