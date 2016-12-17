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

	//	filter := req.GetFilter()
}
