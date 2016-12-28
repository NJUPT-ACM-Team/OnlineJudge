package models

import (
	"OnlineJudge/db"

	"testing"
)

import "time"

func TestInsertContest(t *testing.T) {
	db.InitTest()
	DB := db.New()
	tx := DB.MustBegin()
	cm := NewContestModel()
	con := &Contest{
		Title:            "第一次比赛",
		Description:      "for testing purpose",
		IsVirtual:        true,
		ContestType:      "icpc",
		CreateTime:       time.Now(),
		StartTime:        time.Now(),
		EndTime:          time.Now(),
		LockBoardTime:    time.Now(),
		HideOthersStatus: false,
	}
	id, err := cm.Insert(tx, con)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
