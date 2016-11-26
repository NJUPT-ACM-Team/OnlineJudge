package handler

import (
	"OnlineJudge/models/db"
	locals "OnlineJudge/sessions"
	"github.com/jmoiron/sqlx"
)

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
