package ui

import (
	"budgetAutomation/internal/accounting"
	req "budgetAutomation/internal/requests"
	"budgetAutomation/internal/util"
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func submitDebug() {
	var person int64 = 1
	var month int64 = 6

	// excrpts := debugGetExcrpts()
	debugGetExcrpts()
	updateExcrptsSheetDebug()
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
	_, updateBudgetErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		log.Fatalf("Unable to perform update operation: %v", updateBudgetErr)
	}
}

func submit(month int64, excrptPath string) {
	var person int64 = 1

	updateExcrptSheet(excrptPath, month)

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
	_, updateBudgetErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		log.Fatalf("Unable to perform update operation: %v", updateBudgetErr)
	}
}

/*
Get all excrptGrps from the speadsheet
*/
func getSheetsGrpCol() *sheets.ValueRange {
	// Read all rows of col A in budget sheet.
	sheetsGrpColRange := "A1:A"

	sheet := req.GetSheet()
	sheetsGrpCol, err := sheet.Values.Get(req.GetSpreadsheetId(), sheetsGrpColRange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	return sheetsGrpCol
}

/*
Update the temporary sheet to hold excrpts
*/
func updateExcrptSheet(excrptPath string, month int64) {
	sheet := req.GetSheet()
	ctx := context.Background()

	// Update excerpt sheet, before we begin
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: req.UpdateExcrptSheet(excrptPath, month),
	}

	_, excrptUpdateErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateExcerptSheetReq).Context(ctx).Do()

	if excrptUpdateErr != nil {
		log.Fatalf("Unable to perform update excerpt sheet operation: %v", excrptUpdateErr)
	}
}

func getExcrptsFromSheet() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()
	if readExcrptsErr != nil {
		log.Fatalf("Unable to perform get: %v", readExcrptsErr)
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
					updateReqs = append(updateReqs, req.SingleUpdateReq(total, int64(i), util.MonthToColInd(month, person), 1685114351))
				} else {
					updateReqs = append(updateReqs, req.SingleUpdateReqBlank(int64(i), util.MonthToColInd(month, person), 1685114351))
				}
			} else if strings.EqualFold(strings.Trim(elm[0].(string), " "), "Faktisk balance") {
				updateReqs = append(updateReqs, req.SingleUpdateReq(accBalance, int64(i), util.MonthToColInd(month, person), 1685114351))
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
		log.Fatal("Unable to open JSonExcrptSheets")
	}
	defer f3.Close()                     // //Json decode
	json.NewDecoder(f3).Decode(&excrpts) // if err != nil {

	return excrpts
}

func updateExcrptsSheetDebug() {
	sheet := req.GetSheet()
	ctx := context.Background()
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: req.UpdateExcrptSheet("storage/excrptSheet.csv", 4),
	}

	_, excrptUpdateErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateExcerptSheetReq).Context(ctx).Do()

	if excrptUpdateErr != nil {
		log.Fatalf("Unable to perform update excerpt sheet operation: %v", excrptUpdateErr)
	}
}
