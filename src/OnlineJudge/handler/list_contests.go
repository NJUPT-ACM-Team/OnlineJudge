package handler

import (
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"

	"time"
)

func (this *BasicHandler) ListContests(response *api.ListContestsResponse, req *api.ListContestsRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
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
			ContestId: contest.ContestId,
			Title:     contest.Title,
			StartTime: contest.StartTime.String(),
			EndTime:   contest.EndTime.String(),
			// Status:      "ended",
			// Access:      "public",
			ContestType: contest.ContestType,
		}
		if contest.Password == "" {
			line.Access = "public"
		} else {
			line.Access = "private"
		}
		statusCode := JudgeContestStatus(&contest, time.Now().UTC())
		line.Status = ContestStatusCodeString(statusCode)
		lines = append(lines, line)
	}
	response.Lines = lines
	response.TotalLines = int32(page.TotalLines)
	response.TotalPages = int32(page.TotalPages)
	response.CurrentPage = int32(page.CurrentPage)
}

// -1 in-future, 0 in-progress 1 ended
func JudgeContestStatus(cst *models.Contest, t time.Time) int {
	if t.Before(cst.StartTime) {
		return -1
	}
	if t.After(cst.EndTime) {
		return 1
	}
	return 0
}

func ContestStatusCodeString(c int) string {
	switch c {
	case -1:
		return "in-future"
	case 0:
		return "in-progress"
	case 1:
		return "ended"
	}
	return "unknown"
}
