package websession

import (
	locals "OnlineJudge/session"
	"github.com/gorilla/sessions"
)

type WebSession struct {
	Values map[interface{}]interface{}
}

func (this *WebSession) IsLogin() bool {
	if _, ok := this.Values["username"]; ok {
		return true
	}
	return false
}

func (this *WebSession) GetUserName() string {
	username, _ := this.Values["username"].(string)
	return username
}

func (this *WebSession) GetPrivilege() string {
	privilege, _ := this.Values["privilege"].(string)
	return privilege
}

func (this *WebSession) GetUserId() int64 {
	userid, _ := this.Values["user_id"].(int64)
	return userid
}

func NewWebSession(sess *sessions.Session) locals.Session {
	newsess := &WebSession{
		Values: sess.Values,
	}
	var session locals.Session = newsess
	return session
}
