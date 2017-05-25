package models

import (
	"github.com/jmoiron/sqlx"

	"fmt"
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

func (this *ContestProblemModel) DeleteByContestIdAndLabel(tx *sqlx.Tx, id int64, label string) error {
	sql_del := fmt.Sprintf("DELETE FROM ContestProblems WHERE contest_id_fk=? AND label=?")
	if _, err := tx.Exec(sql_del, id, label); err != nil {
		return err
	}
	return nil
}

func (this *ContestProblemModel) Update(tx *sqlx.Tx, cp *ContestProblem, pk string, required []string, excepts []string) error {
	if pk == "" {
		pk = "cp_id"
	}
	if err := this.InlineUpdate(tx, cp, pk, required, excepts); err != nil {
		return err
	}
	return nil
}
