package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"strings"

	excrptgrps "budgetAutomation/internal/excrptGrps"
	util "budgetAutomation/internal/util"

	req "budgetAutomation/internal/requests"

	"google.golang.org/api/sheets/v4"
)

func main() {
	sheet := req.GetSheet()
	ctx := context.Background()

	// Read all rows of col A in budget sheet.
	budgetColARange := "A1:A"
	sheetsGrpCol, err := sheet.Values.Get(req.GetSpreadsheetId(), budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	// Update excerpt sheet, before we begin
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: excrptgrps.UpdateExcrptSheet("storage/excrptSheet.csv"),
	}

	_, excrptUpdateErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateExcerptSheetReq).Context(ctx).Do()

	if excrptUpdateErr != nil {
		log.Fatalf("Unable to perform update excerpt sheet operation: %v", excrptUpdateErr)
	}

	var excrpts *sheets.ValueRange
	f3, err := os.Open("build/debug/JsonExcrptSheets")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptSheets")
	}
	defer f3.Close()                     // //Json decode
	json.NewDecoder(f3).Decode(&excrpts) // if err != nil {

	// Initial API reqeusts to get Data
	////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

	// Who is doing the budget
	var person int64 = 1

	// Which month from 1-12 should be handled
	var month int64 = 6

	// Initialize and print excerpt groups
	excrptgrps.InitExcrptGrpsDebug()
	// excrptgrps.InitExcrptGrps(sheetsGrpCol, month, person)
	excrptgrps.PrintExcrptGrps()

	// 	return nil, err
	// }
	// defer f.Close()
	// tok := &oauth2.Token{}
	// err = json.NewDecoder(f).Decode(tok)

	////Json decode
	// f1, err := os.OpenFile("JsonExcrptGrps", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600)
	// if err != nil {
	// 	log.Fatal("Unable to open JSonExcrptsGrps")
	// }
	// defer f1.Close()
	// json.NewEncoder(f1).Encode(excrptgrps)

	accBalance := excrptgrps.LoadExcrptTotal(excrpts, month)

	// find Excerpt Total for current month.
	excrptgrps.PrintExcrptGrpTotals()

	// Find excerpt grps to insert at

	updateReqs := updateBudgetReqs(sheetsGrpCol, accBalance, month, person)
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
