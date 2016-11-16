package base

import "testing"

func TestArrayContains(t *testing.T) {
	arrint := []int{1, 2}
	if ArrayContains(arrint[:], 1) == false {
		t.Errorf("1 is in arrint, expected true")
	}
	if ArrayContains(arrint[:], 2) == false {
		t.Errorf("2 is in arrint, expected true")
	}
	if ArrayContains(arrint[:], 3) == true {
		t.Errorf("3 not in arrint, expected false")
	}

	arrstr := []string{"a", "b"}
	if ArrayContains(arrstr[:], "a") == false {
		t.Errorf("1 is in arrstr, expected true")
	}
	if ArrayContains(arrstr[:], "b") == false {
		t.Errorf("b is in arrstr, expected true")
	}
	if ArrayContains(arrstr[:], "c") == true {
		t.Errorf("c not in arrstr, expected false")
	}
}
