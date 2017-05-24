package handler

import (
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
	// "github.com/jmoiron/sqlx"

	"time"
)

func (this *AdminHandler) ContestSubmit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer PanicHandler(response, this.debug)
	Submit_BuildResponse(this.dbu, response, req,
		this.session.GetUserId(), this.session.GetIPAddr(), true, this.debug)
}

func (this *BasicHandler) ContestSubmit(response *api.SubmitResponse, req *api.SubmitRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

func (this *UserHandler) ContestSubmit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Submit {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	Submit_BuildResponse(this.dbu, response, req,
		this.session.GetUserId(), this.session.GetIPAddr(), false, this.debug)
}

func ContestSubmit_BuildResponse(
	dbu *db.DBUtil,
	response *api.SubmitResponse,
	req *api.SubmitRequest,
	user_id int64,
	ip_addr string,
	use_hide bool,
	debug bool) {

	contest_id := req.GetContestId()
	label := req.GetProblemSid()

	tx := dbu.MustBegin()
	defer dbu.Rollback()

	cp, err := models.Query_ContestProblem_By_ContestId_And_Label(
		tx, contest_id, label)
	PanicOnError(err)
	if cp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	mp, err := models.Query_MetaProblem_By_MetaPid(
		tx, cp.MetaPidFK, []string{"meta_pid", "hide", "is_spj"}, nil)
	PanicOnError(err)
	if mp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// Get contestuser by contest_id, user_id
	cu, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	PanicOnError(err)
	if cu == nil {
		MakeResponseError(response, debug, PBUnauthorized, nil)
		return
	}
	cu_id := cu.CUId

	// Add Submission
	subm := models.NewSubmissionModel()
	sub := &models.Submission{
		Status:       "Pending",
		StatusCode:   "wt",
		SubmitTime:   time.Now(),
		Code:         req.GetCode(),
		SubmitIPAddr: ip_addr,
		IsShared:     req.GetIsShared(),

		IsContest: true,
		IsSpj:     mp.IsSpj,
		MetaPidFK: mp.MetaPid,
		CPIdFK:    cp.CPId,
		CUIdFK:    cu_id,
		LangIdFK:  req.GetLanguageId(),
	}
	run_id, err := subm.Insert(tx, sub)
	PanicOnError(err)
	response.RunId = run_id
	dbu.MustCommit()

	go CallJudging(dbu, run_id)
}
