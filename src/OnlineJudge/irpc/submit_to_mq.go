package irpc

import (
	"OnlineJudge/db"
	"OnlineJudge/models"
	"OnlineJudge/mq"
	msgs "OnlineJudge/pbgen/messages"
	"OnlineJudge/pbgen/rpc"

	"github.com/golang/protobuf/proto"
	"github.com/jmoiron/sqlx"

	"log"
)

func MustSetSystemError(tx *sqlx.Tx, run_id int64) {
	sbm := models.NewSubmissionModel()
	if err := sbm.SetSystemError(tx, run_id); err != nil {
		log.Panic(err)
	}
	if err := tx.Commit(); err != nil {
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
	if err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	// Get Info of meta problem
	mp, err := models.Query_MetaProblem_By_MetaPid(
		tx, sub.MetaPidFK, []string{"oj_name", "oj_pid", "is_spj"}, nil)
	if err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	// Get language of submission
	lang, err := models.Query_Language_By_LangId(tx, sub.LangIdFK, nil, nil)
	if err != nil {
		log.Println(err)
		MustSetSystemError(tx, sub.RunId)
		return
	}

	request := &msgs.SubmitMQ{
		RunId:   req.GetRunId(),
		OjName:  mp.OJName,
		OjPid:   mp.OJPid,
		Code:    sub.Code,
		IsLocal: mp.OJName == "local",
		IsSpj:   sub.IsSpj,
		Language: &msgs.SubmitLanguage{
			Compiler:    lang.Compiler,
			Lang:        lang.Language,
			OptionValue: lang.OptionValue,
		},
	}

	// Get testcases of meta problem
	if mp.OJName == "local" {
		tcs, err := models.Query_TestCases_By_MetaPid(tx, sub.MetaPidFK, nil, nil)
		if err != nil {
			log.Println(err)
			MustSetSystemError(tx, sub.RunId)
			return
		}
		testcases := []*msgs.TestCase{}
		for _, tc := range tcs {
			temp := &msgs.TestCase{
				CaseId:     tc.CaseId,
				Input:      tc.Input,
				InputHash:  tc.InputMD5,
				Output:     tc.Input,
				OutputHash: tc.OutputMD5,
			}
			testcases = append(testcases, temp)
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
