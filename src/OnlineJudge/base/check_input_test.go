package base

import (
	"testing"
)

var ntests = []struct {
	Username string
	Valid    bool
}{
	{Username: "afsdfs", Valid: true},
	{Username: "_afsdfs", Valid: false},
	{Username: "0dfas", Valid: false},
	{Username: "abd_2345!", Valid: false},
	{Username: "abc_fd00", Valid: true},
	{Username: "abc_fd0a", Valid: true},
}

func TestCheckUsername(t *testing.T) {
	for _, tt := range ntests {
		if CheckUsername(tt.Username) != tt.Valid {
			t.Fatal(tt.Username, " supposed to be ", tt.Valid)
		}
	}
}
