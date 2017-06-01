package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) SaveProblem(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebPostRequest{}

	var response = &api.SaveProblemResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetSaveProblemRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.SaveProblem(response, request)

	webresponse.SaveProblemResponse = response
}
