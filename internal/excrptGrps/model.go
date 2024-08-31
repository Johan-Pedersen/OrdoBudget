package excrptgrps

type DataExcrpt struct {
	Matches      []string
	FixedExpense bool
}

// Marshal and unmarshal json
type Data struct {
	ExcrptMappings map[string]map[string]DataExcrpt
}

type ExcrptGrp struct {
	// Used to make lookup in excerptMappings array
	Ind int

	// Name of the ExcrptGrp
	Name string

	// Matches for this excrptGrp
	Mappings []string

	// Defines the type of this excerpt
	Parent string

	// Determines if the initial group total value should be read from the sheet or start at 0
	// Default is false
	FixedExpense bool
}

type ExcrptGrpParent struct {
	Name       string
	ExcrptGrps []ExcrptGrp
}
