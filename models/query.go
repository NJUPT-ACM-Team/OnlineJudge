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
func Query_All_Languages(
	tx *sqlx.Tx,
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

// Some more complex methods to query data

/* Problem pages
 	@params
 		filter_status: 0 all, 1 accepted, 2 unsolved, 3 attempted
		orderby_element: 0 pid, 1 title, 2 ac_rate,
	TODO: ac_rate
*/
type ListProblemsPagination struct {
	TotalLines  int
	TotalPages  int
	CurrentPage int
	Problems    []MetaProblem
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
	excepts []string) (*ListProblemsPagination, error) {

	/*-- Func start --*/
	ret := &ListProblemsPagination{}
	mp := MetaProblem{}
	mps := []MetaProblem{}
	str_fields, err := GenerateSelectSQL(&mp, required, excepts)
	if err != nil {
		return nil, err
	}

	if filter_oj == "" {
		filter_oj = "%"
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
		WHERE oj_name like ?
		`
	} else {
		where_sql = `
		WHERE oj_name like ? AND hide=0
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
	if per_page <= 0 {
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
	ret.Problems = mps
	return ret, nil
}

type SubmissionExt struct {
	Username          string
	NumberOfTestcases int `db:"number_of_testcases"`
	Submission
	Language
}

type ListSubmissionsPagination struct {
	TotalLines  int
	TotalPages  int
	CurrentPage int
	Submissions []SubmissionExt
}

/*
	Submission pages
	@params:
		show_private: if true, show all submissions and code both shared or not
					  if false, show only public submissions and shared code.
*/

func XQuery_List_Submissions_With_Filter(
	tx *sqlx.Tx,
	username string,
	show_private bool,
	filter_oj string,
	filter_pid string,
	filter_status_code string,
	filter_language string,
	filter_compiler string,
	per_page int,
	current_page int,
	required []string,
	excepts []string) (*ListSubmissionsPagination, error) {

	/*-- Func start --*/
	need_filter := false
	if username == "" {
		username = "%"
	} else {
		need_filter = true
	}
	if filter_oj == "" {
		filter_oj = "%"
	} else {
		need_filter = true
	}
	if filter_pid == "" {
		filter_pid = "%"
	} else {
		need_filter = true
	}
	if filter_status_code == "" {
		filter_status_code = "%"
	} else {
		need_filter = true
	}
	if filter_language == "" {
		filter_language = "%"
	} else {
		need_filter = true
	}
	if filter_compiler == "" {
		filter_compiler = "%"
	} else {
		need_filter = true
	}

	where_sql := ` WHERE 1 `

	if show_private == false {
		where_sql += `AND is_private = 0 `
	}

	if need_filter {
		where_sql += `
		AND
		user_id_fk IN (SELECT user_id FROM Users WHERE username LIKE ?) AND
		meta_pid_fk IN 
			(SELECT meta_pid FROM MetaProblems 
			 WHERE oj_name LIKE ? AND oj_pid LIKE ?) AND
		status_code like ? AND
		lang_id_fk IN
			(SELECT lang_id FROM Languages
			 WHERE language like ? AND compiler like ?)`
	}

	ret := &ListSubmissionsPagination{}

	// Get count
	count_sql := "SELECT COUNT(*) FROM Submissions " + where_sql
	var count int
	if need_filter {
		if err := tx.Get(&count, count_sql,
			username, filter_oj, filter_pid,
			filter_status_code, filter_language, filter_compiler); err != nil {

			return nil, err
		}
	} else {
		if err := tx.Get(&count, count_sql); err != nil {
			return nil, err
		}
	}
	ret.TotalLines = count

	if per_page <= 0 {
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
	sub := &SubmissionExt{}
	subs := []SubmissionExt{}
	str_fields, err := GenerateSelectSQL(sub, required, []string{"submission", "language"})
	if err != nil {
		return nil, err
	}

	offset := (current_page - 1) * per_page
	sql := `SELECT %s FROM Submissions 
		LEFT JOIN Users ON user_id_fk=user_id 
		LEFT JOIN Languages ON lang_id_fk=lang_id
		LEFT JOIN (SELECT meta_pid, number_of_testcases FROM MetaProblems) AS TP ON meta_pid_fk=meta_pid ` + where_sql + ` LIMIT %d, %d`

	real_sql := fmt.Sprintf(sql, str_fields, offset, per_page)

	if need_filter {
		if err := tx.Select(
			&subs, real_sql,
			username, filter_oj, filter_pid,
			filter_status_code, filter_language, filter_compiler); err != nil {

			return nil, err
		}
	} else {
		if err := tx.Select(&subs, real_sql); err != nil {
			return nil, err
		}
	}
	ret.Submissions = subs
	return ret, nil
}
