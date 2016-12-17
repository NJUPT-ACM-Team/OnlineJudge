package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *Handler) ListProblems(response *api.ListProblemsResponse, req *api.ListProblemsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Format filter information

	var show_hidden bool
	if this.session.GetPrivilege() == "root" {
		show_hidden = true
	} else {
		show_hidden = false
	}

	filter := req.GetFilter()
	if filter.GetPStatus() != 0 {
		if this.session.IsLogin() == false {
			MakeResponseError(response, this.debug, PBLoginRequired, nil)
			return
		}
	}
	page, err := models.XQuery_List_Problems_With_Filter(
		this.tx,
		this.session.GetUsername(),
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
		MakeResponseError(response, this.debug, PBInternalError, err)
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
