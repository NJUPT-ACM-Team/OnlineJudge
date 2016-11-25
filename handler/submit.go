package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
	"errors"
	"fmt"
	"time"
)

// Need to be tested
func (this *Handler) Submit(subreq *api.SubmitRequest) *api.SubmitResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()

	// Parse ProblemSid
	var oj_name string
	var oj_pid int
	fmt.Sscanf(subreq.GetProblemSid(), "%s#%d", &oj_name, &oj_pid)

	sql1 := `
	SELECT meta_pid, hide
	FROM MetaProblems
	WHERE oj_pid=? AND oj_id_fk=(SELECT oj_id FROM OJInfo WHERE name=?)
	`
	type MetaCollection1 struct {
		MetaPid int64 `db:"meta_pid"`
		Hide    int
	}
	mc1 := MetaCollection1{}
	if err := this.tx.Get(&mc1, sql1, oj_pid, oj_name); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	ErrProblemNotFound := errors.New("Problem does not exist or not visible.")
	if mc1.MetaPid == 0 {
		return api.NewSubmitResponseError(true, 404, ErrProblemNotFound)
	}

	// if visible
	if mc1.Hide == 1 && this.session.privilege != "root" {
		return api.NewSubmitResponseError(true, 404, ErrProblemNotFound)
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	sub := &models.Submission{
		Status:     "Pending",
		StatusCode: "wt",
		SubmitTime: time.Now(),
		Code:       subreq.GetCode(),
		IPAddr:     subreq.GetIpAddr(),
		IsShared:   subreq.GetIsShared(),

		IsContest: false,
		MetaPidFK: mc1.MetaPid,
		UserIdFK:  this.session.user_id,
	}
	run_id, err := subm.Insert(this.tx, sub)
	if err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	if err := this.Commit(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}

	// Use RPC to judge the submission

	return &api.SubmitResponse{
		RunId: run_id,
	}
}
