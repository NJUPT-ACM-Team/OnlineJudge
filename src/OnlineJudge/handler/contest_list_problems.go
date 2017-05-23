package handler

import (
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
	// if contest is private, we have no access for problem list
	if cst.IsPrivate() {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	// list problems
	ContestListProblems_BuildResponse(tx, response, req, cst.ContestId, this.session.GetUserId())
}

func (this *UserHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	// Check Access
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	cst, err := models.Query_Contest_By_ContestId(tx,
		req.GetContestId(), nil, nil)
	PanicOnError(err)
	// if contest is private, we need to check if we are in contest_users.
	if cst.IsPrivate() {
		if CheckContestUser(tx, cst.ContestId, this.session.GetUserId()) == false {
			MakeResponseError(response, this.debug, PBUnauthorized, nil)
			return
		}
	}
	// list problems
	ContestListProblems_BuildResponse(tx, response, req, cst.ContestId, this.session.GetUserId())
}

func CheckContestUser(tx *sqlx.Tx, contest_id, user_id int64) bool {
	_, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	if err != nil {
		return false
	}
	return true
}

func ContestListProblems_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestListProblemsResponse,
	req *api.ContestListProblemsRequest,
	contest_id int64,
	user_id int64) {

}
