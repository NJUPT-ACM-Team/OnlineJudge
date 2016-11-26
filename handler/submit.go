package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
	"errors"
	"fmt"
	"time"
)

var (
	ErrProblemNotFound = errors.New("Problem does not exist or not visible.")
	ErrNotLogin        = errors.New("You have not logged in, log in first please.")
)

// Need to be tested
// Depend on MetaProblems, OJInfo,
func (this *Handler) Submit(subreq *api.SubmitRequest) *api.SubmitResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()

	// if login
	if this.session.IsLogin() == false {
		return api.NewSubmitResponseError(true, 403, ErrNotLogin)
	}

	// Parse ProblemSid
	var oj_name string
	var oj_pid int
	fmt.Sscanf(subreq.GetProblemSid(), "%s#%d", &oj_name, &oj_pid)

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, oj_name, oj_pid, []string{"meta_pid", "hide"}, nil)
	if err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}

	if mp.MetaPid == 0 {
		return api.NewSubmitResponseError(true, 404, ErrProblemNotFound)
	}

	// if visible
	privilege, _ := this.session.GetPrivilege()
	if mp.Hide == 1 && privilege != "root" {
		return api.NewSubmitResponseError(true, 404, ErrProblemNotFound)
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	user_id, err := this.session.GetUserId()
	if err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	sub := &models.Submission{
		Status:     "Pending",
		StatusCode: "wt",
		SubmitTime: time.Now(),
		Code:       subreq.GetCode(),
		IPAddr:     subreq.GetIpAddr(),
		IsShared:   subreq.GetIsShared(),

		IsContest: false,
		MetaPidFK: mp.MetaPid,
		UserIdFK:  user_id,
	}
	run_id, err := subm.Insert(this.tx, sub)
	if err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	if err := this.Commit(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}

	// Use RPC to call Daemon to judge the submission

	return &api.SubmitResponse{
		RunId: run_id,
	}
}
