package handler

import (
	"OnlineJudge/pbgen/api"
)

func (this *BasicHandler) Logout(response *api.LogoutResponse, req *api.LogoutRequest) {

}

func (this *UserHandler) Logout(response *api.LogoutResponse, req *api.LogoutRequest) {
	this.session.Logout()
}
