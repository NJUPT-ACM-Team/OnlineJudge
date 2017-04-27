package base

import "testing"

var pbtests = []struct {
	PB    string
	Field string
}{
	{
		PB:    "a_b_c",
		Field: "ABC",
	},
	{
		PB:    "abc",
		Field: "Abc",
	},
	{
		PB:    "abc_d",
		Field: "AbcD",
	},
	{
		PB:    "abcd_de_fg",
		Field: "AbcdDeFg",
	},
}

func TestPBTagToFieldName(t *testing.T) {
	for _, tt := range pbtests {
		res := PBTagToFieldName(tt.PB)
		if res != tt.Field {
			t.Fatal("wrong at", tt.PB, res)
		}
	}
}
