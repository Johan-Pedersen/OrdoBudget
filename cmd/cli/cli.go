package main

import (
	"budgetAutomation/internal/accounting"
	"budgetAutomation/internal/parser"
	"budgetAutomation/ui/cli"
	"flag"
	"log"
	"os"
)

func main() {
	debugMode := flag.Bool("debug", false, "Run in debug mode")

	flag.Parse()

	// Who is doing the budget
	var person int64

	// Which month from 1-12 should be handled
	var month int64

	// var excrpts *sheets.ValueRange
	sheetsGrpCol := cli.GetSheetsGrpCol()
	// Debug mode
	if *debugMode {

		// hard code person + month
		person = 1
		month = 6

		// Update excerpt sheet, before we begin
		cli.UpdateExcrptsSheet(month)
		// excrpts = cli.DebugGetExcrpts()

		// Initialize and print excerpt groups
		accounting.InitGrpsDebug()

	} else {

		cli.GetPersonAndMonth(&person, &month)
		// Update excerpt sheet, before we begin
		// cli.UpdateExcrptsSheet(month)
		// excrpts = cli.GetExcrpts()
		// Initialize and print excerpt groups
		accounting.InitGrps(sheetsGrpCol, month, person)
		cli.PrintExcrptGrps()

	}

	reader, err := os.Open("/home/hanyolo/src/budgetAutomation/storage/excrptSheet.csv")
	if err != nil {
		log.Fatalln("Could not open excrpt file", err)
	}
	// Read excrpts from csv
	excrpts := parser.ReadExcrptCsv(reader, month)

	// Auto find matches
	// create upd requests for match

	matches := accounting.FindUpdMatches(&excrpts)

	// Decide unknown matches

	cli.DecideEntries(matches)

	// Create upd requests for match

	// Update budget -> API kald

	cli.UpdateBudget(sheetsGrpCol, 0.0, month, person)
	// accBalance := cli.LoadExcrptTotal(excrpts, month)

	// find Excerpt Total for current month.
	cli.PrintExcrptGrpTotals()
	cli.PrintResume()
}
