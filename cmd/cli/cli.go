package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	excrptgrps "budgetAutomation/internal/excrptGrps"
	util "budgetAutomation/internal/util"

	req "budgetAutomation/internal/requests"

	"google.golang.org/api/sheets/v4"
)

func main() {
	debugMode := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()

	// Who is doing the budget
	var person int64

	// Which month from 1-12 should be handled
	var month int64

	var excrpts *sheets.ValueRange
	sheetsGrpCol := getSheetsGrpCol()
	// Update excerpt sheet, before we begin
	updateExcrptsSheet()
	// Debug mode
	if *debugMode {

		// hard code person + month
		person = 1
		month = 6

		excrpts = debugGetExcrpts()

		// Initialize and print excerpt groups
		excrptgrps.InitExcrptGrpsDebug()

	} else {

		getPersonAndMonth(&person, &month)
		excrpts = getExcrpts()
		// Initialize and print excerpt groups
		excrptgrps.InitExcrptGrps(sheetsGrpCol, month, person)
		excrptgrps.PrintExcrptGrps()

	}
	accBalance := excrptgrps.LoadExcrptTotal(excrpts, month)

	// find Excerpt Total for current month.
	excrptgrps.PrintExcrptGrpTotals()
	updateBudget(sheetsGrpCol, accBalance, month, person)
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

func getExcrpts() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()
	if readExcrptsErr != nil {
		log.Fatalf("Unable to perform get: %v", readExcrptsErr)
	}
	return excrpts
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

func getSheetsGrpCol() *sheets.ValueRange {
	sheet := req.GetSheet()
	budgetColARange := "A1:A"
	sheetsGrpCol, err := sheet.Values.Get(req.GetSpreadsheetId(), budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	return sheetsGrpCol
}

func updateExcrptsSheet() {
	sheet := req.GetSheet()
	ctx := context.Background()
	batchUpdateExcerptSheetReq := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: excrptgrps.UpdateExcrptSheet("storage/excrptSheet.csv"),
	}

	_, excrptUpdateErr := sheet.BatchUpdate(req.GetSpreadsheetId(), batchUpdateExcerptSheetReq).Context(ctx).Do()

	if excrptUpdateErr != nil {
		log.Fatalf("Unable to perform update excerpt sheet operation: %v", excrptUpdateErr)
	}
}

func getPersonAndMonth(person, month *int64) {
	fmt.Println("Which person is doing the budget: 1 or 2")
	fmt.Scan(person)

	// Which month from 1-12 should be handled
	fmt.Println("Specify month:")
	fmt.Scan(month)
}

func updateBudget(sheetsGrpCol *sheets.ValueRange, accBalance float64, month, person int64) {
	sheet := req.GetSheet()
	ctx := context.Background()
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
}
