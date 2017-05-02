package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ListSubmissions(response *api.ListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	this.OpenDBU()
	defer this.CloseDBU()
	tx := this.dbu.MustBegin()

	ListSubmissions_BuildResponse(
		tx, response, req, this.session.GetUsername(), true, true, this.debug)
}

func (this *BasicHandler) ListSubmissions(response *api.ListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	this.OpenDBU()
	defer this.CloseDBU()
	tx := this.dbu.MustBegin()

	ListSubmissions_BuildResponse(
		tx, response, req, this.session.GetUsername(), false, false, this.debug)
}

func ListSubmissions_BuildResponse(
	tx *sqlx.Tx,
	response *api.ListSubmissionsResponse,
	req *api.ListSubmissionsRequest,
	username string,
	show_private bool,
	show_all_code bool,
	debug bool) {

	// filter := req.GetFilter()
	page, err := models.XQuery_List_Submissions_With_Filter(
		tx,
		username,
		show_private,
		req.GetIsDesc(),
		// filter.GetUsername(),
		req.GetFilterUsername(),
		// filter.GetOj(),
		req.GetFilterOj(),
		// filter.GetPid(),
		req.GetFilterPid(),
		// filter.GetStatusCode(),
		req.GetFilterStatusCode(),
		// filter.GetLanguage(),
		req.GetFilterLanguage(),
		// filter.GetCompiler(),
		req.GetFilterCompiler(),
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		nil,
		nil,
	)
	PanicOnError(err)

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
