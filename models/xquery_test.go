package models

import (
	"OnlineJudge/db"
	"testing"
)

func TestXQuery_List_Problems_With_Filter(t *testing.T) {
	db.Init()
	DB := db.New()
	tx := DB.MustBegin()
	paging, err := XQuery_List_Problems_With_Filter(
		tx,
		"kevince",
		true,
		"zoj",
		3,
		0,
		false,
		0,

		10,
		2,
		nil,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(paging)
}

func TestXQuery_List_Submissions_With_Filter(t *testing.T) {
	db.Init()
	DB := db.New()
	tx := DB.MustBegin()

	paging, err := XQuery_List_Submissions_With_Filter(
		tx,
		"kevince",
		true,
		"kevince",
		"zoj",
		"1000",
		"",
		"",
		"",
		2,
		1,
		nil,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(paging)
}
