package accounting

type Entry struct {
	// Used to make lookup in excerptMappings array
	Ind int

	// Name of Entry
	Name string

	// Matches for this Entry
	Mappings []string

	// Defines the type of this Entry
	GroupName string

	// Determines if the initial group total value should be read from the sheet or start at 0
	// Default is false
	FixedExpense bool
}

type Group struct {
	Name    string
	Entries []Entry
}
