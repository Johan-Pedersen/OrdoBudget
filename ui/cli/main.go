package cli

import (
	excrptgrps "budgetAutomation/internal/excrptGrps"
	req "budgetAutomation/internal/requests"
	"budgetAutomation/internal/util"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func UpdateBudgetReqs(rows *sheets.ValueRange, accBalance float64, month, person int64) []*sheets.Request {
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

func GetExcrpts() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()
	if readExcrptsErr != nil {
		log.Fatalf("Unable to perform get: %v", readExcrptsErr)
	}
	return excrpts
}

func DebugGetExcrpts() *sheets.ValueRange {
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

func GetSheetsGrpCol() *sheets.ValueRange {
	sheet := req.GetSheet()
	budgetColARange := "A1:A"
	sheetsGrpCol, err := sheet.Values.Get(req.GetSpreadsheetId(), budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	return sheetsGrpCol
}

func UpdateExcrptsSheet() {
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

func GetPersonAndMonth(person, month *int64) {
	fmt.Println("Which person is doing the budget: 1 or 2")
	fmt.Scan(person)

	// Which month from 1-12 should be handled
	fmt.Println("Specify month:")
	fmt.Scan(month)
}

func UpdateBudget(sheetsGrpCol *sheets.ValueRange, accBalance float64, month, person int64) {
	sheet := req.GetSheet()
	ctx := context.Background()
	// Find excerpt grps to insert at

	updateReqs := UpdateBudgetReqs(sheetsGrpCol, accBalance, month, person)
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

/*
Select excrptGrp for given match

*/
func selMatchGrp(date, excrpt string, amount float64, excrptGrpMatches []excrptgrps.ExcrptGrp) string {
	grp := ""

	if len(excrptGrpMatches) == 0 {

		// Choose group to match
		ind := -1
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		validInd := false
		for !validInd {
			fmt.Scan(&ind)
			if ind > -1 && ind <= len(excrptgrps.ExcrptGrpTotals) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}
		for _, parent := range excrptgrps.ExcrptGrps {
			for _, excrptGrp := range parent.ExcrptGrps {
				if excrptGrp.Ind == ind {
					grp = excrptGrp.Name
					break
				}
			}

			if grp != "" {
				break
			}

		}
	} else if len(excrptGrpMatches) > 1 {

		// Choose group to match
		ind := -1
		fmt.Println("Matches multiple groups:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		validInd := false

		for _, v := range excrptGrpMatches {
			fmt.Println(v.Ind, ":", v.Name)
		}
		for !validInd {
			fmt.Scan(&ind)
			if ind > -1 && ind <= len(excrptgrps.ExcrptGrpTotals) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}

		for _, parent := range excrptgrps.ExcrptGrps {
			for _, excrptGrp := range parent.ExcrptGrps {
				if excrptGrp.Ind == ind {
					grp = excrptGrp.Name
					break
				}
			}

			if grp != "" {
				break
			}

		}
	} else {
		grp = excrptGrpMatches[0].Name
	}
	return grp
}

func PrintExcrptGrpTotals() {
	fmt.Println("###################################################")
	for k, v := range excrptgrps.ExcrptGrpTotals {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func PrintExcrptGrps() {
	fmt.Println("Excerpt groups")
	fmt.Println("###################################################")
	for _, parent := range excrptgrps.ExcrptGrps {
		fmt.Println("\n************", parent.Name, "************")
		for _, excrptGrp := range parent.ExcrptGrps {
			fmt.Println(excrptGrp.Ind, ":", excrptGrp.Name)
		}
	}
	fmt.Println("###################################################")
}

func PrintResume() {
	fmt.Println("Resume")
	fmt.Println("###################################################")
	for _, str := range excrptgrps.Resume {
		fmt.Println(str)
	}
	fmt.Println("###################################################")
}
