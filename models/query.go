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

// Some more complex methods to query data

/* Problem pages
 	@params
 		filter_status: 0 all, 1 accepted, 2 unsolved, 3 attempted
		orderby_element: 0 pid, 1 title, 2 ac_rate,
	TODO: ac_rate
*/
type Pagination struct {
	TotalLines  int
	TotalPages  int
	CurrentPage int
	Lines       []MetaProblem
}

func XQuery_List_Problems_With_Filter(
	tx *sqlx.Tx,
	username string,
	show_hidden bool,
	filter_oj string,
	filter_status int,
	orderby_element int,
	is_desc bool,
	offset int,
	per_page int,
	current_page int,
	required []string,
	excepts []string) (*Pagination, error) {

	/*-- Func start --*/
	ret := &Pagination{}
	mp := MetaProblem{}
	mps := []MetaProblem{}
	str_fields, err := GenerateSelectSQL(&mp, required, excepts)
	if err != nil {
		return nil, err
	}

	// Build from_sql
	var from_sql string
	switch filter_status {
	case 0:
		from_sql = " MetaProblems "
	case 1:
		// Accepted
		from_sql = fmt.Sprintf(
			` (SELECT %s FROM MetaProblems 
			   WHERE meta_pid IN 
			     (SELECT meta_pid_fk FROM Submissions
				  WHERE status_code="ac" AND
			  		user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s"))) AS ACCEPTED `, str_fields, username)
	case 2:
		// Unsolved
		from_sql = fmt.Sprintf(
			` (SELECT %s FROM MetaProblems 
			   WHERE meta_pid NOT IN 
			     (SELECT meta_pid_fk FROM Submissions
				  WHERE status_code="ac" AND
			  		user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s"))) AS UNSOLVED `, str_fields, username)
	case 3:
		// Attempted
		from_sql = fmt.Sprintf(
			` (SELECT %s FROM MetaProblems 
			   WHERE meta_pid NOT IN 
			     (SELECT meta_pid_fk FROM Submissions
				  WHERE status_code="ac" AND
			  		user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s"))
			   AND meta_pid IN
			     (SELECT meta_pid_fk FROM Submissions
			 	  WHERE user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s")
				  GROUP BY meta_pid_fk
				  HAVING COUNT(run_id) > 0)) AS ATTEMPTED `, str_fields, username, username)
	}

	// Build where_sql
	var where_sql string
	if show_hidden {
		where_sql = `
		WHERE oj_name=?
		`
	} else {
		where_sql = `
		WHERE oj_name=? AND hide=0
		`
	}

	// Get count of lines
	count_sql := `
	SELECT COUNT(*) FROM 
	` + from_sql + where_sql
	var count int
	if err := tx.Get(&count, count_sql, filter_oj); err != nil {
		return nil, err
	}
	ret.TotalLines = count
	if per_page == 0 {
		ret.TotalPages = 1
		per_page = ret.TotalLines
	} else {
		ret.TotalPages = ret.TotalLines / per_page
		if ret.TotalLines%per_page != 0 {
			ret.TotalPages += 1
		}
		if ret.TotalPages == 0 {
			ret.TotalPages = 1
		}
	}
	if current_page == 0 {
		current_page = 1
	}
	if current_page > ret.TotalPages {
		current_page = ret.TotalPages
	}
	ret.CurrentPage = current_page

	// Get lines
	sql :=
		"SELECT %s FROM " +
			from_sql + where_sql + " ORDER BY %s LIMIT %d, %d"
	var orderby string
	switch orderby_element {
	case 0:
		orderby = "oj_pid "
	case 1:
		orderby = "title "
		// TODO: Case 2, ac_rate
	}
	if is_desc == true {
		orderby += "DESC"
	}
	offset = (ret.CurrentPage-1)*per_page + offset
	full_sql := fmt.Sprintf(sql, str_fields, orderby, offset, per_page)
	if err := tx.Select(&mps, full_sql, filter_oj); err != nil {
		return nil, err
	}
	ret.Lines = mps
	return ret, nil
}
