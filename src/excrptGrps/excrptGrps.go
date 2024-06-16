package excrptgrps

import (
	"budgetAutomation/src/requests"
	"budgetAutomation/src/util"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	req "budgetAutomation/src/requests"

	"google.golang.org/api/sheets/v4"
)

var excrptGrpTotals = map[string]float64{}

var excrptGrps = []ExcrptGrpParent{}

var resume = []string{}

type DataExcrpt struct {
	Matches     []string
	IsCommonGrp bool
}

// Marshal and unmarshal json
type Data struct {
	ExcrptMappings map[string]map[string]DataExcrpt
}

type ExcrptGrp struct {
	// Used to make lookup in excerptMappings array
	ind int

	// Name of the ExcrptGrp
	name string

	// Matches for this excrptGrp
	mappings []string

	// Defines the type of this excerpt
	parent string

	// Determines if the initial group total value should be read from the sheet or start at 0
	// Default is false
	isCommonGrp bool
}

type ExcrptGrpParent struct {
	name       string
	excrptGrps []ExcrptGrp
}

func updateExcrptTotal(date, excrpt string, amount float64) {
	excrptGrpName := ""

	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	ind := -1
	// Find correct excrpt grp
	for _, parent := range excrptGrps {

		for i := range parent.excrptGrps {
			for _, match := range parent.excrptGrps[i].mappings {
				match = strings.ToLower(strings.Trim(match, " "))
				if strings.Contains(excrpt, match) {
					excrptGrpName = parent.excrptGrps[i].name
					break
				}
			}
			if excrptGrpName != "" {
				break
			}
		}

		if excrptGrpName != "" {
			break
		}
	}

	// No matches for group
	if excrptGrpName == "" {
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		validInd := false
		for !validInd {
			fmt.Scan(&ind)
			if ind > -1 && ind < len(excrptGrpTotals) {
				validInd = true
			} else {
				fmt.Println("Invalid grp number.\nPlease choose again")
			}
		}

		for _, parent := range excrptGrps {
			for _, excrptGrp := range parent.excrptGrps {
				if excrptGrp.ind == ind {
					excrptGrpName = excrptGrp.name
					break
				}
			}

			if excrptGrpName != "" {
				break
			}

		}
	}

	if excrptGrpName != "Ignored" {
		// Update correct excrpt grp
		excrptGrpTotals[excrptGrpName] += float64(amount)
	}

	// Update resume

	UpdateResume(date, excrpt, excrptGrpName, amount)
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
		fmt.Println("\n************", parent.name, "************")
		for _, excrptGrp := range parent.excrptGrps {
			fmt.Println(excrptGrp.ind, ":", excrptGrp.name)
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
	total := excrptGrpTotals[excrptGrp.name] + 1
	if excrptGrp.parent == "Indkomst efter skat" {
		return total, nil
	} else {
		return -1 * total, nil
	}
}

func InitExcrptGrps(excrptGrps *sheets.ValueRange, month int64) {
	// Open the JSON file
	file, err := os.Open("excrptGrpData.json")
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

	updateCommonGrps(excrptGrps, month)
}

func updateCommonGrps(excrptGrps *sheets.ValueRange, month int64) {
	// Get Date, Amount and description

	A1Not := util.MonthToA1Notation(month)
	for i, elm := range excrptGrps.Values {
		if len(elm) != 0 {
			excrptGrp, notFound := GetExcrptGrp(elm[0].(string), -1)
			if notFound == nil {
				if excrptGrp.isCommonGrp {

					readRangeExrpt := "budget!" + A1Not + fmt.Sprint(i+1)
					excrpts, readExcrptsErr := req.GetSheet().Values.Get(req.GetSpreadsheetId(), readRangeExrpt).Do()

					if readExcrptsErr != nil {
						log.Fatalf("Unable to perform get: %v", readExcrptsErr)
					}

					if len(excrpts.Values) == 0 {
						excrptGrpTotals[excrptGrp.name] += 0.0
					} else {

						val := strings.Trim(excrpts.Values[0][0].(string), " ")
						if len(val) > 0 {
							amount, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(val[:len(val)-4], ".", ""), ",", "."), 64)
							if err != nil {
								log.Fatal(err)
							}
							// updateExcrptTotal("9999-99-99", excrptGrp.name, amount)
							excrptGrpTotals[excrptGrp.name] += -1 * float64(amount)
						} else {
							excrptGrpTotals[excrptGrp.name] += 0.0
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
		ind:         ind,
		name:        name,
		mappings:    data.Matches,
		parent:      parent,
		isCommonGrp: data.IsCommonGrp,
	}
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
		for _, excrptGrp := range parent.excrptGrps {
			if strings.EqualFold(strings.Trim(excrptGrp.name, " "), strings.Trim(name, " ")) || excrptGrp.ind == ind {
				return excrptGrp, nil
			}
		}
	}
	return ExcrptGrp{}, errors.New("Excrpt grp: " + name + ", not found")
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
