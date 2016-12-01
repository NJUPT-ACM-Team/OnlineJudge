package handler

import (
	"testing"

	"OnlineJudge/handler/api"
)

func TestLoginAuth(t *testing.T) {
	res := &api.LoginAuthResponse{}
	handler, session := NewHandlerForTest()
	req := &api.LoginAuthRequest{
		Username: "kevince",
		Password: []byte("abc"),
	}
	handler.LoginAuth(res, req)
	t.Log(res)
	t.Log(session)

	res.Reset()
	req.Password = []byte("123")
	handler.LoginAuth(res, req)
	t.Log(res)
}
