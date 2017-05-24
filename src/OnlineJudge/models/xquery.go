package models

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

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
	filter_status string,
	orderby_element string,
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
	case "ALL":
		from_sql = "MetaProblems"
	case "ACCEPTED":
		// Accepted
		from_sql = fmt.Sprintf(
			`(SELECT %s FROM MetaProblems 
			   WHERE meta_pid IN 
			     (SELECT meta_pid_fk FROM Submissions
				  WHERE status_code="ac" AND
			  		user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s"))) AS ACCEPTED`, str_fields, username)
	case "UNSOLVED":
		// Unsolved
		from_sql = fmt.Sprintf(
			`(SELECT %s FROM MetaProblems 
			   WHERE meta_pid NOT IN 
			     (SELECT meta_pid_fk FROM Submissions
				  WHERE status_code="ac" AND
			  		user_id_fk=(
						SELECT user_id FROM Users
						WHERE username="%s"))) AS UNSOLVED`, str_fields, username)
	case "ATTEMPTED":
		// Attempted
		from_sql = fmt.Sprintf(
			`(SELECT %s FROM MetaProblems 
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
				  HAVING COUNT(run_id) > 0)) AS ATTEMPTED`, str_fields, username, username)
	}

	// Build where_sql
	var where_sql string
	if show_hidden {
		where_sql = `WHERE oj_name like ?`
	} else {
		where_sql = `WHERE oj_name like ? AND hide=0`
	}

	// Get count of lines
	count_sql := JoinSQL("SELECT COUNT(*) FROM", from_sql, where_sql)
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
	var orderby string
	switch orderby_element {
	case "PID":
		orderby = "oj_pid"
	case "TITLE":
		orderby = "title"
		// TODO: Case 2, ac_rate
	}
	if is_desc == true {
		orderby = JoinSQL(orderby, "DESC")
	}
	offset = (ret.CurrentPage-1)*per_page + offset
	// full_sql := fmt.Sprintf(sql, orderby, offset, per_page)
	sql := JoinSQL(
		"SELECT", str_fields, "FROM", from_sql, where_sql,
		"ORDER BY", orderby, fmt.Sprintf("LIMIT %d, %d", offset, per_page))
	if err := tx.Select(&mps, sql, filter_oj); err != nil {
		return nil, err
	}
	ret.Problems = mps
	return ret, nil
}

type SubmissionExt struct {
	Username          string
	NumberOfTestcases int    `db:"number_of_testcases"`
	OJName            string `db:"oj_name"`
	OJPid             string `db:"oj_pid"`

	Submission
	Language
}

type ListSubmissionsPagination struct {
	TotalLines  int
	TotalPages  int
	CurrentPage int
	Submissions []SubmissionExt
}

