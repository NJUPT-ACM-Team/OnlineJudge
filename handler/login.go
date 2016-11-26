package handler

import (
	//"OnlineJudge/models"
	"OnlineJudge/handler/api"
)

func (this *Handler) LoginInit(li *api.LoginInitRequest) *api.LoginInitResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()
}

func (this *Handler) LoginValidation(li *api.LoginValidationRequest) *api.LoginValidationResponse {
	if err := this.OpenDB(); err != nil {
		return api.NewSubmitResponseError(this.debug, 500, err)
	}
	defer this.CloseDB()

	// Validate the login information
	um := NewUserModel()
	r, err := um.Validate(this.tx, li.GetUsername(), li.GetPasswd())
	if err != nil {

	}

	return nil
}
