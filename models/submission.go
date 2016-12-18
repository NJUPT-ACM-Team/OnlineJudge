package models

import (
	"github.com/jmoiron/sqlx"

	"errors"
	"time"
)

type Submission struct {
	RunId           int64 `db:"run_id"`
	Status          string
	StatusCode      string    `db:"status_code"`
	TestcasesPassed int       `db:"testcases_passed"`
	SubmitTime      time.Time `db:"submit_time"`
	TimeUsed        int       `db:"time_used"`
	MemoryUsed      int       `db:"memory_used"`
	Code            string
	LangIdFK        int64  `db:"lang_id_fk"`
	CEInfo          string `db:"ce_info"`
	SubmitIPAddr    string `db:"submit_ip_addr"`
	IsShared        bool   `db:"is_shared"`
	IsPrivate       bool   `db:"is_private"`
	IsSpj           bool   `db:"is_spj"`

	IsContest bool  `db:"is_contest"`
	CPIdFK    int64 `db:"cp_id_fk"`
	CUIdFK    int64 `db:"cu_id_fk"`
	MetaPidFK int64 `db:"meta_pid_fk"`
	UserIdFK  int64 `db:"user_id_fk"`
}

type SubmissionModel struct {
	Model
}

func NewSubmissionModel() *SubmissionModel {
	return &SubmissionModel{Model{Table: "Submissions"}}
}

func (this *SubmissionModel) Validate(sub *Submission) error {
	if sub.IsContest == true {
		if sub.CPIdFK == 0 {
			return errors.New("Lack of foreign key references to ContestsProblems.")
		}
		if sub.CUIdFK == 0 {
			return errors.New("Lack of foreign key references to ContestsUsers.")
		}
	} else {
		if sub.MetaPidFK == 0 {
			return errors.New("Lack of foreign key references to MetaProblems.")
		}
		if sub.UserIdFK == 0 {
			return errors.New("Lack of foreign key references to Users.")
		}
	}
	return nil
}

func (this *SubmissionModel) Insert(tx *sqlx.Tx, sub *Submission) (int64, error) {
	if err := this.Validate(sub); err != nil {
		return 0, err
	}
	last_insert_id, err := this.InlineInsert(tx, sub, nil, []string{"run_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil

}

func (this *SubmissionModel) Update(tx *sqlx.Tx, sub *Submission, pk string, required []string, excepts []string) error {
	if pk == "" {
		pk = "run_id"
	}
	if err := this.InlineUpdate(tx, sub, "run_id", required, excepts); err != nil {
		return err
	}
	return nil
}

func (this *SubmissionModel) UpdateStatus(tx *sqlx.Tx, sub *Submission) error {
	required := []string{"status", "status_code", "testcases_passed"}
	if err := this.Update(tx, sub, "", required, nil); err != nil {
		return err
	}
	return nil
}

func (this *SubmissionModel) SetSystemError(tx *sqlx.Tx, run_id int64) error {
	sub := &Submission{
		RunId:           run_id,
		Status:          "System Error",
		StatusCode:      "se",
		TestcasesPassed: 0,
	}
	if err := this.UpdateStatus(tx, sub); err != nil {
		return err
	}
	return nil
}
