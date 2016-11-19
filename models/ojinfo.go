package models

import (
	"fmt"
	"time"
)

type OJInfo struct {
	OJId         int `db:"oj_id"`
	Name         string
	Version      string
	Int64IO      string
	JavaClass    string
	Status       string
	StatusInfo   string `db:"status_info"`
	LastCheck    time.Time
	CurrentIndex int `db:"current_index"`
}

type OJInfoModel struct {
	Model
}

func NewOJInfoModel() *OJInfoModel {
	return &OJInfoModel{Model{Table: "OJInfo"}}
}

func (this *OJInfoModel) Insert(oj *OJInfo) (int64, error) {
	if err := this.OpenDB(); err != nil {
		return 0, err
	}
	defer this.CloseDB()
	last_insert_id, err := this.InlineInsert(oj, nil, []string{"oj_id"})
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
	str_fields, err := this.GenerateSelectSQL(ojinfo, nil, nil)
	// fmt.Println(str_fields)
	if err != nil {
		return nil, err
	}
	if err := this.DB.Get(&ojinfo, fmt.Sprintf("SELECT %s FROM %s WHERE name=?", str_fields, this.Table), name); err != nil {
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
	str_fields, err := this.GenerateSelectSQL(OJInfo{}, nil, nil)
	if err != nil {
		return nil, err
	}
	if err := this.DB.Select(&ojs, fmt.Sprintf("SELECT %s FROM %s", str_fields, this.Table)); err != nil {
		return nil, err
	}
	return ojs, nil

}
