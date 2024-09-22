package util

import (
	"fmt"
	"log"
	"math"
	"time"
)

/*
Converts colum ID such as 'AD' to it's colum Index
*/
func ColToColInd(col string) int64 {
	fmt.Println("colToColInd ", col)
	return int64(col[len(col)-1]-'A') + int64(25*math.Max(0, float64(len(col)-1)))
}

/*
Month not 0-indexed -> Jan = 1 ... Dec = 12
Currently it matches only my col
*/
func MonthToColInd(month, person int64) int64 {
	var adjustment int64 = 0

	if person == 2 {
		adjustment += 1
	}
	return 1 + ((month - 1) * 2) + adjustment
}

/*
Month not 0-indexed -> Jan = 1 ... Dec = 12
Currently it matches only my col
*/
func MonthToA1Notation(month, person int64) string {
	return string(rune(97 + MonthToColInd(month, person)))
}

/*
Check if date of exerpt is in the prev month
*/
func CheckCurMonth(curMonth, exrptMonth int64) (bool, int64) {
	isNewMonth := false

	if curMonth == -1 {
		curMonth = exrptMonth
	} else if exrptMonth != curMonth {

		isNewMonth = true
		log.Println("Begin new month")
	}

	return isNewMonth, curMonth
}

/*
Converts from str date formate (2006/01/02) to days since Dec 30, 1899

Google sheets calculate dates as - days since Dec 30, 1899
*/
func ConvertDateToFloat(dateStr string) float64 {
	date, err := time.Parse("2006/01/02", dateStr)
	if err != nil {
		log.Fatal("Could not parse date")
	}
	baseDate := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	days := float64(date.Sub(baseDate).Hours() / 24)

	return days
}
