package main

import (
	"budgetAutomation/ui/cli"
	"flag"

	excrptgrps "budgetAutomation/internal/excrptGrps"
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
	// Update excerpt sheet, before we begin
	cli.UpdateExcrptsSheet()
	// Debug mode
	if *debugMode {

		// hard code person + month
		person = 1
		month = 6

		// excrpts = cli.DebugGetExcrpts()

		// Initialize and print excerpt groups
		excrptgrps.InitExcrptGrpsDebug()

	} else {

		cli.GetPersonAndMonth(&person, &month)
		// excrpts = cli.GetExcrpts()
		// Initialize and print excerpt groups
		excrptgrps.InitExcrptGrps(sheetsGrpCol, month, person)
		cli.PrintExcrptGrps()
		cli.PrintExcrptGrps()

	}
	// accBalance := cli.LoadExcrptTotal(excrpts, month)
	accBalance := 0.0

	// find Excerpt Total for current month.
	cli.PrintExcrptGrpTotals()
	cli.UpdateBudget(sheetsGrpCol, accBalance, month, person)
	cli.PrintResume()
}
