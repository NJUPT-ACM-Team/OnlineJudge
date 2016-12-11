package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"

	"time"
)

func CheckRegisterRequest(tx *sqlx.Tx, req *api.RegisterRequest, res *api.RegisterResponse) error {

	// return errors.New(MsgBadRequestError)
	return nil
}

func (this *Handler) Register(response *api.RegisterResponse, req *api.RegisterRequest) {
	if err := this.OpenDB(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}
	defer this.CloseDB()

	// Check Request
	if err := CheckRegisterRequest(this.tx, req, response); err != nil {
		api.MakeResponseError(response, this.debug, api.PBBadRequest, err)
		return
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
		return
	}
	if err := this.Commit(); err != nil {
		api.MakeResponseError(response, this.debug, api.PBInternalError, err)
		return
	}

	// Make response
	response.UserId = user_id
}
