package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) ContestShowProblem(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}

	var response = &api.ShowProblemResponse{}
	var request = &api.ShowProblemRequest{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.ContestShowProblem(response, request)

	webresponse.ShowProblemResponse = response
}
