package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *Handler) ListSubmissions(response *api.ListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	if req.GetNeedLanguagesList() == true {
		var languages []*api.Language
		all, err := models.Query_All_Languages(this.tx, nil, nil)
		if err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
			return
		}
		for _, lang := range all {
			temp := &api.Language{
				Compiler:   lang.Compiler,
				Language:   lang.Language,
				LanguageId: lang.LangId,
				OjName:     lang.OJName,
			}
			languages = append(languages, temp)
		}
		response.LanguagesList = languages
	}

	if req.GetNeedOjsList() == true {
		ojs, err := models.Query_All_OJNames(this.tx)
		if err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
			return
		}
		response.OjsList = ojs
	}

	filter := req.GetFilter()
	var show_private bool
	if this.session.GetPrivilege() == "root" {
		show_private = true
	} else {
		show_private = false
	}
	page, err := models.XQuery_List_Submissions_With_Filter(
		this.tx,
		filter.GetUsername(),
		show_private,
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
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	lines := []*api.ListSubmissionsResponse_PerLine{}
	for _, submission := range page.Submissions {
		line := &api.ListSubmissionsResponse_PerLine{
			Sid:        "",
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
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
