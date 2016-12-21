package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"

	"net/http"
)

func (this *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	var webresponse = &api.WebResponse{}
	var webrequest = &api.WebRequest{}

	defer SetWebResponse(w, nil, webresponse)

	session, err := this.Prepare(webresponse, webrequest, w, r)
	if err != nil {
		return
	}
	defer session.Save(r, w)
	handler := handler.NewHandler(session, this.debug)
	handler.Logout()
}
