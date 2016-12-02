package controller

import (
	// "OnlineJudge/handler"
	"OnlineJudge/handler/api"
	// locals "OnlineJudge/sessions"
	"OnlineJudge/sessions/websession"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	//"fmt"
	"io"
	"net/http"
	"strings"
)

var (
	marshaler   = jsonpb.Marshaler{}
	unmarshaler = jsonpb.Unmarshaler{}
)

type Controller struct {
	store   websession.Store
	debug   bool
	MaxSize int64
}

type Response interface {
	proto.Message
	GetError() *api.Error
}

func (this *Controller) GetSession(r *http.Request) (*websession.WebSession, error) {
	sess, err := this.store.Get(r, "default")
	if err != nil {
		return nil, err
	}
	return websession.NewWebSession(sess), nil
}

func (this *Controller) Prepare(response Response, request proto.Message, w http.ResponseWriter, r *http.Request) (*websession.WebSession, error) {
	// Decode json to pb
	if err := DecodePBFromJsonStream(io.LimitReader(r.Body, this.MaxSize), request); err != nil {
		api.MakeResponseError(response, this.debug, api.PBBadRequest, err)
		return nil, err
	}

	session, err := this.GetSession(r)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return nil, err
	}
	return session, nil
}

func NewController(dbg bool) *Controller {
	return &Controller{
		store:   websession.NewStore(),
		debug:   dbg,
		MaxSize: 1048576,
	}
}

func SetResponse(w http.ResponseWriter, response Response) {
	err := response.GetError()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err == nil || err.Code == 0 {
		// 200 Response
		w.WriteHeader(http.StatusOK)
	} else {
		// Set response code
		w.WriteHeader(int(err.GetCode()))
	}
	EncodePBToJsonStream(w, response)
}

func DecodePBFromJson(json string, pb proto.Message) error {
	if err := unmarshaler.Unmarshal(strings.NewReader(json), pb); err != nil {
		return err
	}
	return nil
}

func EncodePBToJson(pb proto.Message) (string, error) {
	json, err := marshaler.MarshalToString(pb)
	if err != nil {
		return "", err
	}
	return json, nil
}

func EncodePBToJsonStream(out io.Writer, pb proto.Message) error {
	return marshaler.Marshal(out, pb)
}

func DecodePBFromJsonStream(in io.Reader, pb proto.Message) error {
	return unmarshaler.Unmarshal(in, pb)
}

func (this *Controller) HelloWorld(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World")
	pb := &api.LoginAuthResponse{
		Msg:       "Hello Kevince",
		Username:  "Kevince",
		Privilege: "",
		Error:     &api.Error{},
	}
	EncodePBToJsonStream(w, pb)
}
