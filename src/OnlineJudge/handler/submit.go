package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/irpc"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"log"
	"time"
)

func (this *AdminHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer func() {
		if err := recover(); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err.(error))
		}
	}()
	this.OpenDBU()
	defer this.CloseDBU()
	//	tx := this.dbu.MustBegin()
}

func (this *BasicHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

// Need to be tested
// Depend on MetaProblems, OJInfo,
func (this *UserHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
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
	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, pid.OJName, pid.OJPid, []string{"meta_pid", "hide", "is_spj"}, nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	if mp.MetaPid == 0 {
		MakeResponseError(response, this.debug, PBProblemNotFound, nil)
		return
	}

	// if visible
	if mp.Hide == true && this.session.GetPrivilege() != "root" {
		MakeResponseError(response, this.debug, PBProblemNotFound, nil)
		return
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	user_id := this.session.GetUserId()
	sub := &models.Submission{
		Status:       "Pending",
		StatusCode:   "wt",
		SubmitTime:   time.Now(),
		Code:         req.GetCode(),
		SubmitIPAddr: this.session.GetIPAddr(),
		IsShared:     req.GetIsShared(),

		IsContest: false,
		IsSpj:     mp.IsSpj,
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
	response.RunId = run_id

	// Use RPC to call Daemon to judge the submission

	helper := irpc.NewBackendHelper()

	if err := helper.Connect(); err != nil {
		// Log the error
		log.Println(err)
		if err := subm.SetSystemError(this.tx, run_id); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
		}
		if err := this.Commit(); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
		}
		return
	}
	defer helper.Disconnect()

	helper.NewClient()
	res, err := helper.Submit(run_id)

	if err != nil || res.Received != true {
		// Log the error
		log.Println(err)
		if err := subm.SetSystemError(this.tx, run_id); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
		}
		if err := this.Commit(); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
		}
		return
	}
}
