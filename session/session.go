package session

type Session interface {
	IsLogin() bool
	GetUserName() string
	GetUserId() int64
	GetPrivilege() string
}
