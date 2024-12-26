package cli

import (
	"OrdoBudget/internal/accounting"
	"OrdoBudget/internal/logtrace"
	"OrdoBudget/internal/parse"
	req "OrdoBudget/internal/request"
	"OrdoBudget/internal/util"
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

			// balance, notFoundErr := accounting.GetBalance(elm[0].(string))

			// amounts := accounting.Balances[elm[0].(string)]

			// "Faktisk balance", is the last field in the budget to fill out, which means we can break after

			if strings.EqualFold(strings.Trim(elm[0].(string), " "), "Faktisk balance") {
				updateReqs = append(updateReqs, req.SingleUpdateReq(accBalance, int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
				break
			}

			amounts, err := accounting.GetAmounts(elm[0].(string))

			if err == nil {
				if len(amounts) == 0 {
					updateReqs = append(updateReqs, req.SingleUpdateReqBlank(int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
				} else {
					eq := toSumEq(amounts)
					updateReqs = append(updateReqs, req.SingleUpdateReqSum(eq, int64(i), util.MonthToColInd(month, person), req.BudgetSheetId))
				}
			}

		}

	}
	return updateReqs
}

/*
Returns all amounts to a sum equation used in google sheets
*/
func toSumEq(amounts []float64) string {

	sum_str := ""
	if len(amounts) != 0 {

		sum_str = "="
		for i := range len(amounts) {
			amount := strings.ReplaceAll(strconv.FormatFloat(amounts[i], 'f', -1, 64), ".", ",")
			sum_str = sum_str + amount + "+"

		}

		//remove last '+' from equation
		sum_str = sum_str[:len(sum_str)-1]
	}
	return sum_str
}

func GetExcrpts() *sheets.ValueRange {
	sheet := req.GetSheet()
	// Get Date, Amount and description
	readRangeExrpt := "Udtrœk!A2:D"
	excrpts, readExcrptsErr := sheet.Values.Get(req.SpreadSheetId, readRangeExrpt).Do()
	if readExcrptsErr != nil {
		logtrace.Error(readExcrptsErr.Error())
	}
	return excrpts
}

func DebugGetExcrpts() *sheets.ValueRange {
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

func GetSheetsGrpCol() *sheets.ValueRange {
	sheet := req.GetSheet()
	budgetColARange := "A1:A"
	sheetsGrpCol, err := sheet.Values.Get(req.SpreadSheetId, budgetColARange).Do()
	if err != nil {
		logtrace.Error(err.Error())
	}
	return sheetsGrpCol
}

func InputPerson(person *int64) {
	fmt.Println("Which person is doing the budget: 1 or 2")
	fmt.Scan(person)
}

func InputMonth(month *int64) {
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
		logtrace.Error(updateBudgetErr.Error())
	}
	logtrace.Info("Data moved successfully!")
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
	fmt.Println("Balances")
	fmt.Println("###################################################")
	for k := range accounting.Balances {
		balance, err := accounting.GetBalance(k)

		if err != nil {
			logtrace.Info(err.Error())
		} else {

			fmt.Println(k, ": ", balance)
		}
	}
	fmt.Println("###################################################")
}

func PrintEntries() {
	fmt.Println("Groups")
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
func DecideEntries(allMatches map[parse.Excrpt][]accounting.Entry) {
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
