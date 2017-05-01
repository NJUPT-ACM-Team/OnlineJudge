package handler

/*
import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *UserHandler) EditProblem(response *api.EditProblemResponse, req *api.EditProblemRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Judge if authorized
	if this.session.IsLogin() == false {
		MakeResponseError(response, this.debug, PBLoginRequired, nil)
		return
	} else {
		if this.session.GetPrivilege() != "root" {
			MakeResponseError(response, this.debug, PBUnauthorized, nil)
		}
	}

	// If Sid is empty, means this is a add problem request
	if req.Sid == "" {
		return
	}

	// else, we return the current problem information
	pid, err := base.ParseSid(req.GetSid())
	if err != nil {
		MakeResponseError(response, this.debug, PBBadRequest, err)
		return
	}

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, pid.OJName, pid.OJPid, nil, nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	pr := &api.Problem{
		Title:         mp.Title,
		TimeLimit:     int32(mp.TimeLimit),
		CaseTimeLimit: int32(mp.CaseTimeLimit),
		MemoryLimit:   int32(mp.MemoryLimit),
		Description:   mp.Description,
		Input:         mp.Input,
		Output:        mp.Output,
		SampleInput:   mp.SampleIn,
		SampleOutput:  mp.SampleOut,
		Source:        mp.Source,
		Hint:          mp.Hint,
		Hide:          mp.Hide,
	}
	response.Problem = pr
	response.Sid = req.GetSid()
}
*/
