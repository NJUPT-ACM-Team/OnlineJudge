package models

import (
	"OnlineJudge/models/api"

	//	"github.com/golang/protobuf/jsonpb"
	// "github.com/golang/protobuf/proto"
	"testing"
)

func TestAPISubmitResponse(t *testing.T) {
	subres := &api.SubmitResponse{
		Result:       api.SubmitResponse_SUCCESS,
		Hint:         "nothing",
		SubmissionId: 123,
	}
	t.Log(subres.String())
	t.Log(subres.GetResult())
	if subres.GetResult() != 0 {
		t.Errorf("GetResult not right")
	}

}

func TestAPISubmitRequest(t *testing.T) {
	sbreq := api.SubmitRequest{
		ProblemSid:   "local#1000",
		Code:         "hello world",
		LanguageCode: 10,
		IsShared:     false,
	}
	t.Log(sbreq.String())
	if sbreq.GetProblemSid() != "local#1000" {
		t.Errorf("GetProblemId not right")
	}
	if sbreq.GetCode() != "hello world" {
		t.Errorf("GetCode not right")
	}
	if sbreq.GetLanguageCode() != 10 {
		t.Errorf("GetLanguageCode not right")
	}
	if sbreq.GetIsShared() != false {
		t.Errorf("GetIsShared not right")
	}
}
