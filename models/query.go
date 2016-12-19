package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
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
	sql := `
	SELECT %s FROM MetaProblems
	WHERE meta_pid=?
	`
	if err := tx.Get(mp, fmt.Sprintf(sql, str_fields), meta_pid); err != nil {
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
	sql := `
	SELECT %s FROM OJInfo
	`
	if err := tx.Select(&ojs, fmt.Sprintf(sql, str_fields)); err != nil {
		return nil, err
	}
	return ojs, nil
}

// Need to be tested
func Query_All_OJNames(tx *sqlx.Tx) ([]string, error) {
	ojs := []string{}
	sql := `
	SELECT oj_name FROM OJInfo
	`
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
	sql := `
	SELECT %s FROM Languages LEFT JOIN OJInfo ON oj_id_fk=oj_id
	`
	if err := tx.Select(&langs, fmt.Sprintf(sql, str_fields)); err != nil {
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
	sql := `
	SELECT %s FROM MetaProblems
	WHERE oj_pid=? AND oj_name=?
	`
	if err := tx.Get(mp, fmt.Sprintf(sql, str_fields), oj_pid, oj_name); err != nil {
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
	sql := `
	SELECT %s FROM Users
	WHERE username=?
	`
	if err := tx.Get(user, fmt.Sprintf(sql, str_fields), name); err != nil {
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
	sql := `
	SELECT %s FROM Languages
	WHERE lang_id=?
	`
	if err := tx.Get(lang, fmt.Sprintf(sql, str_fields), lang_id); err != nil {
		return nil, err
	}
	return lang, nil
}

func Query_Languages_By_OJIdFK(
	tx *sqlx.Tx,
	oj_id_fk int64,
	required []string,
	excepts []string) ([]Language, error) {

	/*-- Func start --*/
	lang := Language{}
	langs := []Language{}
	str_fields, err := GenerateSelectSQL(&lang, required, excepts)
	if err != nil {
		return nil, err
	}
	sql := `
	SELECT %s FROM Languages 
	WHERE oj_id_fk=?
	`
	if err := tx.Select(&langs, fmt.Sprintf(sql, str_fields), oj_id_fk); err != nil {
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
	sql := `
	SELECT %s FROM Submissions
	WHERE run_id=?
	`
	if err := tx.Get(sub, fmt.Sprintf(sql, str_fields), run_id); err != nil {
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
	sql := `
	SELECT %s FROM TestCases
	WHERE meta_pid_fk=?
	`
	if err := tx.Select(&tcs, fmt.Sprintf(sql, str_fields), meta_pid); err != nil {
		return nil, err
	}
	return tcs, nil
}
