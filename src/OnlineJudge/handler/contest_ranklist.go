package handler

import (
	"OnlineJudge/base"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/api"

	"github.com/jmoiron/sqlx"

	"sort"
	"strings"
	"time"
)

const (
	WATime = 20
)

func (this *AdminHandler) ContestRanklist(response *api.ContestRanklistResponse, req *api.ContestRanklistRequest) {
	defer PanicHandler(response, this.debug)
	// tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
}

func (this *BasicHandler) ContestRanklist(response *api.ContestRanklistResponse, req *api.ContestRanklistRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, true, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Rank {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	// make ranklist
	ContestRanklist_BuildResponse(tx, response, req.GetContestId(), req.GetRankType(), this.debug)
}

func (this *UserHandler) ContestRanklist(response *api.ContestRanklistResponse, req *api.ContestRanklistRequest) {
	defer PanicHandler(response, this.debug)
	tx := this.dbu.MustBegin()
	defer this.dbu.Rollback()
	access, err := CheckContestAccess(
		tx, false, req.GetContestId(), this.session.GetUserId(), this.debug)
	PanicOnError(err)
	if access.If404 {
		MakeResponseError(response, this.debug, PBContestNotFound, nil)
		return
	}
	if !access.Rank {
		MakeResponseError(response, this.debug, PBUnauthorized, nil)
		return
	}
	ContestRanklist_BuildResponse(tx, response, req.GetContestId(), req.GetRankType(), this.debug)
}

type ICPCColumn struct {
	Label   string
	IsAC    bool
	IsFB    bool
	Attempt int
	Minutes int64
	Seconds int64
}

type ICPCRank struct {
	Username     string
	ACNum        int32
	TotalMins    int64
	TotalSeconds int64
	Cols         []ICPCColumn
}

type ICPCRanks []*ICPCRank

func (this ICPCRanks) Len() int {
	return len(this)
}
func (this ICPCRanks) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func (this ICPCRanks) Less(i, j int) bool {
	if this[i].ACNum > this[j].ACNum {
		return true
	} else if this[i].ACNum < this[j].ACNum {
		return false
	}
	if this[i].TotalSeconds < this[j].TotalSeconds {
		return true
	}
	return false
}

func ContestRanklist_BuildResponse(
	tx *sqlx.Tx,
	response *api.ContestRanklistResponse,
	contest_id int64,
	rank_type string,
	debug bool,
) {
	cst, err := models.Query_Contest_By_ContestId(tx, contest_id, nil, nil)
	PanicOnError(err)
	if cst == nil {
		MakeResponseError(response, debug, PBContestNotFound, nil)
		return
	}
	subs, err := models.XQuery_ContestRanklist_Submissions(tx, contest_id, cst.StartTime, cst.EndTime)
	PanicOnError(err)
	labels, err := models.Query_ContestProblemLabels_By_ContestId(tx, contest_id)
	PanicOnError(err)
	usernames, err := models.Query_ContestUsers_By_ContestId(tx, contest_id)
	PanicOnError(err)

	switch strings.ToLower(rank_type) {
	case "icpc":
		MakeICPCRank(response, usernames, labels, subs, cst.StartTime)
		return
	}
}

func MakeICPCRank(
	response *api.ContestRanklistResponse,
	users []string,
	labels []string,
	subs []models.ContestSubmissionExt,
	start_time time.Time,
) {
	response.RankIcpc = &api.RankICPC{}

	response.RankIcpc.Labels = labels
	// init user map
	userMap := make(map[string]*ICPCRank)
	for _, u := range users {
		userMap[u] = &ICPCRank{}
		userMap[u].Username = u
		userMap[u].Cols = make([]ICPCColumn, len(labels))
		for i, l := range labels {
			userMap[u].Cols[i].Label = l
		}
	}
	// init label map
	fbMap := make(map[string]bool)
	for _, l := range labels {
		fbMap[l] = false
	}

	for _, sub := range subs {
		u := sub.Username
		l := sub.Label
		idx := base.LabelToInt(l)
		switch sub.StatusCode {
		case "ac":
			if userMap[u].Cols[idx].IsAC {
				continue
			}
			if fbMap[l] == false {
				fbMap[l] = true
				userMap[u].Cols[idx].IsFB = true
			}
			userMap[u].Cols[idx].IsAC = true
			userMap[u].ACNum += 1
			// caculate time
			dur := sub.SubmitTime.Sub(start_time)
			mins := int64(dur.Minutes())
			secs := int64(dur.Seconds())
			userMap[u].Cols[idx].Minutes += mins
			userMap[u].Cols[idx].Seconds += secs
			userMap[u].TotalMins += userMap[u].Cols[idx].Minutes
			userMap[u].TotalSeconds += userMap[u].Cols[idx].Seconds

		case "pe", "wa", "tle", "re", "mle", "ole", "ce":
			if userMap[u].Cols[idx].IsAC {
				continue
			}
			userMap[u].Cols[idx].Minutes += WATime
			userMap[u].Cols[idx].Seconds += WATime * 20
			userMap[u].Cols[idx].Attempt += 1
		}
	}
	var ranks []*ICPCRank
	for _, v := range userMap {
		ranks = append(ranks, v)
	}
	sort.Sort(ICPCRanks(ranks))
	lines := []*api.RankICPC_PerLine{}
	for k, r := range ranks {
		line := &api.RankICPC_PerLine{
			Rank:      int32(k + 1),
			Username:  r.Username,
			AcNum:     r.ACNum,
			TotalMins: r.TotalMins,
		}
		cols := []*api.RankICPC_PerLine_Column{}
		for _, c := range r.Cols {
			col := &api.RankICPC_PerLine_Column{
				IsAc:  c.IsAC,
				Label: c.Label,
				// IsAc:    true,
				IsFb:    c.IsFB,
				Minutes: c.Minutes,
				Attempt: int32(c.Attempt),
			}
			cols = append(cols, col)
		}
		line.Cols = cols
		lines = append(lines, line)
	}
	response.RankIcpc.Lines = lines
}
