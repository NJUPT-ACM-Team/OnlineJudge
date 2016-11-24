package models_basic

import (
	"OnlineJudge/models/db"
	"testing"
)

func QueryTitle_Model(oj_name string, pid int) (string, error) {
	DB, err := db.NewDB()
	if err != nil {
		return "", err
	}
	tx, err := DB.Beginx()
	if err != nil {
		return "", err
	}
	defer DB.Close()
	mpm := NewMetaProblemModel()
	ojm := NewOJInfoModel()

	oj, err := ojm.QueryByName(tx, oj_name, []string{"oj_id"}, nil)
	if err != nil {
		return "", err
	}
	mp, err := mpm.QueryByOJIdAndPid(tx, oj.OJId, pid, []string{"title"}, nil)
	if err != nil {
		return "", err
	}
	return mp.Title, nil
}

func QueryTitle_SQL(oj_name string, pid int) (string, error) {
	DB, err := db.NewDB()
	if err != nil {
		return "", err
	}
	tx, err := DB.Beginx()
	if err != nil {
		return "", err
	}
	defer DB.Close()
	var title string
	if err := tx.Get(&title, "SELECT title FROM MetaProblems WHERE oj_pid=? AND oj_id_fk=(SELECT oj_id FROM OJInfo WHERE name=?)", pid, oj_name); err != nil {
		return "", err
	}
	return title, nil
}

func QueryTitle_2SQL(oj_name string, pid int) (string, error) {
	DB, err := db.NewDB()
	if err != nil {
		return "", err
	}
	tx, err := DB.Beginx()
	if err != nil {
		return "", err
	}
	defer DB.Close()
	var oj_id int
	if err := tx.Get(&oj_id, "SELECT oj_id FROM OJInfo WHERE name=?", oj_name); err != nil {
		return "", err
	}
	var title string
	if err := tx.Get(&title, "SELECT title FROM MetaProblems WHERE oj_id_fk=? AND oj_pid=?", oj_id, pid); err != nil {
		return "", err
	}
	return title, nil
}

func BenchmarkQueryProblemModel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryTitle_Model("zoj", 1000)
	}
}

func TestQueryProblem(t *testing.T) {
	title, err := QueryTitle_Model("zoj", 1000)
	if err != nil {
		t.Errorf("Failed to query title, %s", err)
	}
	t.Log(title)
}

func BenchmarkQueryProblemSQL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryTitle_SQL("zoj", 1000)
	}
}

func BenchmarkQueryProblemSQL2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QueryTitle_2SQL("zoj", 1000)
	}
}

/*
func TestInsertProblem(t *testing.T) {
	mpm := NewMetaProblemModel()
	mp := MetaProblem{
		Title:             "A+B",
		Description:       "caculate result of a+b",
		Input:             "Two integer",
		Output:            "Sum of two integer a+b",
		SampleIn:          "1 1",
		SampleOut:         "2",
		TimeLimit:         1000,
		CaseTimeLimit:     1000,
		MemoryLimit:       65536,
		NumberOfTestCases: 10,
		Source:            "test",
		Hint:              "for test",
		Hide:              0,
		OJIdFK:            1,
		OJPid:             1000,
	}
	if err := mpm.OpenDB(); err != nil {
		t.Errorf("Failed to open db, %s", err)
		return
	}
	defer mpm.CloseDB()
	id, err := mpm.Insert(&mp)
	if err != nil {
		t.Errorf("Failed to insert meta problem, %s", err)
	}
	t.Log("result id is ", id)
	if err = mpm.Commit(); err != nil {
		t.Errorf("Failed to Commit, %s", err)
	}
}*/
