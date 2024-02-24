package util

import (
	"fmt"
	"log"
)

/*
Converts colum ID such as 'AD' to it's colum Index
*/
func ColToColInd(col string) int64 {
	fmt.Println("colToColInd ", col)
	return int64(col[len(col)-1]-'A') + int64(25*max(0, len(col)-1))
}

/*
Month not 0 indexed
Currently it matches only my col
*/
func MonthToColInd(month int64) int64 {
	return 1 + ((month - 1) * 2)
}

/*
Check if date of exerpt is in the prev month
*/
func CheckCurMonth(haveCurMonth bool, curMonth, exrptMonth int64) (bool, bool, int64) {
	isNewMonth := false

	if !haveCurMonth {
		curMonth = exrptMonth
		haveCurMonth = true
	}
	if exrptMonth != curMonth {

		isNewMonth = true
		log.Println("Begin new month")
	} else {
		curMonth = exrptMonth
	}
	return isNewMonth, haveCurMonth, curMonth
}
