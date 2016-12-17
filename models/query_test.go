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
		"zoj",
		0,
		0,
		false,
		0,

		1,
		2,
		nil,
		nil,
	)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(paging)
}