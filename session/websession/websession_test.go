package websession

import (
	"encoding/base64"
	"github.com/gorilla/sessions"
	"net/http"
	// "net/http/httptest"
	"testing"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

func TestNewWebSession(t *testing.T) {
	req, err := http.NewRequest("GET", "http://www.example.com", nil)
	if err != nil {
		t.Fatal("failed to create request", err)
	}
	// w := httptest.NewRecorder()
	session, err := store.New(req, "my session")
	session.Values["big"] = make([]byte, base64.StdEncoding.DecodedLen(4096*2))

	mysess := NewWebSession(session)
	t.Log(mysess.IsLogin())

	/*
		err = session.Save(req, w)
		if err == nil {
			t.Fatal("expected an error, got nil")
		}
	*/

}
