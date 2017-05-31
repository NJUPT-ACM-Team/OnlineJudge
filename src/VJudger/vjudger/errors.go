package vjudger

import (
	"errors"
)

var (
	CharsetError  = errors.New("Failed to decode charset")
	BadInternet   = errors.New("Can't connect to vjudger server")
	LoginFailed   = errors.New("Failed to login")
	SubmitFailed  = errors.New("Failed to submit code")
	NoSuchProblem = errors.New("No such problem")
	JudgeFailed   = errors.New("Failed to judge solution")
	BadStatus     = errors.New("Can't find status")
)
