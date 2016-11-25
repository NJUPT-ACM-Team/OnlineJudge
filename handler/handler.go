package handler

import (
	"OnlineJudge/models/db"
	"github.com/jmoiron/sqlx"
)

type Session struct {
	username  string
	user_id   int64
	privilege string
}

type Handler struct {
	session *Session
	db      *sqlx.DB
	tx      *sqlx.Tx
	debug   bool
}

func NewHandler(sess *Session, dbg bool) *Handler {
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
