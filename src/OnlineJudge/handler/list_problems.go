package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *AdminHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()

	ListProblems_BuildResponse(
		tx, response, req, this.session.GetUsername(), true, this.debug)
}

func (this *BasicHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()

	// filter := req.GetFilter()
	if req.GetFilterPStatus() != 0 {
		// if filter.GetPStatus() != 0 {
		MakeResponseError(response, this.debug, PBLoginRequired, nil)
		return
	}
	ListProblems_BuildResponse(
		tx, response, req, "", false, this.debug)
}

func (this *UserHandler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()

	ListProblems_BuildResponse(
		tx, response, req, this.session.GetUsername(), false, this.debug)
}

func ListProblems_BuildResponse(
	tx *sqlx.Tx,
	response *api.ListProblemsResponse,
	req *api.ListProblemsRequest,
	username string,
	show_hidden bool,
	debug bool) {

	// filter := req.GetFilter()
	page, err := models.XQuery_List_Problems_With_Filter(
		tx,
		username,
		show_hidden,
		// filter.GetOj(),
		req.GetFilterOj(),
		// filter.GetPStatus().String(),
		req.GetFilterPStatus().String(),
		req.GetOrderBy().String(),
		req.GetIsDesc(),
		int(req.GetOffset()),
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		nil,
		nil,
	)
	PanicOnError(err)

	// Build response
	lines := []*api.ListProblemsResponse_PerLine{}
	for _, problem := range page.Problems {
		total, err := models.Query_Total_Submissions_By_MetaPid(tx, problem.MetaPid)
		if err != nil {
			total = 0
		}
		ac, err := models.Query_Total_AC_Submissions_By_MetaPid(tx, problem.MetaPid)
		if err != nil {
			ac = 0
		}
		line := &api.ListProblemsResponse_PerLine{
			Sid: base.GenSid(&base.Pid{
				OJName: problem.OJName,
				OJPid:  problem.OJPid}),
			Oj:              problem.OJName,
			Pid:             problem.OJPid,
			Title:           problem.Title,
			Source:          problem.Source,
			AcSubmission:    int32(ac),
			TotalSubmission: int32(total),
		}
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
