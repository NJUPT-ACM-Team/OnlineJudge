package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ContestShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
}

func (this *BasicHandler) ContestShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(tx, true, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Problems {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	ContestShowProblem_BuildResponse(tx, response, req, this.debug)
}

func (this *UserHandler) ContestShowProblem(response *api.ShowProblemResponse, req *api.ShowProblemRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Problems {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	ContestShowProblem_BuildResponse(tx, response, req, this.debug)
}

func ContestShowProblem_BuildResponse(
	tx *sqlx.Tx,
	response *api.ShowProblemResponse,
	req *api.ShowProblemRequest,
	debug bool) {

	contest_id := req.GetContestId()
	label := req.GetProblemSid()

	cp, err := models.Query_ContestProblem_By_ContestId_And_Label(tx, contest_id, label)
	PanicOnError(err)
	if cp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	mp, err := models.Query_MetaProblem_By_MetaPid(tx, cp.MetaPidFK, nil, nil)
	PanicOnError(err)
	if mp == nil {
		MakeResponseError(response, debug, PBProblemNotFound, nil)
		return
	}

	// QueryLanuages
	langs, err := models.Query_Languages_By_OJIdFK(tx, mp.OJIdFK, nil, nil)
	PanicOnError(err)

	// Query Limits
	limits, err := models.Query_Limits_By_MetaPid(tx, mp.MetaPid, nil, nil)
	PanicOnError(err)
	r_limits := []*api.Problem_Limit{}
	if limits != nil {
		for _, limit := range limits {
			temp := &api.Problem_Limit{
				Language:      limit.Language,
				TimeLimit:     int32(limit.TimeLimit),
				CaseTimeLimit: int32(limit.CaseTimeLimit),
				MemoryLimit:   int32(limit.MemoryLimit),
			}
			r_limits = append(r_limits, temp)
		}
	}

	// Make response
	response.ProblemSid = req.GetProblemSid()
	problem := &api.Problem{
		Title: cp.Alias,
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
