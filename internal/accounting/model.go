package accounting

type DataExcrpt struct {
	Matches      []string
	FixedExpense bool
}

// Marshal and unmarshal json
type Data struct {
	Mappings map[string]map[string]DataExcrpt
}

type Entry struct {
	// Used to make lookup in excerptMappings array
	Ind int

	// Name of the ExcrptGrp
	Name string

	// Matches for this excrptGrp
	Mappings []string

	// Defines the type of this excerpt
	GroupName string

	// Determines if the initial group total value should be read from the sheet or start at 0
	// Default is false
	FixedExpense bool
}

type Group struct {
	Name    string
	Entries []Entry
}
