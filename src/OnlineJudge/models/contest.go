package models

import (
	"github.com/jmoiron/sqlx"

	"time"
)

type Contest struct {
	ContestId        int64 `db:"contest_id"`
	Title            string
	Description      string
	IsVirtual        bool      `db:"is_virtual"`
	ContestType      string    `db:"contest_type"`
	CreateTime       time.Time `db:"create_time"`
	StartTime        time.Time `db:"start_time"`
	EndTime          time.Time `db:"end_time"`
	LockBoardTime    time.Time `db:"lock_board_time"`
	HideOthersStatus bool      `db:"hide_others_status"`
}

type ContestModel struct {
	Model
}

func NewContestModel() *ContestModel {
	return &ContestModel{Model{Table: "Contests"}}
}

func (this *ContestModel) Insert(tx *sqlx.Tx, con *Contest) (int64, error) {
	id, err := this.InlineInsert(tx, con, nil, []string{"contest_id"})
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (this *ContestModel) Update(
	tx *sqlx.Tx, pk string, con *Contest,
	required []string,
	excepts []string) error {

	if pk == "" {
		pk = "contest_id"
	}
	if err := this.InlineUpdate(tx, con, pk, required, excepts); err != nil {
		return err
	}
	return nil
}
