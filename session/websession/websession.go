package websession

import (
	"OnlineJudge/base"
	locals "OnlineJudge/session"
	"errors"
	"github.com/gorilla/sessions"
)

var (
	KeyUserName  = "username"
	KeyPrivilege = "privilege"
	KeyUserId    = "user_id"

	KEYS = []string{KeyUserName, KeyPrivilege, KeyUserId}

	ErrNotFound   = errors.New("Key not in the session")
	ErrKeyInvalid = errors.New("Key invalid")
)

type WebSession struct {
	Values map[interface{}]interface{}
}

func (this *WebSession) IsLogin() bool {
	if _, ok := this.Values[KeyUserName].(string); ok {
		return true
	}
	return false
}

func (this *WebSession) GetUserName() (string, error) {
	username, ok := this.Values[KeyUserName].(string)
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
		if base.ArrayContains(KEYS, val) {
			return ErrKeyInvalid
		}
	}
	this.Values[key] = val
	return nil
}

func (this *WebSession) SetUserName(username string) {
	this.Values[KeyUserName] = username
}

func (this *WebSession) SetPrivilege(privilege string) {
	this.Values[KeyPrivilege] = privilege
}

func (this *WebSession) SetUserId(user_id int64) {
	this.Values[KeyUserId] = user_id
}

func NewWebSession(sess *sessions.Session) locals.Session {
	newsess := &WebSession{
		Values: sess.Values,
	}
	var session locals.Session = newsess
	return session
}
