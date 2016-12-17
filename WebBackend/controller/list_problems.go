package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) ListProblems(w http.ResponseWriter, r *http.Request) {
	var response = &api.ListProblemsResponse{}
	var request = &api.ListProblemsRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.ListProblems(response, request)
}
