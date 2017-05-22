package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"OnlineJudge/base"
	//"OnlineJudge/models/db"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
)

type Model struct {
	Table string
}

func GetAllFields(v interface{}) ([]string, error) {
	mapper := reflectx.NewMapperTagFunc("db", strings.ToLower, func(value string) string {
		if strings.Contains(value, ",") {
			return strings.Split(value, ",")[0]
		}
		return value
	})
	mapping := mapper.TypeMap(reflect.TypeOf(v))
	if len(mapping.Paths) == 0 {
		return nil, errors.New("Empty struct definition.")
	}
	fields := []string{}
	for k, _ := range mapping.Paths {
		fields = append(fields, k)
	}
	return fields, nil
}

func FilterFields(all_fields []string, required []string, excepts []string) ([]string, error) {
	if !base.IsNilOrZero(excepts) && !base.IsNilOrZero(required) {
		return nil, errors.New("Parameters conficted, using only required or only excepts.")
	}
	fields := []string{}
	if !base.IsNilOrZero(required) {
		for _, v := range required {
			if base.ArrayContains(all_fields, v) {
				fields = append(fields, v)
			}
		}
	} else {
		if !base.IsNilOrZero(excepts) {
			for _, v := range all_fields {
				if !base.ArrayContains(excepts, v) {
					fields = append(fields, v)
				}
			}
		} else {
			fields = all_fields
		}
	}
	return fields, nil
}

func GenerateInsertSQL(st interface{}, table string, required []string, excepts []string) (string, error) {
	all_fields, err := GetAllFields(st)
	if err != nil {
		return "", err
	}
	reversed_fields, err := FilterFields(all_fields, required, excepts)
	if err != nil {
		return "", err
	}
	fields := []string{}
	values := []string{}
	for _, v := range reversed_fields {
		fields = append(fields, v)
		values = append(values, ":"+v)
	}
	cols := "(" + strings.Join(fields, ",") + ")"
	vals := "(" + strings.Join(values, ",") + ")"
	return fmt.Sprintf("INSERT INTO %s %s VALUES %s", table, cols, vals), nil
}

func GenerateSelectSQL(st interface{}, required []string, excepts []string) (string, error) {
	all_fields, err := GetAllFields(st)
	if err != nil {
		return "", err
	}
	fields, err := FilterFields(all_fields, required, excepts)
	if err != nil {
		return "", err
	}
	mapper := reflectx.NewMapperTagFunc("db", strings.ToLower, func(value string) string {
		if strings.Contains(value, ",") {
			return strings.Split(value, ",")[0]
		}
		return value
	})
	for k, v := range fields {
		var dft string
		val := mapper.FieldByName(reflect.ValueOf(st), v)
		switch val.Interface().(type) {
		case int:
			dft = "0"
		case int64:
			dft = "0"
		case string, []byte:
			dft = "''"
		case time.Time:
			dft = "CURRENT_TIMESTAMP"
		case bool:
			dft = "0"
		}
		fields[k] = fmt.Sprintf("COALESCE(%s, %s) AS %s", v, dft, v)
	}
	return strings.Join(fields, ","), nil
}

func GenerateUpdateSQL(st interface{}, pk string, table string, required []string, excepts []string) (string, error) {
	all_fields, err := GetAllFields(st)
	if err != nil {
		return "", err
	}
	fields, err := FilterFields(all_fields, required, excepts)
	values := []string{}
	for _, v := range fields {
		values = append(values, v+"=:"+v)
	}
	str_fields := strings.Join(values, ",")
	/*
		mapper := reflectx.NewMapperTagFunc("db", strings.ToLower, func(value string) string {
			if strings.Contains(value, ",") {
				return strings.Split(value, ",")[0]
			}
			return value
		})
	*/
	return fmt.Sprintf("UPDATE %s SET %s WHERE %s=%s", table, str_fields, pk, ":"+pk), nil
}

func (this *Model) InlineInsert(tx *sqlx.Tx, st interface{}, required []string, excepts []string) (int64, error) {
	sql_insert, err := GenerateInsertSQL(st, this.Table, required, excepts)
	if err != nil {
		return 0, err
	}
	res, err := tx.NamedExec(sql_insert, st)
	if err != nil {
		return 0, err
	}
	last_insert_id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *Model) InlineUpdate(tx *sqlx.Tx, st interface{}, pk string, required []string, excepts []string) error {
	sql_update, err := GenerateUpdateSQL(st, pk, this.Table, required, excepts)
	if err != nil {
		return err
	}
	if _, err := tx.NamedExec(sql_update, st); err != nil {
		return err
	}
	return nil
}

func JoinSQL(sqls ...string) string {
	return strings.Join(sqls, " ")
}

func (this *Model) InlineDelete(tx *sqlx.Tx, st interface{}, pk string) error {
	sql_del := fmt.Sprintf("DELETE FROM %s WHERE %s=:%s", this.Table, pk, pk)
	if _, err := tx.NamedExec(sql_del, st); err != nil {
		return err
	}
	return nil
}
