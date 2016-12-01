package api

import (
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
