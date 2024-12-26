package accounting

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"OrdoBudget/internal/config"
	"OrdoBudget/internal/logtrace"
	"OrdoBudget/internal/parse"
	"OrdoBudget/internal/request"
	"OrdoBudget/internal/util"

	"google.golang.org/api/sheets/v4"
)

var Balances = map[string][]float64{}

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

	Balances[GroupName] = append(Balances[GroupName], float64(tmpAmount))
	UpdateResume(date, excrpt, GroupName, tmpAmount)
}

func UpdateResume(date, excrpt, GroupName string, amount float64) {

	Resume = append(Resume, date+" "+excrpt+" "+strconv.FormatFloat(amount, 'f', -1, 64)+": "+GroupName)
}

func GetAmounts(EntryName string) ([]float64, error) {
	entry, err := GetEntry(EntryName, -1)
	//
	if err != nil {
		return []float64{}, err
	}

	// Balance should always be a positive number
	amounts := Balances[entry.Name]

	if entry.GroupName != "INDKOMST EFTER SKAT" {
		for i := range amounts {
			amounts[i] = amounts[i] * -1
		}
	}
	return amounts, nil
}
func GetBalance(EntryName string) (float64, error) {
	entry, err := GetEntry(EntryName, -1)
	//
	if err != nil {
		return 0, err
	}

	// Balance should always be a positive number
	balance := sum(entry.Name)

	if entry.GroupName == "INDKOMST EFTER SKAT" {
		return balance, nil
	} else {
		return -1 * balance, nil
	}
}

func sum(entryName string) float64 {
	sum := 0.0
	for _, v := range Balances[entryName] {

		sum += v

	}
	return sum
}

func InitGrpsDebug() {
	// Load ExcrptGrps

	f1, err := os.Open("build/debug/JsonEntries")
	if err != nil {
		logtrace.Error(err.Error())
	}
	defer f1.Close()                    // //Json decode
	json.NewDecoder(f1).Decode(&Groups) // if err != nil {

	f2, err := os.Open("build/debug/JsonBalances")
	if err != nil {
		logtrace.Error(err.Error())
	}
	defer f2.Close()                      // //Json decode
	json.NewDecoder(f2).Decode(&Balances) // if err != nil {
}

func InitGrps(sheetsGrpCol *sheets.ValueRange, month, person int64) {
	// Open the JSON file
	config := config.GetConfig()

	// initialize excerpt grps with -1 as total
	createGrps(config[0])

	updateFixedExpenses(sheetsGrpCol, month, person)
}

func createGrps(config *sheets.GridData) {
	// Init excrptGrp Totals

	for i, row := range config.RowData {
		cellData := row.Values
		if isBlank(*cellData[0]) {
			break
		}
		if hasBlueBG(cellData[0].EffectiveFormat.BackgroundColor) {
			grp := Group{
				Name: strings.TrimSpace(strings.ToUpper(*cellData[0].UserEnteredValue.StringValue)),
			}

			Groups = append(Groups, grp)
		} else {

			grp := &Groups[len(Groups)-1]

			entry, err := entryFromCell(cellData, i, grp.Name)

			if err == nil {
				grp.Entries = append(grp.Entries, entry)
				Balances[entry.Name] = []float64{}
			}
		}
	}
}

func hasBlueBG(bg *sheets.Color) bool {
	return bg.Alpha == 0 && bg.Red == 0 && bg.Green == 0
}

func isBlank(cellData sheets.CellData) bool {
	// Cell has white background and no value
	return cellData.UserEnteredValue == nil &&
		cellData.EffectiveFormat.BackgroundColor.Red == 1 &&
		cellData.EffectiveFormat.BackgroundColor.Green == 1 &&
		cellData.EffectiveFormat.BackgroundColor.Blue == 1
}

func entryFromCell(cellData []*sheets.CellData, ind int, grpName string) (Entry, error) {
	var fixedExpense bool
	var entryName string
	var matches []string
	// Kolonne 0 er gruppen / posten

	// No entry name
	if cellData[0].UserEnteredValue == nil {
		return Entry{}, errors.New("no entry defined in row")
	}

	entryName = strings.TrimSpace(strings.ToLower(*cellData[0].UserEnteredValue.StringValue))
	// Kolonne 1 match

	if len(cellData) > 1 && cellData[1].UserEnteredValue != nil {
		tmpMatch := cellData[1].UserEnteredValue.StringValue

		tmp := strings.Split(*tmpMatch, ",")
		for _, v := range tmp {
			v = strings.TrimSpace(v)
			// Make sure accidental whitespace is ignored
			if len(v) != 0 {
				matches = append(matches, strings.ToLower(v))
			}
		}
	}

	// Kolonne 2 fast udgift
	if len(cellData) > 2 && cellData[2].UserEnteredValue != nil {
		fixedExpense = *cellData[2].UserEnteredValue.BoolValue
	}

	return Entry{
		Ind:          ind,
		Name:         entryName,
		Mappings:     matches,
		GroupName:    grpName,
		FixedExpense: fixedExpense,
	}, nil
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
func FindUpdMatches(excrpts *[]parse.Excrpt) map[parse.Excrpt][]Entry {
	ret := make(map[parse.Excrpt][]Entry)

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

	A1Not := util.MonthToA1Notation(month, person)
	for i, elm := range sheetEntries.Values {
		if len(elm) != 0 {
			entry, notFoundtErr := GetEntry(elm[0].(string), -1)
			if notFoundtErr == nil {
				if entry.FixedExpense {

					readRange := A1Not + fmt.Sprint(i+1)

					expense, readExpenseErr := request.GetSheet().Values.Get(request.SpreadSheetId, readRange).Do()

					if readExpenseErr != nil {
						logtrace.Error(readExpenseErr.Error())
					}

					// check size before we try to access element
					if len(expense.Values) > 0 {

						val := strings.Trim(expense.Values[0][0].(string), " ")
						fmt.Printf("val: %v\n", val)
						if len(val) > 0 {
							amount, err := strconv.ParseFloat(strings.ReplaceAll(strings.ReplaceAll(val[:len(val)-4], ".", ""), ",", "."), 64)
							if err != nil {
								logtrace.Error(err.Error())
							}
							Balances[entry.Name] = append(Balances[entry.Name], -1*float64(amount))
						}
					}
				}
			}
		}
	}
}
