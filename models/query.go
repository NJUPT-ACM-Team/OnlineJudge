package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func Query_MetaProblem_By_OJName_OJPid(tx *sqlx.Tx, oj_name string, oj_pid int, required []string, excepts []string) (*MetaProblem, error) {
	mp := MetaProblem{}
	str_fields, err := GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT %s FROM MetaProblems
	WHERE oj_pid=? AND oj_id_fk=(SELECT oj_id FROM OJInfo WHERE name=?)
	`
	if err := tx.Get(&mp, fmt.Sprintf(sql, str_fields), oj_pid, oj_name); err != nil {
		return nil, err
	}
	return &mp, nil
}
