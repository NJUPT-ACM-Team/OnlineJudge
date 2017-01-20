package controller

import (
	"OnlineJudge/captcha"

	"net/http"
)

const (
	CAPTCHA = ".captcha"
)

func handleError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func (this *Controller) Captcha(w http.ResponseWriter, r *http.Request) {
	session, err := this.GetSession(r)
	if err != nil {
		handleError(w, err)
		return
	}
	ca := captcha.Prepare(session)
	// session.SetUsername("hello")
	// captcha.Verify(session, "")
	session.Save(r, w)
	captcha.WriteImage(w, ca, captcha.StdWidth, captcha.StdHeight)
}
