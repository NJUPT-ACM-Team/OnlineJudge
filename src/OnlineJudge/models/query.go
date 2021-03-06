package models

import (
	"github.com/jmoiron/sqlx"

	"time"
)

func Query_MetaProblem_By_MetaPid(
	tx *sqlx.Tx, meta_pid int64,
	required []string,
	excepts []string) (*MetaProblem, error) {

	/*-- Func start --*/
	mp := &MetaProblem{}
	str_fields, err := GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}
	from_sql := "FROM MetaProblems"
	where_sql := "WHERE meta_pid=?"

	select_sql := JoinSQL("SELECT", str_fields, from_sql, where_sql)
	count_sql := JoinSQL("SELECT COUNT(*)", from_sql, where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, meta_pid); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(mp, select_sql, meta_pid); err != nil {
		return nil, err
	}
	return mp, nil
}

func Query_All_OJs(
	tx *sqlx.Tx,
	required []string,
	excepts []string) ([]OJInfo, error) {

	oj := &OJInfo{}
	ojs := []OJInfo{}
	str_fields, err := GenerateSelectSQL(oj, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := JoinSQL("SELECT", str_fields, "FROM OJInfo")
	if err := tx.Select(&ojs, sql); err != nil {
		return nil, err
	}
	return ojs, nil
}

func Query_OJ_By_OJName(
	tx *sqlx.Tx,
	oj_name string,
	required []string,
	excepts []string,
) (*OJInfo, error) {

	from_where_sql := "FROM OJInfo WHERE oj_name=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	oj := &OJInfo{}
	str_fields, err := GenerateSelectSQL(oj, required, excepts)
	if err != nil {
		return nil, err
	}
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, oj_name); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	if err := tx.Get(oj, select_sql, oj_name); err != nil {
		return nil, err
	}
	return oj, nil
}

func Query_ProblemNum_By_OJIdFK(
	tx *sqlx.Tx, id int64) (int32, error) {

	var count int32
	if err := tx.Get(&count, "SELECT COUNT(*) FROM MetaProblems WHERE oj_id_fk=?",
		id); err != nil {

		return 0, err
	}
	return count, nil
}

// Need to be tested
func Query_All_OJNames(tx *sqlx.Tx) ([]string, error) {
	ojs := []string{}
	sql := `SELECT oj_name FROM OJInfo`
	if err := tx.Select(&ojs, sql); err != nil {
		return nil, err
	}
	return ojs, nil
}

//
type LanguageExt struct {
	OJName string `db:"oj_name"`
	Language
}

func Query_All_Languages(
	tx *sqlx.Tx,
	required []string,
	excepts []string) ([]LanguageExt, error) {

	/*-- Func start --*/
	lang := LanguageExt{}
	langs := []LanguageExt{}
	str_fields, err := GenerateSelectSQL(&lang, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := JoinSQL(
		"SELECT", str_fields,
		"FROM Languages LEFT JOIN OJInfo ON oj_id_fk=oj_id")
	if err := tx.Select(&langs, sql); err != nil {
		return nil, err
	}
	return langs, nil

}

func Query_MetaProblem_By_OJName_OJPid(
	tx *sqlx.Tx,
	oj_name string,
	oj_pid string,
	required []string,
	excepts []string) (*MetaProblem, error) {

	/*-- Func start --*/
	mp := &MetaProblem{}
	str_fields, err := GenerateSelectSQL(mp, required, excepts)
	if err != nil {
		return nil, err
	}

	from_where_sql := "FROM MetaProblems WHERE oj_pid=? AND oj_name=?"
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, oj_pid, oj_name); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(mp, select_sql, oj_pid, oj_name); err != nil {
		return nil, err
	}
	return mp, nil
}

func Query_If_User_Exists(
	tx *sqlx.Tx,
	name string,
) (bool, error) {

	/*-- Func start --*/
	var count int
	if err := tx.Get(&count, "SELECT COUNT(*) FROM Users WHERE username=?", name); err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}

func Query_User_By_Username(
	tx *sqlx.Tx,
	name string,
	required []string,
	excepts []string) (*User, error) {

	/*-- Func start --*/
	user := &User{}
	str_fields, err := GenerateSelectSQL(user, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM Users WHERE username=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, name); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(user, select_sql, name); err != nil {
		return nil, err
	}
	return user, nil
}

