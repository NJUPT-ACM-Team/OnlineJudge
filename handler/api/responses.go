package api

func NewSubmitResponseError(debug bool, code int32, err error) *SubmitResponse {
	errmsg := "Internal Error"
	if debug == true {
		errmsg = err.Error()
	}
	theerror := &Error{
		Code: code,
		Msg:  errmsg,
	}
	return &SubmitResponse{
		Error: theerror,
	}
}

func NewLoginValidationResponseError(debug bool, code int32, err error) *LoginValidationResponse {
	errmsg := "Wrong username or password."
	if debug == true && err != nil {
		errmsg = err.Error()
	}
	theerror := &Error{
		Code: code,
		Msg:  errmsg,
	}
	return &LoginValidationResponse{
		Error: theerror,
	}
}
