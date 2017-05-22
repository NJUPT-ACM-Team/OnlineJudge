package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"

	"errors"
	"time"
)

// TODO: Some checks
func CheckRegisterRequest(tx *sqlx.Tx, req *api.RegisterRequest, res *api.RegisterResponse) error {
	is_err := false
	// Check username
	if base.CheckUsername(req.GetUsername()) != true {
		res.CheckUsername = "invalid username format"
		is_err = true
	} else {
		if_exist, err := models.Query_If_User_Exists(tx, req.GetUsername())
		if err != nil {
			res.CheckUsername = "internel error while checking username"
			is_err = true
		} else {
			if if_exist {
				res.CheckUsername = "this one has already been registered"
				is_err = true
			}
		}
	}

	// Check password

	// Check email
	if base.CheckEmail(req.GetEmail()) != true {
		res.CheckEmail = "invalid email format"
		is_err = true
	}

	// Check phone
	if base.CheckPhone(req.GetPhone()) != true {
		res.CheckPhone = "invalid phone number format"
		is_err = true
	}

	// Check school

	// Check motto

	if is_err {
		return errors.New("invalid parameters")
	}
	return nil
}

func (this *BasicHandler) Register(response *api.RegisterResponse, req *api.RegisterRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

	// Check Request
	if err := CheckRegisterRequest(tx, req, response); err != nil {
		MakeResponseError(response, this.debug, PBBadRequest, err)
		return
	}

	// Insert into database
	um := models.NewUserModel()
	user := &models.User{
		Username:      req.GetUsername(),
		Password:      []byte(req.GetPassword()),
		Email:         req.GetEmail(),
		Phone:         req.GetPhone(),
		School:        req.GetSchool(),
		Motto:         req.GetMotto(),
		RegisterTime:  time.Now(),
		LastLoginTime: base.GetDefaultTime(),
	}
	user_id, err := um.Insert(tx, user)
	PanicOnError(err)
	this.dbu.MustCommit()

	// Make response
	response.UserId = user_id
	response.Username = req.GetUsername()
}
