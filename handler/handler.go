package handler

import (
	"OnlineJudge/models/db"
	locals "OnlineJudge/session"
	"github.com/jmoiron/sqlx"
)

// For testing purpose
import (
	"OnlineJudge/session/websession"
	"github.com/gorilla/sessions"
	"net/http"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func NewHandlerForTest() (*Handler, *sessions.Session) {
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	session, _ := store.New(req, "my session")
	sess := websession.NewSession(session)
	return NewHandler(sess, true), session
}

//

type Handler struct {
	session locals.Session
	db      *sqlx.DB
	tx      *sqlx.Tx
	debug   bool
}

func NewHandler(sess locals.Session, dbg bool) *Handler {
	handler := &Handler{
		session: sess,
		debug:   dbg,
	}
	return handler
}

func (this *Handler) OpenDB() error {
	var err error
	this.db, err = db.NewDB()
	if err != nil {
		return err
	}
	this.tx, err = this.db.Beginx()
	if err != nil {
		return err
	}
	return nil
}

// Commit a transaction and start a new one
func (this *Handler) Commit() error {
	err := this.tx.Commit()
	if err != nil {
		return err
	}
	this.tx, err = this.db.Beginx()
	if err != nil {
		return err
	}
	return nil
}

func (this *Handler) CloseDB() {
	this.tx.Rollback()
	this.db.Close()
}
