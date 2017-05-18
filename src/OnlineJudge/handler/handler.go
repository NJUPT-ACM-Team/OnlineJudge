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

func NewHandlerForTest() (*UserHandler, *sessions.Session) {
	db.InitTest()
	var store = sessions.NewCookieStore([]byte("something-very-secret"))
	req, _ := http.NewRequest("GET", "http://www.example.com", nil)
	session, _ := store.New(req, "my session")
	sess := websession.NewSession(session)
	return &UserHandler{BasicHandler{session: sess, debug: true}}, session
}

//

type Handler interface {
	About(*api.AboutResponse, *api.AboutRequest)                      // OK
	ListProblems(*api.ListProblemsResponse, *api.ListProblemsRequest) // OK
	ListContests(*api.ListContestsResponse, *api.ListContestsRequest)
	ListSubmissions(*api.ListSubmissionsResponse, *api.ListSubmissionsRequest) // OK
	LoginInit(*api.LoginInitResponse, *api.LoginInitRequest)
	LoginAuth(*api.LoginAuthResponse, *api.LoginAuthRequest)       // OK
	Logout(*api.LogoutResponse, *api.LogoutRequest)                // OK
	Register(*api.RegisterResponse, *api.RegisterRequest)          // OK
	ShowProblem(*api.ShowProblemResponse, *api.ShowProblemRequest) // OK
	ShowContest(*api.ShowContestResponse, *api.ShowContestRequest) // OK
	Submit(*api.SubmitResponse, *api.SubmitRequest)
}

type BasicHandler struct {
	session locals.Session
	db      *sqlx.DB
	dbu     *db.DBUtil
	tx      *sqlx.Tx
	debug   bool
}

type UserHandler struct {
	BasicHandler
}

type AdminHandler struct {
	UserHandler
}

func CheckAdmin(sess locals.Session) bool {
	return sess.IsRoot()
}

func CheckLogin(sess locals.Session) bool {
	return sess.IsLogin()
}

func NewHandler(sess locals.Session, dbg bool) Handler {
	basic := &BasicHandler{
		session: sess,
		debug:   dbg,
	}
	if !CheckLogin(sess) {
		return basic
	}
	user := &UserHandler{
		BasicHandler: *basic,
	}
	if !CheckAdmin(sess) {
		return user
	}
	return &AdminHandler{
		UserHandler: *user,
	}
}

func (this *BasicHandler) OpenDBU() {
	d, err := db.NewDB()
	if err != nil {
		panic(err)
	}
	this.dbu = db.NewDBU(d)
}

func (this *BasicHandler) CloseDBU() {
	this.dbu.Close()
}

func PanicOnError(err error) {
	if err != nil {
		panic(err)
	}
}

func PanicHandler(response interface{}, dbg bool) {
	if err := recover(); err != nil {
		MakeResponseError(response, dbg, PBInternalError, err.(error))
	}
}

func (this *BasicHandler) OpenDB() error {
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
func (this *BasicHandler) Commit() error {
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

func (this *BasicHandler) CloseDB() {
	this.tx.Rollback()
	this.db.Close()
}
