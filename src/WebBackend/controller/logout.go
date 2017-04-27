package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebPostRequest{}

	var response = &api.LogoutResponse{}

	defer SetWebResponse(w, response, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)

	request := webrequest.LogoutRequest

	handler := handler.NewHandler(session, this.debug)
	handler.Logout(response, request)

	webresponse.LogoutResponse = response
}
