package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *BasicHandler) ContestAuth(response *api.ContestAuthResponse, req *api.ContestAuthRequest) {
	MakeResponseError(response, this.debug, PBLoginRequired, nil)
}

func (this *UserHandler) ContestAuth(response *api.ContestAuthResponse, req *api.ContestAuthRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

	response.Success = false
	cst, err := models.Query_Contest_By_ContestId(tx, req.GetContestId(), nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
	}

	if cst.IsProtected() || cst.IsPrivate() {
		response.ContestId = cst.ContestId
		// check password
		if req.GetPassword() == cst.Password {
			// check and add user to contestusers
			check, err := CheckContestUser(
				tx, cst.ContestId, this.session.GetUserId())
			PanicOnError(err)
			if !check {
				err = AddUserToContest(
					tx, this.session.GetUserId(), cst.ContestId)
				PanicOnError(err)
				this.dbu.MustCommit()
			}
			response.Success = true
		}
	}
}

func CheckContestUser(tx *sqlx.Tx, contest_id, user_id int64) (bool, error) {
	cst, err := models.Query_ContestUser_By_ContestId_And_UserId(
		tx, contest_id, user_id)
	if err != nil {
		return false, err
	}
	if cst == nil {
		return false, nil
	}
	return true, nil
}

func AddUserToContest(tx *sqlx.Tx, user_id, contest_id int64) error {
	cum := models.NewContestUserModel()
	cu := &models.ContestUser{
		UserIdFK:    user_id,
		ContestIdFK: contest_id,
	}
	_, err := cum.Insert(tx, cu)
	return err
}