func initFilters(vars ...*string) bool {
	ret := false
	for _, v := range vars {
		if *v == "" {
			*v = "%"
		} else {
			ret = true
		}
	}
	return ret
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
	is_desc bool,
	filter_username string,
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
	need_filter := initFilters(
		&filter_username,
		&filter_oj,
		&filter_pid,
		&filter_status_code,
		&filter_language,
		&filter_compiler)

	where_sql := `WHERE is_contest=false`

	if show_private == false {
		where_sql = JoinSQL(where_sql, "AND is_private=0")
	}

	if need_filter {
		where_sql = JoinSQL(where_sql,
			`AND
		user_id_fk IN (SELECT user_id FROM Users WHERE username LIKE ?) AND
		meta_pid_fk IN 
			(SELECT meta_pid FROM MetaProblems 
			 WHERE oj_name LIKE ? AND oj_pid LIKE ?) AND
		status_code like ? AND
		lang_id_fk IN
			(SELECT lang_id FROM Languages
			 WHERE language like ? AND compiler like ?)`)
	}

	ret := &ListSubmissionsPagination{}

	// Get count
	count_sql := JoinSQL("SELECT COUNT(*) FROM Submissions", where_sql)
	var count int
	if need_filter {
		if err := tx.Get(&count, count_sql,
			filter_username, filter_oj, filter_pid,
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
	// str_fields, err := GenerateSelectSQL(sub, required, []string{"submission", "language"})
	str_fields, err := GenerateSelectSQL(sub, required, []string{"submission"})
	if err != nil {
		return nil, err
	}

	order_by := "ORDER BY run_id"
	if is_desc == false {
		order_by = JoinSQL(order_by, "DESC")
	}

	offset := (current_page - 1) * per_page
	sql := JoinSQL(
		`SELECT`, str_fields, `FROM Submissions`,
		`LEFT JOIN Users ON user_id_fk=user_id `,
		`LEFT JOIN Languages ON lang_id_fk=lang_id`,
		`LEFT JOIN (SELECT meta_pid, number_of_testcases, oj_name, oj_pid FROM MetaProblems) AS TP ON meta_pid_fk=meta_pid`,
		where_sql, order_by,
		fmt.Sprintf(`LIMIT %d, %d`, offset, per_page))

	if need_filter {
		if err := tx.Select(
			&subs, sql,
			filter_username, filter_oj, filter_pid,
			filter_status_code, filter_language, filter_compiler); err != nil {

			return nil, err
		}
	} else {
		if err := tx.Select(&subs, sql); err != nil {
			return nil, err
		}
	}
	ret.Submissions = subs
	return ret, nil
}

type ListContestsPagination struct {
	TotalLines  int
	TotalPages  int
	CurrentPage int
	Contests    []Contest
}

func XQuery_List_Contests_With_Filter(
	tx *sqlx.Tx,
	per_page int,
	current_page int,
	order_by_element string,
	is_desc bool,
	filter_ctype_element string,
	filter_is_public string,
	filter_is_virtual string,
	required []string,
	excepts []string) (*ListContestsPagination, error) {

	// initFilters(&filter_ctype_element)
	switch strings.ToUpper(filter_ctype_element) {
	case "ICPC":
		filter_ctype_element = "icpc"
	case "OI":
		filter_ctype_element = "oi"
	case "CF":
		filter_ctype_element = "cf"
	default:
		filter_ctype_element = "%"
	}

	switch strings.ToUpper(filter_is_public) {
	case "PUBLIC":
		filter_is_public = "1"
	case "PRIVATE":
		filter_is_public = "0"
	default:
		filter_is_public = "%"
	}

	switch strings.ToUpper(filter_is_virtual) {
	case "VIRTUAL":
		filter_is_virtual = "1"
	case "FORMAL":
		filter_is_virtual = "0"
	default:
		filter_is_virtual = "%"
	}
	// where sql
	where_sql := `WHERE contest_type LIKE ? AND is_virtual LIKE ?`
	if filter_is_public == "1" {
		where_sql = JoinSQL(where_sql, `AND password=""`)
	} else if filter_is_public == "0" {
		where_sql = JoinSQL(where_sql, `AND password!=""`)
	}
	// Get count
	count_sql := JoinSQL("SELECT COUNT(*) FROM Contests", where_sql)
	fmt.Println(filter_ctype_element, filter_is_virtual)
	var count int
	if err := tx.Get(&count, count_sql,
		filter_ctype_element, filter_is_virtual); err != nil {
		return nil, err
	}
	ret := &ListContestsPagination{}
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

	cst := &Contest{}
	csts := []Contest{}
	str_fields, err := GenerateSelectSQL(cst, required, excepts)
	if err != nil {
		return nil, err
	}

	var orderby string
	switch order_by_element {
	case "CID":
		orderby = "contest_id"
	case "TITLE":
		orderby = "title"
	case "STARTTIME":
		orderby = "start_time"
	case "ENDTIME":
		orderby = "end_time"
	case "STATUS":
		orderby = "status"
	default:
		orderby = "contest_id"
	}
	if is_desc {
		orderby = JoinSQL(orderby, "DESC")
	}
	order_by := JoinSQL("ORDER BY", orderby)
	offset := (current_page - 1) * per_page
	sql := JoinSQL(
		"SELECT", str_fields, "FROM Contests", where_sql, order_by,
		fmt.Sprintf(`LIMIT %d, %d`, offset, per_page))
	if err := tx.Select(
		&csts, sql, filter_ctype_element, filter_is_virtual); err != nil {

		return nil, err
	}
	ret.Contests = csts
	return ret, nil
}

type ProblemExt struct {
	MetaProblem
	Status string `db:"status"`
	Label  string `db:"label"`
	Alias  string `db:"alias"`
}

func XQuery_Contest_List_Problems(
	tx *sqlx.Tx,
	contest_id, user_id int64) ([]ProblemExt, error) {

	pext := &ProblemExt{}
	pexts := []ProblemExt{}
	str_fields, err := GenerateSelectSQL(pext, nil, []string{"status", "metaproblem"})
	if err != nil {
		return nil, err
	}
	sql := JoinSQL("SELECT", str_fields,
		"FROM ContestProblems cps",
		"LEFT JOIN MetaProblems mps ON cps.meta_pid_fk=mps.meta_pid",
		"WHERE contest_id_fk=?",
		"ORDER BY cps.label")
	// "ORDER BY cps.Label")
	if err := tx.Select(&pexts, sql, contest_id); err != nil {
		return nil, err
	}
	// TODO: add status check
	return pexts, nil
}
