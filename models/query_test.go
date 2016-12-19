package models

import (
	"OnlineJudge/db"

	"testing"
)

func TestQuery_All_Languages(t *testing.T) {
	db.Init()
	DB := db.New()
	tx := DB.MustBegin()

	langs, err := Query_All_Languages(tx, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(langs)
}

func TestQuery_All_OJNames(t *testing.T) {
	db.Init()
	DB := db.New()
	tx := DB.MustBegin()

	ojs, err := Query_All_OJNames(tx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ojs)
}

func TestQuery_All_OJs(t *testing.T) {
	db.Init()
	DB := db.New()
	tx := DB.MustBegin()

	ojs, err := Query_All_OJs(tx, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ojs)
}
