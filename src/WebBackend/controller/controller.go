package controller

import (
	"OnlineJudge/base"
	"OnlineJudge/handler"
	"OnlineJudge/pbgen/api"
	"OnlineJudge/sessions/websession"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"

	"encoding/json"
	"errors"
	// "fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

var (
	marshaler = jsonpb.Marshaler{
		EmitDefaults: true,
		OrigName:     true,
	}
	unmarshaler = jsonpb.Unmarshaler{
		AllowUnknownFields: true,
	}
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

func makeParamTypeRight(request interface{}, key string, val string) (interface{}, error) {
	key = base.PBTagToFieldName(key)
	r := reflect.ValueOf(request)
	f := reflect.Indirect(r).FieldByName(key)
	// fmt.Println(f)
	if f.IsValid() {
		switch f.Interface().(type) {
		case int32:
			i, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			return int32(i), nil
		case string:
			return val, nil
		case bool:
			switch strings.ToLower(val) {
			case "true", "1":
				return true, nil
			case "false", "0":
				return false, nil
			}
			return nil, errors.New("unknown boolean type")
		}
		return nil, errors.New("unknown type")
	}
	return nil, errors.New("zero value")
}

func (this *Controller) Prepare(response Response, request interface{}, w http.ResponseWriter, r *http.Request) (*websession.WebSession, error) {
	// Decode json to pb
	switch r.Method {
	case "POST":
		if err := DecodePBFromJsonStream(io.LimitReader(r.Body, this.MaxSize), request.(proto.Message)); err != nil {
			handler.MakeResponseError(response, this.debug, handler.PBBadRequest, err)
			return nil, err
		}
	case "GET":
		// TODO: improvement
		// vars := r.URL.Query()
		vars := make(map[string]interface{})
		for k, _ := range r.URL.Query() {
			v, err := makeParamTypeRight(request, k, r.FormValue(k))
			if err != nil {
				continue
			}
			// fmt.Println(request.MessageType(k))
			vars[k] = v
		}
		// fmt.Println(vars)
		js, err := json.Marshal(vars)
		if err != nil {
			handler.MakeResponseError(response, this.debug, handler.PBBadRequest, err)
			return nil, err
		}
		// fmt.Println(string(js))
		if err := DecodePBFromJson(string(js), request.(proto.Message)); err != nil {
			handler.MakeResponseError(response, this.debug, handler.PBBadRequest, err)
			return nil, err
		}
	}

	session, err := this.GetSession(r)
	if err != nil {
		handler.MakeResponseError(response, this.debug, handler.PBInternalError, err)
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

func SetWebResponse(
	w http.ResponseWriter,
	response Response,
	webresponse *api.WebResponse) {

	if webresponse.GetError() == nil && response.GetError() != nil {
		errmsg := *(response.GetError())
		webresponse.Error = &errmsg
		response.GetError().Reset()
	}

	SetResponse(w, webresponse)
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
	if err := EncodePBToJsonStream(w, response); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// TODO: add captcha information to respnose and sessions
func SetResponseWithCAPTCHA(w http.ResponseWriter, response Response) {
	SetResponse(w, response)
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
