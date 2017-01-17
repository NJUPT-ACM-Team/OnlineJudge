package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

// TODO: CAPTCHA
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	var response = &api.RegisterResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	// Do some job to validate CAPTCHA

	request := webrequest.GetRegisterRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.Register(response, request)

	webresponse.RegisterResponse = response
}
