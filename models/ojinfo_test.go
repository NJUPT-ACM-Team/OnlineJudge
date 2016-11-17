package models

import (
	"testing"
)

/*
import "time"

func TestOJInfoInsert(t *testing.T) {
	ojim := NewOJInfoModel()
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
	oj, err := ojim.QueryByName("zoj")
	if err != nil {
		t.Errorf("Failed to query by 'noj', %s", err)
	}
	t.Log(oj)
}

func TestOJInfoQueryALl(t *testing.T) {
	ojim := NewOJInfoModel()
	ojs, err := ojim.QueryAll()
	if err != nil {
		t.Errorf("Failed to query all, %s", err)
	}
	for _, oj := range ojs {
		t.Log(oj)
	}
}
