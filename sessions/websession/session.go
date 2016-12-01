package websession

import (
	"OnlineJudge/base"
	locals "OnlineJudge/sessions"

	"github.com/gorilla/sessions"

	"errors"
)

var (
	KeyUsername  = ".username"
	KeyPrivilege = ".privilege"
	KeyUserId    = ".user_id"
	KeyFlash     = ".flash"

	ErrNotFound   = errors.New("Key not in the session")
	ErrKeyInvalid = errors.New("Key invalid")
)

type WebSession struct {
	Values map[interface{}]interface{}
}

func (this *WebSession) IsLogin() bool {
	if _, ok := this.Values[KeyUsername].(string); ok {
		return true
	}
	return false
}

func (this *WebSession) GetUsername() (string, error) {
	username, ok := this.Values[KeyUsername].(string)
	if !ok {
		return "", ErrNotFound
	}
	return username, nil
}

func (this *WebSession) GetPrivilege() (string, error) {
	privilege, ok := this.Values[KeyPrivilege].(string)
	if !ok {
		return "", ErrNotFound
	}
	return privilege, nil
}

func (this *WebSession) GetUserId() (int64, error) {
	userid, ok := this.Values[KeyUserId].(int64)
	if !ok {
		return 0, ErrNotFound
	}
	return userid, nil
}

func (this *WebSession) Get(key interface{}) (interface{}, error) {
	if _, ok := this.Values[key]; !ok {
		return nil, ErrNotFound
	}
	return this.Values[key], nil
}

func (this *WebSession) Set(key interface{}, val interface{}) error {
	if val, ok := key.(string); ok {
		if !base.IsNilOrZero(val) {
			if val[0] == '.' {
				return ErrKeyInvalid
			}
		}
	}
	this.Values[key] = val
	return nil
}

func (this *WebSession) SetUsername(username string) {
	this.Values[KeyUsername] = username
}

func (this *WebSession) SetPrivilege(privilege string) {
	this.Values[KeyPrivilege] = privilege
}

func (this *WebSession) SetUserId(user_id int64) {
	this.Values[KeyUserId] = user_id
}

func (this *WebSession) Flashes(vars ...string) []interface{} {
	var flashes []interface{}
	key := KeyFlash
	if len(vars) > 0 {
		key = vars[0]
	}
	if v, ok := this.Values[key]; ok {
		// Drop the flashes and return it.
		delete(this.Values, key)
		flashes = v.([]interface{})
	}
	return flashes
}

// AddFlash adds a flash message to the session.
//
// A single variadic argument is accepted, and it is optional: it defines
// the flash key. If not defined "_flash" is used by default.
func (this *WebSession) AddFlash(value interface{}, vars ...string) {
	key := KeyFlash
	if len(vars) > 0 {
		key = vars[0]
	}
	var flashes []interface{}
	if v, ok := this.Values[key]; ok {
		flashes = v.([]interface{})
	}
	this.Values[key] = append(flashes, value)
}

func NewSession(sess *sessions.Session) locals.Session {
	newsess := &WebSession{
		Values: sess.Values,
	}
	var session locals.Session = newsess
	return session
}
