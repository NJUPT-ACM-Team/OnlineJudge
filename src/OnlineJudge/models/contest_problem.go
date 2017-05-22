package models

import (
	"github.com/jmoiron/sqlx"
)

type ContestProblem struct {
	CPId        int64   `db:"cp_id"`
	MetaPidFK   int64   `db:"meta_pid_fk"`
	ContestIdFK int64   `db:"contest_id_fk"`
	Alias       string  `db:"alias"`
	Label       string  `db:"label"`
	Base        int     `db:"base"`
	Minp        int     `db:"minp"`
	ParaA       float32 `db:"para_a"`
	ParaB       float32 `db:"para_b"`
}

type ContestProblemModel struct {
	Model
}

func NewContestProblemModel() *ContestProblemModel {
	return &ContestProblemModel{Model{Table: "ContestProblems"}}
}

func (this *ContestProblemModel) Insert(tx *sqlx.Tx, cp *ContestProblem) (int64, error) {
	id, err := this.InlineInsert(tx, cp, nil, []string{"cp_id"})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *ContestProblemModel) DeleteProblemsByContestId(tx *sqlx.Tx, id int64) error {
	pk := "contest_id_fk"
	return this.InlineDelete(tx, &ContestProblem{ContestIdFK: id}, pk)
}
