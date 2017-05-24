package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"fmt"
	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ContestShow(response *api.ContestShowResponse, req *api.ContestShowRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

	ContestShow_BuildResponse(tx, response, req, int64(req.GetContestId()), true, this.debug)
}

func (this *BasicHandler) ContestShow(response *api.ContestShowResponse, req *api.ContestShowRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

	fmt.Println(req.GetContestId())
	ContestShow_BuildResponse(tx, response, req, int64(req.GetContestId()), false, this.debug)
}

func ContestShow_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestShowResponse,
	req *api.ContestShowRequest,
	contest_id int64,
	is_admin bool,
	debug bool) {

	// Query_Contest_By_ContestId(tx, contest_id, nil, nil)

	// Query Contest
	cst, err := models.Query_Contest_By_ContestId(tx, int64(contest_id), nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, debug, PBContestNotFound, nil)
		return
	}

	contest := &api.Contest{
		ContestId:     contest_id,
		Title:         cst.Title,
		Description:   cst.Description,
		IsVirtual:     cst.IsVirtual,
		ContestType:   cst.ContestType,
		StartTime:     cst.StartTime.String(),
		EndTime:       cst.EndTime.String(),
		LockBoardTime: cst.LockBoardTime.String(),
	}
	if cst.Password == "" {
		contest.Access = "public"
	} else {
		contest.Access = "private"
	}
	contest.Status = "ended"

	response.Contest = contest
}
