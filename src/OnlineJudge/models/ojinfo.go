package models

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

type OJInfo struct {
	OJId         int64  `db:"oj_id"`
	OJName       string `db:"oj_name"`
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

func (this *OJInfoModel) Insert(tx *sqlx.Tx, oj *OJInfo) (int64, error) {
	last_insert_id, err := this.InlineInsert(tx, oj, nil, []string{"oj_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *OJInfoModel) QueryByName(tx *sqlx.Tx, name string, required []string, excepts []string) (*OJInfo, error) {
	ojinfo := OJInfo{}
	str_fields, err := GenerateSelectSQL(ojinfo, required, excepts)
	// fmt.Println(str_fields)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&ojinfo, fmt.Sprintf("SELECT %s FROM %s WHERE oj_name=?", str_fields, this.Table), name); err != nil {
		return nil, err
	}
	return &ojinfo, nil
}

func (this *OJInfoModel) QueryIdByName(tx *sqlx.Tx, name string) (int64, error) {
	ojinfo, err := this.QueryByName(tx, name, []string{"oj_id"}, nil)
	if err != nil {
		return 0, err
	}
	if ojinfo.OJId == 0 {
		return 0, errors.New("Failed to get oj_id")
	}
	return ojinfo.OJId, nil
}

func (this *OJInfoModel) QueryAll(tx *sqlx.Tx, required []string, excepts []string) ([]OJInfo, error) {
	ojs := []OJInfo{}
	str_fields, err := GenerateSelectSQL(OJInfo{}, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Select(&ojs, fmt.Sprintf("SELECT %s FROM %s", str_fields, this.Table)); err != nil {
		return nil, err
	}
	return ojs, nil

}

func (this *OJInfoModel) Update(tx *sqlx.Tx, ojinfo *OJInfo, pk string, required []string, excepts []string) error {
	if pk == "" {
		pk = "oj_id"
	}
	if err := this.InlineUpdate(tx, ojinfo, pk, required, excepts); err != nil {
		return err
	}
	return nil
}
