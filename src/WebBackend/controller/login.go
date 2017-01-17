package controller

import (
	"OnlineJudge/base"
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"
	locals "OnlineJudge/sessions"

	"net/http"
)

// TODO: Set CSRF Token here
func (this *Controller) LoginInit(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	var response = &api.LoginInitResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetLoginInitRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.LoginInit(response, request)

	webresponse.LoginInitResponse = response
}

func (this *Controller) LoginAuth(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	var response = &api.LoginAuthResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetLoginAuthRequest()

	ip_addr := base.GetIPAddress(r)
	locals.Session(session).SetIPAddr(ip_addr)

	handler := handler.NewHandler(session, this.debug)
	handler.LoginAuth(response, request)

	webresponse.LoginAuthResponse = response
}
