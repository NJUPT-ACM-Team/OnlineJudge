package controller

import (
	"OnlineJudge/base"
	"OnlineJudge/handler"
	"OnlineJudge/handler/api"
	locals "OnlineJudge/sessions"

	"net/http"
)

func (this *Controller) Submit(w http.ResponseWriter, r *http.Request) {
	var response = &api.SubmitResponse{}
	var request = &api.SubmitRequest{}
	defer SetResponse(w, response)

	session, err := this.Prepare(response, request, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)
	ip_addr := base.GetIPAddress(r)
	locals.Session(session).SetIPAddr(ip_addr)

	handler := handler.NewHandler(session, this.debug)
	handler.Submit(response, request)
}
