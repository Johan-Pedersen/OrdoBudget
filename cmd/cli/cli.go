package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"OrdoBudget/internal/accounting"
	"OrdoBudget/internal/logtrace"
	"OrdoBudget/internal/parse"
	"OrdoBudget/ui/cli"
)

// Who is doing the budget
var person int64 = 0

// Which month from 1-12 should be handled
var month int64

var multipleUsers bool

var (
	bankStr          string
	multipleUsersStr string
	bank             parse.Bank
)

var parser parse.Parser

func main() {
	debugMode := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()

	// var excrpts *sheets.ValueRange
	sheetsGrpCol := cli.GetSheetsGrpCol()

	// Debug mode
	if *debugMode {

		// hard code person + month
		person = 1
		month = 6

		// Initialize and print excerpt groups
		accounting.InitGrpsDebug()

	} else {

		if multipleUsers {
			cli.InputPerson(&person)
		}

		cli.InputMonth(&month)

		// Update excerpt sheet, before we begin
		// cli.UpdateExcrptsSheet(month)
		// excrpts = cli.GetExcrpts()
		// Initialize and print excerpt groups
		accounting.InitGrps(sheetsGrpCol, month, person)
	}

	var inputFileName string
	fmt.Println("Enter input file name")
	fmt.Scan(&inputFileName)
	inputFileName = strings.TrimSpace(inputFileName)

	excrpts := parser.Parse(inputFileName, month)

	// Auto find matches
	// create upd requests for match

	matches := accounting.FindUpdMatches(&excrpts)

	// Decide unknown matches

	cli.DecideEntries(matches)

	// Create upd requests for match

	// Update budget -> API kald

	cli.UpdateBudget(sheetsGrpCol, excrpts[0].Balance, month, person)
	// accBalance := cli.LoadExcrptTotal(excrpts, month)

	// find Excerpt Total for current month.
	cli.PrintResume()
	cli.PrintBalances()
	fmt.Println("Press Enter to finish")
	fmt.Scanln()
}

// This code converts the -ldflags -X build flags from string to it's actual type

func init() {
	//
	var err error
	multipleUsers, err = strconv.ParseBool(multipleUsersStr)
	if err != nil {
		logtrace.Error(err.Error())
	}

	switch strings.TrimSpace(bankStr) {
	case parse.NordeaBank.String():
		parser = parse.Nordea{}
	case parse.SparKronBank.String():
		parser = parse.SparKron{}

	}
}
