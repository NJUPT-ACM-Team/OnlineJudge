package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"
)

func (this *BasicHandler) About(response *api.AboutResponse, req *api.AboutRequest) {
	if err := this.OpenDB(); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}
	defer this.CloseDB()

	if req.GetNeedOjsList() == true {
		var ojs []*api.OJInfo
		all, err := models.Query_All_OJs(this.tx, nil, nil)
		if err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
			return
		}
		for _, oj := range all {
			temp := &api.OJInfo{
				OjId:       oj.OJId,
				OjName:     oj.OJName,
				Version:    oj.Version,
				Int64Io:    oj.Int64IO,
				Javaclass:  oj.JavaClass,
				Status:     oj.Status,
				StatusInfo: oj.StatusInfo,
				Lastcheck:  oj.LastCheck.String(),
			}
			ojs = append(ojs, temp)
		}
		response.OjsList = ojs
	}

	if req.GetNeedLanguagesList() == true {
		var languages []*api.Language
		all, err := models.Query_All_Languages(this.tx, nil, nil)
		if err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err)
			return
		}
		for _, lang := range all {
			temp := &api.Language{
				Compiler:   lang.Language.Compiler,
				Language:   lang.Language.Language,
				LanguageId: lang.Language.LangId,
				OjName:     lang.OJName,
			}
			languages = append(languages, temp)
		}
		response.LanguagesList = languages
	}
}
