package accounting

import (
	"reflect"
	"testing"

	"google.golang.org/api/sheets/v4"
)

func TestCreatGrps(t *testing.T) {
	valRange := &sheets.ValueRange{
		Values: [][]interface{}{
			{"Grp 0   "},
			{"Entry0", "FALSE", "xxx"},
			{"Entry1", "FALSE", "     , yy"},
			{"Entry 2", "true", ", x , y.2, ~, hteth"},
			{"Entry  3", "FALSE", " "},
			{"ENtry 5", "FALSE", ""},

			{"Grp1"},
			{"Entry1,2", "TRUE"},
			{"Entry1.3", "false", "Entry1.2,HEJ, Entry1.34,       ,false"},
		},
	}

	createGrps(valRange)

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
