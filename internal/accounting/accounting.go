package accounting

import (
	"OrdoBudget/internal/config"
	"OrdoBudget/internal/parser"
	"OrdoBudget/internal/request"
	"OrdoBudget/internal/util"
	"encoding/json"
	"errors"
	"fmt"
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
	return strings.ToUpper(groupName) == "IGNORED"
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

func GetBalance(EntryName string) (float64, error) {
	entry, err := GetEntry(EntryName, -1)
	//
	if err != nil {
		return 0, err
	}

	// Balance should always be a positive number
	balance := Balances[entry.Name] + 1

	if entry.GroupName == "INDKOMST EFTER SKAT" {
		return balance, nil
	} else {
		return -1 * balance, nil
	}
}

func InitGrpsDebug() {
	// Load ExcrptGrps

	f1, err := os.Open("build/debug/JsonEntries")
	if err != nil {
		log.Fatal("Unable to open JsonEntries")
	}
	defer f1.Close()                    // //Json decode
	json.NewDecoder(f1).Decode(&Groups) // if err != nil {

	f2, err := os.Open("build/debug/JsonBalances")
	if err != nil {
		log.Fatal("Unable to open JSonBalances")
	}
	defer f2.Close()                      // //Json decode
	json.NewDecoder(f2).Decode(&Balances) // if err != nil {
}

func InitGrps(sheetsGrpCol *sheets.ValueRange, month, person int64) {
	// Open the JSON file
	config := config.GetConfig()

	// initialize excerpt grps with -1 as total
	createGrps(config)

	updateFixedExpenses(sheetsGrpCol, month, person)
}

func createGrps(config *sheets.ValueRange) {
	// Init excrptGrp Totals
	var grp Group
	for i, elm := range config.Values {
		if len(elm) == 1 {
			if grp.Name != "" {
				Groups = append(Groups, grp)
			}
			grp = Group{}
			grp.Name = strings.TrimSpace(strings.ToUpper(elm[0].(string)))

			// Needed to filter out blanklines
		} else if len(elm) > 1 {

			var matches []string

			// Are there any matches
			if len(elm) == 3 {
				str := strings.ToLower(elm[2].(string))
				tmp := strings.Split(str, ",")
				for _, v := range tmp {
					v = strings.TrimSpace(v)
					// Make sure accidental whitespace is ignored
					if len(v) != 0 {
						matches = append(matches, v)
					}
				}
			}

			fixedExpense, err := strconv.ParseBool(elm[1].(string))
			if err != nil {
				log.Fatal("Cannot parse fixed expense of", elm[0].(string))
			}

			entry := Entry{
				Ind:          i,
				Name:         elm[0].(string),
				Mappings:     matches,
				GroupName:    grp.Name,
				FixedExpense: fixedExpense,
			}
			grp.Entries = append(grp.Entries, entry)
			Balances[entry.Name] = -1.0
		}
	}
	// Make sure to append ignore group
	Groups = append(Groups, grp)
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

func FindMatches(excrpt string) []Entry {
	var matches []Entry

	// ignore case
	excrpt = strings.ToLower(strings.TrimSpace(excrpt))

	// Find correct excrpt grp
	for _, grp := range Groups {
		for i := range grp.Entries {
			for _, match := range grp.Entries[i].Mappings {
				match = strings.ToLower(strings.TrimSpace(match))
				if strings.Contains(excrpt, match) {
					matches = append(matches, grp.Entries[i])
					break
				}
			}
		}
	}
	return matches
}

/*
Find matches for excrpts and updates Ballance iff only 1 match is found, otherwise the found excrpts are added to the return
*/
func FindUpdMatches(excrpts *[]parser.Excrpt) map[parser.Excrpt][]Entry {
	ret := make(map[parser.Excrpt][]Entry)

	for _, excrpt := range *excrpts {
		matches := FindMatches(excrpt.Description)
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

func updateFixedExpenses(sheetEntries *sheets.ValueRange, month, person int64) {
	// Get Date, Amount and description

	A1Not := util.MonthToA1Notation(month, person)
	for i, elm := range sheetEntries.Values {
		if len(elm) != 0 {
			entry, notFound := GetEntry(elm[0].(string), -1)
			if notFound == nil {
				if entry.FixedExpense {

					readRange := "budget!" + A1Not + fmt.Sprint(i+1)
					excrpts, readExcrptsErr := request.GetSheet().Values.Get(request.SpreadSheetId, readRange).Do()

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
