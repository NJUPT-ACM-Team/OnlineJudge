package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.MustCommit()

	ShowProblem_BuildResponse(tx, response, req, true, this.debug)
}

func (this *BasicHandler) ShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.MustCommit()

	ShowProblem_BuildResponse(tx, response, req, false, this.debug)
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
	PanicOnError(err)

	// QueryLanguages
	langs, err := models.Query_Languages_By_OJIdFK(tx, mp.OJIdFK, nil, nil)
	PanicOnError(err)

	// Query Limits
	limits, err := models.Query_Limits_By_MetaPid(tx, mp.MetaPid, nil, nil)
	PanicOnError(err)
	r_limits := []*api.Problem_Limit{}
	for _, limit := range limits {
		temp := &api.Problem_Limit{
			Language:      limit.Language,
			TimeLimit:     int32(limit.TimeLimit),
			CaseTimeLimit: int32(limit.CaseTimeLimit),
			MemoryLimit:   int32(limit.MemoryLimit),
		}
		r_limits = append(r_limits, temp)
	}

	// Judge if authorized
	if mp.Hide == true && !access_hide {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// Make response
	response.ProblemSid = req.GetProblemSid()
	problem := &api.Problem{
		Title: mp.Title,
		// TimeLimit:     int32(mp.TimeLimit),
		// CaseTimeLimit: int32(mp.CaseTimeLimit),
		// MemoryLimit:   int32(mp.MemoryLimit),
		Limits:       r_limits,
		Description:  mp.Description,
		Input:        mp.Input,
		Output:       mp.Output,
		SampleInput:  mp.SampleIn,
		SampleOutput: mp.SampleOut,
		Source:       mp.Source,
		Hint:         mp.Hint,
		Hide:         mp.Hide,
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
