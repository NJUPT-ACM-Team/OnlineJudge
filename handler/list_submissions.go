package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *Handler) AdminListSubmissions(response *api.ListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	if !this.session.IsRoot() {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}

	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	ListSubmissions_BuildResponse(
		this.tx, response, req, this.session.GetUsername(), true, true, this.debug)
}

func (this *Handler) ListSubmissions(response *api.ListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	ListSubmissions_BuildResponse(
		this.tx, response, req, this.session.GetUsername(), false, false, this.debug)
}

func ListSubmissions_BuildResponse(
	tx *sqlx.Tx,
	response *api.ListSubmissionsResponse,
	req *api.ListSubmissionsRequest,
	username string,
	show_private bool,
	show_all_code bool,
	debug bool) {

	filter := req.GetFilter()
	page, err := models.XQuery_List_Submissions_With_Filter(
		tx,
		username,
		show_private,
		filter.GetUsername(),
		filter.GetOj(),
		filter.GetPid(),
		filter.GetStatusCode(),
		filter.GetLanguage(),
		filter.GetCompiler(),
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		nil,
		nil,
	)
	if err != nil {
		MakeResponseError(response, debug, PBInternalError, err)
		return
	}

	lines := []*api.ListSubmissionsResponse_PerLine{}
	for _, submission := range page.Submissions {
		line := &api.ListSubmissionsResponse_PerLine{
			Sid:        base.GenSid(&base.Pid{OJName: submission.OJName, OJPid: submission.OJPid}),
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
			SubmitTime:      submission.SubmitTime.String(),
			IsSpj:           submission.IsSpj,
			// Code,
		}
		if show_all_code || submission.IsShared {
			line.Code = submission.Code
		}
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
