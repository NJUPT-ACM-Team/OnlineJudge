package api

var (
	PBInternalError   = NewPBError(500, "inernal error")
	PBAuthFailure     = NewPBError(401, "username not exists or wrong password")
	PBBadRequest      = NewPBError(400, "bad request")
	PBProblemNotFound = NewPBError(404, "problem does not exist or not visible")
	PBLoginRequired   = NewPBError(401, "login required")
)
