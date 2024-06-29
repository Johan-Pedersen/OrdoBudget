package util

import (
	"fmt"
	"log"
	"math"
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
