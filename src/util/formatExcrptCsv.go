package util

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadExcrptCsv(path string) ([]string, []float64, []string, []float64) {
	// Open the CSV file
	file, err := os.Open("Lønkonto.csv")
	if err != nil {
		log.Fatalln("Coud not open excerpt file.", err)
	}
	defer file.Close()

	// Create a new CSV ReadExcrptCsv
	reader := csv.NewReader(file)

	// Dont test on number of fields
	reader.FieldsPerRecord = -1

	// Set delimite
	reader.Comma = ';'
	// Read all records from the CSV file
	// records, err := reader.ReadAll()

	return getExcrpt(reader)
}

func getExcrpt(reader *csv.Reader) ([]string, []float64, []string, []float64) {
	// Print each record
	var dates []string
	var amounts []float64
	var descriptions []string
	var balances []float64

	i := 0
	for {
		row, err := reader.Read()
		// Dont read header-line
		if i > 0 {
			if err != nil {
				// Check for end of file
				if err.Error() == "EOF" {
					break
				} else {
					log.Fatal("Error:", err)
				}
			}
			dates = append(dates, row[0])
			amount, err := strconv.ParseFloat(strings.ReplaceAll(row[1], ",", "."), 64)
			if err != nil {
				log.Fatal(err)
			}
			amounts = append(amounts, amount)

			descriptions = append(descriptions, row[5])
			balance, err := strconv.ParseFloat(strings.ReplaceAll(row[6], ",", "."), 64)
			if err != nil {
				log.Fatal(err)
			}
			balances = append(balances, balance)
			//
		}

		i++
	}
	return dates, amounts, descriptions, balances
}
