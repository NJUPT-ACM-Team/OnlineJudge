package controller

import (
	"OnlineJudge/captcha"
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

var (
	PBInvalidCAPTCHA = handler.NewPBError(http.StatusInternalServerError, "right captcha needed")
)

// TODO: CAPTCHA
func (this *Controller) Register(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebPostRequest{}

	var response = &api.RegisterResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		handler.MakeResponseError(response, this.debug, handler.PBInternalError, err)
		return
	}
	defer session.Save(r, w)

	// Do some job to validate CAPTCHA
	if !captcha.Verify(session, webrequest.GetCaptcha()) {
		handler.MakeResponseError(webresponse, this.debug, PBInvalidCAPTCHA, err)
		return
	}
	request := webrequest.GetRegisterRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.Register(response, request)

	webresponse.RegisterResponse = response
}
