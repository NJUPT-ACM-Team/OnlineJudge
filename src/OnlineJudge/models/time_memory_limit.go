package models

import (
	"github.com/jmoiron/sqlx"
)

const (
	TimeMemoryLimit_TableName  = "TimeMemoryLimits"
	TimeMemoryLimit_PrimaryKey = "limit_id"
)

type TimeMemoryLimit struct {
	LimitId       int64  `db:"limit_id"`
	TimeLimit     int    `db:"time_limit"`
	MemoryLimit   int    `db:"memory_limit"`
	CaseTimeLimit int    `db:"case_time_limit"`
	Language      string `db:"language"`
	MetaPidFK     int64  `db:"meta_pid_fk"`
}

type TimeMemoryLimitModel struct {
	Model
}

func NewTimeMemoryLimitModel() *TimeMemoryLimitModel {
	return &TimeMemoryLimitModel{Model{Table: TimeMemoryLimit_TableName}}
}

func (this *TimeMemoryLimitModel) Insert(tx *sqlx.Tx, tm *TimeMemoryLimit) (int64, error) {
	last_insert_id, err := this.InlineInsert(tx, tm, nil, []string{TimeMemoryLimit_PrimaryKey})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}
