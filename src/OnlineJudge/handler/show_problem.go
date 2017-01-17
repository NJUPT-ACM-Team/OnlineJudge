package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	ShowProblem_BuildResponse(this.tx, response, req, true, this.debug)
}

func (this *BasicHandler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
}

func (this *UserHandler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()
	ShowProblem_BuildResponse(this.tx, response, req, false, this.debug)
}

func ShowProblem_BuildResponse(
	tx *sqlx.Tx,
	response *api.ShowProblemResponse,
	req *api.ShowProblemRequest,
	access_hide bool,
	debug bool) {

	// Get Sid
	pid, err := base.ParseSid(req.GetProblemSid())
	if err != nil {
		MakeResponseError(response, debug, PBBadRequest, err)
		return
	}

	// Query problem
	mp, err := models.Query_MetaProblem_By_OJName_OJPid(tx, pid.OJName, pid.OJPid, nil, nil)
	if err != nil {
		MakeResponseError(response, debug, PBInternalError, err)
		return
	}

	// QueryLanguages
	langs, err := models.Query_Languages_By_OJIdFK(tx, mp.OJIdFK, nil, nil)
	if err != nil {
		MakeResponseError(response, debug, PBInternalError, err)
		return
	}

	// Judge if authorized
	if mp.Hide == true && !access_hide {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
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
