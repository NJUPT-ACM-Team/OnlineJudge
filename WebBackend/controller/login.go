package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/handler/api"

	"net/http"
)

func (this *Controller) LoginInit(w http.ResponseWriter, r *http.Request) {
	var response = &api.LoginInitResponse{}
	var request = &api.LoginInitRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.LoginInit(response, request)
}

func (this *Controller) LoginAuth(w http.ResponseWriter, r *http.Request) {
	var response = &api.LoginAuthResponse{}
	var request = &api.LoginAuthRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	handler := handler.NewHandler(session, this.debug)
	handler.LoginAuth(response, request)
}
