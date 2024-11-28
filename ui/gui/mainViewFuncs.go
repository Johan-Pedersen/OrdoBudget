package ui

import (
	"OrdoBudget/internal/accounting"
	"OrdoBudget/internal/logtrace"
	req "OrdoBudget/internal/request"
	"OrdoBudget/internal/util"
	"context"
	"encoding/json"
	"os"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func submitDebug() {
	var person int64 = 1
	var month int64 = 6

	// excrpts := debugGetExcrpts()
	debugGetExcrpts()
	accounting.InitGrpsDebug()

	sheetsGrpCol := getSheetsGrpCol()

	// Denne del skal jo kores i et andet window
	// accBalance := excrptgrps.LoadExcrptTotal(excrpts, month)
	accBalance := 0.0

	updateReqs := updateBudgetReqs(sheetsGrpCol, accBalance, month, person)
	batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: updateReqs,
	}
	// Execute the BatchUpdate request
	sheet := req.GetSheet()
	ctx := context.Background()
	_, updateBudgetErr := sheet.BatchUpdate(req.SpreadSheetId, batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		logtrace.Error(updateBudgetErr.Error())
	}
}

func submit(month int64, excrptPath string) {
	var person int64 = 1

	sheetsGrpCol := getSheetsGrpCol()
	accounting.InitGrps(sheetsGrpCol, month, person)

	// excrptsFromSheets := getExcrptsFromSheet()

	// Denne del skal jo kores i et andet window
	// accBalance := excrptgrps.LoadExcrptTotal(excrptsFromSheets, month)
	accBalance := 0.0

	updateReqs := updateBudgetReqs(sheetsGrpCol, accBalance, month, person)
	batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: updateReqs,
	}
	// Execute the BatchUpdate request
	sheet := req.GetSheet()
	ctx := context.Background()
	_, updateBudgetErr := sheet.BatchUpdate(req.SpreadSheetId, batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		logtrace.Error(updateBudgetErr.Error())
	}
}

/*
Get all excrptGrps from the speadsheet
*/
func getSheetsGrpCol() *sheets.ValueRange {
	// Read all rows of col A in budget sheet.
	sheetsGrpColRange := "A1:A"

	sheet := req.GetSheet()
	sheetsGrpCol, err := sheet.Values.Get(req.SpreadSheetId, sheetsGrpColRange).Do()
	if err != nil {
		logtrace.Error(err.Error())
	}
	return sheetsGrpCol
}

func getExcrptsFromSheet() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.SpreadSheetId, readRangeExrpt).Do()
	if readExcrptsErr != nil {
		logtrace.Error(readExcrptsErr.Error())
	}
	return excrpts
}

func updateBudgetReqs(rows *sheets.ValueRange, accBalance float64, month, person int64) []*sheets.Request {
	var updateReqs []*sheets.Request

	for i, elm := range rows.Values {
		if len(elm) != 0 {

			total, notFoundErr := accounting.GetBalance(elm[0].(string))

			if notFoundErr == nil {
				if total != 0.0 {
					updateReqs = append(updateReqs, req.SingleUpdateReq(total, int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
				} else {
					updateReqs = append(updateReqs, req.SingleUpdateReqBlank(int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
				}
			} else if strings.EqualFold(strings.Trim(elm[0].(string), " "), "Faktisk balance") {
				updateReqs = append(updateReqs, req.SingleUpdateReq(accBalance, int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
			}
		}
	}
	return updateReqs
}

func debugGetExcrpts() *sheets.ValueRange {
	// Read all rows of col A in budget sheet.

	var excrpts *sheets.ValueRange
	f3, err := os.Open("build/debug/JsonExcrptSheets")
	if err != nil {
		logtrace.Error(err.Error())
	}
	defer f3.Close()                     // //Json decode
	json.NewDecoder(f3).Decode(&excrpts) // if err != nil {

	return excrpts
}
