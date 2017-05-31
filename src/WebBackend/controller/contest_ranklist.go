package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) ContestRanklist(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}

	var response = &api.ContestRanklistResponse{}
	var request = &api.ContestRanklistRequest{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.ContestRanklist(response, request)

	webresponse.ContestRanklistResponse = response
}
