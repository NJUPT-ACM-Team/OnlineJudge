package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
)

func (this *Handler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Get Sid
	pid, err := base.ParseSid(req.GetProblemSid())
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBBadRequest, err)
		return
	}

	// Query problem
	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, pid.OJName, pid.OJPid, nil, nil)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}

	// QueryLanguages
	langs, err := models.Query_Languages_By_OJIdFK(this.tx, mp.OJIdFK, nil, nil)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}

	// Judge if authorized
	if mp.Hide != 0 {
		if this.session.IsLogin() == false {
			api.MakeResponseError(response, this.debug, api.PBProblemNotFound, err)
			return
		} else {
			if this.session.GetPrivilege() != "root" {
				api.MakeResponseError(response, this.debug, api.PBProblemNotFound, err)
				return
			}
		}
	}

	// Make response
	response.ProblemSid = req.GetProblemSid()
	response.Title = mp.Title
	response.TimeLimit = int32(mp.TimeLimit)
	response.CaseTimeLimit = int32(mp.CaseTimeLimit)
	response.MemoryLimit = int32(mp.MemoryLimit)
	response.Input = mp.Input
	response.Output = mp.Output
	response.SampleInput = mp.SampleIn
	response.SampleOutput = mp.SampleOut
	response.Source = mp.Source
	response.Hint = mp.Hint
	if mp.Hide != 0 {
		response.Hide = true
	} else {
		response.Hide = false
	}

	languages := []*api.Language{}
	for _, lang := range langs {
		temp := &api.Language{
			Compiler:    lang.Compiler,
			Language:    lang.Language,
			SubmitValue: int32(lang.LangId),
		}
		languages = append(languages, temp)
	}
	response.Languages = languages

}
