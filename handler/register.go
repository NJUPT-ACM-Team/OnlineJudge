package handler

import (
	"OnlineJudge/handler/api"
	"OnlineJudge/models"

	"github.com/jmoiron/sqlx"

	"time"
)

func CheckRegisterRequest(tx *sqlx.Tx, req *api.RegisterRequest, res *api.RegisterResponse) error {

	// return errors.New(MsgBadRequestError)
	return nil
}

func (this *Handler) Register(req *api.RegisterRequest) *api.RegisterResponse {
	var response = &api.RegisterResponse{}
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	defer this.CloseDB()

	// Check Request
	if err := CheckRegisterRequest(this.tx, req, response); err != nil {
		api.MakeResponseError(response, this.debug, api.PBBadRequest, err)
		return response
	}

	// Insert into database
	um := models.NewUserModel()
	user := &models.User{
		Username:     req.GetUsername(),
		Password:     req.GetPassword(),
		Email:        req.GetEmail(),
		Phone:        req.GetPhone(),
		School:       req.GetSchool(),
		Motto:        req.GetMotto(),
		RegisterTime: time.Now(),
	}
	user_id, err := um.Insert(this.tx, user)
	if err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}
	if err := this.Commit(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return response
	}

	// Make response
	response.UserId = user_id

	return response
}
