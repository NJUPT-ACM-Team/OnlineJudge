package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) ListContests(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}

	var response = &api.ListContestsResponse{}
	var request = &api.ListContestsRequest{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.ListContests(response, request)

	webresponse.ListContestsResponse = response
}
