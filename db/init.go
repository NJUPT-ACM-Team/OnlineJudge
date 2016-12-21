package db

func Init() {
	params := make(map[string]string)
	params["parseTime"] = "true"
	mysql_config := MySQLConfig{
		Drivername: "mysql",
		Username:   "test",
		DBname:     "ojtest",
		Password:   "abc123",
		Params:     params,
	}
	var config Config
	config = &mysql_config
	UseConfig(config)
}
