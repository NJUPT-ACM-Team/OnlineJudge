package handler

import (
	"OnlineJudge/pbgen/api"
)

func (this *AdminHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

}

func (this *BasicHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

}

func (this *UserHandler) ContestListSubmissions(response *api.ContestListSubmissionsResponse, req *api.ListSubmissionsRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()

}
