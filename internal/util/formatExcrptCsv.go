package util

import (
	excrpt "budgetAutomation/internal/excrpt"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

/*
Denne er lidt overflødig
*/
func ReadExcrptCsv(r io.Reader, month int64) []excrpt.Excrpt {
	monthTime := time.Month(month)

	// Create a new CSV ReadExcrptCsv
	reader := csv.NewReader(r)

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

			elms := strings.Split(row[0][:len(row[0])-1], ",")

			cmpMth := time.Date(0, monthTime, 1, 0, 0, 0, 0, time.UTC)

			//
			date, err := time.Parse("2006/01/02", elms[0])
			if err != nil {
				log.Printf("Could not parse date. Skipping row %d", i)
			} else {

				cmpCurMth := time.Date(0, date.Month(), 1, 0, 0, 0, 0, time.UTC)

				if cmpMth.Equal(cmpCurMth) {

					// Google sheets calculate dates as - days since Dec 30, 1899
					baseDate := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
					days := float64(date.Sub(baseDate).Hours() / 24)
					amount, err := strconv.ParseFloat(elms[1]+"."+elms[2], 64)
					if err != nil {
						log.Fatal(err)
					}

					var balance float64
					var description string
					if len(elms) > 10 {

						balance, err = strconv.ParseFloat(elms[8]+"."+elms[9], 64)
						if err != nil {
							log.Fatal(err)
						}
						description = elms[6] + " " + elms[7]
					} else {

						balance, err = strconv.ParseFloat(elms[7]+"."+elms[8], 64)
						description = elms[6]
						if err != nil {
							log.Fatal(err)
						}
					}

					excrpts = append(excrpts, excrpt.CreateExcrpt(days, amount, balance, description))

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
