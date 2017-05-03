package core

import (
	"testing"
)

func TestMode(t *testing.T) {
	mode := &Mode{}
	mode.SetSPJ()
	if mode.IsSPJ() == false {
		t.Fatal("set spj failed")
	}
	mode.UnsetSPJ()
	if mode.IsSPJ() == true {
		t.Fatal("unset spj failed")
	}

	mode.SetOI()
	if mode.IsOI() == false {
		t.Fatal("set oi failed")
	}
	mode.UnsetOI()
	if mode.IsOI() == true {
		t.Fatal("unset oi failed")
	}

	mode.SetICPC()
	if mode.IsICPC() == false {
		t.Fatal("set icpc failed")
	}
	mode.UnsetICPC()
	if mode.IsICPC() == true {
		t.Fatal("unset icpc failed")
	}
}
