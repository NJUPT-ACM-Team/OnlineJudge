package websession

import (
	//"encoding/base64"
	"github.com/gorilla/sessions"
	"net/http"
	"net/http/httptest"
	"testing"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func TestNewSession(t *testing.T) {
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	w := httptest.NewRecorder()
	session, err := store.New(req, "my session")

	mysess := NewSession(session)
	mysess.SetUsername("kevince")
	t.Log(mysess.IsLogin())
	t.Log(session.Values["username"])
	if session.Values["username"] != "kevince" {
		t.Fatal("Failed to get username")
	}
	if err := mysess.Set("username", "a"); err == nil {
		t.Fatal("Supposed to be error, but nil")
	}
	if err := mysess.Set("test", "this is a test"); err != nil {
		t.Fatal(err)
	}
	if val, err := mysess.Get("test"); err != nil || val != "this is a test" {
		t.Fatal(err)
	}
	err = session.Save(req, w)
	if err != nil {
		t.Fatal("Save error, %s", err)
	}
	t.Log(w)

}
