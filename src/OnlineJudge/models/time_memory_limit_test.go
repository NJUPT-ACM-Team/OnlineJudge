package models

/*
import (
	"OnlineJudge/db"

	"testing"
)

func TestInsertTimeMemoryLimit(t *testing.T) {
	DB := db.New()
	tx := DB.MustBegin()
	tcm := NewTimeMemoryLimitModel()
	tm := &TimeMemoryLimit{
		TimeLimit:     2000,
		MemoryLimit:   65536,
		CaseTimeLimit: 2000,
		Language:      "c++",
		MetaPidFK:     2,
	}
	id, err := tcm.Insert(tx, tm)
	if err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
*/
