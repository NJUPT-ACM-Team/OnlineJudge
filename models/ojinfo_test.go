package models

import (
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
	ojim := NewOJInfoModel()
	if err := ojim.OpenDB(); err != nil {
		t.Errorf("Failed to open db, %s", err)
	}
	defer ojim.CloseDB()
	oj, err := ojim.QueryByName("zoj", nil, nil)
	if err != nil {
		t.Errorf("Failed to query by 'noj', %s", err)
	}
	t.Log(oj)
}

func TestOJInfoQueryALl(t *testing.T) {
	ojim := NewOJInfoModel()
	if err := ojim.OpenDB(); err != nil {
		t.Errorf("Failed to open db, %s", err)
	}
	defer ojim.CloseDB()
	ojs, err := ojim.QueryAll(nil, nil)
	if err != nil {
		t.Errorf("Failed to query all, %s", err)
	}
	for _, oj := range ojs {
		t.Log(oj)
	}
}

func TestOJInfoUpdate(t *testing.T) {
	ojim := NewOJInfoModel()
	if err := ojim.OpenDB(); err != nil {
		t.Errorf("Failed to open db, %s", err)
	}
	defer ojim.CloseDB()
	ojinfo := OJInfo{
		OJId:    1,
		Int64IO: "%I64d",
	}
	if err := ojim.Update(&ojinfo, []string{"int64io"}, nil); err != nil {
		t.Errorf("Failed to update, %s", err)
	}
	if err := ojim.Commit(); err != nil {
		t.Errorf("Failed to commit, %s", err)
	}
}
