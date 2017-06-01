package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) SaveProblem(response *api.SaveProblemResponse, req *api.SaveProblemRequest) {
	defer PanicHandler(response, this.debug)
	SaveProblem_BuildResponse(this.dbu, response, req, this.debug)
}

func (this *BasicHandler) SaveProblem(response *api.SaveProblemResponse, req *api.SaveProblemRequest) {
	MakeResponseError(response, this.debug, PBUnauthorized, nil)
}

func SaveProblem_BuildResponse(
	dbu *db.DBUtil,
	response *api.SaveProblemResponse,
	req *api.SaveProblemRequest,
	debug bool,
) {
	tx := dbu.MustBegin()
	defer dbu.Rollback()

	oj_name := req.GetOjName()
	oj_pid := req.GetOjPid()
	// check oj
	oj, err := models.Query_OJ_By_OJName(tx, oj_name, nil, nil)
	PanicOnError(err)
	if oj == nil {
		MakeResponseError(response, debug, PBOJNotFound, nil)
		return
	}

	nmp := &models.MetaProblem{
		Title:       req.GetTitle(),
		Description: req.GetDescription(),
		Input:       req.GetInput(),
		Output:      req.GetOutput(),
		SampleIn:    req.GetSampleIn(),
		SampleOut:   req.GetSampleOut(),
		Source:      req.GetSource(),
		Hint:        req.GetHint(),
		Hide:        req.GetHide(),
		IsSpj:       req.GetIsSpj(),
		SpjCode:     req.GetSpjCode(),
		OJName:      oj_name,
		OJPid:       oj_pid,
		OJIdFK:      oj.OJId,
	}

	mpm := models.NewMetaProblemModel()
	tmlm := models.NewTimeMemoryLimitModel()
	var meta_pid int64

	mp, err := models.Query_MetaProblem_By_OJName_OJPid(
		tx, oj_name, oj_pid, nil, nil)
	PanicOnError(err)
	if mp == nil {
		// insert
		meta_pid, err = mpm.Insert(tx, nmp)
		PanicOnError(err)
		err = InsertLimits(tx, meta_pid, req.GetLimits(), tmlm)
		PanicOnError(err)
		dbu.MustCommit()
		response.MetaPid = meta_pid
		response.ProblemSid = base.GenSid(&base.Pid{OJName: oj_name, OJPid: oj_pid})
	} else {
		// update
		// TODO: limits
		// delete all limits
		err = tmlm.DeleteLimitsByMetaPid(tx, mp.MetaPid)
		PanicOnError(err)
		err = InsertLimits(tx, mp.MetaPid, req.GetLimits(), tmlm)
		PanicOnError(err)

		nmp.MetaPid = mp.MetaPid
		err = mpm.Update(tx, "", nmp, nil, nil)
		PanicOnError(err)
		dbu.MustCommit()
		response.MetaPid = nmp.MetaPid
		response.ProblemSid = base.GenSid(&base.Pid{OJName: oj_name, OJPid: oj_pid})
	}
}

func InsertLimits(
	tx *sqlx.Tx,
	meta_pid int64,
	limits []*api.Problem_Limit,
	tmlm *models.TimeMemoryLimitModel) error {

	for _, limit := range limits {
		tm := &models.TimeMemoryLimit{
			TimeLimit:   int(limit.GetTimeLimit()),
			MemoryLimit: int(limit.GetMemoryLimit()),
			// CaseTimeLimit:int(limit.GetTimeLim)
			Language:  limit.GetLanguage(),
			MetaPidFK: meta_pid,
		}
		if _, err := tmlm.Insert(tx, tm); err != nil {
			return err
		}
	}
	return nil
}
