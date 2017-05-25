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

func (this *AdminHandler) ContestSave(response *api.ContestSaveResponse, req *api.ContestSaveRequest) {
	defer PanicHandler(response, this.debug)
	ContestSave_BuildResponse(this.dbu, true, this.session.GetUserId(), response, req, this.debug)
}

func (this *BasicHandler) ContestSave(response *api.ContestSaveResponse, req *api.ContestSaveRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

func (this *UserHandler) ContestSave(response *api.ContestSaveResponse, req *api.ContestSaveRequest) {
	defer PanicHandler(response, this.debug)
	// check if able to upate
	if req.GetContestId() != 0 {
		tx := this.dbu.MustBegin()
		defer this.dbu.Rollback()
		access, err := CheckContestAccess(
			tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
		PanicOnError(err)
		if access.If404 {
			MakeResponseError(response, this.debug, PBContestNotFound, nil)
			return
		}
		if !access.Creator {
			MakeResponseError(response, this.debug, PBUnauthorized, errors.New("not the creator"))
			return
		}
	}

	ContestSave_BuildResponse(this.dbu, false, this.session.GetUserId(), response, req, this.debug)
}

func ContestSave_BuildResponse(
	dbu *db.DBUtil,
	is_admin bool,
	user_id int64,
	response *api.ContestSaveResponse,
	req *api.ContestSaveRequest,
	debug bool,
) {
	if is_admin == false {
		// if is not admin, not allowed to create formal contest.
		if req.GetIsVirtual() == false {
			MakeResponseError(response, debug, PBUnauthorized,
				errors.New("only admin can create formal contest"))
			return
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
		CreateTime:       time.Now().UTC(),
		LockBoardTime:    base.GetDefaultTime(),
		HideOthersStatus: req.GetHideOthersStatus(),
		IsHidden:         req.GetIsHidden(),
		Password:         req.GetPassword(),
	}
	startTime, err := base.UnmarshalTime(req.GetStartTime())
	if err != nil {
		MakeResponseError(response, debug, PBBadRequest, errors.New("time format wrong"))
		return
	}
	endTime, err := base.UnmarshalTime(req.GetEndTime())
	if err != nil {
		MakeResponseError(response, debug, PBBadRequest, errors.New("time format wrong"))
		return
	}
	cst.StartTime = startTime
	cst.EndTime = endTime

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
	/*
		if err := cpm.DeleteProblemsByContestId(tx, response.ContestId); err != nil {
			PanicOnError(errors.New("failed to clear problems:" + err.Error()))
		}
	*/

	// insert problems
	cnt := 0
	for _, v := range req.GetProblems() {
		sid := v.GetProblemSid()
		meta_pid, err := Query_MetaPid_By_Sid(tx, sid)
		if err != nil {
			// Log
			log.Println("query meta_pid: " + err.Error())
		}
		if meta_pid == 0 {
			log.Println("problem not exist")
		}

		cp := &models.ContestProblem{
			MetaPidFK:   meta_pid,
			ContestIdFK: response.ContestId,
			Alias:       v.GetAlias(),
			Label:       base.GenerateLabel(cnt),
		}
		cnt++

		// insert
		check, err := Check_Del_ContestProblem(tx, cp)
		if err != nil {
			log.Println("check del contest problem: " + err.Error())
			continue
		}
		if check {
			_, err = cpm.Insert(tx, cp)
			if err != nil {
				// Log
				log.Println("insert contest problem: " + err.Error())
			}
		} else {
			if err := cpm.Update(tx, cp, "", []string{"alias"}, nil); err != nil {
				// Log
				log.Println("update contest problem: " + err.Error())
			}
		}
	}
	for i := cnt; i < 26; i++ {
		label := base.GenerateLabel(cnt)
		Check_Del_ContestProblem(tx,
			&models.ContestProblem{ContestIdFK: response.ContestId, Label: label})
	}

	// Commit changes
	dbu.MustCommit()

}

// if true insert, else update
func Check_Del_ContestProblem(tx *sqlx.Tx, cp *models.ContestProblem) (bool, error) {
	ocp, err := models.Query_ContestProblem_By_ContestId_And_Label(
		tx, cp.ContestIdFK, cp.Label)
	if err != nil {
		return false, err
	}
	if ocp == nil {
		return true, nil
	}
	if ocp.MetaPidFK == cp.MetaPidFK {
		cp.CPId = ocp.CPId
		return false, nil
	}

	// delete current problem
	cpm := models.NewContestProblemModel()
	if err := cpm.DeleteByContestIdAndLabel(
		tx, ocp.ContestIdFK, ocp.Label); err != nil {
		return false, err
	}
	return true, nil
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

	if mp == nil {
		return 0, nil
	}

	return mp.MetaPid, nil
}
