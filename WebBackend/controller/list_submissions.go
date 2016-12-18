package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) ListSubmissions(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	var response = &api.ListSubmissionsResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetListSubmissionsRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.ListSubmissions(response, request)

	webresponse.ListSubmissionsResponse = response
}
