package parser

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
	"strings"
	"time"
)

func ReadExcrptCsv(r io.Reader, month int64) []Excrpt {
	monthTime := time.Month(month)

	// Create a new CSV ReadExcrptCsv
	reader := csv.NewReader(r)

	// Dont test on number of fields
	reader.FieldsPerRecord = -1

	// Set delimite
	reader.Comma = ';'
	// Read all records from the CSV file

	var excrpts []Excrpt

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
				log.Printf("Could not parse date. Skipping row %d in input file", i+1)

				// date can be "reserveret", and we only want to account for excrpts which has been taken from the account
			} else {

				cmpCurMth := time.Date(0, date.Month(), 1, 0, 0, 0, 0, time.UTC)

				if cmpMth.Equal(cmpCurMth) {

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

					excrpts = append(excrpts, CreateExcrpt(amount, balance, date.Format("2006/01/02"), description))

					// All following excrpts will be before our month of interest
				} else if cmpCurMth.Before(cmpMth) {
					break
				}
				//
			}
		}

		i++
	}
	return excrpts
}
