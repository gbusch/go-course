package main

import "testing"

func TestFindIan(t *testing.T) {
	tables := []struct {
		input string
		found bool
	}{
		{"ian", true},
		{"Ian", true},
		{"iuiygaygn", true},
		{"ihhhhn", false},
		{"ina", false},
		{"xian", false},
	}

	for _, table := range tables {
		found := FindIan(table.input)
		if found != table.found {
			t.Errorf("FindIan failed. Input: %v. Got %v, expected %v.", table.input, found, table.found)
		}
	}
}
