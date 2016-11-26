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

func NewLoginInitResponseError(debug bool, code int32, err error) *LoginInitResponse {
	errmsg := "Failed to init."
	if debug == true && err != nil {
		errmsg = err.Error()
	}
	theerror := &Error{
		Code: code,
		Msg:  errmsg,
	}
	return &LoginInitResponse{
		Error: theerror,
	}
}

func NewLoginAuthResponseError(debug bool, code int32, err error) *LoginAuthResponse {
	errmsg := "Wrong username or password."
	if debug == true && err != nil {
		errmsg = err.Error()
	}
	theerror := &Error{
		Code: code,
		Msg:  errmsg,
	}
	return &LoginAuthResponse{
		Error: theerror,
	}
}
