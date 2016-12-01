package controller

import (
	"OnlineJudge/handler"
	"OnlineJudge/handler/api"
	"OnlineJudge/sessions/websession"

	"io"
	"net/http"
)

// Get session
// Parse request
// Use handler to dispose
// Encode response
// Return
func (this *Controller) LoginInit(w http.ResponseWriter, r *http.Request) {

}

func (this *Controller) LoginAuth(w http.ResponseWriter, r *http.Request) {
	var response = &api.LoginAuthResponse{}
	var request = &api.LoginAuthRequest{}
	defer SetResponse(w, response)

	// Decode json to pb
	DecodePBFromJsonStream(io.LimitReader(r.Body, 1048576), request)

	// Get session
	session, err := this.store.Get(r, "default")
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	sess := websession.NewSession(session)
	// Handler
	handler := handler.NewHandler(sess, this.debug)
	handler.LoginAuth(response, request)
}
