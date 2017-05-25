package handler

import (
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
	// "github.com/jmoiron/sqlx"

	// "log"
	"errors"
	"time"
)

func (this *AdminHandler) ContestSubmit(response *api.SubmitResponse, req *api.SubmitRequest) {
	defer PanicHandler(response, this.debug)
	ContestSubmit_BuildResponse(this.dbu, response, req,
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
	if access.Time != 0 && !access.Creator {
		MakeResponseError(response, this.debug, PBUnauthorized, errors.New("contest not started or ended"))
		return
	}
	ContestSubmit_BuildResponse(this.dbu, response, req,
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

	// log.Println("get in contest submit")

	contest_id := req.GetContestId()
	label := req.GetProblemSid()

	tx := dbu.MustBegin()
	defer dbu.Rollback()

	// add user to contestuser
	check, err := CheckContestUser(tx, contest_id, user_id)
	PanicOnError(err)
	if !check {
		err = AddUserToContest(tx, user_id, contest_id)
		PanicOnError(err)
	}

	cp, err := models.Query_ContestProblem_By_ContestId_And_Label(
		tx, contest_id, label)
	// log.Println("query contestproblem:", err.Error())
	PanicOnError(err)
	if cp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	mp, err := models.Query_MetaProblem_By_MetaPid(
		tx, cp.MetaPidFK, []string{"meta_pid", "hide", "is_spj"}, nil)
	PanicOnError(err)
	// log.Println("query metaproblem:", err.Error())
	if mp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// Get contestuser by contest_id, user_id
	cu, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	// log.Println("query contestuser:", err.Error())
	PanicOnError(err)
	// log.Println("get here 0")
	if cu == nil {
		MakeResponseError(response, debug, PBUnauthorized, nil)
		return
	}
	cu_id := cu.CUId

	// log.Println("get here 1")
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
	// log.Println("get here 2")
	dbu.MustCommit()

	go CallJudging(dbu, run_id)
}
