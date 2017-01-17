package websession

import (
	"github.com/gorilla/sessions"

	"net/http"
)

// Copy from github.com/gorilla/sessions/store.go
type Store interface {
	// Get should return a cached session.
	Get(r *http.Request, name string) (*sessions.Session, error)

	// New should create and return a new session.
	//
	// Note that New should never return a nil session, even in the case of
	// an error if using the Registry infrastructure to cache the session.
	New(r *http.Request, name string) (*sessions.Session, error)

	// Save should persist session to the underlying store implementation.
	Save(r *http.Request, w http.ResponseWriter, s *sessions.Session) error
}

func NewStore() Store {
	return NewCookieStore()
}

func NewCookieStore() Store {
	var cookiestore = sessions.NewCookieStore([]byte("something-very-secret"))
	return cookiestore
}
