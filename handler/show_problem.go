package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *Handler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Get Sid
	pid, err := base.ParseSid(req.GetProblemSid())
	if err != nil {
		MakeResponseError(response, this.debug, PBBadRequest, err)
		return
	}

	// Query problem
	mp, err := models.Query_MetaProblem_By_OJName_OJPid(this.tx, pid.OJName, pid.OJPid, nil, nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// QueryLanguages
	langs, err := models.Query_Languages_By_OJIdFK(this.tx, mp.OJIdFK, nil, nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// Judge if authorized
	if mp.Hide != false {
		if this.session.IsLogin() == false {
			MakeResponseError(response, this.debug, PBProblemNotFound, err)
			return
		} else {
			if this.session.GetPrivilege() != "root" {
				MakeResponseError(response, this.debug, PBProblemNotFound, err)
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
	response.Hide = mp.Hide

	languages := []*api.Language{}
	for _, lang := range langs {
		temp := &api.Language{
			Compiler:   lang.Compiler,
			Language:   lang.Language,
			LanguageId: int32(lang.LangId),
		}
		languages = append(languages, temp)
	}
	response.Languages = languages

}
