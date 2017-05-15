package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"
)

func (this *BasicHandler) ListContests(response *api.ListContestsResponse, req *api.ListContestsRequest) {
	defer PanicHandler(response, this.debug)
	this.OpenDBU()
	defer this.CloseDBU()

	tx := this.dbu.MustBegin()
	ListContests_BuildResponse(
		tx,
		response,
		req,
		false,
		this.debug,
	)
}

func ListContests_BuildResponse(
	tx *sqlx.Tx,
	response *api.ListContestsResponse,
	req *api.ListContestsRequest,
	show_hidden bool,
	debug bool) {

	page, err := models.XQuery_List_Contests_With_Filter(
		tx,
		int(req.GetPerPage()),
		int(req.GetCurrentPage()),
		req.GetOrderBy(),
		req.GetIsDesc(),
		req.GetFilterCtype(),
		req.GetFilterIsPublic(),
		req.GetFilterIsVirtual(),
		nil,
		nil,
	)
	PanicOnError(err)

	// Build response
	lines := []*api.ListContestsResponse_PerLine{}
	for _, contest := range page.Contests {
		// TODO: finish status and access
		line := &api.ListContestsResponse_PerLine{
			ContestId:   contest.ContestId,
			Title:       contest.Title,
			StartTime:   contest.StartTime.String(),
			EndTime:     contest.EndTime.String(),
			Status:      "ended",
			Access:      "public",
			ContestType: contest.ContestType,
		}
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}
