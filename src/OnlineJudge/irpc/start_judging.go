package irpc

import (
	"OnlineJudge/base"
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/mq"
	msgs "OnlineJudge/pbgen/messages"
	"OnlineJudge/pbgen/rpc"

	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"

	"golang.org/x/net/context"

	"log"
)

func MustSetSystemError(tx *sqlx.Tx, run_id int64) {
	sbm := models.NewSubmissionModel()
	if err := sbm.SetSystemError(tx, run_id); err != nil {
		// TODO: fatal error notification
		log.Panic(err)
	}
	if err := tx.Commit(); err != nil {
		// TODO: fatal error notification
		log.Panic(err)
	}
}

func SubmitToMQ(jmq *mq.MQ, req *rpc.StartJudgingRequest) {
	log.Println(req)
	DB := db.New()
	defer DB.Close()
	tx := DB.MustBegin()
	defer tx.Rollback()

	// Get Submission info
	sub, err := models.Query_Submission_By_RunId(tx, req.GetRunId(), nil, nil)
	if sub == nil || err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	// Get Info of meta problem
	mp, err := models.Query_MetaProblem_By_MetaPid(
		tx, sub.MetaPidFK, []string{"oj_name", "oj_pid", "is_spj"}, nil)
	if mp == nil || err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	// Get language of submission
	lang, err := models.Query_Language_By_LangId(tx, sub.LangIdFK, nil, nil)
	if lang == nil || err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	// Get limits of language

	request := &msgs.SubmitMQ{
		RunId:   req.GetRunId(),
		OjName:  mp.OJName,
		OjPid:   mp.OJPid,
		Code:    sub.Code,
		IsLocal: mp.OJName == "local",
		IsSpj:   sub.IsSpj,
		// TODO: limits
		TimeLimit:   1000,
		MemoryLimit: 65536,
		SubmitTime:  base.MarshalTime(sub.SubmitTime),
		Language: &msgs.SubmitLanguage{
			Compiler:    lang.Compiler,
			Lang:        lang.Language,
			Suffix:      lang.Suffix,
			OptionValue: lang.OptionValue,
		},
	}
	log.Println("suffix:", lang.Suffix)

	// Get testcases of meta problem
	if mp.OJName == "local" {
		log.Println("use local")
		tcs, err := models.Query_TestCases_By_MetaPid(tx, sub.MetaPidFK, nil, nil)
		if err != nil {
			log.Println(err)
			MustSetSystemError(tx, sub.RunId)
			return
		}
		testcases := []*msgs.TestCase{}
		if tcs != nil {
			for _, tc := range tcs {
				temp := &msgs.TestCase{
					CaseId:     tc.CaseId,
					Input:      tc.Input,
					InputHash:  tc.InputMD5,
					Output:     tc.Output,
					OutputHash: tc.OutputMD5,
				}
				testcases = append(testcases, temp)
			}
		}
		request.Testcases = testcases
	} else {
		request.Testcases = nil
	}

	buffer, err := proto.Marshal(request)
	if err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}
	if request.IsLocal == true {
		if err := jmq.PublishLJ(buffer); err != nil {
			MustSetSystemError(tx, sub.RunId)
		}
	} else {
		if err := jmq.PublishVJ(buffer); err != nil {
			MustSetSystemError(tx, sub.RunId)
		}
	}
}

func (this *helperServer) StartJudging(ctx context.Context, req *rpc.StartJudgingRequest) (*rpc.StartJudgingResponse, error) {
	// Submit the code to MQ
	go SubmitToMQ(this.jmq, req)
	return &rpc.StartJudgingResponse{
		Received: true,
	}, nil
}

func (this *Helper) StartJudging(run_id int64) (*rpc.StartJudgingResponse, error) {
	req := &rpc.StartJudgingRequest{
		RunId: run_id,
	}
	return this.client.StartJudging(context.Background(), req)

}
