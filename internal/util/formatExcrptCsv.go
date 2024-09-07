package util

import (
	excrpt "budgetAutomation/internal/excrpt"
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Denne er lidt overflødig
*/
func ReadExcrptCsv(path string) []excrpt.Excrpt {
	// Open the CSV file
	file, err := os.Open(path)
	if err != nil {
		print("Could not open excerpt file")
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

	var excrpts []excrpt.Excrpt

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
			date, err := time.Parse("2006/01/02", row[0])
			if err != nil {
				log.Println("Skipping this excerpt")
			} else {
				// Google sheets calculate dates as - days since Dec 30, 1899
				baseDate := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
				days := float64(date.Sub(baseDate).Hours() / 24)
				amount, err := strconv.ParseFloat(strings.ReplaceAll(row[1], ",", "."), 64)
				if err != nil {
					log.Fatal(err)
				}

				balance, err := strconv.ParseFloat(strings.ReplaceAll(row[6], ",", "."), 64)
				if err != nil {
					log.Fatal(err)
				}

				excrpts = append(excrpts, excrpt.CreateExcrpt(days, amount, balance, row[5]))
				//
			}
		}

		i++
	}
	return excrpts
}
