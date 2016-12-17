package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

// TODO: CAPCHA
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var response = &api.RegisterResponse{}
	var request = &api.RegisterRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.Register(response, request)
}
