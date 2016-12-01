package sessions

import ()

type Session interface {
	IsLogin() bool

	GetUsername() (string, error)
	GetUserId() (int64, error)
	GetPrivilege() (string, error)
	Get(key interface{}) (interface{}, error)

	SetUsername(username string)
	SetUserId(user_id int64)
	SetPrivilege(privilege string)
	Set(key interface{}, val interface{}) error

	Flashes(vars ...string) []interface{}
	AddFlash(value interface{}, vars ...string)
}
