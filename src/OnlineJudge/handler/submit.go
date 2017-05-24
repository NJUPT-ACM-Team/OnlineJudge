package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/db"
	"OnlineJudge/irpc"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"log"
	"time"
)

func (this *AdminHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer PanicHandler(response, this.debug)
	Submit_BuildResponse(this.dbu, response, req,
		this.session.GetUserId(), this.session.GetIPAddr(), true, this.debug)
}

func (this *BasicHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

// Need to be tested
// Depend on MetaProblems, OJInfo,
func (this *UserHandler) Submit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer PanicHandler(response, this.debug)
	Submit_BuildResponse(this.dbu, response, req,
		this.session.GetUserId(), this.session.GetIPAddr(), false, this.debug)
}

func Submit_BuildResponse(
	dbu *db.DBUtil,
	response *api.SubmitResponse,
	req *api.SubmitRequest,
	user_id int64,
	ip_addr string,
	use_hide bool,
	debug bool) {

	// Parse ProblemSid
	pid, err := base.ParseSid(req.GetProblemSid())
	if err != nil {
		MakeResponseError(response, debug, PBBadRequest, err)
		return
	}
	tx := dbu.MustBegin()
	defer dbu.Rollback()

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(tx, pid.OJName,
		pid.OJPid, []string{"meta_pid", "hide", "is_spj"}, nil)
	PanicOnError(err)
	if mp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// if visible
	if mp.Hide == true && !use_hide {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// Add Submission
	subm := models.NewSubmissionModel()
	sub := &models.Submission{
		Status:       "Pending",
		StatusCode:   "wt",
		SubmitTime:   time.Now(),
		Code:         req.GetCode(),
		SubmitIPAddr: ip_addr,
		IsShared:     req.GetIsShared(),

		IsContest: false,
		IsSpj:     mp.IsSpj,
		MetaPidFK: mp.MetaPid,
		UserIdFK:  user_id,
		LangIdFK:  req.GetLanguageId(),
	}
	run_id, err := subm.Insert(tx, sub)
	PanicOnError(err)

	response.RunId = run_id
	dbu.MustCommit()

	// Use RPC to call Daemon to judge the submission
	go func(dbu *db.DBUtil) {
		tx := dbu.MustBegin()
		defer dbu.Rollback()

		helper := irpc.NewHelper()

		if err := helper.Connect(); err != nil {
			// Log the error
			log.Println("connect error:" + err.Error())
			if err := subm.SetSystemError(tx, run_id); err != nil {
				log.Println("set error:" + err.Error())

			}
			dbu.Commit()
			return
		}
		defer helper.Disconnect()

		helper.NewClient()
		res, err := helper.StartJudging(run_id)

		if err != nil || res.Received != true {
			// Log the error
			log.Println("start judging error:" + err.Error())
			if err := subm.SetSystemError(tx, run_id); err != nil {
				log.Println("set error:" + err.Error())
			}
			dbu.Commit()
			return
		}
	}(dbu)
}
