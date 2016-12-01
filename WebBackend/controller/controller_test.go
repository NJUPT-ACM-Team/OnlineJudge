package controller

import (
	"OnlineJudge/base"
	"OnlineJudge/handler/api"

	"net/http/httptest"
	"testing"
)

var (
	pb = &api.LoginAuthResponse{
		Msg:       "Hello Kevince",
		Username:  "Kevince",
		Privilege: "",
		Error:     &api.Error{Code: 301},
	}
)

func TestEncodePBToJson(t *testing.T) {
	json, err := EncodePBToJson(pb)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(json)
}

func TestDecodePBFromJson(t *testing.T) {
	pb := &api.LoginAuthResponse{}
	json := `{"msg":"Hello Kevince","username":"","error":{}}`
	if err := DecodePBFromJson(json, pb); err != nil {
		t.Fatal(err)
	}
	t.Log(base.IsNilOrZero(pb.Username))
	t.Log(pb)
}

func TestSetResponse(t *testing.T) {
	pb = &api.LoginAuthResponse{
		Msg:       "Hello Kevince",
		Username:  "Kevince",
		Privilege: "",
		//Error:     &api.Error{},
	}
	w := httptest.NewRecorder()
	SetResponse(w, pb)
	t.Log(w)
}
