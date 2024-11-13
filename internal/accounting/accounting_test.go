package accounting

import (
	"reflect"
	"testing"

	"google.golang.org/api/sheets/v4"
)

func TestCreatGrps(t *testing.T) {
	grp0 := "Grp 0   "
	entry0 := "Entry0"
	entry1 := "Entry1"
	entry2 := "Entry 2"
	entry3 := "Entry  3"
	entry5 := "ENtry 5"

	match1 := "xxx"
	match2 := "     , yy"
	match3 := ", x , y.2, ~, hteth"
	match4 := " "
	match5 := ""

	fixed2 := true

	grp1 := "Grp1"
	entry12 := "Entry1.2"
	entry13 := "Entry1.3"

	match13 := "Entry1.2,HEJ, Entry1.34,       ,false"

	fixed12 := true
	config := &sheets.GridData{
		RowData: []*sheets.RowData{
			createRowData(0, 0, 1, &grp0, nil, nil),
			createRowData(1, 1, 1, &entry0, &match1, nil),
			createRowData(1, 1, 1, &entry1, &match2, nil),
			createRowData(1, 1, 1, &entry2, &match3, &fixed2),
			createRowData(1, 1, 1, &entry3, &match4, nil),
			createRowData(1, 1, 1, &entry5, &match5, nil),

			createRowData(0, 0, 1, &grp1, nil, nil),
			createRowData(1, 1, 1, &entry12, nil, &fixed12),
			createRowData(1, 1, 1, &entry13, &match13, nil),
		},
	}
	// Final row

	createGrps(config)

	t.Run("# of groups", func(t *testing.T) {
		if len(Groups) != 2 {
			t.Fatalf("# of element in Groups are %d, should be 2", len(Groups))
		}
	})

	t.Run("# of entries in Group ", func(t *testing.T) {
		if len(Groups[1].Entries) != 2 {
			t.Fatalf("# of entries in Group[1].Entries, shoud be 2, but are %d", len(Groups[1].Entries))
		}
	})

	t.Run("Group name ", func(t *testing.T) {
		if Groups[0].Name != "GRP 0" {
			t.Fatalf("Name of Groups[0] does not match, should be GRP 0, but is %s", Groups[0].Name)
		}
	})
	// Test mappings

	t.Run("Entry mappings 1 ", func(t *testing.T) {
		if !reflect.DeepEqual(Groups[0].Entries[1].Mappings, []string{"yy"}) {
			t.Fatalf("Mappings of Groups[0].Entries[1] does not match []string{\"yy\"}, but was %v", Groups[0].Entries[1].Mappings)
		}
	})

	t.Run("Entry mappings 2 ", func(t *testing.T) {
		if !reflect.DeepEqual(Groups[0].Entries[2].Mappings, []string{"x", "y.2", "~", "hteth"}) {
			t.Fatalf("Mappings of Groups[0].Entries[2] does not match []string{\"x \", \"y.2\", \"~\",\"hteth\"}, but was %v", Groups[0].Entries[2].Mappings)
		}
	})

	t.Run("Entry mappings 3 ", func(t *testing.T) {
		if !reflect.DeepEqual(Groups[1].Entries[1].Mappings, []string{"entry1.2", "hej", "entry1.34", "false"}) {
			t.Fatalf("Mappings of Groups[1].Entries[1] does not match []string{\"entry1.2\", \"hej\", \"entry1.34\", \"false\"}, but was %v", Groups[1].Entries[1].Mappings)
		}
	})

	t.Run("Entry mappings 4 ", func(t *testing.T) {
		if len(Groups[1].Entries[0].Mappings) != 0 {
			t.Fatalf("length of mappings of Groups[1].Entries[0] is not 0, but was %v", len(Groups[1].Entries[0].Mappings))
		}
	})

	// Test fixed expense

	t.Run("Fixed expense 1", func(t *testing.T) {
		if Groups[0].Entries[3].FixedExpense != false {
			t.Fatal("Groups[0].Entries[3].FixedExpense is not false")
		}
	})

	t.Run("Fixed expense 2", func(t *testing.T) {
		if Groups[0].Entries[2].FixedExpense != true {
			t.Fatal("Groups[0].Entries[2].FixedExpense is not true")
		}
	})
}

func createRowData(red, green, blue float64, name *string, matches *string, fixedExpense *bool) *sheets.RowData {
	var cellDataFixed sheets.CellData

	if fixedExpense != nil {
		cellDataFixed = sheets.CellData{
			EffectiveFormat: &sheets.CellFormat{
				BackgroundColor: &sheets.Color{
					Red:   red,
					Blue:  blue,
					Green: green,
				},
			},
			UserEnteredValue: &sheets.ExtendedValue{
				BoolValue: fixedExpense,
			},
		}
	}

	var cellDataMatches sheets.CellData

	if matches != nil {
		cellDataMatches = sheets.CellData{
			EffectiveFormat: &sheets.CellFormat{
				BackgroundColor: &sheets.Color{
					Red:   red,
					Blue:  blue,
					Green: green,
				},
			},
			UserEnteredValue: &sheets.ExtendedValue{
				StringValue: matches,
			},
		}
	}
	return &sheets.RowData{
		Values: []*sheets.CellData{
			// name cell
			{
				EffectiveFormat: &sheets.CellFormat{
					BackgroundColor: &sheets.Color{
						Red:   red,
						Blue:  blue,
						Green: green,
					},
				},
				UserEnteredValue: &sheets.ExtendedValue{
					StringValue: name,
				},
			},
			// matches
			&cellDataMatches,
			// fixed expense
			&cellDataFixed,
		},
	}
}
