package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"OnlineJudge/base"
	"OnlineJudge/models/db"
	"github.com/jmoiron/sqlx"
)

type Model struct {
	DB    *sqlx.DB
	Table string
}

func (this *Model) OpenDB() error {
	db.Init()
	var err error
	this.DB, err = db.NewDB()
	return err
}

func (this *Model) CloseDB() {
	this.DB.Close()
}

func (this *Model) GetAllFields(v interface{}) ([]string, error) {
	mapping := this.DB.Mapper.TypeMap(reflect.TypeOf(v))
	if len(mapping.Paths) == 0 {
		return nil, errors.New("Empty struct definition.")
	}
	fields := []string{}
	for k, _ := range mapping.Paths {
		fields = append(fields, k)
	}
	return fields, nil
}

func (this *Model) GenerateInsertSQL(st interface{}, table string, excepts []string) (string, error) {
	all_fields, err := this.GetAllFields(st)
	if err != nil {
		return "", err
	}
	fields := []string{}
	values := []string{}
	for _, v := range all_fields {
		if excepts == nil || len(excepts) == 0 || !base.ArrayContains(excepts, v) {
			fields = append(fields, v)
			values = append(values, ":"+v)
		}
	}
	cols := "(" + strings.Join(fields, ",") + ")"
	vals := "(" + strings.Join(values, ",") + ")"
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, cols, vals), nil
}

func (this *Model) GenerateSelectSQL(st interface{}, required []string, excepts []string) (string, error) {
	all_fields, err := this.GetAllFields(st)
	if err != nil {
		return "", err
	}
	fields := []string{}
	if required != nil && len(required) > 0 {
		for _, v := range required {
			if base.ArrayContains(all_fields, v) {
				fields = append(fields, v)
			}
		}
	} else {
		if excepts != nil && len(excepts) > 0 {
			for _, v := range all_fields {
				if !base.ArrayContains(excepts, v) {
					fields = append(fields, v)
				}
			}
		} else {
			fields = all_fields
		}
	}
	return strings.Join(fields, ","), nil
}

func (this *Model) InlineInsert(st interface{}, excepts []string) (int, error) {
	sql_insert, err := this.GenerateInsertSQL(st, this.Table, excepts)
	if err != nil {
		return 0, err
	}
	tx, err := this.DB.Beginx()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()
	_, err = tx.NamedExec(sql_insert, st)
	if err != nil {
		return 0, err
	}
	var last_insert_id int
	if err := tx.Get(&last_insert_id, "SELECT LAST_INSERT_ID()"); err != nil {
		return 0, err
	}
	err = tx.Commit()
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}
