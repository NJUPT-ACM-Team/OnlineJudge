package handler

import (
	"OnlineJudge/pbgen/api"

	"testing"
)

func TestListProblems(t *testing.T) {
	res := &api.ListProblemsResponse{}
	handler, session := NewHandlerForTest()
	session.Values[".username"] = "kevince"
	// session.Values[".privilege"] = "root"
	req := &api.ListProblemsRequest{
		PerPage:     1,
		CurrentPage: 2,
		OrderBy:     0,
		IsDesc:      false,
		Filter: &api.ListProblemsRequest_Filter{
			Oj:      "zoj",
			PStatus: 0,
		},
	}
	handler.ListProblems(res, req)
	t.Log(res)
	t.Log(session)
}
