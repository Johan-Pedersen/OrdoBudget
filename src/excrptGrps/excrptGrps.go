package excrptgrps

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Data struct {
	ExcrptGrps     []string
	ExcrptMappings map[string]string
}

var excerptGrpSums = map[string]float64{}

var descToExcerptGrps = map[string]string{}

func UpdateExcerptSum(excerptGrp string, amount float64) {
	excerptGrpSums[descToExcerptGrps[excerptGrp]] += float64(amount)
}

func PrintExcerptGrpSum() {
	fmt.Println("###################################################")
	for k, v := range excerptGrpSums {
		fmt.Println(k, ": ", v+1)
	}
	fmt.Println("###################################################")
}

func GetTotal(excrptGrp string) float64 {
	total := excerptGrpSums[excrptGrp]
	if total != 0.0 {
		return excerptGrpSums[excrptGrp] + 1
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

	// Init excrptGrp Sums
	for _, v := range data.ExcrptGrps {
		excerptGrpSums[v] = -1.0
	}

	descToExcerptGrps = data.ExcrptMappings
}
