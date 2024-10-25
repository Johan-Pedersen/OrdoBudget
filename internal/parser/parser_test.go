package parser

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"
)

func TestNotLatestMth(t *testing.T) {
	reader, err := os.Open("/home/hanyolo/src/OrdoBudget/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 5)

	e1 := CreateExcrpt(-24, 8129.67, "2024/05/12", "test5")
	e2 := CreateExcrpt(-414, 8153.67, "2024/05/07", "test6 test6.2 oeu")

	expectedExcrpts := [2]Excrpt{e1, e2}

	for i := 0; i < len(expectedExcrpts); i++ {
		if !excrpts[i].Equals(expectedExcrpts[i]) {
			fmt.Printf("excrpts[%d]: %v\n", i, excrpts[i])
			fmt.Printf("expectedExcrpts [%d]: %v\n", i, expectedExcrpts[i])
			t.Error("excrpts not as expected")
		}
	}
}

func TestNormCase(t *testing.T) {
	reader, err := os.Open("/home/hanyolo/src/OrdoBudget/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 6)

	e1 := CreateExcrpt(-24, 2988.78, "2024/06/12", "test1")
	e2 := CreateExcrpt(-414, 3012.78, "2024/06/07", "test2 test2.2 oeu")
	e3 := CreateExcrpt(-1399.89, 3426.78, "2024/06/06", "test3")
	e4 := CreateExcrpt(207, 8336.67, "2024/06/06", "Test 4")

	expectedExcrpts := [4]Excrpt{e1, e2, e3, e4}

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
	reader, err := os.Open("/home/hanyolo/src/OrdoBudget/test/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt test file", err)
	}
	excrpts := ReadExcrptCsv(reader, 00)

	if len(excrpts) != 0 {
		t.Error("Number of excrpts should be 0")
	}
}
