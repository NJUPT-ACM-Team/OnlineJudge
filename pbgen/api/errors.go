package api

import (
	"net/http"
	"reflect"
)

func NewPBError(code int32, msg string) *Error {
	return &Error{
		Code: code,
		Msg:  msg,
	}
}

// response, debug, Error, debug err
func MakeResponseError(res interface{}, debug bool, pberr *Error, err error) {
	field := reflect.ValueOf(res).Elem().FieldByName("Error")
	if debug == true && err != nil {
		pberr.Debug = err.Error()
	}
	field.Set(reflect.ValueOf(pberr))
}

var (
	PBInternalError   = NewPBError(http.StatusInternalServerError, "inernal server error")
	PBAuthFailure     = NewPBError(http.StatusUnauthorized, "username not exists or wrong password")
	PBBadRequest      = NewPBError(http.StatusBadRequest, "bad request")
	PBProblemNotFound = NewPBError(http.StatusNotFound, "problem does not exist or not visible")
	PBLoginRequired   = NewPBError(http.StatusUnauthorized, "login required")
)
