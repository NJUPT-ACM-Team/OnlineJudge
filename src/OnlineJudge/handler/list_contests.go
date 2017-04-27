package handler

import (
	"OnlineJudge/pbgen/api"
)

func (this *BasicHandler) ListContests(response *api.ListContestsResponse, req *api.ListContestsRequest) {
	defer PanicHandler(response, this.debug)
	this.OpenDBU()
	defer this.CloseDBU()
}

/*
func ListContests_BuildResponse(
	tx *sqlx, Tx,
	response *api.ListProblemsResponse,
	req *api.ListProblemsRequest,
	show_hidden bool,
	debug bool) {

	filter := req.GetFilter()
	page, err := models.XQuery_List_Contests_With_Filter(
		tx,
		show_hidden,
		filter.GetCtype(),
		filter.GetIsPublic(),
		filter.GetIsVirtual(),
		req.GetOrderBy(),
		req.GetIsDesc(),
		req.GetOffset(),
		req.GetPerPage(),
		req.GetCurrentPage(),
		nil,
		nil,
	)
}
*/
