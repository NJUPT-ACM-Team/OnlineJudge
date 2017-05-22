package models

/*
import (
	"OnlineJudge/db"
	"testing"
)
*/

/*
func TestContestProblemInsert(t *testing.T) {
	DB, err := db.NewDB()
	if err != nil {
		t.Errorf("Failed to open db, %s", err)
		return
	}
	tx, err := DB.Beginx()
	if err != nil {
		t.Errorf("Failed to start transaction, %s", err)
		return
	}
	defer DB.Close()
	cp := &ContestProblem{
		MetaPidFK:   2,
		ContestIdFK: 3,
		Alias:       "the second problem",
		Label:       "B",
	}
	cpm := NewContestProblemModel()
	id, err := cpm.Insert(tx, cp)
	if err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
*/

/*
func TestContestProblemDel(t *testing.T) {
	DB, err := db.NewDB()
	if err != nil {
		t.Fatal(err)
	}
	tx, err := DB.Beginx()
	if err != nil {
		t.Fatal(err)
	}
	defer DB.Close()

	cpm := NewContestProblemModel()
	if err := cpm.DeleteProblemsByContestId(tx, 3); err != nil {
		t.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		t.Fatal(err)
	}
}
*/

/*
func TestLanuageQuery(t *testing.T) {
	DB, err := db.NewDB()
	if err != nil {
		t.Errorf("Failed to open db, %s", err)
		return
	}
	tx, err := DB.Beginx()
	if err != nil {
		t.Errorf("Failed to start transaction, %s", err)
		return
	}
	defer DB.Close()
	lm := NewLanguageModel()
	lang, err := lm.QueryById(tx, 1, nil, nil)
	if err != nil {
		t.Errorf("Failed to query, %s", err)
	}
	t.Log(lang)
}
*/
