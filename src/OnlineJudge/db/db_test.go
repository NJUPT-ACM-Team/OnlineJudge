package db

import "testing"

func TestUseConfig(t *testing.T) {
	mysql_config := MySQLConfig{
		Drivername: "mysql",
		Username:   "kevince",
		DBname:     "testdb",
	}
	var cfg Config
	cfg = &mysql_config
	UseConfig(cfg)
	get_cfg := GetConfig()
	t.Log("Testing UseConifg")
	if dn := get_cfg.GetDriverName(); dn != "mysql" {
		t.Errorf("Failed to pass drivername.")
	}
	if dsn, _ := get_cfg.GetDataSourceName(); dsn != "kevince@/testdb" {
		t.Errorf("Failed to pass data source name parameters.")
	}
}

func TestNewDB(t *testing.T) {
	t.Log("Init database, and trying to connect to db.")
	// Init()

	if _, err := NewDB(); err != nil {
		t.Errorf("Failed to connect to db, because %s", err)
	}
}
