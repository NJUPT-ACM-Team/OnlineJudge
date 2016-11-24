package models_basic

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MetaProblem struct {
	MetaPid           int `db:"meta_pid"`
	Title             string
	Description       string
	Input             string
	Output            string
	SampleIn          string `db:"sample_in"`
	SampleOut         string `db:"sample_out"`
	TimeLimit         int    `db:"time_limit"`
	CaseTimeLimit     int    `db:"case_time_limit"`
	MemoryLimit       int    `db:"memory_limit"`
	NumberOfTestCases int    `db:"number_of_testcases"`
	Source            string
	Hint              string
	Hide              int
	OJIdFK            int `db:"oj_id_fk"`
	OJPid             int `db:"oj_pid"`
}

type MetaProblemModel struct {
	Model
}

func NewMetaProblemModel() *MetaProblemModel {
	return &MetaProblemModel{Model{Table: "MetaProblems"}}
}

func (this *MetaProblemModel) Insert(tx *sqlx.Tx, mp *MetaProblem) (int64, error) {
	last_insert_id, err := this.InlineInsert(tx, mp, nil, []string{"meta_pid"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *MetaProblemModel) QueryById(tx *sqlx.Tx, id int, required []string, excepts []string) (*MetaProblem, error) {
	mp := MetaProblem{}
	str_fields, err := this.GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&mp, fmt.Sprintf("SELECT %s FROM %s WHERE meta_pid=?", str_fields, this.Table), id); err != nil {
		return nil, err
	}
	return &mp, nil
}

func (this *MetaProblemModel) QueryByOJIdAndPid(tx *sqlx.Tx, oj_id int, pid int, required []string, excepts []string) (*MetaProblem, error) {

	mp := MetaProblem{}
	str_fields, err := this.GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}
	if err := tx.Get(&mp, fmt.Sprintf("SELECT %s FROM %s WHERE oj_id_fk=? AND oj_pid=?", str_fields, this.Table), oj_id, pid); err != nil {
		return nil, err
	}
	return &mp, nil
}
