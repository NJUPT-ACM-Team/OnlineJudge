package session

import ()

type Session interface {
	IsLogin() bool

	GetUserName() (string, error)
	GetUserId() (int64, error)
	GetPrivilege() (string, error)
	Get(key interface{}) (interface{}, error)

	SetUserName(username string)
	SetUserId(user_id int64)
	SetPrivilege(privilege string)
	Set(key interface{}, val interface{}) error
}
