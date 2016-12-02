package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/handler/api"

	"net/http"
)

func (this *Controller) ShowProblem(w http.ResponseWriter, r *http.Request) {
	var response = &api.ShowProblemResponse{}
	var request = &api.ShowProblemRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.ShowProblem(response, request)
}
