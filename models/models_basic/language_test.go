package models_basic

import (
	"OnlineJudge/models/db"
	"testing"
)

/*
func TestLanguageInsert(t *testing.T) {
	lm := NewLanguageModel()
	lang := Language{
		Language:    "c++",
		OptionValue: "1",
		Compiler:    "g++4.9",
		OJIdFK:      1,
	}
	id, err := lm.Insert(&lang)
	if err != nil {
		t.Errorf("Failed to insert lang, %s", err)
	}
	t.Log("last insert id: ", id)

}
*/

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
	lang, err := lm.QueryById(tx, 2, nil, nil)
	if err != nil {
		t.Errorf("Failed to query, %s", err)
	}
	t.Log(lang)
}
