package utils

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

func SetSystemError(tx *sqlx.Tx, run_id int64) {
	sbm := models.NewSubmissionModel()
	sub := &models.Submission{
		RunId:           run_id,
		Status:          "System Error",
		StatusCode:      "se",
		TestCasesPassed: 0,
	}
	if err := sbm.UpdateStatus(tx, sub); err != nil {
		log.Panic(err)
	}
	if err := tx.Commit(); err != nil {
		log.Panic(err)
	}
}

func SubmitToMQ(jmq *mq.MQ, req *rpc.SubmitCodeRequest) {
	DB := db.New()
	defer DB.Close()
	tx := DB.MustBegin()
	sub, err := models.Query_Submission_By_RunId(tx, req.GetRunId(), nil, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Get Info of meta problem
	mp, err := models.Query_MetaProblem_By_MetaPid(tx, sub.MetaPidFK, []string{"oj_name", "oj_pid", "is_spj"}, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Get testcases of meta problem
	tcs, err := models.Query_TestCases_By_MetaPid(tx, sub.MetaPidFK, nil, nil)
	if err != nil {
		log.Println(err)
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

	// Get language of submission
	lang, err := models.Query_Language_By_LangId(tx, sub.LangIdFK, nil, nil)
	if err != nil {
		log.Println(err)
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
		Testcases: testcases,
	}

	buffer, err := proto.Marshal(request)
	if err != nil {
		log.Println(err)
		return
	}
	if request.IsLocal == true {
		if err := jmq.PublishLJ(buffer); err != nil {
			SetSystemError(tx, sub.RunId)
		}
	} else {
		if err := jmq.PublishVJ(buffer); err != nil {
			SetSystemError(tx, sub.RunId)
		}
	}
}
