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
		if this.session.IsLogin() == false || this.session.GetPrivilege() != "root" {
			MakeResponseError(response, this.debug, PBProblemNotFound, nil)
			return
		}
	}

	// Make response
	response.ProblemSid = req.GetProblemSid()
	problem := &api.Problem{
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
	response.Problem = problem
	languages := []*api.Language{}
	for _, lang := range langs {
		temp := &api.Language{
			Compiler:   lang.Compiler,
			Language:   lang.Language,
			LanguageId: lang.LangId,
		}
		languages = append(languages, temp)
	}
	response.Languages = languages

}
