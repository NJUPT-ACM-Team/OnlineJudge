package models

import (
	"OnlineJudge/db"
	"testing"
)

func TestXQuery_List_Problems_With_Filter(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	paging, err := XQuery_List_Problems_With_Filter(
		tx,
		"kevince",
		true,
		"zoj",
		"ACCEPTED",
		"TITLE",
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
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()

	paging, err := XQuery_List_Submissions_With_Filter(
		tx,
		"kevince",
		true,
		false,
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

func TestXQuery_List_Contests_With_Filter(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	paging, err := XQuery_List_Contests_With_Filter(
		tx,
		10,
		1,
		"",   // order_by_element
		true, // is_desc
		"",   //filter_ctype_element
		"",   // is_public
		"",   // is_virtual
		nil,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(paging)
}

func TestXQuery_Contest_List_Problems(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	ps, err := XQuery_Contest_List_Problems(
		tx, 5, 1,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ps)
}
