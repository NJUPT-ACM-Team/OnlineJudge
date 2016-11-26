package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"
)

func (this *Handler) LoginInit(li *api.LoginInitRequest) *api.LoginInitResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewLoginInitResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()
	return nil
}

func (this *Handler) LoginAuth(li *api.LoginAuthRequest) *api.LoginAuthResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewLoginAuthResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()

	// Authentic the login information
	um := models.NewUserModel()
	_, err := um.Auth(this.tx, li.GetUsername(), li.GetPasswd())
	if err != nil {

	}

	return nil
}
