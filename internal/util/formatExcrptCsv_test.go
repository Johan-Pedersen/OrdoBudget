package util

import (
	excrpt "budgetAutomation/internal/excrpt"
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNotLatestMth(t *testing.T) {
	reader, err := os.Open("/home/hanyolo/src/budgetAutomation/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 5)

	e1 := excrpt.CreateExcrpt(convertDate("2024/05/12"), -24, 8129.67, "test5")
	e2 := excrpt.CreateExcrpt(convertDate("2024/05/07"), -414, 8153.67, "test6 test6.2 oeu")

	expectedExcrpts := [2]excrpt.Excrpt{e1, e2}

	for i := 0; i < len(expectedExcrpts); i++ {
		if !excrpts[i].Equals(expectedExcrpts[i]) {
			fmt.Printf("excrpts[%d]: %v\n", i, excrpts[i])
			fmt.Printf("expectedExcrpts [%d]: %v\n", i, expectedExcrpts[i])
			t.Error("excrpts not as expected")
		}
	}
}

func TestNormCase(t *testing.T) {
	reader, err := os.Open("/home/hanyolo/src/budgetAutomation/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 6)

	e1 := excrpt.CreateExcrpt(convertDate("2024/06/12"), -24, 2988.78, "test1")
	e2 := excrpt.CreateExcrpt(convertDate("2024/06/07"), -414, 3012.78, "test2 test2.2 oeu")
	e3 := excrpt.CreateExcrpt(convertDate("2024/06/06"), -1399.89, 3426.78, "test3")
	e4 := excrpt.CreateExcrpt(convertDate("2024/06/06"), 207, 8336.67, "Test 4")

	expectedExcrpts := [4]excrpt.Excrpt{e1, e2, e3, e4}

	for i := 0; i < len(expectedExcrpts); i++ {
		if !excrpts[i].Equals(expectedExcrpts[i]) {
			fmt.Printf("excrpts[%d]: %v\n", i, excrpts[i])
			fmt.Printf("expectedExcrpts [%d]: %v\n", i, expectedExcrpts[i])
			t.Error("excrpts not as expected")
		}
	}
}

func TestEmpty(t *testing.T) {
	reader := strings.NewReader("")

	excrpts := ReadExcrptCsv(reader, 6)

	if len(excrpts) != 0 {
		t.Error("num of excrpts given empty input should be 0")
	}
}

func TestNotCorrectMth(t *testing.T) {
	reader, err := os.Open("/home/hanyolo/src/budgetAutomation/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 00)

	if len(excrpts) != 0 {
		t.Error("Number of excrpts should be 0")
	}
}

// Google sheets calculate dates as - days since Dec 30, 1899
func convertDate(dateStr string) float64 {
	date, err := time.Parse("2006/01/02", dateStr)
	if err != nil {
		log.Fatal("Could not parse date")
	}
	baseDate := time.Date(1899, time.December, 30, 0, 0, 0, 0, time.UTC)
	days := float64(date.Sub(baseDate).Hours() / 24)

	return days
}
