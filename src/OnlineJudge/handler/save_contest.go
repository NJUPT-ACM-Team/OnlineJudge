package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
	"github.com/jmoiron/sqlx"

	"errors"
	"log"
	"strings"
	"time"
)

func (this *AdminHandler) SaveContest(response *api.SaveContestResponse, req *api.SaveContestRequest) {
	defer PanicHandler(response, this.debug)
	SaveContest_BuildResponse(this.dbu, true, this.session.GetUserId(), response, req, this.debug)
}

func (this *BasicHandler) SaveContest(response *api.SaveContestResponse, req *api.SaveContestRequest) {
	log.Println("Get here")
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

func (this *UserHandler) SaveContest(response *api.SaveContestResponse, req *api.SaveContestRequest) {
	defer PanicHandler(response, this.debug)
	// check if able to upate
	if req.GetContestId() != 0 {
		tx := this.dbu.MustBegin()
		defer this.dbu.Rollback()
		cst, err := models.Query_Contest_By_ContestId(tx, req.GetContestId(), nil, nil)
		PanicOnError(err)
		if cst.CreatorId != this.session.GetUserId() {
			MakeResponseError(response, this.debug, PBUnauthorized, nil)
			return
		}
	}

	SaveContest_BuildResponse(this.dbu, false, this.session.GetUserId(), response, req, this.debug)
}

func SaveContest_BuildResponse(
	dbu *db.DBUtil,
	is_admin bool,
	user_id int64,
	response *api.SaveContestResponse,
	req *api.SaveContestRequest,
	debug bool,
) {
	if is_admin == false {
		// if is not admin, not allowed to create formal contest.
		if req.GetIsVirtual() == false {
			MakeResponseError(response, debug, PBUnauthorized,
				errors.New("only admin can create formal contest"))
		}
	}

	// check contest type
	switch strings.ToLower(req.GetContestType()) {
	case "oi", "icpc", "cf":
	default:
		MakeResponseError(response, debug, PBBadRequest, errors.New("invalid contest type"))
		return
	}

	// save contest details
	cst := &models.Contest{
		ContestId:        req.GetContestId(),
		Title:            req.GetTitle(),
		Description:      req.GetDescription(),
		IsVirtual:        req.GetIsVirtual(),
		ContestType:      req.GetContestType(),
		CreateTime:       time.Now(),
		StartTime:        base.GetDefaultTime(),
		EndTime:          base.GetDefaultTime(),
		LockBoardTime:    base.GetDefaultTime(),
		HideOthersStatus: req.GetHideOthersStatus(),
		IsHidden:         req.GetIsHidden(),
		Password:         req.GetPassword(),
	}

	cm := models.NewContestModel()
	tx := dbu.MustBegin()
	defer dbu.Rollback()
	// check if insert or update
	if req.GetContestId() == 0 {
		cst.CreatorId = user_id
		id, err := cm.Insert(tx, cst)
		if err != nil {
			PanicOnError(
				errors.New("failed to insert contest detail:" + err.Error()))
		}
		response.ContestId = id
	} else {
		if err := cm.Update(tx, "", cst, nil,
			[]string{"create_time", "creator_id"}); err != nil {
			PanicOnError(
				errors.New("failed to update contest detail:" + err.Error()))
		}
		response.ContestId = cst.ContestId
	}

	// check num of problems
	if len(req.GetProblems()) > 26 {
		MakeResponseError(response, debug, PBBadRequest,
			errors.New("too many problems"))
		return
	}

	// clear problems
	cpm := models.NewContestProblemModel()
	if err := cpm.DeleteProblemsByContestId(tx, response.ContestId); err != nil {
		PanicOnError(errors.New("failed to clear problems:" + err.Error()))
	}

	// insert problems
	cnt := 0
	for _, v := range req.GetProblems() {
		sid := v.GetProblemSid()
		meta_pid, err := Query_MetaPid_By_Sid(tx, sid)
		if err != nil {
			// Log
			log.Println("query meta_pid: " + err.Error())
		}

		cp := &models.ContestProblem{
			MetaPidFK:   meta_pid,
			ContestIdFK: response.ContestId,
			Alias:       v.GetAlias(),
			Label:       base.GenerateLabel(cnt),
		}
		cnt++

		// insert
		_, err = cpm.Insert(tx, cp)
		if err != nil {
			// Log
			log.Println("insert contest problem: " + err.Error())
		}
	}

	// Commit changes
	dbu.MustCommit()

}

func Query_MetaPid_By_Sid(tx *sqlx.Tx, sid string) (int64, error) {
	pid, err := base.ParseSid(sid)
	if err != nil {
		return 0, err
	}

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(
		tx, pid.OJName, pid.OJPid, []string{"meta_pid"}, nil)

	if err != nil {
		return 0, err
	}

	return mp.MetaPid, nil
}
