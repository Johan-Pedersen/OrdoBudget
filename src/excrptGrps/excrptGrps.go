package excrptgrps

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

type Data struct {
	ExcrptGrps     []string
	ExcrptMappings map[string][]string
}

var excrptGrpTotals = map[string]float64{}

var excrptGrpMappings = map[string][]string{}

func UpdateExcrptTotal(excrpt string, amount float64) {
	excrptGrp := ""
	// ignore case
	excrpt = strings.ToLower(excrpt)
	// Find correct excrpt grp
	for curExcrptGrp, excrpts := range excrptGrpMappings {

		for _, excr := range excrpts {
			if strings.Contains(excrpt, excr) {
				excrptGrp = curExcrptGrp
				break
			}
		}

		if excrptGrp != "" {
			break
		}
	}

	// Update correct excrpt grp
	excrptGrpTotals[excrptGrp] += float64(amount)
}

func PrintExcrptGrpTotals() {
	fmt.Println("###################################################")
	for k, v := range excrptGrpTotals {
		fmt.Println(k, ": ", v+1)
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

	// Init excrptGrp Totals
	for _, v := range data.ExcrptGrps {
		excrptGrpTotals[v] = -1.0
	}

	// Init ExcrptGrp mappings
	excrptGrpMappings = data.ExcrptMappings
}
