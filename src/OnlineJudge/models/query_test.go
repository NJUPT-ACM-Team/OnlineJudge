package models

import (
	"OnlineJudge/db"

	"testing"
)

func TestQuery_All_Languages(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	langs, err := Query_All_Languages(tx, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(langs)
}

func TestQuery_All_OJNames(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	ojs, err := Query_All_OJNames(tx)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ojs)
}

func TestQuery_All_OJs(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	ojs, err := Query_All_OJs(tx, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ojs)
}

func TestQuery_Limits_By_MetaPid(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	tms, err := Query_Limits_By_MetaPid(tx, 1, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tms)
}

func TestQuery_Contest_By_ContestId(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	cst, err := Query_Contest_By_ContestId(tx, 3, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cst)
}

func TestQuery_ContestProblemLabels_By_ContestId(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	labels, err := Query_ContestProblemLabels_By_ContestId(tx, 15)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(labels)
}

func TestQuery_ContestUsers_By_ContestId(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	users, err := Query_ContestUsers_By_ContestId(tx, 13)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}
