package parse

import (
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

type SparKron struct{}

func (spar SparKron) Parse(path string, month int64) []Excrpt {
	return spar.parseXlsx(path, month)
}

func (spar SparKron) parseXlsx(path string, month int64) []Excrpt {
	monthTime := time.Month(month)
	println("hj")
	f, err := excelize.OpenFile(path)
	println("hj")
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	var excrpts []Excrpt

	i := 0
	rows, err := f.GetRows("Table 1")
	if err != nil {
		fmt.Println(err)
	}
	for _, row := range rows {
		// Dont read header-line
		if i > 4 {
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				log.Fatal("Error:", err)
			}

			cmpMth := time.Date(0, monthTime, 1, 0, 0, 0, 0, time.UTC)

			//
			// date, err := time.Parse("2006/01/02", row[0])
			date, err := time.Parse("02/01/2006", row[0])
			if err != nil {
				log.Printf("Could not parse date. Skipping row %d in input file", i+1)

				// date can be "reserveret", and we only want to account for excrpts which has been taken from the account
			} else {

				cmpCurMth := time.Date(0, date.Month(), 1, 0, 0, 0, 0, time.UTC)

				if cmpMth.Equal(cmpCurMth) {

					tmp := strings.ReplaceAll(row[2][0:len(row[2])-4], ".", "")
					amount, err := strconv.ParseFloat(strings.ReplaceAll(tmp, ",", "."), 64)
					if err != nil {
						log.Fatal(err)
					}

					var balance float64
					var description string

					tmp2 := strings.ReplaceAll(row[3][0:len(row[3])-4], ".", "")
					balance, err = strconv.ParseFloat(strings.ReplaceAll(tmp2, ",", "."), 64)
					if err != nil {
						log.Fatal(err)
					}
					description = row[1]
					excrpts = append(excrpts, CreateExcrpt(amount, balance, date.Format("02/01/2006"), description))

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
