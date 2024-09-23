package excrptgrps

import (
	"budgetAutomation/internal/excrpt"
	req "budgetAutomation/internal/requests"
	"budgetAutomation/internal/util"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"google.golang.org/api/sheets/v4"
)

var ExcrptGrpTotals = map[string]float64{}

var ExcrptGrps = []ExcrptGrpParent{}

var Resume = []string{}

func isIgnored(parentName string) bool {
	return parentName == "Ignored"
}

func FindExcrptMatches(excrpt string) []ExcrptGrp {
	var excrptGrpMatches []ExcrptGrp

	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	// Find correct excrpt grp
	for _, parent := range ExcrptGrps {
		for i := range parent.ExcrptGrps {
			for _, match := range parent.ExcrptGrps[i].Mappings {
				match = strings.ToLower(strings.Trim(match, " "))
				if strings.Contains(excrpt, match) {
					if !isIgnored(parent.Name) {
						excrptGrpMatches = append(excrptGrpMatches, parent.ExcrptGrps[i])
					}
				}
			}
		}
	}
	return excrptGrpMatches
}

func UpdateExcrptTotal(date, excrpt string, amount float64, excrptGrpName string) {
	tmpAmount := amount

	if excrptGrpName == "Ignored" {
		tmpAmount = 0
	}

	ExcrptGrpTotals[excrptGrpName] += float64(tmpAmount)
	UpdateResume(date, excrpt, excrptGrpName, tmpAmount)
}

func UpdateResume(date, excrpt, excrptGrpName string, amount float64) {
	Resume = append(Resume, date+" "+excrpt+" "+strconv.FormatFloat(amount, 'f', -1, 64)+": "+excrptGrpName)
}

func GetTotal(excrptGrpName string) (float64, error) {
	excrptGrp, err := GetExcrptGrp(excrptGrpName, -1)
	//
	if err != nil {
		return 0, err
	}

	// Total should always be a positive number
	total := ExcrptGrpTotals[excrptGrp.Name] + 1
	if excrptGrp.Parent == "Indkomst efter skat" {
		return total, nil
	} else {
		return -1 * total, nil
	}
}

func InitExcrptGrpsDebug() {
	// Load ExcrptGrps

	f1, err := os.Open("build/debug/JsonExcrptGrps")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptsGrps")
	}
	defer f1.Close()                        // //Json decode
	json.NewDecoder(f1).Decode(&ExcrptGrps) // if err != nil {

	f2, err := os.Open("build/debug/JsonExcrptGrpTotals")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptsGrpTotals")
	}
	defer f2.Close()                             // //Json decode
	json.NewDecoder(f2).Decode(&ExcrptGrpTotals) // if err != nil {
}

func InitExcrptGrps(sheetsGrpCol *sheets.ValueRange, month, person int64) {
	// Open the JSON file
	file, err := os.Open("storage/excrptGrpData.json")
	if err != nil {
		log.Fatalln("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file
	jsonData, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("Error reading JSON data:", err)
		return
	}

	// Create an instance of the struct to hold the unmarshaled data
	var data Data

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatalln("Error unmarshaling JSON data:", err)
		return
	}

	// initialize excerpt grps with -1 as total
	createGrps(data)

	updateCommonGrps(sheetsGrpCol, month, person)
}

func updateCommonGrps(excrptGrps *sheets.ValueRange, month, person int64) {
	// Get Date, Amount and description

	A1Not := util.MonthToA1Notation(month, person)
	for i, elm := range excrptGrps.Values {
		if len(elm) != 0 {
			excrptGrp, notFound := GetExcrptGrp(elm[0].(string), -1)
			if notFound == nil {
				if excrptGrp.FixedExpense {

					readRangeExrpt := "budget!" + A1Not + fmt.Sprint(i+1)
					excrpts, readExcrptsErr := req.GetSheet().Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()

					if readExcrptsErr != nil {
						log.Fatalf("Unable to perform get: %v", readExcrptsErr)
					}

					if len(excrpts.Values) == 0 {
						ExcrptGrpTotals[excrptGrp.Name] += 0.0
					} else {

						val := strings.Trim(excrpts.Values[0][0].(string), " ")
						if len(val) > 0 {
							amount, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(val[:len(val)-4], ".", ""), ",", "."), 64)
							if err != nil {
								log.Fatal(err)
							}
							// updateExcrptTotal("9999-99-99", excrptGrp.name, amount)
							ExcrptGrpTotals[excrptGrp.Name] += -1 * float64(amount)
						} else {
							ExcrptGrpTotals[excrptGrp.Name] += 0.0
						}
					}
				}
			}
		}
	}
}

