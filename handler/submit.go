package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"time"
)

// Need to be tested
// Depend on MetaProblems, OJInfo,
func (this *Handler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// if login
	if this.session.IsLogin() == false {
		MakeResponseError(response, this.debug, PBLoginRequired, nil)
		return
	}

	// Parse ProblemSid
	pid, err := base.ParseSid(req.GetProblemSid())
	if err != nil {
		MakeResponseError(response, this.debug, PBBadRequest, err)
		return
	}
	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, pid.OJName, pid.OJPid, []string{"meta_pid", "hide"}, nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	if mp.MetaPid == 0 {
		MakeResponseError(response, this.debug, PBProblemNotFound, nil)
		return
	}

	// if visible
	if mp.Hide == 1 && this.session.GetPrivilege() != "root" {
		MakeResponseError(response, this.debug, PBProblemNotFound, nil)
		return
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	user_id := this.session.GetUserId()
	sub := &models.Submission{
		Status:     "Pending",
		StatusCode: "wt",
		SubmitTime: time.Now(),
		Code:       req.GetCode(),
		IPAddr:     this.session.GetIPAddr(),
		IsShared:   req.GetIsShared(),

		IsContest: false,
		MetaPidFK: mp.MetaPid,
		UserIdFK:  user_id,
		LangIdFK:  req.GetLanguageId(),
	}
	run_id, err := subm.Insert(this.tx, sub)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	if err := this.Commit(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// Use RPC to call Daemon to judge the submission

	// Return
	response.RunId = run_id
}
