package db

func Init() {
	params := make(map[string]string)
	params["parseTime"] = "true"
	mysql_config := MySQLConfig{
		drivername: "mysql",
		username:   "test",
		dbname:     "ojtest",
		password:   "abc123",
		params:     params,
	}
	var config Config
	config = &mysql_config
	UseConfig(config)
}
