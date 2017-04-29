package models

import (
	"github.com/jmoiron/sqlx"
)

type TimeMemoryLimit struct {
	LimitId       int64 `db:limit_id`
	TimeLimit     int   `db:time_limit`
	MemoryLimit   int   `db:memory_limit`
	CaseTimeLimit int   `db:case_time_limit`
	LangIdFK      int64 `db:lang_id_fk`
	MetaPidFK     int64 `db:meta_pid_fk`
}

type TimeMemoryLimitModel struct {
	Model
}

func NewTimeMemoryLimitModel() *TestCaseModel {
	return &TimeMemoryLimitModel{Model{Table: "TimeMemoryLimits"}}
}
