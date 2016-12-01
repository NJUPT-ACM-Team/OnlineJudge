package handler

import (
	"testing"

	"OnlineJudge/handler/api"
)

func TestLoginAuth(t *testing.T) {
	handler, session := NewHandlerForTest()
	req := &api.LoginAuthRequest{
		Username: "kevince",
		Password: []byte("abc"),
	}
	res := handler.LoginAuth(req)
	t.Log(res)
	t.Log(session)

	req.Password = []byte("123")
	res = handler.LoginAuth(req)
	t.Log(res)
}
