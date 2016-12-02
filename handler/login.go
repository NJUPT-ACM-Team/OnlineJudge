package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
)

func (this *Handler) LoginInit(response *api.LoginInitResponse, req *api.LoginInitRequest) {
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	defer this.CloseDB()

	if this.session.IsLogin() == true {
		response.Version = "login" + req.GetVersion()
	} else {
		response.Version = "notlogin" + req.GetVersion()
	}
}

func (this *Handler) LoginAuth(response *api.LoginAuthResponse, req *api.LoginAuthRequest) {
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Authentic the login information
	um := models.NewUserModel()
	is_login, err := um.Auth(this.tx, req.GetUsername(), []byte(req.GetPassword()))
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	if is_login == false {
		api.MakeResponseError(response, this.debug, api.PBAuthFailure, nil)
		return
	}

	// Query necessary information: username, user_id, privilege
	user, err := models.Query_User_By_Username(this.tx, req.GetUsername(), []string{"username", "user_id", "privilege"}, nil)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
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
}
