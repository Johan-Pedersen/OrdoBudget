package excrptgrps

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

var excrptGrpTotals = map[string]float64{}

var excrptGrpMappings = []ExcrptGrp{}

// Marshal and unmarshal json
type Data struct {
	ExcrptMappings map[string][]string
}

type ExcrptGrp struct {
	// Used to make lookup in excerptMappings array
	ind int

	// Name of the ExcrptGrp
	name string

	// Matches for this excrptGrp
	mappings []string
}

func UpdateExcrptTotal(date, excrpt string, amount float64) {
	excrptGrpName := ""
	// ignore case
	excrpt = strings.ToLower(strings.Trim(excrpt, " "))

	ind := -1
	// Find correct excrpt grp
	for _, excrptGrp := range excrptGrpMappings {

		for _, match := range excrptGrp.mappings {
			if strings.Contains(excrpt, match) {
				excrptGrpName = excrptGrp.name
				break
			}
		}

		if excrptGrpName != "" {
			break
		}
	}

	if excrptGrpName == "" {
		fmt.Println("Can't match to group:", date, excrpt, ":", amount)
		fmt.Println("Select group")
		fmt.Scan(&ind)

		excrptGrpName = excrptGrpMappings[ind].name

	}

	// Update correct excrpt grp
	excrptGrpTotals[excrptGrpName] += float64(amount)
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
	for _, excrptGrp := range excrptGrpMappings {
		fmt.Println(excrptGrp.ind, ":", excrptGrp.name)
	}
	fmt.Println("###################################################")
}

func GetTotal(excrptGrp string) float64 {
	total := excrptGrpTotals[excrptGrp]
	if total != 0.0 {
		return -1 * (excrptGrpTotals[excrptGrp] + 1)
	}
	return 0.0
}

func InitExcrptGrps() {
	// Open the JSON file
	file, err := os.Open("excrptGrpData.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the JSON data from the file
	jsonData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	// Create an instance of the struct to hold the unmarshaled data
	var data Data

	// Unmarshal the JSON data into the struct
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println("Error unmarshaling JSON data:", err)
		return
	}

	i := 0
	// Init excrptGrp Totals
	for excrptGrp, mappings := range data.ExcrptMappings {
		excrptGrpTotals[excrptGrp] = -1.0
		excrptGrpMappings = append(excrptGrpMappings,
			createExcrptGrp(i, excrptGrp, mappings))
		i += 1
	}
}

func createExcrptGrp(ind int, name string, mappings []string) ExcrptGrp {
	return ExcrptGrp{
		ind:      ind,
		name:     name,
		mappings: mappings,
	}
}