func Query_Language_By_LangId(
	tx *sqlx.Tx,
	lang_id int64,
	required []string,
	excepts []string) (*Language, error) {

	/*-- Func start --*/
	lang := &Language{}
	str_fields, err := GenerateSelectSQL(lang, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM Languages WHERE lang_id=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, lang_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(lang, select_sql, lang_id); err != nil {
		return nil, err
	}
	return lang, nil
}

func Query_Languages_By_OJIdFK(
	tx *sqlx.Tx,
	oj_id_fk int64,
	required []string,
	excepts []string) ([]LanguageExt, error) {

	/*-- Func start --*/
	lang := LanguageExt{}
	langs := []LanguageExt{}
	str_fields, err := GenerateSelectSQL(&lang, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := JoinSQL(
		"FROM Languages",
		"LEFT JOIN OJInfo ON oj_id_fk=oj_id",
		"WHERE oj_id_fk=?")
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, oj_id_fk); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Select(&langs, select_sql, oj_id_fk); err != nil {
		return nil, err
	}
	return langs, nil
}

func Query_Submission_By_RunId(
	tx *sqlx.Tx,
	run_id int64,
	required []string,
	excepts []string) (*Submission, error) {

	/*-- Func start --*/
	sub := &Submission{}
	str_fields, err := GenerateSelectSQL(sub, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM Submissions WHERE run_id=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, run_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(sub, select_sql, run_id); err != nil {
		return nil, err
	}
	return sub, nil
}

func Query_TestCases_By_MetaPid(
	tx *sqlx.Tx,
	meta_pid int64,
	required []string,
	excepts []string) ([]TestCase, error) {

	/*-- Func start --*/
	tc := TestCase{}
	tcs := []TestCase{}
	str_fields, err := GenerateSelectSQL(&tc, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM TestCases WHERE meta_pid_fk=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, meta_pid); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Select(&tcs, select_sql, meta_pid); err != nil {
		return nil, err
	}
	return tcs, nil
}

func Query_Limits_By_MetaPid(
	tx *sqlx.Tx,
	meta_pid int64,
	required []string,
	excepts []string) ([]TimeMemoryLimit, error) {

	/* Func start */
	tm := TimeMemoryLimit{}
	tms := []TimeMemoryLimit{}
	str_fields, err := GenerateSelectSQL(&tm, required, excepts)
	if err != nil {
		return nil, err
	}

	from_where_sql := JoinSQL("FROM", TimeMemoryLimit_TableName, "WHERE meta_pid_fk=?")
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, meta_pid); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Select(&tms, select_sql, meta_pid); err != nil {
		return nil, err
	}
	return tms, nil
}

func Query_Total_Submissions_By_MetaPid(
	tx *sqlx.Tx,
	meta_pid int64) (int, error) {

	var count int
	if err := tx.Get(&count, "SELECT COUNT(*) FROM Submissions WHERE meta_pid_fk=?", meta_pid); err != nil {
		return 0, err
	}
	return count, nil
}

func Query_Total_AC_Submissions_By_MetaPid(
	tx *sqlx.Tx,
	meta_pid int64) (int, error) {

	var count int
	if err := tx.Get(&count, "SELECT COUNT(*) FROM Submissions WHERE meta_pid_fk=? AND status_code='ac'", meta_pid); err != nil {
		return 0, err
	}
	return count, nil

}

func Query_Contest_Total_Submissions_By_ContestId_Label(
	tx *sqlx.Tx,
	contest_id int64,
	label string,
	start_time time.Time,
	end_time time.Time,
) (int, error) {
	var cnt int
	sql := JoinSQL("SELECT COUNT(*)",
		"FROM Submissions WHERE is_contest=true",
		"AND submit_time BETWEEN ? AND ?",
		"AND cp_id_fk=(SELECT cp_id FROM ContestProblems WHERE contest_id_fk=? AND label=? LIMIT 1)")
	if err := tx.Get(&cnt, sql, start_time, end_time, contest_id, label); err != nil {
		return 0, err
	}
	return cnt, nil
}

func Query_Contest_AC_Submissions_By_ContestId_Label(
	tx *sqlx.Tx,
	contest_id int64,
	label string,
	start_time time.Time,
	end_time time.Time,
) (int, error) {
	var cnt int
	sql := JoinSQL("SELECT COUNT(*)",
		"FROM Submissions WHERE is_contest=true",
		"AND status_code='ac'",
		"AND submit_time BETWEEN ? AND ?",
		"AND cp_id_fk=(SELECT cp_id FROM ContestProblems WHERE contest_id_fk=? AND label=? LIMIT 1)")
	if err := tx.Get(&cnt, sql, start_time, end_time, contest_id, label); err != nil {
		return 0, err
	}
	return cnt, nil
}

func Query_Contest_By_ContestId(
	tx *sqlx.Tx,
	contest_id int64,
	required []string,
	excepts []string) (*Contest, error) {

	cst := &Contest{}
	str_fields, err := GenerateSelectSQL(cst, required, excepts)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM Contests WHERE contest_id=?"
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, contest_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(cst, select_sql, contest_id); err != nil {
		return nil, err
	}

	return cst, nil
}

func Query_ContestUser_By_ContestId_And_UserId(
	tx *sqlx.Tx,
	contest_id int64,
	user_id int64,
) (*ContestUser, error) {
	cu := &ContestUser{}
	str_fields, err := GenerateSelectSQL(cu, nil, nil)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM ContestUsers WHERE contest_id_fk=? AND user_id_fk=?"
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, contest_id, user_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(cu, select_sql, contest_id, user_id); err != nil {
		return nil, err
	}
	return cu, nil
}

func Query_ContestProblem_By_ContestId_And_Label(
	tx *sqlx.Tx,
	contest_id int64,
	label string,
) (*ContestProblem, error) {
	cp := &ContestProblem{}
	str_fields, err := GenerateSelectSQL(cp, nil, nil)
	if err != nil {
		return nil, err
	}
	from_where_sql := "FROM ContestProblems WHERE contest_id_fk=? AND label=?"
	select_sql := JoinSQL("SELECT", str_fields, from_where_sql)
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, contest_id, label); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}

	if err := tx.Get(cp, select_sql, contest_id, label); err != nil {
		return nil, err
	}
	return cp, nil
}

