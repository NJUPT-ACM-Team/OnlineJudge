package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	ListProblems_BuildResponse(
		this.tx, response, req, this.session.GetUsername(), true, this.debug)
}

func (this *BasicHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
}

func (this *UserHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	filter := req.GetFilter()
	if filter.GetPStatus() != 0 {
		if this.session.IsLogin() == false {
			MakeResponseError(response, this.debug, PBLoginRequired, nil)
			return
		}
	}

	ListProblems_BuildResponse(
		this.tx, response, req, this.session.GetUsername(), false, this.debug)
}

func ListProblems_BuildResponse(
	tx *sqlx.Tx,
	response *api.ListProblemsResponse,
	req *api.ListProblemsRequest,
	username string,
	show_hidden bool,
	debug bool) {

	filter := req.GetFilter()
	page, err := models.XQuery_List_Problems_With_Filter(
		tx,
		username,
		show_hidden,
		filter.GetOj(),
		int(filter.GetPStatus()),
		int(req.GetOrderBy()),
		req.GetIsDesc(),
		int(req.GetOffset()),
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		nil,
		nil,
	)
	if err != nil {
		MakeResponseError(response, debug, PBInternalError, err)
		return
	}

	// Build response
	lines := []*api.ListProblemsResponse_PerLine{}
	for _, problem := range page.Problems {
		line := &api.ListProblemsResponse_PerLine{
			Sid: base.GenSid(&base.Pid{
				OJName: problem.OJName,
				OJPid:  problem.OJPid}),
			Oj:              problem.OJName,
			Pid:             problem.OJPid,
			Title:           problem.Title,
			Source:          problem.Source,
			AcSubmission:    1,
			TotalSubmission: 2,
		}
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
