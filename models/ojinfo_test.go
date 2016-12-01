package models

import (
	"OnlineJudge/models/db"
	"testing"
)

/*
import "time"

func TestOJInfoInsert(t *testing.T) {
	ojim := NewOJInfoModel()
	if err := ojim.OpenDB(); err != nil {
		t.Errorf("Failed to open db, %s", err)
	}
	defer ojim.CloseDB()
	ojinfo := OJInfo{
		Name:       "zoj",
		Version:    "1",
		Int64IO:    "%I64d",
		JavaClass:  "Main",
		Status:     "ok",
		StatusInfo: "OK",
		LastCheck:  time.Now(),
	}
	id, err := ojim.Insert(&ojinfo)
	t.Log("id: ", id)
	if err != nil {
		t.Errorf("TestOjInfo: %s", err)
	}
}
*/

func TestOJInfoQueryByName(t *testing.T) {
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
	ojim := NewOJInfoModel()
	oj, err := ojim.QueryByName(tx, "zoj", nil, nil)
	if err != nil {
		t.Errorf("Failed to query by 'noj', %s", err)
	}
	t.Log(oj)
}

func TestOJInfoQueryALl(t *testing.T) {
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
	ojim := NewOJInfoModel()
	ojs, err := ojim.QueryAll(tx, nil, nil)
	if err != nil {
		t.Errorf("Failed to query all, %s", err)
	}
	for _, oj := range ojs {
		t.Log(oj)
	}
}

func TestOJInfoUpdate(t *testing.T) {
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
	ojim := NewOJInfoModel()
	ojinfo := OJInfo{
		OJId:    1,
		Int64IO: "%I64d",
	}
	if err := ojim.Update(tx, &ojinfo, "", []string{"int64io"}, nil); err != nil {
		t.Errorf("Failed to update, %s", err)
	}
	if err := tx.Commit(); err != nil {
		t.Errorf("Failed to commit, %s", err)
	}
}
