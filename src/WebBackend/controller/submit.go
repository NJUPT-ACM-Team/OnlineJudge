package controller

import (
	"OnlineJudge/base"
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"
	locals "OnlineJudge/sessions"

	"net/http"
)

func (this *Controller) Submit(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebPostRequest{}

	var response = &api.SubmitResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetSubmitRequest()

	ip_addr := base.GetIPAddress(r)
	locals.Session(session).SetIPAddr(ip_addr)

	handler := handler.NewHandler(session, this.debug)
	handler.Submit(response, request)

	webresponse.SubmitResponse = response
}
