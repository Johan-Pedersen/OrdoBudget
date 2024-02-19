package main

import "fmt"

/*
Converts colum ID such as 'AD' to it's colum Index
*/
func colToColInd(col string) int64 {
	fmt.Println("colToColInd ", col)
	return int64(col[len(col)-1]-'A') + int64(25*max(0, len(col)-1))
}
