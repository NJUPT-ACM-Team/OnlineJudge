package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
)

func (this *Handler) LoginInit(req *api.LoginInitRequest) *api.LoginInitResponse {
	var response = &api.LoginInitResponse{}
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	defer this.CloseDB()
	return response
}

func (this *Handler) LoginAuth(req *api.LoginAuthRequest) *api.LoginAuthResponse {
	var response = &api.LoginAuthResponse{}
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	defer this.CloseDB()

	// Authentic the login information
	um := models.NewUserModel()
	is_login, err := um.Auth(this.tx, req.GetUsername(), req.GetPassword())
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	if is_login == false {
		api.MakeResponseError(response, this.debug, api.PBAuthFailure, nil)
		return response
	}

	// Query necessary information: username, user_id, privilege
	user, err := models.Query_User_By_Username(this.tx, req.GetUsername(), []string{"username", "user_id", "privilege"}, nil)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}

	// Set IPAddr into database

	// Set session
	this.session.SetUsername(user.Username)
	this.session.SetUserId(user.UserId)
	this.session.SetPrivilege(user.Privilege)

	// Make response
	response.Msg = "Hello " + user.Username + "!"
	response.Username = user.Username
	response.Privilege = user.Privilege

	return response
}
