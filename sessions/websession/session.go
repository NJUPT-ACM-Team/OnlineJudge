package websession

import (
	"OnlineJudge/base"
	locals "OnlineJudge/sessions"

	"github.com/gorilla/sessions"

	"errors"
	"net/http"
)

var (
	KeyUsername  = ".username"
	KeyPrivilege = ".privilege"
	KeyUserId    = ".user_id"
	KeyFlash     = ".flash"
	KeyIPAddr    = ".ipaddr"

	ErrNotFound   = errors.New("Key not in the session")
	ErrKeyInvalid = errors.New("Key invalid")
)

type WebSession struct {
	session *sessions.Session
}

func (this *WebSession) IsLogin() bool {
	if name, ok := this.session.Values[KeyUsername].(string); ok {
		if name != "" {
			return true
		}
	}
	return false
}

func (this *WebSession) IsRoot() bool {
	if privilege, ok := this.session.Values[KeyPrivilege].(string); ok {
		if privilege == "root" {
			return true
		}
	}
	return false
}

func (this *WebSession) GetUsername() string {
	username, ok := this.session.Values[KeyUsername].(string)
	if !ok {
		return ""
	}
	return username
}

func (this *WebSession) GetPrivilege() string {
	privilege, ok := this.session.Values[KeyPrivilege].(string)
	if !ok {
		return ""
	}
	return privilege
}

func (this *WebSession) GetUserId() int64 {
	userid, ok := this.session.Values[KeyUserId].(int64)
	if !ok {
		return 0
	}
	return userid
}

func (this *WebSession) GetIPAddr() string {
	ip, ok := this.session.Values[KeyIPAddr].(string)
	if !ok {
		return ""
	}
	return ip
}

func (this *WebSession) Get(key interface{}) (interface{}, error) {
	if _, ok := this.session.Values[key]; !ok {
		return nil, ErrNotFound
	}
	return this.session.Values[key], nil
}

func (this *WebSession) Set(key interface{}, val interface{}) error {
	if val, ok := key.(string); ok {
		if !base.IsNilOrZero(val) {
			if val[0] == '.' {
				return ErrKeyInvalid
			}
		}
	}
	this.session.Values[key] = val
	return nil
}

func (this *WebSession) SetUsername(username string) {
	this.session.Values[KeyUsername] = username
}

func (this *WebSession) SetPrivilege(privilege string) {
	this.session.Values[KeyPrivilege] = privilege
}

func (this *WebSession) SetUserId(user_id int64) {
	this.session.Values[KeyUserId] = user_id
}

func (this *WebSession) SetIPAddr(ip string) {
	this.session.Values[KeyIPAddr] = ip
}

func (this *WebSession) Flashes(vars ...string) []interface{} {
	return this.session.Flashes(vars...)
}

func (this *WebSession) AddFlash(value interface{}, vars ...string) {
	this.session.AddFlash(value, vars...)
}

func (this *WebSession) Save(h *http.Request, w http.ResponseWriter) {
	this.session.Save(h, w)
}

func (this *WebSession) Logout() {
	this.session.Options.MaxAge = -1
}

func NewSession(sess *sessions.Session) locals.Session {
	newsess := &WebSession{
		session: sess,
	}
	var session locals.Session = newsess
	return session
}

func NewWebSession(sess *sessions.Session) *WebSession {
	return &WebSession{
		session: sess,
	}
}
