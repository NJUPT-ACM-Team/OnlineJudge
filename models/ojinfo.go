package models

import (
	"fmt"
	"time"
)

type OJInfo struct {
	OJId       int `db:"oj_id"`
	Name       string
	Version    string
	Int64IO    string
	JavaClass  string
	Status     string
	StatusInfo string `db:"status_info"`
	LastCheck  time.Time
}

type OJInfoModel struct {
	Model
}

func NewOJInfoModel() *OJInfoModel {
	return &OJInfoModel{Model{Table: "OJInfo"}}
}

func (this *OJInfoModel) Insert(oj *OJInfo) (int, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()
	last_insert_id, err := this.InlineInsert(oj, []string{"oj_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *OJInfoModel) QueryByName(name string) (*OJInfo, error) {
	if err := this.OpenDB(); err != nil {
		return nil, err
	}
	defer this.CloseDB()
	ojinfo := OJInfo{}
	if err := this.DB.Get(&ojinfo, fmt.Sprintf("SELECT * FROM %s WHERE name=?", this.Table), name); err != nil {
		return nil, err
	}
	return &ojinfo, nil
}

func (this *OJInfoModel) QueryAll() ([]OJInfo, error) {
	if err := this.OpenDB(); err != nil {
		return nil, err
	}
	defer this.CloseDB()
	ojs := []OJInfo{}
	if err := this.DB.Select(&ojs, fmt.Sprintf("SELECT * FROM %s", this.Table)); err != nil {
		return nil, err
	}
	return ojs, nil

}
