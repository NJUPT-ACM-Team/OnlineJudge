DROP DATABASE IF EXISTS ojtest;
CREATE DATABASE ojtest DEFAULT CHARSET=UTF8;
USE ojtest;

-- Users, information for all users.

CREATE TABLE Users (
    user_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    username VARCHAR(64) NOT NULL,
    password VARCHAR(64) NOT NULL,

    email VARCHAR(64) NOT NULL,
    phone VARCHAR(32) NOT NULL DEFAULT '',
    school VARCHAR(64) NOT NULL DEFAULT '',

    motto VARCHAR(1024),

    total_local_submit INTEGER NOT NULL DEFAULT 0,
    total_local_ac INTEGER NOT NULL DEFAULT 0,
    total_submit INTEGER NOT NULL DEFAULT 0,
    total_ac INTEGER NOT NULL DEFAULT 0,

    register_time DATETIME NOT NULL,
    last_login_time DATETIME NOT NULL,
    ip_addr VARCHAR(255) NOT NULL,
    permission VARCHAR(64) NOT NULL DEFAULT '' COMMENT 'ENUM(root)',
    lock_status BOOLEAN NOT NULL DEFAULT 0,

    PRIMARY KEY (user_id),
    UNIQUE KEY (username),
    UNIQUE KEY (email)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE OJInfo (
    oj_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    version VARCHAR(64) NOT NULL,
    int64io VARCHAR(255) NOT NULL,
    javaclass VARCHAR(255) NOT NULL,
    status VARCHAR(16) NOT NULL COMMENT '(ok, down, unstable) NOT NULL',
    status_info VARCHAR(1024) NOT NULL DEFAULT '',
    lastcheck DATETIME NOT NULL,

    PRIMARY KEY (oj_id),
    UNIQUE KEY (name)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

-- References to OJInfo
CREATE TABLE Languages (
    lang_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    language VARCHAR(64) NOT NULL COMMENT 'like c, c++, java, python etc.',
    option_value_int INTEGER COMMENT 'for submit options int',
    option_value_str VARCHAR(64) COMMENT 'for submit options str',
    compiler VARCHAR(255) NOT NULL,
    oj_id_fk INTEGER UNSIGNED NOT NULL,

    PRIMARY KEY (lang_id),
    FOREIGN KEY (oj_id_fk) REFERENCES OJInfo(oj_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE MetaProblems (
    meta_pid INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description LONGTEXT NOT NULL,
    input TEXT NOT NULL,
    output TEXT NOT NULL,
    sample_in TEXT NOT NULL,
    sample_out TEXT NOT NULL,
    time_limit INTEGER UNSIGNED NOT NULL,
    case_time_limit INTEGER UNSIGNED NOT NULL,
    memory_limit INTEGER UNSIGNED NOT NULL,
    number_of_testcases INTEGER UNSIGNED NOT NULL,
    source VARCHAR(1024) NOT NULL,
    hint TEXT NOT NULL,
    hide BOOLEAN NOT NULL DEFAULT 1 COMMENT 'Hide the problem or not, for contest purpose',
    oj_id_fk INTEGER UNSIGNED NOT NULL,
    oj_pid INTEGER UNSIGNED NOT NULL,
    
    PRIMARY KEY (meta_pid),
    FOREIGN KEY (oj_id_fk) REFERENCES OJInfo(oj_id),
    UNIQUE KEY (oj_id_fk, oj_pid)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE LocalProblems (
    local_pid INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    meta_pid_fk INTEGER UNSIGNED NOT NULL,

    PRIMARY KEY (local_pid),
    FOREIGN KEY (meta_pid_fk) REFERENCES MetaProblems(meta_pid)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE Contests (
    contest_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    is_virtual BOOLEAN NOT NULL,
    contest_type ENUM('icpc', 'oi', 'cf') NOT NULL,
    start_time DATETIME NOT NULL,
    end_time DATETIME NOT NULL,
    lock_board_time DATETIME NOT NULL,
    hide_others_status BOOLEAN NOT NULL,

    PRIMARY KEY (contest_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE ContestsUsers (
    cu_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    user_id_fk INTEGER UNSIGNED NOT NULL,
    contest_id_fk INTEGER UNSIGNED NOT NULL,
    motto VARCHAR(1024) NOT NULL DEFAULT '',

    PRIMARY KEY (cu_id),
    FOREIGN KEY (user_id_fk) REFERENCES Users(user_id),
    FOREIGN KEY (contest_id_fk) REFERENCES Contests(contest_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE ContestsProblems (
    cp_id INTEGER UNSIGNED NOT NULL AUTO_INCREMENT,
    meta_pid_fk INTEGER UNSIGNED NOT NULL, 
    contest_id_fk INTEGER UNSIGNED NOT NULL,
    label VARCHAR(64) NOT NULL,
    problem_type ENUM('icpc', 'oi', 'cf', 'cfd') NOT NULL,
    base INTEGER NOT NULL,
    minp INTEGER NOT NULL,
    para_a DOUBLE NOT NULL,
    para_b DOUBLE NOT NULL,

    PRIMARY KEY (cp_id),
    UNIQUE KEY (contest_id_fk, meta_pid_fk),
    UNIQUE KEY (contest_id_fk, label),
    FOREIGN KEY (meta_pid_fk) REFERENCES MetaProblems(meta_pid),
    FOREIGN KEY (contest_id_fk) REFERENCES Contests(contest_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

CREATE TABLE Submissions (
    run_id INTEGER NOT NULL AUTO_INCREMENT,
    status VARCHAR(64) NOT NULL,
    status_code VARCHAR(8) COMMENT 'ENUM(se, wt, ce, tle, mle, ole, re, wa, pe, ac)',
    testcases_passed INTEGER NOT NULL DEFAULT 0,
    submit_time DATETIME NOT NULL,
    time_used INTEGER NOT NULL,
    memory_used INTEGER NOT NULL,
    code TEXT NOT NULL,
    lang_id_fk INTEGER UNSIGNED NOT NULL DEFAULT 0,
    ce_info TEXT NOT NULL,
    ip_addr VARCHAR(255) NOT NULL DEFAULT '',
    is_shared BOOLEAN NOT NULL,

    is_contest BOOLEAN NOT NULL,
    cp_id_fk INTEGER UNSIGNED NOT NULL DEFAULT 0,
    cu_id_fk INTEGER UNSIGNED NOT NULL DEFAULT 0,
    meta_pid_fk INTEGER UNSIGNED NOT NULL DEFAULT 0,
    user_id_fk INTEGER UNSIGNED NOT NULL DEFAULT 0,

    PRIMARY KEY (run_id),
    FOREIGN KEY (lang_id_fk) REFERENCES Languages(lang_id),
    FOREIGN KEY (cp_id_fk) REFERENCES ContestsProblems(cp_id),
    FOREIGN KEY (cu_id_fk) REFERENCES ContestsUsers(cu_id),
    FOREIGN KEY (meta_pid_fk) REFERENCES MetaProblems(meta_pid),
    FOREIGN KEY (user_id_fk) REFERENCES Users(user_id)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

