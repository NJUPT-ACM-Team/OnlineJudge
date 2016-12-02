package handler

import (
	"OnlineJudge/handler/api"

	"testing"
)

func TestShowProblem(t *testing.T) {
	res := &api.ShowProblemResponse{}
	handler, _ := NewHandlerForTest()
	req := &api.ShowProblemRequest{
		ProblemSid: "zoj#1000",
	}
	handler.ShowProblem(res, req)
	t.Log(res)
}
