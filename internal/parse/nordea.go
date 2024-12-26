package parse

import (
	"OrdoBudget/internal/logtrace"
	"encoding/csv"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type Nordea struct{}

func (nor Nordea) Parse(path string, month int64) []Excrpt {
	return nor.parseCsv(path, month)
}

func (nor Nordea) parseCsv(filePath string, month int64) []Excrpt {
	monthTime := time.Month(month)

	file, err := os.Open(filePath)
	if err != nil {
		logtrace.Error(err.Error())
	}
	// Create a new CSV ReadExcrptCsv
	reader := csv.NewReader(file)

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
				logtrace.Error(err.Error())
			}

			cmpMth := time.Date(0, monthTime, 1, 0, 0, 0, 0, time.UTC)

			//
			date, err := time.Parse("2006/01/02", row[0])
			if err != nil {
				logtrace.Info(err.Error())

				// date can be "reserveret", and we only want to account for excrpts which has been taken from the account
			} else {

				cmpCurMth := time.Date(0, date.Month(), 1, 0, 0, 0, 0, time.UTC)

				if cmpMth.Equal(cmpCurMth) {

					amount, err := strconv.ParseFloat(strings.ReplaceAll(row[1], ",", "."), 64)
					if err != nil {
						logtrace.Error(err.Error())
					}

					var balance float64
					var description string

					balance, err = strconv.ParseFloat(strings.ReplaceAll(row[6], ",", "."), 64)
					if err != nil {
						logtrace.Error(err.Error())
					}
					description = row[5]
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
