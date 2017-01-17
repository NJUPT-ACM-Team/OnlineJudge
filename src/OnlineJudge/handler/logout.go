package handler

import (
	"OnlineJudge/pbgen/api"
)

func (this *Handler) Logout(response *api.LogoutResponse, req *api.LogoutRequest) {
	this.session.Logout()
}
