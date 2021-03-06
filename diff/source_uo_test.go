package diff

import (
	"testing"
)

func TestParseArray(t *testing.T) {
	cells := [][]string{
		{"Lorem Ipsum"},
		{},
		{"Account Statement Details"},
		{},
		{"Account Number:", "1234567890123456", "SGD"},
		{"Account Type:", ""},
		{"Statement Date:", "01 Feb 2020"},
		{"Statement Balance:", "2.00", "SGD"},
		{},
		{
			"Transaction Date",
			"Posting Date",
			"Description",
			"Foreign Currency Type",
			"Transaction Amount(Foreign)",
			"Local Currency Type",
			"Transaction Amount(Local)",
		},
		{},
		{"01 Jan 2020", "01 Jan 2020", "both", "", "", "SGD", "1.00"},
		{"02 Jan 2020", "02 Jan 2020", "uo only", "", "", "SGD", "1.00"},
	}

	want := map[string][][]string{
		"2020-01-01 100": {
			{"01 Jan 2020", "1.00", "both"},
		},
		"2020-01-02 100": {
			{"02 Jan 2020", "1.00", "uo only"},
		},
	}

	got, err := parseArray(cells)

	if err != nil {
		t.Errorf("err = %v; want %v", err, nil)
	}

	if !groupsEqual(got, want) {
		t.Errorf("parseArray(cells) = %v; want %v", got, want)
	}
}

func groupsEqual(left, right map[string][][]string) bool {
	if len(left) != len(right) {
		return false
	}

	for leftKey, leftValue := range left {
		rightValue, ok := right[leftKey]

		if !ok {
			return false
		}

		if len(leftValue) != len(rightValue) {
			return false
		}

		for i := 0; i < len(leftValue); i++ {
			for j := 0; j < len(leftValue[i]); j++ {
				if leftValue[i][j] != rightValue[i][j] {
					return false
				}
			}
		}
	}

	return true
}
