package models

import (
	"github.com/jmoiron/sqlx"
)

type ContestUser struct {
	CUId         int64  `db:"cu_id"`
	UserIdFK     int64  `db:"user_id_fk"`
	ContestIdFK  int64  `db:"contest_id_fk"`
	ContestMotto string `db:"contest_motto"`
}

type ContestUserModel struct {
	Model
}

func NewContestUserModel() *ContestUserModel {
	return &ContestUserModel{Model{Table: "ContestUsers"}}
}

func (this *ContestUserModel) Insert(tx *sqlx.Tx, cu *ContestUser) (int64, error) {
	id, err := this.InlineInsert(tx, cu, nil, []string{"cu_id"})
	if err != nil {
		return 0, err
	}
	return id, nil
}
