package ui

import (
	excrptgrps "budgetAutomation/internal/excrptGrps"
	req "budgetAutomation/internal/requests"
	"budgetAutomation/internal/util"
	"context"
	"log"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func submit(month int64, excrptPath string) {
	var person int64 = 1

	updateExcrptSheet(excrptPath)

	budgetColA := getBudgetColA()
	excrptgrps.InitExcrptGrps(budgetColA, month, person)

	excrptsFromSheets := getExcrptsFromSheet()

	// Denne del skal jo kores i et andet window
	accBalance := excrptgrps.LoadExcrptTotal(excrptsFromSheets, month)

	updateReqs := updateBudgetReqs(budgetColA, accBalance, month, person)
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
func getBudgetColA() *sheets.ValueRange {
	// Read all rows of col A in budget sheet.
	budgetColARange := "A1:A"

	sheet := req.GetSheet()
	budgetColA, err := sheet.Values.Get(req.GetSpreadsheetId(), budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	return budgetColA
}

/*
Update the temporary sheet to hold excrpts
*/
func updateExcrptSheet(excrptPath string) {
	sheet := req.GetSheet()
	ctx := context.Background()

	// Update excerpt sheet, before we begin
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: excrptgrps.UpdateExcrptSheet(excrptPath),
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

			total, notFoundErr := excrptgrps.GetTotal(elm[0].(string))

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
