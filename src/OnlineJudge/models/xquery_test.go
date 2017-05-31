package models

import (
	"OnlineJudge/base"
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
		0,
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

func TestXQuery_Contest_List_Submissions_With_Filter(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	ss, err := XQuery_Contest_List_Submissions_With_Filter(
		tx,
		"",    //username
		false, // show_private,
		8,
		false, // is_desc,
		0,     // run_id,
		"",    //username
		"",    //label
		"",    //status_code
		"",    //language
		"",    // compiler
		20,    //per_page
		1,     // current_page
		nil,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ss)
}

func TestXQuery_ContestRanklist_Submissions(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	var contest_id int64
	contest_id = 15
	cst, err := Query_Contest_By_ContestId(tx, contest_id, nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base.MarshalTime(cst.StartTime))
	t.Log(base.MarshalTime(cst.EndTime))
	subs, err := XQuery_ContestRanklist_Submissions(
		tx,
		contest_id,
		cst.StartTime,
		cst.EndTime,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(subs)
}
