package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"

	"fmt"
	"time"
)

// Need to be tested
// Depend on MetaProblems, OJInfo,
func (this *Handler) Submit(subreq *api.SubmitRequest) *api.SubmitResponse {
	var response = &api.SubmitResponse{}
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	defer this.CloseDB()

	// if login
	if this.session.IsLogin() == false {
		api.MakeResponseError(response, this.debug, api.PBLoginRequired, nil)
		return response
	}

	// Parse ProblemSid
	var oj_name string
	var oj_pid int
	fmt.Sscanf(subreq.GetProblemSid(), "%s#%d", &oj_name, &oj_pid)

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, oj_name, oj_pid, []string{"meta_pid", "hide"}, nil)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}

	if mp.MetaPid == 0 {
		api.MakeResponseError(response, this.debug, api.PBProblemNotFound, nil)
		return response
	}

	// if visible
	privilege, _ := this.session.GetPrivilege()
	if mp.Hide == 1 && privilege != "root" {
		api.MakeResponseError(response, this.debug, api.PBProblemNotFound, nil)
		return response
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	user_id, err := this.session.GetUserId()
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	sub := &models.Submission{
		Status:     "Pending",
		StatusCode: "wt",
		SubmitTime: time.Now(),
		Code:       subreq.GetCode(),
		// IPAddr:     subreq.GetIpAddr(),
		IsShared: subreq.GetIsShared(),

		IsContest: false,
		MetaPidFK: mp.MetaPid,
		UserIdFK:  user_id,
	}
	run_id, err := subm.Insert(this.tx, sub)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	if err := this.Commit(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}

	// Use RPC to call Daemon to judge the submission

	// Return
	response.RunId = run_id
	return response

}
