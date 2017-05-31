package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"

	"errors"
	"fmt"
)

func (this *AdminHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ContestListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

}

func (this *BasicHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ContestListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, true, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Status {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	if access.Time < 0 && !access.Creator {
		MakeResponseError(response, this.debug, PBUnauthorized, errors.New("contest not started"))
		return
	}
	ContestListSubmissions_BuildResponse(tx, response, req, this.session.GetUsername(), this.debug)
}

func (this *UserHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ContestListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Status {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	if access.Time < 0 && !access.Creator {
		MakeResponseError(response, this.debug, PBUnauthorized, errors.New("contest not started"))
		return
	}
	ContestListSubmissions_BuildResponse(tx, response, req, this.session.GetUsername(), this.debug)

}

func ContestListSubmissions_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestListSubmissionsResponse,
	req *api.ContestListSubmissionsRequest,
	username string,
	debug bool,
) {
	contest_id := req.GetContestId()
	cst, err := models.Query_Contest_By_ContestId(tx, contest_id, nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, debug, PBContestNotFound, nil)
		return
	}
	filter_username := req.GetFilterUsername()
	if cst.HideOthersStatus {
		if username != "" {
			filter_username = username
		} else {
			filter_username = "-1"
		}

	}
	page, err := models.XQuery_Contest_List_Submissions_With_Filter(
		tx,
		username,
		false,
		contest_id,
		req.GetIsDesc(),
		req.GetFilterRunId(),
		filter_username,
		req.GetFilterLabel(),
		req.GetFilterStatusCode(),
		req.GetFilterLanguage(),
		req.GetFilterCompiler(),
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		nil,
		nil,
	)
	PanicOnError(err)
	lines := []*api.ContestListSubmissionsResponse_PerLine{}
	for _, submission := range page.Submissions {
		fmt.Println(submission)
		line := &api.ContestListSubmissionsResponse_PerLine{
			Label:      submission.Label,
			RunId:      submission.RunId,
			Username:   submission.Username,
			Status:     submission.Status,
			StatusCode: submission.StatusCode,
			CeInfo:     submission.CEInfo,
			Language: &api.Language{
				Language: submission.Language.Language,
				Compiler: submission.Language.Compiler,
			},
			TimeUsed:        int32(submission.TimeUsed),
			MemoryUsed:      int32(submission.MemoryUsed),
			Testcases:       int32(submission.NumberOfTestcases),
			TestcasesPassed: int32(submission.TestcasesPassed),
			CodeLength:      int32(len([]byte(submission.Code))),
			SubmitTime:      base.MarshalTime(submission.SubmitTime),
			IsSpj:           submission.IsSpj,
		}
		if submission.Username == username {
			line.Code = submission.Code
		}
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
