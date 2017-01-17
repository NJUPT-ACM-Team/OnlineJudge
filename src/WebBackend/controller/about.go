package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) About(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	var response = &api.AboutResponse{}
	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.GetAboutRequest()

	handler := handler.NewHandler(session, this.debug)
	handler.About(response, request)

	webresponse.AboutResponse = response
}
