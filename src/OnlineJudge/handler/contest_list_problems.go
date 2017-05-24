package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
}

func (this *BasicHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	// Check Access
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, true, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Problems {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	// list problems
	ContestListProblems_BuildResponse(tx, response, req, false, req.GetContestId(), this.session.GetUserId())
}

func (this *UserHandler) ContestListProblems(response *api.ContestListProblemsResponse, req *api.ContestListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	// Check Access
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Problems {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}

	// list problems
	show_details := false
	if access.Creator {
		show_details = true
	}
	ContestListProblems_BuildResponse(tx, response, req, show_details, req.GetContestId(), this.session.GetUserId())
}

func CheckContestUser(tx *sqlx.Tx, contest_id, user_id int64) (bool, error) {
	cst, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	if err != nil {
		return false, err
	}
	if cst == nil {
		return false, nil
	}
	return true, nil
}

func ContestListProblems_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestListProblemsResponse,
	req *api.ContestListProblemsRequest,
	show_details bool,
	contest_id int64,
	user_id int64) {

	cps, err := models.XQuery_Contest_List_Problems(
		tx, contest_id, user_id)
	PanicOnError(err)

	lines := []*api.ContestListProblemsResponse_PerLine{}
	for _, cp := range cps {
		// TODO: add ac count
		line := &api.ContestListProblemsResponse_PerLine{
			Label:  cp.Label,
			Alias:  cp.Alias,
			Status: cp.Status,
		}
		if show_details {
			line.Title = cp.Title
			line.Sid = base.GenSid(&base.Pid{OJName: cp.OJName, OJPid: cp.OJPid})
		}
		lines = append(lines, line)
	}
	response.Lines = lines
}

type ContestAccess struct {
	If404    bool
	Problems bool
	Status   bool
	Rank     bool
	Submit   bool
	Creator  bool
}

func CheckContestAccess(
	tx *sqlx.Tx, is_guest bool,
	contest_id int64, user_id int64, debug bool) (*ContestAccess, error) {

	cst, err := models.Query_Contest_By_ContestId(tx, contest_id, nil, nil)
	if err != nil {
		return nil, err
	}
	if cst == nil {
		return &ContestAccess{If404: true}, nil
	}

	if is_guest {
		if cst.IsPrivate() {
			return &ContestAccess{}, nil
		} else {
			return &ContestAccess{
				Problems: true,
				Status:   true,
				Rank:     true,
			}, nil
		}
	}

	if cst.CreatorId == user_id {
		return &ContestAccess{
			Problems: true,
			Status:   true,
			Rank:     true,
			Submit:   true,
			Creator:  true,
		}, nil
	}

	cu, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	if err != nil {
		return nil, err
	}

	if cu == nil {
		if cst.IsPrivate() {
			return &ContestAccess{}, nil
		} else if cst.IsProtected() {
			return &ContestAccess{
				Problems: true,
				Status:   true,
				Rank:     true,
			}, nil
		} else if cst.IsPublic() {
			return &ContestAccess{
				Problems: true,
				Status:   true,
				Rank:     true,
				Submit:   true,
			}, nil
		}
	}
	return &ContestAccess{
		Problems: true,
		Status:   true,
		Rank:     true,
		Submit:   true,
	}, nil
}
