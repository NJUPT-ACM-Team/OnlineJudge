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

	return nil
}
