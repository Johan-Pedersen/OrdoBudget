package accounting

import (
	"budgetAutomation/internal/excrpt"
	"budgetAutomation/internal/requests"
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

var Balances = map[string]float64{}

var Groups = []Group{}

var Resume = []string{}

func isIgnored(groupName string) bool {
	return groupName == "Ignored"
}

func FindExcrptMatches(excrpt string) []Entry {
	var matches []Entry

	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	// Find correct excrpt grp
	for _, parent := range Groups {
		for i := range parent.Entries {
			for _, match := range parent.Entries[i].Mappings {
				match = strings.ToLower(strings.Trim(match, " "))
				if strings.Contains(excrpt, match) {
					matches = append(matches, parent.Entries[i])
					break
				}
			}
		}
	}
	return matches
}

func UpdateBalance(date, excrpt string, amount float64, GroupName string) {
	tmpAmount := amount

	if isIgnored(GroupName) {
		tmpAmount = 0
	}

	Balances[GroupName] += float64(tmpAmount)
	UpdateResume(date, excrpt, GroupName, tmpAmount)
}

func UpdateResume(date, excrpt, GroupName string, amount float64) {
	Resume = append(Resume, date+" "+excrpt+" "+strconv.FormatFloat(amount, 'f', -1, 64)+": "+GroupName)
}

func GetTotal(EntryName string) (float64, error) {
	entry, err := GetEntry(EntryName, -1)
	//
	if err != nil {
		return 0, err
	}

	// Total should always be a positive number
	total := Balances[entry.Name] + 1
	if entry.GroupName == "Indkomst efter skat" {
		return total, nil
	} else {
		return -1 * total, nil
	}
}

func InitEntriesDebug() {
	// Load ExcrptGrps

	f1, err := os.Open("build/debug/JsonExcrptGrps")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptsGrps")
	}
	defer f1.Close()                    // //Json decode
	json.NewDecoder(f1).Decode(&Groups) // if err != nil {

	f2, err := os.Open("build/debug/JsonExcrptGrpTotals")
	if err != nil {
		log.Fatal("Unable to open JSonExcrptsGrpTotals")
	}
	defer f2.Close()                      // //Json decode
	json.NewDecoder(f2).Decode(&Balances) // if err != nil {
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

func createGrps(data Data) {
	i := 0
	// Init excrptGrp Totals
	for groupName, excrpts := range data.Mappings {

		// parentName = strings.Trim(parentName, " ")
		entries := []Entry{}
		for entryName, mappings := range excrpts {
			// excrptGrp = strings.Trim(excrptGrp, " ")
			Balances[entryName] = -1.0
			entries = append(entries,
				createEntry(i, entryName, groupName, mappings))
			i += 1
		}
		group := Group{groupName, entries}
		Groups = append(Groups, group)
	}
}

func createEntry(ind int, name, groupName string, data DataExcrpt) Entry {
	return Entry{
		Ind:          ind,
		Name:         name,
		Mappings:     data.Matches,
		GroupName:    groupName,
		FixedExpense: data.FixedExpense,
	}
}

func GetGroups() []Group {
	return Groups
}

/*
Get entry based on name OR index(ind).
Both can be specified, but it's not necessary.
If you don't want to use ind, make ind < 0.
if you don't want to use name, make name = "". \n
returnes err if not found
*/
func GetEntry(name string, ind int) (Entry, error) {
	for _, group := range Groups {
		for _, entry := range group.Entries {
			if strings.EqualFold(strings.Trim(entry.Name, " "), strings.Trim(name, " ")) || entry.Ind == ind {
				return entry, nil
			}
		}
	}
	return Entry{}, errors.New("Entry: " + name + ", not found")
}

/*
Get all entries belonging to given grpName
*/
func GetEntries(grpName string) []Entry {
	for _, grp := range Groups {
		if grp.Name == grpName {
			return grp.Entries
		}
	}
	return nil
}

/*
Find matches for excrpts and updates ExcrptTotal iff only 1 match is found, otherwise the found excrpts are added to the return
*/
func FindUpdMatches(excrpts *[]excrpt.Excrpt) map[excrpt.Excrpt][]Entry {
	ret := make(map[excrpt.Excrpt][]Entry)

	for _, excrpt := range *excrpts {
		matches := FindExcrptMatches(excrpt.Description)
		// If there is only a single match, the update is given
		// Otherwise the correct match has to be made in the ui
		if len(matches) == 1 {
			UpdateBalance(excrpt.Date, excrpt.Description, excrpt.Amount, matches[0].Name)
		} else {
			ret[excrpt] = matches
		}
	}
	return ret
}

func updateCommonGrps(sheetEntries *sheets.ValueRange, month, person int64) {
	// Get Date, Amount and description

	A1Not := util.MonthToA1Notation(month, person)
	for i, elm := range sheetEntries.Values {
		if len(elm) != 0 {
			entry, notFound := GetEntry(elm[0].(string), -1)
			if notFound == nil {
				if entry.FixedExpense {

					readRange := "budget!" + A1Not + fmt.Sprint(i+1)
					excrpts, readExcrptsErr := requests.GetSheet().Values.Get(requests.GetSpreadsheetId(), readRange).Do()

					if readExcrptsErr != nil {
						log.Fatalf("Unable to perform get: %v", readExcrptsErr)
					}

					if len(excrpts.Values) == 0 {
						Balances[entry.Name] += 0.0
					} else {

						val := strings.Trim(excrpts.Values[0][0].(string), " ")
						if len(val) > 0 {
							amount, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(val[:len(val)-4], ".", ""), ",", "."), 64)
							if err != nil {
								log.Fatal(err)
							}
							// updateExcrptTotal("9999-99-99", excrptGrp.name, amount)
							Balances[entry.Name] += -1 * float64(amount)
						} else {
							Balances[entry.Name] += 0.0
						}
					}
				}
			}
		}
	}
}
