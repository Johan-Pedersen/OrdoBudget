package cli

import (
	"budgetAutomation/internal/accounting"
	"budgetAutomation/internal/parser"
	req "budgetAutomation/internal/request"
	"budgetAutomation/internal/util"
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/api/sheets/v4"
)

func UpdateBudgetReqs(rows *sheets.ValueRange, accBalance float64, month, person int64) []*sheets.Request {
	var updateReqs []*sheets.Request

	for i, elm := range rows.Values {
		if len(elm) != 0 {

			balance, notFoundErr := accounting.GetBalance(elm[0].(string))

			if notFoundErr == nil {
				if balance != 0.0 {
					updateReqs = append(updateReqs, req.SingleUpdateReq(balance, int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
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

func GetExcrpts() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.SpreadSheetId, readRangeExrpt).Do()
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
	sheetsGrpCol, err := sheet.Values.Get(req.SpreadSheetId, budgetColARange).Do()
	if err != nil {
		log.Fatalf("Unable to perform get: %v", err)
	}
	return sheetsGrpCol
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
	_, updateBudgetErr := sheet.BatchUpdate(req.SpreadSheetId, batchUpdateReq).Context(ctx).Do()

	if updateBudgetErr != nil {
		log.Fatalf("Unable to perform update operation: %v", updateBudgetErr)
	}
	log.Println("Data moved successfully!")
}

/*
Select excrptGrp for given match

*/
func selMatchGrp(date, excrpt string, amount float64, excrptGrpMatches []accounting.Entry) string {
	grp := ""

	if len(excrptGrpMatches) == 0 {

		// Choose group to match
		ind := -1
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		validInd := false
		for !validInd {
			fmt.Scan(&ind)
			if ind > -1 && ind <= len(accounting.Balances) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}
		for _, parent := range accounting.Groups {
			for _, excrptGrp := range parent.Entries {
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
			if ind > -1 && ind <= len(accounting.Balances) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}

		for _, parent := range accounting.Groups {
			for _, excrptGrp := range parent.Entries {
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

func PrintBalances() {
	fmt.Println("###################################################")
	for k, v := range accounting.Balances {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func PrintEntries() {
	fmt.Println("Excerpt groups")
	fmt.Println("###################################################")
	for _, parent := range accounting.Groups {
		fmt.Println("\n************", parent.Name, "************")
		for _, excrptGrp := range parent.Entries {
			fmt.Println(excrptGrp.Ind, ":", excrptGrp.Name)
		}
	}
	fmt.Println("###################################################")
}

func PrintResume() {
	fmt.Println("Resume")
	fmt.Println("###################################################")
	for _, str := range accounting.Resume {
		fmt.Println(str)
	}
	fmt.Println("###################################################")
}

/*
Decide entries / matchs of the excrpts belonging to multiple or none, pre-defined entries
*/
func DecideEntries(allMatches map[parser.Excrpt][]accounting.Entry) {
	var entry accounting.Entry
	var entryErr error

	for excrpt, matches := range allMatches {
		if len(matches) == 0 {
			PrintEntries()
			fmt.Println("No matches found for ", excrpt, ". Please choose a match")

		} else {
			printFoundEntries(matches)
			fmt.Println(len(matches), " matches found for ", excrpt, ". Please choose a match")
		}

		reader := bufio.NewReader(os.Stdin)
		validInd := false
		for !validInd {
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)

			// Check if user wants to exit
			if strings.ToLower(input) == "exit" {
				os.Exit(0)
			}

			num, err := strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input. Please enter an integer.")
			} else {
				entry, entryErr = accounting.GetEntry("", num)
				if entryErr == nil {
					validInd = true
				} else {
					// Man gaar bare videre til naeste match uden at update noget
					log.Println("Could not find entry.\nPlease choose again")
				}
			}
		}
		accounting.UpdateBalance(excrpt.Date, excrpt.Description, excrpt.Amount, entry.Name)
	}
}

func printFoundEntries(entries []accounting.Entry) {
	fmt.Println("Excerpt groups")
	fmt.Println("###################################################")
	for _, excrptGrp := range entries {
		fmt.Println(excrptGrp.Ind, ":", excrptGrp.Name)
	}
	fmt.Println("###################################################")
}
