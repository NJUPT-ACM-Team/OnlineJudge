package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Query_MetaProblem_By_OJName_OJPid(tx *sqlx.Tx, oj_name string, oj_pid int, required []string, excepts []string) (*MetaProblem, error) {
	mp := &MetaProblem{}
	str_fields, err := GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT %s FROM MetaProblems
	WHERE oj_pid=? AND oj_name=?
	`
	if err := tx.Get(mp, fmt.Sprintf(sql, str_fields), oj_pid, oj_name); err != nil {
		return nil, err
	}
	return mp, nil
}

func Query_User_By_Username(tx *sqlx.Tx, name string, required []string, excepts []string) (*User, error) {
	user := &User{}
	str_fields, err := GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT %s FROM Users
	WHERE username=?
	`
	if err := tx.Get(user, fmt.Sprintf(sql, str_fields), name); err != nil {
		return nil, err
	}
	return user, nil
}

func Query_Languages_By_OJIdFK(tx *sqlx.Tx, oj_id_fk int64, required []string, excepts []string) ([]Language, error) {
	lang := Language{}
	langs := []Language{}
	str_fields, err := GenerateSelectSQL(&lang, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT %s FROM Languages
	WHERE oj_id_fk=?
	`
	if err := tx.Select(&langs, fmt.Sprintf(sql, str_fields), oj_id_fk); err != nil {
		return nil, err
	}
	return langs, nil
}
