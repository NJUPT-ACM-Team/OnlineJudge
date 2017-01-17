package handler

import (
	"OnlineJudge/db"
	"OnlineJudge/pbgen/api"
	locals "OnlineJudge/sessions"
	"github.com/jmoiron/sqlx"
)

// For testing purpose
import (
	"OnlineJudge/sessions/websession"
	"github.com/gorilla/sessions"
	"net/http"
)

func NewHandlerForTest() (*Handler, *sessions.Session) {
	db.InitTest()
	var store = sessions.NewCookieStore([]byte("something-very-secret"))
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	session, _ := store.New(req, "my session")
	sess := websession.NewSession(session)
	return NewHandler(sess, true), session
}

//

type HandlerInterface interface {
	ListProblems(*api.ListProblemsResponse, *api.ListProblemsRequest)
	ListContests(*api.ListContestsResponse, *api.ListContestsRequest)
	ListSubmissions(*api.ListSubmissionsResponse, *api.ListSubmissionsRequest)
	LoginInit(*api.LoginInitResponse, *api.LoginInitRequest)
	LoginAuth(*api.LoginAuthResponse, *api.LoginAuthRequest)
	Logout(*api.LogoutResponse, *api.LogoutRequest)
	Register(*api.RegisterResponse, *api.RegisterRequest)
	ShowProblem(*api.ShowProblemResponse, *api.ShowProblemRequest)
	Submit(*api.SubmitResponse, *api.SubmitRequest)
}

type Handler struct {
	session locals.Session
	db      *sqlx.DB
	tx      *sqlx.Tx
	debug   bool
}

type AdminHandler struct {
	Handler
}

func (this *AdminHandler) Check() bool {
	return this.session.IsRoot()
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
