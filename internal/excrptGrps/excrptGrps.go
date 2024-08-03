package excrptgrps

import (
	"budgetAutomation/internal/requests"
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

var excrptGrpTotals = map[string]float64{}

var excrptGrps = []ExcrptGrpParent{}

var resume = []string{}

type DataExcrpt struct {
	Matches      []string
	FixedExpense bool
}

// Marshal and unmarshal json
type Data struct {
	ExcrptMappings map[string]map[string]DataExcrpt
}

type ExcrptGrp struct {
	// Used to make lookup in excerptMappings array
	Ind int

	// Name of the ExcrptGrp
	Name string

	// Matches for this excrptGrp
	Mappings []string

	// Defines the type of this excerpt
	Parent string

	// Determines if the initial group total value should be read from the sheet or start at 0
	// Default is false
	FixedExpense bool
}

type ExcrptGrpParent struct {
	Name       string
	ExcrptGrps []ExcrptGrp
}

func isIgnored(parentName string) bool {
	return parentName == "Ignored"
}

func updateExcrptTotal(date, excrpt string, amount float64) {
	var excrptGrpMatches []ExcrptGrp
	ignored := false

	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	// Find correct excrpt grp
	for _, parent := range excrptGrps {
		for i := range parent.ExcrptGrps {
			for _, match := range parent.ExcrptGrps[i].Mappings {
				match = strings.ToLower(strings.Trim(match, " "))
				if strings.Contains(excrpt, match) {
					excrptGrpMatches = append(excrptGrpMatches, parent.ExcrptGrps[i])
					if isIgnored(parent.Name) {
						ignored = true
						break
					}
				}
			}
		}
	}

	if !ignored {

		excrptGrpName := selMatchGrp(date, excrpt, amount, excrptGrpMatches)

		excrptGrpTotals[excrptGrpName] += float64(amount)
		UpdateResume(date, excrpt, excrptGrpName, amount)
	} else {
		UpdateResume(date, excrpt, excrptGrpMatches[0].Name, amount)
	}
}

/*
Select excrptGrp for given match
*/
func selMatchGrp(date, excrpt string, amount float64, excrptGrpMatches []ExcrptGrp) string {
	grp := ""

	if len(excrptGrpMatches) == 0 {

		// Choose group to match
		ind := -1
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		validInd := false
		for !validInd {
			fmt.Scan(&ind)
			if ind > -1 && ind <= len(excrptGrpTotals) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}
		for _, parent := range excrptGrps {
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
			if ind > -1 && ind <= len(excrptGrpTotals) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}

		for _, parent := range excrptGrps {
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

func UpdateResume(date, excrpt, excrptGrpName string, amount float64) {
	resume = append(resume, date+" "+excrpt+" "+strconv.FormatFloat(amount, 'f', -1, 64)+": "+excrptGrpName)
}

func PrintExcrptGrpTotals() {
	fmt.Println("###################################################")
	for k, v := range excrptGrpTotals {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func PrintExcrptGrps() {
	fmt.Println("Excerpt groups")
	fmt.Println("###################################################")
	for _, parent := range excrptGrps {
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
	for _, str := range resume {
		fmt.Println(str)
	}
	fmt.Println("###################################################")
}

func GetTotal(excrptGrpName string) (float64, error) {
	excrptGrp, err := GetExcrptGrp(excrptGrpName, -1)
	//
	if err != nil {
		return 0, err
	}

	// Total should always be a positive number
	total := excrptGrpTotals[excrptGrp.Name] + 1
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
	json.NewDecoder(f1).Decode(&excrptGrps) // if err != nil {

	f2, err := os.Open("build/debug/JsonExcrptGrpTotals")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptsGrpTotals")
	}
	defer f2.Close()                             // //Json decode
	json.NewDecoder(f2).Decode(&excrptGrpTotals) // if err != nil {
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

	PrintExcrptGrpTotals()
	PrintExcrptGrps()

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
						excrptGrpTotals[excrptGrp.Name] += 0.0
					} else {

						val := strings.Trim(excrpts.Values[0][0].(string), " ")
						if len(val) > 0 {
							amount, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(val[:len(val)-4], ".", ""), ",", "."), 64)
							if err != nil {
								log.Fatal(err)
							}
							// updateExcrptTotal("9999-99-99", excrptGrp.name, amount)
							excrptGrpTotals[excrptGrp.Name] += -1 * float64(amount)
						} else {
							excrptGrpTotals[excrptGrp.Name] += 0.0
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
			excrptGrpTotals[excrptGrp] = -1.0
			grps = append(grps,
				createExcrptGrp(i, excrptGrp, parentName, mappings))
			i += 1
		}
		parent := ExcrptGrpParent{parentName, grps}
		excrptGrps = append(excrptGrps, parent)
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
	return excrptGrps
}

/*
Get excerpt group based on name OR index(ind).
Both can be specified, but it's not necessary.
If you don't want to use ind, make ind < 0.
if you don't want to use name, make name = "". \n
returnes err if not found
*/
func GetExcrptGrp(name string, ind int) (ExcrptGrp, error) {
	for _, parent := range excrptGrps {
		for _, excrptGrp := range parent.ExcrptGrps {
			if strings.EqualFold(strings.Trim(excrptGrp.Name, " "), strings.Trim(name, " ")) || excrptGrp.Ind == ind {
				return excrptGrp, nil
			}
		}
	}
	return ExcrptGrp{}, errors.New("Excrpt grp: " + name + ", not found")
}

func GetChildren(parentName string) []ExcrptGrp {
	for _, egp := range excrptGrps {
		if egp.Name == parentName {
			return egp.ExcrptGrps
		}
	}
	return nil
}

func UpdateExcrptSheet(path string) []*sheets.Request {
	dates, amounts, descriptions, balances := util.ReadExcrptCsv(path)

	return []*sheets.Request{
		requests.MultiUpdateReqDate(dates, 1, 0, 1472288449),
		requests.MultiUpdateReqNum(amounts, 1, 1, 1472288449),
		requests.MultiUpdateReq(descriptions, 1, 2, 1472288449),
		requests.MultiUpdateReqNum(balances, 1, 3, 1472288449),
	}
}

/*
Reads excrpts to update excrptGrpTotals and returns most recent account balance
*/
func LoadExcrptTotal(excrpts *sheets.ValueRange, month int64) float64 {
	isRightMonth := false
	accBalance := -1.0
	for _, elm := range excrpts.Values {
		date, description := elm[0].(string), elm[2].(string)
		s := strings.ReplaceAll(elm[1].(string), ",", ".")
		amount, err := strconv.ParseFloat(s, 64)

		if err != nil {
			log.Println("Could not read amount for", date, ":", description, ":", elm[1].(string))
		} else {
			// Get excerpt month
			if date != "Reserveret" {

				exrptMonth, err := strconv.ParseInt(strings.Split(date, "/")[1], 0, 64)
				if err != nil {
					log.Fatal("Could not read excerpt date", err)
				}

				if month == exrptMonth {
					isRightMonth = true
					if accBalance == -1.0 {
						s := strings.ReplaceAll(elm[3].(string), ",", ".")

						accBalance, err = strconv.ParseFloat(s, 64)
						if err != nil {
							log.Fatalln("Could not read account balance")
						}
					}

				} else if exrptMonth < month {
					break
				}
				if isRightMonth {
					updateExcrptTotal(date, description, amount)
				} // else {
				// excrptgrps.UpdateResume(date, description, "Not handled", amount)
				//	}
			}
		}
	}
	return accBalance
}
