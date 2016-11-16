package db

import "testing"

func TestGetDriverName(t *testing.T) {
	t.Log("Get driver name.")
	var mysql_config MySQLConfig

	t.Log("Use \"hello\" (expected \"hello\")")
	mysql_config = MySQLConfig{drivername: "hello"}
	if dn := mysql_config.GetDriverName(); dn != "hello" {
		t.Errorf(`Expected "hello", but it was "%s" instead.`, dn)
	}

	t.Log("Use empty string (expected \"mysql\")")
	mysql_config = MySQLConfig{}
	if dn := mysql_config.GetDriverName(); dn != "mysql" {
		t.Errorf(`Expected "mysql", but it was "%s" instead.`, dn)
	}
}

func TestGetDataSourceName(t *testing.T) {
	t.Log("Get data source name.")
	var mysql_config MySQLConfig

	t.Log("Empty username (expected error)")
	mysql_config = MySQLConfig{dbname: "testdb"}
	if _, err := mysql_config.GetDataSourceName(); err == nil {
		t.Errorf("Expected errors, but there is no error.")
	}

	t.Log("Empty dbname (expected error)")
	mysql_config = MySQLConfig{username: "user"}
	if _, err := mysql_config.GetDataSourceName(); err == nil {
		t.Errorf("Expected errors, but there is no error.")
	}

	t.Log(`With password "123456", (expected "user:123456@/testdb")`)
	mysql_config = MySQLConfig{
		username: "user",
		password: "123456",
		dbname:   "testdb",
	}
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@/testdb" {
		t.Errorf(`expected "user:123456@/testdb", but it was "%s" instead`, dsn)
	}

	t.Log(`With protocol "tcp", (expected "user:123456@tcp/testdb")`)
	mysql_config = MySQLConfig{
		username: "user",
		password: "123456",
		dbname:   "testdb",
		protocol: "tcp",
	}
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@tcp/testdb" {
		t.Errorf(`expected "user:123456@tcp/testdb", but it was "%s" instead`, dsn)
	}

	t.Log(`With address "10.10.23.12", (expected "user:123456@tcp(10.10.23.12)/testdb")`)
	mysql_config = MySQLConfig{
		username: "user",
		password: "123456",
		dbname:   "testdb",
		protocol: "tcp",
		address:  "10.10.23.12",
	}
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@tcp(10.10.23.12)/testdb" {
		t.Errorf(`expected "user:123456@tcp(10.10.23.12)/testdb", but it was "%s" instead`, dsn)
	}

	t.Log(`With only address "10.10.23.12", without protocol, (expected "user:123456@/testdb")`)
	mysql_config = MySQLConfig{
		username: "user",
		password: "123456",
		dbname:   "testdb",
		address:  "10.10.23.12",
	}
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@/testdb" {
		t.Errorf(`expected "user:123456@/testdb", but it was "%s" instead`, dsn)
	}

	t.Log(`With one parameter charset=utf8, (expected "user:123456@tcp(10.10.23.12)/testdb?charset=utf8")`)
	params := make(map[string]string)
	params["charset"] = "utf8"
	mysql_config = MySQLConfig{
		username: "user",
		password: "123456",
		dbname:   "testdb",
		protocol: "tcp",
		address:  "10.10.23.12",
		params:   params,
	}
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@tcp(10.10.23.12)/testdb?charset=utf8" {
		t.Errorf(`expected "user:123456@tcp(10.10.23.12)/testdb?charset=utf8", but it was "%s" instead`, dsn)
	}
	t.Log(`With two parameters charset=utf8, engine=innodb\
	(expected "user:123456@tcp(10.10.23.12)/testdb?charset=utf8&engine=innodb")`)
	params["engine"] = "innodb"
	if dsn, _ := mysql_config.GetDataSourceName(); dsn != "user:123456@tcp(10.10.23.12)/testdb?charset=utf8&engine=innodb" &&
		dsn != "user:123456@tcp(10.10.23.12)/testdb?engine=innodb&charset=utf8" {
		t.Errorf(`expected "user:123456@tcp(10.10.23.12)/testdb?charset=utf8&engine=innodb", but it was "%s" instead`, dsn)
	}
}
