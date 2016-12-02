package base

import "testing"

var ptests = []struct {
	Str   string
	Valid bool
	Res   *Pid
}{
	{
		Str:   "LOCAL#1000",
		Valid: true,
		Res:   &Pid{OJName: "LOCAL", OJPid: 1000},
	},
	{
		Str:   "#1000",
		Valid: false,
	},
	{
		Str:   "1000",
		Valid: false,
	},
	{
		Str:   "LOCAL#",
		Valid: false,
	},
	{
		Str:   "#",
		Valid: false,
	},
}

func TestParseSid(t *testing.T) {
	for _, tt := range ptests {
		pid, err := ParseSid(tt.Str)
		if err != nil {
			if tt.Valid == true {
				t.Fatal("Supposed to be invalid, ", pid)
			}
		} else {
			if tt.Valid == false {
				t.Fatal("Supposed to be valid, ", tt)
			}
			if *pid != *tt.Res {
				t.Fatal("parse error, ", tt)
			}
		}
	}
}
