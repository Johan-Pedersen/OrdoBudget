package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	excrptgrps "budgetAutomation/src/excrptGrps"
	util "budgetAutomation/src/util"

	req "budgetAutomation/src/requests"

	"google.golang.org/api/sheets/v4"
)

func main() {
	sheet := req.GetSheet()
	ctx := context.Background()

	// Read all rows of col A in budget sheet.
	budgetColARange := "A1:A"
	budgetColA, err := sheet.Values.Get(req.GetSpreadsheetId(), budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	// Update excerpt sheet, before we begin
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: excrptgrps.UpdateExcrptSheet("Lønkonto.csv"),
	}

	_, excrptUpdateErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateExcerptSheetReq).Context(ctx).Do()

	if excrptUpdateErr != nil {
		log.Fatalf("Unable to perform update excerpt sheet operation: %v", excrptUpdateErr)
	}

	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()
	if readExcrptsErr != nil {
		log.Fatalf("Unable to perform get: %v", readExcrptsErr)
	}

	// Initial API reqeusts to get Data
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Which month from 1-12 should be handled
	var month int64 = -1
	fmt.Println("Specify month:")
	fmt.Scan(&month)

	// Initialize and print excerpt groups
	excrptgrps.InitExcrptGrps(budgetColA, month)
	excrptgrps.PrintExcrptGrps()

	accBalance := excrptgrps.LoadExcrptTotal(excrpts, month)

	// find Excerpt Total for current month.
	excrptgrps.PrintExcrptGrpTotals()

	// Find excerpt grps to insert at

	updateReqs := updateBudgetReqs(budgetColA, accBalance, month)
	batchUpdateReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: updateReqs,
	}

	// Execute the BatchUpdate request
	_, updateBudgetErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		log.Fatalf("Unable to perform update operation: %v", updateBudgetErr)
	}
	log.Println("Data moved successfully!")

	excrptgrps.PrintResume()
}

func updateBudgetReqs(rows *sheets.ValueRange, accBalance float64, month int64) []*sheets.Request {
	var updateReqs []*sheets.Request

	for i, elm := range rows.Values {
		if len(elm) != 0 {

			total, notFoundErr := excrptgrps.GetTotal(elm[0].(string))

			if notFoundErr == nil {
				if total != 0.0 {
					updateReqs = append(updateReqs, req.SingleUpdateReq(total, int64(i), util.MonthToColInd(month), 1685114351))
				} else {
					updateReqs = append(updateReqs, req.SingleUpdateReqBlank(int64(i), util.MonthToColInd(month), 1685114351))
				}
			} else if strings.EqualFold(strings.Trim(elm[0].(string), " "), "Faktisk balance") {
				updateReqs = append(updateReqs, req.SingleUpdateReq(accBalance, int64(i), util.MonthToColInd(month), 1685114351))
			}
		}
	}
	return updateReqs
}
