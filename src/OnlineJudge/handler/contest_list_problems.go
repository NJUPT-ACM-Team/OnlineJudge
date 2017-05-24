package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
}

func (this *BasicHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	// Check Access
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	cst, err := models.Query_Contest_By_ContestId(tx,
		req.GetContestId(), nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	// if contest is private, we have no access for problem list
	if cst.IsPrivate() {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	// list problems
	ContestListProblems_BuildResponse(tx, response, req, false, cst.ContestId, this.session.GetUserId())
}

func (this *UserHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	// Check Access
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	cst, err := models.Query_Contest_By_ContestId(tx,
		req.GetContestId(), nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	// if contest is private, we need to check if we are in contest_users.
	if cst.IsPrivate() {
		check, err := CheckContestUser(tx, cst.ContestId, this.session.GetUserId())
		PanicOnError(err)
		if check == false {
			MakeResponseError(response, this.debug, PBUnauthorized, nil)
			return
		}
	}
	// list problems
	show_details := false
	if cst.CreatorId == this.session.GetUserId() {
		show_details = true
	}
	ContestListProblems_BuildResponse(tx, response, req, show_details, cst.ContestId, this.session.GetUserId())
}

func CheckContestUser(tx *sqlx.Tx, contest_id, user_id int64) (bool, error) {
	cst, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	if err != nil {
		return false, err
	}
	if cst == nil {
		return false, nil
	}
	return true, nil
}

func ContestListProblems_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestListProblemsResponse,
	req *api.ContestListProblemsRequest,
	show_details bool,
	contest_id int64,
	user_id int64) {

	cps, err := models.XQuery_Contest_List_Problems(
		tx, contest_id, user_id)
	PanicOnError(err)

	lines := []*api.ContestListProblemsResponse_PerLine{}
	for _, cp := range cps {
		// TODO: add ac count
		line := &api.ContestListProblemsResponse_PerLine{
			Label:  cp.Label,
			Alias:  cp.Alias,
			Status: cp.Status,
		}
		if show_details {
			line.Title = cp.Title
			line.Sid = base.GenSid(&base.Pid{OJName: cp.OJName, OJPid: cp.OJPid})
		}
		lines = append(lines, line)
	}
	response.Lines = lines
}