func createGrps(data Data) {
	i := 0
	// Init excrptGrp Totals
	for parentName, excrpts := range data.ExcrptMappings {

		// parentName = strings.Trim(parentName, " ")
		grps := []ExcrptGrp{}
		for excrptGrp, mappings := range excrpts {
			// excrptGrp = strings.Trim(excrptGrp, " ")
			ExcrptGrpTotals[excrptGrp] = -1.0
			grps = append(grps,
				createExcrptGrp(i, excrptGrp, parentName, mappings))
			i += 1
		}
		parent := ExcrptGrpParent{parentName, grps}
		ExcrptGrps = append(ExcrptGrps, parent)
	}
}

func createExcrptGrp(ind int, name, parent string, data DataExcrpt) ExcrptGrp {
	return ExcrptGrp{
		Ind:          ind,
		Name:         name,
		Mappings:     data.Matches,
		Parent:       parent,
		FixedExpense: data.FixedExpense,
	}
}

func GetParents() []ExcrptGrpParent {
	return ExcrptGrps
}

/*
Get excerpt group based on name OR index(ind).
Both can be specified, but it's not necessary.
If you don't want to use ind, make ind < 0.
if you don't want to use name, make name = "". \n
returnes err if not found
*/
func GetExcrptGrp(name string, ind int) (ExcrptGrp, error) {
	for _, parent := range ExcrptGrps {
		for _, excrptGrp := range parent.ExcrptGrps {
			if strings.EqualFold(strings.Trim(excrptGrp.Name, " "), strings.Trim(name, " ")) || excrptGrp.Ind == ind {
				return excrptGrp, nil
			}
		}
	}
	return ExcrptGrp{}, errors.New("Excrpt grp: " + name + ", not found")
}

func GetChildren(parentName string) []ExcrptGrp {
	for _, egp := range ExcrptGrps {
		if egp.Name == parentName {
			return egp.ExcrptGrps
		}
	}
	return nil
}

func UpdateExcrptSheet(path string, month int64) []*sheets.Request {
	// open the csv file
	file, err := os.Open(path)
	if err != nil {
		print("could not open excerpt file")
		log.Fatalln("coud not open excerpt file.", err)
	}
	defer file.Close()

	excrpts := util.ReadExcrptCsv(file, month)

	var dates []float64

	for _, exc := range excrpts {
		dates = append(dates, util.ConvertDateToFloat(exc.Date))
	}

	var amounts []float64
	for _, exc := range excrpts {
		amounts = append(amounts, exc.Amount)
	}

	var descriptions []string
	for _, exc := range excrpts {
		descriptions = append(descriptions, exc.Description)
	}

	var balances []float64
	for _, exc := range excrpts {
		balances = append(balances, exc.Balance)
	}

	return []*sheets.Request{
		req.MultiUpdateReqDate(dates, 1, 0, 1472288449),
		req.MultiUpdateReqNum(amounts, 1, 1, 1472288449),
		req.MultiUpdateReq(descriptions, 1, 2, 1472288449),
		req.MultiUpdateReqNum(balances, 1, 3, 1472288449),
	}
}

/*
Find matches for excrpts and updates ExcrptTotal iff only 1 match is found, otherwise the found excrpts are added to the return
*/
func FindUpdMatches(excrpts *[]excrpt.Excrpt) map[excrpt.Excrpt][]ExcrptGrp {
	ret := make(map[excrpt.Excrpt][]ExcrptGrp)

	for _, excrpt := range *excrpts {
		matches := FindExcrptMatches(excrpt.Description)
		// If there is only a single match, the update is given
		// Otherwise the correct match has to be made in the ui
		if len(matches) == 1 {
			UpdateExcrptTotal(excrpt.Date, excrpt.Description, excrpt.Amount, matches[0].Name)
		} else {
			ret[excrpt] = matches
		}
	}
	return ret
}
