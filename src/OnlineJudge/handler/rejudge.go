package handler

import (
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *AdminHandler) ReJudge(response *api.ReJudgeResponse, req *api.ReJudgeRequest) {
	defer PanicHandler(response, this.debug)
	ReJudge_BuildResponse(this.dbu, response, true, req.GetRunId(), this.debug)
}

func (this *BasicHandler) ReJudge(response *api.ReJudgeResponse, req *api.ReJudgeRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

func (this *UserHandler) ReJudge(response *api.ReJudgeResponse, req *api.ReJudgeRequest) {
	defer PanicHandler(response, this.debug)
	ReJudge_BuildResponse(this.dbu, response, false, req.GetRunId(), this.debug)
}

func ReJudge_BuildResponse(
	dbu *db.DBUtil,
	response *api.ReJudgeResponse,
	is_admin bool,
	run_id int64,
	debug bool,
) {
	tx := dbu.MustBegin()
	defer dbu.Rollback()
	if is_admin {
		go CallJudging(dbu, run_id)
		response.Success = true
		return
	}
	sub, err := models.Query_Submission_By_RunId(tx, run_id, nil, nil)
	PanicOnError(err)
	if sub == nil {
		MakeResponseError(response, debug, PBSubmissionNotFound, nil)
		return
	}
	if sub.StatusCode != "se" {
		response.Success = false
		return
	}
	dbu.Rollback()
	go CallJudging(dbu, run_id)
	response.Success = true
}
