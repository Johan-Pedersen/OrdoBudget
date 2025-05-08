package util

import (
	"testing"
)

func TestMonthToColInd(t *testing.T) {
	res := MonthToColInd(12, 1)

	if res != 23 {
		t.Error(res, "!= 23")
	}
}
