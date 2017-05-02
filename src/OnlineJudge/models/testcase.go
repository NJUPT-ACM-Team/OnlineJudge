package models

import (
	"OnlineJudge/base"

	"github.com/jmoiron/sqlx"
)

/*
$OJDATA/testcases/$PID/input/id.in
$OJDATA/testcases/$PID/output/id.out
$OJDATA/testcases/$PID/spj/spj.cpp
*/

type TestCase struct {
	CaseId int64 `db:"case_id"`
	Input  []byte
	// InputPath  string
	InputMD5 []byte `db:"input_md5"`
	Output   []byte
	// OutputPath string
	OutputMD5 []byte `db:"output_md5"`

	MetaPidFK int64 `db:"meta_pid_fk"`
}

type TestCaseModel struct {
	Model
}

func NewTestCaseModel() *TestCaseModel {
	return &TestCaseModel{Model{Table: "TestCases"}}
}

func (this *TestCaseModel) Insert(tx *sqlx.Tx, tc *TestCase) (int64, error) {
	// hash input and output
	last_insert_id, err := this.InlineInsert(tx, tc, nil, []string{"case_id"})
	if err != nil {
		return 0, err
	}
	return last_insert_id, nil
}

func (this *TestCaseModel) InsertTestCase(tx *sqlx.Tx, input []byte, output []byte, meta_pid int64) (int64, error) {
	tc := &TestCase{
		Input:     input,
		Output:    output,
		MetaPidFK: meta_pid,
	}
	tc.InputMD5 = base.MD5Hash(tc.Input)
	tc.OutputMD5 = base.MD5Hash(tc.Output)
	return this.Insert(tx, tc)
}
