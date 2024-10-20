package request

import (
	"log"
	"strconv"

	"google.golang.org/api/sheets/v4"
)

/*
Updates row-wise from (rowInd, colInd) to (rowInd + len(grpSums), colInd)
*/
func MultiUpdateReq(data []string, rowInd, colInd int64, sheetId string) *sheets.Request {
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

	intSheetId, err := strconv.ParseInt(sheetId, 10, 64)
	if err != nil {
		log.Fatalln("Cannot convert sheetId to int")
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     intSheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
Updates row-wise from (rowInd, colInd) to (rowInd + len(grpSums), colInd)
*/
func MultiUpdateReqNum(data []float64, rowInd, colInd int64, sheetId string) *sheets.Request {
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
	intSheetId, err := strconv.ParseInt(sheetId, 10, 64)
	if err != nil {
		log.Fatalln("Cannot convert sheetId to int")
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     intSheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
Updates row-wise from (rowInd, colInd) to (rowInd + len(grpSums), colInd)
*/
func MultiUpdateReqDate(data []float64, rowInd, colInd int64, sheetId string) *sheets.Request {
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
	intSheetId, err := strconv.ParseInt(sheetId, 10, 64)
	if err != nil {
		log.Fatalln("Cannot convert sheetId to int")
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "*",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     intSheetId,
			},
			Rows: rowData,
		},
	}
	return updateReq
}

/*
Update single cell at (rowInd, colInd) with blank
*/

func SingleUpdateReqBlank(rowInd, colInd int64, sheetId string) *sheets.Request {
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
	intSheetId, err := strconv.ParseInt(sheetId, 10, 64)
	if err != nil {
		log.Fatalln("Cannot convert sheetId to int")
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "UserEnteredValue",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     intSheetId,
			},
			Rows: []*sheets.RowData{rowData},
		},
	}
	return updateReq
}

/*
Update single cell at (rowInd, colInd) with amount
*/

func SingleUpdateReq(amount float64, rowInd, colInd int64, sheetId string) *sheets.Request {
	rowData := &sheets.RowData{
		Values: []*sheets.CellData{
			{
				UserEnteredValue: &sheets.ExtendedValue{
					NumberValue: &amount,
				},
			},
		},
	}
	intSheetId, err := strconv.ParseInt(sheetId, 10, 64)
	if err != nil {
		log.Fatalln("Cannot convert sheetId to int")
	}
	updateReq := &sheets.Request{
		UpdateCells: &sheets.UpdateCellsRequest{
			Fields: "UserEnteredValue",
			Start: &sheets.GridCoordinate{
				ColumnIndex: colInd,
				RowIndex:    rowInd,
				SheetId:     intSheetId,
			},
			Rows: []*sheets.RowData{rowData},
		},
	}
	return updateReq
}
