package util

import (
	excrpt "budgetAutomation/internal/excrpt"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
Denne er lidt overflødig
*/
func ReadExcrptCsv(path string, month int64) []excrpt.Excrpt {
	monthTime := time.Month(month)

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
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatal("Error:", err)
			}

			cmpMth := time.Date(0, monthTime, 1, 0, 0, 0, 0, time.UTC)
			//
			date, err := time.Parse("2006/01/02", row[0])
			if err != nil {
				log.Printf("Could not parse date. Skipping row %d", i)
			} else {

				cmpCurMth := time.Date(0, date.Month(), 1, 0, 0, 0, 0, time.UTC)

				if cmpMth.Equal(cmpCurMth) {

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

					// All following excrpts will be before our month of interest
				} else if cmpMth.Before(cmpCurMth) {
					break
				}
				//
			}
		}

		i++
	}
	return excrpts
}