func Query_ContestProblemLabels_By_ContestId(
	tx *sqlx.Tx,
	contest_id int64,
) ([]string, error) {
	from_where_sql := JoinSQL("FROM ContestProblems",
		"WHERE contest_id_fk=?",
		"ORDER BY label")
	count_sql := JoinSQL("SELECT COUNT(*)", from_where_sql)
	select_sql := JoinSQL("SELECT label", from_where_sql)

	var cnt int
	if err := tx.Get(&cnt, count_sql, contest_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	var labels []string
	if err := tx.Select(&labels, select_sql, contest_id); err != nil {
		return nil, err
	}
	return labels, nil
}

func Query_ContestUsers_By_ContestId(
	tx *sqlx.Tx,
	contest_id int64,
) ([]string, error) {
	from_where_sql := JoinSQL(`FROM ContestUsers`,
		`LEFT JOIN Users ON user_id_fk=user_id`,
		`WHERE contest_id_fk=?`)
	count_sql := JoinSQL(`SELECT COUNT(*)`, from_where_sql)
	select_sql := JoinSQL(`SELECT username`, from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, contest_id); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	var users []string
	if err := tx.Select(&users, select_sql, contest_id); err != nil {
		return nil, err
	}
	return users, nil
}

func Query_Limit_By_Language_And_MetaPid(
	tx *sqlx.Tx,
	lang string,
	meta_pid int64,
) (*TimeMemoryLimit, error) {
	tml := &TimeMemoryLimit{}
	str_fields, err := GenerateSelectSQL(tml, nil, nil)
	if err != nil {
		return nil, err
	}
	from_where_sql := JoinSQL(`FROM TimeMemoryLimits`,
		`WHERE meta_pid_fk=? AND language=?`)
	count_sql := JoinSQL(`SELECT COUNT(*)`, from_where_sql)
	select_sql := JoinSQL(`SELECT`, str_fields, from_where_sql)
	var cnt int
	if err := tx.Get(&cnt, count_sql, meta_pid, lang); err != nil {
		return nil, err
	}
	if cnt == 0 {
		return nil, nil
	}
	if err := tx.Get(tml, select_sql, meta_pid, lang); err != nil {
		return nil, err
	}
	return tml, nil
}
