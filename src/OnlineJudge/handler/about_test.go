package handler

import (
	"OnlineJudge/pbgen/api"

	"testing"
)

func TestAbout(t *testing.T) {
	res := &api.AboutResponse{}
	handler, _ := NewHandlerForTest()

	req := &api.AboutRequest{
		NeedOjsList:       true,
		NeedLanguagesList: true,
	}
	handler.About(res, req)
	t.Log(res)
}
