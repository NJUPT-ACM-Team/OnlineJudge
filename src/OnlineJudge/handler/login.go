package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"time"
)

func (this *BasicHandler) LoginInit(response *api.LoginInitResponse, req *api.LoginInitRequest) {
	if this.session.IsLogin() == true {
		response.Version = "login" + req.GetVersion()
	} else {
		response.Version = "notlogin" + req.GetVersion()
	}
}

func (this *BasicHandler) LoginAuth(response *api.LoginAuthResponse, req *api.LoginAuthRequest) {
	defer func() {
		if err := recover(); err != nil {
			MakeResponseError(response, this.debug, PBInternalError, err.(error))
		}
	}()
	this.OpenDBU()
	defer this.CloseDBU()
	tx := this.dbu.MustBegin()

	// Authentic the login information
	um := models.NewUserModel()
	is_login, err := um.Auth(tx, req.GetUsername(), []byte(req.GetPassword()))
	if err != nil {
		MakeResponseError(response, this.debug, PBAuthFailure, err)
		return
	}
	if is_login == false {
		MakeResponseError(response, this.debug, PBAuthFailure, nil)
		return
	}

	// Query necessary information: username, user_id, privilege
	user, err := models.Query_User_By_Username(
		tx, req.GetUsername(),
		[]string{"username", "user_id", "privilege"},
		nil)
	if err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// Save IPAddr into database
	ip_addr := this.session.GetIPAddr()
	if err := um.UpdateIPAddr(tx, user.Username, ip_addr); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// Save last login time
	if err := um.UpdateLastLoginTime(tx, user.Username, time.Now()); err != nil {
		MakeResponseError(response, this.debug, PBInternalError, err)
		return
	}

	// Commit change
	this.dbu.MustCommit()

	// Set session
	this.session.SetUsername(user.Username)
	this.session.SetUserId(user.UserId)
	this.session.SetPrivilege(user.Privilege)

	// Make response
	response.Msg = "Hello " + user.Username + "!"
	response.Username = user.Username
	response.Privilege = user.Privilege
}
