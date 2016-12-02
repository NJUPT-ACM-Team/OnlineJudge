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
		Password: "abc",
	}
	handler.LoginAuth(res, req)
	t.Log(res)
	t.Log(session)

	res.Reset()
	req.Password = "123"
	handler.LoginAuth(res, req)
	t.Log(res)
}
