package irpc

import (
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/pbgen/rpc"

	"golang.org/x/net/context"
	// "log"
)

func (this *helperServer) UpdateSubmissionStatus(ctx context.Context, req *rpc.UpdateSubmissionStatusRequest) (*rpc.UpdateSubmissionStatusResponse, error) {
	DB := db.New()
	defer DB.Close()
	tx := DB.MustBegin()
	defer tx.Rollback()

	sub := &models.Submission{
		RunId:           req.GetRunId(),
		Status:          req.GetStatus(),
		StatusCode:      req.GetStatusCode(),
		TestcasesPassed: int(req.GetTestcasesPassed()),
		TimeUsed:        int(req.GetTimeUsed()),
		MemoryUsed:      int(req.GetMemoryUsed()),
		CEInfo:          req.GetCeInfo(),
	}
	subm := models.NewSubmissionModel()

	if err := subm.UpdateStatus(tx, sub); err != nil {
		return &rpc.UpdateSubmissionStatusResponse{Success: false}, err
	}
	if err := tx.Commit(); err != nil {
		// log.(err)
		return &rpc.UpdateSubmissionStatusResponse{Success: false}, err
	}
	return &rpc.UpdateSubmissionStatusResponse{Success: true}, nil
}

type SubmissionStatus struct {
	RunId           int64
	Status          string
	StatusCode      string
	TestcasesPassed int32
	TimeUsed        int32
	MemoryUsed      int32
	CEInfo          string
}

func (this *Helper) UpdateSubmissionStatus(subs *SubmissionStatus) (*rpc.UpdateSubmissionStatusResponse, error) {
	req := &rpc.UpdateSubmissionStatusRequest{
		RunId:           subs.RunId,
		Status:          subs.Status,
		StatusCode:      subs.StatusCode,
		TestcasesPassed: subs.TestcasesPassed,
		TimeUsed:        subs.TimeUsed,
		MemoryUsed:      subs.MemoryUsed,
		CeInfo:          subs.CEInfo,
	}
	return this.client.UpdateSubmissionStatus(context.Background(), req)
}
